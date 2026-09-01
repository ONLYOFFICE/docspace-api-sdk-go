# \FilesOperationsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AbortUploadSession**](FilesOperationsAPI.md#AbortUploadSession) | **Delete** /api/2.0/files/{folderId}/session/{sessionId} | Aborts an in-progress file upload session.
[**AddFavorites**](FilesOperationsAPI.md#AddFavorites) | **Post** /api/2.0/files/favorites | Add favorite files and folders
[**BulkDownload**](FilesOperationsAPI.md#BulkDownload) | **Put** /api/2.0/files/fileops/bulkdownload | Bulk download
[**CheckConversionStatus**](FilesOperationsAPI.md#CheckConversionStatus) | **Get** /api/2.0/files/file/{fileId}/checkconversion | Get conversion status
[**CheckMoveOrCopyBatchItems**](FilesOperationsAPI.md#CheckMoveOrCopyBatchItems) | **Get** /api/2.0/files/fileops/move | Move or copy files to a folder
[**CheckMoveOrCopyDestFolder**](FilesOperationsAPI.md#CheckMoveOrCopyDestFolder) | **Get** /api/2.0/files/fileops/checkdestfolder | Check for moving or copying files to a folder
[**CopyBatchItems**](FilesOperationsAPI.md#CopyBatchItems) | **Put** /api/2.0/files/fileops/copy | Copy to the folder
[**CreateUploadSession**](FilesOperationsAPI.md#CreateUploadSession) | **Post** /api/2.0/files/{folderId}/upload/create_session | Chunked upload
[**CreateUploadSessionInFolder**](FilesOperationsAPI.md#CreateUploadSessionInFolder) | **Post** /api/2.0/files/{folderId}/session | Creates a session for uploading a file to a specific folder in chunks.
[**DeleteBatchItems**](FilesOperationsAPI.md#DeleteBatchItems) | **Put** /api/2.0/files/fileops/delete | Delete files and folders
[**DeleteFavoritesFromBody**](FilesOperationsAPI.md#DeleteFavoritesFromBody) | **Delete** /api/2.0/files/favorites | Delete favorite files and folders (using body parameters)
[**DeleteFileVersions**](FilesOperationsAPI.md#DeleteFileVersions) | **Put** /api/2.0/files/fileops/deleteversion | Delete file versions
[**DuplicateBatchItems**](FilesOperationsAPI.md#DuplicateBatchItems) | **Put** /api/2.0/files/fileops/duplicate | Duplicate files and folders
[**EmptyTrash**](FilesOperationsAPI.md#EmptyTrash) | **Put** /api/2.0/files/fileops/emptytrash | Empty the Trash folder
[**FinalizeSession**](FilesOperationsAPI.md#FinalizeSession) | **Put** /api/2.0/files/{folderId}/session/{sessionId}/finalize | Finalize an upload session
[**GetOperationStatuses**](FilesOperationsAPI.md#GetOperationStatuses) | **Get** /api/2.0/files/fileops | Get active file operations
[**GetOperationStatusesByType**](FilesOperationsAPI.md#GetOperationStatusesByType) | **Get** /api/2.0/files/fileops/{operationType} | Get file operation statuses
[**MarkAsRead**](FilesOperationsAPI.md#MarkAsRead) | **Put** /api/2.0/files/fileops/markasread | Mark as read
[**MoveBatchItems**](FilesOperationsAPI.md#MoveBatchItems) | **Put** /api/2.0/files/fileops/move | Move or copy to a folder
[**StartFileConversion**](FilesOperationsAPI.md#StartFileConversion) | **Put** /api/2.0/files/file/{fileId}/checkconversion | Start file conversion
[**TerminateTasks**](FilesOperationsAPI.md#TerminateTasks) | **Put** /api/2.0/files/fileops/terminate/{id} | Finish active operations
[**UpdateFileComment**](FilesOperationsAPI.md#UpdateFileComment) | **Put** /api/2.0/files/file/{fileId}/comment | Update a comment
[**UploadAsyncSession**](FilesOperationsAPI.md#UploadAsyncSession) | **Post** /api/2.0/files/{folderId}/session/{sessionId}/upload | Handles the upload of a chunk for an existing upload session.
[**UploadSession**](FilesOperationsAPI.md#UploadSession) | **Post** /api/2.0/files/{folderId}/session/{sessionId} | Resumes an ongoing file upload session for uploading additional chunks of data.



## AbortUploadSession

> AbortUploadSession(ctx, sessionId, folderId).Execute()

Aborts an in-progress file upload session.



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/abort-upload-session/).

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
	sessionId := "session-123-abc" // string | The session ID.
	folderId := int32(1) // int32 | The folder ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesOperationsAPI.AbortUploadSession(context.Background(), sessionId, folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.AbortUploadSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The session ID. | 
**folderId** | **int32** | The folder ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAbortUploadSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddFavorites

> BooleanWrapper AddFavorites(ctx).BaseBatchRequestDto(baseBatchRequestDto).Execute()

Add favorite files and folders



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/add-favorites/).

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
	baseBatchRequestDto := *openapiclient.NewBaseBatchRequestDto() // BaseBatchRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.AddFavorites(context.Background()).BaseBatchRequestDto(baseBatchRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.AddFavorites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddFavorites`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.AddFavorites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddFavoritesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseBatchRequestDto** | [**BaseBatchRequestDto**](BaseBatchRequestDto.md) |  | 

### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkDownload

> FileOperationArrayWrapper BulkDownload(ctx).DownloadRequestDto(downloadRequestDto).Execute()

Bulk download



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/bulk-download/).

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
	downloadRequestDto := *openapiclient.NewDownloadRequestDto() // DownloadRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.BulkDownload(context.Background()).DownloadRequestDto(downloadRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.BulkDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkDownload`: FileOperationArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.BulkDownload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadRequestDto** | [**DownloadRequestDto**](DownloadRequestDto.md) |  | 

### Return type

[**FileOperationArrayWrapper**](FileOperationArrayWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckConversionStatus

> ConversationResultArrayWrapper CheckConversionStatus(ctx, fileId).Start(start).Execute()

Get conversion status



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/check-conversion-status/).

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
	fileId := int32(1) // int32 | The file ID to check conversion status.
	start := false // bool | Specifies whether a conversion operation is started or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.CheckConversionStatus(context.Background(), fileId).Start(start).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.CheckConversionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckConversionStatus`: ConversationResultArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.CheckConversionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID to check conversion status. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckConversionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **bool** | Specifies whether a conversion operation is started or not. | 

### Return type

[**ConversationResultArrayWrapper**](ConversationResultArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckMoveOrCopyBatchItems

> FileEntryBaseArrayWrapper CheckMoveOrCopyBatchItems(ctx).InDto(inDto).Execute()

Move or copy files to a folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/check-move-or-copy-batch-items/).

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
	inDto := *openapiclient.NewBatchRequestDto() // BatchRequestDto | The request parameters for copying/moving files. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.CheckMoveOrCopyBatchItems(context.Background()).InDto(inDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.CheckMoveOrCopyBatchItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckMoveOrCopyBatchItems`: FileEntryBaseArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.CheckMoveOrCopyBatchItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckMoveOrCopyBatchItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inDto** | [**BatchRequestDto**](BatchRequestDto.md) | The request parameters for copying/moving files. | 

### Return type

[**FileEntryBaseArrayWrapper**](FileEntryBaseArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckMoveOrCopyDestFolder

> CheckDestFolderWrapper CheckMoveOrCopyDestFolder(ctx).InDto(inDto).Execute()

Check for moving or copying files to a folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/check-move-or-copy-dest-folder/).

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
	inDto := *openapiclient.NewBatchRequestDto() // BatchRequestDto | The request parameters for copying/moving files. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.CheckMoveOrCopyDestFolder(context.Background()).InDto(inDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.CheckMoveOrCopyDestFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckMoveOrCopyDestFolder`: CheckDestFolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.CheckMoveOrCopyDestFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckMoveOrCopyDestFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inDto** | [**BatchRequestDto**](BatchRequestDto.md) | The request parameters for copying/moving files. | 

### Return type

[**CheckDestFolderWrapper**](CheckDestFolderWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CopyBatchItems

> FileOperationArrayWrapper CopyBatchItems(ctx).BatchRequestDto(batchRequestDto).Execute()

Copy to the folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/copy-batch-items/).

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
	batchRequestDto := *openapiclient.NewBatchRequestDto() // BatchRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.CopyBatchItems(context.Background()).BatchRequestDto(batchRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.CopyBatchItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CopyBatchItems`: FileOperationArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.CopyBatchItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCopyBatchItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchRequestDto** | [**BatchRequestDto**](BatchRequestDto.md) |  | 

### Return type

[**FileOperationArrayWrapper**](FileOperationArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUploadSession

> ChunkedUploadSessionResponseWrapperIntegerWrapper CreateUploadSession(ctx, folderId).SessionRequest(sessionRequest).Execute()

Chunked upload



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-upload-session/).

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
	folderId := int32(1) // int32 | The session folder ID.
	sessionRequest := *openapiclient.NewSessionRequest("My Document.docx") // SessionRequest | The session parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.CreateUploadSession(context.Background(), folderId).SessionRequest(sessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.CreateUploadSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUploadSession`: ChunkedUploadSessionResponseWrapperIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.CreateUploadSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The session folder ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUploadSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sessionRequest** | [**SessionRequest**](SessionRequest.md) | The session parameters. | 

### Return type

[**ChunkedUploadSessionResponseWrapperIntegerWrapper**](ChunkedUploadSessionResponseWrapperIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUploadSessionInFolder

> ChunkedUploadSessionResponseIntegerWrapper CreateUploadSessionInFolder(ctx, folderId).SessionRequest(sessionRequest).Execute()

Creates a session for uploading a file to a specific folder in chunks.



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-upload-session-in-folder/).

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
	folderId := int32(1) // int32 | The session folder ID.
	sessionRequest := *openapiclient.NewSessionRequest("My Document.docx") // SessionRequest | The session parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.CreateUploadSessionInFolder(context.Background(), folderId).SessionRequest(sessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.CreateUploadSessionInFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUploadSessionInFolder`: ChunkedUploadSessionResponseIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.CreateUploadSessionInFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The session folder ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUploadSessionInFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sessionRequest** | [**SessionRequest**](SessionRequest.md) | The session parameters. | 

### Return type

[**ChunkedUploadSessionResponseIntegerWrapper**](ChunkedUploadSessionResponseIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBatchItems

> FileOperationArrayWrapper DeleteBatchItems(ctx).DeleteBatchRequestDto(deleteBatchRequestDto).Execute()

Delete files and folders



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-batch-items/).

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
	deleteBatchRequestDto := *openapiclient.NewDeleteBatchRequestDto() // DeleteBatchRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.DeleteBatchItems(context.Background()).DeleteBatchRequestDto(deleteBatchRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.DeleteBatchItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBatchItems`: FileOperationArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.DeleteBatchItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBatchItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteBatchRequestDto** | [**DeleteBatchRequestDto**](DeleteBatchRequestDto.md) |  | 

### Return type

[**FileOperationArrayWrapper**](FileOperationArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFavoritesFromBody

> BooleanWrapper DeleteFavoritesFromBody(ctx).BaseBatchRequestDto(baseBatchRequestDto).Execute()

Delete favorite files and folders (using body parameters)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-favorites-from-body/).

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
	baseBatchRequestDto := *openapiclient.NewBaseBatchRequestDto() // BaseBatchRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.DeleteFavoritesFromBody(context.Background()).BaseBatchRequestDto(baseBatchRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.DeleteFavoritesFromBody``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFavoritesFromBody`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.DeleteFavoritesFromBody`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFavoritesFromBodyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseBatchRequestDto** | [**BaseBatchRequestDto**](BaseBatchRequestDto.md) |  | 

### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFileVersions

> FileOperationWrapper DeleteFileVersions(ctx).DeleteVersionBatchRequestDto(deleteVersionBatchRequestDto).Execute()

Delete file versions



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-file-versions/).

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
	deleteVersionBatchRequestDto := *openapiclient.NewDeleteVersionBatchRequestDto(int32(1), []int32{int32(123)}) // DeleteVersionBatchRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.DeleteFileVersions(context.Background()).DeleteVersionBatchRequestDto(deleteVersionBatchRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.DeleteFileVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFileVersions`: FileOperationWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.DeleteFileVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteVersionBatchRequestDto** | [**DeleteVersionBatchRequestDto**](DeleteVersionBatchRequestDto.md) |  | 

### Return type

[**FileOperationWrapper**](FileOperationWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DuplicateBatchItems

> FileOperationArrayWrapper DuplicateBatchItems(ctx).DuplicateRequestDto(duplicateRequestDto).Execute()

Duplicate files and folders



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/duplicate-batch-items/).

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
	duplicateRequestDto := *openapiclient.NewDuplicateRequestDto() // DuplicateRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.DuplicateBatchItems(context.Background()).DuplicateRequestDto(duplicateRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.DuplicateBatchItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DuplicateBatchItems`: FileOperationArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.DuplicateBatchItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDuplicateBatchItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **duplicateRequestDto** | [**DuplicateRequestDto**](DuplicateRequestDto.md) |  | 

### Return type

[**FileOperationArrayWrapper**](FileOperationArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmptyTrash

> FileOperationArrayWrapper EmptyTrash(ctx).Single(single).FolderType(folderType).Execute()

Empty the Trash folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/empty-trash/).

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
	single := false // bool | Specifies whether to return only the current operation (optional)
	folderType := []int32{int32(0)} // []int32 | The parent folder types used to empty the trash only from the items originally located in the sections of the specified types. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.EmptyTrash(context.Background()).Single(single).FolderType(folderType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.EmptyTrash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmptyTrash`: FileOperationArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.EmptyTrash`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmptyTrashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **single** | **bool** | Specifies whether to return only the current operation | 
 **folderType** | **[]int32** | The parent folder types used to empty the trash only from the items originally located in the sections of the specified types. | 

### Return type

[**FileOperationArrayWrapper**](FileOperationArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinalizeSession

> UploadSessionResponseIntegerWrapper FinalizeSession(ctx, folderId, sessionId).Execute()

Finalize an upload session



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/finalize-session/).

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
	folderId := int32(1) // int32 | The folder ID.
	sessionId := "doc_key_123" // string | The session ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.FinalizeSession(context.Background(), folderId, sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.FinalizeSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinalizeSession`: UploadSessionResponseIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.FinalizeSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID. | 
**sessionId** | **string** | The session ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinalizeSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UploadSessionResponseIntegerWrapper**](UploadSessionResponseIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOperationStatuses

> FileOperationArrayWrapper GetOperationStatuses(ctx).Id(id).Execute()

Get active file operations



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-operation-statuses/).

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
	id := "operation-123-abc" // string | The ID of the file operation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.GetOperationStatuses(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.GetOperationStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOperationStatuses`: FileOperationArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.GetOperationStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOperationStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The ID of the file operation. | 

### Return type

[**FileOperationArrayWrapper**](FileOperationArrayWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOperationStatusesByType

> FileOperationArrayWrapper GetOperationStatusesByType(ctx, operationType).Id(id).Execute()

Get file operation statuses



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-operation-statuses-by-type/).

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
	operationType := openapiclient.FileOperationType(0) // FileOperationType | Specifies the type of file operation to be retrieved.
	id := "operation-123-abc" // string | The ID of the file operation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.GetOperationStatusesByType(context.Background(), operationType).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.GetOperationStatusesByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOperationStatusesByType`: FileOperationArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.GetOperationStatusesByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationType** | [**FileOperationType**](.md) | Specifies the type of file operation to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOperationStatusesByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **string** | The ID of the file operation. | 

### Return type

[**FileOperationArrayWrapper**](FileOperationArrayWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkAsRead

> FileOperationArrayWrapper MarkAsRead(ctx).BaseBatchRequestDto(baseBatchRequestDto).Execute()

Mark as read



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/mark-as-read/).

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
	baseBatchRequestDto := *openapiclient.NewBaseBatchRequestDto() // BaseBatchRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.MarkAsRead(context.Background()).BaseBatchRequestDto(baseBatchRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.MarkAsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkAsRead`: FileOperationArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.MarkAsRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkAsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseBatchRequestDto** | [**BaseBatchRequestDto**](BaseBatchRequestDto.md) |  | 

### Return type

[**FileOperationArrayWrapper**](FileOperationArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveBatchItems

> FileOperationArrayWrapper MoveBatchItems(ctx).BatchRequestDto(batchRequestDto).Execute()

Move or copy to a folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/move-batch-items/).

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
	batchRequestDto := *openapiclient.NewBatchRequestDto() // BatchRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.MoveBatchItems(context.Background()).BatchRequestDto(batchRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.MoveBatchItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveBatchItems`: FileOperationArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.MoveBatchItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoveBatchItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchRequestDto** | [**BatchRequestDto**](BatchRequestDto.md) |  | 

### Return type

[**FileOperationArrayWrapper**](FileOperationArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartFileConversion

> ConversationResultArrayWrapper StartFileConversion(ctx, fileId).CheckConversionRequestDtoInteger(checkConversionRequestDtoInteger).Execute()

Start file conversion



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/start-file-conversion/).

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
	fileId := int32(1) // int32 | The file ID to start conversion proccess.
	checkConversionRequestDtoInteger := *openapiclient.NewCheckConversionRequestDtoInteger() // CheckConversionRequestDtoInteger | The parameters for checking file conversion. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.StartFileConversion(context.Background(), fileId).CheckConversionRequestDtoInteger(checkConversionRequestDtoInteger).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.StartFileConversion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartFileConversion`: ConversationResultArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.StartFileConversion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID to start conversion proccess. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartFileConversionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkConversionRequestDtoInteger** | [**CheckConversionRequestDtoInteger**](CheckConversionRequestDtoInteger.md) | The parameters for checking file conversion. | 

### Return type

[**ConversationResultArrayWrapper**](ConversationResultArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateTasks

> FileOperationArrayWrapper TerminateTasks(ctx, id).Execute()

Finish active operations



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/terminate-tasks/).

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
	id := "some-operation-id" // string | The operation unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.TerminateTasks(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.TerminateTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TerminateTasks`: FileOperationArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.TerminateTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The operation unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileOperationArrayWrapper**](FileOperationArrayWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFileComment

> StringWrapper UpdateFileComment(ctx, fileId).UpdateComment(updateComment).Execute()

Update a comment



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-file-comment/).

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
	fileId := int32(1) // int32 | The file ID where the comment is located.
	updateComment := *openapiclient.NewUpdateComment(int32(1)) // UpdateComment | The parameters for updating a comment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.UpdateFileComment(context.Background(), fileId).UpdateComment(updateComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.UpdateFileComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFileComment`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.UpdateFileComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID where the comment is located. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFileCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateComment** | [**UpdateComment**](UpdateComment.md) | The parameters for updating a comment. | 

### Return type

[**StringWrapper**](StringWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadAsyncSession

> ChunkedUploadSessionResponseIntegerWrapper UploadAsyncSession(ctx, folderId, sessionId).ChunkNumber(chunkNumber).File(file).Execute()

Handles the upload of a chunk for an existing upload session.



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/upload-async-session/).

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
	folderId := int32(1) // int32 | The folder ID.
	sessionId := "session_abc123" // string | The upload session ID.
	chunkNumber := int32(1) // int32 | The chunk number. (optional)
	file := os.NewFile(1234, "some_file") // *os.File | The file chunk to be uploaded as part of the multipart/form-data request.  This property represents the uploaded file chunk content from the HTTP request form for chunked upload operations.  The file chunk is accessed via the IFormFile interface which provides access to the chunk content and length. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.UploadAsyncSession(context.Background(), folderId, sessionId).ChunkNumber(chunkNumber).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.UploadAsyncSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadAsyncSession`: ChunkedUploadSessionResponseIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.UploadAsyncSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID. | 
**sessionId** | **string** | The upload session ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadAsyncSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **chunkNumber** | **int32** | The chunk number. | 
 **file** | ***os.File** | The file chunk to be uploaded as part of the multipart/form-data request.  This property represents the uploaded file chunk content from the HTTP request form for chunked upload operations.  The file chunk is accessed via the IFormFile interface which provides access to the chunk content and length. | 

### Return type

[**ChunkedUploadSessionResponseIntegerWrapper**](ChunkedUploadSessionResponseIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadSession

> UploadSessionResponseIntegerWrapper UploadSession(ctx, folderId, sessionId).File(file).Execute()

Resumes an ongoing file upload session for uploading additional chunks of data.



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/upload-session/).

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
	folderId := int32(1) // int32 | The folder ID.
	sessionId := "session_abc123" // string | The upload session ID.
	file := os.NewFile(1234, "some_file") // *os.File | The file to be uploaded as part of the multipart/form-data request.  This property represents the uploaded file content from the HTTP request form.  The file is accessed via the IFormFile interface which provides access to the file name, content type, length, and stream. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.UploadSession(context.Background(), folderId, sessionId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.UploadSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadSession`: UploadSessionResponseIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.UploadSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID. | 
**sessionId** | **string** | The upload session ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | The file to be uploaded as part of the multipart/form-data request.  This property represents the uploaded file content from the HTTP request form.  The file is accessed via the IFormFile interface which provides access to the file name, content type, length, and stream. | 

### Return type

[**UploadSessionResponseIntegerWrapper**](UploadSessionResponseIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

