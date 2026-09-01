# \AIPromptsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiPromptsCreate**](AIPromptsAPI.md#AiPromptsCreate) | **Post** /api/2.0/ai/prompts/create | Create
[**AiPromptsCreateFolder**](AIPromptsAPI.md#AiPromptsCreateFolder) | **Post** /api/2.0/ai/prompts/create-folder | Create folder
[**AiPromptsDelete**](AIPromptsAPI.md#AiPromptsDelete) | **Delete** /api/2.0/ai/prompts/delete | Delete
[**AiPromptsDeleteFolder**](AIPromptsAPI.md#AiPromptsDeleteFolder) | **Delete** /api/2.0/ai/prompts/delete-folder | Delete folder
[**AiPromptsExport**](AIPromptsAPI.md#AiPromptsExport) | **Get** /api/2.0/ai/prompts/export | Export
[**AiPromptsGetById**](AIPromptsAPI.md#AiPromptsGetById) | **Get** /api/2.0/ai/prompts/get-by-id | Get by id
[**AiPromptsGetFolderById**](AIPromptsAPI.md#AiPromptsGetFolderById) | **Get** /api/2.0/ai/prompts/get-folder-by-id | Get folder by id
[**AiPromptsImportBundle**](AIPromptsAPI.md#AiPromptsImportBundle) | **Post** /api/2.0/ai/prompts/import-bundle | Import bundle
[**AiPromptsList**](AIPromptsAPI.md#AiPromptsList) | **Get** /api/2.0/ai/prompts/list | List
[**AiPromptsListFolders**](AIPromptsAPI.md#AiPromptsListFolders) | **Get** /api/2.0/ai/prompts/list-folders | List folders
[**AiPromptsMove**](AIPromptsAPI.md#AiPromptsMove) | **Put** /api/2.0/ai/prompts/move | Move
[**AiPromptsRenameFolder**](AIPromptsAPI.md#AiPromptsRenameFolder) | **Put** /api/2.0/ai/prompts/rename-folder | Rename folder
[**AiPromptsUpdate**](AIPromptsAPI.md#AiPromptsUpdate) | **Put** /api/2.0/ai/prompts/update | Update



## AiPromptsCreate

> AiPromptMutationResult AiPromptsCreate(ctx).AiCreatePromptInput(aiCreatePromptInput).Execute()

Create



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-prompts-create/).

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
	aiCreatePromptInput := *openapiclient.NewAiCreatePromptInput("Name_example", "Text_example") // AiCreatePromptInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIPromptsAPI.AiPromptsCreate(context.Background()).AiCreatePromptInput(aiCreatePromptInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIPromptsAPI.AiPromptsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiPromptsCreate`: AiPromptMutationResult
	fmt.Fprintf(os.Stdout, "Response from `AIPromptsAPI.AiPromptsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiPromptsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiCreatePromptInput** | [**AiCreatePromptInput**](AiCreatePromptInput.md) |  | 

### Return type

[**AiPromptMutationResult**](AiPromptMutationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiPromptsCreateFolder

> AiFolderMutationResult AiPromptsCreateFolder(ctx).Body(body).Execute()

Create folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-prompts-create-folder/).

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
	resp, r, err := apiClient.AIPromptsAPI.AiPromptsCreateFolder(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIPromptsAPI.AiPromptsCreateFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiPromptsCreateFolder`: AiFolderMutationResult
	fmt.Fprintf(os.Stdout, "Response from `AIPromptsAPI.AiPromptsCreateFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiPromptsCreateFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

[**AiFolderMutationResult**](AiFolderMutationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiPromptsDelete

> AiSuccessResponse AiPromptsDelete(ctx).Body(body).Execute()

Delete



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-prompts-delete/).

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
	resp, r, err := apiClient.AIPromptsAPI.AiPromptsDelete(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIPromptsAPI.AiPromptsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiPromptsDelete`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIPromptsAPI.AiPromptsDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiPromptsDeleteRequest struct via the builder pattern


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


## AiPromptsDeleteFolder

> AiSuccessResponse AiPromptsDeleteFolder(ctx).Body(body).Execute()

Delete folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-prompts-delete-folder/).

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
	resp, r, err := apiClient.AIPromptsAPI.AiPromptsDeleteFolder(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIPromptsAPI.AiPromptsDeleteFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiPromptsDeleteFolder`: AiSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `AIPromptsAPI.AiPromptsDeleteFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiPromptsDeleteFolderRequest struct via the builder pattern


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


## AiPromptsExport

> AiPromptBundle AiPromptsExport(ctx).Execute()

Export



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-prompts-export/).

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
	resp, r, err := apiClient.AIPromptsAPI.AiPromptsExport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIPromptsAPI.AiPromptsExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiPromptsExport`: AiPromptBundle
	fmt.Fprintf(os.Stdout, "Response from `AIPromptsAPI.AiPromptsExport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiPromptsExportRequest struct via the builder pattern


### Return type

[**AiPromptBundle**](AiPromptBundle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiPromptsGetById

> AiPrompt AiPromptsGetById(ctx).Id(id).Execute()

Get by id



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-prompts-get-by-id/).

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
	id := "id_example" // string | The saved prompt identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIPromptsAPI.AiPromptsGetById(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIPromptsAPI.AiPromptsGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiPromptsGetById`: AiPrompt
	fmt.Fprintf(os.Stdout, "Response from `AIPromptsAPI.AiPromptsGetById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiPromptsGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The saved prompt identifier. | 

### Return type

[**AiPrompt**](AiPrompt.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiPromptsGetFolderById

> AiPromptFolder AiPromptsGetFolderById(ctx).Id(id).Execute()

Get folder by id



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-prompts-get-folder-by-id/).

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
	id := "id_example" // string | The prompt folder identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIPromptsAPI.AiPromptsGetFolderById(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIPromptsAPI.AiPromptsGetFolderById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiPromptsGetFolderById`: AiPromptFolder
	fmt.Fprintf(os.Stdout, "Response from `AIPromptsAPI.AiPromptsGetFolderById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiPromptsGetFolderByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The prompt folder identifier. | 

### Return type

[**AiPromptFolder**](AiPromptFolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiPromptsImportBundle

> AiImportResult AiPromptsImportBundle(ctx).AiPromptsImportBundleRequest(aiPromptsImportBundleRequest).Execute()

Import bundle



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-prompts-import-bundle/).

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
	aiPromptsImportBundleRequest := *openapiclient.NewAiPromptsImportBundleRequest(*openapiclient.NewAiPromptBundle(float32(123), []openapiclient.AiPromptFolder{*openapiclient.NewAiPromptFolder("Id_example", "Name_example", float32(123), float32(123))}, []openapiclient.AiPrompt{*openapiclient.NewAiPrompt("Id_example", "Name_example", "Text_example", float32(123), float32(123))})) // AiPromptsImportBundleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIPromptsAPI.AiPromptsImportBundle(context.Background()).AiPromptsImportBundleRequest(aiPromptsImportBundleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIPromptsAPI.AiPromptsImportBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiPromptsImportBundle`: AiImportResult
	fmt.Fprintf(os.Stdout, "Response from `AIPromptsAPI.AiPromptsImportBundle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiPromptsImportBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiPromptsImportBundleRequest** | [**AiPromptsImportBundleRequest**](AiPromptsImportBundleRequest.md) |  | 

### Return type

[**AiImportResult**](AiImportResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiPromptsList

> []AiPrompt AiPromptsList(ctx).FolderId(folderId).Execute()

List



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-prompts-list/).

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
	folderId := "folderId_example" // string | The prompt folder identifier. Omit to list the prompts that sit outside any folder. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIPromptsAPI.AiPromptsList(context.Background()).FolderId(folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIPromptsAPI.AiPromptsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiPromptsList`: []AiPrompt
	fmt.Fprintf(os.Stdout, "Response from `AIPromptsAPI.AiPromptsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiPromptsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderId** | **string** | The prompt folder identifier. Omit to list the prompts that sit outside any folder. | 

### Return type

[**[]AiPrompt**](AiPrompt.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiPromptsListFolders

> []AiPromptFolder AiPromptsListFolders(ctx).Execute()

List folders



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-prompts-list-folders/).

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
	resp, r, err := apiClient.AIPromptsAPI.AiPromptsListFolders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIPromptsAPI.AiPromptsListFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiPromptsListFolders`: []AiPromptFolder
	fmt.Fprintf(os.Stdout, "Response from `AIPromptsAPI.AiPromptsListFolders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiPromptsListFoldersRequest struct via the builder pattern


### Return type

[**[]AiPromptFolder**](AiPromptFolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiPromptsMove

> AiPromptMutationResult AiPromptsMove(ctx).AiPromptsMoveRequest(aiPromptsMoveRequest).Execute()

Move



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-prompts-move/).

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
	aiPromptsMoveRequest := *openapiclient.NewAiPromptsMoveRequest("Id_example", "FolderId_example") // AiPromptsMoveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIPromptsAPI.AiPromptsMove(context.Background()).AiPromptsMoveRequest(aiPromptsMoveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIPromptsAPI.AiPromptsMove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiPromptsMove`: AiPromptMutationResult
	fmt.Fprintf(os.Stdout, "Response from `AIPromptsAPI.AiPromptsMove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiPromptsMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiPromptsMoveRequest** | [**AiPromptsMoveRequest**](AiPromptsMoveRequest.md) |  | 

### Return type

[**AiPromptMutationResult**](AiPromptMutationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiPromptsRenameFolder

> AiFolderMutationResult AiPromptsRenameFolder(ctx).AiPromptsRenameFolderRequest(aiPromptsRenameFolderRequest).Execute()

Rename folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-prompts-rename-folder/).

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
	aiPromptsRenameFolderRequest := *openapiclient.NewAiPromptsRenameFolderRequest("Id_example", "Name_example") // AiPromptsRenameFolderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIPromptsAPI.AiPromptsRenameFolder(context.Background()).AiPromptsRenameFolderRequest(aiPromptsRenameFolderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIPromptsAPI.AiPromptsRenameFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiPromptsRenameFolder`: AiFolderMutationResult
	fmt.Fprintf(os.Stdout, "Response from `AIPromptsAPI.AiPromptsRenameFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiPromptsRenameFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiPromptsRenameFolderRequest** | [**AiPromptsRenameFolderRequest**](AiPromptsRenameFolderRequest.md) |  | 

### Return type

[**AiFolderMutationResult**](AiFolderMutationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiPromptsUpdate

> AiPromptMutationResult AiPromptsUpdate(ctx).AiPromptsUpdateRequest(aiPromptsUpdateRequest).Execute()

Update



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/ai-prompts-update/).

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
	aiPromptsUpdateRequest := *openapiclient.NewAiPromptsUpdateRequest("Id_example", *openapiclient.NewAiPromptsUpdateRequestUpdates()) // AiPromptsUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AIPromptsAPI.AiPromptsUpdate(context.Background()).AiPromptsUpdateRequest(aiPromptsUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AIPromptsAPI.AiPromptsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiPromptsUpdate`: AiPromptMutationResult
	fmt.Fprintf(os.Stdout, "Response from `AIPromptsAPI.AiPromptsUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiPromptsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aiPromptsUpdateRequest** | [**AiPromptsUpdateRequest**](AiPromptsUpdateRequest.md) |  | 

### Return type

[**AiPromptMutationResult**](AiPromptMutationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

