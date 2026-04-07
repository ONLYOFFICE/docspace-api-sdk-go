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


// PortalUsersAPIService PortalUsersAPI service
type PortalUsersAPIService service

type ApiCreateInvitationLinkRequest struct {
	ctx context.Context
	ApiService *PortalUsersAPIService
	invitationLinkCreateRequestDto *InvitationLinkCreateRequestDto
}

func (r ApiCreateInvitationLinkRequest) InvitationLinkCreateRequestDto(invitationLinkCreateRequestDto InvitationLinkCreateRequestDto) ApiCreateInvitationLinkRequest {	r.invitationLinkCreateRequestDto = &invitationLinkCreateRequestDto
	return r
}

func (r ApiCreateInvitationLinkRequest) Execute() (*InvitationLinkWrapper, *http.Response, error) {
	return r.ApiService.CreateInvitationLinkExecute(r)
}

// CreateInvitationLink Create an invitation link
//
// Returns an invitation link for joining the portal.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/create-invitation-link/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiCreateInvitationLinkRequest
func (a *PortalUsersAPIService) CreateInvitationLink(ctx context.Context) ApiCreateInvitationLinkRequest {
	return ApiCreateInvitationLinkRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InvitationLinkWrapper
func (a *PortalUsersAPIService) CreateInvitationLinkExecute(r ApiCreateInvitationLinkRequest) (*InvitationLinkWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvitationLinkWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalUsersAPIService.CreateInvitationLink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/users/invitationlink"

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
	localVarPostBody = r.invitationLinkCreateRequestDto
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

type ApiDeleteInvitationLinkRequest struct {
	ctx context.Context
	ApiService *PortalUsersAPIService
	invitationLinkDeleteRequestDto *InvitationLinkDeleteRequestDto
}

// The data transfer object containing the details of the invitation link to be deleted.
func (r ApiDeleteInvitationLinkRequest) InvitationLinkDeleteRequestDto(invitationLinkDeleteRequestDto InvitationLinkDeleteRequestDto) ApiDeleteInvitationLinkRequest {	r.invitationLinkDeleteRequestDto = &invitationLinkDeleteRequestDto
	return r
}

func (r ApiDeleteInvitationLinkRequest) Execute() (*StringWrapper, *http.Response, error) {
	return r.ApiService.DeleteInvitationLinkExecute(r)
}

// DeleteInvitationLink Deletes an invitation link.
//
// Ensures that the current user has permission to delete the specified invitation link.  Throws security or not-found exceptions if required conditions are not met.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-invitation-link/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiDeleteInvitationLinkRequest
func (a *PortalUsersAPIService) DeleteInvitationLink(ctx context.Context) ApiDeleteInvitationLinkRequest {
	return ApiDeleteInvitationLinkRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StringWrapper
func (a *PortalUsersAPIService) DeleteInvitationLinkExecute(r ApiDeleteInvitationLinkRequest) (*StringWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StringWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalUsersAPIService.DeleteInvitationLink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/users/invitationlink"

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
	localVarPostBody = r.invitationLinkDeleteRequestDto
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

type ApiGetInvitationLinkRequest struct {
	ctx context.Context
	ApiService *PortalUsersAPIService
	employeeType EmployeeType
}

func (r ApiGetInvitationLinkRequest) Execute() (*StringWrapper, *http.Response, error) {
	return r.ApiService.GetInvitationLinkExecute(r)
}

// GetInvitationLink Get an invitation link
//
// Returns an invitation link for joining the portal.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-invitation-link/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param employeeType The type of employee role for the invitation link (DocSpaceAdmin, RoomAdmin or User).
// @return ApiGetInvitationLinkRequest
//
// Deprecated
func (a *PortalUsersAPIService) GetInvitationLink(ctx context.Context, employeeType EmployeeType) ApiGetInvitationLinkRequest {
	return ApiGetInvitationLinkRequest{
		ApiService: a,
		ctx: ctx,
		employeeType: employeeType,
	}
}

// Execute executes the request
//  @return StringWrapper
// Deprecated
func (a *PortalUsersAPIService) GetInvitationLinkExecute(r ApiGetInvitationLinkRequest) (*StringWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StringWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalUsersAPIService.GetInvitationLink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/users/invite/{employeeType}"
	localVarPath = strings.Replace(localVarPath, "{"+"employeeType"+"}", url.PathEscape(parameterValueToString(r.employeeType, "employeeType")), -1)

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

type ApiGetInvitationLinkByEmployeeTypeRequest struct {
	ctx context.Context
	ApiService *PortalUsersAPIService
	employeeType EmployeeType
}

func (r ApiGetInvitationLinkByEmployeeTypeRequest) Execute() (*InvitationLinkWrapper, *http.Response, error) {
	return r.ApiService.GetInvitationLinkByEmployeeTypeExecute(r)
}

// GetInvitationLinkByEmployeeType Get an invitation link
//
// Returns an invitation link for joining the portal.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-invitation-link-by-employee-type/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param employeeType The type of employee role for the invitation link (DocSpaceAdmin, RoomAdmin or User).
// @return ApiGetInvitationLinkByEmployeeTypeRequest
func (a *PortalUsersAPIService) GetInvitationLinkByEmployeeType(ctx context.Context, employeeType EmployeeType) ApiGetInvitationLinkByEmployeeTypeRequest {
	return ApiGetInvitationLinkByEmployeeTypeRequest{
		ApiService: a,
		ctx: ctx,
		employeeType: employeeType,
	}
}

// Execute executes the request
//  @return InvitationLinkWrapper
func (a *PortalUsersAPIService) GetInvitationLinkByEmployeeTypeExecute(r ApiGetInvitationLinkByEmployeeTypeRequest) (*InvitationLinkWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvitationLinkWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalUsersAPIService.GetInvitationLinkByEmployeeType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/users/invitationlink/{employeeType}"
	localVarPath = strings.Replace(localVarPath, "{"+"employeeType"+"}", url.PathEscape(parameterValueToString(r.employeeType, "employeeType")), -1)

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

type ApiGetPortalUsersCountRequest struct {
	ctx context.Context
	ApiService *PortalUsersAPIService
}

func (r ApiGetPortalUsersCountRequest) Execute() (*Int64Wrapper, *http.Response, error) {
	return r.ApiService.GetPortalUsersCountExecute(r)
}

// GetPortalUsersCount Get a number of portal users
//
// Returns a number of portal users.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-portal-users-count/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetPortalUsersCountRequest
func (a *PortalUsersAPIService) GetPortalUsersCount(ctx context.Context) ApiGetPortalUsersCountRequest {
	return ApiGetPortalUsersCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Int64Wrapper
func (a *PortalUsersAPIService) GetPortalUsersCountExecute(r ApiGetPortalUsersCountRequest) (*Int64Wrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Int64Wrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalUsersAPIService.GetPortalUsersCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/userscount"

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

type ApiGetUserByIdRequest struct {
	ctx context.Context
	ApiService *PortalUsersAPIService
	userID string
}

func (r ApiGetUserByIdRequest) Execute() (*UserInfoWrapper, *http.Response, error) {
	return r.ApiService.GetUserByIdExecute(r)
}

// GetUserById Get a user by ID
//
// Returns a user with the ID specified in the request from the current portal.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-user-by-id/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param userID The user ID extracted from the route parameters.
// @return ApiGetUserByIdRequest
func (a *PortalUsersAPIService) GetUserById(ctx context.Context, userID string) ApiGetUserByIdRequest {
	return ApiGetUserByIdRequest{
		ApiService: a,
		ctx: ctx,
		userID: userID,
	}
}

// Execute executes the request
//  @return UserInfoWrapper
func (a *PortalUsersAPIService) GetUserByIdExecute(r ApiGetUserByIdRequest) (*UserInfoWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserInfoWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalUsersAPIService.GetUserById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/users/{userID}"
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", url.PathEscape(parameterValueToString(r.userID, "userID")), -1)

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

type ApiMarkGiftMessageAsReadRequest struct {
	ctx context.Context
	ApiService *PortalUsersAPIService
}

func (r ApiMarkGiftMessageAsReadRequest) Execute() (*http.Response, error) {
	return r.ApiService.MarkGiftMessageAsReadExecute(r)
}

// MarkGiftMessageAsRead Mark a gift message as read
//
// Marks a gift message as read.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/mark-gift-message-as-read/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiMarkGiftMessageAsReadRequest
func (a *PortalUsersAPIService) MarkGiftMessageAsRead(ctx context.Context) ApiMarkGiftMessageAsReadRequest {
	return ApiMarkGiftMessageAsReadRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PortalUsersAPIService) MarkGiftMessageAsReadExecute(r ApiMarkGiftMessageAsReadRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalUsersAPIService.MarkGiftMessageAsRead")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/present/mark"

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

type ApiSendCongratulationsRequest struct {
	ctx context.Context
	ApiService *PortalUsersAPIService
	userid *string
	key *string
}

// The user ID to receive the congratulatory message.
func (r ApiSendCongratulationsRequest) Userid(userid string) ApiSendCongratulationsRequest {	r.userid = &userid
	return r
}

// The template identifier or email configuration key.
func (r ApiSendCongratulationsRequest) Key(key string) ApiSendCongratulationsRequest {	r.key = &key
	return r
}

func (r ApiSendCongratulationsRequest) Execute() (*http.Response, error) {
	return r.ApiService.SendCongratulationsExecute(r)
}

// SendCongratulations Send congratulations
//
// Sends congratulations to the user after registering a portal.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/send-congratulations/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiSendCongratulationsRequest
func (a *PortalUsersAPIService) SendCongratulations(ctx context.Context) ApiSendCongratulationsRequest {
	return ApiSendCongratulationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PortalUsersAPIService) SendCongratulationsExecute(r ApiSendCongratulationsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalUsersAPIService.SendCongratulations")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/sendcongratulations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userid == nil {
		return nil, reportError("userid is required and must be specified")
	}
	if r.key == nil {
		return nil, reportError("key is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "Userid", r.userid, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "Key", r.key, "form", "")
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

type ApiUpdateInvitationLinkRequest struct {
	ctx context.Context
	ApiService *PortalUsersAPIService
	invitationLinkUpdateRequestDto *InvitationLinkUpdateRequestDto
}

func (r ApiUpdateInvitationLinkRequest) InvitationLinkUpdateRequestDto(invitationLinkUpdateRequestDto InvitationLinkUpdateRequestDto) ApiUpdateInvitationLinkRequest {	r.invitationLinkUpdateRequestDto = &invitationLinkUpdateRequestDto
	return r
}

func (r ApiUpdateInvitationLinkRequest) Execute() (*InvitationLinkWrapper, *http.Response, error) {
	return r.ApiService.UpdateInvitationLinkExecute(r)
}

// UpdateInvitationLink Update an invitation link
//
// Returns an invitation link for joining the portal.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/update-invitation-link/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiUpdateInvitationLinkRequest
func (a *PortalUsersAPIService) UpdateInvitationLink(ctx context.Context) ApiUpdateInvitationLinkRequest {
	return ApiUpdateInvitationLinkRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InvitationLinkWrapper
func (a *PortalUsersAPIService) UpdateInvitationLinkExecute(r ApiUpdateInvitationLinkRequest) (*InvitationLinkWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvitationLinkWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalUsersAPIService.UpdateInvitationLink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/users/invitationlink"

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
	localVarPostBody = r.invitationLinkUpdateRequestDto
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
