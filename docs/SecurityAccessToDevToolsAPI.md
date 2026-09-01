# \SecurityAccessToDevToolsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SetTenantDevToolsAccessSettings**](SecurityAccessToDevToolsAPI.md#SetTenantDevToolsAccessSettings) | **Post** /api/2.0/settings/devtoolsaccess | Set the Developer Tools access settings



## SetTenantDevToolsAccessSettings

> TenantDevToolsAccessSettingsWrapper SetTenantDevToolsAccessSettings(ctx).TenantDevToolsAccessSettingsDto(tenantDevToolsAccessSettingsDto).Execute()

Set the Developer Tools access settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-tenant-dev-tools-access-settings/).

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
	tenantDevToolsAccessSettingsDto := *openapiclient.NewTenantDevToolsAccessSettingsDto() // TenantDevToolsAccessSettingsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAccessToDevToolsAPI.SetTenantDevToolsAccessSettings(context.Background()).TenantDevToolsAccessSettingsDto(tenantDevToolsAccessSettingsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAccessToDevToolsAPI.SetTenantDevToolsAccessSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetTenantDevToolsAccessSettings`: TenantDevToolsAccessSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityAccessToDevToolsAPI.SetTenantDevToolsAccessSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetTenantDevToolsAccessSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantDevToolsAccessSettingsDto** | [**TenantDevToolsAccessSettingsDto**](TenantDevToolsAccessSettingsDto.md) |  | 

### Return type

[**TenantDevToolsAccessSettingsWrapper**](TenantDevToolsAccessSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

