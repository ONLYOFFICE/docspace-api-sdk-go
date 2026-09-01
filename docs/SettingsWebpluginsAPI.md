# \SettingsWebpluginsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddWebPluginFromFile**](SettingsWebpluginsAPI.md#AddWebPluginFromFile) | **Post** /api/2.0/settings/webplugins | Add a web plugin
[**DeleteWebPlugin**](SettingsWebpluginsAPI.md#DeleteWebPlugin) | **Delete** /api/2.0/settings/webplugins/{name} | Delete a web plugin
[**GetWebPlugin**](SettingsWebpluginsAPI.md#GetWebPlugin) | **Get** /api/2.0/settings/webplugins/{name} | Get a web plugin by name
[**GetWebPlugins**](SettingsWebpluginsAPI.md#GetWebPlugins) | **Get** /api/2.0/settings/webplugins | Get web plugins
[**UpdateWebPlugin**](SettingsWebpluginsAPI.md#UpdateWebPlugin) | **Put** /api/2.0/settings/webplugins/{name} | Update a web plugin



## AddWebPluginFromFile

> WebPluginWrapper AddWebPluginFromFile(ctx).System(system).Execute()

Add a web plugin



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/add-web-plugin-from-file/).

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
	system := false // bool | Specifies whether to load the system plugins or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsWebpluginsAPI.AddWebPluginFromFile(context.Background()).System(system).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsWebpluginsAPI.AddWebPluginFromFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddWebPluginFromFile`: WebPluginWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsWebpluginsAPI.AddWebPluginFromFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddWebPluginFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **system** | **bool** | Specifies whether to load the system plugins or not. | 

### Return type

[**WebPluginWrapper**](WebPluginWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebPlugin

> DeleteWebPlugin(ctx, name).Execute()

Delete a web plugin



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-web-plugin/).

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
	name := "example-plugin" // string | The web plugin name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsWebpluginsAPI.DeleteWebPlugin(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsWebpluginsAPI.DeleteWebPlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The web plugin name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetWebPlugin

> WebPluginWrapper GetWebPlugin(ctx, name).Execute()

Get a web plugin by name



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-web-plugin/).

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
	name := "example-plugin" // string | The web plugin name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsWebpluginsAPI.GetWebPlugin(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsWebpluginsAPI.GetWebPlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebPlugin`: WebPluginWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsWebpluginsAPI.GetWebPlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The web plugin name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebPluginWrapper**](WebPluginWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebPlugins

> WebPluginArrayWrapper GetWebPlugins(ctx).Enabled(enabled).Execute()

Get web plugins



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-web-plugins/).

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
	enabled := true // bool | The optional filter for the plugin enabled state. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsWebpluginsAPI.GetWebPlugins(context.Background()).Enabled(enabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsWebpluginsAPI.GetWebPlugins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebPlugins`: WebPluginArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsWebpluginsAPI.GetWebPlugins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebPluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enabled** | **bool** | The optional filter for the plugin enabled state. | 

### Return type

[**WebPluginArrayWrapper**](WebPluginArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebPlugin

> UpdateWebPlugin(ctx, name).WebPluginRequests(webPluginRequests).Execute()

Update a web plugin



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-web-plugin/).

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
	name := "example-plugin" // string | The web plugin name.
	webPluginRequests := *openapiclient.NewWebPluginRequests("{\"theme\":\"dark\",\"autoSave\":true}") // WebPluginRequests | The configuration settings for the web plugin instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsWebpluginsAPI.UpdateWebPlugin(context.Background(), name).WebPluginRequests(webPluginRequests).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsWebpluginsAPI.UpdateWebPlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The web plugin name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webPluginRequests** | [**WebPluginRequests**](WebPluginRequests.md) | The configuration settings for the web plugin instance. | 

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

