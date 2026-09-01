# \AIAssignmentsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiAssignmentsAssign**](AIAssignmentsAPI.md#AiAssignmentsAssign) | **Put** /api/2.0/ai/assignments/assign | Assign
[**AiAssignmentsBulkAssign**](AIAssignmentsAPI.md#AiAssignmentsBulkAssign) | **Put** /api/2.0/ai/assignments/bulk-assign | Bulk assign
[**AiAssignmentsCascadeProfileDelete**](AIAssignmentsAPI.md#AiAssignmentsCascadeProfileDelete) | **Delete** /api/2.0/ai/assignments/cascade-profile-delete | Cascade profile delete
[**AiAssignmentsGetAllAssignments**](AIAssignmentsAPI.md#AiAssignmentsGetAllAssignments) | **Get** /api/2.0/ai/assignments/get-all-assignments | Get all assignments
[**AiAssignmentsGetAssignment**](AIAssignmentsAPI.md#AiAssignmentsGetAssignment) | **Get** /api/2.0/ai/assignments/get-assignment | Get assignment
[**AiAssignmentsResolveForAction**](AIAssignmentsAPI.md#AiAssignmentsResolveForAction) | **Get** /api/2.0/ai/assignments/resolve-for-action | Resolve for action
[**AiAssignmentsTryResolveForAction**](AIAssignmentsAPI.md#AiAssignmentsTryResolveForAction) | **Get** /api/2.0/ai/assignments/try-resolve-for-action | Try resolve for action
[**AiAssignmentsUnassign**](AIAssignmentsAPI.md#AiAssignmentsUnassign) | **Delete** /api/2.0/ai/assignments/unassign | Unassign



## AiAssignmentsAssign

> AiAssignmentMutationResult AiAssignmentsAssign(ctx).AiAssignmentsAssignRequest(aiAssignmentsAssignRequest).Execute()

Assign



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-assignments-assign/).

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
	aiAssignmentsAssignRequest := *openapiclient.NewAiAssignmentsAssignRequest(openapiclient.AiActionType("Default"), "ProfileId_example") // AiAssignmentsAssignRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAssignmentsAPI.AiAssignmentsAssign(context.Background()).AiAssignmentsAssignRequest(aiAssignmentsAssignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAssignmentsAPI.AiAssignmentsAssign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAssignmentsAssign`: AiAssignmentMutationResult
	fmt.Fprintf(os.Stdout, "Response from `AIAssignmentsAPI.AiAssignmentsAssign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAssignmentsAssignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiAssignmentsAssignRequest** | [**AiAssignmentsAssignRequest**](AiAssignmentsAssignRequest.md) |  | 

### Return type

[**AiAssignmentMutationResult**](AiAssignmentMutationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAssignmentsBulkAssign

> AiBulkAssignmentResult AiAssignmentsBulkAssign(ctx).RequestBody(requestBody).Execute()

Bulk assign



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-assignments-bulk-assign/).

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
	requestBody := map[string]string{"key": "Inner_example"} // map[string]string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAssignmentsAPI.AiAssignmentsBulkAssign(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAssignmentsAPI.AiAssignmentsBulkAssign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAssignmentsBulkAssign`: AiBulkAssignmentResult
	fmt.Fprintf(os.Stdout, "Response from `AIAssignmentsAPI.AiAssignmentsBulkAssign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAssignmentsBulkAssignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]string** |  | 

### Return type

[**AiBulkAssignmentResult**](AiBulkAssignmentResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAssignmentsCascadeProfileDelete

> AiSuccessResponse AiAssignmentsCascadeProfileDelete(ctx).Body(body).Execute()

Cascade profile delete



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-assignments-cascade-profile-delete/).

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
	body := "body_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAssignmentsAPI.AiAssignmentsCascadeProfileDelete(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAssignmentsAPI.AiAssignmentsCascadeProfileDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAssignmentsCascadeProfileDelete`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIAssignmentsAPI.AiAssignmentsCascadeProfileDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAssignmentsCascadeProfileDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

[**AiSuccessResponse**](AiSuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAssignmentsGetAllAssignments

> map[string]string AiAssignmentsGetAllAssignments(ctx).EntityId(entityId).Execute()

Get all assignments



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-assignments-get-all-assignments/).

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
	entityId := "entityId_example" // string | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAssignmentsAPI.AiAssignmentsGetAllAssignments(context.Background()).EntityId(entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAssignmentsAPI.AiAssignmentsGetAllAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAssignmentsGetAllAssignments`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `AIAssignmentsAPI.AiAssignmentsGetAllAssignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAssignmentsGetAllAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string** | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAssignmentsGetAssignment

> string AiAssignmentsGetAssignment(ctx).ActionType(actionType).Execute()

Get assignment



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-assignments-get-assignment/).

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
	actionType := "actionType_example" // string | The AI action the request applies to - one of Default, Chat, Code, Summarization, Translation, TextAnalyze, ImageGeneration, OCR, Vision.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAssignmentsAPI.AiAssignmentsGetAssignment(context.Background()).ActionType(actionType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAssignmentsAPI.AiAssignmentsGetAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAssignmentsGetAssignment`: string
	fmt.Fprintf(os.Stdout, "Response from `AIAssignmentsAPI.AiAssignmentsGetAssignment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAssignmentsGetAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actionType** | **string** | The AI action the request applies to - one of Default, Chat, Code, Summarization, Translation, TextAnalyze, ImageGeneration, OCR, Vision. | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAssignmentsResolveForAction

> AiResolvedAssignment AiAssignmentsResolveForAction(ctx).ActionType(actionType).EntityId(entityId).Execute()

Resolve for action



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-assignments-resolve-for-action/).

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
	actionType := "actionType_example" // string | The AI action the request applies to - one of Default, Chat, Code, Summarization, Translation, TextAnalyze, ImageGeneration, OCR, Vision.
	entityId := "entityId_example" // string | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAssignmentsAPI.AiAssignmentsResolveForAction(context.Background()).ActionType(actionType).EntityId(entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAssignmentsAPI.AiAssignmentsResolveForAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAssignmentsResolveForAction`: AiResolvedAssignment
	fmt.Fprintf(os.Stdout, "Response from `AIAssignmentsAPI.AiAssignmentsResolveForAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAssignmentsResolveForActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actionType** | **string** | The AI action the request applies to - one of Default, Chat, Code, Summarization, Translation, TextAnalyze, ImageGeneration, OCR, Vision. | 
 **entityId** | **string** | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. | 

### Return type

[**AiResolvedAssignment**](AiResolvedAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAssignmentsTryResolveForAction

> AiResolvedAssignment AiAssignmentsTryResolveForAction(ctx).ActionType(actionType).EntityId(entityId).Execute()

Try resolve for action



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-assignments-try-resolve-for-action/).

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
	actionType := "actionType_example" // string | The AI action the request applies to - one of Default, Chat, Code, Summarization, Translation, TextAnalyze, ImageGeneration, OCR, Vision.
	entityId := "entityId_example" // string | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAssignmentsAPI.AiAssignmentsTryResolveForAction(context.Background()).ActionType(actionType).EntityId(entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAssignmentsAPI.AiAssignmentsTryResolveForAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAssignmentsTryResolveForAction`: AiResolvedAssignment
	fmt.Fprintf(os.Stdout, "Response from `AIAssignmentsAPI.AiAssignmentsTryResolveForAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAssignmentsTryResolveForActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actionType** | **string** | The AI action the request applies to - one of Default, Chat, Code, Summarization, Translation, TextAnalyze, ImageGeneration, OCR, Vision. | 
 **entityId** | **string** | The DocSpace entity the request is scoped to - the room, folder or agent workspace the chat is invoked from. Omit for the portal-wide scope. | 

### Return type

[**AiResolvedAssignment**](AiResolvedAssignment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiAssignmentsUnassign

> AiSuccessResponse AiAssignmentsUnassign(ctx).Body(body).Execute()

Unassign



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-assignments-unassign/).

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
	body := string(987) // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIAssignmentsAPI.AiAssignmentsUnassign(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIAssignmentsAPI.AiAssignmentsUnassign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiAssignmentsUnassign`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIAssignmentsAPI.AiAssignmentsUnassign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiAssignmentsUnassignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

[**AiSuccessResponse**](AiSuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

