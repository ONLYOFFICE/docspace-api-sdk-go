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


// AIAgentsAPIService AIAgentsAPI service
type AIAgentsAPIService service

type ApiCreateAgentRequest struct {
	ctx context.Context
	ApiService *AIAgentsAPIService
	createAgentRequestDto *CreateAgentRequestDto
}

func (r ApiCreateAgentRequest) CreateAgentRequestDto(createAgentRequestDto CreateAgentRequestDto) ApiCreateAgentRequest {	r.createAgentRequestDto = &createAgentRequestDto
	return r
}

func (r ApiCreateAgentRequest) Execute() (*FolderIntegerWrapper, *http.Response, error) {
	return r.ApiService.CreateAgentExecute(r)
}

// CreateAgent Create an ai agent
//
// Creates an ai agent.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/create-agent/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiCreateAgentRequest
func (a *AIAgentsAPIService) CreateAgent(ctx context.Context) ApiCreateAgentRequest {
	return ApiCreateAgentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FolderIntegerWrapper
func (a *AIAgentsAPIService) CreateAgentExecute(r ApiCreateAgentRequest) (*FolderIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FolderIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAgentsAPIService.CreateAgent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/agents"

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
	localVarPostBody = r.createAgentRequestDto
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

type ApiDeleteAgentRequest struct {
	ctx context.Context
	ApiService *AIAgentsAPIService
	id int32
	deleteRoomRequest *DeleteRoomRequest
}

// The parameters for deleting a room.
func (r ApiDeleteAgentRequest) DeleteRoomRequest(deleteRoomRequest DeleteRoomRequest) ApiDeleteAgentRequest {	r.deleteRoomRequest = &deleteRoomRequest
	return r
}

func (r ApiDeleteAgentRequest) Execute() (*FileOperationWrapper, *http.Response, error) {
	return r.ApiService.DeleteAgentExecute(r)
}

// DeleteAgent Remove an ai agent
//
// Removes an ai agent.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-agent/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The room ID.
// @return ApiDeleteAgentRequest
func (a *AIAgentsAPIService) DeleteAgent(ctx context.Context, id int32) ApiDeleteAgentRequest {
	return ApiDeleteAgentRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return FileOperationWrapper
func (a *AIAgentsAPIService) DeleteAgentExecute(r ApiDeleteAgentRequest) (*FileOperationWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileOperationWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAgentsAPIService.DeleteAgent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/agents/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deleteRoomRequest == nil {
		return localVarReturnValue, nil, reportError("deleteRoomRequest is required and must be specified")
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
	localVarPostBody = r.deleteRoomRequest
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

type ApiGetAgentInfoRequest struct {
	ctx context.Context
	ApiService *AIAgentsAPIService
	id int32
}

func (r ApiGetAgentInfoRequest) Execute() (*FolderIntegerWrapper, *http.Response, error) {
	return r.ApiService.GetAgentInfoExecute(r)
}

// GetAgentInfo Return an ai agent
//
// Returns an ai agent.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-agent-info/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The room ID.
// @return ApiGetAgentInfoRequest
func (a *AIAgentsAPIService) GetAgentInfo(ctx context.Context, id int32) ApiGetAgentInfoRequest {
	return ApiGetAgentInfoRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return FolderIntegerWrapper
func (a *AIAgentsAPIService) GetAgentInfoExecute(r ApiGetAgentInfoRequest) (*FolderIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FolderIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAgentsAPIService.GetAgentInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/agents/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ApiGetAgentsRequest struct {
	ctx context.Context
	ApiService *AIAgentsAPIService
	subjectId *string
	subjectOwnerId *string
	withoutTags *bool
	tags *string
	excludeSubject *bool
	subjectFilter *SubjectFilter
	quotaFilter *QuotaFilter
	count *int32
	startIndex *int32
	sortBy *string
	sortOrder *SortOrder
	filterValue *string
}

// The filter by user ID.
func (r ApiGetAgentsRequest) SubjectId(subjectId string) ApiGetAgentsRequest {	r.subjectId = &subjectId
	return r
}

// The filter by room owner ID.
func (r ApiGetAgentsRequest) SubjectOwnerId(subjectOwnerId string) ApiGetAgentsRequest {	r.subjectOwnerId = &subjectOwnerId
	return r
}

// Specifies whether to search by tags or not.
func (r ApiGetAgentsRequest) WithoutTags(withoutTags bool) ApiGetAgentsRequest {	r.withoutTags = &withoutTags
	return r
}

// The tags in the serialized format.
func (r ApiGetAgentsRequest) Tags(tags string) ApiGetAgentsRequest {	r.tags = &tags
	return r
}

// Specifies whether to exclude search by user or group ID.
func (r ApiGetAgentsRequest) ExcludeSubject(excludeSubject bool) ApiGetAgentsRequest {	r.excludeSubject = &excludeSubject
	return r
}

// The filter by user (Owner - 0, Member - 1).
func (r ApiGetAgentsRequest) SubjectFilter(subjectFilter SubjectFilter) ApiGetAgentsRequest {	r.subjectFilter = &subjectFilter
	return r
}

// The filter by quota (All - 0, Default - 1, Custom - 2).
func (r ApiGetAgentsRequest) QuotaFilter(quotaFilter QuotaFilter) ApiGetAgentsRequest {	r.quotaFilter = &quotaFilter
	return r
}

// Specifies the maximum number of items to retrieve.
func (r ApiGetAgentsRequest) Count(count int32) ApiGetAgentsRequest {	r.count = &count
	return r
}

// The index from which to start retrieving the room content.
func (r ApiGetAgentsRequest) StartIndex(startIndex int32) ApiGetAgentsRequest {	r.startIndex = &startIndex
	return r
}

// Specifies the field by which the room content should be sorted.
func (r ApiGetAgentsRequest) SortBy(sortBy string) ApiGetAgentsRequest {	r.sortBy = &sortBy
	return r
}

// The order in which the results are sorted.
func (r ApiGetAgentsRequest) SortOrder(sortOrder SortOrder) ApiGetAgentsRequest {	r.sortOrder = &sortOrder
	return r
}

// The text filter value used to refine search or query operations.
func (r ApiGetAgentsRequest) FilterValue(filterValue string) ApiGetAgentsRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetAgentsRequest) Execute() (*FolderContentIntegerWrapper, *http.Response, error) {
	return r.ApiService.GetAgentsExecute(r)
}

// GetAgents Get ai agents
//
// Get ai agents
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-agents/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetAgentsRequest
func (a *AIAgentsAPIService) GetAgents(ctx context.Context) ApiGetAgentsRequest {
	return ApiGetAgentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FolderContentIntegerWrapper
func (a *AIAgentsAPIService) GetAgentsExecute(r ApiGetAgentsRequest) (*FolderContentIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FolderContentIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAgentsAPIService.GetAgents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/agents"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.subjectId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subjectId", r.subjectId, "form", "")
			}
	if r.subjectOwnerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subjectOwnerId", r.subjectOwnerId, "form", "")
			}
	if r.withoutTags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "withoutTags", r.withoutTags, "form", "")
			}
	if r.tags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tags", r.tags, "form", "")
			}
	if r.excludeSubject != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "excludeSubject", r.excludeSubject, "form", "")
			}
	if r.subjectFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subjectFilter", r.subjectFilter, "form", "")
			}
	if r.quotaFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "quotaFilter", r.quotaFilter, "form", "")
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

type ApiGetAgentsNewItemsRequest struct {
	ctx context.Context
	ApiService *AIAgentsAPIService
}

func (r ApiGetAgentsNewItemsRequest) Execute() (*NewItemsAgentNewItemsArrayWrapper, *http.Response, error) {
	return r.ApiService.GetAgentsNewItemsExecute(r)
}

// GetAgentsNewItems Get the room new items
//
// Returns the room new items.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-agents-new-items/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetAgentsNewItemsRequest
func (a *AIAgentsAPIService) GetAgentsNewItems(ctx context.Context) ApiGetAgentsNewItemsRequest {
	return ApiGetAgentsNewItemsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NewItemsAgentNewItemsArrayWrapper
func (a *AIAgentsAPIService) GetAgentsNewItemsExecute(r ApiGetAgentsNewItemsRequest) (*NewItemsAgentNewItemsArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NewItemsAgentNewItemsArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAgentsAPIService.GetAgentsNewItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/agents/news"

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

type ApiResetAgentsQuotaRequest struct {
	ctx context.Context
	ApiService *AIAgentsAPIService
	updateRoomsRoomIdsRequestDtoInteger *UpdateRoomsRoomIdsRequestDtoInteger
}

func (r ApiResetAgentsQuotaRequest) UpdateRoomsRoomIdsRequestDtoInteger(updateRoomsRoomIdsRequestDtoInteger UpdateRoomsRoomIdsRequestDtoInteger) ApiResetAgentsQuotaRequest {	r.updateRoomsRoomIdsRequestDtoInteger = &updateRoomsRoomIdsRequestDtoInteger
	return r
}

func (r ApiResetAgentsQuotaRequest) Execute() (*FolderIntegerArrayWrapper, *http.Response, error) {
	return r.ApiService.ResetAgentsQuotaExecute(r)
}

// ResetAgentsQuota Reset the AI agents quota limit
//
// Resets the quota limit for the AI agents with the IDs specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/reset-agents-quota/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiResetAgentsQuotaRequest
func (a *AIAgentsAPIService) ResetAgentsQuota(ctx context.Context) ApiResetAgentsQuotaRequest {
	return ApiResetAgentsQuotaRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FolderIntegerArrayWrapper
func (a *AIAgentsAPIService) ResetAgentsQuotaExecute(r ApiResetAgentsQuotaRequest) (*FolderIntegerArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FolderIntegerArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAgentsAPIService.ResetAgentsQuota")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/agents/resetquota"

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
	localVarPostBody = r.updateRoomsRoomIdsRequestDtoInteger
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

type ApiUpdateAgentRequest struct {
	ctx context.Context
	ApiService *AIAgentsAPIService
	id int32
	updateRoomRequest *UpdateRoomRequest
}

// The request parameters for updating a room.
func (r ApiUpdateAgentRequest) UpdateRoomRequest(updateRoomRequest UpdateRoomRequest) ApiUpdateAgentRequest {	r.updateRoomRequest = &updateRoomRequest
	return r
}

func (r ApiUpdateAgentRequest) Execute() (*FolderIntegerWrapper, *http.Response, error) {
	return r.ApiService.UpdateAgentExecute(r)
}

// UpdateAgent Update an ai agent
//
// Updates an ai agent.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/update-agent/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The room ID.
// @return ApiUpdateAgentRequest
func (a *AIAgentsAPIService) UpdateAgent(ctx context.Context, id int32) ApiUpdateAgentRequest {
	return ApiUpdateAgentRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return FolderIntegerWrapper
func (a *AIAgentsAPIService) UpdateAgentExecute(r ApiUpdateAgentRequest) (*FolderIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FolderIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAgentsAPIService.UpdateAgent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/agents/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateRoomRequest == nil {
		return localVarReturnValue, nil, reportError("updateRoomRequest is required and must be specified")
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
	localVarPostBody = r.updateRoomRequest
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

type ApiUpdateAgentsQuotaRequest struct {
	ctx context.Context
	ApiService *AIAgentsAPIService
	updateRoomsQuotaRequestDtoInteger *UpdateRoomsQuotaRequestDtoInteger
}

func (r ApiUpdateAgentsQuotaRequest) UpdateRoomsQuotaRequestDtoInteger(updateRoomsQuotaRequestDtoInteger UpdateRoomsQuotaRequestDtoInteger) ApiUpdateAgentsQuotaRequest {	r.updateRoomsQuotaRequestDtoInteger = &updateRoomsQuotaRequestDtoInteger
	return r
}

func (r ApiUpdateAgentsQuotaRequest) Execute() (*FolderIntegerArrayWrapper, *http.Response, error) {
	return r.ApiService.UpdateAgentsQuotaExecute(r)
}

// UpdateAgentsQuota Change the AI agent quota limit
//
// Changes the quota limit for the AI agents with the IDs specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/update-agents-quota/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiUpdateAgentsQuotaRequest
func (a *AIAgentsAPIService) UpdateAgentsQuota(ctx context.Context) ApiUpdateAgentsQuotaRequest {
	return ApiUpdateAgentsQuotaRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FolderIntegerArrayWrapper
func (a *AIAgentsAPIService) UpdateAgentsQuotaExecute(r ApiUpdateAgentsQuotaRequest) (*FolderIntegerArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FolderIntegerArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAgentsAPIService.UpdateAgentsQuota")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/agents/agentquota"

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
	localVarPostBody = r.updateRoomsQuotaRequestDtoInteger
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
