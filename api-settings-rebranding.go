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


// SettingsRebrandingAPIService SettingsRebrandingAPI service
type SettingsRebrandingAPIService service

type ApiDeleteAdditionalWhiteLabelSettingsRequest struct {
	ctx context.Context
	ApiService *SettingsRebrandingAPIService
}

func (r ApiDeleteAdditionalWhiteLabelSettingsRequest) Execute() (*AdditionalWhiteLabelSettingsWrapper, *http.Response, error) {
	return r.ApiService.DeleteAdditionalWhiteLabelSettingsExecute(r)
}

// DeleteAdditionalWhiteLabelSettings Delete the additional white label settings
//
// Deletes the additional white label settings.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-additional-white-label-settings/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiDeleteAdditionalWhiteLabelSettingsRequest
func (a *SettingsRebrandingAPIService) DeleteAdditionalWhiteLabelSettings(ctx context.Context) ApiDeleteAdditionalWhiteLabelSettingsRequest {
	return ApiDeleteAdditionalWhiteLabelSettingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AdditionalWhiteLabelSettingsWrapper
func (a *SettingsRebrandingAPIService) DeleteAdditionalWhiteLabelSettingsExecute(r ApiDeleteAdditionalWhiteLabelSettingsRequest) (*AdditionalWhiteLabelSettingsWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AdditionalWhiteLabelSettingsWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsRebrandingAPIService.DeleteAdditionalWhiteLabelSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/rebranding/additional"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiDeleteCompanyWhiteLabelSettingsRequest struct {
	ctx context.Context
	ApiService *SettingsRebrandingAPIService
}

func (r ApiDeleteCompanyWhiteLabelSettingsRequest) Execute() (*CompanyWhiteLabelSettingsWrapper, *http.Response, error) {
	return r.ApiService.DeleteCompanyWhiteLabelSettingsExecute(r)
}

// DeleteCompanyWhiteLabelSettings Delete the company white label settings
//
// Deletes the company white label settings.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-company-white-label-settings/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiDeleteCompanyWhiteLabelSettingsRequest
func (a *SettingsRebrandingAPIService) DeleteCompanyWhiteLabelSettings(ctx context.Context) ApiDeleteCompanyWhiteLabelSettingsRequest {
	return ApiDeleteCompanyWhiteLabelSettingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CompanyWhiteLabelSettingsWrapper
func (a *SettingsRebrandingAPIService) DeleteCompanyWhiteLabelSettingsExecute(r ApiDeleteCompanyWhiteLabelSettingsRequest) (*CompanyWhiteLabelSettingsWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CompanyWhiteLabelSettingsWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsRebrandingAPIService.DeleteCompanyWhiteLabelSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/rebranding/company"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetAdditionalWhiteLabelSettingsRequest struct {
	ctx context.Context
	ApiService *SettingsRebrandingAPIService
}

func (r ApiGetAdditionalWhiteLabelSettingsRequest) Execute() (*AdditionalWhiteLabelSettingsWrapper, *http.Response, error) {
	return r.ApiService.GetAdditionalWhiteLabelSettingsExecute(r)
}

// GetAdditionalWhiteLabelSettings Get the additional white label settings
//
// Returns the additional white label settings.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-additional-white-label-settings/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetAdditionalWhiteLabelSettingsRequest
func (a *SettingsRebrandingAPIService) GetAdditionalWhiteLabelSettings(ctx context.Context) ApiGetAdditionalWhiteLabelSettingsRequest {
	return ApiGetAdditionalWhiteLabelSettingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AdditionalWhiteLabelSettingsWrapper
func (a *SettingsRebrandingAPIService) GetAdditionalWhiteLabelSettingsExecute(r ApiGetAdditionalWhiteLabelSettingsRequest) (*AdditionalWhiteLabelSettingsWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AdditionalWhiteLabelSettingsWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsRebrandingAPIService.GetAdditionalWhiteLabelSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/rebranding/additional"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetCompanyWhiteLabelSettingsRequest struct {
	ctx context.Context
	ApiService *SettingsRebrandingAPIService
}

func (r ApiGetCompanyWhiteLabelSettingsRequest) Execute() (*CompanyWhiteLabelSettingsWrapper, *http.Response, error) {
	return r.ApiService.GetCompanyWhiteLabelSettingsExecute(r)
}

// GetCompanyWhiteLabelSettings Get the company white label settings
//
// Returns the company white label settings.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-company-white-label-settings/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetCompanyWhiteLabelSettingsRequest
func (a *SettingsRebrandingAPIService) GetCompanyWhiteLabelSettings(ctx context.Context) ApiGetCompanyWhiteLabelSettingsRequest {
	return ApiGetCompanyWhiteLabelSettingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CompanyWhiteLabelSettingsWrapper
func (a *SettingsRebrandingAPIService) GetCompanyWhiteLabelSettingsExecute(r ApiGetCompanyWhiteLabelSettingsRequest) (*CompanyWhiteLabelSettingsWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CompanyWhiteLabelSettingsWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsRebrandingAPIService.GetCompanyWhiteLabelSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/rebranding/company"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetEnableWhitelabelRequest struct {
	ctx context.Context
	ApiService *SettingsRebrandingAPIService
}

func (r ApiGetEnableWhitelabelRequest) Execute() (*BooleanWrapper, *http.Response, error) {
	return r.ApiService.GetEnableWhitelabelExecute(r)
}

// GetEnableWhitelabel Check the white label availability
//
// Checks if the white label is enabled or not.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-enable-whitelabel/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetEnableWhitelabelRequest
func (a *SettingsRebrandingAPIService) GetEnableWhitelabel(ctx context.Context) ApiGetEnableWhitelabelRequest {
	return ApiGetEnableWhitelabelRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BooleanWrapper
func (a *SettingsRebrandingAPIService) GetEnableWhitelabelExecute(r ApiGetEnableWhitelabelRequest) (*BooleanWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BooleanWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsRebrandingAPIService.GetEnableWhitelabel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/enablewhitelabel"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetIsDefaultWhiteLabelLogoTextRequest struct {
	ctx context.Context
	ApiService *SettingsRebrandingAPIService
	isDark *bool
	isDefault *bool
}

// Specifies if the white label logo is for the dark theme or not.
func (r ApiGetIsDefaultWhiteLabelLogoTextRequest) IsDark(isDark bool) ApiGetIsDefaultWhiteLabelLogoTextRequest {	r.isDark = &isDark
	return r
}

// Specifies if the logo is for a default tenant or not.
func (r ApiGetIsDefaultWhiteLabelLogoTextRequest) IsDefault(isDefault bool) ApiGetIsDefaultWhiteLabelLogoTextRequest {	r.isDefault = &isDefault
	return r
}

func (r ApiGetIsDefaultWhiteLabelLogoTextRequest) Execute() (*IsDefaultWhiteLabelLogosWrapper, *http.Response, error) {
	return r.ApiService.GetIsDefaultWhiteLabelLogoTextExecute(r)
}

// GetIsDefaultWhiteLabelLogoText Check the default white label logo text
//
// Specifies if the white label logo text is default or not.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-is-default-white-label-logo-text/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetIsDefaultWhiteLabelLogoTextRequest
func (a *SettingsRebrandingAPIService) GetIsDefaultWhiteLabelLogoText(ctx context.Context) ApiGetIsDefaultWhiteLabelLogoTextRequest {
	return ApiGetIsDefaultWhiteLabelLogoTextRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IsDefaultWhiteLabelLogosWrapper
func (a *SettingsRebrandingAPIService) GetIsDefaultWhiteLabelLogoTextExecute(r ApiGetIsDefaultWhiteLabelLogoTextRequest) (*IsDefaultWhiteLabelLogosWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IsDefaultWhiteLabelLogosWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsRebrandingAPIService.GetIsDefaultWhiteLabelLogoText")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/whitelabel/logotext/isdefault"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isDark != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsDark", r.isDark, "form", "")
			}
	if r.isDefault != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsDefault", r.isDefault, "form", "")
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

type ApiGetIsDefaultWhiteLabelLogosRequest struct {
	ctx context.Context
	ApiService *SettingsRebrandingAPIService
	isDark *bool
	isDefault *bool
}

// Specifies if the white label logo is for the dark theme or not.
func (r ApiGetIsDefaultWhiteLabelLogosRequest) IsDark(isDark bool) ApiGetIsDefaultWhiteLabelLogosRequest {	r.isDark = &isDark
	return r
}

// Specifies if the logo is for a default tenant or not.
func (r ApiGetIsDefaultWhiteLabelLogosRequest) IsDefault(isDefault bool) ApiGetIsDefaultWhiteLabelLogosRequest {	r.isDefault = &isDefault
	return r
}

func (r ApiGetIsDefaultWhiteLabelLogosRequest) Execute() (*IsDefaultWhiteLabelLogosArrayWrapper, *http.Response, error) {
	return r.ApiService.GetIsDefaultWhiteLabelLogosExecute(r)
}

// GetIsDefaultWhiteLabelLogos Check the default white label logos
//
// Specifies if the white label logos are default or not.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-is-default-white-label-logos/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetIsDefaultWhiteLabelLogosRequest
func (a *SettingsRebrandingAPIService) GetIsDefaultWhiteLabelLogos(ctx context.Context) ApiGetIsDefaultWhiteLabelLogosRequest {
	return ApiGetIsDefaultWhiteLabelLogosRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IsDefaultWhiteLabelLogosArrayWrapper
func (a *SettingsRebrandingAPIService) GetIsDefaultWhiteLabelLogosExecute(r ApiGetIsDefaultWhiteLabelLogosRequest) (*IsDefaultWhiteLabelLogosArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IsDefaultWhiteLabelLogosArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsRebrandingAPIService.GetIsDefaultWhiteLabelLogos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/whitelabel/logos/isdefault"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isDark != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsDark", r.isDark, "form", "")
			}
	if r.isDefault != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsDefault", r.isDefault, "form", "")
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

type ApiGetLicensorDataRequest struct {
	ctx context.Context
	ApiService *SettingsRebrandingAPIService
}

func (r ApiGetLicensorDataRequest) Execute() (*CompanyWhiteLabelSettingsArrayWrapper, *http.Response, error) {
	return r.ApiService.GetLicensorDataExecute(r)
}

// GetLicensorData Get the licensor data
//
// Returns the licensor data.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-licensor-data/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetLicensorDataRequest
func (a *SettingsRebrandingAPIService) GetLicensorData(ctx context.Context) ApiGetLicensorDataRequest {
	return ApiGetLicensorDataRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CompanyWhiteLabelSettingsArrayWrapper
func (a *SettingsRebrandingAPIService) GetLicensorDataExecute(r ApiGetLicensorDataRequest) (*CompanyWhiteLabelSettingsArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CompanyWhiteLabelSettingsArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsRebrandingAPIService.GetLicensorData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/companywhitelabel"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetWhiteLabelLogoTextRequest struct {
	ctx context.Context
	ApiService *SettingsRebrandingAPIService
	isDark *bool
	isDefault *bool
}

// Specifies if the white label logo is for the dark theme or not.
func (r ApiGetWhiteLabelLogoTextRequest) IsDark(isDark bool) ApiGetWhiteLabelLogoTextRequest {	r.isDark = &isDark
	return r
}

// Specifies if the logo is for a default tenant or not.
func (r ApiGetWhiteLabelLogoTextRequest) IsDefault(isDefault bool) ApiGetWhiteLabelLogoTextRequest {	r.isDefault = &isDefault
	return r
}

func (r ApiGetWhiteLabelLogoTextRequest) Execute() (*StringWrapper, *http.Response, error) {
	return r.ApiService.GetWhiteLabelLogoTextExecute(r)
}

// GetWhiteLabelLogoText Get the white label logo text
//
// Returns the white label logo text.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-white-label-logo-text/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetWhiteLabelLogoTextRequest
func (a *SettingsRebrandingAPIService) GetWhiteLabelLogoText(ctx context.Context) ApiGetWhiteLabelLogoTextRequest {
	return ApiGetWhiteLabelLogoTextRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StringWrapper
func (a *SettingsRebrandingAPIService) GetWhiteLabelLogoTextExecute(r ApiGetWhiteLabelLogoTextRequest) (*StringWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StringWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsRebrandingAPIService.GetWhiteLabelLogoText")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/whitelabel/logotext"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isDark != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsDark", r.isDark, "form", "")
			}
	if r.isDefault != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsDefault", r.isDefault, "form", "")
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

type ApiGetWhiteLabelLogosRequest struct {
	ctx context.Context
	ApiService *SettingsRebrandingAPIService
	isDark *bool
	isDefault *bool
}

// Specifies if the white label logo is for the dark theme or not.
func (r ApiGetWhiteLabelLogosRequest) IsDark(isDark bool) ApiGetWhiteLabelLogosRequest {	r.isDark = &isDark
	return r
}

// Specifies if the logo is for a default tenant or not.
func (r ApiGetWhiteLabelLogosRequest) IsDefault(isDefault bool) ApiGetWhiteLabelLogosRequest {	r.isDefault = &isDefault
	return r
}

func (r ApiGetWhiteLabelLogosRequest) Execute() (*WhiteLabelItemArrayWrapper, *http.Response, error) {
	return r.ApiService.GetWhiteLabelLogosExecute(r)
}

// GetWhiteLabelLogos Get the white label logos
//
// Returns the white label logos.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-white-label-logos/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetWhiteLabelLogosRequest
func (a *SettingsRebrandingAPIService) GetWhiteLabelLogos(ctx context.Context) ApiGetWhiteLabelLogosRequest {
	return ApiGetWhiteLabelLogosRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WhiteLabelItemArrayWrapper
func (a *SettingsRebrandingAPIService) GetWhiteLabelLogosExecute(r ApiGetWhiteLabelLogosRequest) (*WhiteLabelItemArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WhiteLabelItemArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsRebrandingAPIService.GetWhiteLabelLogos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/whitelabel/logos"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isDark != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsDark", r.isDark, "form", "")
			}
	if r.isDefault != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsDefault", r.isDefault, "form", "")
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

type ApiRestoreWhiteLabelLogoTextRequest struct {
	ctx context.Context
	ApiService *SettingsRebrandingAPIService
	isDark *bool
	isDefault *bool
}

// Specifies if the white label logo is for the dark theme or not.
func (r ApiRestoreWhiteLabelLogoTextRequest) IsDark(isDark bool) ApiRestoreWhiteLabelLogoTextRequest {	r.isDark = &isDark
	return r
}

// Specifies if the logo is for a default tenant or not.
func (r ApiRestoreWhiteLabelLogoTextRequest) IsDefault(isDefault bool) ApiRestoreWhiteLabelLogoTextRequest {	r.isDefault = &isDefault
	return r
}

func (r ApiRestoreWhiteLabelLogoTextRequest) Execute() (*BooleanWrapper, *http.Response, error) {
	return r.ApiService.RestoreWhiteLabelLogoTextExecute(r)
}

// RestoreWhiteLabelLogoText Restore the white label logo text
//
// Restores the white label logo text.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/restore-white-label-logo-text/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiRestoreWhiteLabelLogoTextRequest
func (a *SettingsRebrandingAPIService) RestoreWhiteLabelLogoText(ctx context.Context) ApiRestoreWhiteLabelLogoTextRequest {
	return ApiRestoreWhiteLabelLogoTextRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BooleanWrapper
func (a *SettingsRebrandingAPIService) RestoreWhiteLabelLogoTextExecute(r ApiRestoreWhiteLabelLogoTextRequest) (*BooleanWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BooleanWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsRebrandingAPIService.RestoreWhiteLabelLogoText")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/whitelabel/logotext/restore"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isDark != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsDark", r.isDark, "form", "")
			}
	if r.isDefault != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsDefault", r.isDefault, "form", "")
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

type ApiRestoreWhiteLabelLogosRequest struct {
	ctx context.Context
	ApiService *SettingsRebrandingAPIService
	isDark *bool
	isDefault *bool
}

// Specifies if the white label logo is for the dark theme or not.
func (r ApiRestoreWhiteLabelLogosRequest) IsDark(isDark bool) ApiRestoreWhiteLabelLogosRequest {	r.isDark = &isDark
	return r
}

// Specifies if the logo is for a default tenant or not.
func (r ApiRestoreWhiteLabelLogosRequest) IsDefault(isDefault bool) ApiRestoreWhiteLabelLogosRequest {	r.isDefault = &isDefault
	return r
}

func (r ApiRestoreWhiteLabelLogosRequest) Execute() (*BooleanWrapper, *http.Response, error) {
	return r.ApiService.RestoreWhiteLabelLogosExecute(r)
}

// RestoreWhiteLabelLogos Restore the white label logos
//
// Restores the white label logos.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/restore-white-label-logos/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiRestoreWhiteLabelLogosRequest
func (a *SettingsRebrandingAPIService) RestoreWhiteLabelLogos(ctx context.Context) ApiRestoreWhiteLabelLogosRequest {
	return ApiRestoreWhiteLabelLogosRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BooleanWrapper
func (a *SettingsRebrandingAPIService) RestoreWhiteLabelLogosExecute(r ApiRestoreWhiteLabelLogosRequest) (*BooleanWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BooleanWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsRebrandingAPIService.RestoreWhiteLabelLogos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/whitelabel/logos/restore"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isDark != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsDark", r.isDark, "form", "")
			}
	if r.isDefault != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsDefault", r.isDefault, "form", "")
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

type ApiSaveAdditionalWhiteLabelSettingsRequest struct {
	ctx context.Context
	ApiService *SettingsRebrandingAPIService
	additionalWhiteLabelSettingsWrapper *AdditionalWhiteLabelSettingsWrapper
}

func (r ApiSaveAdditionalWhiteLabelSettingsRequest) AdditionalWhiteLabelSettingsWrapper(additionalWhiteLabelSettingsWrapper AdditionalWhiteLabelSettingsWrapper) ApiSaveAdditionalWhiteLabelSettingsRequest {	r.additionalWhiteLabelSettingsWrapper = &additionalWhiteLabelSettingsWrapper
	return r
}

func (r ApiSaveAdditionalWhiteLabelSettingsRequest) Execute() (*BooleanWrapper, *http.Response, error) {
	return r.ApiService.SaveAdditionalWhiteLabelSettingsExecute(r)
}

// SaveAdditionalWhiteLabelSettings Save the additional white label settings
//
// Saves the additional white label settings specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/save-additional-white-label-settings/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiSaveAdditionalWhiteLabelSettingsRequest
func (a *SettingsRebrandingAPIService) SaveAdditionalWhiteLabelSettings(ctx context.Context) ApiSaveAdditionalWhiteLabelSettingsRequest {
	return ApiSaveAdditionalWhiteLabelSettingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BooleanWrapper
func (a *SettingsRebrandingAPIService) SaveAdditionalWhiteLabelSettingsExecute(r ApiSaveAdditionalWhiteLabelSettingsRequest) (*BooleanWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BooleanWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsRebrandingAPIService.SaveAdditionalWhiteLabelSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/rebranding/additional"

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
	localVarPostBody = r.additionalWhiteLabelSettingsWrapper
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

type ApiSaveCompanyWhiteLabelSettingsRequest struct {
	ctx context.Context
	ApiService *SettingsRebrandingAPIService
	companyWhiteLabelSettingsWrapper *CompanyWhiteLabelSettingsWrapper
}

func (r ApiSaveCompanyWhiteLabelSettingsRequest) CompanyWhiteLabelSettingsWrapper(companyWhiteLabelSettingsWrapper CompanyWhiteLabelSettingsWrapper) ApiSaveCompanyWhiteLabelSettingsRequest {	r.companyWhiteLabelSettingsWrapper = &companyWhiteLabelSettingsWrapper
	return r
}

func (r ApiSaveCompanyWhiteLabelSettingsRequest) Execute() (*BooleanWrapper, *http.Response, error) {
	return r.ApiService.SaveCompanyWhiteLabelSettingsExecute(r)
}

// SaveCompanyWhiteLabelSettings Save the company white label settings
//
// Saves the company white label settings specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/save-company-white-label-settings/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiSaveCompanyWhiteLabelSettingsRequest
func (a *SettingsRebrandingAPIService) SaveCompanyWhiteLabelSettings(ctx context.Context) ApiSaveCompanyWhiteLabelSettingsRequest {
	return ApiSaveCompanyWhiteLabelSettingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BooleanWrapper
func (a *SettingsRebrandingAPIService) SaveCompanyWhiteLabelSettingsExecute(r ApiSaveCompanyWhiteLabelSettingsRequest) (*BooleanWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BooleanWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsRebrandingAPIService.SaveCompanyWhiteLabelSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/rebranding/company"

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
	localVarPostBody = r.companyWhiteLabelSettingsWrapper
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

type ApiSaveWhiteLabelLogoTextRequest struct {
	ctx context.Context
	ApiService *SettingsRebrandingAPIService
	isDark *bool
	isDefault *bool
	whiteLabelRequestsDto *WhiteLabelRequestsDto
}

// Specifies if the white label logo is for the dark theme or not.
func (r ApiSaveWhiteLabelLogoTextRequest) IsDark(isDark bool) ApiSaveWhiteLabelLogoTextRequest {	r.isDark = &isDark
	return r
}

// Specifies if the logo is for a default tenant or not.
func (r ApiSaveWhiteLabelLogoTextRequest) IsDefault(isDefault bool) ApiSaveWhiteLabelLogoTextRequest {	r.isDefault = &isDefault
	return r
}

func (r ApiSaveWhiteLabelLogoTextRequest) WhiteLabelRequestsDto(whiteLabelRequestsDto WhiteLabelRequestsDto) ApiSaveWhiteLabelLogoTextRequest {	r.whiteLabelRequestsDto = &whiteLabelRequestsDto
	return r
}

func (r ApiSaveWhiteLabelLogoTextRequest) Execute() (*BooleanWrapper, *http.Response, error) {
	return r.ApiService.SaveWhiteLabelLogoTextExecute(r)
}

// SaveWhiteLabelLogoText Save the white label logo text settings
//
// Saves the white label logo text specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/save-white-label-logo-text/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiSaveWhiteLabelLogoTextRequest
func (a *SettingsRebrandingAPIService) SaveWhiteLabelLogoText(ctx context.Context) ApiSaveWhiteLabelLogoTextRequest {
	return ApiSaveWhiteLabelLogoTextRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BooleanWrapper
func (a *SettingsRebrandingAPIService) SaveWhiteLabelLogoTextExecute(r ApiSaveWhiteLabelLogoTextRequest) (*BooleanWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BooleanWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsRebrandingAPIService.SaveWhiteLabelLogoText")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/whitelabel/logotext/save"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isDark != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsDark", r.isDark, "form", "")
			}
	if r.isDefault != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsDefault", r.isDefault, "form", "")
			}
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
	localVarPostBody = r.whiteLabelRequestsDto
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

type ApiSaveWhiteLabelSettingsRequest struct {
	ctx context.Context
	ApiService *SettingsRebrandingAPIService
	isDark *bool
	isDefault *bool
	whiteLabelRequestsDto *WhiteLabelRequestsDto
}

// Specifies if the white label logo is for the dark theme or not.
func (r ApiSaveWhiteLabelSettingsRequest) IsDark(isDark bool) ApiSaveWhiteLabelSettingsRequest {	r.isDark = &isDark
	return r
}

// Specifies if the logo is for a default tenant or not.
func (r ApiSaveWhiteLabelSettingsRequest) IsDefault(isDefault bool) ApiSaveWhiteLabelSettingsRequest {	r.isDefault = &isDefault
	return r
}

func (r ApiSaveWhiteLabelSettingsRequest) WhiteLabelRequestsDto(whiteLabelRequestsDto WhiteLabelRequestsDto) ApiSaveWhiteLabelSettingsRequest {	r.whiteLabelRequestsDto = &whiteLabelRequestsDto
	return r
}

func (r ApiSaveWhiteLabelSettingsRequest) Execute() (*BooleanWrapper, *http.Response, error) {
	return r.ApiService.SaveWhiteLabelSettingsExecute(r)
}

// SaveWhiteLabelSettings Save the white label logos
//
// Saves the white label logos specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/save-white-label-settings/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiSaveWhiteLabelSettingsRequest
func (a *SettingsRebrandingAPIService) SaveWhiteLabelSettings(ctx context.Context) ApiSaveWhiteLabelSettingsRequest {
	return ApiSaveWhiteLabelSettingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BooleanWrapper
func (a *SettingsRebrandingAPIService) SaveWhiteLabelSettingsExecute(r ApiSaveWhiteLabelSettingsRequest) (*BooleanWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BooleanWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsRebrandingAPIService.SaveWhiteLabelSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/whitelabel/logos/save"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isDark != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsDark", r.isDark, "form", "")
			}
	if r.isDefault != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsDefault", r.isDefault, "form", "")
			}
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
	localVarPostBody = r.whiteLabelRequestsDto
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

type ApiSaveWhiteLabelSettingsFromFilesRequest struct {
	ctx context.Context
	ApiService *SettingsRebrandingAPIService
	isDark *bool
	isDefault *bool
}

// Specifies if the white label logo is for the dark theme or not.
func (r ApiSaveWhiteLabelSettingsFromFilesRequest) IsDark(isDark bool) ApiSaveWhiteLabelSettingsFromFilesRequest {	r.isDark = &isDark
	return r
}

// Specifies if the logo is for a default tenant or not.
func (r ApiSaveWhiteLabelSettingsFromFilesRequest) IsDefault(isDefault bool) ApiSaveWhiteLabelSettingsFromFilesRequest {	r.isDefault = &isDefault
	return r
}

func (r ApiSaveWhiteLabelSettingsFromFilesRequest) Execute() (*BooleanWrapper, *http.Response, error) {
	return r.ApiService.SaveWhiteLabelSettingsFromFilesExecute(r)
}

// SaveWhiteLabelSettingsFromFiles Save the white label logos from files
//
// Saves the white label logos from files.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/save-white-label-settings-from-files/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiSaveWhiteLabelSettingsFromFilesRequest
func (a *SettingsRebrandingAPIService) SaveWhiteLabelSettingsFromFiles(ctx context.Context) ApiSaveWhiteLabelSettingsFromFilesRequest {
	return ApiSaveWhiteLabelSettingsFromFilesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BooleanWrapper
func (a *SettingsRebrandingAPIService) SaveWhiteLabelSettingsFromFilesExecute(r ApiSaveWhiteLabelSettingsFromFilesRequest) (*BooleanWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BooleanWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsRebrandingAPIService.SaveWhiteLabelSettingsFromFiles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/whitelabel/logos/savefromfiles"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isDark != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsDark", r.isDark, "form", "")
			}
	if r.isDefault != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsDefault", r.isDefault, "form", "")
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
