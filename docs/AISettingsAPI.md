# \AISettingsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAiSettings**](AISettingsAPI.md#GetAiSettings) | **Get** /api/2.0/ai/config | Get AI settings
[**GetVectorizationSettings**](AISettingsAPI.md#GetVectorizationSettings) | **Get** /api/2.0/ai/config/vectorization | Get vectorization settings
[**GetWebSearchSettings**](AISettingsAPI.md#GetWebSearchSettings) | **Get** /api/2.0/ai/config/web-search | Get web search settings
[**SetVectorizationSettings**](AISettingsAPI.md#SetVectorizationSettings) | **Put** /api/2.0/ai/config/vectorization | Update vectorization settings
[**SetWebSearchSettings**](AISettingsAPI.md#SetWebSearchSettings) | **Put** /api/2.0/ai/config/web-search | Update web search settings



## GetAiSettings

> AiSettingsWrapper GetAiSettings(ctx).Execute()

Get AI settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-ai-settings/).

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
	resp, r, err := apiClient.AISettingsAPI.GetAiSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISettingsAPI.GetAiSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiSettings`: AiSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `AISettingsAPI.GetAiSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiSettingsRequest struct via the builder pattern


### Return type

[**AiSettingsWrapper**](AiSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVectorizationSettings

> VectorizationSettingsWrapper GetVectorizationSettings(ctx).Execute()

Get vectorization settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-vectorization-settings/).

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
	resp, r, err := apiClient.AISettingsAPI.GetVectorizationSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISettingsAPI.GetVectorizationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVectorizationSettings`: VectorizationSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `AISettingsAPI.GetVectorizationSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVectorizationSettingsRequest struct via the builder pattern


### Return type

[**VectorizationSettingsWrapper**](VectorizationSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebSearchSettings

> WebSearchSettingsWrapper GetWebSearchSettings(ctx).Execute()

Get web search settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-web-search-settings/).

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
	resp, r, err := apiClient.AISettingsAPI.GetWebSearchSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISettingsAPI.GetWebSearchSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebSearchSettings`: WebSearchSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `AISettingsAPI.GetWebSearchSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebSearchSettingsRequest struct via the builder pattern


### Return type

[**WebSearchSettingsWrapper**](WebSearchSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVectorizationSettings

> VectorizationSettingsWrapper SetVectorizationSettings(ctx).SetEmbeddingConfigRequestBody(setEmbeddingConfigRequestBody).Execute()

Update vectorization settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-vectorization-settings/).

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
	setEmbeddingConfigRequestBody := *openapiclient.NewSetEmbeddingConfigRequestBody() // SetEmbeddingConfigRequestBody | The embedding provider configuration parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISettingsAPI.SetVectorizationSettings(context.Background()).SetEmbeddingConfigRequestBody(setEmbeddingConfigRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISettingsAPI.SetVectorizationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetVectorizationSettings`: VectorizationSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `AISettingsAPI.SetVectorizationSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetVectorizationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setEmbeddingConfigRequestBody** | [**SetEmbeddingConfigRequestBody**](SetEmbeddingConfigRequestBody.md) | The embedding provider configuration parameters. | 

### Return type

[**VectorizationSettingsWrapper**](VectorizationSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetWebSearchSettings

> WebSearchSettingsWrapper SetWebSearchSettings(ctx).SetWebSearchSettingsRequestBody(setWebSearchSettingsRequestBody).Execute()

Update web search settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-web-search-settings/).

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
	setWebSearchSettingsRequestBody := *openapiclient.NewSetWebSearchSettingsRequestBody() // SetWebSearchSettingsRequestBody | The web search configuration parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISettingsAPI.SetWebSearchSettings(context.Background()).SetWebSearchSettingsRequestBody(setWebSearchSettingsRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISettingsAPI.SetWebSearchSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetWebSearchSettings`: WebSearchSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `AISettingsAPI.SetWebSearchSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetWebSearchSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setWebSearchSettingsRequestBody** | [**SetWebSearchSettingsRequestBody**](SetWebSearchSettingsRequestBody.md) | The web search configuration parameters. | 

### Return type

[**WebSearchSettingsWrapper**](WebSearchSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

