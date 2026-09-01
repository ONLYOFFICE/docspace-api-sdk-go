# \PeoplePhotosAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMemberPhotoThumbnails**](PeoplePhotosAPI.md#CreateMemberPhotoThumbnails) | **Post** /api/2.0/people/{userid}/photo/thumbnails | Create photo thumbnails
[**DeleteMemberPhoto**](PeoplePhotosAPI.md#DeleteMemberPhoto) | **Delete** /api/2.0/people/{userid}/photo | Delete a user photo
[**GetMemberPhoto**](PeoplePhotosAPI.md#GetMemberPhoto) | **Get** /api/2.0/people/{userid}/photo | Get a user photo
[**UpdateMemberPhoto**](PeoplePhotosAPI.md#UpdateMemberPhoto) | **Put** /api/2.0/people/{userid}/photo | Update a user photo
[**UploadMemberPhoto**](PeoplePhotosAPI.md#UploadMemberPhoto) | **Post** /api/2.0/people/{userid}/photo | Upload a user photo



## CreateMemberPhotoThumbnails

> ThumbnailsDataWrapper CreateMemberPhotoThumbnails(ctx, userid).ThumbnailsRequest(thumbnailsRequest).Execute()

Create photo thumbnails



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-member-photo-thumbnails/).

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
	thumbnailsRequest := *openapiclient.NewThumbnailsRequest() // ThumbnailsRequest | The thumbnail request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeoplePhotosAPI.CreateMemberPhotoThumbnails(context.Background(), userid).ThumbnailsRequest(thumbnailsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeoplePhotosAPI.CreateMemberPhotoThumbnails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMemberPhotoThumbnails`: ThumbnailsDataWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeoplePhotosAPI.CreateMemberPhotoThumbnails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMemberPhotoThumbnailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **thumbnailsRequest** | [**ThumbnailsRequest**](ThumbnailsRequest.md) | The thumbnail request. | 

### Return type

[**ThumbnailsDataWrapper**](ThumbnailsDataWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMemberPhoto

> ThumbnailsDataWrapper DeleteMemberPhoto(ctx, userid).Execute()

Delete a user photo



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-member-photo/).

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
	resp, r, err := apiClient.PeoplePhotosAPI.DeleteMemberPhoto(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeoplePhotosAPI.DeleteMemberPhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMemberPhoto`: ThumbnailsDataWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeoplePhotosAPI.DeleteMemberPhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMemberPhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ThumbnailsDataWrapper**](ThumbnailsDataWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemberPhoto

> ThumbnailsDataWrapper GetMemberPhoto(ctx, userid).Execute()

Get a user photo



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-member-photo/).

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
	resp, r, err := apiClient.PeoplePhotosAPI.GetMemberPhoto(context.Background(), userid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeoplePhotosAPI.GetMemberPhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMemberPhoto`: ThumbnailsDataWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeoplePhotosAPI.GetMemberPhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMemberPhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ThumbnailsDataWrapper**](ThumbnailsDataWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMemberPhoto

> ThumbnailsDataWrapper UpdateMemberPhoto(ctx, userid).UpdatePhotoMemberRequest(updatePhotoMemberRequest).Execute()

Update a user photo



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-member-photo/).

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
	updatePhotoMemberRequest := *openapiclient.NewUpdatePhotoMemberRequest() // UpdatePhotoMemberRequest | The request parameters for updating a photo.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeoplePhotosAPI.UpdateMemberPhoto(context.Background(), userid).UpdatePhotoMemberRequest(updatePhotoMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeoplePhotosAPI.UpdateMemberPhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMemberPhoto`: ThumbnailsDataWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeoplePhotosAPI.UpdateMemberPhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMemberPhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePhotoMemberRequest** | [**UpdatePhotoMemberRequest**](UpdatePhotoMemberRequest.md) | The request parameters for updating a photo. | 

### Return type

[**ThumbnailsDataWrapper**](ThumbnailsDataWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadMemberPhoto

> FileUploadResultWrapper UploadMemberPhoto(ctx, userid).File(file).Autosave(autosave).Execute()

Upload a user photo



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/upload-member-photo/).

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
	file := os.NewFile(1234, "some_file") // *os.File | The image data.
	autosave := true // bool | Specifies whether to autosave a photo or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PeoplePhotosAPI.UploadMemberPhoto(context.Background(), userid).File(file).Autosave(autosave).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PeoplePhotosAPI.UploadMemberPhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadMemberPhoto`: FileUploadResultWrapper
	fmt.Fprintf(os.Stdout, "Response from `PeoplePhotosAPI.UploadMemberPhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userid** | **string** | The user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadMemberPhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The image data. | 
 **autosave** | **bool** | Specifies whether to autosave a photo or not. | 

### Return type

[**FileUploadResultWrapper**](FileUploadResultWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

