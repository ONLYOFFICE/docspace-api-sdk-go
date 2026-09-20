# \AIWebSearchAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiWebSearchClear**](AIWebSearchAPI.md#AiWebSearchClear) | **Delete** /api/2.0/ai/web-search/clear | Clear the web-search configuration
[**AiWebSearchConfigure**](AIWebSearchAPI.md#AiWebSearchConfigure) | **Put** /api/2.0/ai/web-search/configure | Configure and verify web search
[**AiWebSearchGetActiveConfig**](AIWebSearchAPI.md#AiWebSearchGetActiveConfig) | **Get** /api/2.0/ai/web-search/get-active-config | Get active config
[**AiWebSearchIsConfigured**](AIWebSearchAPI.md#AiWebSearchIsConfigured) | **Get** /api/2.0/ai/web-search/is-configured | Is configured
[**AiWebSearchPassthroughContents**](AIWebSearchAPI.md#AiWebSearchPassthroughContents) | **Post** /api/2.0/ai/websearch/v1/contents | Web page contents passthrough
[**AiWebSearchPassthroughSearch**](AIWebSearchAPI.md#AiWebSearchPassthroughSearch) | **Post** /api/2.0/ai/websearch/v1/search | Web search passthrough
[**AiWebSearchSetActiveConfig**](AIWebSearchAPI.md#AiWebSearchSetActiveConfig) | **Put** /api/2.0/ai/web-search/set-active-config | Set active config
[**AiWebSearchTestConnection**](AIWebSearchAPI.md#AiWebSearchTestConnection) | **Post** /api/2.0/ai/web-search/test-connection | Test a web-search provider



## AiWebSearchClear

> AiSuccessResponse AiWebSearchClear(ctx).Body(body).Execute()

Clear the web-search configuration



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-web-search-clear/).

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
	body := "body_example" // string | Ignored. The operation always clears the portal-wide configuration, so send an empty body; a value here does not scope it to a room.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIWebSearchAPI.AiWebSearchClear(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIWebSearchAPI.AiWebSearchClear``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiWebSearchClear`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIWebSearchAPI.AiWebSearchClear`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiWebSearchClearRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | Ignored. The operation always clears the portal-wide configuration, so send an empty body; a value here does not scope it to a room. | 

### Return type

[**AiSuccessResponse**](AiSuccessResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiWebSearchConfigure

> AiWebSearchMutationResult AiWebSearchConfigure(ctx).AiWebSearchConfigureRequest(aiWebSearchConfigureRequest).Execute()

Configure and verify web search



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-web-search-configure/).

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
	aiWebSearchConfigureRequest := *openapiclient.NewAiWebSearchConfigureRequest(*openapiclient.NewAiWebSearchConfig("exa")) // AiWebSearchConfigureRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIWebSearchAPI.AiWebSearchConfigure(context.Background()).AiWebSearchConfigureRequest(aiWebSearchConfigureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIWebSearchAPI.AiWebSearchConfigure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiWebSearchConfigure`: AiWebSearchMutationResult
	fmt.Fprintf(os.Stdout, "Response from `AIWebSearchAPI.AiWebSearchConfigure`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiWebSearchConfigureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiWebSearchConfigureRequest** | [**AiWebSearchConfigureRequest**](AiWebSearchConfigureRequest.md) |  | 

### Return type

[**AiWebSearchMutationResult**](AiWebSearchMutationResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiWebSearchGetActiveConfig

> AiWebSearchConfig AiWebSearchGetActiveConfig(ctx).EntityId(entityId).Execute()

Get active config



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-web-search-get-active-config/).

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
	entityId := "1234" // string | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIWebSearchAPI.AiWebSearchGetActiveConfig(context.Background()).EntityId(entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIWebSearchAPI.AiWebSearchGetActiveConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiWebSearchGetActiveConfig`: AiWebSearchConfig
	fmt.Fprintf(os.Stdout, "Response from `AIWebSearchAPI.AiWebSearchGetActiveConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiWebSearchGetActiveConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string** | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. | 

### Return type

[**AiWebSearchConfig**](AiWebSearchConfig.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiWebSearchIsConfigured

> bool AiWebSearchIsConfigured(ctx).EntityId(entityId).Execute()

Is configured



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-web-search-is-configured/).

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
	entityId := "1234" // string | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIWebSearchAPI.AiWebSearchIsConfigured(context.Background()).EntityId(entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIWebSearchAPI.AiWebSearchIsConfigured``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiWebSearchIsConfigured`: bool
	fmt.Fprintf(os.Stdout, "Response from `AIWebSearchAPI.AiWebSearchIsConfigured`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiWebSearchIsConfiguredRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string** | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. | 

### Return type

**bool**

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiWebSearchPassthroughContents

> map[string]interface{} AiWebSearchPassthroughContents(ctx).RequestBody(requestBody).Execute()

Web page contents passthrough



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-web-search-passthrough-contents/).

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
	requestBody := map[string]*interface{}{"key": interface{}(123)} // map[string]*interface{} | A page-contents request in the shape the portal's active web-search provider expects, forwarded to it unchanged. The endpoint and the key come from the stored configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIWebSearchAPI.AiWebSearchPassthroughContents(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIWebSearchAPI.AiWebSearchPassthroughContents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiWebSearchPassthroughContents`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIWebSearchAPI.AiWebSearchPassthroughContents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiWebSearchPassthroughContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** | A page-contents request in the shape the portal's active web-search provider expects, forwarded to it unchanged. The endpoint and the key come from the stored configuration. | 

### Return type

**map[string]interface{}**

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiWebSearchPassthroughSearch

> map[string]interface{} AiWebSearchPassthroughSearch(ctx).RequestBody(requestBody).Execute()

Web search passthrough



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-web-search-passthrough-search/).

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
	requestBody := map[string]*interface{}{"key": interface{}(123)} // map[string]*interface{} | A search request in the shape the portal's active web-search provider expects, forwarded to it unchanged. The endpoint and the key come from the stored configuration and must not be sent here.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIWebSearchAPI.AiWebSearchPassthroughSearch(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIWebSearchAPI.AiWebSearchPassthroughSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiWebSearchPassthroughSearch`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIWebSearchAPI.AiWebSearchPassthroughSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiWebSearchPassthroughSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** | A search request in the shape the portal's active web-search provider expects, forwarded to it unchanged. The endpoint and the key come from the stored configuration and must not be sent here. | 

### Return type

**map[string]interface{}**

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiWebSearchSetActiveConfig

> AiSuccessResponse AiWebSearchSetActiveConfig(ctx).AiWebSearchConfigureRequest(aiWebSearchConfigureRequest).Execute()

Set active config



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-web-search-set-active-config/).

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
	aiWebSearchConfigureRequest := *openapiclient.NewAiWebSearchConfigureRequest(*openapiclient.NewAiWebSearchConfig("exa")) // AiWebSearchConfigureRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIWebSearchAPI.AiWebSearchSetActiveConfig(context.Background()).AiWebSearchConfigureRequest(aiWebSearchConfigureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIWebSearchAPI.AiWebSearchSetActiveConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiWebSearchSetActiveConfig`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIWebSearchAPI.AiWebSearchSetActiveConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiWebSearchSetActiveConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiWebSearchConfigureRequest** | [**AiWebSearchConfigureRequest**](AiWebSearchConfigureRequest.md) |  | 

### Return type

[**AiSuccessResponse**](AiSuccessResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiWebSearchTestConnection

> AiProfilesTestConnection200Response AiWebSearchTestConnection(ctx).AiWebSearchConfig(aiWebSearchConfig).Execute()

Test a web-search provider



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-web-search-test-connection/).

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
	aiWebSearchConfig := *openapiclient.NewAiWebSearchConfig("exa") // AiWebSearchConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIWebSearchAPI.AiWebSearchTestConnection(context.Background()).AiWebSearchConfig(aiWebSearchConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIWebSearchAPI.AiWebSearchTestConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiWebSearchTestConnection`: AiProfilesTestConnection200Response
	fmt.Fprintf(os.Stdout, "Response from `AIWebSearchAPI.AiWebSearchTestConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiWebSearchTestConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiWebSearchConfig** | [**AiWebSearchConfig**](AiWebSearchConfig.md) |  | 

### Return type

[**AiProfilesTestConnection200Response**](AiProfilesTestConnection200Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

