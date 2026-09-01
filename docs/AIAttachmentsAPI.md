# \AIAttachmentsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiAttachmentsDelete**](AIAttachmentsAPI.md#AiAttachmentsDelete) | **Delete** /api/2.0/ai/attachments/delete | Delete
[**AiAttachmentsDeleteMany**](AIAttachmentsAPI.md#AiAttachmentsDeleteMany) | **Delete** /api/2.0/ai/attachments/delete-many | Delete many
[**AiAttachmentsGet**](AIAttachmentsAPI.md#AiAttachmentsGet) | **Post** /api/2.0/ai/attachments/get | Get
[**AiAttachmentsGetMany**](AIAttachmentsAPI.md#AiAttachmentsGetMany) | **Post** /api/2.0/ai/attachments/get-many | Get many
[**AiAttachmentsLinkToMessage**](AIAttachmentsAPI.md#AiAttachmentsLinkToMessage) | **Post** /api/2.0/ai/attachments/link-to-message | Link to message
[**AiAttachmentsSaveFile**](AIAttachmentsAPI.md#AiAttachmentsSaveFile) | **Post** /api/2.0/ai/attachments/save-file | Save file
[**AiAttachmentsSaveFilesMany**](AIAttachmentsAPI.md#AiAttachmentsSaveFilesMany) | **Post** /api/2.0/ai/attachments/save-files-many | Save files many



## AiAttachmentsDelete

> AiSuccessResponse AiAttachmentsDelete(ctx).Body(body).Execute()

Delete



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-attachments-delete/).

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
	body := "body_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAttachmentsAPI.AiAttachmentsDelete(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAttachmentsAPI.AiAttachmentsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAttachmentsDelete`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIAttachmentsAPI.AiAttachmentsDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAttachmentsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

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


## AiAttachmentsDeleteMany

> AiSuccessResponse AiAttachmentsDeleteMany(ctx).RequestBody(requestBody).Execute()

Delete many



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-attachments-delete-many/).

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
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAttachmentsAPI.AiAttachmentsDeleteMany(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAttachmentsAPI.AiAttachmentsDeleteMany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAttachmentsDeleteMany`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIAttachmentsAPI.AiAttachmentsDeleteMany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAttachmentsDeleteManyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

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


## AiAttachmentsGet

> AiAttachment AiAttachmentsGet(ctx).Body(body).Execute()

Get



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-attachments-get/).

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
	body := "body_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAttachmentsAPI.AiAttachmentsGet(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAttachmentsAPI.AiAttachmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAttachmentsGet`: AiAttachment
	fmt.Fprintf(os.Stdout, "Response from `AIAttachmentsAPI.AiAttachmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAttachmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

[**AiAttachment**](AiAttachment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAttachmentsGetMany

> []*AiAttachment AiAttachmentsGetMany(ctx).RequestBody(requestBody).Execute()

Get many



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-attachments-get-many/).

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
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAttachmentsAPI.AiAttachmentsGetMany(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAttachmentsAPI.AiAttachmentsGetMany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAttachmentsGetMany`: []*AiAttachment
	fmt.Fprintf(os.Stdout, "Response from `AIAttachmentsAPI.AiAttachmentsGetMany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAttachmentsGetManyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 

### Return type

[**[]*AiAttachment**](AiAttachment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAttachmentsLinkToMessage

> AiSuccessResponse AiAttachmentsLinkToMessage(ctx).AiAttachmentsLinkToMessageRequest(aiAttachmentsLinkToMessageRequest).Execute()

Link to message



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-attachments-link-to-message/).

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
	aiAttachmentsLinkToMessageRequest := *openapiclient.NewAiAttachmentsLinkToMessageRequest([]string{"Ids_example"}, "MessageId_example", "ThreadId_example") // AiAttachmentsLinkToMessageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAttachmentsAPI.AiAttachmentsLinkToMessage(context.Background()).AiAttachmentsLinkToMessageRequest(aiAttachmentsLinkToMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAttachmentsAPI.AiAttachmentsLinkToMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAttachmentsLinkToMessage`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIAttachmentsAPI.AiAttachmentsLinkToMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAttachmentsLinkToMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiAttachmentsLinkToMessageRequest** | [**AiAttachmentsLinkToMessageRequest**](AiAttachmentsLinkToMessageRequest.md) |  | 

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


## AiAttachmentsSaveFile

> AiAttachment AiAttachmentsSaveFile(ctx).AiAttachmentsSaveFileRequest(aiAttachmentsSaveFileRequest).Execute()

Save file



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-attachments-save-file/).

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
	aiAttachmentsSaveFileRequest := *openapiclient.NewAiAttachmentsSaveFileRequest(*openapiclient.NewAiAttachmentsSaveFileRequestInput("Path_example", "Content_example", float32(123))) // AiAttachmentsSaveFileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAttachmentsAPI.AiAttachmentsSaveFile(context.Background()).AiAttachmentsSaveFileRequest(aiAttachmentsSaveFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAttachmentsAPI.AiAttachmentsSaveFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAttachmentsSaveFile`: AiAttachment
	fmt.Fprintf(os.Stdout, "Response from `AIAttachmentsAPI.AiAttachmentsSaveFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAttachmentsSaveFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiAttachmentsSaveFileRequest** | [**AiAttachmentsSaveFileRequest**](AiAttachmentsSaveFileRequest.md) |  | 

### Return type

[**AiAttachment**](AiAttachment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAttachmentsSaveFilesMany

> []AiAttachment AiAttachmentsSaveFilesMany(ctx).AiAttachmentsSaveFilesManyRequest(aiAttachmentsSaveFilesManyRequest).Execute()

Save files many



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-attachments-save-files-many/).

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
	aiAttachmentsSaveFilesManyRequest := *openapiclient.NewAiAttachmentsSaveFilesManyRequest([]openapiclient.AiAttachmentsSaveFileRequestInput{*openapiclient.NewAiAttachmentsSaveFileRequestInput("Path_example", "Content_example", float32(123))}) // AiAttachmentsSaveFilesManyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAttachmentsAPI.AiAttachmentsSaveFilesMany(context.Background()).AiAttachmentsSaveFilesManyRequest(aiAttachmentsSaveFilesManyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAttachmentsAPI.AiAttachmentsSaveFilesMany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAttachmentsSaveFilesMany`: []AiAttachment
	fmt.Fprintf(os.Stdout, "Response from `AIAttachmentsAPI.AiAttachmentsSaveFilesMany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAttachmentsSaveFilesManyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiAttachmentsSaveFilesManyRequest** | [**AiAttachmentsSaveFilesManyRequest**](AiAttachmentsSaveFilesManyRequest.md) |  | 

### Return type

[**[]AiAttachment**](AiAttachment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

