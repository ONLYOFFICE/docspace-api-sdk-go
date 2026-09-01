# \PeopleUserDataAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeletePersonalFolderProgress**](PeopleUserDataAPI.md#GetDeletePersonalFolderProgress) | **Get** /api/2.0/people/delete/personal/progress | Get the progress of deleting the personal folder
[**GetReassignProgress**](PeopleUserDataAPI.md#GetReassignProgress) | **Get** /api/2.0/people/reassign/progress/{userid} | Get the reassignment progress
[**GetRemoveProgress**](PeopleUserDataAPI.md#GetRemoveProgress) | **Get** /api/2.0/people/remove/progress/{userid} | Get the deletion progress
[**NecessaryReassign**](PeopleUserDataAPI.md#NecessaryReassign) | **Get** /api/2.0/people/reassign/necessary | Check data for reassignment need
[**SendInstructionsToDelete**](PeopleUserDataAPI.md#SendInstructionsToDelete) | **Put** /api/2.0/people/self/delete | Send the deletion instructions
[**StartDeletePersonalFolder**](PeopleUserDataAPI.md#StartDeletePersonalFolder) | **Post** /api/2.0/people/delete/personal/start | Delete the personal folder
[**StartReassign**](PeopleUserDataAPI.md#StartReassign) | **Post** /api/2.0/people/reassign/start | Start the data reassignment
[**StartRemove**](PeopleUserDataAPI.md#StartRemove) | **Post** /api/2.0/people/remove/start | Start the data deletion
[**TerminateReassign**](PeopleUserDataAPI.md#TerminateReassign) | **Put** /api/2.0/people/reassign/terminate | Terminate the data reassignment
[**TerminateRemove**](PeopleUserDataAPI.md#TerminateRemove) | **Put** /api/2.0/people/remove/terminate | Terminate the data deletion



## GetDeletePersonalFolderProgress

> TaskProgressResponseWrapper GetDeletePersonalFolderProgress(ctx).Execute()

Get the progress of deleting the personal folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-delete-personal-folder-progress/).

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
	resp, r, err := apiClient.PeopleUserDataAPI.GetDeletePersonalFolderProgress(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleUserDataAPI.GetDeletePersonalFolderProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeletePersonalFolderProgress`: TaskProgressResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleUserDataAPI.GetDeletePersonalFolderProgress`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeletePersonalFolderProgressRequest struct via the builder pattern


### Return type

[**TaskProgressResponseWrapper**](TaskProgressResponseWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReassignProgress

> TaskProgressResponseWrapper GetReassignProgress(ctx, userid).Execute()

Get the reassignment progress



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-reassign-progress/).

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
	userid := "00000000-0000-0000-0000-000000000000" // string | The user ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleUserDataAPI.GetReassignProgress(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleUserDataAPI.GetReassignProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReassignProgress`: TaskProgressResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleUserDataAPI.GetReassignProgress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReassignProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskProgressResponseWrapper**](TaskProgressResponseWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoveProgress

> TaskProgressResponseWrapper GetRemoveProgress(ctx, userid).Execute()

Get the deletion progress



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-remove-progress/).

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
	userid := "00000000-0000-0000-0000-000000000000" // string | The user ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleUserDataAPI.GetRemoveProgress(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleUserDataAPI.GetRemoveProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoveProgress`: TaskProgressResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleUserDataAPI.GetRemoveProgress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoveProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskProgressResponseWrapper**](TaskProgressResponseWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NecessaryReassign

> BooleanWrapper NecessaryReassign(ctx).UserId(userId).Type_(type_).Execute()

Check data for reassignment need



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/necessary-reassign/).

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
	userId := "00000000-0000-0000-0000-000000000000" // string | The user ID. (optional)
	type_ := openapiclient.EmployeeType("All") // EmployeeType | The expected user type. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleUserDataAPI.NecessaryReassign(context.Background()).UserId(userId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleUserDataAPI.NecessaryReassign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NecessaryReassign`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleUserDataAPI.NecessaryReassign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNecessaryReassignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | The user ID. | 
 **type_** | [**EmployeeType**](EmployeeType.md) | The expected user type. | 

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


## SendInstructionsToDelete

> StringWrapper SendInstructionsToDelete(ctx).Execute()

Send the deletion instructions



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/send-instructions-to-delete/).

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
	resp, r, err := apiClient.PeopleUserDataAPI.SendInstructionsToDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleUserDataAPI.SendInstructionsToDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendInstructionsToDelete`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleUserDataAPI.SendInstructionsToDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSendInstructionsToDeleteRequest struct via the builder pattern


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


## StartDeletePersonalFolder

> TaskProgressResponseWrapper StartDeletePersonalFolder(ctx).Execute()

Delete the personal folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/start-delete-personal-folder/).

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
	resp, r, err := apiClient.PeopleUserDataAPI.StartDeletePersonalFolder(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleUserDataAPI.StartDeletePersonalFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartDeletePersonalFolder`: TaskProgressResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleUserDataAPI.StartDeletePersonalFolder`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStartDeletePersonalFolderRequest struct via the builder pattern


### Return type

[**TaskProgressResponseWrapper**](TaskProgressResponseWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartReassign

> TaskProgressResponseWrapper StartReassign(ctx).StartReassignRequestDto(startReassignRequestDto).Execute()

Start the data reassignment



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/start-reassign/).

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
	startReassignRequestDto := *openapiclient.NewStartReassignRequestDto("00000000-0000-0000-0000-000000000000", "11111111-1111-1111-1111-111111111111") // StartReassignRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleUserDataAPI.StartReassign(context.Background()).StartReassignRequestDto(startReassignRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleUserDataAPI.StartReassign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartReassign`: TaskProgressResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleUserDataAPI.StartReassign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartReassignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startReassignRequestDto** | [**StartReassignRequestDto**](StartReassignRequestDto.md) |  | 

### Return type

[**TaskProgressResponseWrapper**](TaskProgressResponseWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartRemove

> TaskProgressResponseWrapper StartRemove(ctx).TerminateRequestDto(terminateRequestDto).Execute()

Start the data deletion



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/start-remove/).

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
	terminateRequestDto := *openapiclient.NewTerminateRequestDto("00000000-0000-0000-0000-000000000000") // TerminateRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleUserDataAPI.StartRemove(context.Background()).TerminateRequestDto(terminateRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleUserDataAPI.StartRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartRemove`: TaskProgressResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleUserDataAPI.StartRemove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **terminateRequestDto** | [**TerminateRequestDto**](TerminateRequestDto.md) |  | 

### Return type

[**TaskProgressResponseWrapper**](TaskProgressResponseWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateReassign

> TaskProgressResponseWrapper TerminateReassign(ctx).TerminateRequestDto(terminateRequestDto).Execute()

Terminate the data reassignment



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/terminate-reassign/).

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
	terminateRequestDto := *openapiclient.NewTerminateRequestDto("00000000-0000-0000-0000-000000000000") // TerminateRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleUserDataAPI.TerminateReassign(context.Background()).TerminateRequestDto(terminateRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleUserDataAPI.TerminateReassign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TerminateReassign`: TaskProgressResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleUserDataAPI.TerminateReassign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTerminateReassignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **terminateRequestDto** | [**TerminateRequestDto**](TerminateRequestDto.md) |  | 

### Return type

[**TaskProgressResponseWrapper**](TaskProgressResponseWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateRemove

> TerminateRemove(ctx).TerminateRequestDto(terminateRequestDto).Execute()

Terminate the data deletion



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/terminate-remove/).

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
	terminateRequestDto := *openapiclient.NewTerminateRequestDto("00000000-0000-0000-0000-000000000000") // TerminateRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PeopleUserDataAPI.TerminateRemove(context.Background()).TerminateRequestDto(terminateRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleUserDataAPI.TerminateRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTerminateRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **terminateRequestDto** | [**TerminateRequestDto**](TerminateRequestDto.md) |  | 

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

