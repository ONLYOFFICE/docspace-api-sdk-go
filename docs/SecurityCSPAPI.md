# \SecurityCSPAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigureCsp**](SecurityCSPAPI.md#ConfigureCsp) | **Post** /api/2.0/security/csp | Configure CSP settings
[**GetCspSettings**](SecurityCSPAPI.md#GetCspSettings) | **Get** /api/2.0/security/csp | Get CSP settings



## ConfigureCsp

> CspWrapper ConfigureCsp(ctx).CspRequestsDto(cspRequestsDto).Execute()

Configure CSP settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/configure-csp/).

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
	cspRequestsDto := *openapiclient.NewCspRequestsDto() // CspRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityCSPAPI.ConfigureCsp(context.Background()).CspRequestsDto(cspRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityCSPAPI.ConfigureCsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigureCsp`: CspWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityCSPAPI.ConfigureCsp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigureCspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cspRequestsDto** | [**CspRequestsDto**](CspRequestsDto.md) |  | 

### Return type

[**CspWrapper**](CspWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCspSettings

> CspWrapper GetCspSettings(ctx).Execute()

Get CSP settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-csp-settings/).

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
	resp, r, err := apiClient.SecurityCSPAPI.GetCspSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityCSPAPI.GetCspSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCspSettings`: CspWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityCSPAPI.GetCspSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCspSettingsRequest struct via the builder pattern


### Return type

[**CspWrapper**](CspWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

