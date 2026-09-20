# \AIPreferencesAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiPreferencesClearDeepMode**](AIPreferencesAPI.md#AiPreferencesClearDeepMode) | **Delete** /api/2.0/ai/preferences/clear-deep-mode | Clear deep mode
[**AiPreferencesGetDeepMode**](AIPreferencesAPI.md#AiPreferencesGetDeepMode) | **Get** /api/2.0/ai/preferences/get-deep-mode | Get deep mode
[**AiPreferencesGetReasoningLevel**](AIPreferencesAPI.md#AiPreferencesGetReasoningLevel) | **Get** /api/2.0/ai/preferences/get-reasoning-level | Get reasoning level
[**AiPreferencesIsDeepModeSet**](AIPreferencesAPI.md#AiPreferencesIsDeepModeSet) | **Get** /api/2.0/ai/preferences/is-deep-mode-set | Is deep mode set
[**AiPreferencesSetDeepMode**](AIPreferencesAPI.md#AiPreferencesSetDeepMode) | **Put** /api/2.0/ai/preferences/set-deep-mode | Set deep mode
[**AiPreferencesSetReasoningLevel**](AIPreferencesAPI.md#AiPreferencesSetReasoningLevel) | **Put** /api/2.0/ai/preferences/set-reasoning-level | Set reasoning level



## AiPreferencesClearDeepMode

> AiSuccessResponse AiPreferencesClearDeepMode(ctx).Body(body).Execute()

Clear deep mode



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-preferences-clear-deep-mode/).

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
	body := "body_example" // string | The ID of the room whose preference is cleared, as a bare JSON string. Send an empty body to clear the portal-wide preference.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIPreferencesAPI.AiPreferencesClearDeepMode(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIPreferencesAPI.AiPreferencesClearDeepMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiPreferencesClearDeepMode`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIPreferencesAPI.AiPreferencesClearDeepMode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiPreferencesClearDeepModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | The ID of the room whose preference is cleared, as a bare JSON string. Send an empty body to clear the portal-wide preference. | 

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


## AiPreferencesGetDeepMode

> bool AiPreferencesGetDeepMode(ctx).EntityId(entityId).Execute()

Get deep mode



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-preferences-get-deep-mode/).

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
	resp, r, err := apiClient.AIPreferencesAPI.AiPreferencesGetDeepMode(context.Background()).EntityId(entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIPreferencesAPI.AiPreferencesGetDeepMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiPreferencesGetDeepMode`: bool
	fmt.Fprintf(os.Stdout, "Response from `AIPreferencesAPI.AiPreferencesGetDeepMode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiPreferencesGetDeepModeRequest struct via the builder pattern


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


## AiPreferencesGetReasoningLevel

> AiAiReasoningLevel AiPreferencesGetReasoningLevel(ctx).EntityId(entityId).Execute()

Get reasoning level



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-preferences-get-reasoning-level/).

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
	resp, r, err := apiClient.AIPreferencesAPI.AiPreferencesGetReasoningLevel(context.Background()).EntityId(entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIPreferencesAPI.AiPreferencesGetReasoningLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiPreferencesGetReasoningLevel`: AiAiReasoningLevel
	fmt.Fprintf(os.Stdout, "Response from `AIPreferencesAPI.AiPreferencesGetReasoningLevel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiPreferencesGetReasoningLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string** | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. | 

### Return type

[**AiAiReasoningLevel**](AiAiReasoningLevel.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiPreferencesIsDeepModeSet

> bool AiPreferencesIsDeepModeSet(ctx).EntityId(entityId).Execute()

Is deep mode set



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-preferences-is-deep-mode-set/).

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
	resp, r, err := apiClient.AIPreferencesAPI.AiPreferencesIsDeepModeSet(context.Background()).EntityId(entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIPreferencesAPI.AiPreferencesIsDeepModeSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiPreferencesIsDeepModeSet`: bool
	fmt.Fprintf(os.Stdout, "Response from `AIPreferencesAPI.AiPreferencesIsDeepModeSet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiPreferencesIsDeepModeSetRequest struct via the builder pattern


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


## AiPreferencesSetDeepMode

> AiSuccessResponse AiPreferencesSetDeepMode(ctx).AiPreferencesSetDeepModeRequest(aiPreferencesSetDeepModeRequest).Execute()

Set deep mode



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-preferences-set-deep-mode/).

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
	aiPreferencesSetDeepModeRequest := *openapiclient.NewAiPreferencesSetDeepModeRequest(false) // AiPreferencesSetDeepModeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIPreferencesAPI.AiPreferencesSetDeepMode(context.Background()).AiPreferencesSetDeepModeRequest(aiPreferencesSetDeepModeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIPreferencesAPI.AiPreferencesSetDeepMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiPreferencesSetDeepMode`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIPreferencesAPI.AiPreferencesSetDeepMode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiPreferencesSetDeepModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiPreferencesSetDeepModeRequest** | [**AiPreferencesSetDeepModeRequest**](AiPreferencesSetDeepModeRequest.md) |  | 

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


## AiPreferencesSetReasoningLevel

> AiSuccessResponse AiPreferencesSetReasoningLevel(ctx).AiPreferencesSetReasoningLevelRequest(aiPreferencesSetReasoningLevelRequest).Execute()

Set reasoning level



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-preferences-set-reasoning-level/).

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
	aiPreferencesSetReasoningLevelRequest := *openapiclient.NewAiPreferencesSetReasoningLevelRequest(openapiclient.AiAiReasoningLevel("off")) // AiPreferencesSetReasoningLevelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIPreferencesAPI.AiPreferencesSetReasoningLevel(context.Background()).AiPreferencesSetReasoningLevelRequest(aiPreferencesSetReasoningLevelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIPreferencesAPI.AiPreferencesSetReasoningLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiPreferencesSetReasoningLevel`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIPreferencesAPI.AiPreferencesSetReasoningLevel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiPreferencesSetReasoningLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiPreferencesSetReasoningLevelRequest** | [**AiPreferencesSetReasoningLevelRequest**](AiPreferencesSetReasoningLevelRequest.md) |  | 

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

