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


// PeopleProfilesAPIService PeopleProfilesAPI service
type PeopleProfilesAPIService service

type ApiAddMemberRequest struct {
	ctx context.Context
	ApiService *PeopleProfilesAPIService
	memberRequestDto *MemberRequestDto
}

func (r ApiAddMemberRequest) MemberRequestDto(memberRequestDto MemberRequestDto) ApiAddMemberRequest {	r.memberRequestDto = &memberRequestDto
	return r
}

func (r ApiAddMemberRequest) Execute() (*EmployeeFullWrapper, *http.Response, error) {
	return r.ApiService.AddMemberExecute(r)
}

// AddMember Add a user
//
// Adds a new portal user with the first name, last name, email address, and several optional parameters specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/add-member/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiAddMemberRequest
func (a *PeopleProfilesAPIService) AddMember(ctx context.Context) ApiAddMemberRequest {
	return ApiAddMemberRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EmployeeFullWrapper
func (a *PeopleProfilesAPIService) AddMemberExecute(r ApiAddMemberRequest) (*EmployeeFullWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeFullWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleProfilesAPIService.AddMember")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people"

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
	localVarPostBody = r.memberRequestDto
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

type ApiCheckUserExistsByEmailRequest struct {
	ctx context.Context
	ApiService *PeopleProfilesAPIService
	email *string
	encemail *string
	culture *string
}

// The user email address.
func (r ApiCheckUserExistsByEmailRequest) Email(email string) ApiCheckUserExistsByEmailRequest {	r.email = &email
	return r
}

// The user encrypted email address.
func (r ApiCheckUserExistsByEmailRequest) Encemail(encemail string) ApiCheckUserExistsByEmailRequest {	r.encemail = &encemail
	return r
}

// Culture
func (r ApiCheckUserExistsByEmailRequest) Culture(culture string) ApiCheckUserExistsByEmailRequest {	r.culture = &culture
	return r
}

func (r ApiCheckUserExistsByEmailRequest) Execute() (*UserExistsResponseWrapper, *http.Response, error) {
	return r.ApiService.CheckUserExistsByEmailExecute(r)
}

// CheckUserExistsByEmail Check if a user exists by email
//
// Returns data indicating whether a user with the specified email exists on the portal.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/check-user-exists-by-email/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiCheckUserExistsByEmailRequest
func (a *PeopleProfilesAPIService) CheckUserExistsByEmail(ctx context.Context) ApiCheckUserExistsByEmailRequest {
	return ApiCheckUserExistsByEmailRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UserExistsResponseWrapper
func (a *PeopleProfilesAPIService) CheckUserExistsByEmailExecute(r ApiCheckUserExistsByEmailRequest) (*UserExistsResponseWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserExistsResponseWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleProfilesAPIService.CheckUserExistsByEmail")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/exists"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.email != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
			}
	if r.encemail != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "encemail", r.encemail, "form", "")
			}
	if r.culture != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "culture", r.culture, "form", "")
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

type ApiDeleteMemberRequest struct {
	ctx context.Context
	ApiService *PeopleProfilesAPIService
	userid string
}

func (r ApiDeleteMemberRequest) Execute() (*EmployeeFullWrapper, *http.Response, error) {
	return r.ApiService.DeleteMemberExecute(r)
}

// DeleteMember Delete a user
//
// Deletes a user with the ID specified in the request from the portal.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-member/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param userid The user ID.
// @return ApiDeleteMemberRequest
func (a *PeopleProfilesAPIService) DeleteMember(ctx context.Context, userid string) ApiDeleteMemberRequest {
	return ApiDeleteMemberRequest{
		ApiService: a,
		ctx: ctx,
		userid: userid,
	}
}

// Execute executes the request
//  @return EmployeeFullWrapper
func (a *PeopleProfilesAPIService) DeleteMemberExecute(r ApiDeleteMemberRequest) (*EmployeeFullWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeFullWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleProfilesAPIService.DeleteMember")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/{userid}"
	localVarPath = strings.Replace(localVarPath, "{"+"userid"+"}", url.PathEscape(parameterValueToString(r.userid, "userid")), -1)

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

type ApiDeleteProfileRequest struct {
	ctx context.Context
	ApiService *PeopleProfilesAPIService
}

func (r ApiDeleteProfileRequest) Execute() (*EmployeeFullWrapper, *http.Response, error) {
	return r.ApiService.DeleteProfileExecute(r)
}

// DeleteProfile Delete my profile
//
// Deletes the current user profile.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-profile/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiDeleteProfileRequest
func (a *PeopleProfilesAPIService) DeleteProfile(ctx context.Context) ApiDeleteProfileRequest {
	return ApiDeleteProfileRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EmployeeFullWrapper
func (a *PeopleProfilesAPIService) DeleteProfileExecute(r ApiDeleteProfileRequest) (*EmployeeFullWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeFullWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleProfilesAPIService.DeleteProfile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/@self"

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

type ApiGetAllProfilesRequest struct {
	ctx context.Context
	ApiService *PeopleProfilesAPIService
	count *int32
	startIndex *int32
	filterBy *string
	sortBy *string
	sortOrder *SortOrder
	filterSeparator *string
	filterValue *string
}

// The maximum number of items to be retrieved in the response.
func (r ApiGetAllProfilesRequest) Count(count int32) ApiGetAllProfilesRequest {	r.count = &count
	return r
}

// The zero-based index of the first item to be retrieved in a filtered result set.
func (r ApiGetAllProfilesRequest) StartIndex(startIndex int32) ApiGetAllProfilesRequest {	r.startIndex = &startIndex
	return r
}

// Specifies the filter criteria for user-related queries.
func (r ApiGetAllProfilesRequest) FilterBy(filterBy string) ApiGetAllProfilesRequest {	r.filterBy = &filterBy
	return r
}

// Specifies the property or field name by which the results should be sorted.
func (r ApiGetAllProfilesRequest) SortBy(sortBy string) ApiGetAllProfilesRequest {	r.sortBy = &sortBy
	return r
}

// The order in which the results are sorted.
func (r ApiGetAllProfilesRequest) SortOrder(sortOrder SortOrder) ApiGetAllProfilesRequest {	r.sortOrder = &sortOrder
	return r
}

// The character or string used to separate multiple filter values in a filtering query.
func (r ApiGetAllProfilesRequest) FilterSeparator(filterSeparator string) ApiGetAllProfilesRequest {	r.filterSeparator = &filterSeparator
	return r
}

// The text value used as an additional filter criterion for profiles retrieval.
func (r ApiGetAllProfilesRequest) FilterValue(filterValue string) ApiGetAllProfilesRequest {	r.filterValue = &filterValue
	return r
}

func (r ApiGetAllProfilesRequest) Execute() (*EmployeeFullArrayWrapper, *http.Response, error) {
	return r.ApiService.GetAllProfilesExecute(r)
}

// GetAllProfiles Get profiles
//
// Returns a list of profiles for all the portal users.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-all-profiles/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetAllProfilesRequest
func (a *PeopleProfilesAPIService) GetAllProfiles(ctx context.Context) ApiGetAllProfilesRequest {
	return ApiGetAllProfilesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EmployeeFullArrayWrapper
func (a *PeopleProfilesAPIService) GetAllProfilesExecute(r ApiGetAllProfilesRequest) (*EmployeeFullArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeFullArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleProfilesAPIService.GetAllProfiles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
			}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
			}
	if r.filterBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterBy", r.filterBy, "form", "")
			}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortBy", r.sortBy, "form", "")
			}
	if r.sortOrder != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortOrder", r.sortOrder, "form", "")
			}
	if r.filterSeparator != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterSeparator", r.filterSeparator, "form", "")
			}
	if r.filterValue != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filterValue", r.filterValue, "form", "")
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

type ApiGetClaimsRequest struct {
	ctx context.Context
	ApiService *PeopleProfilesAPIService
}

func (r ApiGetClaimsRequest) Execute() (*ObjectWrapper, *http.Response, error) {
	return r.ApiService.GetClaimsExecute(r)
}

// GetClaims Get user claims
//
// Returns the user claims.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-claims/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetClaimsRequest
func (a *PeopleProfilesAPIService) GetClaims(ctx context.Context) ApiGetClaimsRequest {
	return ApiGetClaimsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ObjectWrapper
func (a *PeopleProfilesAPIService) GetClaimsExecute(r ApiGetClaimsRequest) (*ObjectWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ObjectWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleProfilesAPIService.GetClaims")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/tokendiagnostics"

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

type ApiGetProfileByEmailRequest struct {
	ctx context.Context
	ApiService *PeopleProfilesAPIService
	email *string
	encemail *string
	culture *string
}

// The user email address.
func (r ApiGetProfileByEmailRequest) Email(email string) ApiGetProfileByEmailRequest {	r.email = &email
	return r
}

// The user encrypted email address.
func (r ApiGetProfileByEmailRequest) Encemail(encemail string) ApiGetProfileByEmailRequest {	r.encemail = &encemail
	return r
}

// Culture
func (r ApiGetProfileByEmailRequest) Culture(culture string) ApiGetProfileByEmailRequest {	r.culture = &culture
	return r
}

func (r ApiGetProfileByEmailRequest) Execute() (*EmployeeFullWrapper, *http.Response, error) {
	return r.ApiService.GetProfileByEmailExecute(r)
}

// GetProfileByEmail Get a profile by user email
//
// Returns the detailed information about a profile of the user with the email specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-profile-by-email/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetProfileByEmailRequest
func (a *PeopleProfilesAPIService) GetProfileByEmail(ctx context.Context) ApiGetProfileByEmailRequest {
	return ApiGetProfileByEmailRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EmployeeFullWrapper
func (a *PeopleProfilesAPIService) GetProfileByEmailExecute(r ApiGetProfileByEmailRequest) (*EmployeeFullWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeFullWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleProfilesAPIService.GetProfileByEmail")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/email"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.email != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
			}
	if r.encemail != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "encemail", r.encemail, "form", "")
			}
	if r.culture != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "culture", r.culture, "form", "")
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

type ApiGetProfileByUserIdRequest struct {
	ctx context.Context
	ApiService *PeopleProfilesAPIService
	userid string
}

func (r ApiGetProfileByUserIdRequest) Execute() (*EmployeeFullWrapper, *http.Response, error) {
	return r.ApiService.GetProfileByUserIdExecute(r)
}

// GetProfileByUserId Get a profile by user ID
//
// Returns the detailed information about a profile of the user with the ID specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-profile-by-user-id/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param userid The user ID.
// @return ApiGetProfileByUserIdRequest
func (a *PeopleProfilesAPIService) GetProfileByUserId(ctx context.Context, userid string) ApiGetProfileByUserIdRequest {
	return ApiGetProfileByUserIdRequest{
		ApiService: a,
		ctx: ctx,
		userid: userid,
	}
}

// Execute executes the request
//  @return EmployeeFullWrapper
func (a *PeopleProfilesAPIService) GetProfileByUserIdExecute(r ApiGetProfileByUserIdRequest) (*EmployeeFullWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeFullWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleProfilesAPIService.GetProfileByUserId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/{userid}"
	localVarPath = strings.Replace(localVarPath, "{"+"userid"+"}", url.PathEscape(parameterValueToString(r.userid, "userid")), -1)

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

type ApiGetSelfProfileRequest struct {
	ctx context.Context
	ApiService *PeopleProfilesAPIService
}

func (r ApiGetSelfProfileRequest) Execute() (*EmployeeFullWrapper, *http.Response, error) {
	return r.ApiService.GetSelfProfileExecute(r)
}

// GetSelfProfile Get my profile
//
// Returns the detailed information about the current user profile.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-self-profile/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetSelfProfileRequest
func (a *PeopleProfilesAPIService) GetSelfProfile(ctx context.Context) ApiGetSelfProfileRequest {
	return ApiGetSelfProfileRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EmployeeFullWrapper
func (a *PeopleProfilesAPIService) GetSelfProfileExecute(r ApiGetSelfProfileRequest) (*EmployeeFullWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeFullWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleProfilesAPIService.GetSelfProfile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/@self"

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

type ApiInviteUsersRequest struct {
	ctx context.Context
	ApiService *PeopleProfilesAPIService
	inviteUsersRequestDto *InviteUsersRequestDto
}

func (r ApiInviteUsersRequest) InviteUsersRequestDto(inviteUsersRequestDto InviteUsersRequestDto) ApiInviteUsersRequest {	r.inviteUsersRequestDto = &inviteUsersRequestDto
	return r
}

func (r ApiInviteUsersRequest) Execute() (*EmployeeArrayWrapper, *http.Response, error) {
	return r.ApiService.InviteUsersExecute(r)
}

// InviteUsers Invite users
//
// Invites users specified in the request to the current portal.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/invite-users/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiInviteUsersRequest
func (a *PeopleProfilesAPIService) InviteUsers(ctx context.Context) ApiInviteUsersRequest {
	return ApiInviteUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EmployeeArrayWrapper
func (a *PeopleProfilesAPIService) InviteUsersExecute(r ApiInviteUsersRequest) (*EmployeeArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleProfilesAPIService.InviteUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/invite"

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
	localVarPostBody = r.inviteUsersRequestDto
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

type ApiRemoveUsersRequest struct {
	ctx context.Context
	ApiService *PeopleProfilesAPIService
	updateMembersRequestDto *UpdateMembersRequestDto
}

func (r ApiRemoveUsersRequest) UpdateMembersRequestDto(updateMembersRequestDto UpdateMembersRequestDto) ApiRemoveUsersRequest {	r.updateMembersRequestDto = &updateMembersRequestDto
	return r
}

func (r ApiRemoveUsersRequest) Execute() (*EmployeeFullArrayWrapper, *http.Response, error) {
	return r.ApiService.RemoveUsersExecute(r)
}

// RemoveUsers Delete users
//
// Deletes a list of the users with the IDs specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/remove-users/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiRemoveUsersRequest
func (a *PeopleProfilesAPIService) RemoveUsers(ctx context.Context) ApiRemoveUsersRequest {
	return ApiRemoveUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EmployeeFullArrayWrapper
func (a *PeopleProfilesAPIService) RemoveUsersExecute(r ApiRemoveUsersRequest) (*EmployeeFullArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeFullArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleProfilesAPIService.RemoveUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/delete"

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
	localVarPostBody = r.updateMembersRequestDto
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

type ApiResendUserInvitesRequest struct {
	ctx context.Context
	ApiService *PeopleProfilesAPIService
	updateMembersRequestDto *UpdateMembersRequestDto
}

func (r ApiResendUserInvitesRequest) UpdateMembersRequestDto(updateMembersRequestDto UpdateMembersRequestDto) ApiResendUserInvitesRequest {	r.updateMembersRequestDto = &updateMembersRequestDto
	return r
}

func (r ApiResendUserInvitesRequest) Execute() (*EmployeeFullArrayWrapper, *http.Response, error) {
	return r.ApiService.ResendUserInvitesExecute(r)
}

// ResendUserInvites Resend activation emails
//
// Resends emails to the users who have not activated their emails.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/resend-user-invites/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiResendUserInvitesRequest
func (a *PeopleProfilesAPIService) ResendUserInvites(ctx context.Context) ApiResendUserInvitesRequest {
	return ApiResendUserInvitesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EmployeeFullArrayWrapper
func (a *PeopleProfilesAPIService) ResendUserInvitesExecute(r ApiResendUserInvitesRequest) (*EmployeeFullArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeFullArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleProfilesAPIService.ResendUserInvites")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/invite"

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
	localVarPostBody = r.updateMembersRequestDto
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

type ApiUpdateMemberRequest struct {
	ctx context.Context
	ApiService *PeopleProfilesAPIService
	userid string
	updateMemberRequestDto *UpdateMemberRequestDto
}

// The request parameters for updating the user information.
func (r ApiUpdateMemberRequest) UpdateMemberRequestDto(updateMemberRequestDto UpdateMemberRequestDto) ApiUpdateMemberRequest {	r.updateMemberRequestDto = &updateMemberRequestDto
	return r
}

func (r ApiUpdateMemberRequest) Execute() (*EmployeeFullWrapper, *http.Response, error) {
	return r.ApiService.UpdateMemberExecute(r)
}

// UpdateMember Update a user
//
// Updates the data for the selected portal user with the first name, last name, email address, and/or optional parameters specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/update-member/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param userid The user ID.
// @return ApiUpdateMemberRequest
func (a *PeopleProfilesAPIService) UpdateMember(ctx context.Context, userid string) ApiUpdateMemberRequest {
	return ApiUpdateMemberRequest{
		ApiService: a,
		ctx: ctx,
		userid: userid,
	}
}

// Execute executes the request
//  @return EmployeeFullWrapper
func (a *PeopleProfilesAPIService) UpdateMemberExecute(r ApiUpdateMemberRequest) (*EmployeeFullWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeFullWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleProfilesAPIService.UpdateMember")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/{userid}"
	localVarPath = strings.Replace(localVarPath, "{"+"userid"+"}", url.PathEscape(parameterValueToString(r.userid, "userid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateMemberRequestDto == nil {
		return localVarReturnValue, nil, reportError("updateMemberRequestDto is required and must be specified")
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
	localVarPostBody = r.updateMemberRequestDto
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

type ApiUpdateMemberCultureRequest struct {
	ctx context.Context
	ApiService *PeopleProfilesAPIService
	userid string
	culture *Culture
}

// The culture name parameters.
func (r ApiUpdateMemberCultureRequest) Culture(culture Culture) ApiUpdateMemberCultureRequest {	r.culture = &culture
	return r
}

func (r ApiUpdateMemberCultureRequest) Execute() (*EmployeeFullWrapper, *http.Response, error) {
	return r.ApiService.UpdateMemberCultureExecute(r)
}

// UpdateMemberCulture Update a user culture
//
// Updates the user culture with the parameters specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/update-member-culture/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @param userid The user ID.
// @return ApiUpdateMemberCultureRequest
func (a *PeopleProfilesAPIService) UpdateMemberCulture(ctx context.Context, userid string) ApiUpdateMemberCultureRequest {
	return ApiUpdateMemberCultureRequest{
		ApiService: a,
		ctx: ctx,
		userid: userid,
	}
}

// Execute executes the request
//  @return EmployeeFullWrapper
func (a *PeopleProfilesAPIService) UpdateMemberCultureExecute(r ApiUpdateMemberCultureRequest) (*EmployeeFullWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployeeFullWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PeopleProfilesAPIService.UpdateMemberCulture")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/people/{userid}/culture"
	localVarPath = strings.Replace(localVarPath, "{"+"userid"+"}", url.PathEscape(parameterValueToString(r.userid, "userid")), -1)

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
	localVarPostBody = r.culture
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
