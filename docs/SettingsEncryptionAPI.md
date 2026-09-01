# \SettingsEncryptionAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStorageEncryptionProgress**](SettingsEncryptionAPI.md#GetStorageEncryptionProgress) | **Get** /api/2.0/settings/encryption/progress | Get the storage encryption progress
[**GetStorageEncryptionSettings**](SettingsEncryptionAPI.md#GetStorageEncryptionSettings) | **Get** /api/2.0/settings/encryption/settings | Get the storage encryption settings
[**StartStorageEncryption**](SettingsEncryptionAPI.md#StartStorageEncryption) | **Post** /api/2.0/settings/encryption/start | Start the storage encryption process



## GetStorageEncryptionProgress

> DoubleNullableWrapper GetStorageEncryptionProgress(ctx).Execute()

Get the storage encryption progress



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-storage-encryption-progress/).

### Example

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
	resp, r, err := apiClient.SettingsEncryptionAPI.GetStorageEncryptionProgress(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsEncryptionAPI.GetStorageEncryptionProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageEncryptionProgress`: DoubleNullableWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsEncryptionAPI.GetStorageEncryptionProgress`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageEncryptionProgressRequest struct via the builder pattern


### Return type

[**DoubleNullableWrapper**](DoubleNullableWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageEncryptionSettings

> EncryptionSettingsWrapper GetStorageEncryptionSettings(ctx).Execute()

Get the storage encryption settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-storage-encryption-settings/).

### Example

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
	resp, r, err := apiClient.SettingsEncryptionAPI.GetStorageEncryptionSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsEncryptionAPI.GetStorageEncryptionSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageEncryptionSettings`: EncryptionSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsEncryptionAPI.GetStorageEncryptionSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageEncryptionSettingsRequest struct via the builder pattern


### Return type

[**EncryptionSettingsWrapper**](EncryptionSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartStorageEncryption

> BooleanWrapper StartStorageEncryption(ctx).StorageEncryptionRequestsDto(storageEncryptionRequestsDto).Execute()

Start the storage encryption process



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/start-storage-encryption/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	storageEncryptionRequestsDto := *openapiclient.NewStorageEncryptionRequestsDto() // StorageEncryptionRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsEncryptionAPI.StartStorageEncryption(context.Background()).StorageEncryptionRequestsDto(storageEncryptionRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsEncryptionAPI.StartStorageEncryption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartStorageEncryption`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsEncryptionAPI.StartStorageEncryption`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartStorageEncryptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storageEncryptionRequestsDto** | [**StorageEncryptionRequestsDto**](StorageEncryptionRequestsDto.md) |  | 

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

