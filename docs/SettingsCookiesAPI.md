# \SettingsCookiesAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCookieSettings**](SettingsCookiesAPI.md#GetCookieSettings) | **Get** /api/2.0/settings/cookiesettings | Get cookies lifetime
[**UpdateCookieSettings**](SettingsCookiesAPI.md#UpdateCookieSettings) | **Put** /api/2.0/settings/cookiesettings | Update cookies lifetime



## GetCookieSettings

> CookieSettingsWrapper GetCookieSettings(ctx).Execute()

Get cookies lifetime



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-cookie-settings/).

### Example

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
	resp, r, err := apiClient.SettingsCookiesAPI.GetCookieSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsCookiesAPI.GetCookieSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCookieSettings`: CookieSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsCookiesAPI.GetCookieSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCookieSettingsRequest struct via the builder pattern


### Return type

[**CookieSettingsWrapper**](CookieSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCookieSettings

> StringWrapper UpdateCookieSettings(ctx).CookieSettingsRequestsDto(cookieSettingsRequestsDto).Execute()

Update cookies lifetime



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-cookie-settings/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cookieSettingsRequestsDto := *openapiclient.NewCookieSettingsRequestsDto() // CookieSettingsRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsCookiesAPI.UpdateCookieSettings(context.Background()).CookieSettingsRequestsDto(cookieSettingsRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsCookiesAPI.UpdateCookieSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCookieSettings`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsCookiesAPI.UpdateCookieSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCookieSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookieSettingsRequestsDto** | [**CookieSettingsRequestsDto**](CookieSettingsRequestsDto.md) |  | 

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

