# \AIVectorizationAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiVectorizationStartTask**](AIVectorizationAPI.md#AiVectorizationStartTask) | **Post** /api/2.0/ai/vectorization/tasks | Start a vectorization task



## AiVectorizationStartTask

> AiVectorizationStartTask200Response AiVectorizationStartTask(ctx).AiVectorizationStartTaskRequest(aiVectorizationStartTaskRequest).Execute()

Start a vectorization task



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-vectorization-start-task/).

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
	aiVectorizationStartTaskRequest := *openapiclient.NewAiVectorizationStartTaskRequest([]int32{int32(123)}) // AiVectorizationStartTaskRequest | The files to index, proxied unchanged to the DocSpace AI service, which owns and validates the shape.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIVectorizationAPI.AiVectorizationStartTask(context.Background()).AiVectorizationStartTaskRequest(aiVectorizationStartTaskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIVectorizationAPI.AiVectorizationStartTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiVectorizationStartTask`: AiVectorizationStartTask200Response
	fmt.Fprintf(os.Stdout, "Response from `AIVectorizationAPI.AiVectorizationStartTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiVectorizationStartTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiVectorizationStartTaskRequest** | [**AiVectorizationStartTaskRequest**](AiVectorizationStartTaskRequest.md) | The files to index, proxied unchanged to the DocSpace AI service, which owns and validates the shape. | 

### Return type

[**AiVectorizationStartTask200Response**](AiVectorizationStartTask200Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

