# \SettingsAccessToDevToolsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTenantAccessDevToolsSettings**](SettingsAccessToDevToolsAPI.md#GetTenantAccessDevToolsSettings) | **Get** /api/2.0/settings/devtoolsaccess | Get the Developer Tools access settings



## GetTenantAccessDevToolsSettings

> TenantDevToolsAccessSettingsWrapper GetTenantAccessDevToolsSettings(ctx).Execute()

Get the Developer Tools access settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-tenant-access-dev-tools-settings/).

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
	resp, r, err := apiClient.SettingsAccessToDevToolsAPI.GetTenantAccessDevToolsSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAccessToDevToolsAPI.GetTenantAccessDevToolsSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantAccessDevToolsSettings`: TenantDevToolsAccessSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsAccessToDevToolsAPI.GetTenantAccessDevToolsSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantAccessDevToolsSettingsRequest struct via the builder pattern


### Return type

[**TenantDevToolsAccessSettingsWrapper**](TenantDevToolsAccessSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

