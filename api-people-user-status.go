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
)


// PeopleUserStatusAPIService PeopleUserStatusAPI service
type PeopleUserStatusAPIService service

type ApiGetByStatusRequest struct {
	ctx context.Context
	ApiService *PeopleUserStatusAPIService
	status EmployeeStatus
	filterBy *string
	count *int32
	startIndex *int32
	sortBy *string
	sortOrder *SortOrder
	filterSeparator *string
	filterValue *string
}

// Specifies the criteria used to filter the profiles in the request.
func (r ApiGetByStatusRequest) FilterBy(filterBy string) ApiGetByStatusRequest {	r.filterBy = &filterBy
	return r
}

// The maximum number of user profiles to retrieve.
func (r ApiGetByStatusRequest) Count(count int32) ApiGetByStatusRequest {	r.count = &count
	return r
}

// The starting index for retrieving data in a paginated request.
func (r ApiGetByStatusRequest) StartIndex(startIndex int32) ApiGetByStatusRequest {	r.startIndex = &startIndex
	return r
}

// Specifies the property or field name by which the results should be sorted.
func (r ApiGetByStatusRequest) SortBy(sortBy string) ApiGetByStatusRequest {	r.sortBy = &sortBy
	return r
}

// The order in which the results are sorted.
func (r ApiGetByStatusRequest) SortOrder(sortOrder SortOrder) ApiGetByStatusRequest {	r.sortOrder = &sortOrder
	return r
}

// Represents the separator used to split multiple filter criteria in a query string.
func (r ApiGetByStatusRequest) FilterSeparator(filterSeparator string) ApiGetByStatusRequest {	r.filterSeparator = &filterSeparator
	return r
}

// A string value representing additional filter criteria used in query parameters.
func (r ApiGetByStatusRequest) FilterValue(filterValue string) ApiGetByStatusRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetByStatusRequest) Execute() (*EmployeeFullArrayWrapper, *http.Response, error) {
	return r.ApiService.GetByStatusExecute(r)
}

// GetByStatus Get profiles by status
//
// Returns a list of profiles filtered by the user status.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-by-status/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param status The user status.
// @return ApiGetByStatusRequest
func (a *PeopleUserStatusAPIService) GetByStatus(ctx context.Context, status EmployeeStatus) ApiGetByStatusRequest {
	return ApiGetByStatusRequest{
		ApiService: a,
		ctx: ctx,
		status: status,
	}
}

// Execute executes the request
//  @return EmployeeFullArrayWrapper
func (a *PeopleUserStatusAPIService) GetByStatusExecute(r ApiGetByStatusRequest) (*EmployeeFullArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeFullArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleUserStatusAPIService.GetByStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/status/{status}"
	localVarPath = strings.Replace(localVarPath, "{"+"status"+"}", url.PathEscape(parameterValueToString(r.status, "status")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterBy", r.filterBy, "form", "")
			}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
			}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
			}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortBy", r.sortBy, "form", "")
			}
	if r.sortOrder != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortOrder", r.sortOrder, "form", "")
			}
	if r.filterSeparator != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterSeparator", r.filterSeparator, "form", "")
			}
	if r.filterValue != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterValue", r.filterValue, "form", "")
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

type ApiUpdateUserActivationStatusRequest struct {
	ctx context.Context
	ApiService *PeopleUserStatusAPIService
	activationstatus EmployeeActivationStatus
	updateMembersRequestDto *UpdateMembersRequestDto
}

// The request parameters for updating the user information.
func (r ApiUpdateUserActivationStatusRequest) UpdateMembersRequestDto(updateMembersRequestDto UpdateMembersRequestDto) ApiUpdateUserActivationStatusRequest {	r.updateMembersRequestDto = &updateMembersRequestDto
	return r
}

func (r ApiUpdateUserActivationStatusRequest) Execute() (*EmployeeFullArrayWrapper, *http.Response, error) {
	return r.ApiService.UpdateUserActivationStatusExecute(r)
}

// UpdateUserActivationStatus Set an activation status to the users
//
// Sets the required activation status to the list of users with the IDs specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/update-user-activation-status/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param activationstatus The new user activation status.
// @return ApiUpdateUserActivationStatusRequest
func (a *PeopleUserStatusAPIService) UpdateUserActivationStatus(ctx context.Context, activationstatus EmployeeActivationStatus) ApiUpdateUserActivationStatusRequest {
	return ApiUpdateUserActivationStatusRequest{
		ApiService: a,
		ctx: ctx,
		activationstatus: activationstatus,
	}
}

// Execute executes the request
//  @return EmployeeFullArrayWrapper
func (a *PeopleUserStatusAPIService) UpdateUserActivationStatusExecute(r ApiUpdateUserActivationStatusRequest) (*EmployeeFullArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeFullArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleUserStatusAPIService.UpdateUserActivationStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/activationstatus/{activationstatus}"
	localVarPath = strings.Replace(localVarPath, "{"+"activationstatus"+"}", url.PathEscape(parameterValueToString(r.activationstatus, "activationstatus")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateMembersRequestDto == nil {
		return localVarReturnValue, nil, reportError("updateMembersRequestDto is required and must be specified")
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
	localVarPostBody = r.updateMembersRequestDto
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

type ApiUpdateUserStatusRequest struct {
	ctx context.Context
	ApiService *PeopleUserStatusAPIService
	status EmployeeStatus
	updateMembersRequestDto *UpdateMembersRequestDto
}

// The request parameters for updating the user information.
func (r ApiUpdateUserStatusRequest) UpdateMembersRequestDto(updateMembersRequestDto UpdateMembersRequestDto) ApiUpdateUserStatusRequest {	r.updateMembersRequestDto = &updateMembersRequestDto
	return r
}

func (r ApiUpdateUserStatusRequest) Execute() (*EmployeeFullArrayWrapper, *http.Response, error) {
	return r.ApiService.UpdateUserStatusExecute(r)
}

// UpdateUserStatus Change a user status
//
// Changes a status of the users with the IDs specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/update-user-status/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param status The new user status.
// @return ApiUpdateUserStatusRequest
func (a *PeopleUserStatusAPIService) UpdateUserStatus(ctx context.Context, status EmployeeStatus) ApiUpdateUserStatusRequest {
	return ApiUpdateUserStatusRequest{
		ApiService: a,
		ctx: ctx,
		status: status,
	}
}

// Execute executes the request
//  @return EmployeeFullArrayWrapper
func (a *PeopleUserStatusAPIService) UpdateUserStatusExecute(r ApiUpdateUserStatusRequest) (*EmployeeFullArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeFullArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleUserStatusAPIService.UpdateUserStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/status/{status}"
	localVarPath = strings.Replace(localVarPath, "{"+"status"+"}", url.PathEscape(parameterValueToString(r.status, "status")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateMembersRequestDto == nil {
		return localVarReturnValue, nil, reportError("updateMembersRequestDto is required and must be specified")
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
	localVarPostBody = r.updateMembersRequestDto
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
