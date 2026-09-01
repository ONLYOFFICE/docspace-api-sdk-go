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


// AIOpenAIPassthroughAPIService AIOpenAIPassthroughAPI service
type AIOpenAIPassthroughAPIService service

type ApiAiOpenaiChatCompletionsRequest struct {
	ctx context.Context
	ApiService *AIOpenAIPassthroughAPIService
	profileId string
	requestBody *map[string]interface{}
}

func (r ApiAiOpenaiChatCompletionsRequest) RequestBody(requestBody map[string]interface{}) ApiAiOpenaiChatCompletionsRequest {	r.requestBody = &requestBody
	return r
}

func (r ApiAiOpenaiChatCompletionsRequest) Execute() (*AiSuccessResponse, *http.Response, error) {
	return r.ApiService.AiOpenaiChatCompletionsExecute(r)
}

// AiOpenaiChatCompletions OpenAI-compatible chat completions proxied to the profile's provider
//
// OpenAI-compatible chat completions for the document editor's AI plugin. The profile is resolved server-side, its credentials are attached, and the body is forwarded to the provider verbatim - the payload is owned by the plugin's SDK on one end and the provider on the other. A client disconnect cancels the provider call.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-openai-chat-completions/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param profileId The AI provider profile identifier.
// @return ApiAiOpenaiChatCompletionsRequest
func (a *AIOpenAIPassthroughAPIService) AiOpenaiChatCompletions(ctx context.Context, profileId string) ApiAiOpenaiChatCompletionsRequest {
	return ApiAiOpenaiChatCompletionsRequest{
		ApiService: a,
		ctx: ctx,
		profileId: profileId,
	}
}

// Execute executes the request
//  @return AiSuccessResponse
func (a *AIOpenAIPassthroughAPIService) AiOpenaiChatCompletionsExecute(r ApiAiOpenaiChatCompletionsRequest) (*AiSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIOpenAIPassthroughAPIService.AiOpenaiChatCompletions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/openai/{profileId}/v1/chat/completions"
	localVarPath = strings.Replace(localVarPath, "{"+"profileId"+"}", url.PathEscape(parameterValueToString(r.profileId, "profileId")), -1)

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

type ApiAiOpenaiImagesGenerationsRequest struct {
	ctx context.Context
	ApiService *AIOpenAIPassthroughAPIService
	profileId string
	requestBody *map[string]interface{}
}

func (r ApiAiOpenaiImagesGenerationsRequest) RequestBody(requestBody map[string]interface{}) ApiAiOpenaiImagesGenerationsRequest {	r.requestBody = &requestBody
	return r
}

func (r ApiAiOpenaiImagesGenerationsRequest) Execute() (*AiSuccessResponse, *http.Response, error) {
	return r.ApiService.AiOpenaiImagesGenerationsExecute(r)
}

// AiOpenaiImagesGenerations OpenAI-compatible image generation proxied to the profile's provider
//
// OpenAI-compatible image generation for the document editor's AI plugin. As with the chat-completions passthrough, the profile's credentials are attached server-side and the body reaches the provider unchanged.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-openai-images-generations/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param profileId The AI provider profile identifier.
// @return ApiAiOpenaiImagesGenerationsRequest
func (a *AIOpenAIPassthroughAPIService) AiOpenaiImagesGenerations(ctx context.Context, profileId string) ApiAiOpenaiImagesGenerationsRequest {
	return ApiAiOpenaiImagesGenerationsRequest{
		ApiService: a,
		ctx: ctx,
		profileId: profileId,
	}
}

// Execute executes the request
//  @return AiSuccessResponse
func (a *AIOpenAIPassthroughAPIService) AiOpenaiImagesGenerationsExecute(r ApiAiOpenaiImagesGenerationsRequest) (*AiSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIOpenAIPassthroughAPIService.AiOpenaiImagesGenerations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/openai/{profileId}/v1/images/generations"
	localVarPath = strings.Replace(localVarPath, "{"+"profileId"+"}", url.PathEscape(parameterValueToString(r.profileId, "profileId")), -1)

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
