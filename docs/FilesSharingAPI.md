# \FilesSharingAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyExternalSharePassword**](FilesSharingAPI.md#ApplyExternalSharePassword) | **Post** /api/2.0/files/share/{key}/password | Apply external data password
[**ChangeFileOwner**](FilesSharingAPI.md#ChangeFileOwner) | **Post** /api/2.0/files/owner | Change the file owner
[**GetEncryptionAccess**](FilesSharingAPI.md#GetEncryptionAccess) | **Get** /api/2.0/files/file/{fileId}/publickeys | Get file encryption keys
[**GetExternalShareData**](FilesSharingAPI.md#GetExternalShareData) | **Get** /api/2.0/files/share/{key} | Get the external data
[**GetFileSecurityInfo**](FilesSharingAPI.md#GetFileSecurityInfo) | **Get** /api/2.0/files/file/{id}/share | Get the shared file information
[**GetFolderSecurityInfo**](FilesSharingAPI.md#GetFolderSecurityInfo) | **Get** /api/2.0/files/folder/{id}/share | Get the shared folder information
[**GetGroupsMembersWithFileSecurity**](FilesSharingAPI.md#GetGroupsMembersWithFileSecurity) | **Get** /api/2.0/files/file/{fileId}/group/{groupId}/share | Get file group members with security information
[**GetGroupsMembersWithFolderSecurity**](FilesSharingAPI.md#GetGroupsMembersWithFolderSecurity) | **Get** /api/2.0/files/folder/{folderId}/group/{groupId}/share | Get folder group members with security information
[**GetSecurityInfo**](FilesSharingAPI.md#GetSecurityInfo) | **Post** /api/2.0/files/share | Get the sharing rights
[**GetSharedUsers**](FilesSharingAPI.md#GetSharedUsers) | **Get** /api/2.0/files/file/{fileId}/sharedusers | Get user access rights by file ID
[**RemoveSecurityInfo**](FilesSharingAPI.md#RemoveSecurityInfo) | **Delete** /api/2.0/files/share | Remove the sharing rights
[**SendEditorNotify**](FilesSharingAPI.md#SendEditorNotify) | **Post** /api/2.0/files/file/{fileId}/sendeditornotify | Send the mention message
[**SetFileSecurityInfo**](FilesSharingAPI.md#SetFileSecurityInfo) | **Put** /api/2.0/files/file/{id}/share | Share a file
[**SetFolderSecurityInfo**](FilesSharingAPI.md#SetFolderSecurityInfo) | **Put** /api/2.0/files/folder/{id}/share | Share a folder
[**SetSecurityInfo**](FilesSharingAPI.md#SetSecurityInfo) | **Put** /api/2.0/files/share | Set the sharing rights



## ApplyExternalSharePassword

> ExternalShareWrapper ApplyExternalSharePassword(ctx, key).ExternalShareRequestParam(externalShareRequestParam).Execute()

Apply external data password



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/apply-external-share-password/).

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
	key := "doc_key_123" // string | The unique document identifier.
	externalShareRequestParam := *openapiclient.NewExternalShareRequestParam() // ExternalShareRequestParam | The external data share request parameters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSharingAPI.ApplyExternalSharePassword(context.Background(), key).ExternalShareRequestParam(externalShareRequestParam).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingAPI.ApplyExternalSharePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyExternalSharePassword`: ExternalShareWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSharingAPI.ApplyExternalSharePassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The unique document identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyExternalSharePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalShareRequestParam** | [**ExternalShareRequestParam**](ExternalShareRequestParam.md) | The external data share request parameters. | 

### Return type

[**ExternalShareWrapper**](ExternalShareWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeFileOwner

> FileEntryBaseArrayWrapper ChangeFileOwner(ctx).ChangeOwnerRequestDto(changeOwnerRequestDto).Execute()

Change the file owner



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/change-file-owner/).

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
	changeOwnerRequestDto := *openapiclient.NewChangeOwnerRequestDto("00000000-0000-0000-0000-000000000000") // ChangeOwnerRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSharingAPI.ChangeFileOwner(context.Background()).ChangeOwnerRequestDto(changeOwnerRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingAPI.ChangeFileOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeFileOwner`: FileEntryBaseArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSharingAPI.ChangeFileOwner`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeFileOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changeOwnerRequestDto** | [**ChangeOwnerRequestDto**](ChangeOwnerRequestDto.md) |  | 

### Return type

[**FileEntryBaseArrayWrapper**](FileEntryBaseArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEncryptionAccess

> EncryptionKeyArrayWrapper GetEncryptionAccess(ctx, fileId).Execute()

Get file encryption keys



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-encryption-access/).

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
	fileId := int32(1) // int32 | The file unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSharingAPI.GetEncryptionAccess(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingAPI.GetEncryptionAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEncryptionAccess`: EncryptionKeyArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSharingAPI.GetEncryptionAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncryptionAccessRequest struct via the builder pattern


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


## GetExternalShareData

> ExternalShareWrapper GetExternalShareData(ctx, key).FileId(fileId).FolderId(folderId).Execute()

Get the external data



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-external-share-data/).

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
	key := "doc_key_123" // string | The unique key of the external shared data.
	fileId := "1" // string | The unique document identifier. (optional)
	folderId := "1" // string | The unique folder identifier. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSharingAPI.GetExternalShareData(context.Background(), key).FileId(fileId).FolderId(folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingAPI.GetExternalShareData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalShareData`: ExternalShareWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSharingAPI.GetExternalShareData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The unique key of the external shared data. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalShareDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileId** | **string** | The unique document identifier. | 
 **folderId** | **string** | The unique folder identifier. | 

### Return type

[**ExternalShareWrapper**](ExternalShareWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileSecurityInfo

> FileShareArrayWrapper GetFileSecurityInfo(ctx, id).Count(count).StartIndex(startIndex).Execute()

Get the shared file information



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-file-security-info/).

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
	id := int32(10) // int32 | The file unique identifier.
	count := int32(25) // int32 | The number of items to retrieve in the request. (optional)
	startIndex := int32(0) // int32 | The starting index for the query results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSharingAPI.GetFileSecurityInfo(context.Background(), id).Count(count).StartIndex(startIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingAPI.GetFileSecurityInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileSecurityInfo`: FileShareArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSharingAPI.GetFileSecurityInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The file unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileSecurityInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The number of items to retrieve in the request. | 
 **startIndex** | **int32** | The starting index for the query results. | 

### Return type

[**FileShareArrayWrapper**](FileShareArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolderSecurityInfo

> FileShareArrayWrapper GetFolderSecurityInfo(ctx, id).Count(count).StartIndex(startIndex).Execute()

Get the shared folder information



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-folder-security-info/).

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
	id := int32(10) // int32 | The folder unique identifier.
	count := int32(25) // int32 | The number of items to retrieve in the request. (optional)
	startIndex := int32(0) // int32 | The starting index for the query results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSharingAPI.GetFolderSecurityInfo(context.Background(), id).Count(count).StartIndex(startIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingAPI.GetFolderSecurityInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFolderSecurityInfo`: FileShareArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSharingAPI.GetFolderSecurityInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The folder unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderSecurityInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | The number of items to retrieve in the request. | 
 **startIndex** | **int32** | The starting index for the query results. | 

### Return type

[**FileShareArrayWrapper**](FileShareArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupsMembersWithFileSecurity

> GroupMemberSecurityRequestArrayWrapper GetGroupsMembersWithFileSecurity(ctx, fileId, groupId).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()

Get file group members with security information



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-groups-members-with-file-security/).

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
	groupId := "00000000-0000-0000-0000-000000000000" // string | The group ID.
	count := int32(25) // int32 | The number of items to be retrieved in the current query. (optional)
	startIndex := int32(0) // int32 | The starting index for the query result set. (optional)
	filterValue := "My Document" // string | The filter value used for searching or querying group members based on text input. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSharingAPI.GetGroupsMembersWithFileSecurity(context.Background(), fileId, groupId).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingAPI.GetGroupsMembersWithFileSecurity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupsMembersWithFileSecurity`: GroupMemberSecurityRequestArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSharingAPI.GetGroupsMembersWithFileSecurity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID. | 
**groupId** | **string** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsMembersWithFileSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | The number of items to be retrieved in the current query. | 
 **startIndex** | **int32** | The starting index for the query result set. | 
 **filterValue** | **string** | The filter value used for searching or querying group members based on text input. | 

### Return type

[**GroupMemberSecurityRequestArrayWrapper**](GroupMemberSecurityRequestArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupsMembersWithFolderSecurity

> GroupMemberSecurityRequestArrayWrapper GetGroupsMembersWithFolderSecurity(ctx, folderId, groupId).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()

Get folder group members with security information



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-groups-members-with-folder-security/).

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
	groupId := "00000000-0000-0000-0000-000000000000" // string | The group ID.
	count := int32(25) // int32 | The number of items to be retrieved in the current query. (optional)
	startIndex := int32(0) // int32 | The starting index for the query result set. (optional)
	filterValue := "My Document" // string | The filter value used for searching or querying group members based on text input. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSharingAPI.GetGroupsMembersWithFolderSecurity(context.Background(), folderId, groupId).Count(count).StartIndex(startIndex).FilterValue(filterValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingAPI.GetGroupsMembersWithFolderSecurity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupsMembersWithFolderSecurity`: GroupMemberSecurityRequestArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSharingAPI.GetGroupsMembersWithFolderSecurity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **int32** | The folder ID. | 
**groupId** | **string** | The group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsMembersWithFolderSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | The number of items to be retrieved in the current query. | 
 **startIndex** | **int32** | The starting index for the query result set. | 
 **filterValue** | **string** | The filter value used for searching or querying group members based on text input. | 

### Return type

[**GroupMemberSecurityRequestArrayWrapper**](GroupMemberSecurityRequestArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityInfo

> FileShareArrayWrapper GetSecurityInfo(ctx).BaseBatchRequestDto(baseBatchRequestDto).Execute()

Get the sharing rights



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-security-info/).

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
	baseBatchRequestDto := *openapiclient.NewBaseBatchRequestDto() // BaseBatchRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSharingAPI.GetSecurityInfo(context.Background()).BaseBatchRequestDto(baseBatchRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingAPI.GetSecurityInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityInfo`: FileShareArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSharingAPI.GetSecurityInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseBatchRequestDto** | [**BaseBatchRequestDto**](BaseBatchRequestDto.md) |  | 

### Return type

[**FileShareArrayWrapper**](FileShareArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSharedUsers

> MentionWrapperArrayWrapper GetSharedUsers(ctx, fileId).Execute()

Get user access rights by file ID



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-shared-users/).

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
	fileId := int32(1) // int32 | The file unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSharingAPI.GetSharedUsers(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingAPI.GetSharedUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSharedUsers`: MentionWrapperArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSharingAPI.GetSharedUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSharedUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MentionWrapperArrayWrapper**](MentionWrapperArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSecurityInfo

> BooleanWrapper RemoveSecurityInfo(ctx).BaseBatchRequestDto(baseBatchRequestDto).Execute()

Remove the sharing rights



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/remove-security-info/).

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
	baseBatchRequestDto := *openapiclient.NewBaseBatchRequestDto() // BaseBatchRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSharingAPI.RemoveSecurityInfo(context.Background()).BaseBatchRequestDto(baseBatchRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingAPI.RemoveSecurityInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveSecurityInfo`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSharingAPI.RemoveSecurityInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSecurityInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **baseBatchRequestDto** | [**BaseBatchRequestDto**](BaseBatchRequestDto.md) |  | 

### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendEditorNotify

> AceShortWrapperArrayWrapper SendEditorNotify(ctx, fileId).MentionMessageWrapper(mentionMessageWrapper).Execute()

Send the mention message



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/send-editor-notify/).

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
	fileId := int32(file-id) // int32 | The file ID with the mention message.
	mentionMessageWrapper := *openapiclient.NewMentionMessageWrapper() // MentionMessageWrapper | The mention message. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSharingAPI.SendEditorNotify(context.Background(), fileId).MentionMessageWrapper(mentionMessageWrapper).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingAPI.SendEditorNotify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendEditorNotify`: AceShortWrapperArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSharingAPI.SendEditorNotify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int32** | The file ID with the mention message. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendEditorNotifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mentionMessageWrapper** | [**MentionMessageWrapper**](MentionMessageWrapper.md) | The mention message. | 

### Return type

[**AceShortWrapperArrayWrapper**](AceShortWrapperArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFileSecurityInfo

> FileShareArrayWrapper SetFileSecurityInfo(ctx, id).SecurityInfoSimpleRequestDto(securityInfoSimpleRequestDto).Execute()

Share a file



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-file-security-info/).

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
	id := int32(1) // int32 | The file ID.
	securityInfoSimpleRequestDto := *openapiclient.NewSecurityInfoSimpleRequestDto() // SecurityInfoSimpleRequestDto | The parameters of the security information simple request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSharingAPI.SetFileSecurityInfo(context.Background(), id).SecurityInfoSimpleRequestDto(securityInfoSimpleRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingAPI.SetFileSecurityInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFileSecurityInfo`: FileShareArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSharingAPI.SetFileSecurityInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The file ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFileSecurityInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **securityInfoSimpleRequestDto** | [**SecurityInfoSimpleRequestDto**](SecurityInfoSimpleRequestDto.md) | The parameters of the security information simple request. | 

### Return type

[**FileShareArrayWrapper**](FileShareArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFolderSecurityInfo

> FileShareArrayWrapper SetFolderSecurityInfo(ctx, id).SecurityInfoSimpleRequestDto(securityInfoSimpleRequestDto).Execute()

Share a folder



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-folder-security-info/).

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
	id := int32(1) // int32 | The folder ID.
	securityInfoSimpleRequestDto := *openapiclient.NewSecurityInfoSimpleRequestDto() // SecurityInfoSimpleRequestDto | The parameters of the security information simple request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSharingAPI.SetFolderSecurityInfo(context.Background(), id).SecurityInfoSimpleRequestDto(securityInfoSimpleRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingAPI.SetFolderSecurityInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFolderSecurityInfo`: FileShareArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSharingAPI.SetFolderSecurityInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The folder ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFolderSecurityInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **securityInfoSimpleRequestDto** | [**SecurityInfoSimpleRequestDto**](SecurityInfoSimpleRequestDto.md) | The parameters of the security information simple request. | 

### Return type

[**FileShareArrayWrapper**](FileShareArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSecurityInfo

> FileShareArrayWrapper SetSecurityInfo(ctx).SecurityInfoRequestDto(securityInfoRequestDto).Execute()

Set the sharing rights



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-security-info/).

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
	securityInfoRequestDto := *openapiclient.NewSecurityInfoRequestDto() // SecurityInfoRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesSharingAPI.SetSecurityInfo(context.Background()).SecurityInfoRequestDto(securityInfoRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingAPI.SetSecurityInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSecurityInfo`: FileShareArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `FilesSharingAPI.SetSecurityInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetSecurityInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityInfoRequestDto** | [**SecurityInfoRequestDto**](SecurityInfoRequestDto.md) |  | 

### Return type

[**FileShareArrayWrapper**](FileShareArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

