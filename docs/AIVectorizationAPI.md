# \AIVectorizationAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StartTask**](AIVectorizationAPI.md#StartTask) | **Post** /api/2.0/ai/vectorization/tasks | Start a vectorization task



## StartTask

> StartTask(ctx).VectorizationStartRequestBody(vectorizationStartRequestBody).Execute()

Start a vectorization task



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/start-task/).

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
	vectorizationStartRequestBody := *openapiclient.NewVectorizationStartRequestBody([]int32{int32(123)}) // VectorizationStartRequestBody | The vectorization parameters including file identifiers.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIVectorizationAPI.StartTask(context.Background()).VectorizationStartRequestBody(vectorizationStartRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIVectorizationAPI.StartTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vectorizationStartRequestBody** | [**VectorizationStartRequestBody**](VectorizationStartRequestBody.md) | The vectorization parameters including file identifiers. | 

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

