# \FilesSettingsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeAccessToThirdparty**](FilesSettingsAPI.md#ChangeAccessToThirdparty) | **Put** /api/2.0/files/thirdparty | Change the third-party settings access
[**ChangeAutomaticallyCleanUp**](FilesSettingsAPI.md#ChangeAutomaticallyCleanUp) | **Put** /api/2.0/files/settings/autocleanup | Update the trash bin auto-clearing setting
[**ChangeDefaultAccessRights**](FilesSettingsAPI.md#ChangeDefaultAccessRights) | **Put** /api/2.0/files/settings/dafaultaccessrights | Change the default access rights
[**ChangeDeleteConfirm**](FilesSettingsAPI.md#ChangeDeleteConfirm) | **Put** /api/2.0/files/changedeleteconfrim | Confirm the file deletion
[**ChangeDownloadZipFromBody**](FilesSettingsAPI.md#ChangeDownloadZipFromBody) | **Put** /api/2.0/files/settings/downloadtargz | Change the archive format (using body parameters)
[**CheckDocServiceUrl**](FilesSettingsAPI.md#CheckDocServiceUrl) | **Put** /api/2.0/files/docservice | Check the document service URL
[**DisplayFileExtension**](FilesSettingsAPI.md#DisplayFileExtension) | **Put** /api/2.0/files/displayfileextension | Display a file extension
[**DisplayRecent**](FilesSettingsAPI.md#DisplayRecent) | **Put** /api/2.0/files/displayrecent | Display the Recent folder
[**ExternalShare**](FilesSettingsAPI.md#ExternalShare) | **Put** /api/2.0/files/settings/external | Change the external sharing ability
[**ExternalShareSocialMedia**](FilesSettingsAPI.md#ExternalShareSocialMedia) | **Put** /api/2.0/files/settings/externalsocialmedia | Change the external sharing ability on social networks
[**Forcesave**](FilesSettingsAPI.md#Forcesave) | **Put** /api/2.0/files/forcesave | Change the forcesaving ability
[**GetAutomaticallyCleanUp**](FilesSettingsAPI.md#GetAutomaticallyCleanUp) | **Get** /api/2.0/files/settings/autocleanup | Get the trash bin auto-clearing setting
[**GetDefaultTemplates**](FilesSettingsAPI.md#GetDefaultTemplates) | **Get** /api/2.0/files/settings/defaulttemplate | Get the default template setting
[**GetDocServiceUrl**](FilesSettingsAPI.md#GetDocServiceUrl) | **Get** /api/2.0/files/docservice | Get the document service URL
[**GetFilesModule**](FilesSettingsAPI.md#GetFilesModule) | **Get** /api/2.0/files/info | Get the Documents information
[**GetFilesSettings**](FilesSettingsAPI.md#GetFilesSettings) | **Get** /api/2.0/files/settings | Get file settings
[**HideConfirmCancelOperation**](FilesSettingsAPI.md#HideConfirmCancelOperation) | **Put** /api/2.0/files/hideconfirmcanceloperation | Hide confirmation dialog when canceling operations
[**HideConfirmConvert**](FilesSettingsAPI.md#HideConfirmConvert) | **Put** /api/2.0/files/hideconfirmconvert | Hide the confirmation dialog when converting
[**HideConfirmRoomLifetime**](FilesSettingsAPI.md#HideConfirmRoomLifetime) | **Put** /api/2.0/files/hideconfirmroomlifetime | Hide confirmation dialog when changing room lifetime settings
[**IsAvailablePrivacyRoomSettings**](FilesSettingsAPI.md#IsAvailablePrivacyRoomSettings) | **Get** /api/2.0/files/@privacy/available | Check the Private Room availability
[**KeepNewFileName**](FilesSettingsAPI.md#KeepNewFileName) | **Put** /api/2.0/files/keepnewfilename | Ask a new file name
[**ResetDefaultTemplate**](FilesSettingsAPI.md#ResetDefaultTemplate) | **Delete** /api/2.0/files/settings/defaulttemplate | Reset the default template setting
[**SetDefaultTemplate**](FilesSettingsAPI.md#SetDefaultTemplate) | **Put** /api/2.0/files/settings/defaulttemplate | Change the default template setting
[**SetOpenEditorInSameTab**](FilesSettingsAPI.md#SetOpenEditorInSameTab) | **Put** /api/2.0/files/settings/openeditorinsametab | Open document in the same browser tab
[**SetOrganizeRoomsGrouping**](FilesSettingsAPI.md#SetOrganizeRoomsGrouping) | **Put** /api/2.0/files/settings/organizegrouping | Organize rooms grouping
[**StoreForcesave**](FilesSettingsAPI.md#StoreForcesave) | **Put** /api/2.0/files/storeforcesave | Change the ability to store the forcesaved files
[**StoreOriginal**](FilesSettingsAPI.md#StoreOriginal) | **Put** /api/2.0/files/storeoriginal | Change the ability to upload original formats
[**UpdateFileIfExist**](FilesSettingsAPI.md#UpdateFileIfExist) | **Put** /api/2.0/files/updateifexist | Update a file version if it exists
[**UploadDefaultTemplate**](FilesSettingsAPI.md#UploadDefaultTemplate) | **Post** /api/2.0/files/settings/defaulttemplate | Upload a file as the default template setting



## ChangeAccessToThirdparty

> BooleanWrapper ChangeAccessToThirdparty(ctx).SettingsRequestDto(settingsRequestDto).Execute()

Change the third-party settings access



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/change-access-to-thirdparty/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	settingsRequestDto := *openapiclient.NewSettingsRequestDto() // SettingsRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.ChangeAccessToThirdparty(context.Background()).SettingsRequestDto(settingsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.ChangeAccessToThirdparty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeAccessToThirdparty`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.ChangeAccessToThirdparty`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeAccessToThirdpartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settingsRequestDto** | [**SettingsRequestDto**](SettingsRequestDto.md) |  | 

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


## ChangeAutomaticallyCleanUp

> AutoCleanUpDataWrapper ChangeAutomaticallyCleanUp(ctx).AutoCleanupRequestDto(autoCleanupRequestDto).Execute()

Update the trash bin auto-clearing setting



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/change-automatically-clean-up/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	autoCleanupRequestDto := *openapiclient.NewAutoCleanupRequestDto() // AutoCleanupRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.ChangeAutomaticallyCleanUp(context.Background()).AutoCleanupRequestDto(autoCleanupRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.ChangeAutomaticallyCleanUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeAutomaticallyCleanUp`: AutoCleanUpDataWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.ChangeAutomaticallyCleanUp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeAutomaticallyCleanUpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoCleanupRequestDto** | [**AutoCleanupRequestDto**](AutoCleanupRequestDto.md) |  | 

### Return type

[**AutoCleanUpDataWrapper**](AutoCleanUpDataWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeDefaultAccessRights

> FileShareArrayWrapper ChangeDefaultAccessRights(ctx).RequestBody(requestBody).Execute()

Change the default access rights



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/change-default-access-rights/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	requestBody := []int32{int32(0)} // []int32 | Sharing rights (None, ReadWrite, Read, Restrict, Varies, Review, Comment, FillForms, CustomFilter, RoomAdmin, Editing, Collaborator). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.ChangeDefaultAccessRights(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.ChangeDefaultAccessRights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeDefaultAccessRights`: FileShareArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.ChangeDefaultAccessRights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeDefaultAccessRightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]int32** | Sharing rights (None, ReadWrite, Read, Restrict, Varies, Review, Comment, FillForms, CustomFilter, RoomAdmin, Editing, Collaborator). | 

### Return type

[**FileShareArrayWrapper**](FileShareArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeDeleteConfirm

> BooleanWrapper ChangeDeleteConfirm(ctx).SettingsRequestDto(settingsRequestDto).Execute()

Confirm the file deletion



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/change-delete-confirm/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	settingsRequestDto := *openapiclient.NewSettingsRequestDto() // SettingsRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.ChangeDeleteConfirm(context.Background()).SettingsRequestDto(settingsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.ChangeDeleteConfirm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeDeleteConfirm`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.ChangeDeleteConfirm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeDeleteConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settingsRequestDto** | [**SettingsRequestDto**](SettingsRequestDto.md) |  | 

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


## ChangeDownloadZipFromBody

> ICompressWrapper ChangeDownloadZipFromBody(ctx).DisplayRequestDto(displayRequestDto).Execute()

Change the archive format (using body parameters)



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/change-download-zip-from-body/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	displayRequestDto := *openapiclient.NewDisplayRequestDto() // DisplayRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.ChangeDownloadZipFromBody(context.Background()).DisplayRequestDto(displayRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.ChangeDownloadZipFromBody``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeDownloadZipFromBody`: ICompressWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.ChangeDownloadZipFromBody`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeDownloadZipFromBodyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **displayRequestDto** | [**DisplayRequestDto**](DisplayRequestDto.md) |  | 

### Return type

[**ICompressWrapper**](ICompressWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckDocServiceUrl

> DocServiceUrlWrapper CheckDocServiceUrl(ctx).CheckDocServiceUrlRequestDto(checkDocServiceUrlRequestDto).Execute()

Check the document service URL



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/check-doc-service-url/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	checkDocServiceUrlRequestDto := *openapiclient.NewCheckDocServiceUrlRequestDto("https://documentserver.example.com") // CheckDocServiceUrlRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.CheckDocServiceUrl(context.Background()).CheckDocServiceUrlRequestDto(checkDocServiceUrlRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.CheckDocServiceUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckDocServiceUrl`: DocServiceUrlWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.CheckDocServiceUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckDocServiceUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkDocServiceUrlRequestDto** | [**CheckDocServiceUrlRequestDto**](CheckDocServiceUrlRequestDto.md) |  | 

### Return type

[**DocServiceUrlWrapper**](DocServiceUrlWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisplayFileExtension

> BooleanWrapper DisplayFileExtension(ctx).SettingsRequestDto(settingsRequestDto).Execute()

Display a file extension



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/display-file-extension/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	settingsRequestDto := *openapiclient.NewSettingsRequestDto() // SettingsRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.DisplayFileExtension(context.Background()).SettingsRequestDto(settingsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.DisplayFileExtension``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisplayFileExtension`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.DisplayFileExtension`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisplayFileExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settingsRequestDto** | [**SettingsRequestDto**](SettingsRequestDto.md) |  | 

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


## DisplayRecent

> BooleanWrapper DisplayRecent(ctx).DisplayRequestDto(displayRequestDto).Execute()

Display the Recent folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/display-recent/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	displayRequestDto := *openapiclient.NewDisplayRequestDto() // DisplayRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.DisplayRecent(context.Background()).DisplayRequestDto(displayRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.DisplayRecent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisplayRecent`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.DisplayRecent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisplayRecentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **displayRequestDto** | [**DisplayRequestDto**](DisplayRequestDto.md) |  | 

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


## ExternalShare

> BooleanWrapper ExternalShare(ctx).DisplayRequestDto(displayRequestDto).Execute()

Change the external sharing ability



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/external-share/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	displayRequestDto := *openapiclient.NewDisplayRequestDto() // DisplayRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.ExternalShare(context.Background()).DisplayRequestDto(displayRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.ExternalShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalShare`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.ExternalShare`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExternalShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **displayRequestDto** | [**DisplayRequestDto**](DisplayRequestDto.md) |  | 

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


## ExternalShareSocialMedia

> BooleanWrapper ExternalShareSocialMedia(ctx).DisplayRequestDto(displayRequestDto).Execute()

Change the external sharing ability on social networks



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/external-share-social-media/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	displayRequestDto := *openapiclient.NewDisplayRequestDto() // DisplayRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.ExternalShareSocialMedia(context.Background()).DisplayRequestDto(displayRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.ExternalShareSocialMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalShareSocialMedia`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.ExternalShareSocialMedia`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExternalShareSocialMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **displayRequestDto** | [**DisplayRequestDto**](DisplayRequestDto.md) |  | 

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


## Forcesave

> BooleanWrapper Forcesave(ctx).Execute()

Change the forcesaving ability



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/forcesave/).

### Example

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
	resp, r, err := apiClient.FilesSettingsAPI.Forcesave(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.Forcesave``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Forcesave`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.Forcesave`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiForcesaveRequest struct via the builder pattern


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


## GetAutomaticallyCleanUp

> AutoCleanUpDataWrapper GetAutomaticallyCleanUp(ctx).Execute()

Get the trash bin auto-clearing setting



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-automatically-clean-up/).

### Example

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
	resp, r, err := apiClient.FilesSettingsAPI.GetAutomaticallyCleanUp(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.GetAutomaticallyCleanUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomaticallyCleanUp`: AutoCleanUpDataWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.GetAutomaticallyCleanUp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomaticallyCleanUpRequest struct via the builder pattern


### Return type

[**AutoCleanUpDataWrapper**](AutoCleanUpDataWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultTemplates

> DefaultTemplateSettingsWrapper GetDefaultTemplates(ctx).Execute()

Get the default template setting



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-default-templates/).

### Example

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
	resp, r, err := apiClient.FilesSettingsAPI.GetDefaultTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.GetDefaultTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultTemplates`: DefaultTemplateSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.GetDefaultTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultTemplatesRequest struct via the builder pattern


### Return type

[**DefaultTemplateSettingsWrapper**](DefaultTemplateSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocServiceUrl

> DocServiceUrlWrapper GetDocServiceUrl(ctx).Version(version).Execute()

Get the document service URL



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-doc-service-url/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	version := true // bool | Specifies whether to return the editor version or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.GetDocServiceUrl(context.Background()).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.GetDocServiceUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDocServiceUrl`: DocServiceUrlWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.GetDocServiceUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDocServiceUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **bool** | Specifies whether to return the editor version or not. | 

### Return type

[**DocServiceUrlWrapper**](DocServiceUrlWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilesModule

> ModuleWrapper GetFilesModule(ctx).Execute()

Get the Documents information



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-files-module/).

### Example

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
	resp, r, err := apiClient.FilesSettingsAPI.GetFilesModule(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.GetFilesModule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilesModule`: ModuleWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.GetFilesModule`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesModuleRequest struct via the builder pattern


### Return type

[**ModuleWrapper**](ModuleWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilesSettings

> FilesSettingsWrapper GetFilesSettings(ctx).Execute()

Get file settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-files-settings/).

### Example

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
	resp, r, err := apiClient.FilesSettingsAPI.GetFilesSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.GetFilesSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilesSettings`: FilesSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.GetFilesSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesSettingsRequest struct via the builder pattern


### Return type

[**FilesSettingsWrapper**](FilesSettingsWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HideConfirmCancelOperation

> BooleanWrapper HideConfirmCancelOperation(ctx).SettingsRequestDto(settingsRequestDto).Execute()

Hide confirmation dialog when canceling operations



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/hide-confirm-cancel-operation/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	settingsRequestDto := *openapiclient.NewSettingsRequestDto() // SettingsRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.HideConfirmCancelOperation(context.Background()).SettingsRequestDto(settingsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.HideConfirmCancelOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HideConfirmCancelOperation`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.HideConfirmCancelOperation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHideConfirmCancelOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settingsRequestDto** | [**SettingsRequestDto**](SettingsRequestDto.md) |  | 

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


## HideConfirmConvert

> BooleanWrapper HideConfirmConvert(ctx).HideConfirmConvertRequestDto(hideConfirmConvertRequestDto).Execute()

Hide the confirmation dialog when converting



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/hide-confirm-convert/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	hideConfirmConvertRequestDto := *openapiclient.NewHideConfirmConvertRequestDto() // HideConfirmConvertRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.HideConfirmConvert(context.Background()).HideConfirmConvertRequestDto(hideConfirmConvertRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.HideConfirmConvert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HideConfirmConvert`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.HideConfirmConvert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHideConfirmConvertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hideConfirmConvertRequestDto** | [**HideConfirmConvertRequestDto**](HideConfirmConvertRequestDto.md) |  | 

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


## HideConfirmRoomLifetime

> BooleanWrapper HideConfirmRoomLifetime(ctx).SettingsRequestDto(settingsRequestDto).Execute()

Hide confirmation dialog when changing room lifetime settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/hide-confirm-room-lifetime/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	settingsRequestDto := *openapiclient.NewSettingsRequestDto() // SettingsRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.HideConfirmRoomLifetime(context.Background()).SettingsRequestDto(settingsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.HideConfirmRoomLifetime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HideConfirmRoomLifetime`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.HideConfirmRoomLifetime`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHideConfirmRoomLifetimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settingsRequestDto** | [**SettingsRequestDto**](SettingsRequestDto.md) |  | 

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


## IsAvailablePrivacyRoomSettings

> BooleanWrapper IsAvailablePrivacyRoomSettings(ctx).Execute()

Check the Private Room availability



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/is-available-privacy-room-settings/).

### Example

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
	resp, r, err := apiClient.FilesSettingsAPI.IsAvailablePrivacyRoomSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.IsAvailablePrivacyRoomSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IsAvailablePrivacyRoomSettings`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.IsAvailablePrivacyRoomSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIsAvailablePrivacyRoomSettingsRequest struct via the builder pattern


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


## KeepNewFileName

> BooleanWrapper KeepNewFileName(ctx).SettingsRequestDto(settingsRequestDto).Execute()

Ask a new file name



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/keep-new-file-name/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	settingsRequestDto := *openapiclient.NewSettingsRequestDto() // SettingsRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.KeepNewFileName(context.Background()).SettingsRequestDto(settingsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.KeepNewFileName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeepNewFileName`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.KeepNewFileName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKeepNewFileNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settingsRequestDto** | [**SettingsRequestDto**](SettingsRequestDto.md) |  | 

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


## ResetDefaultTemplate

> DefaultTemplateSettingsWrapper ResetDefaultTemplate(ctx).DefaultTemplateSettingsResetRequestDto(defaultTemplateSettingsResetRequestDto).Execute()

Reset the default template setting



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/reset-default-template/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	defaultTemplateSettingsResetRequestDto := *openapiclient.NewDefaultTemplateSettingsResetRequestDto(".docx") // DefaultTemplateSettingsResetRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.ResetDefaultTemplate(context.Background()).DefaultTemplateSettingsResetRequestDto(defaultTemplateSettingsResetRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.ResetDefaultTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetDefaultTemplate`: DefaultTemplateSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.ResetDefaultTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetDefaultTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **defaultTemplateSettingsResetRequestDto** | [**DefaultTemplateSettingsResetRequestDto**](DefaultTemplateSettingsResetRequestDto.md) |  | 

### Return type

[**DefaultTemplateSettingsWrapper**](DefaultTemplateSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultTemplate

> DefaultTemplateSettingsWrapper SetDefaultTemplate(ctx).DefaultTemplateSettingsRequestDto(defaultTemplateSettingsRequestDto).Execute()

Change the default template setting



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-default-template/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	defaultTemplateSettingsRequestDto := *openapiclient.NewDefaultTemplateSettingsRequestDto(openapiclient.DefaultTemplateSettingsRequestDto_selectedFile{Int32: new(int32)}, ".docx") // DefaultTemplateSettingsRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.SetDefaultTemplate(context.Background()).DefaultTemplateSettingsRequestDto(defaultTemplateSettingsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.SetDefaultTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDefaultTemplate`: DefaultTemplateSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.SetDefaultTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **defaultTemplateSettingsRequestDto** | [**DefaultTemplateSettingsRequestDto**](DefaultTemplateSettingsRequestDto.md) |  | 

### Return type

[**DefaultTemplateSettingsWrapper**](DefaultTemplateSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOpenEditorInSameTab

> BooleanWrapper SetOpenEditorInSameTab(ctx).SettingsRequestDto(settingsRequestDto).Execute()

Open document in the same browser tab



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-open-editor-in-same-tab/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	settingsRequestDto := *openapiclient.NewSettingsRequestDto() // SettingsRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.SetOpenEditorInSameTab(context.Background()).SettingsRequestDto(settingsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.SetOpenEditorInSameTab``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOpenEditorInSameTab`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.SetOpenEditorInSameTab`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetOpenEditorInSameTabRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settingsRequestDto** | [**SettingsRequestDto**](SettingsRequestDto.md) |  | 

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


## SetOrganizeRoomsGrouping

> BooleanWrapper SetOrganizeRoomsGrouping(ctx).SettingsRequestDto(settingsRequestDto).Execute()

Organize rooms grouping



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-organize-rooms-grouping/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	settingsRequestDto := *openapiclient.NewSettingsRequestDto() // SettingsRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.SetOrganizeRoomsGrouping(context.Background()).SettingsRequestDto(settingsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.SetOrganizeRoomsGrouping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOrganizeRoomsGrouping`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.SetOrganizeRoomsGrouping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetOrganizeRoomsGroupingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settingsRequestDto** | [**SettingsRequestDto**](SettingsRequestDto.md) |  | 

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


## StoreForcesave

> BooleanWrapper StoreForcesave(ctx).Execute()

Change the ability to store the forcesaved files



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/store-forcesave/).

### Example

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
	resp, r, err := apiClient.FilesSettingsAPI.StoreForcesave(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.StoreForcesave``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreForcesave`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.StoreForcesave`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStoreForcesaveRequest struct via the builder pattern


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


## StoreOriginal

> BooleanWrapper StoreOriginal(ctx).SettingsRequestDto(settingsRequestDto).Execute()

Change the ability to upload original formats



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/store-original/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	settingsRequestDto := *openapiclient.NewSettingsRequestDto() // SettingsRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.StoreOriginal(context.Background()).SettingsRequestDto(settingsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.StoreOriginal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreOriginal`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.StoreOriginal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreOriginalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settingsRequestDto** | [**SettingsRequestDto**](SettingsRequestDto.md) |  | 

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


## UpdateFileIfExist

> BooleanWrapper UpdateFileIfExist(ctx).SettingsRequestDto(settingsRequestDto).Execute()

Update a file version if it exists



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-file-if-exist/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	settingsRequestDto := *openapiclient.NewSettingsRequestDto() // SettingsRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.UpdateFileIfExist(context.Background()).SettingsRequestDto(settingsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.UpdateFileIfExist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFileIfExist`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.UpdateFileIfExist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFileIfExistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settingsRequestDto** | [**SettingsRequestDto**](SettingsRequestDto.md) |  | 

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


## UploadDefaultTemplate

> DefaultTemplateSettingsWrapper UploadDefaultTemplate(ctx).FileExtension(fileExtension).File(file).Execute()

Upload a file as the default template setting



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/upload-default-template/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	fileExtension := ".docx" // string | File extension of a template to replace
	file := os.NewFile(1234, "some_file") // *os.File | File to replace template with

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSettingsAPI.UploadDefaultTemplate(context.Background()).FileExtension(fileExtension).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSettingsAPI.UploadDefaultTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadDefaultTemplate`: DefaultTemplateSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSettingsAPI.UploadDefaultTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadDefaultTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileExtension** | **string** | File extension of a template to replace | 
 **file** | ***os.File** | File to replace template with | 

### Return type

[**DefaultTemplateSettingsWrapper**](DefaultTemplateSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

