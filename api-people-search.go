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
	"reflect"
)


// PeopleSearchAPIService PeopleSearchAPI service
type PeopleSearchAPIService service

type ApiGetAccountsEntriesWithFilesSharedRequest struct {
	ctx context.Context
	ApiService *PeopleSearchAPIService
	id int32
	employeeStatus *EmployeeStatus
	activationStatus *EmployeeActivationStatus
	excludeShared *bool
	includeShared *bool
	invitedByMe *bool
	inviterId *string
	area *Area
	employeeTypes *[]EmployeeType
	count *int32
	startIndex *int32
	filterSeparator *string
	filterValue *string
}

// The user status.
func (r ApiGetAccountsEntriesWithFilesSharedRequest) EmployeeStatus(employeeStatus EmployeeStatus) ApiGetAccountsEntriesWithFilesSharedRequest {	r.employeeStatus = &employeeStatus
	return r
}

// The user activation status.
func (r ApiGetAccountsEntriesWithFilesSharedRequest) ActivationStatus(activationStatus EmployeeActivationStatus) ApiGetAccountsEntriesWithFilesSharedRequest {	r.activationStatus = &activationStatus
	return r
}

// Specifies whether to exclude the account sharing settings from the response.
func (r ApiGetAccountsEntriesWithFilesSharedRequest) ExcludeShared(excludeShared bool) ApiGetAccountsEntriesWithFilesSharedRequest {	r.excludeShared = &excludeShared
	return r
}

// Specifies whether to include the account sharing settings in the response.
func (r ApiGetAccountsEntriesWithFilesSharedRequest) IncludeShared(includeShared bool) ApiGetAccountsEntriesWithFilesSharedRequest {	r.includeShared = &includeShared
	return r
}

// Specifies whether the user is invited by the current user or not.
func (r ApiGetAccountsEntriesWithFilesSharedRequest) InvitedByMe(invitedByMe bool) ApiGetAccountsEntriesWithFilesSharedRequest {	r.invitedByMe = &invitedByMe
	return r
}

// The inviter ID.
func (r ApiGetAccountsEntriesWithFilesSharedRequest) InviterId(inviterId string) ApiGetAccountsEntriesWithFilesSharedRequest {	r.inviterId = &inviterId
	return r
}

// The area of the account entries.
func (r ApiGetAccountsEntriesWithFilesSharedRequest) Area(area Area) ApiGetAccountsEntriesWithFilesSharedRequest {	r.area = &area
	return r
}

// The list of the user types.
func (r ApiGetAccountsEntriesWithFilesSharedRequest) EmployeeTypes(employeeTypes []EmployeeType) ApiGetAccountsEntriesWithFilesSharedRequest {	r.employeeTypes = &employeeTypes
	return r
}

// The number of items to retrieve in a request.
func (r ApiGetAccountsEntriesWithFilesSharedRequest) Count(count int32) ApiGetAccountsEntriesWithFilesSharedRequest {	r.count = &count
	return r
}

// The starting index for the query results.
func (r ApiGetAccountsEntriesWithFilesSharedRequest) StartIndex(startIndex int32) ApiGetAccountsEntriesWithFilesSharedRequest {	r.startIndex = &startIndex
	return r
}

// Specifies the separator used in filter expressions.
func (r ApiGetAccountsEntriesWithFilesSharedRequest) FilterSeparator(filterSeparator string) ApiGetAccountsEntriesWithFilesSharedRequest {	r.filterSeparator = &filterSeparator
	return r
}

// The text filter applied to the accounts search query.
func (r ApiGetAccountsEntriesWithFilesSharedRequest) FilterValue(filterValue string) ApiGetAccountsEntriesWithFilesSharedRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetAccountsEntriesWithFilesSharedRequest) Execute() (*ObjectArrayWrapper, *http.Response, error) {
	return r.ApiService.GetAccountsEntriesWithFilesSharedExecute(r)
}

// GetAccountsEntriesWithFilesShared Get account entries with file sharing settings
//
// Returns the account entries with their sharing settings for a file with the ID specified in request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-accounts-entries-with-files-shared/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The user ID.
// @return ApiGetAccountsEntriesWithFilesSharedRequest
func (a *PeopleSearchAPIService) GetAccountsEntriesWithFilesShared(ctx context.Context, id int32) ApiGetAccountsEntriesWithFilesSharedRequest {
	return ApiGetAccountsEntriesWithFilesSharedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ObjectArrayWrapper
func (a *PeopleSearchAPIService) GetAccountsEntriesWithFilesSharedExecute(r ApiGetAccountsEntriesWithFilesSharedRequest) (*ObjectArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ObjectArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleSearchAPIService.GetAccountsEntriesWithFilesShared")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/accounts/file/{id}/search"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.employeeStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "employeeStatus", r.employeeStatus, "form", "")
	}
	if r.activationStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "activationStatus", r.activationStatus, "form", "")
	}
	if r.excludeShared != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "excludeShared", r.excludeShared, "form", "")
	}
	if r.includeShared != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeShared", r.includeShared, "form", "")
	}
	if r.invitedByMe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invitedByMe", r.invitedByMe, "form", "")
	}
	if r.inviterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "inviterId", r.inviterId, "form", "")
	}
	if r.area != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "area", r.area, "form", "")
	}
	if r.employeeTypes != nil {
		t := *r.employeeTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "employeeTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "employeeTypes", t, "form", "multi")
		}
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
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
		if localVarHTTPResponse.StatusCode == 401 {
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

type ApiGetAccountsEntriesWithFoldersSharedRequest struct {
	ctx context.Context
	ApiService *PeopleSearchAPIService
	id int32
	employeeStatus *EmployeeStatus
	activationStatus *EmployeeActivationStatus
	excludeShared *bool
	includeShared *bool
	invitedByMe *bool
	inviterId *string
	area *Area
	employeeTypes *[]EmployeeType
	count *int32
	startIndex *int32
	filterSeparator *string
	filterValue *string
}

// The user status.
func (r ApiGetAccountsEntriesWithFoldersSharedRequest) EmployeeStatus(employeeStatus EmployeeStatus) ApiGetAccountsEntriesWithFoldersSharedRequest {	r.employeeStatus = &employeeStatus
	return r
}

// The user activation status.
func (r ApiGetAccountsEntriesWithFoldersSharedRequest) ActivationStatus(activationStatus EmployeeActivationStatus) ApiGetAccountsEntriesWithFoldersSharedRequest {	r.activationStatus = &activationStatus
	return r
}

// Specifies whether to exclude the account sharing settings from the response.
func (r ApiGetAccountsEntriesWithFoldersSharedRequest) ExcludeShared(excludeShared bool) ApiGetAccountsEntriesWithFoldersSharedRequest {	r.excludeShared = &excludeShared
	return r
}

// Specifies whether to include the account sharing settings in the response.
func (r ApiGetAccountsEntriesWithFoldersSharedRequest) IncludeShared(includeShared bool) ApiGetAccountsEntriesWithFoldersSharedRequest {	r.includeShared = &includeShared
	return r
}

// Specifies whether the user is invited by the current user or not.
func (r ApiGetAccountsEntriesWithFoldersSharedRequest) InvitedByMe(invitedByMe bool) ApiGetAccountsEntriesWithFoldersSharedRequest {	r.invitedByMe = &invitedByMe
	return r
}

// The inviter ID.
func (r ApiGetAccountsEntriesWithFoldersSharedRequest) InviterId(inviterId string) ApiGetAccountsEntriesWithFoldersSharedRequest {	r.inviterId = &inviterId
	return r
}

// The area of the account entries.
func (r ApiGetAccountsEntriesWithFoldersSharedRequest) Area(area Area) ApiGetAccountsEntriesWithFoldersSharedRequest {	r.area = &area
	return r
}

// The list of the user types.
func (r ApiGetAccountsEntriesWithFoldersSharedRequest) EmployeeTypes(employeeTypes []EmployeeType) ApiGetAccountsEntriesWithFoldersSharedRequest {	r.employeeTypes = &employeeTypes
	return r
}

// The number of items to retrieve in a request.
func (r ApiGetAccountsEntriesWithFoldersSharedRequest) Count(count int32) ApiGetAccountsEntriesWithFoldersSharedRequest {	r.count = &count
	return r
}

// The starting index for the query results.
func (r ApiGetAccountsEntriesWithFoldersSharedRequest) StartIndex(startIndex int32) ApiGetAccountsEntriesWithFoldersSharedRequest {	r.startIndex = &startIndex
	return r
}

// Specifies the separator used in filter expressions.
func (r ApiGetAccountsEntriesWithFoldersSharedRequest) FilterSeparator(filterSeparator string) ApiGetAccountsEntriesWithFoldersSharedRequest {	r.filterSeparator = &filterSeparator
	return r
}

// The text filter applied to the accounts search query.
func (r ApiGetAccountsEntriesWithFoldersSharedRequest) FilterValue(filterValue string) ApiGetAccountsEntriesWithFoldersSharedRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetAccountsEntriesWithFoldersSharedRequest) Execute() (*ObjectArrayWrapper, *http.Response, error) {
	return r.ApiService.GetAccountsEntriesWithFoldersSharedExecute(r)
}

// GetAccountsEntriesWithFoldersShared Get account entries with folder sharing settings
//
// Returns the account entries with their sharing settings in a folder with the ID specified in request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-accounts-entries-with-folders-shared/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The user ID.
// @return ApiGetAccountsEntriesWithFoldersSharedRequest
func (a *PeopleSearchAPIService) GetAccountsEntriesWithFoldersShared(ctx context.Context, id int32) ApiGetAccountsEntriesWithFoldersSharedRequest {
	return ApiGetAccountsEntriesWithFoldersSharedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ObjectArrayWrapper
func (a *PeopleSearchAPIService) GetAccountsEntriesWithFoldersSharedExecute(r ApiGetAccountsEntriesWithFoldersSharedRequest) (*ObjectArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ObjectArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleSearchAPIService.GetAccountsEntriesWithFoldersShared")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/accounts/folder/{id}/search"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.employeeStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "employeeStatus", r.employeeStatus, "form", "")
	}
	if r.activationStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "activationStatus", r.activationStatus, "form", "")
	}
	if r.excludeShared != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "excludeShared", r.excludeShared, "form", "")
	}
	if r.includeShared != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeShared", r.includeShared, "form", "")
	}
	if r.invitedByMe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invitedByMe", r.invitedByMe, "form", "")
	}
	if r.inviterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "inviterId", r.inviterId, "form", "")
	}
	if r.area != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "area", r.area, "form", "")
	}
	if r.employeeTypes != nil {
		t := *r.employeeTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "employeeTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "employeeTypes", t, "form", "multi")
		}
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
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
		if localVarHTTPResponse.StatusCode == 401 {
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

type ApiGetAccountsEntriesWithRoomsSharedRequest struct {
	ctx context.Context
	ApiService *PeopleSearchAPIService
	id int32
	employeeStatus *EmployeeStatus
	activationStatus *EmployeeActivationStatus
	excludeShared *bool
	includeShared *bool
	invitedByMe *bool
	inviterId *string
	area *Area
	employeeTypes *[]EmployeeType
	count *int32
	startIndex *int32
	filterSeparator *string
	filterValue *string
}

// The user status.
func (r ApiGetAccountsEntriesWithRoomsSharedRequest) EmployeeStatus(employeeStatus EmployeeStatus) ApiGetAccountsEntriesWithRoomsSharedRequest {	r.employeeStatus = &employeeStatus
	return r
}

// The user activation status.
func (r ApiGetAccountsEntriesWithRoomsSharedRequest) ActivationStatus(activationStatus EmployeeActivationStatus) ApiGetAccountsEntriesWithRoomsSharedRequest {	r.activationStatus = &activationStatus
	return r
}

// Specifies whether to exclude the account sharing settings from the response.
func (r ApiGetAccountsEntriesWithRoomsSharedRequest) ExcludeShared(excludeShared bool) ApiGetAccountsEntriesWithRoomsSharedRequest {	r.excludeShared = &excludeShared
	return r
}

// Specifies whether to include the account sharing settings in the response.
func (r ApiGetAccountsEntriesWithRoomsSharedRequest) IncludeShared(includeShared bool) ApiGetAccountsEntriesWithRoomsSharedRequest {	r.includeShared = &includeShared
	return r
}

// Specifies whether the user is invited by the current user or not.
func (r ApiGetAccountsEntriesWithRoomsSharedRequest) InvitedByMe(invitedByMe bool) ApiGetAccountsEntriesWithRoomsSharedRequest {	r.invitedByMe = &invitedByMe
	return r
}

// The inviter ID.
func (r ApiGetAccountsEntriesWithRoomsSharedRequest) InviterId(inviterId string) ApiGetAccountsEntriesWithRoomsSharedRequest {	r.inviterId = &inviterId
	return r
}

// The area of the account entries.
func (r ApiGetAccountsEntriesWithRoomsSharedRequest) Area(area Area) ApiGetAccountsEntriesWithRoomsSharedRequest {	r.area = &area
	return r
}

// The list of the user types.
func (r ApiGetAccountsEntriesWithRoomsSharedRequest) EmployeeTypes(employeeTypes []EmployeeType) ApiGetAccountsEntriesWithRoomsSharedRequest {	r.employeeTypes = &employeeTypes
	return r
}

// The number of items to retrieve in a request.
func (r ApiGetAccountsEntriesWithRoomsSharedRequest) Count(count int32) ApiGetAccountsEntriesWithRoomsSharedRequest {	r.count = &count
	return r
}

// The starting index for the query results.
func (r ApiGetAccountsEntriesWithRoomsSharedRequest) StartIndex(startIndex int32) ApiGetAccountsEntriesWithRoomsSharedRequest {	r.startIndex = &startIndex
	return r
}

// Specifies the separator used in filter expressions.
func (r ApiGetAccountsEntriesWithRoomsSharedRequest) FilterSeparator(filterSeparator string) ApiGetAccountsEntriesWithRoomsSharedRequest {	r.filterSeparator = &filterSeparator
	return r
}

// The text filter applied to the accounts search query.
func (r ApiGetAccountsEntriesWithRoomsSharedRequest) FilterValue(filterValue string) ApiGetAccountsEntriesWithRoomsSharedRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetAccountsEntriesWithRoomsSharedRequest) Execute() (*ObjectArrayWrapper, *http.Response, error) {
	return r.ApiService.GetAccountsEntriesWithRoomsSharedExecute(r)
}

// GetAccountsEntriesWithRoomsShared Get account entries
//
// Returns the account entries with their sharing settings in a room with the ID specified in request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-accounts-entries-with-rooms-shared/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The user ID.
// @return ApiGetAccountsEntriesWithRoomsSharedRequest
func (a *PeopleSearchAPIService) GetAccountsEntriesWithRoomsShared(ctx context.Context, id int32) ApiGetAccountsEntriesWithRoomsSharedRequest {
	return ApiGetAccountsEntriesWithRoomsSharedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ObjectArrayWrapper
func (a *PeopleSearchAPIService) GetAccountsEntriesWithRoomsSharedExecute(r ApiGetAccountsEntriesWithRoomsSharedRequest) (*ObjectArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ObjectArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleSearchAPIService.GetAccountsEntriesWithRoomsShared")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/accounts/room/{id}/search"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.employeeStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "employeeStatus", r.employeeStatus, "form", "")
	}
	if r.activationStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "activationStatus", r.activationStatus, "form", "")
	}
	if r.excludeShared != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "excludeShared", r.excludeShared, "form", "")
	}
	if r.includeShared != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeShared", r.includeShared, "form", "")
	}
	if r.invitedByMe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invitedByMe", r.invitedByMe, "form", "")
	}
	if r.inviterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "inviterId", r.inviterId, "form", "")
	}
	if r.area != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "area", r.area, "form", "")
	}
	if r.employeeTypes != nil {
		t := *r.employeeTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "employeeTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "employeeTypes", t, "form", "multi")
		}
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
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
		if localVarHTTPResponse.StatusCode == 401 {
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

type ApiGetSearchRequest struct {
	ctx context.Context
	ApiService *PeopleSearchAPIService
	query string
	filterBy *string
	filterValue *string
}

// Specifies a filter criteria for the user search query.
func (r ApiGetSearchRequest) FilterBy(filterBy string) ApiGetSearchRequest {	r.filterBy = &filterBy
	return r
}

// The value used for filtering users, allowing additional constraints for the query.
func (r ApiGetSearchRequest) FilterValue(filterValue string) ApiGetSearchRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetSearchRequest) Execute() (*EmployeeFullArrayWrapper, *http.Response, error) {
	return r.ApiService.GetSearchExecute(r)
}

// GetSearch Search users
//
// Returns a list of users matching the search query.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-search/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param query The search query.
// @return ApiGetSearchRequest
func (a *PeopleSearchAPIService) GetSearch(ctx context.Context, query string) ApiGetSearchRequest {
	return ApiGetSearchRequest{
		ApiService: a,
		ctx: ctx,
		query: query,
	}
}

// Execute executes the request
//  @return EmployeeFullArrayWrapper
func (a *PeopleSearchAPIService) GetSearchExecute(r ApiGetSearchRequest) (*EmployeeFullArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeFullArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleSearchAPIService.GetSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/@search/{query}"
	localVarPath = strings.Replace(localVarPath, "{"+"query"+"}", url.PathEscape(parameterValueToString(r.query, "query")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterBy", r.filterBy, "form", "")
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
		if localVarHTTPResponse.StatusCode == 401 {
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

type ApiGetSimpleByFilterRequest struct {
	ctx context.Context
	ApiService *PeopleSearchAPIService
	employeeStatus *EmployeeStatus
	groupId *string
	activationStatus *EmployeeActivationStatus
	employeeType *EmployeeType
	employeeTypes *[]int32
	isAdministrator *bool
	payments *Payments
	accountLoginType *AccountLoginType
	quotaFilter *QuotaFilter
	withoutGroup *bool
	excludeGroup *bool
	invitedByMe *bool
	inviterId *string
	area *Area
	count *int32
	startIndex *int32
	sortBy *string
	sortOrder *SortOrder
	filterSeparator *string
	filterValue *string
}

// The user status.
func (r ApiGetSimpleByFilterRequest) EmployeeStatus(employeeStatus EmployeeStatus) ApiGetSimpleByFilterRequest {	r.employeeStatus = &employeeStatus
	return r
}

// The group ID.
func (r ApiGetSimpleByFilterRequest) GroupId(groupId string) ApiGetSimpleByFilterRequest {	r.groupId = &groupId
	return r
}

// The user activation status.
func (r ApiGetSimpleByFilterRequest) ActivationStatus(activationStatus EmployeeActivationStatus) ApiGetSimpleByFilterRequest {	r.activationStatus = &activationStatus
	return r
}

// The user type.
func (r ApiGetSimpleByFilterRequest) EmployeeType(employeeType EmployeeType) ApiGetSimpleByFilterRequest {	r.employeeType = &employeeType
	return r
}

// The list of user types.
func (r ApiGetSimpleByFilterRequest) EmployeeTypes(employeeTypes []int32) ApiGetSimpleByFilterRequest {	r.employeeTypes = &employeeTypes
	return r
}

// Specifies if the user is an administrator or not.
func (r ApiGetSimpleByFilterRequest) IsAdministrator(isAdministrator bool) ApiGetSimpleByFilterRequest {	r.isAdministrator = &isAdministrator
	return r
}

// The user payment status.
func (r ApiGetSimpleByFilterRequest) Payments(payments Payments) ApiGetSimpleByFilterRequest {	r.payments = &payments
	return r
}

// The account login type.
func (r ApiGetSimpleByFilterRequest) AccountLoginType(accountLoginType AccountLoginType) ApiGetSimpleByFilterRequest {	r.accountLoginType = &accountLoginType
	return r
}

// The quota filter (All - 0, Default - 1, Custom - 2).
func (r ApiGetSimpleByFilterRequest) QuotaFilter(quotaFilter QuotaFilter) ApiGetSimpleByFilterRequest {	r.quotaFilter = &quotaFilter
	return r
}

// Specifies whether the user should be a member of a group or not.
func (r ApiGetSimpleByFilterRequest) WithoutGroup(withoutGroup bool) ApiGetSimpleByFilterRequest {	r.withoutGroup = &withoutGroup
	return r
}

// Specifies whether the user should be a member of the group with the specified ID.
func (r ApiGetSimpleByFilterRequest) ExcludeGroup(excludeGroup bool) ApiGetSimpleByFilterRequest {	r.excludeGroup = &excludeGroup
	return r
}

// Specifies whether the user is invited by the current user or not.
func (r ApiGetSimpleByFilterRequest) InvitedByMe(invitedByMe bool) ApiGetSimpleByFilterRequest {	r.invitedByMe = &invitedByMe
	return r
}

// The inviter ID.
func (r ApiGetSimpleByFilterRequest) InviterId(inviterId string) ApiGetSimpleByFilterRequest {	r.inviterId = &inviterId
	return r
}

// The filter area.
func (r ApiGetSimpleByFilterRequest) Area(area Area) ApiGetSimpleByFilterRequest {	r.area = &area
	return r
}

// The maximum number of items to be retrieved in the response.
func (r ApiGetSimpleByFilterRequest) Count(count int32) ApiGetSimpleByFilterRequest {	r.count = &count
	return r
}

// The zero-based index of the first item to be retrieved in a filtered result set.
func (r ApiGetSimpleByFilterRequest) StartIndex(startIndex int32) ApiGetSimpleByFilterRequest {	r.startIndex = &startIndex
	return r
}

// Specifies the property or field name by which the results should be sorted.
func (r ApiGetSimpleByFilterRequest) SortBy(sortBy string) ApiGetSimpleByFilterRequest {	r.sortBy = &sortBy
	return r
}

// The order in which the results are sorted.
func (r ApiGetSimpleByFilterRequest) SortOrder(sortOrder SortOrder) ApiGetSimpleByFilterRequest {	r.sortOrder = &sortOrder
	return r
}

// Represents the separator used to split filter criteria in query parameters.
func (r ApiGetSimpleByFilterRequest) FilterSeparator(filterSeparator string) ApiGetSimpleByFilterRequest {	r.filterSeparator = &filterSeparator
	return r
}

// The search text used to filter results based on user input.
func (r ApiGetSimpleByFilterRequest) FilterValue(filterValue string) ApiGetSimpleByFilterRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetSimpleByFilterRequest) Execute() (*EmployeeArrayWrapper, *http.Response, error) {
	return r.ApiService.GetSimpleByFilterExecute(r)
}

// GetSimpleByFilter Search users by extended filter
//
// Returns a list of users matching the parameters specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-simple-by-filter/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetSimpleByFilterRequest
func (a *PeopleSearchAPIService) GetSimpleByFilter(ctx context.Context) ApiGetSimpleByFilterRequest {
	return ApiGetSimpleByFilterRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EmployeeArrayWrapper
func (a *PeopleSearchAPIService) GetSimpleByFilterExecute(r ApiGetSimpleByFilterRequest) (*EmployeeArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleSearchAPIService.GetSimpleByFilter")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/simple/filter"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.employeeStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "employeeStatus", r.employeeStatus, "form", "")
	}
	if r.groupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "groupId", r.groupId, "form", "")
	}
	if r.activationStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "activationStatus", r.activationStatus, "form", "")
	}
	if r.employeeType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "employeeType", r.employeeType, "form", "")
	}
	if r.employeeTypes != nil {
		t := *r.employeeTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "employeeTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "employeeTypes", t, "form", "multi")
		}
	}
	if r.isAdministrator != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isAdministrator", r.isAdministrator, "form", "")
	}
	if r.payments != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payments", r.payments, "form", "")
	}
	if r.accountLoginType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "accountLoginType", r.accountLoginType, "form", "")
	}
	if r.quotaFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quotaFilter", r.quotaFilter, "form", "")
	}
	if r.withoutGroup != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "withoutGroup", r.withoutGroup, "form", "")
	}
	if r.excludeGroup != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "excludeGroup", r.excludeGroup, "form", "")
	}
	if r.invitedByMe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invitedByMe", r.invitedByMe, "form", "")
	}
	if r.inviterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "inviterId", r.inviterId, "form", "")
	}
	if r.area != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "area", r.area, "form", "")
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
		if localVarHTTPResponse.StatusCode == 401 {
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

type ApiGetUsersWithFilesSharedRequest struct {
	ctx context.Context
	ApiService *PeopleSearchAPIService
	id int32
	employeeStatus *EmployeeStatus
	activationStatus *EmployeeActivationStatus
	excludeShared *bool
	includeShared *bool
	invitedByMe *bool
	inviterId *string
	area *Area
	employeeTypes *[]EmployeeType
	count *int32
	startIndex *int32
	filterSeparator *string
	filterValue *string
}

// The user status.
func (r ApiGetUsersWithFilesSharedRequest) EmployeeStatus(employeeStatus EmployeeStatus) ApiGetUsersWithFilesSharedRequest {	r.employeeStatus = &employeeStatus
	return r
}

// The user activation status.
func (r ApiGetUsersWithFilesSharedRequest) ActivationStatus(activationStatus EmployeeActivationStatus) ApiGetUsersWithFilesSharedRequest {	r.activationStatus = &activationStatus
	return r
}

// Specifies whether to exclude the user sharing settings or not.
func (r ApiGetUsersWithFilesSharedRequest) ExcludeShared(excludeShared bool) ApiGetUsersWithFilesSharedRequest {	r.excludeShared = &excludeShared
	return r
}

// Specifies whether to include the user sharing settings or not.
func (r ApiGetUsersWithFilesSharedRequest) IncludeShared(includeShared bool) ApiGetUsersWithFilesSharedRequest {	r.includeShared = &includeShared
	return r
}

// Specifies whether the user was invited by the current user or not.
func (r ApiGetUsersWithFilesSharedRequest) InvitedByMe(invitedByMe bool) ApiGetUsersWithFilesSharedRequest {	r.invitedByMe = &invitedByMe
	return r
}

// The inviter ID.
func (r ApiGetUsersWithFilesSharedRequest) InviterId(inviterId string) ApiGetUsersWithFilesSharedRequest {	r.inviterId = &inviterId
	return r
}

// The user area.
func (r ApiGetUsersWithFilesSharedRequest) Area(area Area) ApiGetUsersWithFilesSharedRequest {	r.area = &area
	return r
}

// The list of user types.
func (r ApiGetUsersWithFilesSharedRequest) EmployeeTypes(employeeTypes []EmployeeType) ApiGetUsersWithFilesSharedRequest {	r.employeeTypes = &employeeTypes
	return r
}

// The maximum number of users to be retrieved in the request.
func (r ApiGetUsersWithFilesSharedRequest) Count(count int32) ApiGetUsersWithFilesSharedRequest {	r.count = &count
	return r
}

// The zero-based index of the first record to retrieve in a paged query.
func (r ApiGetUsersWithFilesSharedRequest) StartIndex(startIndex int32) ApiGetUsersWithFilesSharedRequest {	r.startIndex = &startIndex
	return r
}

// The character or string used to separate multiple filter values in a filtering query.
func (r ApiGetUsersWithFilesSharedRequest) FilterSeparator(filterSeparator string) ApiGetUsersWithFilesSharedRequest {	r.filterSeparator = &filterSeparator
	return r
}

// The filter text value used for searching or filtering user results.
func (r ApiGetUsersWithFilesSharedRequest) FilterValue(filterValue string) ApiGetUsersWithFilesSharedRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetUsersWithFilesSharedRequest) Execute() (*EmployeeFullArrayWrapper, *http.Response, error) {
	return r.ApiService.GetUsersWithFilesSharedExecute(r)
}

// GetUsersWithFilesShared Get users with file sharing settings
//
// Returns the users with the sharing settings in a file with the ID specified in request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-users-with-files-shared/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The user ID.
// @return ApiGetUsersWithFilesSharedRequest
func (a *PeopleSearchAPIService) GetUsersWithFilesShared(ctx context.Context, id int32) ApiGetUsersWithFilesSharedRequest {
	return ApiGetUsersWithFilesSharedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return EmployeeFullArrayWrapper
func (a *PeopleSearchAPIService) GetUsersWithFilesSharedExecute(r ApiGetUsersWithFilesSharedRequest) (*EmployeeFullArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeFullArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleSearchAPIService.GetUsersWithFilesShared")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/file/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.employeeStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "employeeStatus", r.employeeStatus, "form", "")
	}
	if r.activationStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "activationStatus", r.activationStatus, "form", "")
	}
	if r.excludeShared != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "excludeShared", r.excludeShared, "form", "")
	}
	if r.includeShared != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeShared", r.includeShared, "form", "")
	}
	if r.invitedByMe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invitedByMe", r.invitedByMe, "form", "")
	}
	if r.inviterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "inviterId", r.inviterId, "form", "")
	}
	if r.area != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "area", r.area, "form", "")
	}
	if r.employeeTypes != nil {
		t := *r.employeeTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "employeeTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "employeeTypes", t, "form", "multi")
		}
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
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
		if localVarHTTPResponse.StatusCode == 401 {
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

type ApiGetUsersWithFoldersSharedRequest struct {
	ctx context.Context
	ApiService *PeopleSearchAPIService
	id int32
	employeeStatus *EmployeeStatus
	activationStatus *EmployeeActivationStatus
	excludeShared *bool
	includeShared *bool
	invitedByMe *bool
	inviterId *string
	area *Area
	employeeTypes *[]EmployeeType
	count *int32
	startIndex *int32
	filterSeparator *string
	filterValue *string
}

// The user status.
func (r ApiGetUsersWithFoldersSharedRequest) EmployeeStatus(employeeStatus EmployeeStatus) ApiGetUsersWithFoldersSharedRequest {	r.employeeStatus = &employeeStatus
	return r
}

// The user activation status.
func (r ApiGetUsersWithFoldersSharedRequest) ActivationStatus(activationStatus EmployeeActivationStatus) ApiGetUsersWithFoldersSharedRequest {	r.activationStatus = &activationStatus
	return r
}

// Specifies whether to exclude the user sharing settings or not.
func (r ApiGetUsersWithFoldersSharedRequest) ExcludeShared(excludeShared bool) ApiGetUsersWithFoldersSharedRequest {	r.excludeShared = &excludeShared
	return r
}

// Specifies whether to include the user sharing settings or not.
func (r ApiGetUsersWithFoldersSharedRequest) IncludeShared(includeShared bool) ApiGetUsersWithFoldersSharedRequest {	r.includeShared = &includeShared
	return r
}

// Specifies whether the user was invited by the current user or not.
func (r ApiGetUsersWithFoldersSharedRequest) InvitedByMe(invitedByMe bool) ApiGetUsersWithFoldersSharedRequest {	r.invitedByMe = &invitedByMe
	return r
}

// The inviter ID.
func (r ApiGetUsersWithFoldersSharedRequest) InviterId(inviterId string) ApiGetUsersWithFoldersSharedRequest {	r.inviterId = &inviterId
	return r
}

// The user area.
func (r ApiGetUsersWithFoldersSharedRequest) Area(area Area) ApiGetUsersWithFoldersSharedRequest {	r.area = &area
	return r
}

// The list of user types.
func (r ApiGetUsersWithFoldersSharedRequest) EmployeeTypes(employeeTypes []EmployeeType) ApiGetUsersWithFoldersSharedRequest {	r.employeeTypes = &employeeTypes
	return r
}

// The maximum number of users to be retrieved in the request.
func (r ApiGetUsersWithFoldersSharedRequest) Count(count int32) ApiGetUsersWithFoldersSharedRequest {	r.count = &count
	return r
}

// The zero-based index of the first record to retrieve in a paged query.
func (r ApiGetUsersWithFoldersSharedRequest) StartIndex(startIndex int32) ApiGetUsersWithFoldersSharedRequest {	r.startIndex = &startIndex
	return r
}

// The character or string used to separate multiple filter values in a filtering query.
func (r ApiGetUsersWithFoldersSharedRequest) FilterSeparator(filterSeparator string) ApiGetUsersWithFoldersSharedRequest {	r.filterSeparator = &filterSeparator
	return r
}

// The filter text value used for searching or filtering user results.
func (r ApiGetUsersWithFoldersSharedRequest) FilterValue(filterValue string) ApiGetUsersWithFoldersSharedRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetUsersWithFoldersSharedRequest) Execute() (*EmployeeFullArrayWrapper, *http.Response, error) {
	return r.ApiService.GetUsersWithFoldersSharedExecute(r)
}

// GetUsersWithFoldersShared Get users with folder sharing settings
//
// Returns the users with the sharing settings in a folder with the ID specified in request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-users-with-folders-shared/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The user ID.
// @return ApiGetUsersWithFoldersSharedRequest
func (a *PeopleSearchAPIService) GetUsersWithFoldersShared(ctx context.Context, id int32) ApiGetUsersWithFoldersSharedRequest {
	return ApiGetUsersWithFoldersSharedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return EmployeeFullArrayWrapper
func (a *PeopleSearchAPIService) GetUsersWithFoldersSharedExecute(r ApiGetUsersWithFoldersSharedRequest) (*EmployeeFullArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeFullArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleSearchAPIService.GetUsersWithFoldersShared")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/folder/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.employeeStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "employeeStatus", r.employeeStatus, "form", "")
	}
	if r.activationStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "activationStatus", r.activationStatus, "form", "")
	}
	if r.excludeShared != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "excludeShared", r.excludeShared, "form", "")
	}
	if r.includeShared != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeShared", r.includeShared, "form", "")
	}
	if r.invitedByMe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invitedByMe", r.invitedByMe, "form", "")
	}
	if r.inviterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "inviterId", r.inviterId, "form", "")
	}
	if r.area != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "area", r.area, "form", "")
	}
	if r.employeeTypes != nil {
		t := *r.employeeTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "employeeTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "employeeTypes", t, "form", "multi")
		}
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
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
		if localVarHTTPResponse.StatusCode == 401 {
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

type ApiGetUsersWithRoomSharedRequest struct {
	ctx context.Context
	ApiService *PeopleSearchAPIService
	id int32
	employeeStatus *EmployeeStatus
	activationStatus *EmployeeActivationStatus
	excludeShared *bool
	includeShared *bool
	invitedByMe *bool
	inviterId *string
	area *Area
	employeeTypes *[]EmployeeType
	count *int32
	startIndex *int32
	filterSeparator *string
	filterValue *string
}

// The user status.
func (r ApiGetUsersWithRoomSharedRequest) EmployeeStatus(employeeStatus EmployeeStatus) ApiGetUsersWithRoomSharedRequest {	r.employeeStatus = &employeeStatus
	return r
}

// The user activation status.
func (r ApiGetUsersWithRoomSharedRequest) ActivationStatus(activationStatus EmployeeActivationStatus) ApiGetUsersWithRoomSharedRequest {	r.activationStatus = &activationStatus
	return r
}

// Specifies whether to exclude the user sharing settings or not.
func (r ApiGetUsersWithRoomSharedRequest) ExcludeShared(excludeShared bool) ApiGetUsersWithRoomSharedRequest {	r.excludeShared = &excludeShared
	return r
}

// Specifies whether to include the user sharing settings or not.
func (r ApiGetUsersWithRoomSharedRequest) IncludeShared(includeShared bool) ApiGetUsersWithRoomSharedRequest {	r.includeShared = &includeShared
	return r
}

// Specifies whether the user was invited by the current user or not.
func (r ApiGetUsersWithRoomSharedRequest) InvitedByMe(invitedByMe bool) ApiGetUsersWithRoomSharedRequest {	r.invitedByMe = &invitedByMe
	return r
}

// The inviter ID.
func (r ApiGetUsersWithRoomSharedRequest) InviterId(inviterId string) ApiGetUsersWithRoomSharedRequest {	r.inviterId = &inviterId
	return r
}

// The user area.
func (r ApiGetUsersWithRoomSharedRequest) Area(area Area) ApiGetUsersWithRoomSharedRequest {	r.area = &area
	return r
}

// The list of user types.
func (r ApiGetUsersWithRoomSharedRequest) EmployeeTypes(employeeTypes []EmployeeType) ApiGetUsersWithRoomSharedRequest {	r.employeeTypes = &employeeTypes
	return r
}

// The maximum number of users to be retrieved in the request.
func (r ApiGetUsersWithRoomSharedRequest) Count(count int32) ApiGetUsersWithRoomSharedRequest {	r.count = &count
	return r
}

// The zero-based index of the first record to retrieve in a paged query.
func (r ApiGetUsersWithRoomSharedRequest) StartIndex(startIndex int32) ApiGetUsersWithRoomSharedRequest {	r.startIndex = &startIndex
	return r
}

// The character or string used to separate multiple filter values in a filtering query.
func (r ApiGetUsersWithRoomSharedRequest) FilterSeparator(filterSeparator string) ApiGetUsersWithRoomSharedRequest {	r.filterSeparator = &filterSeparator
	return r
}

// The filter text value used for searching or filtering user results.
func (r ApiGetUsersWithRoomSharedRequest) FilterValue(filterValue string) ApiGetUsersWithRoomSharedRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetUsersWithRoomSharedRequest) Execute() (*EmployeeFullArrayWrapper, *http.Response, error) {
	return r.ApiService.GetUsersWithRoomSharedExecute(r)
}

// GetUsersWithRoomShared Get users with room sharing settings
//
// Returns the users with the sharing settings in a room with the ID specified in request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-users-with-room-shared/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The user ID.
// @return ApiGetUsersWithRoomSharedRequest
func (a *PeopleSearchAPIService) GetUsersWithRoomShared(ctx context.Context, id int32) ApiGetUsersWithRoomSharedRequest {
	return ApiGetUsersWithRoomSharedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return EmployeeFullArrayWrapper
func (a *PeopleSearchAPIService) GetUsersWithRoomSharedExecute(r ApiGetUsersWithRoomSharedRequest) (*EmployeeFullArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeFullArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleSearchAPIService.GetUsersWithRoomShared")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/room/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.employeeStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "employeeStatus", r.employeeStatus, "form", "")
	}
	if r.activationStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "activationStatus", r.activationStatus, "form", "")
	}
	if r.excludeShared != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "excludeShared", r.excludeShared, "form", "")
	}
	if r.includeShared != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeShared", r.includeShared, "form", "")
	}
	if r.invitedByMe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invitedByMe", r.invitedByMe, "form", "")
	}
	if r.inviterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "inviterId", r.inviterId, "form", "")
	}
	if r.area != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "area", r.area, "form", "")
	}
	if r.employeeTypes != nil {
		t := *r.employeeTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "employeeTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "employeeTypes", t, "form", "multi")
		}
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
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
		if localVarHTTPResponse.StatusCode == 401 {
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

type ApiSearchUsersByExtendedFilterRequest struct {
	ctx context.Context
	ApiService *PeopleSearchAPIService
	employeeStatus *EmployeeStatus
	groupId *string
	activationStatus *EmployeeActivationStatus
	employeeType *EmployeeType
	employeeTypes *[]int32
	isAdministrator *bool
	payments *Payments
	accountLoginType *AccountLoginType
	quotaFilter *QuotaFilter
	withoutGroup *bool
	excludeGroup *bool
	invitedByMe *bool
	inviterId *string
	area *Area
	count *int32
	startIndex *int32
	sortBy *string
	sortOrder *SortOrder
	filterSeparator *string
	filterValue *string
}

// The user status.
func (r ApiSearchUsersByExtendedFilterRequest) EmployeeStatus(employeeStatus EmployeeStatus) ApiSearchUsersByExtendedFilterRequest {	r.employeeStatus = &employeeStatus
	return r
}

// The group ID.
func (r ApiSearchUsersByExtendedFilterRequest) GroupId(groupId string) ApiSearchUsersByExtendedFilterRequest {	r.groupId = &groupId
	return r
}

// The user activation status.
func (r ApiSearchUsersByExtendedFilterRequest) ActivationStatus(activationStatus EmployeeActivationStatus) ApiSearchUsersByExtendedFilterRequest {	r.activationStatus = &activationStatus
	return r
}

// The user type.
func (r ApiSearchUsersByExtendedFilterRequest) EmployeeType(employeeType EmployeeType) ApiSearchUsersByExtendedFilterRequest {	r.employeeType = &employeeType
	return r
}

// The list of user types.
func (r ApiSearchUsersByExtendedFilterRequest) EmployeeTypes(employeeTypes []int32) ApiSearchUsersByExtendedFilterRequest {	r.employeeTypes = &employeeTypes
	return r
}

// Specifies if the user is an administrator or not.
func (r ApiSearchUsersByExtendedFilterRequest) IsAdministrator(isAdministrator bool) ApiSearchUsersByExtendedFilterRequest {	r.isAdministrator = &isAdministrator
	return r
}

// The user payment status.
func (r ApiSearchUsersByExtendedFilterRequest) Payments(payments Payments) ApiSearchUsersByExtendedFilterRequest {	r.payments = &payments
	return r
}

// The account login type.
func (r ApiSearchUsersByExtendedFilterRequest) AccountLoginType(accountLoginType AccountLoginType) ApiSearchUsersByExtendedFilterRequest {	r.accountLoginType = &accountLoginType
	return r
}

// The quota filter (All - 0, Default - 1, Custom - 2).
func (r ApiSearchUsersByExtendedFilterRequest) QuotaFilter(quotaFilter QuotaFilter) ApiSearchUsersByExtendedFilterRequest {	r.quotaFilter = &quotaFilter
	return r
}

// Specifies whether the user should be a member of a group or not.
func (r ApiSearchUsersByExtendedFilterRequest) WithoutGroup(withoutGroup bool) ApiSearchUsersByExtendedFilterRequest {	r.withoutGroup = &withoutGroup
	return r
}

// Specifies whether the user should be a member of the group with the specified ID.
func (r ApiSearchUsersByExtendedFilterRequest) ExcludeGroup(excludeGroup bool) ApiSearchUsersByExtendedFilterRequest {	r.excludeGroup = &excludeGroup
	return r
}

// Specifies whether the user is invited by the current user or not.
func (r ApiSearchUsersByExtendedFilterRequest) InvitedByMe(invitedByMe bool) ApiSearchUsersByExtendedFilterRequest {	r.invitedByMe = &invitedByMe
	return r
}

// The inviter ID.
func (r ApiSearchUsersByExtendedFilterRequest) InviterId(inviterId string) ApiSearchUsersByExtendedFilterRequest {	r.inviterId = &inviterId
	return r
}

// The filter area.
func (r ApiSearchUsersByExtendedFilterRequest) Area(area Area) ApiSearchUsersByExtendedFilterRequest {	r.area = &area
	return r
}

// The maximum number of items to be retrieved in the response.
func (r ApiSearchUsersByExtendedFilterRequest) Count(count int32) ApiSearchUsersByExtendedFilterRequest {	r.count = &count
	return r
}

// The zero-based index of the first item to be retrieved in a filtered result set.
func (r ApiSearchUsersByExtendedFilterRequest) StartIndex(startIndex int32) ApiSearchUsersByExtendedFilterRequest {	r.startIndex = &startIndex
	return r
}

// Specifies the property or field name by which the results should be sorted.
func (r ApiSearchUsersByExtendedFilterRequest) SortBy(sortBy string) ApiSearchUsersByExtendedFilterRequest {	r.sortBy = &sortBy
	return r
}

// The order in which the results are sorted.
func (r ApiSearchUsersByExtendedFilterRequest) SortOrder(sortOrder SortOrder) ApiSearchUsersByExtendedFilterRequest {	r.sortOrder = &sortOrder
	return r
}

// Represents the separator used to split filter criteria in query parameters.
func (r ApiSearchUsersByExtendedFilterRequest) FilterSeparator(filterSeparator string) ApiSearchUsersByExtendedFilterRequest {	r.filterSeparator = &filterSeparator
	return r
}

// The search text used to filter results based on user input.
func (r ApiSearchUsersByExtendedFilterRequest) FilterValue(filterValue string) ApiSearchUsersByExtendedFilterRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiSearchUsersByExtendedFilterRequest) Execute() (*EmployeeFullArrayWrapper, *http.Response, error) {
	return r.ApiService.SearchUsersByExtendedFilterExecute(r)
}

// SearchUsersByExtendedFilter Search users with detailed information by extended filter
//
// Returns a list of users with full information about them matching the parameters specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/search-users-by-extended-filter/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiSearchUsersByExtendedFilterRequest
func (a *PeopleSearchAPIService) SearchUsersByExtendedFilter(ctx context.Context) ApiSearchUsersByExtendedFilterRequest {
	return ApiSearchUsersByExtendedFilterRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EmployeeFullArrayWrapper
func (a *PeopleSearchAPIService) SearchUsersByExtendedFilterExecute(r ApiSearchUsersByExtendedFilterRequest) (*EmployeeFullArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeFullArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleSearchAPIService.SearchUsersByExtendedFilter")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/filter"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.employeeStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "employeeStatus", r.employeeStatus, "form", "")
	}
	if r.groupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "groupId", r.groupId, "form", "")
	}
	if r.activationStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "activationStatus", r.activationStatus, "form", "")
	}
	if r.employeeType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "employeeType", r.employeeType, "form", "")
	}
	if r.employeeTypes != nil {
		t := *r.employeeTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "employeeTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "employeeTypes", t, "form", "multi")
		}
	}
	if r.isAdministrator != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isAdministrator", r.isAdministrator, "form", "")
	}
	if r.payments != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payments", r.payments, "form", "")
	}
	if r.accountLoginType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "accountLoginType", r.accountLoginType, "form", "")
	}
	if r.quotaFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quotaFilter", r.quotaFilter, "form", "")
	}
	if r.withoutGroup != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "withoutGroup", r.withoutGroup, "form", "")
	}
	if r.excludeGroup != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "excludeGroup", r.excludeGroup, "form", "")
	}
	if r.invitedByMe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invitedByMe", r.invitedByMe, "form", "")
	}
	if r.inviterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "inviterId", r.inviterId, "form", "")
	}
	if r.area != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "area", r.area, "form", "")
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
		if localVarHTTPResponse.StatusCode == 401 {
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

type ApiSearchUsersByQueryRequest struct {
	ctx context.Context
	ApiService *PeopleSearchAPIService
	query *string
}

// The search query.
func (r ApiSearchUsersByQueryRequest) Query(query string) ApiSearchUsersByQueryRequest {	r.query = &query
	return r
}

func (r ApiSearchUsersByQueryRequest) Execute() (*EmployeeArrayWrapper, *http.Response, error) {
	return r.ApiService.SearchUsersByQueryExecute(r)
}

// SearchUsersByQuery Search users (using query parameters)
//
// Returns a list of users matching the search query. This method uses the query parameters.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/search-users-by-query/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiSearchUsersByQueryRequest
func (a *PeopleSearchAPIService) SearchUsersByQuery(ctx context.Context) ApiSearchUsersByQueryRequest {
	return ApiSearchUsersByQueryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EmployeeArrayWrapper
func (a *PeopleSearchAPIService) SearchUsersByQueryExecute(r ApiSearchUsersByQueryRequest) (*EmployeeArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleSearchAPIService.SearchUsersByQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.query != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "query", r.query, "form", "")
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
		if localVarHTTPResponse.StatusCode == 401 {
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

type ApiSearchUsersByStatusRequest struct {
	ctx context.Context
	ApiService *PeopleSearchAPIService
	status EmployeeStatus
	query *string
	filterBy *string
	filterValue *string
}

// The advanced search query.
func (r ApiSearchUsersByStatusRequest) Query(query string) ApiSearchUsersByStatusRequest {	r.query = &query
	return r
}

// Specifies the criteria used to filter search results in advanced queries.
func (r ApiSearchUsersByStatusRequest) FilterBy(filterBy string) ApiSearchUsersByStatusRequest {	r.filterBy = &filterBy
	return r
}

// The value used to filter the search query.
func (r ApiSearchUsersByStatusRequest) FilterValue(filterValue string) ApiSearchUsersByStatusRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiSearchUsersByStatusRequest) Execute() (*EmployeeFullArrayWrapper, *http.Response, error) {
	return r.ApiService.SearchUsersByStatusExecute(r)
}

// SearchUsersByStatus Search users by status filter
//
// Returns a list of users matching the status filter and search query.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/search-users-by-status/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param status The user status.
// @return ApiSearchUsersByStatusRequest
func (a *PeopleSearchAPIService) SearchUsersByStatus(ctx context.Context, status EmployeeStatus) ApiSearchUsersByStatusRequest {
	return ApiSearchUsersByStatusRequest{
		ApiService: a,
		ctx: ctx,
		status: status,
	}
}

// Execute executes the request
//  @return EmployeeFullArrayWrapper
func (a *PeopleSearchAPIService) SearchUsersByStatusExecute(r ApiSearchUsersByStatusRequest) (*EmployeeFullArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeFullArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleSearchAPIService.SearchUsersByStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/status/{status}/search"
	localVarPath = strings.Replace(localVarPath, "{"+"status"+"}", url.PathEscape(parameterValueToString(r.status, "status")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.query != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "query", r.query, "form", "")
	}
	if r.filterBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterBy", r.filterBy, "form", "")
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
		if localVarHTTPResponse.StatusCode == 401 {
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
