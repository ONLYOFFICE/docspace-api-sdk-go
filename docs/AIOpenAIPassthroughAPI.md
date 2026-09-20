# \AIOpenAIPassthroughAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiOpenaiChatCompletions**](AIOpenAIPassthroughAPI.md#AiOpenaiChatCompletions) | **Post** /api/2.0/ai/openai/{profileId}/v1/chat/completions | OpenAI chat completions passthrough
[**AiOpenaiImagesGenerations**](AIOpenAIPassthroughAPI.md#AiOpenaiImagesGenerations) | **Post** /api/2.0/ai/openai/{profileId}/v1/images/generations | OpenAI image generation passthrough



## AiOpenaiChatCompletions

> map[string]*interface{} AiOpenaiChatCompletions(ctx, profileId).RequestBody(requestBody).Execute()

OpenAI chat completions passthrough



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-openai-chat-completions/).

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
	profileId := "00000000-0000-0000-0000-000000000000" // string | The AI provider profile identifier.
	requestBody := map[string]*interface{}{"key": interface{}(123)} // map[string]*interface{} | An OpenAI Chat Completions request, forwarded to the provider byte for byte. The shape is the provider's, not this API's, so consult the provider's own reference; the model and the credentials come from the profile in the path and must not be sent here.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIOpenAIPassthroughAPI.AiOpenaiChatCompletions(context.Background(), profileId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIOpenAIPassthroughAPI.AiOpenaiChatCompletions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiOpenaiChatCompletions`: map[string]*interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIOpenAIPassthroughAPI.AiOpenaiChatCompletions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The AI provider profile identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiOpenaiChatCompletionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** | An OpenAI Chat Completions request, forwarded to the provider byte for byte. The shape is the provider's, not this API's, so consult the provider's own reference; the model and the credentials come from the profile in the path and must not be sent here. | 

### Return type

**map[string]*interface{}**

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiOpenaiImagesGenerations

> map[string]*interface{} AiOpenaiImagesGenerations(ctx, profileId).RequestBody(requestBody).Execute()

OpenAI image generation passthrough



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-openai-images-generations/).

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
	profileId := "00000000-0000-0000-0000-000000000000" // string | The AI provider profile identifier.
	requestBody := map[string]*interface{}{"key": interface{}(123)} // map[string]*interface{} | An OpenAI image-generation request, forwarded to the provider byte for byte. The shape is the provider's, not this API's, and the credentials come from the profile in the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIOpenAIPassthroughAPI.AiOpenaiImagesGenerations(context.Background(), profileId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIOpenAIPassthroughAPI.AiOpenaiImagesGenerations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiOpenaiImagesGenerations`: map[string]*interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIOpenAIPassthroughAPI.AiOpenaiImagesGenerations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | The AI provider profile identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiOpenaiImagesGenerationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** | An OpenAI image-generation request, forwarded to the provider byte for byte. The shape is the provider's, not this API's, and the credentials come from the profile in the path. | 

### Return type

**map[string]*interface{}**

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

