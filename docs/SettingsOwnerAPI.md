# \SettingsOwnerAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendOwnerChangeInstructions**](SettingsOwnerAPI.md#SendOwnerChangeInstructions) | **Post** /api/2.0/settings/owner | Send the owner change instructions
[**UpdatePortalOwner**](SettingsOwnerAPI.md#UpdatePortalOwner) | **Put** /api/2.0/settings/owner | Update the portal owner



## SendOwnerChangeInstructions

> OwnerChangeInstructionsWrapper SendOwnerChangeInstructions(ctx).OwnerIdSettingsRequestDto(ownerIdSettingsRequestDto).Execute()

Send the owner change instructions



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/send-owner-change-instructions/).

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
	ownerIdSettingsRequestDto := *openapiclient.NewOwnerIdSettingsRequestDto("00000000-0000-0000-0000-000000000001") // OwnerIdSettingsRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsOwnerAPI.SendOwnerChangeInstructions(context.Background()).OwnerIdSettingsRequestDto(ownerIdSettingsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsOwnerAPI.SendOwnerChangeInstructions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendOwnerChangeInstructions`: OwnerChangeInstructionsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsOwnerAPI.SendOwnerChangeInstructions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendOwnerChangeInstructionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerIdSettingsRequestDto** | [**OwnerIdSettingsRequestDto**](OwnerIdSettingsRequestDto.md) |  | 

### Return type

[**OwnerChangeInstructionsWrapper**](OwnerChangeInstructionsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePortalOwner

> UpdatePortalOwner(ctx).OwnerIdSettingsRequestDto(ownerIdSettingsRequestDto).Execute()

Update the portal owner



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-portal-owner/).

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
	ownerIdSettingsRequestDto := *openapiclient.NewOwnerIdSettingsRequestDto("00000000-0000-0000-0000-000000000001") // OwnerIdSettingsRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsOwnerAPI.UpdatePortalOwner(context.Background()).OwnerIdSettingsRequestDto(ownerIdSettingsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsOwnerAPI.UpdatePortalOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePortalOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerIdSettingsRequestDto** | [**OwnerIdSettingsRequestDto**](OwnerIdSettingsRequestDto.md) |  | 

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

