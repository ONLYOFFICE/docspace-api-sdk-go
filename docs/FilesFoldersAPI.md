# \FilesFoldersAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckUpload**](FilesFoldersAPI.md#CheckUpload) | **Post** /api/2.0/files/{folderId}/upload/check | Check for upload conflicts
[**CreateFolder**](FilesFoldersAPI.md#CreateFolder) | **Post** /api/2.0/files/folder/{folderId} | Create a folder
[**CreateFolderPrimaryExternalLink**](FilesFoldersAPI.md#CreateFolderPrimaryExternalLink) | **Post** /api/2.0/files/folder/{id}/link | Create the folder primary external link
[**CreateReportFolderHistory**](FilesFoldersAPI.md#CreateReportFolderHistory) | **Post** /api/2.0/files/folder/{folderId}/log/report | Start the folder history report generation
[**DeleteFolder**](FilesFoldersAPI.md#DeleteFolder) | **Delete** /api/2.0/files/folder/{folderId} | Delete a folder
[**GenerateXlsxByFolder**](FilesFoldersAPI.md#GenerateXlsxByFolder) | **Post** /api/2.0/files/folder/{folderId}/xlsx | Generate XLSX report by folder
[**GetFavoritesFolder**](FilesFoldersAPI.md#GetFavoritesFolder) | **Get** /api/2.0/files/@favorites | Get the Favorites section
[**GetFilesUsedSpace**](FilesFoldersAPI.md#GetFilesUsedSpace) | **Get** /api/2.0/files/filesusedspace | Get used space of files
[**GetFolder**](FilesFoldersAPI.md#GetFolder) | **Get** /api/2.0/files/{folderId}/formfilter | Get folder form filter
[**GetFolderByFolderId**](FilesFoldersAPI.md#GetFolderByFolderId) | **Get** /api/2.0/files/{folderId} | Get a folder by ID
[**GetFolderHistory**](FilesFoldersAPI.md#GetFolderHistory) | **Get** /api/2.0/files/folder/{folderId}/log | Get folder history
[**GetFolderInfo**](FilesFoldersAPI.md#GetFolderInfo) | **Get** /api/2.0/files/folder/{folderId} | Get folder information
[**GetFolderLinks**](FilesFoldersAPI.md#GetFolderLinks) | **Get** /api/2.0/files/folder/{id}/links | Get folder external links
[**GetFolderPath**](FilesFoldersAPI.md#GetFolderPath) | **Get** /api/2.0/files/folder/{folderId}/path | Get the folder path
[**GetFolderPrimaryExternalLink**](FilesFoldersAPI.md#GetFolderPrimaryExternalLink) | **Get** /api/2.0/files/folder/{id}/link | Get the folder primary external link
[**GetFolders**](FilesFoldersAPI.md#GetFolders) | **Get** /api/2.0/files/{folderId}/subfolders | Get subfolders
[**GetFormsFolder**](FilesFoldersAPI.md#GetFormsFolder) | **Get** /api/2.0/files/@forms | Get the Forms section
[**GetMyFolder**](FilesFoldersAPI.md#GetMyFolder) | **Get** /api/2.0/files/@my | Get the My documents section
[**GetNewFolderItems**](FilesFoldersAPI.md#GetNewFolderItems) | **Get** /api/2.0/files/{folderId}/news | Get new folder items
[**GetRecentFolder**](FilesFoldersAPI.md#GetRecentFolder) | **Get** /api/2.0/files/recent | Get the Recent section
[**GetReportFolderHistory**](FilesFoldersAPI.md#GetReportFolderHistory) | **Get** /api/2.0/files/folder/{folderId}/log/report | Get the folder history report generation status
[**GetRootFolders**](FilesFoldersAPI.md#GetRootFolders) | **Get** /api/2.0/files/@root | Get filtered sections
[**GetTrashFolder**](FilesFoldersAPI.md#GetTrashFolder) | **Get** /api/2.0/files/@trash | Get the Trash section
[**InsertFile**](FilesFoldersAPI.md#InsertFile) | **Post** /api/2.0/files/{folderId}/insert | Insert a file
[**InsertFileToMyFromBody**](FilesFoldersAPI.md#InsertFileToMyFromBody) | **Post** /api/2.0/files/@my/insert | Insert a file into My documents
[**RenameFolder**](FilesFoldersAPI.md#RenameFolder) | **Put** /api/2.0/files/folder/{folderId} | Rename a folder
[**SetFolderOrder**](FilesFoldersAPI.md#SetFolderOrder) | **Put** /api/2.0/files/folder/{folderId}/order | Set folder order
[**SetFolderPrimaryExternalLink**](FilesFoldersAPI.md#SetFolderPrimaryExternalLink) | **Put** /api/2.0/files/folder/{id}/links | Set the folder external link
[**TerminateReportFolderHistory**](FilesFoldersAPI.md#TerminateReportFolderHistory) | **Delete** /api/2.0/files/folder/{folderId}/log/report | Terminate the folder history report generation
[**UploadFile**](FilesFoldersAPI.md#UploadFile) | **Post** /api/2.0/files/{folderId}/upload | Upload a file
[**UploadFileToMy**](FilesFoldersAPI.md#UploadFileToMy) | **Post** /api/2.0/files/@my/upload | Upload a file to My documents



## CheckUpload

> STRINGArrayWrapper CheckUpload(ctx, folderId).CheckUploadRequest(checkUploadRequest).Execute()

Check for upload conflicts



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
	folderId := int32(1) // int32 | The folder whose contents the names are tested against; take the id from a listing such as  `GET api/2.0/files/@root`.
	checkUploadRequest := *openapiclient.NewCheckUploadRequest() // CheckUploadRequest | The names to test against the files the folder already holds.

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
**folderId** | **int32** | The folder whose contents the names are tested against; take the id from a listing such as  `GET api/2.0/files/@root`. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkUploadRequest** | [**CheckUploadRequest**](CheckUploadRequest.md) | The names to test against the files the folder already holds. | 

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

> FolderWrapper CreateFolder(ctx, folderId).CreateFolder(createFolder).Execute()

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
	folderId := int32(1) // int32 | The folder the request is addressed to: when a folder is created it is the parent that receives the new  folder, and when a folder is renamed it is the folder that gets the new title.
	createFolder := *openapiclient.NewCreateFolder("New Folder") // CreateFolder | The title carried by the request body.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.CreateFolder(context.Background(), folderId).CreateFolder(createFolder).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// folderId := "sbox-42"
	// thirdPartyResp, r, err := apiClient.FilesFoldersAPI.CreateFolder(context.Background(), folderId).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.CreateFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFolder`: FolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.CreateFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder the request is addressed to: when a folder is created it is the parent that receives the new  folder, and when a folder is renamed it is the folder that gets the new title. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFolder** | [**CreateFolder**](CreateFolder.md) | The title carried by the request body. | 

### Return type

[**FolderWrapper**](FolderWrapper.md)

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

Create the folder primary external link



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
	id := int32(1) // int32 | The folder or room the link belongs to.
	folderLinkRequest := *openapiclient.NewFolderLinkRequest() // FolderLinkRequest | The link and the way it is to be shaped.

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
**id** | **int32** | The folder or room the link belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFolderPrimaryExternalLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderLinkRequest** | [**FolderLinkRequest**](FolderLinkRequest.md) | The link and the way it is to be shaped. | 

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
	folderId := int32(1) // int32 | The folder whose history is exported; the report covers the folder itself and the entries inside it.
	format := openapiclient.AuditReportFormat(0) // AuditReportFormat | The shape the report is written in: `Xlsx` produces a spreadsheet that is saved as a file of the portal, while  `Csv` produces a comma-separated text file that is uploaded to My documents without being reported back with  a file identifier. (optional)
	from := time.Now() // time.Time | The earliest moment an exported entry may have, read in the time zone of the portal; left out, the report  starts at the oldest entry the portal still keeps. (optional)
	to := time.Now() // time.Time | The latest moment an exported entry may have, read in the time zone of the portal; left out, the report ends  at the newest entry. (optional)

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
**folderId** | **int32** | The folder whose history is exported; the report covers the folder itself and the entries inside it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReportFolderHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**AuditReportFormat**](AuditReportFormat.md) | The shape the report is written in: `Xlsx` produces a spreadsheet that is saved as a file of the portal, while  `Csv` produces a comma-separated text file that is uploaded to My documents without being reported back with  a file identifier. | 
 **from** | **time.Time** | The earliest moment an exported entry may have, read in the time zone of the portal; left out, the report  starts at the oldest entry the portal still keeps. | 
 **to** | **time.Time** | The latest moment an exported entry may have, read in the time zone of the portal; left out, the report ends  at the newest entry. | 

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
	folderId := int32(10) // int32 | The folder to delete, together with everything it holds.
	deleteFolder := *openapiclient.NewDeleteFolder() // DeleteFolder | How the deletion is to be carried out.

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
**folderId** | **int32** | The folder to delete, together with everything it holds. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteFolder** | [**DeleteFolder**](DeleteFolder.md) | How the deletion is to be carried out. | 

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
	folderId := int32(1) // int32 | The folder the operation acts on. Take the identifier from a listing such as `GET api/2.0/files/@root` or  `GET api/2.0/files/{folderId}`: a folder stored in the portal is numbered, while a folder in a connected  third-party account is named by an opaque string.

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
**folderId** | **int32** | The folder the operation acts on. Take the identifier from a listing such as `GET api/2.0/files/@root` or  `GET api/2.0/files/{folderId}`: a folder stored in the portal is numbered, while a folder in a connected  third-party account is named by an opaque string. | 

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

> FolderContentWrapper GetFavoritesFolder(ctx).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()

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
	userIdOrGroupId := "00000000-0000-0000-0000-000000000000" // string | Restricts the listing to the entries authored by this portal member, or by the members of this group; the same  parameter accepts either kind of identifier. Omit it to list everything the caller can read. (optional)
	filterType := openapiclient.FilterType(0) // FilterType | Narrows the listing to a single kind of entry, such as documents, images or one type of room. Omit it to list  every kind the section holds. (optional)
	count := int32(25) // int32 | The size of one page of section content. Pair it with `startIndex` to walk the listing, and compare the two  with `total` in the response to see when the last page has been read. (optional)
	startIndex := int32(0) // int32 | The number of matching entries to skip before the returned page begins; add `count` to it to ask for the next  page. (optional)
	sortBy := "DateAndTime" // string | The name of the field the entries are ordered by, matched case-insensitively against the file sort fields:  `DateAndTime`, `AZ`, `Size`, `Author`, `Type`, `New`, `DateAndTimeCreation`, `RoomType`, `Tags`, `Room`,  `CustomOrder`, `LastOpened` and `UsedSpace`. A recognized value is also saved as the default order of the  account and reused by later listings that omit the parameter, while a value matching none of the fields leaves  that saved order in place. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The direction in which the `sortBy` field is ordered. It is saved together with `sortBy` as the default order  of the account. (optional)
	filterValue := "My Document" // string | The search string the section is filtered by: it is matched as a substring of entry titles and, for files,  against the indexed document content as well. Omit it to list the section unfiltered. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetFavoritesFolder(context.Background()).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetFavoritesFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFavoritesFolder`: FolderContentWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetFavoritesFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFavoritesFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userIdOrGroupId** | **string** | Restricts the listing to the entries authored by this portal member, or by the members of this group; the same  parameter accepts either kind of identifier. Omit it to list everything the caller can read. | 
 **filterType** | [**FilterType**](FilterType.md) | Narrows the listing to a single kind of entry, such as documents, images or one type of room. Omit it to list  every kind the section holds. | 
 **count** | **int32** | The size of one page of section content. Pair it with `startIndex` to walk the listing, and compare the two  with `total` in the response to see when the last page has been read. | 
 **startIndex** | **int32** | The number of matching entries to skip before the returned page begins; add `count` to it to ask for the next  page. | 
 **sortBy** | **string** | The name of the field the entries are ordered by, matched case-insensitively against the file sort fields:  `DateAndTime`, `AZ`, `Size`, `Author`, `Type`, `New`, `DateAndTimeCreation`, `RoomType`, `Tags`, `Room`,  `CustomOrder`, `LastOpened` and `UsedSpace`. A recognized value is also saved as the default order of the  account and reused by later listings that omit the parameter, while a value matching none of the fields leaves  that saved order in place. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The direction in which the `sortBy` field is ordered. It is saved together with `sortBy` as the default order  of the account. | 
 **filterValue** | **string** | The search string the section is filtered by: it is matched as a substring of entry titles and, for files,  against the indexed document content as well. Omit it to list the section unfiltered. | 

### Return type

[**FolderContentWrapper**](FolderContentWrapper.md)

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
	folderId := int32(1) // int32 | The folder the operation acts on. Take the identifier from a listing such as `GET api/2.0/files/@root` or  `GET api/2.0/files/{folderId}`: a folder stored in the portal is numbered, while a folder in a connected  third-party account is named by an opaque string.

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
**folderId** | **int32** | The folder the operation acts on. Take the identifier from a listing such as `GET api/2.0/files/@root` or  `GET api/2.0/files/{folderId}`: a folder stored in the portal is numbered, while a folder in a connected  third-party account is named by an opaque string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FormsItemArrayWrapper**](FormsItemArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolderByFolderId

> FolderContentWrapper GetFolderByFolderId(ctx, folderId).UserIdOrGroupId(userIdOrGroupId).SharedBy(sharedBy).FilterType(filterType).RoomId(roomId).FolderType(folderType).ExcludeSubject(excludeSubject).ApplyFilterOption(applyFilterOption).WithSubFolders(withSubFolders).Extension(extension).SearchArea(searchArea).FormsItemKey(formsItemKey).FormsItemType(formsItemType).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Location(location).Execute()

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
	folderId := int32(1) // int32 | The folder whose contents are listed. Each section root has an operation of its own, such as  `GET api/2.0/files/@my`, and every other folder is opened by the identifier a listing gave for it.
	userIdOrGroupId := "00000000-0000-0000-0000-000000000000" // string | Restricts the listing to the entries authored by this portal member, or by the members of this group; the same  parameter accepts either kind of identifier. Omit it to list everything the caller can read. (optional)
	sharedBy := "00000000-0000-0000-0000-000000000000" // string | Restricts the listing to the entries this member shared, which narrows a shared listing down to what one  person handed out. (optional)
	filterType := openapiclient.FilterType(0) // FilterType | Narrows the listing to a single kind of entry, such as documents, spreadsheets, images or one type of room.  Omit it to list every kind the folder holds. (optional)
	roomId := int32(1) // int32 | Keeps only the entries that lie in this room, which matters when the listing being read gathers entries from  more than one of them. (optional)
	folderType := []int32{int32(0)} // []int32 | Keeps only the folders of these kinds, each given as the number of a folder type; it is how a listing is  narrowed down to, say, the form-filling folders of a room. (optional)
	excludeSubject := false // bool | Turns `userIdOrGroupId` around: with true the entries of that member or group are the ones left out, with  false they are the only ones kept. (optional)
	applyFilterOption := openapiclient.ApplyFilterOption(0) // ApplyFilterOption | Chooses which half of the listing `filterType` and `filterValue` are applied to: with `Files` the folders come  back unfiltered, with `Folders` the files do, and with `All` both halves are filtered. (optional)
	withSubFolders := true // bool | Whether a narrowed request reaches into the subfolders: with true, which is what an omitted parameter means,  matching entries are gathered from the whole subtree, with false only the top level is read. It makes a  difference only once `filterType`, `userIdOrGroupId` or `filterValue` narrows the request, because an  unfiltered listing always shows the top level alone. (optional)
	extension := "docx,pdf" // string | Keeps only the files carrying one of these extensions, several of them separated by commas; the leading dot is  optional. (optional)
	searchArea := openapiclient.SearchArea("Active") // SearchArea | Which area a listing that spans several of them is taken from - the active rooms, the archive, the room  templates or the form-filling rooms. A folder that belongs to one area only settles the area itself and  ignores the parameter. (optional)
	formsItemKey := "first_name" // string | Keeps only the completed forms whose form field of this name holds a value. Take the name from  `GET api/2.0/files/{folderId}/formfilter`, and use it in the folder that gathers the completed copies of a  form-filling room. (optional)
	formsItemType := "text" // string | The kind of the form field named by `formsItemKey`, taken from the same list; the two are sent together. (optional)
	count := int32(25) // int32 | The size of one page of the listing. Pair it with `startIndex` to walk through the result, and compare the two  with `total` in the response to see when the last page has been read. (optional)
	startIndex := int32(0) // int32 | The number of matching entries to skip before the returned page begins; add `count` to it to ask for the next  page. (optional)
	sortBy := "DateAndTime" // string | The name of the field the entries are ordered by, matched case-insensitively against the file sort fields:  `DateAndTime`, `AZ`, `Size`, `Author`, `Type`, `New`, `DateAndTimeCreation`, `RoomType`, `Tags`, `Room`,  `CustomOrder`, `LastOpened` and `UsedSpace`. A recognized value is also saved as the default order of the  account and reused by later listings that omit the parameter, while a value matching none of the fields leaves  that saved order in place. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The direction in which the `sortBy` field is ordered. It is saved together with `sortBy` as the default order  of the account. (optional)
	filterValue := "My Document" // string | The search string the listing is filtered by: it is matched as a substring of entry titles and, for files,  against the indexed document content as well. Omit it to list the folder unfiltered. (optional)
	location := openapiclient.Location(1) // Location | Where the entries of a tag-based listing have to live to be kept: `Room` keeps what lies in a room,  `Documents` what lies in a personal section, and `Link` what was reached through an external link that is  still valid. It shapes the Favorites and Recent listings and does nothing in an ordinary folder. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetFolderByFolderId(context.Background(), folderId).UserIdOrGroupId(userIdOrGroupId).SharedBy(sharedBy).FilterType(filterType).RoomId(roomId).FolderType(folderType).ExcludeSubject(excludeSubject).ApplyFilterOption(applyFilterOption).WithSubFolders(withSubFolders).Extension(extension).SearchArea(searchArea).FormsItemKey(formsItemKey).FormsItemType(formsItemType).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Location(location).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// folderId := "sbox-42"
	// thirdPartyResp, r, err := apiClient.FilesFoldersAPI.GetFolderByFolderId(context.Background(), folderId).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetFolderByFolderId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolderByFolderId`: FolderContentWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetFolderByFolderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder whose contents are listed. Each section root has an operation of its own, such as  `GET api/2.0/files/@my`, and every other folder is opened by the identifier a listing gave for it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderByFolderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userIdOrGroupId** | **string** | Restricts the listing to the entries authored by this portal member, or by the members of this group; the same  parameter accepts either kind of identifier. Omit it to list everything the caller can read. | 
 **sharedBy** | **string** | Restricts the listing to the entries this member shared, which narrows a shared listing down to what one  person handed out. | 
 **filterType** | [**FilterType**](FilterType.md) | Narrows the listing to a single kind of entry, such as documents, spreadsheets, images or one type of room.  Omit it to list every kind the folder holds. | 
 **roomId** | **int32** | Keeps only the entries that lie in this room, which matters when the listing being read gathers entries from  more than one of them. | 
 **folderType** | **[]int32** | Keeps only the folders of these kinds, each given as the number of a folder type; it is how a listing is  narrowed down to, say, the form-filling folders of a room. | 
 **excludeSubject** | **bool** | Turns `userIdOrGroupId` around: with true the entries of that member or group are the ones left out, with  false they are the only ones kept. | 
 **applyFilterOption** | [**ApplyFilterOption**](ApplyFilterOption.md) | Chooses which half of the listing `filterType` and `filterValue` are applied to: with `Files` the folders come  back unfiltered, with `Folders` the files do, and with `All` both halves are filtered. | 
 **withSubFolders** | **bool** | Whether a narrowed request reaches into the subfolders: with true, which is what an omitted parameter means,  matching entries are gathered from the whole subtree, with false only the top level is read. It makes a  difference only once `filterType`, `userIdOrGroupId` or `filterValue` narrows the request, because an  unfiltered listing always shows the top level alone. | 
 **extension** | **string** | Keeps only the files carrying one of these extensions, several of them separated by commas; the leading dot is  optional. | 
 **searchArea** | [**SearchArea**](SearchArea.md) | Which area a listing that spans several of them is taken from - the active rooms, the archive, the room  templates or the form-filling rooms. A folder that belongs to one area only settles the area itself and  ignores the parameter. | 
 **formsItemKey** | **string** | Keeps only the completed forms whose form field of this name holds a value. Take the name from  `GET api/2.0/files/{folderId}/formfilter`, and use it in the folder that gathers the completed copies of a  form-filling room. | 
 **formsItemType** | **string** | The kind of the form field named by `formsItemKey`, taken from the same list; the two are sent together. | 
 **count** | **int32** | The size of one page of the listing. Pair it with `startIndex` to walk through the result, and compare the two  with `total` in the response to see when the last page has been read. | 
 **startIndex** | **int32** | The number of matching entries to skip before the returned page begins; add `count` to it to ask for the next  page. | 
 **sortBy** | **string** | The name of the field the entries are ordered by, matched case-insensitively against the file sort fields:  `DateAndTime`, `AZ`, `Size`, `Author`, `Type`, `New`, `DateAndTimeCreation`, `RoomType`, `Tags`, `Room`,  `CustomOrder`, `LastOpened` and `UsedSpace`. A recognized value is also saved as the default order of the  account and reused by later listings that omit the parameter, while a value matching none of the fields leaves  that saved order in place. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The direction in which the `sortBy` field is ordered. It is saved together with `sortBy` as the default order  of the account. | 
 **filterValue** | **string** | The search string the listing is filtered by: it is matched as a substring of entry titles and, for files,  against the indexed document content as well. Omit it to list the folder unfiltered. | 
 **location** | [**Location**](Location.md) | Where the entries of a tag-based listing have to live to be kept: `Room` keeps what lies in a room,  `Documents` what lies in a personal section, and `Link` what was reached through an external link that is  still valid. It shapes the Favorites and Recent listings and does nothing in an ordinary folder. | 

### Return type

[**FolderContentWrapper**](FolderContentWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

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
	folderId := int32(1) // int32 | The folder whose activity log is read; the log covers the folder itself and the entries inside it.
	fromDate := time.Now() // time.Time | The earliest moment an entry may have, read in the time zone of the portal; left out, the log starts at the  oldest entry the portal still keeps. (optional)
	toDate := time.Now() // time.Time | The latest moment an entry may have, read in the time zone of the portal; left out, the log ends at the newest  entry. (optional)
	count := int32(25) // int32 | How many entries one page holds. The number of entries that match the query is reported in the response  headers, not in the body. (optional)
	startIndex := int32(0) // int32 | How many entries to skip before the page begins, counted from the newest one, so pages are taken by adding the  page size to it. (optional)

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
**folderId** | **int32** | The folder whose activity log is read; the log covers the folder itself and the entries inside it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **time.Time** | The earliest moment an entry may have, read in the time zone of the portal; left out, the log starts at the  oldest entry the portal still keeps. | 
 **toDate** | **time.Time** | The latest moment an entry may have, read in the time zone of the portal; left out, the log ends at the newest  entry. | 
 **count** | **int32** | How many entries one page holds. The number of entries that match the query is reported in the response  headers, not in the body. | 
 **startIndex** | **int32** | How many entries to skip before the page begins, counted from the newest one, so pages are taken by adding the  page size to it. | 

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

> FolderWrapper GetFolderInfo(ctx, folderId).Execute()

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
	folderId := int32(1) // int32 | The folder the operation acts on. Take the identifier from a listing such as `GET api/2.0/files/@root` or  `GET api/2.0/files/{folderId}`: a folder stored in the portal is numbered, while a folder in a connected  third-party account is named by an opaque string.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetFolderInfo(context.Background(), folderId).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// folderId := "sbox-42"
	// thirdPartyResp, r, err := apiClient.FilesFoldersAPI.GetFolderInfo(context.Background(), folderId).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetFolderInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolderInfo`: FolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetFolderInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder the operation acts on. Take the identifier from a listing such as `GET api/2.0/files/@root` or  `GET api/2.0/files/{folderId}`: a folder stored in the portal is numbered, while a folder in a connected  third-party account is named by an opaque string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderWrapper**](FolderWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolderLinks

> FileShareArrayWrapper GetFolderLinks(ctx, id).Execute()

Get folder external links



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
	id := int32(1) // int32 | The folder or room whose external links are listed.

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
**id** | **int32** | The folder or room whose external links are listed. | 

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
	folderId := int32(1) // int32 | The folder the operation acts on. Take the identifier from a listing such as `GET api/2.0/files/@root` or  `GET api/2.0/files/{folderId}`: a folder stored in the portal is numbered, while a folder in a connected  third-party account is named by an opaque string.

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
**folderId** | **int32** | The folder the operation acts on. Take the identifier from a listing such as `GET api/2.0/files/@root` or  `GET api/2.0/files/{folderId}`: a folder stored in the portal is numbered, while a folder in a connected  third-party account is named by an opaque string. | 

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

Get the folder primary external link



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
	id := int32(10) // int32 | The folder or room the operation addresses. A folder stored on the portal is numbered, while a folder in a  connected third-party account is named by an opaque string.
	count := int32(25) // int32 | How many entries at most to answer with, in the operations of this folder that return a list; an operation  that answers with a single object is not affected by it. (optional)
	startIndex := int32(0) // int32 | How many entries of such a list to skip before answering, used together with `count` to walk through it page  by page. (optional)

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
**id** | **int32** | The folder or room the operation addresses. A folder stored on the portal is numbered, while a folder in a  connected third-party account is named by an opaque string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderPrimaryExternalLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | How many entries at most to answer with, in the operations of this folder that return a list; an operation  that answers with a single object is not affected by it. | 
 **startIndex** | **int32** | How many entries of such a list to skip before answering, used together with `count` to walk through it page  by page. | 

### Return type

[**FileShareWrapper**](FileShareWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

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
	folderId := int32(1) // int32 | The folder the operation acts on. Take the identifier from a listing such as `GET api/2.0/files/@root` or  `GET api/2.0/files/{folderId}`: a folder stored in the portal is numbered, while a folder in a connected  third-party account is named by an opaque string.

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
**folderId** | **int32** | The folder the operation acts on. Take the identifier from a listing such as `GET api/2.0/files/@root` or  `GET api/2.0/files/{folderId}`: a folder stored in the portal is numbered, while a folder in a connected  third-party account is named by an opaque string. | 

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

> FolderContentWrapper GetFormsFolder(ctx).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()

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
	userIdOrGroupId := "00000000-0000-0000-0000-000000000000" // string | Restricts the listing to the entries authored by this portal member, or by the members of this group; the same  parameter accepts either kind of identifier. Omit it to list everything the caller can read. (optional)
	filterType := openapiclient.FilterType(0) // FilterType | Narrows the listing to a single kind of entry, such as documents, images or one type of room. Omit it to list  every kind the section holds. (optional)
	count := int32(25) // int32 | The size of one page of section content. Pair it with `startIndex` to walk the listing, and compare the two  with `total` in the response to see when the last page has been read. (optional)
	startIndex := int32(0) // int32 | The number of matching entries to skip before the returned page begins; add `count` to it to ask for the next  page. (optional)
	sortBy := "DateAndTime" // string | The name of the field the entries are ordered by, matched case-insensitively against the file sort fields:  `DateAndTime`, `AZ`, `Size`, `Author`, `Type`, `New`, `DateAndTimeCreation`, `RoomType`, `Tags`, `Room`,  `CustomOrder`, `LastOpened` and `UsedSpace`. A recognized value is also saved as the default order of the  account and reused by later listings that omit the parameter, while a value matching none of the fields leaves  that saved order in place. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The direction in which the `sortBy` field is ordered. It is saved together with `sortBy` as the default order  of the account. (optional)
	filterValue := "My Document" // string | The search string the section is filtered by: it is matched as a substring of entry titles and, for files,  against the indexed document content as well. Omit it to list the section unfiltered. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetFormsFolder(context.Background()).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetFormsFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFormsFolder`: FolderContentWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetFormsFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFormsFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userIdOrGroupId** | **string** | Restricts the listing to the entries authored by this portal member, or by the members of this group; the same  parameter accepts either kind of identifier. Omit it to list everything the caller can read. | 
 **filterType** | [**FilterType**](FilterType.md) | Narrows the listing to a single kind of entry, such as documents, images or one type of room. Omit it to list  every kind the section holds. | 
 **count** | **int32** | The size of one page of section content. Pair it with `startIndex` to walk the listing, and compare the two  with `total` in the response to see when the last page has been read. | 
 **startIndex** | **int32** | The number of matching entries to skip before the returned page begins; add `count` to it to ask for the next  page. | 
 **sortBy** | **string** | The name of the field the entries are ordered by, matched case-insensitively against the file sort fields:  `DateAndTime`, `AZ`, `Size`, `Author`, `Type`, `New`, `DateAndTimeCreation`, `RoomType`, `Tags`, `Room`,  `CustomOrder`, `LastOpened` and `UsedSpace`. A recognized value is also saved as the default order of the  account and reused by later listings that omit the parameter, while a value matching none of the fields leaves  that saved order in place. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The direction in which the `sortBy` field is ordered. It is saved together with `sortBy` as the default order  of the account. | 
 **filterValue** | **string** | The search string the section is filtered by: it is matched as a substring of entry titles and, for files,  against the indexed document content as well. Omit it to list the section unfiltered. | 

### Return type

[**FolderContentWrapper**](FolderContentWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyFolder

> FolderContentWrapper GetMyFolder(ctx).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).ApplyFilterOption(applyFilterOption).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()

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
	userIdOrGroupId := "00000000-0000-0000-0000-000000000000" // string | Restricts the listing to the entries authored by this portal member, or by the members of this group; the same  parameter accepts either kind of identifier. Omit it to list everything the caller can read. (optional)
	filterType := openapiclient.FilterType(0) // FilterType | Narrows the listing to a single kind of entry, such as documents, images or one type of room. Omit it to list  every kind the section holds. (optional)
	applyFilterOption := openapiclient.ApplyFilterOption(0) // ApplyFilterOption | Chooses which half of the listing `filterType` and `filterValue` are applied to: with `Files` the folders come  back unfiltered, with `Folders` the files do, and with `All` both halves are filtered. (optional)
	count := int32(25) // int32 | The size of one page of section content. Pair it with `startIndex` to walk the listing, and compare the two  with `total` in the response to see when the last page has been read. (optional)
	startIndex := int32(0) // int32 | The number of matching entries to skip before the returned page begins; add `count` to it to ask for the next  page. (optional)
	sortBy := "DateAndTime" // string | The name of the field the entries are ordered by, matched case-insensitively against the file sort fields:  `DateAndTime`, `AZ`, `Size`, `Author`, `Type`, `New`, `DateAndTimeCreation`, `RoomType`, `Tags`, `Room`,  `CustomOrder`, `LastOpened` and `UsedSpace`. A recognized value is also saved as the default order of the  account and reused by later listings that omit the parameter, while a value matching none of the fields leaves  that saved order in place. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The direction in which the `sortBy` field is ordered. It is saved together with `sortBy` as the default order  of the account. (optional)
	filterValue := "My Document" // string | The search string the section is filtered by, matched as a substring of entry titles. Omit it to list the  section unfiltered. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetMyFolder(context.Background()).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).ApplyFilterOption(applyFilterOption).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetMyFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyFolder`: FolderContentWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetMyFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userIdOrGroupId** | **string** | Restricts the listing to the entries authored by this portal member, or by the members of this group; the same  parameter accepts either kind of identifier. Omit it to list everything the caller can read. | 
 **filterType** | [**FilterType**](FilterType.md) | Narrows the listing to a single kind of entry, such as documents, images or one type of room. Omit it to list  every kind the section holds. | 
 **applyFilterOption** | [**ApplyFilterOption**](ApplyFilterOption.md) | Chooses which half of the listing `filterType` and `filterValue` are applied to: with `Files` the folders come  back unfiltered, with `Folders` the files do, and with `All` both halves are filtered. | 
 **count** | **int32** | The size of one page of section content. Pair it with `startIndex` to walk the listing, and compare the two  with `total` in the response to see when the last page has been read. | 
 **startIndex** | **int32** | The number of matching entries to skip before the returned page begins; add `count` to it to ask for the next  page. | 
 **sortBy** | **string** | The name of the field the entries are ordered by, matched case-insensitively against the file sort fields:  `DateAndTime`, `AZ`, `Size`, `Author`, `Type`, `New`, `DateAndTimeCreation`, `RoomType`, `Tags`, `Room`,  `CustomOrder`, `LastOpened` and `UsedSpace`. A recognized value is also saved as the default order of the  account and reused by later listings that omit the parameter, while a value matching none of the fields leaves  that saved order in place. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The direction in which the `sortBy` field is ordered. It is saved together with `sortBy` as the default order  of the account. | 
 **filterValue** | **string** | The search string the section is filtered by, matched as a substring of entry titles. Omit it to list the  section unfiltered. | 

### Return type

[**FolderContentWrapper**](FolderContentWrapper.md)

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
	folderId := int32(1) // int32 | The folder the operation acts on. Take the identifier from a listing such as `GET api/2.0/files/@root` or  `GET api/2.0/files/{folderId}`: a folder stored in the portal is numbered, while a folder in a connected  third-party account is named by an opaque string.

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
**folderId** | **int32** | The folder the operation acts on. Take the identifier from a listing such as `GET api/2.0/files/@root` or  `GET api/2.0/files/{folderId}`: a folder stored in the portal is numbered, while a folder in a connected  third-party account is named by an opaque string. | 

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

> FolderContentWrapper GetRecentFolder(ctx).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).ExcludeSubject(excludeSubject).ApplyFilterOption(applyFilterOption).SearchArea(searchArea).Extension(extension).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()

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
	userIdOrGroupId := "00000000-0000-0000-0000-000000000000" // string | Restricts the listing to the files authored by this portal member, or by the members of this group; the same  parameter accepts either kind of identifier. Omit it to list the whole history. (optional)
	filterType := openapiclient.FilterType(0) // FilterType | Narrows the listing to a single kind of file, such as documents, spreadsheets or images. Omit it to list every  kind the history holds. (optional)
	excludeSubject := false // bool | Inverts `userIdOrGroupId`: with `true` the files of that member or group are the ones left out of the listing  instead of the only ones kept. (optional)
	applyFilterOption := openapiclient.ApplyFilterOption(0) // ApplyFilterOption | Chooses which half of a listing `filterType` and `filterValue` are applied to. The Recent section holds  files only, so the value does not change what comes back. (optional)
	searchArea := openapiclient.SearchArea("Active") // SearchArea | The area a listing is taken from. The Recent section is assembled from the caller's own open history rather  than from an area, so the value does not change which files are returned. (optional)
	extension := []string{"Inner_example"} // []string | The file extensions the listing is limited to, matched against the end of the file name. The leading dot is  optional, and the parameter is repeated once per extension. (optional)
	count := int32(25) // int32 | The size of one page of section content. Pair it with `startIndex` to walk the listing, and compare the two  with `total` in the response to see when the last page has been read. (optional)
	startIndex := int32(0) // int32 | The number of matching entries to skip before the returned page begins; add `count` to it to ask for the next  page. (optional)
	sortBy := "DateAndTime" // string | The name of the field the entries are ordered by, matched case-insensitively against the file sort fields:  `DateAndTime`, `AZ`, `Size`, `Author`, `Type`, `New`, `DateAndTimeCreation`, `RoomType`, `Tags`, `Room`,  `CustomOrder`, `LastOpened` and `UsedSpace`. A recognized value is also saved as the default order of the  account and reused by later listings that omit the parameter, while a value matching none of the fields leaves  that saved order in place. The Recent section keeps its own newest-first order, so the value does not  reorder this listing. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The direction in which the `sortBy` field is ordered. It is saved together with `sortBy` as the default order  of the account. The Recent section keeps its own newest-first order, so the value does not reorder this  listing. (optional)
	filterValue := "My Document" // string | The search string the history is filtered by: it is matched as a substring of file titles and against the  indexed document content as well. Omit it to list the whole history. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetRecentFolder(context.Background()).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).ExcludeSubject(excludeSubject).ApplyFilterOption(applyFilterOption).SearchArea(searchArea).Extension(extension).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetRecentFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecentFolder`: FolderContentWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetRecentFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecentFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userIdOrGroupId** | **string** | Restricts the listing to the files authored by this portal member, or by the members of this group; the same  parameter accepts either kind of identifier. Omit it to list the whole history. | 
 **filterType** | [**FilterType**](FilterType.md) | Narrows the listing to a single kind of file, such as documents, spreadsheets or images. Omit it to list every  kind the history holds. | 
 **excludeSubject** | **bool** | Inverts `userIdOrGroupId`: with `true` the files of that member or group are the ones left out of the listing  instead of the only ones kept. | 
 **applyFilterOption** | [**ApplyFilterOption**](ApplyFilterOption.md) | Chooses which half of a listing `filterType` and `filterValue` are applied to. The Recent section holds  files only, so the value does not change what comes back. | 
 **searchArea** | [**SearchArea**](SearchArea.md) | The area a listing is taken from. The Recent section is assembled from the caller's own open history rather  than from an area, so the value does not change which files are returned. | 
 **extension** | **[]string** | The file extensions the listing is limited to, matched against the end of the file name. The leading dot is  optional, and the parameter is repeated once per extension. | 
 **count** | **int32** | The size of one page of section content. Pair it with `startIndex` to walk the listing, and compare the two  with `total` in the response to see when the last page has been read. | 
 **startIndex** | **int32** | The number of matching entries to skip before the returned page begins; add `count` to it to ask for the next  page. | 
 **sortBy** | **string** | The name of the field the entries are ordered by, matched case-insensitively against the file sort fields:  `DateAndTime`, `AZ`, `Size`, `Author`, `Type`, `New`, `DateAndTimeCreation`, `RoomType`, `Tags`, `Room`,  `CustomOrder`, `LastOpened` and `UsedSpace`. A recognized value is also saved as the default order of the  account and reused by later listings that omit the parameter, while a value matching none of the fields leaves  that saved order in place. The Recent section keeps its own newest-first order, so the value does not  reorder this listing. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The direction in which the `sortBy` field is ordered. It is saved together with `sortBy` as the default order  of the account. The Recent section keeps its own newest-first order, so the value does not reorder this  listing. | 
 **filterValue** | **string** | The search string the history is filtered by: it is matched as a substring of file titles and against the  indexed document content as well. Omit it to list the whole history. | 

### Return type

[**FolderContentWrapper**](FolderContentWrapper.md)

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
	folderId := int32(56) // int32 | The folder whose history report is being polled. It is the folder that was              passed to the operation that started the report.

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
**folderId** | **int32** | The folder whose history report is being polled. It is the folder that was              passed to the operation that started the report. | 

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

> FolderContentArrayWrapper GetRootFolders(ctx).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).WithoutTrash(withoutTrash).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()

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
	userIdOrGroupId := "00000000-0000-0000-0000-000000000000" // string | Restricts the listing to the entries authored by this portal member, or by the members of this group; the same  parameter accepts either kind of identifier. Omit it to list everything the caller can read. (optional)
	filterType := openapiclient.FilterType(0) // FilterType | Narrows the content listed inside every returned section to a single kind of entry, such as documents, images  or one type of room. Omit it to list every kind the sections hold. (optional)
	withoutTrash := false // bool | Set it to `true` to leave the Trash section out of the returned set of sections; with `false`, or when the  parameter is omitted, the section is returned whenever the account has one of its own. (optional)
	count := int32(25) // int32 | The size of the content page returned for each section separately, so a value of 1 yields one entry per  section rather than one entry in total. (optional)
	startIndex := int32(0) // int32 | The number of matching entries skipped in each section before its page begins; add `count` to it to ask for  the next page of every section. (optional)
	sortBy := "DateAndTime" // string | The name of the field the entries are ordered by, matched case-insensitively against the file sort fields:  `DateAndTime`, `AZ`, `Size`, `Author`, `Type`, `New`, `DateAndTimeCreation`, `RoomType`, `Tags`, `Room`,  `CustomOrder`, `LastOpened` and `UsedSpace`. A recognized value is also saved as the default order of the  account and reused by later listings that omit the parameter, while a value matching none of the fields leaves  that saved order in place. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The direction in which the `sortBy` field is ordered. It is saved together with `sortBy` as the default order  of the account. (optional)
	filterValue := "My Document" // string | The search string the content of every section is filtered by: it is matched as a substring of entry titles  and, for files, against the indexed document content as well. Omit it to list the sections unfiltered. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetRootFolders(context.Background()).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).WithoutTrash(withoutTrash).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetRootFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRootFolders`: FolderContentArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetRootFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRootFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userIdOrGroupId** | **string** | Restricts the listing to the entries authored by this portal member, or by the members of this group; the same  parameter accepts either kind of identifier. Omit it to list everything the caller can read. | 
 **filterType** | [**FilterType**](FilterType.md) | Narrows the content listed inside every returned section to a single kind of entry, such as documents, images  or one type of room. Omit it to list every kind the sections hold. | 
 **withoutTrash** | **bool** | Set it to `true` to leave the Trash section out of the returned set of sections; with `false`, or when the  parameter is omitted, the section is returned whenever the account has one of its own. | 
 **count** | **int32** | The size of the content page returned for each section separately, so a value of 1 yields one entry per  section rather than one entry in total. | 
 **startIndex** | **int32** | The number of matching entries skipped in each section before its page begins; add `count` to it to ask for  the next page of every section. | 
 **sortBy** | **string** | The name of the field the entries are ordered by, matched case-insensitively against the file sort fields:  `DateAndTime`, `AZ`, `Size`, `Author`, `Type`, `New`, `DateAndTimeCreation`, `RoomType`, `Tags`, `Room`,  `CustomOrder`, `LastOpened` and `UsedSpace`. A recognized value is also saved as the default order of the  account and reused by later listings that omit the parameter, while a value matching none of the fields leaves  that saved order in place. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The direction in which the `sortBy` field is ordered. It is saved together with `sortBy` as the default order  of the account. | 
 **filterValue** | **string** | The search string the content of every section is filtered by: it is matched as a substring of entry titles  and, for files, against the indexed document content as well. Omit it to list the sections unfiltered. | 

### Return type

[**FolderContentArrayWrapper**](FolderContentArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrashFolder

> FolderContentWrapper GetTrashFolder(ctx).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).ApplyFilterOption(applyFilterOption).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()

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
	userIdOrGroupId := "00000000-0000-0000-0000-000000000000" // string | Restricts the listing to the entries authored by this portal member, or by the members of this group; the same  parameter accepts either kind of identifier. Omit it to list everything the caller can read. (optional)
	filterType := openapiclient.FilterType(0) // FilterType | Narrows the listing to a single kind of entry, such as documents, images or one type of room. Omit it to list  every kind the section holds. (optional)
	applyFilterOption := openapiclient.ApplyFilterOption(0) // ApplyFilterOption | Chooses which half of the listing `filterType` and `filterValue` are applied to: with `Files` the folders come  back unfiltered, with `Folders` the files do, and with `All` both halves are filtered. (optional)
	count := int32(25) // int32 | The size of one page of section content. Pair it with `startIndex` to walk the listing, and compare the two  with `total` in the response to see when the last page has been read. (optional)
	startIndex := int32(0) // int32 | The number of matching entries to skip before the returned page begins; add `count` to it to ask for the next  page. (optional)
	sortBy := "DateAndTime" // string | The name of the field the entries are ordered by, matched case-insensitively against the file sort fields:  `DateAndTime`, `AZ`, `Size`, `Author`, `Type`, `New`, `DateAndTimeCreation`, `RoomType`, `Tags`, `Room`,  `CustomOrder`, `LastOpened` and `UsedSpace`. A recognized value is also saved as the default order of the  account and reused by later listings that omit the parameter, while a value matching none of the fields leaves  that saved order in place. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The direction in which the `sortBy` field is ordered. It is saved together with `sortBy` as the default order  of the account. (optional)
	filterValue := "My Document" // string | The search string the section is filtered by, matched as a substring of entry titles. Omit it to list the  section unfiltered. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.GetTrashFolder(context.Background()).UserIdOrGroupId(userIdOrGroupId).FilterType(filterType).ApplyFilterOption(applyFilterOption).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.GetTrashFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrashFolder`: FolderContentWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.GetTrashFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrashFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userIdOrGroupId** | **string** | Restricts the listing to the entries authored by this portal member, or by the members of this group; the same  parameter accepts either kind of identifier. Omit it to list everything the caller can read. | 
 **filterType** | [**FilterType**](FilterType.md) | Narrows the listing to a single kind of entry, such as documents, images or one type of room. Omit it to list  every kind the section holds. | 
 **applyFilterOption** | [**ApplyFilterOption**](ApplyFilterOption.md) | Chooses which half of the listing `filterType` and `filterValue` are applied to: with `Files` the folders come  back unfiltered, with `Folders` the files do, and with `All` both halves are filtered. | 
 **count** | **int32** | The size of one page of section content. Pair it with `startIndex` to walk the listing, and compare the two  with `total` in the response to see when the last page has been read. | 
 **startIndex** | **int32** | The number of matching entries to skip before the returned page begins; add `count` to it to ask for the next  page. | 
 **sortBy** | **string** | The name of the field the entries are ordered by, matched case-insensitively against the file sort fields:  `DateAndTime`, `AZ`, `Size`, `Author`, `Type`, `New`, `DateAndTimeCreation`, `RoomType`, `Tags`, `Room`,  `CustomOrder`, `LastOpened` and `UsedSpace`. A recognized value is also saved as the default order of the  account and reused by later listings that omit the parameter, while a value matching none of the fields leaves  that saved order in place. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The direction in which the `sortBy` field is ordered. It is saved together with `sortBy` as the default order  of the account. | 
 **filterValue** | **string** | The search string the section is filtered by, matched as a substring of entry titles. Omit it to list the  section unfiltered. | 

### Return type

[**FolderContentWrapper**](FolderContentWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsertFile

> FileWrapper InsertFile(ctx, folderId).InsertFileFile(insertFileFile).InsertFileTitle(insertFileTitle).InsertFileCreateNewIfExist(insertFileCreateNewIfExist).InsertFileKeepConvertStatus(insertFileKeepConvertStatus).InsertFileStreamCanRead(insertFileStreamCanRead).InsertFileStreamCanWrite(insertFileStreamCanWrite).InsertFileStreamCanSeek(insertFileStreamCanSeek).InsertFileStreamCanTimeout(insertFileStreamCanTimeout).InsertFileStreamLength(insertFileStreamLength).InsertFileStreamPosition(insertFileStreamPosition).InsertFileStreamReadTimeout(insertFileStreamReadTimeout).InsertFileStreamWriteTimeout(insertFileStreamWriteTimeout).Execute()

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
	folderId := int32(1) // int32 | The folder that receives the file; take the id from a listing such as `GET api/2.0/files/@root`. A room or an  ordinary folder inside one is accepted, a section root is not.
	insertFileFile := os.NewFile(1234, "some_file") // *os.File | The content to store, sent as a `multipart/form-data` part. The same content may instead be sent as the raw  request body, which is what a client that cannot build a form does; when both are present the form part wins. (optional)
	insertFileTitle := "insertFileTitle_example" // string | The name to store the file under, extension included. It wins over the name of the uploaded part, which is the  reason to choose this operation over the plain upload, and it is the only name available when the content  arrives as a raw body. Characters a title cannot hold are replaced with underscores and the name is cut to 170  characters before the file is stored. (optional)
	insertFileCreateNewIfExist := true // bool | Settles the clash with a file already carrying that title: left out, the content is written as the next  version of that file; set to true, both survive and the new one gets a numeric suffix in its title. (optional)
	insertFileKeepConvertStatus := true // bool | Decides whether the outcome of the background conversion outlives the conversion itself. True keeps the queue  record, so `GET api/2.0/files/file/{fileId}/checkconversion` can still report the result or the error; left  out, the record is cleared the moment the conversion ends and that call finds nothing. (optional)
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
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// folderId := "sbox-42"
	// thirdPartyResp, r, err := apiClient.FilesFoldersAPI.InsertFile(context.Background(), folderId).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.InsertFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InsertFile`: FileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.InsertFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder that receives the file; take the id from a listing such as `GET api/2.0/files/@root`. A room or an  ordinary folder inside one is accepted, a section root is not. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsertFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **insertFileFile** | ***os.File** | The content to store, sent as a `multipart/form-data` part. The same content may instead be sent as the raw  request body, which is what a client that cannot build a form does; when both are present the form part wins. | 
 **insertFileTitle** | **string** | The name to store the file under, extension included. It wins over the name of the uploaded part, which is the  reason to choose this operation over the plain upload, and it is the only name available when the content  arrives as a raw body. Characters a title cannot hold are replaced with underscores and the name is cut to 170  characters before the file is stored. | 
 **insertFileCreateNewIfExist** | **bool** | Settles the clash with a file already carrying that title: left out, the content is written as the next  version of that file; set to true, both survive and the new one gets a numeric suffix in its title. | 
 **insertFileKeepConvertStatus** | **bool** | Decides whether the outcome of the background conversion outlives the conversion itself. True keeps the queue  record, so `GET api/2.0/files/file/{fileId}/checkconversion` can still report the result or the error; left  out, the record is cleared the moment the conversion ends and that call finds nothing. | 
 **insertFileStreamCanRead** | **bool** |  | 
 **insertFileStreamCanWrite** | **bool** |  | 
 **insertFileStreamCanSeek** | **bool** |  | 
 **insertFileStreamCanTimeout** | **bool** |  | 
 **insertFileStreamLength** | **int64** |  | 
 **insertFileStreamPosition** | **int64** |  | 
 **insertFileStreamReadTimeout** | **int32** |  | 
 **insertFileStreamWriteTimeout** | **int32** |  | 

### Return type

[**FileWrapper**](FileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsertFileToMyFromBody

> FileWrapper InsertFileToMyFromBody(ctx).File(file).Title(title).CreateNewIfExist(createNewIfExist).KeepConvertStatus(keepConvertStatus).StreamCanRead(streamCanRead).StreamCanWrite(streamCanWrite).StreamCanSeek(streamCanSeek).StreamCanTimeout(streamCanTimeout).StreamLength(streamLength).StreamPosition(streamPosition).StreamReadTimeout(streamReadTimeout).StreamWriteTimeout(streamWriteTimeout).Execute()

Insert a file into My documents



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
	file := os.NewFile(1234, "some_file") // *os.File | The content to store, sent as a `multipart/form-data` part. The same content may instead be sent as the raw  request body, which is what a client that cannot build a form does; when both are present the form part wins. (optional)
	title := "title_example" // string | The name to store the file under, extension included. It wins over the name of the uploaded part, which is the  reason to choose this operation over the plain upload, and it is the only name available when the content  arrives as a raw body. Characters a title cannot hold are replaced with underscores and the name is cut to 170  characters before the file is stored. (optional)
	createNewIfExist := true // bool | Settles the clash with a file already carrying that title: left out, the content is written as the next  version of that file; set to true, both survive and the new one gets a numeric suffix in its title. (optional)
	keepConvertStatus := true // bool | Decides whether the outcome of the background conversion outlives the conversion itself. True keeps the queue  record, so `GET api/2.0/files/file/{fileId}/checkconversion` can still report the result or the error; left  out, the record is cleared the moment the conversion ends and that call finds nothing. (optional)
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
	// response from `InsertFileToMyFromBody`: FileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.InsertFileToMyFromBody`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsertFileToMyFromBodyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The content to store, sent as a `multipart/form-data` part. The same content may instead be sent as the raw  request body, which is what a client that cannot build a form does; when both are present the form part wins. | 
 **title** | **string** | The name to store the file under, extension included. It wins over the name of the uploaded part, which is the  reason to choose this operation over the plain upload, and it is the only name available when the content  arrives as a raw body. Characters a title cannot hold are replaced with underscores and the name is cut to 170  characters before the file is stored. | 
 **createNewIfExist** | **bool** | Settles the clash with a file already carrying that title: left out, the content is written as the next  version of that file; set to true, both survive and the new one gets a numeric suffix in its title. | 
 **keepConvertStatus** | **bool** | Decides whether the outcome of the background conversion outlives the conversion itself. True keeps the queue  record, so `GET api/2.0/files/file/{fileId}/checkconversion` can still report the result or the error; left  out, the record is cleared the moment the conversion ends and that call finds nothing. | 
 **streamCanRead** | **bool** |  | 
 **streamCanWrite** | **bool** |  | 
 **streamCanSeek** | **bool** |  | 
 **streamCanTimeout** | **bool** |  | 
 **streamLength** | **int64** |  | 
 **streamPosition** | **int64** |  | 
 **streamReadTimeout** | **int32** |  | 
 **streamWriteTimeout** | **int32** |  | 

### Return type

[**FileWrapper**](FileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameFolder

> FolderWrapper RenameFolder(ctx, folderId).CreateFolder(createFolder).Execute()

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
	folderId := int32(1) // int32 | The folder the request is addressed to: when a folder is created it is the parent that receives the new  folder, and when a folder is renamed it is the folder that gets the new title.
	createFolder := *openapiclient.NewCreateFolder("New Folder") // CreateFolder | The title carried by the request body.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.RenameFolder(context.Background(), folderId).CreateFolder(createFolder).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// folderId := "sbox-42"
	// thirdPartyResp, r, err := apiClient.FilesFoldersAPI.RenameFolder(context.Background(), folderId).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.RenameFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenameFolder`: FolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.RenameFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder the request is addressed to: when a folder is created it is the parent that receives the new  folder, and when a folder is renamed it is the folder that gets the new title. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFolder** | [**CreateFolder**](CreateFolder.md) | The title carried by the request body. | 

### Return type

[**FolderWrapper**](FolderWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFolderOrder

> FolderWrapper SetFolderOrder(ctx, folderId).OrderRequestDto(orderRequestDto).Execute()

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
	folderId := int32(1) // int32 | The folder to move.
	orderRequestDto := *openapiclient.NewOrderRequestDto() // OrderRequestDto | The position the folder is to take. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.SetFolderOrder(context.Background(), folderId).OrderRequestDto(orderRequestDto).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// folderId := "sbox-42"
	// thirdPartyResp, r, err := apiClient.FilesFoldersAPI.SetFolderOrder(context.Background(), folderId).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.SetFolderOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFolderOrder`: FolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.SetFolderOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder to move. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFolderOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderRequestDto** | [**OrderRequestDto**](OrderRequestDto.md) | The position the folder is to take. | 

### Return type

[**FolderWrapper**](FolderWrapper.md)

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
	id := int32(1) // int32 | The folder or room the link belongs to.
	folderLinkRequest := *openapiclient.NewFolderLinkRequest() // FolderLinkRequest | The link and the way it is to be shaped.

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
**id** | **int32** | The folder or room the link belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFolderPrimaryExternalLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderLinkRequest** | [**FolderLinkRequest**](FolderLinkRequest.md) | The link and the way it is to be shaped. | 

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
	folderId := int32(56) // int32 | The folder whose running history report is to be given up. It is the folder that              was passed to the operation that started the report.

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
**folderId** | **int32** | The folder whose running history report is to be given up. It is the folder that              was passed to the operation that started the report. | 

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

> FileArrayWrapper UploadFile(ctx, folderId).CreateNewIfExist(createNewIfExist).StoreOriginalFile(storeOriginalFile).KeepConvertStatus(keepConvertStatus).File(file).Execute()

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
	folderId := int32(1) // int32 | The folder that receives the file; take the id from a listing such as `GET api/2.0/files/@root`. A room or an  ordinary folder inside one is accepted, a section root is not.
	createNewIfExist := true // bool | Settles the clash with a file already carrying that title: left out, the content is written as the next  version of that file; set to true, both survive and the new one gets a numeric suffix in its title. (optional)
	storeOriginalFile := true // bool | Reaches further than this request: it writes a setting on the calling account, the same one  `PUT api/2.0/files/storeoriginal` writes, and it stays in force for later uploads. True keeps both the  uploaded file and the copy the portal converts it into, false replaces the uploaded file with the converted  one, and leaving it out keeps whatever the account already has. (optional)
	keepConvertStatus := true // bool | Decides whether the outcome of the background conversion outlives the conversion itself. True keeps the queue  record, so `GET api/2.0/files/file/{fileId}/checkconversion` can still report the result or the error; left  out, the record is cleared the moment the conversion ends and that call finds nothing. (optional)
	file := os.NewFile(1234, "some_file") // *os.File | The content to store, sent as a `multipart/form-data` part; the name of that part becomes the title of the  stored file, with characters a title cannot hold replaced and the name cut to 170 characters. A request  without it is rejected as invalid. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.UploadFile(context.Background(), folderId).CreateNewIfExist(createNewIfExist).StoreOriginalFile(storeOriginalFile).KeepConvertStatus(keepConvertStatus).File(file).Execute()
	// for an entry in a connected third-party storage (a string id such as "sbox-42"):
	// folderId := "sbox-42"
	// thirdPartyResp, r, err := apiClient.FilesFoldersAPI.UploadFile(context.Background(), folderId).ExecuteThirdParty()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.UploadFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadFile`: FileArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.UploadFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder that receives the file; take the id from a listing such as `GET api/2.0/files/@root`. A room or an  ordinary folder inside one is accepted, a section root is not. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNewIfExist** | **bool** | Settles the clash with a file already carrying that title: left out, the content is written as the next  version of that file; set to true, both survive and the new one gets a numeric suffix in its title. | 
 **storeOriginalFile** | **bool** | Reaches further than this request: it writes a setting on the calling account, the same one  `PUT api/2.0/files/storeoriginal` writes, and it stays in force for later uploads. True keeps both the  uploaded file and the copy the portal converts it into, false replaces the uploaded file with the converted  one, and leaving it out keeps whatever the account already has. | 
 **keepConvertStatus** | **bool** | Decides whether the outcome of the background conversion outlives the conversion itself. True keeps the queue  record, so `GET api/2.0/files/file/{fileId}/checkconversion` can still report the result or the error; left  out, the record is cleared the moment the conversion ends and that call finds nothing. | 
 **file** | ***os.File** | The content to store, sent as a `multipart/form-data` part; the name of that part becomes the title of the  stored file, with characters a title cannot hold replaced and the name cut to 170 characters. A request  without it is rejected as invalid. | 

### Return type

[**FileArrayWrapper**](FileArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFileToMy

> FileArrayWrapper UploadFileToMy(ctx).CreateNewIfExist(createNewIfExist).StoreOriginalFile(storeOriginalFile).KeepConvertStatus(keepConvertStatus).File(file).Execute()

Upload a file to My documents



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
	createNewIfExist := true // bool | Settles the clash with a file already carrying that title: left out, the content is written as the next  version of that file; set to true, both survive and the new one gets a numeric suffix in its title. (optional)
	storeOriginalFile := true // bool | Reaches further than this request: it writes a setting on the calling account, the same one  `PUT api/2.0/files/storeoriginal` writes, and it stays in force for later uploads. True keeps both the  uploaded file and the copy the portal converts it into, false replaces the uploaded file with the converted  one, and leaving it out keeps whatever the account already has. (optional)
	keepConvertStatus := true // bool | Decides whether the outcome of the background conversion outlives the conversion itself. True keeps the queue  record, so `GET api/2.0/files/file/{fileId}/checkconversion` can still report the result or the error; left  out, the record is cleared the moment the conversion ends and that call finds nothing. (optional)
	file := os.NewFile(1234, "some_file") // *os.File | The content to store, sent as a `multipart/form-data` part; the name of that part becomes the title of the  stored file, with characters a title cannot hold replaced and the name cut to 170 characters. A request  without it is rejected as invalid. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFoldersAPI.UploadFileToMy(context.Background()).CreateNewIfExist(createNewIfExist).StoreOriginalFile(storeOriginalFile).KeepConvertStatus(keepConvertStatus).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFoldersAPI.UploadFileToMy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadFileToMy`: FileArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFoldersAPI.UploadFileToMy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileToMyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNewIfExist** | **bool** | Settles the clash with a file already carrying that title: left out, the content is written as the next  version of that file; set to true, both survive and the new one gets a numeric suffix in its title. | 
 **storeOriginalFile** | **bool** | Reaches further than this request: it writes a setting on the calling account, the same one  `PUT api/2.0/files/storeoriginal` writes, and it stays in force for later uploads. True keeps both the  uploaded file and the copy the portal converts it into, false replaces the uploaded file with the converted  one, and leaving it out keeps whatever the account already has. | 
 **keepConvertStatus** | **bool** | Decides whether the outcome of the background conversion outlives the conversion itself. True keeps the queue  record, so `GET api/2.0/files/file/{fileId}/checkconversion` can still report the result or the error; left  out, the record is cleared the moment the conversion ends and that call finds nothing. | 
 **file** | ***os.File** | The content to store, sent as a `multipart/form-data` part; the name of that part becomes the title of the  stored file, with characters a title cannot hold replaced and the name cut to 170 characters. A request  without it is rejected as invalid. | 

### Return type

[**FileArrayWrapper**](FileArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

