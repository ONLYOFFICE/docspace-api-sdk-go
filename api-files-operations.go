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
	"os"
)


// FilesOperationsAPIService FilesOperationsAPI service
type FilesOperationsAPIService service

type ApiAbortUploadSessionRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	sessionId string
	folderId int32
}

func (r ApiAbortUploadSessionRequest) Execute() (*http.Response, error) {
	return r.ApiService.AbortUploadSessionExecute(r)
}

// AbortUploadSession Aborts an in-progress file upload session.
//
// This method allows users to cancel an ongoing upload session identified by the session ID.  Once the session is aborted, the associated resources will be cleaned up, and the session will no longer accept further uploads.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/abort-upload-session/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param sessionId The session ID.
// @param folderId The folder ID.
// @return ApiAbortUploadSessionRequest
func (a *FilesOperationsAPIService) AbortUploadSession(ctx context.Context, sessionId string, folderId int32) ApiAbortUploadSessionRequest {
	return ApiAbortUploadSessionRequest{
		ApiService: a,
		ctx: ctx,
		sessionId: sessionId,
		folderId: folderId,
	}
}

// Execute executes the request
func (a *FilesOperationsAPIService) AbortUploadSessionExecute(r ApiAbortUploadSessionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.AbortUploadSession")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/{folderId}/session/{sessionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", url.PathEscape(parameterValueToString(r.sessionId, "sessionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiAddFavoritesRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	baseBatchRequestDto *BaseBatchRequestDto
}

func (r ApiAddFavoritesRequest) BaseBatchRequestDto(baseBatchRequestDto BaseBatchRequestDto) ApiAddFavoritesRequest {	r.baseBatchRequestDto = &baseBatchRequestDto
	return r
}

func (r ApiAddFavoritesRequest) Execute() (*BooleanWrapper, *http.Response, error) {
	return r.ApiService.AddFavoritesExecute(r)
}

// AddFavorites Add favorite files and folders
//
// Adds files and folders with the IDs specified in the request to the favorite list.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/add-favorites/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAddFavoritesRequest
func (a *FilesOperationsAPIService) AddFavorites(ctx context.Context) ApiAddFavoritesRequest {
	return ApiAddFavoritesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BooleanWrapper
func (a *FilesOperationsAPIService) AddFavoritesExecute(r ApiAddFavoritesRequest) (*BooleanWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BooleanWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.AddFavorites")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/favorites"

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

type ApiBulkDownloadRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	downloadRequestDto *DownloadRequestDto
}

func (r ApiBulkDownloadRequest) DownloadRequestDto(downloadRequestDto DownloadRequestDto) ApiBulkDownloadRequest {	r.downloadRequestDto = &downloadRequestDto
	return r
}

func (r ApiBulkDownloadRequest) Execute() (*FileOperationArrayWrapper, *http.Response, error) {
	return r.ApiService.BulkDownloadExecute(r)
}

// BulkDownload Bulk download
//
// Starts the download process of files and folders with the IDs specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/bulk-download/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiBulkDownloadRequest
func (a *FilesOperationsAPIService) BulkDownload(ctx context.Context) ApiBulkDownloadRequest {
	return ApiBulkDownloadRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FileOperationArrayWrapper
func (a *FilesOperationsAPIService) BulkDownloadExecute(r ApiBulkDownloadRequest) (*FileOperationArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileOperationArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.BulkDownload")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/fileops/bulkdownload"

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
	localVarPostBody = r.downloadRequestDto
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

type ApiCheckConversionStatusRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	fileId int32
	start *bool
}

// Specifies whether a conversion operation is started or not.
func (r ApiCheckConversionStatusRequest) Start(start bool) ApiCheckConversionStatusRequest {	r.start = &start
	return r
}

func (r ApiCheckConversionStatusRequest) Execute() (*ConversationResultArrayWrapper, *http.Response, error) {
	return r.ApiService.CheckConversionStatusExecute(r)
}

// CheckConversionStatus Get conversion status
//
// Checks the conversion status of a file with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/check-conversion-status/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param fileId The file ID to check conversion status.
// @return ApiCheckConversionStatusRequest
func (a *FilesOperationsAPIService) CheckConversionStatus(ctx context.Context, fileId int32) ApiCheckConversionStatusRequest {
	return ApiCheckConversionStatusRequest{
		ApiService: a,
		ctx: ctx,
		fileId: fileId,
	}
}

// Execute executes the request
//  @return ConversationResultArrayWrapper
func (a *FilesOperationsAPIService) CheckConversionStatusExecute(r ApiCheckConversionStatusRequest) (*ConversationResultArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConversationResultArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.CheckConversionStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/file/{fileId}/checkconversion"
	localVarPath = strings.Replace(localVarPath, "{"+"fileId"+"}", url.PathEscape(parameterValueToString(r.fileId, "fileId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "form", "")
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

type ApiCheckMoveOrCopyBatchItemsRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	inDto *BatchRequestDto
}

// The request parameters for copying/moving files.
func (r ApiCheckMoveOrCopyBatchItemsRequest) InDto(inDto BatchRequestDto) ApiCheckMoveOrCopyBatchItemsRequest {	r.inDto = &inDto
	return r
}

func (r ApiCheckMoveOrCopyBatchItemsRequest) Execute() (*FileEntryBaseArrayWrapper, *http.Response, error) {
	return r.ApiService.CheckMoveOrCopyBatchItemsExecute(r)
}

// CheckMoveOrCopyBatchItems Move or copy files to a folder
//
// Checks if files or folders can be moved or copied to the specified folder, moves or copies them, and returns their information.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/check-move-or-copy-batch-items/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiCheckMoveOrCopyBatchItemsRequest
func (a *FilesOperationsAPIService) CheckMoveOrCopyBatchItems(ctx context.Context) ApiCheckMoveOrCopyBatchItemsRequest {
	return ApiCheckMoveOrCopyBatchItemsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FileEntryBaseArrayWrapper
func (a *FilesOperationsAPIService) CheckMoveOrCopyBatchItemsExecute(r ApiCheckMoveOrCopyBatchItemsRequest) (*FileEntryBaseArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileEntryBaseArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.CheckMoveOrCopyBatchItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/fileops/move"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.inDto != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "inDto", r.inDto, "form", "")
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

type ApiCheckMoveOrCopyDestFolderRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	inDto *BatchRequestDto
}

// The request parameters for copying/moving files.
func (r ApiCheckMoveOrCopyDestFolderRequest) InDto(inDto BatchRequestDto) ApiCheckMoveOrCopyDestFolderRequest {	r.inDto = &inDto
	return r
}

func (r ApiCheckMoveOrCopyDestFolderRequest) Execute() (*CheckDestFolderWrapper, *http.Response, error) {
	return r.ApiService.CheckMoveOrCopyDestFolderExecute(r)
}

// CheckMoveOrCopyDestFolder Check for moving or copying files to a folder
//
// Checks if files can be moved or copied to the specified folder.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/check-move-or-copy-dest-folder/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiCheckMoveOrCopyDestFolderRequest
func (a *FilesOperationsAPIService) CheckMoveOrCopyDestFolder(ctx context.Context) ApiCheckMoveOrCopyDestFolderRequest {
	return ApiCheckMoveOrCopyDestFolderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CheckDestFolderWrapper
func (a *FilesOperationsAPIService) CheckMoveOrCopyDestFolderExecute(r ApiCheckMoveOrCopyDestFolderRequest) (*CheckDestFolderWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CheckDestFolderWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.CheckMoveOrCopyDestFolder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/fileops/checkdestfolder"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.inDto != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "inDto", r.inDto, "form", "")
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

type ApiCopyBatchItemsRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	batchRequestDto *BatchRequestDto
}

func (r ApiCopyBatchItemsRequest) BatchRequestDto(batchRequestDto BatchRequestDto) ApiCopyBatchItemsRequest {	r.batchRequestDto = &batchRequestDto
	return r
}

func (r ApiCopyBatchItemsRequest) Execute() (*FileOperationArrayWrapper, *http.Response, error) {
	return r.ApiService.CopyBatchItemsExecute(r)
}

// CopyBatchItems Copy to the folder
//
// Copies all the selected files and folders to the folder with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/copy-batch-items/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiCopyBatchItemsRequest
func (a *FilesOperationsAPIService) CopyBatchItems(ctx context.Context) ApiCopyBatchItemsRequest {
	return ApiCopyBatchItemsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FileOperationArrayWrapper
func (a *FilesOperationsAPIService) CopyBatchItemsExecute(r ApiCopyBatchItemsRequest) (*FileOperationArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileOperationArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.CopyBatchItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/fileops/copy"

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
	localVarPostBody = r.batchRequestDto
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

type ApiCreateUploadSessionRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	folderId int32
	sessionRequest *SessionRequest
}

// The session parameters.
func (r ApiCreateUploadSessionRequest) SessionRequest(sessionRequest SessionRequest) ApiCreateUploadSessionRequest {	r.sessionRequest = &sessionRequest
	return r
}

func (r ApiCreateUploadSessionRequest) Execute() (*ChunkedUploadSessionResponseWrapperIntegerWrapper, *http.Response, error) {
	return r.ApiService.CreateUploadSessionExecute(r)
}

// CreateUploadSession Chunked upload
//
// Creates the session to upload large files in multiple chunks to the folder with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/create-upload-session/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The session folder ID.
// @return ApiCreateUploadSessionRequest
//
// Deprecated
func (a *FilesOperationsAPIService) CreateUploadSession(ctx context.Context, folderId int32) ApiCreateUploadSessionRequest {
	return ApiCreateUploadSessionRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return ChunkedUploadSessionResponseWrapperIntegerWrapper
// Deprecated
func (a *FilesOperationsAPIService) CreateUploadSessionExecute(r ApiCreateUploadSessionRequest) (*ChunkedUploadSessionResponseWrapperIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ChunkedUploadSessionResponseWrapperIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.CreateUploadSession")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/{folderId}/upload/create_session"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sessionRequest == nil {
		return localVarReturnValue, nil, reportError("sessionRequest is required and must be specified")
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
	localVarPostBody = r.sessionRequest
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

type ApiCreateUploadSessionInFolderRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	folderId int32
	sessionRequest *SessionRequest
}

// The session parameters.
func (r ApiCreateUploadSessionInFolderRequest) SessionRequest(sessionRequest SessionRequest) ApiCreateUploadSessionInFolderRequest {	r.sessionRequest = &sessionRequest
	return r
}

func (r ApiCreateUploadSessionInFolderRequest) Execute() (*ChunkedUploadSessionResponseIntegerWrapper, *http.Response, error) {
	return r.ApiService.CreateUploadSessionInFolderExecute(r)
}

// CreateUploadSessionInFolder Creates a session for uploading a file to a specific folder in chunks.
//
// The session allows the user to upload a file in smaller chunks to the folder identified by its ID.  The file information, such as name, size, and additional metadata, must be provided in the request.  This method facilitates large file upload scenarios by enabling chunked file uploads.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/create-upload-session-in-folder/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The session folder ID.
// @return ApiCreateUploadSessionInFolderRequest
func (a *FilesOperationsAPIService) CreateUploadSessionInFolder(ctx context.Context, folderId int32) ApiCreateUploadSessionInFolderRequest {
	return ApiCreateUploadSessionInFolderRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return ChunkedUploadSessionResponseIntegerWrapper
func (a *FilesOperationsAPIService) CreateUploadSessionInFolderExecute(r ApiCreateUploadSessionInFolderRequest) (*ChunkedUploadSessionResponseIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ChunkedUploadSessionResponseIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.CreateUploadSessionInFolder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/{folderId}/session"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sessionRequest == nil {
		return localVarReturnValue, nil, reportError("sessionRequest is required and must be specified")
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
	localVarPostBody = r.sessionRequest
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

type ApiDeleteBatchItemsRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	deleteBatchRequestDto *DeleteBatchRequestDto
}

func (r ApiDeleteBatchItemsRequest) DeleteBatchRequestDto(deleteBatchRequestDto DeleteBatchRequestDto) ApiDeleteBatchItemsRequest {	r.deleteBatchRequestDto = &deleteBatchRequestDto
	return r
}

func (r ApiDeleteBatchItemsRequest) Execute() (*FileOperationArrayWrapper, *http.Response, error) {
	return r.ApiService.DeleteBatchItemsExecute(r)
}

// DeleteBatchItems Delete files and folders
//
// Deletes the files and folders with the IDs specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-batch-items/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiDeleteBatchItemsRequest
func (a *FilesOperationsAPIService) DeleteBatchItems(ctx context.Context) ApiDeleteBatchItemsRequest {
	return ApiDeleteBatchItemsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FileOperationArrayWrapper
func (a *FilesOperationsAPIService) DeleteBatchItemsExecute(r ApiDeleteBatchItemsRequest) (*FileOperationArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileOperationArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.DeleteBatchItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/fileops/delete"

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
	localVarPostBody = r.deleteBatchRequestDto
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

type ApiDeleteFavoritesFromBodyRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	baseBatchRequestDto *BaseBatchRequestDto
}

func (r ApiDeleteFavoritesFromBodyRequest) BaseBatchRequestDto(baseBatchRequestDto BaseBatchRequestDto) ApiDeleteFavoritesFromBodyRequest {	r.baseBatchRequestDto = &baseBatchRequestDto
	return r
}

func (r ApiDeleteFavoritesFromBodyRequest) Execute() (*BooleanWrapper, *http.Response, error) {
	return r.ApiService.DeleteFavoritesFromBodyExecute(r)
}

// DeleteFavoritesFromBody Delete favorite files and folders (using body parameters)
//
// Removes files and folders with the IDs specified in the request from the favorite list. This method uses the body parameters.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-favorites-from-body/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiDeleteFavoritesFromBodyRequest
func (a *FilesOperationsAPIService) DeleteFavoritesFromBody(ctx context.Context) ApiDeleteFavoritesFromBodyRequest {
	return ApiDeleteFavoritesFromBodyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BooleanWrapper
func (a *FilesOperationsAPIService) DeleteFavoritesFromBodyExecute(r ApiDeleteFavoritesFromBodyRequest) (*BooleanWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BooleanWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.DeleteFavoritesFromBody")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/favorites"

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

type ApiDeleteFileVersionsRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	deleteVersionBatchRequestDto *DeleteVersionBatchRequestDto
}

func (r ApiDeleteFileVersionsRequest) DeleteVersionBatchRequestDto(deleteVersionBatchRequestDto DeleteVersionBatchRequestDto) ApiDeleteFileVersionsRequest {	r.deleteVersionBatchRequestDto = &deleteVersionBatchRequestDto
	return r
}

func (r ApiDeleteFileVersionsRequest) Execute() (*FileOperationWrapper, *http.Response, error) {
	return r.ApiService.DeleteFileVersionsExecute(r)
}

// DeleteFileVersions Delete file versions
//
// Deletes the file versions with the IDs specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-file-versions/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiDeleteFileVersionsRequest
func (a *FilesOperationsAPIService) DeleteFileVersions(ctx context.Context) ApiDeleteFileVersionsRequest {
	return ApiDeleteFileVersionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FileOperationWrapper
func (a *FilesOperationsAPIService) DeleteFileVersionsExecute(r ApiDeleteFileVersionsRequest) (*FileOperationWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileOperationWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.DeleteFileVersions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/fileops/deleteversion"

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
	localVarPostBody = r.deleteVersionBatchRequestDto
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

type ApiDuplicateBatchItemsRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	duplicateRequestDto *DuplicateRequestDto
}

func (r ApiDuplicateBatchItemsRequest) DuplicateRequestDto(duplicateRequestDto DuplicateRequestDto) ApiDuplicateBatchItemsRequest {	r.duplicateRequestDto = &duplicateRequestDto
	return r
}

func (r ApiDuplicateBatchItemsRequest) Execute() (*FileOperationArrayWrapper, *http.Response, error) {
	return r.ApiService.DuplicateBatchItemsExecute(r)
}

// DuplicateBatchItems Duplicate files and folders
//
// Duplicates all the selected files and folders.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/duplicate-batch-items/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiDuplicateBatchItemsRequest
func (a *FilesOperationsAPIService) DuplicateBatchItems(ctx context.Context) ApiDuplicateBatchItemsRequest {
	return ApiDuplicateBatchItemsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FileOperationArrayWrapper
func (a *FilesOperationsAPIService) DuplicateBatchItemsExecute(r ApiDuplicateBatchItemsRequest) (*FileOperationArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileOperationArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.DuplicateBatchItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/fileops/duplicate"

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
	localVarPostBody = r.duplicateRequestDto
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

type ApiEmptyTrashRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	single *bool
}

// Specifies whether to return only the current operation
func (r ApiEmptyTrashRequest) Single(single bool) ApiEmptyTrashRequest {	r.single = &single
	return r
}

func (r ApiEmptyTrashRequest) Execute() (*FileOperationArrayWrapper, *http.Response, error) {
	return r.ApiService.EmptyTrashExecute(r)
}

// EmptyTrash Empty the Trash folder
//
// Deletes all the files and folders from the Trash folder.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/empty-trash/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiEmptyTrashRequest
func (a *FilesOperationsAPIService) EmptyTrash(ctx context.Context) ApiEmptyTrashRequest {
	return ApiEmptyTrashRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FileOperationArrayWrapper
func (a *FilesOperationsAPIService) EmptyTrashExecute(r ApiEmptyTrashRequest) (*FileOperationArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileOperationArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.EmptyTrash")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/fileops/emptytrash"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.single != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Single", r.single, "form", "")
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

type ApiFinalizeSessionRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	folderId int32
	sessionId string
}

func (r ApiFinalizeSessionRequest) Execute() (*UploadSessionResponseIntegerWrapper, *http.Response, error) {
	return r.ApiService.FinalizeSessionExecute(r)
}

// FinalizeSession Finalize an upload session
//
// Finalizes the upload session by processing the uploaded file chunks and marking the upload as complete.  This method consolidates chunked uploads into a complete file if required, sends notifications about the upload event,  and performs any additional cleanup or related actions, such as socket updates and webhook publishing.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/finalize-session/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The folder ID.
// @param sessionId The session ID.
// @return ApiFinalizeSessionRequest
func (a *FilesOperationsAPIService) FinalizeSession(ctx context.Context, folderId int32, sessionId string) ApiFinalizeSessionRequest {
	return ApiFinalizeSessionRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
		sessionId: sessionId,
	}
}

// Execute executes the request
//  @return UploadSessionResponseIntegerWrapper
func (a *FilesOperationsAPIService) FinalizeSessionExecute(r ApiFinalizeSessionRequest) (*UploadSessionResponseIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UploadSessionResponseIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.FinalizeSession")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/{folderId}/session/{sessionId}/finalize"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", url.PathEscape(parameterValueToString(r.sessionId, "sessionId")), -1)

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

type ApiGetOperationStatusesRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	id *string
}

// The ID of the file operation.
func (r ApiGetOperationStatusesRequest) Id(id string) ApiGetOperationStatusesRequest {	r.id = &id
	return r
}

func (r ApiGetOperationStatusesRequest) Execute() (*FileOperationArrayWrapper, *http.Response, error) {
	return r.ApiService.GetOperationStatusesExecute(r)
}

// GetOperationStatuses Get active file operations
//
// Returns a list of all the active file operations.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-operation-statuses/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetOperationStatusesRequest
func (a *FilesOperationsAPIService) GetOperationStatuses(ctx context.Context) ApiGetOperationStatusesRequest {
	return ApiGetOperationStatusesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FileOperationArrayWrapper
func (a *FilesOperationsAPIService) GetOperationStatusesExecute(r ApiGetOperationStatusesRequest) (*FileOperationArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileOperationArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.GetOperationStatuses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/fileops"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "")
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

type ApiGetOperationStatusesByTypeRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	operationType FileOperationType
	id *string
}

// The ID of the file operation.
func (r ApiGetOperationStatusesByTypeRequest) Id(id string) ApiGetOperationStatusesByTypeRequest {	r.id = &id
	return r
}

func (r ApiGetOperationStatusesByTypeRequest) Execute() (*FileOperationArrayWrapper, *http.Response, error) {
	return r.ApiService.GetOperationStatusesByTypeExecute(r)
}

// GetOperationStatusesByType Get file operation statuses
//
// Retrieves the statuses of operations filtered by the specified operation type.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-operation-statuses-by-type/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param operationType Specifies the type of file operation to be retrieved.
// @return ApiGetOperationStatusesByTypeRequest
func (a *FilesOperationsAPIService) GetOperationStatusesByType(ctx context.Context, operationType FileOperationType) ApiGetOperationStatusesByTypeRequest {
	return ApiGetOperationStatusesByTypeRequest{
		ApiService: a,
		ctx: ctx,
		operationType: operationType,
	}
}

// Execute executes the request
//  @return FileOperationArrayWrapper
func (a *FilesOperationsAPIService) GetOperationStatusesByTypeExecute(r ApiGetOperationStatusesByTypeRequest) (*FileOperationArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileOperationArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.GetOperationStatusesByType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/fileops/{operationType}"
	localVarPath = strings.Replace(localVarPath, "{"+"operationType"+"}", url.PathEscape(parameterValueToString(r.operationType, "operationType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "")
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

type ApiMarkAsReadRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	baseBatchRequestDto *BaseBatchRequestDto
}

func (r ApiMarkAsReadRequest) BaseBatchRequestDto(baseBatchRequestDto BaseBatchRequestDto) ApiMarkAsReadRequest {	r.baseBatchRequestDto = &baseBatchRequestDto
	return r
}

func (r ApiMarkAsReadRequest) Execute() (*FileOperationArrayWrapper, *http.Response, error) {
	return r.ApiService.MarkAsReadExecute(r)
}

// MarkAsRead Mark as read
//
// Marks the files and folders with the IDs specified in the request as read.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/mark-as-read/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiMarkAsReadRequest
func (a *FilesOperationsAPIService) MarkAsRead(ctx context.Context) ApiMarkAsReadRequest {
	return ApiMarkAsReadRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FileOperationArrayWrapper
func (a *FilesOperationsAPIService) MarkAsReadExecute(r ApiMarkAsReadRequest) (*FileOperationArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileOperationArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.MarkAsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/fileops/markasread"

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

type ApiMoveBatchItemsRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	batchRequestDto *BatchRequestDto
}

func (r ApiMoveBatchItemsRequest) BatchRequestDto(batchRequestDto BatchRequestDto) ApiMoveBatchItemsRequest {	r.batchRequestDto = &batchRequestDto
	return r
}

func (r ApiMoveBatchItemsRequest) Execute() (*FileOperationArrayWrapper, *http.Response, error) {
	return r.ApiService.MoveBatchItemsExecute(r)
}

// MoveBatchItems Move or copy to a folder
//
// Moves or copies all the selected files and folders to the folder with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/move-batch-items/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiMoveBatchItemsRequest
func (a *FilesOperationsAPIService) MoveBatchItems(ctx context.Context) ApiMoveBatchItemsRequest {
	return ApiMoveBatchItemsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FileOperationArrayWrapper
func (a *FilesOperationsAPIService) MoveBatchItemsExecute(r ApiMoveBatchItemsRequest) (*FileOperationArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileOperationArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.MoveBatchItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/fileops/move"

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
	localVarPostBody = r.batchRequestDto
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

type ApiStartFileConversionRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	fileId int32
	checkConversionRequestDtoInteger *CheckConversionRequestDtoInteger
}

// The parameters for checking file conversion.
func (r ApiStartFileConversionRequest) CheckConversionRequestDtoInteger(checkConversionRequestDtoInteger CheckConversionRequestDtoInteger) ApiStartFileConversionRequest {	r.checkConversionRequestDtoInteger = &checkConversionRequestDtoInteger
	return r
}

func (r ApiStartFileConversionRequest) Execute() (*ConversationResultArrayWrapper, *http.Response, error) {
	return r.ApiService.StartFileConversionExecute(r)
}

// StartFileConversion Start file conversion
//
// Starts a conversion operation of a file with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/start-file-conversion/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param fileId The file ID to start conversion proccess.
// @return ApiStartFileConversionRequest
func (a *FilesOperationsAPIService) StartFileConversion(ctx context.Context, fileId int32) ApiStartFileConversionRequest {
	return ApiStartFileConversionRequest{
		ApiService: a,
		ctx: ctx,
		fileId: fileId,
	}
}

// Execute executes the request
//  @return ConversationResultArrayWrapper
func (a *FilesOperationsAPIService) StartFileConversionExecute(r ApiStartFileConversionRequest) (*ConversationResultArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConversationResultArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.StartFileConversion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/file/{fileId}/checkconversion"
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
	localVarPostBody = r.checkConversionRequestDtoInteger
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

type ApiTerminateTasksRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	id string
}

func (r ApiTerminateTasksRequest) Execute() (*FileOperationArrayWrapper, *http.Response, error) {
	return r.ApiService.TerminateTasksExecute(r)
}

// TerminateTasks Finish active operations
//
// Finishes an operation with the ID specified in the request or all the active operations.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/terminate-tasks/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The operation unique identifier.
// @return ApiTerminateTasksRequest
func (a *FilesOperationsAPIService) TerminateTasks(ctx context.Context, id string) ApiTerminateTasksRequest {
	return ApiTerminateTasksRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return FileOperationArrayWrapper
func (a *FilesOperationsAPIService) TerminateTasksExecute(r ApiTerminateTasksRequest) (*FileOperationArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileOperationArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.TerminateTasks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/fileops/terminate/{id}"
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

type ApiUpdateFileCommentRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	fileId int32
	updateComment *UpdateComment
}

// The parameters for updating a comment.
func (r ApiUpdateFileCommentRequest) UpdateComment(updateComment UpdateComment) ApiUpdateFileCommentRequest {	r.updateComment = &updateComment
	return r
}

func (r ApiUpdateFileCommentRequest) Execute() (*StringWrapper, *http.Response, error) {
	return r.ApiService.UpdateFileCommentExecute(r)
}

// UpdateFileComment Update a comment
//
// Updates a comment in a file with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/update-file-comment/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param fileId The file ID where the comment is located.
// @return ApiUpdateFileCommentRequest
func (a *FilesOperationsAPIService) UpdateFileComment(ctx context.Context, fileId int32) ApiUpdateFileCommentRequest {
	return ApiUpdateFileCommentRequest{
		ApiService: a,
		ctx: ctx,
		fileId: fileId,
	}
}

// Execute executes the request
//  @return StringWrapper
func (a *FilesOperationsAPIService) UpdateFileCommentExecute(r ApiUpdateFileCommentRequest) (*StringWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StringWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.UpdateFileComment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/file/{fileId}/comment"
	localVarPath = strings.Replace(localVarPath, "{"+"fileId"+"}", url.PathEscape(parameterValueToString(r.fileId, "fileId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateComment == nil {
		return localVarReturnValue, nil, reportError("updateComment is required and must be specified")
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
	localVarPostBody = r.updateComment
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

type ApiUploadAsyncSessionRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	folderId int32
	sessionId string
	chunkNumber *int32
	file *os.File
}

// The chunk number.
func (r ApiUploadAsyncSessionRequest) ChunkNumber(chunkNumber int32) ApiUploadAsyncSessionRequest {	r.chunkNumber = &chunkNumber
	return r
}

// The file chunk to be uploaded as part of the multipart/form-data request.  This property represents the uploaded file chunk content from the HTTP request form for chunked upload operations.  The file chunk is accessed via the IFormFile interface which provides access to the chunk content and length.
func (r ApiUploadAsyncSessionRequest) File(file *os.File) ApiUploadAsyncSessionRequest {	r.file = file
	return r
}

func (r ApiUploadAsyncSessionRequest) Execute() (*ChunkedUploadSessionResponseIntegerWrapper, *http.Response, error) {
	return r.ApiService.UploadAsyncSessionExecute(r)
}

// UploadAsyncSession Handles the upload of a chunk for an existing upload session.
//
// This method allows the caller to upload a specific chunk of a file to an ongoing upload session.  The session is identified by the session ID provided in the request. The chunk can be of any size  within the limits allowed during the session initialization. Each chunk must be uploaded in the  correct order for the server to process it appropriately.  The server updates the upload session status and stores the progress information after processing  each chunk. The updated session details are returned in the response.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/upload-async-session/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The folder ID.
// @param sessionId The upload session ID.
// @return ApiUploadAsyncSessionRequest
func (a *FilesOperationsAPIService) UploadAsyncSession(ctx context.Context, folderId int32, sessionId string) ApiUploadAsyncSessionRequest {
	return ApiUploadAsyncSessionRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
		sessionId: sessionId,
	}
}

// Execute executes the request
//  @return ChunkedUploadSessionResponseIntegerWrapper
func (a *FilesOperationsAPIService) UploadAsyncSessionExecute(r ApiUploadAsyncSessionRequest) (*ChunkedUploadSessionResponseIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ChunkedUploadSessionResponseIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.UploadAsyncSession")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/{folderId}/session/{sessionId}/upload"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", url.PathEscape(parameterValueToString(r.sessionId, "sessionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.chunkNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ChunkNumber", r.chunkNumber, "form", "")
			}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "File"
	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
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

type ApiUploadSessionRequest struct {
	ctx context.Context
	ApiService *FilesOperationsAPIService
	folderId int32
	sessionId string
	file *os.File
}

// The file to be uploaded as part of the multipart/form-data request.  This property represents the uploaded file content from the HTTP request form.  The file is accessed via the IFormFile interface which provides access to the file name, content type, length, and stream.
func (r ApiUploadSessionRequest) File(file *os.File) ApiUploadSessionRequest {	r.file = file
	return r
}

func (r ApiUploadSessionRequest) Execute() (*UploadSessionResponseIntegerWrapper, *http.Response, error) {
	return r.ApiService.UploadSessionExecute(r)
}

// UploadSession Resumes an ongoing file upload session for uploading additional chunks of data.
//
// This method allows continuing an interrupted or partially completed file upload session by uploading subsequent data chunks.  The server will validate each uploaded chunk, update the session state, and respond with the status of the current upload. Once  the total bytes uploaded match the total file size, the file upload process is finalized and related events are triggered.  If the file is newly uploaded, the server responds with a 201 Created status upon completion. If it overwrites an existing file,  versioning information is updated accordingly. The method also triggers associated webhooks and socket notifications to reflect  the updated file state.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/upload-session/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The folder ID.
// @param sessionId The upload session ID.
// @return ApiUploadSessionRequest
func (a *FilesOperationsAPIService) UploadSession(ctx context.Context, folderId int32, sessionId string) ApiUploadSessionRequest {
	return ApiUploadSessionRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
		sessionId: sessionId,
	}
}

// Execute executes the request
//  @return UploadSessionResponseIntegerWrapper
func (a *FilesOperationsAPIService) UploadSessionExecute(r ApiUploadSessionRequest) (*UploadSessionResponseIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UploadSessionResponseIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOperationsAPIService.UploadSession")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/{folderId}/session/{sessionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"sessionId"+"}", url.PathEscape(parameterValueToString(r.sessionId, "sessionId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "File"
	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
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
