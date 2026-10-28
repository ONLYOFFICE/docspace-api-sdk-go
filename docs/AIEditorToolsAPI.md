# \AIEditorToolsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiEditorToolsCall**](AIEditorToolsAPI.md#AiEditorToolsCall) | **Post** /api/2.0/ai/editor-tools/call | Call an editor tool
[**AiEditorToolsList**](AIEditorToolsAPI.md#AiEditorToolsList) | **Get** /api/2.0/ai/editor-tools/list | List editor tools



## AiEditorToolsCall

> AiEditorToolsCall200Response AiEditorToolsCall(ctx).AiEditorToolsCallRequest(aiEditorToolsCallRequest).Execute()

Call an editor tool



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-editor-tools-call/).

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
	aiEditorToolsCallRequest := *openapiclient.NewAiEditorToolsCallRequest("docspace_get_folder") // AiEditorToolsCallRequest | The tool to run: `name` from `GET api/2.0/ai/editor-tools/list`, `arguments` matching that tool's input schema, and an optional `entityId` for the room to run it in.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIEditorToolsAPI.AiEditorToolsCall(context.Background()).AiEditorToolsCallRequest(aiEditorToolsCallRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIEditorToolsAPI.AiEditorToolsCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiEditorToolsCall`: AiEditorToolsCall200Response
	fmt.Fprintf(os.Stdout, "Response from `AIEditorToolsAPI.AiEditorToolsCall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiEditorToolsCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiEditorToolsCallRequest** | [**AiEditorToolsCallRequest**](AiEditorToolsCallRequest.md) | The tool to run: `name` from `GET api/2.0/ai/editor-tools/list`, `arguments` matching that tool's input schema, and an optional `entityId` for the room to run it in. | 

### Return type

[**AiEditorToolsCall200Response**](AiEditorToolsCall200Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiEditorToolsList

> AiEditorToolsList200Response AiEditorToolsList(ctx).Execute()

List editor tools



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-editor-tools-list/).

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
	resp, r, err := apiClient.AIEditorToolsAPI.AiEditorToolsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIEditorToolsAPI.AiEditorToolsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiEditorToolsList`: AiEditorToolsList200Response
	fmt.Fprintf(os.Stdout, "Response from `AIEditorToolsAPI.AiEditorToolsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiEditorToolsListRequest struct via the builder pattern


### Return type

[**AiEditorToolsList200Response**](AiEditorToolsList200Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

