# \AIContextAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiContextGetContextFolders**](AIContextAPI.md#AiContextGetContextFolders) | **Get** /api/2.0/ai/context/get-context-folders | Get context folders
[**AiContextGetRoomSkill**](AIContextAPI.md#AiContextGetRoomSkill) | **Get** /api/2.0/ai/context/get-room-skill | Get room skill
[**AiContextGetRoomSkills**](AIContextAPI.md#AiContextGetRoomSkills) | **Get** /api/2.0/ai/context/get-room-skills | Get room skills



## AiContextGetContextFolders

> AiSuccessResponse AiContextGetContextFolders(ctx).Execute()

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
	// response from `AiContextGetContextFolders`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIContextAPI.AiContextGetContextFolders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiContextGetContextFoldersRequest struct via the builder pattern


### Return type

[**AiSuccessResponse**](AiSuccessResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiContextGetRoomSkill

> AiSuccessResponse AiContextGetRoomSkill(ctx).Cloud(cloud).RoomId(roomId).SkillId(skillId).Execute()

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
	cloud := "cloud_example" // string | 
	roomId := "roomId_example" // string | 
	skillId := "skillId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIContextAPI.AiContextGetRoomSkill(context.Background()).Cloud(cloud).RoomId(roomId).SkillId(skillId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIContextAPI.AiContextGetRoomSkill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiContextGetRoomSkill`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIContextAPI.AiContextGetRoomSkill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiContextGetRoomSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud** | **string** |  | 
 **roomId** | **string** |  | 
 **skillId** | **string** |  | 

### Return type

[**AiSuccessResponse**](AiSuccessResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiContextGetRoomSkills

> AiSuccessResponse AiContextGetRoomSkills(ctx).Cloud(cloud).RoomId(roomId).Execute()

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
	cloud := "cloud_example" // string | 
	roomId := "roomId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIContextAPI.AiContextGetRoomSkills(context.Background()).Cloud(cloud).RoomId(roomId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIContextAPI.AiContextGetRoomSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiContextGetRoomSkills`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIContextAPI.AiContextGetRoomSkills`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiContextGetRoomSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud** | **string** |  | 
 **roomId** | **string** |  | 

### Return type

[**AiSuccessResponse**](AiSuccessResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

