# \SecurityLoginHistoryAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLoginHistoryReport**](SecurityLoginHistoryAPI.md#CreateLoginHistoryReport) | **Post** /api/2.0/security/audit/login/report | Start the login history report generation
[**GetLastLoginEvents**](SecurityLoginHistoryAPI.md#GetLastLoginEvents) | **Get** /api/2.0/security/audit/login/last | Get login history
[**GetLoginEventsByFilter**](SecurityLoginHistoryAPI.md#GetLoginEventsByFilter) | **Get** /api/2.0/security/audit/login/filter | Get filtered login events
[**GetLoginHistoryReport**](SecurityLoginHistoryAPI.md#GetLoginHistoryReport) | **Get** /api/2.0/security/audit/login/report | Get the login history report generation status
[**TerminateLoginHistoryReport**](SecurityLoginHistoryAPI.md#TerminateLoginHistoryReport) | **Delete** /api/2.0/security/audit/login/report | Terminate the login history report generation



## CreateLoginHistoryReport

> DocumentBuilderTaskWrapper CreateLoginHistoryReport(ctx).Format(format).Execute()

Start the login history report generation



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-login-history-report/).

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
	format := openapiclient.AuditReportFormat(0) // AuditReportFormat | The output file format of the report. Defaults to XLSX. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityLoginHistoryAPI.CreateLoginHistoryReport(context.Background()).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityLoginHistoryAPI.CreateLoginHistoryReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLoginHistoryReport`: DocumentBuilderTaskWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityLoginHistoryAPI.CreateLoginHistoryReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoginHistoryReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | [**AuditReportFormat**](AuditReportFormat.md) | The output file format of the report. Defaults to XLSX. | 

### Return type

[**DocumentBuilderTaskWrapper**](DocumentBuilderTaskWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastLoginEvents

> LoginEventArrayWrapper GetLastLoginEvents(ctx).Execute()

Get login history



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-last-login-events/).

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
	resp, r, err := apiClient.SecurityLoginHistoryAPI.GetLastLoginEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityLoginHistoryAPI.GetLastLoginEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLastLoginEvents`: LoginEventArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityLoginHistoryAPI.GetLastLoginEvents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLastLoginEventsRequest struct via the builder pattern


### Return type

[**LoginEventArrayWrapper**](LoginEventArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoginEventsByFilter

> LoginEventArrayWrapper GetLoginEventsByFilter(ctx).UserId(userId).Action(action).From(from).To(to).Count(count).StartIndex(startIndex).Execute()

Get filtered login events



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-login-events-by-filter/).

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	userId := "00000000-0000-0000-0000-000000000000" // string | The ID of the user whose login events are being queried. (optional)
	action := openapiclient.MessageAction(1000) // MessageAction | The login-related action to filter events by. (optional)
	from := time.Now() // time.Time | The starting date and time for filtering login events. (optional)
	to := time.Now() // time.Time | The ending date and time for filtering login events. (optional)
	count := int32(1) // int32 | The number of login events to retrieve in the query. (optional)
	startIndex := int32(1) // int32 | The starting index for fetching a subset of login events from the query results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityLoginHistoryAPI.GetLoginEventsByFilter(context.Background()).UserId(userId).Action(action).From(from).To(to).Count(count).StartIndex(startIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityLoginHistoryAPI.GetLoginEventsByFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoginEventsByFilter`: LoginEventArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityLoginHistoryAPI.GetLoginEventsByFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoginEventsByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | The ID of the user whose login events are being queried. | 
 **action** | [**MessageAction**](MessageAction.md) | The login-related action to filter events by. | 
 **from** | **time.Time** | The starting date and time for filtering login events. | 
 **to** | **time.Time** | The ending date and time for filtering login events. | 
 **count** | **int32** | The number of login events to retrieve in the query. | 
 **startIndex** | **int32** | The starting index for fetching a subset of login events from the query results. | 

### Return type

[**LoginEventArrayWrapper**](LoginEventArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoginHistoryReport

> DocumentBuilderTaskWrapper GetLoginHistoryReport(ctx).Execute()

Get the login history report generation status



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-login-history-report/).

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
	resp, r, err := apiClient.SecurityLoginHistoryAPI.GetLoginHistoryReport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityLoginHistoryAPI.GetLoginHistoryReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoginHistoryReport`: DocumentBuilderTaskWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityLoginHistoryAPI.GetLoginHistoryReport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoginHistoryReportRequest struct via the builder pattern


### Return type

[**DocumentBuilderTaskWrapper**](DocumentBuilderTaskWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateLoginHistoryReport

> TerminateLoginHistoryReport(ctx).Execute()

Terminate the login history report generation



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/terminate-login-history-report/).

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
	r, err := apiClient.SecurityLoginHistoryAPI.TerminateLoginHistoryReport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityLoginHistoryAPI.TerminateLoginHistoryReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateLoginHistoryReportRequest struct via the builder pattern


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

