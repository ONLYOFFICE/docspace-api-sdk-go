# \AIContextAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiContextGetContextFolders**](AIContextAPI.md#AiContextGetContextFolders) | **Get** /api/2.0/ai/context/get-context-folders | Get context folders
[**AiContextGetKnowledgeFolders**](AIContextAPI.md#AiContextGetKnowledgeFolders) | **Get** /api/2.0/ai/context/get-knowledge-folders | Get knowledge folders
[**AiContextGetRoomSkill**](AIContextAPI.md#AiContextGetRoomSkill) | **Get** /api/2.0/ai/context/get-room-skill | Get room skill
[**AiContextGetRoomSkills**](AIContextAPI.md#AiContextGetRoomSkills) | **Get** /api/2.0/ai/context/get-room-skills | Get room skills



## AiContextGetContextFolders

> []AiContextFolder AiContextGetContextFolders(ctx).Execute()

Get context folders



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-context-get-context-folders/).

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
	resp, r, err := apiClient.AIContextAPI.AiContextGetContextFolders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIContextAPI.AiContextGetContextFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiContextGetContextFolders`: []AiContextFolder
	fmt.Fprintf(os.Stdout, "Response from `AIContextAPI.AiContextGetContextFolders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiContextGetContextFoldersRequest struct via the builder pattern


### Return type

[**[]AiContextFolder**](AiContextFolder.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiContextGetKnowledgeFolders

> []AiContextFolder AiContextGetKnowledgeFolders(ctx).Execute()

Get knowledge folders



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-context-get-knowledge-folders/).

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
	resp, r, err := apiClient.AIContextAPI.AiContextGetKnowledgeFolders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIContextAPI.AiContextGetKnowledgeFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiContextGetKnowledgeFolders`: []AiContextFolder
	fmt.Fprintf(os.Stdout, "Response from `AIContextAPI.AiContextGetKnowledgeFolders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiContextGetKnowledgeFoldersRequest struct via the builder pattern


### Return type

[**[]AiContextFolder**](AiContextFolder.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiContextGetRoomSkill

> string AiContextGetRoomSkill(ctx).RoomId(roomId).SkillId(skillId).Cloud(cloud).Execute()

Get room skill



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-context-get-room-skill/).

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
	roomId := "roomId_example" // string | The Instructions room identifier.
	skillId := "skillId_example" // string | The skill identifier, as `GET api/2.0/ai/context/get-room-skills` lists it.
	cloud := "cloud_example" // string | The account the room belongs to. Only `docspace`, this portal, exists, and the value is not read; omit it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIContextAPI.AiContextGetRoomSkill(context.Background()).RoomId(roomId).SkillId(skillId).Cloud(cloud).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIContextAPI.AiContextGetRoomSkill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiContextGetRoomSkill`: string
	fmt.Fprintf(os.Stdout, "Response from `AIContextAPI.AiContextGetRoomSkill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiContextGetRoomSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roomId** | **string** | The Instructions room identifier. | 
 **skillId** | **string** | The skill identifier, as `GET api/2.0/ai/context/get-room-skills` lists it. | 
 **cloud** | **string** | The account the room belongs to. Only `docspace`, this portal, exists, and the value is not read; omit it. | 

### Return type

**string**

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiContextGetRoomSkills

> []AiRoomSkill AiContextGetRoomSkills(ctx).RoomId(roomId).Cloud(cloud).Execute()

Get room skills



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-context-get-room-skills/).

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
	roomId := "roomId_example" // string | The Instructions room identifier.
	cloud := "cloud_example" // string | The account the room belongs to. Only `docspace`, this portal, exists, and the value is not read; omit it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIContextAPI.AiContextGetRoomSkills(context.Background()).RoomId(roomId).Cloud(cloud).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIContextAPI.AiContextGetRoomSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiContextGetRoomSkills`: []AiRoomSkill
	fmt.Fprintf(os.Stdout, "Response from `AIContextAPI.AiContextGetRoomSkills`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiContextGetRoomSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roomId** | **string** | The Instructions room identifier. | 
 **cloud** | **string** | The account the room belongs to. Only `docspace`, this portal, exists, and the value is not read; omit it. | 

### Return type

[**[]AiRoomSkill**](AiRoomSkill.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

