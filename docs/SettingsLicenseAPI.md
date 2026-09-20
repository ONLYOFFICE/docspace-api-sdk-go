# \SettingsLicenseAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptLicense**](SettingsLicenseAPI.md#AcceptLicense) | **Post** /api/2.0/settings/license/accept | Activate a license
[**GetIsLicenseRequired**](SettingsLicenseAPI.md#GetIsLicenseRequired) | **Get** /api/2.0/settings/license/required | Check if a license is required
[**RefreshLicense**](SettingsLicenseAPI.md#RefreshLicense) | **Get** /api/2.0/settings/license/refresh | Refresh the license
[**UploadLicense**](SettingsLicenseAPI.md#UploadLicense) | **Post** /api/2.0/settings/license | Upload a license



## AcceptLicense

> StringWrapper AcceptLicense(ctx).Execute()

Activate a license



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/accept-license/).

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
	resp, r, err := apiClient.SettingsLicenseAPI.AcceptLicense(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsLicenseAPI.AcceptLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcceptLicense`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsLicenseAPI.AcceptLicense`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptLicenseRequest struct via the builder pattern


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


## GetIsLicenseRequired

> BooleanWrapper GetIsLicenseRequired(ctx).Execute()

Check if a license is required



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-is-license-required/).

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
	resp, r, err := apiClient.SettingsLicenseAPI.GetIsLicenseRequired(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsLicenseAPI.GetIsLicenseRequired``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIsLicenseRequired`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsLicenseAPI.GetIsLicenseRequired`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIsLicenseRequiredRequest struct via the builder pattern


### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshLicense

> BooleanWrapper RefreshLicense(ctx).Execute()

Refresh the license



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/refresh-license/).

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
	resp, r, err := apiClient.SettingsLicenseAPI.RefreshLicense(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsLicenseAPI.RefreshLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshLicense`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsLicenseAPI.RefreshLicense`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshLicenseRequest struct via the builder pattern


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


## UploadLicense

> StringWrapper UploadLicense(ctx).Files(files).Execute()

Upload a license



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/upload-license/).

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
	files := []*os.File{"TODO"} // []*os.File | The license file, sent as `multipart/form-data`. Only the first entry is read and the rest are ignored, and a  request carrying none is refused with 400. A file that cannot be read as a license, that carries no customer  id or signature, or that was issued for the other edition fails the call; one whose start date has not  arrived yet is refused, while one already past its due date is still accepted. Staging only stores the file -  `POST api/2.0/settings/license/accept` puts it in force - and a file staged earlier is overwritten.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsLicenseAPI.UploadLicense(context.Background()).Files(files).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsLicenseAPI.UploadLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadLicense`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsLicenseAPI.UploadLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **files** | **[]*os.File** | The license file, sent as `multipart/form-data`. Only the first entry is read and the rest are ignored, and a  request carrying none is refused with 400. A file that cannot be read as a license, that carries no customer  id or signature, or that was issued for the other edition fails the call; one whose start date has not  arrived yet is refused, while one already past its due date is still accepted. Staging only stores the file -  `POST api/2.0/settings/license/accept` puts it in force - and a file staged earlier is overwritten. | 

### Return type

[**StringWrapper**](StringWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

