# \FilesQuotaAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResetRoomQuota**](FilesQuotaAPI.md#ResetRoomQuota) | **Put** /api/2.0/files/rooms/resetquota | Reset the room quota limit
[**UpdateRoomsQuota**](FilesQuotaAPI.md#UpdateRoomsQuota) | **Put** /api/2.0/files/rooms/roomquota | Change the room quota limit



## ResetRoomQuota

> FolderIntegerArrayWrapper ResetRoomQuota(ctx).UpdateRoomsRoomIdsRequestDtoInteger(updateRoomsRoomIdsRequestDtoInteger).Execute()

Reset the room quota limit



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/reset-room-quota/).

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
	updateRoomsRoomIdsRequestDtoInteger := *openapiclient.NewUpdateRoomsRoomIdsRequestDtoInteger() // UpdateRoomsRoomIdsRequestDtoInteger |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesQuotaAPI.ResetRoomQuota(context.Background()).UpdateRoomsRoomIdsRequestDtoInteger(updateRoomsRoomIdsRequestDtoInteger).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesQuotaAPI.ResetRoomQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetRoomQuota`: FolderIntegerArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesQuotaAPI.ResetRoomQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetRoomQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRoomsRoomIdsRequestDtoInteger** | [**UpdateRoomsRoomIdsRequestDtoInteger**](UpdateRoomsRoomIdsRequestDtoInteger.md) |  | 

### Return type

[**FolderIntegerArrayWrapper**](FolderIntegerArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoomsQuota

> FolderIntegerArrayWrapper UpdateRoomsQuota(ctx).UpdateRoomsQuotaRequestDtoInteger(updateRoomsQuotaRequestDtoInteger).Execute()

Change the room quota limit



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-rooms-quota/).

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
	updateRoomsQuotaRequestDtoInteger := *openapiclient.NewUpdateRoomsQuotaRequestDtoInteger() // UpdateRoomsQuotaRequestDtoInteger |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesQuotaAPI.UpdateRoomsQuota(context.Background()).UpdateRoomsQuotaRequestDtoInteger(updateRoomsQuotaRequestDtoInteger).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesQuotaAPI.UpdateRoomsQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoomsQuota`: FolderIntegerArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesQuotaAPI.UpdateRoomsQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoomsQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRoomsQuotaRequestDtoInteger** | [**UpdateRoomsQuotaRequestDtoInteger**](UpdateRoomsQuotaRequestDtoInteger.md) |  | 

### Return type

[**FolderIntegerArrayWrapper**](FolderIntegerArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

