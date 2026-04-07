# \SecurityFirebaseAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DocRegisterPusnNotificationDevice**](SecurityFirebaseAPI.md#DocRegisterPusnNotificationDevice) | **Post** /api/2.0/settings/push/docregisterdevice | Save the Documents Firebase device token
[**SubscribeDocumentsPushNotification**](SecurityFirebaseAPI.md#SubscribeDocumentsPushNotification) | **Put** /api/2.0/settings/push/docsubscribe | Subscribe to Documents push notification



## DocRegisterPusnNotificationDevice

> FireBaseUserWrapper DocRegisterPusnNotificationDevice(ctx).FirebaseRequestsDto(firebaseRequestsDto).Execute()

Save the Documents Firebase device token



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/doc-register-pusn-notification-device/).

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
	firebaseRequestsDto := *openapiclient.NewFirebaseRequestsDto() // FirebaseRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityFirebaseAPI.DocRegisterPusnNotificationDevice(context.Background()).FirebaseRequestsDto(firebaseRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityFirebaseAPI.DocRegisterPusnNotificationDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocRegisterPusnNotificationDevice`: FireBaseUserWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityFirebaseAPI.DocRegisterPusnNotificationDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocRegisterPusnNotificationDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firebaseRequestsDto** | [**FirebaseRequestsDto**](FirebaseRequestsDto.md) |  | 

### Return type

[**FireBaseUserWrapper**](FireBaseUserWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribeDocumentsPushNotification

> FireBaseUserWrapper SubscribeDocumentsPushNotification(ctx).FirebaseRequestsDto(firebaseRequestsDto).Execute()

Subscribe to Documents push notification



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/subscribe-documents-push-notification/).

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
	firebaseRequestsDto := *openapiclient.NewFirebaseRequestsDto() // FirebaseRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityFirebaseAPI.SubscribeDocumentsPushNotification(context.Background()).FirebaseRequestsDto(firebaseRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityFirebaseAPI.SubscribeDocumentsPushNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscribeDocumentsPushNotification`: FireBaseUserWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityFirebaseAPI.SubscribeDocumentsPushNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeDocumentsPushNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firebaseRequestsDto** | [**FirebaseRequestsDto**](FirebaseRequestsDto.md) |  | 

### Return type

[**FireBaseUserWrapper**](FireBaseUserWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

