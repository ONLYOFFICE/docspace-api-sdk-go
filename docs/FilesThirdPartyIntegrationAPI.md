# \FilesThirdPartyIntegrationAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteThirdParty**](FilesThirdPartyIntegrationAPI.md#DeleteThirdParty) | **Delete** /api/2.0/files/thirdparty/{providerId} | Remove a third-party account
[**GetAllProviders**](FilesThirdPartyIntegrationAPI.md#GetAllProviders) | **Get** /api/2.0/files/thirdparty/providers | Get all third-party providers
[**GetBackupThirdPartyAccount**](FilesThirdPartyIntegrationAPI.md#GetBackupThirdPartyAccount) | **Get** /api/2.0/files/thirdparty/backup | Get the third-party backup folder
[**GetCapabilities**](FilesThirdPartyIntegrationAPI.md#GetCapabilities) | **Get** /api/2.0/files/thirdparty/capabilities | Get third-party provider capabilities
[**GetCommonThirdPartyFolders**](FilesThirdPartyIntegrationAPI.md#GetCommonThirdPartyFolders) | **Get** /api/2.0/files/thirdparty/common | Get common third-party folders
[**GetThirdPartyAccounts**](FilesThirdPartyIntegrationAPI.md#GetThirdPartyAccounts) | **Get** /api/2.0/files/thirdparty | Get the third-party accounts
[**SaveThirdParty**](FilesThirdPartyIntegrationAPI.md#SaveThirdParty) | **Post** /api/2.0/files/thirdparty | Connect a third-party account
[**SaveThirdPartyBackup**](FilesThirdPartyIntegrationAPI.md#SaveThirdPartyBackup) | **Post** /api/2.0/files/thirdparty/backup | Connect the third-party backup storage



## DeleteThirdParty

> StringWrapper DeleteThirdParty(ctx, providerId).Execute()

Remove a third-party account



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-third-party/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	providerId := int32(12) // int32 | The ID of the connected third-party storage account, as `providerId` of `GET api/2.0/files/thirdparty`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesThirdPartyIntegrationAPI.DeleteThirdParty(context.Background(), providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesThirdPartyIntegrationAPI.DeleteThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteThirdParty`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesThirdPartyIntegrationAPI.DeleteThirdParty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **int32** | The ID of the connected third-party storage account, as `providerId` of `GET api/2.0/files/thirdparty`. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteThirdPartyRequest struct via the builder pattern


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


## GetAllProviders

> ProviderArrayWrapper GetAllProviders(ctx).Excludewebdav(excludewebdav).Execute()

Get all third-party providers



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-all-providers/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	excludewebdav := false // bool | Set to true to leave out the whole WebDAV family, the kDrive and Yandex presets included, and keep only the  services that authenticate through OAuth 2.0; false lists all of them. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesThirdPartyIntegrationAPI.GetAllProviders(context.Background()).Excludewebdav(excludewebdav).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesThirdPartyIntegrationAPI.GetAllProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllProviders`: ProviderArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesThirdPartyIntegrationAPI.GetAllProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **excludewebdav** | **bool** | Set to true to leave out the whole WebDAV family, the kDrive and Yandex presets included, and keep only the  services that authenticate through OAuth 2.0; false lists all of them. | 

### Return type

[**ProviderArrayWrapper**](ProviderArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupThirdPartyAccount

> ThirdPartyFolderWrapper GetBackupThirdPartyAccount(ctx).Execute()

Get the third-party backup folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-backup-third-party-account/).

### Example

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
	resp, r, err := apiClient.FilesThirdPartyIntegrationAPI.GetBackupThirdPartyAccount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesThirdPartyIntegrationAPI.GetBackupThirdPartyAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackupThirdPartyAccount`: ThirdPartyFolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesThirdPartyIntegrationAPI.GetBackupThirdPartyAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupThirdPartyAccountRequest struct via the builder pattern


### Return type

[**ThirdPartyFolderWrapper**](ThirdPartyFolderWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCapabilities

> ArrayArrayWrapper GetCapabilities(ctx).Execute()

Get third-party provider capabilities



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-capabilities/).

### Example

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
	resp, r, err := apiClient.FilesThirdPartyIntegrationAPI.GetCapabilities(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesThirdPartyIntegrationAPI.GetCapabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCapabilities`: ArrayArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesThirdPartyIntegrationAPI.GetCapabilities`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCapabilitiesRequest struct via the builder pattern


### Return type

[**ArrayArrayWrapper**](ArrayArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommonThirdPartyFolders

> ThirdPartyFolderArrayWrapper GetCommonThirdPartyFolders(ctx).Execute()

Get common third-party folders



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-common-third-party-folders/).

### Example

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
	resp, r, err := apiClient.FilesThirdPartyIntegrationAPI.GetCommonThirdPartyFolders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesThirdPartyIntegrationAPI.GetCommonThirdPartyFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommonThirdPartyFolders`: ThirdPartyFolderArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesThirdPartyIntegrationAPI.GetCommonThirdPartyFolders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommonThirdPartyFoldersRequest struct via the builder pattern


### Return type

[**ThirdPartyFolderArrayWrapper**](ThirdPartyFolderArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThirdPartyAccounts

> ThirdPartyParamsArrayWrapper GetThirdPartyAccounts(ctx).Execute()

Get the third-party accounts



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-third-party-accounts/).

### Example

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
	resp, r, err := apiClient.FilesThirdPartyIntegrationAPI.GetThirdPartyAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesThirdPartyIntegrationAPI.GetThirdPartyAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThirdPartyAccounts`: ThirdPartyParamsArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesThirdPartyIntegrationAPI.GetThirdPartyAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetThirdPartyAccountsRequest struct via the builder pattern


### Return type

[**ThirdPartyParamsArrayWrapper**](ThirdPartyParamsArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveThirdParty

> ThirdPartyFolderWrapper SaveThirdParty(ctx).ThirdPartyRequestDto(thirdPartyRequestDto).Execute()

Connect a third-party account



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/save-third-party/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	thirdPartyRequestDto := *openapiclient.NewThirdPartyRequestDto("Nextcloud storage", "Nextcloud") // ThirdPartyRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesThirdPartyIntegrationAPI.SaveThirdParty(context.Background()).ThirdPartyRequestDto(thirdPartyRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesThirdPartyIntegrationAPI.SaveThirdParty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveThirdParty`: ThirdPartyFolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesThirdPartyIntegrationAPI.SaveThirdParty`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveThirdPartyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **thirdPartyRequestDto** | [**ThirdPartyRequestDto**](ThirdPartyRequestDto.md) |  | 

### Return type

[**ThirdPartyFolderWrapper**](ThirdPartyFolderWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveThirdPartyBackup

> ThirdPartyFolderWrapper SaveThirdPartyBackup(ctx).ThirdPartyBackupRequestDto(thirdPartyBackupRequestDto).Execute()

Connect the third-party backup storage



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/save-third-party-backup/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	thirdPartyBackupRequestDto := *openapiclient.NewThirdPartyBackupRequestDto() // ThirdPartyBackupRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesThirdPartyIntegrationAPI.SaveThirdPartyBackup(context.Background()).ThirdPartyBackupRequestDto(thirdPartyBackupRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesThirdPartyIntegrationAPI.SaveThirdPartyBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveThirdPartyBackup`: ThirdPartyFolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesThirdPartyIntegrationAPI.SaveThirdPartyBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveThirdPartyBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **thirdPartyBackupRequestDto** | [**ThirdPartyBackupRequestDto**](ThirdPartyBackupRequestDto.md) |  | 

### Return type

[**ThirdPartyFolderWrapper**](ThirdPartyFolderWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

