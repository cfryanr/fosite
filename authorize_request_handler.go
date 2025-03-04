/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 *
 */

package fosite

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gopkg.in/square/go-jose.v2"

	"github.com/ory/fosite/i18n"
	"github.com/ory/fosite/token/jwt"
	"github.com/ory/x/errorsx"

	"github.com/pkg/errors"

	"github.com/ory/go-convenience/stringslice"
)

func wrapSigningKeyFailure(outer *RFC6749Error, inner error) *RFC6749Error {
	outer = outer.WithWrap(inner).WithDebug(inner.Error())
	if e := new(RFC6749Error); errors.As(inner, &e) {
		return outer.WithHintf("%s %s", outer.Reason(), e.Reason())
	}
	return outer
}

func (f *Fosite) authorizeRequestParametersFromOpenIDConnectRequest(ctx context.Context, request *AuthorizeRequest) error {
	var scope Arguments = RemoveEmpty(strings.Split(request.Form.Get("scope"), " "))

	// Even if a scope parameter is present in the Request Object value, a scope parameter MUST always be passed using
	// the OAuth 2.0 request syntax containing the openid scope value to indicate to the underlying OAuth 2.0 logic that this is an OpenID Connect request.
	// Source: http://openid.net/specs/openid-connect-core-1_0.html#CodeFlowAuth
	if !scope.Has("openid") {
		return nil
	}

	if len(request.Form.Get("request")+request.Form.Get("request_uri")) == 0 {
		return nil
	} else if len(request.Form.Get("request")) > 0 && len(request.Form.Get("request_uri")) > 0 {
		return errorsx.WithStack(ErrInvalidRequest.WithHint("OpenID Connect parameters 'request' and 'request_uri' were both given, but you can use at most one."))
	}

	oidcClient, ok := request.Client.(OpenIDConnectClient)
	if !ok {
		if len(request.Form.Get("request_uri")) > 0 {
			return errorsx.WithStack(ErrRequestURINotSupported.WithHint("OpenID Connect 'request_uri' context was given, but the OAuth 2.0 Client does not implement advanced OpenID Connect capabilities."))
		}
		return errorsx.WithStack(ErrRequestNotSupported.WithHint("OpenID Connect 'request' context was given, but the OAuth 2.0 Client does not implement advanced OpenID Connect capabilities."))
	}

	if oidcClient.GetJSONWebKeys() == nil && len(oidcClient.GetJSONWebKeysURI()) == 0 {
		return errorsx.WithStack(ErrInvalidRequest.WithHint("OpenID Connect 'request' or 'request_uri' context was given, but the OAuth 2.0 Client does not have any JSON Web Keys registered."))
	}

	assertion := request.Form.Get("request")
	if location := request.Form.Get("request_uri"); len(location) > 0 {
		if !stringslice.Has(oidcClient.GetRequestURIs(), location) {
			return errorsx.WithStack(ErrInvalidRequestURI.WithHintf("Request URI '%s' is not whitelisted by the OAuth 2.0 Client.", location))
		}

		hc := f.Config.GetHTTPClient(ctx)
		response, err := hc.Get(location)
		if err != nil {
			return errorsx.WithStack(ErrInvalidRequestURI.WithHintf("Unable to fetch OpenID Connect request parameters from 'request_uri' because: %s.", err.Error()).WithWrap(err).WithDebug(err.Error()))
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			return errorsx.WithStack(ErrInvalidRequestURI.WithHintf("Unable to fetch OpenID Connect request parameters from 'request_uri' because status code '%d' was expected, but got '%d'.", http.StatusOK, response.StatusCode))
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return errorsx.WithStack(ErrInvalidRequestURI.WithHintf("Unable to fetch OpenID Connect request parameters from 'request_uri' because body parsing failed with: %s.", err).WithWrap(err).WithDebug(err.Error()))
		}

		assertion = string(body)
	}

	token, err := jwt.ParseWithClaims(assertion, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		// request_object_signing_alg - OPTIONAL.
		//  JWS [JWS] alg algorithm [JWA] that MUST be used for signing Request Objects sent to the OP. All Request Objects from this Client MUST be rejected,
		// 	if not signed with this algorithm. Request Objects are described in Section 6.1 of OpenID Connect Core 1.0 [OpenID.Core]. This algorithm MUST
		//	be used both when the Request Object is passed by value (using the request parameter) and when it is passed by reference (using the request_uri parameter).
		//	Servers SHOULD support RS256. The value none MAY be used. The default, if omitted, is that any algorithm supported by the OP and the RP MAY be used.
		if oidcClient.GetRequestObjectSigningAlgorithm() != "" && oidcClient.GetRequestObjectSigningAlgorithm() != fmt.Sprintf("%s", t.Header["alg"]) {
			return nil, errorsx.WithStack(ErrInvalidRequestObject.WithHintf("The request object uses signing algorithm '%s', but the requested OAuth 2.0 Client enforces signing algorithm '%s'.", t.Header["alg"], oidcClient.GetRequestObjectSigningAlgorithm()))
		}

		if t.Method == jwt.SigningMethodNone {
			return jwt.UnsafeAllowNoneSignatureType, nil
		}

		switch t.Method {
		case jose.RS256, jose.RS384, jose.RS512:
			key, err := f.findClientPublicJWK(ctx, oidcClient, t, true)
			if err != nil {
				return nil, wrapSigningKeyFailure(
					ErrInvalidRequestObject.WithHint("Unable to retrieve RSA signing key from OAuth 2.0 Client."), err)
			}
			return key, nil
		case jose.ES256, jose.ES384, jose.ES512:
			key, err := f.findClientPublicJWK(ctx, oidcClient, t, false)
			if err != nil {
				return nil, wrapSigningKeyFailure(
					ErrInvalidRequestObject.WithHint("Unable to retrieve ECDSA signing key from OAuth 2.0 Client."), err)
			}
			return key, nil
		case jose.PS256, jose.PS384, jose.PS512:
			key, err := f.findClientPublicJWK(ctx, oidcClient, t, true)
			if err != nil {
				return nil, wrapSigningKeyFailure(
					ErrInvalidRequestObject.WithHint("Unable to retrieve RSA signing key from OAuth 2.0 Client."), err)
			}
			return key, nil
		default:
			return nil, errorsx.WithStack(ErrInvalidRequestObject.WithHintf("This request object uses unsupported signing algorithm '%s'.", t.Header["alg"]))
		}
	})
	if err != nil {
		// Do not re-process already enhanced errors
		var e *jwt.ValidationError
		if errors.As(err, &e) {
			if e.Inner != nil {
				return e.Inner
			}
			return errorsx.WithStack(ErrInvalidRequestObject.WithHint("Unable to verify the request object's signature.").WithWrap(err).WithDebug(err.Error()))
		}
		return err
	} else if err := token.Claims.Valid(); err != nil {
		return errorsx.WithStack(ErrInvalidRequestObject.WithHint("Unable to verify the request object because its claims could not be validated, check if the expiry time is set correctly.").WithWrap(err).WithDebug(err.Error()))
	}

	claims := token.Claims
	for k, v := range claims {
		request.Form.Set(k, fmt.Sprintf("%s", v))
	}

	claimScope := RemoveEmpty(strings.Split(request.Form.Get("scope"), " "))
	for _, s := range scope {
		if !stringslice.Has(claimScope, s) {
			claimScope = append(claimScope, s)
		}
	}

	request.State = request.Form.Get("state")
	request.Form.Set("scope", strings.Join(claimScope, " "))
	return nil
}

func (f *Fosite) validateAuthorizeRedirectURI(_ *http.Request, request *AuthorizeRequest) error {
	// Fetch redirect URI from request
	rawRedirURI := request.Form.Get("redirect_uri")

	// Validate redirect uri
	redirectURI, err := MatchRedirectURIWithClientRedirectURIs(rawRedirURI, request.Client)
	if err != nil {
		return err
	} else if !IsValidRedirectURI(redirectURI) {
		return errorsx.WithStack(ErrInvalidRequest.WithHintf("The redirect URI '%s' contains an illegal character (for example #) or is otherwise invalid.", redirectURI))
	}
	request.RedirectURI = redirectURI
	return nil
}

func (f *Fosite) validateAuthorizeScope(ctx context.Context, _ *http.Request, request *AuthorizeRequest) error {
	scope := RemoveEmpty(strings.Split(request.Form.Get("scope"), " "))
	for _, permission := range scope {
		if !f.Config.GetScopeStrategy(ctx)(request.Client.GetScopes(), permission) {
			return errorsx.WithStack(ErrInvalidScope.WithHintf("The OAuth 2.0 Client is not allowed to request scope '%s'.", permission))
		}
	}
	request.SetRequestedScopes(scope)

	return nil
}

func (f *Fosite) validateResponseTypes(r *http.Request, request *AuthorizeRequest) error {
	// https://tools.ietf.org/html/rfc6749#section-3.1.1
	// Extension response types MAY contain a space-delimited (%x20) list of
	// values, where the order of values does not matter (e.g., response
	// type "a b" is the same as "b a").  The meaning of such composite
	// response types is defined by their respective specifications.
	responseTypes := RemoveEmpty(strings.Split(r.Form.Get("response_type"), " "))
	if len(responseTypes) == 0 {
		return errorsx.WithStack(ErrUnsupportedResponseType.WithHint("`The request is missing the 'response_type' parameter."))
	}

	var found bool
	for _, t := range request.GetClient().GetResponseTypes() {
		if Arguments(responseTypes).Matches(RemoveEmpty(strings.Split(t, " "))...) {
			found = true
			break
		}
	}

	if !found {
		return errorsx.WithStack(ErrUnsupportedResponseType.WithHintf("The client is not allowed to request response_type '%s'.", r.Form.Get("response_type")))
	}

	request.ResponseTypes = responseTypes
	return nil
}

func (f *Fosite) ParseResponseMode(ctx context.Context, r *http.Request, request *AuthorizeRequest) error {
	switch responseMode := r.Form.Get("response_mode"); responseMode {
	case string(ResponseModeDefault):
		request.ResponseMode = ResponseModeDefault
	case string(ResponseModeFragment):
		request.ResponseMode = ResponseModeFragment
	case string(ResponseModeQuery):
		request.ResponseMode = ResponseModeQuery
	case string(ResponseModeFormPost):
		request.ResponseMode = ResponseModeFormPost
	default:
		rm := ResponseModeType(responseMode)
		if f.ResponseModeHandler(ctx).ResponseModes().Has(rm) {
			request.ResponseMode = rm
			break
		}
		return errorsx.WithStack(ErrUnsupportedResponseMode.WithHintf("Request with unsupported response_mode \"%s\".", responseMode))
	}

	return nil
}

func (f *Fosite) validateResponseMode(r *http.Request, request *AuthorizeRequest) error {
	if request.ResponseMode == ResponseModeDefault {
		return nil
	}

	responseModeClient, ok := request.GetClient().(ResponseModeClient)
	if !ok {
		return errorsx.WithStack(ErrUnsupportedResponseMode.WithHintf("The request has response_mode \"%s\". set but registered OAuth 2.0 client doesn't support response_mode", r.Form.Get("response_mode")))
	}

	var found bool
	for _, t := range responseModeClient.GetResponseModes() {
		if request.ResponseMode == t {
			found = true
			break
		}
	}

	if !found {
		return errorsx.WithStack(ErrUnsupportedResponseMode.WithHintf("The client is not allowed to request response_mode '%s'.", r.Form.Get("response_mode")))
	}

	return nil
}

func (f *Fosite) NewAuthorizeRequest(ctx context.Context, r *http.Request) (AuthorizeRequester, error) {
	request := NewAuthorizeRequest()
	request.Request.Lang = i18n.GetLangFromRequest(f.Config.GetMessageCatalog(ctx), r)

	ctx = context.WithValue(ctx, RequestContextKey, r)
	ctx = context.WithValue(ctx, AuthorizeRequestContextKey, request)

	if err := r.ParseMultipartForm(1 << 20); err != nil && err != http.ErrNotMultipart {
		return request, errorsx.WithStack(ErrInvalidRequest.WithHint("Unable to parse HTTP body, make sure to send a properly formatted form request body.").WithWrap(err).WithDebug(err.Error()))
	}
	request.Form = r.Form

	// Save state to the request to be returned in error conditions (https://github.com/ory/hydra/issues/1642)
	request.State = request.Form.Get("state")

	client, err := f.Store.GetClient(ctx, request.GetRequestForm().Get("client_id"))
	if err != nil {
		return request, errorsx.WithStack(ErrInvalidClient.WithHint("The requested OAuth 2.0 Client does not exist.").WithWrap(err).WithDebug(err.Error()))
	}
	request.Client = client

	// Now that the base fields (state and client) are populated, we extract all the information
	// from the request object or request object uri, if one is set.
	//
	// All other parse methods should come afterwards so that we ensure that the data is taken
	// from the request_object if set.
	if err := f.authorizeRequestParametersFromOpenIDConnectRequest(ctx, request); err != nil {
		return request, err
	}

	// The request context is now fully available and we can start processing the individual
	// fields.
	if err := f.ParseResponseMode(ctx, r, request); err != nil {
		return request, err
	}

	if err := f.validateAuthorizeRedirectURI(r, request); err != nil {
		return request, err
	}

	if err := f.validateAuthorizeScope(ctx, r, request); err != nil {
		return request, err
	}

	if err := f.validateAuthorizeAudience(ctx, r, request); err != nil {
		return request, err
	}

	if len(request.Form.Get("registration")) > 0 {
		return request, errorsx.WithStack(ErrRegistrationNotSupported)
	}

	if err := f.validateResponseTypes(r, request); err != nil {
		return request, err
	}

	if err := f.validateResponseMode(r, request); err != nil {
		return request, err
	}

	// A fallback handler to set the default response mode in cases where we can not reach the Authorize Handlers
	// but still need the e.g. correct error response mode.
	if request.GetResponseMode() == ResponseModeDefault {
		if request.ResponseTypes.ExactOne("code") {
			request.SetDefaultResponseMode(ResponseModeQuery)
		} else {
			// If the response type is not `code` it is an implicit/hybrid (fragment) response mode.
			request.SetDefaultResponseMode(ResponseModeFragment)
		}
	}

	// rfc6819 4.4.1.8.  Threat: CSRF Attack against redirect-uri
	// The "state" parameter should be used to link the authorization
	// request with the redirect URI used to deliver the access token (Section 5.3.5).
	//
	// https://tools.ietf.org/html/rfc6819#section-4.4.1.8
	// The "state" parameter should not	be guessable
	if len(request.State) < f.GetMinParameterEntropy(ctx) {
		// We're assuming that using less then, by default, 8 characters for the state can not be considered "unguessable"
		return request, errorsx.WithStack(ErrInvalidState.WithHintf("Request parameter 'state' must be at least be %d characters long to ensure sufficient entropy.", f.GetMinParameterEntropy(ctx)))
	}

	return request, nil
}
