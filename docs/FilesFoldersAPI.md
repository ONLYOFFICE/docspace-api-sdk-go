# \FilesFoldersAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckUpload**](FilesFoldersAPI.md#CheckUpload) | **Post** /api/2.0/files/{folderId}/upload/check | Check file uploads
[**CreateFolder**](FilesFoldersAPI.md#CreateFolder) | **Post** /api/2.0/files/folder/{folderId} | Create a folder
[**CreateFolderPrimaryExternalLink**](FilesFoldersAPI.md#CreateFolderPrimaryExternalLink) | **Post** /api/2.0/files/folder/{id}/link | Create primary external link
[**CreateReportFolderHistory**](FilesFoldersAPI.md#CreateReportFolderHistory) | **Post** /api/2.0/files/folder/{folderId}/log/report | Start the folder history report generation
[**DeleteFolder**](FilesFoldersAPI.md#DeleteFolder) | **Delete** /api/2.0/files/folder/{folderId} | Delete a folder
[**GenerateXlsxByFolder**](FilesFoldersAPI.md#GenerateXlsxByFolder) | **Post** /api/2.0/files/folder/{folderId}/xlsx | Generate XLSX report by folder
[**GetFavoritesFolder**](FilesFoldersAPI.md#GetFavoritesFolder) | **Get** /api/2.0/files/@favorites | Get the Favorites section
[**GetFilesUsedSpace**](FilesFoldersAPI.md#GetFilesUsedSpace) | **Get** /api/2.0/files/filesusedspace | Get used space of files
[**GetFolder**](FilesFoldersAPI.md#GetFolder) | **Get** /api/2.0/files/{folderId}/formfilter | Get folder form filter
[**GetFolderByFolderId**](FilesFoldersAPI.md#GetFolderByFolderId) | **Get** /api/2.0/files/{folderId} | Get a folder by ID
[**GetFolderHistory**](FilesFoldersAPI.md#GetFolderHistory) | **Get** /api/2.0/files/folder/{folderId}/log | Get folder history
[**GetFolderInfo**](FilesFoldersAPI.md#GetFolderInfo) | **Get** /api/2.0/files/folder/{folderId} | Get folder information
[**GetFolderLinks**](FilesFoldersAPI.md#GetFolderLinks) | **Get** /api/2.0/files/folder/{id}/links | Get the folder links
[**GetFolderPath**](FilesFoldersAPI.md#GetFolderPath) | **Get** /api/2.0/files/folder/{folderId}/path | Get the folder path
[**GetFolderPrimaryExternalLink**](FilesFoldersAPI.md#GetFolderPrimaryExternalLink) | **Get** /api/2.0/files/folder/{id}/link | Get primary external link
[**GetFolders**](FilesFoldersAPI.md#GetFolders) | **Get** /api/2.0/files/{folderId}/subfolders | Get subfolders
[**GetFormsFolder**](FilesFoldersAPI.md#GetFormsFolder) | **Get** /api/2.0/files/@forms | Get the Forms section
[**GetMyFolder**](FilesFoldersAPI.md#GetMyFolder) | **Get** /api/2.0/files/@my | Get the My documents section
[**GetNewFolderItems**](FilesFoldersAPI.md#GetNewFolderItems) | **Get** /api/2.0/files/{folderId}/news | Get new folder items
[**GetRecentFolder**](FilesFoldersAPI.md#GetRecentFolder) | **Get** /api/2.0/files/recent | Get the Recent section
[**GetReportFolderHistory**](FilesFoldersAPI.md#GetReportFolderHistory) | **Get** /api/2.0/files/folder/{folderId}/log/report | Get the folder history report generation status
[**GetRootFolders**](FilesFoldersAPI.md#GetRootFolders) | **Get** /api/2.0/files/@root | Get filtered sections
[**GetTrashFolder**](FilesFoldersAPI.md#GetTrashFolder) | **Get** /api/2.0/files/@trash | Get the Trash section
[**InsertFile**](FilesFoldersAPI.md#InsertFile) | **Post** /api/2.0/files/{folderId}/insert | Insert a file
[**InsertFileToMyFromBody**](FilesFoldersAPI.md#InsertFileToMyFromBody) | **Post** /api/2.0/files/@my/insert | Insert a file to the My documents section
[**RenameFolder**](FilesFoldersAPI.md#RenameFolder) | **Put** /api/2.0/files/folder/{folderId} | Rename a folder
[**SetFolderOrder**](FilesFoldersAPI.md#SetFolderOrder) | **Put** /api/2.0/files/folder/{folderId}/order | Set folder order
[**SetFolderPrimaryExternalLink**](FilesFoldersAPI.md#SetFolderPrimaryExternalLink) | **Put** /api/2.0/files/folder/{id}/links | Set the folder external link
[**TerminateReportFolderHistory**](FilesFoldersAPI.md#TerminateReportFolderHistory) | **Delete** /api/2.0/files/folder/{folderId}/log/report | Terminate the folder history report generation
[**UploadFile**](FilesFoldersAPI.md#UploadFile) | **Post** /api/2.0/files/{folderId}/upload | Upload a file
[**UploadFileToMy**](FilesFoldersAPI.md#UploadFileToMy) | **Post** /api/2.0/files/@my/upload | Upload a file to the My documents section



## CheckUpload

> STRINGArrayWrapper CheckUpload(ctx, folderId).CheckUploadRequest(checkUploadRequest).Execute()

Check file uploads



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/check-upload/).

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
	checkUploadRequest := *openapiclient.NewCheckUploadRequest() // CheckUploadRequest | The request parameters for checking file uploads.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.CheckUpload(context.Background(), folderId).CheckUploadRequest(checkUploadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.CheckUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckUpload`: STRINGArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.CheckUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkUploadRequest** | [**CheckUploadRequest**](CheckUploadRequest.md) | The request parameters for checking file uploads. | 

### Return type

[**STRINGArrayWrapper**](STRINGArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFolder

> FolderIntegerWrapper CreateFolder(ctx, folderId).CreateFolder(createFolder).Execute()

Create a folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-folder/).

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
	folderId := int32(1) // int32 | The folder ID for the folder creation.
	createFolder := *openapiclient.NewCreateFolder("New Folder") // CreateFolder | The parameters for creating a folder.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.CreateFolder(context.Background(), folderId).CreateFolder(createFolder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.CreateFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFolder`: FolderIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.CreateFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID for the folder creation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFolder** | [**CreateFolder**](CreateFolder.md) | The parameters for creating a folder. | 

### Return type

[**FolderIntegerWrapper**](FolderIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFolderPrimaryExternalLink

> FileShareWrapper CreateFolderPrimaryExternalLink(ctx, id).FolderLinkRequest(folderLinkRequest).Execute()

Create primary external link



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-folder-primary-external-link/).

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
	id := int32(1) // int32 | The folder ID.
	folderLinkRequest := *openapiclient.NewFolderLinkRequest() // FolderLinkRequest | The folder link parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.CreateFolderPrimaryExternalLink(context.Background(), id).FolderLinkRequest(folderLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.CreateFolderPrimaryExternalLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFolderPrimaryExternalLink`: FileShareWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.CreateFolderPrimaryExternalLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The folder ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFolderPrimaryExternalLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderLinkRequest** | [**FolderLinkRequest**](FolderLinkRequest.md) | The folder link parameters. | 

### Return type

[**FileShareWrapper**](FileShareWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReportFolderHistory

> DocumentBuilderTaskWrapper CreateReportFolderHistory(ctx, folderId).Format(format).From(from).To(to).Execute()

Start the folder history report generation



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-report-folder-history/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	folderId := int32(1) // int32 | The folder ID whose history is exported.
	format := openapiclient.AuditReportFormat(0) // AuditReportFormat | The output file format of the report. Defaults to XLSX. (optional)
	from := time.Now() // time.Time | The start date of the history period to export. (optional)
	to := time.Now() // time.Time | The end date of the history period to export. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.CreateReportFolderHistory(context.Background(), folderId).Format(format).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.CreateReportFolderHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReportFolderHistory`: DocumentBuilderTaskWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.CreateReportFolderHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID whose history is exported. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReportFolderHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**AuditReportFormat**](AuditReportFormat.md) | The output file format of the report. Defaults to XLSX. | 
 **from** | **time.Time** | The start date of the history period to export. | 
 **to** | **time.Time** | The end date of the history period to export. | 

### Return type

[**DocumentBuilderTaskWrapper**](DocumentBuilderTaskWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFolder

> FileOperationArrayWrapper DeleteFolder(ctx, folderId).DeleteFolder(deleteFolder).Execute()

Delete a folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-folder/).

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
	folderId := int32(10) // int32 | The folder ID to delete.
	deleteFolder := *openapiclient.NewDeleteFolder() // DeleteFolder | The parameters for deleting a folder.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.DeleteFolder(context.Background(), folderId).DeleteFolder(deleteFolder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.DeleteFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFolder`: FileOperationArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.DeleteFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteFolder** | [**DeleteFolder**](DeleteFolder.md) | The parameters for deleting a folder. | 

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


## GenerateXlsxByFolder

> XlsxReportResponseWrapper GenerateXlsxByFolder(ctx, folderId).Execute()

Generate XLSX report by folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/generate-xlsx-by-folder/).

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
	folderId := int32(1) // int32 | The folder unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GenerateXlsxByFolder(context.Background(), folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GenerateXlsxByFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateXlsxByFolder`: XlsxReportResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GenerateXlsxByFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateXlsxByFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**XlsxReportResponseWrapper**](XlsxReportResponseWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFavoritesFolder

> FolderContentIntegerWrapper GetFavoritesFolder(ctx).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()

Get the Favorites section



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-favorites-folder/).

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
	userIdOrGroupId := "00000000-0000-0000-0000-000000000000" // string | The user or group ID. (optional)
	filterType := openapiclient.FilterType(0) // FilterType | The filter type. (optional)
	count := int32(25) // int32 | The maximum number of items to retrieve in the request. (optional)
	startIndex := int32(0) // int32 | The zero-based index of the first item to retrieve in a paginated list. (optional)
	sortBy := "DateAndTime" // string | Specifies the field by which the folder content should be sorted. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The order in which the results are sorted. (optional)
	filterValue := "My Document" // string | The text used as a filter or search criterion for folder content queries. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetFavoritesFolder(context.Background()).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetFavoritesFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFavoritesFolder`: FolderContentIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetFavoritesFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFavoritesFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userIdOrGroupId** | **string** | The user or group ID. | 
 **filterType** | [**FilterType**](FilterType.md) | The filter type. | 
 **count** | **int32** | The maximum number of items to retrieve in the request. | 
 **startIndex** | **int32** | The zero-based index of the first item to retrieve in a paginated list. | 
 **sortBy** | **string** | Specifies the field by which the folder content should be sorted. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The order in which the results are sorted. | 
 **filterValue** | **string** | The text used as a filter or search criterion for folder content queries. | 

### Return type

[**FolderContentIntegerWrapper**](FolderContentIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilesUsedSpace

> FilesStatisticsResultWrapper GetFilesUsedSpace(ctx).Execute()

Get used space of files



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-files-used-space/).

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetFilesUsedSpace(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetFilesUsedSpace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilesUsedSpace`: FilesStatisticsResultWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetFilesUsedSpace`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesUsedSpaceRequest struct via the builder pattern


### Return type

[**FilesStatisticsResultWrapper**](FilesStatisticsResultWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolder

> FormsItemArrayWrapper GetFolder(ctx, folderId).Execute()

Get folder form filter



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-folder/).

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
	folderId := int32(1) // int32 | The folder unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetFolder(context.Background(), folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolder`: FormsItemArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FormsItemArrayWrapper**](FormsItemArrayWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolderByFolderId

> FolderContentIntegerWrapper GetFolderByFolderId(ctx, folderId).UserIdOrGroupId(userIdOrGroupId).SharedBy(sharedBy).FilterType(filterType).RoomId(roomId).FolderType(folderType).ExcludeSubject(excludeSubject).ApplyFilterOption(applyFilterOption).WithSubFolders(withSubFolders).Extension(extension).SearchArea(searchArea).FormsItemKey(formsItemKey).FormsItemType(formsItemType).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Location(location).Execute()

Get a folder by ID



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-folder-by-folder-id/).

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
	userIdOrGroupId := "00000000-0000-0000-0000-000000000000" // string | The user or group ID. (optional)
	sharedBy := "00000000-0000-0000-0000-000000000000" // string | The identifier of the user who shared the folder or file. (optional)
	filterType := openapiclient.FilterType(0) // FilterType | The filter type. (optional)
	roomId := int32(1) // int32 | The room ID. (optional)
	folderType := []int32{int32(0)} // []int32 | The parent folder types used to filter the folder contents by folder type. (optional)
	excludeSubject := false // bool | Specifies whether to exclude search by user or group ID. (optional)
	applyFilterOption := openapiclient.ApplyFilterOption(0) // ApplyFilterOption | Specifies whether to return only files, only folders, or all elements from the specified folder. (optional)
	withSubFolders := true // bool | Specifies whether to include files from subfolders in the results. (optional)
	extension := ".docx" // string | Specifies whether to search for the specific file extension. (optional)
	searchArea := openapiclient.SearchArea(0) // SearchArea | The search area. (optional)
	formsItemKey := "doc_key_123" // string | The forms item key. (optional)
	formsItemType := "text" // string | The forms item type. (optional)
	count := int32(25) // int32 | The maximum number of items to retrieve in the request. (optional)
	startIndex := int32(0) // int32 | The zero-based index of the first item to retrieve in a paginated request. (optional)
	sortBy := "DateAndTime" // string | The property used for sorting the folder request results. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The order in which the results are sorted. (optional)
	filterValue := "My Document" // string | The text value used as a filter parameter for folder content queries. (optional)
	location := openapiclient.Location(1) // Location | The location context of the request, specifying the area  where the operation is performed, such as a room, documents, or a link. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetFolderByFolderId(context.Background(), folderId).UserIdOrGroupId(userIdOrGroupId).SharedBy(sharedBy).FilterType(filterType).RoomId(roomId).FolderType(folderType).ExcludeSubject(excludeSubject).ApplyFilterOption(applyFilterOption).WithSubFolders(withSubFolders).Extension(extension).SearchArea(searchArea).FormsItemKey(formsItemKey).FormsItemType(formsItemType).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Location(location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetFolderByFolderId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolderByFolderId`: FolderContentIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetFolderByFolderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderByFolderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userIdOrGroupId** | **string** | The user or group ID. | 
 **sharedBy** | **string** | The identifier of the user who shared the folder or file. | 
 **filterType** | [**FilterType**](FilterType.md) | The filter type. | 
 **roomId** | **int32** | The room ID. | 
 **folderType** | **[]int32** | The parent folder types used to filter the folder contents by folder type. | 
 **excludeSubject** | **bool** | Specifies whether to exclude search by user or group ID. | 
 **applyFilterOption** | [**ApplyFilterOption**](ApplyFilterOption.md) | Specifies whether to return only files, only folders, or all elements from the specified folder. | 
 **withSubFolders** | **bool** | Specifies whether to include files from subfolders in the results. | 
 **extension** | **string** | Specifies whether to search for the specific file extension. | 
 **searchArea** | [**SearchArea**](SearchArea.md) | The search area. | 
 **formsItemKey** | **string** | The forms item key. | 
 **formsItemType** | **string** | The forms item type. | 
 **count** | **int32** | The maximum number of items to retrieve in the request. | 
 **startIndex** | **int32** | The zero-based index of the first item to retrieve in a paginated request. | 
 **sortBy** | **string** | The property used for sorting the folder request results. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The order in which the results are sorted. | 
 **filterValue** | **string** | The text value used as a filter parameter for folder content queries. | 
 **location** | [**Location**](Location.md) | The location context of the request, specifying the area  where the operation is performed, such as a room, documents, or a link. | 

### Return type

[**FolderContentIntegerWrapper**](FolderContentIntegerWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolderHistory

> HistoryArrayWrapper GetFolderHistory(ctx, folderId).FromDate(fromDate).ToDate(toDate).Count(count).StartIndex(startIndex).Execute()

Get folder history



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-folder-history/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	folderId := int32(1) // int32 | The folder ID of the history request.
	fromDate := time.Now() // time.Time | The start date of the history request. (optional)
	toDate := time.Now() // time.Time | The end date of the history request. (optional)
	count := int32(25) // int32 | The number of records to retrieve for the folder history. (optional)
	startIndex := int32(0) // int32 | The starting index from which the history records are retrieved in the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetFolderHistory(context.Background(), folderId).FromDate(fromDate).ToDate(toDate).Count(count).StartIndex(startIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetFolderHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolderHistory`: HistoryArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetFolderHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID of the history request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **time.Time** | The start date of the history request. | 
 **toDate** | **time.Time** | The end date of the history request. | 
 **count** | **int32** | The number of records to retrieve for the folder history. | 
 **startIndex** | **int32** | The starting index from which the history records are retrieved in the request. | 

### Return type

[**HistoryArrayWrapper**](HistoryArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolderInfo

> FolderIntegerWrapper GetFolderInfo(ctx, folderId).Execute()

Get folder information



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-folder-info/).

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
	folderId := int32(1) // int32 | The folder unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetFolderInfo(context.Background(), folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetFolderInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolderInfo`: FolderIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetFolderInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderIntegerWrapper**](FolderIntegerWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolderLinks

> FileShareArrayWrapper GetFolderLinks(ctx, id).Execute()

Get the folder links



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-folder-links/).

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
	id := int32(1) // int32 | The folder ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetFolderLinks(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetFolderLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolderLinks`: FileShareArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetFolderLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The folder ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileShareArrayWrapper**](FileShareArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolderPath

> FileEntryBaseArrayWrapper GetFolderPath(ctx, folderId).Execute()

Get the folder path



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-folder-path/).

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
	folderId := int32(1) // int32 | The folder unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetFolderPath(context.Background(), folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetFolderPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolderPath`: FileEntryBaseArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetFolderPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetFolderPrimaryExternalLink

> FileShareWrapper GetFolderPrimaryExternalLink(ctx, id).Count(count).StartIndex(startIndex).Execute()

Get primary external link



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-folder-primary-external-link/).

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
	id := int32(10) // int32 | The folder unique identifier.
	count := int32(25) // int32 | The number of items to retrieve in the request. (optional)
	startIndex := int32(0) // int32 | The starting index for the query results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetFolderPrimaryExternalLink(context.Background(), id).Count(count).StartIndex(startIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetFolderPrimaryExternalLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolderPrimaryExternalLink`: FileShareWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetFolderPrimaryExternalLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The folder unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderPrimaryExternalLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The number of items to retrieve in the request. | 
 **startIndex** | **int32** | The starting index for the query results. | 

### Return type

[**FileShareWrapper**](FileShareWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolders

> FileEntryBaseArrayWrapper GetFolders(ctx, folderId).Execute()

Get subfolders



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-folders/).

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
	folderId := int32(1) // int32 | The folder unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetFolders(context.Background(), folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolders`: FileEntryBaseArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetFormsFolder

> FolderContentIntegerWrapper GetFormsFolder(ctx).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()

Get the Forms section



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-forms-folder/).

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
	userIdOrGroupId := "00000000-0000-0000-0000-000000000000" // string | The user or group ID. (optional)
	filterType := openapiclient.FilterType(0) // FilterType | The filter type. (optional)
	count := int32(25) // int32 | The maximum number of items to retrieve in the request. (optional)
	startIndex := int32(0) // int32 | The zero-based index of the first item to retrieve in a paginated list. (optional)
	sortBy := "DateAndTime" // string | Specifies the field by which the folder content should be sorted. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The order in which the results are sorted. (optional)
	filterValue := "My Document" // string | The text used as a filter or search criterion for folder content queries. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetFormsFolder(context.Background()).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetFormsFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFormsFolder`: FolderContentIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetFormsFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFormsFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userIdOrGroupId** | **string** | The user or group ID. | 
 **filterType** | [**FilterType**](FilterType.md) | The filter type. | 
 **count** | **int32** | The maximum number of items to retrieve in the request. | 
 **startIndex** | **int32** | The zero-based index of the first item to retrieve in a paginated list. | 
 **sortBy** | **string** | Specifies the field by which the folder content should be sorted. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The order in which the results are sorted. | 
 **filterValue** | **string** | The text used as a filter or search criterion for folder content queries. | 

### Return type

[**FolderContentIntegerWrapper**](FolderContentIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyFolder

> FolderContentIntegerWrapper GetMyFolder(ctx).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).ApplyFilterOption(applyFilterOption).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()

Get the My documents section



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-my-folder/).

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
	userIdOrGroupId := "00000000-0000-0000-0000-000000000000" // string | The user or group ID. (optional)
	filterType := openapiclient.FilterType(0) // FilterType | The filter type. (optional)
	applyFilterOption := openapiclient.ApplyFilterOption(0) // ApplyFilterOption | Specifies whether to return only files, only folders or all elements. (optional)
	count := int32(25) // int32 | The maximum number of items to retrieve in the response. (optional)
	startIndex := int32(0) // int32 | The starting position of the items to be retrieved. (optional)
	sortBy := "DateAndTime" // string | The property used to specify the sorting criteria for folder contents. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The order in which the results are sorted. (optional)
	filterValue := "My Document" // string | The text used for filtering or searching folder contents. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetMyFolder(context.Background()).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).ApplyFilterOption(applyFilterOption).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetMyFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyFolder`: FolderContentIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetMyFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userIdOrGroupId** | **string** | The user or group ID. | 
 **filterType** | [**FilterType**](FilterType.md) | The filter type. | 
 **applyFilterOption** | [**ApplyFilterOption**](ApplyFilterOption.md) | Specifies whether to return only files, only folders or all elements. | 
 **count** | **int32** | The maximum number of items to retrieve in the response. | 
 **startIndex** | **int32** | The starting position of the items to be retrieved. | 
 **sortBy** | **string** | The property used to specify the sorting criteria for folder contents. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The order in which the results are sorted. | 
 **filterValue** | **string** | The text used for filtering or searching folder contents. | 

### Return type

[**FolderContentIntegerWrapper**](FolderContentIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNewFolderItems

> FileEntryBaseArrayWrapper GetNewFolderItems(ctx, folderId).Execute()

Get new folder items



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-new-folder-items/).

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
	folderId := int32(1) // int32 | The folder unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetNewFolderItems(context.Background(), folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetNewFolderItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNewFolderItems`: FileEntryBaseArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetNewFolderItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNewFolderItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetRecentFolder

> FolderContentIntegerWrapper GetRecentFolder(ctx).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).ExcludeSubject(excludeSubject).ApplyFilterOption(applyFilterOption).SearchArea(searchArea).Extension(extension).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()

Get the Recent section



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-recent-folder/).

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
	userIdOrGroupId := "00000000-0000-0000-0000-000000000000" // string | The user or group ID. (optional)
	filterType := openapiclient.FilterType(0) // FilterType | The filter type. (optional)
	excludeSubject := false // bool | Specifies whether to exclude search by user or group ID. (optional)
	applyFilterOption := openapiclient.ApplyFilterOption(0) // ApplyFilterOption | Specifies whether to return only files, only folders or all elements. (optional)
	searchArea := openapiclient.SearchArea(0) // SearchArea | The search area. (optional)
	extension := []string{"Inner_example"} // []string | Specifies whether to search for a specific file extension in the Recent folder. (optional)
	count := int32(25) // int32 | The maximum number of items to return. (optional)
	startIndex := int32(0) // int32 | The starting position of the results to be returned in the query response. (optional)
	sortBy := "DateAndTime" // string | Specifies the sorting criteria for the folder request. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The order in which the results are sorted. (optional)
	filterValue := "My Document" // string | The text used for filtering or searching folder contents. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetRecentFolder(context.Background()).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).ExcludeSubject(excludeSubject).ApplyFilterOption(applyFilterOption).SearchArea(searchArea).Extension(extension).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetRecentFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecentFolder`: FolderContentIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetRecentFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecentFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userIdOrGroupId** | **string** | The user or group ID. | 
 **filterType** | [**FilterType**](FilterType.md) | The filter type. | 
 **excludeSubject** | **bool** | Specifies whether to exclude search by user or group ID. | 
 **applyFilterOption** | [**ApplyFilterOption**](ApplyFilterOption.md) | Specifies whether to return only files, only folders or all elements. | 
 **searchArea** | [**SearchArea**](SearchArea.md) | The search area. | 
 **extension** | **[]string** | Specifies whether to search for a specific file extension in the Recent folder. | 
 **count** | **int32** | The maximum number of items to return. | 
 **startIndex** | **int32** | The starting position of the results to be returned in the query response. | 
 **sortBy** | **string** | Specifies the sorting criteria for the folder request. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The order in which the results are sorted. | 
 **filterValue** | **string** | The text used for filtering or searching folder contents. | 

### Return type

[**FolderContentIntegerWrapper**](FolderContentIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportFolderHistory

> DocumentBuilderTaskWrapper GetReportFolderHistory(ctx, folderId).Execute()

Get the folder history report generation status



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-report-folder-history/).

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
	folderId := int32(56) // int32 | The folder unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetReportFolderHistory(context.Background(), folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetReportFolderHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportFolderHistory`: DocumentBuilderTaskWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetReportFolderHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportFolderHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DocumentBuilderTaskWrapper**](DocumentBuilderTaskWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRootFolders

> FolderContentIntegerArrayWrapper GetRootFolders(ctx).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).WithoutTrash(withoutTrash).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()

Get filtered sections



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-root-folders/).

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
	userIdOrGroupId := "00000000-0000-0000-0000-000000000000" // string | The user or group ID. (optional)
	filterType := openapiclient.FilterType(0) // FilterType | The filter type. (optional)
	withoutTrash := false // bool | Specifies whether to return the Trash section or not. (optional)
	count := int32(25) // int32 | The maximum number of items to retrieve in the response. (optional)
	startIndex := int32(0) // int32 | The starting position of the items to be retrieved. (optional)
	sortBy := "DateAndTime" // string | Specifies the field by which the folder content should be sorted. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The order in which the results are sorted. (optional)
	filterValue := "My Document" // string | The text used as a filter for searching or retrieving folder contents. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetRootFolders(context.Background()).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).WithoutTrash(withoutTrash).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetRootFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRootFolders`: FolderContentIntegerArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetRootFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRootFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userIdOrGroupId** | **string** | The user or group ID. | 
 **filterType** | [**FilterType**](FilterType.md) | The filter type. | 
 **withoutTrash** | **bool** | Specifies whether to return the Trash section or not. | 
 **count** | **int32** | The maximum number of items to retrieve in the response. | 
 **startIndex** | **int32** | The starting position of the items to be retrieved. | 
 **sortBy** | **string** | Specifies the field by which the folder content should be sorted. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The order in which the results are sorted. | 
 **filterValue** | **string** | The text used as a filter for searching or retrieving folder contents. | 

### Return type

[**FolderContentIntegerArrayWrapper**](FolderContentIntegerArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrashFolder

> FolderContentIntegerWrapper GetTrashFolder(ctx).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).ApplyFilterOption(applyFilterOption).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()

Get the Trash section



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-trash-folder/).

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
	userIdOrGroupId := "00000000-0000-0000-0000-000000000000" // string | The user or group ID. (optional)
	filterType := openapiclient.FilterType(0) // FilterType | The filter type. (optional)
	applyFilterOption := openapiclient.ApplyFilterOption(0) // ApplyFilterOption | Specifies whether to return only files, only folders or all elements. (optional)
	count := int32(25) // int32 | The maximum number of items to retrieve in the response. (optional)
	startIndex := int32(0) // int32 | The starting position of the items to be retrieved. (optional)
	sortBy := "DateAndTime" // string | The property used to specify the sorting criteria for folder contents. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The order in which the results are sorted. (optional)
	filterValue := "My Document" // string | The text used for filtering or searching folder contents. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetTrashFolder(context.Background()).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).ApplyFilterOption(applyFilterOption).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetTrashFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrashFolder`: FolderContentIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetTrashFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrashFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userIdOrGroupId** | **string** | The user or group ID. | 
 **filterType** | [**FilterType**](FilterType.md) | The filter type. | 
 **applyFilterOption** | [**ApplyFilterOption**](ApplyFilterOption.md) | Specifies whether to return only files, only folders or all elements. | 
 **count** | **int32** | The maximum number of items to retrieve in the response. | 
 **startIndex** | **int32** | The starting position of the items to be retrieved. | 
 **sortBy** | **string** | The property used to specify the sorting criteria for folder contents. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The order in which the results are sorted. | 
 **filterValue** | **string** | The text used for filtering or searching folder contents. | 

### Return type

[**FolderContentIntegerWrapper**](FolderContentIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsertFile

> FileIntegerWrapper InsertFile(ctx, folderId).InsertFileFile(insertFileFile).InsertFileTitle(insertFileTitle).InsertFileCreateNewIfExist(insertFileCreateNewIfExist).InsertFileKeepConvertStatus(insertFileKeepConvertStatus).InsertFileStreamCanRead(insertFileStreamCanRead).InsertFileStreamCanWrite(insertFileStreamCanWrite).InsertFileStreamCanSeek(insertFileStreamCanSeek).InsertFileStreamCanTimeout(insertFileStreamCanTimeout).InsertFileStreamLength(insertFileStreamLength).InsertFileStreamPosition(insertFileStreamPosition).InsertFileStreamReadTimeout(insertFileStreamReadTimeout).InsertFileStreamWriteTimeout(insertFileStreamWriteTimeout).Execute()

Insert a file



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/insert-file/).

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
	folderId := int32(1) // int32 | The folder ID for inserting a file.
	insertFileFile := os.NewFile(1234, "some_file") // *os.File | The file to be inserted. (optional)
	insertFileTitle := "insertFileTitle_example" // string | The file title to be inserted. (optional)
	insertFileCreateNewIfExist := true // bool | Specifies whether to create a new file if it already exists or not. (optional)
	insertFileKeepConvertStatus := true // bool | Specifies whether to keep the file converting status or not. (optional)
	insertFileStreamCanRead := true // bool |  (optional)
	insertFileStreamCanWrite := true // bool |  (optional)
	insertFileStreamCanSeek := true // bool |  (optional)
	insertFileStreamCanTimeout := true // bool |  (optional)
	insertFileStreamLength := int64(789) // int64 |  (optional)
	insertFileStreamPosition := int64(789) // int64 |  (optional)
	insertFileStreamReadTimeout := int32(56) // int32 |  (optional)
	insertFileStreamWriteTimeout := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.InsertFile(context.Background(), folderId).InsertFileFile(insertFileFile).InsertFileTitle(insertFileTitle).InsertFileCreateNewIfExist(insertFileCreateNewIfExist).InsertFileKeepConvertStatus(insertFileKeepConvertStatus).InsertFileStreamCanRead(insertFileStreamCanRead).InsertFileStreamCanWrite(insertFileStreamCanWrite).InsertFileStreamCanSeek(insertFileStreamCanSeek).InsertFileStreamCanTimeout(insertFileStreamCanTimeout).InsertFileStreamLength(insertFileStreamLength).InsertFileStreamPosition(insertFileStreamPosition).InsertFileStreamReadTimeout(insertFileStreamReadTimeout).InsertFileStreamWriteTimeout(insertFileStreamWriteTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.InsertFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InsertFile`: FileIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.InsertFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID for inserting a file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsertFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **insertFileFile** | ***os.File** | The file to be inserted. | 
 **insertFileTitle** | **string** | The file title to be inserted. | 
 **insertFileCreateNewIfExist** | **bool** | Specifies whether to create a new file if it already exists or not. | 
 **insertFileKeepConvertStatus** | **bool** | Specifies whether to keep the file converting status or not. | 
 **insertFileStreamCanRead** | **bool** |  | 
 **insertFileStreamCanWrite** | **bool** |  | 
 **insertFileStreamCanSeek** | **bool** |  | 
 **insertFileStreamCanTimeout** | **bool** |  | 
 **insertFileStreamLength** | **int64** |  | 
 **insertFileStreamPosition** | **int64** |  | 
 **insertFileStreamReadTimeout** | **int32** |  | 
 **insertFileStreamWriteTimeout** | **int32** |  | 

### Return type

[**FileIntegerWrapper**](FileIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsertFileToMyFromBody

> FileIntegerWrapper InsertFileToMyFromBody(ctx).File(file).Title(title).CreateNewIfExist(createNewIfExist).KeepConvertStatus(keepConvertStatus).StreamCanRead(streamCanRead).StreamCanWrite(streamCanWrite).StreamCanSeek(streamCanSeek).StreamCanTimeout(streamCanTimeout).StreamLength(streamLength).StreamPosition(streamPosition).StreamReadTimeout(streamReadTimeout).StreamWriteTimeout(streamWriteTimeout).Execute()

Insert a file to the My documents section



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/insert-file-to-my-from-body/).

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
	file := os.NewFile(1234, "some_file") // *os.File | The file to be inserted. (optional)
	title := "title_example" // string | The file title to be inserted. (optional)
	createNewIfExist := true // bool | Specifies whether to create a new file if it already exists or not. (optional)
	keepConvertStatus := true // bool | Specifies whether to keep the file converting status or not. (optional)
	streamCanRead := true // bool |  (optional)
	streamCanWrite := true // bool |  (optional)
	streamCanSeek := true // bool |  (optional)
	streamCanTimeout := true // bool |  (optional)
	streamLength := int64(789) // int64 |  (optional)
	streamPosition := int64(789) // int64 |  (optional)
	streamReadTimeout := int32(56) // int32 |  (optional)
	streamWriteTimeout := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.InsertFileToMyFromBody(context.Background()).File(file).Title(title).CreateNewIfExist(createNewIfExist).KeepConvertStatus(keepConvertStatus).StreamCanRead(streamCanRead).StreamCanWrite(streamCanWrite).StreamCanSeek(streamCanSeek).StreamCanTimeout(streamCanTimeout).StreamLength(streamLength).StreamPosition(streamPosition).StreamReadTimeout(streamReadTimeout).StreamWriteTimeout(streamWriteTimeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.InsertFileToMyFromBody``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InsertFileToMyFromBody`: FileIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.InsertFileToMyFromBody`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsertFileToMyFromBodyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The file to be inserted. | 
 **title** | **string** | The file title to be inserted. | 
 **createNewIfExist** | **bool** | Specifies whether to create a new file if it already exists or not. | 
 **keepConvertStatus** | **bool** | Specifies whether to keep the file converting status or not. | 
 **streamCanRead** | **bool** |  | 
 **streamCanWrite** | **bool** |  | 
 **streamCanSeek** | **bool** |  | 
 **streamCanTimeout** | **bool** |  | 
 **streamLength** | **int64** |  | 
 **streamPosition** | **int64** |  | 
 **streamReadTimeout** | **int32** |  | 
 **streamWriteTimeout** | **int32** |  | 

### Return type

[**FileIntegerWrapper**](FileIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameFolder

> FolderIntegerWrapper RenameFolder(ctx, folderId).CreateFolder(createFolder).Execute()

Rename a folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/rename-folder/).

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
	folderId := int32(1) // int32 | The folder ID for the folder creation.
	createFolder := *openapiclient.NewCreateFolder("New Folder") // CreateFolder | The parameters for creating a folder.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.RenameFolder(context.Background(), folderId).CreateFolder(createFolder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.RenameFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenameFolder`: FolderIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.RenameFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID for the folder creation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFolder** | [**CreateFolder**](CreateFolder.md) | The parameters for creating a folder. | 

### Return type

[**FolderIntegerWrapper**](FolderIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFolderOrder

> FolderIntegerWrapper SetFolderOrder(ctx, folderId).OrderRequestDto(orderRequestDto).Execute()

Set folder order



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-folder-order/).

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
	folderId := int32(1) // int32 | The folder unique identifier.
	orderRequestDto := *openapiclient.NewOrderRequestDto() // OrderRequestDto | The folder order information. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.SetFolderOrder(context.Background(), folderId).OrderRequestDto(orderRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.SetFolderOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFolderOrder`: FolderIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.SetFolderOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFolderOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderRequestDto** | [**OrderRequestDto**](OrderRequestDto.md) | The folder order information. | 

### Return type

[**FolderIntegerWrapper**](FolderIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFolderPrimaryExternalLink

> FileShareWrapper SetFolderPrimaryExternalLink(ctx, id).FolderLinkRequest(folderLinkRequest).Execute()

Set the folder external link



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-folder-primary-external-link/).

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
	id := int32(1) // int32 | The folder ID.
	folderLinkRequest := *openapiclient.NewFolderLinkRequest() // FolderLinkRequest | The folder link parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.SetFolderPrimaryExternalLink(context.Background(), id).FolderLinkRequest(folderLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.SetFolderPrimaryExternalLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFolderPrimaryExternalLink`: FileShareWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.SetFolderPrimaryExternalLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The folder ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFolderPrimaryExternalLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderLinkRequest** | [**FolderLinkRequest**](FolderLinkRequest.md) | The folder link parameters. | 

### Return type

[**FileShareWrapper**](FileShareWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateReportFolderHistory

> TerminateReportFolderHistory(ctx, folderId).Execute()

Terminate the folder history report generation



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/terminate-report-folder-history/).

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
	folderId := int32(56) // int32 | The folder unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesFoldersAPI.TerminateReportFolderHistory(context.Background(), folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.TerminateReportFolderHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateReportFolderHistoryRequest struct via the builder pattern


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


## UploadFile

> FileIntegerArrayWrapper UploadFile(ctx, folderId).CreateNewIfExist(createNewIfExist).StoreOriginalFile(storeOriginalFile).KeepConvertStatus(keepConvertStatus).File(file).Execute()

Upload a file



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/upload-file/).

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
	folderId := int32(1) // int32 | The folder ID to upload a file.
	createNewIfExist := true // bool | Specifies whether to create the new file if it already exists or not. (optional)
	storeOriginalFile := true // bool | Specifies whether to upload documents in the original formats as well or not. (optional)
	keepConvertStatus := false // bool | Specifies whether to keep the file converting status or not. (optional)
	file := os.NewFile(1234, "some_file") // *os.File | The file to be uploaded. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.UploadFile(context.Background(), folderId).CreateNewIfExist(createNewIfExist).StoreOriginalFile(storeOriginalFile).KeepConvertStatus(keepConvertStatus).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.UploadFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadFile`: FileIntegerArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.UploadFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID to upload a file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNewIfExist** | **bool** | Specifies whether to create the new file if it already exists or not. | 
 **storeOriginalFile** | **bool** | Specifies whether to upload documents in the original formats as well or not. | 
 **keepConvertStatus** | **bool** | Specifies whether to keep the file converting status or not. | 
 **file** | ***os.File** | The file to be uploaded. | 

### Return type

[**FileIntegerArrayWrapper**](FileIntegerArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFileToMy

> FileIntegerArrayWrapper UploadFileToMy(ctx).CreateNewIfExist(createNewIfExist).StoreOriginalFile(storeOriginalFile).KeepConvertStatus(keepConvertStatus).File(file).Execute()

Upload a file to the My documents section



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/upload-file-to-my/).

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
	createNewIfExist := true // bool | Specifies whether to create the new file if it already exists or not. (optional)
	storeOriginalFile := true // bool | Specifies whether to upload documents in the original formats as well or not. (optional)
	keepConvertStatus := false // bool | Specifies whether to keep the file converting status or not. (optional)
	file := os.NewFile(1234, "some_file") // *os.File | The file to be uploaded. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.UploadFileToMy(context.Background()).CreateNewIfExist(createNewIfExist).StoreOriginalFile(storeOriginalFile).KeepConvertStatus(keepConvertStatus).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.UploadFileToMy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadFileToMy`: FileIntegerArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.UploadFileToMy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileToMyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNewIfExist** | **bool** | Specifies whether to create the new file if it already exists or not. | 
 **storeOriginalFile** | **bool** | Specifies whether to upload documents in the original formats as well or not. | 
 **keepConvertStatus** | **bool** | Specifies whether to keep the file converting status or not. | 
 **file** | ***os.File** | The file to be uploaded. | 

### Return type

[**FileIntegerArrayWrapper**](FileIntegerArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

