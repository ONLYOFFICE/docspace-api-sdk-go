# \BackupAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelBackup**](BackupAPI.md#CancelBackup) | **Post** /api/2.0/backup/cancelbackup | Cancel current backup
[**CreateBackupSchedule**](BackupAPI.md#CreateBackupSchedule) | **Post** /api/2.0/backup/createbackupschedule | Create the backup schedule
[**DeleteBackup**](BackupAPI.md#DeleteBackup) | **Delete** /api/2.0/backup/deletebackup/{id} | Delete the backup
[**DeleteBackupHistory**](BackupAPI.md#DeleteBackupHistory) | **Delete** /api/2.0/backup/deletebackuphistory | Delete the backup history
[**DeleteBackupSchedule**](BackupAPI.md#DeleteBackupSchedule) | **Delete** /api/2.0/backup/deletebackupschedule | Delete the backup schedule
[**GetBackupHistory**](BackupAPI.md#GetBackupHistory) | **Get** /api/2.0/backup/getbackuphistory | Get the backup history
[**GetBackupProgress**](BackupAPI.md#GetBackupProgress) | **Get** /api/2.0/backup/getbackupprogress | Get the backup progress
[**GetBackupSchedule**](BackupAPI.md#GetBackupSchedule) | **Get** /api/2.0/backup/getbackupschedule | Get the backup schedule
[**GetBackupsCount**](BackupAPI.md#GetBackupsCount) | **Get** /api/2.0/backup/getbackupscount | Get the number of backups
[**GetBackupsCounts**](BackupAPI.md#GetBackupsCounts) | **Get** /api/2.0/backup/getbackupscountbypaid | Get the number of free and paid backups
[**GetBackupsServiceState**](BackupAPI.md#GetBackupsServiceState) | **Get** /api/2.0/backup/getservicestate | Get the backup service state
[**GetRestoreProgress**](BackupAPI.md#GetRestoreProgress) | **Get** /api/2.0/backup/getrestoreprogress | Get the restoring progress
[**StartBackup**](BackupAPI.md#StartBackup) | **Post** /api/2.0/backup/startbackup | Start the backup
[**StartBackupRestore**](BackupAPI.md#StartBackupRestore) | **Post** /api/2.0/backup/startrestore | Start the restoring process



## CancelBackup

> BooleanWrapper CancelBackup(ctx).Execute()

Cancel current backup



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/cancel-backup/).

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
	resp, r, err := apiClient.BackupAPI.CancelBackup(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.CancelBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelBackup`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.CancelBackup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCancelBackupRequest struct via the builder pattern


### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBackupSchedule

> BooleanWrapper CreateBackupSchedule(ctx).BackupScheduleDto(backupScheduleDto).Execute()

Create the backup schedule



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-backup-schedule/).

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
	backupScheduleDto := *openapiclient.NewBackupScheduleDto() // BackupScheduleDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.CreateBackupSchedule(context.Background()).BackupScheduleDto(backupScheduleDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.CreateBackupSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBackupSchedule`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.CreateBackupSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupScheduleDto** | [**BackupScheduleDto**](BackupScheduleDto.md) |  | 

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


## DeleteBackup

> BooleanWrapper DeleteBackup(ctx, id).Execute()

Delete the backup



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-backup/).

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
	id := "00000000-0000-0000-0000-000000000000" // string | The backup ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.DeleteBackup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.DeleteBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBackup`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.DeleteBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The backup ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBackupHistory

> BooleanWrapper DeleteBackupHistory(ctx).Dump(dump).Execute()

Delete the backup history



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-backup-history/).

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
	dump := true // bool | Specifies if a dump will be created or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.DeleteBackupHistory(context.Background()).Dump(dump).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.DeleteBackupHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBackupHistory`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.DeleteBackupHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackupHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dump** | **bool** | Specifies if a dump will be created or not. | 

### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBackupSchedule

> BooleanWrapper DeleteBackupSchedule(ctx).Dump(dump).Execute()

Delete the backup schedule



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/delete-backup-schedule/).

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
	dump := true // bool | Specifies if a dump will be created or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.DeleteBackupSchedule(context.Background()).Dump(dump).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.DeleteBackupSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBackupSchedule`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.DeleteBackupSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dump** | **bool** | Specifies if a dump will be created or not. | 

### Return type

[**BooleanWrapper**](BooleanWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupHistory

> BackupHistoryRecordArrayWrapper GetBackupHistory(ctx).Dump(dump).Execute()

Get the backup history



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-backup-history/).

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
	dump := true // bool | Specifies if a dump will be created or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.GetBackupHistory(context.Background()).Dump(dump).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.GetBackupHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackupHistory`: BackupHistoryRecordArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.GetBackupHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dump** | **bool** | Specifies if a dump will be created or not. | 

### Return type

[**BackupHistoryRecordArrayWrapper**](BackupHistoryRecordArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupProgress

> BackupProgressWrapper GetBackupProgress(ctx).Dump(dump).Execute()

Get the backup progress



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-backup-progress/).

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
	dump := true // bool | Specifies if a dump will be created or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.GetBackupProgress(context.Background()).Dump(dump).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.GetBackupProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackupProgress`: BackupProgressWrapper
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.GetBackupProgress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dump** | **bool** | Specifies if a dump will be created or not. | 

### Return type

[**BackupProgressWrapper**](BackupProgressWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupSchedule

> ScheduleWrapper GetBackupSchedule(ctx).Dump(dump).Execute()

Get the backup schedule



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-backup-schedule/).

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
	dump := true // bool | Specifies if a dump will be created or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.GetBackupSchedule(context.Background()).Dump(dump).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.GetBackupSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackupSchedule`: ScheduleWrapper
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.GetBackupSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dump** | **bool** | Specifies if a dump will be created or not. | 

### Return type

[**ScheduleWrapper**](ScheduleWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupsCount

> Int32Wrapper GetBackupsCount(ctx).From(from).To(to).Paid(paid).Execute()

Get the number of backups



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-backups-count/).

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
	from := time.Now() // time.Time | The from date. (optional)
	to := time.Now() // time.Time | The to date. (optional)
	paid := false // bool | Specifies if the backups are paid or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.GetBackupsCount(context.Background()).From(from).To(to).Paid(paid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.GetBackupsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackupsCount`: Int32Wrapper
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.GetBackupsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **time.Time** | The from date. | 
 **to** | **time.Time** | The to date. | 
 **paid** | **bool** | Specifies if the backups are paid or not. | 

### Return type

[**Int32Wrapper**](Int32Wrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupsCounts

> BackupsCountResultWrapper GetBackupsCounts(ctx).From(from).To(to).Paid(paid).Execute()

Get the number of free and paid backups



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-backups-counts/).

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
	from := time.Now() // time.Time | The from date. (optional)
	to := time.Now() // time.Time | The to date. (optional)
	paid := false // bool | Specifies if the backups are paid or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.GetBackupsCounts(context.Background()).From(from).To(to).Paid(paid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.GetBackupsCounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackupsCounts`: BackupsCountResultWrapper
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.GetBackupsCounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupsCountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **time.Time** | The from date. | 
 **to** | **time.Time** | The to date. | 
 **paid** | **bool** | Specifies if the backups are paid or not. | 

### Return type

[**BackupsCountResultWrapper**](BackupsCountResultWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupsServiceState

> BackupServiceStateWrapper GetBackupsServiceState(ctx).Execute()

Get the backup service state



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-backups-service-state/).

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
	resp, r, err := apiClient.BackupAPI.GetBackupsServiceState(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.GetBackupsServiceState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackupsServiceState`: BackupServiceStateWrapper
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.GetBackupsServiceState`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupsServiceStateRequest struct via the builder pattern


### Return type

[**BackupServiceStateWrapper**](BackupServiceStateWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRestoreProgress

> BackupProgressWrapper GetRestoreProgress(ctx).Dump(dump).Execute()

Get the restoring progress



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-restore-progress/).

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
	dump := false // bool | Specifies if a dump will be created or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.GetRestoreProgress(context.Background()).Dump(dump).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.GetRestoreProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRestoreProgress`: BackupProgressWrapper
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.GetRestoreProgress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRestoreProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dump** | **bool** | Specifies if a dump will be created or not. | 

### Return type

[**BackupProgressWrapper**](BackupProgressWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartBackup

> BackupProgressWrapper StartBackup(ctx).BackupDto(backupDto).Execute()

Start the backup



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/start-backup/).

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
	backupDto := *openapiclient.NewBackupDto() // BackupDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.StartBackup(context.Background()).BackupDto(backupDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.StartBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartBackup`: BackupProgressWrapper
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.StartBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupDto** | [**BackupDto**](BackupDto.md) |  | 

### Return type

[**BackupProgressWrapper**](BackupProgressWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartBackupRestore

> BackupProgressWrapper StartBackupRestore(ctx).BackupRestoreDto(backupRestoreDto).Execute()

Start the restoring process



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/start-backup-restore/).

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
	backupRestoreDto := *openapiclient.NewBackupRestoreDto("00000000-0000-0000-0000-000000000000") // BackupRestoreDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.StartBackupRestore(context.Background()).BackupRestoreDto(backupRestoreDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.StartBackupRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartBackupRestore`: BackupProgressWrapper
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.StartBackupRestore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartBackupRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupRestoreDto** | [**BackupRestoreDto**](BackupRestoreDto.md) |  | 

### Return type

[**BackupProgressWrapper**](BackupProgressWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

