# \PeopleUserTypeAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserTypeUpdateProgress**](PeopleUserTypeAPI.md#GetUserTypeUpdateProgress) | **Get** /api/2.0/people/type/progress/{userid} | Get the progress of updating user type
[**StartUserTypeUpdate**](PeopleUserTypeAPI.md#StartUserTypeUpdate) | **Post** /api/2.0/people/type | Start updating user type
[**TerminateUserTypeUpdate**](PeopleUserTypeAPI.md#TerminateUserTypeUpdate) | **Put** /api/2.0/people/type/terminate | Terminate updating user type
[**UpdateUserType**](PeopleUserTypeAPI.md#UpdateUserType) | **Put** /api/2.0/people/type/{type} | Change a user type



## GetUserTypeUpdateProgress

> TaskProgressResponseWrapper GetUserTypeUpdateProgress(ctx, userid).Execute()

Get the progress of updating user type



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-user-type-update-progress/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	userid := "00000000-0000-0000-0000-000000000000" // string | The user ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleUserTypeAPI.GetUserTypeUpdateProgress(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleUserTypeAPI.GetUserTypeUpdateProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserTypeUpdateProgress`: TaskProgressResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleUserTypeAPI.GetUserTypeUpdateProgress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTypeUpdateProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskProgressResponseWrapper**](TaskProgressResponseWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartUserTypeUpdate

> TaskProgressResponseWrapper StartUserTypeUpdate(ctx).StartUpdateUserTypeDto(startUpdateUserTypeDto).Execute()

Start updating user type



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/start-user-type-update/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	startUpdateUserTypeDto := *openapiclient.NewStartUpdateUserTypeDto() // StartUpdateUserTypeDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleUserTypeAPI.StartUserTypeUpdate(context.Background()).StartUpdateUserTypeDto(startUpdateUserTypeDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleUserTypeAPI.StartUserTypeUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartUserTypeUpdate`: TaskProgressResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleUserTypeAPI.StartUserTypeUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartUserTypeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startUpdateUserTypeDto** | [**StartUpdateUserTypeDto**](StartUpdateUserTypeDto.md) |  | 

### Return type

[**TaskProgressResponseWrapper**](TaskProgressResponseWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateUserTypeUpdate

> TaskProgressResponseWrapper TerminateUserTypeUpdate(ctx).TerminateRequestDto(terminateRequestDto).Execute()

Terminate updating user type



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/terminate-user-type-update/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	terminateRequestDto := *openapiclient.NewTerminateRequestDto("00000000-0000-0000-0000-000000000000") // TerminateRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleUserTypeAPI.TerminateUserTypeUpdate(context.Background()).TerminateRequestDto(terminateRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleUserTypeAPI.TerminateUserTypeUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TerminateUserTypeUpdate`: TaskProgressResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleUserTypeAPI.TerminateUserTypeUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTerminateUserTypeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **terminateRequestDto** | [**TerminateRequestDto**](TerminateRequestDto.md) |  | 

### Return type

[**TaskProgressResponseWrapper**](TaskProgressResponseWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserType

> EmployeeFullArrayWrapper UpdateUserType(ctx, type_).UpdateMembersRequestDto(updateMembersRequestDto).Execute()

Change a user type



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-user-type/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	type_ := openapiclient.EmployeeType("All") // EmployeeType | The new user type.
	updateMembersRequestDto := *openapiclient.NewUpdateMembersRequestDto() // UpdateMembersRequestDto | The request parameters for updating the user information.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleUserTypeAPI.UpdateUserType(context.Background(), type_).UpdateMembersRequestDto(updateMembersRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleUserTypeAPI.UpdateUserType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserType`: EmployeeFullArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleUserTypeAPI.UpdateUserType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**EmployeeType**](.md) | The new user type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMembersRequestDto** | [**UpdateMembersRequestDto**](UpdateMembersRequestDto.md) | The request parameters for updating the user information. | 

### Return type

[**EmployeeFullArrayWrapper**](EmployeeFullArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

