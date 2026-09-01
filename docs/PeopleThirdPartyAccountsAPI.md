# \PeopleThirdPartyAccountsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetThirdPartyAuthProviders**](PeopleThirdPartyAccountsAPI.md#GetThirdPartyAuthProviders) | **Get** /api/2.0/people/thirdparty/providers | Get third-party accounts
[**LinkThirdPartyAccount**](PeopleThirdPartyAccountsAPI.md#LinkThirdPartyAccount) | **Put** /api/2.0/people/thirdparty/linkaccount | Link a third-pary account
[**SignupThirdPartyAccount**](PeopleThirdPartyAccountsAPI.md#SignupThirdPartyAccount) | **Post** /api/2.0/people/thirdparty/signup | Create a third-pary account
[**UnlinkThirdPartyAccount**](PeopleThirdPartyAccountsAPI.md#UnlinkThirdPartyAccount) | **Delete** /api/2.0/people/thirdparty/unlinkaccount | Unlink a third-pary account



## GetThirdPartyAuthProviders

> AccountInfoArrayWrapper GetThirdPartyAuthProviders(ctx).InviteView(inviteView).SettingsView(settingsView).ClientCallback(clientCallback).FromOnly(fromOnly).Execute()

Get third-party accounts



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-third-party-auth-providers/).

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
	inviteView := false // bool | Specifies whether to return providers that are available for invitation links, i.e. the user can login or register through these providers. (optional)
	settingsView := false // bool | Specifies whether to display the provider settings in a pop-up window (true) or redirect them to the desktop application (false). (optional)
	clientCallback := "onAuthCallback" // string | The method that is called after authentication. (optional)
	fromOnly := "Google" // string | The provider name if a response is required only from this provider. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleThirdPartyAccountsAPI.GetThirdPartyAuthProviders(context.Background()).InviteView(inviteView).SettingsView(settingsView).ClientCallback(clientCallback).FromOnly(fromOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleThirdPartyAccountsAPI.GetThirdPartyAuthProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThirdPartyAuthProviders`: AccountInfoArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleThirdPartyAccountsAPI.GetThirdPartyAuthProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetThirdPartyAuthProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inviteView** | **bool** | Specifies whether to return providers that are available for invitation links, i.e. the user can login or register through these providers. | 
 **settingsView** | **bool** | Specifies whether to display the provider settings in a pop-up window (true) or redirect them to the desktop application (false). | 
 **clientCallback** | **string** | The method that is called after authentication. | 
 **fromOnly** | **string** | The provider name if a response is required only from this provider. | 

### Return type

[**AccountInfoArrayWrapper**](AccountInfoArrayWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkThirdPartyAccount

> LinkThirdPartyAccount(ctx).LinkAccountRequestDto(linkAccountRequestDto).Execute()

Link a third-pary account



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/link-third-party-account/).

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
	linkAccountRequestDto := *openapiclient.NewLinkAccountRequestDto() // LinkAccountRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PeopleThirdPartyAccountsAPI.LinkThirdPartyAccount(context.Background()).LinkAccountRequestDto(linkAccountRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleThirdPartyAccountsAPI.LinkThirdPartyAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkThirdPartyAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkAccountRequestDto** | [**LinkAccountRequestDto**](LinkAccountRequestDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignupThirdPartyAccount

> EmployeeWrapper SignupThirdPartyAccount(ctx).SignupAccountRequestDto(signupAccountRequestDto).Execute()

Create a third-pary account



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/signup-third-party-account/).

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
	signupAccountRequestDto := *openapiclient.NewSignupAccountRequestDto("invite_key_123456", "{\"provider\":\"Google\",\"id\":\"123456\"}") // SignupAccountRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleThirdPartyAccountsAPI.SignupThirdPartyAccount(context.Background()).SignupAccountRequestDto(signupAccountRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleThirdPartyAccountsAPI.SignupThirdPartyAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignupThirdPartyAccount`: EmployeeWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleThirdPartyAccountsAPI.SignupThirdPartyAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignupThirdPartyAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signupAccountRequestDto** | [**SignupAccountRequestDto**](SignupAccountRequestDto.md) |  | 

### Return type

[**EmployeeWrapper**](EmployeeWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkThirdPartyAccount

> UnlinkThirdPartyAccount(ctx).Provider(provider).Execute()

Unlink a third-pary account



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/unlink-third-party-account/).

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
	provider := "Google" // string | The provider name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PeopleThirdPartyAccountsAPI.UnlinkThirdPartyAccount(context.Background()).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleThirdPartyAccountsAPI.UnlinkThirdPartyAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkThirdPartyAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | The provider name. | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

