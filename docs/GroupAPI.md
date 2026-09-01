# \GroupAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroup**](GroupAPI.md#AddGroup) | **Post** /api/2.0/group | Add a new group
[**AddMembersTo**](GroupAPI.md#AddMembersTo) | **Put** /api/2.0/group/{id}/members | Add group members
[**DeleteGroup**](GroupAPI.md#DeleteGroup) | **Delete** /api/2.0/group/{id} | Delete a group
[**GetGroup**](GroupAPI.md#GetGroup) | **Get** /api/2.0/group/{id} | Get a group
[**GetGroupByUserId**](GroupAPI.md#GetGroupByUserId) | **Get** /api/2.0/group/user/{userid} | Get user groups
[**GetGroups**](GroupAPI.md#GetGroups) | **Get** /api/2.0/group | Get groups
[**MoveMembersTo**](GroupAPI.md#MoveMembersTo) | **Put** /api/2.0/group/{fromId}/members/{toId} | Move group members
[**RemoveMembersFrom**](GroupAPI.md#RemoveMembersFrom) | **Delete** /api/2.0/group/{id}/members | Remove group members
[**SetGroupManager**](GroupAPI.md#SetGroupManager) | **Put** /api/2.0/group/{id}/manager | Set a group manager
[**SetMembersTo**](GroupAPI.md#SetMembersTo) | **Post** /api/2.0/group/{id}/members | Replace group members
[**UpdateGroup**](GroupAPI.md#UpdateGroup) | **Put** /api/2.0/group/{id} | Update a group



## AddGroup

> GroupWrapper AddGroup(ctx).GroupRequestDto(groupRequestDto).Execute()

Add a new group



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/add-group/).

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
	groupRequestDto := *openapiclient.NewGroupRequestDto("Marketing Team") // GroupRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.AddGroup(context.Background()).GroupRequestDto(groupRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.AddGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGroup`: GroupWrapper
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.AddGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupRequestDto** | [**GroupRequestDto**](GroupRequestDto.md) |  | 

### Return type

[**GroupWrapper**](GroupWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddMembersTo

> GroupWrapper AddMembersTo(ctx, id).MembersRequest(membersRequest).Execute()

Add group members



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/add-members-to/).

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
	id := "00000000-0000-0000-0000-000000000000" // string | The group ID.
	membersRequest := *openapiclient.NewMembersRequest() // MembersRequest | The member request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.AddMembersTo(context.Background(), id).MembersRequest(membersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.AddMembersTo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddMembersTo`: GroupWrapper
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.AddMembersTo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMembersToRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **membersRequest** | [**MembersRequest**](MembersRequest.md) | The member request. | 

### Return type

[**GroupWrapper**](GroupWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> DeleteGroup(ctx, id).Execute()

Delete a group



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-group/).

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
	id := "00000000-0000-0000-0000-000000000000" // string | The group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupAPI.DeleteGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.DeleteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


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


## GetGroup

> GroupWrapper GetGroup(ctx, id).IncludeMembers(includeMembers).Execute()

Get a group



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-group/).

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
	id := "00000000-0000-0000-0000-000000000000" // string | The group ID.
	includeMembers := true // bool | Specifies whether to include the group members or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.GetGroup(context.Background(), id).IncludeMembers(includeMembers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.GetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroup`: GroupWrapper
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeMembers** | **bool** | Specifies whether to include the group members or not. | 

### Return type

[**GroupWrapper**](GroupWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupByUserId

> GroupSummaryArrayWrapper GetGroupByUserId(ctx, userid).Execute()

Get user groups



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-group-by-user-id/).

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
	resp, r, err := apiClient.GroupAPI.GetGroupByUserId(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.GetGroupByUserId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupByUserId`: GroupSummaryArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.GetGroupByUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupByUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupSummaryArrayWrapper**](GroupSummaryArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroups

> GroupArrayWrapper GetGroups(ctx).UserId(userId).Manager(manager).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()

Get groups



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-groups/).

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
	manager := false // bool | Specifies if the user is a manager or not. (optional)
	count := int32(25) // int32 | The number of records to retrieve. (optional)
	startIndex := int32(0) // int32 | The starting index for paginated results. (optional)
	sortBy := "displayName" // string | Specifies the property used to sort the query results. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The order in which the results are sorted. (optional)
	filterValue := "John" // string | The text used for filtering or searching group data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.GetGroups(context.Background()).UserId(userId).Manager(manager).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.GetGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroups`: GroupArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.GetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | The user ID. | 
 **manager** | **bool** | Specifies if the user is a manager or not. | 
 **count** | **int32** | The number of records to retrieve. | 
 **startIndex** | **int32** | The starting index for paginated results. | 
 **sortBy** | **string** | Specifies the property used to sort the query results. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The order in which the results are sorted. | 
 **filterValue** | **string** | The text used for filtering or searching group data. | 

### Return type

[**GroupArrayWrapper**](GroupArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveMembersTo

> GroupWrapper MoveMembersTo(ctx, fromId, toId).Execute()

Move group members



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/move-members-to/).

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
	fromId := "00000000-0000-0000-0000-000000000000" // string | The group ID to move from.
	toId := "11111111-1111-1111-1111-111111111111" // string | The group ID to move to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.MoveMembersTo(context.Background(), fromId, toId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.MoveMembersTo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveMembersTo`: GroupWrapper
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.MoveMembersTo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromId** | **string** | The group ID to move from. | 
**toId** | **string** | The group ID to move to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveMembersToRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupWrapper**](GroupWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMembersFrom

> GroupWrapper RemoveMembersFrom(ctx, id).MembersRequest(membersRequest).Execute()

Remove group members



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/remove-members-from/).

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
	id := "00000000-0000-0000-0000-000000000000" // string | The group ID.
	membersRequest := *openapiclient.NewMembersRequest() // MembersRequest | The member request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.RemoveMembersFrom(context.Background(), id).MembersRequest(membersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.RemoveMembersFrom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveMembersFrom`: GroupWrapper
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.RemoveMembersFrom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMembersFromRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **membersRequest** | [**MembersRequest**](MembersRequest.md) | The member request. | 

### Return type

[**GroupWrapper**](GroupWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGroupManager

> GroupWrapper SetGroupManager(ctx, id).SetManagerRequest(setManagerRequest).Execute()

Set a group manager



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-group-manager/).

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
	id := "00000000-0000-0000-0000-000000000000" // string | The group ID.
	setManagerRequest := *openapiclient.NewSetManagerRequest("00000000-0000-0000-0000-000000000000") // SetManagerRequest | The request for setting a group manager.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.SetGroupManager(context.Background(), id).SetManagerRequest(setManagerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.SetGroupManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGroupManager`: GroupWrapper
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.SetGroupManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setManagerRequest** | [**SetManagerRequest**](SetManagerRequest.md) | The request for setting a group manager. | 

### Return type

[**GroupWrapper**](GroupWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMembersTo

> GroupWrapper SetMembersTo(ctx, id).MembersRequest(membersRequest).Execute()

Replace group members



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-members-to/).

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
	id := "00000000-0000-0000-0000-000000000000" // string | The group ID.
	membersRequest := *openapiclient.NewMembersRequest() // MembersRequest | The member request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.SetMembersTo(context.Background(), id).MembersRequest(membersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.SetMembersTo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetMembersTo`: GroupWrapper
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.SetMembersTo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMembersToRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **membersRequest** | [**MembersRequest**](MembersRequest.md) | The member request. | 

### Return type

[**GroupWrapper**](GroupWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> GroupWrapper UpdateGroup(ctx, id).UpdateGroupRequest(updateGroupRequest).Execute()

Update a group



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-group/).

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
	id := "00000000-0000-0000-0000-000000000000" // string | The group ID.
	updateGroupRequest := *openapiclient.NewUpdateGroupRequest() // UpdateGroupRequest | The request for updating a group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAPI.UpdateGroup(context.Background(), id).UpdateGroupRequest(updateGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAPI.UpdateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroup`: GroupWrapper
	fmt.Fprintf(os.Stdout, "Response from `GroupAPI.UpdateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGroupRequest** | [**UpdateGroupRequest**](UpdateGroupRequest.md) | The request for updating a group. | 

### Return type

[**GroupWrapper**](GroupWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

