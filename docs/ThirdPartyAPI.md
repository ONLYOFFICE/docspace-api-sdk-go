# \ThirdPartyAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetThirdPartyCode**](ThirdPartyAPI.md#GetThirdPartyCode) | **Get** /api/2.0/thirdparty/{provider} | Get provider consent URL



## GetThirdPartyCode

> StringWrapper GetThirdPartyCode(ctx, provider).Execute()

Get provider consent URL



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-third-party-code/).

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
	provider := openapiclient.LoginProvider(0) // LoginProvider | The provider whose consent screen is wanted. Only Google, Dropbox, Docusign, Box, OneDrive, Wordpress and  Github produce a URL; any other provider is answered with 200 and no URL rather than an error. The provider  credentials have to be saved with `POST api/2.0/settings/authservice` first, or the URL comes back without a  client identifier and the provider refuses it.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThirdPartyAPI.GetThirdPartyCode(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThirdPartyAPI.GetThirdPartyCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThirdPartyCode`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `ThirdPartyAPI.GetThirdPartyCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | [**LoginProvider**](.md) | The provider whose consent screen is wanted. Only Google, Dropbox, Docusign, Box, OneDrive, Wordpress and  Github produce a URL; any other provider is answered with 200 and no URL rather than an error. The provider  credentials have to be saved with `POST api/2.0/settings/authservice` first, or the URL comes back without a  client identifier and the provider refuses it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThirdPartyCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

