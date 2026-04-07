# \GroupSearchAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGroupsWithFilesShared**](GroupSearchAPI.md#GetGroupsWithFilesShared) | **Get** /api/2.0/group/file/{id} | Get groups with file sharing settings
[**GetGroupsWithFoldersShared**](GroupSearchAPI.md#GetGroupsWithFoldersShared) | **Get** /api/2.0/group/folder/{id} | Get groups with folder sharing settings
[**GetGroupsWithRoomsShared**](GroupSearchAPI.md#GetGroupsWithRoomsShared) | **Get** /api/2.0/group/room/{id} | Get groups with room sharing settings



## GetGroupsWithFilesShared

> GroupArrayWrapper GetGroupsWithFilesShared(ctx, id).ExcludeShared(excludeShared).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()

Get groups with file sharing settings



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
	id := int32(56) // int32 | The group ID.
	excludeShared := false // bool | Specifies whether to exclude the group sharing settings from the response. (optional)
	count := int32(25) // int32 | The number of groups to retrieve in the request. (optional)
	startIndex := int32(0) // int32 | The starting index from which to begin retrieving groups with their sharing settings. (optional)
	filterValue := "John" // string | The text used as a filter for retrieving groups with their sharing settings. (optional)

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
**id** | **int32** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsWithFilesSharedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeShared** | **bool** | Specifies whether to exclude the group sharing settings from the response. | 
 **count** | **int32** | The number of groups to retrieve in the request. | 
 **startIndex** | **int32** | The starting index from which to begin retrieving groups with their sharing settings. | 
 **filterValue** | **string** | The text used as a filter for retrieving groups with their sharing settings. | 

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

Get groups with folder sharing settings



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
	id := int32(56) // int32 | The group ID.
	excludeShared := false // bool | Specifies whether to exclude the group sharing settings from the response. (optional)
	count := int32(25) // int32 | The number of groups to retrieve in the request. (optional)
	startIndex := int32(0) // int32 | The starting index from which to begin retrieving groups with their sharing settings. (optional)
	filterValue := "John" // string | The text used as a filter for retrieving groups with their sharing settings. (optional)

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
**id** | **int32** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsWithFoldersSharedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeShared** | **bool** | Specifies whether to exclude the group sharing settings from the response. | 
 **count** | **int32** | The number of groups to retrieve in the request. | 
 **startIndex** | **int32** | The starting index from which to begin retrieving groups with their sharing settings. | 
 **filterValue** | **string** | The text used as a filter for retrieving groups with their sharing settings. | 

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

Get groups with room sharing settings



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
	id := int32(56) // int32 | The group ID.
	excludeShared := false // bool | Specifies whether to exclude the group sharing settings from the response. (optional)
	count := int32(25) // int32 | The number of groups to retrieve in the request. (optional)
	startIndex := int32(0) // int32 | The starting index from which to begin retrieving groups with their sharing settings. (optional)
	filterValue := "John" // string | The text used as a filter for retrieving groups with their sharing settings. (optional)

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
**id** | **int32** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsWithRoomsSharedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeShared** | **bool** | Specifies whether to exclude the group sharing settings from the response. | 
 **count** | **int32** | The number of groups to retrieve in the request. | 
 **startIndex** | **int32** | The starting index from which to begin retrieving groups with their sharing settings. | 
 **filterValue** | **string** | The text used as a filter for retrieving groups with their sharing settings. | 

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

