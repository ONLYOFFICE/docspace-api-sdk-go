# \PeopleThirdPartyAccountsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetThirdPartyAuthProviders**](PeopleThirdPartyAccountsAPI.md#GetThirdPartyAuthProviders) | **Get** /api/2.0/people/thirdparty/providers | Get third-party providers
[**LinkThirdPartyAccount**](PeopleThirdPartyAccountsAPI.md#LinkThirdPartyAccount) | **Put** /api/2.0/people/thirdparty/linkaccount | Link a third-party account
[**SignupThirdPartyAccount**](PeopleThirdPartyAccountsAPI.md#SignupThirdPartyAccount) | **Post** /api/2.0/people/thirdparty/signup | Sign up with a provider
[**UnlinkThirdPartyAccount**](PeopleThirdPartyAccountsAPI.md#UnlinkThirdPartyAccount) | **Delete** /api/2.0/people/thirdparty/unlinkaccount | Unlink a third-party account



## GetThirdPartyAuthProviders

> AccountInfoArrayWrapper GetThirdPartyAuthProviders(ctx).InviteView(inviteView).SettingsView(settingsView).ClientCallback(clientCallback).FromOnly(fromOnly).Execute()

Get third-party providers



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
	inviteView := false // bool | Set it to true when the list is rendered on an invitation page: the providers that cannot be used to accept an  invitation, `twitter` and `appleid`, are then left out. It defaults to false, which returns every enabled  provider. (optional)
	settingsView := false // bool | Set it to true when the list is rendered on a settings page, to get login URLs that open in a popup window.  With the default false the URL still opens in a popup for a desktop browser, and switches to a redirect only  for a mobile browser or for the DocSpace desktop application. (optional)
	clientCallback := "onAuthCallback" // string | The name of the client-side function the popup calls back when the provider authorization finishes. It is  placed into the returned URLs as they are, and it is only used by the popup mode. (optional)
	fromOnly := "google" // string | Keeps only the named provider, compared case-insensitively against the lowercase provider names such as  `google` or `microsoft`; the special value `openid` selects `google`. Omit it to get every enabled provider. (optional)

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
 **inviteView** | **bool** | Set it to true when the list is rendered on an invitation page: the providers that cannot be used to accept an  invitation, `twitter` and `appleid`, are then left out. It defaults to false, which returns every enabled  provider. | 
 **settingsView** | **bool** | Set it to true when the list is rendered on a settings page, to get login URLs that open in a popup window.  With the default false the URL still opens in a popup for a desktop browser, and switches to a redirect only  for a mobile browser or for the DocSpace desktop application. | 
 **clientCallback** | **string** | The name of the client-side function the popup calls back when the provider authorization finishes. It is  placed into the returned URLs as they are, and it is only used by the popup mode. | 
 **fromOnly** | **string** | Keeps only the named provider, compared case-insensitively against the lowercase provider names such as  `google` or `microsoft`; the special value `openid` selects `google`. Omit it to get every enabled provider. | 

### Return type

[**AccountInfoArrayWrapper**](AccountInfoArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkThirdPartyAccount

> LinkThirdPartyAccount(ctx).LinkAccountRequestDto(linkAccountRequestDto).Execute()

Link a third-party account



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

Sign up with a provider



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
	signupAccountRequestDto := *openapiclient.NewSignupAccountRequestDto("invite_key_123456", "{\"provider\":\"google\",\"id\":\"123456\"}") // SignupAccountRequestDto |  (optional)

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

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkThirdPartyAccount

> UnlinkThirdPartyAccount(ctx).Provider(provider).Execute()

Unlink a third-party account



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
	provider := "google" // string | The name of the provider to unlink, in the lowercase form `GET api/2.0/people/thirdparty/providers` returns,  such as `google` or `microsoft`. A name that is not linked to the calling profile is accepted and changes  nothing. (optional)

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
 **provider** | **string** | The name of the provider to unlink, in the lowercase form `GET api/2.0/people/thirdparty/providers` returns,  such as `google` or `microsoft`. A name that is not linked to the calling profile is accepted and changes  nothing. | 

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

