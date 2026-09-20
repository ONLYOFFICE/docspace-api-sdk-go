# \RoomsGroupsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRoomGroup**](RoomsGroupsAPI.md#AddRoomGroup) | **Post** /api/2.0/files/group | Add a new room group
[**ChangeRoomGroupIcon**](RoomsGroupsAPI.md#ChangeRoomGroupIcon) | **Post** /api/2.0/files/group/{id}/icon | Change room group icon
[**DeleteRoomGroup**](RoomsGroupsAPI.md#DeleteRoomGroup) | **Delete** /api/2.0/files/group/{id} | Delete a room group
[**GetRoomGroupInfo**](RoomsGroupsAPI.md#GetRoomGroupInfo) | **Get** /api/2.0/files/group/{id} | Get room group info
[**GetRoomGroups**](RoomsGroupsAPI.md#GetRoomGroups) | **Get** /api/2.0/files/group | List room groups
[**UpdateRoomGroup**](RoomsGroupsAPI.md#UpdateRoomGroup) | **Put** /api/2.0/files/group/{id} | Update room group



## AddRoomGroup

> RoomGroupWrapper AddRoomGroup(ctx).RoomGroupRequestDto(roomGroupRequestDto).Execute()

Add a new room group



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/add-room-group/).

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
	roomGroupRequestDto := *openapiclient.NewRoomGroupRequestDto("Client projects", "star", []openapiclient.DuplicateRequestDtoAllOfFileIds{*openapiclient.NewDuplicateRequestDtoAllOfFileIds()}) // RoomGroupRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsGroupsAPI.AddRoomGroup(context.Background()).RoomGroupRequestDto(roomGroupRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsGroupsAPI.AddRoomGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddRoomGroup`: RoomGroupWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsGroupsAPI.AddRoomGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRoomGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roomGroupRequestDto** | [**RoomGroupRequestDto**](RoomGroupRequestDto.md) |  | 

### Return type

[**RoomGroupWrapper**](RoomGroupWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeRoomGroupIcon

> RoomGroupWrapper ChangeRoomGroupIcon(ctx, id).IconRequest(iconRequest).Execute()

Change room group icon



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/change-room-group-icon/).

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
	id := int32(42) // int32 | The room group to re-icon, identified by the value `GET api/2.0/files/group` reports for it. A group of  another account cannot be addressed and reads as missing.
	iconRequest := *openapiclient.NewIconRequest() // IconRequest | The icon to give the group. A body that leaves the icon out is accepted and changes nothing. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsGroupsAPI.ChangeRoomGroupIcon(context.Background(), id).IconRequest(iconRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsGroupsAPI.ChangeRoomGroupIcon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeRoomGroupIcon`: RoomGroupWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsGroupsAPI.ChangeRoomGroupIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room group to re-icon, identified by the value `GET api/2.0/files/group` reports for it. A group of  another account cannot be addressed and reads as missing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeRoomGroupIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iconRequest** | [**IconRequest**](IconRequest.md) | The icon to give the group. A body that leaves the icon out is accepted and changes nothing. | 

### Return type

[**RoomGroupWrapper**](RoomGroupWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoomGroup

> DeleteRoomGroup(ctx, id).IncludeMembers(includeMembers).Execute()

Delete a room group



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-room-group/).

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
	id := int32(42) // int32 | The room group to act on, identified by the value `GET api/2.0/files/group` reports for it. A group of another  account cannot be addressed and reads as missing.
	includeMembers := true // bool | Whether the rooms of the group are listed in the answer: true fills the `rooms` array, false leaves it out and  reports only how many there are in `totalRooms`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoomsGroupsAPI.DeleteRoomGroup(context.Background(), id).IncludeMembers(includeMembers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsGroupsAPI.DeleteRoomGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room group to act on, identified by the value `GET api/2.0/files/group` reports for it. A group of another  account cannot be addressed and reads as missing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoomGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeMembers** | **bool** | Whether the rooms of the group are listed in the answer: true fills the `rooms` array, false leaves it out and  reports only how many there are in `totalRooms`. | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomGroupInfo

> RoomGroupWrapper GetRoomGroupInfo(ctx, id).IncludeMembers(includeMembers).Execute()

Get room group info



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-room-group-info/).

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
	id := int32(42) // int32 | The room group to act on, identified by the value `GET api/2.0/files/group` reports for it. A group of another  account cannot be addressed and reads as missing.
	includeMembers := true // bool | Whether the rooms of the group are listed in the answer: true fills the `rooms` array, false leaves it out and  reports only how many there are in `totalRooms`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsGroupsAPI.GetRoomGroupInfo(context.Background(), id).IncludeMembers(includeMembers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsGroupsAPI.GetRoomGroupInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomGroupInfo`: RoomGroupWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsGroupsAPI.GetRoomGroupInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room group to act on, identified by the value `GET api/2.0/files/group` reports for it. A group of another  account cannot be addressed and reads as missing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomGroupInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeMembers** | **bool** | Whether the rooms of the group are listed in the answer: true fills the `rooms` array, false leaves it out and  reports only how many there are in `totalRooms`. | 

### Return type

[**RoomGroupWrapper**](RoomGroupWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomGroups

> RoomGroupArrayWrapper GetRoomGroups(ctx).IncludeMembers(includeMembers).Execute()

List room groups



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-room-groups/).

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
	includeMembers := true // bool | Whether the rooms of each group are listed in the answer: true fills the `rooms` array of every entry, false  leaves it out and reports only how many there are in `totalRooms`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsGroupsAPI.GetRoomGroups(context.Background()).IncludeMembers(includeMembers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsGroupsAPI.GetRoomGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomGroups`: RoomGroupArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsGroupsAPI.GetRoomGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeMembers** | **bool** | Whether the rooms of each group are listed in the answer: true fills the `rooms` array of every entry, false  leaves it out and reports only how many there are in `totalRooms`. | 

### Return type

[**RoomGroupArrayWrapper**](RoomGroupArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoomGroup

> RoomGroupWrapper UpdateRoomGroup(ctx, id).UpdateRoomGroupRequest(updateRoomGroupRequest).Execute()

Update room group



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-room-group/).

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
	id := int32(42) // int32 | The room group to change, identified by the value `GET api/2.0/files/group` reports for it. A group of another  account cannot be addressed and reads as missing.
	updateRoomGroupRequest := *openapiclient.NewUpdateRoomGroupRequest() // UpdateRoomGroupRequest | The changes to apply. Carrying none of them leaves the group as it is, and each of them may be sent on its own  or together with the others.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsGroupsAPI.UpdateRoomGroup(context.Background(), id).UpdateRoomGroupRequest(updateRoomGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsGroupsAPI.UpdateRoomGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoomGroup`: RoomGroupWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsGroupsAPI.UpdateRoomGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room group to change, identified by the value `GET api/2.0/files/group` reports for it. A group of another  account cannot be addressed and reads as missing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoomGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRoomGroupRequest** | [**UpdateRoomGroupRequest**](UpdateRoomGroupRequest.md) | The changes to apply. Carrying none of them leaves the group as it is, and each of them may be sent on its own  or together with the others. | 

### Return type

[**RoomGroupWrapper**](RoomGroupWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

