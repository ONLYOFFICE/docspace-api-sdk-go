# \AIAgentsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiAgentsCreate**](AIAgentsAPI.md#AiAgentsCreate) | **Post** /api/2.0/ai/agents | Create an agent
[**AiAgentsDelete**](AIAgentsAPI.md#AiAgentsDelete) | **Delete** /api/2.0/ai/agents/{id} | Delete an agent
[**AiAgentsGet**](AIAgentsAPI.md#AiAgentsGet) | **Get** /api/2.0/ai/agents/{id} | Get an agent
[**AiAgentsList**](AIAgentsAPI.md#AiAgentsList) | **Get** /api/2.0/ai/agents | List agents
[**AiAgentsNews**](AIAgentsAPI.md#AiAgentsNews) | **Get** /api/2.0/ai/agents/news | List agent news items
[**AiAgentsResetQuota**](AIAgentsAPI.md#AiAgentsResetQuota) | **Put** /api/2.0/ai/agents/resetquota | Reset agents' quota
[**AiAgentsUpdate**](AIAgentsAPI.md#AiAgentsUpdate) | **Put** /api/2.0/ai/agents/{id} | Update an agent
[**AiAgentsUpdateQuota**](AIAgentsAPI.md#AiAgentsUpdateQuota) | **Put** /api/2.0/ai/agents/agentquota | Update agents' quota



## AiAgentsCreate

> AiFolderWrapper AiAgentsCreate(ctx).AiAgentsCreateRequest(aiAgentsCreateRequest).Execute()

Create an agent



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-agents-create/).

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
	aiAgentsCreateRequest := *openapiclient.NewAiAgentsCreateRequest("ProfileId_example", "Prompt_example") // AiAgentsCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.AiAgentsCreate(context.Background()).AiAgentsCreateRequest(aiAgentsCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.AiAgentsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAgentsCreate`: AiFolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.AiAgentsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAgentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiAgentsCreateRequest** | [**AiAgentsCreateRequest**](AiAgentsCreateRequest.md) |  | 

### Return type

[**AiFolderWrapper**](AiFolderWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAgentsDelete

> AiFileOperationWrapper AiAgentsDelete(ctx, id).AiAgentsDeleteRequest(aiAgentsDeleteRequest).Execute()

Delete an agent



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-agents-delete/).

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
	id := "1234" // string | The agent identifier.
	aiAgentsDeleteRequest := *openapiclient.NewAiAgentsDeleteRequest() // AiAgentsDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.AiAgentsDelete(context.Background(), id).AiAgentsDeleteRequest(aiAgentsDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.AiAgentsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAgentsDelete`: AiFileOperationWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.AiAgentsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The agent identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiAgentsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aiAgentsDeleteRequest** | [**AiAgentsDeleteRequest**](AiAgentsDeleteRequest.md) |  | 

### Return type

[**AiFileOperationWrapper**](AiFileOperationWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAgentsGet

> AiAgentsGet200Response AiAgentsGet(ctx, id).Execute()

Get an agent



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-agents-get/).

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
	id := "1234" // string | The agent identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.AiAgentsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.AiAgentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAgentsGet`: AiAgentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.AiAgentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The agent identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiAgentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AiAgentsGet200Response**](AiAgentsGet200Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAgentsList

> AiFolderContentWrapper AiAgentsList(ctx).SubjectId(subjectId).SubjectOwnerId(subjectOwnerId).ExcludeSubject(excludeSubject).Tags(tags).WithoutTags(withoutTags).QuotaFilter(quotaFilter).FilterValue(filterValue).SortBy(sortBy).SortOrder(sortOrder).StartIndex(startIndex).Count(count).Execute()

List agents



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-agents-list/).

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
	subjectId := "00000000-0000-0000-0000-000000000000" // string | Show only the agent rooms this user takes part in. (optional)
	subjectOwnerId := "00000000-0000-0000-0000-000000000000" // string | Show only the agent rooms owned by this user. (optional)
	excludeSubject := false // bool | Invert the user filter: leave out what `subjectId` selects instead of keeping it. (optional)
	tags := "ai,assistant" // string | Show only the agent rooms carrying these tags, comma-separated. (optional)
	withoutTags := false // bool | Show only the agent rooms that carry no tags at all. (optional)
	quotaFilter := int32(0) // int32 | Filter by quota kind: 0 for all, 1 for the default quota, 2 for a custom one. (optional)
	filterValue := "assistant" // string | Show only the agent rooms whose title matches this text. (optional)
	sortBy := "DateAndTime" // string | Field to sort by, for example `DateAndTime`. (optional)
	sortOrder := "descending" // string | Sort direction, `ascending` or `descending`. (optional)
	startIndex := int32(0) // int32 | Index of the first entry to return; 0 starts at the beginning. (optional)
	count := int32(25) // int32 | How many entries to return. The internal service applies its own default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.AiAgentsList(context.Background()).SubjectId(subjectId).SubjectOwnerId(subjectOwnerId).ExcludeSubject(excludeSubject).Tags(tags).WithoutTags(withoutTags).QuotaFilter(quotaFilter).FilterValue(filterValue).SortBy(sortBy).SortOrder(sortOrder).StartIndex(startIndex).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.AiAgentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAgentsList`: AiFolderContentWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.AiAgentsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAgentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subjectId** | **string** | Show only the agent rooms this user takes part in. | 
 **subjectOwnerId** | **string** | Show only the agent rooms owned by this user. | 
 **excludeSubject** | **bool** | Invert the user filter: leave out what `subjectId` selects instead of keeping it. | 
 **tags** | **string** | Show only the agent rooms carrying these tags, comma-separated. | 
 **withoutTags** | **bool** | Show only the agent rooms that carry no tags at all. | 
 **quotaFilter** | **int32** | Filter by quota kind: 0 for all, 1 for the default quota, 2 for a custom one. | 
 **filterValue** | **string** | Show only the agent rooms whose title matches this text. | 
 **sortBy** | **string** | Field to sort by, for example `DateAndTime`. | 
 **sortOrder** | **string** | Sort direction, `ascending` or `descending`. | 
 **startIndex** | **int32** | Index of the first entry to return; 0 starts at the beginning. | 
 **count** | **int32** | How many entries to return. The internal service applies its own default. | 

### Return type

[**AiFolderContentWrapper**](AiFolderContentWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAgentsNews

> AiNewItemsAgentNewItemsArrayWrapper AiAgentsNews(ctx).Execute()

List agent news items



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-agents-news/).

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
	resp, r, err := apiClient.AIAgentsAPI.AiAgentsNews(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.AiAgentsNews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAgentsNews`: AiNewItemsAgentNewItemsArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.AiAgentsNews`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiAgentsNewsRequest struct via the builder pattern


### Return type

[**AiNewItemsAgentNewItemsArrayWrapper**](AiNewItemsAgentNewItemsArrayWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAgentsResetQuota

> AiFolderArrayWrapper AiAgentsResetQuota(ctx).AiAgentsResetQuotaRequest(aiAgentsResetQuotaRequest).Execute()

Reset agents' quota



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-agents-reset-quota/).

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
	aiAgentsResetQuotaRequest := *openapiclient.NewAiAgentsResetQuotaRequest([]openapiclient.AiAgentsUpdateQuotaRequestRoomIdsInner{*openapiclient.NewAiAgentsUpdateQuotaRequestRoomIdsInner()}) // AiAgentsResetQuotaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.AiAgentsResetQuota(context.Background()).AiAgentsResetQuotaRequest(aiAgentsResetQuotaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.AiAgentsResetQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAgentsResetQuota`: AiFolderArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.AiAgentsResetQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAgentsResetQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiAgentsResetQuotaRequest** | [**AiAgentsResetQuotaRequest**](AiAgentsResetQuotaRequest.md) |  | 

### Return type

[**AiFolderArrayWrapper**](AiFolderArrayWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAgentsUpdate

> AiFolderWrapper AiAgentsUpdate(ctx, id).AiAgentsUpdateRequest(aiAgentsUpdateRequest).Execute()

Update an agent



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-agents-update/).

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
	id := "1234" // string | The agent identifier.
	aiAgentsUpdateRequest := *openapiclient.NewAiAgentsUpdateRequest() // AiAgentsUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.AiAgentsUpdate(context.Background(), id).AiAgentsUpdateRequest(aiAgentsUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.AiAgentsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAgentsUpdate`: AiFolderWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.AiAgentsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The agent identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiAgentsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aiAgentsUpdateRequest** | [**AiAgentsUpdateRequest**](AiAgentsUpdateRequest.md) |  | 

### Return type

[**AiFolderWrapper**](AiFolderWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAgentsUpdateQuota

> AiFolderArrayWrapper AiAgentsUpdateQuota(ctx).AiAgentsUpdateQuotaRequest(aiAgentsUpdateQuotaRequest).Execute()

Update agents' quota



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-agents-update-quota/).

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
	aiAgentsUpdateQuotaRequest := *openapiclient.NewAiAgentsUpdateQuotaRequest([]openapiclient.AiAgentsUpdateQuotaRequestRoomIdsInner{*openapiclient.NewAiAgentsUpdateQuotaRequestRoomIdsInner()}, float32(123)) // AiAgentsUpdateQuotaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.AiAgentsUpdateQuota(context.Background()).AiAgentsUpdateQuotaRequest(aiAgentsUpdateQuotaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.AiAgentsUpdateQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAgentsUpdateQuota`: AiFolderArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.AiAgentsUpdateQuota`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAgentsUpdateQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiAgentsUpdateQuotaRequest** | [**AiAgentsUpdateQuotaRequest**](AiAgentsUpdateQuotaRequest.md) |  | 

### Return type

[**AiFolderArrayWrapper**](AiFolderArrayWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

