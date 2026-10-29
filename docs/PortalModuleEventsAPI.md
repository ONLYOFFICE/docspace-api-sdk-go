# \PortalModuleEventsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendModuleEvent**](PortalModuleEventsAPI.md#SendModuleEvent) | **Post** /api/2.0/portal/moduleevents | Send a module event



## SendModuleEvent

> SendModuleEvent(ctx).XModuleKey(xModuleKey).SendModuleEventRequest(sendModuleEventRequest).Execute()

Send a module event



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/send-module-event/).

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
	xModuleKey := "mk_Q2hhdE1vZHVsZUtleTIwMjYtMTAtMjM" // string | The secret key the installation administrator issued to the module. It is not a user or portal credential:  the module the event belongs to is the one this key is configured for, so the key is the module's identity. (optional)
	sendModuleEventRequest := *openapiclient.NewSendModuleEventRequest() // SendModuleEventRequest | The event to deliver. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PortalModuleEventsAPI.SendModuleEvent(context.Background()).XModuleKey(xModuleKey).SendModuleEventRequest(sendModuleEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalModuleEventsAPI.SendModuleEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendModuleEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xModuleKey** | **string** | The secret key the installation administrator issued to the module. It is not a user or portal credential:  the module the event belongs to is the one this key is configured for, so the key is the module's identity. | 
 **sendModuleEventRequest** | [**SendModuleEventRequest**](SendModuleEventRequest.md) | The event to deliver. | 

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

