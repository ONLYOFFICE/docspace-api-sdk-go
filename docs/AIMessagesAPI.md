# \AIMessagesAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportMessage**](AIMessagesAPI.md#ExportMessage) | **Post** /api/2.0/ai/messages/{messageId}/export | Export a single AI message to a document



## ExportMessage

> ExportMessage(ctx, messageId).ExportMessageRequestBody(exportMessageRequestBody).Execute()

Export a single AI message to a document



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/export-message/).

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
	messageId := int32(1) // int32 | The unique identifier of the AI chat message to export.
	exportMessageRequestBody := *openapiclient.NewExportMessageRequestBody(openapiclient.ExportChatRequestBody_folderId{Int32: new(int32)}, "Message Export") // ExportMessageRequestBody | The export parameters including destination folder and file title.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIMessagesAPI.ExportMessage(context.Background(), messageId).ExportMessageRequestBody(exportMessageRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIMessagesAPI.ExportMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **int32** | The unique identifier of the AI chat message to export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportMessageRequestBody** | [**ExportMessageRequestBody**](ExportMessageRequestBody.md) | The export parameters including destination folder and file title. | 

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

