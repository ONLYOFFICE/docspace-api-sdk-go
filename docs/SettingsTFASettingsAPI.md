# \SettingsTFASettingsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTfaAppCodes**](SettingsTFASettingsAPI.md#GetTfaAppCodes) | **Get** /api/2.0/settings/tfaappcodes | Get the TFA backup codes
[**GetTfaConfirmData**](SettingsTFASettingsAPI.md#GetTfaConfirmData) | **Get** /api/2.0/settings/tfaapp/confirm | Get TFA confirmation data
[**GetTfaSettings**](SettingsTFASettingsAPI.md#GetTfaSettings) | **Get** /api/2.0/settings/tfaapp | Get the TFA settings
[**TfaAppGenerateSetupCode**](SettingsTFASettingsAPI.md#TfaAppGenerateSetupCode) | **Get** /api/2.0/settings/tfaapp/setup | Generate the TFA setup code
[**TfaValidateAuthCode**](SettingsTFASettingsAPI.md#TfaValidateAuthCode) | **Post** /api/2.0/settings/tfaapp/validate | Validate the TFA code
[**UnlinkTfaApp**](SettingsTFASettingsAPI.md#UnlinkTfaApp) | **Put** /api/2.0/settings/tfaappnewapp | Unlink the TFA application
[**UpdateTfaAppCodes**](SettingsTFASettingsAPI.md#UpdateTfaAppCodes) | **Put** /api/2.0/settings/tfaappnewcodes | Regenerate the TFA backup codes
[**UpdateTfaSettings**](SettingsTFASettingsAPI.md#UpdateTfaSettings) | **Put** /api/2.0/settings/tfaapp | Update the TFA settings
[**UpdateTfaSettingsLink**](SettingsTFASettingsAPI.md#UpdateTfaSettingsLink) | **Put** /api/2.0/settings/tfaappwithlink | Update TFA settings with a link



## GetTfaAppCodes

> TfaAppCodeArrayWrapper GetTfaAppCodes(ctx).Execute()

Get the TFA backup codes



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-tfa-app-codes/).

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
	resp, r, err := apiClient.SettingsTFASettingsAPI.GetTfaAppCodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsTFASettingsAPI.GetTfaAppCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTfaAppCodes`: TfaAppCodeArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsTFASettingsAPI.GetTfaAppCodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTfaAppCodesRequest struct via the builder pattern


### Return type

[**TfaAppCodeArrayWrapper**](TfaAppCodeArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTfaConfirmData

> TfaConfirmDataWrapper GetTfaConfirmData(ctx).Execute()

Get TFA confirmation data



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-tfa-confirm-data/).

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
	resp, r, err := apiClient.SettingsTFASettingsAPI.GetTfaConfirmData(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsTFASettingsAPI.GetTfaConfirmData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTfaConfirmData`: TfaConfirmDataWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsTFASettingsAPI.GetTfaConfirmData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTfaConfirmDataRequest struct via the builder pattern


### Return type

[**TfaConfirmDataWrapper**](TfaConfirmDataWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTfaSettings

> TfaSettingsArrayWrapper GetTfaSettings(ctx).Execute()

Get the TFA settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-tfa-settings/).

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
	resp, r, err := apiClient.SettingsTFASettingsAPI.GetTfaSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsTFASettingsAPI.GetTfaSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTfaSettings`: TfaSettingsArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsTFASettingsAPI.GetTfaSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTfaSettingsRequest struct via the builder pattern


### Return type

[**TfaSettingsArrayWrapper**](TfaSettingsArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TfaAppGenerateSetupCode

> TfaSetupCodeWrapper TfaAppGenerateSetupCode(ctx).Execute()

Generate the TFA setup code



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/tfa-app-generate-setup-code/).

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
	resp, r, err := apiClient.SettingsTFASettingsAPI.TfaAppGenerateSetupCode(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsTFASettingsAPI.TfaAppGenerateSetupCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TfaAppGenerateSetupCode`: TfaSetupCodeWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsTFASettingsAPI.TfaAppGenerateSetupCode`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTfaAppGenerateSetupCodeRequest struct via the builder pattern


### Return type

[**TfaSetupCodeWrapper**](TfaSetupCodeWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TfaValidateAuthCode

> BooleanWrapper TfaValidateAuthCode(ctx).TfaValidateRequestsDto(tfaValidateRequestsDto).Execute()

Validate the TFA code



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/tfa-validate-auth-code/).

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
	tfaValidateRequestsDto := *openapiclient.NewTfaValidateRequestsDto("123456") // TfaValidateRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsTFASettingsAPI.TfaValidateAuthCode(context.Background()).TfaValidateRequestsDto(tfaValidateRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsTFASettingsAPI.TfaValidateAuthCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TfaValidateAuthCode`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsTFASettingsAPI.TfaValidateAuthCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTfaValidateAuthCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tfaValidateRequestsDto** | [**TfaValidateRequestsDto**](TfaValidateRequestsDto.md) |  | 

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


## UnlinkTfaApp

> StringWrapper UnlinkTfaApp(ctx).TfaRequestsDto(tfaRequestsDto).Execute()

Unlink the TFA application



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/unlink-tfa-app/).

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
	tfaRequestsDto := *openapiclient.NewTfaRequestsDto() // TfaRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsTFASettingsAPI.UnlinkTfaApp(context.Background()).TfaRequestsDto(tfaRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsTFASettingsAPI.UnlinkTfaApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnlinkTfaApp`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsTFASettingsAPI.UnlinkTfaApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkTfaAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tfaRequestsDto** | [**TfaRequestsDto**](TfaRequestsDto.md) |  | 

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


## UpdateTfaAppCodes

> TfaAppCodeArrayWrapper UpdateTfaAppCodes(ctx).Execute()

Regenerate the TFA backup codes



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-tfa-app-codes/).

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
	resp, r, err := apiClient.SettingsTFASettingsAPI.UpdateTfaAppCodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsTFASettingsAPI.UpdateTfaAppCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTfaAppCodes`: TfaAppCodeArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsTFASettingsAPI.UpdateTfaAppCodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTfaAppCodesRequest struct via the builder pattern


### Return type

[**TfaAppCodeArrayWrapper**](TfaAppCodeArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTfaSettings

> BooleanWrapper UpdateTfaSettings(ctx).TfaRequestsDto(tfaRequestsDto).Execute()

Update the TFA settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-tfa-settings/).

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
	tfaRequestsDto := *openapiclient.NewTfaRequestsDto() // TfaRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsTFASettingsAPI.UpdateTfaSettings(context.Background()).TfaRequestsDto(tfaRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsTFASettingsAPI.UpdateTfaSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTfaSettings`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsTFASettingsAPI.UpdateTfaSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTfaSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tfaRequestsDto** | [**TfaRequestsDto**](TfaRequestsDto.md) |  | 

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


## UpdateTfaSettingsLink

> StringWrapper UpdateTfaSettingsLink(ctx).TfaRequestsDto(tfaRequestsDto).Execute()

Update TFA settings with a link



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-tfa-settings-link/).

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
	tfaRequestsDto := *openapiclient.NewTfaRequestsDto() // TfaRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsTFASettingsAPI.UpdateTfaSettingsLink(context.Background()).TfaRequestsDto(tfaRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsTFASettingsAPI.UpdateTfaSettingsLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTfaSettingsLink`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsTFASettingsAPI.UpdateTfaSettingsLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTfaSettingsLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tfaRequestsDto** | [**TfaRequestsDto**](TfaRequestsDto.md) |  | 

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

