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


// AIAIAPIService AIAIAPI service
type AIAIAPIService service

type ApiAiAiApproveToolCallRequest struct {
	ctx context.Context
	ApiService *AIAIAPIService
	aiAiApproveToolCallRequest *AiAiApproveToolCallRequest
}

func (r ApiAiAiApproveToolCallRequest) AiAiApproveToolCallRequest(aiAiApproveToolCallRequest AiAiApproveToolCallRequest) ApiAiAiApproveToolCallRequest {	r.aiAiApproveToolCallRequest = &aiAiApproveToolCallRequest
	return r
}

func (r ApiAiAiApproveToolCallRequest) Execute() (*AiChatEvent, *http.Response, error) {
	return r.ApiService.AiAiApproveToolCallExecute(r)
}

// AiAiApproveToolCall Approve tool call
//
// Resumes a chat round paused on a tool call. The supplied result is persisted onto the assistant message that issued the call and the stream continues with the augmented history.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-ai-approve-tool-call/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAiApproveToolCallRequest
func (a *AIAIAPIService) AiAiApproveToolCall(ctx context.Context) ApiAiAiApproveToolCallRequest {
	return ApiAiAiApproveToolCallRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiChatEvent
func (a *AIAIAPIService) AiAiApproveToolCallExecute(r ApiAiAiApproveToolCallRequest) (*AiChatEvent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiChatEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAIAPIService.AiAiApproveToolCall")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/ai/approve-tool-call"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiAiApproveToolCallRequest == nil {
		return localVarReturnValue, nil, reportError("aiAiApproveToolCallRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/x-ndjson", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.aiAiApproveToolCallRequest
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

type ApiAiAiDenyToolCallRequest struct {
	ctx context.Context
	ApiService *AIAIAPIService
	aiAiToolCallData *AiAiToolCallData
}

func (r ApiAiAiDenyToolCallRequest) AiAiToolCallData(aiAiToolCallData AiAiToolCallData) ApiAiAiDenyToolCallRequest {	r.aiAiToolCallData = &aiAiToolCallData
	return r
}

func (r ApiAiAiDenyToolCallRequest) Execute() (*AiChatEvent, *http.Response, error) {
	return r.ApiService.AiAiDenyToolCallExecute(r)
}

// AiAiDenyToolCall Deny tool call
//
// Denies the pending tool call and resumes the chat immediately, with `User deny tool call` standing in for the tool result.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-ai-deny-tool-call/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAiDenyToolCallRequest
func (a *AIAIAPIService) AiAiDenyToolCall(ctx context.Context) ApiAiAiDenyToolCallRequest {
	return ApiAiAiDenyToolCallRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiChatEvent
func (a *AIAIAPIService) AiAiDenyToolCallExecute(r ApiAiAiDenyToolCallRequest) (*AiChatEvent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiChatEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAIAPIService.AiAiDenyToolCall")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/ai/deny-tool-call"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiAiToolCallData == nil {
		return localVarReturnValue, nil, reportError("aiAiToolCallData is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/x-ndjson", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.aiAiToolCallData
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

type ApiAiAiRegenerateStreamRequest struct {
	ctx context.Context
	ApiService *AIAIAPIService
	aiAiRegenerateStreamRequest *AiAiRegenerateStreamRequest
}

func (r ApiAiAiRegenerateStreamRequest) AiAiRegenerateStreamRequest(aiAiRegenerateStreamRequest AiAiRegenerateStreamRequest) ApiAiAiRegenerateStreamRequest {	r.aiAiRegenerateStreamRequest = &aiAiRegenerateStreamRequest
	return r
}

func (r ApiAiAiRegenerateStreamRequest) Execute() (*AiChatEvent, *http.Response, error) {
	return r.ApiService.AiAiRegenerateStreamExecute(r)
}

// AiAiRegenerateStream Regenerate stream
//
// Re-rolls the last assistant reply in an existing thread: every message after the last user message (the previous reply plus any tool-call hops) is dropped and a fresh reply is streamed against the unchanged prompt. The thread must already exist and no title is generated.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-ai-regenerate-stream/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAiRegenerateStreamRequest
func (a *AIAIAPIService) AiAiRegenerateStream(ctx context.Context) ApiAiAiRegenerateStreamRequest {
	return ApiAiAiRegenerateStreamRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiChatEvent
func (a *AIAIAPIService) AiAiRegenerateStreamExecute(r ApiAiAiRegenerateStreamRequest) (*AiChatEvent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiChatEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAIAPIService.AiAiRegenerateStream")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/ai/regenerate-stream"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiAiRegenerateStreamRequest == nil {
		return localVarReturnValue, nil, reportError("aiAiRegenerateStreamRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/x-ndjson", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.aiAiRegenerateStreamRequest
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

type ApiAiAiSendRequest struct {
	ctx context.Context
	ApiService *AIAIAPIService
	aiAiSendRequest *AiAiSendRequest
}

func (r ApiAiAiSendRequest) AiAiSendRequest(aiAiSendRequest AiAiSendRequest) ApiAiAiSendRequest {	r.aiAiSendRequest = &aiAiSendRequest
	return r
}

func (r ApiAiAiSendRequest) Execute() (*AiThreadMessageLike, *http.Response, error) {
	return r.ApiService.AiAiSendExecute(r)
}

// AiAiSend Send
//
// Runs one AI action: the profile bound to `actionType` (falling back to the `Default` slot) is dispatched against a single-message history. Nothing is persisted - no thread, no title generation, no storage writes.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-ai-send/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAiSendRequest
func (a *AIAIAPIService) AiAiSend(ctx context.Context) ApiAiAiSendRequest {
	return ApiAiAiSendRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiThreadMessageLike
func (a *AIAIAPIService) AiAiSendExecute(r ApiAiAiSendRequest) (*AiThreadMessageLike, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiThreadMessageLike
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAIAPIService.AiAiSend")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/ai/send"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiAiSendRequest == nil {
		return localVarReturnValue, nil, reportError("aiAiSendRequest is required and must be specified")
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
	localVarPostBody = r.aiAiSendRequest
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

type ApiAiAiSendCustomRequest struct {
	ctx context.Context
	ApiService *AIAIAPIService
	aiAiSendCustomRequest *AiAiSendCustomRequest
}

func (r ApiAiAiSendCustomRequest) AiAiSendCustomRequest(aiAiSendCustomRequest AiAiSendCustomRequest) ApiAiAiSendCustomRequest {	r.aiAiSendCustomRequest = &aiAiSendCustomRequest
	return r
}

func (r ApiAiAiSendCustomRequest) Execute() (*AiThreadMessageLike, *http.Response, error) {
	return r.ApiService.AiAiSendCustomExecute(r)
}

// AiAiSendCustom Send custom
//
// Runs a free-form one-turn call against a caller-supplied system prompt. No thread, no history and no persistence. The profile is the explicit `profileId` when it resolves, otherwise the `Default` assignment slot.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-ai-send-custom/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAiSendCustomRequest
func (a *AIAIAPIService) AiAiSendCustom(ctx context.Context) ApiAiAiSendCustomRequest {
	return ApiAiAiSendCustomRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiThreadMessageLike
func (a *AIAIAPIService) AiAiSendCustomExecute(r ApiAiAiSendCustomRequest) (*AiThreadMessageLike, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiThreadMessageLike
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAIAPIService.AiAiSendCustom")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/ai/send-custom"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiAiSendCustomRequest == nil {
		return localVarReturnValue, nil, reportError("aiAiSendCustomRequest is required and must be specified")
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
	localVarPostBody = r.aiAiSendCustomRequest
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

type ApiAiAiSendWithStreamRequest struct {
	ctx context.Context
	ApiService *AIAIAPIService
	aiAiSendStreamBody *AiAiSendStreamBody
}

func (r ApiAiAiSendWithStreamRequest) AiAiSendStreamBody(aiAiSendStreamBody AiAiSendStreamBody) ApiAiAiSendWithStreamRequest {	r.aiAiSendStreamBody = &aiAiSendStreamBody
	return r
}

func (r ApiAiAiSendWithStreamRequest) Execute() (*AiChatEvent, *http.Response, error) {
	return r.ApiService.AiAiSendWithStreamExecute(r)
}

// AiAiSendWithStream Send with stream
//
// Starts a chat round and streams it back as newline-delimited `ChatEvent` objects. The thread is opened or created, the user message and the reply are persisted, a new thread gets a generated title, and a tool call pauses the round until it is approved or denied.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-ai-send-with-stream/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAiSendWithStreamRequest
func (a *AIAIAPIService) AiAiSendWithStream(ctx context.Context) ApiAiAiSendWithStreamRequest {
	return ApiAiAiSendWithStreamRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiChatEvent
func (a *AIAIAPIService) AiAiSendWithStreamExecute(r ApiAiAiSendWithStreamRequest) (*AiChatEvent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiChatEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAIAPIService.AiAiSendWithStream")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/ai/send-with-stream"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiAiSendStreamBody == nil {
		return localVarReturnValue, nil, reportError("aiAiSendStreamBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/x-ndjson", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.aiAiSendStreamBody
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

type ApiAiAiSendWithStreamOpenAIRequest struct {
	ctx context.Context
	ApiService *AIAIAPIService
	aiAiSendStreamBody *AiAiSendStreamBody
}

func (r ApiAiAiSendWithStreamOpenAIRequest) AiAiSendStreamBody(aiAiSendStreamBody AiAiSendStreamBody) ApiAiAiSendWithStreamOpenAIRequest {	r.aiAiSendStreamBody = &aiAiSendStreamBody
	return r
}

func (r ApiAiAiSendWithStreamOpenAIRequest) Execute() (*AiOpenAIStreamChunk, *http.Response, error) {
	return r.ApiService.AiAiSendWithStreamOpenAIExecute(r)
}

// AiAiSendWithStreamOpenAI Send with stream open ai
//
// The same chat round as `send-with-stream`, re-encoded as an OpenAI Chat Completions stream of `chat.completion.chunk` objects. Storage, title generation and tool-call pauses are identical - only the wire shape differs; a tool call ends the stream with `finish_reason: tool_calls`.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-ai-send-with-stream-open-ai/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAiSendWithStreamOpenAIRequest
func (a *AIAIAPIService) AiAiSendWithStreamOpenAI(ctx context.Context) ApiAiAiSendWithStreamOpenAIRequest {
	return ApiAiAiSendWithStreamOpenAIRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiOpenAIStreamChunk
func (a *AIAIAPIService) AiAiSendWithStreamOpenAIExecute(r ApiAiAiSendWithStreamOpenAIRequest) (*AiOpenAIStreamChunk, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiOpenAIStreamChunk
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAIAPIService.AiAiSendWithStreamOpenAI")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/ai/send-with-stream-openai"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiAiSendStreamBody == nil {
		return localVarReturnValue, nil, reportError("aiAiSendStreamBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/event-stream", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.aiAiSendStreamBody
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
