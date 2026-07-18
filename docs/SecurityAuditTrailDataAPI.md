# \SecurityAuditTrailDataAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuditTrailReport**](SecurityAuditTrailDataAPI.md#CreateAuditTrailReport) | **Post** /api/2.0/security/audit/events/report | Generate the audit trail report
[**GetAuditEventsByFilter**](SecurityAuditTrailDataAPI.md#GetAuditEventsByFilter) | **Get** /api/2.0/security/audit/events/filter | Get filtered audit trail data
[**GetAuditSettings**](SecurityAuditTrailDataAPI.md#GetAuditSettings) | **Get** /api/2.0/security/audit/settings/lifetime | Get the audit trail settings
[**GetAuditTrailMappers**](SecurityAuditTrailDataAPI.md#GetAuditTrailMappers) | **Get** /api/2.0/security/audit/mappers | Get audit trail mappers
[**GetAuditTrailTypes**](SecurityAuditTrailDataAPI.md#GetAuditTrailTypes) | **Get** /api/2.0/security/audit/types | Get audit trail types
[**GetLastAuditEvents**](SecurityAuditTrailDataAPI.md#GetLastAuditEvents) | **Get** /api/2.0/security/audit/events/last | Get audit trail data
[**SetAuditSettings**](SecurityAuditTrailDataAPI.md#SetAuditSettings) | **Post** /api/2.0/security/audit/settings/lifetime | Set the audit trail settings



## CreateAuditTrailReport

> StringWrapper CreateAuditTrailReport(ctx).Execute()

Generate the audit trail report



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-audit-trail-report/).

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
	resp, r, err := apiClient.SecurityAuditTrailDataAPI.CreateAuditTrailReport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAuditTrailDataAPI.CreateAuditTrailReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAuditTrailReport`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityAuditTrailDataAPI.CreateAuditTrailReport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuditTrailReportRequest struct via the builder pattern


### Return type

[**StringWrapper**](StringWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditEventsByFilter

> AuditEventArrayWrapper GetAuditEventsByFilter(ctx).UserId(userId).ModuleType(moduleType).ActionType(actionType).Action(action).EntryType(entryType).Target(target).From(from).To(to).Count(count).StartIndex(startIndex).Execute()

Get filtered audit trail data



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-audit-events-by-filter/).

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
	userId := "00000000-0000-0000-0000-000000000001" // string | The ID of the user who triggered the audit event. (optional)
	moduleType := openapiclient.LocationType(0) // LocationType | The location where the audit event occurred. (optional)
	actionType := openapiclient.ActionType(0) // ActionType | The type of action performed in the audit event (e.g., Create, Update, Delete). (optional)
	action := openapiclient.MessageAction(1000) // MessageAction | The specific action that occurred within the audit event. (optional)
	entryType := openapiclient.EntryType(0) // EntryType | The type of audit entry (e.g., Folder, User, File). (optional)
	target := "document.docx" // string | The target object affected by the audit event (e.g., document ID, user account). (optional)
	from := *openapiclient.NewApiDateTime() // ApiDateTime | The starting date and time for filtering audit events. (optional)
	to := *openapiclient.NewApiDateTime() // ApiDateTime | The ending date and time for filtering audit events. (optional)
	count := int32(100) // int32 | The maximum number of audit event records to retrieve. (optional)
	startIndex := int32(0) // int32 | The index of the first audit event record to retrieve in a paged query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAuditTrailDataAPI.GetAuditEventsByFilter(context.Background()).UserId(userId).ModuleType(moduleType).ActionType(actionType).Action(action).EntryType(entryType).Target(target).From(from).To(to).Count(count).StartIndex(startIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAuditTrailDataAPI.GetAuditEventsByFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditEventsByFilter`: AuditEventArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityAuditTrailDataAPI.GetAuditEventsByFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditEventsByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | The ID of the user who triggered the audit event. | 
 **moduleType** | [**LocationType**](LocationType.md) | The location where the audit event occurred. | 
 **actionType** | [**ActionType**](ActionType.md) | The type of action performed in the audit event (e.g., Create, Update, Delete). | 
 **action** | [**MessageAction**](MessageAction.md) | The specific action that occurred within the audit event. | 
 **entryType** | [**EntryType**](EntryType.md) | The type of audit entry (e.g., Folder, User, File). | 
 **target** | **string** | The target object affected by the audit event (e.g., document ID, user account). | 
 **from** | [**ApiDateTime**](ApiDateTime.md) | The starting date and time for filtering audit events. | 
 **to** | [**ApiDateTime**](ApiDateTime.md) | The ending date and time for filtering audit events. | 
 **count** | **int32** | The maximum number of audit event records to retrieve. | 
 **startIndex** | **int32** | The index of the first audit event record to retrieve in a paged query. | 

### Return type

[**AuditEventArrayWrapper**](AuditEventArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditSettings

> TenantAuditSettingsWrapper GetAuditSettings(ctx).Execute()

Get the audit trail settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-audit-settings/).

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
	resp, r, err := apiClient.SecurityAuditTrailDataAPI.GetAuditSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAuditTrailDataAPI.GetAuditSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditSettings`: TenantAuditSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityAuditTrailDataAPI.GetAuditSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditSettingsRequest struct via the builder pattern


### Return type

[**TenantAuditSettingsWrapper**](TenantAuditSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditTrailMappers

> ObjectWrapper GetAuditTrailMappers(ctx).ProductType(productType).ModuleType(moduleType).Execute()

Get audit trail mappers



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-audit-trail-mappers/).

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
	productType := openapiclient.ProductType(2) // ProductType | The type of product related to the audit trail. (optional)
	moduleType := openapiclient.LocationType(0) // LocationType | The location associated with the audit trail. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAuditTrailDataAPI.GetAuditTrailMappers(context.Background()).ProductType(productType).ModuleType(moduleType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAuditTrailDataAPI.GetAuditTrailMappers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditTrailMappers`: ObjectWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityAuditTrailDataAPI.GetAuditTrailMappers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditTrailMappersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productType** | [**ProductType**](ProductType.md) | The type of product related to the audit trail. | 
 **moduleType** | [**LocationType**](LocationType.md) | The location associated with the audit trail. | 

### Return type

[**ObjectWrapper**](ObjectWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditTrailTypes

> ObjectWrapper GetAuditTrailTypes(ctx).Execute()

Get audit trail types



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-audit-trail-types/).

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
	resp, r, err := apiClient.SecurityAuditTrailDataAPI.GetAuditTrailTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAuditTrailDataAPI.GetAuditTrailTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditTrailTypes`: ObjectWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityAuditTrailDataAPI.GetAuditTrailTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditTrailTypesRequest struct via the builder pattern


### Return type

[**ObjectWrapper**](ObjectWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastAuditEvents

> AuditEventArrayWrapper GetLastAuditEvents(ctx).Execute()

Get audit trail data



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-last-audit-events/).

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
	resp, r, err := apiClient.SecurityAuditTrailDataAPI.GetLastAuditEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAuditTrailDataAPI.GetLastAuditEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLastAuditEvents`: AuditEventArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityAuditTrailDataAPI.GetLastAuditEvents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLastAuditEventsRequest struct via the builder pattern


### Return type

[**AuditEventArrayWrapper**](AuditEventArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAuditSettings

> TenantAuditSettingsWrapper SetAuditSettings(ctx).TenantAuditSettingsWrapper(tenantAuditSettingsWrapper).Execute()

Set the audit trail settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-audit-settings/).

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
	tenantAuditSettingsWrapper := *openapiclient.NewTenantAuditSettingsWrapper() // TenantAuditSettingsWrapper |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAuditTrailDataAPI.SetAuditSettings(context.Background()).TenantAuditSettingsWrapper(tenantAuditSettingsWrapper).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAuditTrailDataAPI.SetAuditSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAuditSettings`: TenantAuditSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `SecurityAuditTrailDataAPI.SetAuditSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetAuditSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantAuditSettingsWrapper** | [**TenantAuditSettingsWrapper**](TenantAuditSettingsWrapper.md) |  | 

### Return type

[**TenantAuditSettingsWrapper**](TenantAuditSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

