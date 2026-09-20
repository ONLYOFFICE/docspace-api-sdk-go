# \SettingsRebrandingAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAdditionalWhiteLabelSettings**](SettingsRebrandingAPI.md#DeleteAdditionalWhiteLabelSettings) | **Delete** /api/2.0/settings/rebranding/additional | Delete the additional white label settings
[**DeleteCompanyWhiteLabelSettings**](SettingsRebrandingAPI.md#DeleteCompanyWhiteLabelSettings) | **Delete** /api/2.0/settings/rebranding/company | Delete the company white label settings
[**GetAdditionalWhiteLabelSettings**](SettingsRebrandingAPI.md#GetAdditionalWhiteLabelSettings) | **Get** /api/2.0/settings/rebranding/additional | Get the additional white label settings
[**GetCompanyWhiteLabelSettings**](SettingsRebrandingAPI.md#GetCompanyWhiteLabelSettings) | **Get** /api/2.0/settings/rebranding/company | Get the company white label settings
[**GetEnableWhitelabel**](SettingsRebrandingAPI.md#GetEnableWhitelabel) | **Get** /api/2.0/settings/enablewhitelabel | Check the white label availability
[**GetIsDefaultWhiteLabelLogoText**](SettingsRebrandingAPI.md#GetIsDefaultWhiteLabelLogoText) | **Get** /api/2.0/settings/whitelabel/logotext/isdefault | Check the default logo text
[**GetIsDefaultWhiteLabelLogos**](SettingsRebrandingAPI.md#GetIsDefaultWhiteLabelLogos) | **Get** /api/2.0/settings/whitelabel/logos/isdefault | Check the default white label logos
[**GetLicensorData**](SettingsRebrandingAPI.md#GetLicensorData) | **Get** /api/2.0/settings/companywhitelabel | Get the licensor data
[**GetWhiteLabelLogoText**](SettingsRebrandingAPI.md#GetWhiteLabelLogoText) | **Get** /api/2.0/settings/whitelabel/logotext | Get the white label logo text
[**GetWhiteLabelLogos**](SettingsRebrandingAPI.md#GetWhiteLabelLogos) | **Get** /api/2.0/settings/whitelabel/logos | Get the white label logos
[**RestoreWhiteLabelLogoText**](SettingsRebrandingAPI.md#RestoreWhiteLabelLogoText) | **Put** /api/2.0/settings/whitelabel/logotext/restore | Restore the white label logo text
[**RestoreWhiteLabelLogos**](SettingsRebrandingAPI.md#RestoreWhiteLabelLogos) | **Put** /api/2.0/settings/whitelabel/logos/restore | Restore the white label logos
[**SaveAdditionalWhiteLabelSettings**](SettingsRebrandingAPI.md#SaveAdditionalWhiteLabelSettings) | **Post** /api/2.0/settings/rebranding/additional | Save the additional white label settings
[**SaveCompanyWhiteLabelSettings**](SettingsRebrandingAPI.md#SaveCompanyWhiteLabelSettings) | **Post** /api/2.0/settings/rebranding/company | Save the company white label settings
[**SaveWhiteLabelLogoText**](SettingsRebrandingAPI.md#SaveWhiteLabelLogoText) | **Post** /api/2.0/settings/whitelabel/logotext/save | Save the white label logo text
[**SaveWhiteLabelSettings**](SettingsRebrandingAPI.md#SaveWhiteLabelSettings) | **Post** /api/2.0/settings/whitelabel/logos/save | Save the white label logos
[**SaveWhiteLabelSettingsFromFiles**](SettingsRebrandingAPI.md#SaveWhiteLabelSettingsFromFiles) | **Post** /api/2.0/settings/whitelabel/logos/savefromfiles | Save the logos from files



## DeleteAdditionalWhiteLabelSettings

> AdditionalWhiteLabelSettingsResponseWrapper DeleteAdditionalWhiteLabelSettings(ctx).Execute()

Delete the additional white label settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-additional-white-label-settings/).

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
	resp, r, err := apiClient.SettingsRebrandingAPI.DeleteAdditionalWhiteLabelSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsRebrandingAPI.DeleteAdditionalWhiteLabelSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAdditionalWhiteLabelSettings`: AdditionalWhiteLabelSettingsResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsRebrandingAPI.DeleteAdditionalWhiteLabelSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdditionalWhiteLabelSettingsRequest struct via the builder pattern


### Return type

[**AdditionalWhiteLabelSettingsResponseWrapper**](AdditionalWhiteLabelSettingsResponseWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCompanyWhiteLabelSettings

> CompanyWhiteLabelSettingsResponseWrapper DeleteCompanyWhiteLabelSettings(ctx).Execute()

Delete the company white label settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-company-white-label-settings/).

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
	resp, r, err := apiClient.SettingsRebrandingAPI.DeleteCompanyWhiteLabelSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsRebrandingAPI.DeleteCompanyWhiteLabelSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCompanyWhiteLabelSettings`: CompanyWhiteLabelSettingsResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsRebrandingAPI.DeleteCompanyWhiteLabelSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyWhiteLabelSettingsRequest struct via the builder pattern


### Return type

[**CompanyWhiteLabelSettingsResponseWrapper**](CompanyWhiteLabelSettingsResponseWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdditionalWhiteLabelSettings

> AdditionalWhiteLabelSettingsDtoWrapper GetAdditionalWhiteLabelSettings(ctx).Execute()

Get the additional white label settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-additional-white-label-settings/).

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
	resp, r, err := apiClient.SettingsRebrandingAPI.GetAdditionalWhiteLabelSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsRebrandingAPI.GetAdditionalWhiteLabelSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdditionalWhiteLabelSettings`: AdditionalWhiteLabelSettingsDtoWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsRebrandingAPI.GetAdditionalWhiteLabelSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdditionalWhiteLabelSettingsRequest struct via the builder pattern


### Return type

[**AdditionalWhiteLabelSettingsDtoWrapper**](AdditionalWhiteLabelSettingsDtoWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyWhiteLabelSettings

> CompanyWhiteLabelSettingsDtoWrapper GetCompanyWhiteLabelSettings(ctx).Execute()

Get the company white label settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-company-white-label-settings/).

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
	resp, r, err := apiClient.SettingsRebrandingAPI.GetCompanyWhiteLabelSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsRebrandingAPI.GetCompanyWhiteLabelSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyWhiteLabelSettings`: CompanyWhiteLabelSettingsDtoWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsRebrandingAPI.GetCompanyWhiteLabelSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyWhiteLabelSettingsRequest struct via the builder pattern


### Return type

[**CompanyWhiteLabelSettingsDtoWrapper**](CompanyWhiteLabelSettingsDtoWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnableWhitelabel

> BooleanWrapper GetEnableWhitelabel(ctx).Execute()

Check the white label availability



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-enable-whitelabel/).

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
	resp, r, err := apiClient.SettingsRebrandingAPI.GetEnableWhitelabel(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsRebrandingAPI.GetEnableWhitelabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnableWhitelabel`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsRebrandingAPI.GetEnableWhitelabel`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnableWhitelabelRequest struct via the builder pattern


### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIsDefaultWhiteLabelLogoText

> IsDefaultWhiteLabelLogosWrapper GetIsDefaultWhiteLabelLogoText(ctx).IsDark(isDark).IsDefault(isDefault).Execute()

Check the default logo text



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-is-default-white-label-logo-text/).

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
	isDark := true // bool | Which theme the answer is filled in for: `true` fills the dark image only, `false` the light one only.  Omitting it fills both, leaving the dark one empty for the slots that have no separate dark image. (optional)
	isDefault := true // bool | Whether the installation-wide default branding is addressed instead of this portal own. Writing the default  branding is only allowed on a self-hosted installation; elsewhere it is refused with 403. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsRebrandingAPI.GetIsDefaultWhiteLabelLogoText(context.Background()).IsDark(isDark).IsDefault(isDefault).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsRebrandingAPI.GetIsDefaultWhiteLabelLogoText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIsDefaultWhiteLabelLogoText`: IsDefaultWhiteLabelLogosWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsRebrandingAPI.GetIsDefaultWhiteLabelLogoText`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIsDefaultWhiteLabelLogoTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isDark** | **bool** | Which theme the answer is filled in for: `true` fills the dark image only, `false` the light one only.  Omitting it fills both, leaving the dark one empty for the slots that have no separate dark image. | 
 **isDefault** | **bool** | Whether the installation-wide default branding is addressed instead of this portal own. Writing the default  branding is only allowed on a self-hosted installation; elsewhere it is refused with 403. | 

### Return type

[**IsDefaultWhiteLabelLogosWrapper**](IsDefaultWhiteLabelLogosWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIsDefaultWhiteLabelLogos

> IsDefaultWhiteLabelLogosArrayWrapper GetIsDefaultWhiteLabelLogos(ctx).IsDark(isDark).IsDefault(isDefault).Execute()

Check the default white label logos



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-is-default-white-label-logos/).

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
	isDark := true // bool | Which theme the answer is filled in for: `true` fills the dark image only, `false` the light one only.  Omitting it fills both, leaving the dark one empty for the slots that have no separate dark image. (optional)
	isDefault := true // bool | Whether the installation-wide default branding is addressed instead of this portal own. Writing the default  branding is only allowed on a self-hosted installation; elsewhere it is refused with 403. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsRebrandingAPI.GetIsDefaultWhiteLabelLogos(context.Background()).IsDark(isDark).IsDefault(isDefault).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsRebrandingAPI.GetIsDefaultWhiteLabelLogos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIsDefaultWhiteLabelLogos`: IsDefaultWhiteLabelLogosArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsRebrandingAPI.GetIsDefaultWhiteLabelLogos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIsDefaultWhiteLabelLogosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isDark** | **bool** | Which theme the answer is filled in for: `true` fills the dark image only, `false` the light one only.  Omitting it fills both, leaving the dark one empty for the slots that have no separate dark image. | 
 **isDefault** | **bool** | Whether the installation-wide default branding is addressed instead of this portal own. Writing the default  branding is only allowed on a self-hosted installation; elsewhere it is refused with 403. | 

### Return type

[**IsDefaultWhiteLabelLogosArrayWrapper**](IsDefaultWhiteLabelLogosArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicensorData

> CompanyWhiteLabelSettingsArrayWrapper GetLicensorData(ctx).Execute()

Get the licensor data



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-licensor-data/).

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
	resp, r, err := apiClient.SettingsRebrandingAPI.GetLicensorData(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsRebrandingAPI.GetLicensorData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicensorData`: CompanyWhiteLabelSettingsArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsRebrandingAPI.GetLicensorData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicensorDataRequest struct via the builder pattern


### Return type

[**CompanyWhiteLabelSettingsArrayWrapper**](CompanyWhiteLabelSettingsArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhiteLabelLogoText

> StringWrapper GetWhiteLabelLogoText(ctx).IsDark(isDark).IsDefault(isDefault).Execute()

Get the white label logo text



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-white-label-logo-text/).

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
	isDark := true // bool | Which theme the answer is filled in for: `true` fills the dark image only, `false` the light one only.  Omitting it fills both, leaving the dark one empty for the slots that have no separate dark image. (optional)
	isDefault := true // bool | Whether the installation-wide default branding is addressed instead of this portal own. Writing the default  branding is only allowed on a self-hosted installation; elsewhere it is refused with 403. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsRebrandingAPI.GetWhiteLabelLogoText(context.Background()).IsDark(isDark).IsDefault(isDefault).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsRebrandingAPI.GetWhiteLabelLogoText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhiteLabelLogoText`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsRebrandingAPI.GetWhiteLabelLogoText`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWhiteLabelLogoTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isDark** | **bool** | Which theme the answer is filled in for: `true` fills the dark image only, `false` the light one only.  Omitting it fills both, leaving the dark one empty for the slots that have no separate dark image. | 
 **isDefault** | **bool** | Whether the installation-wide default branding is addressed instead of this portal own. Writing the default  branding is only allowed on a self-hosted installation; elsewhere it is refused with 403. | 

### Return type

[**StringWrapper**](StringWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhiteLabelLogos

> WhiteLabelItemArrayWrapper GetWhiteLabelLogos(ctx).IsDark(isDark).IsDefault(isDefault).Execute()

Get the white label logos



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-white-label-logos/).

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
	isDark := true // bool | Which theme the answer is filled in for: `true` fills the dark image only, `false` the light one only.  Omitting it fills both, leaving the dark one empty for the slots that have no separate dark image. (optional)
	isDefault := true // bool | Whether the installation-wide default branding is addressed instead of this portal own. Writing the default  branding is only allowed on a self-hosted installation; elsewhere it is refused with 403. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsRebrandingAPI.GetWhiteLabelLogos(context.Background()).IsDark(isDark).IsDefault(isDefault).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsRebrandingAPI.GetWhiteLabelLogos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWhiteLabelLogos`: WhiteLabelItemArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsRebrandingAPI.GetWhiteLabelLogos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWhiteLabelLogosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isDark** | **bool** | Which theme the answer is filled in for: `true` fills the dark image only, `false` the light one only.  Omitting it fills both, leaving the dark one empty for the slots that have no separate dark image. | 
 **isDefault** | **bool** | Whether the installation-wide default branding is addressed instead of this portal own. Writing the default  branding is only allowed on a self-hosted installation; elsewhere it is refused with 403. | 

### Return type

[**WhiteLabelItemArrayWrapper**](WhiteLabelItemArrayWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreWhiteLabelLogoText

> BooleanWrapper RestoreWhiteLabelLogoText(ctx).IsDark(isDark).IsDefault(isDefault).Execute()

Restore the white label logo text



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/restore-white-label-logo-text/).

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
	isDark := true // bool | Which theme the answer is filled in for: `true` fills the dark image only, `false` the light one only.  Omitting it fills both, leaving the dark one empty for the slots that have no separate dark image. (optional)
	isDefault := true // bool | Whether the installation-wide default branding is addressed instead of this portal own. Writing the default  branding is only allowed on a self-hosted installation; elsewhere it is refused with 403. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsRebrandingAPI.RestoreWhiteLabelLogoText(context.Background()).IsDark(isDark).IsDefault(isDefault).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsRebrandingAPI.RestoreWhiteLabelLogoText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreWhiteLabelLogoText`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsRebrandingAPI.RestoreWhiteLabelLogoText`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestoreWhiteLabelLogoTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isDark** | **bool** | Which theme the answer is filled in for: `true` fills the dark image only, `false` the light one only.  Omitting it fills both, leaving the dark one empty for the slots that have no separate dark image. | 
 **isDefault** | **bool** | Whether the installation-wide default branding is addressed instead of this portal own. Writing the default  branding is only allowed on a self-hosted installation; elsewhere it is refused with 403. | 

### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreWhiteLabelLogos

> BooleanWrapper RestoreWhiteLabelLogos(ctx).IsDark(isDark).IsDefault(isDefault).Execute()

Restore the white label logos



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/restore-white-label-logos/).

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
	isDark := true // bool | Which theme the answer is filled in for: `true` fills the dark image only, `false` the light one only.  Omitting it fills both, leaving the dark one empty for the slots that have no separate dark image. (optional)
	isDefault := true // bool | Whether the installation-wide default branding is addressed instead of this portal own. Writing the default  branding is only allowed on a self-hosted installation; elsewhere it is refused with 403. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsRebrandingAPI.RestoreWhiteLabelLogos(context.Background()).IsDark(isDark).IsDefault(isDefault).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsRebrandingAPI.RestoreWhiteLabelLogos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreWhiteLabelLogos`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsRebrandingAPI.RestoreWhiteLabelLogos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestoreWhiteLabelLogosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isDark** | **bool** | Which theme the answer is filled in for: `true` fills the dark image only, `false` the light one only.  Omitting it fills both, leaving the dark one empty for the slots that have no separate dark image. | 
 **isDefault** | **bool** | Whether the installation-wide default branding is addressed instead of this portal own. Writing the default  branding is only allowed on a self-hosted installation; elsewhere it is refused with 403. | 

### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveAdditionalWhiteLabelSettings

> BooleanWrapper SaveAdditionalWhiteLabelSettings(ctx).AdditionalWhiteLabelSettingsWrapper(additionalWhiteLabelSettingsWrapper).Execute()

Save the additional white label settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/save-additional-white-label-settings/).

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
	additionalWhiteLabelSettingsWrapper := *openapiclient.NewAdditionalWhiteLabelSettingsWrapper() // AdditionalWhiteLabelSettingsWrapper |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsRebrandingAPI.SaveAdditionalWhiteLabelSettings(context.Background()).AdditionalWhiteLabelSettingsWrapper(additionalWhiteLabelSettingsWrapper).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsRebrandingAPI.SaveAdditionalWhiteLabelSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveAdditionalWhiteLabelSettings`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsRebrandingAPI.SaveAdditionalWhiteLabelSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveAdditionalWhiteLabelSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **additionalWhiteLabelSettingsWrapper** | [**AdditionalWhiteLabelSettingsWrapper**](AdditionalWhiteLabelSettingsWrapper.md) |  | 

### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveCompanyWhiteLabelSettings

> BooleanWrapper SaveCompanyWhiteLabelSettings(ctx).CompanyWhiteLabelSettingsWrapper(companyWhiteLabelSettingsWrapper).Execute()

Save the company white label settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/save-company-white-label-settings/).

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
	companyWhiteLabelSettingsWrapper := *openapiclient.NewCompanyWhiteLabelSettingsWrapper() // CompanyWhiteLabelSettingsWrapper |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsRebrandingAPI.SaveCompanyWhiteLabelSettings(context.Background()).CompanyWhiteLabelSettingsWrapper(companyWhiteLabelSettingsWrapper).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsRebrandingAPI.SaveCompanyWhiteLabelSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveCompanyWhiteLabelSettings`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsRebrandingAPI.SaveCompanyWhiteLabelSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveCompanyWhiteLabelSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyWhiteLabelSettingsWrapper** | [**CompanyWhiteLabelSettingsWrapper**](CompanyWhiteLabelSettingsWrapper.md) |  | 

### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveWhiteLabelLogoText

> BooleanWrapper SaveWhiteLabelLogoText(ctx).IsDark(isDark).IsDefault(isDefault).WhiteLabelRequestsDto(whiteLabelRequestsDto).Execute()

Save the white label logo text



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/save-white-label-logo-text/).

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
	isDark := true // bool | Which theme the answer is filled in for: `true` fills the dark image only, `false` the light one only.  Omitting it fills both, leaving the dark one empty for the slots that have no separate dark image. (optional)
	isDefault := true // bool | Whether the installation-wide default branding is addressed instead of this portal own. Writing the default  branding is only allowed on a self-hosted installation; elsewhere it is refused with 403. (optional)
	whiteLabelRequestsDto := *openapiclient.NewWhiteLabelRequestsDto() // WhiteLabelRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsRebrandingAPI.SaveWhiteLabelLogoText(context.Background()).IsDark(isDark).IsDefault(isDefault).WhiteLabelRequestsDto(whiteLabelRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsRebrandingAPI.SaveWhiteLabelLogoText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveWhiteLabelLogoText`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsRebrandingAPI.SaveWhiteLabelLogoText`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveWhiteLabelLogoTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isDark** | **bool** | Which theme the answer is filled in for: `true` fills the dark image only, `false` the light one only.  Omitting it fills both, leaving the dark one empty for the slots that have no separate dark image. | 
 **isDefault** | **bool** | Whether the installation-wide default branding is addressed instead of this portal own. Writing the default  branding is only allowed on a self-hosted installation; elsewhere it is refused with 403. | 
 **whiteLabelRequestsDto** | [**WhiteLabelRequestsDto**](WhiteLabelRequestsDto.md) |  | 

### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveWhiteLabelSettings

> BooleanWrapper SaveWhiteLabelSettings(ctx).IsDark(isDark).IsDefault(isDefault).WhiteLabelRequestsDto(whiteLabelRequestsDto).Execute()

Save the white label logos



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/save-white-label-settings/).

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
	isDark := true // bool | Which theme the answer is filled in for: `true` fills the dark image only, `false` the light one only.  Omitting it fills both, leaving the dark one empty for the slots that have no separate dark image. (optional)
	isDefault := true // bool | Whether the installation-wide default branding is addressed instead of this portal own. Writing the default  branding is only allowed on a self-hosted installation; elsewhere it is refused with 403. (optional)
	whiteLabelRequestsDto := *openapiclient.NewWhiteLabelRequestsDto() // WhiteLabelRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsRebrandingAPI.SaveWhiteLabelSettings(context.Background()).IsDark(isDark).IsDefault(isDefault).WhiteLabelRequestsDto(whiteLabelRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsRebrandingAPI.SaveWhiteLabelSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveWhiteLabelSettings`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsRebrandingAPI.SaveWhiteLabelSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveWhiteLabelSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isDark** | **bool** | Which theme the answer is filled in for: `true` fills the dark image only, `false` the light one only.  Omitting it fills both, leaving the dark one empty for the slots that have no separate dark image. | 
 **isDefault** | **bool** | Whether the installation-wide default branding is addressed instead of this portal own. Writing the default  branding is only allowed on a self-hosted installation; elsewhere it is refused with 403. | 
 **whiteLabelRequestsDto** | [**WhiteLabelRequestsDto**](WhiteLabelRequestsDto.md) |  | 

### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveWhiteLabelSettingsFromFiles

> BooleanWrapper SaveWhiteLabelSettingsFromFiles(ctx).IsDark(isDark).IsDefault(isDefault).Execute()

Save the logos from files



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/save-white-label-settings-from-files/).

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
	isDark := true // bool | Which theme the answer is filled in for: `true` fills the dark image only, `false` the light one only.  Omitting it fills both, leaving the dark one empty for the slots that have no separate dark image. (optional)
	isDefault := true // bool | Whether the installation-wide default branding is addressed instead of this portal own. Writing the default  branding is only allowed on a self-hosted installation; elsewhere it is refused with 403. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsRebrandingAPI.SaveWhiteLabelSettingsFromFiles(context.Background()).IsDark(isDark).IsDefault(isDefault).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsRebrandingAPI.SaveWhiteLabelSettingsFromFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveWhiteLabelSettingsFromFiles`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsRebrandingAPI.SaveWhiteLabelSettingsFromFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveWhiteLabelSettingsFromFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isDark** | **bool** | Which theme the answer is filled in for: `true` fills the dark image only, `false` the light one only.  Omitting it fills both, leaving the dark one empty for the slots that have no separate dark image. | 
 **isDefault** | **bool** | Whether the installation-wide default branding is addressed instead of this portal own. Writing the default  branding is only allowed on a self-hosted installation; elsewhere it is refused with 403. | 

### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

