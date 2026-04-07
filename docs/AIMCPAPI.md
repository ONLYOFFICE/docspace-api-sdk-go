# \AIMCPAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRoomServers**](AIMCPAPI.md#AddRoomServers) | **Post** /api/2.0/ai/rooms/{roomId}/servers | Assign MCP servers to a room
[**AddServer**](AIMCPAPI.md#AddServer) | **Post** /api/2.0/ai/servers | Register a custom MCP server
[**ConnectServer**](AIMCPAPI.md#ConnectServer) | **Post** /api/2.0/ai/rooms/{roomId}/servers/{serverId}/connect | Connect an OAuth-based MCP server in a room
[**DeleteRoomServers**](AIMCPAPI.md#DeleteRoomServers) | **Delete** /api/2.0/ai/rooms/{roomId}/servers | Remove MCP servers from a room
[**DeleteServer**](AIMCPAPI.md#DeleteServer) | **Delete** /api/2.0/ai/servers | Delete MCP servers
[**DisconnectServer**](AIMCPAPI.md#DisconnectServer) | **Post** /api/2.0/ai/rooms/{roomId}/servers/{serverId}/disconnect | Disconnect an MCP server in a room
[**GetAvailableServers**](AIMCPAPI.md#GetAvailableServers) | **Get** /api/2.0/ai/servers/available | Get available MCP servers
[**GetRoomServers**](AIMCPAPI.md#GetRoomServers) | **Get** /api/2.0/ai/rooms/{roomId}/servers | Get MCP servers assigned to a room
[**GetServer**](AIMCPAPI.md#GetServer) | **Get** /api/2.0/ai/servers/{id} | Get an MCP server by ID
[**GetServers**](AIMCPAPI.md#GetServers) | **Get** /api/2.0/ai/servers | Get all MCP servers
[**GetTools**](AIMCPAPI.md#GetTools) | **Get** /api/2.0/ai/rooms/{roomId}/servers/{serverId}/tools | Get MCP server tools in a room
[**SetServerStatus**](AIMCPAPI.md#SetServerStatus) | **Put** /api/2.0/ai/servers/{id}/status | Enable or disable an MCP server
[**SetTools**](AIMCPAPI.md#SetTools) | **Put** /api/2.0/ai/rooms/{roomId}/servers/{serverId}/tools | Configure MCP server tools in a room
[**UpdateServer**](AIMCPAPI.md#UpdateServer) | **Put** /api/2.0/ai/servers/{id} | Update a custom MCP server



## AddRoomServers

> McpServerStatusArrayWrapper AddRoomServers(ctx, roomId).AddRoomServersRequestBody(addRoomServersRequestBody).Execute()

Assign MCP servers to a room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/add-room-servers/).

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
	roomId := int32(42) // int32 | Identifier of the room to which MCP servers will be assigned.
	addRoomServersRequestBody := *openapiclient.NewAddRoomServersRequestBody([]string{"Servers_example"}) // AddRoomServersRequestBody | Server identifiers to assign.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIMCPAPI.AddRoomServers(context.Background(), roomId).AddRoomServersRequestBody(addRoomServersRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIMCPAPI.AddRoomServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddRoomServers`: McpServerStatusArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIMCPAPI.AddRoomServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **int32** | Identifier of the room to which MCP servers will be assigned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRoomServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addRoomServersRequestBody** | [**AddRoomServersRequestBody**](AddRoomServersRequestBody.md) | Server identifiers to assign. | 

### Return type

[**McpServerStatusArrayWrapper**](McpServerStatusArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddServer

> McpServerWrapper AddServer(ctx).AddMcpServerRequestBody(addMcpServerRequestBody).Execute()

Register a custom MCP server



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/add-server/).

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
	addMcpServerRequestBody := *openapiclient.NewAddMcpServerRequestBody("my-custom-server", "Custom MCP server for project management tools", "https://mcp.example.com/sse") // AddMcpServerRequestBody | MCP server registration parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIMCPAPI.AddServer(context.Background()).AddMcpServerRequestBody(addMcpServerRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIMCPAPI.AddServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddServer`: McpServerWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIMCPAPI.AddServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addMcpServerRequestBody** | [**AddMcpServerRequestBody**](AddMcpServerRequestBody.md) | MCP server registration parameters. | 

### Return type

[**McpServerWrapper**](McpServerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectServer

> McpServerStatusWrapper ConnectServer(ctx, roomId, serverId).ConnectServerRequestBody(connectServerRequestBody).Execute()

Connect an OAuth-based MCP server in a room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/connect-server/).

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
	roomId := int32(42) // int32 | Identifier of the room containing the MCP server.
	serverId := "00000000-0000-0000-0000-000000000000" // string | Unique identifier of the MCP server to connect.
	connectServerRequestBody := *openapiclient.NewConnectServerRequestBody("abc123def456") // ConnectServerRequestBody | The request body containing additional data necessary for connecting to the server,  such as authentication or operation-specific information.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIMCPAPI.ConnectServer(context.Background(), roomId, serverId).ConnectServerRequestBody(connectServerRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIMCPAPI.ConnectServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectServer`: McpServerStatusWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIMCPAPI.ConnectServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **int32** | Identifier of the room containing the MCP server. | 
**serverId** | **string** | Unique identifier of the MCP server to connect. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **connectServerRequestBody** | [**ConnectServerRequestBody**](ConnectServerRequestBody.md) | The request body containing additional data necessary for connecting to the server,  such as authentication or operation-specific information. | 

### Return type

[**McpServerStatusWrapper**](McpServerStatusWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoomServers

> DeleteRoomServers(ctx, roomId).DeleteRoomServersRequestBody(deleteRoomServersRequestBody).Execute()

Remove MCP servers from a room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-room-servers/).

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
	roomId := int32(42) // int32 | Identifier of the room from which MCP servers will be removed.
	deleteRoomServersRequestBody := *openapiclient.NewDeleteRoomServersRequestBody([]string{"Servers_example"}) // DeleteRoomServersRequestBody | Server identifiers to remove.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIMCPAPI.DeleteRoomServers(context.Background(), roomId).DeleteRoomServersRequestBody(deleteRoomServersRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIMCPAPI.DeleteRoomServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **int32** | Identifier of the room from which MCP servers will be removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoomServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteRoomServersRequestBody** | [**DeleteRoomServersRequestBody**](DeleteRoomServersRequestBody.md) | Server identifiers to remove. | 

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


## DeleteServer

> DeleteServer(ctx).DeleteServersRequestBody(deleteServersRequestBody).Execute()

Delete MCP servers



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-server/).

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
	deleteServersRequestBody := *openapiclient.NewDeleteServersRequestBody([]string{"Servers_example"}) // DeleteServersRequestBody | Server identifiers to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AIMCPAPI.DeleteServer(context.Background()).DeleteServersRequestBody(deleteServersRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIMCPAPI.DeleteServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteServersRequestBody** | [**DeleteServersRequestBody**](DeleteServersRequestBody.md) | Server identifiers to delete. | 

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


## DisconnectServer

> McpServerStatusWrapper DisconnectServer(ctx, roomId, serverId).Execute()

Disconnect an MCP server in a room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/disconnect-server/).

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
	roomId := int32(42) // int32 | Identifier of the room containing the MCP server.
	serverId := "00000000-0000-0000-0000-000000000000" // string | Unique identifier of the MCP server to disconnect from.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIMCPAPI.DisconnectServer(context.Background(), roomId, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIMCPAPI.DisconnectServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisconnectServer`: McpServerStatusWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIMCPAPI.DisconnectServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **int32** | Identifier of the room containing the MCP server. | 
**serverId** | **string** | Unique identifier of the MCP server to disconnect from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**McpServerStatusWrapper**](McpServerStatusWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableServers

> McpServerShortArrayWrapper GetAvailableServers(ctx).StartIndex(startIndex).Count(count).Execute()

Get available MCP servers



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-available-servers/).

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
	resp, r, err := apiClient.AIMCPAPI.GetAvailableServers(context.Background()).StartIndex(startIndex).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIMCPAPI.GetAvailableServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableServers`: McpServerShortArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIMCPAPI.GetAvailableServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startIndex** | **int32** | The number of items to skip before returning results (zero-based offset). Defaults to 0. | 
 **count** | **int32** | The maximum number of items to return per page. Defaults to 100. | 

### Return type

[**McpServerShortArrayWrapper**](McpServerShortArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomServers

> McpServerStatusArrayWrapper GetRoomServers(ctx, roomId).Execute()

Get MCP servers assigned to a room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-room-servers/).

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
	roomId := int32(42) // int32 | Identifier of the room whose assigned MCP servers are being retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIMCPAPI.GetRoomServers(context.Background(), roomId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIMCPAPI.GetRoomServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomServers`: McpServerStatusArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIMCPAPI.GetRoomServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **int32** | Identifier of the room whose assigned MCP servers are being retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**McpServerStatusArrayWrapper**](McpServerStatusArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServer

> McpServerShortWrapper GetServer(ctx, id).Execute()

Get an MCP server by ID



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-server/).

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
	id := "00000000-0000-0000-0000-000000000000" // string | Unique identifier of the MCP server to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIMCPAPI.GetServer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIMCPAPI.GetServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServer`: McpServerShortWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIMCPAPI.GetServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the MCP server to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**McpServerShortWrapper**](McpServerShortWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServers

> McpServerArrayWrapper GetServers(ctx).StartIndex(startIndex).Count(count).Execute()

Get all MCP servers



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-servers/).

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
	resp, r, err := apiClient.AIMCPAPI.GetServers(context.Background()).StartIndex(startIndex).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIMCPAPI.GetServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServers`: McpServerArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIMCPAPI.GetServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startIndex** | **int32** | The number of items to skip before returning results (zero-based offset). Defaults to 0. | 
 **count** | **int32** | The maximum number of items to return per page. Defaults to 100. | 

### Return type

[**McpServerArrayWrapper**](McpServerArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTools

> McpToolArrayWrapper GetTools(ctx, roomId, serverId).Execute()

Get MCP server tools in a room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-tools/).

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
	roomId := int32(42) // int32 | Identifier of the room containing the MCP server.
	serverId := "00000000-0000-0000-0000-000000000000" // string | Unique identifier of the MCP server whose tools are being retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIMCPAPI.GetTools(context.Background(), roomId, serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIMCPAPI.GetTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTools`: McpToolArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIMCPAPI.GetTools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **int32** | Identifier of the room containing the MCP server. | 
**serverId** | **string** | Unique identifier of the MCP server whose tools are being retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**McpToolArrayWrapper**](McpToolArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetServerStatus

> McpServerWrapper SetServerStatus(ctx, id).SetServerStatusRequestBody(setServerStatusRequestBody).Execute()

Enable or disable an MCP server



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-server-status/).

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
	id := "00000000-0000-0000-0000-000000000000" // string | Unique identifier of the MCP server whose status is being changed.
	setServerStatusRequestBody := *openapiclient.NewSetServerStatusRequestBody() // SetServerStatusRequestBody | New status value.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIMCPAPI.SetServerStatus(context.Background(), id).SetServerStatusRequestBody(setServerStatusRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIMCPAPI.SetServerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetServerStatus`: McpServerWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIMCPAPI.SetServerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the MCP server whose status is being changed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetServerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setServerStatusRequestBody** | [**SetServerStatusRequestBody**](SetServerStatusRequestBody.md) | New status value. | 

### Return type

[**McpServerWrapper**](McpServerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTools

> McpToolArrayWrapper SetTools(ctx, roomId, serverId).SetMcpToolsRequestBody(setMcpToolsRequestBody).Execute()

Configure MCP server tools in a room



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-tools/).

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
	roomId := int32(42) // int32 | Identifier of the room containing the MCP server.
	serverId := "00000000-0000-0000-0000-000000000000" // string | Unique identifier of the MCP server whose tools are being configured.
	setMcpToolsRequestBody := *openapiclient.NewSetMcpToolsRequestBody([]string{"DisabledTools_example"}) // SetMcpToolsRequestBody | Tool configuration parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIMCPAPI.SetTools(context.Background(), roomId, serverId).SetMcpToolsRequestBody(setMcpToolsRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIMCPAPI.SetTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetTools`: McpToolArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIMCPAPI.SetTools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **int32** | Identifier of the room containing the MCP server. | 
**serverId** | **string** | Unique identifier of the MCP server whose tools are being configured. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetToolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setMcpToolsRequestBody** | [**SetMcpToolsRequestBody**](SetMcpToolsRequestBody.md) | Tool configuration parameters. | 

### Return type

[**McpToolArrayWrapper**](McpToolArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServer

> McpServerWrapper UpdateServer(ctx, id).UpdateServerRequestBody(updateServerRequestBody).Execute()

Update a custom MCP server



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-server/).

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
	id := "00000000-0000-0000-0000-000000000000" // string | Unique identifier of the MCP server to update.
	updateServerRequestBody := *openapiclient.NewUpdateServerRequestBody() // UpdateServerRequestBody | Updated server configuration fields.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIMCPAPI.UpdateServer(context.Background(), id).UpdateServerRequestBody(updateServerRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIMCPAPI.UpdateServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServer`: McpServerWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIMCPAPI.UpdateServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the MCP server to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServerRequestBody** | [**UpdateServerRequestBody**](UpdateServerRequestBody.md) | Updated server configuration fields. | 

### Return type

[**McpServerWrapper**](McpServerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

