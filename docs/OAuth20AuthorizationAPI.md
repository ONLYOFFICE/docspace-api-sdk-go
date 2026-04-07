# \OAuth20AuthorizationAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizeOAuth**](OAuth20AuthorizationAPI.md#AuthorizeOAuth) | **Get** /oauth2/authorize | OAuth2 Authorization Endpoint
[**ExchangeToken**](OAuth20AuthorizationAPI.md#ExchangeToken) | **Post** /oauth2/token | OAuth2 Token Endpoint
[**SubmitConsent**](OAuth20AuthorizationAPI.md#SubmitConsent) | **Post** /oauth2/authorize | OAuth2 consent endpoint



## AuthorizeOAuth

> AuthorizeOAuth(ctx).ResponseType(responseType).ClientId(clientId).RedirectUri(redirectUri).Scope(scope).Execute()

OAuth2 Authorization Endpoint



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/authorize-o-auth/).

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
	responseType := "code" // string | The OAuth 2.0 response type, must be 'code' for authorization code flow.
	clientId := "6c7cf17b-1bd3-47d5-94c6-be2d3570e168" // string | The client identifier issued to the client during registration.
	redirectUri := "https://example.com" // string | The URL to redirect to after authorization is complete.
	scope := "files:read" // string | The space-separated list of requested scope permissions.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OAuth20AuthorizationAPI.AuthorizeOAuth(context.Background()).ResponseType(responseType).ClientId(clientId).RedirectUri(redirectUri).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20AuthorizationAPI.AuthorizeOAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeOAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **responseType** | **string** | The OAuth 2.0 response type, must be 'code' for authorization code flow. | 
 **clientId** | **string** | The client identifier issued to the client during registration. | 
 **redirectUri** | **string** | The URL to redirect to after authorization is complete. | 
 **scope** | **string** | The space-separated list of requested scope permissions. | 

### Return type

 (empty response body)

### Authorization

[x-signature](../README.md#x-signature)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExchangeToken

> ExchangeToken200Response ExchangeToken(ctx).GrantType(grantType).Code(code).RedirectUri(redirectUri).ClientId(clientId).ClientSecret(clientSecret).Execute()

OAuth2 Token Endpoint



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/exchange-token/).

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
	grantType := "grantType_example" // string | The OAuth2 grant type, must be 'authorization_code' for the authorization code flow. (optional)
	code := "code_example" // string | A temporary authorization code that is sent to the client to be exchanged for a token. (optional)
	redirectUri := "redirectUri_example" // string | The URL where the user will be redirected after successful or unsuccessful authentication. (optional)
	clientId := "clientId_example" // string | The client identifier issued to the client during registration. (optional)
	clientSecret := "clientSecret_example" // string | The client secret issued to the client during registration. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth20AuthorizationAPI.ExchangeToken(context.Background()).GrantType(grantType).Code(code).RedirectUri(redirectUri).ClientId(clientId).ClientSecret(clientSecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20AuthorizationAPI.ExchangeToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExchangeToken`: ExchangeToken200Response
	fmt.Fprintf(os.Stdout, "Response from `OAuth20AuthorizationAPI.ExchangeToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExchangeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** | The OAuth2 grant type, must be 'authorization_code' for the authorization code flow. | 
 **code** | **string** | A temporary authorization code that is sent to the client to be exchanged for a token. | 
 **redirectUri** | **string** | The URL where the user will be redirected after successful or unsuccessful authentication. | 
 **clientId** | **string** | The client identifier issued to the client during registration. | 
 **clientSecret** | **string** | The client secret issued to the client during registration. | 

### Return type

[**ExchangeToken200Response**](ExchangeToken200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitConsent

> SubmitConsent(ctx).ClientId(clientId).State(state).Scope(scope).Execute()

OAuth2 consent endpoint



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/submit-consent/).

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
	clientId := "clientId_example" // string | The client identifier issued to the client during registration. (optional)
	state := "state_example" // string | The random string used to solve the CSRF vulnerability problem. (optional)
	scope := "scope_example" // string | The space-separated list of requested scope permissions. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OAuth20AuthorizationAPI.SubmitConsent(context.Background()).ClientId(clientId).State(state).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20AuthorizationAPI.SubmitConsent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitConsentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | The client identifier issued to the client during registration. | 
 **state** | **string** | The random string used to solve the CSRF vulnerability problem. | 
 **scope** | **string** | The space-separated list of requested scope permissions. | 

### Return type

 (empty response body)

### Authorization

[x-signature](../README.md#x-signature)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

