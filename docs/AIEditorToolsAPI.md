# \AIEditorToolsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiEditorToolsCall**](AIEditorToolsAPI.md#AiEditorToolsCall) | **Post** /api/2.0/ai/editor-tools/call | Execute a DocSpace tool on behalf of the editor AI plugin
[**AiEditorToolsList**](AIEditorToolsAPI.md#AiEditorToolsList) | **Get** /api/2.0/ai/editor-tools/list | Sanitized DocSpace tool catalog for the editor AI plugin



## AiEditorToolsCall

> AiSuccessResponse AiEditorToolsCall(ctx).RequestBody(requestBody).Execute()

Execute a DocSpace tool on behalf of the editor AI plugin



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
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIEditorToolsAPI.AiEditorToolsCall(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIEditorToolsAPI.AiEditorToolsCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiEditorToolsCall`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIEditorToolsAPI.AiEditorToolsCall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiEditorToolsCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**AiSuccessResponse**](AiSuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiEditorToolsList

> AiSuccessResponse AiEditorToolsList(ctx).Execute()

Sanitized DocSpace tool catalog for the editor AI plugin



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
	// response from `AiEditorToolsList`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIEditorToolsAPI.AiEditorToolsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiEditorToolsListRequest struct via the builder pattern


### Return type

[**AiSuccessResponse**](AiSuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

