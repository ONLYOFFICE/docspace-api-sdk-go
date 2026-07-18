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


// AIChatAPIService AIChatAPI service
type AIChatAPIService service

type ApiContinueChatRequest struct {
	ctx context.Context
	ApiService *AIChatAPIService
	chatId string
	continueChatBody *ContinueChatBody
}

// The message and optional file attachments.
func (r ApiContinueChatRequest) ContinueChatBody(continueChatBody ContinueChatBody) ApiContinueChatRequest {	r.continueChatBody = &continueChatBody
	return r
}

func (r ApiContinueChatRequest) Execute() (*http.Response, error) {
	return r.ApiService.ContinueChatExecute(r)
}

// ContinueChat Send a message to an existing AI chat
//
// Appends a new user message to an existing chat session and streams the AI assistant's response.  The full conversation history of the chat is sent to the AI provider to maintain context.  The response is delivered as a Server-Sent Events (SSE) stream with periodic keep-alive pings.  File references can optionally be attached to provide additional context.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/continue-chat/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param chatId The unique identifier of the existing AI chat session to continue.
// @return ApiContinueChatRequest
func (a *AIChatAPIService) ContinueChat(ctx context.Context, chatId string) ApiContinueChatRequest {
	return ApiContinueChatRequest{
		ApiService: a,
		ctx: ctx,
		chatId: chatId,
	}
}

// Execute executes the request
func (a *AIChatAPIService) ContinueChatExecute(r ApiContinueChatRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIChatAPIService.ContinueChat")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/chats/{chatId}/messages"
	localVarPath = strings.Replace(localVarPath, "{"+"chatId"+"}", url.PathEscape(parameterValueToString(r.chatId, "chatId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.continueChatBody == nil {
		return nil, reportError("continueChatBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.continueChatBody
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

type ApiDeleteChatRequest struct {
	ctx context.Context
	ApiService *AIChatAPIService
	chatId string
}

func (r ApiDeleteChatRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteChatExecute(r)
}

// DeleteChat Delete an AI chat
//
// Permanently deletes an AI chat session along with all of its messages.  Only the chat owner can delete their own chat sessions. This action cannot be undone.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-chat/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param chatId The unique identifier of the AI chat session to delete.
// @return ApiDeleteChatRequest
func (a *AIChatAPIService) DeleteChat(ctx context.Context, chatId string) ApiDeleteChatRequest {
	return ApiDeleteChatRequest{
		ApiService: a,
		ctx: ctx,
		chatId: chatId,
	}
}

// Execute executes the request
func (a *AIChatAPIService) DeleteChatExecute(r ApiDeleteChatRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIChatAPIService.DeleteChat")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/chats/{chatId}"
	localVarPath = strings.Replace(localVarPath, "{"+"chatId"+"}", url.PathEscape(parameterValueToString(r.chatId, "chatId")), -1)

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

type ApiExportChatRequest struct {
	ctx context.Context
	ApiService *AIChatAPIService
	chatId string
	exportChatRequestBody *ExportChatRequestBody
}

// The export parameters including destination folder and file title.
func (r ApiExportChatRequest) ExportChatRequestBody(exportChatRequestBody ExportChatRequestBody) ApiExportChatRequest {	r.exportChatRequestBody = &exportChatRequestBody
	return r
}

func (r ApiExportChatRequest) Execute() (*http.Response, error) {
	return r.ApiService.ExportChatExecute(r)
}

// ExportChat Export AI chat messages to a file
//
// Exports the entire message history of an AI chat session and saves it as a document in the specified folder.  The exported file is created with the provided title. Only the chat owner can export their own chat sessions.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/export-chat/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param chatId The unique identifier of the AI chat session to export.
// @return ApiExportChatRequest
func (a *AIChatAPIService) ExportChat(ctx context.Context, chatId string) ApiExportChatRequest {
	return ApiExportChatRequest{
		ApiService: a,
		ctx: ctx,
		chatId: chatId,
	}
}

// Execute executes the request
func (a *AIChatAPIService) ExportChatExecute(r ApiExportChatRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIChatAPIService.ExportChat")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/chats/{chatId}/messages/export"
	localVarPath = strings.Replace(localVarPath, "{"+"chatId"+"}", url.PathEscape(parameterValueToString(r.chatId, "chatId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.exportChatRequestBody == nil {
		return nil, reportError("exportChatRequestBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.exportChatRequestBody
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

type ApiGetChatRequest struct {
	ctx context.Context
	ApiService *AIChatAPIService
	chatId string
}

func (r ApiGetChatRequest) Execute() (*ChatWrapper, *http.Response, error) {
	return r.ApiService.GetChatExecute(r)
}

// GetChat Get an AI chat by ID
//
// Retrieves the metadata of a single AI chat session, including its title, creation date, and the user who created it.  Only the chat owner can access their own chat sessions.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-chat/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param chatId The unique identifier of the AI chat session to retrieve.
// @return ApiGetChatRequest
func (a *AIChatAPIService) GetChat(ctx context.Context, chatId string) ApiGetChatRequest {
	return ApiGetChatRequest{
		ApiService: a,
		ctx: ctx,
		chatId: chatId,
	}
}

// Execute executes the request
//  @return ChatWrapper
func (a *AIChatAPIService) GetChatExecute(r ApiGetChatRequest) (*ChatWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ChatWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIChatAPIService.GetChat")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/chats/{chatId}"
	localVarPath = strings.Replace(localVarPath, "{"+"chatId"+"}", url.PathEscape(parameterValueToString(r.chatId, "chatId")), -1)

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

type ApiGetChatModelsRequest struct {
	ctx context.Context
	ApiService *AIChatAPIService
	provider *int32
}

// The optional AI provider identifier to filter models by. When set to 0, models from all providers are returned.
func (r ApiGetChatModelsRequest) Provider(provider int32) ApiGetChatModelsRequest {	r.provider = &provider
	return r
}

func (r ApiGetChatModelsRequest) Execute() (*ModelArrayWrapper, *http.Response, error) {
	return r.ApiService.GetChatModelsExecute(r)
}

// GetChatModels Get available AI models
//
// Returns the list of AI models available for chat conversations.  Optionally filters the results to models from a specific provider when the provider query parameter is specified.  Each model entry includes the provider ID, provider display name, and the model identifier.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-chat-models/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetChatModelsRequest
func (a *AIChatAPIService) GetChatModels(ctx context.Context) ApiGetChatModelsRequest {
	return ApiGetChatModelsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ModelArrayWrapper
func (a *AIChatAPIService) GetChatModelsExecute(r ApiGetChatModelsRequest) (*ModelArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIChatAPIService.GetChatModels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/chats/models"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.provider != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "provider", r.provider, "form", "")
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

type ApiGetChatsRequest struct {
	ctx context.Context
	ApiService *AIChatAPIService
	roomId int32
	startIndex *int32
	count *int32
}

// The number of items to skip before returning results (zero-based offset). Defaults to 0.
func (r ApiGetChatsRequest) StartIndex(startIndex int32) ApiGetChatsRequest {	r.startIndex = &startIndex
	return r
}

// The maximum number of items to return per page. Defaults to 100.
func (r ApiGetChatsRequest) Count(count int32) ApiGetChatsRequest {	r.count = &count
	return r
}

func (r ApiGetChatsRequest) Execute() (*ChatArrayWrapper, *http.Response, error) {
	return r.ApiService.GetChatsExecute(r)
}

// GetChats Get AI chats in a room
//
// Returns a paginated list of AI chat sessions that belong to the current user within the specified room.  Supports pagination via the startIndex and count query parameters. The total number of chats is included in the response metadata.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-chats/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param roomId The identifier of the room whose AI chat sessions are to be listed.
// @return ApiGetChatsRequest
func (a *AIChatAPIService) GetChats(ctx context.Context, roomId int32) ApiGetChatsRequest {
	return ApiGetChatsRequest{
		ApiService: a,
		ctx: ctx,
		roomId: roomId,
	}
}

// Execute executes the request
//  @return ChatArrayWrapper
func (a *AIChatAPIService) GetChatsExecute(r ApiGetChatsRequest) (*ChatArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ChatArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIChatAPIService.GetChats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/rooms/{roomId}/chats"
	localVarPath = strings.Replace(localVarPath, "{"+"roomId"+"}", url.PathEscape(parameterValueToString(r.roomId, "roomId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
			}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
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

type ApiGetMessagesRequest struct {
	ctx context.Context
	ApiService *AIChatAPIService
	chatId string
	startIndex *int32
	count *int32
}

// The number of items to skip before returning results (zero-based offset). Defaults to 0.
func (r ApiGetMessagesRequest) StartIndex(startIndex int32) ApiGetMessagesRequest {	r.startIndex = &startIndex
	return r
}

// The maximum number of items to return per page. Defaults to 100.
func (r ApiGetMessagesRequest) Count(count int32) ApiGetMessagesRequest {	r.count = &count
	return r
}

func (r ApiGetMessagesRequest) Execute() (*MessageArrayWrapper, *http.Response, error) {
	return r.ApiService.GetMessagesExecute(r)
}

// GetMessages Get messages of an AI chat
//
// Returns a paginated list of messages from an AI chat session owned by the current user.  Each message includes its role (user or assistant), content blocks (text, tool calls, attachments), and timestamp.  Supports pagination via the startIndex and count query parameters. The total number of messages is included in the response metadata.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-messages/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param chatId The unique identifier of the AI chat session whose messages are to be listed.
// @return ApiGetMessagesRequest
func (a *AIChatAPIService) GetMessages(ctx context.Context, chatId string) ApiGetMessagesRequest {
	return ApiGetMessagesRequest{
		ApiService: a,
		ctx: ctx,
		chatId: chatId,
	}
}

// Execute executes the request
//  @return MessageArrayWrapper
func (a *AIChatAPIService) GetMessagesExecute(r ApiGetMessagesRequest) (*MessageArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MessageArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIChatAPIService.GetMessages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/chats/{chatId}/messages"
	localVarPath = strings.Replace(localVarPath, "{"+"chatId"+"}", url.PathEscape(parameterValueToString(r.chatId, "chatId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
			}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
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

type ApiGetUserChatsSettingsRequest struct {
	ctx context.Context
	ApiService *AIChatAPIService
	roomId int32
}

func (r ApiGetUserChatsSettingsRequest) Execute() (*UserChatSettingsWrapper, *http.Response, error) {
	return r.ApiService.GetUserChatsSettingsExecute(r)
}

// GetUserChatsSettings Get user chat settings for a room
//
// Retrieves the current user's personal AI chat preferences for the specified room,  including whether web search is enabled for AI-assisted responses.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-user-chats-settings/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param roomId The identifier of the room whose chat settings are to be retrieved.
// @return ApiGetUserChatsSettingsRequest
func (a *AIChatAPIService) GetUserChatsSettings(ctx context.Context, roomId int32) ApiGetUserChatsSettingsRequest {
	return ApiGetUserChatsSettingsRequest{
		ApiService: a,
		ctx: ctx,
		roomId: roomId,
	}
}

// Execute executes the request
//  @return UserChatSettingsWrapper
func (a *AIChatAPIService) GetUserChatsSettingsExecute(r ApiGetUserChatsSettingsRequest) (*UserChatSettingsWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserChatSettingsWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIChatAPIService.GetUserChatsSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/rooms/{roomId}/chats/config"
	localVarPath = strings.Replace(localVarPath, "{"+"roomId"+"}", url.PathEscape(parameterValueToString(r.roomId, "roomId")), -1)

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

type ApiProvidePermissionRequest struct {
	ctx context.Context
	ApiService *AIChatAPIService
	callId string
	toolDecisionRequestBody *ToolDecisionRequestBody
}

// The permission decision parameters.
func (r ApiProvidePermissionRequest) ToolDecisionRequestBody(toolDecisionRequestBody ToolDecisionRequestBody) ApiProvidePermissionRequest {	r.toolDecisionRequestBody = &toolDecisionRequestBody
	return r
}

func (r ApiProvidePermissionRequest) Execute() (*http.Response, error) {
	return r.ApiService.ProvidePermissionExecute(r)
}

// ProvidePermission Submit a tool execution permission decision
//
// Provides the user's approval or denial decision for a pending MCP (Model Context Protocol) tool execution request.  When an AI assistant attempts to invoke an external tool that requires explicit user consent,  the client receives a permission prompt via the SSE stream. This endpoint is used to submit the user's decision  so that the AI chat session can proceed accordingly.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/provide-permission/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param callId The unique identifier of the pending tool execution call awaiting a permission decision.
// @return ApiProvidePermissionRequest
func (a *AIChatAPIService) ProvidePermission(ctx context.Context, callId string) ApiProvidePermissionRequest {
	return ApiProvidePermissionRequest{
		ApiService: a,
		ctx: ctx,
		callId: callId,
	}
}

// Execute executes the request
func (a *AIChatAPIService) ProvidePermissionExecute(r ApiProvidePermissionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIChatAPIService.ProvidePermission")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/chats/tool-permissions/{callId}/decision"
	localVarPath = strings.Replace(localVarPath, "{"+"callId"+"}", url.PathEscape(parameterValueToString(r.callId, "callId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.toolDecisionRequestBody == nil {
		return nil, reportError("toolDecisionRequestBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.toolDecisionRequestBody
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

type ApiRenameChatRequest struct {
	ctx context.Context
	ApiService *AIChatAPIService
	chatId string
	renameChatBody *RenameChatBody
}

// The new chat name.
func (r ApiRenameChatRequest) RenameChatBody(renameChatBody RenameChatBody) ApiRenameChatRequest {	r.renameChatBody = &renameChatBody
	return r
}

func (r ApiRenameChatRequest) Execute() (*ChatWrapper, *http.Response, error) {
	return r.ApiService.RenameChatExecute(r)
}

// RenameChat Rename an AI chat
//
// Updates the display title of an existing AI chat session owned by the current user.  The new name must not exceed 255 characters.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/rename-chat/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param chatId The unique identifier of the AI chat session to rename.
// @return ApiRenameChatRequest
func (a *AIChatAPIService) RenameChat(ctx context.Context, chatId string) ApiRenameChatRequest {
	return ApiRenameChatRequest{
		ApiService: a,
		ctx: ctx,
		chatId: chatId,
	}
}

// Execute executes the request
//  @return ChatWrapper
func (a *AIChatAPIService) RenameChatExecute(r ApiRenameChatRequest) (*ChatWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ChatWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIChatAPIService.RenameChat")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/chats/{chatId}"
	localVarPath = strings.Replace(localVarPath, "{"+"chatId"+"}", url.PathEscape(parameterValueToString(r.chatId, "chatId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.renameChatBody == nil {
		return localVarReturnValue, nil, reportError("renameChatBody is required and must be specified")
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
	localVarPostBody = r.renameChatBody
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

type ApiResolveEditorToolRequest struct {
	ctx context.Context
	ApiService *AIChatAPIService
	callId string
	editorToolDecisionRequestBody *EditorToolDecisionRequestBody
}

// The decision parameters.
func (r ApiResolveEditorToolRequest) EditorToolDecisionRequestBody(editorToolDecisionRequestBody EditorToolDecisionRequestBody) ApiResolveEditorToolRequest {	r.editorToolDecisionRequestBody = &editorToolDecisionRequestBody
	return r
}

func (r ApiResolveEditorToolRequest) Execute() (*GeneratedFileWrapper, *http.Response, error) {
	return r.ApiService.ResolveEditorToolExecute(r)
}

// ResolveEditorTool Resolve a pending editor file-generation tool
//
// Submits the user's approval or denial for a pending editor generation tool call (docx, form, presentation).  On approval the file is created from the original tool arguments and information about it is returned,  while the suspended chat tool is resumed with the same result so the AI session can continue.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/resolve-editor-tool/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param callId The unique identifier of the pending tool call awaiting the user's decision.
// @return ApiResolveEditorToolRequest
func (a *AIChatAPIService) ResolveEditorTool(ctx context.Context, callId string) ApiResolveEditorToolRequest {
	return ApiResolveEditorToolRequest{
		ApiService: a,
		ctx: ctx,
		callId: callId,
	}
}

// Execute executes the request
//  @return GeneratedFileWrapper
func (a *AIChatAPIService) ResolveEditorToolExecute(r ApiResolveEditorToolRequest) (*GeneratedFileWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GeneratedFileWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIChatAPIService.ResolveEditorTool")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/chats/tool-files/{callId}/decision"
	localVarPath = strings.Replace(localVarPath, "{"+"callId"+"}", url.PathEscape(parameterValueToString(r.callId, "callId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.editorToolDecisionRequestBody == nil {
		return localVarReturnValue, nil, reportError("editorToolDecisionRequestBody is required and must be specified")
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
	localVarPostBody = r.editorToolDecisionRequestBody
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

type ApiSetUserChatsSettingsRequest struct {
	ctx context.Context
	ApiService *AIChatAPIService
	roomId int32
	setUserChatSettingsRequestBody *SetUserChatSettingsRequestBody
}

// The chat settings to apply.
func (r ApiSetUserChatsSettingsRequest) SetUserChatSettingsRequestBody(setUserChatSettingsRequestBody SetUserChatSettingsRequestBody) ApiSetUserChatsSettingsRequest {	r.setUserChatSettingsRequestBody = &setUserChatSettingsRequestBody
	return r
}

func (r ApiSetUserChatsSettingsRequest) Execute() (*UserChatSettingsWrapper, *http.Response, error) {
	return r.ApiService.SetUserChatsSettingsExecute(r)
}

// SetUserChatsSettings Update user chat settings for a room
//
// Saves the current user's personal AI chat preferences for the specified room.  Currently supports toggling the web search capability, which allows the AI assistant to search the internet when generating responses.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/set-user-chats-settings/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param roomId The identifier of the room whose chat settings are to be updated.
// @return ApiSetUserChatsSettingsRequest
func (a *AIChatAPIService) SetUserChatsSettings(ctx context.Context, roomId int32) ApiSetUserChatsSettingsRequest {
	return ApiSetUserChatsSettingsRequest{
		ApiService: a,
		ctx: ctx,
		roomId: roomId,
	}
}

// Execute executes the request
//  @return UserChatSettingsWrapper
func (a *AIChatAPIService) SetUserChatsSettingsExecute(r ApiSetUserChatsSettingsRequest) (*UserChatSettingsWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserChatSettingsWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIChatAPIService.SetUserChatsSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/rooms/{roomId}/chats/config"
	localVarPath = strings.Replace(localVarPath, "{"+"roomId"+"}", url.PathEscape(parameterValueToString(r.roomId, "roomId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.setUserChatSettingsRequestBody == nil {
		return localVarReturnValue, nil, reportError("setUserChatSettingsRequestBody is required and must be specified")
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
	localVarPostBody = r.setUserChatSettingsRequestBody
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

type ApiStartNewChatRequest struct {
	ctx context.Context
	ApiService *AIChatAPIService
	roomId int32
	startNewChatBody *StartNewChatBody
}

// The initial message and optional file attachments.
func (r ApiStartNewChatRequest) StartNewChatBody(startNewChatBody StartNewChatBody) ApiStartNewChatRequest {	r.startNewChatBody = &startNewChatBody
	return r
}

func (r ApiStartNewChatRequest) Execute() (*http.Response, error) {
	return r.ApiService.StartNewChatExecute(r)
}

// StartNewChat Start a new AI chat
//
// Creates a new AI chat session within the specified room and sends the initial message to the configured AI provider.  The response is delivered as a Server-Sent Events (SSE) stream containing completion chunks (text deltas, tool calls, tool results, and message lifecycle events)  with periodic keep-alive pings every 5 seconds. File references can be included as context for the AI model.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/start-new-chat/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param roomId The identifier of the room in which to create the new AI chat session.
// @return ApiStartNewChatRequest
func (a *AIChatAPIService) StartNewChat(ctx context.Context, roomId int32) ApiStartNewChatRequest {
	return ApiStartNewChatRequest{
		ApiService: a,
		ctx: ctx,
		roomId: roomId,
	}
}

// Execute executes the request
func (a *AIChatAPIService) StartNewChatExecute(r ApiStartNewChatRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIChatAPIService.StartNewChat")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/rooms/{roomId}/chats"
	localVarPath = strings.Replace(localVarPath, "{"+"roomId"+"}", url.PathEscape(parameterValueToString(r.roomId, "roomId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.startNewChatBody == nil {
		return nil, reportError("startNewChatBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.startNewChatBody
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
