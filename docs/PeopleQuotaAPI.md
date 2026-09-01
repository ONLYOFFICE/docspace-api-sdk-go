# \PeopleQuotaAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResetUsersQuota**](PeopleQuotaAPI.md#ResetUsersQuota) | **Put** /api/2.0/people/resetquota | Reset a user quota limit
[**UpdateUserQuota**](PeopleQuotaAPI.md#UpdateUserQuota) | **Put** /api/2.0/people/userquota | Change a user quota limit



## ResetUsersQuota

> EmployeeFullArrayWrapper ResetUsersQuota(ctx).UpdateMembersQuotaRequestDto(updateMembersQuotaRequestDto).Execute()

Reset a user quota limit



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/reset-users-quota/).

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
	updateMembersQuotaRequestDto := *openapiclient.NewUpdateMembersQuotaRequestDto() // UpdateMembersQuotaRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleQuotaAPI.ResetUsersQuota(context.Background()).UpdateMembersQuotaRequestDto(updateMembersQuotaRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleQuotaAPI.ResetUsersQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetUsersQuota`: EmployeeFullArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleQuotaAPI.ResetUsersQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetUsersQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateMembersQuotaRequestDto** | [**UpdateMembersQuotaRequestDto**](UpdateMembersQuotaRequestDto.md) |  | 

### Return type

[**EmployeeFullArrayWrapper**](EmployeeFullArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserQuota

> EmployeeFullArrayWrapper UpdateUserQuota(ctx).UpdateMembersQuotaRequestDto(updateMembersQuotaRequestDto).Execute()

Change a user quota limit



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-user-quota/).

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
	updateMembersQuotaRequestDto := *openapiclient.NewUpdateMembersQuotaRequestDto() // UpdateMembersQuotaRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleQuotaAPI.UpdateUserQuota(context.Background()).UpdateMembersQuotaRequestDto(updateMembersQuotaRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleQuotaAPI.UpdateUserQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserQuota`: EmployeeFullArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleQuotaAPI.UpdateUserQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateMembersQuotaRequestDto** | [**UpdateMembersQuotaRequestDto**](UpdateMembersQuotaRequestDto.md) |  | 

### Return type

[**EmployeeFullArrayWrapper**](EmployeeFullArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

