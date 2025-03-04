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
	"reflect"
)

const MinParameterEntropy = 8

var defaultResponseModeHandler = &DefaultResponseModeHandler{}

// AuthorizeEndpointHandlers is a list of AuthorizeEndpointHandler
type AuthorizeEndpointHandlers []AuthorizeEndpointHandler

// Append adds an AuthorizeEndpointHandler to this list. Ignores duplicates based on reflect.TypeOf.
func (a *AuthorizeEndpointHandlers) Append(h AuthorizeEndpointHandler) {
	for _, this := range *a {
		if reflect.TypeOf(this) == reflect.TypeOf(h) {
			return
		}
	}

	*a = append(*a, h)
}

// TokenEndpointHandlers is a list of TokenEndpointHandler
type TokenEndpointHandlers []TokenEndpointHandler

// Append adds an TokenEndpointHandler to this list. Ignores duplicates based on reflect.TypeOf.
func (t *TokenEndpointHandlers) Append(h TokenEndpointHandler) {
	for _, this := range *t {
		if reflect.TypeOf(this) == reflect.TypeOf(h) {
			return
		}
	}

	*t = append(*t, h)
}

// TokenIntrospectionHandlers is a list of TokenValidator
type TokenIntrospectionHandlers []TokenIntrospector

// Append adds an AccessTokenValidator to this list. Ignores duplicates based on reflect.TypeOf.
func (t *TokenIntrospectionHandlers) Append(h TokenIntrospector) {
	for _, this := range *t {
		if reflect.TypeOf(this) == reflect.TypeOf(h) {
			return
		}
	}

	*t = append(*t, h)
}

// RevocationHandlers is a list of RevocationHandler
type RevocationHandlers []RevocationHandler

// Append adds an RevocationHandler to this list. Ignores duplicates based on reflect.TypeOf.
func (t *RevocationHandlers) Append(h RevocationHandler) {
	for _, this := range *t {
		if reflect.TypeOf(this) == reflect.TypeOf(h) {
			return
		}
	}

	*t = append(*t, h)
}

var _ OAuth2Provider = (*Fosite)(nil)

type Configurator interface {
	IDTokenIssuerProvider
	IDTokenLifespanProvider
	AllowedPromptsProvider
	EnforcePKCEProvider
	EnforcePKCEForPublicClientsProvider
	EnablePKCEPlainChallengeMethodProvider
	GrantTypeJWTBearerCanSkipClientAuthProvider
	GrantTypeJWTBearerIDOptionalProvider
	GrantTypeJWTBearerIssuedDateOptionalProvider
	GetJWTMaxDurationProvider
	AudienceStrategyProvider
	ScopeStrategyProvider
	RedirectSecureCheckerProvider
	OmitRedirectScopeParamProvider
	SanitationAllowedProvider
	JWTScopeFieldProvider
	AccessTokenIssuerProvider
	DisableRefreshTokenValidationProvider
	RefreshTokenScopesProvider
	AccessTokenLifespanProvider
	RefreshTokenLifespanProvider
	AuthorizeCodeLifespanProvider
	TokenEntropyProvider
	RotatedGlobalSecretsProvider
	GlobalSecretProvider
	JWKSFetcherStrategyProvider
	HTTPClientProvider
	ScopeStrategyProvider
	AudienceStrategyProvider
	MinParameterEntropyProvider
	HMACHashingProvider
	ClientAuthenticationStrategyProvider
	ResponseModeHandlerExtensionProvider
	SendDebugMessagesToClientsProvider
	JWKSFetcherStrategyProvider
	ClientAuthenticationStrategyProvider
	ResponseModeHandlerExtensionProvider
	MessageCatalogProvider
	FormPostHTMLTemplateProvider
	TokenURLProvider
	GetSecretsHashingProvider
	AuthorizeEndpointHandlersProvider
	TokenEndpointHandlersProvider
	TokenIntrospectionHandlersProvider
	RevocationHandlersProvider
	UseLegacyErrorFormatProvider
}

func NewOAuth2Provider(s Storage, c Configurator) *Fosite {
	return &Fosite{Store: s, Config: c}
}

// Fosite implements OAuth2Provider.
type Fosite struct {
	Store Storage

	Config Configurator
}

// GetMinParameterEntropy returns MinParameterEntropy if set. Defaults to fosite.MinParameterEntropy.
func (f *Fosite) GetMinParameterEntropy(ctx context.Context) int {
	if mp := f.Config.GetMinParameterEntropy(ctx); mp > 0 {
		return mp
	}

	return MinParameterEntropy
}

func (f *Fosite) ResponseModeHandler(ctx context.Context) ResponseModeHandler {
	if ext := f.Config.GetResponseModeHandlerExtension(ctx); ext != nil {
		return ext
	}
	return defaultResponseModeHandler
}
