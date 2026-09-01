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


// AIThreadsAPIService AIThreadsAPI service
type AIThreadsAPIService service

type ApiAiThreadsAppendUserMessageRequest struct {
	ctx context.Context
	ApiService *AIThreadsAPIService
	aiThreadsAppendUserMessageRequest *AiThreadsAppendUserMessageRequest
}

func (r ApiAiThreadsAppendUserMessageRequest) AiThreadsAppendUserMessageRequest(aiThreadsAppendUserMessageRequest AiThreadsAppendUserMessageRequest) ApiAiThreadsAppendUserMessageRequest {	r.aiThreadsAppendUserMessageRequest = &aiThreadsAppendUserMessageRequest
	return r
}

func (r ApiAiThreadsAppendUserMessageRequest) Execute() (*AiThreadMessageLike, *http.Response, error) {
	return r.ApiService.AiThreadsAppendUserMessageExecute(r)
}

// AiThreadsAppendUserMessage Append user message
//
// Persists a user message in a thread and bumps the thread's last-edit date so it resurfaces in the sidebar. Optionally rebinds the thread to another profile when the model changed mid-conversation.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-append-user-message/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiThreadsAppendUserMessageRequest
func (a *AIThreadsAPIService) AiThreadsAppendUserMessage(ctx context.Context) ApiAiThreadsAppendUserMessageRequest {
	return ApiAiThreadsAppendUserMessageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiThreadMessageLike
func (a *AIThreadsAPIService) AiThreadsAppendUserMessageExecute(r ApiAiThreadsAppendUserMessageRequest) (*AiThreadMessageLike, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiThreadMessageLike
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIThreadsAPIService.AiThreadsAppendUserMessage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/threads/append-user-message"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiThreadsAppendUserMessageRequest == nil {
		return localVarReturnValue, nil, reportError("aiThreadsAppendUserMessageRequest is required and must be specified")
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
	localVarPostBody = r.aiThreadsAppendUserMessageRequest
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

type ApiAiThreadsClearMessagesRequest struct {
	ctx context.Context
	ApiService *AIThreadsAPIService
	body *string
}

func (r ApiAiThreadsClearMessagesRequest) Body(body string) ApiAiThreadsClearMessagesRequest {	r.body = &body
	return r
}

func (r ApiAiThreadsClearMessagesRequest) Execute() (*AiSuccessResponse, *http.Response, error) {
	return r.ApiService.AiThreadsClearMessagesExecute(r)
}

// AiThreadsClearMessages Clear messages
//
// Drops every message of a thread while keeping the thread itself, and bumps its last-edit date.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-clear-messages/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiThreadsClearMessagesRequest
func (a *AIThreadsAPIService) AiThreadsClearMessages(ctx context.Context) ApiAiThreadsClearMessagesRequest {
	return ApiAiThreadsClearMessagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiSuccessResponse
func (a *AIThreadsAPIService) AiThreadsClearMessagesExecute(r ApiAiThreadsClearMessagesRequest) (*AiSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIThreadsAPIService.AiThreadsClearMessages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/threads/clear-messages"

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

type ApiAiThreadsCreateRequest struct {
	ctx context.Context
	ApiService *AIThreadsAPIService
	aiThreadsCreateRequest *AiThreadsCreateRequest
}

func (r ApiAiThreadsCreateRequest) AiThreadsCreateRequest(aiThreadsCreateRequest AiThreadsCreateRequest) ApiAiThreadsCreateRequest {	r.aiThreadsCreateRequest = &aiThreadsCreateRequest
	return r
}

func (r ApiAiThreadsCreateRequest) Execute() (*AiThread, *http.Response, error) {
	return r.ApiService.AiThreadsCreateExecute(r)
}

// AiThreadsCreate Create
//
// Creates a chat thread with a caller-supplied title. Use `open-or-create` instead when the title should be generated from the first user message.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-create/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiThreadsCreateRequest
func (a *AIThreadsAPIService) AiThreadsCreate(ctx context.Context) ApiAiThreadsCreateRequest {
	return ApiAiThreadsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiThread
func (a *AIThreadsAPIService) AiThreadsCreateExecute(r ApiAiThreadsCreateRequest) (*AiThread, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiThread
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIThreadsAPIService.AiThreadsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/threads/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiThreadsCreateRequest == nil {
		return localVarReturnValue, nil, reportError("aiThreadsCreateRequest is required and must be specified")
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
	localVarPostBody = r.aiThreadsCreateRequest
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

type ApiAiThreadsDeleteRequest struct {
	ctx context.Context
	ApiService *AIThreadsAPIService
	body *string
}

func (r ApiAiThreadsDeleteRequest) Body(body string) ApiAiThreadsDeleteRequest {	r.body = &body
	return r
}

func (r ApiAiThreadsDeleteRequest) Execute() (*AiSuccessResponse, *http.Response, error) {
	return r.ApiService.AiThreadsDeleteExecute(r)
}

// AiThreadsDelete Delete
//
// Deletes a chat thread together with its messages.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-delete/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiThreadsDeleteRequest
func (a *AIThreadsAPIService) AiThreadsDelete(ctx context.Context) ApiAiThreadsDeleteRequest {
	return ApiAiThreadsDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiSuccessResponse
func (a *AIThreadsAPIService) AiThreadsDeleteExecute(r ApiAiThreadsDeleteRequest) (*AiSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIThreadsAPIService.AiThreadsDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/threads/delete"

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

type ApiAiThreadsDeleteMessageRequest struct {
	ctx context.Context
	ApiService *AIThreadsAPIService
	body *string
}

func (r ApiAiThreadsDeleteMessageRequest) Body(body string) ApiAiThreadsDeleteMessageRequest {	r.body = &body
	return r
}

func (r ApiAiThreadsDeleteMessageRequest) Execute() (*AiSuccessResponse, *http.Response, error) {
	return r.ApiService.AiThreadsDeleteMessageExecute(r)
}

// AiThreadsDeleteMessage Delete message
//
// Deletes one chat message, leaving the rest of the thread untouched.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-delete-message/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiThreadsDeleteMessageRequest
func (a *AIThreadsAPIService) AiThreadsDeleteMessage(ctx context.Context) ApiAiThreadsDeleteMessageRequest {
	return ApiAiThreadsDeleteMessageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiSuccessResponse
func (a *AIThreadsAPIService) AiThreadsDeleteMessageExecute(r ApiAiThreadsDeleteMessageRequest) (*AiSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIThreadsAPIService.AiThreadsDeleteMessage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/threads/delete-message"

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

type ApiAiThreadsGetByIdRequest struct {
	ctx context.Context
	ApiService *AIThreadsAPIService
	threadId *string
}

// The chat thread identifier.
func (r ApiAiThreadsGetByIdRequest) ThreadId(threadId string) ApiAiThreadsGetByIdRequest {	r.threadId = &threadId
	return r
}

func (r ApiAiThreadsGetByIdRequest) Execute() (*AiThread, *http.Response, error) {
	return r.ApiService.AiThreadsGetByIdExecute(r)
}

// AiThreadsGetById Get by id
//
// Returns one chat thread, or an empty result when the identifier is unknown.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-get-by-id/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiThreadsGetByIdRequest
func (a *AIThreadsAPIService) AiThreadsGetById(ctx context.Context) ApiAiThreadsGetByIdRequest {
	return ApiAiThreadsGetByIdRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiThread
func (a *AIThreadsAPIService) AiThreadsGetByIdExecute(r ApiAiThreadsGetByIdRequest) (*AiThread, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiThread
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIThreadsAPIService.AiThreadsGetById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/threads/get-by-id"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.threadId == nil {
		return localVarReturnValue, nil, reportError("threadId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "threadId", r.threadId, "form", "")
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

type ApiAiThreadsGetMessageByIdRequest struct {
	ctx context.Context
	ApiService *AIThreadsAPIService
	messageId *string
}

// The globally unique chat message identifier.
func (r ApiAiThreadsGetMessageByIdRequest) MessageId(messageId string) ApiAiThreadsGetMessageByIdRequest {	r.messageId = &messageId
	return r
}

func (r ApiAiThreadsGetMessageByIdRequest) Execute() (*AiThreadMessageLike, *http.Response, error) {
	return r.ApiService.AiThreadsGetMessageByIdExecute(r)
}

// AiThreadsGetMessageById Get message by id
//
// Returns one chat message by its globally unique identifier.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-get-message-by-id/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiThreadsGetMessageByIdRequest
func (a *AIThreadsAPIService) AiThreadsGetMessageById(ctx context.Context) ApiAiThreadsGetMessageByIdRequest {
	return ApiAiThreadsGetMessageByIdRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiThreadMessageLike
func (a *AIThreadsAPIService) AiThreadsGetMessageByIdExecute(r ApiAiThreadsGetMessageByIdRequest) (*AiThreadMessageLike, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiThreadMessageLike
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIThreadsAPIService.AiThreadsGetMessageById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/threads/get-message-by-id"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.messageId == nil {
		return localVarReturnValue, nil, reportError("messageId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "messageId", r.messageId, "form", "")
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

type ApiAiThreadsListRequest struct {
	ctx context.Context
	ApiService *AIThreadsAPIService
	entityId *string
	count *string
	cursor *string
	query *string
}

// The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope.
func (r ApiAiThreadsListRequest) EntityId(entityId string) ApiAiThreadsListRequest {	r.entityId = &entityId
	return r
}

// The maximum number of items to return in one page.
func (r ApiAiThreadsListRequest) Count(count string) ApiAiThreadsListRequest {	r.count = &count
	return r
}

// The keyset pagination cursor: the JSON-encoded sort key of the last item already received. Omit for the first page.
func (r ApiAiThreadsListRequest) Cursor(cursor string) ApiAiThreadsListRequest {	r.cursor = &cursor
	return r
}

// The full-text query the thread list is filtered by.
func (r ApiAiThreadsListRequest) Query(query string) ApiAiThreadsListRequest {	r.query = &query
	return r
}

func (r ApiAiThreadsListRequest) Execute() ([]AiThread, *http.Response, error) {
	return r.ApiService.AiThreadsListExecute(r)
}

// AiThreadsList List
//
// Lists the chat threads of the scope, most recently edited first. Supports cursor pagination and a server-side case-insensitive title search.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-list/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiThreadsListRequest
func (a *AIThreadsAPIService) AiThreadsList(ctx context.Context) ApiAiThreadsListRequest {
	return ApiAiThreadsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []AiThread
func (a *AIThreadsAPIService) AiThreadsListExecute(r ApiAiThreadsListRequest) ([]AiThread, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AiThread
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIThreadsAPIService.AiThreadsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/threads/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.entityId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "entityId", r.entityId, "form", "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	}
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

type ApiAiThreadsOpenOrCreateRequest struct {
	ctx context.Context
	ApiService *AIThreadsAPIService
	aiThreadsOpenOrCreateRequest *AiThreadsOpenOrCreateRequest
}

func (r ApiAiThreadsOpenOrCreateRequest) AiThreadsOpenOrCreateRequest(aiThreadsOpenOrCreateRequest AiThreadsOpenOrCreateRequest) ApiAiThreadsOpenOrCreateRequest {	r.aiThreadsOpenOrCreateRequest = &aiThreadsOpenOrCreateRequest
	return r
}

func (r ApiAiThreadsOpenOrCreateRequest) Execute() (*AiOpenOrCreateResult, *http.Response, error) {
	return r.ApiService.AiThreadsOpenOrCreateExecute(r)
}

// AiThreadsOpenOrCreate Open or create
//
// Opens a chat thread and returns its history, or creates one with a title generated from the supplied first message. That first message is not persisted - the caller decides whether to follow up with `append-user-message`.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-open-or-create/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiThreadsOpenOrCreateRequest
func (a *AIThreadsAPIService) AiThreadsOpenOrCreate(ctx context.Context) ApiAiThreadsOpenOrCreateRequest {
	return ApiAiThreadsOpenOrCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiOpenOrCreateResult
func (a *AIThreadsAPIService) AiThreadsOpenOrCreateExecute(r ApiAiThreadsOpenOrCreateRequest) (*AiOpenOrCreateResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiOpenOrCreateResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIThreadsAPIService.AiThreadsOpenOrCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/threads/open-or-create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiThreadsOpenOrCreateRequest == nil {
		return localVarReturnValue, nil, reportError("aiThreadsOpenOrCreateRequest is required and must be specified")
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
	localVarPostBody = r.aiThreadsOpenOrCreateRequest
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

type ApiAiThreadsReadMessagesRequest struct {
	ctx context.Context
	ApiService *AIThreadsAPIService
	threadId *string
	count *string
	cursor *string
	direction *string
}

// The chat thread identifier.
func (r ApiAiThreadsReadMessagesRequest) ThreadId(threadId string) ApiAiThreadsReadMessagesRequest {	r.threadId = &threadId
	return r
}

// The maximum number of items to return in one page.
func (r ApiAiThreadsReadMessagesRequest) Count(count string) ApiAiThreadsReadMessagesRequest {	r.count = &count
	return r
}

// The keyset pagination cursor: the JSON-encoded sort key of the last item already received. Omit for the first page.
func (r ApiAiThreadsReadMessagesRequest) Cursor(cursor string) ApiAiThreadsReadMessagesRequest {	r.cursor = &cursor
	return r
}

// The order the message page is read in. Only desc turns the read around and pages back from the newest message; omit for the forward read.
func (r ApiAiThreadsReadMessagesRequest) Direction(direction string) ApiAiThreadsReadMessagesRequest {	r.direction = &direction
	return r
}

func (r ApiAiThreadsReadMessagesRequest) Execute() ([]AiThreadMessageLike, *http.Response, error) {
	return r.ApiService.AiThreadsReadMessagesExecute(r)
}

// AiThreadsReadMessages Read messages
//
// Reads the messages of a thread, with the same cursor pagination as the thread list.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-read-messages/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiThreadsReadMessagesRequest
func (a *AIThreadsAPIService) AiThreadsReadMessages(ctx context.Context) ApiAiThreadsReadMessagesRequest {
	return ApiAiThreadsReadMessagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []AiThreadMessageLike
func (a *AIThreadsAPIService) AiThreadsReadMessagesExecute(r ApiAiThreadsReadMessagesRequest) ([]AiThreadMessageLike, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AiThreadMessageLike
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIThreadsAPIService.AiThreadsReadMessages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/threads/read-messages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.threadId == nil {
		return localVarReturnValue, nil, reportError("threadId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "threadId", r.threadId, "form", "")
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	}
	if r.direction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direction", r.direction, "form", "")
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

type ApiAiThreadsRegenerateTitleRequest struct {
	ctx context.Context
	ApiService *AIThreadsAPIService
	aiThreadsRegenerateTitleRequest *AiThreadsRegenerateTitleRequest
}

func (r ApiAiThreadsRegenerateTitleRequest) AiThreadsRegenerateTitleRequest(aiThreadsRegenerateTitleRequest AiThreadsRegenerateTitleRequest) ApiAiThreadsRegenerateTitleRequest {	r.aiThreadsRegenerateTitleRequest = &aiThreadsRegenerateTitleRequest
	return r
}

func (r ApiAiThreadsRegenerateTitleRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.AiThreadsRegenerateTitleExecute(r)
}

// AiThreadsRegenerateTitle Regenerate title
//
// Generates a fresh title from the thread's first user message and persists it. Fails when the thread has no user message yet.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-regenerate-title/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiThreadsRegenerateTitleRequest
func (a *AIThreadsAPIService) AiThreadsRegenerateTitle(ctx context.Context) ApiAiThreadsRegenerateTitleRequest {
	return ApiAiThreadsRegenerateTitleRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *AIThreadsAPIService) AiThreadsRegenerateTitleExecute(r ApiAiThreadsRegenerateTitleRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIThreadsAPIService.AiThreadsRegenerateTitle")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/threads/regenerate-title"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiThreadsRegenerateTitleRequest == nil {
		return localVarReturnValue, nil, reportError("aiThreadsRegenerateTitleRequest is required and must be specified")
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
	localVarPostBody = r.aiThreadsRegenerateTitleRequest
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

type ApiAiThreadsRenameRequest struct {
	ctx context.Context
	ApiService *AIThreadsAPIService
	aiThreadsRenameRequest *AiThreadsRenameRequest
}

func (r ApiAiThreadsRenameRequest) AiThreadsRenameRequest(aiThreadsRenameRequest AiThreadsRenameRequest) ApiAiThreadsRenameRequest {	r.aiThreadsRenameRequest = &aiThreadsRenameRequest
	return r
}

func (r ApiAiThreadsRenameRequest) Execute() (*AiSuccessResponse, *http.Response, error) {
	return r.ApiService.AiThreadsRenameExecute(r)
}

// AiThreadsRename Rename
//
// Renames a chat thread and bumps its last-edit date so the new title shows up in the sidebar.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-rename/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiThreadsRenameRequest
func (a *AIThreadsAPIService) AiThreadsRename(ctx context.Context) ApiAiThreadsRenameRequest {
	return ApiAiThreadsRenameRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiSuccessResponse
func (a *AIThreadsAPIService) AiThreadsRenameExecute(r ApiAiThreadsRenameRequest) (*AiSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIThreadsAPIService.AiThreadsRename")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/threads/rename"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiThreadsRenameRequest == nil {
		return localVarReturnValue, nil, reportError("aiThreadsRenameRequest is required and must be specified")
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
	localVarPostBody = r.aiThreadsRenameRequest
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

type ApiAiThreadsTouchRequest struct {
	ctx context.Context
	ApiService *AIThreadsAPIService
	aiThreadsTouchRequest *AiThreadsTouchRequest
}

func (r ApiAiThreadsTouchRequest) AiThreadsTouchRequest(aiThreadsTouchRequest AiThreadsTouchRequest) ApiAiThreadsTouchRequest {	r.aiThreadsTouchRequest = &aiThreadsTouchRequest
	return r
}

func (r ApiAiThreadsTouchRequest) Execute() (*AiSuccessResponse, *http.Response, error) {
	return r.ApiService.AiThreadsTouchExecute(r)
}

// AiThreadsTouch Touch
//
// Bumps a thread's last-edit date, and optionally rebinds it to another profile, when something other than a new message - a model switch, say - should resurface it.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-touch/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiThreadsTouchRequest
func (a *AIThreadsAPIService) AiThreadsTouch(ctx context.Context) ApiAiThreadsTouchRequest {
	return ApiAiThreadsTouchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiSuccessResponse
func (a *AIThreadsAPIService) AiThreadsTouchExecute(r ApiAiThreadsTouchRequest) (*AiSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIThreadsAPIService.AiThreadsTouch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/threads/touch"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiThreadsTouchRequest == nil {
		return localVarReturnValue, nil, reportError("aiThreadsTouchRequest is required and must be specified")
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
	localVarPostBody = r.aiThreadsTouchRequest
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

type ApiAiThreadsUpdateMessageRequest struct {
	ctx context.Context
	ApiService *AIThreadsAPIService
	aiThreadsUpdateMessageRequest *AiThreadsUpdateMessageRequest
}

func (r ApiAiThreadsUpdateMessageRequest) AiThreadsUpdateMessageRequest(aiThreadsUpdateMessageRequest AiThreadsUpdateMessageRequest) ApiAiThreadsUpdateMessageRequest {	r.aiThreadsUpdateMessageRequest = &aiThreadsUpdateMessageRequest
	return r
}

func (r ApiAiThreadsUpdateMessageRequest) Execute() (*AiSuccessResponse, *http.Response, error) {
	return r.ApiService.AiThreadsUpdateMessageExecute(r)
}

// AiThreadsUpdateMessage Update message
//
// Replaces the content of a chat message - used by the edit and regenerate flows that change a message outside the streaming lifecycle.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-update-message/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiThreadsUpdateMessageRequest
func (a *AIThreadsAPIService) AiThreadsUpdateMessage(ctx context.Context) ApiAiThreadsUpdateMessageRequest {
	return ApiAiThreadsUpdateMessageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiSuccessResponse
func (a *AIThreadsAPIService) AiThreadsUpdateMessageExecute(r ApiAiThreadsUpdateMessageRequest) (*AiSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIThreadsAPIService.AiThreadsUpdateMessage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/threads/update-message"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiThreadsUpdateMessageRequest == nil {
		return localVarReturnValue, nil, reportError("aiThreadsUpdateMessageRequest is required and must be specified")
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
	localVarPostBody = r.aiThreadsUpdateMessageRequest
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
