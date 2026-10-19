# \AISettingsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiSettingsGet**](AISettingsAPI.md#AiSettingsGet) | **Get** /api/2.0/ai/config | Get AI settings
[**AiSettingsGetToolMode**](AISettingsAPI.md#AiSettingsGetToolMode) | **Get** /api/2.0/ai/config/tool-mode | Get the tool permission mode
[**AiSettingsGetUser**](AISettingsAPI.md#AiSettingsGetUser) | **Get** /api/2.0/ai/config/user | Get user AI settings
[**AiSettingsGetVectorization**](AISettingsAPI.md#AiSettingsGetVectorization) | **Get** /api/2.0/ai/config/vectorization | Get vectorization settings
[**AiSettingsSetToolMode**](AISettingsAPI.md#AiSettingsSetToolMode) | **Put** /api/2.0/ai/config/tool-mode | Set the tool permission mode
[**AiSettingsSetUser**](AISettingsAPI.md#AiSettingsSetUser) | **Put** /api/2.0/ai/config/user | Update user AI settings
[**AiSettingsSetVectorization**](AISettingsAPI.md#AiSettingsSetVectorization) | **Put** /api/2.0/ai/config/vectorization | Update vectorization settings



## AiSettingsGet

> AiSettingsWrapper AiSettingsGet(ctx).Execute()

Get AI settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-settings-get/).

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
	resp, r, err := apiClient.AISettingsAPI.AiSettingsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISettingsAPI.AiSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiSettingsGet`: AiSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `AISettingsAPI.AiSettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiSettingsGetRequest struct via the builder pattern


### Return type

[**AiSettingsWrapper**](AiSettingsWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiSettingsGetToolMode

> AiSuccessResponse AiSettingsGetToolMode(ctx).Execute()

Get the tool permission mode



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-settings-get-tool-mode/).

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
	resp, r, err := apiClient.AISettingsAPI.AiSettingsGetToolMode(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISettingsAPI.AiSettingsGetToolMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiSettingsGetToolMode`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AISettingsAPI.AiSettingsGetToolMode`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiSettingsGetToolModeRequest struct via the builder pattern


### Return type

[**AiSuccessResponse**](AiSuccessResponse.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiSettingsGetUser

> AiUserSettingsWrapper AiSettingsGetUser(ctx).Execute()

Get user AI settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-settings-get-user/).

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
	resp, r, err := apiClient.AISettingsAPI.AiSettingsGetUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISettingsAPI.AiSettingsGetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiSettingsGetUser`: AiUserSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `AISettingsAPI.AiSettingsGetUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiSettingsGetUserRequest struct via the builder pattern


### Return type

[**AiUserSettingsWrapper**](AiUserSettingsWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiSettingsGetVectorization

> AiVectorizationSettingsWrapper AiSettingsGetVectorization(ctx).Execute()

Get vectorization settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-settings-get-vectorization/).

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
	resp, r, err := apiClient.AISettingsAPI.AiSettingsGetVectorization(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISettingsAPI.AiSettingsGetVectorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiSettingsGetVectorization`: AiVectorizationSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `AISettingsAPI.AiSettingsGetVectorization`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiSettingsGetVectorizationRequest struct via the builder pattern


### Return type

[**AiVectorizationSettingsWrapper**](AiVectorizationSettingsWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiSettingsSetToolMode

> AiSuccessResponse AiSettingsSetToolMode(ctx).AiSettingsSetToolModeRequest(aiSettingsSetToolModeRequest).Execute()

Set the tool permission mode



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-settings-set-tool-mode/).

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
	aiSettingsSetToolModeRequest := map[string]*interface{}{"key": interface{}(123)} // map[string]*interface{} | `{ mode }` with the AI service's `ToolPermissionMode` enum, proxied unchanged; the service rejects anything outside the enum.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISettingsAPI.AiSettingsSetToolMode(context.Background()).AiSettingsSetToolModeRequest(aiSettingsSetToolModeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISettingsAPI.AiSettingsSetToolMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiSettingsSetToolMode`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AISettingsAPI.AiSettingsSetToolMode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiSettingsSetToolModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiSettingsSetToolModeRequest** | **map[string]interface{}** | `{ mode }` with the AI service's `ToolPermissionMode` enum, proxied unchanged; the service rejects anything outside the enum. | 

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


## AiSettingsSetUser

> AiUserSettingsWrapper AiSettingsSetUser(ctx).AiSettingsSetUserRequest(aiSettingsSetUserRequest).Execute()

Update user AI settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-settings-set-user/).

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
	aiSettingsSetUserRequest := map[string]*interface{}{"key": interface{}(123)} // map[string]*interface{} | The user's AI settings: `chatRecommendedModelVisible` as a boolean, required. Read the current one with `GET api/2.0/ai/config/user` and send it back changed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISettingsAPI.AiSettingsSetUser(context.Background()).AiSettingsSetUserRequest(aiSettingsSetUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISettingsAPI.AiSettingsSetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiSettingsSetUser`: AiUserSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `AISettingsAPI.AiSettingsSetUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiSettingsSetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiSettingsSetUserRequest** | **map[string]interface{}** | The user's AI settings: `chatRecommendedModelVisible` as a boolean, required. Read the current one with `GET api/2.0/ai/config/user` and send it back changed. | 

### Return type

[**AiUserSettingsWrapper**](AiUserSettingsWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiSettingsSetVectorization

> AiVectorizationSettingsWrapper AiSettingsSetVectorization(ctx).AiSettingsSetVectorizationRequest(aiSettingsSetVectorizationRequest).Execute()

Update vectorization settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-settings-set-vectorization/).

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
	aiSettingsSetVectorizationRequest := map[string]*interface{}{"key": interface{}(123)} // map[string]*interface{} | The portal's vectorization settings, proxied unchanged to the DocSpace AI service, which owns and validates the shape. Read the current one with `GET api/2.0/ai/config/vectorization` and send it back changed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AISettingsAPI.AiSettingsSetVectorization(context.Background()).AiSettingsSetVectorizationRequest(aiSettingsSetVectorizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AISettingsAPI.AiSettingsSetVectorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiSettingsSetVectorization`: AiVectorizationSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `AISettingsAPI.AiSettingsSetVectorization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiSettingsSetVectorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiSettingsSetVectorizationRequest** | **map[string]interface{}** | The portal's vectorization settings, proxied unchanged to the DocSpace AI service, which owns and validates the shape. Read the current one with `GET api/2.0/ai/config/vectorization` and send it back changed. | 

### Return type

[**AiVectorizationSettingsWrapper**](AiVectorizationSettingsWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

