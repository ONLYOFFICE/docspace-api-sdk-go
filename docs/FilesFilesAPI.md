# \FilesFilesAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFileToRecent**](FilesFilesAPI.md#AddFileToRecent) | **Post** /api/2.0/files/file/{fileId}/recent | Add a file to Recent
[**AddFileToRecentThirdParty**](FilesFilesAPI.md#AddFileToRecentThirdParty) | **Post** /api/2.0/files/file/{fileId}/recent | Add a file to Recent (third-party storage)
[**AddTemplates**](FilesFilesAPI.md#AddTemplates) | **Post** /api/2.0/files/templates | Add template files
[**ChangeVersionHistory**](FilesFilesAPI.md#ChangeVersionHistory) | **Put** /api/2.0/files/file/{fileId}/history | Change version history
[**ChangeVersionHistoryThirdParty**](FilesFilesAPI.md#ChangeVersionHistoryThirdParty) | **Put** /api/2.0/files/file/{fileId}/history | Change version history (third-party storage)
[**CheckFillFormDraft**](FilesFilesAPI.md#CheckFillFormDraft) | **Post** /api/2.0/files/masterform/{fileId}/checkfillformdraft | Open a form draft for filling
[**CheckFillFormDraftThirdParty**](FilesFilesAPI.md#CheckFillFormDraftThirdParty) | **Post** /api/2.0/files/masterform/{fileId}/checkfillformdraft | Open a form draft for filling (third-party storage)
[**CopyFileAs**](FilesFilesAPI.md#CopyFileAs) | **Post** /api/2.0/files/file/{fileId}/copyas | Copy a file
[**CopyFileAsThirdParty**](FilesFilesAPI.md#CopyFileAsThirdParty) | **Post** /api/2.0/files/file/{fileId}/copyas | Copy a file (third-party storage)
[**CreateEditSession**](FilesFilesAPI.md#CreateEditSession) | **Post** /api/2.0/files/file/{fileId}/edit_session | Create the editing session
[**CreateEditSessionThirdParty**](FilesFilesAPI.md#CreateEditSessionThirdParty) | **Post** /api/2.0/files/file/{fileId}/edit_session | Create the editing session (third-party storage)
[**CreateFile**](FilesFilesAPI.md#CreateFile) | **Post** /api/2.0/files/{folderId}/file | Create a file
[**CreateFileThirdParty**](FilesFilesAPI.md#CreateFileThirdParty) | **Post** /api/2.0/files/{folderId}/file | Create a file (third-party storage)
[**CreateFileInMyDocuments**](FilesFilesAPI.md#CreateFileInMyDocuments) | **Post** /api/2.0/files/@my/file | Create a file in My documents
[**CreateFilePrimaryExternalLink**](FilesFilesAPI.md#CreateFilePrimaryExternalLink) | **Post** /api/2.0/files/file/{id}/link | Create the file primary external link
[**CreateFilePrimaryExternalLinkThirdParty**](FilesFilesAPI.md#CreateFilePrimaryExternalLinkThirdParty) | **Post** /api/2.0/files/file/{id}/link | Create the file primary external link (third-party storage)
[**CreateHtmlFile**](FilesFilesAPI.md#CreateHtmlFile) | **Post** /api/2.0/files/{folderId}/html | Create an HTML file
[**CreateHtmlFileThirdParty**](FilesFilesAPI.md#CreateHtmlFileThirdParty) | **Post** /api/2.0/files/{folderId}/html | Create an HTML file (third-party storage)
[**CreateHtmlFileInMyDocuments**](FilesFilesAPI.md#CreateHtmlFileInMyDocuments) | **Post** /api/2.0/files/@my/html | Create an HTML file in My documents
[**CreateTextFile**](FilesFilesAPI.md#CreateTextFile) | **Post** /api/2.0/files/{folderId}/text | Create a text file
[**CreateTextFileThirdParty**](FilesFilesAPI.md#CreateTextFileThirdParty) | **Post** /api/2.0/files/{folderId}/text | Create a text file (third-party storage)
[**CreateTextFileInMyDocuments**](FilesFilesAPI.md#CreateTextFileInMyDocuments) | **Post** /api/2.0/files/@my/text | Create a text file in My documents
[**CreateThumbnails**](FilesFilesAPI.md#CreateThumbnails) | **Post** /api/2.0/files/thumbnails | Queue file thumbnails
[**DeleteFile**](FilesFilesAPI.md#DeleteFile) | **Delete** /api/2.0/files/file/{fileId} | Delete a file
[**DeleteFileThirdParty**](FilesFilesAPI.md#DeleteFileThirdParty) | **Delete** /api/2.0/files/file/{fileId} | Delete a file (third-party storage)
[**DeleteRecent**](FilesFilesAPI.md#DeleteRecent) | **Delete** /api/2.0/files/recent | Delete recent files
[**DeleteTemplates**](FilesFilesAPI.md#DeleteTemplates) | **Delete** /api/2.0/files/templates | Delete template files
[**GenerateXlsx**](FilesFilesAPI.md#GenerateXlsx) | **Post** /api/2.0/files/file/{fileId}/xlsx | Generate a form answers report
[**GetAllFormRoles**](FilesFilesAPI.md#GetAllFormRoles) | **Get** /api/2.0/files/file/{fileId}/formroles | Get form roles
[**GetAllFormRolesThirdParty**](FilesFilesAPI.md#GetAllFormRolesThirdParty) | **Get** /api/2.0/files/file/{fileId}/formroles | Get form roles (third-party storage)
[**GetEditDiffUrl**](FilesFilesAPI.md#GetEditDiffUrl) | **Get** /api/2.0/files/file/{fileId}/edit/diff | Get changes URL
[**GetEditDiffUrlThirdParty**](FilesFilesAPI.md#GetEditDiffUrlThirdParty) | **Get** /api/2.0/files/file/{fileId}/edit/diff | Get changes URL (third-party storage)
[**GetEditHistory**](FilesFilesAPI.md#GetEditHistory) | **Get** /api/2.0/files/file/{fileId}/edit/history | Get version history
[**GetEditHistoryThirdParty**](FilesFilesAPI.md#GetEditHistoryThirdParty) | **Get** /api/2.0/files/file/{fileId}/edit/history | Get version history (third-party storage)
[**GetEncryptionInfo**](FilesFilesAPI.md#GetEncryptionInfo) | **Get** /api/2.0/files/{fileId}/access | Get file encryption information
[**GetEncryptionInfoThirdParty**](FilesFilesAPI.md#GetEncryptionInfoThirdParty) | **Get** /api/2.0/files/{fileId}/access | Get file encryption information (third-party storage)
[**GetFileHistory**](FilesFilesAPI.md#GetFileHistory) | **Get** /api/2.0/files/file/{fileId}/log | Get file history
[**GetFileInfo**](FilesFilesAPI.md#GetFileInfo) | **Get** /api/2.0/files/file/{fileId} | Get file information
[**GetFileInfoThirdParty**](FilesFilesAPI.md#GetFileInfoThirdParty) | **Get** /api/2.0/files/file/{fileId} | Get file information (third-party storage)
[**GetFileLinks**](FilesFilesAPI.md#GetFileLinks) | **Get** /api/2.0/files/file/{id}/links | Get file external links
[**GetFileLinksThirdParty**](FilesFilesAPI.md#GetFileLinksThirdParty) | **Get** /api/2.0/files/file/{id}/links | Get file external links (third-party storage)
[**GetFilePrimaryExternalLink**](FilesFilesAPI.md#GetFilePrimaryExternalLink) | **Get** /api/2.0/files/file/{id}/link | Get the file primary external link
[**GetFilePrimaryExternalLinkThirdParty**](FilesFilesAPI.md#GetFilePrimaryExternalLinkThirdParty) | **Get** /api/2.0/files/file/{id}/link | Get the file primary external link (third-party storage)
[**GetFileVersionInfo**](FilesFilesAPI.md#GetFileVersionInfo) | **Get** /api/2.0/files/file/{fileId}/history | Get file versions
[**GetFileVersionInfoThirdParty**](FilesFilesAPI.md#GetFileVersionInfoThirdParty) | **Get** /api/2.0/files/file/{fileId}/history | Get file versions (third-party storage)
[**GetFillResult**](FilesFilesAPI.md#GetFillResult) | **Get** /api/2.0/files/file/fillresult | Get form-filling result
[**GetFormSubmissions**](FilesFilesAPI.md#GetFormSubmissions) | **Get** /api/2.0/files/file/{fileId}/submissions | Get form submission results
[**GetPresignedFileUri**](FilesFilesAPI.md#GetPresignedFileUri) | **Get** /api/2.0/files/file/{fileId}/presigned | Get a signed download address
[**GetPresignedFileUriThirdParty**](FilesFilesAPI.md#GetPresignedFileUriThirdParty) | **Get** /api/2.0/files/file/{fileId}/presigned | Get a signed download address (third-party storage)
[**GetPresignedUri**](FilesFilesAPI.md#GetPresignedUri) | **Get** /api/2.0/files/file/{fileId}/presigneduri | Get file download link
[**GetPresignedUriThirdParty**](FilesFilesAPI.md#GetPresignedUriThirdParty) | **Get** /api/2.0/files/file/{fileId}/presigneduri | Get file download link (third-party storage)
[**GetProtectedFileUsers**](FilesFilesAPI.md#GetProtectedFileUsers) | **Get** /api/2.0/files/file/{fileId}/protectusers | Get users for document protection
[**GetProtectedFileUsersThirdParty**](FilesFilesAPI.md#GetProtectedFileUsersThirdParty) | **Get** /api/2.0/files/file/{fileId}/protectusers | Get users for document protection (third-party storage)
[**GetReferenceData**](FilesFilesAPI.md#GetReferenceData) | **Post** /api/2.0/files/file/referencedata | Resolve a spreadsheet reference
[**GetXlsx**](FilesFilesAPI.md#GetXlsx) | **Get** /api/2.0/files/file/{fileId}/xlsx | Get form report generation status
[**IsFormPDF**](FilesFilesAPI.md#IsFormPDF) | **Get** /api/2.0/files/file/{fileId}/isformpdf | Check the PDF file
[**IsFormPDFThirdParty**](FilesFilesAPI.md#IsFormPDFThirdParty) | **Get** /api/2.0/files/file/{fileId}/isformpdf | Check the PDF file (third-party storage)
[**LockFile**](FilesFilesAPI.md#LockFile) | **Put** /api/2.0/files/file/{fileId}/lock | Lock a file
[**LockFileThirdParty**](FilesFilesAPI.md#LockFileThirdParty) | **Put** /api/2.0/files/file/{fileId}/lock | Lock a file (third-party storage)
[**ManageFormFilling**](FilesFilesAPI.md#ManageFormFilling) | **Put** /api/2.0/files/file/{fileId}/manageformfilling | Perform form filling action
[**OpenEditFile**](FilesFilesAPI.md#OpenEditFile) | **Get** /api/2.0/files/file/{fileId}/openedit | Get the editor configuration
[**OpenEditFileThirdParty**](FilesFilesAPI.md#OpenEditFileThirdParty) | **Get** /api/2.0/files/file/{fileId}/openedit | Get the editor configuration (third-party storage)
[**RestoreFileVersion**](FilesFilesAPI.md#RestoreFileVersion) | **Post** /api/2.0/files/file/{fileId}/restoreversion | Restore a file version
[**RestoreFileVersionThirdParty**](FilesFilesAPI.md#RestoreFileVersionThirdParty) | **Post** /api/2.0/files/file/{fileId}/restoreversion | Restore a file version (third-party storage)
[**SaveEditingFileFromForm**](FilesFilesAPI.md#SaveEditingFileFromForm) | **Put** /api/2.0/files/file/{fileId}/saveediting | Save edited file content
[**SaveEditingFileFromFormThirdParty**](FilesFilesAPI.md#SaveEditingFileFromFormThirdParty) | **Put** /api/2.0/files/file/{fileId}/saveediting | Save edited file content (third-party storage)
[**SaveFileAsPdf**](FilesFilesAPI.md#SaveFileAsPdf) | **Post** /api/2.0/files/file/{id}/saveaspdf | Save a file as PDF
[**SaveFileAsPdfThirdParty**](FilesFilesAPI.md#SaveFileAsPdfThirdParty) | **Post** /api/2.0/files/file/{id}/saveaspdf | Save a file as PDF (third-party storage)
[**SaveFormRoleMapping**](FilesFilesAPI.md#SaveFormRoleMapping) | **Post** /api/2.0/files/file/{fileId}/formrolemapping | Save form role mapping
[**SetCustomFilterTag**](FilesFilesAPI.md#SetCustomFilterTag) | **Put** /api/2.0/files/file/{fileId}/customfilter | Set the Custom Filter editing mode
[**SetCustomFilterTagThirdParty**](FilesFilesAPI.md#SetCustomFilterTagThirdParty) | **Put** /api/2.0/files/file/{fileId}/customfilter | Set the Custom Filter editing mode (third-party storage)
[**SetEncryptionInfo**](FilesFilesAPI.md#SetEncryptionInfo) | **Put** /api/2.0/files/{fileId}/access | Set file encryption information
[**SetEncryptionInfoThirdParty**](FilesFilesAPI.md#SetEncryptionInfoThirdParty) | **Put** /api/2.0/files/{fileId}/access | Set file encryption information (third-party storage)
[**SetFileExternalLink**](FilesFilesAPI.md#SetFileExternalLink) | **Put** /api/2.0/files/file/{id}/links | Set a file external link
[**SetFileExternalLinkThirdParty**](FilesFilesAPI.md#SetFileExternalLinkThirdParty) | **Put** /api/2.0/files/file/{id}/links | Set a file external link (third-party storage)
[**SetFileOrder**](FilesFilesAPI.md#SetFileOrder) | **Put** /api/2.0/files/{fileId}/order | Set file order
[**SetFileOrderThirdParty**](FilesFilesAPI.md#SetFileOrderThirdParty) | **Put** /api/2.0/files/{fileId}/order | Set file order (third-party storage)
[**SetFilesOrder**](FilesFilesAPI.md#SetFilesOrder) | **Put** /api/2.0/files/order | Set order of files
[**StartEditFile**](FilesFilesAPI.md#StartEditFile) | **Post** /api/2.0/files/file/{fileId}/startedit | Open an editing session
[**StartEditFileThirdParty**](FilesFilesAPI.md#StartEditFileThirdParty) | **Post** /api/2.0/files/file/{fileId}/startedit | Open an editing session (third-party storage)
[**StartFillingFile**](FilesFilesAPI.md#StartFillingFile) | **Put** /api/2.0/files/file/{fileId}/startfilling | Start filling a form
[**StartFillingFileThirdParty**](FilesFilesAPI.md#StartFillingFileThirdParty) | **Put** /api/2.0/files/file/{fileId}/startfilling | Start filling a form (third-party storage)
[**ToggleFileFavorite**](FilesFilesAPI.md#ToggleFileFavorite) | **Get** /api/2.0/files/favorites/{fileId} | Set the file favorite status
[**ToggleFileFavoriteThirdParty**](FilesFilesAPI.md#ToggleFileFavoriteThirdParty) | **Get** /api/2.0/files/favorites/{fileId} | Set the file favorite status (third-party storage)
[**TrackEditFile**](FilesFilesAPI.md#TrackEditFile) | **Get** /api/2.0/files/file/{fileId}/trackeditfile | Track an editing session
[**TrackEditFileThirdParty**](FilesFilesAPI.md#TrackEditFileThirdParty) | **Get** /api/2.0/files/file/{fileId}/trackeditfile | Track an editing session (third-party storage)
[**UpdateFile**](FilesFilesAPI.md#UpdateFile) | **Put** /api/2.0/files/file/{fileId} | Update a file
[**UpdateFileThirdParty**](FilesFilesAPI.md#UpdateFileThirdParty) | **Put** /api/2.0/files/file/{fileId} | Update a file (third-party storage)



## AddFileToRecent

> FileWrapper AddFileToRecent(ctx, fileId).Execute()

Add a file to Recent



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
	fileId := int32(10) // int32 | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.AddFileToRecent(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.AddFileToRecent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddFileToRecent`: FileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.AddFileToRecent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFileToRecentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileWrapper**](FileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddFileToRecentThirdParty

> ThirdPartyFileWrapper AddFileToRecentThirdParty(ctx, fileId).Execute()

Add a file to Recent (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/add-file-to-recent-third-party/).

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
	fileId := "10" // string | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.AddFileToRecentThirdParty(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.AddFileToRecentThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddFileToRecentThirdParty`: ThirdPartyFileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.AddFileToRecentThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFileToRecentThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ThirdPartyFileWrapper**](ThirdPartyFileWrapper.md)

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

> FileArrayWrapper ChangeVersionHistory(ctx, fileId).ChangeHistory(changeHistory).Execute()

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
	fileId := int32(1) // int32 | The file whose version history is changed.
	changeHistory := *openapiclient.NewChangeHistory(int32(1)) // ChangeHistory | The change to make to the revision group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.ChangeVersionHistory(context.Background(), fileId).ChangeHistory(changeHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.ChangeVersionHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeVersionHistory`: FileArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.ChangeVersionHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file whose version history is changed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeVersionHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeHistory** | [**ChangeHistory**](ChangeHistory.md) | The change to make to the revision group. | 

### Return type

[**FileArrayWrapper**](FileArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeVersionHistoryThirdParty

> ThirdPartyFileArrayWrapper ChangeVersionHistoryThirdParty(ctx, fileId).ChangeHistory(changeHistory).Execute()

Change version history (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/change-version-history-third-party/).

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
	fileId := "1" // string | The file whose version history is changed.
	changeHistory := *openapiclient.NewChangeHistory(int32(1)) // ChangeHistory | The change to make to the revision group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.ChangeVersionHistoryThirdParty(context.Background(), fileId).ChangeHistory(changeHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.ChangeVersionHistoryThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeVersionHistoryThirdParty`: ThirdPartyFileArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.ChangeVersionHistoryThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file whose version history is changed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeVersionHistoryThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeHistory** | [**ChangeHistory**](ChangeHistory.md) | The change to make to the revision group. | 

### Return type

[**ThirdPartyFileArrayWrapper**](ThirdPartyFileArrayWrapper.md)

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

Open a form draft for filling



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
	fileId := int32(1) // int32 | The identifier of the PDF form to open, as it is returned by a room listing such as  `GET api/2.0/files/{folderId}`. The identifier of an already created draft is accepted here as well.
	checkFillFormDraft := *openapiclient.NewCheckFillFormDraft(int32(0)) // CheckFillFormDraft | The revision of the form to open and what the caller intends to do with it.

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
**fileId** | **int32** | The identifier of the PDF form to open, as it is returned by a room listing such as  `GET api/2.0/files/{folderId}`. The identifier of an already created draft is accepted here as well. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckFillFormDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkFillFormDraft** | [**CheckFillFormDraft**](CheckFillFormDraft.md) | The revision of the form to open and what the caller intends to do with it. | 

### Return type

[**StringWrapper**](StringWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckFillFormDraftThirdParty

> StringWrapper CheckFillFormDraftThirdParty(ctx, fileId).CheckFillFormDraft(checkFillFormDraft).Execute()

Open a form draft for filling (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/check-fill-form-draft-third-party/).

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
	fileId := "1" // string | The identifier of the PDF form to open, as it is returned by a room listing such as  `GET api/2.0/files/{folderId}`. The identifier of an already created draft is accepted here as well.
	checkFillFormDraft := *openapiclient.NewCheckFillFormDraft(int32(0)) // CheckFillFormDraft | The revision of the form to open and what the caller intends to do with it.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CheckFillFormDraftThirdParty(context.Background(), fileId).CheckFillFormDraft(checkFillFormDraft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CheckFillFormDraftThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckFillFormDraftThirdParty`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CheckFillFormDraftThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The identifier of the PDF form to open, as it is returned by a room listing such as  `GET api/2.0/files/{folderId}`. The identifier of an already created draft is accepted here as well. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckFillFormDraftThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkFillFormDraft** | [**CheckFillFormDraft**](CheckFillFormDraft.md) | The revision of the form to open and what the caller intends to do with it. | 

### Return type

[**StringWrapper**](StringWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

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
	fileId := int32(1) // int32 | The file to copy.
	copyAsJsonElement := *openapiclient.NewCopyAsJsonElement("Document Copy.docx", openapiclient.CopyAsJsonElement_destFolderId{Int32: new(int32)}) // CopyAsJsonElement | The title, the destination and the conversion options of the copy.

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
**fileId** | **int32** | The file to copy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyFileAsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **copyAsJsonElement** | [**CopyAsJsonElement**](CopyAsJsonElement.md) | The title, the destination and the conversion options of the copy. | 

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


## CopyFileAsThirdParty

> FileEntryBaseWrapper CopyFileAsThirdParty(ctx, fileId).CopyAsJsonElement(copyAsJsonElement).Execute()

Copy a file (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/copy-file-as-third-party/).

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
	fileId := "1" // string | The file to copy.
	copyAsJsonElement := *openapiclient.NewCopyAsJsonElement("Document Copy.docx", openapiclient.CopyAsJsonElement_destFolderId{Int32: new(int32)}) // CopyAsJsonElement | The title, the destination and the conversion options of the copy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CopyFileAsThirdParty(context.Background(), fileId).CopyAsJsonElement(copyAsJsonElement).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CopyFileAsThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CopyFileAsThirdParty`: FileEntryBaseWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CopyFileAsThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file to copy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyFileAsThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **copyAsJsonElement** | [**CopyAsJsonElement**](CopyAsJsonElement.md) | The title, the destination and the conversion options of the copy. | 

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

> ChunkedUploadSessionResponseWrapperWrapper CreateEditSession(ctx, fileId).FileSize(fileSize).Execute()

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
	fileId := int32(1) // int32 | The file whose content the session will replace; take the id from a folder listing or from the file itself.
	fileSize := int64(1024) // int64 | The number of bytes the new content will take. It is checked against the portal limit for chunked uploads  before the session opens, and a session left at 0 takes the whole content in a single part. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CreateEditSession(context.Background(), fileId).FileSize(fileSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CreateEditSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEditSession`: ChunkedUploadSessionResponseWrapperWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CreateEditSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file whose content the session will replace; take the id from a folder listing or from the file itself. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEditSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileSize** | **int64** | The number of bytes the new content will take. It is checked against the portal limit for chunked uploads  before the session opens, and a session left at 0 takes the whole content in a single part. | 

### Return type

[**ChunkedUploadSessionResponseWrapperWrapper**](ChunkedUploadSessionResponseWrapperWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEditSessionThirdParty

> ThirdPartyChunkedUploadSessionResponseWrapperWrapper CreateEditSessionThirdParty(ctx, fileId).FileSize(fileSize).Execute()

Create the editing session (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-edit-session-third-party/).

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
	fileId := "1" // string | The file whose content the session will replace; take the id from a folder listing or from the file itself.
	fileSize := int64(1024) // int64 | The number of bytes the new content will take. It is checked against the portal limit for chunked uploads  before the session opens, and a session left at 0 takes the whole content in a single part. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CreateEditSessionThirdParty(context.Background(), fileId).FileSize(fileSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CreateEditSessionThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEditSessionThirdParty`: ThirdPartyChunkedUploadSessionResponseWrapperWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CreateEditSessionThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file whose content the session will replace; take the id from a folder listing or from the file itself. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEditSessionThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileSize** | **int64** | The number of bytes the new content will take. It is checked against the portal limit for chunked uploads  before the session opens, and a session left at 0 takes the whole content in a single part. | 

### Return type

[**ThirdPartyChunkedUploadSessionResponseWrapperWrapper**](ThirdPartyChunkedUploadSessionResponseWrapperWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFile

> FileWrapper CreateFile(ctx, folderId).CreateFileJsonElement(createFileJsonElement).Execute()

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
	folderId := int32(1) // int32 | The folder the file is created in.
	createFileJsonElement := *openapiclient.NewCreateFileJsonElement("New Document.docx") // CreateFileJsonElement | The title of the new file and the source of its content.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CreateFile(context.Background(), folderId).CreateFileJsonElement(createFileJsonElement).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CreateFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFile`: FileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CreateFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder the file is created in. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFileJsonElement** | [**CreateFileJsonElement**](CreateFileJsonElement.md) | The title of the new file and the source of its content. | 

### Return type

[**FileWrapper**](FileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFileThirdParty

> ThirdPartyFileWrapper CreateFileThirdParty(ctx, folderId).CreateFileJsonElement(createFileJsonElement).Execute()

Create a file (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-file-third-party/).

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
	folderId := "1" // string | The folder the file is created in.
	createFileJsonElement := *openapiclient.NewCreateFileJsonElement("New Document.docx") // CreateFileJsonElement | The title of the new file and the source of its content.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CreateFileThirdParty(context.Background(), folderId).CreateFileJsonElement(createFileJsonElement).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CreateFileThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFileThirdParty`: ThirdPartyFileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CreateFileThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | The folder the file is created in. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFileThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFileJsonElement** | [**CreateFileJsonElement**](CreateFileJsonElement.md) | The title of the new file and the source of its content. | 

### Return type

[**ThirdPartyFileWrapper**](ThirdPartyFileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFileInMyDocuments

> FileWrapper CreateFileInMyDocuments(ctx).CreateFileJsonElement(createFileJsonElement).Execute()

Create a file in My documents



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
	// response from `CreateFileInMyDocuments`: FileWrapper
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

[**FileWrapper**](FileWrapper.md)

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

Create the file primary external link



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
	id := int32(1) // int32 | The file the link points at.
	fileLinkRequest := *openapiclient.NewFileLinkRequest() // FileLinkRequest | The settings of the link. They are applied in full, so a field left out is reset rather than kept.

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
**id** | **int32** | The file the link points at. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFilePrimaryExternalLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileLinkRequest** | [**FileLinkRequest**](FileLinkRequest.md) | The settings of the link. They are applied in full, so a field left out is reset rather than kept. | 

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


## CreateFilePrimaryExternalLinkThirdParty

> FileShareWrapper CreateFilePrimaryExternalLinkThirdParty(ctx, id).FileLinkRequest(fileLinkRequest).Execute()

Create the file primary external link (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-file-primary-external-link-third-party/).

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
	id := "1" // string | The file the link points at.
	fileLinkRequest := *openapiclient.NewFileLinkRequest() // FileLinkRequest | The settings of the link. They are applied in full, so a field left out is reset rather than kept.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CreateFilePrimaryExternalLinkThirdParty(context.Background(), id).FileLinkRequest(fileLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CreateFilePrimaryExternalLinkThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFilePrimaryExternalLinkThirdParty`: FileShareWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CreateFilePrimaryExternalLinkThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The file the link points at. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFilePrimaryExternalLinkThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileLinkRequest** | [**FileLinkRequest**](FileLinkRequest.md) | The settings of the link. They are applied in full, so a field left out is reset rather than kept. | 

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

> FileWrapper CreateHtmlFile(ctx, folderId).CreateTextOrHtmlFile(createTextOrHtmlFile).Execute()

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
	folderId := int32(1) // int32 | The folder the file is created in.
	createTextOrHtmlFile := *openapiclient.NewCreateTextOrHtmlFile("Document.txt") // CreateTextOrHtmlFile | The title, the content and the collision behaviour of the new file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CreateHtmlFile(context.Background(), folderId).CreateTextOrHtmlFile(createTextOrHtmlFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CreateHtmlFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHtmlFile`: FileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CreateHtmlFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder the file is created in. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHtmlFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTextOrHtmlFile** | [**CreateTextOrHtmlFile**](CreateTextOrHtmlFile.md) | The title, the content and the collision behaviour of the new file. | 

### Return type

[**FileWrapper**](FileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHtmlFileThirdParty

> ThirdPartyFileWrapper CreateHtmlFileThirdParty(ctx, folderId).CreateTextOrHtmlFile(createTextOrHtmlFile).Execute()

Create an HTML file (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-html-file-third-party/).

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
	folderId := "1" // string | The folder the file is created in.
	createTextOrHtmlFile := *openapiclient.NewCreateTextOrHtmlFile("Document.txt") // CreateTextOrHtmlFile | The title, the content and the collision behaviour of the new file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CreateHtmlFileThirdParty(context.Background(), folderId).CreateTextOrHtmlFile(createTextOrHtmlFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CreateHtmlFileThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHtmlFileThirdParty`: ThirdPartyFileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CreateHtmlFileThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | The folder the file is created in. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHtmlFileThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTextOrHtmlFile** | [**CreateTextOrHtmlFile**](CreateTextOrHtmlFile.md) | The title, the content and the collision behaviour of the new file. | 

### Return type

[**ThirdPartyFileWrapper**](ThirdPartyFileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHtmlFileInMyDocuments

> FileWrapper CreateHtmlFileInMyDocuments(ctx).CreateTextOrHtmlFile(createTextOrHtmlFile).Execute()

Create an HTML file in My documents



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
	// response from `CreateHtmlFileInMyDocuments`: FileWrapper
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

[**FileWrapper**](FileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTextFile

> FileWrapper CreateTextFile(ctx, folderId).CreateTextOrHtmlFile(createTextOrHtmlFile).Execute()

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
	folderId := int32(1) // int32 | The folder the file is created in.
	createTextOrHtmlFile := *openapiclient.NewCreateTextOrHtmlFile("Document.txt") // CreateTextOrHtmlFile | The title, the content and the collision behaviour of the new file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CreateTextFile(context.Background(), folderId).CreateTextOrHtmlFile(createTextOrHtmlFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CreateTextFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTextFile`: FileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CreateTextFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder the file is created in. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTextFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTextOrHtmlFile** | [**CreateTextOrHtmlFile**](CreateTextOrHtmlFile.md) | The title, the content and the collision behaviour of the new file. | 

### Return type

[**FileWrapper**](FileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTextFileThirdParty

> ThirdPartyFileWrapper CreateTextFileThirdParty(ctx, folderId).CreateTextOrHtmlFile(createTextOrHtmlFile).Execute()

Create a text file (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-text-file-third-party/).

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
	folderId := "1" // string | The folder the file is created in.
	createTextOrHtmlFile := *openapiclient.NewCreateTextOrHtmlFile("Document.txt") // CreateTextOrHtmlFile | The title, the content and the collision behaviour of the new file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.CreateTextFileThirdParty(context.Background(), folderId).CreateTextOrHtmlFile(createTextOrHtmlFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.CreateTextFileThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTextFileThirdParty`: ThirdPartyFileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.CreateTextFileThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | The folder the file is created in. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTextFileThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTextOrHtmlFile** | [**CreateTextOrHtmlFile**](CreateTextOrHtmlFile.md) | The title, the content and the collision behaviour of the new file. | 

### Return type

[**ThirdPartyFileWrapper**](ThirdPartyFileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTextFileInMyDocuments

> FileWrapper CreateTextFileInMyDocuments(ctx).CreateTextOrHtmlFile(createTextOrHtmlFile).Execute()

Create a text file in My documents



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
	// response from `CreateTextFileInMyDocuments`: FileWrapper
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

[**FileWrapper**](FileWrapper.md)

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

Queue file thumbnails



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

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

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
	fileId := int32(1) // int32 | The file to delete.
	delete := *openapiclient.NewDelete() // Delete | When and how the file is deleted.
	returnSingleOperation := false // bool | Which operations the answer carries: `true` returns the operation this call started and nothing else, `false`  returns every operation of the same kind that the caller has running or unread. When nothing was queued, which  happens for an empty selection, `true` falls back to the full list. (optional)

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
**fileId** | **int32** | The file to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **delete** | [**Delete**](Delete.md) | When and how the file is deleted. | 
 **returnSingleOperation** | **bool** | Which operations the answer carries: `true` returns the operation this call started and nothing else, `false`  returns every operation of the same kind that the caller has running or unread. When nothing was queued, which  happens for an empty selection, `true` falls back to the full list. | 

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


## DeleteFileThirdParty

> FileOperationArrayWrapper DeleteFileThirdParty(ctx, fileId).Delete(delete).ReturnSingleOperation(returnSingleOperation).Execute()

Delete a file (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-file-third-party/).

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
	fileId := "1" // string | The file to delete.
	delete := *openapiclient.NewDelete() // Delete | When and how the file is deleted.
	returnSingleOperation := false // bool | Which operations the answer carries: `true` returns the operation this call started and nothing else, `false`  returns every operation of the same kind that the caller has running or unread. When nothing was queued, which  happens for an empty selection, `true` falls back to the full list. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.DeleteFileThirdParty(context.Background(), fileId).Delete(delete).ReturnSingleOperation(returnSingleOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.DeleteFileThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFileThirdParty`: FileOperationArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.DeleteFileThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **delete** | [**Delete**](Delete.md) | When and how the file is deleted. | 
 **returnSingleOperation** | **bool** | Which operations the answer carries: `true` returns the operation this call started and nothing else, `false`  returns every operation of the same kind that the caller has running or unread. When nothing was queued, which  happens for an empty selection, `true` falls back to the full list. | 

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

> DeleteRecent(ctx).BaseBatchRequestDto(baseBatchRequestDto).Execute()

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
	r, err := apiClient.FilesFilesAPI.DeleteRecent(context.Background()).BaseBatchRequestDto(baseBatchRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.DeleteRecent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseBatchRequestDto** | [**BaseBatchRequestDto**](BaseBatchRequestDto.md) |  | 

### Return type

 (empty response body)

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
	requestBody := []int32{int32(123)} // []int32 | The files to take off the template list, by id; this array is the whole request body. Only a file stored in  the portal itself can be a template, which is why an id here is always numeric. (optional)

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
 **requestBody** | **[]int32** | The files to take off the template list, by id; this array is the whole request body. Only a file stored in  the portal itself can be a template, which is why an id here is always numeric. | 

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

Generate a form answers report



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
	fileId := int32(10) // int32 | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

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
**fileId** | **int32** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

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
	fileId := int32(10) // int32 | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

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
**fileId** | **int32** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

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


## GetAllFormRolesThirdParty

> FormRoleArrayWrapper GetAllFormRolesThirdParty(ctx, fileId).Execute()

Get form roles (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-all-form-roles-third-party/).

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
	fileId := "10" // string | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetAllFormRolesThirdParty(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetAllFormRolesThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllFormRolesThirdParty`: FormRoleArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetAllFormRolesThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllFormRolesThirdPartyRequest struct via the builder pattern


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
	fileId := int32(1) // int32 | The file whose changes are read.
	version := int32(1) // int32 | The version to show the changes of, as reported by `GET api/2.0/files/file/{fileId}/edit/history`; 0 means the  current version. (optional)

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
**fileId** | **int32** | The file whose changes are read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEditDiffUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** | The version to show the changes of, as reported by `GET api/2.0/files/file/{fileId}/edit/history`; 0 means the  current version. | 

### Return type

[**EditHistoryDataWrapper**](EditHistoryDataWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEditDiffUrlThirdParty

> EditHistoryDataWrapper GetEditDiffUrlThirdParty(ctx, fileId).Version(version).Execute()

Get changes URL (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-edit-diff-url-third-party/).

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
	fileId := "1" // string | The file whose changes are read.
	version := int32(1) // int32 | The version to show the changes of, as reported by `GET api/2.0/files/file/{fileId}/edit/history`; 0 means the  current version. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetEditDiffUrlThirdParty(context.Background(), fileId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetEditDiffUrlThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEditDiffUrlThirdParty`: EditHistoryDataWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetEditDiffUrlThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file whose changes are read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEditDiffUrlThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** | The version to show the changes of, as reported by `GET api/2.0/files/file/{fileId}/edit/history`; 0 means the  current version. | 

### Return type

[**EditHistoryDataWrapper**](EditHistoryDataWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

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
	fileId := int32(10) // int32 | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

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
**fileId** | **int32** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEditHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EditHistoryArrayWrapper**](EditHistoryArrayWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEditHistoryThirdParty

> EditHistoryArrayWrapper GetEditHistoryThirdParty(ctx, fileId).Execute()

Get version history (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-edit-history-third-party/).

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
	fileId := "10" // string | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetEditHistoryThirdParty(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetEditHistoryThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEditHistoryThirdParty`: EditHistoryArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetEditHistoryThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEditHistoryThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EditHistoryArrayWrapper**](EditHistoryArrayWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEncryptionInfo

> FileEncryptionInfoWrapper GetEncryptionInfo(ctx, fileId).Execute()

Get file encryption information



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-encryption-info/).

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
	fileId := int32(56) // int32 | The file whose encryption keys are read. Only a file in an end-to-end encrypted              private room has any.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetEncryptionInfo(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetEncryptionInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEncryptionInfo`: FileEncryptionInfoWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetEncryptionInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file whose encryption keys are read. Only a file in an end-to-end encrypted              private room has any. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncryptionInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileEncryptionInfoWrapper**](FileEncryptionInfoWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEncryptionInfoThirdParty

> FileEncryptionInfoWrapper GetEncryptionInfoThirdParty(ctx, fileId).Execute()

Get file encryption information (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-encryption-info-third-party/).

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
	fileId := "fileId_example" // string | The file whose encryption keys are read. Only a file in an end-to-end encrypted              private room has any.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetEncryptionInfoThirdParty(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetEncryptionInfoThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEncryptionInfoThirdParty`: FileEncryptionInfoWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetEncryptionInfoThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file whose encryption keys are read. Only a file in an end-to-end encrypted              private room has any. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncryptionInfoThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileEncryptionInfoWrapper**](FileEncryptionInfoWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

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
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fileId := int32(1) // int32 | The file whose activity log is read; only files stored in the portal itself have one.
	fromDate := time.Now() // time.Time | The earliest moment an entry may have, read in the time zone of the portal; left out, the log starts at the  oldest entry the portal still keeps. (optional)
	toDate := time.Now() // time.Time | The latest moment an entry may have, read in the time zone of the portal; left out, the log ends at the newest  entry. (optional)
	count := int32(25) // int32 | How many entries one page holds. The number of entries that match the query is reported in the response  headers, not in the body. (optional)
	startIndex := int32(0) // int32 | How many entries to skip before the page begins, counted from the newest one, so pages are taken by adding the  page size to it. (optional)

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
**fileId** | **int32** | The file whose activity log is read; only files stored in the portal itself have one. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileHistoryRequest struct via the builder pattern


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


## GetFileInfo

> FileWrapper GetFileInfo(ctx, fileId).Version(version).Execute()

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
	fileId := int32(1) // int32 | The file to read.
	version := int32(1) // int32 | The version to read, as reported by `GET api/2.0/files/file/{fileId}/history`; -1, the default, reads the  current version. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetFileInfo(context.Background(), fileId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetFileInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileInfo`: FileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetFileInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** | The version to read, as reported by `GET api/2.0/files/file/{fileId}/history`; -1, the default, reads the  current version. | 

### Return type

[**FileWrapper**](FileWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileInfoThirdParty

> ThirdPartyFileWrapper GetFileInfoThirdParty(ctx, fileId).Version(version).Execute()

Get file information (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-file-info-third-party/).

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
	fileId := "1" // string | The file to read.
	version := int32(1) // int32 | The version to read, as reported by `GET api/2.0/files/file/{fileId}/history`; -1, the default, reads the  current version. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetFileInfoThirdParty(context.Background(), fileId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetFileInfoThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileInfoThirdParty`: ThirdPartyFileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetFileInfoThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file to read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileInfoThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** | The version to read, as reported by `GET api/2.0/files/file/{fileId}/history`; -1, the default, reads the  current version. | 

### Return type

[**ThirdPartyFileWrapper**](ThirdPartyFileWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

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
	id := int32(10) // int32 | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.
	count := int32(25) // int32 | How many entries at most to answer with, in the operations of this file that return a list; an operation that  answers with a single object is not affected by it. (optional)
	startIndex := int32(0) // int32 | How many entries of such a list to skip before answering, used together with `count` to walk through it page  by page. (optional)

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
**id** | **int32** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | How many entries at most to answer with, in the operations of this file that return a list; an operation that  answers with a single object is not affected by it. | 
 **startIndex** | **int32** | How many entries of such a list to skip before answering, used together with `count` to walk through it page  by page. | 

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


## GetFileLinksThirdParty

> FileShareArrayWrapper GetFileLinksThirdParty(ctx, id).Count(count).StartIndex(startIndex).Execute()

Get file external links (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-file-links-third-party/).

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
	id := "10" // string | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.
	count := int32(25) // int32 | How many entries at most to answer with, in the operations of this file that return a list; an operation that  answers with a single object is not affected by it. (optional)
	startIndex := int32(0) // int32 | How many entries of such a list to skip before answering, used together with `count` to walk through it page  by page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetFileLinksThirdParty(context.Background(), id).Count(count).StartIndex(startIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetFileLinksThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileLinksThirdParty`: FileShareArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetFileLinksThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileLinksThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | How many entries at most to answer with, in the operations of this file that return a list; an operation that  answers with a single object is not affected by it. | 
 **startIndex** | **int32** | How many entries of such a list to skip before answering, used together with `count` to walk through it page  by page. | 

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

Get the file primary external link



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
	id := int32(10) // int32 | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.
	count := int32(25) // int32 | How many entries at most to answer with, in the operations of this file that return a list; an operation that  answers with a single object is not affected by it. (optional)
	startIndex := int32(0) // int32 | How many entries of such a list to skip before answering, used together with `count` to walk through it page  by page. (optional)

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
**id** | **int32** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilePrimaryExternalLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | How many entries at most to answer with, in the operations of this file that return a list; an operation that  answers with a single object is not affected by it. | 
 **startIndex** | **int32** | How many entries of such a list to skip before answering, used together with `count` to walk through it page  by page. | 

### Return type

[**FileShareWrapper**](FileShareWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilePrimaryExternalLinkThirdParty

> FileShareWrapper GetFilePrimaryExternalLinkThirdParty(ctx, id).Count(count).StartIndex(startIndex).Execute()

Get the file primary external link (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-file-primary-external-link-third-party/).

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
	id := "10" // string | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.
	count := int32(25) // int32 | How many entries at most to answer with, in the operations of this file that return a list; an operation that  answers with a single object is not affected by it. (optional)
	startIndex := int32(0) // int32 | How many entries of such a list to skip before answering, used together with `count` to walk through it page  by page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetFilePrimaryExternalLinkThirdParty(context.Background(), id).Count(count).StartIndex(startIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetFilePrimaryExternalLinkThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilePrimaryExternalLinkThirdParty`: FileShareWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetFilePrimaryExternalLinkThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilePrimaryExternalLinkThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | How many entries at most to answer with, in the operations of this file that return a list; an operation that  answers with a single object is not affected by it. | 
 **startIndex** | **int32** | How many entries of such a list to skip before answering, used together with `count` to walk through it page  by page. | 

### Return type

[**FileShareWrapper**](FileShareWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileVersionInfo

> FileArrayWrapper GetFileVersionInfo(ctx, fileId).Execute()

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
	fileId := int32(10) // int32 | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetFileVersionInfo(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetFileVersionInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileVersionInfo`: FileArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetFileVersionInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileVersionInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileArrayWrapper**](FileArrayWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileVersionInfoThirdParty

> ThirdPartyFileArrayWrapper GetFileVersionInfoThirdParty(ctx, fileId).Execute()

Get file versions (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-file-version-info-third-party/).

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
	fileId := "10" // string | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetFileVersionInfoThirdParty(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetFileVersionInfoThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileVersionInfoThirdParty`: ThirdPartyFileArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetFileVersionInfoThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileVersionInfoThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ThirdPartyFileArrayWrapper**](ThirdPartyFileArrayWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFillResult

> FillingFormResultWrapper GetFillResult(ctx).FillingSessionId(fillingSessionId).Execute()

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
	fillingSessionId := "11111111-2222-3333-4444-555555555555" // string | The identifier of the finished filling session, the value the document service reports when the filling ends.  The portal remembers it only for a while afterwards, so an older session is answered as not found. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetFillResult(context.Background()).FillingSessionId(fillingSessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetFillResult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFillResult`: FillingFormResultWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetFillResult`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFillResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fillingSessionId** | **string** | The identifier of the finished filling session, the value the document service reports when the filling ends.  The portal remembers it only for a while afterwards, so an older session is answered as not found. | 

### Return type

[**FillingFormResultWrapper**](FillingFormResultWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

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
	fileId := int32(10) // int32 | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

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
**fileId** | **int32** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

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

Get a signed download address



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
	fileId := int32(10) // int32 | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

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
**fileId** | **int32** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

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


## GetPresignedFileUriThirdParty

> FileLinkWrapper GetPresignedFileUriThirdParty(ctx, fileId).Execute()

Get a signed download address (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-presigned-file-uri-third-party/).

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
	fileId := "10" // string | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetPresignedFileUriThirdParty(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetPresignedFileUriThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPresignedFileUriThirdParty`: FileLinkWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetPresignedFileUriThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPresignedFileUriThirdPartyRequest struct via the builder pattern


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
	fileId := int32(10) // int32 | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

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
**fileId** | **int32** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

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


## GetPresignedUriThirdParty

> StringWrapper GetPresignedUriThirdParty(ctx, fileId).Execute()

Get file download link (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-presigned-uri-third-party/).

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
	fileId := "10" // string | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetPresignedUriThirdParty(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetPresignedUriThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPresignedUriThirdParty`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetPresignedUriThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPresignedUriThirdPartyRequest struct via the builder pattern


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

Get users for document protection



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
	fileId := int32(10) // int32 | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

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
**fileId** | **int32** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

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


## GetProtectedFileUsersThirdParty

> MentionWrapperArrayWrapper GetProtectedFileUsersThirdParty(ctx, fileId).Execute()

Get users for document protection (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-protected-file-users-third-party/).

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
	fileId := "10" // string | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetProtectedFileUsersThirdParty(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.GetProtectedFileUsersThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProtectedFileUsersThirdParty`: MentionWrapperArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.GetProtectedFileUsersThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProtectedFileUsersThirdPartyRequest struct via the builder pattern


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

> FileReferenceWrapper GetReferenceData(ctx).GetReferenceDataDto(getReferenceDataDto).Execute()

Resolve a spreadsheet reference



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
	getReferenceDataDto := *openapiclient.NewGetReferenceDataDto("512", "1") // GetReferenceDataDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.GetReferenceData(context.Background()).GetReferenceDataDto(getReferenceDataDto).Execute()
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
 **getReferenceDataDto** | [**GetReferenceDataDto**](GetReferenceDataDto.md) |  | 

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

Get form report generation status



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
	fileId := int32(10) // int32 | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

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
**fileId** | **int32** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

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
	fileId := int32(10) // int32 | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

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
**fileId** | **int32** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

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


## IsFormPDFThirdParty

> BooleanWrapper IsFormPDFThirdParty(ctx, fileId).Execute()

Check the PDF file (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/is-form-pdf-third-party/).

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
	fileId := "10" // string | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.IsFormPDFThirdParty(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.IsFormPDFThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IsFormPDFThirdParty`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.IsFormPDFThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIsFormPDFThirdPartyRequest struct via the builder pattern


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

> FileWrapper LockFile(ctx, fileId).LockFileParameters(lockFileParameters).Execute()

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
	fileId := int32(1) // int32 | The file to lock or unlock.
	lockFileParameters := *openapiclient.NewLockFileParameters() // LockFileParameters | The lock state to reach.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.LockFile(context.Background(), fileId).LockFileParameters(lockFileParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.LockFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LockFile`: FileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.LockFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file to lock or unlock. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lockFileParameters** | [**LockFileParameters**](LockFileParameters.md) | The lock state to reach. | 

### Return type

[**FileWrapper**](FileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LockFileThirdParty

> ThirdPartyFileWrapper LockFileThirdParty(ctx, fileId).LockFileParameters(lockFileParameters).Execute()

Lock a file (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/lock-file-third-party/).

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
	fileId := "1" // string | The file to lock or unlock.
	lockFileParameters := *openapiclient.NewLockFileParameters() // LockFileParameters | The lock state to reach.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.LockFileThirdParty(context.Background(), fileId).LockFileParameters(lockFileParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.LockFileThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LockFileThirdParty`: ThirdPartyFileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.LockFileThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file to lock or unlock. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockFileThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lockFileParameters** | [**LockFileParameters**](LockFileParameters.md) | The lock state to reach. | 

### Return type

[**ThirdPartyFileWrapper**](ThirdPartyFileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageFormFilling

> ManageFormFilling(ctx, fileId).ManageFormFillingDto(manageFormFillingDto).Execute()

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
	fileId := "fileId_example" // string | The form the action applies to. Send the same value as the `formId` of the request body, which is the one the handler reads.
	manageFormFillingDto := *openapiclient.NewManageFormFillingDto(int32(1)) // ManageFormFillingDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesFilesAPI.ManageFormFilling(context.Background(), fileId).ManageFormFillingDto(manageFormFillingDto).Execute()
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
**fileId** | **string** | The form the action applies to. Send the same value as the `formId` of the request body, which is the one the handler reads. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManageFormFillingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **manageFormFillingDto** | [**ManageFormFillingDto**](ManageFormFillingDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenEditFile

> ConfigurationWrapper OpenEditFile(ctx, fileId).Version(version).View(view).EditorType(editorType).Edit(edit).Fill(fill).Execute()

Get the editor configuration



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
	fileId := int32(1) // int32 | The file the editor configuration is built for. Take the id from a folder listing such as  `GET api/2.0/files/{folderId}`.
	version := int32(1) // int32 | Which entry of the file history to open, numbered the way the file versions are. Left out, the current  revision is opened; naming a version requires access to the history of the file. (optional)
	view := false // bool | Asks for a read-only configuration. Left off, the configuration is built for editing as far as the caller's  rights and the room the file lies in allow. (optional)
	editorType := openapiclient.EditorType(0) // EditorType | Which editor layout the configuration is built for: the full desktop interface, the reduced mobile one, or the  embedded viewer meant to be framed inside another page. (optional)
	edit := false // bool | Asks for editing rather than viewing. On a form in a form-filling room this also records that the form is  being edited; the room may still turn the request into viewing or into filling. (optional)
	fill := false // bool | Asks for a PDF form to open for filling out rather than for editing. It has no effect on a file that is not a  form. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.OpenEditFile(context.Background(), fileId).Version(version).View(view).EditorType(editorType).Edit(edit).Fill(fill).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.OpenEditFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenEditFile`: ConfigurationWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.OpenEditFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file the editor configuration is built for. Take the id from a folder listing such as  `GET api/2.0/files/{folderId}`. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpenEditFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** | Which entry of the file history to open, numbered the way the file versions are. Left out, the current  revision is opened; naming a version requires access to the history of the file. | 
 **view** | **bool** | Asks for a read-only configuration. Left off, the configuration is built for editing as far as the caller's  rights and the room the file lies in allow. | 
 **editorType** | [**EditorType**](EditorType.md) | Which editor layout the configuration is built for: the full desktop interface, the reduced mobile one, or the  embedded viewer meant to be framed inside another page. | 
 **edit** | **bool** | Asks for editing rather than viewing. On a form in a form-filling room this also records that the form is  being edited; the room may still turn the request into viewing or into filling. | 
 **fill** | **bool** | Asks for a PDF form to open for filling out rather than for editing. It has no effect on a file that is not a  form. | 

### Return type

[**ConfigurationWrapper**](ConfigurationWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenEditFileThirdParty

> ThirdPartyConfigurationWrapper OpenEditFileThirdParty(ctx, fileId).Version(version).View(view).EditorType(editorType).Edit(edit).Fill(fill).Execute()

Get the editor configuration (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/open-edit-file-third-party/).

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
	fileId := "1" // string | The file the editor configuration is built for. Take the id from a folder listing such as  `GET api/2.0/files/{folderId}`.
	version := int32(1) // int32 | Which entry of the file history to open, numbered the way the file versions are. Left out, the current  revision is opened; naming a version requires access to the history of the file. (optional)
	view := false // bool | Asks for a read-only configuration. Left off, the configuration is built for editing as far as the caller's  rights and the room the file lies in allow. (optional)
	editorType := openapiclient.EditorType(0) // EditorType | Which editor layout the configuration is built for: the full desktop interface, the reduced mobile one, or the  embedded viewer meant to be framed inside another page. (optional)
	edit := false // bool | Asks for editing rather than viewing. On a form in a form-filling room this also records that the form is  being edited; the room may still turn the request into viewing or into filling. (optional)
	fill := false // bool | Asks for a PDF form to open for filling out rather than for editing. It has no effect on a file that is not a  form. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.OpenEditFileThirdParty(context.Background(), fileId).Version(version).View(view).EditorType(editorType).Edit(edit).Fill(fill).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.OpenEditFileThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenEditFileThirdParty`: ThirdPartyConfigurationWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.OpenEditFileThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file the editor configuration is built for. Take the id from a folder listing such as  `GET api/2.0/files/{folderId}`. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpenEditFileThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** | Which entry of the file history to open, numbered the way the file versions are. Left out, the current  revision is opened; naming a version requires access to the history of the file. | 
 **view** | **bool** | Asks for a read-only configuration. Left off, the configuration is built for editing as far as the caller's  rights and the room the file lies in allow. | 
 **editorType** | [**EditorType**](EditorType.md) | Which editor layout the configuration is built for: the full desktop interface, the reduced mobile one, or the  embedded viewer meant to be framed inside another page. | 
 **edit** | **bool** | Asks for editing rather than viewing. On a form in a form-filling room this also records that the form is  being edited; the room may still turn the request into viewing or into filling. | 
 **fill** | **bool** | Asks for a PDF form to open for filling out rather than for editing. It has no effect on a file that is not a  form. | 

### Return type

[**ThirdPartyConfigurationWrapper**](ThirdPartyConfigurationWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

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
	fileId := int32(1) // int32 | The file whose version is restored.
	version := int32(1) // int32 | The version to restore, as reported by `GET api/2.0/files/file/{fileId}/edit/history`. It has to name an  existing version that is not the current one. (optional)
	url := "https://document-server.example.com/cache/files/conv_1_docx/output.docx" // string | The address the content of the new version is fetched from instead of the stored version, which is how the  document service hands back a document with a set of changes rolled back; left out, the stored version is  used. (optional)

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
**fileId** | **int32** | The file whose version is restored. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreFileVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** | The version to restore, as reported by `GET api/2.0/files/file/{fileId}/edit/history`. It has to name an  existing version that is not the current one. | 
 **url** | **string** | The address the content of the new version is fetched from instead of the stored version, which is how the  document service hands back a document with a set of changes rolled back; left out, the stored version is  used. | 

### Return type

[**EditHistoryArrayWrapper**](EditHistoryArrayWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreFileVersionThirdParty

> EditHistoryArrayWrapper RestoreFileVersionThirdParty(ctx, fileId).Version(version).Url(url).Execute()

Restore a file version (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/restore-file-version-third-party/).

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
	fileId := "1" // string | The file whose version is restored.
	version := int32(1) // int32 | The version to restore, as reported by `GET api/2.0/files/file/{fileId}/edit/history`. It has to name an  existing version that is not the current one. (optional)
	url := "https://document-server.example.com/cache/files/conv_1_docx/output.docx" // string | The address the content of the new version is fetched from instead of the stored version, which is how the  document service hands back a document with a set of changes rolled back; left out, the stored version is  used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.RestoreFileVersionThirdParty(context.Background(), fileId).Version(version).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.RestoreFileVersionThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreFileVersionThirdParty`: EditHistoryArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.RestoreFileVersionThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file whose version is restored. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreFileVersionThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** | The version to restore, as reported by `GET api/2.0/files/file/{fileId}/edit/history`. It has to name an  existing version that is not the current one. | 
 **url** | **string** | The address the content of the new version is fetched from instead of the stored version, which is how the  document service hands back a document with a set of changes rolled back; left out, the stored version is  used. | 

### Return type

[**EditHistoryArrayWrapper**](EditHistoryArrayWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveEditingFileFromForm

> FileWrapper SaveEditingFileFromForm(ctx, fileId).DownloadUri(downloadUri).FileExtension(fileExtension).File(file).Forcesave(forcesave).Execute()

Save edited file content



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
	fileId := int32(1) // int32 | The file whose content is replaced. The submitted content is written onto this file, so it has to be the file  the editing session was opened on rather than a copy of it.
	downloadUri := "https://example.com/file.txt" // string | An address the document service saved the document at. This operation does not fetch the content from it - the  content always comes from the request body - and reads it only for the extension, when no file extension is  given. (optional)
	fileExtension := "fileExtension_example" // string | The format the submitted content is in, with the leading dot, as in `.docx`. When it differs from the format  the file is stored in, the portal converts the content before saving it. Left empty, the extension is read off  the download address, and failing that the stored format is assumed. (optional)
	file := os.NewFile(1234, "some_file") // *os.File | The edited content, sent as the `File` part of a `multipart/form-data` body. When the part is missing the raw  request body is saved as the content instead, so an empty body empties the file. (optional)
	forcesave := true // bool | Records the write as an editor autosave: the file keeps its running editing session and the previous autosave  revision is overwritten. Left off, the write closes the solo editing session, is refused while somebody else  has the file open, and adds a version to the history. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.SaveEditingFileFromForm(context.Background(), fileId).DownloadUri(downloadUri).FileExtension(fileExtension).File(file).Forcesave(forcesave).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.SaveEditingFileFromForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveEditingFileFromForm`: FileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.SaveEditingFileFromForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file whose content is replaced. The submitted content is written onto this file, so it has to be the file  the editing session was opened on rather than a copy of it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveEditingFileFromFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **downloadUri** | **string** | An address the document service saved the document at. This operation does not fetch the content from it - the  content always comes from the request body - and reads it only for the extension, when no file extension is  given. | 
 **fileExtension** | **string** | The format the submitted content is in, with the leading dot, as in `.docx`. When it differs from the format  the file is stored in, the portal converts the content before saving it. Left empty, the extension is read off  the download address, and failing that the stored format is assumed. | 
 **file** | ***os.File** | The edited content, sent as the `File` part of a `multipart/form-data` body. When the part is missing the raw  request body is saved as the content instead, so an empty body empties the file. | 
 **forcesave** | **bool** | Records the write as an editor autosave: the file keeps its running editing session and the previous autosave  revision is overwritten. Left off, the write closes the solo editing session, is refused while somebody else  has the file open, and adds a version to the history. | 

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


## SaveEditingFileFromFormThirdParty

> ThirdPartyFileWrapper SaveEditingFileFromFormThirdParty(ctx, fileId).DownloadUri(downloadUri).FileExtension(fileExtension).File(file).Forcesave(forcesave).Execute()

Save edited file content (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/save-editing-file-from-form-third-party/).

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
	fileId := "1" // string | The file whose content is replaced. The submitted content is written onto this file, so it has to be the file  the editing session was opened on rather than a copy of it.
	downloadUri := "https://example.com/file.txt" // string | An address the document service saved the document at. This operation does not fetch the content from it - the  content always comes from the request body - and reads it only for the extension, when no file extension is  given. (optional)
	fileExtension := "fileExtension_example" // string | The format the submitted content is in, with the leading dot, as in `.docx`. When it differs from the format  the file is stored in, the portal converts the content before saving it. Left empty, the extension is read off  the download address, and failing that the stored format is assumed. (optional)
	file := os.NewFile(1234, "some_file") // *os.File | The edited content, sent as the `File` part of a `multipart/form-data` body. When the part is missing the raw  request body is saved as the content instead, so an empty body empties the file. (optional)
	forcesave := true // bool | Records the write as an editor autosave: the file keeps its running editing session and the previous autosave  revision is overwritten. Left off, the write closes the solo editing session, is refused while somebody else  has the file open, and adds a version to the history. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.SaveEditingFileFromFormThirdParty(context.Background(), fileId).DownloadUri(downloadUri).FileExtension(fileExtension).File(file).Forcesave(forcesave).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.SaveEditingFileFromFormThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveEditingFileFromFormThirdParty`: ThirdPartyFileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.SaveEditingFileFromFormThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file whose content is replaced. The submitted content is written onto this file, so it has to be the file  the editing session was opened on rather than a copy of it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveEditingFileFromFormThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **downloadUri** | **string** | An address the document service saved the document at. This operation does not fetch the content from it - the  content always comes from the request body - and reads it only for the extension, when no file extension is  given. | 
 **fileExtension** | **string** | The format the submitted content is in, with the leading dot, as in `.docx`. When it differs from the format  the file is stored in, the portal converts the content before saving it. Left empty, the extension is read off  the download address, and failing that the stored format is assumed. | 
 **file** | ***os.File** | The edited content, sent as the `File` part of a `multipart/form-data` body. When the part is missing the raw  request body is saved as the content instead, so an empty body empties the file. | 
 **forcesave** | **bool** | Records the write as an editor autosave: the file keeps its running editing session and the previous autosave  revision is overwritten. Left off, the write closes the solo editing session, is refused while somebody else  has the file open, and adds a version to the history. | 

### Return type

[**ThirdPartyFileWrapper**](ThirdPartyFileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveFileAsPdf

> FileWrapper SaveFileAsPdf(ctx, id).SaveAsPdf(saveAsPdf).Execute()

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
	id := int32(1) // int32 | The file to convert; it is left untouched.
	saveAsPdf := *openapiclient.NewSaveAsPdf(int32(1), "My Document") // SaveAsPdf | The destination folder and the name of the PDF.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.SaveFileAsPdf(context.Background(), id).SaveAsPdf(saveAsPdf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.SaveFileAsPdf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveFileAsPdf`: FileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.SaveFileAsPdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The file to convert; it is left untouched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveFileAsPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **saveAsPdf** | [**SaveAsPdf**](SaveAsPdf.md) | The destination folder and the name of the PDF. | 

### Return type

[**FileWrapper**](FileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveFileAsPdfThirdParty

> ThirdPartyFileWrapper SaveFileAsPdfThirdParty(ctx, id).ThirdPartySaveAsPdf(thirdPartySaveAsPdf).Execute()

Save a file as PDF (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/save-file-as-pdf-third-party/).

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
	id := "1" // string | The file to convert; it is left untouched.
	thirdPartySaveAsPdf := *openapiclient.NewThirdPartySaveAsPdf("1", "My Document") // ThirdPartySaveAsPdf | The destination folder and the name of the PDF.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.SaveFileAsPdfThirdParty(context.Background(), id).ThirdPartySaveAsPdf(thirdPartySaveAsPdf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.SaveFileAsPdfThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveFileAsPdfThirdParty`: ThirdPartyFileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.SaveFileAsPdfThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The file to convert; it is left untouched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveFileAsPdfThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **thirdPartySaveAsPdf** | [**ThirdPartySaveAsPdf**](ThirdPartySaveAsPdf.md) | The destination folder and the name of the PDF. | 

### Return type

[**ThirdPartyFileWrapper**](ThirdPartyFileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveFormRoleMapping

> SaveFormRoleMapping(ctx, fileId).SaveFormRoleMappingDto(saveFormRoleMappingDto).Execute()

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
	fileId := "fileId_example" // string | The form the role mapping belongs to. Send the same value as the `formId` of the request body, which is the one the handler reads.
	saveFormRoleMappingDto := *openapiclient.NewSaveFormRoleMappingDto(int32(1), []openapiclient.FormRole{*openapiclient.NewFormRole()}) // SaveFormRoleMappingDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesFilesAPI.SaveFormRoleMapping(context.Background(), fileId).SaveFormRoleMappingDto(saveFormRoleMappingDto).Execute()
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
**fileId** | **string** | The form the role mapping belongs to. Send the same value as the `formId` of the request body, which is the one the handler reads. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveFormRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **saveFormRoleMappingDto** | [**SaveFormRoleMappingDto**](SaveFormRoleMappingDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCustomFilterTag

> FileWrapper SetCustomFilterTag(ctx, fileId).CustomFilterParameters(customFilterParameters).Execute()

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
	fileId := int32(1) // int32 | The spreadsheet whose Custom Filter mode is switched.
	customFilterParameters := *openapiclient.NewCustomFilterParameters() // CustomFilterParameters | The Custom Filter state to reach.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.SetCustomFilterTag(context.Background(), fileId).CustomFilterParameters(customFilterParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.SetCustomFilterTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetCustomFilterTag`: FileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.SetCustomFilterTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The spreadsheet whose Custom Filter mode is switched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCustomFilterTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customFilterParameters** | [**CustomFilterParameters**](CustomFilterParameters.md) | The Custom Filter state to reach. | 

### Return type

[**FileWrapper**](FileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCustomFilterTagThirdParty

> ThirdPartyFileWrapper SetCustomFilterTagThirdParty(ctx, fileId).CustomFilterParameters(customFilterParameters).Execute()

Set the Custom Filter editing mode (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-custom-filter-tag-third-party/).

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
	fileId := "1" // string | The spreadsheet whose Custom Filter mode is switched.
	customFilterParameters := *openapiclient.NewCustomFilterParameters() // CustomFilterParameters | The Custom Filter state to reach.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.SetCustomFilterTagThirdParty(context.Background(), fileId).CustomFilterParameters(customFilterParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.SetCustomFilterTagThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetCustomFilterTagThirdParty`: ThirdPartyFileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.SetCustomFilterTagThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The spreadsheet whose Custom Filter mode is switched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCustomFilterTagThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customFilterParameters** | [**CustomFilterParameters**](CustomFilterParameters.md) | The Custom Filter state to reach. | 

### Return type

[**ThirdPartyFileWrapper**](ThirdPartyFileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetEncryptionInfo

> SetEncryptionInfo(ctx, fileId).AccessRequestKeyDto(accessRequestKeyDto).Execute()

Set file encryption information



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-encryption-info/).

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
	fileId := int32(12345) // int32 | The file the keys are issued for; it has to lie in a private room.
	accessRequestKeyDto := []openapiclient.AccessRequestKeyDto{*openapiclient.NewAccessRequestKeyDto()} // []AccessRequestKeyDto | One key per account that is to open the file. The keys of the accounts named here are replaced and the keys of  everybody else are left as they are, so sending no entry for a person does not revoke that person's key. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesFilesAPI.SetEncryptionInfo(context.Background(), fileId).AccessRequestKeyDto(accessRequestKeyDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.SetEncryptionInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file the keys are issued for; it has to lie in a private room. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetEncryptionInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessRequestKeyDto** | [**[]AccessRequestKeyDto**](AccessRequestKeyDto.md) | One key per account that is to open the file. The keys of the accounts named here are replaced and the keys of  everybody else are left as they are, so sending no entry for a person does not revoke that person's key. | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetEncryptionInfoThirdParty

> SetEncryptionInfoThirdParty(ctx, fileId).AccessRequestKeyDto(accessRequestKeyDto).Execute()

Set file encryption information (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-encryption-info-third-party/).

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
	fileId := "12345" // string | The file the keys are issued for; it has to lie in a private room.
	accessRequestKeyDto := []openapiclient.AccessRequestKeyDto{*openapiclient.NewAccessRequestKeyDto()} // []AccessRequestKeyDto | One key per account that is to open the file. The keys of the accounts named here are replaced and the keys of  everybody else are left as they are, so sending no entry for a person does not revoke that person's key. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesFilesAPI.SetEncryptionInfoThirdParty(context.Background(), fileId).AccessRequestKeyDto(accessRequestKeyDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.SetEncryptionInfoThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file the keys are issued for; it has to lie in a private room. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetEncryptionInfoThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessRequestKeyDto** | [**[]AccessRequestKeyDto**](AccessRequestKeyDto.md) | One key per account that is to open the file. The keys of the accounts named here are replaced and the keys of  everybody else are left as they are, so sending no entry for a person does not revoke that person's key. | 

### Return type

 (empty response body)

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

Set a file external link



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
	id := int32(1) // int32 | The file the link points at.
	fileLinkRequest := *openapiclient.NewFileLinkRequest() // FileLinkRequest | The settings of the link. They are applied in full, so a field left out is reset rather than kept.

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
**id** | **int32** | The file the link points at. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFileExternalLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileLinkRequest** | [**FileLinkRequest**](FileLinkRequest.md) | The settings of the link. They are applied in full, so a field left out is reset rather than kept. | 

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


## SetFileExternalLinkThirdParty

> FileShareWrapper SetFileExternalLinkThirdParty(ctx, id).FileLinkRequest(fileLinkRequest).Execute()

Set a file external link (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-file-external-link-third-party/).

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
	id := "1" // string | The file the link points at.
	fileLinkRequest := *openapiclient.NewFileLinkRequest() // FileLinkRequest | The settings of the link. They are applied in full, so a field left out is reset rather than kept.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.SetFileExternalLinkThirdParty(context.Background(), id).FileLinkRequest(fileLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.SetFileExternalLinkThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFileExternalLinkThirdParty`: FileShareWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.SetFileExternalLinkThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The file the link points at. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFileExternalLinkThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileLinkRequest** | [**FileLinkRequest**](FileLinkRequest.md) | The settings of the link. They are applied in full, so a field left out is reset rather than kept. | 

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

> FileWrapper SetFileOrder(ctx, fileId).OrderRequestDto(orderRequestDto).Execute()

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
	fileId := int32(1) // int32 | The file to move.
	orderRequestDto := *openapiclient.NewOrderRequestDto() // OrderRequestDto | The position the file is to take. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.SetFileOrder(context.Background(), fileId).OrderRequestDto(orderRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.SetFileOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFileOrder`: FileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.SetFileOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file to move. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFileOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderRequestDto** | [**OrderRequestDto**](OrderRequestDto.md) | The position the file is to take. | 

### Return type

[**FileWrapper**](FileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFileOrderThirdParty

> ThirdPartyFileWrapper SetFileOrderThirdParty(ctx, fileId).OrderRequestDto(orderRequestDto).Execute()

Set file order (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-file-order-third-party/).

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
	fileId := "1" // string | The file to move.
	orderRequestDto := *openapiclient.NewOrderRequestDto() // OrderRequestDto | The position the file is to take. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.SetFileOrderThirdParty(context.Background(), fileId).OrderRequestDto(orderRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.SetFileOrderThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFileOrderThirdParty`: ThirdPartyFileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.SetFileOrderThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file to move. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFileOrderThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderRequestDto** | [**OrderRequestDto**](OrderRequestDto.md) | The position the file is to take. | 

### Return type

[**ThirdPartyFileWrapper**](ThirdPartyFileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFilesOrder

> FileEntryArrayWrapper SetFilesOrder(ctx).OrdersRequestDto(ordersRequestDto).Execute()

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
	ordersRequestDto := *openapiclient.NewOrdersRequestDto([]openapiclient.OrdersItemRequestDto{*openapiclient.NewOrdersItemRequestDto(int32(1), openapiclient.FileEntryType(1), int32(1))}) // OrdersRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.SetFilesOrder(context.Background()).OrdersRequestDto(ordersRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.SetFilesOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFilesOrder`: FileEntryArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.SetFilesOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetFilesOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ordersRequestDto** | [**OrdersRequestDto**](OrdersRequestDto.md) |  | 

### Return type

[**FileEntryArrayWrapper**](FileEntryArrayWrapper.md)

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

Open an editing session



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
	fileId := int32(1) // int32 | The file to open the editing session on. The caller needs edit access to it.
	startEdit := *openapiclient.NewStartEdit() // StartEdit | The session options. The body is required even when it only carries the default, so send an empty object to  open an ordinary co-editing session.

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
**fileId** | **int32** | The file to open the editing session on. The caller needs edit access to it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartEditFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startEdit** | [**StartEdit**](StartEdit.md) | The session options. The body is required even when it only carries the default, so send an empty object to  open an ordinary co-editing session. | 

### Return type

[**StringWrapper**](StringWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartEditFileThirdParty

> StringWrapper StartEditFileThirdParty(ctx, fileId).StartEdit(startEdit).Execute()

Open an editing session (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/start-edit-file-third-party/).

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
	fileId := "1" // string | The file to open the editing session on. The caller needs edit access to it.
	startEdit := *openapiclient.NewStartEdit() // StartEdit | The session options. The body is required even when it only carries the default, so send an empty object to  open an ordinary co-editing session.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.StartEditFileThirdParty(context.Background(), fileId).StartEdit(startEdit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.StartEditFileThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartEditFileThirdParty`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.StartEditFileThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file to open the editing session on. The caller needs edit access to it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartEditFileThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startEdit** | [**StartEdit**](StartEdit.md) | The session options. The body is required even when it only carries the default, so send an empty object to  open an ordinary co-editing session. | 

### Return type

[**StringWrapper**](StringWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartFillingFile

> FileWrapper StartFillingFile(ctx, fileId).Execute()

Start filling a form



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
	fileId := int32(1) // int32 | The PDF form to open for filling. It has to be the form as it lies in the form-filling room itself, not a copy  kept elsewhere and not a submitted result.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.StartFillingFile(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.StartFillingFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartFillingFile`: FileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.StartFillingFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The PDF form to open for filling. It has to be the form as it lies in the form-filling room itself, not a copy  kept elsewhere and not a submitted result. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartFillingFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileWrapper**](FileWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartFillingFileThirdParty

> ThirdPartyFileWrapper StartFillingFileThirdParty(ctx, fileId).Execute()

Start filling a form (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/start-filling-file-third-party/).

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
	fileId := "1" // string | The PDF form to open for filling. It has to be the form as it lies in the form-filling room itself, not a copy  kept elsewhere and not a submitted result.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.StartFillingFileThirdParty(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.StartFillingFileThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartFillingFileThirdParty`: ThirdPartyFileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.StartFillingFileThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The PDF form to open for filling. It has to be the form as it lies in the form-filling room itself, not a copy  kept elsewhere and not a submitted result. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartFillingFileThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ThirdPartyFileWrapper**](ThirdPartyFileWrapper.md)

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

Set the file favorite status



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
	fileId := int32(10) // int32 | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.
	favorite := true // bool | Which state to put the mark in: `true` adds the file to the favorites of the calling account, `false` removes  it from them. Leaving the field out of the request removes the mark rather than setting it. (optional)

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
**fileId** | **int32** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleFileFavoriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **favorite** | **bool** | Which state to put the mark in: `true` adds the file to the favorites of the calling account, `false` removes  it from them. Leaving the field out of the request removes the mark rather than setting it. | 

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


## ToggleFileFavoriteThirdParty

> BooleanWrapper ToggleFileFavoriteThirdParty(ctx, fileId).Favorite(favorite).Execute()

Set the file favorite status (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/toggle-file-favorite-third-party/).

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
	fileId := "10" // string | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.
	favorite := true // bool | Which state to put the mark in: `true` adds the file to the favorites of the calling account, `false` removes  it from them. Leaving the field out of the request removes the mark rather than setting it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.ToggleFileFavoriteThirdParty(context.Background(), fileId).Favorite(favorite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.ToggleFileFavoriteThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToggleFileFavoriteThirdParty`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.ToggleFileFavoriteThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleFileFavoriteThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **favorite** | **bool** | Which state to put the mark in: `true` adds the file to the favorites of the calling account, `false` removes  it from them. Leaving the field out of the request removes the mark rather than setting it. | 

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

> ItemKeyValuePairBooleanStringWrapper TrackEditFile(ctx, fileId).TabId(tabId).DocKeyForTrack(docKeyForTrack).IsFinish(isFinish).Execute()

Track an editing session



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
	fileId := int32(1) // int32 | The file whose editing session is being tracked.
	tabId := "00000000-0000-0000-0000-000000000000" // string | The client tab that holds the session, a value the client makes up once and repeats on every call about that  tab. Two tabs sending different values are tracked as two sessions on the same file, while the all-zero value  belongs to a session claimed for a single editor. (optional)
	docKeyForTrack := "abc123" // string | The document key of the revision being edited, as `POST api/2.0/files/file/{fileId}/startedit` returned it. It  is checked against the file's current key on every call, so a key left over from an older revision is refused. (optional)
	isFinish := true // bool | Ends the session for this tab and tells the other clients that editing has stopped. Left off, the session is  refreshed and the file stays marked as being edited. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.TrackEditFile(context.Background(), fileId).TabId(tabId).DocKeyForTrack(docKeyForTrack).IsFinish(isFinish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.TrackEditFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackEditFile`: ItemKeyValuePairBooleanStringWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.TrackEditFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file whose editing session is being tracked. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrackEditFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tabId** | **string** | The client tab that holds the session, a value the client makes up once and repeats on every call about that  tab. Two tabs sending different values are tracked as two sessions on the same file, while the all-zero value  belongs to a session claimed for a single editor. | 
 **docKeyForTrack** | **string** | The document key of the revision being edited, as `POST api/2.0/files/file/{fileId}/startedit` returned it. It  is checked against the file's current key on every call, so a key left over from an older revision is refused. | 
 **isFinish** | **bool** | Ends the session for this tab and tells the other clients that editing has stopped. Left off, the session is  refreshed and the file stays marked as being edited. | 

### Return type

[**ItemKeyValuePairBooleanStringWrapper**](ItemKeyValuePairBooleanStringWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackEditFileThirdParty

> ItemKeyValuePairBooleanStringWrapper TrackEditFileThirdParty(ctx, fileId).TabId(tabId).DocKeyForTrack(docKeyForTrack).IsFinish(isFinish).Execute()

Track an editing session (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/track-edit-file-third-party/).

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
	fileId := "1" // string | The file whose editing session is being tracked.
	tabId := "00000000-0000-0000-0000-000000000000" // string | The client tab that holds the session, a value the client makes up once and repeats on every call about that  tab. Two tabs sending different values are tracked as two sessions on the same file, while the all-zero value  belongs to a session claimed for a single editor. (optional)
	docKeyForTrack := "abc123" // string | The document key of the revision being edited, as `POST api/2.0/files/file/{fileId}/startedit` returned it. It  is checked against the file's current key on every call, so a key left over from an older revision is refused. (optional)
	isFinish := true // bool | Ends the session for this tab and tells the other clients that editing has stopped. Left off, the session is  refreshed and the file stays marked as being edited. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.TrackEditFileThirdParty(context.Background(), fileId).TabId(tabId).DocKeyForTrack(docKeyForTrack).IsFinish(isFinish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.TrackEditFileThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackEditFileThirdParty`: ItemKeyValuePairBooleanStringWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.TrackEditFileThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file whose editing session is being tracked. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrackEditFileThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tabId** | **string** | The client tab that holds the session, a value the client makes up once and repeats on every call about that  tab. Two tabs sending different values are tracked as two sessions on the same file, while the all-zero value  belongs to a session claimed for a single editor. | 
 **docKeyForTrack** | **string** | The document key of the revision being edited, as `POST api/2.0/files/file/{fileId}/startedit` returned it. It  is checked against the file's current key on every call, so a key left over from an older revision is refused. | 
 **isFinish** | **bool** | Ends the session for this tab and tells the other clients that editing has stopped. Left off, the session is  refreshed and the file stays marked as being edited. | 

### Return type

[**ItemKeyValuePairBooleanStringWrapper**](ItemKeyValuePairBooleanStringWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFile

> FileWrapper UpdateFile(ctx, fileId).UpdateFile(updateFile).Execute()

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
	fileId := int32(1) // int32 | The file to update.
	updateFile := *openapiclient.NewUpdateFile() // UpdateFile | The new title and the version to restore.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.UpdateFile(context.Background(), fileId).UpdateFile(updateFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.UpdateFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFile`: FileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.UpdateFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFile** | [**UpdateFile**](UpdateFile.md) | The new title and the version to restore. | 

### Return type

[**FileWrapper**](FileWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFileThirdParty

> ThirdPartyFileWrapper UpdateFileThirdParty(ctx, fileId).UpdateFile(updateFile).Execute()

Update a file (third-party storage)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-file-third-party/).

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
	fileId := "1" // string | The file to update.
	updateFile := *openapiclient.NewUpdateFile() // UpdateFile | The new title and the version to restore.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesFilesAPI.UpdateFileThirdParty(context.Background(), fileId).UpdateFile(updateFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesFilesAPI.UpdateFileThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFileThirdParty`: ThirdPartyFileWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesFilesAPI.UpdateFileThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFileThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFile** | [**UpdateFile**](UpdateFile.md) | The new title and the version to restore. | 

### Return type

[**ThirdPartyFileWrapper**](ThirdPartyFileWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

