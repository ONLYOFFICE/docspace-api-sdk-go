# \PeopleProfilesAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMember**](PeopleProfilesAPI.md#AddMember) | **Post** /api/2.0/people | Add a user
[**CheckUserExistsByEmail**](PeopleProfilesAPI.md#CheckUserExistsByEmail) | **Get** /api/2.0/people/exists | Check if a user exists by email
[**DeleteMember**](PeopleProfilesAPI.md#DeleteMember) | **Delete** /api/2.0/people/{userid} | Delete a user
[**DeleteProfile**](PeopleProfilesAPI.md#DeleteProfile) | **Delete** /api/2.0/people/@self | Delete my profile
[**GetAllProfiles**](PeopleProfilesAPI.md#GetAllProfiles) | **Get** /api/2.0/people | Get profiles
[**GetClaims**](PeopleProfilesAPI.md#GetClaims) | **Get** /api/2.0/people/tokendiagnostics | Get user claims
[**GetProfileByEmail**](PeopleProfilesAPI.md#GetProfileByEmail) | **Get** /api/2.0/people/email | Get a profile by user email
[**GetProfileByUserId**](PeopleProfilesAPI.md#GetProfileByUserId) | **Get** /api/2.0/people/{userid} | Get a profile by user ID
[**GetSelfProfile**](PeopleProfilesAPI.md#GetSelfProfile) | **Get** /api/2.0/people/@self | Get my profile
[**InviteUsers**](PeopleProfilesAPI.md#InviteUsers) | **Post** /api/2.0/people/invite | Invite users
[**RemoveUsers**](PeopleProfilesAPI.md#RemoveUsers) | **Put** /api/2.0/people/delete | Delete users
[**ResendUserInvites**](PeopleProfilesAPI.md#ResendUserInvites) | **Put** /api/2.0/people/invite | Resend activation emails
[**UpdateMember**](PeopleProfilesAPI.md#UpdateMember) | **Put** /api/2.0/people/{userid} | Update a user
[**UpdateMemberCulture**](PeopleProfilesAPI.md#UpdateMemberCulture) | **Put** /api/2.0/people/{userid}/culture | Update a user culture



## AddMember

> EmployeeFullWrapper AddMember(ctx).MemberRequestDto(memberRequestDto).Execute()

Add a user



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/add-member/).

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
	memberRequestDto := *openapiclient.NewMemberRequestDto() // MemberRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleProfilesAPI.AddMember(context.Background()).MemberRequestDto(memberRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleProfilesAPI.AddMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddMember`: EmployeeFullWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleProfilesAPI.AddMember`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **memberRequestDto** | [**MemberRequestDto**](MemberRequestDto.md) |  | 

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


## CheckUserExistsByEmail

> BooleanWrapper CheckUserExistsByEmail(ctx).Email(email).Encemail(encemail).Culture(culture).Execute()

Check if a user exists by email



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/check-user-exists-by-email/).

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
	email := "john.doe@example.com" // string | The user email address. (optional)
	encemail := "encrypted_email_string" // string | The user encrypted email address. (optional)
	culture := "en-US" // string | Culture (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleProfilesAPI.CheckUserExistsByEmail(context.Background()).Email(email).Encemail(encemail).Culture(culture).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleProfilesAPI.CheckUserExistsByEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckUserExistsByEmail`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleProfilesAPI.CheckUserExistsByEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckUserExistsByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | The user email address. | 
 **encemail** | **string** | The user encrypted email address. | 
 **culture** | **string** | Culture | 

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


## DeleteMember

> EmployeeFullWrapper DeleteMember(ctx, userid).Execute()

Delete a user



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-member/).

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
	resp, r, err := apiClient.PeopleProfilesAPI.DeleteMember(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleProfilesAPI.DeleteMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMember`: EmployeeFullWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleProfilesAPI.DeleteMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmployeeFullWrapper**](EmployeeFullWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProfile

> EmployeeFullWrapper DeleteProfile(ctx).Execute()

Delete my profile



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-profile/).

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
	resp, r, err := apiClient.PeopleProfilesAPI.DeleteProfile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleProfilesAPI.DeleteProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProfile`: EmployeeFullWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleProfilesAPI.DeleteProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProfileRequest struct via the builder pattern


### Return type

[**EmployeeFullWrapper**](EmployeeFullWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllProfiles

> EmployeeFullArrayWrapper GetAllProfiles(ctx).Count(count).StartIndex(startIndex).FilterBy(filterBy).SortBy(sortBy).SortOrder(sortOrder).FilterSeparator(filterSeparator).FilterValue(filterValue).Execute()

Get profiles



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-all-profiles/).

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
	count := int32(25) // int32 | The maximum number of items to be retrieved in the response. (optional)
	startIndex := int32(0) // int32 | The zero-based index of the first item to be retrieved in a filtered result set. (optional)
	filterBy := "displayName" // string | Specifies the filter criteria for user-related queries. (optional)
	sortBy := "displayName" // string | Specifies the property or field name by which the results should be sorted. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The order in which the results are sorted. (optional)
	filterSeparator := "," // string | The character or string used to separate multiple filter values in a filtering query. (optional)
	filterValue := "John" // string | The text value used as an additional filter criterion for profiles retrieval. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleProfilesAPI.GetAllProfiles(context.Background()).Count(count).StartIndex(startIndex).FilterBy(filterBy).SortBy(sortBy).SortOrder(sortOrder).FilterSeparator(filterSeparator).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleProfilesAPI.GetAllProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllProfiles`: EmployeeFullArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleProfilesAPI.GetAllProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | The maximum number of items to be retrieved in the response. | 
 **startIndex** | **int32** | The zero-based index of the first item to be retrieved in a filtered result set. | 
 **filterBy** | **string** | Specifies the filter criteria for user-related queries. | 
 **sortBy** | **string** | Specifies the property or field name by which the results should be sorted. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The order in which the results are sorted. | 
 **filterSeparator** | **string** | The character or string used to separate multiple filter values in a filtering query. | 
 **filterValue** | **string** | The text value used as an additional filter criterion for profiles retrieval. | 

### Return type

[**EmployeeFullArrayWrapper**](EmployeeFullArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClaims

> ObjectWrapper GetClaims(ctx).Execute()

Get user claims



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-claims/).

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
	resp, r, err := apiClient.PeopleProfilesAPI.GetClaims(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleProfilesAPI.GetClaims``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClaims`: ObjectWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleProfilesAPI.GetClaims`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClaimsRequest struct via the builder pattern


### Return type

[**ObjectWrapper**](ObjectWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfileByEmail

> EmployeeFullWrapper GetProfileByEmail(ctx).Email(email).Encemail(encemail).Culture(culture).Execute()

Get a profile by user email



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-profile-by-email/).

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
	email := "john.doe@example.com" // string | The user email address. (optional)
	encemail := "encrypted_email_string" // string | The user encrypted email address. (optional)
	culture := "en-US" // string | Culture (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleProfilesAPI.GetProfileByEmail(context.Background()).Email(email).Encemail(encemail).Culture(culture).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleProfilesAPI.GetProfileByEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfileByEmail`: EmployeeFullWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleProfilesAPI.GetProfileByEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | The user email address. | 
 **encemail** | **string** | The user encrypted email address. | 
 **culture** | **string** | Culture | 

### Return type

[**EmployeeFullWrapper**](EmployeeFullWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfileByUserId

> EmployeeFullWrapper GetProfileByUserId(ctx, userid).Execute()

Get a profile by user ID



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-profile-by-user-id/).

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
	resp, r, err := apiClient.PeopleProfilesAPI.GetProfileByUserId(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleProfilesAPI.GetProfileByUserId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfileByUserId`: EmployeeFullWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleProfilesAPI.GetProfileByUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileByUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmployeeFullWrapper**](EmployeeFullWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelfProfile

> EmployeeFullWrapper GetSelfProfile(ctx).Execute()

Get my profile



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-self-profile/).

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
	resp, r, err := apiClient.PeopleProfilesAPI.GetSelfProfile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleProfilesAPI.GetSelfProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSelfProfile`: EmployeeFullWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleProfilesAPI.GetSelfProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfProfileRequest struct via the builder pattern


### Return type

[**EmployeeFullWrapper**](EmployeeFullWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteUsers

> EmployeeArrayWrapper InviteUsers(ctx).InviteUsersRequestDto(inviteUsersRequestDto).Execute()

Invite users



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/invite-users/).

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
	inviteUsersRequestDto := *openapiclient.NewInviteUsersRequestDto([]openapiclient.UserInvitationRequestDto{*openapiclient.NewUserInvitationRequestDto()}) // InviteUsersRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleProfilesAPI.InviteUsers(context.Background()).InviteUsersRequestDto(inviteUsersRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleProfilesAPI.InviteUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InviteUsers`: EmployeeArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleProfilesAPI.InviteUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInviteUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inviteUsersRequestDto** | [**InviteUsersRequestDto**](InviteUsersRequestDto.md) |  | 

### Return type

[**EmployeeArrayWrapper**](EmployeeArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUsers

> EmployeeFullArrayWrapper RemoveUsers(ctx).UpdateMembersRequestDto(updateMembersRequestDto).Execute()

Delete users



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/remove-users/).

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
	resp, r, err := apiClient.PeopleProfilesAPI.RemoveUsers(context.Background()).UpdateMembersRequestDto(updateMembersRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleProfilesAPI.RemoveUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveUsers`: EmployeeFullArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleProfilesAPI.RemoveUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateMembersRequestDto** | [**UpdateMembersRequestDto**](UpdateMembersRequestDto.md) |  | 

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


## ResendUserInvites

> EmployeeFullArrayWrapper ResendUserInvites(ctx).UpdateMembersRequestDto(updateMembersRequestDto).Execute()

Resend activation emails



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/resend-user-invites/).

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
	resp, r, err := apiClient.PeopleProfilesAPI.ResendUserInvites(context.Background()).UpdateMembersRequestDto(updateMembersRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleProfilesAPI.ResendUserInvites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResendUserInvites`: EmployeeFullArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleProfilesAPI.ResendUserInvites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResendUserInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateMembersRequestDto** | [**UpdateMembersRequestDto**](UpdateMembersRequestDto.md) |  | 

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


## UpdateMember

> EmployeeFullWrapper UpdateMember(ctx, userid).UpdateMemberRequestDto(updateMemberRequestDto).Execute()

Update a user



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-member/).

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
	updateMemberRequestDto := *openapiclient.NewUpdateMemberRequestDto() // UpdateMemberRequestDto | The request parameters for updating the user information.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleProfilesAPI.UpdateMember(context.Background(), userid).UpdateMemberRequestDto(updateMemberRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleProfilesAPI.UpdateMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMember`: EmployeeFullWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleProfilesAPI.UpdateMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMemberRequestDto** | [**UpdateMemberRequestDto**](UpdateMemberRequestDto.md) | The request parameters for updating the user information. | 

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


## UpdateMemberCulture

> EmployeeFullWrapper UpdateMemberCulture(ctx, userid).Culture(culture).Execute()

Update a user culture



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-member-culture/).

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
	culture := *openapiclient.NewCulture("en-US") // Culture | The culture name parameters. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeopleProfilesAPI.UpdateMemberCulture(context.Background(), userid).Culture(culture).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeopleProfilesAPI.UpdateMemberCulture``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMemberCulture`: EmployeeFullWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeopleProfilesAPI.UpdateMemberCulture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemberCultureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **culture** | [**Culture**](Culture.md) | The culture name parameters. | 

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

