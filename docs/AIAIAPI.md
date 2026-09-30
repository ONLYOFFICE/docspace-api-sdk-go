# \AIAIAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiAiApproveToolCall**](AIAIAPI.md#AiAiApproveToolCall) | **Post** /api/2.0/ai/ai/approve-tool-call | Approve tool call
[**AiAiDenyToolCall**](AIAIAPI.md#AiAiDenyToolCall) | **Post** /api/2.0/ai/ai/deny-tool-call | Deny tool call
[**AiAiRegenerateStream**](AIAIAPI.md#AiAiRegenerateStream) | **Post** /api/2.0/ai/ai/regenerate-stream | Regenerate stream
[**AiAiSend**](AIAIAPI.md#AiAiSend) | **Post** /api/2.0/ai/ai/send | Run an AI action
[**AiAiSendCustom**](AIAIAPI.md#AiAiSendCustom) | **Post** /api/2.0/ai/ai/send-custom | Send custom
[**AiAiSendWithStream**](AIAIAPI.md#AiAiSendWithStream) | **Post** /api/2.0/ai/ai/send-with-stream | Send with stream
[**AiAiSendWithStreamOpenAI**](AIAIAPI.md#AiAiSendWithStreamOpenAI) | **Post** /api/2.0/ai/ai/send-with-stream-openai | Stream a chat in OpenAI format



## AiAiApproveToolCall

> AiChatEvent AiAiApproveToolCall(ctx).AiAiApproveToolCallRequest(aiAiApproveToolCallRequest).Execute()

Approve tool call



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-ai-approve-tool-call/).

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
	aiAiApproveToolCallRequest := *openapiclient.NewAiAiApproveToolCallRequest(interface{}(123), "ThreadId_example", "MessageId_example", float32(123), *openapiclient.NewAiThreadMessageLike("user", *openapiclient.NewAiThreadMessageLikeContent())) // AiAiApproveToolCallRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAIAPI.AiAiApproveToolCall(context.Background()).AiAiApproveToolCallRequest(aiAiApproveToolCallRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAIAPI.AiAiApproveToolCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAiApproveToolCall`: AiChatEvent
	fmt.Fprintf(os.Stdout, "Response from `AIAIAPI.AiAiApproveToolCall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAiApproveToolCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiAiApproveToolCallRequest** | [**AiAiApproveToolCallRequest**](AiAiApproveToolCallRequest.md) |  | 

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


## AiAiDenyToolCall

> AiChatEvent AiAiDenyToolCall(ctx).AiAiToolCallData(aiAiToolCallData).Execute()

Deny tool call



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-ai-deny-tool-call/).

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
	aiAiToolCallData := *openapiclient.NewAiAiToolCallData("11111111-1111-1111-1111-111111111111", "22222222-2222-2222-2222-222222222222", float32(0), *openapiclient.NewAiThreadMessageLike("user", *openapiclient.NewAiThreadMessageLikeContent())) // AiAiToolCallData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAIAPI.AiAiDenyToolCall(context.Background()).AiAiToolCallData(aiAiToolCallData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAIAPI.AiAiDenyToolCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAiDenyToolCall`: AiChatEvent
	fmt.Fprintf(os.Stdout, "Response from `AIAIAPI.AiAiDenyToolCall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAiDenyToolCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiAiToolCallData** | [**AiAiToolCallData**](AiAiToolCallData.md) |  | 

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


## AiAiRegenerateStream

> AiChatEvent AiAiRegenerateStream(ctx).AiAiRegenerateStreamRequest(aiAiRegenerateStreamRequest).Execute()

Regenerate stream



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-ai-regenerate-stream/).

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
	aiAiRegenerateStreamRequest := *openapiclient.NewAiAiRegenerateStreamRequest("ThreadId_example") // AiAiRegenerateStreamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAIAPI.AiAiRegenerateStream(context.Background()).AiAiRegenerateStreamRequest(aiAiRegenerateStreamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAIAPI.AiAiRegenerateStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAiRegenerateStream`: AiChatEvent
	fmt.Fprintf(os.Stdout, "Response from `AIAIAPI.AiAiRegenerateStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAiRegenerateStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiAiRegenerateStreamRequest** | [**AiAiRegenerateStreamRequest**](AiAiRegenerateStreamRequest.md) |  | 

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


## AiAiSend

> AiThreadMessageLike AiAiSend(ctx).AiAiSendRequest(aiAiSendRequest).Execute()

Run an AI action



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-ai-send/).

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
	aiAiSendRequest := *openapiclient.NewAiAiSendRequest(openapiclient.AiActionType("Default"), *openapiclient.NewAiThreadMessageLike("user", *openapiclient.NewAiThreadMessageLikeContent())) // AiAiSendRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAIAPI.AiAiSend(context.Background()).AiAiSendRequest(aiAiSendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAIAPI.AiAiSend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAiSend`: AiThreadMessageLike
	fmt.Fprintf(os.Stdout, "Response from `AIAIAPI.AiAiSend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAiSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiAiSendRequest** | [**AiAiSendRequest**](AiAiSendRequest.md) |  | 

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


## AiAiSendCustom

> AiThreadMessageLike AiAiSendCustom(ctx).AiAiSendCustomRequest(aiAiSendCustomRequest).Execute()

Send custom



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-ai-send-custom/).

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
	aiAiSendCustomRequest := *openapiclient.NewAiAiSendCustomRequest(false, "SystemPrompt_example", *openapiclient.NewAiThreadMessageLike("user", *openapiclient.NewAiThreadMessageLikeContent())) // AiAiSendCustomRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAIAPI.AiAiSendCustom(context.Background()).AiAiSendCustomRequest(aiAiSendCustomRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAIAPI.AiAiSendCustom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAiSendCustom`: AiThreadMessageLike
	fmt.Fprintf(os.Stdout, "Response from `AIAIAPI.AiAiSendCustom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAiSendCustomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiAiSendCustomRequest** | [**AiAiSendCustomRequest**](AiAiSendCustomRequest.md) |  | 

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


## AiAiSendWithStream

> AiChatEvent AiAiSendWithStream(ctx).AiAiSendStreamBody(aiAiSendStreamBody).Execute()

Send with stream



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-ai-send-with-stream/).

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
	aiAiSendStreamBody := *openapiclient.NewAiAiSendStreamBody(*openapiclient.NewAiThreadMessageLike("user", *openapiclient.NewAiThreadMessageLikeContent())) // AiAiSendStreamBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAIAPI.AiAiSendWithStream(context.Background()).AiAiSendStreamBody(aiAiSendStreamBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAIAPI.AiAiSendWithStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAiSendWithStream`: AiChatEvent
	fmt.Fprintf(os.Stdout, "Response from `AIAIAPI.AiAiSendWithStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAiSendWithStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiAiSendStreamBody** | [**AiAiSendStreamBody**](AiAiSendStreamBody.md) |  | 

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


## AiAiSendWithStreamOpenAI

> AiOpenAIStreamChunk AiAiSendWithStreamOpenAI(ctx).AiAiSendStreamBody(aiAiSendStreamBody).Execute()

Stream a chat in OpenAI format



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-ai-send-with-stream-open-ai/).

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
	aiAiSendStreamBody := *openapiclient.NewAiAiSendStreamBody(*openapiclient.NewAiThreadMessageLike("user", *openapiclient.NewAiThreadMessageLikeContent())) // AiAiSendStreamBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAIAPI.AiAiSendWithStreamOpenAI(context.Background()).AiAiSendStreamBody(aiAiSendStreamBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAIAPI.AiAiSendWithStreamOpenAI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAiSendWithStreamOpenAI`: AiOpenAIStreamChunk
	fmt.Fprintf(os.Stdout, "Response from `AIAIAPI.AiAiSendWithStreamOpenAI`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAiSendWithStreamOpenAIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiAiSendStreamBody** | [**AiAiSendStreamBody**](AiAiSendStreamBody.md) |  | 

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

