# \SettingsIPRestrictionsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIpRestrictions**](SettingsIPRestrictionsAPI.md#GetIpRestrictions) | **Get** /api/2.0/settings/iprestrictions | Get the IP portal restrictions
[**ReadIpRestrictionsSettings**](SettingsIPRestrictionsAPI.md#ReadIpRestrictionsSettings) | **Get** /api/2.0/settings/iprestrictions/settings | Get the IP restriction settings
[**SaveIpRestrictions**](SettingsIPRestrictionsAPI.md#SaveIpRestrictions) | **Put** /api/2.0/settings/iprestrictions | Update the IP restrictions
[**UpdateIpRestrictionsSettings**](SettingsIPRestrictionsAPI.md#UpdateIpRestrictionsSettings) | **Put** /api/2.0/settings/iprestrictions/settings | Update the IP restriction settings



## GetIpRestrictions

> IPRestrictionArrayWrapper GetIpRestrictions(ctx).Execute()

Get the IP portal restrictions



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-ip-restrictions/).

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
	resp, r, err := apiClient.SettingsIPRestrictionsAPI.GetIpRestrictions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsIPRestrictionsAPI.GetIpRestrictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpRestrictions`: IPRestrictionArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsIPRestrictionsAPI.GetIpRestrictions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpRestrictionsRequest struct via the builder pattern


### Return type

[**IPRestrictionArrayWrapper**](IPRestrictionArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadIpRestrictionsSettings

> IPRestrictionsSettingsWrapper ReadIpRestrictionsSettings(ctx).Execute()

Get the IP restriction settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/read-ip-restrictions-settings/).

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
	resp, r, err := apiClient.SettingsIPRestrictionsAPI.ReadIpRestrictionsSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsIPRestrictionsAPI.ReadIpRestrictionsSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadIpRestrictionsSettings`: IPRestrictionsSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsIPRestrictionsAPI.ReadIpRestrictionsSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReadIpRestrictionsSettingsRequest struct via the builder pattern


### Return type

[**IPRestrictionsSettingsWrapper**](IPRestrictionsSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveIpRestrictions

> IpRestrictionsWrapper SaveIpRestrictions(ctx).IpRestrictionsDto(ipRestrictionsDto).Execute()

Update the IP restrictions



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/save-ip-restrictions/).

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
	ipRestrictionsDto := *openapiclient.NewIpRestrictionsDto([]openapiclient.IpRestrictionBase{*openapiclient.NewIpRestrictionBase("Ip_example")}) // IpRestrictionsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsIPRestrictionsAPI.SaveIpRestrictions(context.Background()).IpRestrictionsDto(ipRestrictionsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsIPRestrictionsAPI.SaveIpRestrictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveIpRestrictions`: IpRestrictionsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsIPRestrictionsAPI.SaveIpRestrictions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveIpRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ipRestrictionsDto** | [**IpRestrictionsDto**](IpRestrictionsDto.md) |  | 

### Return type

[**IpRestrictionsWrapper**](IpRestrictionsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIpRestrictionsSettings

> IpRestrictionsWrapper UpdateIpRestrictionsSettings(ctx).IpRestrictionsDto(ipRestrictionsDto).Execute()

Update the IP restriction settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-ip-restrictions-settings/).

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
	ipRestrictionsDto := *openapiclient.NewIpRestrictionsDto([]openapiclient.IpRestrictionBase{*openapiclient.NewIpRestrictionBase("Ip_example")}) // IpRestrictionsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsIPRestrictionsAPI.UpdateIpRestrictionsSettings(context.Background()).IpRestrictionsDto(ipRestrictionsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsIPRestrictionsAPI.UpdateIpRestrictionsSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIpRestrictionsSettings`: IpRestrictionsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsIPRestrictionsAPI.UpdateIpRestrictionsSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIpRestrictionsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ipRestrictionsDto** | [**IpRestrictionsDto**](IpRestrictionsDto.md) |  | 

### Return type

[**IpRestrictionsWrapper**](IpRestrictionsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

