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


// AIContextAPIService AIContextAPI service
type AIContextAPIService service

type ApiAiContextGetContextFoldersRequest struct {
	ctx context.Context
	ApiService *AIContextAPIService
}

func (r ApiAiContextGetContextFoldersRequest) Execute() (*AiSuccessResponse, *http.Response, error) {
	return r.ApiService.AiContextGetContextFoldersExecute(r)
}

// AiContextGetContextFolders Get context folders
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-context-get-context-folders/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiContextGetContextFoldersRequest
func (a *AIContextAPIService) AiContextGetContextFolders(ctx context.Context) ApiAiContextGetContextFoldersRequest {
	return ApiAiContextGetContextFoldersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiSuccessResponse
func (a *AIContextAPIService) AiContextGetContextFoldersExecute(r ApiAiContextGetContextFoldersRequest) (*AiSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIContextAPIService.AiContextGetContextFolders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/context/get-context-folders"

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
		if localVarHTTPResponse.StatusCode == 401 {
			var v AiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v AiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
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

type ApiAiContextGetRoomSkillRequest struct {
	ctx context.Context
	ApiService *AIContextAPIService
	cloud *string
	roomId *string
	skillId *string
}

func (r ApiAiContextGetRoomSkillRequest) Cloud(cloud string) ApiAiContextGetRoomSkillRequest {	r.cloud = &cloud
	return r
}

func (r ApiAiContextGetRoomSkillRequest) RoomId(roomId string) ApiAiContextGetRoomSkillRequest {	r.roomId = &roomId
	return r
}

func (r ApiAiContextGetRoomSkillRequest) SkillId(skillId string) ApiAiContextGetRoomSkillRequest {	r.skillId = &skillId
	return r
}

func (r ApiAiContextGetRoomSkillRequest) Execute() (*AiSuccessResponse, *http.Response, error) {
	return r.ApiService.AiContextGetRoomSkillExecute(r)
}

// AiContextGetRoomSkill Get room skill
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-context-get-room-skill/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiContextGetRoomSkillRequest
func (a *AIContextAPIService) AiContextGetRoomSkill(ctx context.Context) ApiAiContextGetRoomSkillRequest {
	return ApiAiContextGetRoomSkillRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiSuccessResponse
func (a *AIContextAPIService) AiContextGetRoomSkillExecute(r ApiAiContextGetRoomSkillRequest) (*AiSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIContextAPIService.AiContextGetRoomSkill")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/context/get-room-skill"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cloud == nil {
		return localVarReturnValue, nil, reportError("cloud is required and must be specified")
	}
	if r.roomId == nil {
		return localVarReturnValue, nil, reportError("roomId is required and must be specified")
	}
	if r.skillId == nil {
		return localVarReturnValue, nil, reportError("skillId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "cloud", r.cloud, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "roomId", r.roomId, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "skillId", r.skillId, "form", "")
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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v AiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
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

type ApiAiContextGetRoomSkillsRequest struct {
	ctx context.Context
	ApiService *AIContextAPIService
	cloud *string
	roomId *string
}

func (r ApiAiContextGetRoomSkillsRequest) Cloud(cloud string) ApiAiContextGetRoomSkillsRequest {	r.cloud = &cloud
	return r
}

func (r ApiAiContextGetRoomSkillsRequest) RoomId(roomId string) ApiAiContextGetRoomSkillsRequest {	r.roomId = &roomId
	return r
}

func (r ApiAiContextGetRoomSkillsRequest) Execute() (*AiSuccessResponse, *http.Response, error) {
	return r.ApiService.AiContextGetRoomSkillsExecute(r)
}

// AiContextGetRoomSkills Get room skills
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-context-get-room-skills/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiContextGetRoomSkillsRequest
func (a *AIContextAPIService) AiContextGetRoomSkills(ctx context.Context) ApiAiContextGetRoomSkillsRequest {
	return ApiAiContextGetRoomSkillsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiSuccessResponse
func (a *AIContextAPIService) AiContextGetRoomSkillsExecute(r ApiAiContextGetRoomSkillsRequest) (*AiSuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiSuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIContextAPIService.AiContextGetRoomSkills")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/context/get-room-skills"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cloud == nil {
		return localVarReturnValue, nil, reportError("cloud is required and must be specified")
	}
	if r.roomId == nil {
		return localVarReturnValue, nil, reportError("roomId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "cloud", r.cloud, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "roomId", r.roomId, "form", "")
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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v AiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
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
