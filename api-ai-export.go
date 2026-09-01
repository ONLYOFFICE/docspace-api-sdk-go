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


// AIExportAPIService AIExportAPI service
type AIExportAPIService service

type ApiAiExportTextToDocxRequest struct {
	ctx context.Context
	ApiService *AIExportAPIService
	aiExportTextToDocxRequest *AiExportTextToDocxRequest
}

func (r ApiAiExportTextToDocxRequest) AiExportTextToDocxRequest(aiExportTextToDocxRequest AiExportTextToDocxRequest) ApiAiExportTextToDocxRequest {	r.aiExportTextToDocxRequest = &aiExportTextToDocxRequest
	return r
}

func (r ApiAiExportTextToDocxRequest) Execute() (*AiExportTextToDocx200Response, *http.Response, error) {
	return r.ApiService.AiExportTextToDocxExecute(r)
}

// AiExportTextToDocx Start markdown → docx export
//
// Starts an asynchronous markdown-to-docx export. The response only acknowledges the task: the AI Worker converts the content and saves the .docx into the target folder (an agent room resolves to its result-storage subfolder), and completion reaches the client as the usual folder-modified socket event.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-export-text-to-docx/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAiExportTextToDocxRequest
func (a *AIExportAPIService) AiExportTextToDocx(ctx context.Context) ApiAiExportTextToDocxRequest {
	return ApiAiExportTextToDocxRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiExportTextToDocx200Response
func (a *AIExportAPIService) AiExportTextToDocxExecute(r ApiAiExportTextToDocxRequest) (*AiExportTextToDocx200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiExportTextToDocx200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIExportAPIService.AiExportTextToDocx")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/text-to-docx"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.aiExportTextToDocxRequest == nil {
		return localVarReturnValue, nil, reportError("aiExportTextToDocxRequest is required and must be specified")
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
	localVarPostBody = r.aiExportTextToDocxRequest
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
