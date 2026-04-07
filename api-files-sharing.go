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


// FilesSharingAPIService FilesSharingAPI service
type FilesSharingAPIService service

type ApiApplyExternalSharePasswordRequest struct {
	ctx context.Context
	ApiService *FilesSharingAPIService
	key string
	externalShareRequestParam *ExternalShareRequestParam
}

// The external data share request parameters.
func (r ApiApplyExternalSharePasswordRequest) ExternalShareRequestParam(externalShareRequestParam ExternalShareRequestParam) ApiApplyExternalSharePasswordRequest {	r.externalShareRequestParam = &externalShareRequestParam
	return r
}

func (r ApiApplyExternalSharePasswordRequest) Execute() (*ExternalShareWrapper, *http.Response, error) {
	return r.ApiService.ApplyExternalSharePasswordExecute(r)
}

// ApplyExternalSharePassword Apply external data password
//
// Applies a password specified in the request to get the external data.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/apply-external-share-password/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param key The unique document identifier.
// @return ApiApplyExternalSharePasswordRequest
func (a *FilesSharingAPIService) ApplyExternalSharePassword(ctx context.Context, key string) ApiApplyExternalSharePasswordRequest {
	return ApiApplyExternalSharePasswordRequest{
		ApiService: a,
		ctx: ctx,
		key: key,
	}
}

// Execute executes the request
//  @return ExternalShareWrapper
func (a *FilesSharingAPIService) ApplyExternalSharePasswordExecute(r ApiApplyExternalSharePasswordRequest) (*ExternalShareWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExternalShareWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingAPIService.ApplyExternalSharePassword")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/share/{key}/password"
	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", url.PathEscape(parameterValueToString(r.key, "key")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.externalShareRequestParam == nil {
		return localVarReturnValue, nil, reportError("externalShareRequestParam is required and must be specified")
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
	localVarPostBody = r.externalShareRequestParam
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

type ApiChangeFileOwnerRequest struct {
	ctx context.Context
	ApiService *FilesSharingAPIService
	changeOwnerRequestDto *ChangeOwnerRequestDto
}

func (r ApiChangeFileOwnerRequest) ChangeOwnerRequestDto(changeOwnerRequestDto ChangeOwnerRequestDto) ApiChangeFileOwnerRequest {	r.changeOwnerRequestDto = &changeOwnerRequestDto
	return r
}

func (r ApiChangeFileOwnerRequest) Execute() (*FileEntryBaseArrayWrapper, *http.Response, error) {
	return r.ApiService.ChangeFileOwnerExecute(r)
}

// ChangeFileOwner Change the file owner
//
// Changes the owner of the file with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/change-file-owner/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiChangeFileOwnerRequest
func (a *FilesSharingAPIService) ChangeFileOwner(ctx context.Context) ApiChangeFileOwnerRequest {
	return ApiChangeFileOwnerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FileEntryBaseArrayWrapper
func (a *FilesSharingAPIService) ChangeFileOwnerExecute(r ApiChangeFileOwnerRequest) (*FileEntryBaseArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileEntryBaseArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingAPIService.ChangeFileOwner")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/owner"

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
	localVarPostBody = r.changeOwnerRequestDto
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

type ApiGetExternalShareDataRequest struct {
	ctx context.Context
	ApiService *FilesSharingAPIService
	key string
	fileId *string
	folderId *string
}

// The unique document identifier.
func (r ApiGetExternalShareDataRequest) FileId(fileId string) ApiGetExternalShareDataRequest {	r.fileId = &fileId
	return r
}

// The unique folder identifier.
func (r ApiGetExternalShareDataRequest) FolderId(folderId string) ApiGetExternalShareDataRequest {	r.folderId = &folderId
	return r
}

func (r ApiGetExternalShareDataRequest) Execute() (*ExternalShareWrapper, *http.Response, error) {
	return r.ApiService.GetExternalShareDataExecute(r)
}

// GetExternalShareData Get the external data
//
// Returns the external data by the key specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-external-share-data/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param key The unique key of the external shared data.
// @return ApiGetExternalShareDataRequest
func (a *FilesSharingAPIService) GetExternalShareData(ctx context.Context, key string) ApiGetExternalShareDataRequest {
	return ApiGetExternalShareDataRequest{
		ApiService: a,
		ctx: ctx,
		key: key,
	}
}

// Execute executes the request
//  @return ExternalShareWrapper
func (a *FilesSharingAPIService) GetExternalShareDataExecute(r ApiGetExternalShareDataRequest) (*ExternalShareWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExternalShareWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingAPIService.GetExternalShareData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/share/{key}"
	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", url.PathEscape(parameterValueToString(r.key, "key")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fileId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fileId", r.fileId, "form", "")
			}
	if r.folderId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "folderId", r.folderId, "form", "")
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

type ApiGetFileSecurityInfoRequest struct {
	ctx context.Context
	ApiService *FilesSharingAPIService
	id int32
	count *int32
	startIndex *int32
}

// The number of items to retrieve in the request.
func (r ApiGetFileSecurityInfoRequest) Count(count int32) ApiGetFileSecurityInfoRequest {	r.count = &count
	return r
}

// The starting index for the query results.
func (r ApiGetFileSecurityInfoRequest) StartIndex(startIndex int32) ApiGetFileSecurityInfoRequest {	r.startIndex = &startIndex
	return r
}

func (r ApiGetFileSecurityInfoRequest) Execute() (*FileShareArrayWrapper, *http.Response, error) {
	return r.ApiService.GetFileSecurityInfoExecute(r)
}

// GetFileSecurityInfo Get the shared file information
//
// Returns the detailed information about the shared file with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-file-security-info/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The file unique identifier.
// @return ApiGetFileSecurityInfoRequest
func (a *FilesSharingAPIService) GetFileSecurityInfo(ctx context.Context, id int32) ApiGetFileSecurityInfoRequest {
	return ApiGetFileSecurityInfoRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return FileShareArrayWrapper
func (a *FilesSharingAPIService) GetFileSecurityInfoExecute(r ApiGetFileSecurityInfoRequest) (*FileShareArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileShareArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingAPIService.GetFileSecurityInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/file/{id}/share"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
			}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
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

type ApiGetFolderSecurityInfoRequest struct {
	ctx context.Context
	ApiService *FilesSharingAPIService
	id int32
	count *int32
	startIndex *int32
}

// The number of items to retrieve in the request.
func (r ApiGetFolderSecurityInfoRequest) Count(count int32) ApiGetFolderSecurityInfoRequest {	r.count = &count
	return r
}

// The starting index for the query results.
func (r ApiGetFolderSecurityInfoRequest) StartIndex(startIndex int32) ApiGetFolderSecurityInfoRequest {	r.startIndex = &startIndex
	return r
}

func (r ApiGetFolderSecurityInfoRequest) Execute() (*FileShareArrayWrapper, *http.Response, error) {
	return r.ApiService.GetFolderSecurityInfoExecute(r)
}

// GetFolderSecurityInfo Get the shared folder information
//
// Returns the detailed information about the shared folder with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-folder-security-info/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The folder unique identifier.
// @return ApiGetFolderSecurityInfoRequest
func (a *FilesSharingAPIService) GetFolderSecurityInfo(ctx context.Context, id int32) ApiGetFolderSecurityInfoRequest {
	return ApiGetFolderSecurityInfoRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return FileShareArrayWrapper
func (a *FilesSharingAPIService) GetFolderSecurityInfoExecute(r ApiGetFolderSecurityInfoRequest) (*FileShareArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileShareArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingAPIService.GetFolderSecurityInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/folder/{id}/share"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
			}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
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

type ApiGetGroupsMembersWithFileSecurityRequest struct {
	ctx context.Context
	ApiService *FilesSharingAPIService
	fileId int32
	groupId string
	count *int32
	startIndex *int32
	filterValue *string
}

// The number of items to be retrieved in the current query.
func (r ApiGetGroupsMembersWithFileSecurityRequest) Count(count int32) ApiGetGroupsMembersWithFileSecurityRequest {	r.count = &count
	return r
}

// The starting index for the query result set.
func (r ApiGetGroupsMembersWithFileSecurityRequest) StartIndex(startIndex int32) ApiGetGroupsMembersWithFileSecurityRequest {	r.startIndex = &startIndex
	return r
}

// The filter value used for searching or querying group members based on text input.
func (r ApiGetGroupsMembersWithFileSecurityRequest) FilterValue(filterValue string) ApiGetGroupsMembersWithFileSecurityRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetGroupsMembersWithFileSecurityRequest) Execute() (*GroupMemberSecurityRequestArrayWrapper, *http.Response, error) {
	return r.ApiService.GetGroupsMembersWithFileSecurityExecute(r)
}

// GetGroupsMembersWithFileSecurity Get file group members with security information
//
// Returns the group members with their file security information.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-groups-members-with-file-security/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param fileId The file ID.
// @param groupId The group ID.
// @return ApiGetGroupsMembersWithFileSecurityRequest
func (a *FilesSharingAPIService) GetGroupsMembersWithFileSecurity(ctx context.Context, fileId int32, groupId string) ApiGetGroupsMembersWithFileSecurityRequest {
	return ApiGetGroupsMembersWithFileSecurityRequest{
		ApiService: a,
		ctx: ctx,
		fileId: fileId,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return GroupMemberSecurityRequestArrayWrapper
func (a *FilesSharingAPIService) GetGroupsMembersWithFileSecurityExecute(r ApiGetGroupsMembersWithFileSecurityRequest) (*GroupMemberSecurityRequestArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupMemberSecurityRequestArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingAPIService.GetGroupsMembersWithFileSecurity")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/file/{fileId}/group/{groupId}/share"
	localVarPath = strings.Replace(localVarPath, "{"+"fileId"+"}", url.PathEscape(parameterValueToString(r.fileId, "fileId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetGroupsMembersWithFolderSecurityRequest struct {
	ctx context.Context
	ApiService *FilesSharingAPIService
	folderId int32
	groupId string
	count *int32
	startIndex *int32
	filterValue *string
}

// The number of items to be retrieved in the current query.
func (r ApiGetGroupsMembersWithFolderSecurityRequest) Count(count int32) ApiGetGroupsMembersWithFolderSecurityRequest {	r.count = &count
	return r
}

// The starting index for the query result set.
func (r ApiGetGroupsMembersWithFolderSecurityRequest) StartIndex(startIndex int32) ApiGetGroupsMembersWithFolderSecurityRequest {	r.startIndex = &startIndex
	return r
}

// The filter value used for searching or querying group members based on text input.
func (r ApiGetGroupsMembersWithFolderSecurityRequest) FilterValue(filterValue string) ApiGetGroupsMembersWithFolderSecurityRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetGroupsMembersWithFolderSecurityRequest) Execute() (*GroupMemberSecurityRequestArrayWrapper, *http.Response, error) {
	return r.ApiService.GetGroupsMembersWithFolderSecurityExecute(r)
}

// GetGroupsMembersWithFolderSecurity Get folder group members with security information
//
// Returns the group members with their folder security information.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-groups-members-with-folder-security/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The folder ID.
// @param groupId The group ID.
// @return ApiGetGroupsMembersWithFolderSecurityRequest
func (a *FilesSharingAPIService) GetGroupsMembersWithFolderSecurity(ctx context.Context, folderId int32, groupId string) ApiGetGroupsMembersWithFolderSecurityRequest {
	return ApiGetGroupsMembersWithFolderSecurityRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return GroupMemberSecurityRequestArrayWrapper
func (a *FilesSharingAPIService) GetGroupsMembersWithFolderSecurityExecute(r ApiGetGroupsMembersWithFolderSecurityRequest) (*GroupMemberSecurityRequestArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupMemberSecurityRequestArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingAPIService.GetGroupsMembersWithFolderSecurity")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/folder/{folderId}/group/{groupId}/share"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiGetSecurityInfoRequest struct {
	ctx context.Context
	ApiService *FilesSharingAPIService
	baseBatchRequestDto *BaseBatchRequestDto
}

func (r ApiGetSecurityInfoRequest) BaseBatchRequestDto(baseBatchRequestDto BaseBatchRequestDto) ApiGetSecurityInfoRequest {	r.baseBatchRequestDto = &baseBatchRequestDto
	return r
}

func (r ApiGetSecurityInfoRequest) Execute() (*FileShareArrayWrapper, *http.Response, error) {
	return r.ApiService.GetSecurityInfoExecute(r)
}

// GetSecurityInfo Get the sharing rights
//
// Returns the sharing rights for all the files and folders specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-security-info/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetSecurityInfoRequest
func (a *FilesSharingAPIService) GetSecurityInfo(ctx context.Context) ApiGetSecurityInfoRequest {
	return ApiGetSecurityInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FileShareArrayWrapper
func (a *FilesSharingAPIService) GetSecurityInfoExecute(r ApiGetSecurityInfoRequest) (*FileShareArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileShareArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingAPIService.GetSecurityInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/share"

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
	localVarPostBody = r.baseBatchRequestDto
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

type ApiGetSharedUsersRequest struct {
	ctx context.Context
	ApiService *FilesSharingAPIService
	fileId int32
}

func (r ApiGetSharedUsersRequest) Execute() (*MentionWrapperArrayWrapper, *http.Response, error) {
	return r.ApiService.GetSharedUsersExecute(r)
}

// GetSharedUsers Get user access rights by file ID
//
// Returns a list of users with their access rights to the file with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-shared-users/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param fileId The file unique identifier.
// @return ApiGetSharedUsersRequest
func (a *FilesSharingAPIService) GetSharedUsers(ctx context.Context, fileId int32) ApiGetSharedUsersRequest {
	return ApiGetSharedUsersRequest{
		ApiService: a,
		ctx: ctx,
		fileId: fileId,
	}
}

// Execute executes the request
//  @return MentionWrapperArrayWrapper
func (a *FilesSharingAPIService) GetSharedUsersExecute(r ApiGetSharedUsersRequest) (*MentionWrapperArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MentionWrapperArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingAPIService.GetSharedUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/file/{fileId}/sharedusers"
	localVarPath = strings.Replace(localVarPath, "{"+"fileId"+"}", url.PathEscape(parameterValueToString(r.fileId, "fileId")), -1)

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

type ApiRemoveSecurityInfoRequest struct {
	ctx context.Context
	ApiService *FilesSharingAPIService
	baseBatchRequestDto *BaseBatchRequestDto
}

func (r ApiRemoveSecurityInfoRequest) BaseBatchRequestDto(baseBatchRequestDto BaseBatchRequestDto) ApiRemoveSecurityInfoRequest {	r.baseBatchRequestDto = &baseBatchRequestDto
	return r
}

func (r ApiRemoveSecurityInfoRequest) Execute() (*BooleanWrapper, *http.Response, error) {
	return r.ApiService.RemoveSecurityInfoExecute(r)
}

// RemoveSecurityInfo Remove the sharing rights
//
// Removes the sharing rights from all the files and folders specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/remove-security-info/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiRemoveSecurityInfoRequest
func (a *FilesSharingAPIService) RemoveSecurityInfo(ctx context.Context) ApiRemoveSecurityInfoRequest {
	return ApiRemoveSecurityInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BooleanWrapper
func (a *FilesSharingAPIService) RemoveSecurityInfoExecute(r ApiRemoveSecurityInfoRequest) (*BooleanWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BooleanWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingAPIService.RemoveSecurityInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/share"

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
	localVarPostBody = r.baseBatchRequestDto
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

type ApiSendEditorNotifyRequest struct {
	ctx context.Context
	ApiService *FilesSharingAPIService
	fileId int32
	mentionMessageWrapper *MentionMessageWrapper
}

// The mention message.
func (r ApiSendEditorNotifyRequest) MentionMessageWrapper(mentionMessageWrapper MentionMessageWrapper) ApiSendEditorNotifyRequest {	r.mentionMessageWrapper = &mentionMessageWrapper
	return r
}

func (r ApiSendEditorNotifyRequest) Execute() (*AceShortWrapperArrayWrapper, *http.Response, error) {
	return r.ApiService.SendEditorNotifyExecute(r)
}

// SendEditorNotify Send the mention message
//
// Sends a message to the users who are mentioned in the file with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/send-editor-notify/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param fileId The file ID with the mention message.
// @return ApiSendEditorNotifyRequest
func (a *FilesSharingAPIService) SendEditorNotify(ctx context.Context, fileId int32) ApiSendEditorNotifyRequest {
	return ApiSendEditorNotifyRequest{
		ApiService: a,
		ctx: ctx,
		fileId: fileId,
	}
}

// Execute executes the request
//  @return AceShortWrapperArrayWrapper
func (a *FilesSharingAPIService) SendEditorNotifyExecute(r ApiSendEditorNotifyRequest) (*AceShortWrapperArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AceShortWrapperArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingAPIService.SendEditorNotify")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/file/{fileId}/sendeditornotify"
	localVarPath = strings.Replace(localVarPath, "{"+"fileId"+"}", url.PathEscape(parameterValueToString(r.fileId, "fileId")), -1)

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
	localVarPostBody = r.mentionMessageWrapper
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

type ApiSetFileSecurityInfoRequest struct {
	ctx context.Context
	ApiService *FilesSharingAPIService
	fileId int32
	securityInfoSimpleRequestDto *SecurityInfoSimpleRequestDto
}

// The parameters of the security information simple request.
func (r ApiSetFileSecurityInfoRequest) SecurityInfoSimpleRequestDto(securityInfoSimpleRequestDto SecurityInfoSimpleRequestDto) ApiSetFileSecurityInfoRequest {	r.securityInfoSimpleRequestDto = &securityInfoSimpleRequestDto
	return r
}

func (r ApiSetFileSecurityInfoRequest) Execute() (*FileShareArrayWrapper, *http.Response, error) {
	return r.ApiService.SetFileSecurityInfoExecute(r)
}

// SetFileSecurityInfo Share a file
//
// Sets the sharing settings to a file with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/set-file-security-info/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param fileId The file ID.
// @return ApiSetFileSecurityInfoRequest
func (a *FilesSharingAPIService) SetFileSecurityInfo(ctx context.Context, fileId int32) ApiSetFileSecurityInfoRequest {
	return ApiSetFileSecurityInfoRequest{
		ApiService: a,
		ctx: ctx,
		fileId: fileId,
	}
}

// Execute executes the request
//  @return FileShareArrayWrapper
func (a *FilesSharingAPIService) SetFileSecurityInfoExecute(r ApiSetFileSecurityInfoRequest) (*FileShareArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileShareArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingAPIService.SetFileSecurityInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/file/{fileId}/share"
	localVarPath = strings.Replace(localVarPath, "{"+"fileId"+"}", url.PathEscape(parameterValueToString(r.fileId, "fileId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.securityInfoSimpleRequestDto == nil {
		return localVarReturnValue, nil, reportError("securityInfoSimpleRequestDto is required and must be specified")
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
	localVarPostBody = r.securityInfoSimpleRequestDto
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

type ApiSetFolderSecurityInfoRequest struct {
	ctx context.Context
	ApiService *FilesSharingAPIService
	folderId int32
	securityInfoSimpleRequestDto *SecurityInfoSimpleRequestDto
}

// The parameters of the security information simple request.
func (r ApiSetFolderSecurityInfoRequest) SecurityInfoSimpleRequestDto(securityInfoSimpleRequestDto SecurityInfoSimpleRequestDto) ApiSetFolderSecurityInfoRequest {	r.securityInfoSimpleRequestDto = &securityInfoSimpleRequestDto
	return r
}

func (r ApiSetFolderSecurityInfoRequest) Execute() (*FileShareArrayWrapper, *http.Response, error) {
	return r.ApiService.SetFolderSecurityInfoExecute(r)
}

// SetFolderSecurityInfo Share a folder
//
// Sets the sharing settings to a folder with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/set-folder-security-info/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The folder ID.
// @return ApiSetFolderSecurityInfoRequest
func (a *FilesSharingAPIService) SetFolderSecurityInfo(ctx context.Context, folderId int32) ApiSetFolderSecurityInfoRequest {
	return ApiSetFolderSecurityInfoRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return FileShareArrayWrapper
func (a *FilesSharingAPIService) SetFolderSecurityInfoExecute(r ApiSetFolderSecurityInfoRequest) (*FileShareArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileShareArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingAPIService.SetFolderSecurityInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/folder/{folderId}/share"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.securityInfoSimpleRequestDto == nil {
		return localVarReturnValue, nil, reportError("securityInfoSimpleRequestDto is required and must be specified")
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
	localVarPostBody = r.securityInfoSimpleRequestDto
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

type ApiSetSecurityInfoRequest struct {
	ctx context.Context
	ApiService *FilesSharingAPIService
	securityInfoRequestDto *SecurityInfoRequestDto
}

func (r ApiSetSecurityInfoRequest) SecurityInfoRequestDto(securityInfoRequestDto SecurityInfoRequestDto) ApiSetSecurityInfoRequest {	r.securityInfoRequestDto = &securityInfoRequestDto
	return r
}

func (r ApiSetSecurityInfoRequest) Execute() (*FileShareArrayWrapper, *http.Response, error) {
	return r.ApiService.SetSecurityInfoExecute(r)
}

// SetSecurityInfo Set the sharing rights
//
// Sets the sharing rights to all the files and folders specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/set-security-info/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiSetSecurityInfoRequest
func (a *FilesSharingAPIService) SetSecurityInfo(ctx context.Context) ApiSetSecurityInfoRequest {
	return ApiSetSecurityInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FileShareArrayWrapper
func (a *FilesSharingAPIService) SetSecurityInfoExecute(r ApiSetSecurityInfoRequest) (*FileShareArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileShareArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingAPIService.SetSecurityInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/share"

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
	localVarPostBody = r.securityInfoRequestDto
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
