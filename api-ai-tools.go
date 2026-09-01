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


// AIToolsAPIService AIToolsAPI service
type AIToolsAPIService service

type ApiAiToolsAddCustomServerRequest struct {
	ctx context.Context
	ApiService *AIToolsAPIService
	aiToolsAddCustomServerRequest *AiToolsAddCustomServerRequest
}

func (r ApiAiToolsAddCustomServerRequest) AiToolsAddCustomServerRequest(aiToolsAddCustomServerRequest AiToolsAddCustomServerRequest) ApiAiToolsAddCustomServerRequest {	r.aiToolsAddCustomServerRequest = &aiToolsAddCustomServerRequest
	return r
}

func (r ApiAiToolsAddCustomServerRequest) Execute() (*AiToolsMutationResult, *http.Response, error) {
	return r.ApiService.AiToolsAddCustomServerExecute(r)
}

// AiToolsAddCustomServer Add custom server
//
// Registers a custom MCP server in the scope under the given name.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-add-custom-server/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiToolsAddCustomServerRequest
func (a *AIToolsAPIService) AiToolsAddCustomServer(ctx context.Context) ApiAiToolsAddCustomServerRequest {
	return ApiAiToolsAddCustomServerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiToolsMutationResult
func (a *AIToolsAPIService) AiToolsAddCustomServerExecute(r ApiAiToolsAddCustomServerRequest) (*AiToolsMutationResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiToolsMutationResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIToolsAPIService.AiToolsAddCustomServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/tools/add-custom-server"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiToolsAddCustomServerRequest == nil {
		return localVarReturnValue, nil, reportError("aiToolsAddCustomServerRequest is required and must be specified")
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
	localVarPostBody = r.aiToolsAddCustomServerRequest
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

type ApiAiToolsGetAllowAlwaysRequest struct {
	ctx context.Context
	ApiService *AIToolsAPIService
	entityId *string
}

// The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope.
func (r ApiAiToolsGetAllowAlwaysRequest) EntityId(entityId string) ApiAiToolsGetAllowAlwaysRequest {	r.entityId = &entityId
	return r
}

func (r ApiAiToolsGetAllowAlwaysRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.AiToolsGetAllowAlwaysExecute(r)
}

// AiToolsGetAllowAlways Get allow always
//
// Lists the tools on the always-allow list of the scope.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-get-allow-always/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiToolsGetAllowAlwaysRequest
func (a *AIToolsAPIService) AiToolsGetAllowAlways(ctx context.Context) ApiAiToolsGetAllowAlwaysRequest {
	return ApiAiToolsGetAllowAlwaysRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []string
func (a *AIToolsAPIService) AiToolsGetAllowAlwaysExecute(r ApiAiToolsGetAllowAlwaysRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIToolsAPIService.AiToolsGetAllowAlways")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/tools/get-allow-always"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.entityId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "entityId", r.entityId, "form", "")
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

type ApiAiToolsGetCustomServerRequest struct {
	ctx context.Context
	ApiService *AIToolsAPIService
	name *string
	entityId *string
}

// The custom MCP server name.
func (r ApiAiToolsGetCustomServerRequest) Name(name string) ApiAiToolsGetCustomServerRequest {	r.name = &name
	return r
}

// The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope.
func (r ApiAiToolsGetCustomServerRequest) EntityId(entityId string) ApiAiToolsGetCustomServerRequest {	r.entityId = &entityId
	return r
}

func (r ApiAiToolsGetCustomServerRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AiToolsGetCustomServerExecute(r)
}

// AiToolsGetCustomServer Get custom server
//
// Returns the configuration of one custom MCP server, or an empty result when it is not registered.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-get-custom-server/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiToolsGetCustomServerRequest
func (a *AIToolsAPIService) AiToolsGetCustomServer(ctx context.Context) ApiAiToolsGetCustomServerRequest {
	return ApiAiToolsGetCustomServerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AIToolsAPIService) AiToolsGetCustomServerExecute(r ApiAiToolsGetCustomServerRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIToolsAPIService.AiToolsGetCustomServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/tools/get-custom-server"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.name == nil {
		return localVarReturnValue, nil, reportError("name is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	if r.entityId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "entityId", r.entityId, "form", "")
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

type ApiAiToolsGetDisabledRequest struct {
	ctx context.Context
	ApiService *AIToolsAPIService
	entityId *string
}

// The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope.
func (r ApiAiToolsGetDisabledRequest) EntityId(entityId string) ApiAiToolsGetDisabledRequest {	r.entityId = &entityId
	return r
}

func (r ApiAiToolsGetDisabledRequest) Execute() (*map[string][]string, *http.Response, error) {
	return r.ApiService.AiToolsGetDisabledExecute(r)
}

// AiToolsGetDisabled Get disabled
//
// Returns the switched-off tools of the scope, grouped by server type.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-get-disabled/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiToolsGetDisabledRequest
func (a *AIToolsAPIService) AiToolsGetDisabled(ctx context.Context) ApiAiToolsGetDisabledRequest {
	return ApiAiToolsGetDisabledRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string][]string
func (a *AIToolsAPIService) AiToolsGetDisabledExecute(r ApiAiToolsGetDisabledRequest) (*map[string][]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *map[string][]string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIToolsAPIService.AiToolsGetDisabled")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/tools/get-disabled"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.entityId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "entityId", r.entityId, "form", "")
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

type ApiAiToolsIsAllowAlwaysRequest struct {
	ctx context.Context
	ApiService *AIToolsAPIService
	serverType *string
	toolName *string
	entityId *string
}

// The MCP server type the tool belongs to.
func (r ApiAiToolsIsAllowAlwaysRequest) ServerType(serverType string) ApiAiToolsIsAllowAlwaysRequest {	r.serverType = &serverType
	return r
}

// The tool name.
func (r ApiAiToolsIsAllowAlwaysRequest) ToolName(toolName string) ApiAiToolsIsAllowAlwaysRequest {	r.toolName = &toolName
	return r
}

// The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope.
func (r ApiAiToolsIsAllowAlwaysRequest) EntityId(entityId string) ApiAiToolsIsAllowAlwaysRequest {	r.entityId = &entityId
	return r
}

func (r ApiAiToolsIsAllowAlwaysRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.AiToolsIsAllowAlwaysExecute(r)
}

// AiToolsIsAllowAlways Is allow always
//
// Tells whether one tool is on the always-allow list.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-is-allow-always/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiToolsIsAllowAlwaysRequest
func (a *AIToolsAPIService) AiToolsIsAllowAlways(ctx context.Context) ApiAiToolsIsAllowAlwaysRequest {
	return ApiAiToolsIsAllowAlwaysRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return bool
func (a *AIToolsAPIService) AiToolsIsAllowAlwaysExecute(r ApiAiToolsIsAllowAlwaysRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIToolsAPIService.AiToolsIsAllowAlways")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/tools/is-allow-always"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serverType == nil {
		return localVarReturnValue, nil, reportError("serverType is required and must be specified")
	}
	if r.toolName == nil {
		return localVarReturnValue, nil, reportError("toolName is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "serverType", r.serverType, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "toolName", r.toolName, "form", "")
	if r.entityId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "entityId", r.entityId, "form", "")
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

type ApiAiToolsIsToolDisabledRequest struct {
	ctx context.Context
	ApiService *AIToolsAPIService
	serverType *string
	toolName *string
	entityId *string
}

// The MCP server type the tool belongs to.
func (r ApiAiToolsIsToolDisabledRequest) ServerType(serverType string) ApiAiToolsIsToolDisabledRequest {	r.serverType = &serverType
	return r
}

// The tool name.
func (r ApiAiToolsIsToolDisabledRequest) ToolName(toolName string) ApiAiToolsIsToolDisabledRequest {	r.toolName = &toolName
	return r
}

// The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope.
func (r ApiAiToolsIsToolDisabledRequest) EntityId(entityId string) ApiAiToolsIsToolDisabledRequest {	r.entityId = &entityId
	return r
}

func (r ApiAiToolsIsToolDisabledRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.AiToolsIsToolDisabledExecute(r)
}

// AiToolsIsToolDisabled Is tool disabled
//
// Tells whether one tool of a server type is switched off.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-is-tool-disabled/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiToolsIsToolDisabledRequest
func (a *AIToolsAPIService) AiToolsIsToolDisabled(ctx context.Context) ApiAiToolsIsToolDisabledRequest {
	return ApiAiToolsIsToolDisabledRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return bool
func (a *AIToolsAPIService) AiToolsIsToolDisabledExecute(r ApiAiToolsIsToolDisabledRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIToolsAPIService.AiToolsIsToolDisabled")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/tools/is-tool-disabled"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serverType == nil {
		return localVarReturnValue, nil, reportError("serverType is required and must be specified")
	}
	if r.toolName == nil {
		return localVarReturnValue, nil, reportError("toolName is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "serverType", r.serverType, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "toolName", r.toolName, "form", "")
	if r.entityId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "entityId", r.entityId, "form", "")
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

type ApiAiToolsListCustomServersRequest struct {
	ctx context.Context
	ApiService *AIToolsAPIService
	entityId *string
}

// The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope.
func (r ApiAiToolsListCustomServersRequest) EntityId(entityId string) ApiAiToolsListCustomServersRequest {	r.entityId = &entityId
	return r
}

func (r ApiAiToolsListCustomServersRequest) Execute() (map[string]map[string]interface{}, *http.Response, error) {
	return r.ApiService.AiToolsListCustomServersExecute(r)
}

// AiToolsListCustomServers List custom servers
//
// Lists the custom MCP servers registered in the scope, keyed by name.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-list-custom-servers/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiToolsListCustomServersRequest
func (a *AIToolsAPIService) AiToolsListCustomServers(ctx context.Context) ApiAiToolsListCustomServersRequest {
	return ApiAiToolsListCustomServersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]map[string]interface{}
func (a *AIToolsAPIService) AiToolsListCustomServersExecute(r ApiAiToolsListCustomServersRequest) (map[string]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIToolsAPIService.AiToolsListCustomServers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/tools/list-custom-servers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.entityId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "entityId", r.entityId, "form", "")
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

type ApiAiToolsListSystemToolsRequest struct {
	ctx context.Context
	ApiService *AIToolsAPIService
	entityId *string
}

// The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope.
func (r ApiAiToolsListSystemToolsRequest) EntityId(entityId string) ApiAiToolsListSystemToolsRequest {	r.entityId = &entityId
	return r
}

func (r ApiAiToolsListSystemToolsRequest) Execute() (*map[string][]AiTMCPItem, *http.Response, error) {
	return r.ApiService.AiToolsListSystemToolsExecute(r)
}

// AiToolsListSystemTools List system tools
//
// Lists the tools of the host-configured system MCP servers, grouped by server type. The servers are connected and listed server-side, so the client renders its permission cards from one request and never opens an MCP connection of its own.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-list-system-tools/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiToolsListSystemToolsRequest
func (a *AIToolsAPIService) AiToolsListSystemTools(ctx context.Context) ApiAiToolsListSystemToolsRequest {
	return ApiAiToolsListSystemToolsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string][]AiTMCPItem
func (a *AIToolsAPIService) AiToolsListSystemToolsExecute(r ApiAiToolsListSystemToolsRequest) (*map[string][]AiTMCPItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *map[string][]AiTMCPItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIToolsAPIService.AiToolsListSystemTools")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/tools/list-system-tools"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.entityId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "entityId", r.entityId, "form", "")
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

type ApiAiToolsRemoveCustomServerRequest struct {
	ctx context.Context
	ApiService *AIToolsAPIService
	aiToolsRemoveCustomServerRequest *AiToolsRemoveCustomServerRequest
}

func (r ApiAiToolsRemoveCustomServerRequest) AiToolsRemoveCustomServerRequest(aiToolsRemoveCustomServerRequest AiToolsRemoveCustomServerRequest) ApiAiToolsRemoveCustomServerRequest {	r.aiToolsRemoveCustomServerRequest = &aiToolsRemoveCustomServerRequest
	return r
}

func (r ApiAiToolsRemoveCustomServerRequest) Execute() (*AiSuccessResponse, *http.Response, error) {
	return r.ApiService.AiToolsRemoveCustomServerExecute(r)
}

// AiToolsRemoveCustomServer Remove custom server
//
// Removes a custom MCP server from the registry.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-remove-custom-server/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiToolsRemoveCustomServerRequest
func (a *AIToolsAPIService) AiToolsRemoveCustomServer(ctx context.Context) ApiAiToolsRemoveCustomServerRequest {
	return ApiAiToolsRemoveCustomServerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiSuccessResponse
func (a *AIToolsAPIService) AiToolsRemoveCustomServerExecute(r ApiAiToolsRemoveCustomServerRequest) (*AiSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIToolsAPIService.AiToolsRemoveCustomServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/tools/remove-custom-server"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiToolsRemoveCustomServerRequest == nil {
		return localVarReturnValue, nil, reportError("aiToolsRemoveCustomServerRequest is required and must be specified")
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
	localVarPostBody = r.aiToolsRemoveCustomServerRequest
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

type ApiAiToolsReplaceAllCustomServersRequest struct {
	ctx context.Context
	ApiService *AIToolsAPIService
	aiToolsReplaceAllCustomServersRequest *AiToolsReplaceAllCustomServersRequest
}

func (r ApiAiToolsReplaceAllCustomServersRequest) AiToolsReplaceAllCustomServersRequest(aiToolsReplaceAllCustomServersRequest AiToolsReplaceAllCustomServersRequest) ApiAiToolsReplaceAllCustomServersRequest {	r.aiToolsReplaceAllCustomServersRequest = &aiToolsReplaceAllCustomServersRequest
	return r
}

func (r ApiAiToolsReplaceAllCustomServersRequest) Execute() (*AiToolsBulkResult, *http.Response, error) {
	return r.ApiService.AiToolsReplaceAllCustomServersExecute(r)
}

// AiToolsReplaceAllCustomServers Replace all custom servers
//
// Replaces the whole custom MCP server registry of the scope with the supplied map.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-replace-all-custom-servers/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiToolsReplaceAllCustomServersRequest
func (a *AIToolsAPIService) AiToolsReplaceAllCustomServers(ctx context.Context) ApiAiToolsReplaceAllCustomServersRequest {
	return ApiAiToolsReplaceAllCustomServersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiToolsBulkResult
func (a *AIToolsAPIService) AiToolsReplaceAllCustomServersExecute(r ApiAiToolsReplaceAllCustomServersRequest) (*AiToolsBulkResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiToolsBulkResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIToolsAPIService.AiToolsReplaceAllCustomServers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/tools/replace-all-custom-servers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiToolsReplaceAllCustomServersRequest == nil {
		return localVarReturnValue, nil, reportError("aiToolsReplaceAllCustomServersRequest is required and must be specified")
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
	localVarPostBody = r.aiToolsReplaceAllCustomServersRequest
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

type ApiAiToolsSetAllowAlwaysRequest struct {
	ctx context.Context
	ApiService *AIToolsAPIService
	aiToolsSetAllowAlwaysRequest *AiToolsSetAllowAlwaysRequest
}

func (r ApiAiToolsSetAllowAlwaysRequest) AiToolsSetAllowAlwaysRequest(aiToolsSetAllowAlwaysRequest AiToolsSetAllowAlwaysRequest) ApiAiToolsSetAllowAlwaysRequest {	r.aiToolsSetAllowAlwaysRequest = &aiToolsSetAllowAlwaysRequest
	return r
}

func (r ApiAiToolsSetAllowAlwaysRequest) Execute() (*AiSuccessResponse, *http.Response, error) {
	return r.ApiService.AiToolsSetAllowAlwaysExecute(r)
}

// AiToolsSetAllowAlways Set allow always
//
// Adds a tool to the always-allow list, or removes it - the tools on that list run without an approval dialog.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-set-allow-always/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiToolsSetAllowAlwaysRequest
func (a *AIToolsAPIService) AiToolsSetAllowAlways(ctx context.Context) ApiAiToolsSetAllowAlwaysRequest {
	return ApiAiToolsSetAllowAlwaysRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiSuccessResponse
func (a *AIToolsAPIService) AiToolsSetAllowAlwaysExecute(r ApiAiToolsSetAllowAlwaysRequest) (*AiSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIToolsAPIService.AiToolsSetAllowAlways")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/tools/set-allow-always"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiToolsSetAllowAlwaysRequest == nil {
		return localVarReturnValue, nil, reportError("aiToolsSetAllowAlwaysRequest is required and must be specified")
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
	localVarPostBody = r.aiToolsSetAllowAlwaysRequest
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

type ApiAiToolsSetDisabledRequest struct {
	ctx context.Context
	ApiService *AIToolsAPIService
	aiToolsSetDisabledRequest *AiToolsSetDisabledRequest
}

func (r ApiAiToolsSetDisabledRequest) AiToolsSetDisabledRequest(aiToolsSetDisabledRequest AiToolsSetDisabledRequest) ApiAiToolsSetDisabledRequest {	r.aiToolsSetDisabledRequest = &aiToolsSetDisabledRequest
	return r
}

func (r ApiAiToolsSetDisabledRequest) Execute() (*AiSuccessResponse, *http.Response, error) {
	return r.ApiService.AiToolsSetDisabledExecute(r)
}

// AiToolsSetDisabled Set disabled
//
// Marks the listed tools of one server type as switched off, so the model is no longer offered them.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-set-disabled/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiToolsSetDisabledRequest
func (a *AIToolsAPIService) AiToolsSetDisabled(ctx context.Context) ApiAiToolsSetDisabledRequest {
	return ApiAiToolsSetDisabledRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiSuccessResponse
func (a *AIToolsAPIService) AiToolsSetDisabledExecute(r ApiAiToolsSetDisabledRequest) (*AiSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIToolsAPIService.AiToolsSetDisabled")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/tools/set-disabled"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiToolsSetDisabledRequest == nil {
		return localVarReturnValue, nil, reportError("aiToolsSetDisabledRequest is required and must be specified")
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
	localVarPostBody = r.aiToolsSetDisabledRequest
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

type ApiAiToolsUpdateCustomServerRequest struct {
	ctx context.Context
	ApiService *AIToolsAPIService
	aiToolsUpdateCustomServerRequest *AiToolsUpdateCustomServerRequest
}

func (r ApiAiToolsUpdateCustomServerRequest) AiToolsUpdateCustomServerRequest(aiToolsUpdateCustomServerRequest AiToolsUpdateCustomServerRequest) ApiAiToolsUpdateCustomServerRequest {	r.aiToolsUpdateCustomServerRequest = &aiToolsUpdateCustomServerRequest
	return r
}

func (r ApiAiToolsUpdateCustomServerRequest) Execute() (*AiToolsMutationResult, *http.Response, error) {
	return r.ApiService.AiToolsUpdateCustomServerExecute(r)
}

// AiToolsUpdateCustomServer Update custom server
//
// Updates the configuration of a registered custom MCP server.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-update-custom-server/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiToolsUpdateCustomServerRequest
func (a *AIToolsAPIService) AiToolsUpdateCustomServer(ctx context.Context) ApiAiToolsUpdateCustomServerRequest {
	return ApiAiToolsUpdateCustomServerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiToolsMutationResult
func (a *AIToolsAPIService) AiToolsUpdateCustomServerExecute(r ApiAiToolsUpdateCustomServerRequest) (*AiToolsMutationResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiToolsMutationResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIToolsAPIService.AiToolsUpdateCustomServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/tools/update-custom-server"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiToolsUpdateCustomServerRequest == nil {
		return localVarReturnValue, nil, reportError("aiToolsUpdateCustomServerRequest is required and must be specified")
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
	localVarPostBody = r.aiToolsUpdateCustomServerRequest
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
