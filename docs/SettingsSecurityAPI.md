# \SettingsSecurityAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnabledModules**](SettingsSecurityAPI.md#GetEnabledModules) | **Get** /api/2.0/settings/security/modules | Get the enabled modules
[**GetIsProductAdministrator**](SettingsSecurityAPI.md#GetIsProductAdministrator) | **Get** /api/2.0/settings/security/administrator | Check a product administrator
[**GetPasswordSettings**](SettingsSecurityAPI.md#GetPasswordSettings) | **Get** /api/2.0/settings/security/password | Get the password settings
[**GetProductAdministrators**](SettingsSecurityAPI.md#GetProductAdministrators) | **Get** /api/2.0/settings/security/administrator/{productid} | Get the product administrators
[**GetWebItemSecurityInfo**](SettingsSecurityAPI.md#GetWebItemSecurityInfo) | **Get** /api/2.0/settings/security/{id} | Get the module availability
[**GetWebItemSettingsSecurityInfo**](SettingsSecurityAPI.md#GetWebItemSettingsSecurityInfo) | **Get** /api/2.0/settings/security | Get the security settings
[**SetAccessToWebItems**](SettingsSecurityAPI.md#SetAccessToWebItems) | **Put** /api/2.0/settings/security/access | Set the security settings to modules
[**SetProductAdministrator**](SettingsSecurityAPI.md#SetProductAdministrator) | **Put** /api/2.0/settings/security/administrator | Set a product administrator
[**SetWebItemSecurity**](SettingsSecurityAPI.md#SetWebItemSecurity) | **Put** /api/2.0/settings/security | Set the module security settings
[**UpdatePasswordSettings**](SettingsSecurityAPI.md#UpdatePasswordSettings) | **Put** /api/2.0/settings/security/password | Set the password settings



## GetEnabledModules

> ObjectWrapper GetEnabledModules(ctx).Execute()

Get the enabled modules



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-enabled-modules/).

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
	resp, r, err := apiClient.SettingsSecurityAPI.GetEnabledModules(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsSecurityAPI.GetEnabledModules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnabledModules`: ObjectWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsSecurityAPI.GetEnabledModules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnabledModulesRequest struct via the builder pattern


### Return type

[**ObjectWrapper**](ObjectWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIsProductAdministrator

> ProductAdministratorWrapper GetIsProductAdministrator(ctx).Productid(productid).Userid(userid).Execute()

Check a product administrator



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-is-product-administrator/).

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
	productid := "00000000-0000-0000-0000-000000000000" // string | The ID of the product extracted from the query parameters.
	userid := "00000000-0000-0000-0000-000000000000" // string | The user ID extracted from the query parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsSecurityAPI.GetIsProductAdministrator(context.Background()).Productid(productid).Userid(userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsSecurityAPI.GetIsProductAdministrator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIsProductAdministrator`: ProductAdministratorWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsSecurityAPI.GetIsProductAdministrator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIsProductAdministratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productid** | **string** | The ID of the product extracted from the query parameters. | 
 **userid** | **string** | The user ID extracted from the query parameters. | 

### Return type

[**ProductAdministratorWrapper**](ProductAdministratorWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPasswordSettings

> PasswordSettingsWrapper GetPasswordSettings(ctx).Execute()

Get the password settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-password-settings/).

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
	resp, r, err := apiClient.SettingsSecurityAPI.GetPasswordSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsSecurityAPI.GetPasswordSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPasswordSettings`: PasswordSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsSecurityAPI.GetPasswordSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordSettingsRequest struct via the builder pattern


### Return type

[**PasswordSettingsWrapper**](PasswordSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductAdministrators

> EmployeeArrayWrapper GetProductAdministrators(ctx, productid).Execute()

Get the product administrators



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-product-administrators/).

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
	productid := "00000000-0000-0000-0000-000000000000" // string | The ID of the product extracted from the route parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsSecurityAPI.GetProductAdministrators(context.Background(), productid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsSecurityAPI.GetProductAdministrators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductAdministrators`: EmployeeArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsSecurityAPI.GetProductAdministrators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productid** | **string** | The ID of the product extracted from the route parameters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductAdministratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmployeeArrayWrapper**](EmployeeArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebItemSecurityInfo

> BooleanWrapper GetWebItemSecurityInfo(ctx, id).Execute()

Get the module availability



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-web-item-security-info/).

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
	id := "1" // string | The ID extracted from the route parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsSecurityAPI.GetWebItemSecurityInfo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsSecurityAPI.GetWebItemSecurityInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebItemSecurityInfo`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsSecurityAPI.GetWebItemSecurityInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID extracted from the route parameters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebItemSecurityInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetWebItemSettingsSecurityInfo

> SecurityArrayWrapper GetWebItemSettingsSecurityInfo(ctx).Ids(ids).Execute()

Get the security settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-web-item-settings-security-info/).

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
	ids := []string{"Inner_example"} // []string | The list of module identifiers for which to retrieve the security settings. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsSecurityAPI.GetWebItemSettingsSecurityInfo(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsSecurityAPI.GetWebItemSettingsSecurityInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebItemSettingsSecurityInfo`: SecurityArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsSecurityAPI.GetWebItemSettingsSecurityInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebItemSettingsSecurityInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | The list of module identifiers for which to retrieve the security settings. | 

### Return type

[**SecurityArrayWrapper**](SecurityArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAccessToWebItems

> SecurityArrayWrapper SetAccessToWebItems(ctx).WebItemsSecurityRequestsDto(webItemsSecurityRequestsDto).Execute()

Set the security settings to modules



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-access-to-web-items/).

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
	webItemsSecurityRequestsDto := *openapiclient.NewWebItemsSecurityRequestsDto() // WebItemsSecurityRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsSecurityAPI.SetAccessToWebItems(context.Background()).WebItemsSecurityRequestsDto(webItemsSecurityRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsSecurityAPI.SetAccessToWebItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAccessToWebItems`: SecurityArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsSecurityAPI.SetAccessToWebItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetAccessToWebItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webItemsSecurityRequestsDto** | [**WebItemsSecurityRequestsDto**](WebItemsSecurityRequestsDto.md) |  | 

### Return type

[**SecurityArrayWrapper**](SecurityArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetProductAdministrator

> ProductAdministratorWrapper SetProductAdministrator(ctx).SecurityRequestsDto(securityRequestsDto).Execute()

Set a product administrator



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-product-administrator/).

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
	securityRequestsDto := *openapiclient.NewSecurityRequestsDto("00000000-0000-0000-0000-000000000000", "00000000-0000-0000-0000-000000000000") // SecurityRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsSecurityAPI.SetProductAdministrator(context.Background()).SecurityRequestsDto(securityRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsSecurityAPI.SetProductAdministrator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetProductAdministrator`: ProductAdministratorWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsSecurityAPI.SetProductAdministrator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetProductAdministratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityRequestsDto** | [**SecurityRequestsDto**](SecurityRequestsDto.md) |  | 

### Return type

[**ProductAdministratorWrapper**](ProductAdministratorWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetWebItemSecurity

> SecurityArrayWrapper SetWebItemSecurity(ctx).WebItemSecurityRequestsDto(webItemSecurityRequestsDto).Execute()

Set the module security settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-web-item-security/).

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
	webItemSecurityRequestsDto := *openapiclient.NewWebItemSecurityRequestsDto("00000000-0000-0000-0000-000000000000") // WebItemSecurityRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsSecurityAPI.SetWebItemSecurity(context.Background()).WebItemSecurityRequestsDto(webItemSecurityRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsSecurityAPI.SetWebItemSecurity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetWebItemSecurity`: SecurityArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsSecurityAPI.SetWebItemSecurity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetWebItemSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webItemSecurityRequestsDto** | [**WebItemSecurityRequestsDto**](WebItemSecurityRequestsDto.md) |  | 

### Return type

[**SecurityArrayWrapper**](SecurityArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePasswordSettings

> PasswordSettingsWrapper UpdatePasswordSettings(ctx).PasswordSettingsRequestsDto(passwordSettingsRequestsDto).Execute()

Set the password settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-password-settings/).

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
	passwordSettingsRequestsDto := *openapiclient.NewPasswordSettingsRequestsDto(int32(8)) // PasswordSettingsRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsSecurityAPI.UpdatePasswordSettings(context.Background()).PasswordSettingsRequestsDto(passwordSettingsRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsSecurityAPI.UpdatePasswordSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePasswordSettings`: PasswordSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsSecurityAPI.UpdatePasswordSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordSettingsRequestsDto** | [**PasswordSettingsRequestsDto**](PasswordSettingsRequestsDto.md) |  | 

### Return type

[**PasswordSettingsWrapper**](PasswordSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

