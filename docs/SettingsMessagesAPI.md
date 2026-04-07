# \SettingsMessagesAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnableAdminMessageSettings**](SettingsMessagesAPI.md#EnableAdminMessageSettings) | **Post** /api/2.0/settings/messagesettings | Enable the administrator message settings
[**SendAdminMail**](SettingsMessagesAPI.md#SendAdminMail) | **Post** /api/2.0/settings/sendadmmail | Send a message to the administrator
[**SendJoinInviteMail**](SettingsMessagesAPI.md#SendJoinInviteMail) | **Post** /api/2.0/settings/sendjoininvite | Sends an invitation email



## EnableAdminMessageSettings

> StringWrapper EnableAdminMessageSettings(ctx).TurnOnAdminMessageSettingsRequestDto(turnOnAdminMessageSettingsRequestDto).Execute()

Enable the administrator message settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/enable-admin-message-settings/).

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
	turnOnAdminMessageSettingsRequestDto := *openapiclient.NewTurnOnAdminMessageSettingsRequestDto() // TurnOnAdminMessageSettingsRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsMessagesAPI.EnableAdminMessageSettings(context.Background()).TurnOnAdminMessageSettingsRequestDto(turnOnAdminMessageSettingsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsMessagesAPI.EnableAdminMessageSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableAdminMessageSettings`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsMessagesAPI.EnableAdminMessageSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableAdminMessageSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **turnOnAdminMessageSettingsRequestDto** | [**TurnOnAdminMessageSettingsRequestDto**](TurnOnAdminMessageSettingsRequestDto.md) |  | 

### Return type

[**StringWrapper**](StringWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendAdminMail

> StringWrapper SendAdminMail(ctx).AdminMessageSettingsRequestsDto(adminMessageSettingsRequestsDto).Execute()

Send a message to the administrator



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/send-admin-mail/).

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
	adminMessageSettingsRequestsDto := *openapiclient.NewAdminMessageSettingsRequestsDto("Hello, this is a test message from the administrator.", "user@example.com") // AdminMessageSettingsRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsMessagesAPI.SendAdminMail(context.Background()).AdminMessageSettingsRequestsDto(adminMessageSettingsRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsMessagesAPI.SendAdminMail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendAdminMail`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsMessagesAPI.SendAdminMail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendAdminMailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminMessageSettingsRequestsDto** | [**AdminMessageSettingsRequestsDto**](AdminMessageSettingsRequestsDto.md) |  | 

### Return type

[**StringWrapper**](StringWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendJoinInviteMail

> StringWrapper SendJoinInviteMail(ctx).AdminMessageBaseSettingsRequestsDto(adminMessageBaseSettingsRequestsDto).Execute()

Sends an invitation email



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/send-join-invite-mail/).

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
	adminMessageBaseSettingsRequestsDto := *openapiclient.NewAdminMessageBaseSettingsRequestsDto("admin@example.com") // AdminMessageBaseSettingsRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsMessagesAPI.SendJoinInviteMail(context.Background()).AdminMessageBaseSettingsRequestsDto(adminMessageBaseSettingsRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsMessagesAPI.SendJoinInviteMail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendJoinInviteMail`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsMessagesAPI.SendJoinInviteMail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendJoinInviteMailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminMessageBaseSettingsRequestsDto** | [**AdminMessageBaseSettingsRequestsDto**](AdminMessageBaseSettingsRequestsDto.md) |  | 

### Return type

[**StringWrapper**](StringWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

