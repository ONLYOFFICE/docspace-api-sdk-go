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


// AIAssignmentsAPIService AIAssignmentsAPI service
type AIAssignmentsAPIService service

type ApiAiAssignmentsAssignRequest struct {
	ctx context.Context
	ApiService *AIAssignmentsAPIService
	aiAssignmentsAssignRequest *AiAssignmentsAssignRequest
}

func (r ApiAiAssignmentsAssignRequest) AiAssignmentsAssignRequest(aiAssignmentsAssignRequest AiAssignmentsAssignRequest) ApiAiAssignmentsAssignRequest {	r.aiAssignmentsAssignRequest = &aiAssignmentsAssignRequest
	return r
}

func (r ApiAiAssignmentsAssignRequest) Execute() (*AiAssignmentMutationResult, *http.Response, error) {
	return r.ApiService.AiAssignmentsAssignExecute(r)
}

// AiAssignmentsAssign Assign
//
// Binds a profile to an AI action, creating the assignment or updating it in place. The profile's declared capabilities are validated against the action, except for the `Default` slot.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-assignments-assign/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAssignmentsAssignRequest
func (a *AIAssignmentsAPIService) AiAssignmentsAssign(ctx context.Context) ApiAiAssignmentsAssignRequest {
	return ApiAiAssignmentsAssignRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiAssignmentMutationResult
func (a *AIAssignmentsAPIService) AiAssignmentsAssignExecute(r ApiAiAssignmentsAssignRequest) (*AiAssignmentMutationResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiAssignmentMutationResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAssignmentsAPIService.AiAssignmentsAssign")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/assignments/assign"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiAssignmentsAssignRequest == nil {
		return localVarReturnValue, nil, reportError("aiAssignmentsAssignRequest is required and must be specified")
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
	localVarPostBody = r.aiAssignmentsAssignRequest
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

type ApiAiAssignmentsBulkAssignRequest struct {
	ctx context.Context
	ApiService *AIAssignmentsAPIService
	requestBody *map[string]string
}

func (r ApiAiAssignmentsBulkAssignRequest) RequestBody(requestBody map[string]string) ApiAiAssignmentsBulkAssignRequest {	r.requestBody = &requestBody
	return r
}

func (r ApiAiAssignmentsBulkAssignRequest) Execute() (*AiBulkAssignmentResult, *http.Response, error) {
	return r.ApiService.AiAssignmentsBulkAssignExecute(r)
}

// AiAssignmentsBulkAssign Bulk assign
//
// Applies many action-to-profile bindings at once. Every entry is validated first and nothing is written if any of them fails, so the assignment set is never left half-written.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-assignments-bulk-assign/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAssignmentsBulkAssignRequest
func (a *AIAssignmentsAPIService) AiAssignmentsBulkAssign(ctx context.Context) ApiAiAssignmentsBulkAssignRequest {
	return ApiAiAssignmentsBulkAssignRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiBulkAssignmentResult
func (a *AIAssignmentsAPIService) AiAssignmentsBulkAssignExecute(r ApiAiAssignmentsBulkAssignRequest) (*AiBulkAssignmentResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiBulkAssignmentResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAssignmentsAPIService.AiAssignmentsBulkAssign")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/assignments/bulk-assign"

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

type ApiAiAssignmentsCascadeProfileDeleteRequest struct {
	ctx context.Context
	ApiService *AIAssignmentsAPIService
	body *string
}

func (r ApiAiAssignmentsCascadeProfileDeleteRequest) Body(body string) ApiAiAssignmentsCascadeProfileDeleteRequest {	r.body = &body
	return r
}

func (r ApiAiAssignmentsCascadeProfileDeleteRequest) Execute() (*AiSuccessResponse, *http.Response, error) {
	return r.ApiService.AiAssignmentsCascadeProfileDeleteExecute(r)
}

// AiAssignmentsCascadeProfileDelete Cascade profile delete
//
// Cleans up the assignments pointing at a profile that is about to be deleted: the `Default` slot is promoted to the first remaining profile (or dropped when none is left), and every other slot holding that profile is unbound.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-assignments-cascade-profile-delete/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAssignmentsCascadeProfileDeleteRequest
func (a *AIAssignmentsAPIService) AiAssignmentsCascadeProfileDelete(ctx context.Context) ApiAiAssignmentsCascadeProfileDeleteRequest {
	return ApiAiAssignmentsCascadeProfileDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiSuccessResponse
func (a *AIAssignmentsAPIService) AiAssignmentsCascadeProfileDeleteExecute(r ApiAiAssignmentsCascadeProfileDeleteRequest) (*AiSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAssignmentsAPIService.AiAssignmentsCascadeProfileDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/assignments/cascade-profile-delete"

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

type ApiAiAssignmentsGetAllAssignmentsRequest struct {
	ctx context.Context
	ApiService *AIAssignmentsAPIService
	entityId *string
}

// The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope.
func (r ApiAiAssignmentsGetAllAssignmentsRequest) EntityId(entityId string) ApiAiAssignmentsGetAllAssignmentsRequest {	r.entityId = &entityId
	return r
}

func (r ApiAiAssignmentsGetAllAssignmentsRequest) Execute() (map[string]string, *http.Response, error) {
	return r.ApiService.AiAssignmentsGetAllAssignmentsExecute(r)
}

// AiAssignmentsGetAllAssignments Get all assignments
//
// Returns the full action-to-profile assignment map of the scope.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-assignments-get-all-assignments/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAssignmentsGetAllAssignmentsRequest
func (a *AIAssignmentsAPIService) AiAssignmentsGetAllAssignments(ctx context.Context) ApiAiAssignmentsGetAllAssignmentsRequest {
	return ApiAiAssignmentsGetAllAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]string
func (a *AIAssignmentsAPIService) AiAssignmentsGetAllAssignmentsExecute(r ApiAiAssignmentsGetAllAssignmentsRequest) (map[string]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAssignmentsAPIService.AiAssignmentsGetAllAssignments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/assignments/get-all-assignments"

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

type ApiAiAssignmentsGetAssignmentRequest struct {
	ctx context.Context
	ApiService *AIAssignmentsAPIService
	actionType *string
}

// The AI action the request applies to - one of Default, Chat, Code, Summarization, Translation, TextAnalyze, ImageGeneration, OCR, Vision.
func (r ApiAiAssignmentsGetAssignmentRequest) ActionType(actionType string) ApiAiAssignmentsGetAssignmentRequest {	r.actionType = &actionType
	return r
}

func (r ApiAiAssignmentsGetAssignmentRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.AiAssignmentsGetAssignmentExecute(r)
}

// AiAssignmentsGetAssignment Get assignment
//
// Returns the profile bound to one AI action, without the `Default` fallback.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-assignments-get-assignment/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAssignmentsGetAssignmentRequest
func (a *AIAssignmentsAPIService) AiAssignmentsGetAssignment(ctx context.Context) ApiAiAssignmentsGetAssignmentRequest {
	return ApiAiAssignmentsGetAssignmentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *AIAssignmentsAPIService) AiAssignmentsGetAssignmentExecute(r ApiAiAssignmentsGetAssignmentRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAssignmentsAPIService.AiAssignmentsGetAssignment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/assignments/get-assignment"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.actionType == nil {
		return localVarReturnValue, nil, reportError("actionType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "actionType", r.actionType, "form", "")
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

type ApiAiAssignmentsResolveForActionRequest struct {
	ctx context.Context
	ApiService *AIAssignmentsAPIService
	actionType *string
	entityId *string
}

// The AI action the request applies to - one of Default, Chat, Code, Summarization, Translation, TextAnalyze, ImageGeneration, OCR, Vision.
func (r ApiAiAssignmentsResolveForActionRequest) ActionType(actionType string) ApiAiAssignmentsResolveForActionRequest {	r.actionType = &actionType
	return r
}

// The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope.
func (r ApiAiAssignmentsResolveForActionRequest) EntityId(entityId string) ApiAiAssignmentsResolveForActionRequest {	r.entityId = &entityId
	return r
}

func (r ApiAiAssignmentsResolveForActionRequest) Execute() (*AiResolvedAssignment, *http.Response, error) {
	return r.ApiService.AiAssignmentsResolveForActionExecute(r)
}

// AiAssignmentsResolveForAction Resolve for action
//
// Resolves the profile bound to an AI action, falling back to the `Default` slot when the action itself has none. Fails when neither slot is set or the bound profile no longer exists - use `try-resolve-for-action` for an empty answer instead.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-assignments-resolve-for-action/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAssignmentsResolveForActionRequest
func (a *AIAssignmentsAPIService) AiAssignmentsResolveForAction(ctx context.Context) ApiAiAssignmentsResolveForActionRequest {
	return ApiAiAssignmentsResolveForActionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiResolvedAssignment
func (a *AIAssignmentsAPIService) AiAssignmentsResolveForActionExecute(r ApiAiAssignmentsResolveForActionRequest) (*AiResolvedAssignment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiResolvedAssignment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAssignmentsAPIService.AiAssignmentsResolveForAction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/assignments/resolve-for-action"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.actionType == nil {
		return localVarReturnValue, nil, reportError("actionType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "actionType", r.actionType, "form", "")
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

type ApiAiAssignmentsTryResolveForActionRequest struct {
	ctx context.Context
	ApiService *AIAssignmentsAPIService
	actionType *string
	entityId *string
}

// The AI action the request applies to - one of Default, Chat, Code, Summarization, Translation, TextAnalyze, ImageGeneration, OCR, Vision.
func (r ApiAiAssignmentsTryResolveForActionRequest) ActionType(actionType string) ApiAiAssignmentsTryResolveForActionRequest {	r.actionType = &actionType
	return r
}

// The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope.
func (r ApiAiAssignmentsTryResolveForActionRequest) EntityId(entityId string) ApiAiAssignmentsTryResolveForActionRequest {	r.entityId = &entityId
	return r
}

func (r ApiAiAssignmentsTryResolveForActionRequest) Execute() (*AiResolvedAssignment, *http.Response, error) {
	return r.ApiService.AiAssignmentsTryResolveForActionExecute(r)
}

// AiAssignmentsTryResolveForAction Try resolve for action
//
// Resolves the profile bound to an AI action exactly like `resolve-for-action`, but answers with an empty result instead of failing when nothing is configured.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-assignments-try-resolve-for-action/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAssignmentsTryResolveForActionRequest
func (a *AIAssignmentsAPIService) AiAssignmentsTryResolveForAction(ctx context.Context) ApiAiAssignmentsTryResolveForActionRequest {
	return ApiAiAssignmentsTryResolveForActionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiResolvedAssignment
func (a *AIAssignmentsAPIService) AiAssignmentsTryResolveForActionExecute(r ApiAiAssignmentsTryResolveForActionRequest) (*AiResolvedAssignment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiResolvedAssignment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAssignmentsAPIService.AiAssignmentsTryResolveForAction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/assignments/try-resolve-for-action"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.actionType == nil {
		return localVarReturnValue, nil, reportError("actionType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "actionType", r.actionType, "form", "")
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

type ApiAiAssignmentsUnassignRequest struct {
	ctx context.Context
	ApiService *AIAssignmentsAPIService
	body *string
}

func (r ApiAiAssignmentsUnassignRequest) Body(body string) ApiAiAssignmentsUnassignRequest {	r.body = &body
	return r
}

func (r ApiAiAssignmentsUnassignRequest) Execute() (*AiSuccessResponse, *http.Response, error) {
	return r.ApiService.AiAssignmentsUnassignExecute(r)
}

// AiAssignmentsUnassign Unassign
//
// Removes the profile binding of an AI action. Does nothing when that slot is already empty.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-assignments-unassign/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiAssignmentsUnassignRequest
func (a *AIAssignmentsAPIService) AiAssignmentsUnassign(ctx context.Context) ApiAiAssignmentsUnassignRequest {
	return ApiAiAssignmentsUnassignRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiSuccessResponse
func (a *AIAssignmentsAPIService) AiAssignmentsUnassignExecute(r ApiAiAssignmentsUnassignRequest) (*AiSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIAssignmentsAPIService.AiAssignmentsUnassign")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/assignments/unassign"

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
