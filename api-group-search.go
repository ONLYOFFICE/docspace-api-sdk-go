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


// GroupSearchAPIService GroupSearchAPI service
type GroupSearchAPIService service

type ApiGetGroupsWithFilesSharedRequest struct {
	ctx context.Context
	ApiService *GroupSearchAPIService
	id int32
	excludeShared *bool
	count *int32
	startIndex *int32
	filterValue *string
}

// Specifies whether to exclude the group sharing settings from the response.
func (r ApiGetGroupsWithFilesSharedRequest) ExcludeShared(excludeShared bool) ApiGetGroupsWithFilesSharedRequest {	r.excludeShared = &excludeShared
	return r
}

// The number of groups to retrieve in the request.
func (r ApiGetGroupsWithFilesSharedRequest) Count(count int32) ApiGetGroupsWithFilesSharedRequest {	r.count = &count
	return r
}

// The starting index from which to begin retrieving groups with their sharing settings.
func (r ApiGetGroupsWithFilesSharedRequest) StartIndex(startIndex int32) ApiGetGroupsWithFilesSharedRequest {	r.startIndex = &startIndex
	return r
}

// The text used as a filter for retrieving groups with their sharing settings.
func (r ApiGetGroupsWithFilesSharedRequest) FilterValue(filterValue string) ApiGetGroupsWithFilesSharedRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetGroupsWithFilesSharedRequest) Execute() (*GroupArrayWrapper, *http.Response, error) {
	return r.ApiService.GetGroupsWithFilesSharedExecute(r)
}

// GetGroupsWithFilesShared Get groups with file sharing settings
//
// Returns groups with their sharing settings for a file with the ID specified in request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-groups-with-files-shared/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The group ID.
// @return ApiGetGroupsWithFilesSharedRequest
func (a *GroupSearchAPIService) GetGroupsWithFilesShared(ctx context.Context, id int32) ApiGetGroupsWithFilesSharedRequest {
	return ApiGetGroupsWithFilesSharedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GroupArrayWrapper
func (a *GroupSearchAPIService) GetGroupsWithFilesSharedExecute(r ApiGetGroupsWithFilesSharedRequest) (*GroupArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupSearchAPIService.GetGroupsWithFilesShared")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/group/file/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.excludeShared != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "excludeShared", r.excludeShared, "form", "")
			}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
			}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
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

type ApiGetGroupsWithFoldersSharedRequest struct {
	ctx context.Context
	ApiService *GroupSearchAPIService
	id int32
	excludeShared *bool
	count *int32
	startIndex *int32
	filterValue *string
}

// Specifies whether to exclude the group sharing settings from the response.
func (r ApiGetGroupsWithFoldersSharedRequest) ExcludeShared(excludeShared bool) ApiGetGroupsWithFoldersSharedRequest {	r.excludeShared = &excludeShared
	return r
}

// The number of groups to retrieve in the request.
func (r ApiGetGroupsWithFoldersSharedRequest) Count(count int32) ApiGetGroupsWithFoldersSharedRequest {	r.count = &count
	return r
}

// The starting index from which to begin retrieving groups with their sharing settings.
func (r ApiGetGroupsWithFoldersSharedRequest) StartIndex(startIndex int32) ApiGetGroupsWithFoldersSharedRequest {	r.startIndex = &startIndex
	return r
}

// The text used as a filter for retrieving groups with their sharing settings.
func (r ApiGetGroupsWithFoldersSharedRequest) FilterValue(filterValue string) ApiGetGroupsWithFoldersSharedRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetGroupsWithFoldersSharedRequest) Execute() (*GroupArrayWrapper, *http.Response, error) {
	return r.ApiService.GetGroupsWithFoldersSharedExecute(r)
}

// GetGroupsWithFoldersShared Get groups with folder sharing settings
//
// Returns groups with their sharing settings in a folder with the ID specified in request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-groups-with-folders-shared/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The group ID.
// @return ApiGetGroupsWithFoldersSharedRequest
func (a *GroupSearchAPIService) GetGroupsWithFoldersShared(ctx context.Context, id int32) ApiGetGroupsWithFoldersSharedRequest {
	return ApiGetGroupsWithFoldersSharedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GroupArrayWrapper
func (a *GroupSearchAPIService) GetGroupsWithFoldersSharedExecute(r ApiGetGroupsWithFoldersSharedRequest) (*GroupArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupSearchAPIService.GetGroupsWithFoldersShared")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/group/folder/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.excludeShared != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "excludeShared", r.excludeShared, "form", "")
			}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
			}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
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

type ApiGetGroupsWithRoomsSharedRequest struct {
	ctx context.Context
	ApiService *GroupSearchAPIService
	id int32
	excludeShared *bool
	count *int32
	startIndex *int32
	filterValue *string
}

// Specifies whether to exclude the group sharing settings from the response.
func (r ApiGetGroupsWithRoomsSharedRequest) ExcludeShared(excludeShared bool) ApiGetGroupsWithRoomsSharedRequest {	r.excludeShared = &excludeShared
	return r
}

// The number of groups to retrieve in the request.
func (r ApiGetGroupsWithRoomsSharedRequest) Count(count int32) ApiGetGroupsWithRoomsSharedRequest {	r.count = &count
	return r
}

// The starting index from which to begin retrieving groups with their sharing settings.
func (r ApiGetGroupsWithRoomsSharedRequest) StartIndex(startIndex int32) ApiGetGroupsWithRoomsSharedRequest {	r.startIndex = &startIndex
	return r
}

// The text used as a filter for retrieving groups with their sharing settings.
func (r ApiGetGroupsWithRoomsSharedRequest) FilterValue(filterValue string) ApiGetGroupsWithRoomsSharedRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetGroupsWithRoomsSharedRequest) Execute() (*GroupArrayWrapper, *http.Response, error) {
	return r.ApiService.GetGroupsWithRoomsSharedExecute(r)
}

// GetGroupsWithRoomsShared Get groups with room sharing settings
//
// Returns groups with their sharing settings in a room with the ID specified in request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-groups-with-rooms-shared/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The group ID.
// @return ApiGetGroupsWithRoomsSharedRequest
func (a *GroupSearchAPIService) GetGroupsWithRoomsShared(ctx context.Context, id int32) ApiGetGroupsWithRoomsSharedRequest {
	return ApiGetGroupsWithRoomsSharedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GroupArrayWrapper
func (a *GroupSearchAPIService) GetGroupsWithRoomsSharedExecute(r ApiGetGroupsWithRoomsSharedRequest) (*GroupArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupSearchAPIService.GetGroupsWithRoomsShared")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/group/room/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.excludeShared != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "excludeShared", r.excludeShared, "form", "")
			}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
			}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
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
