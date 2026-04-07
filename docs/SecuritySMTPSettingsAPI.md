# \SecuritySMTPSettingsAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSmtpOperationStatus**](SecuritySMTPSettingsAPI.md#GetSmtpOperationStatus) | **Get** /api/2.0/smtpsettings/smtp/test/status | Get the SMTP testing process status
[**GetSmtpSettings**](SecuritySMTPSettingsAPI.md#GetSmtpSettings) | **Get** /api/2.0/smtpsettings/smtp | Get the SMTP settings
[**ResetSmtpSettings**](SecuritySMTPSettingsAPI.md#ResetSmtpSettings) | **Delete** /api/2.0/smtpsettings/smtp | Reset the SMTP settings
[**SaveSmtpSettings**](SecuritySMTPSettingsAPI.md#SaveSmtpSettings) | **Post** /api/2.0/smtpsettings/smtp | Save the SMTP settings
[**TestSmtpSettings**](SecuritySMTPSettingsAPI.md#TestSmtpSettings) | **Get** /api/2.0/smtpsettings/smtp/test | Test the SMTP settings



## GetSmtpOperationStatus

> SmtpOperationStatusRequestsWrapper GetSmtpOperationStatus(ctx).Execute()

Get the SMTP testing process status



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-smtp-operation-status/).

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
	resp, r, err := apiClient.SecuritySMTPSettingsAPI.GetSmtpOperationStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecuritySMTPSettingsAPI.GetSmtpOperationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSmtpOperationStatus`: SmtpOperationStatusRequestsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecuritySMTPSettingsAPI.GetSmtpOperationStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmtpOperationStatusRequest struct via the builder pattern


### Return type

[**SmtpOperationStatusRequestsWrapper**](SmtpOperationStatusRequestsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSmtpSettings

> SmtpSettingsWrapper GetSmtpSettings(ctx).Execute()

Get the SMTP settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-smtp-settings/).

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
	resp, r, err := apiClient.SecuritySMTPSettingsAPI.GetSmtpSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecuritySMTPSettingsAPI.GetSmtpSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSmtpSettings`: SmtpSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecuritySMTPSettingsAPI.GetSmtpSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSmtpSettingsRequest struct via the builder pattern


### Return type

[**SmtpSettingsWrapper**](SmtpSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetSmtpSettings

> SmtpSettingsWrapper ResetSmtpSettings(ctx).Execute()

Reset the SMTP settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/reset-smtp-settings/).

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
	resp, r, err := apiClient.SecuritySMTPSettingsAPI.ResetSmtpSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecuritySMTPSettingsAPI.ResetSmtpSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetSmtpSettings`: SmtpSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecuritySMTPSettingsAPI.ResetSmtpSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiResetSmtpSettingsRequest struct via the builder pattern


### Return type

[**SmtpSettingsWrapper**](SmtpSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveSmtpSettings

> SmtpSettingsWrapper SaveSmtpSettings(ctx).SmtpSettingsDto(smtpSettingsDto).Execute()

Save the SMTP settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/save-smtp-settings/).

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
	smtpSettingsDto := *openapiclient.NewSmtpSettingsDto() // SmtpSettingsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecuritySMTPSettingsAPI.SaveSmtpSettings(context.Background()).SmtpSettingsDto(smtpSettingsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecuritySMTPSettingsAPI.SaveSmtpSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveSmtpSettings`: SmtpSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecuritySMTPSettingsAPI.SaveSmtpSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveSmtpSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smtpSettingsDto** | [**SmtpSettingsDto**](SmtpSettingsDto.md) |  | 

### Return type

[**SmtpSettingsWrapper**](SmtpSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestSmtpSettings

> SmtpOperationStatusRequestsWrapper TestSmtpSettings(ctx).Execute()

Test the SMTP settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/test-smtp-settings/).

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
	resp, r, err := apiClient.SecuritySMTPSettingsAPI.TestSmtpSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecuritySMTPSettingsAPI.TestSmtpSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestSmtpSettings`: SmtpOperationStatusRequestsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecuritySMTPSettingsAPI.TestSmtpSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestSmtpSettingsRequest struct via the builder pattern


### Return type

[**SmtpOperationStatusRequestsWrapper**](SmtpOperationStatusRequestsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

