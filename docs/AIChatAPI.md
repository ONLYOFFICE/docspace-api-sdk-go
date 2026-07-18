# \AIChatAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContinueChat**](AIChatAPI.md#ContinueChat) | **Post** /api/2.0/ai/chats/{chatId}/messages | Send a message to an existing AI chat
[**DeleteChat**](AIChatAPI.md#DeleteChat) | **Delete** /api/2.0/ai/chats/{chatId} | Delete an AI chat
[**ExportChat**](AIChatAPI.md#ExportChat) | **Post** /api/2.0/ai/chats/{chatId}/messages/export | Export AI chat messages to a file
[**GetChat**](AIChatAPI.md#GetChat) | **Get** /api/2.0/ai/chats/{chatId} | Get an AI chat by ID
[**GetChatModels**](AIChatAPI.md#GetChatModels) | **Get** /api/2.0/ai/chats/models | Get available AI models
[**GetChats**](AIChatAPI.md#GetChats) | **Get** /api/2.0/ai/rooms/{roomId}/chats | Get AI chats in a room
[**GetMessages**](AIChatAPI.md#GetMessages) | **Get** /api/2.0/ai/chats/{chatId}/messages | Get messages of an AI chat
[**GetUserChatsSettings**](AIChatAPI.md#GetUserChatsSettings) | **Get** /api/2.0/ai/rooms/{roomId}/chats/config | Get user chat settings for a room
[**ProvidePermission**](AIChatAPI.md#ProvidePermission) | **Post** /api/2.0/ai/chats/tool-permissions/{callId}/decision | Submit a tool execution permission decision
[**RenameChat**](AIChatAPI.md#RenameChat) | **Put** /api/2.0/ai/chats/{chatId} | Rename an AI chat
[**ResolveEditorTool**](AIChatAPI.md#ResolveEditorTool) | **Post** /api/2.0/ai/chats/tool-files/{callId}/decision | Resolve a pending editor file-generation tool
[**SetUserChatsSettings**](AIChatAPI.md#SetUserChatsSettings) | **Put** /api/2.0/ai/rooms/{roomId}/chats/config | Update user chat settings for a room
[**StartNewChat**](AIChatAPI.md#StartNewChat) | **Post** /api/2.0/ai/rooms/{roomId}/chats | Start a new AI chat



## ContinueChat

> ContinueChat(ctx, chatId).ContinueChatBody(continueChatBody).Execute()

Send a message to an existing AI chat



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/continue-chat/).

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
	chatId := "00000000-0000-0000-0000-000000000000" // string | The unique identifier of the existing AI chat session to continue.
	continueChatBody := *openapiclient.NewContinueChatBody("Summarize this document for me") // ContinueChatBody | The message and optional file attachments.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIChatAPI.ContinueChat(context.Background(), chatId).ContinueChatBody(continueChatBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIChatAPI.ContinueChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | The unique identifier of the existing AI chat session to continue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiContinueChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **continueChatBody** | [**ContinueChatBody**](ContinueChatBody.md) | The message and optional file attachments. | 

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


## DeleteChat

> DeleteChat(ctx, chatId).Execute()

Delete an AI chat



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-chat/).

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
	chatId := "00000000-0000-0000-0000-000000000000" // string | The unique identifier of the AI chat session to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIChatAPI.DeleteChat(context.Background(), chatId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIChatAPI.DeleteChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | The unique identifier of the AI chat session to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportChat

> ExportChat(ctx, chatId).ExportChatRequestBody(exportChatRequestBody).Execute()

Export AI chat messages to a file



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/export-chat/).

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
	chatId := "00000000-0000-0000-0000-000000000000" // string | The unique identifier of the AI chat session to export.
	exportChatRequestBody := *openapiclient.NewExportChatRequestBody(openapiclient.ExportChatRequestBody_folderId{Int32: new(int32)}, "Chat Export") // ExportChatRequestBody | The export parameters including destination folder and file title.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIChatAPI.ExportChat(context.Background(), chatId).ExportChatRequestBody(exportChatRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIChatAPI.ExportChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | The unique identifier of the AI chat session to export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportChatRequestBody** | [**ExportChatRequestBody**](ExportChatRequestBody.md) | The export parameters including destination folder and file title. | 

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


## GetChat

> ChatWrapper GetChat(ctx, chatId).Execute()

Get an AI chat by ID



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-chat/).

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
	chatId := "00000000-0000-0000-0000-000000000000" // string | The unique identifier of the AI chat session to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIChatAPI.GetChat(context.Background(), chatId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIChatAPI.GetChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChat`: ChatWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIChatAPI.GetChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | The unique identifier of the AI chat session to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatWrapper**](ChatWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChatModels

> ModelArrayWrapper GetChatModels(ctx).Provider(provider).Execute()

Get available AI models



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-chat-models/).

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
	provider := int32(1) // int32 | The optional AI provider identifier to filter models by. When set to 0, models from all providers are returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIChatAPI.GetChatModels(context.Background()).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIChatAPI.GetChatModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatModels`: ModelArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIChatAPI.GetChatModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChatModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **int32** | The optional AI provider identifier to filter models by. When set to 0, models from all providers are returned. | 

### Return type

[**ModelArrayWrapper**](ModelArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChats

> ChatArrayWrapper GetChats(ctx, roomId).StartIndex(startIndex).Count(count).Execute()

Get AI chats in a room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-chats/).

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
	roomId := int32(42) // int32 | The identifier of the room whose AI chat sessions are to be listed.
	startIndex := int32(0) // int32 | The number of items to skip before returning results (zero-based offset). Defaults to 0. (optional)
	count := int32(100) // int32 | The maximum number of items to return per page. Defaults to 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIChatAPI.GetChats(context.Background(), roomId).StartIndex(startIndex).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIChatAPI.GetChats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChats`: ChatArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIChatAPI.GetChats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **int32** | The identifier of the room whose AI chat sessions are to be listed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | The number of items to skip before returning results (zero-based offset). Defaults to 0. | 
 **count** | **int32** | The maximum number of items to return per page. Defaults to 100. | 

### Return type

[**ChatArrayWrapper**](ChatArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessages

> MessageArrayWrapper GetMessages(ctx, chatId).StartIndex(startIndex).Count(count).Execute()

Get messages of an AI chat



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-messages/).

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
	chatId := "00000000-0000-0000-0000-000000000000" // string | The unique identifier of the AI chat session whose messages are to be listed.
	startIndex := int32(0) // int32 | The number of items to skip before returning results (zero-based offset). Defaults to 0. (optional)
	count := int32(100) // int32 | The maximum number of items to return per page. Defaults to 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIChatAPI.GetMessages(context.Background(), chatId).StartIndex(startIndex).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIChatAPI.GetMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessages`: MessageArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIChatAPI.GetMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | The unique identifier of the AI chat session whose messages are to be listed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | The number of items to skip before returning results (zero-based offset). Defaults to 0. | 
 **count** | **int32** | The maximum number of items to return per page. Defaults to 100. | 

### Return type

[**MessageArrayWrapper**](MessageArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserChatsSettings

> UserChatSettingsWrapper GetUserChatsSettings(ctx, roomId).Execute()

Get user chat settings for a room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-user-chats-settings/).

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
	roomId := int32(42) // int32 | The identifier of the room whose chat settings are to be retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIChatAPI.GetUserChatsSettings(context.Background(), roomId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIChatAPI.GetUserChatsSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserChatsSettings`: UserChatSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIChatAPI.GetUserChatsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **int32** | The identifier of the room whose chat settings are to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserChatsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserChatSettingsWrapper**](UserChatSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvidePermission

> ProvidePermission(ctx, callId).ToolDecisionRequestBody(toolDecisionRequestBody).Execute()

Submit a tool execution permission decision



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/provide-permission/).

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
	callId := "call_abc123" // string | The unique identifier of the pending tool execution call awaiting a permission decision.
	toolDecisionRequestBody := *openapiclient.NewToolDecisionRequestBody() // ToolDecisionRequestBody | The permission decision parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIChatAPI.ProvidePermission(context.Background(), callId).ToolDecisionRequestBody(toolDecisionRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIChatAPI.ProvidePermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | The unique identifier of the pending tool execution call awaiting a permission decision. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvidePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **toolDecisionRequestBody** | [**ToolDecisionRequestBody**](ToolDecisionRequestBody.md) | The permission decision parameters. | 

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


## RenameChat

> ChatWrapper RenameChat(ctx, chatId).RenameChatBody(renameChatBody).Execute()

Rename an AI chat



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/rename-chat/).

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
	chatId := "00000000-0000-0000-0000-000000000000" // string | The unique identifier of the AI chat session to rename.
	renameChatBody := *openapiclient.NewRenameChatBody("Project Discussion") // RenameChatBody | The new chat name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIChatAPI.RenameChat(context.Background(), chatId).RenameChatBody(renameChatBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIChatAPI.RenameChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenameChat`: ChatWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIChatAPI.RenameChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chatId** | **string** | The unique identifier of the AI chat session to rename. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renameChatBody** | [**RenameChatBody**](RenameChatBody.md) | The new chat name. | 

### Return type

[**ChatWrapper**](ChatWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveEditorTool

> GeneratedFileWrapper ResolveEditorTool(ctx, callId).EditorToolDecisionRequestBody(editorToolDecisionRequestBody).Execute()

Resolve a pending editor file-generation tool



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/resolve-editor-tool/).

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
	callId := "call_abc123" // string | The unique identifier of the pending tool call awaiting the user's decision.
	editorToolDecisionRequestBody := *openapiclient.NewEditorToolDecisionRequestBody() // EditorToolDecisionRequestBody | The decision parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIChatAPI.ResolveEditorTool(context.Background(), callId).EditorToolDecisionRequestBody(editorToolDecisionRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIChatAPI.ResolveEditorTool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveEditorTool`: GeneratedFileWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIChatAPI.ResolveEditorTool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callId** | **string** | The unique identifier of the pending tool call awaiting the user's decision. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveEditorToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editorToolDecisionRequestBody** | [**EditorToolDecisionRequestBody**](EditorToolDecisionRequestBody.md) | The decision parameters. | 

### Return type

[**GeneratedFileWrapper**](GeneratedFileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetUserChatsSettings

> UserChatSettingsWrapper SetUserChatsSettings(ctx, roomId).SetUserChatSettingsRequestBody(setUserChatSettingsRequestBody).Execute()

Update user chat settings for a room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-user-chats-settings/).

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
	roomId := int32(42) // int32 | The identifier of the room whose chat settings are to be updated.
	setUserChatSettingsRequestBody := *openapiclient.NewSetUserChatSettingsRequestBody() // SetUserChatSettingsRequestBody | The chat settings to apply.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIChatAPI.SetUserChatsSettings(context.Background(), roomId).SetUserChatSettingsRequestBody(setUserChatSettingsRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIChatAPI.SetUserChatsSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetUserChatsSettings`: UserChatSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIChatAPI.SetUserChatsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **int32** | The identifier of the room whose chat settings are to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetUserChatsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setUserChatSettingsRequestBody** | [**SetUserChatSettingsRequestBody**](SetUserChatSettingsRequestBody.md) | The chat settings to apply. | 

### Return type

[**UserChatSettingsWrapper**](UserChatSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartNewChat

> StartNewChat(ctx, roomId).StartNewChatBody(startNewChatBody).Execute()

Start a new AI chat



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/start-new-chat/).

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
	roomId := int32(42) // int32 | The identifier of the room in which to create the new AI chat session.
	startNewChatBody := *openapiclient.NewStartNewChatBody("Hello, can you help me with this document?") // StartNewChatBody | The initial message and optional file attachments.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIChatAPI.StartNewChat(context.Background(), roomId).StartNewChatBody(startNewChatBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIChatAPI.StartNewChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **int32** | The identifier of the room in which to create the new AI chat session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartNewChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startNewChatBody** | [**StartNewChatBody**](StartNewChatBody.md) | The initial message and optional file attachments. | 

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

