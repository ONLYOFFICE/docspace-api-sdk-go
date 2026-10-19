# \SettingsAuthorizationAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthServices**](SettingsAuthorizationAPI.md#GetAuthServices) | **Get** /api/2.0/settings/authservice | Get the authorization services
[**SaveAuthKeys**](SettingsAuthorizationAPI.md#SaveAuthKeys) | **Post** /api/2.0/settings/authservice | Save the authorization keys
[**TestExternalDatabaseConnection**](SettingsAuthorizationAPI.md#TestExternalDatabaseConnection) | **Post** /api/2.0/settings/authservice/externaldb/test | Test external database connection



## GetAuthServices

> AuthServiceArrayWrapper GetAuthServices(ctx).Execute()

Get the authorization services



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-auth-services/).

### Example

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
	resp, r, err := apiClient.SettingsAuthorizationAPI.GetAuthServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAuthorizationAPI.GetAuthServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthServices`: AuthServiceArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsAuthorizationAPI.GetAuthServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthServicesRequest struct via the builder pattern


### Return type

[**AuthServiceArrayWrapper**](AuthServiceArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveAuthKeys

> BooleanWrapper SaveAuthKeys(ctx).SaveAuthKeysRequestDto(saveAuthKeysRequestDto).Execute()

Save the authorization keys



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/save-auth-keys/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	saveAuthKeysRequestDto := *openapiclient.NewSaveAuthKeysRequestDto() // SaveAuthKeysRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAuthorizationAPI.SaveAuthKeys(context.Background()).SaveAuthKeysRequestDto(saveAuthKeysRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAuthorizationAPI.SaveAuthKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveAuthKeys`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsAuthorizationAPI.SaveAuthKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveAuthKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saveAuthKeysRequestDto** | [**SaveAuthKeysRequestDto**](SaveAuthKeysRequestDto.md) |  | 

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


## TestExternalDatabaseConnection

> ConnectionTestResultWrapper TestExternalDatabaseConnection(ctx).ExternalDatabaseConnectionRequestDto(externalDatabaseConnectionRequestDto).Execute()

Test external database connection



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/test-external-database-connection/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	externalDatabaseConnectionRequestDto := *openapiclient.NewExternalDatabaseConnectionRequestDto() // ExternalDatabaseConnectionRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAuthorizationAPI.TestExternalDatabaseConnection(context.Background()).ExternalDatabaseConnectionRequestDto(externalDatabaseConnectionRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAuthorizationAPI.TestExternalDatabaseConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestExternalDatabaseConnection`: ConnectionTestResultWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsAuthorizationAPI.TestExternalDatabaseConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestExternalDatabaseConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalDatabaseConnectionRequestDto** | [**ExternalDatabaseConnectionRequestDto**](ExternalDatabaseConnectionRequestDto.md) |  | 

### Return type

[**ConnectionTestResultWrapper**](ConnectionTestResultWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

