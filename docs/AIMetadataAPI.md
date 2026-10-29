# \AIMetadataAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiMetaFill**](AIMetadataAPI.md#AiMetaFill) | **Post** /api/2.0/ai/meta/fill | Fill metadata values
[**AiMetaGenerate**](AIMetadataAPI.md#AiMetaGenerate) | **Post** /api/2.0/ai/meta/generate | Generate metadata fields



## AiMetaFill

> AiMetaFill200Response AiMetaFill(ctx).AiMetaFillRequest(aiMetaFillRequest).Execute()

Fill metadata values



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-meta-fill/).

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
	aiMetaFillRequest := *openapiclient.NewAiMetaFillRequest("Kind_example", *openapiclient.NewAiMetaGenerateRequestId()) // AiMetaFillRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIMetadataAPI.AiMetaFill(context.Background()).AiMetaFillRequest(aiMetaFillRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIMetadataAPI.AiMetaFill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiMetaFill`: AiMetaFill200Response
	fmt.Fprintf(os.Stdout, "Response from `AIMetadataAPI.AiMetaFill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiMetaFillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiMetaFillRequest** | [**AiMetaFillRequest**](AiMetaFillRequest.md) |  | 

### Return type

[**AiMetaFill200Response**](AiMetaFill200Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiMetaGenerate

> AiMetaGenerate200Response AiMetaGenerate(ctx).AiMetaGenerateRequest(aiMetaGenerateRequest).Execute()

Generate metadata fields



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-meta-generate/).

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
	aiMetaGenerateRequest := *openapiclient.NewAiMetaGenerateRequest("Kind_example", *openapiclient.NewAiMetaGenerateRequestId()) // AiMetaGenerateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIMetadataAPI.AiMetaGenerate(context.Background()).AiMetaGenerateRequest(aiMetaGenerateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIMetadataAPI.AiMetaGenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiMetaGenerate`: AiMetaGenerate200Response
	fmt.Fprintf(os.Stdout, "Response from `AIMetadataAPI.AiMetaGenerate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiMetaGenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiMetaGenerateRequest** | [**AiMetaGenerateRequest**](AiMetaGenerateRequest.md) |  | 

### Return type

[**AiMetaGenerate200Response**](AiMetaGenerate200Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

