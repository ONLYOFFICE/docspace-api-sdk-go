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


// AIMessagesAPIService AIMessagesAPI service
type AIMessagesAPIService service

type ApiExportMessageRequest struct {
	ctx context.Context
	ApiService *AIMessagesAPIService
	messageId int32
	exportMessageRequestBodyInteger *ExportMessageRequestBodyInteger
}

// The export parameters including destination folder and file title.
func (r ApiExportMessageRequest) ExportMessageRequestBodyInteger(exportMessageRequestBodyInteger ExportMessageRequestBodyInteger) ApiExportMessageRequest {	r.exportMessageRequestBodyInteger = &exportMessageRequestBodyInteger
	return r
}

func (r ApiExportMessageRequest) Execute() (*http.Response, error) {
	return r.ApiService.ExportMessageExecute(r)
}

// ExportMessage Export a single AI message to a document
//
// Exports a specific AI chat message as a document into the specified folder. The system verifies that the message exists  and belongs to a chat accessible by the current user, then publishes an asynchronous export task to the event bus.  The exported document will be created in the target folder with the given title once the background task completes.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/export-message/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param messageId The unique identifier of the AI chat message to export.
// @return ApiExportMessageRequest
func (a *AIMessagesAPIService) ExportMessage(ctx context.Context, messageId int32) ApiExportMessageRequest {
	return ApiExportMessageRequest{
		ApiService: a,
		ctx: ctx,
		messageId: messageId,
	}
}

// Execute executes the request
func (a *AIMessagesAPIService) ExportMessageExecute(r ApiExportMessageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AIMessagesAPIService.ExportMessage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/ai/messages/{messageId}/export"
	localVarPath = strings.Replace(localVarPath, "{"+"messageId"+"}", url.PathEscape(parameterValueToString(r.messageId, "messageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.messageId < 1 {
		return nil, reportError("messageId must be greater than 1")
	}
	if r.messageId > 2147483647 {
		return nil, reportError("messageId must be less than 2147483647")
	}
	if r.exportMessageRequestBodyInteger == nil {
		return nil, reportError("exportMessageRequestBodyInteger is required and must be specified")
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
	localVarPostBody = r.exportMessageRequestBodyInteger
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
