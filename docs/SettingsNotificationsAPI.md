# \SettingsNotificationsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNotificationChannels**](SettingsNotificationsAPI.md#GetNotificationChannels) | **Get** /api/2.0/settings/notification/channels | Get notification channels
[**GetNotificationSettings**](SettingsNotificationsAPI.md#GetNotificationSettings) | **Get** /api/2.0/settings/notification/{type} | Check notification availability
[**GetRoomsNotificationSettings**](SettingsNotificationsAPI.md#GetRoomsNotificationSettings) | **Get** /api/2.0/settings/notification/rooms | Get room notification settings
[**SetNotificationSettings**](SettingsNotificationsAPI.md#SetNotificationSettings) | **Post** /api/2.0/settings/notification | Enable notifications
[**SetRoomsNotificationStatus**](SettingsNotificationsAPI.md#SetRoomsNotificationStatus) | **Post** /api/2.0/settings/notification/rooms | Set room notification status



## GetNotificationChannels

> NotificationChannelStatusWrapper GetNotificationChannels(ctx).Execute()

Get notification channels



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-notification-channels/).

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
	resp, r, err := apiClient.SettingsNotificationsAPI.GetNotificationChannels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsNotificationsAPI.GetNotificationChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationChannels`: NotificationChannelStatusWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsNotificationsAPI.GetNotificationChannels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationChannelsRequest struct via the builder pattern


### Return type

[**NotificationChannelStatusWrapper**](NotificationChannelStatusWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationSettings

> NotificationSettingsWrapper GetNotificationSettings(ctx, type_).Execute()

Check notification availability



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-notification-settings/).

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
	type_ := openapiclient.NotificationType(0) // NotificationType | The type of notification to query, specified in the route.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsNotificationsAPI.GetNotificationSettings(context.Background(), type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsNotificationsAPI.GetNotificationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationSettings`: NotificationSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsNotificationsAPI.GetNotificationSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | [**NotificationType**](.md) | The type of notification to query, specified in the route. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationSettingsWrapper**](NotificationSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoomsNotificationSettings

> RoomsNotificationSettingsWrapper GetRoomsNotificationSettings(ctx).Execute()

Get room notification settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-rooms-notification-settings/).

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
	resp, r, err := apiClient.SettingsNotificationsAPI.GetRoomsNotificationSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsNotificationsAPI.GetRoomsNotificationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoomsNotificationSettings`: RoomsNotificationSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsNotificationsAPI.GetRoomsNotificationSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoomsNotificationSettingsRequest struct via the builder pattern


### Return type

[**RoomsNotificationSettingsWrapper**](RoomsNotificationSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNotificationSettings

> NotificationSettingsWrapper SetNotificationSettings(ctx).NotificationSettingsRequestsDto(notificationSettingsRequestsDto).Execute()

Enable notifications



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-notification-settings/).

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
	notificationSettingsRequestsDto := *openapiclient.NewNotificationSettingsRequestsDto(openapiclient.NotificationType(0)) // NotificationSettingsRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsNotificationsAPI.SetNotificationSettings(context.Background()).NotificationSettingsRequestsDto(notificationSettingsRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsNotificationsAPI.SetNotificationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetNotificationSettings`: NotificationSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsNotificationsAPI.SetNotificationSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetNotificationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationSettingsRequestsDto** | [**NotificationSettingsRequestsDto**](NotificationSettingsRequestsDto.md) |  | 

### Return type

[**NotificationSettingsWrapper**](NotificationSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRoomsNotificationStatus

> RoomsNotificationSettingsWrapper SetRoomsNotificationStatus(ctx).RoomsNotificationsSettingsRequestDto(roomsNotificationsSettingsRequestDto).Execute()

Set room notification status



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-rooms-notification-status/).

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
	roomsNotificationsSettingsRequestDto := *openapiclient.NewRoomsNotificationsSettingsRequestDto() // RoomsNotificationsSettingsRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsNotificationsAPI.SetRoomsNotificationStatus(context.Background()).RoomsNotificationsSettingsRequestDto(roomsNotificationsSettingsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsNotificationsAPI.SetRoomsNotificationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetRoomsNotificationStatus`: RoomsNotificationSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SettingsNotificationsAPI.SetRoomsNotificationStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetRoomsNotificationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roomsNotificationsSettingsRequestDto** | [**RoomsNotificationsSettingsRequestDto**](RoomsNotificationsSettingsRequestDto.md) |  | 

### Return type

[**RoomsNotificationSettingsWrapper**](RoomsNotificationSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

