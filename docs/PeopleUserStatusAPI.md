# \PeopleUserStatusAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetByStatus**](PeopleUserStatusAPI.md#GetByStatus) | **Get** /api/2.0/people/status/{status} | Get profiles by status
[**UpdateUserActivationStatus**](PeopleUserStatusAPI.md#UpdateUserActivationStatus) | **Put** /api/2.0/people/activationstatus/{activationstatus} | Set an activation status to the users
[**UpdateUserStatus**](PeopleUserStatusAPI.md#UpdateUserStatus) | **Put** /api/2.0/people/status/{status} | Change a user status



## GetByStatus

> EmployeeFullArrayWrapper GetByStatus(ctx, status).FilterBy(filterBy).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterSeparator(filterSeparator).FilterValue(filterValue).Execute()

Get profiles by status



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-by-status/).

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
	status := openapiclient.EmployeeStatus(1) // EmployeeStatus | The user status.
	filterBy := "displayName" // string | Specifies the criteria used to filter the profiles in the request. (optional)
	count := int32(25) // int32 | The maximum number of user profiles to retrieve. (optional)
	startIndex := int32(0) // int32 | The starting index for retrieving data in a paginated request. (optional)
	sortBy := "displayName" // string | Specifies the property or field name by which the results should be sorted. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The order in which the results are sorted. (optional)
	filterSeparator := "," // string | Represents the separator used to split multiple filter criteria in a query string. (optional)
	filterValue := "John" // string | A string value representing additional filter criteria used in query parameters. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleUserStatusAPI.GetByStatus(context.Background(), status).FilterBy(filterBy).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterSeparator(filterSeparator).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleUserStatusAPI.GetByStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetByStatus`: EmployeeFullArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleUserStatusAPI.GetByStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**status** | [**EmployeeStatus**](.md) | The user status. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterBy** | **string** | Specifies the criteria used to filter the profiles in the request. | 
 **count** | **int32** | The maximum number of user profiles to retrieve. | 
 **startIndex** | **int32** | The starting index for retrieving data in a paginated request. | 
 **sortBy** | **string** | Specifies the property or field name by which the results should be sorted. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The order in which the results are sorted. | 
 **filterSeparator** | **string** | Represents the separator used to split multiple filter criteria in a query string. | 
 **filterValue** | **string** | A string value representing additional filter criteria used in query parameters. | 

### Return type

[**EmployeeFullArrayWrapper**](EmployeeFullArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserActivationStatus

> EmployeeFullArrayWrapper UpdateUserActivationStatus(ctx, activationstatus).UpdateMembersRequestDto(updateMembersRequestDto).Execute()

Set an activation status to the users



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-user-activation-status/).

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
	activationstatus := openapiclient.EmployeeActivationStatus(0) // EmployeeActivationStatus | The new user activation status.
	updateMembersRequestDto := *openapiclient.NewUpdateMembersRequestDto() // UpdateMembersRequestDto | The request parameters for updating the user information.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleUserStatusAPI.UpdateUserActivationStatus(context.Background(), activationstatus).UpdateMembersRequestDto(updateMembersRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleUserStatusAPI.UpdateUserActivationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserActivationStatus`: EmployeeFullArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleUserStatusAPI.UpdateUserActivationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activationstatus** | [**EmployeeActivationStatus**](.md) | The new user activation status. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserActivationStatusRequest struct via the builder pattern


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


## UpdateUserStatus

> EmployeeFullArrayWrapper UpdateUserStatus(ctx, status).UpdateMembersRequestDto(updateMembersRequestDto).Execute()

Change a user status



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-user-status/).

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
	status := openapiclient.EmployeeStatus(1) // EmployeeStatus | The new user status.
	updateMembersRequestDto := *openapiclient.NewUpdateMembersRequestDto() // UpdateMembersRequestDto | The request parameters for updating the user information.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleUserStatusAPI.UpdateUserStatus(context.Background(), status).UpdateMembersRequestDto(updateMembersRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleUserStatusAPI.UpdateUserStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserStatus`: EmployeeFullArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleUserStatusAPI.UpdateUserStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**status** | [**EmployeeStatus**](.md) | The new user status. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserStatusRequest struct via the builder pattern


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

