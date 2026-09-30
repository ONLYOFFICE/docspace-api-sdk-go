# \AIToolsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiToolsAddCustomServer**](AIToolsAPI.md#AiToolsAddCustomServer) | **Post** /api/2.0/ai/tools/add-custom-server | Add custom server
[**AiToolsGetAllowAlways**](AIToolsAPI.md#AiToolsGetAllowAlways) | **Get** /api/2.0/ai/tools/get-allow-always | Get allow always
[**AiToolsGetCustomServer**](AIToolsAPI.md#AiToolsGetCustomServer) | **Get** /api/2.0/ai/tools/get-custom-server | Get custom server
[**AiToolsGetDisabled**](AIToolsAPI.md#AiToolsGetDisabled) | **Get** /api/2.0/ai/tools/get-disabled | Get disabled
[**AiToolsIsAllowAlways**](AIToolsAPI.md#AiToolsIsAllowAlways) | **Get** /api/2.0/ai/tools/is-allow-always | Is allow always
[**AiToolsIsToolDisabled**](AIToolsAPI.md#AiToolsIsToolDisabled) | **Get** /api/2.0/ai/tools/is-tool-disabled | Is tool disabled
[**AiToolsListCustomServers**](AIToolsAPI.md#AiToolsListCustomServers) | **Get** /api/2.0/ai/tools/list-custom-servers | List custom servers
[**AiToolsListSystemTools**](AIToolsAPI.md#AiToolsListSystemTools) | **Get** /api/2.0/ai/tools/list-system-tools | List system tools
[**AiToolsRemoveCustomServer**](AIToolsAPI.md#AiToolsRemoveCustomServer) | **Delete** /api/2.0/ai/tools/remove-custom-server | Remove custom server
[**AiToolsReplaceAllCustomServers**](AIToolsAPI.md#AiToolsReplaceAllCustomServers) | **Put** /api/2.0/ai/tools/replace-all-custom-servers | Replace all custom servers
[**AiToolsSetAllowAlways**](AIToolsAPI.md#AiToolsSetAllowAlways) | **Put** /api/2.0/ai/tools/set-allow-always | Set allow always
[**AiToolsSetDisabled**](AIToolsAPI.md#AiToolsSetDisabled) | **Put** /api/2.0/ai/tools/set-disabled | Set disabled
[**AiToolsUpdateCustomServer**](AIToolsAPI.md#AiToolsUpdateCustomServer) | **Put** /api/2.0/ai/tools/update-custom-server | Update custom server



## AiToolsAddCustomServer

> AiToolsMutationResult AiToolsAddCustomServer(ctx).AiToolsAddCustomServerRequest(aiToolsAddCustomServerRequest).Execute()

Add custom server



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-add-custom-server/).

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
	aiToolsAddCustomServerRequest := *openapiclient.NewAiToolsAddCustomServerRequest("Name_example", map[string]interface{}(123)) // AiToolsAddCustomServerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIToolsAPI.AiToolsAddCustomServer(context.Background()).AiToolsAddCustomServerRequest(aiToolsAddCustomServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIToolsAPI.AiToolsAddCustomServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiToolsAddCustomServer`: AiToolsMutationResult
	fmt.Fprintf(os.Stdout, "Response from `AIToolsAPI.AiToolsAddCustomServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiToolsAddCustomServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiToolsAddCustomServerRequest** | [**AiToolsAddCustomServerRequest**](AiToolsAddCustomServerRequest.md) |  | 

### Return type

[**AiToolsMutationResult**](AiToolsMutationResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiToolsGetAllowAlways

> []string AiToolsGetAllowAlways(ctx).EntityId(entityId).Execute()

Get allow always



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-get-allow-always/).

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
	resp, r, err := apiClient.AIToolsAPI.AiToolsGetAllowAlways(context.Background()).EntityId(entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIToolsAPI.AiToolsGetAllowAlways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiToolsGetAllowAlways`: []string
	fmt.Fprintf(os.Stdout, "Response from `AIToolsAPI.AiToolsGetAllowAlways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiToolsGetAllowAlwaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string** | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. | 

### Return type

**[]string**

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiToolsGetCustomServer

> map[string]interface{} AiToolsGetCustomServer(ctx).Name(name).EntityId(entityId).Execute()

Get custom server



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-get-custom-server/).

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
	name := "acme-mcp" // string | The custom MCP server name.
	entityId := "1234" // string | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIToolsAPI.AiToolsGetCustomServer(context.Background()).Name(name).EntityId(entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIToolsAPI.AiToolsGetCustomServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiToolsGetCustomServer`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIToolsAPI.AiToolsGetCustomServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiToolsGetCustomServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The custom MCP server name. | 
 **entityId** | **string** | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. | 

### Return type

**map[string]interface{}**

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiToolsGetDisabled

> map[string][]string AiToolsGetDisabled(ctx).EntityId(entityId).Execute()

Get disabled



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-get-disabled/).

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
	resp, r, err := apiClient.AIToolsAPI.AiToolsGetDisabled(context.Background()).EntityId(entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIToolsAPI.AiToolsGetDisabled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiToolsGetDisabled`: map[string][]string
	fmt.Fprintf(os.Stdout, "Response from `AIToolsAPI.AiToolsGetDisabled`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiToolsGetDisabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string** | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. | 

### Return type

[**map[string][]string**](array.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiToolsIsAllowAlways

> bool AiToolsIsAllowAlways(ctx).ServerType(serverType).ToolName(toolName).EntityId(entityId).Execute()

Is allow always



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-is-allow-always/).

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
	serverType := "docspace" // string | The MCP server type the tool belongs to.
	toolName := "docspace_get_folder" // string | The tool name.
	entityId := "1234" // string | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIToolsAPI.AiToolsIsAllowAlways(context.Background()).ServerType(serverType).ToolName(toolName).EntityId(entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIToolsAPI.AiToolsIsAllowAlways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiToolsIsAllowAlways`: bool
	fmt.Fprintf(os.Stdout, "Response from `AIToolsAPI.AiToolsIsAllowAlways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiToolsIsAllowAlwaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverType** | **string** | The MCP server type the tool belongs to. | 
 **toolName** | **string** | The tool name. | 
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


## AiToolsIsToolDisabled

> bool AiToolsIsToolDisabled(ctx).ServerType(serverType).ToolName(toolName).EntityId(entityId).Execute()

Is tool disabled



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-is-tool-disabled/).

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
	serverType := "docspace" // string | The MCP server type the tool belongs to.
	toolName := "docspace_get_folder" // string | The tool name.
	entityId := "1234" // string | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIToolsAPI.AiToolsIsToolDisabled(context.Background()).ServerType(serverType).ToolName(toolName).EntityId(entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIToolsAPI.AiToolsIsToolDisabled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiToolsIsToolDisabled`: bool
	fmt.Fprintf(os.Stdout, "Response from `AIToolsAPI.AiToolsIsToolDisabled`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiToolsIsToolDisabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverType** | **string** | The MCP server type the tool belongs to. | 
 **toolName** | **string** | The tool name. | 
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


## AiToolsListCustomServers

> map[string]map[string]interface{} AiToolsListCustomServers(ctx).EntityId(entityId).Execute()

List custom servers



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-list-custom-servers/).

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
	resp, r, err := apiClient.AIToolsAPI.AiToolsListCustomServers(context.Background()).EntityId(entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIToolsAPI.AiToolsListCustomServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiToolsListCustomServers`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AIToolsAPI.AiToolsListCustomServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiToolsListCustomServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string** | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. | 

### Return type

**map[string]map[string]interface{}**

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiToolsListSystemTools

> AiToolsListSystemTools200Response AiToolsListSystemTools(ctx).EntityId(entityId).Execute()

List system tools



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-list-system-tools/).

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
	resp, r, err := apiClient.AIToolsAPI.AiToolsListSystemTools(context.Background()).EntityId(entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIToolsAPI.AiToolsListSystemTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiToolsListSystemTools`: AiToolsListSystemTools200Response
	fmt.Fprintf(os.Stdout, "Response from `AIToolsAPI.AiToolsListSystemTools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiToolsListSystemToolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string** | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. | 

### Return type

[**AiToolsListSystemTools200Response**](AiToolsListSystemTools200Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiToolsRemoveCustomServer

> AiSuccessResponse AiToolsRemoveCustomServer(ctx).AiToolsRemoveCustomServerRequest(aiToolsRemoveCustomServerRequest).Execute()

Remove custom server



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-remove-custom-server/).

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
	aiToolsRemoveCustomServerRequest := *openapiclient.NewAiToolsRemoveCustomServerRequest("Name_example") // AiToolsRemoveCustomServerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIToolsAPI.AiToolsRemoveCustomServer(context.Background()).AiToolsRemoveCustomServerRequest(aiToolsRemoveCustomServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIToolsAPI.AiToolsRemoveCustomServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiToolsRemoveCustomServer`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIToolsAPI.AiToolsRemoveCustomServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiToolsRemoveCustomServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiToolsRemoveCustomServerRequest** | [**AiToolsRemoveCustomServerRequest**](AiToolsRemoveCustomServerRequest.md) |  | 

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


## AiToolsReplaceAllCustomServers

> AiToolsBulkResult AiToolsReplaceAllCustomServers(ctx).AiToolsReplaceAllCustomServersRequest(aiToolsReplaceAllCustomServersRequest).Execute()

Replace all custom servers



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-replace-all-custom-servers/).

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
	aiToolsReplaceAllCustomServersRequest := *openapiclient.NewAiToolsReplaceAllCustomServersRequest(map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // AiToolsReplaceAllCustomServersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIToolsAPI.AiToolsReplaceAllCustomServers(context.Background()).AiToolsReplaceAllCustomServersRequest(aiToolsReplaceAllCustomServersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIToolsAPI.AiToolsReplaceAllCustomServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiToolsReplaceAllCustomServers`: AiToolsBulkResult
	fmt.Fprintf(os.Stdout, "Response from `AIToolsAPI.AiToolsReplaceAllCustomServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiToolsReplaceAllCustomServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiToolsReplaceAllCustomServersRequest** | [**AiToolsReplaceAllCustomServersRequest**](AiToolsReplaceAllCustomServersRequest.md) |  | 

### Return type

[**AiToolsBulkResult**](AiToolsBulkResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiToolsSetAllowAlways

> AiSuccessResponse AiToolsSetAllowAlways(ctx).AiToolsSetAllowAlwaysRequest(aiToolsSetAllowAlwaysRequest).Execute()

Set allow always



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-set-allow-always/).

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
	aiToolsSetAllowAlwaysRequest := *openapiclient.NewAiToolsSetAllowAlwaysRequest("ServerType_example", "ToolName_example", false) // AiToolsSetAllowAlwaysRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIToolsAPI.AiToolsSetAllowAlways(context.Background()).AiToolsSetAllowAlwaysRequest(aiToolsSetAllowAlwaysRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIToolsAPI.AiToolsSetAllowAlways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiToolsSetAllowAlways`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIToolsAPI.AiToolsSetAllowAlways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiToolsSetAllowAlwaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiToolsSetAllowAlwaysRequest** | [**AiToolsSetAllowAlwaysRequest**](AiToolsSetAllowAlwaysRequest.md) |  | 

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


## AiToolsSetDisabled

> AiSuccessResponse AiToolsSetDisabled(ctx).AiToolsSetDisabledRequest(aiToolsSetDisabledRequest).Execute()

Set disabled



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-set-disabled/).

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
	aiToolsSetDisabledRequest := *openapiclient.NewAiToolsSetDisabledRequest("ServerType_example", []string{"ToolNames_example"}) // AiToolsSetDisabledRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIToolsAPI.AiToolsSetDisabled(context.Background()).AiToolsSetDisabledRequest(aiToolsSetDisabledRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIToolsAPI.AiToolsSetDisabled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiToolsSetDisabled`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIToolsAPI.AiToolsSetDisabled`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiToolsSetDisabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiToolsSetDisabledRequest** | [**AiToolsSetDisabledRequest**](AiToolsSetDisabledRequest.md) |  | 

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


## AiToolsUpdateCustomServer

> AiToolsMutationResult AiToolsUpdateCustomServer(ctx).AiToolsUpdateCustomServerRequest(aiToolsUpdateCustomServerRequest).Execute()

Update custom server



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-tools-update-custom-server/).

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
	aiToolsUpdateCustomServerRequest := *openapiclient.NewAiToolsUpdateCustomServerRequest("Name_example", map[string]interface{}(123)) // AiToolsUpdateCustomServerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIToolsAPI.AiToolsUpdateCustomServer(context.Background()).AiToolsUpdateCustomServerRequest(aiToolsUpdateCustomServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIToolsAPI.AiToolsUpdateCustomServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiToolsUpdateCustomServer`: AiToolsMutationResult
	fmt.Fprintf(os.Stdout, "Response from `AIToolsAPI.AiToolsUpdateCustomServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiToolsUpdateCustomServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiToolsUpdateCustomServerRequest** | [**AiToolsUpdateCustomServerRequest**](AiToolsUpdateCustomServerRequest.md) |  | 

### Return type

[**AiToolsMutationResult**](AiToolsMutationResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

