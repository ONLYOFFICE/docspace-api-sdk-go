# \PeopleSearchAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountsEntriesWithFilesShared**](PeopleSearchAPI.md#GetAccountsEntriesWithFilesShared) | **Get** /api/2.0/accounts/file/{id}/search | Get account entries with file sharing settings
[**GetAccountsEntriesWithFoldersShared**](PeopleSearchAPI.md#GetAccountsEntriesWithFoldersShared) | **Get** /api/2.0/accounts/folder/{id}/search | Get account entries with folder sharing settings
[**GetAccountsEntriesWithRoomsShared**](PeopleSearchAPI.md#GetAccountsEntriesWithRoomsShared) | **Get** /api/2.0/accounts/room/{id}/search | Get account entries
[**GetSearch**](PeopleSearchAPI.md#GetSearch) | **Get** /api/2.0/people/@search/{query} | Search users
[**GetSimpleByFilter**](PeopleSearchAPI.md#GetSimpleByFilter) | **Get** /api/2.0/people/simple/filter | Search users by extended filter
[**GetUsersWithFilesShared**](PeopleSearchAPI.md#GetUsersWithFilesShared) | **Get** /api/2.0/people/file/{id} | Get users with file sharing settings
[**GetUsersWithFoldersShared**](PeopleSearchAPI.md#GetUsersWithFoldersShared) | **Get** /api/2.0/people/folder/{id} | Get users with folder sharing settings
[**GetUsersWithRoomShared**](PeopleSearchAPI.md#GetUsersWithRoomShared) | **Get** /api/2.0/people/room/{id} | Get users with room sharing settings
[**SearchUsersByExtendedFilter**](PeopleSearchAPI.md#SearchUsersByExtendedFilter) | **Get** /api/2.0/people/filter | Search users with detailed information by extended filter
[**SearchUsersByQuery**](PeopleSearchAPI.md#SearchUsersByQuery) | **Get** /api/2.0/people/search | Search users (using query parameters)
[**SearchUsersByStatus**](PeopleSearchAPI.md#SearchUsersByStatus) | **Get** /api/2.0/people/status/{status}/search | Search users by status filter



## GetAccountsEntriesWithFilesShared

> ObjectArrayWrapper GetAccountsEntriesWithFilesShared(ctx, id).EmployeeStatus(employeeStatus).ActivationStatus(activationStatus).ExcludeShared(excludeShared).IncludeShared(includeShared).InvitedByMe(invitedByMe).InviterId(inviterId).Area(area).EmployeeTypes(employeeTypes).Count(count).StartIndex(startIndex).FilterSeparator(filterSeparator).FilterValue(filterValue).Execute()

Get account entries with file sharing settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-accounts-entries-with-files-shared/).

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
	id := int32(00000000-0000-0000-0000-000000000000) // int32 | The user ID.
	employeeStatus := openapiclient.EmployeeStatus(1) // EmployeeStatus | The user status. (optional)
	activationStatus := openapiclient.EmployeeActivationStatus(0) // EmployeeActivationStatus | The user activation status. (optional)
	excludeShared := false // bool | Specifies whether to exclude the account sharing settings from the response. (optional)
	includeShared := false // bool | Specifies whether to include the account sharing settings in the response. (optional)
	invitedByMe := false // bool | Specifies whether the user is invited by the current user or not. (optional)
	inviterId := "00000000-0000-0000-0000-000000000000" // string | The inviter ID. (optional)
	area := openapiclient.Area(0) // Area | The area of the account entries. (optional)
	employeeTypes := []openapiclient.EmployeeType{openapiclient.EmployeeType("All")} // []EmployeeType | The list of the user types. (optional)
	count := int32(25) // int32 | The number of items to retrieve in a request. (optional)
	startIndex := int32(0) // int32 | The starting index for the query results. (optional)
	filterSeparator := "," // string | Specifies the separator used in filter expressions. (optional)
	filterValue := "John" // string | The text filter applied to the accounts search query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleSearchAPI.GetAccountsEntriesWithFilesShared(context.Background(), id).EmployeeStatus(employeeStatus).ActivationStatus(activationStatus).ExcludeShared(excludeShared).IncludeShared(includeShared).InvitedByMe(invitedByMe).InviterId(inviterId).Area(area).EmployeeTypes(employeeTypes).Count(count).StartIndex(startIndex).FilterSeparator(filterSeparator).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleSearchAPI.GetAccountsEntriesWithFilesShared``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountsEntriesWithFilesShared`: ObjectArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleSearchAPI.GetAccountsEntriesWithFilesShared`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsEntriesWithFilesSharedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **employeeStatus** | [**EmployeeStatus**](EmployeeStatus.md) | The user status. | 
 **activationStatus** | [**EmployeeActivationStatus**](EmployeeActivationStatus.md) | The user activation status. | 
 **excludeShared** | **bool** | Specifies whether to exclude the account sharing settings from the response. | 
 **includeShared** | **bool** | Specifies whether to include the account sharing settings in the response. | 
 **invitedByMe** | **bool** | Specifies whether the user is invited by the current user or not. | 
 **inviterId** | **string** | The inviter ID. | 
 **area** | [**Area**](Area.md) | The area of the account entries. | 
 **employeeTypes** | [**[]EmployeeType**](EmployeeType.md) | The list of the user types. | 
 **count** | **int32** | The number of items to retrieve in a request. | 
 **startIndex** | **int32** | The starting index for the query results. | 
 **filterSeparator** | **string** | Specifies the separator used in filter expressions. | 
 **filterValue** | **string** | The text filter applied to the accounts search query. | 

### Return type

[**ObjectArrayWrapper**](ObjectArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountsEntriesWithFoldersShared

> ObjectArrayWrapper GetAccountsEntriesWithFoldersShared(ctx, id).EmployeeStatus(employeeStatus).ActivationStatus(activationStatus).ExcludeShared(excludeShared).IncludeShared(includeShared).InvitedByMe(invitedByMe).InviterId(inviterId).Area(area).EmployeeTypes(employeeTypes).Count(count).StartIndex(startIndex).FilterSeparator(filterSeparator).FilterValue(filterValue).Execute()

Get account entries with folder sharing settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-accounts-entries-with-folders-shared/).

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
	id := int32(00000000-0000-0000-0000-000000000000) // int32 | The user ID.
	employeeStatus := openapiclient.EmployeeStatus(1) // EmployeeStatus | The user status. (optional)
	activationStatus := openapiclient.EmployeeActivationStatus(0) // EmployeeActivationStatus | The user activation status. (optional)
	excludeShared := false // bool | Specifies whether to exclude the account sharing settings from the response. (optional)
	includeShared := false // bool | Specifies whether to include the account sharing settings in the response. (optional)
	invitedByMe := false // bool | Specifies whether the user is invited by the current user or not. (optional)
	inviterId := "00000000-0000-0000-0000-000000000000" // string | The inviter ID. (optional)
	area := openapiclient.Area(0) // Area | The area of the account entries. (optional)
	employeeTypes := []openapiclient.EmployeeType{openapiclient.EmployeeType("All")} // []EmployeeType | The list of the user types. (optional)
	count := int32(25) // int32 | The number of items to retrieve in a request. (optional)
	startIndex := int32(0) // int32 | The starting index for the query results. (optional)
	filterSeparator := "," // string | Specifies the separator used in filter expressions. (optional)
	filterValue := "John" // string | The text filter applied to the accounts search query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleSearchAPI.GetAccountsEntriesWithFoldersShared(context.Background(), id).EmployeeStatus(employeeStatus).ActivationStatus(activationStatus).ExcludeShared(excludeShared).IncludeShared(includeShared).InvitedByMe(invitedByMe).InviterId(inviterId).Area(area).EmployeeTypes(employeeTypes).Count(count).StartIndex(startIndex).FilterSeparator(filterSeparator).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleSearchAPI.GetAccountsEntriesWithFoldersShared``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountsEntriesWithFoldersShared`: ObjectArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleSearchAPI.GetAccountsEntriesWithFoldersShared`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsEntriesWithFoldersSharedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **employeeStatus** | [**EmployeeStatus**](EmployeeStatus.md) | The user status. | 
 **activationStatus** | [**EmployeeActivationStatus**](EmployeeActivationStatus.md) | The user activation status. | 
 **excludeShared** | **bool** | Specifies whether to exclude the account sharing settings from the response. | 
 **includeShared** | **bool** | Specifies whether to include the account sharing settings in the response. | 
 **invitedByMe** | **bool** | Specifies whether the user is invited by the current user or not. | 
 **inviterId** | **string** | The inviter ID. | 
 **area** | [**Area**](Area.md) | The area of the account entries. | 
 **employeeTypes** | [**[]EmployeeType**](EmployeeType.md) | The list of the user types. | 
 **count** | **int32** | The number of items to retrieve in a request. | 
 **startIndex** | **int32** | The starting index for the query results. | 
 **filterSeparator** | **string** | Specifies the separator used in filter expressions. | 
 **filterValue** | **string** | The text filter applied to the accounts search query. | 

### Return type

[**ObjectArrayWrapper**](ObjectArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountsEntriesWithRoomsShared

> ObjectArrayWrapper GetAccountsEntriesWithRoomsShared(ctx, id).EmployeeStatus(employeeStatus).ActivationStatus(activationStatus).ExcludeShared(excludeShared).IncludeShared(includeShared).InvitedByMe(invitedByMe).InviterId(inviterId).Area(area).EmployeeTypes(employeeTypes).Count(count).StartIndex(startIndex).FilterSeparator(filterSeparator).FilterValue(filterValue).Execute()

Get account entries



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-accounts-entries-with-rooms-shared/).

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
	id := int32(00000000-0000-0000-0000-000000000000) // int32 | The user ID.
	employeeStatus := openapiclient.EmployeeStatus(1) // EmployeeStatus | The user status. (optional)
	activationStatus := openapiclient.EmployeeActivationStatus(0) // EmployeeActivationStatus | The user activation status. (optional)
	excludeShared := false // bool | Specifies whether to exclude the account sharing settings from the response. (optional)
	includeShared := false // bool | Specifies whether to include the account sharing settings in the response. (optional)
	invitedByMe := false // bool | Specifies whether the user is invited by the current user or not. (optional)
	inviterId := "00000000-0000-0000-0000-000000000000" // string | The inviter ID. (optional)
	area := openapiclient.Area(0) // Area | The area of the account entries. (optional)
	employeeTypes := []openapiclient.EmployeeType{openapiclient.EmployeeType("All")} // []EmployeeType | The list of the user types. (optional)
	count := int32(25) // int32 | The number of items to retrieve in a request. (optional)
	startIndex := int32(0) // int32 | The starting index for the query results. (optional)
	filterSeparator := "," // string | Specifies the separator used in filter expressions. (optional)
	filterValue := "John" // string | The text filter applied to the accounts search query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleSearchAPI.GetAccountsEntriesWithRoomsShared(context.Background(), id).EmployeeStatus(employeeStatus).ActivationStatus(activationStatus).ExcludeShared(excludeShared).IncludeShared(includeShared).InvitedByMe(invitedByMe).InviterId(inviterId).Area(area).EmployeeTypes(employeeTypes).Count(count).StartIndex(startIndex).FilterSeparator(filterSeparator).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleSearchAPI.GetAccountsEntriesWithRoomsShared``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountsEntriesWithRoomsShared`: ObjectArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleSearchAPI.GetAccountsEntriesWithRoomsShared`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsEntriesWithRoomsSharedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **employeeStatus** | [**EmployeeStatus**](EmployeeStatus.md) | The user status. | 
 **activationStatus** | [**EmployeeActivationStatus**](EmployeeActivationStatus.md) | The user activation status. | 
 **excludeShared** | **bool** | Specifies whether to exclude the account sharing settings from the response. | 
 **includeShared** | **bool** | Specifies whether to include the account sharing settings in the response. | 
 **invitedByMe** | **bool** | Specifies whether the user is invited by the current user or not. | 
 **inviterId** | **string** | The inviter ID. | 
 **area** | [**Area**](Area.md) | The area of the account entries. | 
 **employeeTypes** | [**[]EmployeeType**](EmployeeType.md) | The list of the user types. | 
 **count** | **int32** | The number of items to retrieve in a request. | 
 **startIndex** | **int32** | The starting index for the query results. | 
 **filterSeparator** | **string** | Specifies the separator used in filter expressions. | 
 **filterValue** | **string** | The text filter applied to the accounts search query. | 

### Return type

[**ObjectArrayWrapper**](ObjectArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearch

> EmployeeFullArrayWrapper GetSearch(ctx, query).FilterBy(filterBy).FilterValue(filterValue).Execute()

Search users



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-search/).

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
	query := "John" // string | The search query.
	filterBy := "displayName" // string | Specifies a filter criteria for the user search query. (optional)
	filterValue := "John" // string | The value used for filtering users, allowing additional constraints for the query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleSearchAPI.GetSearch(context.Background(), query).FilterBy(filterBy).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleSearchAPI.GetSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearch`: EmployeeFullArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleSearchAPI.GetSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**query** | **string** | The search query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterBy** | **string** | Specifies a filter criteria for the user search query. | 
 **filterValue** | **string** | The value used for filtering users, allowing additional constraints for the query. | 

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


## GetSimpleByFilter

> EmployeeArrayWrapper GetSimpleByFilter(ctx).EmployeeStatus(employeeStatus).GroupId(groupId).ActivationStatus(activationStatus).EmployeeType(employeeType).EmployeeTypes(employeeTypes).IsAdministrator(isAdministrator).Payments(payments).AccountLoginType(accountLoginType).QuotaFilter(quotaFilter).WithoutGroup(withoutGroup).ExcludeGroup(excludeGroup).InvitedByMe(invitedByMe).InviterId(inviterId).Area(area).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterSeparator(filterSeparator).FilterValue(filterValue).Execute()

Search users by extended filter



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-simple-by-filter/).

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
	employeeStatus := openapiclient.EmployeeStatus(1) // EmployeeStatus | The user status. (optional)
	groupId := "00000000-0000-0000-0000-000000000000" // string | The group ID. (optional)
	activationStatus := openapiclient.EmployeeActivationStatus(0) // EmployeeActivationStatus | The user activation status. (optional)
	employeeType := openapiclient.EmployeeType("All") // EmployeeType | The user type. (optional)
	employeeTypes := []int32{int32(0)} // []int32 | The list of user types. (optional)
	isAdministrator := false // bool | Specifies if the user is an administrator or not. (optional)
	payments := openapiclient.Payments(0) // Payments | The user payment status. (optional)
	accountLoginType := openapiclient.AccountLoginType(0) // AccountLoginType | The account login type. (optional)
	quotaFilter := openapiclient.QuotaFilter(0) // QuotaFilter | The quota filter (All - 0, Default - 1, Custom - 2). (optional)
	withoutGroup := false // bool | Specifies whether the user should be a member of a group or not. (optional)
	excludeGroup := false // bool | Specifies whether the user should be a member of the group with the specified ID. (optional)
	invitedByMe := false // bool | Specifies whether the user is invited by the current user or not. (optional)
	inviterId := "00000000-0000-0000-0000-000000000000" // string | The inviter ID. (optional)
	area := openapiclient.Area(0) // Area | The filter area. (optional)
	count := int32(25) // int32 | The maximum number of items to be retrieved in the response. (optional)
	startIndex := int32(0) // int32 | The zero-based index of the first item to be retrieved in a filtered result set. (optional)
	sortBy := "displayName" // string | Specifies the property or field name by which the results should be sorted. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The order in which the results are sorted. (optional)
	filterSeparator := "," // string | Represents the separator used to split filter criteria in query parameters. (optional)
	filterValue := "John" // string | The search text used to filter results based on user input. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleSearchAPI.GetSimpleByFilter(context.Background()).EmployeeStatus(employeeStatus).GroupId(groupId).ActivationStatus(activationStatus).EmployeeType(employeeType).EmployeeTypes(employeeTypes).IsAdministrator(isAdministrator).Payments(payments).AccountLoginType(accountLoginType).QuotaFilter(quotaFilter).WithoutGroup(withoutGroup).ExcludeGroup(excludeGroup).InvitedByMe(invitedByMe).InviterId(inviterId).Area(area).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterSeparator(filterSeparator).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleSearchAPI.GetSimpleByFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimpleByFilter`: EmployeeArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleSearchAPI.GetSimpleByFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **employeeStatus** | [**EmployeeStatus**](EmployeeStatus.md) | The user status. | 
 **groupId** | **string** | The group ID. | 
 **activationStatus** | [**EmployeeActivationStatus**](EmployeeActivationStatus.md) | The user activation status. | 
 **employeeType** | [**EmployeeType**](EmployeeType.md) | The user type. | 
 **employeeTypes** | **[]int32** | The list of user types. | 
 **isAdministrator** | **bool** | Specifies if the user is an administrator or not. | 
 **payments** | [**Payments**](Payments.md) | The user payment status. | 
 **accountLoginType** | [**AccountLoginType**](AccountLoginType.md) | The account login type. | 
 **quotaFilter** | [**QuotaFilter**](QuotaFilter.md) | The quota filter (All - 0, Default - 1, Custom - 2). | 
 **withoutGroup** | **bool** | Specifies whether the user should be a member of a group or not. | 
 **excludeGroup** | **bool** | Specifies whether the user should be a member of the group with the specified ID. | 
 **invitedByMe** | **bool** | Specifies whether the user is invited by the current user or not. | 
 **inviterId** | **string** | The inviter ID. | 
 **area** | [**Area**](Area.md) | The filter area. | 
 **count** | **int32** | The maximum number of items to be retrieved in the response. | 
 **startIndex** | **int32** | The zero-based index of the first item to be retrieved in a filtered result set. | 
 **sortBy** | **string** | Specifies the property or field name by which the results should be sorted. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The order in which the results are sorted. | 
 **filterSeparator** | **string** | Represents the separator used to split filter criteria in query parameters. | 
 **filterValue** | **string** | The search text used to filter results based on user input. | 

### Return type

[**EmployeeArrayWrapper**](EmployeeArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersWithFilesShared

> EmployeeFullArrayWrapper GetUsersWithFilesShared(ctx, id).EmployeeStatus(employeeStatus).ActivationStatus(activationStatus).ExcludeShared(excludeShared).IncludeShared(includeShared).InvitedByMe(invitedByMe).InviterId(inviterId).Area(area).EmployeeTypes(employeeTypes).Count(count).StartIndex(startIndex).FilterSeparator(filterSeparator).FilterValue(filterValue).Execute()

Get users with file sharing settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-users-with-files-shared/).

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
	id := int32(00000000-0000-0000-0000-000000000000) // int32 | The user ID.
	employeeStatus := openapiclient.EmployeeStatus(1) // EmployeeStatus | The user status. (optional)
	activationStatus := openapiclient.EmployeeActivationStatus(0) // EmployeeActivationStatus | The user activation status. (optional)
	excludeShared := false // bool | Specifies whether to exclude the user sharing settings or not. (optional)
	includeShared := false // bool | Specifies whether to include the user sharing settings or not. (optional)
	invitedByMe := false // bool | Specifies whether the user was invited by the current user or not. (optional)
	inviterId := "00000000-0000-0000-0000-000000000000" // string | The inviter ID. (optional)
	area := openapiclient.Area(0) // Area | The user area. (optional)
	employeeTypes := []openapiclient.EmployeeType{openapiclient.EmployeeType("All")} // []EmployeeType | The list of user types. (optional)
	count := int32(25) // int32 | The maximum number of users to be retrieved in the request. (optional)
	startIndex := int32(0) // int32 | The zero-based index of the first record to retrieve in a paged query. (optional)
	filterSeparator := "," // string | The character or string used to separate multiple filter values in a filtering query. (optional)
	filterValue := "John" // string | The filter text value used for searching or filtering user results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleSearchAPI.GetUsersWithFilesShared(context.Background(), id).EmployeeStatus(employeeStatus).ActivationStatus(activationStatus).ExcludeShared(excludeShared).IncludeShared(includeShared).InvitedByMe(invitedByMe).InviterId(inviterId).Area(area).EmployeeTypes(employeeTypes).Count(count).StartIndex(startIndex).FilterSeparator(filterSeparator).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleSearchAPI.GetUsersWithFilesShared``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersWithFilesShared`: EmployeeFullArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleSearchAPI.GetUsersWithFilesShared`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersWithFilesSharedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **employeeStatus** | [**EmployeeStatus**](EmployeeStatus.md) | The user status. | 
 **activationStatus** | [**EmployeeActivationStatus**](EmployeeActivationStatus.md) | The user activation status. | 
 **excludeShared** | **bool** | Specifies whether to exclude the user sharing settings or not. | 
 **includeShared** | **bool** | Specifies whether to include the user sharing settings or not. | 
 **invitedByMe** | **bool** | Specifies whether the user was invited by the current user or not. | 
 **inviterId** | **string** | The inviter ID. | 
 **area** | [**Area**](Area.md) | The user area. | 
 **employeeTypes** | [**[]EmployeeType**](EmployeeType.md) | The list of user types. | 
 **count** | **int32** | The maximum number of users to be retrieved in the request. | 
 **startIndex** | **int32** | The zero-based index of the first record to retrieve in a paged query. | 
 **filterSeparator** | **string** | The character or string used to separate multiple filter values in a filtering query. | 
 **filterValue** | **string** | The filter text value used for searching or filtering user results. | 

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


## GetUsersWithFoldersShared

> EmployeeFullArrayWrapper GetUsersWithFoldersShared(ctx, id).EmployeeStatus(employeeStatus).ActivationStatus(activationStatus).ExcludeShared(excludeShared).IncludeShared(includeShared).InvitedByMe(invitedByMe).InviterId(inviterId).Area(area).EmployeeTypes(employeeTypes).Count(count).StartIndex(startIndex).FilterSeparator(filterSeparator).FilterValue(filterValue).Execute()

Get users with folder sharing settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-users-with-folders-shared/).

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
	id := int32(00000000-0000-0000-0000-000000000000) // int32 | The user ID.
	employeeStatus := openapiclient.EmployeeStatus(1) // EmployeeStatus | The user status. (optional)
	activationStatus := openapiclient.EmployeeActivationStatus(0) // EmployeeActivationStatus | The user activation status. (optional)
	excludeShared := false // bool | Specifies whether to exclude the user sharing settings or not. (optional)
	includeShared := false // bool | Specifies whether to include the user sharing settings or not. (optional)
	invitedByMe := false // bool | Specifies whether the user was invited by the current user or not. (optional)
	inviterId := "00000000-0000-0000-0000-000000000000" // string | The inviter ID. (optional)
	area := openapiclient.Area(0) // Area | The user area. (optional)
	employeeTypes := []openapiclient.EmployeeType{openapiclient.EmployeeType("All")} // []EmployeeType | The list of user types. (optional)
	count := int32(25) // int32 | The maximum number of users to be retrieved in the request. (optional)
	startIndex := int32(0) // int32 | The zero-based index of the first record to retrieve in a paged query. (optional)
	filterSeparator := "," // string | The character or string used to separate multiple filter values in a filtering query. (optional)
	filterValue := "John" // string | The filter text value used for searching or filtering user results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleSearchAPI.GetUsersWithFoldersShared(context.Background(), id).EmployeeStatus(employeeStatus).ActivationStatus(activationStatus).ExcludeShared(excludeShared).IncludeShared(includeShared).InvitedByMe(invitedByMe).InviterId(inviterId).Area(area).EmployeeTypes(employeeTypes).Count(count).StartIndex(startIndex).FilterSeparator(filterSeparator).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleSearchAPI.GetUsersWithFoldersShared``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersWithFoldersShared`: EmployeeFullArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleSearchAPI.GetUsersWithFoldersShared`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersWithFoldersSharedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **employeeStatus** | [**EmployeeStatus**](EmployeeStatus.md) | The user status. | 
 **activationStatus** | [**EmployeeActivationStatus**](EmployeeActivationStatus.md) | The user activation status. | 
 **excludeShared** | **bool** | Specifies whether to exclude the user sharing settings or not. | 
 **includeShared** | **bool** | Specifies whether to include the user sharing settings or not. | 
 **invitedByMe** | **bool** | Specifies whether the user was invited by the current user or not. | 
 **inviterId** | **string** | The inviter ID. | 
 **area** | [**Area**](Area.md) | The user area. | 
 **employeeTypes** | [**[]EmployeeType**](EmployeeType.md) | The list of user types. | 
 **count** | **int32** | The maximum number of users to be retrieved in the request. | 
 **startIndex** | **int32** | The zero-based index of the first record to retrieve in a paged query. | 
 **filterSeparator** | **string** | The character or string used to separate multiple filter values in a filtering query. | 
 **filterValue** | **string** | The filter text value used for searching or filtering user results. | 

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


## GetUsersWithRoomShared

> EmployeeFullArrayWrapper GetUsersWithRoomShared(ctx, id).EmployeeStatus(employeeStatus).ActivationStatus(activationStatus).ExcludeShared(excludeShared).IncludeShared(includeShared).InvitedByMe(invitedByMe).InviterId(inviterId).Area(area).EmployeeTypes(employeeTypes).Count(count).StartIndex(startIndex).FilterSeparator(filterSeparator).FilterValue(filterValue).Execute()

Get users with room sharing settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-users-with-room-shared/).

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
	id := int32(00000000-0000-0000-0000-000000000000) // int32 | The user ID.
	employeeStatus := openapiclient.EmployeeStatus(1) // EmployeeStatus | The user status. (optional)
	activationStatus := openapiclient.EmployeeActivationStatus(0) // EmployeeActivationStatus | The user activation status. (optional)
	excludeShared := false // bool | Specifies whether to exclude the user sharing settings or not. (optional)
	includeShared := false // bool | Specifies whether to include the user sharing settings or not. (optional)
	invitedByMe := false // bool | Specifies whether the user was invited by the current user or not. (optional)
	inviterId := "00000000-0000-0000-0000-000000000000" // string | The inviter ID. (optional)
	area := openapiclient.Area(0) // Area | The user area. (optional)
	employeeTypes := []openapiclient.EmployeeType{openapiclient.EmployeeType("All")} // []EmployeeType | The list of user types. (optional)
	count := int32(25) // int32 | The maximum number of users to be retrieved in the request. (optional)
	startIndex := int32(0) // int32 | The zero-based index of the first record to retrieve in a paged query. (optional)
	filterSeparator := "," // string | The character or string used to separate multiple filter values in a filtering query. (optional)
	filterValue := "John" // string | The filter text value used for searching or filtering user results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleSearchAPI.GetUsersWithRoomShared(context.Background(), id).EmployeeStatus(employeeStatus).ActivationStatus(activationStatus).ExcludeShared(excludeShared).IncludeShared(includeShared).InvitedByMe(invitedByMe).InviterId(inviterId).Area(area).EmployeeTypes(employeeTypes).Count(count).StartIndex(startIndex).FilterSeparator(filterSeparator).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleSearchAPI.GetUsersWithRoomShared``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersWithRoomShared`: EmployeeFullArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleSearchAPI.GetUsersWithRoomShared`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersWithRoomSharedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **employeeStatus** | [**EmployeeStatus**](EmployeeStatus.md) | The user status. | 
 **activationStatus** | [**EmployeeActivationStatus**](EmployeeActivationStatus.md) | The user activation status. | 
 **excludeShared** | **bool** | Specifies whether to exclude the user sharing settings or not. | 
 **includeShared** | **bool** | Specifies whether to include the user sharing settings or not. | 
 **invitedByMe** | **bool** | Specifies whether the user was invited by the current user or not. | 
 **inviterId** | **string** | The inviter ID. | 
 **area** | [**Area**](Area.md) | The user area. | 
 **employeeTypes** | [**[]EmployeeType**](EmployeeType.md) | The list of user types. | 
 **count** | **int32** | The maximum number of users to be retrieved in the request. | 
 **startIndex** | **int32** | The zero-based index of the first record to retrieve in a paged query. | 
 **filterSeparator** | **string** | The character or string used to separate multiple filter values in a filtering query. | 
 **filterValue** | **string** | The filter text value used for searching or filtering user results. | 

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


## SearchUsersByExtendedFilter

> EmployeeFullArrayWrapper SearchUsersByExtendedFilter(ctx).EmployeeStatus(employeeStatus).GroupId(groupId).ActivationStatus(activationStatus).EmployeeType(employeeType).EmployeeTypes(employeeTypes).IsAdministrator(isAdministrator).Payments(payments).AccountLoginType(accountLoginType).QuotaFilter(quotaFilter).WithoutGroup(withoutGroup).ExcludeGroup(excludeGroup).InvitedByMe(invitedByMe).InviterId(inviterId).Area(area).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterSeparator(filterSeparator).FilterValue(filterValue).Execute()

Search users with detailed information by extended filter



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/search-users-by-extended-filter/).

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
	employeeStatus := openapiclient.EmployeeStatus(1) // EmployeeStatus | The user status. (optional)
	groupId := "00000000-0000-0000-0000-000000000000" // string | The group ID. (optional)
	activationStatus := openapiclient.EmployeeActivationStatus(0) // EmployeeActivationStatus | The user activation status. (optional)
	employeeType := openapiclient.EmployeeType("All") // EmployeeType | The user type. (optional)
	employeeTypes := []int32{int32(0)} // []int32 | The list of user types. (optional)
	isAdministrator := false // bool | Specifies if the user is an administrator or not. (optional)
	payments := openapiclient.Payments(0) // Payments | The user payment status. (optional)
	accountLoginType := openapiclient.AccountLoginType(0) // AccountLoginType | The account login type. (optional)
	quotaFilter := openapiclient.QuotaFilter(0) // QuotaFilter | The quota filter (All - 0, Default - 1, Custom - 2). (optional)
	withoutGroup := false // bool | Specifies whether the user should be a member of a group or not. (optional)
	excludeGroup := false // bool | Specifies whether the user should be a member of the group with the specified ID. (optional)
	invitedByMe := false // bool | Specifies whether the user is invited by the current user or not. (optional)
	inviterId := "00000000-0000-0000-0000-000000000000" // string | The inviter ID. (optional)
	area := openapiclient.Area(0) // Area | The filter area. (optional)
	count := int32(25) // int32 | The maximum number of items to be retrieved in the response. (optional)
	startIndex := int32(0) // int32 | The zero-based index of the first item to be retrieved in a filtered result set. (optional)
	sortBy := "displayName" // string | Specifies the property or field name by which the results should be sorted. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The order in which the results are sorted. (optional)
	filterSeparator := "," // string | Represents the separator used to split filter criteria in query parameters. (optional)
	filterValue := "John" // string | The search text used to filter results based on user input. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleSearchAPI.SearchUsersByExtendedFilter(context.Background()).EmployeeStatus(employeeStatus).GroupId(groupId).ActivationStatus(activationStatus).EmployeeType(employeeType).EmployeeTypes(employeeTypes).IsAdministrator(isAdministrator).Payments(payments).AccountLoginType(accountLoginType).QuotaFilter(quotaFilter).WithoutGroup(withoutGroup).ExcludeGroup(excludeGroup).InvitedByMe(invitedByMe).InviterId(inviterId).Area(area).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterSeparator(filterSeparator).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleSearchAPI.SearchUsersByExtendedFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUsersByExtendedFilter`: EmployeeFullArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleSearchAPI.SearchUsersByExtendedFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUsersByExtendedFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **employeeStatus** | [**EmployeeStatus**](EmployeeStatus.md) | The user status. | 
 **groupId** | **string** | The group ID. | 
 **activationStatus** | [**EmployeeActivationStatus**](EmployeeActivationStatus.md) | The user activation status. | 
 **employeeType** | [**EmployeeType**](EmployeeType.md) | The user type. | 
 **employeeTypes** | **[]int32** | The list of user types. | 
 **isAdministrator** | **bool** | Specifies if the user is an administrator or not. | 
 **payments** | [**Payments**](Payments.md) | The user payment status. | 
 **accountLoginType** | [**AccountLoginType**](AccountLoginType.md) | The account login type. | 
 **quotaFilter** | [**QuotaFilter**](QuotaFilter.md) | The quota filter (All - 0, Default - 1, Custom - 2). | 
 **withoutGroup** | **bool** | Specifies whether the user should be a member of a group or not. | 
 **excludeGroup** | **bool** | Specifies whether the user should be a member of the group with the specified ID. | 
 **invitedByMe** | **bool** | Specifies whether the user is invited by the current user or not. | 
 **inviterId** | **string** | The inviter ID. | 
 **area** | [**Area**](Area.md) | The filter area. | 
 **count** | **int32** | The maximum number of items to be retrieved in the response. | 
 **startIndex** | **int32** | The zero-based index of the first item to be retrieved in a filtered result set. | 
 **sortBy** | **string** | Specifies the property or field name by which the results should be sorted. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The order in which the results are sorted. | 
 **filterSeparator** | **string** | Represents the separator used to split filter criteria in query parameters. | 
 **filterValue** | **string** | The search text used to filter results based on user input. | 

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


## SearchUsersByQuery

> EmployeeArrayWrapper SearchUsersByQuery(ctx).Query(query).Execute()

Search users (using query parameters)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/search-users-by-query/).

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
	query := "John" // string | The search query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleSearchAPI.SearchUsersByQuery(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleSearchAPI.SearchUsersByQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUsersByQuery`: EmployeeArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleSearchAPI.SearchUsersByQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUsersByQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The search query. | 

### Return type

[**EmployeeArrayWrapper**](EmployeeArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUsersByStatus

> EmployeeFullArrayWrapper SearchUsersByStatus(ctx, status).Query(query).FilterBy(filterBy).FilterValue(filterValue).Execute()

Search users by status filter



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/search-users-by-status/).

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
	query := "John" // string | The advanced search query. (optional)
	filterBy := "displayName" // string | Specifies the criteria used to filter search results in advanced queries. (optional)
	filterValue := "John" // string | The value used to filter the search query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleSearchAPI.SearchUsersByStatus(context.Background(), status).Query(query).FilterBy(filterBy).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleSearchAPI.SearchUsersByStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUsersByStatus`: EmployeeFullArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleSearchAPI.SearchUsersByStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**status** | [**EmployeeStatus**](.md) | The user status. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchUsersByStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **string** | The advanced search query. | 
 **filterBy** | **string** | Specifies the criteria used to filter search results in advanced queries. | 
 **filterValue** | **string** | The value used to filter the search query. | 

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

