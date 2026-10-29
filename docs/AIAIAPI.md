# \AIAIAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiApproveToolCall**](AIAIAPI.md#AiApproveToolCall) | **Post** /api/2.0/ai/ai/approve-tool-call | Approve tool call
[**AiDenyToolCall**](AIAIAPI.md#AiDenyToolCall) | **Post** /api/2.0/ai/ai/deny-tool-call | Deny tool call
[**AiRegenerateStream**](AIAIAPI.md#AiRegenerateStream) | **Post** /api/2.0/ai/ai/regenerate-stream | Regenerate stream
[**AiSend**](AIAIAPI.md#AiSend) | **Post** /api/2.0/ai/ai/send | Run an AI action
[**AiSendCustom**](AIAIAPI.md#AiSendCustom) | **Post** /api/2.0/ai/ai/send-custom | Send custom
[**AiSendWithStream**](AIAIAPI.md#AiSendWithStream) | **Post** /api/2.0/ai/ai/send-with-stream | Send with stream
[**AiSendWithStreamOpenAI**](AIAIAPI.md#AiSendWithStreamOpenAI) | **Post** /api/2.0/ai/ai/send-with-stream-openai | Stream a chat in OpenAI format



## AiApproveToolCall

> AiChatEvent AiApproveToolCall(ctx).AiApproveToolCallRequest(aiApproveToolCallRequest).Execute()

Approve tool call



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-approve-tool-call/).

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
	aiApproveToolCallRequest := *openapiclient.NewAiApproveToolCallRequest(interface{}(123), "ThreadId_example", "MessageId_example", float32(123)) // AiApproveToolCallRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAIAPI.AiApproveToolCall(context.Background()).AiApproveToolCallRequest(aiApproveToolCallRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAIAPI.AiApproveToolCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiApproveToolCall`: AiChatEvent
	fmt.Fprintf(os.Stdout, "Response from `AIAIAPI.AiApproveToolCall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiApproveToolCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiApproveToolCallRequest** | [**AiApproveToolCallRequest**](AiApproveToolCallRequest.md) |  | 

### Return type

[**AiChatEvent**](AiChatEvent.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/x-ndjson, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiDenyToolCall

> AiChatEvent AiDenyToolCall(ctx).AiToolCallData(aiToolCallData).Execute()

Deny tool call



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-deny-tool-call/).

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
	aiToolCallData := *openapiclient.NewAiToolCallData("11111111-1111-1111-1111-111111111111", "22222222-2222-2222-2222-222222222222", float32(0)) // AiToolCallData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAIAPI.AiDenyToolCall(context.Background()).AiToolCallData(aiToolCallData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAIAPI.AiDenyToolCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiDenyToolCall`: AiChatEvent
	fmt.Fprintf(os.Stdout, "Response from `AIAIAPI.AiDenyToolCall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiDenyToolCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiToolCallData** | [**AiToolCallData**](AiToolCallData.md) |  | 

### Return type

[**AiChatEvent**](AiChatEvent.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/x-ndjson, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiRegenerateStream

> AiChatEvent AiRegenerateStream(ctx).AiRegenerateStreamRequest(aiRegenerateStreamRequest).Execute()

Regenerate stream



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-regenerate-stream/).

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
	aiRegenerateStreamRequest := *openapiclient.NewAiRegenerateStreamRequest("ThreadId_example") // AiRegenerateStreamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAIAPI.AiRegenerateStream(context.Background()).AiRegenerateStreamRequest(aiRegenerateStreamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAIAPI.AiRegenerateStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiRegenerateStream`: AiChatEvent
	fmt.Fprintf(os.Stdout, "Response from `AIAIAPI.AiRegenerateStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiRegenerateStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiRegenerateStreamRequest** | [**AiRegenerateStreamRequest**](AiRegenerateStreamRequest.md) |  | 

### Return type

[**AiChatEvent**](AiChatEvent.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/x-ndjson, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiSend

> AiThreadMessageLike AiSend(ctx).AiSendRequest(aiSendRequest).Execute()

Run an AI action



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-send/).

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
	aiSendRequest := *openapiclient.NewAiSendRequest(openapiclient.AiActionType("Default"), *openapiclient.NewAiThreadMessageLike("user", *openapiclient.NewAiThreadMessageLikeContent())) // AiSendRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAIAPI.AiSend(context.Background()).AiSendRequest(aiSendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAIAPI.AiSend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiSend`: AiThreadMessageLike
	fmt.Fprintf(os.Stdout, "Response from `AIAIAPI.AiSend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiSendRequest** | [**AiSendRequest**](AiSendRequest.md) |  | 

### Return type

[**AiThreadMessageLike**](AiThreadMessageLike.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiSendCustom

> AiThreadMessageLike AiSendCustom(ctx).AiSendCustomRequest(aiSendCustomRequest).Execute()

Send custom



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-send-custom/).

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
	aiSendCustomRequest := *openapiclient.NewAiSendCustomRequest(false, "SystemPrompt_example", *openapiclient.NewAiThreadMessageLike("user", *openapiclient.NewAiThreadMessageLikeContent())) // AiSendCustomRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAIAPI.AiSendCustom(context.Background()).AiSendCustomRequest(aiSendCustomRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAIAPI.AiSendCustom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiSendCustom`: AiThreadMessageLike
	fmt.Fprintf(os.Stdout, "Response from `AIAIAPI.AiSendCustom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiSendCustomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiSendCustomRequest** | [**AiSendCustomRequest**](AiSendCustomRequest.md) |  | 

### Return type

[**AiThreadMessageLike**](AiThreadMessageLike.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiSendWithStream

> AiChatEvent AiSendWithStream(ctx).AiSendStreamBody(aiSendStreamBody).Execute()

Send with stream



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-send-with-stream/).

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
	aiSendStreamBody := *openapiclient.NewAiSendStreamBody(*openapiclient.NewAiThreadMessageLike("user", *openapiclient.NewAiThreadMessageLikeContent())) // AiSendStreamBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAIAPI.AiSendWithStream(context.Background()).AiSendStreamBody(aiSendStreamBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAIAPI.AiSendWithStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiSendWithStream`: AiChatEvent
	fmt.Fprintf(os.Stdout, "Response from `AIAIAPI.AiSendWithStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiSendWithStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiSendStreamBody** | [**AiSendStreamBody**](AiSendStreamBody.md) |  | 

### Return type

[**AiChatEvent**](AiChatEvent.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/x-ndjson, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiSendWithStreamOpenAI

> AiOpenAIStreamChunk AiSendWithStreamOpenAI(ctx).AiSendStreamBody(aiSendStreamBody).Execute()

Stream a chat in OpenAI format



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-send-with-stream-open-ai/).

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
	aiSendStreamBody := *openapiclient.NewAiSendStreamBody(*openapiclient.NewAiThreadMessageLike("user", *openapiclient.NewAiThreadMessageLikeContent())) // AiSendStreamBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAIAPI.AiSendWithStreamOpenAI(context.Background()).AiSendStreamBody(aiSendStreamBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAIAPI.AiSendWithStreamOpenAI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiSendWithStreamOpenAI`: AiOpenAIStreamChunk
	fmt.Fprintf(os.Stdout, "Response from `AIAIAPI.AiSendWithStreamOpenAI`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiSendWithStreamOpenAIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiSendStreamBody** | [**AiSendStreamBody**](AiSendStreamBody.md) |  | 

### Return type

[**AiOpenAIStreamChunk**](AiOpenAIStreamChunk.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

