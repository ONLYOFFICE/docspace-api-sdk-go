# \PeopleGuestsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveGuestShareLink**](PeopleGuestsAPI.md#ApproveGuestShareLink) | **Post** /api/2.0/people/guests/share/approve | Approve a guest sharing link
[**DeleteGuests**](PeopleGuestsAPI.md#DeleteGuests) | **Delete** /api/2.0/people/guests | Remove guest relations



## ApproveGuestShareLink

> EmployeeFullWrapper ApproveGuestShareLink(ctx).EmailMemberRequestDto(emailMemberRequestDto).Execute()

Approve a guest sharing link



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/approve-guest-share-link/).

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
	resp, r, err := apiClient.PeopleGuestsAPI.ApproveGuestShareLink(context.Background()).EmailMemberRequestDto(emailMemberRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleGuestsAPI.ApproveGuestShareLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveGuestShareLink`: EmployeeFullWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleGuestsAPI.ApproveGuestShareLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApproveGuestShareLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailMemberRequestDto** | [**EmailMemberRequestDto**](EmailMemberRequestDto.md) |  | 

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


## DeleteGuests

> DeleteGuests(ctx).UpdateMembersRequestDto(updateMembersRequestDto).Execute()

Remove guest relations



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-guests/).

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
	updateMembersRequestDto := *openapiclient.NewUpdateMembersRequestDto() // UpdateMembersRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PeopleGuestsAPI.DeleteGuests(context.Background()).UpdateMembersRequestDto(updateMembersRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleGuestsAPI.DeleteGuests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGuestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateMembersRequestDto** | [**UpdateMembersRequestDto**](UpdateMembersRequestDto.md) |  | 

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

