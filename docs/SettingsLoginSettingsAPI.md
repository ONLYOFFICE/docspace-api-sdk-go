# \SettingsLoginSettingsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLoginSettings**](SettingsLoginSettingsAPI.md#GetLoginSettings) | **Get** /api/2.0/settings/security/loginsettings | Get the login settings
[**SetDefaultLoginSettings**](SettingsLoginSettingsAPI.md#SetDefaultLoginSettings) | **Delete** /api/2.0/settings/security/loginsettings | Reset the login settings
[**UpdateLoginSettings**](SettingsLoginSettingsAPI.md#UpdateLoginSettings) | **Put** /api/2.0/settings/security/loginsettings | Update the login settings



## GetLoginSettings

> LoginSettingsWrapper GetLoginSettings(ctx).Execute()

Get the login settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-login-settings/).

### Example

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
	resp, r, err := apiClient.SettingsLoginSettingsAPI.GetLoginSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsLoginSettingsAPI.GetLoginSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoginSettings`: LoginSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsLoginSettingsAPI.GetLoginSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoginSettingsRequest struct via the builder pattern


### Return type

[**LoginSettingsWrapper**](LoginSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultLoginSettings

> LoginSettingsWrapper SetDefaultLoginSettings(ctx).Execute()

Reset the login settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-default-login-settings/).

### Example

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
	resp, r, err := apiClient.SettingsLoginSettingsAPI.SetDefaultLoginSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsLoginSettingsAPI.SetDefaultLoginSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDefaultLoginSettings`: LoginSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsLoginSettingsAPI.SetDefaultLoginSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultLoginSettingsRequest struct via the builder pattern


### Return type

[**LoginSettingsWrapper**](LoginSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoginSettings

> LoginSettingsWrapper UpdateLoginSettings(ctx).LoginSettingsRequestDto(loginSettingsRequestDto).Execute()

Update the login settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-login-settings/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	loginSettingsRequestDto := *openapiclient.NewLoginSettingsRequestDto() // LoginSettingsRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsLoginSettingsAPI.UpdateLoginSettings(context.Background()).LoginSettingsRequestDto(loginSettingsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsLoginSettingsAPI.UpdateLoginSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLoginSettings`: LoginSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsLoginSettingsAPI.UpdateLoginSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoginSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginSettingsRequestDto** | [**LoginSettingsRequestDto**](LoginSettingsRequestDto.md) |  | 

### Return type

[**LoginSettingsWrapper**](LoginSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

