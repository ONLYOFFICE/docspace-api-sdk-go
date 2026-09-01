# \PortalUsersAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInvitationLink**](PortalUsersAPI.md#CreateInvitationLink) | **Post** /api/2.0/portal/users/invitationlink | Create an invitation link
[**DeleteInvitationLink**](PortalUsersAPI.md#DeleteInvitationLink) | **Delete** /api/2.0/portal/users/invitationlink | Deletes an invitation link.
[**GetInvitationLink**](PortalUsersAPI.md#GetInvitationLink) | **Get** /api/2.0/portal/users/invite/{employeeType} | Get an invitation link
[**GetInvitationLinkByEmployeeType**](PortalUsersAPI.md#GetInvitationLinkByEmployeeType) | **Get** /api/2.0/portal/users/invitationlink/{employeeType} | Get an invitation link
[**GetPortalUsersCount**](PortalUsersAPI.md#GetPortalUsersCount) | **Get** /api/2.0/portal/userscount | Get a number of portal users
[**GetUserById**](PortalUsersAPI.md#GetUserById) | **Get** /api/2.0/portal/users/{userID} | Get a user by ID
[**MarkGiftMessageAsRead**](PortalUsersAPI.md#MarkGiftMessageAsRead) | **Post** /api/2.0/portal/present/mark | Mark a gift message as read
[**SendCongratulations**](PortalUsersAPI.md#SendCongratulations) | **Post** /api/2.0/portal/sendcongratulations | Send congratulations
[**UpdateInvitationLink**](PortalUsersAPI.md#UpdateInvitationLink) | **Put** /api/2.0/portal/users/invitationlink | Update an invitation link



## CreateInvitationLink

> InvitationLinkWrapper CreateInvitationLink(ctx).InvitationLinkCreateRequestDto(invitationLinkCreateRequestDto).Execute()

Create an invitation link



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-invitation-link/).

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
	invitationLinkCreateRequestDto := *openapiclient.NewInvitationLinkCreateRequestDto(openapiclient.EmployeeType("All")) // InvitationLinkCreateRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalUsersAPI.CreateInvitationLink(context.Background()).InvitationLinkCreateRequestDto(invitationLinkCreateRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalUsersAPI.CreateInvitationLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvitationLink`: InvitationLinkWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalUsersAPI.CreateInvitationLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvitationLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationLinkCreateRequestDto** | [**InvitationLinkCreateRequestDto**](InvitationLinkCreateRequestDto.md) |  | 

### Return type

[**InvitationLinkWrapper**](InvitationLinkWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInvitationLink

> StringWrapper DeleteInvitationLink(ctx).InvitationLinkDeleteRequestDto(invitationLinkDeleteRequestDto).Execute()

Deletes an invitation link.



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-invitation-link/).

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
	invitationLinkDeleteRequestDto := *openapiclient.NewInvitationLinkDeleteRequestDto("00000000-0000-0000-0000-000000000000") // InvitationLinkDeleteRequestDto | The data transfer object containing the details of the invitation link to be deleted. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalUsersAPI.DeleteInvitationLink(context.Background()).InvitationLinkDeleteRequestDto(invitationLinkDeleteRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalUsersAPI.DeleteInvitationLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInvitationLink`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalUsersAPI.DeleteInvitationLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInvitationLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationLinkDeleteRequestDto** | [**InvitationLinkDeleteRequestDto**](InvitationLinkDeleteRequestDto.md) | The data transfer object containing the details of the invitation link to be deleted. | 

### Return type

[**StringWrapper**](StringWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvitationLink

> StringWrapper GetInvitationLink(ctx, employeeType).Execute()

Get an invitation link



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-invitation-link/).

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
	employeeType := openapiclient.EmployeeType("All") // EmployeeType | The type of employee role for the invitation link (DocSpaceAdmin, RoomAdmin or User).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalUsersAPI.GetInvitationLink(context.Background(), employeeType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalUsersAPI.GetInvitationLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvitationLink`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalUsersAPI.GetInvitationLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employeeType** | [**EmployeeType**](.md) | The type of employee role for the invitation link (DocSpaceAdmin, RoomAdmin or User). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvitationLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetInvitationLinkByEmployeeType

> InvitationLinkWrapper GetInvitationLinkByEmployeeType(ctx, employeeType).Execute()

Get an invitation link



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-invitation-link-by-employee-type/).

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
	employeeType := openapiclient.EmployeeType("All") // EmployeeType | The type of employee role for the invitation link (DocSpaceAdmin, RoomAdmin or User).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalUsersAPI.GetInvitationLinkByEmployeeType(context.Background(), employeeType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalUsersAPI.GetInvitationLinkByEmployeeType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvitationLinkByEmployeeType`: InvitationLinkWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalUsersAPI.GetInvitationLinkByEmployeeType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**employeeType** | [**EmployeeType**](.md) | The type of employee role for the invitation link (DocSpaceAdmin, RoomAdmin or User). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvitationLinkByEmployeeTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvitationLinkWrapper**](InvitationLinkWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortalUsersCount

> Int64Wrapper GetPortalUsersCount(ctx).Execute()

Get a number of portal users



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-portal-users-count/).

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
	resp, r, err := apiClient.PortalUsersAPI.GetPortalUsersCount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalUsersAPI.GetPortalUsersCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortalUsersCount`: Int64Wrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalUsersAPI.GetPortalUsersCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortalUsersCountRequest struct via the builder pattern


### Return type

[**Int64Wrapper**](Int64Wrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserById

> UserInfoWrapper GetUserById(ctx, userID).Execute()

Get a user by ID



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-user-by-id/).

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
	userID := "00000000-0000-0000-0000-000000000000" // string | The user ID extracted from the route parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalUsersAPI.GetUserById(context.Background(), userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalUsersAPI.GetUserById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserById`: UserInfoWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalUsersAPI.GetUserById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | The user ID extracted from the route parameters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserInfoWrapper**](UserInfoWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkGiftMessageAsRead

> MarkGiftMessageAsRead(ctx).Execute()

Mark a gift message as read



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/mark-gift-message-as-read/).

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
	r, err := apiClient.PortalUsersAPI.MarkGiftMessageAsRead(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalUsersAPI.MarkGiftMessageAsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMarkGiftMessageAsReadRequest struct via the builder pattern


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


## SendCongratulations

> SendCongratulations(ctx).Userid(userid).Key(key).Execute()

Send congratulations



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/send-congratulations/).

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
	userid := "00000000-0000-0000-0000-000000000000" // string | The user ID to receive the congratulatory message.
	key := "birthday" // string | The template identifier or email configuration key.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PortalUsersAPI.SendCongratulations(context.Background()).Userid(userid).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalUsersAPI.SendCongratulations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendCongratulationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userid** | **string** | The user ID to receive the congratulatory message. | 
 **key** | **string** | The template identifier or email configuration key. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvitationLink

> InvitationLinkWrapper UpdateInvitationLink(ctx).InvitationLinkUpdateRequestDto(invitationLinkUpdateRequestDto).Execute()

Update an invitation link



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-invitation-link/).

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
	invitationLinkUpdateRequestDto := *openapiclient.NewInvitationLinkUpdateRequestDto("00000000-0000-0000-0000-000000000000") // InvitationLinkUpdateRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalUsersAPI.UpdateInvitationLink(context.Background()).InvitationLinkUpdateRequestDto(invitationLinkUpdateRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalUsersAPI.UpdateInvitationLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInvitationLink`: InvitationLinkWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalUsersAPI.UpdateInvitationLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvitationLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationLinkUpdateRequestDto** | [**InvitationLinkUpdateRequestDto**](InvitationLinkUpdateRequestDto.md) |  | 

### Return type

[**InvitationLinkWrapper**](InvitationLinkWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

