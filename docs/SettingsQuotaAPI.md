# \SettingsQuotaAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserQuotaSettings**](SettingsQuotaAPI.md#GetUserQuotaSettings) | **Get** /api/2.0/settings/userquotasettings | Get the user quota settings
[**SaveAiAgentQuotaSettings**](SettingsQuotaAPI.md#SaveAiAgentQuotaSettings) | **Post** /api/2.0/settings/aiagentquotasettings | Save the AI Agent quota settings
[**SaveRoomQuotaSettings**](SettingsQuotaAPI.md#SaveRoomQuotaSettings) | **Post** /api/2.0/settings/roomquotasettings | Save the room quota settings
[**SetTenantQuotaSettings**](SettingsQuotaAPI.md#SetTenantQuotaSettings) | **Put** /api/2.0/settings/tenantquotasettings | Save the tenant quota settings



## GetUserQuotaSettings

> TenantUserQuotaSettingsWrapper GetUserQuotaSettings(ctx).Execute()

Get the user quota settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-user-quota-settings/).

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
	resp, r, err := apiClient.SettingsQuotaAPI.GetUserQuotaSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsQuotaAPI.GetUserQuotaSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserQuotaSettings`: TenantUserQuotaSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsQuotaAPI.GetUserQuotaSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserQuotaSettingsRequest struct via the builder pattern


### Return type

[**TenantUserQuotaSettingsWrapper**](TenantUserQuotaSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveAiAgentQuotaSettings

> TenantAiAgentQuotaSettingsWrapper SaveAiAgentQuotaSettings(ctx).QuotaSettingsRequestsDto(quotaSettingsRequestsDto).Execute()

Save the AI Agent quota settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/save-ai-agent-quota-settings/).

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
	quotaSettingsRequestsDto := *openapiclient.NewQuotaSettingsRequestsDto(openapiclient.QuotaSettingsRequestsDto_defaultQuota{Int32: new(int32)}) // QuotaSettingsRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsQuotaAPI.SaveAiAgentQuotaSettings(context.Background()).QuotaSettingsRequestsDto(quotaSettingsRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsQuotaAPI.SaveAiAgentQuotaSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveAiAgentQuotaSettings`: TenantAiAgentQuotaSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsQuotaAPI.SaveAiAgentQuotaSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveAiAgentQuotaSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quotaSettingsRequestsDto** | [**QuotaSettingsRequestsDto**](QuotaSettingsRequestsDto.md) |  | 

### Return type

[**TenantAiAgentQuotaSettingsWrapper**](TenantAiAgentQuotaSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveRoomQuotaSettings

> TenantRoomQuotaSettingsWrapper SaveRoomQuotaSettings(ctx).QuotaSettingsRequestsDto(quotaSettingsRequestsDto).Execute()

Save the room quota settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/save-room-quota-settings/).

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
	quotaSettingsRequestsDto := *openapiclient.NewQuotaSettingsRequestsDto(openapiclient.QuotaSettingsRequestsDto_defaultQuota{Int32: new(int32)}) // QuotaSettingsRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsQuotaAPI.SaveRoomQuotaSettings(context.Background()).QuotaSettingsRequestsDto(quotaSettingsRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsQuotaAPI.SaveRoomQuotaSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveRoomQuotaSettings`: TenantRoomQuotaSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsQuotaAPI.SaveRoomQuotaSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveRoomQuotaSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quotaSettingsRequestsDto** | [**QuotaSettingsRequestsDto**](QuotaSettingsRequestsDto.md) |  | 

### Return type

[**TenantRoomQuotaSettingsWrapper**](TenantRoomQuotaSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTenantQuotaSettings

> TenantQuotaSettingsWrapper SetTenantQuotaSettings(ctx).TenantQuotaSettingsRequestsDto(tenantQuotaSettingsRequestsDto).Execute()

Save the tenant quota settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-tenant-quota-settings/).

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
	tenantQuotaSettingsRequestsDto := *openapiclient.NewTenantQuotaSettingsRequestsDto(int32(1)) // TenantQuotaSettingsRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsQuotaAPI.SetTenantQuotaSettings(context.Background()).TenantQuotaSettingsRequestsDto(tenantQuotaSettingsRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsQuotaAPI.SetTenantQuotaSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetTenantQuotaSettings`: TenantQuotaSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsQuotaAPI.SetTenantQuotaSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetTenantQuotaSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantQuotaSettingsRequestsDto** | [**TenantQuotaSettingsRequestsDto**](TenantQuotaSettingsRequestsDto.md) |  | 

### Return type

[**TenantQuotaSettingsWrapper**](TenantQuotaSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

