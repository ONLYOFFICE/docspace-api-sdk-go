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


// AIAttachmentsAPIService AIAttachmentsAPI service
type AIAttachmentsAPIService service

type ApiAiAttachmentsDeleteRequest struct {
	ctx context.Context
	ApiService *AIAttachmentsAPIService
	body *string
}

func (r ApiAiAttachmentsDeleteRequest) Body(body string) ApiAiAttachmentsDeleteRequest {	r.body = &body
	return r
}

func (r ApiAiAttachmentsDeleteRequest) Execute() (*AiSuccessResponse, *http.Response, error) {
	return r.ApiService.AiAttachmentsDeleteExecute(r)
}

// AiAttachmentsDelete Delete
//
// Permanently deletes one attachment, whether it is still a draft or already linked to a message.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-attachments-delete/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAttachmentsDeleteRequest
func (a *AIAttachmentsAPIService) AiAttachmentsDelete(ctx context.Context) ApiAiAttachmentsDeleteRequest {
	return ApiAiAttachmentsDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiSuccessResponse
func (a *AIAttachmentsAPIService) AiAttachmentsDeleteExecute(r ApiAiAttachmentsDeleteRequest) (*AiSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAttachmentsAPIService.AiAttachmentsDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/attachments/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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
			var v AiErrorResponse
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

type ApiAiAttachmentsDeleteManyRequest struct {
	ctx context.Context
	ApiService *AIAttachmentsAPIService
	requestBody *[]string
}

func (r ApiAiAttachmentsDeleteManyRequest) RequestBody(requestBody []string) ApiAiAttachmentsDeleteManyRequest {	r.requestBody = &requestBody
	return r
}

func (r ApiAiAttachmentsDeleteManyRequest) Execute() (*AiSuccessResponse, *http.Response, error) {
	return r.ApiService.AiAttachmentsDeleteManyExecute(r)
}

// AiAttachmentsDeleteMany Delete many
//
// Permanently deletes a batch of attachments in a single round trip.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-attachments-delete-many/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAttachmentsDeleteManyRequest
func (a *AIAttachmentsAPIService) AiAttachmentsDeleteMany(ctx context.Context) ApiAiAttachmentsDeleteManyRequest {
	return ApiAiAttachmentsDeleteManyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiSuccessResponse
func (a *AIAttachmentsAPIService) AiAttachmentsDeleteManyExecute(r ApiAiAttachmentsDeleteManyRequest) (*AiSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAttachmentsAPIService.AiAttachmentsDeleteMany")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/attachments/delete-many"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestBody == nil {
		return localVarReturnValue, nil, reportError("requestBody is required and must be specified")
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
	localVarPostBody = r.requestBody
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
			var v AiErrorResponse
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

type ApiAiAttachmentsGetRequest struct {
	ctx context.Context
	ApiService *AIAttachmentsAPIService
	body *string
}

func (r ApiAiAttachmentsGetRequest) Body(body string) ApiAiAttachmentsGetRequest {	r.body = &body
	return r
}

func (r ApiAiAttachmentsGetRequest) Execute() (*AiAttachment, *http.Response, error) {
	return r.ApiService.AiAttachmentsGetExecute(r)
}

// AiAttachmentsGet Get
//
// Returns one attachment by identifier.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-attachments-get/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAttachmentsGetRequest
func (a *AIAttachmentsAPIService) AiAttachmentsGet(ctx context.Context) ApiAiAttachmentsGetRequest {
	return ApiAiAttachmentsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiAttachment
func (a *AIAttachmentsAPIService) AiAttachmentsGetExecute(r ApiAiAttachmentsGetRequest) (*AiAttachment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiAttachment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAttachmentsAPIService.AiAttachmentsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/attachments/get"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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
			var v AiErrorResponse
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

type ApiAiAttachmentsGetManyRequest struct {
	ctx context.Context
	ApiService *AIAttachmentsAPIService
	requestBody *[]string
}

func (r ApiAiAttachmentsGetManyRequest) RequestBody(requestBody []string) ApiAiAttachmentsGetManyRequest {	r.requestBody = &requestBody
	return r
}

func (r ApiAiAttachmentsGetManyRequest) Execute() ([]*AiAttachment, *http.Response, error) {
	return r.ApiService.AiAttachmentsGetManyExecute(r)
}

// AiAttachmentsGetMany Get many
//
// Returns a batch of attachments, preserving the requested order; an identifier that no longer exists comes back empty.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-attachments-get-many/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAttachmentsGetManyRequest
func (a *AIAttachmentsAPIService) AiAttachmentsGetMany(ctx context.Context) ApiAiAttachmentsGetManyRequest {
	return ApiAiAttachmentsGetManyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []*AiAttachment
func (a *AIAttachmentsAPIService) AiAttachmentsGetManyExecute(r ApiAiAttachmentsGetManyRequest) ([]*AiAttachment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []*AiAttachment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAttachmentsAPIService.AiAttachmentsGetMany")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/attachments/get-many"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestBody == nil {
		return localVarReturnValue, nil, reportError("requestBody is required and must be specified")
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
	localVarPostBody = r.requestBody
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
			var v AiErrorResponse
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

type ApiAiAttachmentsLinkToMessageRequest struct {
	ctx context.Context
	ApiService *AIAttachmentsAPIService
	aiAttachmentsLinkToMessageRequest *AiAttachmentsLinkToMessageRequest
}

func (r ApiAiAttachmentsLinkToMessageRequest) AiAttachmentsLinkToMessageRequest(aiAttachmentsLinkToMessageRequest AiAttachmentsLinkToMessageRequest) ApiAiAttachmentsLinkToMessageRequest {	r.aiAttachmentsLinkToMessageRequest = &aiAttachmentsLinkToMessageRequest
	return r
}

func (r ApiAiAttachmentsLinkToMessageRequest) Execute() (*AiSuccessResponse, *http.Response, error) {
	return r.ApiService.AiAttachmentsLinkToMessageExecute(r)
}

// AiAttachmentsLinkToMessage Link to message
//
// Binds draft attachments to the chat message that owns them, once that message has been persisted, so deleting the message removes them too. Identifiers that no longer exist are skipped.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-attachments-link-to-message/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAttachmentsLinkToMessageRequest
func (a *AIAttachmentsAPIService) AiAttachmentsLinkToMessage(ctx context.Context) ApiAiAttachmentsLinkToMessageRequest {
	return ApiAiAttachmentsLinkToMessageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiSuccessResponse
func (a *AIAttachmentsAPIService) AiAttachmentsLinkToMessageExecute(r ApiAiAttachmentsLinkToMessageRequest) (*AiSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAttachmentsAPIService.AiAttachmentsLinkToMessage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/attachments/link-to-message"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiAttachmentsLinkToMessageRequest == nil {
		return localVarReturnValue, nil, reportError("aiAttachmentsLinkToMessageRequest is required and must be specified")
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
	localVarPostBody = r.aiAttachmentsLinkToMessageRequest
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
			var v AiErrorResponse
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

type ApiAiAttachmentsSaveFileRequest struct {
	ctx context.Context
	ApiService *AIAttachmentsAPIService
	aiAttachmentsSaveFileRequest *AiAttachmentsSaveFileRequest
}

func (r ApiAiAttachmentsSaveFileRequest) AiAttachmentsSaveFileRequest(aiAttachmentsSaveFileRequest AiAttachmentsSaveFileRequest) ApiAiAttachmentsSaveFileRequest {	r.aiAttachmentsSaveFileRequest = &aiAttachmentsSaveFileRequest
	return r
}

func (r ApiAiAttachmentsSaveFileRequest) Execute() (*AiAttachment, *http.Response, error) {
	return r.ApiService.AiAttachmentsSaveFileExecute(r)
}

// AiAttachmentsSaveFile Save file
//
// Stores one file attachment as a draft, carrying the host-extracted text of the file. Prefer `save-files-many` when adding several files at once so they land as one round trip.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-attachments-save-file/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAttachmentsSaveFileRequest
func (a *AIAttachmentsAPIService) AiAttachmentsSaveFile(ctx context.Context) ApiAiAttachmentsSaveFileRequest {
	return ApiAiAttachmentsSaveFileRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiAttachment
func (a *AIAttachmentsAPIService) AiAttachmentsSaveFileExecute(r ApiAiAttachmentsSaveFileRequest) (*AiAttachment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiAttachment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAttachmentsAPIService.AiAttachmentsSaveFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/attachments/save-file"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiAttachmentsSaveFileRequest == nil {
		return localVarReturnValue, nil, reportError("aiAttachmentsSaveFileRequest is required and must be specified")
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
	localVarPostBody = r.aiAttachmentsSaveFileRequest
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
			var v AiErrorResponse
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

type ApiAiAttachmentsSaveFilesManyRequest struct {
	ctx context.Context
	ApiService *AIAttachmentsAPIService
	aiAttachmentsSaveFilesManyRequest *AiAttachmentsSaveFilesManyRequest
}

func (r ApiAiAttachmentsSaveFilesManyRequest) AiAttachmentsSaveFilesManyRequest(aiAttachmentsSaveFilesManyRequest AiAttachmentsSaveFilesManyRequest) ApiAiAttachmentsSaveFilesManyRequest {	r.aiAttachmentsSaveFilesManyRequest = &aiAttachmentsSaveFilesManyRequest
	return r
}

func (r ApiAiAttachmentsSaveFilesManyRequest) Execute() ([]AiAttachment, *http.Response, error) {
	return r.ApiService.AiAttachmentsSaveFilesManyExecute(r)
}

// AiAttachmentsSaveFilesMany Save files many
//
// Stores a batch of file attachments as drafts in a single round trip. The returned records keep the order of the input.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-attachments-save-files-many/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAttachmentsSaveFilesManyRequest
func (a *AIAttachmentsAPIService) AiAttachmentsSaveFilesMany(ctx context.Context) ApiAiAttachmentsSaveFilesManyRequest {
	return ApiAiAttachmentsSaveFilesManyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []AiAttachment
func (a *AIAttachmentsAPIService) AiAttachmentsSaveFilesManyExecute(r ApiAiAttachmentsSaveFilesManyRequest) ([]AiAttachment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AiAttachment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAttachmentsAPIService.AiAttachmentsSaveFilesMany")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/attachments/save-files-many"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiAttachmentsSaveFilesManyRequest == nil {
		return localVarReturnValue, nil, reportError("aiAttachmentsSaveFilesManyRequest is required and must be specified")
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
	localVarPostBody = r.aiAttachmentsSaveFilesManyRequest
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
			var v AiErrorResponse
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
