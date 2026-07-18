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


// FilesFoldersAPIService FilesFoldersAPI service
type FilesFoldersAPIService service

type ApiCheckUploadRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	folderId int32
	checkUploadRequest *CheckUploadRequest
}

// The request parameters for checking file uploads.
func (r ApiCheckUploadRequest) CheckUploadRequest(checkUploadRequest CheckUploadRequest) ApiCheckUploadRequest {	r.checkUploadRequest = &checkUploadRequest
	return r
}

func (r ApiCheckUploadRequest) Execute() (*STRINGArrayWrapper, *http.Response, error) {
	return r.ApiService.CheckUploadExecute(r)
}

// CheckUpload Check file uploads
//
// Checks the file uploads to the folder with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/check-upload/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The folder ID.
// @return ApiCheckUploadRequest
func (a *FilesFoldersAPIService) CheckUpload(ctx context.Context, folderId int32) ApiCheckUploadRequest {
	return ApiCheckUploadRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return STRINGArrayWrapper
func (a *FilesFoldersAPIService) CheckUploadExecute(r ApiCheckUploadRequest) (*STRINGArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *STRINGArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.CheckUpload")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/{folderId}/upload/check"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.checkUploadRequest == nil {
		return localVarReturnValue, nil, reportError("checkUploadRequest is required and must be specified")
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
	localVarPostBody = r.checkUploadRequest
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

type ApiCreateFolderRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	folderId int32
	createFolder *CreateFolder
}

// The parameters for creating a folder.
func (r ApiCreateFolderRequest) CreateFolder(createFolder CreateFolder) ApiCreateFolderRequest {	r.createFolder = &createFolder
	return r
}

func (r ApiCreateFolderRequest) Execute() (*FolderIntegerWrapper, *http.Response, error) {
	return r.ApiService.CreateFolderExecute(r)
}

// CreateFolder Create a folder
//
// Creates a new folder with the title specified in the request. The parent folder ID can be also specified.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/create-folder/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The folder ID for the folder creation.
// @return ApiCreateFolderRequest
func (a *FilesFoldersAPIService) CreateFolder(ctx context.Context, folderId int32) ApiCreateFolderRequest {
	return ApiCreateFolderRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return FolderIntegerWrapper
func (a *FilesFoldersAPIService) CreateFolderExecute(r ApiCreateFolderRequest) (*FolderIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FolderIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.CreateFolder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/folder/{folderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createFolder == nil {
		return localVarReturnValue, nil, reportError("createFolder is required and must be specified")
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
	localVarPostBody = r.createFolder
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

type ApiCreateFolderPrimaryExternalLinkRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	id int32
	folderLinkRequest *FolderLinkRequest
}

// The folder link parameters.
func (r ApiCreateFolderPrimaryExternalLinkRequest) FolderLinkRequest(folderLinkRequest FolderLinkRequest) ApiCreateFolderPrimaryExternalLinkRequest {	r.folderLinkRequest = &folderLinkRequest
	return r
}

func (r ApiCreateFolderPrimaryExternalLinkRequest) Execute() (*FileShareWrapper, *http.Response, error) {
	return r.ApiService.CreateFolderPrimaryExternalLinkExecute(r)
}

// CreateFolderPrimaryExternalLink Create primary external link
//
// Creates a primary external link by the identifier specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/create-folder-primary-external-link/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The folder ID.
// @return ApiCreateFolderPrimaryExternalLinkRequest
func (a *FilesFoldersAPIService) CreateFolderPrimaryExternalLink(ctx context.Context, id int32) ApiCreateFolderPrimaryExternalLinkRequest {
	return ApiCreateFolderPrimaryExternalLinkRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return FileShareWrapper
func (a *FilesFoldersAPIService) CreateFolderPrimaryExternalLinkExecute(r ApiCreateFolderPrimaryExternalLinkRequest) (*FileShareWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileShareWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.CreateFolderPrimaryExternalLink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/folder/{id}/link"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.folderLinkRequest == nil {
		return localVarReturnValue, nil, reportError("folderLinkRequest is required and must be specified")
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
	localVarPostBody = r.folderLinkRequest
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

type ApiCreateReportFolderHistoryRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	folderId int32
}

func (r ApiCreateReportFolderHistoryRequest) Execute() (*StringWrapper, *http.Response, error) {
	return r.ApiService.CreateReportFolderHistoryExecute(r)
}

// CreateReportFolderHistory Generates folder history
//
// Generates the activity history of a folder.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/create-report-folder-history/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId
// @return ApiCreateReportFolderHistoryRequest
func (a *FilesFoldersAPIService) CreateReportFolderHistory(ctx context.Context, folderId int32) ApiCreateReportFolderHistoryRequest {
	return ApiCreateReportFolderHistoryRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return StringWrapper
func (a *FilesFoldersAPIService) CreateReportFolderHistoryExecute(r ApiCreateReportFolderHistoryRequest) (*StringWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StringWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.CreateReportFolderHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/folder/{folderId}/log/report"
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

type ApiDeleteFolderRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	folderId int32
	deleteFolder *DeleteFolder
}

// The parameters for deleting a folder.
func (r ApiDeleteFolderRequest) DeleteFolder(deleteFolder DeleteFolder) ApiDeleteFolderRequest {	r.deleteFolder = &deleteFolder
	return r
}

func (r ApiDeleteFolderRequest) Execute() (*FileOperationArrayWrapper, *http.Response, error) {
	return r.ApiService.DeleteFolderExecute(r)
}

// DeleteFolder Delete a folder
//
// Deletes a folder with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-folder/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The folder ID to delete.
// @return ApiDeleteFolderRequest
func (a *FilesFoldersAPIService) DeleteFolder(ctx context.Context, folderId int32) ApiDeleteFolderRequest {
	return ApiDeleteFolderRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return FileOperationArrayWrapper
func (a *FilesFoldersAPIService) DeleteFolderExecute(r ApiDeleteFolderRequest) (*FileOperationArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileOperationArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.DeleteFolder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/folder/{folderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deleteFolder == nil {
		return localVarReturnValue, nil, reportError("deleteFolder is required and must be specified")
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
	localVarPostBody = r.deleteFolder
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

type ApiGenerateXlsxByFolderRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	folderId int32
}

func (r ApiGenerateXlsxByFolderRequest) Execute() (*XlsxReportResponseWrapper, *http.Response, error) {
	return r.ApiService.GenerateXlsxByFolderExecute(r)
}

// GenerateXlsxByFolder Generate XLSX report by folder
//
// Triggers asynchronous XLSX report generation for the specified form results folder.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/generate-xlsx-by-folder/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The folder unique identifier.
// @return ApiGenerateXlsxByFolderRequest
func (a *FilesFoldersAPIService) GenerateXlsxByFolder(ctx context.Context, folderId int32) ApiGenerateXlsxByFolderRequest {
	return ApiGenerateXlsxByFolderRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return XlsxReportResponseWrapper
func (a *FilesFoldersAPIService) GenerateXlsxByFolderExecute(r ApiGenerateXlsxByFolderRequest) (*XlsxReportResponseWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *XlsxReportResponseWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.GenerateXlsxByFolder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/folder/{folderId}/xlsx"
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

type ApiGetFavoritesFolderRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	userIdOrGroupId *string
	filterType *FilterType
	count *int32
	startIndex *int32
	sortBy *string
	sortOrder *SortOrder
	filterValue *string
}

// The user or group ID.
func (r ApiGetFavoritesFolderRequest) UserIdOrGroupId(userIdOrGroupId string) ApiGetFavoritesFolderRequest {	r.userIdOrGroupId = &userIdOrGroupId
	return r
}

// The filter type.
func (r ApiGetFavoritesFolderRequest) FilterType(filterType FilterType) ApiGetFavoritesFolderRequest {	r.filterType = &filterType
	return r
}

// The maximum number of items to retrieve in the request.
func (r ApiGetFavoritesFolderRequest) Count(count int32) ApiGetFavoritesFolderRequest {	r.count = &count
	return r
}

// The zero-based index of the first item to retrieve in a paginated list.
func (r ApiGetFavoritesFolderRequest) StartIndex(startIndex int32) ApiGetFavoritesFolderRequest {	r.startIndex = &startIndex
	return r
}

// Specifies the field by which the folder content should be sorted.
func (r ApiGetFavoritesFolderRequest) SortBy(sortBy string) ApiGetFavoritesFolderRequest {	r.sortBy = &sortBy
	return r
}

// The order in which the results are sorted.
func (r ApiGetFavoritesFolderRequest) SortOrder(sortOrder SortOrder) ApiGetFavoritesFolderRequest {	r.sortOrder = &sortOrder
	return r
}

// The text used as a filter or search criterion for folder content queries.
func (r ApiGetFavoritesFolderRequest) FilterValue(filterValue string) ApiGetFavoritesFolderRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetFavoritesFolderRequest) Execute() (*FolderContentIntegerWrapper, *http.Response, error) {
	return r.ApiService.GetFavoritesFolderExecute(r)
}

// GetFavoritesFolder Get the Favorites section
//
// Returns the detailed list of files and folders located in the Favorites section.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-favorites-folder/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetFavoritesFolderRequest
func (a *FilesFoldersAPIService) GetFavoritesFolder(ctx context.Context) ApiGetFavoritesFolderRequest {
	return ApiGetFavoritesFolderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FolderContentIntegerWrapper
func (a *FilesFoldersAPIService) GetFavoritesFolderExecute(r ApiGetFavoritesFolderRequest) (*FolderContentIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FolderContentIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.GetFavoritesFolder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/@favorites"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userIdOrGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userIdOrGroupId", r.userIdOrGroupId, "form", "")
			}
	if r.filterType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterType", r.filterType, "form", "")
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

type ApiGetFilesUsedSpaceRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
}

func (r ApiGetFilesUsedSpaceRequest) Execute() (*FilesStatisticsResultWrapper, *http.Response, error) {
	return r.ApiService.GetFilesUsedSpaceExecute(r)
}

// GetFilesUsedSpace Get used space of files
//
// Returns the used space of files in the root folders.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-files-used-space/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetFilesUsedSpaceRequest
func (a *FilesFoldersAPIService) GetFilesUsedSpace(ctx context.Context) ApiGetFilesUsedSpaceRequest {
	return ApiGetFilesUsedSpaceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FilesStatisticsResultWrapper
func (a *FilesFoldersAPIService) GetFilesUsedSpaceExecute(r ApiGetFilesUsedSpaceRequest) (*FilesStatisticsResultWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FilesStatisticsResultWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.GetFilesUsedSpace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/filesusedspace"

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

type ApiGetFolderRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	folderId int32
}

func (r ApiGetFolderRequest) Execute() (*FormsItemArrayWrapper, *http.Response, error) {
	return r.ApiService.GetFolderExecute(r)
}

// GetFolder Get folder form filter
//
// Returns the form filter of a folder with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-folder/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The folder unique identifier.
// @return ApiGetFolderRequest
func (a *FilesFoldersAPIService) GetFolder(ctx context.Context, folderId int32) ApiGetFolderRequest {
	return ApiGetFolderRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return FormsItemArrayWrapper
func (a *FilesFoldersAPIService) GetFolderExecute(r ApiGetFolderRequest) (*FormsItemArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FormsItemArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.GetFolder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/{folderId}/formfilter"
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

type ApiGetFolderByFolderIdRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	folderId int32
	userIdOrGroupId *string
	sharedBy *string
	filterType *FilterType
	roomId *int32
	excludeSubject *bool
	applyFilterOption *ApplyFilterOption
	withSubFolders *bool
	extension *string
	searchArea *SearchArea
	formsItemKey *string
	formsItemType *string
	count *int32
	startIndex *int32
	sortBy *string
	sortOrder *SortOrder
	filterValue *string
	location *Location
}

// The user or group ID.
func (r ApiGetFolderByFolderIdRequest) UserIdOrGroupId(userIdOrGroupId string) ApiGetFolderByFolderIdRequest {	r.userIdOrGroupId = &userIdOrGroupId
	return r
}

// The identifier of the user who shared the folder or file.
func (r ApiGetFolderByFolderIdRequest) SharedBy(sharedBy string) ApiGetFolderByFolderIdRequest {	r.sharedBy = &sharedBy
	return r
}

// The filter type.
func (r ApiGetFolderByFolderIdRequest) FilterType(filterType FilterType) ApiGetFolderByFolderIdRequest {	r.filterType = &filterType
	return r
}

// The room ID.
func (r ApiGetFolderByFolderIdRequest) RoomId(roomId int32) ApiGetFolderByFolderIdRequest {	r.roomId = &roomId
	return r
}

// Specifies whether to exclude search by user or group ID.
func (r ApiGetFolderByFolderIdRequest) ExcludeSubject(excludeSubject bool) ApiGetFolderByFolderIdRequest {	r.excludeSubject = &excludeSubject
	return r
}

// Specifies whether to return only files, only folders, or all elements from the specified folder.
func (r ApiGetFolderByFolderIdRequest) ApplyFilterOption(applyFilterOption ApplyFilterOption) ApiGetFolderByFolderIdRequest {	r.applyFilterOption = &applyFilterOption
	return r
}

// Specifies whether to include files from subfolders in the results.
func (r ApiGetFolderByFolderIdRequest) WithSubFolders(withSubFolders bool) ApiGetFolderByFolderIdRequest {	r.withSubFolders = &withSubFolders
	return r
}

// Specifies whether to search for the specific file extension.
func (r ApiGetFolderByFolderIdRequest) Extension(extension string) ApiGetFolderByFolderIdRequest {	r.extension = &extension
	return r
}

// The search area.
func (r ApiGetFolderByFolderIdRequest) SearchArea(searchArea SearchArea) ApiGetFolderByFolderIdRequest {	r.searchArea = &searchArea
	return r
}

// The forms item key.
func (r ApiGetFolderByFolderIdRequest) FormsItemKey(formsItemKey string) ApiGetFolderByFolderIdRequest {	r.formsItemKey = &formsItemKey
	return r
}

// The forms item type.
func (r ApiGetFolderByFolderIdRequest) FormsItemType(formsItemType string) ApiGetFolderByFolderIdRequest {	r.formsItemType = &formsItemType
	return r
}

// The maximum number of items to retrieve in the request.
func (r ApiGetFolderByFolderIdRequest) Count(count int32) ApiGetFolderByFolderIdRequest {	r.count = &count
	return r
}

// The zero-based index of the first item to retrieve in a paginated request.
func (r ApiGetFolderByFolderIdRequest) StartIndex(startIndex int32) ApiGetFolderByFolderIdRequest {	r.startIndex = &startIndex
	return r
}

// The property used for sorting the folder request results.
func (r ApiGetFolderByFolderIdRequest) SortBy(sortBy string) ApiGetFolderByFolderIdRequest {	r.sortBy = &sortBy
	return r
}

// The order in which the results are sorted.
func (r ApiGetFolderByFolderIdRequest) SortOrder(sortOrder SortOrder) ApiGetFolderByFolderIdRequest {	r.sortOrder = &sortOrder
	return r
}

// The text value used as a filter parameter for folder content queries.
func (r ApiGetFolderByFolderIdRequest) FilterValue(filterValue string) ApiGetFolderByFolderIdRequest {	r.filterValue = &filterValue
	return r
}

// The location context of the request, specifying the area  where the operation is performed, such as a room, documents, or a link.
func (r ApiGetFolderByFolderIdRequest) Location(location Location) ApiGetFolderByFolderIdRequest {	r.location = &location
	return r
}

func (r ApiGetFolderByFolderIdRequest) Execute() (*FolderContentIntegerWrapper, *http.Response, error) {
	return r.ApiService.GetFolderByFolderIdExecute(r)
}

// GetFolderByFolderId Get a folder by ID
//
// Returns the detailed list of files and folders located in the folder with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-folder-by-folder-id/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The folder ID.
// @return ApiGetFolderByFolderIdRequest
func (a *FilesFoldersAPIService) GetFolderByFolderId(ctx context.Context, folderId int32) ApiGetFolderByFolderIdRequest {
	return ApiGetFolderByFolderIdRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return FolderContentIntegerWrapper
func (a *FilesFoldersAPIService) GetFolderByFolderIdExecute(r ApiGetFolderByFolderIdRequest) (*FolderContentIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FolderContentIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.GetFolderByFolderId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/{folderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userIdOrGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userIdOrGroupId", r.userIdOrGroupId, "form", "")
			}
	if r.sharedBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sharedBy", r.sharedBy, "form", "")
			}
	if r.filterType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterType", r.filterType, "form", "")
			}
	if r.roomId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "roomId", r.roomId, "form", "")
			}
	if r.excludeSubject != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "excludeSubject", r.excludeSubject, "form", "")
			}
	if r.applyFilterOption != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "applyFilterOption", r.applyFilterOption, "form", "")
			}
	if r.withSubFolders != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "withSubFolders", r.withSubFolders, "form", "")
			}
	if r.extension != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "extension", r.extension, "form", "")
			}
	if r.searchArea != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchArea", r.searchArea, "form", "")
			}
	if r.formsItemKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "formsItemKey", r.formsItemKey, "form", "")
			}
	if r.formsItemType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "formsItemType", r.formsItemType, "form", "")
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
	if r.location != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Location", r.location, "form", "")
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

type ApiGetFolderHistoryRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	folderId int32
	fromDate *ApiDateTime
	toDate *ApiDateTime
	count *int32
	startIndex *int32
}

// The start date of the history request.
func (r ApiGetFolderHistoryRequest) FromDate(fromDate ApiDateTime) ApiGetFolderHistoryRequest {	r.fromDate = &fromDate
	return r
}

// The end date of the history request.
func (r ApiGetFolderHistoryRequest) ToDate(toDate ApiDateTime) ApiGetFolderHistoryRequest {	r.toDate = &toDate
	return r
}

// The number of records to retrieve for the folder history.
func (r ApiGetFolderHistoryRequest) Count(count int32) ApiGetFolderHistoryRequest {	r.count = &count
	return r
}

// The starting index from which the history records are retrieved in the request.
func (r ApiGetFolderHistoryRequest) StartIndex(startIndex int32) ApiGetFolderHistoryRequest {	r.startIndex = &startIndex
	return r
}

func (r ApiGetFolderHistoryRequest) Execute() (*HistoryArrayWrapper, *http.Response, error) {
	return r.ApiService.GetFolderHistoryExecute(r)
}

// GetFolderHistory Get folder history
//
// Returns the activity history of a folder with a specified identifier.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-folder-history/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The folder ID of the history request.
// @return ApiGetFolderHistoryRequest
func (a *FilesFoldersAPIService) GetFolderHistory(ctx context.Context, folderId int32) ApiGetFolderHistoryRequest {
	return ApiGetFolderHistoryRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return HistoryArrayWrapper
func (a *FilesFoldersAPIService) GetFolderHistoryExecute(r ApiGetFolderHistoryRequest) (*HistoryArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HistoryArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.GetFolderHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/folder/{folderId}/log"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fromDate", r.fromDate, "deepObject", "")
			}
	if r.toDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "toDate", r.toDate, "deepObject", "")
			}
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

type ApiGetFolderInfoRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	folderId int32
}

func (r ApiGetFolderInfoRequest) Execute() (*FolderIntegerWrapper, *http.Response, error) {
	return r.ApiService.GetFolderInfoExecute(r)
}

// GetFolderInfo Get folder information
//
// Returns the detailed information about a folder with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-folder-info/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The folder unique identifier.
// @return ApiGetFolderInfoRequest
func (a *FilesFoldersAPIService) GetFolderInfo(ctx context.Context, folderId int32) ApiGetFolderInfoRequest {
	return ApiGetFolderInfoRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return FolderIntegerWrapper
func (a *FilesFoldersAPIService) GetFolderInfoExecute(r ApiGetFolderInfoRequest) (*FolderIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FolderIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.GetFolderInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/folder/{folderId}"
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

type ApiGetFolderLinksRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	id int32
}

func (r ApiGetFolderLinksRequest) Execute() (*FileShareArrayWrapper, *http.Response, error) {
	return r.ApiService.GetFolderLinksExecute(r)
}

// GetFolderLinks Get the folder links
//
// Returns the links of the folder with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-folder-links/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The folder ID.
// @return ApiGetFolderLinksRequest
func (a *FilesFoldersAPIService) GetFolderLinks(ctx context.Context, id int32) ApiGetFolderLinksRequest {
	return ApiGetFolderLinksRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return FileShareArrayWrapper
func (a *FilesFoldersAPIService) GetFolderLinksExecute(r ApiGetFolderLinksRequest) (*FileShareArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileShareArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.GetFolderLinks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/folder/{id}/links"
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

type ApiGetFolderPathRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	folderId int32
}

func (r ApiGetFolderPathRequest) Execute() (*FileEntryBaseArrayWrapper, *http.Response, error) {
	return r.ApiService.GetFolderPathExecute(r)
}

// GetFolderPath Get the folder path
//
// Returns a path to the folder with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-folder-path/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The folder unique identifier.
// @return ApiGetFolderPathRequest
func (a *FilesFoldersAPIService) GetFolderPath(ctx context.Context, folderId int32) ApiGetFolderPathRequest {
	return ApiGetFolderPathRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return FileEntryBaseArrayWrapper
func (a *FilesFoldersAPIService) GetFolderPathExecute(r ApiGetFolderPathRequest) (*FileEntryBaseArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileEntryBaseArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.GetFolderPath")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/folder/{folderId}/path"
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

type ApiGetFolderPrimaryExternalLinkRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	id int32
	count *int32
	startIndex *int32
}

// The number of items to retrieve in the request.
func (r ApiGetFolderPrimaryExternalLinkRequest) Count(count int32) ApiGetFolderPrimaryExternalLinkRequest {	r.count = &count
	return r
}

// The starting index for the query results.
func (r ApiGetFolderPrimaryExternalLinkRequest) StartIndex(startIndex int32) ApiGetFolderPrimaryExternalLinkRequest {	r.startIndex = &startIndex
	return r
}

func (r ApiGetFolderPrimaryExternalLinkRequest) Execute() (*FileShareWrapper, *http.Response, error) {
	return r.ApiService.GetFolderPrimaryExternalLinkExecute(r)
}

// GetFolderPrimaryExternalLink Get primary external link
//
// Returns the primary external link by the identifier specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-folder-primary-external-link/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The folder unique identifier.
// @return ApiGetFolderPrimaryExternalLinkRequest
func (a *FilesFoldersAPIService) GetFolderPrimaryExternalLink(ctx context.Context, id int32) ApiGetFolderPrimaryExternalLinkRequest {
	return ApiGetFolderPrimaryExternalLinkRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return FileShareWrapper
func (a *FilesFoldersAPIService) GetFolderPrimaryExternalLinkExecute(r ApiGetFolderPrimaryExternalLinkRequest) (*FileShareWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileShareWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.GetFolderPrimaryExternalLink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/folder/{id}/link"
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

type ApiGetFoldersRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	folderId int32
}

func (r ApiGetFoldersRequest) Execute() (*FileEntryBaseArrayWrapper, *http.Response, error) {
	return r.ApiService.GetFoldersExecute(r)
}

// GetFolders Get subfolders
//
// Returns a list of all the subfolders from a folder with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-folders/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The folder unique identifier.
// @return ApiGetFoldersRequest
func (a *FilesFoldersAPIService) GetFolders(ctx context.Context, folderId int32) ApiGetFoldersRequest {
	return ApiGetFoldersRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return FileEntryBaseArrayWrapper
func (a *FilesFoldersAPIService) GetFoldersExecute(r ApiGetFoldersRequest) (*FileEntryBaseArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileEntryBaseArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.GetFolders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/{folderId}/subfolders"
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

type ApiGetMyFolderRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	userIdOrGroupId *string
	filterType *FilterType
	applyFilterOption *ApplyFilterOption
	count *int32
	startIndex *int32
	sortBy *string
	sortOrder *SortOrder
	filterValue *string
}

// The user or group ID.
func (r ApiGetMyFolderRequest) UserIdOrGroupId(userIdOrGroupId string) ApiGetMyFolderRequest {	r.userIdOrGroupId = &userIdOrGroupId
	return r
}

// The filter type.
func (r ApiGetMyFolderRequest) FilterType(filterType FilterType) ApiGetMyFolderRequest {	r.filterType = &filterType
	return r
}

// Specifies whether to return only files, only folders or all elements.
func (r ApiGetMyFolderRequest) ApplyFilterOption(applyFilterOption ApplyFilterOption) ApiGetMyFolderRequest {	r.applyFilterOption = &applyFilterOption
	return r
}

// The maximum number of items to retrieve in the response.
func (r ApiGetMyFolderRequest) Count(count int32) ApiGetMyFolderRequest {	r.count = &count
	return r
}

// The starting position of the items to be retrieved.
func (r ApiGetMyFolderRequest) StartIndex(startIndex int32) ApiGetMyFolderRequest {	r.startIndex = &startIndex
	return r
}

// The property used to specify the sorting criteria for folder contents.
func (r ApiGetMyFolderRequest) SortBy(sortBy string) ApiGetMyFolderRequest {	r.sortBy = &sortBy
	return r
}

// The order in which the results are sorted.
func (r ApiGetMyFolderRequest) SortOrder(sortOrder SortOrder) ApiGetMyFolderRequest {	r.sortOrder = &sortOrder
	return r
}

// The text used for filtering or searching folder contents.
func (r ApiGetMyFolderRequest) FilterValue(filterValue string) ApiGetMyFolderRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetMyFolderRequest) Execute() (*FolderContentIntegerWrapper, *http.Response, error) {
	return r.ApiService.GetMyFolderExecute(r)
}

// GetMyFolder Get the My documents section
//
// Returns the detailed list of files and folders located in the My documents section.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-my-folder/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetMyFolderRequest
func (a *FilesFoldersAPIService) GetMyFolder(ctx context.Context) ApiGetMyFolderRequest {
	return ApiGetMyFolderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FolderContentIntegerWrapper
func (a *FilesFoldersAPIService) GetMyFolderExecute(r ApiGetMyFolderRequest) (*FolderContentIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FolderContentIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.GetMyFolder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/@my"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userIdOrGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userIdOrGroupId", r.userIdOrGroupId, "form", "")
			}
	if r.filterType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterType", r.filterType, "form", "")
			}
	if r.applyFilterOption != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "applyFilterOption", r.applyFilterOption, "form", "")
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

type ApiGetNewFolderItemsRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	folderId int32
}

func (r ApiGetNewFolderItemsRequest) Execute() (*FileEntryBaseArrayWrapper, *http.Response, error) {
	return r.ApiService.GetNewFolderItemsExecute(r)
}

// GetNewFolderItems Get new folder items
//
// Returns a list of all the new items from a folder with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-new-folder-items/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The folder unique identifier.
// @return ApiGetNewFolderItemsRequest
func (a *FilesFoldersAPIService) GetNewFolderItems(ctx context.Context, folderId int32) ApiGetNewFolderItemsRequest {
	return ApiGetNewFolderItemsRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return FileEntryBaseArrayWrapper
func (a *FilesFoldersAPIService) GetNewFolderItemsExecute(r ApiGetNewFolderItemsRequest) (*FileEntryBaseArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileEntryBaseArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.GetNewFolderItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/{folderId}/news"
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

type ApiGetPrivacyFolderRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	userIdOrGroupId *string
	filterType *FilterType
	count *int32
	startIndex *int32
	sortBy *string
	sortOrder *SortOrder
	filterValue *string
}

// The user or group ID.
func (r ApiGetPrivacyFolderRequest) UserIdOrGroupId(userIdOrGroupId string) ApiGetPrivacyFolderRequest {	r.userIdOrGroupId = &userIdOrGroupId
	return r
}

// The filter type.
func (r ApiGetPrivacyFolderRequest) FilterType(filterType FilterType) ApiGetPrivacyFolderRequest {	r.filterType = &filterType
	return r
}

// The maximum number of items to retrieve in the request.
func (r ApiGetPrivacyFolderRequest) Count(count int32) ApiGetPrivacyFolderRequest {	r.count = &count
	return r
}

// The zero-based index of the first item to retrieve in a paginated list.
func (r ApiGetPrivacyFolderRequest) StartIndex(startIndex int32) ApiGetPrivacyFolderRequest {	r.startIndex = &startIndex
	return r
}

// Specifies the field by which the folder content should be sorted.
func (r ApiGetPrivacyFolderRequest) SortBy(sortBy string) ApiGetPrivacyFolderRequest {	r.sortBy = &sortBy
	return r
}

// The order in which the results are sorted.
func (r ApiGetPrivacyFolderRequest) SortOrder(sortOrder SortOrder) ApiGetPrivacyFolderRequest {	r.sortOrder = &sortOrder
	return r
}

// The text used as a filter or search criterion for folder content queries.
func (r ApiGetPrivacyFolderRequest) FilterValue(filterValue string) ApiGetPrivacyFolderRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetPrivacyFolderRequest) Execute() (*FolderContentIntegerWrapper, *http.Response, error) {
	return r.ApiService.GetPrivacyFolderExecute(r)
}

// GetPrivacyFolder Get the Private Room section
//
// Returns the detailed list of files and folders located in the Private Room section.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-privacy-folder/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetPrivacyFolderRequest
func (a *FilesFoldersAPIService) GetPrivacyFolder(ctx context.Context) ApiGetPrivacyFolderRequest {
	return ApiGetPrivacyFolderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FolderContentIntegerWrapper
func (a *FilesFoldersAPIService) GetPrivacyFolderExecute(r ApiGetPrivacyFolderRequest) (*FolderContentIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FolderContentIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.GetPrivacyFolder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/@privacy"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userIdOrGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userIdOrGroupId", r.userIdOrGroupId, "form", "")
			}
	if r.filterType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterType", r.filterType, "form", "")
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

type ApiGetRecentFolderRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	userIdOrGroupId *string
	filterType *FilterType
	excludeSubject *bool
	applyFilterOption *ApplyFilterOption
	searchArea *SearchArea
	extension *[]string
	count *int32
	startIndex *int32
	sortBy *string
	sortOrder *SortOrder
	filterValue *string
}

// The user or group ID.
func (r ApiGetRecentFolderRequest) UserIdOrGroupId(userIdOrGroupId string) ApiGetRecentFolderRequest {	r.userIdOrGroupId = &userIdOrGroupId
	return r
}

// The filter type.
func (r ApiGetRecentFolderRequest) FilterType(filterType FilterType) ApiGetRecentFolderRequest {	r.filterType = &filterType
	return r
}

// Specifies whether to exclude search by user or group ID.
func (r ApiGetRecentFolderRequest) ExcludeSubject(excludeSubject bool) ApiGetRecentFolderRequest {	r.excludeSubject = &excludeSubject
	return r
}

// Specifies whether to return only files, only folders or all elements.
func (r ApiGetRecentFolderRequest) ApplyFilterOption(applyFilterOption ApplyFilterOption) ApiGetRecentFolderRequest {	r.applyFilterOption = &applyFilterOption
	return r
}

// The search area.
func (r ApiGetRecentFolderRequest) SearchArea(searchArea SearchArea) ApiGetRecentFolderRequest {	r.searchArea = &searchArea
	return r
}

// Specifies whether to search for a specific file extension in the Recent folder.
func (r ApiGetRecentFolderRequest) Extension(extension []string) ApiGetRecentFolderRequest {	r.extension = &extension
	return r
}

// The maximum number of items to return.
func (r ApiGetRecentFolderRequest) Count(count int32) ApiGetRecentFolderRequest {	r.count = &count
	return r
}

// The starting position of the results to be returned in the query response.
func (r ApiGetRecentFolderRequest) StartIndex(startIndex int32) ApiGetRecentFolderRequest {	r.startIndex = &startIndex
	return r
}

// Specifies the sorting criteria for the folder request.
func (r ApiGetRecentFolderRequest) SortBy(sortBy string) ApiGetRecentFolderRequest {	r.sortBy = &sortBy
	return r
}

// The order in which the results are sorted.
func (r ApiGetRecentFolderRequest) SortOrder(sortOrder SortOrder) ApiGetRecentFolderRequest {	r.sortOrder = &sortOrder
	return r
}

// The text used for filtering or searching folder contents.
func (r ApiGetRecentFolderRequest) FilterValue(filterValue string) ApiGetRecentFolderRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetRecentFolderRequest) Execute() (*FolderContentIntegerWrapper, *http.Response, error) {
	return r.ApiService.GetRecentFolderExecute(r)
}

// GetRecentFolder Get the Recent section
//
// Returns the detailed list of files located in the Recent section.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-recent-folder/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetRecentFolderRequest
func (a *FilesFoldersAPIService) GetRecentFolder(ctx context.Context) ApiGetRecentFolderRequest {
	return ApiGetRecentFolderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FolderContentIntegerWrapper
func (a *FilesFoldersAPIService) GetRecentFolderExecute(r ApiGetRecentFolderRequest) (*FolderContentIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FolderContentIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.GetRecentFolder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/recent"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userIdOrGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userIdOrGroupId", r.userIdOrGroupId, "form", "")
			}
	if r.filterType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterType", r.filterType, "form", "")
			}
	if r.excludeSubject != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "excludeSubject", r.excludeSubject, "form", "")
			}
	if r.applyFilterOption != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "applyFilterOption", r.applyFilterOption, "form", "")
			}
	if r.searchArea != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchArea", r.searchArea, "form", "")
			}
	if r.extension != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "extension", r.extension, "deepObject", "csv")
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

type ApiGetRootFoldersRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	userIdOrGroupId *string
	filterType *FilterType
	withoutTrash *bool
	count *int32
	startIndex *int32
	sortBy *string
	sortOrder *SortOrder
	filterValue *string
}

// The user or group ID.
func (r ApiGetRootFoldersRequest) UserIdOrGroupId(userIdOrGroupId string) ApiGetRootFoldersRequest {	r.userIdOrGroupId = &userIdOrGroupId
	return r
}

// The filter type.
func (r ApiGetRootFoldersRequest) FilterType(filterType FilterType) ApiGetRootFoldersRequest {	r.filterType = &filterType
	return r
}

// Specifies whether to return the Trash section or not.
func (r ApiGetRootFoldersRequest) WithoutTrash(withoutTrash bool) ApiGetRootFoldersRequest {	r.withoutTrash = &withoutTrash
	return r
}

// The maximum number of items to retrieve in the response.
func (r ApiGetRootFoldersRequest) Count(count int32) ApiGetRootFoldersRequest {	r.count = &count
	return r
}

// The starting position of the items to be retrieved.
func (r ApiGetRootFoldersRequest) StartIndex(startIndex int32) ApiGetRootFoldersRequest {	r.startIndex = &startIndex
	return r
}

// Specifies the field by which the folder content should be sorted.
func (r ApiGetRootFoldersRequest) SortBy(sortBy string) ApiGetRootFoldersRequest {	r.sortBy = &sortBy
	return r
}

// The order in which the results are sorted.
func (r ApiGetRootFoldersRequest) SortOrder(sortOrder SortOrder) ApiGetRootFoldersRequest {	r.sortOrder = &sortOrder
	return r
}

// The text used as a filter for searching or retrieving folder contents.
func (r ApiGetRootFoldersRequest) FilterValue(filterValue string) ApiGetRootFoldersRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetRootFoldersRequest) Execute() (*FolderContentIntegerArrayWrapper, *http.Response, error) {
	return r.ApiService.GetRootFoldersExecute(r)
}

// GetRootFolders Get filtered sections
//
// Returns all the sections matching the parameters specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-root-folders/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetRootFoldersRequest
func (a *FilesFoldersAPIService) GetRootFolders(ctx context.Context) ApiGetRootFoldersRequest {
	return ApiGetRootFoldersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FolderContentIntegerArrayWrapper
func (a *FilesFoldersAPIService) GetRootFoldersExecute(r ApiGetRootFoldersRequest) (*FolderContentIntegerArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FolderContentIntegerArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.GetRootFolders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/@root"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userIdOrGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userIdOrGroupId", r.userIdOrGroupId, "form", "")
			}
	if r.filterType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterType", r.filterType, "form", "")
			}
	if r.withoutTrash != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "withoutTrash", r.withoutTrash, "form", "")
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

type ApiGetTrashFolderRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	userIdOrGroupId *string
	filterType *FilterType
	applyFilterOption *ApplyFilterOption
	count *int32
	startIndex *int32
	sortBy *string
	sortOrder *SortOrder
	filterValue *string
}

// The user or group ID.
func (r ApiGetTrashFolderRequest) UserIdOrGroupId(userIdOrGroupId string) ApiGetTrashFolderRequest {	r.userIdOrGroupId = &userIdOrGroupId
	return r
}

// The filter type.
func (r ApiGetTrashFolderRequest) FilterType(filterType FilterType) ApiGetTrashFolderRequest {	r.filterType = &filterType
	return r
}

// Specifies whether to return only files, only folders or all elements.
func (r ApiGetTrashFolderRequest) ApplyFilterOption(applyFilterOption ApplyFilterOption) ApiGetTrashFolderRequest {	r.applyFilterOption = &applyFilterOption
	return r
}

// The maximum number of items to retrieve in the response.
func (r ApiGetTrashFolderRequest) Count(count int32) ApiGetTrashFolderRequest {	r.count = &count
	return r
}

// The starting position of the items to be retrieved.
func (r ApiGetTrashFolderRequest) StartIndex(startIndex int32) ApiGetTrashFolderRequest {	r.startIndex = &startIndex
	return r
}

// The property used to specify the sorting criteria for folder contents.
func (r ApiGetTrashFolderRequest) SortBy(sortBy string) ApiGetTrashFolderRequest {	r.sortBy = &sortBy
	return r
}

// The order in which the results are sorted.
func (r ApiGetTrashFolderRequest) SortOrder(sortOrder SortOrder) ApiGetTrashFolderRequest {	r.sortOrder = &sortOrder
	return r
}

// The text used for filtering or searching folder contents.
func (r ApiGetTrashFolderRequest) FilterValue(filterValue string) ApiGetTrashFolderRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetTrashFolderRequest) Execute() (*FolderContentIntegerWrapper, *http.Response, error) {
	return r.ApiService.GetTrashFolderExecute(r)
}

// GetTrashFolder Get the Trash section
//
// Returns the detailed list of files and folders located in the Trash section.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-trash-folder/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetTrashFolderRequest
func (a *FilesFoldersAPIService) GetTrashFolder(ctx context.Context) ApiGetTrashFolderRequest {
	return ApiGetTrashFolderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FolderContentIntegerWrapper
func (a *FilesFoldersAPIService) GetTrashFolderExecute(r ApiGetTrashFolderRequest) (*FolderContentIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FolderContentIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.GetTrashFolder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/@trash"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userIdOrGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userIdOrGroupId", r.userIdOrGroupId, "form", "")
			}
	if r.filterType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterType", r.filterType, "form", "")
			}
	if r.applyFilterOption != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "applyFilterOption", r.applyFilterOption, "form", "")
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

type ApiInsertFileRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	folderId int32
	insertFileFile *os.File
	insertFileTitle *string
	insertFileCreateNewIfExist *bool
	insertFileKeepConvertStatus *bool
	insertFileStreamCanRead *bool
	insertFileStreamCanWrite *bool
	insertFileStreamCanSeek *bool
	insertFileStreamCanTimeout *bool
	insertFileStreamLength *int64
	insertFileStreamPosition *int64
	insertFileStreamReadTimeout *int32
	insertFileStreamWriteTimeout *int32
}

// The file to be inserted.
func (r ApiInsertFileRequest) InsertFileFile(insertFileFile *os.File) ApiInsertFileRequest {	r.insertFileFile = insertFileFile
	return r
}

// The file title to be inserted.
func (r ApiInsertFileRequest) InsertFileTitle(insertFileTitle string) ApiInsertFileRequest {	r.insertFileTitle = &insertFileTitle
	return r
}

// Specifies whether to create a new file if it already exists or not.
func (r ApiInsertFileRequest) InsertFileCreateNewIfExist(insertFileCreateNewIfExist bool) ApiInsertFileRequest {	r.insertFileCreateNewIfExist = &insertFileCreateNewIfExist
	return r
}

// Specifies whether to keep the file converting status or not.
func (r ApiInsertFileRequest) InsertFileKeepConvertStatus(insertFileKeepConvertStatus bool) ApiInsertFileRequest {	r.insertFileKeepConvertStatus = &insertFileKeepConvertStatus
	return r
}

func (r ApiInsertFileRequest) InsertFileStreamCanRead(insertFileStreamCanRead bool) ApiInsertFileRequest {	r.insertFileStreamCanRead = &insertFileStreamCanRead
	return r
}

func (r ApiInsertFileRequest) InsertFileStreamCanWrite(insertFileStreamCanWrite bool) ApiInsertFileRequest {	r.insertFileStreamCanWrite = &insertFileStreamCanWrite
	return r
}

func (r ApiInsertFileRequest) InsertFileStreamCanSeek(insertFileStreamCanSeek bool) ApiInsertFileRequest {	r.insertFileStreamCanSeek = &insertFileStreamCanSeek
	return r
}

func (r ApiInsertFileRequest) InsertFileStreamCanTimeout(insertFileStreamCanTimeout bool) ApiInsertFileRequest {	r.insertFileStreamCanTimeout = &insertFileStreamCanTimeout
	return r
}

func (r ApiInsertFileRequest) InsertFileStreamLength(insertFileStreamLength int64) ApiInsertFileRequest {	r.insertFileStreamLength = &insertFileStreamLength
	return r
}

func (r ApiInsertFileRequest) InsertFileStreamPosition(insertFileStreamPosition int64) ApiInsertFileRequest {	r.insertFileStreamPosition = &insertFileStreamPosition
	return r
}

func (r ApiInsertFileRequest) InsertFileStreamReadTimeout(insertFileStreamReadTimeout int32) ApiInsertFileRequest {	r.insertFileStreamReadTimeout = &insertFileStreamReadTimeout
	return r
}

func (r ApiInsertFileRequest) InsertFileStreamWriteTimeout(insertFileStreamWriteTimeout int32) ApiInsertFileRequest {	r.insertFileStreamWriteTimeout = &insertFileStreamWriteTimeout
	return r
}

func (r ApiInsertFileRequest) Execute() (*FileIntegerWrapper, *http.Response, error) {
	return r.ApiService.InsertFileExecute(r)
}

// InsertFile Insert a file
//
// Inserts a file specified in the request to the selected folder by single file uploading.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/insert-file/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The folder ID for inserting a file.
// @return ApiInsertFileRequest
func (a *FilesFoldersAPIService) InsertFile(ctx context.Context, folderId int32) ApiInsertFileRequest {
	return ApiInsertFileRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return FileIntegerWrapper
func (a *FilesFoldersAPIService) InsertFileExecute(r ApiInsertFileRequest) (*FileIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.InsertFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/{folderId}/insert"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)

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
	var insertFileFileLocalVarFormFileName string
	var insertFileFileLocalVarFileName     string
	var insertFileFileLocalVarFileBytes    []byte

	insertFileFileLocalVarFormFileName = "InsertFile.File"
	insertFileFileLocalVarFile := r.insertFileFile

	if insertFileFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(insertFileFileLocalVarFile)

		insertFileFileLocalVarFileBytes = fbs
		insertFileFileLocalVarFileName = insertFileFileLocalVarFile.Name()
		insertFileFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: insertFileFileLocalVarFileBytes, fileName: insertFileFileLocalVarFileName, formFileName: insertFileFileLocalVarFormFileName})
	}
	if r.insertFileTitle != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "InsertFile.Title", r.insertFileTitle, "form", "")
	}
	if r.insertFileCreateNewIfExist != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "InsertFile.CreateNewIfExist", r.insertFileCreateNewIfExist, "form", "")
	}
	if r.insertFileKeepConvertStatus != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "InsertFile.KeepConvertStatus", r.insertFileKeepConvertStatus, "form", "")
	}
	if r.insertFileStreamCanRead != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "InsertFile.Stream.CanRead", r.insertFileStreamCanRead, "form", "")
	}
	if r.insertFileStreamCanWrite != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "InsertFile.Stream.CanWrite", r.insertFileStreamCanWrite, "form", "")
	}
	if r.insertFileStreamCanSeek != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "InsertFile.Stream.CanSeek", r.insertFileStreamCanSeek, "form", "")
	}
	if r.insertFileStreamCanTimeout != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "InsertFile.Stream.CanTimeout", r.insertFileStreamCanTimeout, "form", "")
	}
	if r.insertFileStreamLength != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "InsertFile.Stream.Length", r.insertFileStreamLength, "form", "")
	}
	if r.insertFileStreamPosition != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "InsertFile.Stream.Position", r.insertFileStreamPosition, "form", "")
	}
	if r.insertFileStreamReadTimeout != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "InsertFile.Stream.ReadTimeout", r.insertFileStreamReadTimeout, "form", "")
	}
	if r.insertFileStreamWriteTimeout != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "InsertFile.Stream.WriteTimeout", r.insertFileStreamWriteTimeout, "form", "")
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

type ApiInsertFileToMyFromBodyRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	file *os.File
	title *string
	createNewIfExist *bool
	keepConvertStatus *bool
	streamCanRead *bool
	streamCanWrite *bool
	streamCanSeek *bool
	streamCanTimeout *bool
	streamLength *int64
	streamPosition *int64
	streamReadTimeout *int32
	streamWriteTimeout *int32
}

// The file to be inserted.
func (r ApiInsertFileToMyFromBodyRequest) File(file *os.File) ApiInsertFileToMyFromBodyRequest {	r.file = file
	return r
}

// The file title to be inserted.
func (r ApiInsertFileToMyFromBodyRequest) Title(title string) ApiInsertFileToMyFromBodyRequest {	r.title = &title
	return r
}

// Specifies whether to create a new file if it already exists or not.
func (r ApiInsertFileToMyFromBodyRequest) CreateNewIfExist(createNewIfExist bool) ApiInsertFileToMyFromBodyRequest {	r.createNewIfExist = &createNewIfExist
	return r
}

// Specifies whether to keep the file converting status or not.
func (r ApiInsertFileToMyFromBodyRequest) KeepConvertStatus(keepConvertStatus bool) ApiInsertFileToMyFromBodyRequest {	r.keepConvertStatus = &keepConvertStatus
	return r
}

func (r ApiInsertFileToMyFromBodyRequest) StreamCanRead(streamCanRead bool) ApiInsertFileToMyFromBodyRequest {	r.streamCanRead = &streamCanRead
	return r
}

func (r ApiInsertFileToMyFromBodyRequest) StreamCanWrite(streamCanWrite bool) ApiInsertFileToMyFromBodyRequest {	r.streamCanWrite = &streamCanWrite
	return r
}

func (r ApiInsertFileToMyFromBodyRequest) StreamCanSeek(streamCanSeek bool) ApiInsertFileToMyFromBodyRequest {	r.streamCanSeek = &streamCanSeek
	return r
}

func (r ApiInsertFileToMyFromBodyRequest) StreamCanTimeout(streamCanTimeout bool) ApiInsertFileToMyFromBodyRequest {	r.streamCanTimeout = &streamCanTimeout
	return r
}

func (r ApiInsertFileToMyFromBodyRequest) StreamLength(streamLength int64) ApiInsertFileToMyFromBodyRequest {	r.streamLength = &streamLength
	return r
}

func (r ApiInsertFileToMyFromBodyRequest) StreamPosition(streamPosition int64) ApiInsertFileToMyFromBodyRequest {	r.streamPosition = &streamPosition
	return r
}

func (r ApiInsertFileToMyFromBodyRequest) StreamReadTimeout(streamReadTimeout int32) ApiInsertFileToMyFromBodyRequest {	r.streamReadTimeout = &streamReadTimeout
	return r
}

func (r ApiInsertFileToMyFromBodyRequest) StreamWriteTimeout(streamWriteTimeout int32) ApiInsertFileToMyFromBodyRequest {	r.streamWriteTimeout = &streamWriteTimeout
	return r
}

func (r ApiInsertFileToMyFromBodyRequest) Execute() (*FileIntegerWrapper, *http.Response, error) {
	return r.ApiService.InsertFileToMyFromBodyExecute(r)
}

// InsertFileToMyFromBody Insert a file to the My documents section
//
// Inserts a file specified in the request to the My documents section by single file uploading.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/insert-file-to-my-from-body/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiInsertFileToMyFromBodyRequest
func (a *FilesFoldersAPIService) InsertFileToMyFromBody(ctx context.Context) ApiInsertFileToMyFromBodyRequest {
	return ApiInsertFileToMyFromBodyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FileIntegerWrapper
func (a *FilesFoldersAPIService) InsertFileToMyFromBodyExecute(r ApiInsertFileToMyFromBodyRequest) (*FileIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.InsertFileToMyFromBody")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/@my/insert"

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
	if r.title != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "Title", r.title, "form", "")
	}
	if r.createNewIfExist != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "CreateNewIfExist", r.createNewIfExist, "form", "")
	}
	if r.keepConvertStatus != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "KeepConvertStatus", r.keepConvertStatus, "form", "")
	}
	if r.streamCanRead != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "Stream.CanRead", r.streamCanRead, "form", "")
	}
	if r.streamCanWrite != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "Stream.CanWrite", r.streamCanWrite, "form", "")
	}
	if r.streamCanSeek != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "Stream.CanSeek", r.streamCanSeek, "form", "")
	}
	if r.streamCanTimeout != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "Stream.CanTimeout", r.streamCanTimeout, "form", "")
	}
	if r.streamLength != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "Stream.Length", r.streamLength, "form", "")
	}
	if r.streamPosition != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "Stream.Position", r.streamPosition, "form", "")
	}
	if r.streamReadTimeout != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "Stream.ReadTimeout", r.streamReadTimeout, "form", "")
	}
	if r.streamWriteTimeout != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "Stream.WriteTimeout", r.streamWriteTimeout, "form", "")
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

type ApiRenameFolderRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	folderId int32
	createFolder *CreateFolder
}

// The parameters for creating a folder.
func (r ApiRenameFolderRequest) CreateFolder(createFolder CreateFolder) ApiRenameFolderRequest {	r.createFolder = &createFolder
	return r
}

func (r ApiRenameFolderRequest) Execute() (*FolderIntegerWrapper, *http.Response, error) {
	return r.ApiService.RenameFolderExecute(r)
}

// RenameFolder Rename a folder
//
// Renames the selected folder with a new title specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/rename-folder/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The folder ID for the folder creation.
// @return ApiRenameFolderRequest
func (a *FilesFoldersAPIService) RenameFolder(ctx context.Context, folderId int32) ApiRenameFolderRequest {
	return ApiRenameFolderRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return FolderIntegerWrapper
func (a *FilesFoldersAPIService) RenameFolderExecute(r ApiRenameFolderRequest) (*FolderIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FolderIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.RenameFolder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/folder/{folderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createFolder == nil {
		return localVarReturnValue, nil, reportError("createFolder is required and must be specified")
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
	localVarPostBody = r.createFolder
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

type ApiSetFolderOrderRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	folderId int32
	orderRequestDto *OrderRequestDto
}

// The folder order information.
func (r ApiSetFolderOrderRequest) OrderRequestDto(orderRequestDto OrderRequestDto) ApiSetFolderOrderRequest {	r.orderRequestDto = &orderRequestDto
	return r
}

func (r ApiSetFolderOrderRequest) Execute() (*FolderIntegerWrapper, *http.Response, error) {
	return r.ApiService.SetFolderOrderExecute(r)
}

// SetFolderOrder Set folder order
//
// Sets the order of a folder with ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/set-folder-order/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The folder unique identifier.
// @return ApiSetFolderOrderRequest
func (a *FilesFoldersAPIService) SetFolderOrder(ctx context.Context, folderId int32) ApiSetFolderOrderRequest {
	return ApiSetFolderOrderRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return FolderIntegerWrapper
func (a *FilesFoldersAPIService) SetFolderOrderExecute(r ApiSetFolderOrderRequest) (*FolderIntegerWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FolderIntegerWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.SetFolderOrder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/folder/{folderId}/order"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)

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
	localVarPostBody = r.orderRequestDto
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

type ApiSetFolderPrimaryExternalLinkRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	id int32
	folderLinkRequest *FolderLinkRequest
}

// The folder link parameters.
func (r ApiSetFolderPrimaryExternalLinkRequest) FolderLinkRequest(folderLinkRequest FolderLinkRequest) ApiSetFolderPrimaryExternalLinkRequest {	r.folderLinkRequest = &folderLinkRequest
	return r
}

func (r ApiSetFolderPrimaryExternalLinkRequest) Execute() (*FileShareWrapper, *http.Response, error) {
	return r.ApiService.SetFolderPrimaryExternalLinkExecute(r)
}

// SetFolderPrimaryExternalLink Set the folder external link
//
// Sets the folder external link with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/set-folder-primary-external-link/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The folder ID.
// @return ApiSetFolderPrimaryExternalLinkRequest
func (a *FilesFoldersAPIService) SetFolderPrimaryExternalLink(ctx context.Context, id int32) ApiSetFolderPrimaryExternalLinkRequest {
	return ApiSetFolderPrimaryExternalLinkRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return FileShareWrapper
func (a *FilesFoldersAPIService) SetFolderPrimaryExternalLinkExecute(r ApiSetFolderPrimaryExternalLinkRequest) (*FileShareWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileShareWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.SetFolderPrimaryExternalLink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/folder/{id}/links"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.folderLinkRequest == nil {
		return localVarReturnValue, nil, reportError("folderLinkRequest is required and must be specified")
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
	localVarPostBody = r.folderLinkRequest
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

type ApiUploadFileRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	folderId int32
	createNewIfExist *bool
	storeOriginalFile *bool
	keepConvertStatus *bool
	file *os.File
}

// Specifies whether to create the new file if it already exists or not.
func (r ApiUploadFileRequest) CreateNewIfExist(createNewIfExist bool) ApiUploadFileRequest {	r.createNewIfExist = &createNewIfExist
	return r
}

// Specifies whether to upload documents in the original formats as well or not.
func (r ApiUploadFileRequest) StoreOriginalFile(storeOriginalFile bool) ApiUploadFileRequest {	r.storeOriginalFile = &storeOriginalFile
	return r
}

// Specifies whether to keep the file converting status or not.
func (r ApiUploadFileRequest) KeepConvertStatus(keepConvertStatus bool) ApiUploadFileRequest {	r.keepConvertStatus = &keepConvertStatus
	return r
}

// The file to be uploaded.
func (r ApiUploadFileRequest) File(file *os.File) ApiUploadFileRequest {	r.file = file
	return r
}

func (r ApiUploadFileRequest) Execute() (*FileIntegerArrayWrapper, *http.Response, error) {
	return r.ApiService.UploadFileExecute(r)
}

// UploadFile Upload a file
//
// Uploads a file specified in the request to the selected folder by single file uploading or standart multipart/form-data method.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/upload-file/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param folderId The folder ID to upload a file.
// @return ApiUploadFileRequest
func (a *FilesFoldersAPIService) UploadFile(ctx context.Context, folderId int32) ApiUploadFileRequest {
	return ApiUploadFileRequest{
		ApiService: a,
		ctx: ctx,
		folderId: folderId,
	}
}

// Execute executes the request
//  @return FileIntegerArrayWrapper
func (a *FilesFoldersAPIService) UploadFileExecute(r ApiUploadFileRequest) (*FileIntegerArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileIntegerArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.UploadFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/{folderId}/upload"
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createNewIfExist != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "createNewIfExist", r.createNewIfExist, "form", "")
			}
	if r.storeOriginalFile != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "storeOriginalFile", r.storeOriginalFile, "form", "")
			}
	if r.keepConvertStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "keepConvertStatus", r.keepConvertStatus, "form", "")
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

type ApiUploadFileToMyRequest struct {
	ctx context.Context
	ApiService *FilesFoldersAPIService
	createNewIfExist *bool
	storeOriginalFile *bool
	keepConvertStatus *bool
	file *os.File
}

// Specifies whether to create the new file if it already exists or not.
func (r ApiUploadFileToMyRequest) CreateNewIfExist(createNewIfExist bool) ApiUploadFileToMyRequest {	r.createNewIfExist = &createNewIfExist
	return r
}

// Specifies whether to upload documents in the original formats as well or not.
func (r ApiUploadFileToMyRequest) StoreOriginalFile(storeOriginalFile bool) ApiUploadFileToMyRequest {	r.storeOriginalFile = &storeOriginalFile
	return r
}

// Specifies whether to keep the file converting status or not.
func (r ApiUploadFileToMyRequest) KeepConvertStatus(keepConvertStatus bool) ApiUploadFileToMyRequest {	r.keepConvertStatus = &keepConvertStatus
	return r
}

// The file to be uploaded.
func (r ApiUploadFileToMyRequest) File(file *os.File) ApiUploadFileToMyRequest {	r.file = file
	return r
}

func (r ApiUploadFileToMyRequest) Execute() (*FileIntegerArrayWrapper, *http.Response, error) {
	return r.ApiService.UploadFileToMyExecute(r)
}

// UploadFileToMy Upload a file to the My documents section
//
// Uploads a file specified in the request to the My documents section by single file uploading or standart multipart/form-data method.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/upload-file-to-my/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiUploadFileToMyRequest
func (a *FilesFoldersAPIService) UploadFileToMy(ctx context.Context) ApiUploadFileToMyRequest {
	return ApiUploadFileToMyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FileIntegerArrayWrapper
func (a *FilesFoldersAPIService) UploadFileToMyExecute(r ApiUploadFileToMyRequest) (*FileIntegerArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FileIntegerArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesFoldersAPIService.UploadFileToMy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/files/@my/upload"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createNewIfExist != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "createNewIfExist", r.createNewIfExist, "form", "")
			}
	if r.storeOriginalFile != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "storeOriginalFile", r.storeOriginalFile, "form", "")
			}
	if r.keepConvertStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "keepConvertStatus", r.keepConvertStatus, "form", "")
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
