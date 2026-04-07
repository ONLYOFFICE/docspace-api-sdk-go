# \PeoplePasswordAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeUserPassword**](PeoplePasswordAPI.md#ChangeUserPassword) | **Put** /api/2.0/people/{userid}/password | Change a user password
[**SendUserPassword**](PeoplePasswordAPI.md#SendUserPassword) | **Post** /api/2.0/people/password | Remind a user password



## ChangeUserPassword

> EmployeeFullWrapper ChangeUserPassword(ctx, userid).ChangePasswordRequest(changePasswordRequest).Execute()

Change a user password



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/change-user-password/).

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
	changePasswordRequest := *openapiclient.NewChangePasswordRequest() // ChangePasswordRequest | The request parameters for updating a user password.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeoplePasswordAPI.ChangeUserPassword(context.Background(), userid).ChangePasswordRequest(changePasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeoplePasswordAPI.ChangeUserPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeUserPassword`: EmployeeFullWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeoplePasswordAPI.ChangeUserPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changePasswordRequest** | [**ChangePasswordRequest**](ChangePasswordRequest.md) | The request parameters for updating a user password. | 

### Return type

[**EmployeeFullWrapper**](EmployeeFullWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendUserPassword

> StringWrapper SendUserPassword(ctx).EmailMemberRequestDto(emailMemberRequestDto).Execute()

Remind a user password



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/send-user-password/).

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
	emailMemberRequestDto := *openapiclient.NewEmailMemberRequestDto("john.doe@example.com") // EmailMemberRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeoplePasswordAPI.SendUserPassword(context.Background()).EmailMemberRequestDto(emailMemberRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeoplePasswordAPI.SendUserPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendUserPassword`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeoplePasswordAPI.SendUserPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailMemberRequestDto** | [**EmailMemberRequestDto**](EmailMemberRequestDto.md) |  | 

### Return type

[**StringWrapper**](StringWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

