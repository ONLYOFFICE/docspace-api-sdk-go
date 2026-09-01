# \AIProfilesAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiProfilesCreate**](AIProfilesAPI.md#AiProfilesCreate) | **Post** /api/2.0/ai/profiles/create | Create
[**AiProfilesDelete**](AIProfilesAPI.md#AiProfilesDelete) | **Delete** /api/2.0/ai/profiles/delete | Delete
[**AiProfilesGetById**](AIProfilesAPI.md#AiProfilesGetById) | **Get** /api/2.0/ai/profiles/get-by-id | Get by id
[**AiProfilesList**](AIProfilesAPI.md#AiProfilesList) | **Get** /api/2.0/ai/profiles/list | List
[**AiProfilesListModels**](AIProfilesAPI.md#AiProfilesListModels) | **Get** /api/2.0/ai/profiles/list-models | List models
[**AiProfilesListProviderModels**](AIProfilesAPI.md#AiProfilesListProviderModels) | **Post** /api/2.0/ai/profiles/list-provider-models | List provider models
[**AiProfilesTestConnection**](AIProfilesAPI.md#AiProfilesTestConnection) | **Post** /api/2.0/ai/profiles/test-connection | Test connection
[**AiProfilesUpdate**](AIProfilesAPI.md#AiProfilesUpdate) | **Put** /api/2.0/ai/profiles/update | Update



## AiProfilesCreate

> AiProfileMutationResult AiProfilesCreate(ctx).AiCreateProfileInput(aiCreateProfileInput).Execute()

Create



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-profiles-create/).

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
	aiCreateProfileInput := *openapiclient.NewAiCreateProfileInput("Name_example", *openapiclient.NewAiProviderType(), "BaseUrl_example", "ModelId_example") // AiCreateProfileInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIProfilesAPI.AiProfilesCreate(context.Background()).AiCreateProfileInput(aiCreateProfileInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIProfilesAPI.AiProfilesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiProfilesCreate`: AiProfileMutationResult
	fmt.Fprintf(os.Stdout, "Response from `AIProfilesAPI.AiProfilesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiProfilesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiCreateProfileInput** | [**AiCreateProfileInput**](AiCreateProfileInput.md) |  | 

### Return type

[**AiProfileMutationResult**](AiProfileMutationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiProfilesDelete

> AiSuccessResponse AiProfilesDelete(ctx).Body(body).Execute()

Delete



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-profiles-delete/).

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
	resp, r, err := apiClient.AIProfilesAPI.AiProfilesDelete(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIProfilesAPI.AiProfilesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiProfilesDelete`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIProfilesAPI.AiProfilesDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiProfilesDeleteRequest struct via the builder pattern


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


## AiProfilesGetById

> AiProfilesGetById200Response AiProfilesGetById(ctx).Id(id).Execute()

Get by id



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-profiles-get-by-id/).

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
	id := "id_example" // string | The AI provider profile identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIProfilesAPI.AiProfilesGetById(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIProfilesAPI.AiProfilesGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiProfilesGetById`: AiProfilesGetById200Response
	fmt.Fprintf(os.Stdout, "Response from `AIProfilesAPI.AiProfilesGetById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiProfilesGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The AI provider profile identifier. | 

### Return type

[**AiProfilesGetById200Response**](AiProfilesGetById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiProfilesList

> []AiProfile AiProfilesList(ctx).Execute()

List



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-profiles-list/).

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
	resp, r, err := apiClient.AIProfilesAPI.AiProfilesList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIProfilesAPI.AiProfilesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiProfilesList`: []AiProfile
	fmt.Fprintf(os.Stdout, "Response from `AIProfilesAPI.AiProfilesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiProfilesListRequest struct via the builder pattern


### Return type

[**[]AiProfile**](AiProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiProfilesListModels

> []AiModel AiProfilesListModels(ctx).ProfileId(profileId).Execute()

List models



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-profiles-list-models/).

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
	profileId := "profileId_example" // string | The AI provider profile identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIProfilesAPI.AiProfilesListModels(context.Background()).ProfileId(profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIProfilesAPI.AiProfilesListModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiProfilesListModels`: []AiModel
	fmt.Fprintf(os.Stdout, "Response from `AIProfilesAPI.AiProfilesListModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiProfilesListModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileId** | **string** | The AI provider profile identifier. | 

### Return type

[**[]AiModel**](AiModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiProfilesListProviderModels

> []AiModel AiProfilesListProviderModels(ctx).AiProfilesListProviderModelsRequest(aiProfilesListProviderModelsRequest).Execute()

List provider models



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-profiles-list-provider-models/).

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
	aiProfilesListProviderModelsRequest := *openapiclient.NewAiProfilesListProviderModelsRequest(*openapiclient.NewAiProviderType(), "BaseUrl_example", "ApiKey_example") // AiProfilesListProviderModelsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIProfilesAPI.AiProfilesListProviderModels(context.Background()).AiProfilesListProviderModelsRequest(aiProfilesListProviderModelsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIProfilesAPI.AiProfilesListProviderModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiProfilesListProviderModels`: []AiModel
	fmt.Fprintf(os.Stdout, "Response from `AIProfilesAPI.AiProfilesListProviderModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiProfilesListProviderModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiProfilesListProviderModelsRequest** | [**AiProfilesListProviderModelsRequest**](AiProfilesListProviderModelsRequest.md) |  | 

### Return type

[**[]AiModel**](AiModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiProfilesTestConnection

> AiProfilesTestConnection200Response AiProfilesTestConnection(ctx).Body(body).Execute()

Test connection



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-profiles-test-connection/).

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
	resp, r, err := apiClient.AIProfilesAPI.AiProfilesTestConnection(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIProfilesAPI.AiProfilesTestConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiProfilesTestConnection`: AiProfilesTestConnection200Response
	fmt.Fprintf(os.Stdout, "Response from `AIProfilesAPI.AiProfilesTestConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiProfilesTestConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

[**AiProfilesTestConnection200Response**](AiProfilesTestConnection200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiProfilesUpdate

> AiProfileMutationResult AiProfilesUpdate(ctx).AiProfile(aiProfile).Execute()

Update



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-profiles-update/).

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
	aiProfile := *openapiclient.NewAiProfile("Id_example", "Name_example", *openapiclient.NewAiProviderType(), "BaseUrl_example", "ModelId_example") // AiProfile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIProfilesAPI.AiProfilesUpdate(context.Background()).AiProfile(aiProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIProfilesAPI.AiProfilesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiProfilesUpdate`: AiProfileMutationResult
	fmt.Fprintf(os.Stdout, "Response from `AIProfilesAPI.AiProfilesUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiProfilesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiProfile** | [**AiProfile**](AiProfile.md) |  | 

### Return type

[**AiProfileMutationResult**](AiProfileMutationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

