# \OAuth20ClientQueryingAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClient**](OAuth20ClientQueryingAPI.md#GetClient) | **Get** /api/2.0/clients/{clientId} | Get client details
[**GetClientInfo**](OAuth20ClientQueryingAPI.md#GetClientInfo) | **Get** /api/2.0/clients/{clientId}/info | Retrieves detailed information for a specific client
[**GetClients**](OAuth20ClientQueryingAPI.md#GetClients) | **Get** /api/2.0/clients | List clients
[**GetClientsInfo**](OAuth20ClientQueryingAPI.md#GetClientsInfo) | **Get** /api/2.0/clients/info | Retrieves a pageable list of client information
[**GetConsents**](OAuth20ClientQueryingAPI.md#GetConsents) | **Get** /api/2.0/clients/consents | Retrieves a pageable list of consents
[**GetPublicClientInfo**](OAuth20ClientQueryingAPI.md#GetPublicClientInfo) | **Get** /api/2.0/clients/{clientId}/public/info | Handles the GET request for public client information



## GetClient

> ClientResponse GetClient(ctx, clientId).Execute()

Get client details



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-client/).

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
	clientId := "6c7cf17b-1bd3-47d5-94c6-be2d3570e168" // string | ID of the client to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth20ClientQueryingAPI.GetClient(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20ClientQueryingAPI.GetClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClient`: ClientResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuth20ClientQueryingAPI.GetClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | ID of the client to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientResponse**](ClientResponse.md)

### Authorization

[x-signature](../README.md#x-signature)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientInfo

> ClientInfoResponse GetClientInfo(ctx, clientId).Execute()

Retrieves detailed information for a specific client



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-client-info/).

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
	clientId := "6c7cf17b-1bd3-47d5-94c6-be2d3570e168" // string | ID of the client to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth20ClientQueryingAPI.GetClientInfo(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20ClientQueryingAPI.GetClientInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientInfo`: ClientInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuth20ClientQueryingAPI.GetClientInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | ID of the client to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientInfoResponse**](ClientInfoResponse.md)

### Authorization

[x-signature](../README.md#x-signature)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClients

> PageableResponse GetClients(ctx).Limit(limit).LastClientId(lastClientId).LastCreatedOn(lastCreatedOn).Execute()

List clients



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-clients/).

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
	limit := int32(1) // int32 | Pagination limit (default to 30)
	lastClientId := "6c7cf17b-1bd3-47d5-94c6-be2d3570e168" // string | ID of the last retrieved client (optional)
	lastCreatedOn := time.Now() // time.Time | Date of the last retrieved client (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth20ClientQueryingAPI.GetClients(context.Background()).Limit(limit).LastClientId(lastClientId).LastCreatedOn(lastCreatedOn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20ClientQueryingAPI.GetClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClients`: PageableResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuth20ClientQueryingAPI.GetClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Pagination limit | [default to 30]
 **lastClientId** | **string** | ID of the last retrieved client | 
 **lastCreatedOn** | **time.Time** | Date of the last retrieved client | 

### Return type

[**PageableResponse**](PageableResponse.md)

### Authorization

[x-signature](../README.md#x-signature)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientsInfo

> PageableResponseClientInfoResponse GetClientsInfo(ctx).Limit(limit).LastClientId(lastClientId).LastCreatedOn(lastCreatedOn).Execute()

Retrieves a pageable list of client information



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-clients-info/).

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
	limit := int32(1) // int32 | Pagination limit
	lastClientId := "6c7cf17b-1bd3-47d5-94c6-be2d3570e168" // string | ID of the last retrieved client (optional)
	lastCreatedOn := time.Now() // time.Time | Date of the last retrieved client (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth20ClientQueryingAPI.GetClientsInfo(context.Background()).Limit(limit).LastClientId(lastClientId).LastCreatedOn(lastCreatedOn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20ClientQueryingAPI.GetClientsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientsInfo`: PageableResponseClientInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuth20ClientQueryingAPI.GetClientsInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Pagination limit | 
 **lastClientId** | **string** | ID of the last retrieved client | 
 **lastCreatedOn** | **time.Time** | Date of the last retrieved client | 

### Return type

[**PageableResponseClientInfoResponse**](PageableResponseClientInfoResponse.md)

### Authorization

[x-signature](../README.md#x-signature)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsents

> PageableModificationResponse GetConsents(ctx).Limit(limit).LastModifiedOn(lastModifiedOn).Execute()

Retrieves a pageable list of consents



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-consents/).

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
	limit := int32(1) // int32 | Pagination limit
	lastModifiedOn := time.Now() // time.Time | Date of the last retrieved consent (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth20ClientQueryingAPI.GetConsents(context.Background()).Limit(limit).LastModifiedOn(lastModifiedOn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20ClientQueryingAPI.GetConsents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConsents`: PageableModificationResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuth20ClientQueryingAPI.GetConsents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConsentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Pagination limit | 
 **lastModifiedOn** | **time.Time** | Date of the last retrieved consent | 

### Return type

[**PageableModificationResponse**](PageableModificationResponse.md)

### Authorization

[x-signature](../README.md#x-signature)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicClientInfo

> ClientInfoResponse GetPublicClientInfo(ctx, clientId).Execute()

Handles the GET request for public client information

For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-public-client-info/).

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
	clientId := "6c7cf17b-1bd3-47d5-94c6-be2d3570e168" // string | ID of the client to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth20ClientQueryingAPI.GetPublicClientInfo(context.Background(), clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20ClientQueryingAPI.GetPublicClientInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicClientInfo`: ClientInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuth20ClientQueryingAPI.GetPublicClientInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** | ID of the client to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicClientInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientInfoResponse**](ClientInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

