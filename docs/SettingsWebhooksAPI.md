# \SettingsWebhooksAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](SettingsWebhooksAPI.md#CreateWebhook) | **Post** /api/2.0/settings/webhook | Create a webhook
[**EnableWebhook**](SettingsWebhooksAPI.md#EnableWebhook) | **Put** /api/2.0/settings/webhook/enable | Enable a webhook
[**GetTenantWebhooks**](SettingsWebhooksAPI.md#GetTenantWebhooks) | **Get** /api/2.0/settings/webhook | Get webhooks
[**GetWebhookTriggers**](SettingsWebhooksAPI.md#GetWebhookTriggers) | **Get** /api/2.0/settings/webhook/triggers | Get webhook triggers
[**GetWebhooksLogs**](SettingsWebhooksAPI.md#GetWebhooksLogs) | **Get** /api/2.0/settings/webhooks/log | Get webhook logs
[**RemoveWebhook**](SettingsWebhooksAPI.md#RemoveWebhook) | **Delete** /api/2.0/settings/webhook/{id} | Remove a webhook
[**RetryWebhook**](SettingsWebhooksAPI.md#RetryWebhook) | **Put** /api/2.0/settings/webhook/{id}/retry | Retry a webhook
[**RetryWebhooks**](SettingsWebhooksAPI.md#RetryWebhooks) | **Put** /api/2.0/settings/webhook/retry | Retry webhooks
[**UpdateWebhook**](SettingsWebhooksAPI.md#UpdateWebhook) | **Put** /api/2.0/settings/webhook | Update a webhook



## CreateWebhook

> WebhooksConfigWrapper CreateWebhook(ctx).CreateWebhooksConfigRequestsDto(createWebhooksConfigRequestsDto).Execute()

Create a webhook



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-webhook/).

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
	createWebhooksConfigRequestsDto := *openapiclient.NewCreateWebhooksConfigRequestsDto("Production Webhook", "https://example.com/webhook") // CreateWebhooksConfigRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsWebhooksAPI.CreateWebhook(context.Background()).CreateWebhooksConfigRequestsDto(createWebhooksConfigRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsWebhooksAPI.CreateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhook`: WebhooksConfigWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsWebhooksAPI.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWebhooksConfigRequestsDto** | [**CreateWebhooksConfigRequestsDto**](CreateWebhooksConfigRequestsDto.md) |  | 

### Return type

[**WebhooksConfigWrapper**](WebhooksConfigWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableWebhook

> WebhooksConfigWrapper EnableWebhook(ctx).UpdateWebhooksConfigRequestsDto(updateWebhooksConfigRequestsDto).Execute()

Enable a webhook



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/enable-webhook/).

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
	updateWebhooksConfigRequestsDto := *openapiclient.NewUpdateWebhooksConfigRequestsDto("Production Webhook", "https://example.com/webhook", int32(1)) // UpdateWebhooksConfigRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsWebhooksAPI.EnableWebhook(context.Background()).UpdateWebhooksConfigRequestsDto(updateWebhooksConfigRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsWebhooksAPI.EnableWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableWebhook`: WebhooksConfigWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsWebhooksAPI.EnableWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateWebhooksConfigRequestsDto** | [**UpdateWebhooksConfigRequestsDto**](UpdateWebhooksConfigRequestsDto.md) |  | 

### Return type

[**WebhooksConfigWrapper**](WebhooksConfigWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantWebhooks

> WebhooksConfigWithStatusArrayWrapper GetTenantWebhooks(ctx).Execute()

Get webhooks



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-tenant-webhooks/).

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
	resp, r, err := apiClient.SettingsWebhooksAPI.GetTenantWebhooks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsWebhooksAPI.GetTenantWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantWebhooks`: WebhooksConfigWithStatusArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsWebhooksAPI.GetTenantWebhooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantWebhooksRequest struct via the builder pattern


### Return type

[**WebhooksConfigWithStatusArrayWrapper**](WebhooksConfigWithStatusArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookTriggers

> WebhookTriggerArrayWrapper GetWebhookTriggers(ctx).Execute()

Get webhook triggers



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-webhook-triggers/).

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
	resp, r, err := apiClient.SettingsWebhooksAPI.GetWebhookTriggers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsWebhooksAPI.GetWebhookTriggers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookTriggers`: WebhookTriggerArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsWebhooksAPI.GetWebhookTriggers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookTriggersRequest struct via the builder pattern


### Return type

[**WebhookTriggerArrayWrapper**](WebhookTriggerArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooksLogs

> WebhooksLogArrayWrapper GetWebhooksLogs(ctx).DeliveryFrom(deliveryFrom).DeliveryTo(deliveryTo).HookUri(hookUri).ConfigId(configId).EventId(eventId).GroupStatus(groupStatus).UserId(userId).Trigger(trigger).Count(count).StartIndex(startIndex).Execute()

Get webhook logs



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-webhooks-logs/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	deliveryFrom := time.Now() // time.Time | The delivery start time for filtering webhook logs. (optional)
	deliveryTo := time.Now() // time.Time | The delivery end time for filtering webhook logs. (optional)
	hookUri := "https://example.com/webhook" // string | The destination URL where webhooks are delivered. (optional)
	configId := int32(1) // int32 | The webhook configuration identifier. (optional)
	eventId := int32(1) // int32 | The unique identifier of the event that triggered the webhook. (optional)
	groupStatus := openapiclient.WebhookGroupStatus(0) // WebhookGroupStatus | The status of the webhook delivery group. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the user associated with the webhook event. (optional)
	trigger := openapiclient.WebhookTrigger(0) // WebhookTrigger | The type of event that triggered the webhook. (optional)
	count := int32(1) // int32 | The maximum number of webhook log records to return in the query response. (optional)
	startIndex := int32(1) // int32 | Specifies the starting index for retrieving webhook logs.  Used for pagination in the webhook delivery log queries. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsWebhooksAPI.GetWebhooksLogs(context.Background()).DeliveryFrom(deliveryFrom).DeliveryTo(deliveryTo).HookUri(hookUri).ConfigId(configId).EventId(eventId).GroupStatus(groupStatus).UserId(userId).Trigger(trigger).Count(count).StartIndex(startIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsWebhooksAPI.GetWebhooksLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhooksLogs`: WebhooksLogArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsWebhooksAPI.GetWebhooksLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deliveryFrom** | **time.Time** | The delivery start time for filtering webhook logs. | 
 **deliveryTo** | **time.Time** | The delivery end time for filtering webhook logs. | 
 **hookUri** | **string** | The destination URL where webhooks are delivered. | 
 **configId** | **int32** | The webhook configuration identifier. | 
 **eventId** | **int32** | The unique identifier of the event that triggered the webhook. | 
 **groupStatus** | [**WebhookGroupStatus**](WebhookGroupStatus.md) | The status of the webhook delivery group. | 
 **userId** | **string** | The identifier of the user associated with the webhook event. | 
 **trigger** | [**WebhookTrigger**](WebhookTrigger.md) | The type of event that triggered the webhook. | 
 **count** | **int32** | The maximum number of webhook log records to return in the query response. | 
 **startIndex** | **int32** | Specifies the starting index for retrieving webhook logs.  Used for pagination in the webhook delivery log queries. | 

### Return type

[**WebhooksLogArrayWrapper**](WebhooksLogArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveWebhook

> WebhooksConfigWrapper RemoveWebhook(ctx, id).Execute()

Remove a webhook



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/remove-webhook/).

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
	id := int32(1) // int32 | The ID extracted from the route parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsWebhooksAPI.RemoveWebhook(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsWebhooksAPI.RemoveWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveWebhook`: WebhooksConfigWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsWebhooksAPI.RemoveWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID extracted from the route parameters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhooksConfigWrapper**](WebhooksConfigWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryWebhook

> WebhooksLogWrapper RetryWebhook(ctx, id).Execute()

Retry a webhook



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/retry-webhook/).

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
	id := int32(1) // int32 | The ID extracted from the route parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsWebhooksAPI.RetryWebhook(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsWebhooksAPI.RetryWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryWebhook`: WebhooksLogWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsWebhooksAPI.RetryWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID extracted from the route parameters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetryWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhooksLogWrapper**](WebhooksLogWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetryWebhooks

> WebhooksLogArrayWrapper RetryWebhooks(ctx).WebhookRetryRequestsDto(webhookRetryRequestsDto).Execute()

Retry webhooks



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/retry-webhooks/).

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
	webhookRetryRequestsDto := *openapiclient.NewWebhookRetryRequestsDto() // WebhookRetryRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsWebhooksAPI.RetryWebhooks(context.Background()).WebhookRetryRequestsDto(webhookRetryRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsWebhooksAPI.RetryWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetryWebhooks`: WebhooksLogArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsWebhooksAPI.RetryWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetryWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookRetryRequestsDto** | [**WebhookRetryRequestsDto**](WebhookRetryRequestsDto.md) |  | 

### Return type

[**WebhooksLogArrayWrapper**](WebhooksLogArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook

> WebhooksConfigWrapper UpdateWebhook(ctx).UpdateWebhooksConfigRequestsDto(updateWebhooksConfigRequestsDto).Execute()

Update a webhook



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-webhook/).

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
	updateWebhooksConfigRequestsDto := *openapiclient.NewUpdateWebhooksConfigRequestsDto("Production Webhook", "https://example.com/webhook", int32(1)) // UpdateWebhooksConfigRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsWebhooksAPI.UpdateWebhook(context.Background()).UpdateWebhooksConfigRequestsDto(updateWebhooksConfigRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsWebhooksAPI.UpdateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhook`: WebhooksConfigWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsWebhooksAPI.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateWebhooksConfigRequestsDto** | [**UpdateWebhooksConfigRequestsDto**](UpdateWebhooksConfigRequestsDto.md) |  | 

### Return type

[**WebhooksConfigWrapper**](WebhooksConfigWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

