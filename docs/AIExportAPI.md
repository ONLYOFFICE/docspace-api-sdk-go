# \AIExportAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiExportTextToDocx**](AIExportAPI.md#AiExportTextToDocx) | **Post** /api/2.0/ai/text-to-docx | Start markdown → docx export



## AiExportTextToDocx

> AiExportTextToDocx202Response AiExportTextToDocx(ctx).AiExportTextToDocxRequest(aiExportTextToDocxRequest).Execute()

Start markdown → docx export



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-export-text-to-docx/).

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
	aiExportTextToDocxRequest := *openapiclient.NewAiExportTextToDocxRequest("Title_example", "Content_example", *openapiclient.NewAiExportTextToDocxRequestFolderId()) // AiExportTextToDocxRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIExportAPI.AiExportTextToDocx(context.Background()).AiExportTextToDocxRequest(aiExportTextToDocxRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIExportAPI.AiExportTextToDocx``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiExportTextToDocx`: AiExportTextToDocx202Response
	fmt.Fprintf(os.Stdout, "Response from `AIExportAPI.AiExportTextToDocx`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiExportTextToDocxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiExportTextToDocxRequest** | [**AiExportTextToDocxRequest**](AiExportTextToDocxRequest.md) |  | 

### Return type

[**AiExportTextToDocx202Response**](AiExportTextToDocx202Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

