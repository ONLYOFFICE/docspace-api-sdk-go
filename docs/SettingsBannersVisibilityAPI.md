# \SettingsBannersVisibilityAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTenantBannerSettings**](SettingsBannersVisibilityAPI.md#GetTenantBannerSettings) | **Get** /api/2.0/settings/banner | Get the banners visibility



## GetTenantBannerSettings

> TenantBannerSettingsWrapper GetTenantBannerSettings(ctx).Execute()

Get the banners visibility



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-tenant-banner-settings/).

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
	resp, r, err := apiClient.SettingsBannersVisibilityAPI.GetTenantBannerSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsBannersVisibilityAPI.GetTenantBannerSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantBannerSettings`: TenantBannerSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsBannersVisibilityAPI.GetTenantBannerSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantBannerSettingsRequest struct via the builder pattern


### Return type

[**TenantBannerSettingsWrapper**](TenantBannerSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

