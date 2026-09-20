# \PeopleEmailAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeUserEmail**](PeopleEmailAPI.md#ChangeUserEmail) | **Put** /api/2.0/people/{userid}/email | Change a user email
[**SendEmailChangeInstructions**](PeopleEmailAPI.md#SendEmailChangeInstructions) | **Post** /api/2.0/people/email | Send instructions to change email



## ChangeUserEmail

> EmployeeFullWrapper ChangeUserEmail(ctx, userid).ChangeEmailRequest(changeEmailRequest).Execute()

Change a user email



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/change-user-email/).

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
	userid := "00000000-0000-0000-0000-000000000000" // string | The ID of the account whose address is set, taken from the route. It has to match the account the  confirmation token was issued for, and the account has to be active.
	changeEmailRequest := *openapiclient.NewChangeEmailRequest() // ChangeEmailRequest | The new address, in plain text or in the encrypted form the confirmation link carries.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleEmailAPI.ChangeUserEmail(context.Background(), userid).ChangeEmailRequest(changeEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleEmailAPI.ChangeUserEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeUserEmail`: EmployeeFullWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleEmailAPI.ChangeUserEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** | The ID of the account whose address is set, taken from the route. It has to match the account the  confirmation token was issued for, and the account has to be active. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeUserEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeEmailRequest** | [**ChangeEmailRequest**](ChangeEmailRequest.md) | The new address, in plain text or in the encrypted form the confirmation link carries. | 

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


## SendEmailChangeInstructions

> StringWrapper SendEmailChangeInstructions(ctx).UpdateMemberRequestDto(updateMemberRequestDto).Execute()

Send instructions to change email



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/send-email-change-instructions/).

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
	updateMemberRequestDto := *openapiclient.NewUpdateMemberRequestDto() // UpdateMemberRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleEmailAPI.SendEmailChangeInstructions(context.Background()).UpdateMemberRequestDto(updateMemberRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleEmailAPI.SendEmailChangeInstructions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendEmailChangeInstructions`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleEmailAPI.SendEmailChangeInstructions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendEmailChangeInstructionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateMemberRequestDto** | [**UpdateMemberRequestDto**](UpdateMemberRequestDto.md) |  | 

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

