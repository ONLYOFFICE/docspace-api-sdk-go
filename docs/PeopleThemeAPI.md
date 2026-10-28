# \PeopleThemeAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangePortalTheme**](PeopleThemeAPI.md#ChangePortalTheme) | **Put** /api/2.0/people/theme | Change the portal theme
[**GetPortalTheme**](PeopleThemeAPI.md#GetPortalTheme) | **Get** /api/2.0/people/theme | Get the portal theme



## ChangePortalTheme

> DarkThemeSettingsWrapper ChangePortalTheme(ctx).DarkThemeSettingsRequestDto(darkThemeSettingsRequestDto).Execute()

Change the portal theme



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/change-portal-theme/).

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
	darkThemeSettingsRequestDto := *openapiclient.NewDarkThemeSettingsRequestDto(openapiclient.DarkThemeSettingsType("Base")) // DarkThemeSettingsRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleThemeAPI.ChangePortalTheme(context.Background()).DarkThemeSettingsRequestDto(darkThemeSettingsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleThemeAPI.ChangePortalTheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangePortalTheme`: DarkThemeSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleThemeAPI.ChangePortalTheme`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangePortalThemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **darkThemeSettingsRequestDto** | [**DarkThemeSettingsRequestDto**](DarkThemeSettingsRequestDto.md) |  | 

### Return type

[**DarkThemeSettingsWrapper**](DarkThemeSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortalTheme

> DarkThemeSettingsWrapper GetPortalTheme(ctx).Execute()

Get the portal theme



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-portal-theme/).

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
	resp, r, err := apiClient.PeopleThemeAPI.GetPortalTheme(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleThemeAPI.GetPortalTheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortalTheme`: DarkThemeSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleThemeAPI.GetPortalTheme`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortalThemeRequest struct via the builder pattern


### Return type

[**DarkThemeSettingsWrapper**](DarkThemeSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

