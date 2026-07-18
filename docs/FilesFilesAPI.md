# \FilesFilesAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFileToRecent**](FilesFilesAPI.md#AddFileToRecent) | **Post** /api/2.0/files/file/{fileId}/recent | Add a file to the Recent section
[**AddTemplates**](FilesFilesAPI.md#AddTemplates) | **Post** /api/2.0/files/templates | Add template files
[**ChangeVersionHistory**](FilesFilesAPI.md#ChangeVersionHistory) | **Put** /api/2.0/files/file/{fileId}/history | Change version history
[**CheckFillFormDraft**](FilesFilesAPI.md#CheckFillFormDraft) | **Post** /api/2.0/files/masterform/{fileId}/checkfillformdraft | Check the form draft filling
[**CopyFileAs**](FilesFilesAPI.md#CopyFileAs) | **Post** /api/2.0/files/file/{fileId}/copyas | Copy a file
[**CreateEditSession**](FilesFilesAPI.md#CreateEditSession) | **Post** /api/2.0/files/file/{fileId}/edit_session | Create the editing session
[**CreateFile**](FilesFilesAPI.md#CreateFile) | **Post** /api/2.0/files/{folderId}/file | Create a file
[**CreateFileInMyDocuments**](FilesFilesAPI.md#CreateFileInMyDocuments) | **Post** /api/2.0/files/@my/file | Create a file in the My documents section
[**CreateFilePrimaryExternalLink**](FilesFilesAPI.md#CreateFilePrimaryExternalLink) | **Post** /api/2.0/files/file/{id}/link | Create primary external link
[**CreateHtmlFile**](FilesFilesAPI.md#CreateHtmlFile) | **Post** /api/2.0/files/{folderId}/html | Create an HTML file
[**CreateHtmlFileInMyDocuments**](FilesFilesAPI.md#CreateHtmlFileInMyDocuments) | **Post** /api/2.0/files/@my/html | Create an HTML file in the My documents section
[**CreateTextFile**](FilesFilesAPI.md#CreateTextFile) | **Post** /api/2.0/files/{folderId}/text | Create a text file
[**CreateTextFileInMyDocuments**](FilesFilesAPI.md#CreateTextFileInMyDocuments) | **Post** /api/2.0/files/@my/text | Create a text file in the My documents section
[**CreateThumbnails**](FilesFilesAPI.md#CreateThumbnails) | **Post** /api/2.0/files/thumbnails | Create file thumbnails
[**DeleteFile**](FilesFilesAPI.md#DeleteFile) | **Delete** /api/2.0/files/file/{fileId} | Delete a file
[**DeleteRecent**](FilesFilesAPI.md#DeleteRecent) | **Delete** /api/2.0/files/recent | Delete recent files
[**DeleteTemplates**](FilesFilesAPI.md#DeleteTemplates) | **Delete** /api/2.0/files/templates | Delete template files
[**GenerateXlsx**](FilesFilesAPI.md#GenerateXlsx) | **Post** /api/2.0/files/file/{fileId}/xlsx | Generate XLSX report
[**GetAllFormRoles**](FilesFilesAPI.md#GetAllFormRoles) | **Get** /api/2.0/files/file/{fileId}/formroles | Get form roles
[**GetEditDiffUrl**](FilesFilesAPI.md#GetEditDiffUrl) | **Get** /api/2.0/files/file/{fileId}/edit/diff | Get changes URL
[**GetEditHistory**](FilesFilesAPI.md#GetEditHistory) | **Get** /api/2.0/files/file/{fileId}/edit/history | Get version history
[**GetFileHistory**](FilesFilesAPI.md#GetFileHistory) | **Get** /api/2.0/files/file/{fileId}/log | Get file history
[**GetFileInfo**](FilesFilesAPI.md#GetFileInfo) | **Get** /api/2.0/files/file/{fileId} | Get file information
[**GetFileLinks**](FilesFilesAPI.md#GetFileLinks) | **Get** /api/2.0/files/file/{id}/links | Get file external links
[**GetFilePrimaryExternalLink**](FilesFilesAPI.md#GetFilePrimaryExternalLink) | **Get** /api/2.0/files/file/{id}/link | Get primary external link
[**GetFileVersionInfo**](FilesFilesAPI.md#GetFileVersionInfo) | **Get** /api/2.0/files/file/{fileId}/history | Get file versions
[**GetFillResult**](FilesFilesAPI.md#GetFillResult) | **Get** /api/2.0/files/file/fillresult | Get form-filling result
[**GetFormSubmissions**](FilesFilesAPI.md#GetFormSubmissions) | **Get** /api/2.0/files/file/{fileId}/submissions | Get form submission results
[**GetPresignedFileUri**](FilesFilesAPI.md#GetPresignedFileUri) | **Get** /api/2.0/files/file/{fileId}/presigned | Get file download link asynchronously
[**GetPresignedUri**](FilesFilesAPI.md#GetPresignedUri) | **Get** /api/2.0/files/file/{fileId}/presigneduri | Get file download link
[**GetProtectedFileUsers**](FilesFilesAPI.md#GetProtectedFileUsers) | **Get** /api/2.0/files/file/{fileId}/protectusers | Get users access rights to the protected file
[**GetReferenceData**](FilesFilesAPI.md#GetReferenceData) | **Post** /api/2.0/files/file/referencedata | Get reference data
[**GetXlsx**](FilesFilesAPI.md#GetXlsx) | **Get** /api/2.0/files/file/{fileId}/xlsx | Get XLSX report generation status
[**IsFormPDF**](FilesFilesAPI.md#IsFormPDF) | **Get** /api/2.0/files/file/{fileId}/isformpdf | Check the PDF file
[**LockFile**](FilesFilesAPI.md#LockFile) | **Put** /api/2.0/files/file/{fileId}/lock | Lock a file
[**ManageFormFilling**](FilesFilesAPI.md#ManageFormFilling) | **Put** /api/2.0/files/file/{fileId}/manageformfilling | Perform form filling action
[**OpenEditFile**](FilesFilesAPI.md#OpenEditFile) | **Get** /api/2.0/files/file/{fileId}/openedit | Open a file configuration
[**RestoreFileVersion**](FilesFilesAPI.md#RestoreFileVersion) | **Post** /api/2.0/files/file/{fileId}/restoreversion | Restore a file version
[**SaveEditingFileFromForm**](FilesFilesAPI.md#SaveEditingFileFromForm) | **Put** /api/2.0/files/file/{fileId}/saveediting | Save file edits
[**SaveFileAsPdf**](FilesFilesAPI.md#SaveFileAsPdf) | **Post** /api/2.0/files/file/{id}/saveaspdf | Save a file as PDF
[**SaveFormRoleMapping**](FilesFilesAPI.md#SaveFormRoleMapping) | **Post** /api/2.0/files/file/{fileId}/formrolemapping | Save form role mapping
[**SetCustomFilterTag**](FilesFilesAPI.md#SetCustomFilterTag) | **Put** /api/2.0/files/file/{fileId}/customfilter | Set the Custom Filter editing mode
[**SetFileExternalLink**](FilesFilesAPI.md#SetFileExternalLink) | **Put** /api/2.0/files/file/{id}/links | Set an external link
[**SetFileOrder**](FilesFilesAPI.md#SetFileOrder) | **Put** /api/2.0/files/{fileId}/order | Set file order
[**SetFilesOrder**](FilesFilesAPI.md#SetFilesOrder) | **Put** /api/2.0/files/order | Set order of files
[**StartEditFile**](FilesFilesAPI.md#StartEditFile) | **Post** /api/2.0/files/file/{fileId}/startedit | Start file editing
[**StartFillingFile**](FilesFilesAPI.md#StartFillingFile) | **Put** /api/2.0/files/file/{fileId}/startfilling | Start file filling
[**ToggleFileFavorite**](FilesFilesAPI.md#ToggleFileFavorite) | **Get** /api/2.0/files/favorites/{fileId} | Change the file favorite status
[**TrackEditFile**](FilesFilesAPI.md#TrackEditFile) | **Get** /api/2.0/files/file/{fileId}/trackeditfile | Track file editing
[**UpdateFile**](FilesFilesAPI.md#UpdateFile) | **Put** /api/2.0/files/file/{fileId} | Update a file



## AddFileToRecent

> FileIntegerWrapper AddFileToRecent(ctx, fileId).Execute()

Add a file to the Recent section



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/add-file-to-recent/).

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
	fileId := int32(1) // int32 | The file unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.AddFileToRecent(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.AddFileToRecent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddFileToRecent`: FileIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.AddFileToRecent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFileToRecentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileIntegerWrapper**](FileIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTemplates

> BooleanWrapper AddTemplates(ctx).TemplatesRequestDto(templatesRequestDto).Execute()

Add template files



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/add-templates/).

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
	templatesRequestDto := *openapiclient.NewTemplatesRequestDto() // TemplatesRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.AddTemplates(context.Background()).TemplatesRequestDto(templatesRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.AddTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTemplates`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.AddTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templatesRequestDto** | [**TemplatesRequestDto**](TemplatesRequestDto.md) |  | 

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


## ChangeVersionHistory

> FileIntegerArrayWrapper ChangeVersionHistory(ctx, fileId).ChangeHistory(changeHistory).Execute()

Change version history



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/change-version-history/).

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
	fileId := int32(1) // int32 | The file Id to change its version history.
	changeHistory := *openapiclient.NewChangeHistory(int32(1)) // ChangeHistory | The parameters for changing version history.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.ChangeVersionHistory(context.Background(), fileId).ChangeHistory(changeHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.ChangeVersionHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeVersionHistory`: FileIntegerArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.ChangeVersionHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file Id to change its version history. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeVersionHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeHistory** | [**ChangeHistory**](ChangeHistory.md) | The parameters for changing version history. | 

### Return type

[**FileIntegerArrayWrapper**](FileIntegerArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckFillFormDraft

> StringWrapper CheckFillFormDraft(ctx, fileId).CheckFillFormDraft(checkFillFormDraft).Execute()

Check the form draft filling



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/check-fill-form-draft/).

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
	fileId := int32(1) // int32 | The file ID of the form draft.
	checkFillFormDraft := *openapiclient.NewCheckFillFormDraft(int32(1)) // CheckFillFormDraft | The parameters for checking the form draft filling.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CheckFillFormDraft(context.Background(), fileId).CheckFillFormDraft(checkFillFormDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CheckFillFormDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckFillFormDraft`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CheckFillFormDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID of the form draft. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckFillFormDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkFillFormDraft** | [**CheckFillFormDraft**](CheckFillFormDraft.md) | The parameters for checking the form draft filling. | 

### Return type

[**StringWrapper**](StringWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CopyFileAs

> FileEntryBaseWrapper CopyFileAs(ctx, fileId).CopyAsJsonElement(copyAsJsonElement).Execute()

Copy a file



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/copy-file-as/).

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
	fileId := int32(1) // int32 | The file ID to copy.
	copyAsJsonElement := *openapiclient.NewCopyAsJsonElement("Document Copy.docx", openapiclient.CopyAsJsonElement_destFolderId{Int32: new(int32)}) // CopyAsJsonElement | The parameters for copying a file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CopyFileAs(context.Background(), fileId).CopyAsJsonElement(copyAsJsonElement).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CopyFileAs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CopyFileAs`: FileEntryBaseWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CopyFileAs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID to copy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyFileAsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **copyAsJsonElement** | [**CopyAsJsonElement**](CopyAsJsonElement.md) | The parameters for copying a file. | 

### Return type

[**FileEntryBaseWrapper**](FileEntryBaseWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEditSession

> ChunkedUploadSessionResponseWrapperIntegerWrapper CreateEditSession(ctx, fileId).FileSize(fileSize).Execute()

Create the editing session



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-edit-session/).

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
	fileId := int32(1) // int32 | The file ID.
	fileSize := int64(1024) // int64 | The file size in bytes. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CreateEditSession(context.Background(), fileId).FileSize(fileSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CreateEditSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEditSession`: ChunkedUploadSessionResponseWrapperIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CreateEditSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEditSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileSize** | **int64** | The file size in bytes. | 

### Return type

[**ChunkedUploadSessionResponseWrapperIntegerWrapper**](ChunkedUploadSessionResponseWrapperIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFile

> FileIntegerWrapper CreateFile(ctx, folderId).CreateFileJsonElement(createFileJsonElement).Execute()

Create a file



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-file/).

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
	folderId := int32(1) // int32 | The folder ID for the file creation.
	createFileJsonElement := *openapiclient.NewCreateFileJsonElement("New Document.docx") // CreateFileJsonElement | The parameters for creating a file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CreateFile(context.Background(), folderId).CreateFileJsonElement(createFileJsonElement).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CreateFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFile`: FileIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CreateFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID for the file creation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFileJsonElement** | [**CreateFileJsonElement**](CreateFileJsonElement.md) | The parameters for creating a file. | 

### Return type

[**FileIntegerWrapper**](FileIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFileInMyDocuments

> FileIntegerWrapper CreateFileInMyDocuments(ctx).CreateFileJsonElement(createFileJsonElement).Execute()

Create a file in the My documents section



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-file-in-my-documents/).

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
	createFileJsonElement := *openapiclient.NewCreateFileJsonElement("New Document.docx") // CreateFileJsonElement |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CreateFileInMyDocuments(context.Background()).CreateFileJsonElement(createFileJsonElement).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CreateFileInMyDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFileInMyDocuments`: FileIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CreateFileInMyDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFileInMyDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFileJsonElement** | [**CreateFileJsonElement**](CreateFileJsonElement.md) |  | 

### Return type

[**FileIntegerWrapper**](FileIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFilePrimaryExternalLink

> FileShareWrapper CreateFilePrimaryExternalLink(ctx, id).FileLinkRequest(fileLinkRequest).Execute()

Create primary external link



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-file-primary-external-link/).

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
	id := int32(1) // int32 | The file ID.
	fileLinkRequest := *openapiclient.NewFileLinkRequest() // FileLinkRequest | The file external link parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CreateFilePrimaryExternalLink(context.Background(), id).FileLinkRequest(fileLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CreateFilePrimaryExternalLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFilePrimaryExternalLink`: FileShareWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CreateFilePrimaryExternalLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The file ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFilePrimaryExternalLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileLinkRequest** | [**FileLinkRequest**](FileLinkRequest.md) | The file external link parameters. | 

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


## CreateHtmlFile

> FileIntegerWrapper CreateHtmlFile(ctx, folderId).CreateTextOrHtmlFile(createTextOrHtmlFile).Execute()

Create an HTML file



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-html-file/).

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
	folderId := int32(1) // int32 | The folder ID to create the text or HTML file.
	createTextOrHtmlFile := *openapiclient.NewCreateTextOrHtmlFile("Document.txt") // CreateTextOrHtmlFile | The parameters for creating an HTML or text file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CreateHtmlFile(context.Background(), folderId).CreateTextOrHtmlFile(createTextOrHtmlFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CreateHtmlFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHtmlFile`: FileIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CreateHtmlFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID to create the text or HTML file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHtmlFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTextOrHtmlFile** | [**CreateTextOrHtmlFile**](CreateTextOrHtmlFile.md) | The parameters for creating an HTML or text file. | 

### Return type

[**FileIntegerWrapper**](FileIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHtmlFileInMyDocuments

> FileIntegerWrapper CreateHtmlFileInMyDocuments(ctx).CreateTextOrHtmlFile(createTextOrHtmlFile).Execute()

Create an HTML file in the My documents section



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-html-file-in-my-documents/).

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
	createTextOrHtmlFile := *openapiclient.NewCreateTextOrHtmlFile("Document.txt") // CreateTextOrHtmlFile |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CreateHtmlFileInMyDocuments(context.Background()).CreateTextOrHtmlFile(createTextOrHtmlFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CreateHtmlFileInMyDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHtmlFileInMyDocuments`: FileIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CreateHtmlFileInMyDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHtmlFileInMyDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTextOrHtmlFile** | [**CreateTextOrHtmlFile**](CreateTextOrHtmlFile.md) |  | 

### Return type

[**FileIntegerWrapper**](FileIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTextFile

> FileIntegerWrapper CreateTextFile(ctx, folderId).CreateTextOrHtmlFile(createTextOrHtmlFile).Execute()

Create a text file



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-text-file/).

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
	folderId := int32(1) // int32 | The folder ID to create the text or HTML file.
	createTextOrHtmlFile := *openapiclient.NewCreateTextOrHtmlFile("Document.txt") // CreateTextOrHtmlFile | The parameters for creating an HTML or text file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CreateTextFile(context.Background(), folderId).CreateTextOrHtmlFile(createTextOrHtmlFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CreateTextFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTextFile`: FileIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CreateTextFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID to create the text or HTML file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTextFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTextOrHtmlFile** | [**CreateTextOrHtmlFile**](CreateTextOrHtmlFile.md) | The parameters for creating an HTML or text file. | 

### Return type

[**FileIntegerWrapper**](FileIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTextFileInMyDocuments

> FileIntegerWrapper CreateTextFileInMyDocuments(ctx).CreateTextOrHtmlFile(createTextOrHtmlFile).Execute()

Create a text file in the My documents section



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-text-file-in-my-documents/).

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
	createTextOrHtmlFile := *openapiclient.NewCreateTextOrHtmlFile("Document.txt") // CreateTextOrHtmlFile |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CreateTextFileInMyDocuments(context.Background()).CreateTextOrHtmlFile(createTextOrHtmlFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CreateTextFileInMyDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTextFileInMyDocuments`: FileIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CreateTextFileInMyDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTextFileInMyDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTextOrHtmlFile** | [**CreateTextOrHtmlFile**](CreateTextOrHtmlFile.md) |  | 

### Return type

[**FileIntegerWrapper**](FileIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateThumbnails

> ObjectArrayWrapper CreateThumbnails(ctx).BaseBatchRequestDto(baseBatchRequestDto).Execute()

Create file thumbnails



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-thumbnails/).

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
	resp, r, err := apiClient.FilesFilesAPI.CreateThumbnails(context.Background()).BaseBatchRequestDto(baseBatchRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CreateThumbnails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateThumbnails`: ObjectArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CreateThumbnails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateThumbnailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseBatchRequestDto** | [**BaseBatchRequestDto**](BaseBatchRequestDto.md) |  | 

### Return type

[**ObjectArrayWrapper**](ObjectArrayWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFile

> FileOperationArrayWrapper DeleteFile(ctx, fileId).Delete(delete).ReturnSingleOperation(returnSingleOperation).Execute()

Delete a file



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-file/).

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
	fileId := int32(1) // int32 | The file ID to delete.
	delete := *openapiclient.NewDelete() // Delete | The parameters for deleting a file.
	returnSingleOperation := false // bool | Specifies whether to return only the current operation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.DeleteFile(context.Background(), fileId).Delete(delete).ReturnSingleOperation(returnSingleOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.DeleteFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFile`: FileOperationArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.DeleteFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **delete** | [**Delete**](Delete.md) | The parameters for deleting a file. | 
 **returnSingleOperation** | **bool** | Specifies whether to return only the current operation | 

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


## DeleteRecent

> NoContentResultWrapper DeleteRecent(ctx).BaseBatchRequestDto(baseBatchRequestDto).Execute()

Delete recent files



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-recent/).

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
	resp, r, err := apiClient.FilesFilesAPI.DeleteRecent(context.Background()).BaseBatchRequestDto(baseBatchRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.DeleteRecent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRecent`: NoContentResultWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.DeleteRecent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseBatchRequestDto** | [**BaseBatchRequestDto**](BaseBatchRequestDto.md) |  | 

### Return type

[**NoContentResultWrapper**](NoContentResultWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTemplates

> BooleanWrapper DeleteTemplates(ctx).RequestBody(requestBody).Execute()

Delete template files



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-templates/).

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
	requestBody := []int32{int32(123)} // []int32 | The file IDs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.DeleteTemplates(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.DeleteTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTemplates`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.DeleteTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]int32** | The file IDs. | 

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


## GenerateXlsx

> XlsxReportResponseWrapper GenerateXlsx(ctx, fileId).Execute()

Generate XLSX report



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/generate-xlsx/).

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
	fileId := int32(1) // int32 | The file unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GenerateXlsx(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GenerateXlsx``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateXlsx`: XlsxReportResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GenerateXlsx`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateXlsxRequest struct via the builder pattern


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


## GetAllFormRoles

> FormRoleArrayWrapper GetAllFormRoles(ctx, fileId).Execute()

Get form roles



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-all-form-roles/).

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
	fileId := int32(1) // int32 | The file unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetAllFormRoles(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetAllFormRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllFormRoles`: FormRoleArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetAllFormRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllFormRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FormRoleArrayWrapper**](FormRoleArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEditDiffUrl

> EditHistoryDataWrapper GetEditDiffUrl(ctx, fileId).Version(version).Execute()

Get changes URL



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-edit-diff-url/).

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
	fileId := int32(1) // int32 | The file ID.
	version := int32(1) // int32 | The file version. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetEditDiffUrl(context.Background(), fileId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetEditDiffUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEditDiffUrl`: EditHistoryDataWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetEditDiffUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEditDiffUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** | The file version. | 

### Return type

[**EditHistoryDataWrapper**](EditHistoryDataWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEditHistory

> EditHistoryArrayWrapper GetEditHistory(ctx, fileId).Execute()

Get version history



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-edit-history/).

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
	fileId := int32(1) // int32 | The file unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetEditHistory(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetEditHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEditHistory`: EditHistoryArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetEditHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEditHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EditHistoryArrayWrapper**](EditHistoryArrayWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileHistory

> HistoryArrayWrapper GetFileHistory(ctx, fileId).FromDate(fromDate).ToDate(toDate).Count(count).StartIndex(startIndex).Execute()

Get file history



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-file-history/).

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
	fileId := int32(1) // int32 | The file ID of the history request.
	fromDate := *openapiclient.NewApiDateTime() // ApiDateTime | The start date of the history. (optional)
	toDate := *openapiclient.NewApiDateTime() // ApiDateTime | The end date of the history. (optional)
	count := int32(25) // int32 | The number of history entries to retrieve for the file log. (optional)
	startIndex := int32(0) // int32 | The starting index for retrieving a subset of file history entries. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetFileHistory(context.Background(), fileId).FromDate(fromDate).ToDate(toDate).Count(count).StartIndex(startIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetFileHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileHistory`: HistoryArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetFileHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID of the history request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | [**ApiDateTime**](ApiDateTime.md) | The start date of the history. | 
 **toDate** | [**ApiDateTime**](ApiDateTime.md) | The end date of the history. | 
 **count** | **int32** | The number of history entries to retrieve for the file log. | 
 **startIndex** | **int32** | The starting index for retrieving a subset of file history entries. | 

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


## GetFileInfo

> FileIntegerWrapper GetFileInfo(ctx, fileId).Version(version).Execute()

Get file information



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-file-info/).

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
	fileId := int32(1) // int32 | The file ID.
	version := int32(1) // int32 | The file version. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetFileInfo(context.Background(), fileId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetFileInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileInfo`: FileIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetFileInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** | The file version. | 

### Return type

[**FileIntegerWrapper**](FileIntegerWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileLinks

> FileShareArrayWrapper GetFileLinks(ctx, id).Count(count).StartIndex(startIndex).Execute()

Get file external links



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-file-links/).

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
	id := int32(10) // int32 | The file unique identifier.
	count := int32(25) // int32 | The number of items to retrieve in the request. (optional)
	startIndex := int32(0) // int32 | The starting index for the query results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetFileLinks(context.Background(), id).Count(count).StartIndex(startIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetFileLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileLinks`: FileShareArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetFileLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The file unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The number of items to retrieve in the request. | 
 **startIndex** | **int32** | The starting index for the query results. | 

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


## GetFilePrimaryExternalLink

> FileShareWrapper GetFilePrimaryExternalLink(ctx, id).Count(count).StartIndex(startIndex).Execute()

Get primary external link



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-file-primary-external-link/).

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
	id := int32(10) // int32 | The file unique identifier.
	count := int32(25) // int32 | The number of items to retrieve in the request. (optional)
	startIndex := int32(0) // int32 | The starting index for the query results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetFilePrimaryExternalLink(context.Background(), id).Count(count).StartIndex(startIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetFilePrimaryExternalLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilePrimaryExternalLink`: FileShareWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetFilePrimaryExternalLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The file unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilePrimaryExternalLinkRequest struct via the builder pattern


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


## GetFileVersionInfo

> FileIntegerArrayWrapper GetFileVersionInfo(ctx, fileId).Execute()

Get file versions



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-file-version-info/).

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
	fileId := int32(1) // int32 | The file unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetFileVersionInfo(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetFileVersionInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileVersionInfo`: FileIntegerArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetFileVersionInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileVersionInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileIntegerArrayWrapper**](FileIntegerArrayWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFillResult

> FillingFormResultIntegerWrapper GetFillResult(ctx).FillingSessionId(fillingSessionId).Execute()

Get form-filling result



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-fill-result/).

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
	fillingSessionId := "doc_key_123" // string | The form-filling session ID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetFillResult(context.Background()).FillingSessionId(fillingSessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetFillResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFillResult`: FillingFormResultIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetFillResult`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFillResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fillingSessionId** | **string** | The form-filling session ID. | 

### Return type

[**FillingFormResultIntegerWrapper**](FillingFormResultIntegerWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFormSubmissions

> FormSubmissionsWrapper GetFormSubmissions(ctx, fileId).Execute()

Get form submission results



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-form-submissions/).

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
	fileId := int32(1) // int32 | The file unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetFormSubmissions(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetFormSubmissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFormSubmissions`: FormSubmissionsWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetFormSubmissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFormSubmissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FormSubmissionsWrapper**](FormSubmissionsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPresignedFileUri

> FileLinkWrapper GetPresignedFileUri(ctx, fileId).Execute()

Get file download link asynchronously



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-presigned-file-uri/).

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
	fileId := int32(1) // int32 | The file unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetPresignedFileUri(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetPresignedFileUri``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPresignedFileUri`: FileLinkWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetPresignedFileUri`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPresignedFileUriRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileLinkWrapper**](FileLinkWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPresignedUri

> StringWrapper GetPresignedUri(ctx, fileId).Execute()

Get file download link



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-presigned-uri/).

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
	fileId := int32(1) // int32 | The file unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetPresignedUri(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetPresignedUri``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPresignedUri`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetPresignedUri`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPresignedUriRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StringWrapper**](StringWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProtectedFileUsers

> MentionWrapperArrayWrapper GetProtectedFileUsers(ctx, fileId).Execute()

Get users access rights to the protected file



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-protected-file-users/).

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
	fileId := int32(1) // int32 | The file unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetProtectedFileUsers(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetProtectedFileUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtectedFileUsers`: MentionWrapperArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetProtectedFileUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectedFileUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MentionWrapperArrayWrapper**](MentionWrapperArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReferenceData

> FileReferenceWrapper GetReferenceData(ctx).GetReferenceDataDtoInteger(getReferenceDataDtoInteger).Execute()

Get reference data



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-reference-data/).

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
	getReferenceDataDtoInteger := *openapiclient.NewGetReferenceDataDtoInteger("doc_key_123", "doc_key_123") // GetReferenceDataDtoInteger |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetReferenceData(context.Background()).GetReferenceDataDtoInteger(getReferenceDataDtoInteger).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetReferenceData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReferenceData`: FileReferenceWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetReferenceData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReferenceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getReferenceDataDtoInteger** | [**GetReferenceDataDtoInteger**](GetReferenceDataDtoInteger.md) |  | 

### Return type

[**FileReferenceWrapper**](FileReferenceWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetXlsx

> DocumentBuilderTaskWrapper GetXlsx(ctx, fileId).Execute()

Get XLSX report generation status



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-xlsx/).

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
	fileId := int32(1) // int32 | The file unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetXlsx(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetXlsx``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetXlsx`: DocumentBuilderTaskWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetXlsx`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetXlsxRequest struct via the builder pattern


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


## IsFormPDF

> BooleanWrapper IsFormPDF(ctx, fileId).Execute()

Check the PDF file



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/is-form-pdf/).

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
	fileId := int32(1) // int32 | The file unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.IsFormPDF(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.IsFormPDF``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IsFormPDF`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.IsFormPDF`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIsFormPDFRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockFile

> FileIntegerWrapper LockFile(ctx, fileId).LockFileParameters(lockFileParameters).Execute()

Lock a file



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/lock-file/).

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
	fileId := int32(1) // int32 | The file ID for locking.
	lockFileParameters := *openapiclient.NewLockFileParameters() // LockFileParameters | The parameters for locking a file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.LockFile(context.Background(), fileId).LockFileParameters(lockFileParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.LockFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LockFile`: FileIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.LockFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID for locking. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lockFileParameters** | [**LockFileParameters**](LockFileParameters.md) | The parameters for locking a file. | 

### Return type

[**FileIntegerWrapper**](FileIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageFormFilling

> ManageFormFilling(ctx, fileId).ManageFormFillingDtoInteger(manageFormFillingDtoInteger).Execute()

Perform form filling action



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/manage-form-filling/).

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
	fileId := "fileId_example" // string | 
	manageFormFillingDtoInteger := *openapiclient.NewManageFormFillingDtoInteger(int32(1)) // ManageFormFillingDtoInteger |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesFilesAPI.ManageFormFilling(context.Background(), fileId).ManageFormFillingDtoInteger(manageFormFillingDtoInteger).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.ManageFormFilling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageFormFillingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **manageFormFillingDtoInteger** | [**ManageFormFillingDtoInteger**](ManageFormFillingDtoInteger.md) |  | 

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


## OpenEditFile

> ConfigurationIntegerWrapper OpenEditFile(ctx, fileId).Version(version).View(view).EditorType(editorType).Edit(edit).Fill(fill).Execute()

Open a file configuration



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/open-edit-file/).

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
	fileId := int32(1) // int32 | The file ID to open.
	version := int32(1) // int32 | The file version to open. (optional)
	view := false // bool | Specifies if the document will be opened for viewing only or not. (optional)
	editorType := openapiclient.EditorType(0) // EditorType | The editor type to open the file. (optional)
	edit := false // bool | Specifies if the document is opened in the editing mode or not. (optional)
	fill := false // bool | Specifies if the document is opened in the form-filling mode or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.OpenEditFile(context.Background(), fileId).Version(version).View(view).EditorType(editorType).Edit(edit).Fill(fill).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.OpenEditFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenEditFile`: ConfigurationIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.OpenEditFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID to open. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpenEditFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** | The file version to open. | 
 **view** | **bool** | Specifies if the document will be opened for viewing only or not. | 
 **editorType** | [**EditorType**](EditorType.md) | The editor type to open the file. | 
 **edit** | **bool** | Specifies if the document is opened in the editing mode or not. | 
 **fill** | **bool** | Specifies if the document is opened in the form-filling mode or not. | 

### Return type

[**ConfigurationIntegerWrapper**](ConfigurationIntegerWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreFileVersion

> EditHistoryArrayWrapper RestoreFileVersion(ctx, fileId).Version(version).Url(url).Execute()

Restore a file version



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/restore-file-version/).

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
	fileId := int32(1) // int32 | The file ID of the restore version.
	version := int32(1) // int32 | The file version of the restore. (optional)
	url := "https://example.com" // string | The file version URL of the restore. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.RestoreFileVersion(context.Background(), fileId).Version(version).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.RestoreFileVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreFileVersion`: EditHistoryArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.RestoreFileVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID of the restore version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreFileVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** | The file version of the restore. | 
 **url** | **string** | The file version URL of the restore. | 

### Return type

[**EditHistoryArrayWrapper**](EditHistoryArrayWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveEditingFileFromForm

> FileIntegerWrapper SaveEditingFileFromForm(ctx, fileId).DownloadUri(downloadUri).FileExtension(fileExtension).File(file).Forcesave(forcesave).Execute()

Save file edits



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/save-editing-file-from-form/).

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
	fileId := int32(1) // int32 | The editing file ID from the request.
	downloadUri := "https://example.com/file.txt" // string | The URI to download the editing file. (optional)
	fileExtension := "fileExtension_example" // string | The editing file extension from the request. (optional)
	file := os.NewFile(1234, "some_file") // *os.File | The edited file to be saved, uploaded as part of the multipart/form-data request.  This property represents the modified file content from the HTTP request form after editing operations.  The file is accessed via the IFormFile interface which provides access to the file name, content type, length, and stream. (optional)
	forcesave := true // bool | Specifies whether to force save the file or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.SaveEditingFileFromForm(context.Background(), fileId).DownloadUri(downloadUri).FileExtension(fileExtension).File(file).Forcesave(forcesave).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.SaveEditingFileFromForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveEditingFileFromForm`: FileIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.SaveEditingFileFromForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The editing file ID from the request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveEditingFileFromFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **downloadUri** | **string** | The URI to download the editing file. | 
 **fileExtension** | **string** | The editing file extension from the request. | 
 **file** | ***os.File** | The edited file to be saved, uploaded as part of the multipart/form-data request.  This property represents the modified file content from the HTTP request form after editing operations.  The file is accessed via the IFormFile interface which provides access to the file name, content type, length, and stream. | 
 **forcesave** | **bool** | Specifies whether to force save the file or not. | 

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


## SaveFileAsPdf

> FileIntegerWrapper SaveFileAsPdf(ctx, id).SaveAsPdfInteger(saveAsPdfInteger).Execute()

Save a file as PDF



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/save-file-as-pdf/).

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
	id := int32(1) // int32 | The file ID to save as PDF.
	saveAsPdfInteger := *openapiclient.NewSaveAsPdfInteger(int32(1), "My Document") // SaveAsPdfInteger | The parameters for saving the file as PDF.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.SaveFileAsPdf(context.Background(), id).SaveAsPdfInteger(saveAsPdfInteger).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.SaveFileAsPdf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveFileAsPdf`: FileIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.SaveFileAsPdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The file ID to save as PDF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveFileAsPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **saveAsPdfInteger** | [**SaveAsPdfInteger**](SaveAsPdfInteger.md) | The parameters for saving the file as PDF. | 

### Return type

[**FileIntegerWrapper**](FileIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveFormRoleMapping

> SaveFormRoleMapping(ctx, fileId).SaveFormRoleMappingDtoInteger(saveFormRoleMappingDtoInteger).Execute()

Save form role mapping



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/save-form-role-mapping/).

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
	fileId := "fileId_example" // string | 
	saveFormRoleMappingDtoInteger := *openapiclient.NewSaveFormRoleMappingDtoInteger(int32(1), []openapiclient.FormRole{*openapiclient.NewFormRole()}) // SaveFormRoleMappingDtoInteger |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesFilesAPI.SaveFormRoleMapping(context.Background(), fileId).SaveFormRoleMappingDtoInteger(saveFormRoleMappingDtoInteger).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.SaveFormRoleMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveFormRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **saveFormRoleMappingDtoInteger** | [**SaveFormRoleMappingDtoInteger**](SaveFormRoleMappingDtoInteger.md) |  | 

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


## SetCustomFilterTag

> FileIntegerWrapper SetCustomFilterTag(ctx, fileId).CustomFilterParameters(customFilterParameters).Execute()

Set the Custom Filter editing mode



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-custom-filter-tag/).

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
	fileId := int32(1) // int32 | The file ID.
	customFilterParameters := *openapiclient.NewCustomFilterParameters() // CustomFilterParameters | The parameters for setting the Custom Filter editing mode.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.SetCustomFilterTag(context.Background(), fileId).CustomFilterParameters(customFilterParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.SetCustomFilterTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetCustomFilterTag`: FileIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.SetCustomFilterTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCustomFilterTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customFilterParameters** | [**CustomFilterParameters**](CustomFilterParameters.md) | The parameters for setting the Custom Filter editing mode. | 

### Return type

[**FileIntegerWrapper**](FileIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFileExternalLink

> FileShareWrapper SetFileExternalLink(ctx, id).FileLinkRequest(fileLinkRequest).Execute()

Set an external link



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-file-external-link/).

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
	id := int32(1) // int32 | The file ID.
	fileLinkRequest := *openapiclient.NewFileLinkRequest() // FileLinkRequest | The file external link parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.SetFileExternalLink(context.Background(), id).FileLinkRequest(fileLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.SetFileExternalLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFileExternalLink`: FileShareWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.SetFileExternalLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The file ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFileExternalLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileLinkRequest** | [**FileLinkRequest**](FileLinkRequest.md) | The file external link parameters. | 

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


## SetFileOrder

> FileIntegerWrapper SetFileOrder(ctx, fileId).OrderRequestDto(orderRequestDto).Execute()

Set file order



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-file-order/).

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
	fileId := int32(1) // int32 | The file unique identifier.
	orderRequestDto := *openapiclient.NewOrderRequestDto() // OrderRequestDto | The file order information. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.SetFileOrder(context.Background(), fileId).OrderRequestDto(orderRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.SetFileOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFileOrder`: FileIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.SetFileOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFileOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderRequestDto** | [**OrderRequestDto**](OrderRequestDto.md) | The file order information. | 

### Return type

[**FileIntegerWrapper**](FileIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFilesOrder

> FileEntryIntegerArrayWrapper SetFilesOrder(ctx).OrdersRequestDtoInteger(ordersRequestDtoInteger).Execute()

Set order of files



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-files-order/).

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
	ordersRequestDtoInteger := *openapiclient.NewOrdersRequestDtoInteger([]openapiclient.OrdersItemRequestDtoInteger{*openapiclient.NewOrdersItemRequestDtoInteger(int32(1), openapiclient.FileEntryType(1), int32(1))}) // OrdersRequestDtoInteger |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.SetFilesOrder(context.Background()).OrdersRequestDtoInteger(ordersRequestDtoInteger).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.SetFilesOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFilesOrder`: FileEntryIntegerArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.SetFilesOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetFilesOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordersRequestDtoInteger** | [**OrdersRequestDtoInteger**](OrdersRequestDtoInteger.md) |  | 

### Return type

[**FileEntryIntegerArrayWrapper**](FileEntryIntegerArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartEditFile

> StringWrapper StartEditFile(ctx, fileId).StartEdit(startEdit).Execute()

Start file editing



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/start-edit-file/).

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
	fileId := int32(1) // int32 | The file ID to start editing.
	startEdit := *openapiclient.NewStartEdit() // StartEdit | The file parameters to start editing.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.StartEditFile(context.Background(), fileId).StartEdit(startEdit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.StartEditFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartEditFile`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.StartEditFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID to start editing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartEditFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startEdit** | [**StartEdit**](StartEdit.md) | The file parameters to start editing. | 

### Return type

[**StringWrapper**](StringWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartFillingFile

> FileIntegerWrapper StartFillingFile(ctx, fileId).Execute()

Start file filling



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/start-filling-file/).

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
	fileId := int32(1) // int32 | The file ID to start filling.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.StartFillingFile(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.StartFillingFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartFillingFile`: FileIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.StartFillingFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID to start filling. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartFillingFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileIntegerWrapper**](FileIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleFileFavorite

> BooleanWrapper ToggleFileFavorite(ctx, fileId).Favorite(favorite).Execute()

Change the file favorite status



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/toggle-file-favorite/).

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
	fileId := int32(1) // int32 | The file ID.
	favorite := true // bool | Specifies if the file is marked as favorite or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.ToggleFileFavorite(context.Background(), fileId).Favorite(favorite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.ToggleFileFavorite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToggleFileFavorite`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.ToggleFileFavorite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleFileFavoriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **favorite** | **bool** | Specifies if the file is marked as favorite or not. | 

### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackEditFile

> KeyValuePairBooleanStringWrapper TrackEditFile(ctx, fileId).TabId(tabId).DocKeyForTrack(docKeyForTrack).IsFinish(isFinish).Execute()

Track file editing



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/track-edit-file/).

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
	fileId := int32(1) // int32 | The file ID to track editing changes.
	tabId := "00000000-0000-0000-0000-000000000000" // string | The tab ID to track editing changes. (optional)
	docKeyForTrack := "abc123" // string | The document key for tracking changes. (optional)
	isFinish := true // bool | Specifies whether to finish file tracking or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.TrackEditFile(context.Background(), fileId).TabId(tabId).DocKeyForTrack(docKeyForTrack).IsFinish(isFinish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.TrackEditFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackEditFile`: KeyValuePairBooleanStringWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.TrackEditFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID to track editing changes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrackEditFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tabId** | **string** | The tab ID to track editing changes. | 
 **docKeyForTrack** | **string** | The document key for tracking changes. | 
 **isFinish** | **bool** | Specifies whether to finish file tracking or not. | 

### Return type

[**KeyValuePairBooleanStringWrapper**](KeyValuePairBooleanStringWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFile

> FileIntegerWrapper UpdateFile(ctx, fileId).UpdateFile(updateFile).Execute()

Update a file



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-file/).

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
	fileId := int32(1) // int32 | The file ID to update.
	updateFile := *openapiclient.NewUpdateFile() // UpdateFile | The parameters for updating a file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.UpdateFile(context.Background(), fileId).UpdateFile(updateFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.UpdateFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFile`: FileIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.UpdateFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFile** | [**UpdateFile**](UpdateFile.md) | The parameters for updating a file. | 

### Return type

[**FileIntegerWrapper**](FileIntegerWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

