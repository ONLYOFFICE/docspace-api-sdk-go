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
	"time"
)


// SettingsWebhooksAPIService SettingsWebhooksAPI service
type SettingsWebhooksAPIService service

type ApiCreateWebhookRequest struct {
	ctx context.Context
	ApiService *SettingsWebhooksAPIService
	createWebhooksConfigRequestsDto *CreateWebhooksConfigRequestsDto
}

func (r ApiCreateWebhookRequest) CreateWebhooksConfigRequestsDto(createWebhooksConfigRequestsDto CreateWebhooksConfigRequestsDto) ApiCreateWebhookRequest {	r.createWebhooksConfigRequestsDto = &createWebhooksConfigRequestsDto
	return r
}

func (r ApiCreateWebhookRequest) Execute() (*WebhooksConfigWrapper, *http.Response, error) {
	return r.ApiService.CreateWebhookExecute(r)
}

// CreateWebhook Create a webhook
//
// Creates a new tenant webhook with the parameters specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/create-webhook/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiCreateWebhookRequest
func (a *SettingsWebhooksAPIService) CreateWebhook(ctx context.Context) ApiCreateWebhookRequest {
	return ApiCreateWebhookRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WebhooksConfigWrapper
func (a *SettingsWebhooksAPIService) CreateWebhookExecute(r ApiCreateWebhookRequest) (*WebhooksConfigWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhooksConfigWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsWebhooksAPIService.CreateWebhook")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/webhook"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.createWebhooksConfigRequestsDto
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

type ApiEnableWebhookRequest struct {
	ctx context.Context
	ApiService *SettingsWebhooksAPIService
	updateWebhooksConfigRequestsDto *UpdateWebhooksConfigRequestsDto
}

func (r ApiEnableWebhookRequest) UpdateWebhooksConfigRequestsDto(updateWebhooksConfigRequestsDto UpdateWebhooksConfigRequestsDto) ApiEnableWebhookRequest {	r.updateWebhooksConfigRequestsDto = &updateWebhooksConfigRequestsDto
	return r
}

func (r ApiEnableWebhookRequest) Execute() (*WebhooksConfigWrapper, *http.Response, error) {
	return r.ApiService.EnableWebhookExecute(r)
}

// EnableWebhook Enable a webhook
//
// Enables or disables a tenant webhook with the parameters specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/enable-webhook/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiEnableWebhookRequest
func (a *SettingsWebhooksAPIService) EnableWebhook(ctx context.Context) ApiEnableWebhookRequest {
	return ApiEnableWebhookRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WebhooksConfigWrapper
func (a *SettingsWebhooksAPIService) EnableWebhookExecute(r ApiEnableWebhookRequest) (*WebhooksConfigWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhooksConfigWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsWebhooksAPIService.EnableWebhook")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/webhook/enable"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.updateWebhooksConfigRequestsDto
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

type ApiGetTenantWebhooksRequest struct {
	ctx context.Context
	ApiService *SettingsWebhooksAPIService
}

func (r ApiGetTenantWebhooksRequest) Execute() (*WebhooksConfigWithStatusArrayWrapper, *http.Response, error) {
	return r.ApiService.GetTenantWebhooksExecute(r)
}

// GetTenantWebhooks Get webhooks
//
// Returns a list of the tenant webhooks.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-tenant-webhooks/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetTenantWebhooksRequest
func (a *SettingsWebhooksAPIService) GetTenantWebhooks(ctx context.Context) ApiGetTenantWebhooksRequest {
	return ApiGetTenantWebhooksRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WebhooksConfigWithStatusArrayWrapper
func (a *SettingsWebhooksAPIService) GetTenantWebhooksExecute(r ApiGetTenantWebhooksRequest) (*WebhooksConfigWithStatusArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhooksConfigWithStatusArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsWebhooksAPIService.GetTenantWebhooks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/webhook"

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

type ApiGetWebhookTriggersRequest struct {
	ctx context.Context
	ApiService *SettingsWebhooksAPIService
}

func (r ApiGetWebhookTriggersRequest) Execute() (*GetWebhookTriggers200Response, *http.Response, error) {
	return r.ApiService.GetWebhookTriggersExecute(r)
}

// GetWebhookTriggers Get webhook triggers
//
// Returns a list of triggers for a webhook.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-webhook-triggers/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetWebhookTriggersRequest
func (a *SettingsWebhooksAPIService) GetWebhookTriggers(ctx context.Context) ApiGetWebhookTriggersRequest {
	return ApiGetWebhookTriggersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetWebhookTriggers200Response
func (a *SettingsWebhooksAPIService) GetWebhookTriggersExecute(r ApiGetWebhookTriggersRequest) (*GetWebhookTriggers200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetWebhookTriggers200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsWebhooksAPIService.GetWebhookTriggers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/webhook/triggers"

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

type ApiGetWebhooksLogsRequest struct {
	ctx context.Context
	ApiService *SettingsWebhooksAPIService
	deliveryFrom *time.Time
	deliveryTo *time.Time
	hookUri *string
	configId *int32
	eventId *int32
	groupStatus *WebhookGroupStatus
	userId *string
	trigger *WebhookTrigger
	count *int32
	startIndex *int32
}

// The delivery start time for filtering webhook logs.
func (r ApiGetWebhooksLogsRequest) DeliveryFrom(deliveryFrom time.Time) ApiGetWebhooksLogsRequest {	r.deliveryFrom = &deliveryFrom
	return r
}

// The delivery end time for filtering webhook logs.
func (r ApiGetWebhooksLogsRequest) DeliveryTo(deliveryTo time.Time) ApiGetWebhooksLogsRequest {	r.deliveryTo = &deliveryTo
	return r
}

// The destination URL where webhooks are delivered.
func (r ApiGetWebhooksLogsRequest) HookUri(hookUri string) ApiGetWebhooksLogsRequest {	r.hookUri = &hookUri
	return r
}

// The webhook configuration identifier.
func (r ApiGetWebhooksLogsRequest) ConfigId(configId int32) ApiGetWebhooksLogsRequest {	r.configId = &configId
	return r
}

// The unique identifier of the event that triggered the webhook.
func (r ApiGetWebhooksLogsRequest) EventId(eventId int32) ApiGetWebhooksLogsRequest {	r.eventId = &eventId
	return r
}

// The status of the webhook delivery group.
func (r ApiGetWebhooksLogsRequest) GroupStatus(groupStatus WebhookGroupStatus) ApiGetWebhooksLogsRequest {	r.groupStatus = &groupStatus
	return r
}

// The identifier of the user associated with the webhook event.
func (r ApiGetWebhooksLogsRequest) UserId(userId string) ApiGetWebhooksLogsRequest {	r.userId = &userId
	return r
}

// The type of event that triggered the webhook.
func (r ApiGetWebhooksLogsRequest) Trigger(trigger WebhookTrigger) ApiGetWebhooksLogsRequest {	r.trigger = &trigger
	return r
}

// The maximum number of webhook log records to return in the query response.
func (r ApiGetWebhooksLogsRequest) Count(count int32) ApiGetWebhooksLogsRequest {	r.count = &count
	return r
}

// Specifies the starting index for retrieving webhook logs.  Used for pagination in the webhook delivery log queries.
func (r ApiGetWebhooksLogsRequest) StartIndex(startIndex int32) ApiGetWebhooksLogsRequest {	r.startIndex = &startIndex
	return r
}

func (r ApiGetWebhooksLogsRequest) Execute() (*WebhooksLogArrayWrapper, *http.Response, error) {
	return r.ApiService.GetWebhooksLogsExecute(r)
}

// GetWebhooksLogs Get webhook logs
//
// Returns the logs of the webhook activities.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-webhooks-logs/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetWebhooksLogsRequest
func (a *SettingsWebhooksAPIService) GetWebhooksLogs(ctx context.Context) ApiGetWebhooksLogsRequest {
	return ApiGetWebhooksLogsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WebhooksLogArrayWrapper
func (a *SettingsWebhooksAPIService) GetWebhooksLogsExecute(r ApiGetWebhooksLogsRequest) (*WebhooksLogArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhooksLogArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsWebhooksAPIService.GetWebhooksLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/webhooks/log"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.deliveryFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deliveryFrom", r.deliveryFrom, "form", "")
			}
	if r.deliveryTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "deliveryTo", r.deliveryTo, "form", "")
			}
	if r.hookUri != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "hookUri", r.hookUri, "form", "")
			}
	if r.configId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "configId", r.configId, "form", "")
			}
	if r.eventId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eventId", r.eventId, "form", "")
			}
	if r.groupStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "groupStatus", r.groupStatus, "form", "")
			}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
			}
	if r.trigger != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "trigger", r.trigger, "form", "")
			}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
			}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
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

type ApiRemoveWebhookRequest struct {
	ctx context.Context
	ApiService *SettingsWebhooksAPIService
	id int32
}

func (r ApiRemoveWebhookRequest) Execute() (*WebhooksConfigWrapper, *http.Response, error) {
	return r.ApiService.RemoveWebhookExecute(r)
}

// RemoveWebhook Remove a webhook
//
// Removes a tenant webhook with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/remove-webhook/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The ID extracted from the route parameters.
// @return ApiRemoveWebhookRequest
func (a *SettingsWebhooksAPIService) RemoveWebhook(ctx context.Context, id int32) ApiRemoveWebhookRequest {
	return ApiRemoveWebhookRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return WebhooksConfigWrapper
func (a *SettingsWebhooksAPIService) RemoveWebhookExecute(r ApiRemoveWebhookRequest) (*WebhooksConfigWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhooksConfigWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsWebhooksAPIService.RemoveWebhook")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/webhook/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ApiRetryWebhookRequest struct {
	ctx context.Context
	ApiService *SettingsWebhooksAPIService
	id int32
}

func (r ApiRetryWebhookRequest) Execute() (*WebhooksLogWrapper, *http.Response, error) {
	return r.ApiService.RetryWebhookExecute(r)
}

// RetryWebhook Retry a webhook
//
// Retries a webhook with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/retry-webhook/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param id The ID extracted from the route parameters.
// @return ApiRetryWebhookRequest
func (a *SettingsWebhooksAPIService) RetryWebhook(ctx context.Context, id int32) ApiRetryWebhookRequest {
	return ApiRetryWebhookRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return WebhooksLogWrapper
func (a *SettingsWebhooksAPIService) RetryWebhookExecute(r ApiRetryWebhookRequest) (*WebhooksLogWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhooksLogWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsWebhooksAPIService.RetryWebhook")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/webhook/{id}/retry"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ApiRetryWebhooksRequest struct {
	ctx context.Context
	ApiService *SettingsWebhooksAPIService
	webhookRetryRequestsDto *WebhookRetryRequestsDto
}

func (r ApiRetryWebhooksRequest) WebhookRetryRequestsDto(webhookRetryRequestsDto WebhookRetryRequestsDto) ApiRetryWebhooksRequest {	r.webhookRetryRequestsDto = &webhookRetryRequestsDto
	return r
}

func (r ApiRetryWebhooksRequest) Execute() (*WebhooksLogArrayWrapper, *http.Response, error) {
	return r.ApiService.RetryWebhooksExecute(r)
}

// RetryWebhooks Retry webhooks
//
// Retries all the webhooks with the IDs specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/retry-webhooks/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiRetryWebhooksRequest
func (a *SettingsWebhooksAPIService) RetryWebhooks(ctx context.Context) ApiRetryWebhooksRequest {
	return ApiRetryWebhooksRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WebhooksLogArrayWrapper
func (a *SettingsWebhooksAPIService) RetryWebhooksExecute(r ApiRetryWebhooksRequest) (*WebhooksLogArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhooksLogArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsWebhooksAPIService.RetryWebhooks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/webhook/retry"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.webhookRetryRequestsDto
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

type ApiUpdateWebhookRequest struct {
	ctx context.Context
	ApiService *SettingsWebhooksAPIService
	updateWebhooksConfigRequestsDto *UpdateWebhooksConfigRequestsDto
}

func (r ApiUpdateWebhookRequest) UpdateWebhooksConfigRequestsDto(updateWebhooksConfigRequestsDto UpdateWebhooksConfigRequestsDto) ApiUpdateWebhookRequest {	r.updateWebhooksConfigRequestsDto = &updateWebhooksConfigRequestsDto
	return r
}

func (r ApiUpdateWebhookRequest) Execute() (*WebhooksConfigWrapper, *http.Response, error) {
	return r.ApiService.UpdateWebhookExecute(r)
}

// UpdateWebhook Update a webhook
//
// Updates a tenant webhook with the parameters specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/update-webhook/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiUpdateWebhookRequest
func (a *SettingsWebhooksAPIService) UpdateWebhook(ctx context.Context) ApiUpdateWebhookRequest {
	return ApiUpdateWebhookRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WebhooksConfigWrapper
func (a *SettingsWebhooksAPIService) UpdateWebhookExecute(r ApiUpdateWebhookRequest) (*WebhooksConfigWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhooksConfigWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingsWebhooksAPIService.UpdateWebhook")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/settings/webhook"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.updateWebhooksConfigRequestsDto
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
