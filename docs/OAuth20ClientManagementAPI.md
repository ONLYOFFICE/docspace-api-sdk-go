# \OAuth20ClientManagementAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeActivation**](OAuth20ClientManagementAPI.md#ChangeActivation) | **Patch** /api/2.0/oauth2/clients/{clientId}/activation | Change client activation status
[**CreateClient**](OAuth20ClientManagementAPI.md#CreateClient) | **Post** /api/2.0/oauth2/clients | Create a new OAuth2 client
[**DeleteClient**](OAuth20ClientManagementAPI.md#DeleteClient) | **Delete** /api/2.0/oauth2/clients/{clientId} | Delete an OAuth2 client
[**DeleteTenantClients**](OAuth20ClientManagementAPI.md#DeleteTenantClients) | **Delete** /api/2.0/oauth2/clients/tenant | Delete all tenant OAuth2 clients
[**DeleteUserClients**](OAuth20ClientManagementAPI.md#DeleteUserClients) | **Delete** /api/2.0/oauth2/clients | Delete all user OAuth2 clients
[**RegenerateSecret**](OAuth20ClientManagementAPI.md#RegenerateSecret) | **Patch** /api/2.0/oauth2/clients/{clientId}/regenerate | Regenerate client secret
[**RevokeUserClient**](OAuth20ClientManagementAPI.md#RevokeUserClient) | **Delete** /api/2.0/oauth2/clients/{clientId}/revoke | Revoke client consent
[**UpdateClient**](OAuth20ClientManagementAPI.md#UpdateClient) | **Put** /api/2.0/oauth2/clients/{clientId} | Update an existing OAuth2 client



## ChangeActivation

> ChangeActivation(ctx, clientId).ChangeClientActivationRequest(changeClientActivationRequest).Execute()

Change client activation status



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/change-activation/).

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
	clientId := "6c7cf17b-1bd3-47d5-94c6-be2d3570e168" // string | ID of the client to change activation for
	changeClientActivationRequest := *openapiclient.NewChangeClientActivationRequest(true) // ChangeClientActivationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OAuth20ClientManagementAPI.ChangeActivation(context.Background(), clientId).ChangeClientActivationRequest(changeClientActivationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20ClientManagementAPI.ChangeActivation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | ID of the client to change activation for | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeActivationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeClientActivationRequest** | [**ChangeClientActivationRequest**](ChangeClientActivationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[x-signature](../README.md#x-signature)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateClient

> ClientResponse CreateClient(ctx).CreateClientRequest(createClientRequest).Execute()

Create a new OAuth2 client



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-client/).

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
	createClientRequest := *openapiclient.NewCreateClientRequest("Example Client", "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg==", []string{"files:read"}, "http://example.com", "http://example.com/terms", "http://example.com/policy", []string{"http://example.com/redirect"}, []string{"http://example.com"}, "http://example.com/logout") // CreateClientRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth20ClientManagementAPI.CreateClient(context.Background()).CreateClientRequest(createClientRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20ClientManagementAPI.CreateClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateClient`: ClientResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuth20ClientManagementAPI.CreateClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createClientRequest** | [**CreateClientRequest**](CreateClientRequest.md) |  | 

### Return type

[**ClientResponse**](ClientResponse.md)

### Authorization

[x-signature](../README.md#x-signature)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClient

> DeleteClient(ctx, clientId).Execute()

Delete an OAuth2 client



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-client/).

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
	clientId := "6c7cf17b-1bd3-47d5-94c6-be2d3570e168" // string | ID of the client to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OAuth20ClientManagementAPI.DeleteClient(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20ClientManagementAPI.DeleteClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | ID of the client to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[x-signature](../README.md#x-signature)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTenantClients

> DeleteTenantClients(ctx).Execute()

Delete all tenant OAuth2 clients



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-tenant-clients/).

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
	r, err := apiClient.OAuth20ClientManagementAPI.DeleteTenantClients(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20ClientManagementAPI.DeleteTenantClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantClientsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[x-signature](../README.md#x-signature)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserClients

> DeleteUserClients(ctx).Execute()

Delete all user OAuth2 clients



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-user-clients/).

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
	r, err := apiClient.OAuth20ClientManagementAPI.DeleteUserClients(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20ClientManagementAPI.DeleteUserClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserClientsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[x-signature](../README.md#x-signature)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegenerateSecret

> ClientSecretResponse RegenerateSecret(ctx, clientId).Execute()

Regenerate client secret



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/regenerate-secret/).

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
	clientId := "6c7cf17b-1bd3-47d5-94c6-be2d3570e168" // string | ID of the client to regenerate secret for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth20ClientManagementAPI.RegenerateSecret(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20ClientManagementAPI.RegenerateSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegenerateSecret`: ClientSecretResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuth20ClientManagementAPI.RegenerateSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | ID of the client to regenerate secret for | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientSecretResponse**](ClientSecretResponse.md)

### Authorization

[x-signature](../README.md#x-signature)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeUserClient

> RevokeUserClient(ctx, clientId).Execute()

Revoke client consent



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/revoke-user-client/).

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
	clientId := "6c7cf17b-1bd3-47d5-94c6-be2d3570e168" // string | ID of the client to revoke consent for

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OAuth20ClientManagementAPI.RevokeUserClient(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20ClientManagementAPI.RevokeUserClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | ID of the client to revoke consent for | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeUserClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[x-signature](../README.md#x-signature)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClient

> UpdateClient(ctx, clientId).UpdateClientRequest(updateClientRequest).Execute()

Update an existing OAuth2 client



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-client/).

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
	clientId := "6c7cf17b-1bd3-47d5-94c6-be2d3570e168" // string | ID of the client to update
	updateClientRequest := *openapiclient.NewUpdateClientRequest("Updated Client", "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg==", []string{"files:read"}, []string{"http://example.com"}, []string{"https://example.com/callback"}) // UpdateClientRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OAuth20ClientManagementAPI.UpdateClient(context.Background(), clientId).UpdateClientRequest(updateClientRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20ClientManagementAPI.UpdateClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | ID of the client to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateClientRequest** | [**UpdateClientRequest**](UpdateClientRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[x-signature](../README.md#x-signature)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

