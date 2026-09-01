# \AIOpenAIPassthroughAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiOpenaiChatCompletions**](AIOpenAIPassthroughAPI.md#AiOpenaiChatCompletions) | **Post** /api/2.0/ai/openai/{profileId}/v1/chat/completions | OpenAI-compatible chat completions proxied to the profile's provider
[**AiOpenaiImagesGenerations**](AIOpenAIPassthroughAPI.md#AiOpenaiImagesGenerations) | **Post** /api/2.0/ai/openai/{profileId}/v1/images/generations | OpenAI-compatible image generation proxied to the profile's provider



## AiOpenaiChatCompletions

> AiSuccessResponse AiOpenaiChatCompletions(ctx, profileId).RequestBody(requestBody).Execute()

OpenAI-compatible chat completions proxied to the profile's provider



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
	profileId := "profileId_example" // string | The AI provider profile identifier.
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIOpenAIPassthroughAPI.AiOpenaiChatCompletions(context.Background(), profileId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIOpenAIPassthroughAPI.AiOpenaiChatCompletions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiOpenaiChatCompletions`: AiSuccessResponse
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


## AiOpenaiImagesGenerations

> AiSuccessResponse AiOpenaiImagesGenerations(ctx, profileId).RequestBody(requestBody).Execute()

OpenAI-compatible image generation proxied to the profile's provider



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
	profileId := "profileId_example" // string | The AI provider profile identifier.
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIOpenAIPassthroughAPI.AiOpenaiImagesGenerations(context.Background(), profileId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIOpenAIPassthroughAPI.AiOpenaiImagesGenerations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiOpenaiImagesGenerations`: AiSuccessResponse
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

