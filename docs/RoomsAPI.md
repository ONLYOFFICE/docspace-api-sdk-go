# \RoomsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRoomTags**](RoomsAPI.md#AddRoomTags) | **Put** /api/2.0/files/rooms/{id}/tags | Add the room tags
[**ArchiveRoom**](RoomsAPI.md#ArchiveRoom) | **Put** /api/2.0/files/rooms/{id}/archive | Archive a room
[**ChangeRoomCover**](RoomsAPI.md#ChangeRoomCover) | **Post** /api/2.0/files/rooms/{id}/cover | Change the room cover
[**CreateRoom**](RoomsAPI.md#CreateRoom) | **Post** /api/2.0/files/rooms | Create a room
[**CreateRoomFromTemplate**](RoomsAPI.md#CreateRoomFromTemplate) | **Post** /api/2.0/files/rooms/fromtemplate | Create a room from the template
[**CreateRoomLogo**](RoomsAPI.md#CreateRoomLogo) | **Post** /api/2.0/files/rooms/{id}/logo | Create a room logo
[**CreateRoomTag**](RoomsAPI.md#CreateRoomTag) | **Post** /api/2.0/files/tags | Create a room tag
[**CreateRoomTemplate**](RoomsAPI.md#CreateRoomTemplate) | **Post** /api/2.0/files/roomtemplate | Start creating room template
[**CreateRoomThirdParty**](RoomsAPI.md#CreateRoomThirdParty) | **Post** /api/2.0/files/rooms/thirdparty/{id} | Create a third-party room
[**DeleteCustomTags**](RoomsAPI.md#DeleteCustomTags) | **Delete** /api/2.0/files/tags | Delete the custom room tags
[**DeleteRoom**](RoomsAPI.md#DeleteRoom) | **Delete** /api/2.0/files/rooms/{id} | Remove a room
[**DeleteRoomLogo**](RoomsAPI.md#DeleteRoomLogo) | **Delete** /api/2.0/files/rooms/{id}/logo | Remove a room logo
[**DeleteRoomTags**](RoomsAPI.md#DeleteRoomTags) | **Delete** /api/2.0/files/rooms/{id}/tags | Remove the room tags
[**GetExternalDbSyncStatus**](RoomsAPI.md#GetExternalDbSyncStatus) | **Get** /api/2.0/files/rooms/{id}/externaldbsync | Get external DB sync status
[**GetNewRoomItems**](RoomsAPI.md#GetNewRoomItems) | **Get** /api/2.0/files/rooms/{id}/news | Get the new room items
[**GetPublicSettings**](RoomsAPI.md#GetPublicSettings) | **Get** /api/2.0/files/roomtemplate/{id}/public | Get public settings
[**GetRoomCovers**](RoomsAPI.md#GetRoomCovers) | **Get** /api/2.0/files/rooms/covers | Get covers
[**GetRoomCreatingStatus**](RoomsAPI.md#GetRoomCreatingStatus) | **Get** /api/2.0/files/rooms/fromtemplate/status | Get the room creation progress
[**GetRoomIndexExport**](RoomsAPI.md#GetRoomIndexExport) | **Get** /api/2.0/files/rooms/indexexport | Get the room index export
[**GetRoomInfo**](RoomsAPI.md#GetRoomInfo) | **Get** /api/2.0/files/rooms/{id} | Get room information
[**GetRoomLinks**](RoomsAPI.md#GetRoomLinks) | **Get** /api/2.0/files/rooms/{id}/links | Get the room links
[**GetRoomSecurityInfo**](RoomsAPI.md#GetRoomSecurityInfo) | **Get** /api/2.0/files/rooms/{id}/share | Get the room access rights
[**GetRoomTagsInfo**](RoomsAPI.md#GetRoomTagsInfo) | **Get** /api/2.0/files/tags | Get the room tags
[**GetRoomTemplateCreatingStatus**](RoomsAPI.md#GetRoomTemplateCreatingStatus) | **Get** /api/2.0/files/roomtemplate/status | Get status of room template creation
[**GetRoomsFolder**](RoomsAPI.md#GetRoomsFolder) | **Get** /api/2.0/files/rooms | Get rooms
[**GetRoomsNewItems**](RoomsAPI.md#GetRoomsNewItems) | **Get** /api/2.0/files/rooms/news | Get the room new items
[**GetRoomsPrimaryExternalLink**](RoomsAPI.md#GetRoomsPrimaryExternalLink) | **Get** /api/2.0/files/rooms/{id}/link | Get the room primary external link
[**HasTagLinks**](RoomsAPI.md#HasTagLinks) | **Get** /api/2.0/files/tags/{tagName}/haslinks | Has tag links
[**PinRoom**](RoomsAPI.md#PinRoom) | **Put** /api/2.0/files/rooms/{id}/pin | Pin a room
[**ReorderRoom**](RoomsAPI.md#ReorderRoom) | **Put** /api/2.0/files/rooms/{id}/reorder | Reorder the room
[**ResendEmailInvitations**](RoomsAPI.md#ResendEmailInvitations) | **Post** /api/2.0/files/rooms/{id}/resend | Resend the room invitations
[**SetPublicSettings**](RoomsAPI.md#SetPublicSettings) | **Put** /api/2.0/files/roomtemplate/public | Set public settings
[**SetRoomLink**](RoomsAPI.md#SetRoomLink) | **Put** /api/2.0/files/rooms/{id}/links | Set the room external or invitation link
[**SetRoomSecurity**](RoomsAPI.md#SetRoomSecurity) | **Put** /api/2.0/files/rooms/{id}/share | Set the room access rights
[**StartExternalDbSync**](RoomsAPI.md#StartExternalDbSync) | **Post** /api/2.0/files/rooms/{id}/externaldbsync | Start external DB sync
[**StartRoomIndexExport**](RoomsAPI.md#StartRoomIndexExport) | **Post** /api/2.0/files/rooms/{id}/indexexport | Start the room index export
[**TerminateRoomIndexExport**](RoomsAPI.md#TerminateRoomIndexExport) | **Delete** /api/2.0/files/rooms/indexexport | Terminate the room index export
[**UnarchiveRoom**](RoomsAPI.md#UnarchiveRoom) | **Put** /api/2.0/files/rooms/{id}/unarchive | Unarchive a room
[**UnpinRoom**](RoomsAPI.md#UnpinRoom) | **Put** /api/2.0/files/rooms/{id}/unpin | Unpin a room
[**UpdateRoom**](RoomsAPI.md#UpdateRoom) | **Put** /api/2.0/files/rooms/{id} | Update a room
[**UpdateRoomTag**](RoomsAPI.md#UpdateRoomTag) | **Put** /api/2.0/files/tags | Update tag
[**UploadRoomLogo**](RoomsAPI.md#UploadRoomLogo) | **Post** /api/2.0/files/logos | Upload a room logo image



## AddRoomTags

> FolderIntegerWrapper AddRoomTags(ctx, id).BatchTagsRequestDto(batchTagsRequestDto).Execute()

Add the room tags



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/add-room-tags/).

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
	id := int32(1) // int32 | The room Id.
	batchTagsRequestDto := *openapiclient.NewBatchTagsRequestDto([]string{"Names_example"}) // BatchTagsRequestDto | The parameters for managing tags. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.AddRoomTags(context.Background(), id).BatchTagsRequestDto(batchTagsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.AddRoomTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddRoomTags`: FolderIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.AddRoomTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRoomTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchTagsRequestDto** | [**BatchTagsRequestDto**](BatchTagsRequestDto.md) | The parameters for managing tags. | 

### Return type

[**FolderIntegerWrapper**](FolderIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchiveRoom

> FileOperationWrapper ArchiveRoom(ctx, id).ArchiveRoomRequest(archiveRoomRequest).Execute()

Archive a room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/archive-room/).

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
	id := int32(1) // int32 | The room ID.
	archiveRoomRequest := *openapiclient.NewArchiveRoomRequest() // ArchiveRoomRequest | The parameters for archiving a room. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.ArchiveRoom(context.Background(), id).ArchiveRoomRequest(archiveRoomRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.ArchiveRoom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchiveRoom`: FileOperationWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.ArchiveRoom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archiveRoomRequest** | [**ArchiveRoomRequest**](ArchiveRoomRequest.md) | The parameters for archiving a room. | 

### Return type

[**FileOperationWrapper**](FileOperationWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeRoomCover

> FolderIntegerWrapper ChangeRoomCover(ctx, id).CoverRequestDto(coverRequestDto).Execute()

Change the room cover



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/change-room-cover/).

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
	id := int32(1) // int32 | The room ID.
	coverRequestDto := *openapiclient.NewCoverRequestDto() // CoverRequestDto | The request parameters to change the room cover.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.ChangeRoomCover(context.Background(), id).CoverRequestDto(coverRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.ChangeRoomCover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeRoomCover`: FolderIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.ChangeRoomCover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeRoomCoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **coverRequestDto** | [**CoverRequestDto**](CoverRequestDto.md) | The request parameters to change the room cover. | 

### Return type

[**FolderIntegerWrapper**](FolderIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoom

> FolderIntegerWrapper CreateRoom(ctx).CreateRoomRequestDto(createRoomRequestDto).Execute()

Create a room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-room/).

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
	createRoomRequestDto := *openapiclient.NewCreateRoomRequestDto("My Room", openapiclient.RoomType(1)) // CreateRoomRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.CreateRoom(context.Background()).CreateRoomRequestDto(createRoomRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.CreateRoom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoom`: FolderIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.CreateRoom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRoomRequestDto** | [**CreateRoomRequestDto**](CreateRoomRequestDto.md) |  | 

### Return type

[**FolderIntegerWrapper**](FolderIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoomFromTemplate

> RoomFromTemplateStatusWrapper CreateRoomFromTemplate(ctx).CreateRoomFromTemplateDto(createRoomFromTemplateDto).Execute()

Create a room from the template



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-room-from-template/).

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
	createRoomFromTemplateDto := *openapiclient.NewCreateRoomFromTemplateDto(int32(1), "My Room From Template") // CreateRoomFromTemplateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.CreateRoomFromTemplate(context.Background()).CreateRoomFromTemplateDto(createRoomFromTemplateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.CreateRoomFromTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoomFromTemplate`: RoomFromTemplateStatusWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.CreateRoomFromTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoomFromTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRoomFromTemplateDto** | [**CreateRoomFromTemplateDto**](CreateRoomFromTemplateDto.md) |  | 

### Return type

[**RoomFromTemplateStatusWrapper**](RoomFromTemplateStatusWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoomLogo

> FolderIntegerWrapper CreateRoomLogo(ctx, id).LogoRequest(logoRequest).Execute()

Create a room logo



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-room-logo/).

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
	id := int32(1) // int32 | The room ID.
	logoRequest := *openapiclient.NewLogoRequest("/tmp/logo.png") // LogoRequest | The logo request parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.CreateRoomLogo(context.Background(), id).LogoRequest(logoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.CreateRoomLogo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoomLogo`: FolderIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.CreateRoomLogo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoomLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logoRequest** | [**LogoRequest**](LogoRequest.md) | The logo request parameters. | 

### Return type

[**FolderIntegerWrapper**](FolderIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoomTag

> StringWrapper CreateRoomTag(ctx).CreateTagRequestDto(createTagRequestDto).Execute()

Create a room tag



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-room-tag/).

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
	createTagRequestDto := *openapiclient.NewCreateTagRequestDto("Important") // CreateTagRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.CreateRoomTag(context.Background()).CreateTagRequestDto(createTagRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.CreateRoomTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoomTag`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.CreateRoomTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoomTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTagRequestDto** | [**CreateTagRequestDto**](CreateTagRequestDto.md) |  | 

### Return type

[**StringWrapper**](StringWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoomTemplate

> RoomTemplateStatusWrapper CreateRoomTemplate(ctx).RoomTemplateDto(roomTemplateDto).Execute()

Start creating room template



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-room-template/).

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
	roomTemplateDto := *openapiclient.NewRoomTemplateDto(int32(1)) // RoomTemplateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.CreateRoomTemplate(context.Background()).RoomTemplateDto(roomTemplateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.CreateRoomTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoomTemplate`: RoomTemplateStatusWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.CreateRoomTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoomTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roomTemplateDto** | [**RoomTemplateDto**](RoomTemplateDto.md) |  | 

### Return type

[**RoomTemplateStatusWrapper**](RoomTemplateStatusWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoomThirdParty

> FolderStringWrapper CreateRoomThirdParty(ctx, id).CreateThirdPartyRoom(createThirdPartyRoom).Execute()

Create a third-party room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-room-third-party/).

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
	id := "folder-123-abc" // string | The ID of the folder in the third-party storage in which the contents of the room will be stored.
	createThirdPartyRoom := *openapiclient.NewCreateThirdPartyRoom("My Third-Party Room", openapiclient.RoomType(1)) // CreateThirdPartyRoom | The third-party room information.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.CreateRoomThirdParty(context.Background(), id).CreateThirdPartyRoom(createThirdPartyRoom).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.CreateRoomThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoomThirdParty`: FolderStringWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.CreateRoomThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the folder in the third-party storage in which the contents of the room will be stored. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoomThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createThirdPartyRoom** | [**CreateThirdPartyRoom**](CreateThirdPartyRoom.md) | The third-party room information. | 

### Return type

[**FolderStringWrapper**](FolderStringWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomTags

> DeleteCustomTags(ctx).BatchTagsRequestDto(batchTagsRequestDto).Execute()

Delete the custom room tags



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-custom-tags/).

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
	batchTagsRequestDto := *openapiclient.NewBatchTagsRequestDto([]string{"Names_example"}) // BatchTagsRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoomsAPI.DeleteCustomTags(context.Background()).BatchTagsRequestDto(batchTagsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.DeleteCustomTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchTagsRequestDto** | [**BatchTagsRequestDto**](BatchTagsRequestDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoom

> FileOperationWrapper DeleteRoom(ctx, id).DeleteRoomRequest(deleteRoomRequest).Execute()

Remove a room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-room/).

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
	id := int32(10) // int32 | The room ID.
	deleteRoomRequest := *openapiclient.NewDeleteRoomRequest() // DeleteRoomRequest | The parameters for deleting a room.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.DeleteRoom(context.Background(), id).DeleteRoomRequest(deleteRoomRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.DeleteRoom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRoom`: FileOperationWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.DeleteRoom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteRoomRequest** | [**DeleteRoomRequest**](DeleteRoomRequest.md) | The parameters for deleting a room. | 

### Return type

[**FileOperationWrapper**](FileOperationWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoomLogo

> FolderIntegerWrapper DeleteRoomLogo(ctx, id).Execute()

Remove a room logo



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-room-logo/).

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
	id := int32(1) // int32 | The room ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.DeleteRoomLogo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.DeleteRoomLogo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRoomLogo`: FolderIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.DeleteRoomLogo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoomLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderIntegerWrapper**](FolderIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoomTags

> FolderIntegerWrapper DeleteRoomTags(ctx, id).BatchTagsRequestDto(batchTagsRequestDto).Execute()

Remove the room tags



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-room-tags/).

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
	id := int32(1) // int32 | The room Id.
	batchTagsRequestDto := *openapiclient.NewBatchTagsRequestDto([]string{"Names_example"}) // BatchTagsRequestDto | The parameters for managing tags. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.DeleteRoomTags(context.Background(), id).BatchTagsRequestDto(batchTagsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.DeleteRoomTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRoomTags`: FolderIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.DeleteRoomTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoomTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchTagsRequestDto** | [**BatchTagsRequestDto**](BatchTagsRequestDto.md) | The parameters for managing tags. | 

### Return type

[**FolderIntegerWrapper**](FolderIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalDbSyncStatus

> ExternalDbSyncTaskWrapper GetExternalDbSyncStatus(ctx, id).Execute()

Get external DB sync status



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-external-db-sync-status/).

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
	id := int32(1) // int32 | The room ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.GetExternalDbSyncStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.GetExternalDbSyncStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalDbSyncStatus`: ExternalDbSyncTaskWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.GetExternalDbSyncStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalDbSyncStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalDbSyncTaskWrapper**](ExternalDbSyncTaskWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNewRoomItems

> NewItemsFileEntryBaseArrayWrapper GetNewRoomItems(ctx, id).Execute()

Get the new room items



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-new-room-items/).

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
	id := int32(1) // int32 | The room ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.GetNewRoomItems(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.GetNewRoomItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNewRoomItems`: NewItemsFileEntryBaseArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.GetNewRoomItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNewRoomItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NewItemsFileEntryBaseArrayWrapper**](NewItemsFileEntryBaseArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicSettings

> BooleanWrapper GetPublicSettings(ctx, id).Execute()

Get public settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-public-settings/).

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
	id := int32(1) // int32 | The room template ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.GetPublicSettings(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.GetPublicSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicSettings`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.GetPublicSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room template ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomCovers

> CoversResultArrayWrapper GetRoomCovers(ctx).Execute()

Get covers



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-room-covers/).

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.GetRoomCovers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.GetRoomCovers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomCovers`: CoversResultArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.GetRoomCovers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomCoversRequest struct via the builder pattern


### Return type

[**CoversResultArrayWrapper**](CoversResultArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomCreatingStatus

> RoomFromTemplateStatusWrapper GetRoomCreatingStatus(ctx).Execute()

Get the room creation progress



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-room-creating-status/).

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.GetRoomCreatingStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.GetRoomCreatingStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomCreatingStatus`: RoomFromTemplateStatusWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.GetRoomCreatingStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomCreatingStatusRequest struct via the builder pattern


### Return type

[**RoomFromTemplateStatusWrapper**](RoomFromTemplateStatusWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomIndexExport

> DocumentBuilderTaskWrapper GetRoomIndexExport(ctx).Execute()

Get the room index export



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-room-index-export/).

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.GetRoomIndexExport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.GetRoomIndexExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomIndexExport`: DocumentBuilderTaskWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.GetRoomIndexExport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomIndexExportRequest struct via the builder pattern


### Return type

[**DocumentBuilderTaskWrapper**](DocumentBuilderTaskWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomInfo

> FolderIntegerWrapper GetRoomInfo(ctx, id).Execute()

Get room information



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-room-info/).

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
	id := int32(1) // int32 | The room ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.GetRoomInfo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.GetRoomInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomInfo`: FolderIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.GetRoomInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderIntegerWrapper**](FolderIntegerWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomLinks

> FileShareArrayWrapper GetRoomLinks(ctx, id).Type_(type_).Execute()

Get the room links



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-room-links/).

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
	id := int32(1) // int32 | The room ID.
	type_ := openapiclient.LinkType(0) // LinkType | The link type. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.GetRoomLinks(context.Background(), id).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.GetRoomLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomLinks`: FileShareArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.GetRoomLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**LinkType**](LinkType.md) | The link type. | 

### Return type

[**FileShareArrayWrapper**](FileShareArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomSecurityInfo

> FileShareArrayWrapper GetRoomSecurityInfo(ctx, id).FilterType(filterType).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()

Get the room access rights



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-room-security-info/).

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
	id := int32(1) // int32 | The room ID.
	filterType := openapiclient.ShareFilterType(0) // ShareFilterType | The filter type of the access rights. (optional)
	count := int32(25) // int32 | The number of items to be retrieved or processed. (optional)
	startIndex := int32(0) // int32 | The starting index of the items to retrieve in a paginated request. (optional)
	filterValue := "Sample filter" // string | The text filter value used for filtering room security information. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.GetRoomSecurityInfo(context.Background(), id).FilterType(filterType).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.GetRoomSecurityInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomSecurityInfo`: FileShareArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.GetRoomSecurityInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomSecurityInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterType** | [**ShareFilterType**](ShareFilterType.md) | The filter type of the access rights. | 
 **count** | **int32** | The number of items to be retrieved or processed. | 
 **startIndex** | **int32** | The starting index of the items to retrieve in a paginated request. | 
 **filterValue** | **string** | The text filter value used for filtering room security information. | 

### Return type

[**FileShareArrayWrapper**](FileShareArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomTagsInfo

> ObjectArrayWrapper GetRoomTagsInfo(ctx).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()

Get the room tags



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-room-tags-info/).

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
	count := int32(25) // int32 | Gets or sets the number of tag results to retrieve.  This property specifies the maximum amount of tag data to be included in the result set. (optional)
	startIndex := int32(0) // int32 | Represents the starting index from which the tags' information will be retrieved.  This property is used to define the offset for pagination when retrieving a list of tags. It determines  the point in the data set from which the retrieval begins. (optional)
	filterValue := "My Document" // string | Gets or sets the text value used for searching tags.  This property is typically used as a filter value when retrieving tag information. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.GetRoomTagsInfo(context.Background()).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.GetRoomTagsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomTagsInfo`: ObjectArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.GetRoomTagsInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomTagsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | Gets or sets the number of tag results to retrieve.  This property specifies the maximum amount of tag data to be included in the result set. | 
 **startIndex** | **int32** | Represents the starting index from which the tags' information will be retrieved.  This property is used to define the offset for pagination when retrieving a list of tags. It determines  the point in the data set from which the retrieval begins. | 
 **filterValue** | **string** | Gets or sets the text value used for searching tags.  This property is typically used as a filter value when retrieving tag information. | 

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


## GetRoomTemplateCreatingStatus

> RoomTemplateStatusWrapper GetRoomTemplateCreatingStatus(ctx).Execute()

Get status of room template creation



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-room-template-creating-status/).

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.GetRoomTemplateCreatingStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.GetRoomTemplateCreatingStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomTemplateCreatingStatus`: RoomTemplateStatusWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.GetRoomTemplateCreatingStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomTemplateCreatingStatusRequest struct via the builder pattern


### Return type

[**RoomTemplateStatusWrapper**](RoomTemplateStatusWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomsFolder

> FolderContentIntegerWrapper GetRoomsFolder(ctx).Type_(type_).SubjectId(subjectId).SubjectOwnerId(subjectOwnerId).SearchArea(searchArea).WithoutTags(withoutTags).Tags(tags).ExcludeSubject(excludeSubject).Provider(provider).SubjectFilter(subjectFilter).QuotaFilter(quotaFilter).StorageFilter(storageFilter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).GroupId(groupId).Execute()

Get rooms



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-rooms-folder/).

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
	type_ := []openapiclient.RoomType{[]openapiclient.RoomType{openapiclient.RoomType(1)}} // []RoomType | The filter by room type. (optional)
	subjectId := "00000000-0000-0000-0000-000000000000" // string | The filter by user ID. (optional)
	subjectOwnerId := "00000000-0000-0000-0000-000000000000" // string | The filter by room owner ID. (optional)
	searchArea := openapiclient.SearchArea(0) // SearchArea | The room search area (Active, Archive, Any, Recent by links). (optional)
	withoutTags := false // bool | Specifies whether to search by tags or not. (optional)
	tags := "tag1" // string | The tags in the serialized format. (optional)
	excludeSubject := false // bool | Specifies whether to exclude search by user or group ID. (optional)
	provider := openapiclient.ProviderFilter(0) // ProviderFilter | The filter by provider name (None, Box, DropBox, GoogleDrive, kDrive, OneDrive, SharePoint, WebDav, Yandex, Storage). (optional)
	subjectFilter := openapiclient.SubjectFilter(0) // SubjectFilter | The filter by user (Owner - 0, Member - 1). (optional)
	quotaFilter := openapiclient.QuotaFilter(0) // QuotaFilter | The filter by quota (All - 0, Default - 1, Custom - 2). (optional)
	storageFilter := openapiclient.StorageFilter(0) // StorageFilter | The filter by storage (None - 0, Internal - 1, ThirdParty - 2). (optional)
	count := int32(25) // int32 | Specifies the maximum number of items to retrieve. (optional)
	startIndex := int32(0) // int32 | The index from which to start retrieving the room content. (optional)
	sortBy := "DateAndTime" // string | Specifies the field by which the room content should be sorted. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The order in which the results are sorted. (optional)
	filterValue := "My Document" // string | The text filter value used to refine search or query operations. (optional)
	groupId := int32(1) // int32 | The group ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.GetRoomsFolder(context.Background()).Type_(type_).SubjectId(subjectId).SubjectOwnerId(subjectOwnerId).SearchArea(searchArea).WithoutTags(withoutTags).Tags(tags).ExcludeSubject(excludeSubject).Provider(provider).SubjectFilter(subjectFilter).QuotaFilter(quotaFilter).StorageFilter(storageFilter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).GroupId(groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.GetRoomsFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomsFolder`: FolderContentIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.GetRoomsFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomsFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**[][]RoomType**](array.md) | The filter by room type. | 
 **subjectId** | **string** | The filter by user ID. | 
 **subjectOwnerId** | **string** | The filter by room owner ID. | 
 **searchArea** | [**SearchArea**](SearchArea.md) | The room search area (Active, Archive, Any, Recent by links). | 
 **withoutTags** | **bool** | Specifies whether to search by tags or not. | 
 **tags** | **string** | The tags in the serialized format. | 
 **excludeSubject** | **bool** | Specifies whether to exclude search by user or group ID. | 
 **provider** | [**ProviderFilter**](ProviderFilter.md) | The filter by provider name (None, Box, DropBox, GoogleDrive, kDrive, OneDrive, SharePoint, WebDav, Yandex, Storage). | 
 **subjectFilter** | [**SubjectFilter**](SubjectFilter.md) | The filter by user (Owner - 0, Member - 1). | 
 **quotaFilter** | [**QuotaFilter**](QuotaFilter.md) | The filter by quota (All - 0, Default - 1, Custom - 2). | 
 **storageFilter** | [**StorageFilter**](StorageFilter.md) | The filter by storage (None - 0, Internal - 1, ThirdParty - 2). | 
 **count** | **int32** | Specifies the maximum number of items to retrieve. | 
 **startIndex** | **int32** | The index from which to start retrieving the room content. | 
 **sortBy** | **string** | Specifies the field by which the room content should be sorted. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The order in which the results are sorted. | 
 **filterValue** | **string** | The text filter value used to refine search or query operations. | 
 **groupId** | **int32** | The group ID | 

### Return type

[**FolderContentIntegerWrapper**](FolderContentIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomsNewItems

> NewItemsRoomNewItemsArrayWrapper GetRoomsNewItems(ctx).Execute()

Get the room new items



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-rooms-new-items/).

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.GetRoomsNewItems(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.GetRoomsNewItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomsNewItems`: NewItemsRoomNewItemsArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.GetRoomsNewItems`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomsNewItemsRequest struct via the builder pattern


### Return type

[**NewItemsRoomNewItemsArrayWrapper**](NewItemsRoomNewItemsArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomsPrimaryExternalLink

> FileShareWrapper GetRoomsPrimaryExternalLink(ctx, id).Execute()

Get the room primary external link



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-rooms-primary-external-link/).

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
	id := int32(1) // int32 | The room ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.GetRoomsPrimaryExternalLink(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.GetRoomsPrimaryExternalLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomsPrimaryExternalLink`: FileShareWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.GetRoomsPrimaryExternalLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomsPrimaryExternalLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileShareWrapper**](FileShareWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HasTagLinks

> BooleanWrapper HasTagLinks(ctx, tagName2).TagName(tagName).Execute()

Has tag links



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/has-tag-links/).

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
	tagName2 := "tagName_example" // string | 
	tagName := "tag1" // string | Represents the name of a tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.HasTagLinks(context.Background(), tagName2).TagName(tagName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.HasTagLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HasTagLinks`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.HasTagLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagName2** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHasTagLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagName** | **string** | Represents the name of a tag | 

### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PinRoom

> FolderIntegerWrapper PinRoom(ctx, id).Execute()

Pin a room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/pin-room/).

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
	id := int32(1) // int32 | The room ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.PinRoom(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.PinRoom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PinRoom`: FolderIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.PinRoom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPinRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderIntegerWrapper**](FolderIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReorderRoom

> FolderIntegerWrapper ReorderRoom(ctx, id).Execute()

Reorder the room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/reorder-room/).

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
	id := int32(1) // int32 | The room ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.ReorderRoom(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.ReorderRoom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReorderRoom`: FolderIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.ReorderRoom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReorderRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderIntegerWrapper**](FolderIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendEmailInvitations

> ResendEmailInvitations(ctx, id).UserInvitation(userInvitation).Execute()

Resend the room invitations



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/resend-email-invitations/).

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
	id := int32(1) // int32 | The room ID.
	userInvitation := *openapiclient.NewUserInvitation() // UserInvitation | The user invitation parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoomsAPI.ResendEmailInvitations(context.Background(), id).UserInvitation(userInvitation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.ResendEmailInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendEmailInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userInvitation** | [**UserInvitation**](UserInvitation.md) | The user invitation parameters. | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPublicSettings

> SetPublicSettings(ctx).SetPublicDto(setPublicDto).Execute()

Set public settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-public-settings/).

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
	setPublicDto := *openapiclient.NewSetPublicDto(int32(1)) // SetPublicDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoomsAPI.SetPublicSettings(context.Background()).SetPublicDto(setPublicDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.SetPublicSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPublicSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setPublicDto** | [**SetPublicDto**](SetPublicDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRoomLink

> FileShareWrapper SetRoomLink(ctx, id).RoomLinkRequest(roomLinkRequest).Execute()

Set the room external or invitation link



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-room-link/).

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
	id := int32(1) // int32 | The room ID.
	roomLinkRequest := *openapiclient.NewRoomLinkRequest() // RoomLinkRequest | The room link parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.SetRoomLink(context.Background(), id).RoomLinkRequest(roomLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.SetRoomLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetRoomLink`: FileShareWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.SetRoomLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetRoomLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roomLinkRequest** | [**RoomLinkRequest**](RoomLinkRequest.md) | The room link parameters. | 

### Return type

[**FileShareWrapper**](FileShareWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRoomSecurity

> RoomSecurityWrapper SetRoomSecurity(ctx, id).RoomInvitationRequest(roomInvitationRequest).Execute()

Set the room access rights



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-room-security/).

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
	id := int32(1) // int32 | The room ID.
	roomInvitationRequest := *openapiclient.NewRoomInvitationRequest() // RoomInvitationRequest | The room invitation request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.SetRoomSecurity(context.Background(), id).RoomInvitationRequest(roomInvitationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.SetRoomSecurity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetRoomSecurity`: RoomSecurityWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.SetRoomSecurity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetRoomSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roomInvitationRequest** | [**RoomInvitationRequest**](RoomInvitationRequest.md) | The room invitation request. | 

### Return type

[**RoomSecurityWrapper**](RoomSecurityWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartExternalDbSync

> ExternalDbSyncTaskWrapper StartExternalDbSync(ctx, id).Execute()

Start external DB sync



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/start-external-db-sync/).

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
	id := int32(1) // int32 | The room ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.StartExternalDbSync(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.StartExternalDbSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartExternalDbSync`: ExternalDbSyncTaskWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.StartExternalDbSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartExternalDbSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalDbSyncTaskWrapper**](ExternalDbSyncTaskWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartRoomIndexExport

> DocumentBuilderTaskWrapper StartRoomIndexExport(ctx, id).Execute()

Start the room index export



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/start-room-index-export/).

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
	id := int32(1) // int32 | The room ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.StartRoomIndexExport(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.StartRoomIndexExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartRoomIndexExport`: DocumentBuilderTaskWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.StartRoomIndexExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartRoomIndexExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DocumentBuilderTaskWrapper**](DocumentBuilderTaskWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateRoomIndexExport

> TerminateRoomIndexExport(ctx).Execute()

Terminate the room index export



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/terminate-room-index-export/).

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoomsAPI.TerminateRoomIndexExport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.TerminateRoomIndexExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateRoomIndexExportRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveRoom

> FileOperationWrapper UnarchiveRoom(ctx, id).ArchiveRoomRequest(archiveRoomRequest).Execute()

Unarchive a room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/unarchive-room/).

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
	id := int32(1) // int32 | The room ID.
	archiveRoomRequest := *openapiclient.NewArchiveRoomRequest() // ArchiveRoomRequest | The parameters for archiving a room. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.UnarchiveRoom(context.Background(), id).ArchiveRoomRequest(archiveRoomRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.UnarchiveRoom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnarchiveRoom`: FileOperationWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.UnarchiveRoom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archiveRoomRequest** | [**ArchiveRoomRequest**](ArchiveRoomRequest.md) | The parameters for archiving a room. | 

### Return type

[**FileOperationWrapper**](FileOperationWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnpinRoom

> FolderIntegerWrapper UnpinRoom(ctx, id).Execute()

Unpin a room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/unpin-room/).

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
	id := int32(1) // int32 | The room ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.UnpinRoom(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.UnpinRoom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnpinRoom`: FolderIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.UnpinRoom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpinRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderIntegerWrapper**](FolderIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoom

> FolderIntegerWrapper UpdateRoom(ctx, id).UpdateRoomRequest(updateRoomRequest).Execute()

Update a room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-room/).

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
	id := int32(56) // int32 | The room ID.
	updateRoomRequest := *openapiclient.NewUpdateRoomRequest() // UpdateRoomRequest | The request parameters for updating a room.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.UpdateRoom(context.Background(), id).UpdateRoomRequest(updateRoomRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.UpdateRoom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoom`: FolderIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.UpdateRoom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRoomRequest** | [**UpdateRoomRequest**](UpdateRoomRequest.md) | The request parameters for updating a room. | 

### Return type

[**FolderIntegerWrapper**](FolderIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoomTag

> StringWrapper UpdateRoomTag(ctx).UpdateTagRequestDto(updateTagRequestDto).Execute()

Update tag



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-room-tag/).

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
	updateTagRequestDto := *openapiclient.NewUpdateTagRequestDto("old-tag", "new-tag") // UpdateTagRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.UpdateRoomTag(context.Background()).UpdateTagRequestDto(updateTagRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.UpdateRoomTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoomTag`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.UpdateRoomTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoomTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTagRequestDto** | [**UpdateTagRequestDto**](UpdateTagRequestDto.md) |  | 

### Return type

[**StringWrapper**](StringWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadRoomLogo

> UploadResultWrapper UploadRoomLogo(ctx).File(file).Execute()

Upload a room logo image



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/upload-room-logo/).

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
	file := os.NewFile(1234, "some_file") // *os.File | The image data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.UploadRoomLogo(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.UploadRoomLogo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadRoomLogo`: UploadResultWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.UploadRoomLogo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadRoomLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The image data. | 

### Return type

[**UploadResultWrapper**](UploadResultWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

