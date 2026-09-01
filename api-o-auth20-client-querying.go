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
	"strings"
	"time"
)


// OAuth20ClientQueryingAPIService OAuth20ClientQueryingAPI service
type OAuth20ClientQueryingAPIService service

type ApiGetClientRequest struct {
	ctx context.Context
	ApiService *OAuth20ClientQueryingAPIService
	clientId string
}

func (r ApiGetClientRequest) Execute() (*ClientResponse, *http.Response, error) {
	return r.ApiService.GetClientExecute(r)
}

// GetClient Get client details
//
// Retrieves detailed information about a specific OAuth2 client including its name, description, redirect URIs, and scopes.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-client/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param clientId ID of the client to retrieve
// @return ApiGetClientRequest
func (a *OAuth20ClientQueryingAPIService) GetClient(ctx context.Context, clientId string) ApiGetClientRequest {
	return ApiGetClientRequest{
		ApiService: a,
		ctx: ctx,
		clientId: clientId,
	}
}

// Execute executes the request
//  @return ClientResponse
func (a *OAuth20ClientQueryingAPIService) GetClientExecute(r ApiGetClientRequest) (*ClientResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ClientResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuth20ClientQueryingAPIService.GetClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/clients/{clientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterValueToString(r.clientId, "clientId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.clientId) < 1 {
		return localVarReturnValue, nil, reportError("clientId must have at least 1 elements")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetail
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetail
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetail
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ProblemDetail
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
			var v ProblemDetail
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetClientInfoRequest struct {
	ctx context.Context
	ApiService *OAuth20ClientQueryingAPIService
	clientId string
}

func (r ApiGetClientInfoRequest) Execute() (*ClientInfoResponse, *http.Response, error) {
	return r.ApiService.GetClientInfoExecute(r)
}

// GetClientInfo Retrieves detailed information for a specific client
//
// Retrieves the detailed information for a client with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-client-info/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param clientId ID of the client to retrieve
// @return ApiGetClientInfoRequest
func (a *OAuth20ClientQueryingAPIService) GetClientInfo(ctx context.Context, clientId string) ApiGetClientInfoRequest {
	return ApiGetClientInfoRequest{
		ApiService: a,
		ctx: ctx,
		clientId: clientId,
	}
}

// Execute executes the request
//  @return ClientInfoResponse
func (a *OAuth20ClientQueryingAPIService) GetClientInfoExecute(r ApiGetClientInfoRequest) (*ClientInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ClientInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuth20ClientQueryingAPIService.GetClientInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/clients/{clientId}/info"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterValueToString(r.clientId, "clientId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.clientId) < 1 {
		return localVarReturnValue, nil, reportError("clientId must have at least 1 elements")
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
			var v ProblemDetail
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

type ApiGetClientsRequest struct {
	ctx context.Context
	ApiService *OAuth20ClientQueryingAPIService
	limit *int32
	lastClientId *string
	lastCreatedOn *time.Time
}

// Pagination limit
func (r ApiGetClientsRequest) Limit(limit int32) ApiGetClientsRequest {	r.limit = &limit
	return r
}

// ID of the last retrieved client
func (r ApiGetClientsRequest) LastClientId(lastClientId string) ApiGetClientsRequest {	r.lastClientId = &lastClientId
	return r
}

// Date of the last retrieved client
func (r ApiGetClientsRequest) LastCreatedOn(lastCreatedOn time.Time) ApiGetClientsRequest {	r.lastCreatedOn = &lastCreatedOn
	return r
}

func (r ApiGetClientsRequest) Execute() (*PageableResponse, *http.Response, error) {
	return r.ApiService.GetClientsExecute(r)
}

// GetClients List clients
//
// Retrieves a paginated list of OAuth2 clients. The results can be paginated using the limit parameter and last seen client ID/creation date.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-clients/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetClientsRequest
func (a *OAuth20ClientQueryingAPIService) GetClients(ctx context.Context) ApiGetClientsRequest {
	return ApiGetClientsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PageableResponse
func (a *OAuth20ClientQueryingAPIService) GetClientsExecute(r ApiGetClientsRequest) (*PageableResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PageableResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuth20ClientQueryingAPIService.GetClients")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/clients"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.limit == nil {
		return localVarReturnValue, nil, reportError("limit is required and must be specified")
	}
	if *r.limit < 1 {
		return localVarReturnValue, nil, reportError("limit must be greater than 1")
	}
	if *r.limit > 50 {
		return localVarReturnValue, nil, reportError("limit must be less than 50")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	if r.lastClientId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_client_id", r.lastClientId, "form", "")
	}
	if r.lastCreatedOn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_created_on", r.lastCreatedOn, "form", "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetail
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetail
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ProblemDetail
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
			var v ProblemDetail
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetClientsInfoRequest struct {
	ctx context.Context
	ApiService *OAuth20ClientQueryingAPIService
	limit *int32
	lastClientId *string
	lastCreatedOn *time.Time
}

// Pagination limit
func (r ApiGetClientsInfoRequest) Limit(limit int32) ApiGetClientsInfoRequest {	r.limit = &limit
	return r
}

// ID of the last retrieved client
func (r ApiGetClientsInfoRequest) LastClientId(lastClientId string) ApiGetClientsInfoRequest {	r.lastClientId = &lastClientId
	return r
}

// Date of the last retrieved client
func (r ApiGetClientsInfoRequest) LastCreatedOn(lastCreatedOn time.Time) ApiGetClientsInfoRequest {	r.lastCreatedOn = &lastCreatedOn
	return r
}

func (r ApiGetClientsInfoRequest) Execute() (*PageableResponseClientInfoResponse, *http.Response, error) {
	return r.ApiService.GetClientsInfoExecute(r)
}

// GetClientsInfo Retrieves a pageable list of client information
//
// Retrieves a paginated list of information for all clients.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-clients-info/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetClientsInfoRequest
func (a *OAuth20ClientQueryingAPIService) GetClientsInfo(ctx context.Context) ApiGetClientsInfoRequest {
	return ApiGetClientsInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PageableResponseClientInfoResponse
func (a *OAuth20ClientQueryingAPIService) GetClientsInfoExecute(r ApiGetClientsInfoRequest) (*PageableResponseClientInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PageableResponseClientInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuth20ClientQueryingAPIService.GetClientsInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/clients/info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.limit == nil {
		return localVarReturnValue, nil, reportError("limit is required and must be specified")
	}
	if *r.limit < 1 {
		return localVarReturnValue, nil, reportError("limit must be greater than 1")
	}
	if *r.limit > 50 {
		return localVarReturnValue, nil, reportError("limit must be less than 50")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	if r.lastClientId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_client_id", r.lastClientId, "form", "")
	}
	if r.lastCreatedOn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_created_on", r.lastCreatedOn, "form", "")
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
			var v ProblemDetail
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

type ApiGetConsentsRequest struct {
	ctx context.Context
	ApiService *OAuth20ClientQueryingAPIService
	limit *int32
	lastModifiedOn *time.Time
}

// Pagination limit
func (r ApiGetConsentsRequest) Limit(limit int32) ApiGetConsentsRequest {	r.limit = &limit
	return r
}

// Date of the last retrieved consent
func (r ApiGetConsentsRequest) LastModifiedOn(lastModifiedOn time.Time) ApiGetConsentsRequest {	r.lastModifiedOn = &lastModifiedOn
	return r
}

func (r ApiGetConsentsRequest) Execute() (*PageableModificationResponse, *http.Response, error) {
	return r.ApiService.GetConsentsExecute(r)
}

// GetConsents Retrieves a pageable list of consents
//
// Retrieves a paginated list of user consents.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-consents/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetConsentsRequest
func (a *OAuth20ClientQueryingAPIService) GetConsents(ctx context.Context) ApiGetConsentsRequest {
	return ApiGetConsentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PageableModificationResponse
func (a *OAuth20ClientQueryingAPIService) GetConsentsExecute(r ApiGetConsentsRequest) (*PageableModificationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PageableModificationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuth20ClientQueryingAPIService.GetConsents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/clients/consents"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.limit == nil {
		return localVarReturnValue, nil, reportError("limit is required and must be specified")
	}
	if *r.limit < 1 {
		return localVarReturnValue, nil, reportError("limit must be greater than 1")
	}
	if *r.limit > 50 {
		return localVarReturnValue, nil, reportError("limit must be less than 50")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	if r.lastModifiedOn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_modified_on", r.lastModifiedOn, "form", "")
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

type ApiGetPublicClientInfoRequest struct {
	ctx context.Context
	ApiService *OAuth20ClientQueryingAPIService
	clientId string
}

func (r ApiGetPublicClientInfoRequest) Execute() (*ClientInfoResponse, *http.Response, error) {
	return r.ApiService.GetPublicClientInfoExecute(r)
}

// GetPublicClientInfo Handles the GET request for public client information
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-public-client-info/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param clientId ID of the client to retrieve
// @return ApiGetPublicClientInfoRequest
func (a *OAuth20ClientQueryingAPIService) GetPublicClientInfo(ctx context.Context, clientId string) ApiGetPublicClientInfoRequest {
	return ApiGetPublicClientInfoRequest{
		ApiService: a,
		ctx: ctx,
		clientId: clientId,
	}
}

// Execute executes the request
//  @return ClientInfoResponse
func (a *OAuth20ClientQueryingAPIService) GetPublicClientInfoExecute(r ApiGetPublicClientInfoRequest) (*ClientInfoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ClientInfoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuth20ClientQueryingAPIService.GetPublicClientInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/clients/{clientId}/public/info"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterValueToString(r.clientId, "clientId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.clientId) < 1 {
		return localVarReturnValue, nil, reportError("clientId must have at least 1 elements")
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
			var v ProblemDetail
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
