# \AIAgentsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAgent**](AIAgentsAPI.md#CreateAgent) | **Post** /api/2.0/ai/agents | Create an ai agent
[**DeleteAgent**](AIAgentsAPI.md#DeleteAgent) | **Delete** /api/2.0/ai/agents/{id} | Remove an ai agent
[**GetAgentInfo**](AIAgentsAPI.md#GetAgentInfo) | **Get** /api/2.0/ai/agents/{id} | Return an ai agent
[**GetAgents**](AIAgentsAPI.md#GetAgents) | **Get** /api/2.0/ai/agents | Get ai agents
[**GetAgentsNewItems**](AIAgentsAPI.md#GetAgentsNewItems) | **Get** /api/2.0/ai/agents/news | Get the room new items
[**ResetAgentsQuota**](AIAgentsAPI.md#ResetAgentsQuota) | **Put** /api/2.0/ai/agents/resetquota | Reset the AI agents quota limit
[**UpdateAgent**](AIAgentsAPI.md#UpdateAgent) | **Put** /api/2.0/ai/agents/{id} | Update an ai agent
[**UpdateAgentsQuota**](AIAgentsAPI.md#UpdateAgentsQuota) | **Put** /api/2.0/ai/agents/agentquota | Change the AI agent quota limit



## CreateAgent

> FolderIntegerWrapper CreateAgent(ctx).CreateAgentRequestDto(createAgentRequestDto).Execute()

Create an ai agent



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-agent/).

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
	createAgentRequestDto := *openapiclient.NewCreateAgentRequestDto("My AI Agent Room", *openapiclient.NewChatSettings()) // CreateAgentRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.CreateAgent(context.Background()).CreateAgentRequestDto(createAgentRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.CreateAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAgent`: FolderIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.CreateAgent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAgentRequestDto** | [**CreateAgentRequestDto**](CreateAgentRequestDto.md) |  | 

### Return type

[**FolderIntegerWrapper**](FolderIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgent

> FileOperationWrapper DeleteAgent(ctx, id).DeleteRoomRequest(deleteRoomRequest).Execute()

Remove an ai agent



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-agent/).

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
	id := int32(10) // int32 | The room ID.
	deleteRoomRequest := *openapiclient.NewDeleteRoomRequest() // DeleteRoomRequest | The parameters for deleting a room.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.DeleteAgent(context.Background(), id).DeleteRoomRequest(deleteRoomRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.DeleteAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAgent`: FileOperationWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.DeleteAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteRoomRequest** | [**DeleteRoomRequest**](DeleteRoomRequest.md) | The parameters for deleting a room. | 

### Return type

[**FileOperationWrapper**](FileOperationWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentInfo

> FolderIntegerWrapper GetAgentInfo(ctx, id).Execute()

Return an ai agent



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-agent-info/).

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
	id := int32(1) // int32 | The room ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.GetAgentInfo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.GetAgentInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentInfo`: FolderIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.GetAgentInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderIntegerWrapper**](FolderIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgents

> FolderContentIntegerWrapper GetAgents(ctx).SubjectId(subjectId).SubjectOwnerId(subjectOwnerId).WithoutTags(withoutTags).Tags(tags).ExcludeSubject(excludeSubject).SubjectFilter(subjectFilter).QuotaFilter(quotaFilter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()

Get ai agents



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-agents/).

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
	subjectId := "00000000-0000-0000-0000-000000000000" // string | The filter by user ID. (optional)
	subjectOwnerId := "00000000-0000-0000-0000-000000000000" // string | The filter by room owner ID. (optional)
	withoutTags := false // bool | Specifies whether to search by tags or not. (optional)
	tags := "ai,assistant" // string | The tags in the serialized format. (optional)
	excludeSubject := false // bool | Specifies whether to exclude search by user or group ID. (optional)
	subjectFilter := openapiclient.SubjectFilter(0) // SubjectFilter | The filter by user (Owner - 0, Member - 1). (optional)
	quotaFilter := openapiclient.QuotaFilter(0) // QuotaFilter | The filter by quota (All - 0, Default - 1, Custom - 2). (optional)
	count := int32(25) // int32 | Specifies the maximum number of items to retrieve. (optional)
	startIndex := int32(0) // int32 | The index from which to start retrieving the room content. (optional)
	sortBy := "DateAndTime" // string | Specifies the field by which the room content should be sorted. (optional)
	sortOrder := openapiclient.SortOrder(0) // SortOrder | The order in which the results are sorted. (optional)
	filterValue := "my agent" // string | The text filter value used to refine search or query operations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.GetAgents(context.Background()).SubjectId(subjectId).SubjectOwnerId(subjectOwnerId).WithoutTags(withoutTags).Tags(tags).ExcludeSubject(excludeSubject).SubjectFilter(subjectFilter).QuotaFilter(quotaFilter).Count(count).StartIndex(startIndex).SortBy(sortBy).SortOrder(sortOrder).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.GetAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgents`: FolderContentIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.GetAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subjectId** | **string** | The filter by user ID. | 
 **subjectOwnerId** | **string** | The filter by room owner ID. | 
 **withoutTags** | **bool** | Specifies whether to search by tags or not. | 
 **tags** | **string** | The tags in the serialized format. | 
 **excludeSubject** | **bool** | Specifies whether to exclude search by user or group ID. | 
 **subjectFilter** | [**SubjectFilter**](SubjectFilter.md) | The filter by user (Owner - 0, Member - 1). | 
 **quotaFilter** | [**QuotaFilter**](QuotaFilter.md) | The filter by quota (All - 0, Default - 1, Custom - 2). | 
 **count** | **int32** | Specifies the maximum number of items to retrieve. | 
 **startIndex** | **int32** | The index from which to start retrieving the room content. | 
 **sortBy** | **string** | Specifies the field by which the room content should be sorted. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | The order in which the results are sorted. | 
 **filterValue** | **string** | The text filter value used to refine search or query operations. | 

### Return type

[**FolderContentIntegerWrapper**](FolderContentIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentsNewItems

> NewItemsAgentNewItemsArrayWrapper GetAgentsNewItems(ctx).Execute()

Get the room new items



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-agents-new-items/).

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
	resp, r, err := apiClient.AIAgentsAPI.GetAgentsNewItems(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.GetAgentsNewItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAgentsNewItems`: NewItemsAgentNewItemsArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.GetAgentsNewItems`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentsNewItemsRequest struct via the builder pattern


### Return type

[**NewItemsAgentNewItemsArrayWrapper**](NewItemsAgentNewItemsArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetAgentsQuota

> FolderIntegerArrayWrapper ResetAgentsQuota(ctx).UpdateRoomsRoomIdsRequestDtoInteger(updateRoomsRoomIdsRequestDtoInteger).Execute()

Reset the AI agents quota limit



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/reset-agents-quota/).

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
	updateRoomsRoomIdsRequestDtoInteger := *openapiclient.NewUpdateRoomsRoomIdsRequestDtoInteger() // UpdateRoomsRoomIdsRequestDtoInteger |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.ResetAgentsQuota(context.Background()).UpdateRoomsRoomIdsRequestDtoInteger(updateRoomsRoomIdsRequestDtoInteger).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.ResetAgentsQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetAgentsQuota`: FolderIntegerArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.ResetAgentsQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetAgentsQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRoomsRoomIdsRequestDtoInteger** | [**UpdateRoomsRoomIdsRequestDtoInteger**](UpdateRoomsRoomIdsRequestDtoInteger.md) |  | 

### Return type

[**FolderIntegerArrayWrapper**](FolderIntegerArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgent

> FolderIntegerWrapper UpdateAgent(ctx, id).UpdateRoomRequest(updateRoomRequest).Execute()

Update an ai agent



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-agent/).

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
	id := int32(56) // int32 | The room ID.
	updateRoomRequest := *openapiclient.NewUpdateRoomRequest() // UpdateRoomRequest | The request parameters for updating a room.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.UpdateAgent(context.Background(), id).UpdateRoomRequest(updateRoomRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.UpdateAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAgent`: FolderIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.UpdateAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The room ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRoomRequest** | [**UpdateRoomRequest**](UpdateRoomRequest.md) | The request parameters for updating a room. | 

### Return type

[**FolderIntegerWrapper**](FolderIntegerWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgentsQuota

> FolderIntegerArrayWrapper UpdateAgentsQuota(ctx).UpdateRoomsQuotaRequestDtoInteger(updateRoomsQuotaRequestDtoInteger).Execute()

Change the AI agent quota limit



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-agents-quota/).

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
	updateRoomsQuotaRequestDtoInteger := *openapiclient.NewUpdateRoomsQuotaRequestDtoInteger() // UpdateRoomsQuotaRequestDtoInteger |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.UpdateAgentsQuota(context.Background()).UpdateRoomsQuotaRequestDtoInteger(updateRoomsQuotaRequestDtoInteger).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.UpdateAgentsQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAgentsQuota`: FolderIntegerArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.UpdateAgentsQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgentsQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRoomsQuotaRequestDtoInteger** | [**UpdateRoomsQuotaRequestDtoInteger**](UpdateRoomsQuotaRequestDtoInteger.md) |  | 

### Return type

[**FolderIntegerArrayWrapper**](FolderIntegerArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

