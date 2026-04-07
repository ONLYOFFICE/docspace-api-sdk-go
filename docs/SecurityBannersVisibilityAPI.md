# \SecurityBannersVisibilityAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SetTenantBannerSettings**](SecurityBannersVisibilityAPI.md#SetTenantBannerSettings) | **Post** /api/2.0/settings/banner | Set the banners visibility



## SetTenantBannerSettings

> TenantBannerSettingsWrapper SetTenantBannerSettings(ctx).TenantBannerSettingsDto(tenantBannerSettingsDto).Execute()

Set the banners visibility



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-tenant-banner-settings/).

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
	tenantBannerSettingsDto := *openapiclient.NewTenantBannerSettingsDto() // TenantBannerSettingsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityBannersVisibilityAPI.SetTenantBannerSettings(context.Background()).TenantBannerSettingsDto(tenantBannerSettingsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityBannersVisibilityAPI.SetTenantBannerSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetTenantBannerSettings`: TenantBannerSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityBannersVisibilityAPI.SetTenantBannerSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetTenantBannerSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantBannerSettingsDto** | [**TenantBannerSettingsDto**](TenantBannerSettingsDto.md) |  | 

### Return type

[**TenantBannerSettingsWrapper**](TenantBannerSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

