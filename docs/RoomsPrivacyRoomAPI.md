# \RoomsPrivacyRoomAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKeys**](RoomsPrivacyRoomAPI.md#DeleteKeys) | **Delete** /api/2.0/privacyroom/keys/{id} | Delete an encryption key
[**GetUserKeys**](RoomsPrivacyRoomAPI.md#GetUserKeys) | **Get** /api/2.0/privacyroom/keys | Get own encryption keys
[**GetUserKeysForRoom**](RoomsPrivacyRoomAPI.md#GetUserKeysForRoom) | **Get** /api/2.0/privacyroom/{roomId}/access | Get private room access keys
[**ReplaceKey**](RoomsPrivacyRoomAPI.md#ReplaceKey) | **Put** /api/2.0/privacyroom/keys | Rotate an encryption key
[**SetKeys**](RoomsPrivacyRoomAPI.md#SetKeys) | **Post** /api/2.0/privacyroom/keys | Create an encryption key



## DeleteKeys

> DeleteKeys(ctx, id).Execute()

Delete an encryption key



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-keys/).

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
	id := "9924256B-447C-4F19-9dbd-8ad8c39e8ff5" // string | The pair to delete, taken from the `id` of an entry of `GET api/2.0/privacyroom/keys`. Only the caller's own  pairs can be named here.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoomsPrivacyRoomAPI.DeleteKeys(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsPrivacyRoomAPI.DeleteKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The pair to delete, taken from the `id` of an entry of `GET api/2.0/privacyroom/keys`. Only the caller's own  pairs can be named here. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeysRequest struct via the builder pattern


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


## GetUserKeys

> EncryptionKeyArrayWrapper GetUserKeys(ctx).Execute()

Get own encryption keys



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-user-keys/).

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
	resp, r, err := apiClient.RoomsPrivacyRoomAPI.GetUserKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsPrivacyRoomAPI.GetUserKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserKeys`: EncryptionKeyArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsPrivacyRoomAPI.GetUserKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserKeysRequest struct via the builder pattern


### Return type

[**EncryptionKeyArrayWrapper**](EncryptionKeyArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserKeysForRoom

> EncryptionKeyArrayWrapper GetUserKeysForRoom(ctx, roomId).Execute()

Get private room access keys



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-user-keys-for-room/).

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
	roomId := int32(56) // int32 | The private room whose access keys are read. Take it from the `id` of the room returned by  `POST api/2.0/files/rooms` or listed by `GET api/2.0/files/rooms`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsPrivacyRoomAPI.GetUserKeysForRoom(context.Background(), roomId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsPrivacyRoomAPI.GetUserKeysForRoom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserKeysForRoom`: EncryptionKeyArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsPrivacyRoomAPI.GetUserKeysForRoom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomId** | **int32** | The private room whose access keys are read. Take it from the `id` of the room returned by  `POST api/2.0/files/rooms` or listed by `GET api/2.0/files/rooms`. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserKeysForRoomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EncryptionKeyArrayWrapper**](EncryptionKeyArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceKey

> EncryptionKeyArrayWrapper ReplaceKey(ctx).EncryptionKeyRequestDto(encryptionKeyRequestDto).Execute()

Rotate an encryption key



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/replace-key/).

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
	encryptionKeyRequestDto := *openapiclient.NewEncryptionKeyRequestDto() // EncryptionKeyRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsPrivacyRoomAPI.ReplaceKey(context.Background()).EncryptionKeyRequestDto(encryptionKeyRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsPrivacyRoomAPI.ReplaceKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceKey`: EncryptionKeyArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsPrivacyRoomAPI.ReplaceKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **encryptionKeyRequestDto** | [**EncryptionKeyRequestDto**](EncryptionKeyRequestDto.md) |  | 

### Return type

[**EncryptionKeyArrayWrapper**](EncryptionKeyArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetKeys

> EncryptionKeyArrayWrapper SetKeys(ctx).EncryptionKeyRequestDto(encryptionKeyRequestDto).Execute()

Create an encryption key



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-keys/).

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
	encryptionKeyRequestDto := *openapiclient.NewEncryptionKeyRequestDto() // EncryptionKeyRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoomsPrivacyRoomAPI.SetKeys(context.Background()).EncryptionKeyRequestDto(encryptionKeyRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoomsPrivacyRoomAPI.SetKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetKeys`: EncryptionKeyArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `RoomsPrivacyRoomAPI.SetKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **encryptionKeyRequestDto** | [**EncryptionKeyRequestDto**](EncryptionKeyRequestDto.md) |  | 

### Return type

[**EncryptionKeyArrayWrapper**](EncryptionKeyArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

