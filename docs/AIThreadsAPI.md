# \AIThreadsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiThreadsAppendUserMessage**](AIThreadsAPI.md#AiThreadsAppendUserMessage) | **Post** /api/2.0/ai/threads/append-user-message | Append user message
[**AiThreadsClearMessages**](AIThreadsAPI.md#AiThreadsClearMessages) | **Delete** /api/2.0/ai/threads/clear-messages | Clear messages
[**AiThreadsCreate**](AIThreadsAPI.md#AiThreadsCreate) | **Post** /api/2.0/ai/threads/create | Create a chat thread
[**AiThreadsDelete**](AIThreadsAPI.md#AiThreadsDelete) | **Delete** /api/2.0/ai/threads/delete | Delete a chat thread
[**AiThreadsDeleteMessage**](AIThreadsAPI.md#AiThreadsDeleteMessage) | **Delete** /api/2.0/ai/threads/delete-message | Delete message
[**AiThreadsGetById**](AIThreadsAPI.md#AiThreadsGetById) | **Get** /api/2.0/ai/threads/get-by-id | Get a chat thread
[**AiThreadsGetMessageById**](AIThreadsAPI.md#AiThreadsGetMessageById) | **Get** /api/2.0/ai/threads/get-message-by-id | Get one chat message
[**AiThreadsList**](AIThreadsAPI.md#AiThreadsList) | **Get** /api/2.0/ai/threads/list | List chat threads
[**AiThreadsOpenOrCreate**](AIThreadsAPI.md#AiThreadsOpenOrCreate) | **Post** /api/2.0/ai/threads/open-or-create | Open or create
[**AiThreadsReadMessages**](AIThreadsAPI.md#AiThreadsReadMessages) | **Get** /api/2.0/ai/threads/read-messages | Read messages
[**AiThreadsRegenerateTitle**](AIThreadsAPI.md#AiThreadsRegenerateTitle) | **Post** /api/2.0/ai/threads/regenerate-title | Regenerate title
[**AiThreadsRename**](AIThreadsAPI.md#AiThreadsRename) | **Put** /api/2.0/ai/threads/rename | Rename a chat thread
[**AiThreadsTouch**](AIThreadsAPI.md#AiThreadsTouch) | **Post** /api/2.0/ai/threads/touch | Bump a thread's activity
[**AiThreadsUpdateMessage**](AIThreadsAPI.md#AiThreadsUpdateMessage) | **Put** /api/2.0/ai/threads/update-message | Update message



## AiThreadsAppendUserMessage

> AiThreadsAppendUserMessage200Response AiThreadsAppendUserMessage(ctx).AiThreadsAppendUserMessageRequest(aiThreadsAppendUserMessageRequest).Execute()

Append user message



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-append-user-message/).

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
	aiThreadsAppendUserMessageRequest := *openapiclient.NewAiThreadsAppendUserMessageRequest("ThreadId_example", *openapiclient.NewAiThreadMessageLike("user", *openapiclient.NewAiThreadMessageLikeContent())) // AiThreadsAppendUserMessageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIThreadsAPI.AiThreadsAppendUserMessage(context.Background()).AiThreadsAppendUserMessageRequest(aiThreadsAppendUserMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIThreadsAPI.AiThreadsAppendUserMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiThreadsAppendUserMessage`: AiThreadsAppendUserMessage200Response
	fmt.Fprintf(os.Stdout, "Response from `AIThreadsAPI.AiThreadsAppendUserMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiThreadsAppendUserMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiThreadsAppendUserMessageRequest** | [**AiThreadsAppendUserMessageRequest**](AiThreadsAppendUserMessageRequest.md) |  | 

### Return type

[**AiThreadsAppendUserMessage200Response**](AiThreadsAppendUserMessage200Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiThreadsClearMessages

> AiSuccessResponse AiThreadsClearMessages(ctx).Body(body).Execute()

Clear messages



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-clear-messages/).

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
	body := "body_example" // string | The ID of the thread to empty, as a bare JSON string.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIThreadsAPI.AiThreadsClearMessages(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIThreadsAPI.AiThreadsClearMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiThreadsClearMessages`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIThreadsAPI.AiThreadsClearMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiThreadsClearMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | The ID of the thread to empty, as a bare JSON string. | 

### Return type

[**AiSuccessResponse**](AiSuccessResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiThreadsCreate

> AiThread AiThreadsCreate(ctx).AiThreadsCreateRequest(aiThreadsCreateRequest).Execute()

Create a chat thread



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-create/).

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
	aiThreadsCreateRequest := *openapiclient.NewAiThreadsCreateRequest("Title_example") // AiThreadsCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIThreadsAPI.AiThreadsCreate(context.Background()).AiThreadsCreateRequest(aiThreadsCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIThreadsAPI.AiThreadsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiThreadsCreate`: AiThread
	fmt.Fprintf(os.Stdout, "Response from `AIThreadsAPI.AiThreadsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiThreadsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiThreadsCreateRequest** | [**AiThreadsCreateRequest**](AiThreadsCreateRequest.md) |  | 

### Return type

[**AiThread**](AiThread.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiThreadsDelete

> AiSuccessResponse AiThreadsDelete(ctx).Body(body).Execute()

Delete a chat thread



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-delete/).

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
	body := "body_example" // string | The ID of the thread to delete, as a bare JSON string.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIThreadsAPI.AiThreadsDelete(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIThreadsAPI.AiThreadsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiThreadsDelete`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIThreadsAPI.AiThreadsDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiThreadsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | The ID of the thread to delete, as a bare JSON string. | 

### Return type

[**AiSuccessResponse**](AiSuccessResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiThreadsDeleteMessage

> AiSuccessResponse AiThreadsDeleteMessage(ctx).Body(body).Execute()

Delete message



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-delete-message/).

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
	body := "body_example" // string | The ID of the message to delete, as a bare JSON string.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIThreadsAPI.AiThreadsDeleteMessage(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIThreadsAPI.AiThreadsDeleteMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiThreadsDeleteMessage`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIThreadsAPI.AiThreadsDeleteMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiThreadsDeleteMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | The ID of the message to delete, as a bare JSON string. | 

### Return type

[**AiSuccessResponse**](AiSuccessResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiThreadsGetById

> AiThread AiThreadsGetById(ctx).ThreadId(threadId).Execute()

Get a chat thread



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-get-by-id/).

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
	threadId := "11111111-1111-1111-1111-111111111111" // string | The chat thread identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIThreadsAPI.AiThreadsGetById(context.Background()).ThreadId(threadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIThreadsAPI.AiThreadsGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiThreadsGetById`: AiThread
	fmt.Fprintf(os.Stdout, "Response from `AIThreadsAPI.AiThreadsGetById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiThreadsGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threadId** | **string** | The chat thread identifier. | 

### Return type

[**AiThread**](AiThread.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiThreadsGetMessageById

> AiThreadMessageLike AiThreadsGetMessageById(ctx).MessageId(messageId).Execute()

Get one chat message



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-get-message-by-id/).

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
	messageId := "22222222-2222-2222-2222-222222222222" // string | The globally unique chat message identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIThreadsAPI.AiThreadsGetMessageById(context.Background()).MessageId(messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIThreadsAPI.AiThreadsGetMessageById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiThreadsGetMessageById`: AiThreadMessageLike
	fmt.Fprintf(os.Stdout, "Response from `AIThreadsAPI.AiThreadsGetMessageById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiThreadsGetMessageByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageId** | **string** | The globally unique chat message identifier. | 

### Return type

[**AiThreadMessageLike**](AiThreadMessageLike.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiThreadsList

> []AiThread AiThreadsList(ctx).EntityId(entityId).Count(count).Cursor(cursor).Query(query).Execute()

List chat threads



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-list/).

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
	entityId := "1234" // string | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. (optional)
	count := int32(20) // int32 | The maximum number of items to return in one page. (optional)
	cursor := "{\"id\":\"11111111-1111-1111-1111-111111111111\",\"lastEditDate\":1767225600000}" // string | The keyset pagination cursor: the JSON-encoded sort key of the last item already received. Omit for the first page. (optional)
	query := "contract" // string | The full-text query the thread list is filtered by. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIThreadsAPI.AiThreadsList(context.Background()).EntityId(entityId).Count(count).Cursor(cursor).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIThreadsAPI.AiThreadsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiThreadsList`: []AiThread
	fmt.Fprintf(os.Stdout, "Response from `AIThreadsAPI.AiThreadsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiThreadsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string** | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. | 
 **count** | **int32** | The maximum number of items to return in one page. | 
 **cursor** | **string** | The keyset pagination cursor: the JSON-encoded sort key of the last item already received. Omit for the first page. | 
 **query** | **string** | The full-text query the thread list is filtered by. | 

### Return type

[**[]AiThread**](AiThread.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiThreadsOpenOrCreate

> AiOpenOrCreateResult AiThreadsOpenOrCreate(ctx).AiThreadsOpenOrCreateRequest(aiThreadsOpenOrCreateRequest).Execute()

Open or create



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-open-or-create/).

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
	aiThreadsOpenOrCreateRequest := *openapiclient.NewAiThreadsOpenOrCreateRequest(*openapiclient.NewAiProfile("00000000-0000-0000-0000-000000000000", "OpenAI GPT-4o", *openapiclient.NewAiProviderType(), "https://api.openai.com/v1", "gpt-4o"), "ProfileId_example", *openapiclient.NewAiThreadMessageLike("user", *openapiclient.NewAiThreadMessageLikeContent())) // AiThreadsOpenOrCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIThreadsAPI.AiThreadsOpenOrCreate(context.Background()).AiThreadsOpenOrCreateRequest(aiThreadsOpenOrCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIThreadsAPI.AiThreadsOpenOrCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiThreadsOpenOrCreate`: AiOpenOrCreateResult
	fmt.Fprintf(os.Stdout, "Response from `AIThreadsAPI.AiThreadsOpenOrCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiThreadsOpenOrCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiThreadsOpenOrCreateRequest** | [**AiThreadsOpenOrCreateRequest**](AiThreadsOpenOrCreateRequest.md) |  | 

### Return type

[**AiOpenOrCreateResult**](AiOpenOrCreateResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiThreadsReadMessages

> []AiThreadMessageLike AiThreadsReadMessages(ctx).ThreadId(threadId).Count(count).Cursor(cursor).Direction(direction).Execute()

Read messages



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-read-messages/).

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
	threadId := "11111111-1111-1111-1111-111111111111" // string | The chat thread identifier.
	count := int32(20) // int32 | The maximum number of items to return in one page. (optional)
	cursor := "{\"id\":\"11111111-1111-1111-1111-111111111111\",\"lastEditDate\":1767225600000}" // string | The keyset pagination cursor: the JSON-encoded sort key of the last item already received. Omit for the first page. (optional)
	direction := "desc" // string | The order the message page is read in. Only desc turns the read around and pages back from the newest message; omit for the forward read. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIThreadsAPI.AiThreadsReadMessages(context.Background()).ThreadId(threadId).Count(count).Cursor(cursor).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIThreadsAPI.AiThreadsReadMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiThreadsReadMessages`: []AiThreadMessageLike
	fmt.Fprintf(os.Stdout, "Response from `AIThreadsAPI.AiThreadsReadMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiThreadsReadMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threadId** | **string** | The chat thread identifier. | 
 **count** | **int32** | The maximum number of items to return in one page. | 
 **cursor** | **string** | The keyset pagination cursor: the JSON-encoded sort key of the last item already received. Omit for the first page. | 
 **direction** | **string** | The order the message page is read in. Only desc turns the read around and pages back from the newest message; omit for the forward read. | 

### Return type

[**[]AiThreadMessageLike**](AiThreadMessageLike.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiThreadsRegenerateTitle

> AiThreadsRegenerateTitle200Response AiThreadsRegenerateTitle(ctx).AiThreadsRegenerateTitleRequest(aiThreadsRegenerateTitleRequest).Execute()

Regenerate title



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-regenerate-title/).

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
	aiThreadsRegenerateTitleRequest := *openapiclient.NewAiThreadsRegenerateTitleRequest("ThreadId_example", *openapiclient.NewAiProfile("00000000-0000-0000-0000-000000000000", "OpenAI GPT-4o", *openapiclient.NewAiProviderType(), "https://api.openai.com/v1", "gpt-4o")) // AiThreadsRegenerateTitleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIThreadsAPI.AiThreadsRegenerateTitle(context.Background()).AiThreadsRegenerateTitleRequest(aiThreadsRegenerateTitleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIThreadsAPI.AiThreadsRegenerateTitle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiThreadsRegenerateTitle`: AiThreadsRegenerateTitle200Response
	fmt.Fprintf(os.Stdout, "Response from `AIThreadsAPI.AiThreadsRegenerateTitle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiThreadsRegenerateTitleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiThreadsRegenerateTitleRequest** | [**AiThreadsRegenerateTitleRequest**](AiThreadsRegenerateTitleRequest.md) |  | 

### Return type

[**AiThreadsRegenerateTitle200Response**](AiThreadsRegenerateTitle200Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiThreadsRename

> AiSuccessResponse AiThreadsRename(ctx).AiThreadsRenameRequest(aiThreadsRenameRequest).Execute()

Rename a chat thread



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-rename/).

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
	aiThreadsRenameRequest := *openapiclient.NewAiThreadsRenameRequest("ThreadId_example", "Title_example") // AiThreadsRenameRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIThreadsAPI.AiThreadsRename(context.Background()).AiThreadsRenameRequest(aiThreadsRenameRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIThreadsAPI.AiThreadsRename``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiThreadsRename`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIThreadsAPI.AiThreadsRename`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiThreadsRenameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiThreadsRenameRequest** | [**AiThreadsRenameRequest**](AiThreadsRenameRequest.md) |  | 

### Return type

[**AiSuccessResponse**](AiSuccessResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiThreadsTouch

> AiSuccessResponse AiThreadsTouch(ctx).AiThreadsTouchRequest(aiThreadsTouchRequest).Execute()

Bump a thread's activity



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-touch/).

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
	aiThreadsTouchRequest := *openapiclient.NewAiThreadsTouchRequest("ThreadId_example") // AiThreadsTouchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIThreadsAPI.AiThreadsTouch(context.Background()).AiThreadsTouchRequest(aiThreadsTouchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIThreadsAPI.AiThreadsTouch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiThreadsTouch`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIThreadsAPI.AiThreadsTouch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiThreadsTouchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiThreadsTouchRequest** | [**AiThreadsTouchRequest**](AiThreadsTouchRequest.md) |  | 

### Return type

[**AiSuccessResponse**](AiSuccessResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiThreadsUpdateMessage

> AiSuccessResponse AiThreadsUpdateMessage(ctx).AiThreadsUpdateMessageRequest(aiThreadsUpdateMessageRequest).Execute()

Update message



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-threads-update-message/).

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
	aiThreadsUpdateMessageRequest := *openapiclient.NewAiThreadsUpdateMessageRequest("MessageId_example", *openapiclient.NewAiThreadMessageLike("user", *openapiclient.NewAiThreadMessageLikeContent())) // AiThreadsUpdateMessageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIThreadsAPI.AiThreadsUpdateMessage(context.Background()).AiThreadsUpdateMessageRequest(aiThreadsUpdateMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIThreadsAPI.AiThreadsUpdateMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiThreadsUpdateMessage`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIThreadsAPI.AiThreadsUpdateMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiThreadsUpdateMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiThreadsUpdateMessageRequest** | [**AiThreadsUpdateMessageRequest**](AiThreadsUpdateMessageRequest.md) |  | 

### Return type

[**AiSuccessResponse**](AiSuccessResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

