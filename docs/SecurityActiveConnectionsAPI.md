# \SecurityActiveConnectionsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllActiveConnections**](SecurityActiveConnectionsAPI.md#GetAllActiveConnections) | **Get** /api/2.0/security/activeconnections | Get active connections
[**LogOutActiveConnection**](SecurityActiveConnectionsAPI.md#LogOutActiveConnection) | **Put** /api/2.0/security/activeconnections/logout/{loginEventId} | Log out from the connection
[**LogOutAllActiveConnectionsChangePassword**](SecurityActiveConnectionsAPI.md#LogOutAllActiveConnectionsChangePassword) | **Put** /api/2.0/security/activeconnections/logoutallchangepassword | Log out and change password
[**LogOutAllActiveConnectionsForUser**](SecurityActiveConnectionsAPI.md#LogOutAllActiveConnectionsForUser) | **Put** /api/2.0/security/activeconnections/logoutall/{userId} | Log out for the user by ID
[**LogOutAllExceptThisConnection**](SecurityActiveConnectionsAPI.md#LogOutAllExceptThisConnection) | **Put** /api/2.0/security/activeconnections/logoutallexceptthis | Log out from all connections except the current one



## GetAllActiveConnections

> ActiveConnectionsWrapper GetAllActiveConnections(ctx).Execute()

Get active connections



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-all-active-connections/).

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
	resp, r, err := apiClient.SecurityActiveConnectionsAPI.GetAllActiveConnections(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityActiveConnectionsAPI.GetAllActiveConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllActiveConnections`: ActiveConnectionsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityActiveConnectionsAPI.GetAllActiveConnections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllActiveConnectionsRequest struct via the builder pattern


### Return type

[**ActiveConnectionsWrapper**](ActiveConnectionsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogOutActiveConnection

> BooleanWrapper LogOutActiveConnection(ctx, loginEventId).Execute()

Log out from the connection



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/log-out-active-connection/).

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
	loginEventId := int32(12345) // int32 | The ID of the specific login event.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityActiveConnectionsAPI.LogOutActiveConnection(context.Background(), loginEventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityActiveConnectionsAPI.LogOutActiveConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LogOutActiveConnection`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityActiveConnectionsAPI.LogOutActiveConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loginEventId** | **int32** | The ID of the specific login event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogOutActiveConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogOutAllActiveConnectionsChangePassword

> StringWrapper LogOutAllActiveConnectionsChangePassword(ctx).Execute()

Log out and change password



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/log-out-all-active-connections-change-password/).

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
	resp, r, err := apiClient.SecurityActiveConnectionsAPI.LogOutAllActiveConnectionsChangePassword(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityActiveConnectionsAPI.LogOutAllActiveConnectionsChangePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LogOutAllActiveConnectionsChangePassword`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityActiveConnectionsAPI.LogOutAllActiveConnectionsChangePassword`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogOutAllActiveConnectionsChangePasswordRequest struct via the builder pattern


### Return type

[**StringWrapper**](StringWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogOutAllActiveConnectionsForUser

> LogOutAllActiveConnectionsForUser(ctx, userId).Execute()

Log out for the user by ID



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/log-out-all-active-connections-for-user/).

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
	userId := "00000000-0000-0000-0000-000000000000" // string | The user ID extracted from the route parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityActiveConnectionsAPI.LogOutAllActiveConnectionsForUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityActiveConnectionsAPI.LogOutAllActiveConnectionsForUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID extracted from the route parameters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogOutAllActiveConnectionsForUserRequest struct via the builder pattern


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


## LogOutAllExceptThisConnection

> StringWrapper LogOutAllExceptThisConnection(ctx).Execute()

Log out from all connections except the current one



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/log-out-all-except-this-connection/).

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
	resp, r, err := apiClient.SecurityActiveConnectionsAPI.LogOutAllExceptThisConnection(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityActiveConnectionsAPI.LogOutAllExceptThisConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LogOutAllExceptThisConnection`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityActiveConnectionsAPI.LogOutAllExceptThisConnection`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogOutAllExceptThisConnectionRequest struct via the builder pattern


### Return type

[**StringWrapper**](StringWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

