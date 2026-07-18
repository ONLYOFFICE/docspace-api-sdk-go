# \PortalPaymentAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculateWalletPayment**](PortalPaymentAPI.md#CalculateWalletPayment) | **Put** /api/2.0/portal/payment/calculatewallet | Calculate the wallet payment amount
[**ChangeTenantWalletServiceState**](PortalPaymentAPI.md#ChangeTenantWalletServiceState) | **Post** /api/2.0/portal/payment/servicestate | Change tenant wallet service state
[**CreateCustomerOperationsReport**](PortalPaymentAPI.md#CreateCustomerOperationsReport) | **Post** /api/2.0/portal/payment/customer/operationsreport | Start the customer operations report generation
[**CreditAiBalance**](PortalPaymentAPI.md#CreditAiBalance) | **Post** /api/2.0/portal/payment/creditaibalance | Credit AI balance
[**GetAiPrices**](PortalPaymentAPI.md#GetAiPrices) | **Get** /api/2.0/portal/payment/ai-prices | Get AI model prices
[**GetCheckoutSetupUrl**](PortalPaymentAPI.md#GetCheckoutSetupUrl) | **Get** /api/2.0/portal/payment/checkoutsetupurl | Get the checkout setup page URL
[**GetCustomerAiBalance**](PortalPaymentAPI.md#GetCustomerAiBalance) | **Get** /api/2.0/portal/payment/customer/aibalance | Get the customer AI balance
[**GetCustomerBalance**](PortalPaymentAPI.md#GetCustomerBalance) | **Get** /api/2.0/portal/payment/customer/balance | Get the customer balance
[**GetCustomerInfo**](PortalPaymentAPI.md#GetCustomerInfo) | **Get** /api/2.0/portal/payment/customerinfo | Get the customer information
[**GetCustomerOperations**](PortalPaymentAPI.md#GetCustomerOperations) | **Get** /api/2.0/portal/payment/customer/operations | Get the customer operations
[**GetCustomerOperationsReport**](PortalPaymentAPI.md#GetCustomerOperationsReport) | **Get** /api/2.0/portal/payment/customer/operationsreport | Get the status of the customer operations report generation
[**GetPaymentAccount**](PortalPaymentAPI.md#GetPaymentAccount) | **Get** /api/2.0/portal/payment/account | Get the payment account
[**GetPaymentCurrencies**](PortalPaymentAPI.md#GetPaymentCurrencies) | **Get** /api/2.0/portal/payment/currencies | Get currencies
[**GetPaymentQuotas**](PortalPaymentAPI.md#GetPaymentQuotas) | **Get** /api/2.0/portal/payment/quotas | Get quotas
[**GetPaymentUrl**](PortalPaymentAPI.md#GetPaymentUrl) | **Put** /api/2.0/portal/payment/url | Get the payment page URL
[**GetPortalPrices**](PortalPaymentAPI.md#GetPortalPrices) | **Get** /api/2.0/portal/payment/prices | Get prices
[**GetQuotaPaymentInformation**](PortalPaymentAPI.md#GetQuotaPaymentInformation) | **Get** /api/2.0/portal/payment/quota | Get quota payment information
[**GetRestrictedAiModels**](PortalPaymentAPI.md#GetRestrictedAiModels) | **Get** /api/2.0/portal/payment/ai-model/restrictions | Get restricted AI models
[**GetTenantWalletServiceSettings**](PortalPaymentAPI.md#GetTenantWalletServiceSettings) | **Get** /api/2.0/portal/payment/servicessettings | Gets the wallet service settings for the tenant.
[**GetTenantWalletSettings**](PortalPaymentAPI.md#GetTenantWalletSettings) | **Get** /api/2.0/portal/payment/topupsettings | Gets the tenant wallet auto top up settings
[**GetWalletService**](PortalPaymentAPI.md#GetWalletService) | **Get** /api/2.0/portal/payment/walletservice | Get wallet service
[**GetWalletServices**](PortalPaymentAPI.md#GetWalletServices) | **Get** /api/2.0/portal/payment/walletservices | Get wallet services
[**SendPaymentRequest**](PortalPaymentAPI.md#SendPaymentRequest) | **Post** /api/2.0/portal/payment/request | Send a payment request
[**SetRestrictedAiModels**](PortalPaymentAPI.md#SetRestrictedAiModels) | **Put** /api/2.0/portal/payment/ai-model/restrictions | Set restricted AI models
[**SetTenantWalletSettings**](PortalPaymentAPI.md#SetTenantWalletSettings) | **Post** /api/2.0/portal/payment/topupsettings | Set the wallet auto top up settings
[**TerminateCustomerOperationsReport**](PortalPaymentAPI.md#TerminateCustomerOperationsReport) | **Delete** /api/2.0/portal/payment/customer/operationsreport | Terminate the customer operations report generation
[**TopUpDeposit**](PortalPaymentAPI.md#TopUpDeposit) | **Post** /api/2.0/portal/payment/deposit | Put money on deposit
[**UpdatePayment**](PortalPaymentAPI.md#UpdatePayment) | **Put** /api/2.0/portal/payment/update | Update the payment quantity
[**UpdateWalletPayment**](PortalPaymentAPI.md#UpdateWalletPayment) | **Put** /api/2.0/portal/payment/updatewallet | Update the wallet payment quantity



## CalculateWalletPayment

> PaymentCalculationWrapper CalculateWalletPayment(ctx).WalletQuantityRequestDto(walletQuantityRequestDto).Execute()

Calculate the wallet payment amount



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/calculate-wallet-payment/).

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
	walletQuantityRequestDto := *openapiclient.NewWalletQuantityRequestDto() // WalletQuantityRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.CalculateWalletPayment(context.Background()).WalletQuantityRequestDto(walletQuantityRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.CalculateWalletPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalculateWalletPayment`: PaymentCalculationWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.CalculateWalletPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCalculateWalletPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletQuantityRequestDto** | [**WalletQuantityRequestDto**](WalletQuantityRequestDto.md) |  | 

### Return type

[**PaymentCalculationWrapper**](PaymentCalculationWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChangeTenantWalletServiceState

> TenantWalletServiceSettingsWrapper ChangeTenantWalletServiceState(ctx).ChangeWalletServiceStateRequestDto(changeWalletServiceStateRequestDto).Execute()

Change tenant wallet service state



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/change-tenant-wallet-service-state/).

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
	changeWalletServiceStateRequestDto := *openapiclient.NewChangeWalletServiceStateRequestDto() // ChangeWalletServiceStateRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.ChangeTenantWalletServiceState(context.Background()).ChangeWalletServiceStateRequestDto(changeWalletServiceStateRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.ChangeTenantWalletServiceState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeTenantWalletServiceState`: TenantWalletServiceSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.ChangeTenantWalletServiceState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeTenantWalletServiceStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changeWalletServiceStateRequestDto** | [**ChangeWalletServiceStateRequestDto**](ChangeWalletServiceStateRequestDto.md) |  | 

### Return type

[**TenantWalletServiceSettingsWrapper**](TenantWalletServiceSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomerOperationsReport

> DocumentBuilderTaskWrapper CreateCustomerOperationsReport(ctx).CustomerOperationsReportRequestDto(customerOperationsReportRequestDto).Execute()

Start the customer operations report generation



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-customer-operations-report/).

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
	customerOperationsReportRequestDto := *openapiclient.NewCustomerOperationsReportRequestDto() // CustomerOperationsReportRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.CreateCustomerOperationsReport(context.Background()).CustomerOperationsReportRequestDto(customerOperationsReportRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.CreateCustomerOperationsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomerOperationsReport`: DocumentBuilderTaskWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.CreateCustomerOperationsReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerOperationsReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerOperationsReportRequestDto** | [**CustomerOperationsReportRequestDto**](CustomerOperationsReportRequestDto.md) |  | 

### Return type

[**DocumentBuilderTaskWrapper**](DocumentBuilderTaskWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditAiBalance

> ServicePaymentWrapper CreditAiBalance(ctx).CreditAiBalanceRequestDto(creditAiBalanceRequestDto).Execute()

Credit AI balance



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/credit-ai-balance/).

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
	creditAiBalanceRequestDto := *openapiclient.NewCreditAiBalanceRequestDto() // CreditAiBalanceRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.CreditAiBalance(context.Background()).CreditAiBalanceRequestDto(creditAiBalanceRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.CreditAiBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditAiBalance`: ServicePaymentWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.CreditAiBalance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditAiBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **creditAiBalanceRequestDto** | [**CreditAiBalanceRequestDto**](CreditAiBalanceRequestDto.md) |  | 

### Return type

[**ServicePaymentWrapper**](ServicePaymentWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiPrices

> AiPricesResponseWrapper GetAiPrices(ctx).Execute()

Get AI model prices



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-ai-prices/).

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
	resp, r, err := apiClient.PortalPaymentAPI.GetAiPrices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetAiPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAiPrices`: AiPricesResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetAiPrices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiPricesRequest struct via the builder pattern


### Return type

[**AiPricesResponseWrapper**](AiPricesResponseWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCheckoutSetupUrl

> StringWrapper GetCheckoutSetupUrl(ctx).BackUrl(backUrl).SuccessUrl(successUrl).Execute()

Get the checkout setup page URL



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-checkout-setup-url/).

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
	backUrl := "https://example.com/payment/back" // string | The URL where the user will be redirected after setup cancellation.
	successUrl := "https://example.com/payment/success" // string | The URL where the user will be redirected after successful payment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.GetCheckoutSetupUrl(context.Background()).BackUrl(backUrl).SuccessUrl(successUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetCheckoutSetupUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCheckoutSetupUrl`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetCheckoutSetupUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckoutSetupUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backUrl** | **string** | The URL where the user will be redirected after setup cancellation. | 
 **successUrl** | **string** | The URL where the user will be redirected after successful payment. | 

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


## GetCustomerAiBalance

> BalanceWrapper GetCustomerAiBalance(ctx).Refresh(refresh).Execute()

Get the customer AI balance



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-customer-ai-balance/).

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
	refresh := true // bool | Specifies whether to refresh the payment information cache or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.GetCustomerAiBalance(context.Background()).Refresh(refresh).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetCustomerAiBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerAiBalance`: BalanceWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetCustomerAiBalance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerAiBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refresh** | **bool** | Specifies whether to refresh the payment information cache or not. | 

### Return type

[**BalanceWrapper**](BalanceWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerBalance

> BalanceWrapper GetCustomerBalance(ctx).Refresh(refresh).Execute()

Get the customer balance



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-customer-balance/).

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
	refresh := true // bool | Specifies whether to refresh the payment information cache or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.GetCustomerBalance(context.Background()).Refresh(refresh).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetCustomerBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerBalance`: BalanceWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetCustomerBalance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refresh** | **bool** | Specifies whether to refresh the payment information cache or not. | 

### Return type

[**BalanceWrapper**](BalanceWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerInfo

> CustomerInfoWrapper GetCustomerInfo(ctx).Refresh(refresh).Execute()

Get the customer information



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-customer-info/).

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
	refresh := true // bool | Specifies whether to refresh the payment information cache or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.GetCustomerInfo(context.Background()).Refresh(refresh).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetCustomerInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerInfo`: CustomerInfoWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetCustomerInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refresh** | **bool** | Specifies whether to refresh the payment information cache or not. | 

### Return type

[**CustomerInfoWrapper**](CustomerInfoWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerOperations

> ReportWrapper GetCustomerOperations(ctx).Offset(offset).Limit(limit).ServiceName(serviceName).StartDate(startDate).EndDate(endDate).ParticipantName(participantName).Credit(credit).Debit(debit).Type_(type_).Status(status).OrderBy(orderBy).OrderType(orderType).Execute()

Get the customer operations



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-customer-operations/).

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
	offset := int32(0) // int32 | The number of items to skip for pagination. The default value is 0. (optional)
	limit := int32(25) // int32 | The maximum number of items to return for pagination. The default value is 25. (optional)
	serviceName := "backup" // string | The service name. (optional)
	startDate := time.Now() // time.Time | The report start date. (optional)
	endDate := time.Now() // time.Time | The report end date. (optional)
	participantName := "My Own Corporation" // string | The participant name. (optional)
	credit := true // bool | Specifies whether to include credit operations in the report. (optional)
	debit := false // bool | Specifies whether to include debit operations in the report. (optional)
	type_ := openapiclient.OperationType(0) // OperationType | The operation type to filter by. (optional)
	status := openapiclient.OperationStatus(0) // OperationStatus | The operation status to filter by. (optional)
	orderBy := "StartDate" // string | The field to order by. (optional)
	orderType := openapiclient.OperationOrderType(0) // OperationOrderType | Order direction: Ascending or Descending. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.GetCustomerOperations(context.Background()).Offset(offset).Limit(limit).ServiceName(serviceName).StartDate(startDate).EndDate(endDate).ParticipantName(participantName).Credit(credit).Debit(debit).Type_(type_).Status(status).OrderBy(orderBy).OrderType(orderType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetCustomerOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerOperations`: ReportWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetCustomerOperations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of items to skip for pagination. The default value is 0. | 
 **limit** | **int32** | The maximum number of items to return for pagination. The default value is 25. | 
 **serviceName** | **string** | The service name. | 
 **startDate** | **time.Time** | The report start date. | 
 **endDate** | **time.Time** | The report end date. | 
 **participantName** | **string** | The participant name. | 
 **credit** | **bool** | Specifies whether to include credit operations in the report. | 
 **debit** | **bool** | Specifies whether to include debit operations in the report. | 
 **type_** | [**OperationType**](OperationType.md) | The operation type to filter by. | 
 **status** | [**OperationStatus**](OperationStatus.md) | The operation status to filter by. | 
 **orderBy** | **string** | The field to order by. | 
 **orderType** | [**OperationOrderType**](OperationOrderType.md) | Order direction: Ascending or Descending. | 

### Return type

[**ReportWrapper**](ReportWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerOperationsReport

> DocumentBuilderTaskWrapper GetCustomerOperationsReport(ctx).Execute()

Get the status of the customer operations report generation



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-customer-operations-report/).

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
	resp, r, err := apiClient.PortalPaymentAPI.GetCustomerOperationsReport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetCustomerOperationsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerOperationsReport`: DocumentBuilderTaskWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetCustomerOperationsReport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerOperationsReportRequest struct via the builder pattern


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


## GetPaymentAccount

> StringWrapper GetPaymentAccount(ctx).BackUrl(backUrl).Execute()

Get the payment account



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-payment-account/).

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
	backUrl := "https://example.com" // string | The URL where the user will be redirected after payment processing. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.GetPaymentAccount(context.Background()).BackUrl(backUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetPaymentAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentAccount`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetPaymentAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backUrl** | **string** | The URL where the user will be redirected after payment processing. | 

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


## GetPaymentCurrencies

> CurrenciesArrayWrapper GetPaymentCurrencies(ctx).Execute()

Get currencies



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-payment-currencies/).

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
	resp, r, err := apiClient.PortalPaymentAPI.GetPaymentCurrencies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetPaymentCurrencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentCurrencies`: CurrenciesArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetPaymentCurrencies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentCurrenciesRequest struct via the builder pattern


### Return type

[**CurrenciesArrayWrapper**](CurrenciesArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentQuotas

> QuotaArrayWrapper GetPaymentQuotas(ctx).Wallet(wallet).Execute()

Get quotas



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-payment-quotas/).

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
	wallet := true // bool | Specifies whether to return the wallet quotas only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.GetPaymentQuotas(context.Background()).Wallet(wallet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetPaymentQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentQuotas`: QuotaArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetPaymentQuotas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wallet** | **bool** | Specifies whether to return the wallet quotas only. | 

### Return type

[**QuotaArrayWrapper**](QuotaArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentUrl

> StringWrapper GetPaymentUrl(ctx).PaymentUrlRequestDto(paymentUrlRequestDto).Execute()

Get the payment page URL



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-payment-url/).

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
	paymentUrlRequestDto := *openapiclient.NewPaymentUrlRequestDto("https://example.com/payment/back", "https://example.com/payment/success") // PaymentUrlRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.GetPaymentUrl(context.Background()).PaymentUrlRequestDto(paymentUrlRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetPaymentUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentUrl`: StringWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetPaymentUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentUrlRequestDto** | [**PaymentUrlRequestDto**](PaymentUrlRequestDto.md) |  | 

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


## GetPortalPrices

> GetPortalPrices200Response GetPortalPrices(ctx).Execute()

Get prices



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-portal-prices/).

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
	resp, r, err := apiClient.PortalPaymentAPI.GetPortalPrices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetPortalPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortalPrices`: GetPortalPrices200Response
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetPortalPrices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortalPricesRequest struct via the builder pattern


### Return type

[**GetPortalPrices200Response**](GetPortalPrices200Response.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuotaPaymentInformation

> QuotaWrapper GetQuotaPaymentInformation(ctx).Refresh(refresh).Execute()

Get quota payment information



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-quota-payment-information/).

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
	refresh := true // bool | Specifies whether to refresh the payment information cache or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.GetQuotaPaymentInformation(context.Background()).Refresh(refresh).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetQuotaPaymentInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuotaPaymentInformation`: QuotaWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetQuotaPaymentInformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuotaPaymentInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refresh** | **bool** | Specifies whether to refresh the payment information cache or not. | 

### Return type

[**QuotaWrapper**](QuotaWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRestrictedAiModels

> RestrictedModelsResponseWrapper GetRestrictedAiModels(ctx).Execute()

Get restricted AI models



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-restricted-ai-models/).

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
	resp, r, err := apiClient.PortalPaymentAPI.GetRestrictedAiModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetRestrictedAiModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRestrictedAiModels`: RestrictedModelsResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetRestrictedAiModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRestrictedAiModelsRequest struct via the builder pattern


### Return type

[**RestrictedModelsResponseWrapper**](RestrictedModelsResponseWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantWalletServiceSettings

> TenantWalletServiceSettingsWrapper GetTenantWalletServiceSettings(ctx).Execute()

Gets the wallet service settings for the tenant.



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-tenant-wallet-service-settings/).

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
	resp, r, err := apiClient.PortalPaymentAPI.GetTenantWalletServiceSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetTenantWalletServiceSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantWalletServiceSettings`: TenantWalletServiceSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetTenantWalletServiceSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantWalletServiceSettingsRequest struct via the builder pattern


### Return type

[**TenantWalletServiceSettingsWrapper**](TenantWalletServiceSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantWalletSettings

> TenantWalletSettingsWrapper GetTenantWalletSettings(ctx).Execute()

Gets the tenant wallet auto top up settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-tenant-wallet-settings/).

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
	resp, r, err := apiClient.PortalPaymentAPI.GetTenantWalletSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetTenantWalletSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantWalletSettings`: TenantWalletSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetTenantWalletSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantWalletSettingsRequest struct via the builder pattern


### Return type

[**TenantWalletSettingsWrapper**](TenantWalletSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletService

> WalletServiceWrapper GetWalletService(ctx).Service(service).Execute()

Get wallet service



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-wallet-service/).

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
	service := openapiclient.TenantWalletService(-13) // TenantWalletService | The wallet service type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.GetWalletService(context.Background()).Service(service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetWalletService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletService`: WalletServiceWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetWalletService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | [**TenantWalletService**](TenantWalletService.md) | The wallet service type. | 

### Return type

[**WalletServiceWrapper**](WalletServiceWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletServices

> WalletServiceArrayWrapper GetWalletServices(ctx).Execute()

Get wallet services



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-wallet-services/).

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
	resp, r, err := apiClient.PortalPaymentAPI.GetWalletServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetWalletServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletServices`: WalletServiceArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetWalletServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletServicesRequest struct via the builder pattern


### Return type

[**WalletServiceArrayWrapper**](WalletServiceArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendPaymentRequest

> SendPaymentRequest(ctx).SalesRequestsDto(salesRequestsDto).Execute()

Send a payment request



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/send-payment-request/).

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
	salesRequestsDto := *openapiclient.NewSalesRequestsDto("user@example.com", "I would like to inquire about pricing") // SalesRequestsDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PortalPaymentAPI.SendPaymentRequest(context.Background()).SalesRequestsDto(salesRequestsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.SendPaymentRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendPaymentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **salesRequestsDto** | [**SalesRequestsDto**](SalesRequestsDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRestrictedAiModels

> RestrictedModelsResponseWrapper SetRestrictedAiModels(ctx).SetRestrictedAiModelsRequestDto(setRestrictedAiModelsRequestDto).Execute()

Set restricted AI models



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-restricted-ai-models/).

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
	setRestrictedAiModelsRequestDto := *openapiclient.NewSetRestrictedAiModelsRequestDto([]string{"Models_example"}) // SetRestrictedAiModelsRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.SetRestrictedAiModels(context.Background()).SetRestrictedAiModelsRequestDto(setRestrictedAiModelsRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.SetRestrictedAiModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetRestrictedAiModels`: RestrictedModelsResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.SetRestrictedAiModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetRestrictedAiModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setRestrictedAiModelsRequestDto** | [**SetRestrictedAiModelsRequestDto**](SetRestrictedAiModelsRequestDto.md) |  | 

### Return type

[**RestrictedModelsResponseWrapper**](RestrictedModelsResponseWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTenantWalletSettings

> TenantWalletSettingsWrapper SetTenantWalletSettings(ctx).TenantWalletSettingsWrapper(tenantWalletSettingsWrapper).Execute()

Set the wallet auto top up settings



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/set-tenant-wallet-settings/).

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
	tenantWalletSettingsWrapper := *openapiclient.NewTenantWalletSettingsWrapper() // TenantWalletSettingsWrapper |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.SetTenantWalletSettings(context.Background()).TenantWalletSettingsWrapper(tenantWalletSettingsWrapper).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.SetTenantWalletSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetTenantWalletSettings`: TenantWalletSettingsWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.SetTenantWalletSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetTenantWalletSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantWalletSettingsWrapper** | [**TenantWalletSettingsWrapper**](TenantWalletSettingsWrapper.md) |  | 

### Return type

[**TenantWalletSettingsWrapper**](TenantWalletSettingsWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateCustomerOperationsReport

> TerminateCustomerOperationsReport(ctx).Execute()

Terminate the customer operations report generation



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/terminate-customer-operations-report/).

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
	r, err := apiClient.PortalPaymentAPI.TerminateCustomerOperationsReport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.TerminateCustomerOperationsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateCustomerOperationsReportRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TopUpDeposit

> BooleanWrapper TopUpDeposit(ctx).TopUpDepositRequestDto(topUpDepositRequestDto).Execute()

Put money on deposit



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/top-up-deposit/).

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
	topUpDepositRequestDto := *openapiclient.NewTopUpDepositRequestDto() // TopUpDepositRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.TopUpDeposit(context.Background()).TopUpDepositRequestDto(topUpDepositRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.TopUpDeposit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TopUpDeposit`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.TopUpDeposit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTopUpDepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topUpDepositRequestDto** | [**TopUpDepositRequestDto**](TopUpDepositRequestDto.md) |  | 

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


## UpdatePayment

> BooleanWrapper UpdatePayment(ctx).QuantityRequestDto(quantityRequestDto).Execute()

Update the payment quantity



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-payment/).

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
	quantityRequestDto := *openapiclient.NewQuantityRequestDto() // QuantityRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.UpdatePayment(context.Background()).QuantityRequestDto(quantityRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.UpdatePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePayment`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.UpdatePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quantityRequestDto** | [**QuantityRequestDto**](QuantityRequestDto.md) |  | 

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


## UpdateWalletPayment

> BooleanWrapper UpdateWalletPayment(ctx).WalletQuantityRequestDto(walletQuantityRequestDto).Execute()

Update the wallet payment quantity



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/update-wallet-payment/).

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
	walletQuantityRequestDto := *openapiclient.NewWalletQuantityRequestDto() // WalletQuantityRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.UpdateWalletPayment(context.Background()).WalletQuantityRequestDto(walletQuantityRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.UpdateWalletPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWalletPayment`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.UpdateWalletPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWalletPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletQuantityRequestDto** | [**WalletQuantityRequestDto**](WalletQuantityRequestDto.md) |  | 

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

