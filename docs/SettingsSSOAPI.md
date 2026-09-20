# \SettingsSSOAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDefaultSsoSettingsV2**](SettingsSSOAPI.md#GetDefaultSsoSettingsV2) | **Get** /api/2.0/settings/ssov2/default | Get the default SSO settings
[**GetSsoSettingsV2**](SettingsSSOAPI.md#GetSsoSettingsV2) | **Get** /api/2.0/settings/ssov2 | Get the SSO settings
[**GetSsoSettingsV2Constants**](SettingsSSOAPI.md#GetSsoSettingsV2Constants) | **Get** /api/2.0/settings/ssov2/constants | Get the SSO settings constants
[**ResetSsoSettingsV2**](SettingsSSOAPI.md#ResetSsoSettingsV2) | **Delete** /api/2.0/settings/ssov2 | Reset the SSO settings
[**SaveSsoSettingsV2**](SettingsSSOAPI.md#SaveSsoSettingsV2) | **Post** /api/2.0/settings/ssov2 | Save the SSO settings



## GetDefaultSsoSettingsV2

> SsoSettingsV2Wrapper GetDefaultSsoSettingsV2(ctx).Execute()

Get the default SSO settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-default-sso-settings-v2/).

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
	resp, r, err := apiClient.SettingsSSOAPI.GetDefaultSsoSettingsV2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsSSOAPI.GetDefaultSsoSettingsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultSsoSettingsV2`: SsoSettingsV2Wrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsSSOAPI.GetDefaultSsoSettingsV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultSsoSettingsV2Request struct via the builder pattern


### Return type

[**SsoSettingsV2Wrapper**](SsoSettingsV2Wrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSsoSettingsV2

> SsoSettingsV2Wrapper GetSsoSettingsV2(ctx).Execute()

Get the SSO settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-sso-settings-v2/).

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
	resp, r, err := apiClient.SettingsSSOAPI.GetSsoSettingsV2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsSSOAPI.GetSsoSettingsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSsoSettingsV2`: SsoSettingsV2Wrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsSSOAPI.GetSsoSettingsV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSsoSettingsV2Request struct via the builder pattern


### Return type

[**SsoSettingsV2Wrapper**](SsoSettingsV2Wrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSsoSettingsV2Constants

> SsoSettingsV2ConstantsWrapper GetSsoSettingsV2Constants(ctx).Execute()

Get the SSO settings constants



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-sso-settings-v2constants/).

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
	resp, r, err := apiClient.SettingsSSOAPI.GetSsoSettingsV2Constants(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsSSOAPI.GetSsoSettingsV2Constants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSsoSettingsV2Constants`: SsoSettingsV2ConstantsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsSSOAPI.GetSsoSettingsV2Constants`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSsoSettingsV2ConstantsRequest struct via the builder pattern


### Return type

[**SsoSettingsV2ConstantsWrapper**](SsoSettingsV2ConstantsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetSsoSettingsV2

> SsoSettingsV2Wrapper ResetSsoSettingsV2(ctx).Execute()

Reset the SSO settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/reset-sso-settings-v2/).

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
	resp, r, err := apiClient.SettingsSSOAPI.ResetSsoSettingsV2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsSSOAPI.ResetSsoSettingsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetSsoSettingsV2`: SsoSettingsV2Wrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsSSOAPI.ResetSsoSettingsV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiResetSsoSettingsV2Request struct via the builder pattern


### Return type

[**SsoSettingsV2Wrapper**](SsoSettingsV2Wrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveSsoSettingsV2

> SsoSettingsV2Wrapper SaveSsoSettingsV2(ctx).SsoSettingsRequestsDto(ssoSettingsRequestsDto).Execute()

Save the SSO settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/save-sso-settings-v2/).

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
	ssoSettingsRequestsDto := *openapiclient.NewSsoSettingsRequestsDto("{\"enableSso\":true,\"idpSettings\":{\"entityId\":\"https://idp.example.com\"}}") // SsoSettingsRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsSSOAPI.SaveSsoSettingsV2(context.Background()).SsoSettingsRequestsDto(ssoSettingsRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsSSOAPI.SaveSsoSettingsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveSsoSettingsV2`: SsoSettingsV2Wrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsSSOAPI.SaveSsoSettingsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveSsoSettingsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ssoSettingsRequestsDto** | [**SsoSettingsRequestsDto**](SsoSettingsRequestsDto.md) |  | 

### Return type

[**SsoSettingsV2Wrapper**](SsoSettingsV2Wrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

