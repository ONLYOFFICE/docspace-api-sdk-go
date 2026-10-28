# \SecurityFirebaseAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DocRegisterPusnNotificationDevice**](SecurityFirebaseAPI.md#DocRegisterPusnNotificationDevice) | **Post** /api/2.0/settings/push/docregisterdevice | Register a push device
[**SubscribeDocumentsPushNotification**](SecurityFirebaseAPI.md#SubscribeDocumentsPushNotification) | **Put** /api/2.0/settings/push/docsubscribe | Set push subscription



## DocRegisterPusnNotificationDevice

> FirebaseDeviceWrapper DocRegisterPusnNotificationDevice(ctx).FirebaseRequestDto(firebaseRequestDto).Execute()

Register a push device



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
	firebaseRequestDto := *openapiclient.NewFirebaseRequestDto() // FirebaseRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityFirebaseAPI.DocRegisterPusnNotificationDevice(context.Background()).FirebaseRequestDto(firebaseRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityFirebaseAPI.DocRegisterPusnNotificationDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DocRegisterPusnNotificationDevice`: FirebaseDeviceWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityFirebaseAPI.DocRegisterPusnNotificationDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocRegisterPusnNotificationDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firebaseRequestDto** | [**FirebaseRequestDto**](FirebaseRequestDto.md) |  | 

### Return type

[**FirebaseDeviceWrapper**](FirebaseDeviceWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribeDocumentsPushNotification

> FirebaseDeviceWrapper SubscribeDocumentsPushNotification(ctx).FirebaseRequestDto(firebaseRequestDto).Execute()

Set push subscription



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
	firebaseRequestDto := *openapiclient.NewFirebaseRequestDto() // FirebaseRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityFirebaseAPI.SubscribeDocumentsPushNotification(context.Background()).FirebaseRequestDto(firebaseRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityFirebaseAPI.SubscribeDocumentsPushNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscribeDocumentsPushNotification`: FirebaseDeviceWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityFirebaseAPI.SubscribeDocumentsPushNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeDocumentsPushNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firebaseRequestDto** | [**FirebaseRequestDto**](FirebaseRequestDto.md) |  | 

### Return type

[**FirebaseDeviceWrapper**](FirebaseDeviceWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

