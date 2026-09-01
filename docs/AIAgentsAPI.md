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

> AiFolderIntegerWrapper AiAgentsCreate(ctx).AiAgentsCreateRequest(aiAgentsCreateRequest).Execute()

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
	// response from `AiAgentsCreate`: AiFolderIntegerWrapper
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

[**AiFolderIntegerWrapper**](AiFolderIntegerWrapper.md)

### Authorization

No authorization required

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
	id := "id_example" // string | The agent identifier.
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAgentsGet

> AiFolderIntegerWrapper AiAgentsGet(ctx, id).Execute()

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
	id := "id_example" // string | The agent identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.AiAgentsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.AiAgentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAgentsGet`: AiFolderIntegerWrapper
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

[**AiFolderIntegerWrapper**](AiFolderIntegerWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAgentsList

> AiFolderContentIntegerWrapper AiAgentsList(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.AiAgentsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.AiAgentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAgentsList`: AiFolderContentIntegerWrapper
	fmt.Fprintf(os.Stdout, "Response from `AIAgentsAPI.AiAgentsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiAgentsListRequest struct via the builder pattern


### Return type

[**AiFolderContentIntegerWrapper**](AiFolderContentIntegerWrapper.md)

### Authorization

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAgentsResetQuota

> AiFolderIntegerArrayWrapper AiAgentsResetQuota(ctx).AiAgentsResetQuotaRequest(aiAgentsResetQuotaRequest).Execute()

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
	// response from `AiAgentsResetQuota`: AiFolderIntegerArrayWrapper
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

[**AiFolderIntegerArrayWrapper**](AiFolderIntegerArrayWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAgentsUpdate

> AiFolderIntegerWrapper AiAgentsUpdate(ctx, id).AiAgentsUpdateRequest(aiAgentsUpdateRequest).Execute()

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
	id := "id_example" // string | The agent identifier.
	aiAgentsUpdateRequest := *openapiclient.NewAiAgentsUpdateRequest() // AiAgentsUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAgentsAPI.AiAgentsUpdate(context.Background(), id).AiAgentsUpdateRequest(aiAgentsUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAgentsAPI.AiAgentsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAgentsUpdate`: AiFolderIntegerWrapper
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

[**AiFolderIntegerWrapper**](AiFolderIntegerWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAgentsUpdateQuota

> AiFolderIntegerArrayWrapper AiAgentsUpdateQuota(ctx).AiAgentsUpdateQuotaRequest(aiAgentsUpdateQuotaRequest).Execute()

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
	// response from `AiAgentsUpdateQuota`: AiFolderIntegerArrayWrapper
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

[**AiFolderIntegerArrayWrapper**](AiFolderIntegerArrayWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

