# \OAuth20AuthorizationAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizeOAuth**](OAuth20AuthorizationAPI.md#AuthorizeOAuth) | **Get** /oauth2/authorize | Start the authorization flow
[**ExchangeToken**](OAuth20AuthorizationAPI.md#ExchangeToken) | **Post** /oauth2/token | Exchange the authorization code
[**SubmitConsent**](OAuth20AuthorizationAPI.md#SubmitConsent) | **Post** /oauth2/authorize | Submit the consent decision



## AuthorizeOAuth

> AuthorizeOAuth(ctx).ResponseType(responseType).ClientId(clientId).RedirectUri(redirectUri).Scope(scope).Execute()

Start the authorization flow



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
	responseType := "code" // string | The OAuth 2.0 response type. Only code is supported: this server issues an authorization code, never a token, from this endpoint.
	clientId := "6c7cf17b-1bd3-47d5-94c6-be2d3570e168" // string | The identifier the client was given when it was registered. It selects both the client shown on the consent screen and the set of redirect URIs the request is checked against.
	redirectUri := "https://example.com" // string | Where to send the user once authorization is complete. It has to be one of the redirect URIs registered for the client, otherwise the request is refused.
	scope := "files:read" // string | The permissions being asked for, as a space-separated list. Every scope has to be one the client is registered for, and the consent screen lists exactly these.

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
 **responseType** | **string** | The OAuth 2.0 response type. Only code is supported: this server issues an authorization code, never a token, from this endpoint. | 
 **clientId** | **string** | The identifier the client was given when it was registered. It selects both the client shown on the consent screen and the set of redirect URIs the request is checked against. | 
 **redirectUri** | **string** | Where to send the user once authorization is complete. It has to be one of the redirect URIs registered for the client, otherwise the request is refused. | 
 **scope** | **string** | The permissions being asked for, as a space-separated list. Every scope has to be one the client is registered for, and the consent screen lists exactly these. | 

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

Exchange the authorization code



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
	grantType := "grantType_example" // string | Which exchange is being performed: authorization_code to redeem a code, refresh_token to renew an access token. (optional)
	code := "code_example" // string | The authorization code returned by the authorization endpoint. It may be redeemed once. (optional)
	redirectUri := "redirectUri_example" // string | The same redirect URI that was used to obtain the code. The exchange fails when it differs. (optional)
	clientId := "clientId_example" // string | The identifier of the client redeeming the code. (optional)
	clientSecret := "clientSecret_example" // string | The secret of the client redeeming the code. It is omitted by a public client, which proves itself with a PKCE code verifier instead. (optional)

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
 **grantType** | **string** | Which exchange is being performed: authorization_code to redeem a code, refresh_token to renew an access token. | 
 **code** | **string** | The authorization code returned by the authorization endpoint. It may be redeemed once. | 
 **redirectUri** | **string** | The same redirect URI that was used to obtain the code. The exchange fails when it differs. | 
 **clientId** | **string** | The identifier of the client redeeming the code. | 
 **clientSecret** | **string** | The secret of the client redeeming the code. It is omitted by a public client, which proves itself with a PKCE code verifier instead. | 

### Return type

[**ExchangeToken200Response**](ExchangeToken200Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitConsent

> SubmitConsent(ctx).ClientId(clientId).State(state).Scope(scope).Execute()

Submit the consent decision



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
	clientId := "clientId_example" // string | The client the consent is being given to. It has to be the same client the authorization request named. (optional)
	state := "state_example" // string | The opaque value carried through from the authorization request, returned unchanged on the redirect so the client can match the answer to its request. (optional)
	scope := "scope_example" // string | The scopes the user agreed to, as a space-separated list. Anything the user declined is left out, so this may be narrower than what was requested. (optional)

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
 **clientId** | **string** | The client the consent is being given to. It has to be the same client the authorization request named. | 
 **state** | **string** | The opaque value carried through from the authorization request, returned unchanged on the redirect so the client can match the answer to its request. | 
 **scope** | **string** | The scopes the user agreed to, as a space-separated list. Anything the user declined is left out, so this may be narrower than what was requested. | 

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

