# \RoomsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRoomTags**](RoomsAPI.md#AddRoomTags) | **Put** /api/2.0/files/rooms/{id}/tags | Attach tags to a room
[**ArchiveRoom**](RoomsAPI.md#ArchiveRoom) | **Put** /api/2.0/files/rooms/{id}/archive | Archive a room
[**ChangeRoomCover**](RoomsAPI.md#ChangeRoomCover) | **Post** /api/2.0/files/rooms/{id}/cover | Change the room cover
[**CreateRoom**](RoomsAPI.md#CreateRoom) | **Post** /api/2.0/files/rooms | Create a room
[**CreateRoomFromTemplate**](RoomsAPI.md#CreateRoomFromTemplate) | **Post** /api/2.0/files/rooms/fromtemplate | Create a room from the template
[**CreateRoomLogo**](RoomsAPI.md#CreateRoomLogo) | **Post** /api/2.0/files/rooms/{id}/logo | Set the room logo
[**CreateRoomTag**](RoomsAPI.md#CreateRoomTag) | **Post** /api/2.0/files/tags | Create a room tag
[**CreateRoomTemplate**](RoomsAPI.md#CreateRoomTemplate) | **Post** /api/2.0/files/roomtemplate | Create a room template
[**CreateRoomThirdParty**](RoomsAPI.md#CreateRoomThirdParty) | **Post** /api/2.0/files/rooms/thirdparty/{id} | Create a third-party room
[**DeleteCustomTags**](RoomsAPI.md#DeleteCustomTags) | **Delete** /api/2.0/files/tags | Delete the custom room tags
[**DeleteRoom**](RoomsAPI.md#DeleteRoom) | **Delete** /api/2.0/files/rooms/{id} | Remove a room
[**DeleteRoomLogo**](RoomsAPI.md#DeleteRoomLogo) | **Delete** /api/2.0/files/rooms/{id}/logo | Remove a room logo
[**DeleteRoomTags**](RoomsAPI.md#DeleteRoomTags) | **Delete** /api/2.0/files/rooms/{id}/tags | Detach tags from a room
[**GetExternalDbSyncStatus**](RoomsAPI.md#GetExternalDbSyncStatus) | **Get** /api/2.0/files/rooms/{id}/externaldbsync | Get external DB sync status
[**GetNewRoomItems**](RoomsAPI.md#GetNewRoomItems) | **Get** /api/2.0/files/rooms/{id}/news | Get new items in a room
[**GetPublicSettings**](RoomsAPI.md#GetPublicSettings) | **Get** /api/2.0/files/roomtemplate/{id}/public | Get room template public access
[**GetRoomCovers**](RoomsAPI.md#GetRoomCovers) | **Get** /api/2.0/files/rooms/covers | Get room cover gallery
[**GetRoomCreatingStatus**](RoomsAPI.md#GetRoomCreatingStatus) | **Get** /api/2.0/files/rooms/fromtemplate/status | Get the room creation progress
[**GetRoomIndexExport**](RoomsAPI.md#GetRoomIndexExport) | **Get** /api/2.0/files/rooms/indexexport | Get the room index export
[**GetRoomInfo**](RoomsAPI.md#GetRoomInfo) | **Get** /api/2.0/files/rooms/{id} | Get room information
[**GetRoomLinks**](RoomsAPI.md#GetRoomLinks) | **Get** /api/2.0/files/rooms/{id}/links | Get the room links
[**GetRoomSecurityInfo**](RoomsAPI.md#GetRoomSecurityInfo) | **Get** /api/2.0/files/rooms/{id}/share | Get the room access rights
[**GetRoomTagsInfo**](RoomsAPI.md#GetRoomTagsInfo) | **Get** /api/2.0/files/tags | Get available room tags
[**GetRoomTemplateCreatingStatus**](RoomsAPI.md#GetRoomTemplateCreatingStatus) | **Get** /api/2.0/files/roomtemplate/status | Get room template creation status
[**GetRoomsFolder**](RoomsAPI.md#GetRoomsFolder) | **Get** /api/2.0/files/rooms | Get rooms
[**GetRoomsNewItems**](RoomsAPI.md#GetRoomsNewItems) | **Get** /api/2.0/files/rooms/news | Get new items in all rooms
[**GetRoomsPrimaryExternalLink**](RoomsAPI.md#GetRoomsPrimaryExternalLink) | **Get** /api/2.0/files/rooms/{id}/link | Get the room primary external link
[**HasTagLinks**](RoomsAPI.md#HasTagLinks) | **Get** /api/2.0/files/tags/{tagName}/haslinks | Check room tag usage
[**PinRoom**](RoomsAPI.md#PinRoom) | **Put** /api/2.0/files/rooms/{id}/pin | Pin a room
[**ReorderRoom**](RoomsAPI.md#ReorderRoom) | **Put** /api/2.0/files/rooms/{id}/reorder | Reorder room contents
[**ResendEmailInvitations**](RoomsAPI.md#ResendEmailInvitations) | **Post** /api/2.0/files/rooms/{id}/resend | Resend the room invitations
[**SetPublicSettings**](RoomsAPI.md#SetPublicSettings) | **Put** /api/2.0/files/roomtemplate/public | Set room template public access
[**SetRoomLink**](RoomsAPI.md#SetRoomLink) | **Put** /api/2.0/files/rooms/{id}/links | Set the room external or invitation link
[**SetRoomSecurity**](RoomsAPI.md#SetRoomSecurity) | **Put** /api/2.0/files/rooms/{id}/share | Set the room access rights
[**StartExternalDbSync**](RoomsAPI.md#StartExternalDbSync) | **Post** /api/2.0/files/rooms/{id}/externaldbsync | Start external DB sync
[**StartRoomIndexExport**](RoomsAPI.md#StartRoomIndexExport) | **Post** /api/2.0/files/rooms/{id}/indexexport | Start the room index export
[**TerminateRoomIndexExport**](RoomsAPI.md#TerminateRoomIndexExport) | **Delete** /api/2.0/files/rooms/indexexport | Terminate the room index export
[**UnarchiveRoom**](RoomsAPI.md#UnarchiveRoom) | **Put** /api/2.0/files/rooms/{id}/unarchive | Unarchive a room
[**UnpinRoom**](RoomsAPI.md#UnpinRoom) | **Put** /api/2.0/files/rooms/{id}/unpin | Unpin a room
[**UpdateRoom**](RoomsAPI.md#UpdateRoom) | **Put** /api/2.0/files/rooms/{id} | Update a room
[**UpdateRoomTag**](RoomsAPI.md#UpdateRoomTag) | **Put** /api/2.0/files/tags | Rename a room tag
[**UploadRoomLogo**](RoomsAPI.md#UploadRoomLogo) | **Post** /api/2.0/files/logos | Upload a room logo image



## AddRoomTags

> FolderWrapper AddRoomTags(ctx, id).BatchTagsRequestDto(batchTagsRequestDto).Execute()

Attach tags to a room



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
	id := int32(1) // int32 | The room whose tags are changed, named by the identifier that `GET api/2.0/files/rooms` reports for it.
	batchTagsRequestDto := *openapiclient.NewBatchTagsRequestDto([]string{"Names_example"}) // BatchTagsRequestDto | The names to attach or to detach. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.AddRoomTags(context.Background(), id).BatchTagsRequestDto(batchTagsRequestDto).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// id := "sbox-42"
	// thirdPartyResp, r, err := apiClient.RoomsAPI.AddRoomTags(context.Background(), id).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.AddRoomTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddRoomTags`: FolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.AddRoomTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room whose tags are changed, named by the identifier that `GET api/2.0/files/rooms` reports for it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRoomTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchTagsRequestDto** | [**BatchTagsRequestDto**](BatchTagsRequestDto.md) | The names to attach or to detach. | 

### Return type

[**FolderWrapper**](FolderWrapper.md)

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
	id := int32(1) // int32 | The room to move, named by the identifier that `GET api/2.0/files/rooms` reports for it.
	archiveRoomRequest := *openapiclient.NewArchiveRoomRequest() // ArchiveRoomRequest | The body of the request. It carries only the lifetime of the job record, so an empty object is a normal  request. (optional)

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
**id** | **int32** | The room to move, named by the identifier that `GET api/2.0/files/rooms` reports for it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archiveRoomRequest** | [**ArchiveRoomRequest**](ArchiveRoomRequest.md) | The body of the request. It carries only the lifetime of the job record, so an empty object is a normal  request. | 

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

> FolderWrapper ChangeRoomCover(ctx, id).CoverRequestDto(coverRequestDto).Execute()

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
	id := int32(1) // int32 | The room to change, named by the identifier that `GET api/2.0/files/rooms` reports for it.
	coverRequestDto := *openapiclient.NewCoverRequestDto() // CoverRequestDto | The cover and the colour to apply. Either half may be sent on its own, and an empty object leaves the room as  it is.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.ChangeRoomCover(context.Background(), id).CoverRequestDto(coverRequestDto).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// id := "sbox-42"
	// thirdPartyResp, r, err := apiClient.RoomsAPI.ChangeRoomCover(context.Background(), id).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.ChangeRoomCover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeRoomCover`: FolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.ChangeRoomCover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room to change, named by the identifier that `GET api/2.0/files/rooms` reports for it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeRoomCoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **coverRequestDto** | [**CoverRequestDto**](CoverRequestDto.md) | The cover and the colour to apply. Either half may be sent on its own, and an empty object leaves the room as  it is. | 

### Return type

[**FolderWrapper**](FolderWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRoom

> FolderWrapper CreateRoom(ctx).CreateRoomRequestDto(createRoomRequestDto).Execute()

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
	createRoomRequestDto := *openapiclient.NewCreateRoomRequestDto("Project Alpha", openapiclient.RoomType(1)) // CreateRoomRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.CreateRoom(context.Background()).CreateRoomRequestDto(createRoomRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.CreateRoom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoom`: FolderWrapper
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

[**FolderWrapper**](FolderWrapper.md)

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
	createRoomFromTemplateDto := *openapiclient.NewCreateRoomFromTemplateDto(int32(42), "Project Alpha") // CreateRoomFromTemplateDto |  (optional)

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

> FolderWrapper CreateRoomLogo(ctx, id).LogoRequest(logoRequest).Execute()

Set the room logo



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
	id := int32(1) // int32 | The room the logo is set on.
	logoRequest := *openapiclient.NewLogoRequest("/temp/logo_a1b2c3.png") // LogoRequest | The uploaded picture and the piece of it to use.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.CreateRoomLogo(context.Background(), id).LogoRequest(logoRequest).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// id := "sbox-42"
	// thirdPartyResp, r, err := apiClient.RoomsAPI.CreateRoomLogo(context.Background(), id).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.CreateRoomLogo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoomLogo`: FolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.CreateRoomLogo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room the logo is set on. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoomLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logoRequest** | [**LogoRequest**](LogoRequest.md) | The uploaded picture and the piece of it to use. | 

### Return type

[**FolderWrapper**](FolderWrapper.md)

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

Create a room template



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
	roomTemplateDto := *openapiclient.NewRoomTemplateDto(int32(1234), "Sales agreement room") // RoomTemplateDto |  (optional)

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

> ThirdPartyFolderWrapper CreateRoomThirdParty(ctx, id).CreateThirdPartyRoom(createThirdPartyRoom).Execute()

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
	id := "box-12-|280143035119" // string | The identifier of the folder in the connected third-party storage that becomes the room, or receives it as a  subfolder. Folder identifiers of a connected account are strings and are returned by the folder listings of  that account.
	createThirdPartyRoom := *openapiclient.NewCreateThirdPartyRoom("Third-party project room", openapiclient.RoomType(1)) // CreateThirdPartyRoom | The settings of the room to be created out of the folder.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.CreateRoomThirdParty(context.Background(), id).CreateThirdPartyRoom(createThirdPartyRoom).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.CreateRoomThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoomThirdParty`: ThirdPartyFolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.CreateRoomThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identifier of the folder in the connected third-party storage that becomes the room, or receives it as a  subfolder. Folder identifiers of a connected account are strings and are returned by the folder listings of  that account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoomThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createThirdPartyRoom** | [**CreateThirdPartyRoom**](CreateThirdPartyRoom.md) | The settings of the room to be created out of the folder. | 

### Return type

[**ThirdPartyFolderWrapper**](ThirdPartyFolderWrapper.md)

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
- **Accept**: application/json

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
	id := int32(10) // int32 | The room to delete, named by the identifier that `GET api/2.0/files/rooms` reports for it.
	deleteRoomRequest := *openapiclient.NewDeleteRoomRequest() // DeleteRoomRequest | The body of the request. It is required even though the deletion does not depend on what it holds.

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
**id** | **int32** | The room to delete, named by the identifier that `GET api/2.0/files/rooms` reports for it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteRoomRequest** | [**DeleteRoomRequest**](DeleteRoomRequest.md) | The body of the request. It is required even though the deletion does not depend on what it holds. | 

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

> FolderWrapper DeleteRoomLogo(ctx, id).Execute()

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
	id := int32(1) // int32 | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.DeleteRoomLogo(context.Background(), id).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// id := "sbox-42"
	// thirdPartyResp, r, err := apiClient.RoomsAPI.DeleteRoomLogo(context.Background(), id).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.DeleteRoomLogo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRoomLogo`: FolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.DeleteRoomLogo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoomLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderWrapper**](FolderWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoomTags

> FolderWrapper DeleteRoomTags(ctx, id).BatchTagsRequestDto(batchTagsRequestDto).Execute()

Detach tags from a room



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
	id := int32(1) // int32 | The room whose tags are changed, named by the identifier that `GET api/2.0/files/rooms` reports for it.
	batchTagsRequestDto := *openapiclient.NewBatchTagsRequestDto([]string{"Names_example"}) // BatchTagsRequestDto | The names to attach or to detach. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.DeleteRoomTags(context.Background(), id).BatchTagsRequestDto(batchTagsRequestDto).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// id := "sbox-42"
	// thirdPartyResp, r, err := apiClient.RoomsAPI.DeleteRoomTags(context.Background(), id).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.DeleteRoomTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRoomTags`: FolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.DeleteRoomTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room whose tags are changed, named by the identifier that `GET api/2.0/files/rooms` reports for it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoomTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchTagsRequestDto** | [**BatchTagsRequestDto**](BatchTagsRequestDto.md) | The names to attach or to detach. | 

### Return type

[**FolderWrapper**](FolderWrapper.md)

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
	id := int32(1) // int32 | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing.

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
**id** | **int32** | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing. | 

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

Get new items in a room



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
	id := int32(1) // int32 | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing.

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
**id** | **int32** | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing. | 

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

Get room template public access



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
	id := int32(1234) // int32 | The identifier of the room template. Take it from `templateId` of `GET api/2.0/files/roomtemplate/status`, or  from the folder list of `GET api/2.0/files/rooms` called with `searchArea` set to 4; an identifier of an  ordinary room is not accepted.

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
**id** | **int32** | The identifier of the room template. Take it from `templateId` of `GET api/2.0/files/roomtemplate/status`, or  from the folder list of `GET api/2.0/files/rooms` called with `searchArea` set to 4; an identifier of an  ordinary room is not accepted. | 

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

Get room cover gallery



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

> FolderWrapper GetRoomInfo(ctx, id).Execute()

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
	id := int32(1) // int32 | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.GetRoomInfo(context.Background(), id).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// id := "sbox-42"
	// thirdPartyResp, r, err := apiClient.RoomsAPI.GetRoomInfo(context.Background(), id).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.GetRoomInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomInfo`: FolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.GetRoomInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderWrapper**](FolderWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

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
	id := int32(1) // int32 | The room whose links are listed, named by the identifier that `GET api/2.0/files/rooms` reports for it.
	type_ := openapiclient.LinkType(0) // LinkType | Narrows the answer to one kind of link: invitation links, which turn whoever opens them into a member, or  external links, which open the room without an account. Leaving it out returns both kinds together. (optional)

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
**id** | **int32** | The room whose links are listed, named by the identifier that `GET api/2.0/files/rooms` reports for it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**LinkType**](LinkType.md) | Narrows the answer to one kind of link: invitation links, which turn whoever opens them into a member, or  external links, which open the room without an account. Leaving it out returns both kinds together. | 

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
	id := int32(1) // int32 | The room whose access list is read, named by the identifier that `GET api/2.0/files/rooms` reports for it.
	filterType := openapiclient.ShareFilterType(0) // ShareFilterType | What kind of access entries to list. The default covers accounts and groups and leaves the sharing links of  the room out; those are read with `GET api/2.0/files/rooms/{id}/links`. (optional)
	count := int32(25) // int32 | How many entries to return in one answer. The total number of matching entries comes back in the response  headers, so it is what tells the caller whether another page is needed. (optional)
	startIndex := int32(0) // int32 | How many matching entries to skip before the page starts. Together with the page size it walks the list, which  is ordered by role and then by name and is therefore stable between calls. (optional)
	filterValue := "Smith" // string | Keeps only the entries whose displayed name contains this text. An invitation that has not been accepted yet  is listed under the email address it was sent to, so that is what has to be searched for. (optional)

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
**id** | **int32** | The room whose access list is read, named by the identifier that `GET api/2.0/files/rooms` reports for it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomSecurityInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterType** | [**ShareFilterType**](ShareFilterType.md) | What kind of access entries to list. The default covers accounts and groups and leaves the sharing links of  the room out; those are read with `GET api/2.0/files/rooms/{id}/links`. | 
 **count** | **int32** | How many entries to return in one answer. The total number of matching entries comes back in the response  headers, so it is what tells the caller whether another page is needed. | 
 **startIndex** | **int32** | How many matching entries to skip before the page starts. Together with the page size it walks the list, which  is ordered by role and then by name and is therefore stable between calls. | 
 **filterValue** | **string** | Keeps only the entries whose displayed name contains this text. An invitation that has not been accepted yet  is listed under the email address it was sent to, so that is what has to be searched for. | 

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

> STRINGArrayWrapper GetRoomTagsInfo(ctx).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()

Get available room tags



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
	count := int32(25) // int32 | How many tag names one page may carry. The answer reports no total, so a page shorter than this is the sign  that the list is exhausted. (optional)
	startIndex := int32(0) // int32 | How many tag names to skip before the page begins. Raise it by the number of names already received to read  the next page. (optional)
	filterValue := "conf" // string | Keeps only the tag names that contain this text, ignoring case. It is a substring match, so a fragment from  the middle of a name is enough. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.GetRoomTagsInfo(context.Background()).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.GetRoomTagsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomTagsInfo`: STRINGArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.GetRoomTagsInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomTagsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | How many tag names one page may carry. The answer reports no total, so a page shorter than this is the sign  that the list is exhausted. | 
 **startIndex** | **int32** | How many tag names to skip before the page begins. Raise it by the number of names already received to read  the next page. | 
 **filterValue** | **string** | Keeps only the tag names that contain this text, ignoring case. It is a substring match, so a fragment from  the middle of a name is enough. | 

### Return type

[**STRINGArrayWrapper**](STRINGArrayWrapper.md)

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

Get room template creation status



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

> FolderContentWrapper GetRoomsFolder(ctx).Type_(type_).SubjectId(subjectId).SubjectOwnerId(subjectOwnerId).SearchArea(searchArea).WithoutTags(withoutTags).Tags(tags).ExcludeSubject(excludeSubject).Provider(provider).QuotaFilter(quotaFilter).StorageFilter(storageFilter).PrivacyFilter(privacyFilter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).GroupId(groupId).Execute()

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
	type_ := []openapiclient.RoomType{openapiclient.RoomType(1)} // []RoomType | Keeps only the rooms of the listed kinds. Repeat the parameter to pass more than one value; they are combined  with OR, and omitting it returns the rooms of every kind. (optional)
	subjectId := "9a1b2c3d-4e5f-6071-8293-a4b5c6d7e8f9" // string | Keeps only the rooms this account or group has access to, which is how the rooms of one member are listed. The  identifier comes from the portal people and group listings, and the exclude flag turns the filter into its  opposite. (optional)
	subjectOwnerId := "9a1b2c3d-4e5f-6071-8293-a4b5c6d7e8f9" // string | Keeps only the rooms created by this account, regardless of who else was invited to them. The identifier comes  from the portal people listing, and the exclude flag turns the filter into its opposite. (optional)
	searchArea := openapiclient.SearchArea("Active") // SearchArea | The section to list. Every section is a separate root and a room belongs to exactly one of them at a time, so  archiving a room moves it out of the active section. The default is the active section, which leaves the  form-filling rooms to their own value. (optional)
	withoutTags := false // bool | When true, keeps only the rooms that carry no tag at all, which is the complement of the tag filter. When  false or omitted, tags play no part in the selection. (optional)
	tags := "[\"Important\"]" // string | A JSON array of tag names serialized into a single query value, for example [Important,Legal]. A room  matches when it carries any one of them. Take the names from `GET api/2.0/files/tags`; a name that is not in  the catalog simply matches nothing. (optional)
	excludeSubject := false // bool | Inverts the two subject filters: when true, the rooms of the named account are the ones left out of the answer  instead of the only ones kept. It does nothing on its own. (optional)
	provider := openapiclient.ProviderFilter(0) // ProviderFilter | Keeps only the rooms whose content lives in the named third-party service, for portals where rooms may be  connected to external storage. The default keeps rooms of every origin. (optional)
	quotaFilter := openapiclient.QuotaFilter(0) // QuotaFilter | Splits the rooms by whether a storage quota was set on the room itself or it follows the portal default, which  is how rooms with a custom limit are found. (optional)
	storageFilter := openapiclient.StorageFilter(0) // StorageFilter | Splits the rooms by where their content is stored, in the portal itself or in a connected third-party account.  It is the coarse form of the provider filter. (optional)
	privacyFilter := openapiclient.RoomPrivacyFilter(0) // RoomPrivacyFilter | Splits the rooms by whether they are private, that is encrypted rooms whose content the portal cannot read.  Omitting it returns both kinds. (optional)
	count := int32(25) // int32 | How many rooms one page may carry. Ask for the next page by raising the start index by the number of rooms  already received. (optional)
	startIndex := int32(0) // int32 | How many matching rooms to skip before the page begins. Page through the answer until the skip plus the rooms  received reaches the total it reports. (optional)
	sortBy := "DateAndTime" // string | The field to order the rooms by, named as in the file listings: `AZ` for the title, `DateAndTime` for the last  change, `DateAndTimeCreation`, `Author`, `Size`, `Type`, `RoomType`, `Tags`, `UsedSpace`, `LastOpened`. The  name is matched ignoring case, an unknown one is rejected rather than ignored, and the accepted one also  becomes this account's stored order. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The direction of the order chosen by the sort field. It has no effect when no sort field is given and the  stored order of the account is used. (optional)
	filterValue := "Sales" // string | Keeps only the rooms whose title contains this text, ignoring case. It is a substring match over the title  alone: room content and tags are not searched. (optional)
	groupId := int32(1) // int32 | Keeps only the rooms that belong to this room group. The identifier comes from `GET api/2.0/files/group`; the  groups of portal members are a different concept and their identifiers do not match here. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.GetRoomsFolder(context.Background()).Type_(type_).SubjectId(subjectId).SubjectOwnerId(subjectOwnerId).SearchArea(searchArea).WithoutTags(withoutTags).Tags(tags).ExcludeSubject(excludeSubject).Provider(provider).QuotaFilter(quotaFilter).StorageFilter(storageFilter).PrivacyFilter(privacyFilter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).GroupId(groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.GetRoomsFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomsFolder`: FolderContentWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.GetRoomsFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomsFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**[]RoomType**](RoomType.md) | Keeps only the rooms of the listed kinds. Repeat the parameter to pass more than one value; they are combined  with OR, and omitting it returns the rooms of every kind. | 
 **subjectId** | **string** | Keeps only the rooms this account or group has access to, which is how the rooms of one member are listed. The  identifier comes from the portal people and group listings, and the exclude flag turns the filter into its  opposite. | 
 **subjectOwnerId** | **string** | Keeps only the rooms created by this account, regardless of who else was invited to them. The identifier comes  from the portal people listing, and the exclude flag turns the filter into its opposite. | 
 **searchArea** | [**SearchArea**](SearchArea.md) | The section to list. Every section is a separate root and a room belongs to exactly one of them at a time, so  archiving a room moves it out of the active section. The default is the active section, which leaves the  form-filling rooms to their own value. | 
 **withoutTags** | **bool** | When true, keeps only the rooms that carry no tag at all, which is the complement of the tag filter. When  false or omitted, tags play no part in the selection. | 
 **tags** | **string** | A JSON array of tag names serialized into a single query value, for example [Important,Legal]. A room  matches when it carries any one of them. Take the names from `GET api/2.0/files/tags`; a name that is not in  the catalog simply matches nothing. | 
 **excludeSubject** | **bool** | Inverts the two subject filters: when true, the rooms of the named account are the ones left out of the answer  instead of the only ones kept. It does nothing on its own. | 
 **provider** | [**ProviderFilter**](ProviderFilter.md) | Keeps only the rooms whose content lives in the named third-party service, for portals where rooms may be  connected to external storage. The default keeps rooms of every origin. | 
 **quotaFilter** | [**QuotaFilter**](QuotaFilter.md) | Splits the rooms by whether a storage quota was set on the room itself or it follows the portal default, which  is how rooms with a custom limit are found. | 
 **storageFilter** | [**StorageFilter**](StorageFilter.md) | Splits the rooms by where their content is stored, in the portal itself or in a connected third-party account.  It is the coarse form of the provider filter. | 
 **privacyFilter** | [**RoomPrivacyFilter**](RoomPrivacyFilter.md) | Splits the rooms by whether they are private, that is encrypted rooms whose content the portal cannot read.  Omitting it returns both kinds. | 
 **count** | **int32** | How many rooms one page may carry. Ask for the next page by raising the start index by the number of rooms  already received. | 
 **startIndex** | **int32** | How many matching rooms to skip before the page begins. Page through the answer until the skip plus the rooms  received reaches the total it reports. | 
 **sortBy** | **string** | The field to order the rooms by, named as in the file listings: `AZ` for the title, `DateAndTime` for the last  change, `DateAndTimeCreation`, `Author`, `Size`, `Type`, `RoomType`, `Tags`, `UsedSpace`, `LastOpened`. The  name is matched ignoring case, an unknown one is rejected rather than ignored, and the accepted one also  becomes this account's stored order. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The direction of the order chosen by the sort field. It has no effect when no sort field is given and the  stored order of the account is used. | 
 **filterValue** | **string** | Keeps only the rooms whose title contains this text, ignoring case. It is a substring match over the title  alone: room content and tags are not searched. | 
 **groupId** | **int32** | Keeps only the rooms that belong to this room group. The identifier comes from `GET api/2.0/files/group`; the  groups of portal members are a different concept and their identifiers do not match here. | 

### Return type

[**FolderContentWrapper**](FolderContentWrapper.md)

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

Get new items in all rooms



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
	id := int32(1) // int32 | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing.

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
**id** | **int32** | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing. | 

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

Check room tag usage



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
	tagName2 := "tagName_example" // string | The tag being checked. Send the same value as the `tagName` query parameter, which is the one the handler reads.
	tagName := "Important" // string | The tag to check, spelled exactly as it is stored in the catalog. This query value is the one the handler  reads, so the path segment of the same name has to repeat it. (optional)

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
**tagName2** | **string** | The tag being checked. Send the same value as the `tagName` query parameter, which is the one the handler reads. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHasTagLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagName** | **string** | The tag to check, spelled exactly as it is stored in the catalog. This query value is the one the handler  reads, so the path segment of the same name has to repeat it. | 

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

> FolderWrapper PinRoom(ctx, id).Execute()

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
	id := int32(1) // int32 | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.PinRoom(context.Background(), id).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// id := "sbox-42"
	// thirdPartyResp, r, err := apiClient.RoomsAPI.PinRoom(context.Background(), id).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.PinRoom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PinRoom`: FolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.PinRoom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPinRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderWrapper**](FolderWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReorderRoom

> FolderWrapper ReorderRoom(ctx, id).Execute()

Reorder room contents



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
	id := int32(1) // int32 | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.ReorderRoom(context.Background(), id).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// id := "sbox-42"
	// thirdPartyResp, r, err := apiClient.RoomsAPI.ReorderRoom(context.Background(), id).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.ReorderRoom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReorderRoom`: FolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.ReorderRoom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReorderRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderWrapper**](FolderWrapper.md)

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
	id := int32(1) // int32 | The room whose invitations are resent, named by the identifier that `GET api/2.0/files/rooms` reports for it.
	userInvitation := *openapiclient.NewUserInvitation() // UserInvitation | Which pending invitations to send again.

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
**id** | **int32** | The room whose invitations are resent, named by the identifier that `GET api/2.0/files/rooms` reports for it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendEmailInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userInvitation** | [**UserInvitation**](UserInvitation.md) | Which pending invitations to send again. | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPublicSettings

> SetPublicSettings(ctx).SetPublicDto(setPublicDto).Execute()

Set room template public access



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
	setPublicDto := *openapiclient.NewSetPublicDto(int32(1234)) // SetPublicDto |  (optional)

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
- **Accept**: application/json

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
	id := int32(1) // int32 | The room the link belongs to, named by the identifier that `GET api/2.0/files/rooms` reports for it.
	roomLinkRequest := *openapiclient.NewRoomLinkRequest() // RoomLinkRequest | The link to create, change or revoke.

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
**id** | **int32** | The room the link belongs to, named by the identifier that `GET api/2.0/files/rooms` reports for it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetRoomLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roomLinkRequest** | [**RoomLinkRequest**](RoomLinkRequest.md) | The link to create, change or revoke. | 

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
	id := int32(1) // int32 | The room whose membership changes, named by the identifier that `GET api/2.0/files/rooms` reports for it.
	roomInvitationRequest := *openapiclient.NewRoomInvitationRequest() // RoomInvitationRequest | The membership changes to apply, together with how the people concerned are notified.

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
**id** | **int32** | The room whose membership changes, named by the identifier that `GET api/2.0/files/rooms` reports for it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetRoomSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roomInvitationRequest** | [**RoomInvitationRequest**](RoomInvitationRequest.md) | The membership changes to apply, together with how the people concerned are notified. | 

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
	id := int32(1) // int32 | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing.

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
**id** | **int32** | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing. | 

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
	id := int32(1) // int32 | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing.

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
**id** | **int32** | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing. | 

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
- **Accept**: application/json

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
	id := int32(1) // int32 | The room to move, named by the identifier that `GET api/2.0/files/rooms` reports for it.
	archiveRoomRequest := *openapiclient.NewArchiveRoomRequest() // ArchiveRoomRequest | The body of the request. It carries only the lifetime of the job record, so an empty object is a normal  request. (optional)

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
**id** | **int32** | The room to move, named by the identifier that `GET api/2.0/files/rooms` reports for it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archiveRoomRequest** | [**ArchiveRoomRequest**](ArchiveRoomRequest.md) | The body of the request. It carries only the lifetime of the job record, so an empty object is a normal  request. | 

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

> FolderWrapper UnpinRoom(ctx, id).Execute()

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
	id := int32(1) // int32 | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.UnpinRoom(context.Background(), id).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// id := "sbox-42"
	// thirdPartyResp, r, err := apiClient.RoomsAPI.UnpinRoom(context.Background(), id).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.UnpinRoom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnpinRoom`: FolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.UnpinRoom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpinRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderWrapper**](FolderWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoom

> FolderWrapper UpdateRoom(ctx, id).UpdateRoomRequest(updateRoomRequest).Execute()

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
	id := int32(1) // int32 | The room to update, named by the identifier that `GET api/2.0/files/rooms` reports for it.
	updateRoomRequest := *openapiclient.NewUpdateRoomRequest() // UpdateRoomRequest | The fields to change. Only the properties present in the object are applied, and a property that the object  does not define is rejected instead of being ignored.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAPI.UpdateRoom(context.Background(), id).UpdateRoomRequest(updateRoomRequest).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// id := "sbox-42"
	// thirdPartyResp, r, err := apiClient.RoomsAPI.UpdateRoom(context.Background(), id).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAPI.UpdateRoom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoom`: FolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAPI.UpdateRoom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room to update, named by the identifier that `GET api/2.0/files/rooms` reports for it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRoomRequest** | [**UpdateRoomRequest**](UpdateRoomRequest.md) | The fields to change. Only the properties present in the object are applied, and a property that the object  does not define is rejected instead of being ignored. | 

### Return type

[**FolderWrapper**](FolderWrapper.md)

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

Rename a room tag



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
	updateTagRequestDto := *openapiclient.NewUpdateTagRequestDto("Confidential", "Restricted") // UpdateTagRequestDto |  (optional)

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

