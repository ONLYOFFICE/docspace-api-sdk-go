# \GroupSearchAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGroupsWithFilesShared**](GroupSearchAPI.md#GetGroupsWithFilesShared) | **Get** /api/2.0/group/file/{id} | Search groups for a file
[**GetGroupsWithFilesSharedThirdParty**](GroupSearchAPI.md#GetGroupsWithFilesSharedThirdParty) | **Get** /api/2.0/group/file/{id} | Search groups for a file (third-party storage)
[**GetGroupsWithFoldersShared**](GroupSearchAPI.md#GetGroupsWithFoldersShared) | **Get** /api/2.0/group/folder/{id} | Search groups for a folder
[**GetGroupsWithFoldersSharedThirdParty**](GroupSearchAPI.md#GetGroupsWithFoldersSharedThirdParty) | **Get** /api/2.0/group/folder/{id} | Search groups for a folder (third-party storage)
[**GetGroupsWithRoomsShared**](GroupSearchAPI.md#GetGroupsWithRoomsShared) | **Get** /api/2.0/group/room/{id} | Search groups for a room
[**GetGroupsWithRoomsSharedThirdParty**](GroupSearchAPI.md#GetGroupsWithRoomsSharedThirdParty) | **Get** /api/2.0/group/room/{id} | Search groups for a room (third-party storage)



## GetGroupsWithFilesShared

> GroupArrayWrapper GetGroupsWithFilesShared(ctx, id).ExcludeShared(excludeShared).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()

Search groups for a file



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-groups-with-files-shared/).

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
	id := int32(1234) // int32 | The ID of the room, folder or file whose access the search is run against, taken from the route. It is an  integer for an entry stored in DocSpace and a provider-specific string for an entry in a connected  third-party storage.
	excludeShared := false // bool | Keeps only the groups that do not have access to the entry yet, which is the set to offer when granting  access. Every returned entry then has `shared` set to false; without the flag every matching group comes back  and `shared` tells them apart. (optional)
	count := int32(25) // int32 | The size of the page. It defaults to 100, which is also the largest value the operation accepts. (optional)
	startIndex := int32(0) // int32 | The number of matching groups to skip before the page starts. It defaults to 0, and the total number of  matches is reported in the total count of the response. (optional)
	filterValue := "Marketing" // string | The text to match against the group name. Omit it to get every group the caller may grant access to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupSearchAPI.GetGroupsWithFilesShared(context.Background(), id).ExcludeShared(excludeShared).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupSearchAPI.GetGroupsWithFilesShared``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupsWithFilesShared`: GroupArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `GroupSearchAPI.GetGroupsWithFilesShared`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the room, folder or file whose access the search is run against, taken from the route. It is an  integer for an entry stored in DocSpace and a provider-specific string for an entry in a connected  third-party storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsWithFilesSharedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeShared** | **bool** | Keeps only the groups that do not have access to the entry yet, which is the set to offer when granting  access. Every returned entry then has `shared` set to false; without the flag every matching group comes back  and `shared` tells them apart. | 
 **count** | **int32** | The size of the page. It defaults to 100, which is also the largest value the operation accepts. | 
 **startIndex** | **int32** | The number of matching groups to skip before the page starts. It defaults to 0, and the total number of  matches is reported in the total count of the response. | 
 **filterValue** | **string** | The text to match against the group name. Omit it to get every group the caller may grant access to. | 

### Return type

[**GroupArrayWrapper**](GroupArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupsWithFilesSharedThirdParty

> GroupArrayWrapper GetGroupsWithFilesSharedThirdParty(ctx, id).ExcludeShared(excludeShared).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()

Search groups for a file (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-groups-with-files-shared-third-party/).

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
	id := "1234" // string | The ID of the room, folder or file whose access the search is run against, taken from the route. It is an  integer for an entry stored in DocSpace and a provider-specific string for an entry in a connected  third-party storage.
	excludeShared := false // bool | Keeps only the groups that do not have access to the entry yet, which is the set to offer when granting  access. Every returned entry then has `shared` set to false; without the flag every matching group comes back  and `shared` tells them apart. (optional)
	count := int32(25) // int32 | The size of the page. It defaults to 100, which is also the largest value the operation accepts. (optional)
	startIndex := int32(0) // int32 | The number of matching groups to skip before the page starts. It defaults to 0, and the total number of  matches is reported in the total count of the response. (optional)
	filterValue := "Marketing" // string | The text to match against the group name. Omit it to get every group the caller may grant access to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupSearchAPI.GetGroupsWithFilesSharedThirdParty(context.Background(), id).ExcludeShared(excludeShared).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupSearchAPI.GetGroupsWithFilesSharedThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupsWithFilesSharedThirdParty`: GroupArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `GroupSearchAPI.GetGroupsWithFilesSharedThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the room, folder or file whose access the search is run against, taken from the route. It is an  integer for an entry stored in DocSpace and a provider-specific string for an entry in a connected  third-party storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsWithFilesSharedThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeShared** | **bool** | Keeps only the groups that do not have access to the entry yet, which is the set to offer when granting  access. Every returned entry then has `shared` set to false; without the flag every matching group comes back  and `shared` tells them apart. | 
 **count** | **int32** | The size of the page. It defaults to 100, which is also the largest value the operation accepts. | 
 **startIndex** | **int32** | The number of matching groups to skip before the page starts. It defaults to 0, and the total number of  matches is reported in the total count of the response. | 
 **filterValue** | **string** | The text to match against the group name. Omit it to get every group the caller may grant access to. | 

### Return type

[**GroupArrayWrapper**](GroupArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupsWithFoldersShared

> GroupArrayWrapper GetGroupsWithFoldersShared(ctx, id).ExcludeShared(excludeShared).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()

Search groups for a folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-groups-with-folders-shared/).

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
	id := int32(1234) // int32 | The ID of the room, folder or file whose access the search is run against, taken from the route. It is an  integer for an entry stored in DocSpace and a provider-specific string for an entry in a connected  third-party storage.
	excludeShared := false // bool | Keeps only the groups that do not have access to the entry yet, which is the set to offer when granting  access. Every returned entry then has `shared` set to false; without the flag every matching group comes back  and `shared` tells them apart. (optional)
	count := int32(25) // int32 | The size of the page. It defaults to 100, which is also the largest value the operation accepts. (optional)
	startIndex := int32(0) // int32 | The number of matching groups to skip before the page starts. It defaults to 0, and the total number of  matches is reported in the total count of the response. (optional)
	filterValue := "Marketing" // string | The text to match against the group name. Omit it to get every group the caller may grant access to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupSearchAPI.GetGroupsWithFoldersShared(context.Background(), id).ExcludeShared(excludeShared).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupSearchAPI.GetGroupsWithFoldersShared``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupsWithFoldersShared`: GroupArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `GroupSearchAPI.GetGroupsWithFoldersShared`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the room, folder or file whose access the search is run against, taken from the route. It is an  integer for an entry stored in DocSpace and a provider-specific string for an entry in a connected  third-party storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsWithFoldersSharedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeShared** | **bool** | Keeps only the groups that do not have access to the entry yet, which is the set to offer when granting  access. Every returned entry then has `shared` set to false; without the flag every matching group comes back  and `shared` tells them apart. | 
 **count** | **int32** | The size of the page. It defaults to 100, which is also the largest value the operation accepts. | 
 **startIndex** | **int32** | The number of matching groups to skip before the page starts. It defaults to 0, and the total number of  matches is reported in the total count of the response. | 
 **filterValue** | **string** | The text to match against the group name. Omit it to get every group the caller may grant access to. | 

### Return type

[**GroupArrayWrapper**](GroupArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupsWithFoldersSharedThirdParty

> GroupArrayWrapper GetGroupsWithFoldersSharedThirdParty(ctx, id).ExcludeShared(excludeShared).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()

Search groups for a folder (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-groups-with-folders-shared-third-party/).

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
	id := "1234" // string | The ID of the room, folder or file whose access the search is run against, taken from the route. It is an  integer for an entry stored in DocSpace and a provider-specific string for an entry in a connected  third-party storage.
	excludeShared := false // bool | Keeps only the groups that do not have access to the entry yet, which is the set to offer when granting  access. Every returned entry then has `shared` set to false; without the flag every matching group comes back  and `shared` tells them apart. (optional)
	count := int32(25) // int32 | The size of the page. It defaults to 100, which is also the largest value the operation accepts. (optional)
	startIndex := int32(0) // int32 | The number of matching groups to skip before the page starts. It defaults to 0, and the total number of  matches is reported in the total count of the response. (optional)
	filterValue := "Marketing" // string | The text to match against the group name. Omit it to get every group the caller may grant access to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupSearchAPI.GetGroupsWithFoldersSharedThirdParty(context.Background(), id).ExcludeShared(excludeShared).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupSearchAPI.GetGroupsWithFoldersSharedThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupsWithFoldersSharedThirdParty`: GroupArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `GroupSearchAPI.GetGroupsWithFoldersSharedThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the room, folder or file whose access the search is run against, taken from the route. It is an  integer for an entry stored in DocSpace and a provider-specific string for an entry in a connected  third-party storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsWithFoldersSharedThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeShared** | **bool** | Keeps only the groups that do not have access to the entry yet, which is the set to offer when granting  access. Every returned entry then has `shared` set to false; without the flag every matching group comes back  and `shared` tells them apart. | 
 **count** | **int32** | The size of the page. It defaults to 100, which is also the largest value the operation accepts. | 
 **startIndex** | **int32** | The number of matching groups to skip before the page starts. It defaults to 0, and the total number of  matches is reported in the total count of the response. | 
 **filterValue** | **string** | The text to match against the group name. Omit it to get every group the caller may grant access to. | 

### Return type

[**GroupArrayWrapper**](GroupArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupsWithRoomsShared

> GroupArrayWrapper GetGroupsWithRoomsShared(ctx, id).ExcludeShared(excludeShared).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()

Search groups for a room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-groups-with-rooms-shared/).

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
	id := int32(1234) // int32 | The ID of the room, folder or file whose access the search is run against, taken from the route. It is an  integer for an entry stored in DocSpace and a provider-specific string for an entry in a connected  third-party storage.
	excludeShared := false // bool | Keeps only the groups that do not have access to the entry yet, which is the set to offer when granting  access. Every returned entry then has `shared` set to false; without the flag every matching group comes back  and `shared` tells them apart. (optional)
	count := int32(25) // int32 | The size of the page. It defaults to 100, which is also the largest value the operation accepts. (optional)
	startIndex := int32(0) // int32 | The number of matching groups to skip before the page starts. It defaults to 0, and the total number of  matches is reported in the total count of the response. (optional)
	filterValue := "Marketing" // string | The text to match against the group name. Omit it to get every group the caller may grant access to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupSearchAPI.GetGroupsWithRoomsShared(context.Background(), id).ExcludeShared(excludeShared).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupSearchAPI.GetGroupsWithRoomsShared``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupsWithRoomsShared`: GroupArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `GroupSearchAPI.GetGroupsWithRoomsShared`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the room, folder or file whose access the search is run against, taken from the route. It is an  integer for an entry stored in DocSpace and a provider-specific string for an entry in a connected  third-party storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsWithRoomsSharedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeShared** | **bool** | Keeps only the groups that do not have access to the entry yet, which is the set to offer when granting  access. Every returned entry then has `shared` set to false; without the flag every matching group comes back  and `shared` tells them apart. | 
 **count** | **int32** | The size of the page. It defaults to 100, which is also the largest value the operation accepts. | 
 **startIndex** | **int32** | The number of matching groups to skip before the page starts. It defaults to 0, and the total number of  matches is reported in the total count of the response. | 
 **filterValue** | **string** | The text to match against the group name. Omit it to get every group the caller may grant access to. | 

### Return type

[**GroupArrayWrapper**](GroupArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupsWithRoomsSharedThirdParty

> GroupArrayWrapper GetGroupsWithRoomsSharedThirdParty(ctx, id).ExcludeShared(excludeShared).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()

Search groups for a room (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-groups-with-rooms-shared-third-party/).

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
	id := "1234" // string | The ID of the room, folder or file whose access the search is run against, taken from the route. It is an  integer for an entry stored in DocSpace and a provider-specific string for an entry in a connected  third-party storage.
	excludeShared := false // bool | Keeps only the groups that do not have access to the entry yet, which is the set to offer when granting  access. Every returned entry then has `shared` set to false; without the flag every matching group comes back  and `shared` tells them apart. (optional)
	count := int32(25) // int32 | The size of the page. It defaults to 100, which is also the largest value the operation accepts. (optional)
	startIndex := int32(0) // int32 | The number of matching groups to skip before the page starts. It defaults to 0, and the total number of  matches is reported in the total count of the response. (optional)
	filterValue := "Marketing" // string | The text to match against the group name. Omit it to get every group the caller may grant access to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupSearchAPI.GetGroupsWithRoomsSharedThirdParty(context.Background(), id).ExcludeShared(excludeShared).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupSearchAPI.GetGroupsWithRoomsSharedThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupsWithRoomsSharedThirdParty`: GroupArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `GroupSearchAPI.GetGroupsWithRoomsSharedThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the room, folder or file whose access the search is run against, taken from the route. It is an  integer for an entry stored in DocSpace and a provider-specific string for an entry in a connected  third-party storage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsWithRoomsSharedThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeShared** | **bool** | Keeps only the groups that do not have access to the entry yet, which is the set to offer when granting  access. Every returned entry then has `shared` set to false; without the flag every matching group comes back  and `shared` tells them apart. | 
 **count** | **int32** | The size of the page. It defaults to 100, which is also the largest value the operation accepts. | 
 **startIndex** | **int32** | The number of matching groups to skip before the page starts. It defaults to 0, and the total number of  matches is reported in the total count of the response. | 
 **filterValue** | **string** | The text to match against the group name. Omit it to get every group the caller may grant access to. | 

### Return type

[**GroupArrayWrapper**](GroupArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

