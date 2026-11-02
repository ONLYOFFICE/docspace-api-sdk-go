# \RoomsAIAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProjectModel**](RoomsAIAPI.md#GetProjectModel) | **Get** /api/2.0/files/rooms/{id}/model | Get the AI project model
[**GetProjectSources**](RoomsAIAPI.md#GetProjectSources) | **Get** /api/2.0/files/rooms/{id}/sources | Get the AI project sources
[**GetRoomSkills**](RoomsAIAPI.md#GetRoomSkills) | **Get** /api/2.0/files/rooms/{id}/skills | Get the skills of an AI room
[**GetRoomsSkills**](RoomsAIAPI.md#GetRoomsSkills) | **Get** /api/2.0/files/rooms/skills | Get the skills of Instructions rooms
[**LinkProjectSources**](RoomsAIAPI.md#LinkProjectSources) | **Post** /api/2.0/files/rooms/{id}/sources | Link sources to an AI project
[**SetProjectModel**](RoomsAIAPI.md#SetProjectModel) | **Put** /api/2.0/files/rooms/{id}/model | Set the AI project model
[**UnlinkProjectSources**](RoomsAIAPI.md#UnlinkProjectSources) | **Delete** /api/2.0/files/rooms/{id}/sources | Unlink sources from an AI project



## GetProjectModel

> AiProjectModelWrapper GetProjectModel(ctx, id).Execute()

Get the AI project model



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-project-model/).

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
	resp, r, err := apiClient.RoomsAIAPI.GetProjectModel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAIAPI.GetProjectModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectModel`: AiProjectModelWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAIAPI.GetProjectModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AiProjectModelWrapper**](AiProjectModelWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectSources

> AiProjectSourcesWrapper GetProjectSources(ctx, id).Execute()

Get the AI project sources



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-project-sources/).

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
	resp, r, err := apiClient.RoomsAIAPI.GetProjectSources(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAIAPI.GetProjectSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectSources`: AiProjectSourcesWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAIAPI.GetProjectSources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AiProjectSourcesWrapper**](AiProjectSourcesWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomSkills

> AiSkillsWrapper GetRoomSkills(ctx, id).Execute()

Get the skills of an AI room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-room-skills/).

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
	resp, r, err := apiClient.RoomsAIAPI.GetRoomSkills(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAIAPI.GetRoomSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomSkills`: AiSkillsWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAIAPI.GetRoomSkills`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room to act on, named by the identifier that `GET api/2.0/files/rooms` reports for it. Rooms kept in the  portal itself use whole numbers, while a room backed by a connected third-party account uses the string form  of the same listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AiSkillsWrapper**](AiSkillsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomsSkills

> AiSkillsWrapper GetRoomsSkills(ctx).RoomIds(roomIds).Execute()

Get the skills of Instructions rooms



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-rooms-skills/).

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
	roomIds := []int32{int32(123)} // []int32 | IDs of the Instructions rooms, one `roomIds` query item per room; 1 to 20 items. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAIAPI.GetRoomsSkills(context.Background()).RoomIds(roomIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAIAPI.GetRoomsSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomsSkills`: AiSkillsWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAIAPI.GetRoomsSkills`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomsSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roomIds** | **[]int32** | IDs of the Instructions rooms, one `roomIds` query item per room; 1 to 20 items. | 

### Return type

[**AiSkillsWrapper**](AiSkillsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkProjectSources

> LinkProjectSources(ctx, id).LinkAiProjectSourcesRequest(linkAiProjectSourcesRequest).Execute()

Link sources to an AI project



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/link-project-sources/).

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
	id := int32(1) // int32 | ID of the project.
	linkAiProjectSourcesRequest := *openapiclient.NewLinkAiProjectSourcesRequest([]int32{int32(123)}) // LinkAiProjectSourcesRequest | The rooms to link.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoomsAIAPI.LinkProjectSources(context.Background(), id).LinkAiProjectSourcesRequest(linkAiProjectSourcesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAIAPI.LinkProjectSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLinkProjectSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **linkAiProjectSourcesRequest** | [**LinkAiProjectSourcesRequest**](LinkAiProjectSourcesRequest.md) | The rooms to link. | 

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


## SetProjectModel

> AiProjectModelWrapper SetProjectModel(ctx, id).UpdateAiProjectModelRequest(updateAiProjectModelRequest).Execute()

Set the AI project model



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-project-model/).

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
	id := int32(1) // int32 | ID of the project.
	updateAiProjectModelRequest := *openapiclient.NewUpdateAiProjectModelRequest("0198c3f4-6a2e-7c41-9b8d-2f5e7a1c3d90") // UpdateAiProjectModelRequest | The model to set.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsAIAPI.SetProjectModel(context.Background(), id).UpdateAiProjectModelRequest(updateAiProjectModelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAIAPI.SetProjectModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetProjectModel`: AiProjectModelWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsAIAPI.SetProjectModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetProjectModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAiProjectModelRequest** | [**UpdateAiProjectModelRequest**](UpdateAiProjectModelRequest.md) | The model to set. | 

### Return type

[**AiProjectModelWrapper**](AiProjectModelWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkProjectSources

> UnlinkProjectSources(ctx, id).UnlinkAiProjectSourcesRequest(unlinkAiProjectSourcesRequest).Execute()

Unlink sources from an AI project



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/unlink-project-sources/).

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
	id := int32(1) // int32 | ID of the project.
	unlinkAiProjectSourcesRequest := *openapiclient.NewUnlinkAiProjectSourcesRequest([]int32{int32(123)}) // UnlinkAiProjectSourcesRequest | The rooms to unlink.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoomsAIAPI.UnlinkProjectSources(context.Background(), id).UnlinkAiProjectSourcesRequest(unlinkAiProjectSourcesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsAIAPI.UnlinkProjectSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkProjectSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unlinkAiProjectSourcesRequest** | [**UnlinkAiProjectSourcesRequest**](UnlinkAiProjectSourcesRequest.md) | The rooms to unlink. | 

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

