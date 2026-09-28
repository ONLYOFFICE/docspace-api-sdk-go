# \FilesOperationsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AbortUploadSession**](FilesOperationsAPI.md#AbortUploadSession) | **Delete** /api/2.0/files/{folderId}/session/{sessionId} | Abort an upload session
[**AddFavorites**](FilesOperationsAPI.md#AddFavorites) | **Post** /api/2.0/files/favorites | Add favorite files and folders
[**BulkDownload**](FilesOperationsAPI.md#BulkDownload) | **Put** /api/2.0/files/fileops/bulkdownload | Bulk download
[**CheckConversionStatus**](FilesOperationsAPI.md#CheckConversionStatus) | **Get** /api/2.0/files/file/{fileId}/checkconversion | Get conversion status
[**CheckMoveOrCopyBatchItems**](FilesOperationsAPI.md#CheckMoveOrCopyBatchItems) | **Get** /api/2.0/files/fileops/move | Check move or copy conflicts
[**CheckMoveOrCopyDestFolder**](FilesOperationsAPI.md#CheckMoveOrCopyDestFolder) | **Get** /api/2.0/files/fileops/checkdestfolder | Check the destination folder
[**CopyBatchItems**](FilesOperationsAPI.md#CopyBatchItems) | **Put** /api/2.0/files/fileops/copy | Copy files and folders
[**CreateUploadSession**](FilesOperationsAPI.md#CreateUploadSession) | **Post** /api/2.0/files/{folderId}/upload/create_session | Chunked upload
[**CreateUploadSessionInFolder**](FilesOperationsAPI.md#CreateUploadSessionInFolder) | **Post** /api/2.0/files/{folderId}/session | Create an upload session
[**DeleteBatchItems**](FilesOperationsAPI.md#DeleteBatchItems) | **Put** /api/2.0/files/fileops/delete | Delete files and folders
[**DeleteFavoritesFromBody**](FilesOperationsAPI.md#DeleteFavoritesFromBody) | **Delete** /api/2.0/files/favorites | Delete favorite files and folders
[**DeleteFileVersions**](FilesOperationsAPI.md#DeleteFileVersions) | **Put** /api/2.0/files/fileops/deleteversion | Delete file versions
[**DuplicateBatchItems**](FilesOperationsAPI.md#DuplicateBatchItems) | **Put** /api/2.0/files/fileops/duplicate | Duplicate files and folders
[**EmptyTrash**](FilesOperationsAPI.md#EmptyTrash) | **Put** /api/2.0/files/fileops/emptytrash | Empty the Trash folder
[**FinalizeSession**](FilesOperationsAPI.md#FinalizeSession) | **Put** /api/2.0/files/{folderId}/session/{sessionId}/finalize | Finalize an upload session
[**GetOperationStatuses**](FilesOperationsAPI.md#GetOperationStatuses) | **Get** /api/2.0/files/fileops | Get active file operations
[**GetOperationStatusesByType**](FilesOperationsAPI.md#GetOperationStatusesByType) | **Get** /api/2.0/files/fileops/{operationType} | Get file operations by type
[**MarkAsRead**](FilesOperationsAPI.md#MarkAsRead) | **Put** /api/2.0/files/fileops/markasread | Mark files and folders as read
[**MoveBatchItems**](FilesOperationsAPI.md#MoveBatchItems) | **Put** /api/2.0/files/fileops/move | Move files and folders
[**StartFileConversion**](FilesOperationsAPI.md#StartFileConversion) | **Put** /api/2.0/files/file/{fileId}/checkconversion | Start file conversion
[**TerminateTasks**](FilesOperationsAPI.md#TerminateTasks) | **Put** /api/2.0/files/fileops/terminate/{id} | Cancel file operations
[**UpdateFileComment**](FilesOperationsAPI.md#UpdateFileComment) | **Put** /api/2.0/files/file/{fileId}/comment | Update a comment
[**UploadAsyncSession**](FilesOperationsAPI.md#UploadAsyncSession) | **Post** /api/2.0/files/{folderId}/session/{sessionId}/upload | Upload a numbered chunk
[**UploadSession**](FilesOperationsAPI.md#UploadSession) | **Post** /api/2.0/files/{folderId}/session/{sessionId} | Upload the next chunk



## AbortUploadSession

> AbortUploadSession(ctx, sessionId, folderId).Execute()

Abort an upload session



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
	sessionId := "9f1c7a2b4d3e4f5a8b6c0d1e2f3a4b5c" // string | The session to cancel, as returned in `id` when it was created: a 32-character hexadecimal string that  identifies the session on its own.
	folderId := int32(1) // int32 | The folder the session was opened against. It is part of the route only and is not matched against the  session, which is found by its own id.

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
**sessionId** | **string** | The session to cancel, as returned in `id` when it was created: a 32-character hexadecimal string that  identifies the session on its own. | 
**folderId** | **int32** | The folder the session was opened against. It is part of the route only and is not matched against the  session, which is found by its own id. | 

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

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

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
	fileId := int32(1) // int32 | The file whose conversion is asked about.
	start := false // bool | Whether to start the conversion as well: `true` queues it with the default output format and no password,  `false` only reports what the portal already knows. (optional)

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
**fileId** | **int32** | The file whose conversion is asked about. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckConversionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **bool** | Whether to start the conversion as well: `true` queues it with the default output format and no password,  `false` only reports what the portal already knows. | 

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

Check move or copy conflicts



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
	inDto := *openapiclient.NewBatchRequestDto() // BatchRequestDto | The files and folders to move or copy, the folder they go to, and the way name clashes are settled. (optional)

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
 **inDto** | [**BatchRequestDto**](BatchRequestDto.md) | The files and folders to move or copy, the folder they go to, and the way name clashes are settled. | 

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

Check the destination folder



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
	inDto := *openapiclient.NewBatchRequestDto() // BatchRequestDto | The files and folders to move or copy, the folder they go to, and the way name clashes are settled. (optional)

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
 **inDto** | [**BatchRequestDto**](BatchRequestDto.md) | The files and folders to move or copy, the folder they go to, and the way name clashes are settled. | 

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

Copy files and folders



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

> ChunkedUploadSessionResponseWrapperWrapper CreateUploadSession(ctx, folderId).SessionRequest(sessionRequest).Execute()

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
	folderId := int32(1) // int32 | The folder that receives the file; take the id from a listing such as `GET api/2.0/files/@root`. A room or an  ordinary folder inside one is accepted, a section root is not.
	sessionRequest := *openapiclient.NewSessionRequest("My Document.docx") // SessionRequest | The file the session is opened for, and how a clash with an existing name is settled.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.CreateUploadSession(context.Background(), folderId).SessionRequest(sessionRequest).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// folderId := "sbox-42"
	// thirdPartyResp, r, err := apiClient.FilesOperationsAPI.CreateUploadSession(context.Background(), folderId).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.CreateUploadSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUploadSession`: ChunkedUploadSessionResponseWrapperWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.CreateUploadSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder that receives the file; take the id from a listing such as `GET api/2.0/files/@root`. A room or an  ordinary folder inside one is accepted, a section root is not. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUploadSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sessionRequest** | [**SessionRequest**](SessionRequest.md) | The file the session is opened for, and how a clash with an existing name is settled. | 

### Return type

[**ChunkedUploadSessionResponseWrapperWrapper**](ChunkedUploadSessionResponseWrapperWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUploadSessionInFolder

> ChunkedUploadSessionResponseResponseWrapper CreateUploadSessionInFolder(ctx, folderId).SessionRequest(sessionRequest).Execute()

Create an upload session



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
	folderId := int32(1) // int32 | The folder that receives the file; take the id from a listing such as `GET api/2.0/files/@root`. A room or an  ordinary folder inside one is accepted, a section root is not.
	sessionRequest := *openapiclient.NewSessionRequest("My Document.docx") // SessionRequest | The file the session is opened for, and how a clash with an existing name is settled.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.CreateUploadSessionInFolder(context.Background(), folderId).SessionRequest(sessionRequest).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// folderId := "sbox-42"
	// thirdPartyResp, r, err := apiClient.FilesOperationsAPI.CreateUploadSessionInFolder(context.Background(), folderId).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.CreateUploadSessionInFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUploadSessionInFolder`: ChunkedUploadSessionResponseResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.CreateUploadSessionInFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder that receives the file; take the id from a listing such as `GET api/2.0/files/@root`. A room or an  ordinary folder inside one is accepted, a section root is not. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUploadSessionInFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sessionRequest** | [**SessionRequest**](SessionRequest.md) | The file the session is opened for, and how a clash with an existing name is settled. | 

### Return type

[**ChunkedUploadSessionResponseResponseWrapper**](ChunkedUploadSessionResponseResponseWrapper.md)

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

Delete favorite files and folders



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

> FileOperationArrayWrapper DeleteFileVersions(ctx).DeleteVersionBatchRequestDto(deleteVersionBatchRequestDto).Execute()

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
	// response from `DeleteFileVersions`: FileOperationArrayWrapper
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

[**FileOperationArrayWrapper**](FileOperationArrayWrapper.md)

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
	single := false // bool | Which operations the answer carries: `true` returns the operation this call started and nothing else, `false`  returns every delete operation that the caller has running or unread. (optional)
	folderType := []int32{int32(0)} // []int32 | Limits the sweep to the items whose original location was inside a section or a room of one of the named  types, leaving the rest of the Trash untouched; without the parameter the whole Trash is emptied. `5` covers  what was deleted from personal documents, `14` what was deleted from rooms. (optional)

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
 **single** | **bool** | Which operations the answer carries: `true` returns the operation this call started and nothing else, `false`  returns every delete operation that the caller has running or unread. | 
 **folderType** | **[]int32** | Limits the sweep to the items whose original location was inside a section or a room of one of the named  types, leaving the rest of the Trash untouched; without the parameter the whole Trash is emptied. `5` covers  what was deleted from personal documents, `14` what was deleted from rooms. | 

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

> UploadSessionResponseWrapper FinalizeSession(ctx, folderId, sessionId).Execute()

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
	folderId := int32(1) // int32 | The folder the session was opened against. It is part of the route only and is not matched against the  session, which is found by its own id.
	sessionId := "9f1c7a2b4d3e4f5a8b6c0d1e2f3a4b5c" // string | The session to assemble, as returned in `id` when it was created: a 32-character hexadecimal string that  identifies the session on its own.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.FinalizeSession(context.Background(), folderId, sessionId).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// folderId := "sbox-42"sessionId := "sbox-42"
	// thirdPartyResp, r, err := apiClient.FilesOperationsAPI.FinalizeSession(context.Background(), folderId, sessionId).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.FinalizeSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinalizeSession`: UploadSessionResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.FinalizeSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder the session was opened against. It is part of the route only and is not matched against the  session, which is found by its own id. | 
**sessionId** | **string** | The session to assemble, as returned in `id` when it was created: a 32-character hexadecimal string that  identifies the session on its own. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinalizeSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UploadSessionResponseWrapper**](UploadSessionResponseWrapper.md)

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
	id := "b2f3e9a4-7c15-4d8e-9f60-3a1c5e7d0b42" // string | The operation to report on, as returned in `id` when it was started; without it every operation of the caller  is reported. An id that is not among the caller's operations gives an empty answer rather than an error. (optional)

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
 **id** | **string** | The operation to report on, as returned in `id` when it was started; without it every operation of the caller  is reported. An id that is not among the caller's operations gives an empty answer rather than an error. | 

### Return type

[**FileOperationArrayWrapper**](FileOperationArrayWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOperationStatusesByType

> FileOperationArrayWrapper GetOperationStatusesByType(ctx, operationType).Id(id).Execute()

Get file operations by type



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
	operationType := openapiclient.FileOperationType(0) // FileOperationType | The kind of operation the answer is limited to. Only the kinds that have a queue of their own ever carry  records — a copy, a deletion, a download, a mark-as-read and a duplication — and moves cannot be read through  this route at all, because its address belongs to another operation.
	id := "b2f3e9a4-7c15-4d8e-9f60-3a1c5e7d0b42" // string | The operation to report on, as returned in `id` when it was started; without it every operation of the caller  is reported. An id that is not among the caller's operations gives an empty answer rather than an error. (optional)

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
**operationType** | [**FileOperationType**](.md) | The kind of operation the answer is limited to. Only the kinds that have a queue of their own ever carry  records — a copy, a deletion, a download, a mark-as-read and a duplication — and moves cannot be read through  this route at all, because its address belongs to another operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOperationStatusesByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **string** | The operation to report on, as returned in `id` when it was started; without it every operation of the caller  is reported. An id that is not among the caller's operations gives an empty answer rather than an error. | 

### Return type

[**FileOperationArrayWrapper**](FileOperationArrayWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkAsRead

> FileOperationArrayWrapper MarkAsRead(ctx).BaseBatchRequestDto(baseBatchRequestDto).Execute()

Mark files and folders as read



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

Move files and folders



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

> ConversationResultArrayWrapper StartFileConversion(ctx, fileId).CheckConversionRequestDto(checkConversionRequestDto).Execute()

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
	fileId := int32(1) // int32 | The file to convert.
	checkConversionRequestDto := *openapiclient.NewCheckConversionRequestDto() // CheckConversionRequestDto | The parameters of the conversion. The whole body may be omitted, in which case the defaults of the portal  apply. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.StartFileConversion(context.Background(), fileId).CheckConversionRequestDto(checkConversionRequestDto).Execute()
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
**fileId** | **int32** | The file to convert. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartFileConversionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkConversionRequestDto** | [**CheckConversionRequestDto**](CheckConversionRequestDto.md) | The parameters of the conversion. The whole body may be omitted, in which case the defaults of the portal  apply. | 

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

Cancel file operations



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
	id := "b2f3e9a4-7c15-4d8e-9f60-3a1c5e7d0b42" // string | The operation to cancel, as returned in `id` when it was started. A call that leaves the route segment out  cancels every operation of the caller, and an id that is not among their operations cancels nothing without  being an error.

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
**id** | **string** | The operation to cancel, as returned in `id` when it was started. A call that leaves the route segment out  cancels every operation of the caller, and an id that is not among their operations cancels nothing without  being an error. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileOperationArrayWrapper**](FileOperationArrayWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

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
	fileId := int32(1) // int32 | The file whose version comment is replaced.
	updateComment := *openapiclient.NewUpdateComment(int32(1)) // UpdateComment | The version and the comment to store on it.

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
**fileId** | **int32** | The file whose version comment is replaced. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFileCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateComment** | [**UpdateComment**](UpdateComment.md) | The version and the comment to store on it. | 

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

> ChunkedUploadSessionResponseResponseWrapper UploadAsyncSession(ctx, folderId, sessionId).ChunkNumber(chunkNumber).File(file).Execute()

Upload a numbered chunk



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
	folderId := int32(1) // int32 | The folder the session was opened against. It is part of the route only and is not matched against the  session, which is found by its own id.
	sessionId := "9f1c7a2b4d3e4f5a8b6c0d1e2f3a4b5c" // string | The session this part belongs to, as returned in `id` when it was created; a 32-character hexadecimal string.
	chunkNumber := int32(1) // int32 | The position of this part in the file, counted from 1. Sending the same number again replaces that part  instead of adding one, which is how a failed part is retried; leaving the number out makes the server count  the parts itself. (optional)
	file := os.NewFile(1234, "some_file") // *os.File | The part of the file to store, sent as the multipart field of the same name. It is kept under the number given  beside it, and a part larger than the portal chunk size is refused. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.UploadAsyncSession(context.Background(), folderId, sessionId).ChunkNumber(chunkNumber).File(file).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// folderId := "sbox-42"sessionId := "sbox-42"
	// thirdPartyResp, r, err := apiClient.FilesOperationsAPI.UploadAsyncSession(context.Background(), folderId, sessionId).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.UploadAsyncSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadAsyncSession`: ChunkedUploadSessionResponseResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.UploadAsyncSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder the session was opened against. It is part of the route only and is not matched against the  session, which is found by its own id. | 
**sessionId** | **string** | The session this part belongs to, as returned in `id` when it was created; a 32-character hexadecimal string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadAsyncSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **chunkNumber** | **int32** | The position of this part in the file, counted from 1. Sending the same number again replaces that part  instead of adding one, which is how a failed part is retried; leaving the number out makes the server count  the parts itself. | 
 **file** | ***os.File** | The part of the file to store, sent as the multipart field of the same name. It is kept under the number given  beside it, and a part larger than the portal chunk size is refused. | 

### Return type

[**ChunkedUploadSessionResponseResponseWrapper**](ChunkedUploadSessionResponseResponseWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadSession

> UploadSessionResponseWrapper UploadSession(ctx, folderId, sessionId).File(file).Execute()

Upload the next chunk



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
	folderId := int32(1) // int32 | The folder the session was opened against. It is part of the route only and is not matched against the  session, which is found by its own id.
	sessionId := "9f1c7a2b4d3e4f5a8b6c0d1e2f3a4b5c" // string | The session this part belongs to, as returned in `id` when it was created; the parts of one session must be  sent one after another, not in parallel.
	file := os.NewFile(1234, "some_file") // *os.File | The next part of the file, sent as the multipart field of the same name. Parts are appended in the order they  arrive, and a part larger than the portal chunk size is refused. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesOperationsAPI.UploadSession(context.Background(), folderId, sessionId).File(file).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// folderId := "sbox-42"sessionId := "sbox-42"
	// thirdPartyResp, r, err := apiClient.FilesOperationsAPI.UploadSession(context.Background(), folderId, sessionId).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesOperationsAPI.UploadSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadSession`: UploadSessionResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesOperationsAPI.UploadSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder the session was opened against. It is part of the route only and is not matched against the  session, which is found by its own id. | 
**sessionId** | **string** | The session this part belongs to, as returned in `id` when it was created; the parts of one session must be  sent one after another, not in parallel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | The next part of the file, sent as the multipart field of the same name. Parts are appended in the order they  arrive, and a part larger than the portal chunk size is refused. | 

### Return type

[**UploadSessionResponseWrapper**](UploadSessionResponseWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

