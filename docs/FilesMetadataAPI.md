# \FilesMetadataAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignFileTemplates**](FilesMetadataAPI.md#AssignFileTemplates) | **Put** /api/2.0/files/metadata/file/{fileId}/templates | Assign templates to a file
[**AssignFolderTemplates**](FilesMetadataAPI.md#AssignFolderTemplates) | **Put** /api/2.0/files/metadata/folder/{folderId}/templates | Assign templates to a folder
[**CreateField**](FilesMetadataAPI.md#CreateField) | **Post** /api/2.0/files/metadata/templates/{templateId}/fields | Add a metadata field
[**CreateTemplate**](FilesMetadataAPI.md#CreateTemplate) | **Post** /api/2.0/files/metadata/templates | Create a metadata template
[**DeleteField**](FilesMetadataAPI.md#DeleteField) | **Delete** /api/2.0/files/metadata/templates/{templateId}/fields/{fieldId} | Delete a metadata field
[**DeleteTemplate**](FilesMetadataAPI.md#DeleteTemplate) | **Delete** /api/2.0/files/metadata/templates/{templateId} | Delete a metadata template
[**GetCascadeProgress**](FilesMetadataAPI.md#GetCascadeProgress) | **Get** /api/2.0/files/metadata/folder/{folderId}/templates/progress | Get cascade progress
[**GetFileMetadata**](FilesMetadataAPI.md#GetFileMetadata) | **Get** /api/2.0/files/metadata/file/{fileId} | Get file metadata
[**GetFolderMetadata**](FilesMetadataAPI.md#GetFolderMetadata) | **Get** /api/2.0/files/metadata/folder/{folderId} | Get folder metadata
[**GetTemplate**](FilesMetadataAPI.md#GetTemplate) | **Get** /api/2.0/files/metadata/templates/{templateId} | Get a metadata template
[**GetTemplates**](FilesMetadataAPI.md#GetTemplates) | **Get** /api/2.0/files/metadata/templates | Get metadata templates
[**SetFileCustomFields**](FilesMetadataAPI.md#SetFileCustomFields) | **Put** /api/2.0/files/metadata/file/{fileId}/customfields | Set file custom fields
[**SetFileValues**](FilesMetadataAPI.md#SetFileValues) | **Put** /api/2.0/files/metadata/file/{fileId}/values | Set file metadata values
[**SetFolderCustomFields**](FilesMetadataAPI.md#SetFolderCustomFields) | **Put** /api/2.0/files/metadata/folder/{folderId}/customfields | Set folder custom fields
[**SetFolderValues**](FilesMetadataAPI.md#SetFolderValues) | **Put** /api/2.0/files/metadata/folder/{folderId}/values | Set folder metadata values
[**UnassignFileTemplate**](FilesMetadataAPI.md#UnassignFileTemplate) | **Delete** /api/2.0/files/metadata/file/{fileId}/templates/{templateId} | Unassign a template from a file
[**UnassignFolderTemplate**](FilesMetadataAPI.md#UnassignFolderTemplate) | **Delete** /api/2.0/files/metadata/folder/{folderId}/templates/{templateId} | Unassign a template from a folder
[**UpdateField**](FilesMetadataAPI.md#UpdateField) | **Put** /api/2.0/files/metadata/templates/{templateId}/fields/{fieldId} | Update a metadata field
[**UpdateFolderTemplate**](FilesMetadataAPI.md#UpdateFolderTemplate) | **Put** /api/2.0/files/metadata/folder/{folderId}/templates/{templateId} | Update a template assignment on a folder
[**UpdateTemplate**](FilesMetadataAPI.md#UpdateTemplate) | **Put** /api/2.0/files/metadata/templates/{templateId} | Update a metadata template



## AssignFileTemplates

> AssignFileTemplates(ctx, fileId).AssignMetadataTemplates(assignMetadataTemplates).Execute()

Assign templates to a file



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/assign-file-templates/).

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
	fileId := int32(1) // int32 | The file ID.
	assignMetadataTemplates := *openapiclient.NewAssignMetadataTemplates([]int32{int32(123)}) // AssignMetadataTemplates | The parameters for assigning templates.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesMetadataAPI.AssignFileTemplates(context.Background(), fileId).AssignMetadataTemplates(assignMetadataTemplates).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.AssignFileTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignFileTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignMetadataTemplates** | [**AssignMetadataTemplates**](AssignMetadataTemplates.md) | The parameters for assigning templates. | 

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


## AssignFolderTemplates

> MetadataOperationWrapper AssignFolderTemplates(ctx, folderId).AssignMetadataTemplates(assignMetadataTemplates).Execute()

Assign templates to a folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/assign-folder-templates/).

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
	folderId := int32(1) // int32 | The folder ID.
	assignMetadataTemplates := *openapiclient.NewAssignMetadataTemplates([]int32{int32(123)}) // AssignMetadataTemplates | The parameters for assigning templates.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesMetadataAPI.AssignFolderTemplates(context.Background(), folderId).AssignMetadataTemplates(assignMetadataTemplates).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.AssignFolderTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignFolderTemplates`: MetadataOperationWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesMetadataAPI.AssignFolderTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignFolderTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignMetadataTemplates** | [**AssignMetadataTemplates**](AssignMetadataTemplates.md) | The parameters for assigning templates. | 

### Return type

[**MetadataOperationWrapper**](MetadataOperationWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateField

> MetadataFieldWrapper CreateField(ctx, templateId).MetadataFieldRequest(metadataFieldRequest).Execute()

Add a metadata field



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-field/).

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
	templateId := int32(1) // int32 | The template ID.
	metadataFieldRequest := *openapiclient.NewMetadataFieldRequest() // MetadataFieldRequest | The parameters of the field.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesMetadataAPI.CreateField(context.Background(), templateId).MetadataFieldRequest(metadataFieldRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.CreateField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateField`: MetadataFieldWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesMetadataAPI.CreateField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** | The template ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metadataFieldRequest** | [**MetadataFieldRequest**](MetadataFieldRequest.md) | The parameters of the field. | 

### Return type

[**MetadataFieldWrapper**](MetadataFieldWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTemplate

> MetadataTemplateWrapper CreateTemplate(ctx).CreateMetadataTemplateRequestDto(createMetadataTemplateRequestDto).Execute()

Create a metadata template



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-template/).

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
	createMetadataTemplateRequestDto := *openapiclient.NewCreateMetadataTemplateRequestDto("Contracts") // CreateMetadataTemplateRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesMetadataAPI.CreateTemplate(context.Background()).CreateMetadataTemplateRequestDto(createMetadataTemplateRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.CreateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTemplate`: MetadataTemplateWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesMetadataAPI.CreateTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMetadataTemplateRequestDto** | [**CreateMetadataTemplateRequestDto**](CreateMetadataTemplateRequestDto.md) |  | 

### Return type

[**MetadataTemplateWrapper**](MetadataTemplateWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteField

> DeleteField(ctx, templateId, fieldId).Execute()

Delete a metadata field



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-field/).

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
	templateId := int32(1) // int32 | The template ID.
	fieldId := int32(1) // int32 | The field ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesMetadataAPI.DeleteField(context.Background(), templateId, fieldId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.DeleteField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** | The template ID. | 
**fieldId** | **int32** | The field ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFieldRequest struct via the builder pattern


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


## DeleteTemplate

> DeleteTemplate(ctx, templateId).Execute()

Delete a metadata template



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-template/).

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
	templateId := int32(1) // int32 | The template ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesMetadataAPI.DeleteTemplate(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.DeleteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** | The template ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateRequest struct via the builder pattern


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


## GetCascadeProgress

> MetadataOperationWrapper GetCascadeProgress(ctx, folderId).Execute()

Get cascade progress



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-cascade-progress/).

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
	folderId := int32(1) // int32 | The folder the operation acts on. Take the identifier from a listing such as `GET api/2.0/files/@root` or  `GET api/2.0/files/{folderId}`: a folder stored in the portal is numbered, while a folder in a connected  third-party account is named by an opaque string.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesMetadataAPI.GetCascadeProgress(context.Background(), folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.GetCascadeProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCascadeProgress`: MetadataOperationWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesMetadataAPI.GetCascadeProgress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder the operation acts on. Take the identifier from a listing such as `GET api/2.0/files/@root` or  `GET api/2.0/files/{folderId}`: a folder stored in the portal is numbered, while a folder in a connected  third-party account is named by an opaque string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCascadeProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetadataOperationWrapper**](MetadataOperationWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileMetadata

> EntryMetadataWrapper GetFileMetadata(ctx, fileId).Execute()

Get file metadata



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-file-metadata/).

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
	fileId := int32(10) // int32 | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesMetadataAPI.GetFileMetadata(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.GetFileMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileMetadata`: EntryMetadataWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesMetadataAPI.GetFileMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntryMetadataWrapper**](EntryMetadataWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolderMetadata

> EntryMetadataWrapper GetFolderMetadata(ctx, folderId).Execute()

Get folder metadata



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-folder-metadata/).

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
	folderId := int32(1) // int32 | The folder the operation acts on. Take the identifier from a listing such as `GET api/2.0/files/@root` or  `GET api/2.0/files/{folderId}`: a folder stored in the portal is numbered, while a folder in a connected  third-party account is named by an opaque string.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesMetadataAPI.GetFolderMetadata(context.Background(), folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.GetFolderMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolderMetadata`: EntryMetadataWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesMetadataAPI.GetFolderMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder the operation acts on. Take the identifier from a listing such as `GET api/2.0/files/@root` or  `GET api/2.0/files/{folderId}`: a folder stored in the portal is numbered, while a folder in a connected  third-party account is named by an opaque string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntryMetadataWrapper**](EntryMetadataWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplate

> MetadataTemplateWrapper GetTemplate(ctx, templateId).Execute()

Get a metadata template



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-template/).

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
	templateId := int32(1) // int32 | The template ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesMetadataAPI.GetTemplate(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.GetTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplate`: MetadataTemplateWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesMetadataAPI.GetTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** | The template ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetadataTemplateWrapper**](MetadataTemplateWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplates

> MetadataTemplateArrayWrapper GetTemplates(ctx).Visible(visible).Execute()

Get metadata templates



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-templates/).

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
	visible := true // bool | Filters the templates by their visibility. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesMetadataAPI.GetTemplates(context.Background()).Visible(visible).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.GetTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplates`: MetadataTemplateArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesMetadataAPI.GetTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **visible** | **bool** | Filters the templates by their visibility. | 

### Return type

[**MetadataTemplateArrayWrapper**](MetadataTemplateArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFileCustomFields

> CustomFieldValueArrayWrapper SetFileCustomFields(ctx, fileId).SetCustomFields(setCustomFields).Execute()

Set file custom fields



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-file-custom-fields/).

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
	fileId := int32(1) // int32 | The file ID.
	setCustomFields := *openapiclient.NewSetCustomFields([]openapiclient.CustomFieldRequest{*openapiclient.NewCustomFieldRequest("Project code")}) // SetCustomFields | The custom fields to set.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesMetadataAPI.SetFileCustomFields(context.Background(), fileId).SetCustomFields(setCustomFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.SetFileCustomFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFileCustomFields`: CustomFieldValueArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesMetadataAPI.SetFileCustomFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFileCustomFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setCustomFields** | [**SetCustomFields**](SetCustomFields.md) | The custom fields to set. | 

### Return type

[**CustomFieldValueArrayWrapper**](CustomFieldValueArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFileValues

> EntryMetadataWrapper SetFileValues(ctx, fileId).SetMetadataValues(setMetadataValues).Execute()

Set file metadata values



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-file-values/).

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
	fileId := int32(1) // int32 | The file ID.
	setMetadataValues := *openapiclient.NewSetMetadataValues([]openapiclient.MetadataValueRequest{*openapiclient.NewMetadataValueRequest(int32(1))}) // SetMetadataValues | The parameters for setting values.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesMetadataAPI.SetFileValues(context.Background(), fileId).SetMetadataValues(setMetadataValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.SetFileValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFileValues`: EntryMetadataWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesMetadataAPI.SetFileValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFileValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setMetadataValues** | [**SetMetadataValues**](SetMetadataValues.md) | The parameters for setting values. | 

### Return type

[**EntryMetadataWrapper**](EntryMetadataWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFolderCustomFields

> CustomFieldValueArrayWrapper SetFolderCustomFields(ctx, folderId).SetCustomFields(setCustomFields).Execute()

Set folder custom fields



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-folder-custom-fields/).

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
	folderId := int32(1) // int32 | The folder ID.
	setCustomFields := *openapiclient.NewSetCustomFields([]openapiclient.CustomFieldRequest{*openapiclient.NewCustomFieldRequest("Project code")}) // SetCustomFields | The custom fields to set.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesMetadataAPI.SetFolderCustomFields(context.Background(), folderId).SetCustomFields(setCustomFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.SetFolderCustomFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFolderCustomFields`: CustomFieldValueArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesMetadataAPI.SetFolderCustomFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFolderCustomFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setCustomFields** | [**SetCustomFields**](SetCustomFields.md) | The custom fields to set. | 

### Return type

[**CustomFieldValueArrayWrapper**](CustomFieldValueArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFolderValues

> EntryMetadataWrapper SetFolderValues(ctx, folderId).SetMetadataValues(setMetadataValues).Execute()

Set folder metadata values



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-folder-values/).

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
	folderId := int32(1) // int32 | The folder ID.
	setMetadataValues := *openapiclient.NewSetMetadataValues([]openapiclient.MetadataValueRequest{*openapiclient.NewMetadataValueRequest(int32(1))}) // SetMetadataValues | The parameters for setting values.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesMetadataAPI.SetFolderValues(context.Background(), folderId).SetMetadataValues(setMetadataValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.SetFolderValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFolderValues`: EntryMetadataWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesMetadataAPI.SetFolderValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFolderValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setMetadataValues** | [**SetMetadataValues**](SetMetadataValues.md) | The parameters for setting values. | 

### Return type

[**EntryMetadataWrapper**](EntryMetadataWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignFileTemplate

> UnassignFileTemplate(ctx, fileId, templateId).Execute()

Unassign a template from a file



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/unassign-file-template/).

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
	fileId := int32(1) // int32 | The file ID.
	templateId := int32(1) // int32 | The template ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesMetadataAPI.UnassignFileTemplate(context.Background(), fileId, templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.UnassignFileTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID. | 
**templateId** | **int32** | The template ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignFileTemplateRequest struct via the builder pattern


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


## UnassignFolderTemplate

> UnassignFolderTemplate(ctx, folderId, templateId).Execute()

Unassign a template from a folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/unassign-folder-template/).

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
	folderId := int32(1) // int32 | The folder ID.
	templateId := int32(1) // int32 | The template ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesMetadataAPI.UnassignFolderTemplate(context.Background(), folderId, templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.UnassignFolderTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID. | 
**templateId** | **int32** | The template ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignFolderTemplateRequest struct via the builder pattern


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


## UpdateField

> MetadataFieldWrapper UpdateField(ctx, templateId, fieldId).UpdateMetadataFieldRequest(updateMetadataFieldRequest).Execute()

Update a metadata field



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-field/).

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
	templateId := int32(1) // int32 | The template ID.
	fieldId := int32(1) // int32 | The field ID.
	updateMetadataFieldRequest := *openapiclient.NewUpdateMetadataFieldRequest() // UpdateMetadataFieldRequest | The parameters of the field update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesMetadataAPI.UpdateField(context.Background(), templateId, fieldId).UpdateMetadataFieldRequest(updateMetadataFieldRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.UpdateField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateField`: MetadataFieldWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesMetadataAPI.UpdateField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** | The template ID. | 
**fieldId** | **int32** | The field ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMetadataFieldRequest** | [**UpdateMetadataFieldRequest**](UpdateMetadataFieldRequest.md) | The parameters of the field update. | 

### Return type

[**MetadataFieldWrapper**](MetadataFieldWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFolderTemplate

> MetadataOperationWrapper UpdateFolderTemplate(ctx, folderId, templateId).UpdateFolderMetadataTemplateRequest(updateFolderMetadataTemplateRequest).Execute()

Update a template assignment on a folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-folder-template/).

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
	folderId := int32(1) // int32 | The folder ID.
	templateId := int32(1) // int32 | The ID of a template the folder already carries.
	updateFolderMetadataTemplateRequest := *openapiclient.NewUpdateFolderMetadataTemplateRequest() // UpdateFolderMetadataTemplateRequest | The new state of the assignment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesMetadataAPI.UpdateFolderTemplate(context.Background(), folderId, templateId).UpdateFolderMetadataTemplateRequest(updateFolderMetadataTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.UpdateFolderTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFolderTemplate`: MetadataOperationWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesMetadataAPI.UpdateFolderTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID. | 
**templateId** | **int32** | The ID of a template the folder already carries. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFolderTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateFolderMetadataTemplateRequest** | [**UpdateFolderMetadataTemplateRequest**](UpdateFolderMetadataTemplateRequest.md) | The new state of the assignment. | 

### Return type

[**MetadataOperationWrapper**](MetadataOperationWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTemplate

> MetadataTemplateWrapper UpdateTemplate(ctx, templateId).UpdateMetadataTemplate(updateMetadataTemplate).Execute()

Update a metadata template



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-template/).

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
	templateId := int32(1) // int32 | The template ID.
	updateMetadataTemplate := *openapiclient.NewUpdateMetadataTemplate() // UpdateMetadataTemplate | The parameters for updating the template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesMetadataAPI.UpdateTemplate(context.Background(), templateId).UpdateMetadataTemplate(updateMetadataTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesMetadataAPI.UpdateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTemplate`: MetadataTemplateWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesMetadataAPI.UpdateTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** | The template ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMetadataTemplate** | [**UpdateMetadataTemplate**](UpdateMetadataTemplate.md) | The parameters for updating the template. | 

### Return type

[**MetadataTemplateWrapper**](MetadataTemplateWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

