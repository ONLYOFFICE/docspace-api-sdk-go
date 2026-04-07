# \AIProvidersAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProvider**](AIProvidersAPI.md#AddProvider) | **Post** /api/2.0/ai/providers | Add an AI provider
[**DeleteProviders**](AIProvidersAPI.md#DeleteProviders) | **Delete** /api/2.0/ai/providers | Delete AI providers
[**GetAvailableProviders**](AIProvidersAPI.md#GetAvailableProviders) | **Get** /api/2.0/ai/providers/available | Get available AI provider types
[**GetDefaultProvider**](AIProvidersAPI.md#GetDefaultProvider) | **Get** /api/2.0/ai/providers/default | Get the default AI provider
[**GetProviders**](AIProvidersAPI.md#GetProviders) | **Get** /api/2.0/ai/providers | Get AI providers
[**SetDefaultProvider**](AIProvidersAPI.md#SetDefaultProvider) | **Put** /api/2.0/ai/providers/default | Set the default AI provider
[**UpdateProvider**](AIProvidersAPI.md#UpdateProvider) | **Put** /api/2.0/ai/providers/{id} | Update an AI provider



## AddProvider

> AiProviderWrapper AddProvider(ctx).CreateProviderRequestDto(createProviderRequestDto).Execute()

Add an AI provider



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/add-provider/).

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
	createProviderRequestDto := *openapiclient.NewCreateProviderRequestDto("OpenAI Provider", "sk-example-key-123") // CreateProviderRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIProvidersAPI.AddProvider(context.Background()).CreateProviderRequestDto(createProviderRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIProvidersAPI.AddProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddProvider`: AiProviderWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIProvidersAPI.AddProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProviderRequestDto** | [**CreateProviderRequestDto**](CreateProviderRequestDto.md) |  | 

### Return type

[**AiProviderWrapper**](AiProviderWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProviders

> DeleteProviders(ctx).RemoveProviderRequestDto(removeProviderRequestDto).Execute()

Delete AI providers



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-providers/).

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
	removeProviderRequestDto := *openapiclient.NewRemoveProviderRequestDto([]int32{int32(123)}) // RemoveProviderRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIProvidersAPI.DeleteProviders(context.Background()).RemoveProviderRequestDto(removeProviderRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIProvidersAPI.DeleteProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeProviderRequestDto** | [**RemoveProviderRequestDto**](RemoveProviderRequestDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableProviders

> ProviderSettingsArrayWrapper GetAvailableProviders(ctx).Execute()

Get available AI provider types



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-available-providers/).

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
	resp, r, err := apiClient.AIProvidersAPI.GetAvailableProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIProvidersAPI.GetAvailableProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableProviders`: ProviderSettingsArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIProvidersAPI.GetAvailableProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableProvidersRequest struct via the builder pattern


### Return type

[**ProviderSettingsArrayWrapper**](ProviderSettingsArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultProvider

> DefaultProviderWrapper GetDefaultProvider(ctx).Execute()

Get the default AI provider



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-default-provider/).

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
	resp, r, err := apiClient.AIProvidersAPI.GetDefaultProvider(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIProvidersAPI.GetDefaultProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultProvider`: DefaultProviderWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIProvidersAPI.GetDefaultProvider`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultProviderRequest struct via the builder pattern


### Return type

[**DefaultProviderWrapper**](DefaultProviderWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProviders

> AiProviderArrayWrapper GetProviders(ctx).StartIndex(startIndex).Count(count).Execute()

Get AI providers



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-providers/).

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
	startIndex := int32(0) // int32 | The number of items to skip before returning results (zero-based offset). Defaults to 0. (optional)
	count := int32(100) // int32 | The maximum number of items to return per page. Defaults to 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIProvidersAPI.GetProviders(context.Background()).StartIndex(startIndex).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIProvidersAPI.GetProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProviders`: AiProviderArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIProvidersAPI.GetProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startIndex** | **int32** | The number of items to skip before returning results (zero-based offset). Defaults to 0. | 
 **count** | **int32** | The maximum number of items to return per page. Defaults to 100. | 

### Return type

[**AiProviderArrayWrapper**](AiProviderArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultProvider

> DefaultProviderWrapper SetDefaultProvider(ctx).SetDefaultProviderRequestDto(setDefaultProviderRequestDto).Execute()

Set the default AI provider



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-default-provider/).

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
	setDefaultProviderRequestDto := *openapiclient.NewSetDefaultProviderRequestDto("gpt-4") // SetDefaultProviderRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIProvidersAPI.SetDefaultProvider(context.Background()).SetDefaultProviderRequestDto(setDefaultProviderRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIProvidersAPI.SetDefaultProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDefaultProvider`: DefaultProviderWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIProvidersAPI.SetDefaultProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setDefaultProviderRequestDto** | [**SetDefaultProviderRequestDto**](SetDefaultProviderRequestDto.md) |  | 

### Return type

[**DefaultProviderWrapper**](DefaultProviderWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProvider

> AiProviderWrapper UpdateProvider(ctx, id).UpdateProviderBody(updateProviderBody).Execute()

Update an AI provider



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-provider/).

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
	id := int32(1) // int32 | The identifier of the AI provider to update.
	updateProviderBody := *openapiclient.NewUpdateProviderBody() // UpdateProviderBody | The AI provider configuration parameters to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIProvidersAPI.UpdateProvider(context.Background(), id).UpdateProviderBody(updateProviderBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIProvidersAPI.UpdateProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProvider`: AiProviderWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIProvidersAPI.UpdateProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The identifier of the AI provider to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProviderBody** | [**UpdateProviderBody**](UpdateProviderBody.md) | The AI provider configuration parameters to update. | 

### Return type

[**AiProviderWrapper**](AiProviderWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

