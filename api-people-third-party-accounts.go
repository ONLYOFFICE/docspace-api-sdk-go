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


// PeopleThirdPartyAccountsAPIService PeopleThirdPartyAccountsAPI service
type PeopleThirdPartyAccountsAPIService service

type ApiGetThirdPartyAuthProvidersRequest struct {
	ctx context.Context
	ApiService *PeopleThirdPartyAccountsAPIService
	inviteView *bool
	settingsView *bool
	clientCallback *string
	fromOnly *string
}

// Specifies whether to return providers that are available for invitation links, i.e. the user can login or register through these providers.
func (r ApiGetThirdPartyAuthProvidersRequest) InviteView(inviteView bool) ApiGetThirdPartyAuthProvidersRequest {	r.inviteView = &inviteView
	return r
}

// Specifies whether to display the provider settings in a pop-up window (true) or redirect them to the desktop application (false).
func (r ApiGetThirdPartyAuthProvidersRequest) SettingsView(settingsView bool) ApiGetThirdPartyAuthProvidersRequest {	r.settingsView = &settingsView
	return r
}

// The method that is called after authentication.
func (r ApiGetThirdPartyAuthProvidersRequest) ClientCallback(clientCallback string) ApiGetThirdPartyAuthProvidersRequest {	r.clientCallback = &clientCallback
	return r
}

// The provider name if a response is required only from this provider.
func (r ApiGetThirdPartyAuthProvidersRequest) FromOnly(fromOnly string) ApiGetThirdPartyAuthProvidersRequest {	r.fromOnly = &fromOnly
	return r
}

func (r ApiGetThirdPartyAuthProvidersRequest) Execute() (*AccountInfoArrayWrapper, *http.Response, error) {
	return r.ApiService.GetThirdPartyAuthProvidersExecute(r)
}

// GetThirdPartyAuthProviders Get third-party accounts
//
// Returns a list of the available third-party accounts.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-third-party-auth-providers/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetThirdPartyAuthProvidersRequest
func (a *PeopleThirdPartyAccountsAPIService) GetThirdPartyAuthProviders(ctx context.Context) ApiGetThirdPartyAuthProvidersRequest {
	return ApiGetThirdPartyAuthProvidersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AccountInfoArrayWrapper
func (a *PeopleThirdPartyAccountsAPIService) GetThirdPartyAuthProvidersExecute(r ApiGetThirdPartyAuthProvidersRequest) (*AccountInfoArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountInfoArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleThirdPartyAccountsAPIService.GetThirdPartyAuthProviders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/thirdparty/providers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.inviteView != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "inviteView", r.inviteView, "form", "")
	}
	if r.settingsView != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "settingsView", r.settingsView, "form", "")
	}
	if r.clientCallback != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientCallback", r.clientCallback, "form", "")
	}
	if r.fromOnly != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fromOnly", r.fromOnly, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiLinkThirdPartyAccountRequest struct {
	ctx context.Context
	ApiService *PeopleThirdPartyAccountsAPIService
	linkAccountRequestDto *LinkAccountRequestDto
}

func (r ApiLinkThirdPartyAccountRequest) LinkAccountRequestDto(linkAccountRequestDto LinkAccountRequestDto) ApiLinkThirdPartyAccountRequest {	r.linkAccountRequestDto = &linkAccountRequestDto
	return r
}

func (r ApiLinkThirdPartyAccountRequest) Execute() (*http.Response, error) {
	return r.ApiService.LinkThirdPartyAccountExecute(r)
}

// LinkThirdPartyAccount Link a third-pary account
//
// Links a third-party account specified in the request to the user profile.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/link-third-party-account/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiLinkThirdPartyAccountRequest
func (a *PeopleThirdPartyAccountsAPIService) LinkThirdPartyAccount(ctx context.Context) ApiLinkThirdPartyAccountRequest {
	return ApiLinkThirdPartyAccountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PeopleThirdPartyAccountsAPIService) LinkThirdPartyAccountExecute(r ApiLinkThirdPartyAccountRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleThirdPartyAccountsAPIService.LinkThirdPartyAccount")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/thirdparty/linkaccount"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.linkAccountRequestDto
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiSignupThirdPartyAccountRequest struct {
	ctx context.Context
	ApiService *PeopleThirdPartyAccountsAPIService
	signupAccountRequestDto *SignupAccountRequestDto
}

func (r ApiSignupThirdPartyAccountRequest) SignupAccountRequestDto(signupAccountRequestDto SignupAccountRequestDto) ApiSignupThirdPartyAccountRequest {	r.signupAccountRequestDto = &signupAccountRequestDto
	return r
}

func (r ApiSignupThirdPartyAccountRequest) Execute() (*EmployeeWrapper, *http.Response, error) {
	return r.ApiService.SignupThirdPartyAccountExecute(r)
}

// SignupThirdPartyAccount Create a third-pary account
//
// Creates a third-party account with the parameters specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/signup-third-party-account/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiSignupThirdPartyAccountRequest
func (a *PeopleThirdPartyAccountsAPIService) SignupThirdPartyAccount(ctx context.Context) ApiSignupThirdPartyAccountRequest {
	return ApiSignupThirdPartyAccountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EmployeeWrapper
func (a *PeopleThirdPartyAccountsAPIService) SignupThirdPartyAccountExecute(r ApiSignupThirdPartyAccountRequest) (*EmployeeWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleThirdPartyAccountsAPIService.SignupThirdPartyAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/thirdparty/signup"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.signupAccountRequestDto
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
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiUnlinkThirdPartyAccountRequest struct {
	ctx context.Context
	ApiService *PeopleThirdPartyAccountsAPIService
	provider *string
}

// The provider name.
func (r ApiUnlinkThirdPartyAccountRequest) Provider(provider string) ApiUnlinkThirdPartyAccountRequest {	r.provider = &provider
	return r
}

func (r ApiUnlinkThirdPartyAccountRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnlinkThirdPartyAccountExecute(r)
}

// UnlinkThirdPartyAccount Unlink a third-pary account
//
// Unlinks a third-party account specified in the request from the user profile.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/unlink-third-party-account/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiUnlinkThirdPartyAccountRequest
func (a *PeopleThirdPartyAccountsAPIService) UnlinkThirdPartyAccount(ctx context.Context) ApiUnlinkThirdPartyAccountRequest {
	return ApiUnlinkThirdPartyAccountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PeopleThirdPartyAccountsAPIService) UnlinkThirdPartyAccountExecute(r ApiUnlinkThirdPartyAccountRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleThirdPartyAccountsAPIService.UnlinkThirdPartyAccount")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/thirdparty/unlinkaccount"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.provider != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "provider", r.provider, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorApiResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
