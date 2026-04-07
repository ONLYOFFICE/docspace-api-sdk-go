// (c) Copyright Ascensio System SIA 2026
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docspace_api_sdk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// OAuth20AuthorizationAPIService OAuth20AuthorizationAPI service
type OAuth20AuthorizationAPIService service

type ApiAuthorizeOAuthRequest struct {
	ctx context.Context
	ApiService *OAuth20AuthorizationAPIService
	responseType *string
	clientId *string
	redirectUri *string
	scope *string
}

// The OAuth 2.0 response type, must be 'code' for authorization code flow.
func (r ApiAuthorizeOAuthRequest) ResponseType(responseType string) ApiAuthorizeOAuthRequest {	r.responseType = &responseType
	return r
}

// The client identifier issued to the client during registration.
func (r ApiAuthorizeOAuthRequest) ClientId(clientId string) ApiAuthorizeOAuthRequest {	r.clientId = &clientId
	return r
}

// The URL to redirect to after authorization is complete.
func (r ApiAuthorizeOAuthRequest) RedirectUri(redirectUri string) ApiAuthorizeOAuthRequest {	r.redirectUri = &redirectUri
	return r
}

// The space-separated list of requested scope permissions.
func (r ApiAuthorizeOAuthRequest) Scope(scope string) ApiAuthorizeOAuthRequest {	r.scope = &scope
	return r
}

func (r ApiAuthorizeOAuthRequest) Execute() (*http.Response, error) {
	return r.ApiService.AuthorizeOAuthExecute(r)
}

// AuthorizeOAuth OAuth2 Authorization Endpoint
//
// Initiates the OAuth2 authorization flow.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/authorize-o-auth/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAuthorizeOAuthRequest
func (a *OAuth20AuthorizationAPIService) AuthorizeOAuth(ctx context.Context) ApiAuthorizeOAuthRequest {
	return ApiAuthorizeOAuthRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *OAuth20AuthorizationAPIService) AuthorizeOAuthExecute(r ApiAuthorizeOAuthRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuth20AuthorizationAPIService.AuthorizeOAuth")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/authorize"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.responseType == nil {
		return nil, reportError("responseType is required and must be specified")
	}
	if r.clientId == nil {
		return nil, reportError("clientId is required and must be specified")
	}
	if r.redirectUri == nil {
		return nil, reportError("redirectUri is required and must be specified")
	}
	if r.scope == nil {
		return nil, reportError("scope is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "response_type", r.responseType, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "client_id", r.clientId, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "redirect_uri", r.redirectUri, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "scope", r.scope, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiExchangeTokenRequest struct {
	ctx context.Context
	ApiService *OAuth20AuthorizationAPIService
	grantType *string
	code *string
	redirectUri *string
	clientId *string
	clientSecret *string
}

// The OAuth2 grant type, must be 'authorization_code' for the authorization code flow.
func (r ApiExchangeTokenRequest) GrantType(grantType string) ApiExchangeTokenRequest {	r.grantType = &grantType
	return r
}

// A temporary authorization code that is sent to the client to be exchanged for a token.
func (r ApiExchangeTokenRequest) Code(code string) ApiExchangeTokenRequest {	r.code = &code
	return r
}

// The URL where the user will be redirected after successful or unsuccessful authentication.
func (r ApiExchangeTokenRequest) RedirectUri(redirectUri string) ApiExchangeTokenRequest {	r.redirectUri = &redirectUri
	return r
}

// The client identifier issued to the client during registration.
func (r ApiExchangeTokenRequest) ClientId(clientId string) ApiExchangeTokenRequest {	r.clientId = &clientId
	return r
}

// The client secret issued to the client during registration.
func (r ApiExchangeTokenRequest) ClientSecret(clientSecret string) ApiExchangeTokenRequest {	r.clientSecret = &clientSecret
	return r
}

func (r ApiExchangeTokenRequest) Execute() (*ExchangeToken200Response, *http.Response, error) {
	return r.ApiService.ExchangeTokenExecute(r)
}

// ExchangeToken OAuth2 Token Endpoint
//
// Exchange authorization code for access token
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/exchange-token/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiExchangeTokenRequest
func (a *OAuth20AuthorizationAPIService) ExchangeToken(ctx context.Context) ApiExchangeTokenRequest {
	return ApiExchangeTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ExchangeToken200Response
func (a *OAuth20AuthorizationAPIService) ExchangeTokenExecute(r ApiExchangeTokenRequest) (*ExchangeToken200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExchangeToken200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuth20AuthorizationAPIService.ExchangeToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.grantType != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "grant_type", r.grantType, "", "")
	}
	if r.code != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "code", r.code, "", "")
	}
	if r.redirectUri != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "redirect_uri", r.redirectUri, "", "")
	}
	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "client_id", r.clientId, "", "")
	}
	if r.clientSecret != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "client_secret", r.clientSecret, "", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSubmitConsentRequest struct {
	ctx context.Context
	ApiService *OAuth20AuthorizationAPIService
	clientId *string
	state *string
	scope *string
}

// The client identifier issued to the client during registration.
func (r ApiSubmitConsentRequest) ClientId(clientId string) ApiSubmitConsentRequest {	r.clientId = &clientId
	return r
}

// The random string used to solve the CSRF vulnerability problem.
func (r ApiSubmitConsentRequest) State(state string) ApiSubmitConsentRequest {	r.state = &state
	return r
}

// The space-separated list of requested scope permissions.
func (r ApiSubmitConsentRequest) Scope(scope string) ApiSubmitConsentRequest {	r.scope = &scope
	return r
}

func (r ApiSubmitConsentRequest) Execute() (*http.Response, error) {
	return r.ApiService.SubmitConsentExecute(r)
}

// SubmitConsent OAuth2 consent endpoint
//
// Sends consent approval
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/submit-consent/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiSubmitConsentRequest
func (a *OAuth20AuthorizationAPIService) SubmitConsent(ctx context.Context) ApiSubmitConsentRequest {
	return ApiSubmitConsentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *OAuth20AuthorizationAPIService) SubmitConsentExecute(r ApiSubmitConsentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuth20AuthorizationAPIService.SubmitConsent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/authorize"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "client_id", r.clientId, "", "")
	}
	if r.state != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "state", r.state, "", "")
	}
	if r.scope != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "scope", r.scope, "", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
