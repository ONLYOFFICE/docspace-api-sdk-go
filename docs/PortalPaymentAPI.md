# \PortalPaymentAPI

All URIs are relative to *https://your-docspace.onlyoffice.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculateWalletPayment**](PortalPaymentAPI.md#CalculateWalletPayment) | **Put** /api/2.0/portal/payment/calculatewallet | Calculate the wallet payment amount
[**ChangeTenantWalletServiceState**](PortalPaymentAPI.md#ChangeTenantWalletServiceState) | **Post** /api/2.0/portal/payment/servicestate | Change tenant wallet service state
[**CreateCustomerMonthlyUsageReport**](PortalPaymentAPI.md#CreateCustomerMonthlyUsageReport) | **Post** /api/2.0/portal/payment/customer/usage/monthly/report | Start the customer monthly usage report generation
[**CreateCustomerOperationsReport**](PortalPaymentAPI.md#CreateCustomerOperationsReport) | **Post** /api/2.0/portal/payment/customer/operationsreport | Start the customer operations report generation
[**CreateCustomerServiceUsageReport**](PortalPaymentAPI.md#CreateCustomerServiceUsageReport) | **Post** /api/2.0/portal/payment/customer/usage/report | Start the customer service usage report generation
[**GetActiveServices**](PortalPaymentAPI.md#GetActiveServices) | **Get** /api/2.0/portal/payment/activeservices | Get the active wallet services
[**GetAiPrices**](PortalPaymentAPI.md#GetAiPrices) | **Get** /api/2.0/portal/payment/ai-prices | Get AI model prices
[**GetCheckoutSetupUrl**](PortalPaymentAPI.md#GetCheckoutSetupUrl) | **Get** /api/2.0/portal/payment/checkoutsetupurl | Get the checkout setup page URL
[**GetCustomerBalance**](PortalPaymentAPI.md#GetCustomerBalance) | **Get** /api/2.0/portal/payment/customer/balance | Get the customer balance
[**GetCustomerInfo**](PortalPaymentAPI.md#GetCustomerInfo) | **Get** /api/2.0/portal/payment/customerinfo | Get the customer information
[**GetCustomerMonthlyUsage**](PortalPaymentAPI.md#GetCustomerMonthlyUsage) | **Get** /api/2.0/portal/payment/customer/usage/monthly | Get the customer monthly usage
[**GetCustomerMonthlyUsageReport**](PortalPaymentAPI.md#GetCustomerMonthlyUsageReport) | **Get** /api/2.0/portal/payment/customer/usage/monthly/report | Get the status of the customer monthly usage report generation
[**GetCustomerOperations**](PortalPaymentAPI.md#GetCustomerOperations) | **Get** /api/2.0/portal/payment/customer/operations | Get the customer operations
[**GetCustomerOperationsReport**](PortalPaymentAPI.md#GetCustomerOperationsReport) | **Get** /api/2.0/portal/payment/customer/operationsreport | Get the status of the customer operations report generation
[**GetCustomerServiceUsage**](PortalPaymentAPI.md#GetCustomerServiceUsage) | **Get** /api/2.0/portal/payment/customer/usage | Get the customer service usage
[**GetCustomerServiceUsageReport**](PortalPaymentAPI.md#GetCustomerServiceUsageReport) | **Get** /api/2.0/portal/payment/customer/usage/report | Get the status of the customer service usage report generation
[**GetPaymentAccount**](PortalPaymentAPI.md#GetPaymentAccount) | **Get** /api/2.0/portal/payment/account | Get the payment account
[**GetPaymentCurrencies**](PortalPaymentAPI.md#GetPaymentCurrencies) | **Get** /api/2.0/portal/payment/currencies | Get currencies
[**GetPaymentQuotas**](PortalPaymentAPI.md#GetPaymentQuotas) | **Get** /api/2.0/portal/payment/quotas | Get quotas
[**GetPaymentUrl**](PortalPaymentAPI.md#GetPaymentUrl) | **Put** /api/2.0/portal/payment/url | Get the payment page URL
[**GetPortalPrices**](PortalPaymentAPI.md#GetPortalPrices) | **Get** /api/2.0/portal/payment/prices | Get prices
[**GetQuotaPaymentInformation**](PortalPaymentAPI.md#GetQuotaPaymentInformation) | **Get** /api/2.0/portal/payment/quota | Get quota payment information
[**GetRestrictedAiModels**](PortalPaymentAPI.md#GetRestrictedAiModels) | **Get** /api/2.0/portal/payment/ai-model/restrictions | Get restricted AI models
[**GetSubscriptionBalanceInfo**](PortalPaymentAPI.md#GetSubscriptionBalanceInfo) | **Get** /api/2.0/portal/payment/subscription/balance | Get the subscription balance information
[**GetTenantWalletServiceSettings**](PortalPaymentAPI.md#GetTenantWalletServiceSettings) | **Get** /api/2.0/portal/payment/servicessettings | Gets the wallet service settings for the tenant.
[**GetTenantWalletSettings**](PortalPaymentAPI.md#GetTenantWalletSettings) | **Get** /api/2.0/portal/payment/topupsettings | Gets the tenant wallet auto top up settings
[**GetWalletService**](PortalPaymentAPI.md#GetWalletService) | **Get** /api/2.0/portal/payment/walletservice | Get wallet service
[**GetWalletServices**](PortalPaymentAPI.md#GetWalletServices) | **Get** /api/2.0/portal/payment/walletservices | Get wallet services
[**MoveSubscriptionToWallet**](PortalPaymentAPI.md#MoveSubscriptionToWallet) | **Post** /api/2.0/portal/payment/subscription/movetowallet | Move the subscription balance to the wallet and purchase admins
[**SendPaymentRequest**](PortalPaymentAPI.md#SendPaymentRequest) | **Post** /api/2.0/portal/payment/request | Send a payment request
[**SetRestrictedAiModels**](PortalPaymentAPI.md#SetRestrictedAiModels) | **Put** /api/2.0/portal/payment/ai-model/restrictions | Set restricted AI models
[**SetTenantWalletSettings**](PortalPaymentAPI.md#SetTenantWalletSettings) | **Post** /api/2.0/portal/payment/topupsettings | Set the wallet auto top up settings
[**TerminateCustomerMonthlyUsageReport**](PortalPaymentAPI.md#TerminateCustomerMonthlyUsageReport) | **Delete** /api/2.0/portal/payment/customer/usage/monthly/report | Terminate the customer monthly usage report generation
[**TerminateCustomerOperationsReport**](PortalPaymentAPI.md#TerminateCustomerOperationsReport) | **Delete** /api/2.0/portal/payment/customer/operationsreport | Terminate the customer operations report generation
[**TerminateCustomerServiceUsageReport**](PortalPaymentAPI.md#TerminateCustomerServiceUsageReport) | **Delete** /api/2.0/portal/payment/customer/usage/report | Terminate the customer service usage report generation
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
	walletQuantityRequestDto := *openapiclient.NewWalletQuantityRequestDto(map[string]*int32{"key": int32(123)}) // WalletQuantityRequestDto |  (optional)

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


## CreateCustomerMonthlyUsageReport

> DocumentBuilderTaskWrapper CreateCustomerMonthlyUsageReport(ctx).CustomerMonthlyUsageReportRequestDto(customerMonthlyUsageReportRequestDto).Execute()

Start the customer monthly usage report generation



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-customer-monthly-usage-report/).

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
	customerMonthlyUsageReportRequestDto := *openapiclient.NewCustomerMonthlyUsageReportRequestDto() // CustomerMonthlyUsageReportRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.CreateCustomerMonthlyUsageReport(context.Background()).CustomerMonthlyUsageReportRequestDto(customerMonthlyUsageReportRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.CreateCustomerMonthlyUsageReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomerMonthlyUsageReport`: DocumentBuilderTaskWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.CreateCustomerMonthlyUsageReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerMonthlyUsageReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerMonthlyUsageReportRequestDto** | [**CustomerMonthlyUsageReportRequestDto**](CustomerMonthlyUsageReportRequestDto.md) |  | 

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


## CreateCustomerServiceUsageReport

> DocumentBuilderTaskWrapper CreateCustomerServiceUsageReport(ctx).CustomerServiceUsageReportRequestDto(customerServiceUsageReportRequestDto).Execute()

Start the customer service usage report generation



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/create-customer-service-usage-report/).

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
	customerServiceUsageReportRequestDto := *openapiclient.NewCustomerServiceUsageReportRequestDto() // CustomerServiceUsageReportRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.CreateCustomerServiceUsageReport(context.Background()).CustomerServiceUsageReportRequestDto(customerServiceUsageReportRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.CreateCustomerServiceUsageReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomerServiceUsageReport`: DocumentBuilderTaskWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.CreateCustomerServiceUsageReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerServiceUsageReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerServiceUsageReportRequestDto** | [**CustomerServiceUsageReportRequestDto**](CustomerServiceUsageReportRequestDto.md) |  | 

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


## GetActiveServices

> ActiveServiceArrayWrapper GetActiveServices(ctx).Execute()

Get the active wallet services



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-active-services/).

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
	resp, r, err := apiClient.PortalPaymentAPI.GetActiveServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetActiveServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveServices`: ActiveServiceArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetActiveServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveServicesRequest struct via the builder pattern


### Return type

[**ActiveServiceArrayWrapper**](ActiveServiceArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
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


## GetCustomerMonthlyUsage

> CustomerMonthlyUsageArrayWrapper GetCustomerMonthlyUsage(ctx).StartDate(startDate).EndDate(endDate).Execute()

Get the customer monthly usage



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-customer-monthly-usage/).

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
	startDate := time.Now() // time.Time | Start of the period (inclusive). (optional)
	endDate := time.Now() // time.Time | End of the period (inclusive). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.GetCustomerMonthlyUsage(context.Background()).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetCustomerMonthlyUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerMonthlyUsage`: CustomerMonthlyUsageArrayWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetCustomerMonthlyUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerMonthlyUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** | Start of the period (inclusive). | 
 **endDate** | **time.Time** | End of the period (inclusive). | 

### Return type

[**CustomerMonthlyUsageArrayWrapper**](CustomerMonthlyUsageArrayWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerMonthlyUsageReport

> DocumentBuilderTaskWrapper GetCustomerMonthlyUsageReport(ctx).Execute()

Get the status of the customer monthly usage report generation



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-customer-monthly-usage-report/).

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
	resp, r, err := apiClient.PortalPaymentAPI.GetCustomerMonthlyUsageReport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetCustomerMonthlyUsageReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerMonthlyUsageReport`: DocumentBuilderTaskWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetCustomerMonthlyUsageReport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerMonthlyUsageReportRequest struct via the builder pattern


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
	serviceName := []string{"Inner_example"} // []string | The service name list. A single string is also accepted for backward compatibility. (optional)
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
 **serviceName** | **[]string** | The service name list. A single string is also accepted for backward compatibility. | 
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


## GetCustomerServiceUsage

> CustomerServiceUsageReportWrapper GetCustomerServiceUsage(ctx).ServiceName(serviceName).ParticipantName(participantName).Status(status).StartDate(startDate).EndDate(endDate).Metadata(metadata).Offset(offset).Limit(limit).OrderBy(orderBy).OrderType(orderType).Execute()

Get the customer service usage



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-customer-service-usage/).

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
	serviceName := []string{"Inner_example"} // []string | The service name list. (optional)
	participantName := "My Own Corporation" // string | The participant name. (optional)
	status := openapiclient.OperationStatus(0) // OperationStatus | The operation status to filter by. (optional)
	startDate := time.Now() // time.Time | Start of the period (inclusive). (optional)
	endDate := time.Now() // time.Time | End of the period (inclusive). (optional)
	metadata := map[string]*string{"key": map[string]*string{"key": "Inner_example"}} // map[string]*string | Metadata key-value pairs to filter by. (optional)
	offset := int32(0) // int32 | The number of items to skip for pagination. The default value is 0. (optional)
	limit := int32(25) // int32 | The maximum number of items to return for pagination. The default value is 25. (optional)
	orderBy := "ServiceName" // string | The field to order by. (optional)
	orderType := openapiclient.OperationOrderType(0) // OperationOrderType | Order direction: Ascending or Descending. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.GetCustomerServiceUsage(context.Background()).ServiceName(serviceName).ParticipantName(participantName).Status(status).StartDate(startDate).EndDate(endDate).Metadata(metadata).Offset(offset).Limit(limit).OrderBy(orderBy).OrderType(orderType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetCustomerServiceUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerServiceUsage`: CustomerServiceUsageReportWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetCustomerServiceUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerServiceUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceName** | **[]string** | The service name list. | 
 **participantName** | **string** | The participant name. | 
 **status** | [**OperationStatus**](OperationStatus.md) | The operation status to filter by. | 
 **startDate** | **time.Time** | Start of the period (inclusive). | 
 **endDate** | **time.Time** | End of the period (inclusive). | 
 **metadata** | **map[string]map[string]*string** | Metadata key-value pairs to filter by. | 
 **offset** | **int32** | The number of items to skip for pagination. The default value is 0. | 
 **limit** | **int32** | The maximum number of items to return for pagination. The default value is 25. | 
 **orderBy** | **string** | The field to order by. | 
 **orderType** | [**OperationOrderType**](OperationOrderType.md) | Order direction: Ascending or Descending. | 

### Return type

[**CustomerServiceUsageReportWrapper**](CustomerServiceUsageReportWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerServiceUsageReport

> DocumentBuilderTaskWrapper GetCustomerServiceUsageReport(ctx).Execute()

Get the status of the customer service usage report generation



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-customer-service-usage-report/).

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
	resp, r, err := apiClient.PortalPaymentAPI.GetCustomerServiceUsageReport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetCustomerServiceUsageReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerServiceUsageReport`: DocumentBuilderTaskWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetCustomerServiceUsageReport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerServiceUsageReportRequest struct via the builder pattern


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

> QuotaArrayWrapper GetPaymentQuotas(ctx).Wallet(wallet).Additional(additional).Execute()

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
	additional := true // bool | Specifies whether to return additional quotas only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.GetPaymentQuotas(context.Background()).Wallet(wallet).Additional(additional).Execute()
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
 **additional** | **bool** | Specifies whether to return additional quotas only. | 

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
	paymentUrlRequestDto := *openapiclient.NewPaymentUrlRequestDto("https://example.com/payment/back", "https://example.com/payment/success", map[string]int32{"key": int32(123)}) // PaymentUrlRequestDto |  (optional)

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


## GetSubscriptionBalanceInfo

> SubscriptionBalanceInfoWrapper GetSubscriptionBalanceInfo(ctx).Execute()

Get the subscription balance information



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/get-subscription-balance-info/).

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
	resp, r, err := apiClient.PortalPaymentAPI.GetSubscriptionBalanceInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.GetSubscriptionBalanceInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscriptionBalanceInfo`: SubscriptionBalanceInfoWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetSubscriptionBalanceInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionBalanceInfoRequest struct via the builder pattern


### Return type

[**SubscriptionBalanceInfoWrapper**](SubscriptionBalanceInfoWrapper.md)

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

> TenantWalletSettingsResponseWrapper GetTenantWalletSettings(ctx).Execute()

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
	// response from `GetTenantWalletSettings`: TenantWalletSettingsResponseWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.GetTenantWalletSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantWalletSettingsRequest struct via the builder pattern


### Return type

[**TenantWalletSettingsResponseWrapper**](TenantWalletSettingsResponseWrapper.md)

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
	service := openapiclient.TenantWalletService(-18) // TenantWalletService | The wallet service type.

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


## MoveSubscriptionToWallet

> BooleanWrapper MoveSubscriptionToWallet(ctx).QuantityRequestDto(quantityRequestDto).Execute()

Move the subscription balance to the wallet and purchase admins



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/move-subscription-to-wallet/).

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
	quantityRequestDto := *openapiclient.NewQuantityRequestDto(map[string]int32{"key": int32(123)}) // QuantityRequestDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalPaymentAPI.MoveSubscriptionToWallet(context.Background()).QuantityRequestDto(quantityRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.MoveSubscriptionToWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveSubscriptionToWallet`: BooleanWrapper
	fmt.Fprintf(os.Stdout, "Response from `PortalPaymentAPI.MoveSubscriptionToWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoveSubscriptionToWalletRequest struct via the builder pattern


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
	salesRequestsDto := *openapiclient.NewSalesRequestsDto("John Doe", "user@example.com", "I would like to inquire about pricing") // SalesRequestsDto |  (optional)

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
- **Accept**: application/json

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

> TenantWalletSettingsResponseWrapper SetTenantWalletSettings(ctx).TenantWalletSettingsWrapper(tenantWalletSettingsWrapper).Execute()

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
	// response from `SetTenantWalletSettings`: TenantWalletSettingsResponseWrapper
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

[**TenantWalletSettingsResponseWrapper**](TenantWalletSettingsResponseWrapper.md)

### Authorization

[Basic](../README.md#Basic), [OAuth2](../README.md#OAuth2), [ApiKeyBearer](../README.md#ApiKeyBearer), [asc_auth_key](../README.md#asc_auth_key), [Bearer](../README.md#Bearer), [OpenId](../README.md#OpenId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateCustomerMonthlyUsageReport

> TerminateCustomerMonthlyUsageReport(ctx).Execute()

Terminate the customer monthly usage report generation



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/terminate-customer-monthly-usage-report/).

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
	r, err := apiClient.PortalPaymentAPI.TerminateCustomerMonthlyUsageReport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.TerminateCustomerMonthlyUsageReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateCustomerMonthlyUsageReportRequest struct via the builder pattern


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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateCustomerServiceUsageReport

> TerminateCustomerServiceUsageReport(ctx).Execute()

Terminate the customer service usage report generation



For more information, see [api.onlyoffice.com](https://api.onlyoffice.com/docspace/api-backend/usage-api/terminate-customer-service-usage-report/).

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
	r, err := apiClient.PortalPaymentAPI.TerminateCustomerServiceUsageReport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalPaymentAPI.TerminateCustomerServiceUsageReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateCustomerServiceUsageReportRequest struct via the builder pattern


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
	quantityRequestDto := *openapiclient.NewQuantityRequestDto(map[string]int32{"key": int32(123)}) // QuantityRequestDto |  (optional)

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
	walletQuantityRequestDto := *openapiclient.NewWalletQuantityRequestDto(map[string]*int32{"key": int32(123)}) // WalletQuantityRequestDto |  (optional)

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

