# \RoomsPrivacyRoomAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKeys**](RoomsPrivacyRoomAPI.md#DeleteKeys) | **Delete** /api/2.0/privacyroom/keys/{id} | Deletes an encryption key and removes it from the system.
[**GetUserKeys**](RoomsPrivacyRoomAPI.md#GetUserKeys) | **Get** /api/2.0/privacyroom/keys | Retrieves encryption keys associated with the current user.
[**GetUserKeysForRoom**](RoomsPrivacyRoomAPI.md#GetUserKeysForRoom) | **Get** /api/2.0/privacyroom/{roomId}/access | Retrieves the encryption keys associated with a specific privacy room.
[**ReplaceKey**](RoomsPrivacyRoomAPI.md#ReplaceKey) | **Put** /api/2.0/privacyroom/keys | Replaces an existing encryption key with a new one for the user.
[**SetKeys**](RoomsPrivacyRoomAPI.md#SetKeys) | **Post** /api/2.0/privacyroom/keys | Creates and sets encryption keys for the user.



## DeleteKeys

> DeleteKeys(ctx, id).Execute()

Deletes an encryption key and removes it from the system.



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
	id := "00000000-0000-0000-0000-000000000000" // string | The unique identifier of the encryption key to be deleted.

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
**id** | **string** | The unique identifier of the encryption key to be deleted. | 

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

Retrieves encryption keys associated with the current user.



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

Retrieves the encryption keys associated with a specific privacy room.



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
	roomId := int32(56) // int32 | The identifier of the privacy room.

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
**roomId** | **int32** | The identifier of the privacy room. | 

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

Replaces an existing encryption key with a new one for the user.



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
	encryptionKeyRequestDto := *openapiclient.NewEncryptionKeyRequestDto() // EncryptionKeyRequestDto | The request object containing the public and private key information to replace the existing key. (optional)

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
 **encryptionKeyRequestDto** | [**EncryptionKeyRequestDto**](EncryptionKeyRequestDto.md) | The request object containing the public and private key information to replace the existing key. | 

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

Creates and sets encryption keys for the user.



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
	encryptionKeyRequestDto := *openapiclient.NewEncryptionKeyRequestDto() // EncryptionKeyRequestDto | The request object containing public and private key information. (optional)

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
 **encryptionKeyRequestDto** | [**EncryptionKeyRequestDto**](EncryptionKeyRequestDto.md) | The request object containing public and private key information. | 

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

