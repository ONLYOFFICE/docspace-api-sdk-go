# \FilesSharingAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyExternalSharePassword**](FilesSharingAPI.md#ApplyExternalSharePassword) | **Post** /api/2.0/files/share/{key}/password | Unlock a password-protected link
[**ChangeFileOwner**](FilesSharingAPI.md#ChangeFileOwner) | **Post** /api/2.0/files/owner | Change the room or file owner
[**GetEncryptionAccess**](FilesSharingAPI.md#GetEncryptionAccess) | **Get** /api/2.0/files/file/{fileId}/publickeys | Get file encryption keys
[**GetExternalShareData**](FilesSharingAPI.md#GetExternalShareData) | **Get** /api/2.0/files/share/{key} | Resolve an external share link
[**GetFileSecurityInfo**](FilesSharingAPI.md#GetFileSecurityInfo) | **Get** /api/2.0/files/file/{id}/share | Get file sharing rights
[**GetFolderSecurityInfo**](FilesSharingAPI.md#GetFolderSecurityInfo) | **Get** /api/2.0/files/folder/{id}/share | Get folder sharing rights
[**GetGroupsMembersWithFileSecurity**](FilesSharingAPI.md#GetGroupsMembersWithFileSecurity) | **Get** /api/2.0/files/file/{fileId}/group/{groupId}/share | Get file access of group members
[**GetGroupsMembersWithFolderSecurity**](FilesSharingAPI.md#GetGroupsMembersWithFolderSecurity) | **Get** /api/2.0/files/folder/{folderId}/group/{groupId}/share | Get folder access of group members
[**GetSecurityInfo**](FilesSharingAPI.md#GetSecurityInfo) | **Post** /api/2.0/files/share | Get sharing rights in batch
[**GetSharedUsers**](FilesSharingAPI.md#GetSharedUsers) | **Get** /api/2.0/files/file/{fileId}/sharedusers | Get users to mention in a file
[**RemoveSecurityInfo**](FilesSharingAPI.md#RemoveSecurityInfo) | **Delete** /api/2.0/files/share | Remove sharing rights in batch
[**SendEditorNotify**](FilesSharingAPI.md#SendEditorNotify) | **Post** /api/2.0/files/file/{fileId}/sendeditornotify | Notify mentioned users
[**SetFileSecurityInfo**](FilesSharingAPI.md#SetFileSecurityInfo) | **Put** /api/2.0/files/file/{id}/share | Share a file
[**SetFolderSecurityInfo**](FilesSharingAPI.md#SetFolderSecurityInfo) | **Put** /api/2.0/files/folder/{id}/share | Share a folder
[**SetSecurityInfo**](FilesSharingAPI.md#SetSecurityInfo) | **Put** /api/2.0/files/share | Set sharing rights in batch



## ApplyExternalSharePassword

> ExternalShareWrapper ApplyExternalSharePassword(ctx, key).ExternalShareRequestParam(externalShareRequestParam).Execute()

Unlock a password-protected link



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
	key := "q7Ry8cQ1lZ0dP3sK2mXfA9tBnV6hJ4uE8wCz5oLg" // string | The token of the external share link, taken verbatim from the `requestToken` of a link returned by the link  operations of an entry, such as `GET api/2.0/files/rooms/{id}/link`. It is an opaque URL-safe string that  carries the link's own identifier, so it cannot be assembled by hand.
	externalShareRequestParam := *openapiclient.NewExternalShareRequestParam() // ExternalShareRequestParam | The body of the request, holding the password to check.

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
**key** | **string** | The token of the external share link, taken verbatim from the `requestToken` of a link returned by the link  operations of an entry, such as `GET api/2.0/files/rooms/{id}/link`. It is an opaque URL-safe string that  carries the link's own identifier, so it cannot be assembled by hand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyExternalSharePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalShareRequestParam** | [**ExternalShareRequestParam**](ExternalShareRequestParam.md) | The body of the request, holding the password to check. | 

### Return type

[**ExternalShareWrapper**](ExternalShareWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeFileOwner

> FileEntryBaseArrayWrapper ChangeFileOwner(ctx).ChangeOwnerRequestDto(changeOwnerRequestDto).Execute()

Change the room or file owner



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
	changeOwnerRequestDto := *openapiclient.NewChangeOwnerRequestDto("9924256a-739c-462b-af15-e652a3b1b6eb") // ChangeOwnerRequestDto |  (optional)

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
	fileId := int32(10) // int32 | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

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
**fileId** | **int32** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

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

Resolve an external share link



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
	key := "q7Ry8cQ1lZ0dP3sK2mXfA9tBnV6hJ4uE8wCz5oLg" // string | The token of the external share link, taken verbatim from the `requestToken` of a link returned by the link  operations of an entry, such as `GET api/2.0/files/rooms/{id}/link`. It is an opaque URL-safe string that  carries the link's own identifier, so it cannot be assembled by hand.
	fileId := "9" // string | A file inside the room the link points at, echoed back in the answer's entity fields so a client can show what  was opened. The value is ignored when the file does not sit under the link's target, and passing it together  with a folder has no effect - the file wins. (optional)
	folderId := "3" // string | A folder inside the room the link points at, echoed back in the answer's entity fields. It is ignored when the  folder does not sit under the link's target, and when a file is passed as well. (optional)

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
**key** | **string** | The token of the external share link, taken verbatim from the `requestToken` of a link returned by the link  operations of an entry, such as `GET api/2.0/files/rooms/{id}/link`. It is an opaque URL-safe string that  carries the link's own identifier, so it cannot be assembled by hand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalShareDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileId** | **string** | A file inside the room the link points at, echoed back in the answer's entity fields so a client can show what  was opened. The value is ignored when the file does not sit under the link's target, and passing it together  with a folder has no effect - the file wins. | 
 **folderId** | **string** | A folder inside the room the link points at, echoed back in the answer's entity fields. It is ignored when the  folder does not sit under the link's target, and when a file is passed as well. | 

### Return type

[**ExternalShareWrapper**](ExternalShareWrapper.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileSecurityInfo

> FileShareArrayWrapper GetFileSecurityInfo(ctx, id).Count(count).StartIndex(startIndex).Execute()

Get file sharing rights



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
	id := int32(10) // int32 | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.
	count := int32(25) // int32 | How many entries at most to answer with, in the operations of this file that return a list; an operation that  answers with a single object is not affected by it. (optional)
	startIndex := int32(0) // int32 | How many entries of such a list to skip before answering, used together with `count` to walk through it page  by page. (optional)

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
**id** | **int32** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileSecurityInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | How many entries at most to answer with, in the operations of this file that return a list; an operation that  answers with a single object is not affected by it. | 
 **startIndex** | **int32** | How many entries of such a list to skip before answering, used together with `count` to walk through it page  by page. | 

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

Get folder sharing rights



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
	id := int32(10) // int32 | The folder or room the operation addresses. A folder stored on the portal is numbered, while a folder in a  connected third-party account is named by an opaque string.
	count := int32(25) // int32 | How many entries at most to answer with, in the operations of this folder that return a list; an operation  that answers with a single object is not affected by it. (optional)
	startIndex := int32(0) // int32 | How many entries of such a list to skip before answering, used together with `count` to walk through it page  by page. (optional)

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
**id** | **int32** | The folder or room the operation addresses. A folder stored on the portal is numbered, while a folder in a  connected third-party account is named by an opaque string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderSecurityInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | How many entries at most to answer with, in the operations of this folder that return a list; an operation  that answers with a single object is not affected by it. | 
 **startIndex** | **int32** | How many entries of such a list to skip before answering, used together with `count` to walk through it page  by page. | 

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

Get file access of group members



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
	fileId := int32(10) // int32 | The file whose access is being read. A file stored on the portal is numbered, while a file in a connected  third-party account is named by an opaque string.
	groupId := "9924256a-739c-462b-af15-e652a3b1b6eb" // string | The group whose members are listed. Take it from the entries of `GET api/2.0/files/file/{id}/share` that stand  for a group; a group that holds no rights on this file is answered with an empty list.
	count := int32(25) // int32 | How many members at most to answer with. (optional)
	startIndex := int32(0) // int32 | How many members to skip before answering, used together with `count` to page through a large group. (optional)
	filterValue := "john" // string | Keeps only the members whose first name, last name or email contains this value. The value is matched in lower  case, so an uppercase one finds nothing. (optional)

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
**fileId** | **int32** | The file whose access is being read. A file stored on the portal is numbered, while a file in a connected  third-party account is named by an opaque string. | 
**groupId** | **string** | The group whose members are listed. Take it from the entries of `GET api/2.0/files/file/{id}/share` that stand  for a group; a group that holds no rights on this file is answered with an empty list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsMembersWithFileSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | How many members at most to answer with. | 
 **startIndex** | **int32** | How many members to skip before answering, used together with `count` to page through a large group. | 
 **filterValue** | **string** | Keeps only the members whose first name, last name or email contains this value. The value is matched in lower  case, so an uppercase one finds nothing. | 

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

Get folder access of group members



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
	folderId := int32(10) // int32 | The folder or room whose access is being read. A folder stored on the portal is numbered, while a folder in a  connected third-party account is named by an opaque string.
	groupId := "9924256a-739c-462b-af15-e652a3b1b6eb" // string | The group whose members are listed. Take it from the entries of `GET api/2.0/files/folder/{id}/share` that  stand for a group; a group that holds no rights on this folder is answered with an empty list.
	count := int32(25) // int32 | How many members at most to answer with. (optional)
	startIndex := int32(0) // int32 | How many members to skip before answering, used together with `count` to page through a large group. (optional)
	filterValue := "john" // string | Keeps only the members whose first name, last name or email contains this value. The value is matched in lower  case, so an uppercase one finds nothing. (optional)

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
**folderId** | **int32** | The folder or room whose access is being read. A folder stored on the portal is numbered, while a folder in a  connected third-party account is named by an opaque string. | 
**groupId** | **string** | The group whose members are listed. Take it from the entries of `GET api/2.0/files/folder/{id}/share` that  stand for a group; a group that holds no rights on this folder is answered with an empty list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsMembersWithFolderSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **int32** | How many members at most to answer with. | 
 **startIndex** | **int32** | How many members to skip before answering, used together with `count` to page through a large group. | 
 **filterValue** | **string** | Keeps only the members whose first name, last name or email contains this value. The value is matched in lower  case, so an uppercase one finds nothing. | 

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

Get sharing rights in batch



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

Get users to mention in a file



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
	fileId := int32(10) // int32 | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string.

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
**fileId** | **int32** | The file the operation addresses. Take the identifier from a listing such as `GET api/2.0/files/{folderId}`: a  file stored on the portal is numbered, while a file in a connected third-party account is named by an opaque  string. | 

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

Remove sharing rights in batch



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

Notify mentioned users



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
	fileId := int32(10) // int32 | The file the mention was made in. A file stored on the portal is numbered, while a file in a connected  third-party account is named by an opaque string.
	mentionMessageWrapper := *openapiclient.NewMentionMessageWrapper() // MentionMessageWrapper | The notification to send. (optional)

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
**fileId** | **int32** | The file the mention was made in. A file stored on the portal is numbered, while a file in a connected  third-party account is named by an opaque string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendEditorNotifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mentionMessageWrapper** | [**MentionMessageWrapper**](MentionMessageWrapper.md) | The notification to send. | 

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
	id := int32(10) // int32 | The file whose sharing is being changed. A file stored on the portal is numbered, while a file in a connected  third-party account is named by an opaque string.
	securityInfoSimpleRequestDto := *openapiclient.NewSecurityInfoSimpleRequestDto() // SecurityInfoSimpleRequestDto | The rights to apply to the file, and whether to announce them by mail.

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
**id** | **int32** | The file whose sharing is being changed. A file stored on the portal is numbered, while a file in a connected  third-party account is named by an opaque string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFileSecurityInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **securityInfoSimpleRequestDto** | [**SecurityInfoSimpleRequestDto**](SecurityInfoSimpleRequestDto.md) | The rights to apply to the file, and whether to announce them by mail. | 

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
	id := int32(10) // int32 | The folder whose sharing is being changed. A folder stored on the portal is numbered, while a folder in a  connected third-party account is named by an opaque string.
	securityInfoSimpleRequestDto := *openapiclient.NewSecurityInfoSimpleRequestDto() // SecurityInfoSimpleRequestDto | The rights to apply to the folder, and whether to announce them by mail.

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
**id** | **int32** | The folder whose sharing is being changed. A folder stored on the portal is numbered, while a folder in a  connected third-party account is named by an opaque string. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFolderSecurityInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **securityInfoSimpleRequestDto** | [**SecurityInfoSimpleRequestDto**](SecurityInfoSimpleRequestDto.md) | The rights to apply to the folder, and whether to announce them by mail. | 

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

Set sharing rights in batch



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

