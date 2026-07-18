// (c) Copyright Ascensio System SIA 2026
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docspace_api_sdk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"time"
)


// PortalPaymentAPIService PortalPaymentAPI service
type PortalPaymentAPIService service

type ApiCalculateWalletPaymentRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
	walletQuantityRequestDto *WalletQuantityRequestDto
}

func (r ApiCalculateWalletPaymentRequest) WalletQuantityRequestDto(walletQuantityRequestDto WalletQuantityRequestDto) ApiCalculateWalletPaymentRequest {	r.walletQuantityRequestDto = &walletQuantityRequestDto
	return r
}

func (r ApiCalculateWalletPaymentRequest) Execute() (*PaymentCalculationWrapper, *http.Response, error) {
	return r.ApiService.CalculateWalletPaymentExecute(r)
}

// CalculateWalletPayment Calculate the wallet payment amount
//
// Calculates an amount of the wallet payment with the parameters specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/calculate-wallet-payment/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiCalculateWalletPaymentRequest
func (a *PortalPaymentAPIService) CalculateWalletPayment(ctx context.Context) ApiCalculateWalletPaymentRequest {
	return ApiCalculateWalletPaymentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaymentCalculationWrapper
func (a *PortalPaymentAPIService) CalculateWalletPaymentExecute(r ApiCalculateWalletPaymentRequest) (*PaymentCalculationWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentCalculationWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.CalculateWalletPayment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/calculatewallet"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.walletQuantityRequestDto
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiChangeTenantWalletServiceStateRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
	changeWalletServiceStateRequestDto *ChangeWalletServiceStateRequestDto
}

func (r ApiChangeTenantWalletServiceStateRequest) ChangeWalletServiceStateRequestDto(changeWalletServiceStateRequestDto ChangeWalletServiceStateRequestDto) ApiChangeTenantWalletServiceStateRequest {	r.changeWalletServiceStateRequestDto = &changeWalletServiceStateRequestDto
	return r
}

func (r ApiChangeTenantWalletServiceStateRequest) Execute() (*TenantWalletServiceSettingsWrapper, *http.Response, error) {
	return r.ApiService.ChangeTenantWalletServiceStateExecute(r)
}

// ChangeTenantWalletServiceState Change tenant wallet service state
//
// Changes the state of a wallet service for the current tenant.  Requires permission to edit portal settings and a configured tariff service.  Adds or removes the specified service from the enabled services list based on the enabled flag.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/change-tenant-wallet-service-state/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiChangeTenantWalletServiceStateRequest
func (a *PortalPaymentAPIService) ChangeTenantWalletServiceState(ctx context.Context) ApiChangeTenantWalletServiceStateRequest {
	return ApiChangeTenantWalletServiceStateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TenantWalletServiceSettingsWrapper
func (a *PortalPaymentAPIService) ChangeTenantWalletServiceStateExecute(r ApiChangeTenantWalletServiceStateRequest) (*TenantWalletServiceSettingsWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TenantWalletServiceSettingsWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.ChangeTenantWalletServiceState")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/servicestate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.changeWalletServiceStateRequestDto
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateCustomerOperationsReportRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
	customerOperationsReportRequestDto *CustomerOperationsReportRequestDto
}

func (r ApiCreateCustomerOperationsReportRequest) CustomerOperationsReportRequestDto(customerOperationsReportRequestDto CustomerOperationsReportRequestDto) ApiCreateCustomerOperationsReportRequest {	r.customerOperationsReportRequestDto = &customerOperationsReportRequestDto
	return r
}

func (r ApiCreateCustomerOperationsReportRequest) Execute() (*DocumentBuilderTaskWrapper, *http.Response, error) {
	return r.ApiService.CreateCustomerOperationsReportExecute(r)
}

// CreateCustomerOperationsReport Start the customer operations report generation
//
// Starts generating a customer operations report as an xlsx file and saves it in Documents.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/create-customer-operations-report/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiCreateCustomerOperationsReportRequest
func (a *PortalPaymentAPIService) CreateCustomerOperationsReport(ctx context.Context) ApiCreateCustomerOperationsReportRequest {
	return ApiCreateCustomerOperationsReportRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DocumentBuilderTaskWrapper
func (a *PortalPaymentAPIService) CreateCustomerOperationsReportExecute(r ApiCreateCustomerOperationsReportRequest) (*DocumentBuilderTaskWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DocumentBuilderTaskWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.CreateCustomerOperationsReport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/customer/operationsreport"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.customerOperationsReportRequestDto
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreditAiBalanceRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
	creditAiBalanceRequestDto *CreditAiBalanceRequestDto
}

func (r ApiCreditAiBalanceRequest) CreditAiBalanceRequestDto(creditAiBalanceRequestDto CreditAiBalanceRequestDto) ApiCreditAiBalanceRequest {	r.creditAiBalanceRequestDto = &creditAiBalanceRequestDto
	return r
}

func (r ApiCreditAiBalanceRequest) Execute() (*ServicePaymentWrapper, *http.Response, error) {
	return r.ApiService.CreditAiBalanceExecute(r)
}

// CreditAiBalance Credit AI balance
//
// Credits AI quota to the customer AI sub-account from their main balance.  Requires the customer to have a configured payment method.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/credit-ai-balance/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiCreditAiBalanceRequest
func (a *PortalPaymentAPIService) CreditAiBalance(ctx context.Context) ApiCreditAiBalanceRequest {
	return ApiCreditAiBalanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ServicePaymentWrapper
func (a *PortalPaymentAPIService) CreditAiBalanceExecute(r ApiCreditAiBalanceRequest) (*ServicePaymentWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServicePaymentWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.CreditAiBalance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/creditaibalance"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.creditAiBalanceRequestDto
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAiPricesRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
}

func (r ApiGetAiPricesRequest) Execute() (*AiPricesResponseWrapper, *http.Response, error) {
	return r.ApiService.GetAiPricesExecute(r)
}

// GetAiPrices Get AI model prices
//
// Retrieves the pricing information for AI models including chat, embedding, and web search services.  The prices are returned in the configured currency and normalized per million tokens.  Requires administrator permissions to access.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-ai-prices/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetAiPricesRequest
func (a *PortalPaymentAPIService) GetAiPrices(ctx context.Context) ApiGetAiPricesRequest {
	return ApiGetAiPricesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AiPricesResponseWrapper
func (a *PortalPaymentAPIService) GetAiPricesExecute(r ApiGetAiPricesRequest) (*AiPricesResponseWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiPricesResponseWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.GetAiPrices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/ai-prices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCheckoutSetupUrlRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
	backUrl *string
	successUrl *string
}

// The URL where the user will be redirected after setup cancellation.
func (r ApiGetCheckoutSetupUrlRequest) BackUrl(backUrl string) ApiGetCheckoutSetupUrlRequest {	r.backUrl = &backUrl
	return r
}

// The URL where the user will be redirected after successful payment.
func (r ApiGetCheckoutSetupUrlRequest) SuccessUrl(successUrl string) ApiGetCheckoutSetupUrlRequest {	r.successUrl = &successUrl
	return r
}

func (r ApiGetCheckoutSetupUrlRequest) Execute() (*StringWrapper, *http.Response, error) {
	return r.ApiService.GetCheckoutSetupUrlExecute(r)
}

// GetCheckoutSetupUrl Get the checkout setup page URL
//
// Returns the URL to the checkout setup page.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-checkout-setup-url/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetCheckoutSetupUrlRequest
func (a *PortalPaymentAPIService) GetCheckoutSetupUrl(ctx context.Context) ApiGetCheckoutSetupUrlRequest {
	return ApiGetCheckoutSetupUrlRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StringWrapper
func (a *PortalPaymentAPIService) GetCheckoutSetupUrlExecute(r ApiGetCheckoutSetupUrlRequest) (*StringWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StringWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.GetCheckoutSetupUrl")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/checkoutsetupurl"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.backUrl == nil {
		return localVarReturnValue, nil, reportError("backUrl is required and must be specified")
	}
	if strlen(*r.backUrl) < 0 {
		return localVarReturnValue, nil, reportError("backUrl must have at least 0 elements")
	}
	if strlen(*r.backUrl) > 255 {
		return localVarReturnValue, nil, reportError("backUrl must have less than 255 elements")
	}
	if r.successUrl == nil {
		return localVarReturnValue, nil, reportError("successUrl is required and must be specified")
	}
	if strlen(*r.successUrl) < 0 {
		return localVarReturnValue, nil, reportError("successUrl must have at least 0 elements")
	}
	if strlen(*r.successUrl) > 255 {
		return localVarReturnValue, nil, reportError("successUrl must have less than 255 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "BackUrl", r.backUrl, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "SuccessUrl", r.successUrl, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCustomerAiBalanceRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
	refresh *bool
}

// Specifies whether to refresh the payment information cache or not.
func (r ApiGetCustomerAiBalanceRequest) Refresh(refresh bool) ApiGetCustomerAiBalanceRequest {	r.refresh = &refresh
	return r
}

func (r ApiGetCustomerAiBalanceRequest) Execute() (*BalanceWrapper, *http.Response, error) {
	return r.ApiService.GetCustomerAiBalanceExecute(r)
}

// GetCustomerAiBalance Get the customer AI balance
//
// Returns the AI quota balance of a customer from the accounting service.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-customer-ai-balance/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetCustomerAiBalanceRequest
func (a *PortalPaymentAPIService) GetCustomerAiBalance(ctx context.Context) ApiGetCustomerAiBalanceRequest {
	return ApiGetCustomerAiBalanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BalanceWrapper
func (a *PortalPaymentAPIService) GetCustomerAiBalanceExecute(r ApiGetCustomerAiBalanceRequest) (*BalanceWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BalanceWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.GetCustomerAiBalance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/customer/aibalance"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.refresh != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "refresh", r.refresh, "form", "")
			}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCustomerBalanceRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
	refresh *bool
}

// Specifies whether to refresh the payment information cache or not.
func (r ApiGetCustomerBalanceRequest) Refresh(refresh bool) ApiGetCustomerBalanceRequest {	r.refresh = &refresh
	return r
}

func (r ApiGetCustomerBalanceRequest) Execute() (*BalanceWrapper, *http.Response, error) {
	return r.ApiService.GetCustomerBalanceExecute(r)
}

// GetCustomerBalance Get the customer balance
//
// Returns the customer balance from the accounting service.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-customer-balance/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetCustomerBalanceRequest
func (a *PortalPaymentAPIService) GetCustomerBalance(ctx context.Context) ApiGetCustomerBalanceRequest {
	return ApiGetCustomerBalanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BalanceWrapper
func (a *PortalPaymentAPIService) GetCustomerBalanceExecute(r ApiGetCustomerBalanceRequest) (*BalanceWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BalanceWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.GetCustomerBalance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/customer/balance"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.refresh != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "refresh", r.refresh, "form", "")
			}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCustomerInfoRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
	refresh *bool
}

// Specifies whether to refresh the payment information cache or not.
func (r ApiGetCustomerInfoRequest) Refresh(refresh bool) ApiGetCustomerInfoRequest {	r.refresh = &refresh
	return r
}

func (r ApiGetCustomerInfoRequest) Execute() (*CustomerInfoWrapper, *http.Response, error) {
	return r.ApiService.GetCustomerInfoExecute(r)
}

// GetCustomerInfo Get the customer information
//
// Returns the customer information.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-customer-info/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetCustomerInfoRequest
func (a *PortalPaymentAPIService) GetCustomerInfo(ctx context.Context) ApiGetCustomerInfoRequest {
	return ApiGetCustomerInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CustomerInfoWrapper
func (a *PortalPaymentAPIService) GetCustomerInfoExecute(r ApiGetCustomerInfoRequest) (*CustomerInfoWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomerInfoWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.GetCustomerInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/customerinfo"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.refresh != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "refresh", r.refresh, "form", "")
			}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCustomerOperationsRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
	offset *int32
	limit *int32
	serviceName *string
	startDate *time.Time
	endDate *time.Time
	participantName *string
	credit *bool
	debit *bool
	type_ *OperationType
	status *OperationStatus
	orderBy *string
	orderType *OperationOrderType
}

// The number of items to skip for pagination. The default value is 0.
func (r ApiGetCustomerOperationsRequest) Offset(offset int32) ApiGetCustomerOperationsRequest {	r.offset = &offset
	return r
}

// The maximum number of items to return for pagination. The default value is 25.
func (r ApiGetCustomerOperationsRequest) Limit(limit int32) ApiGetCustomerOperationsRequest {	r.limit = &limit
	return r
}

// The service name.
func (r ApiGetCustomerOperationsRequest) ServiceName(serviceName string) ApiGetCustomerOperationsRequest {	r.serviceName = &serviceName
	return r
}

// The report start date.
func (r ApiGetCustomerOperationsRequest) StartDate(startDate time.Time) ApiGetCustomerOperationsRequest {	r.startDate = &startDate
	return r
}

// The report end date.
func (r ApiGetCustomerOperationsRequest) EndDate(endDate time.Time) ApiGetCustomerOperationsRequest {	r.endDate = &endDate
	return r
}

// The participant name.
func (r ApiGetCustomerOperationsRequest) ParticipantName(participantName string) ApiGetCustomerOperationsRequest {	r.participantName = &participantName
	return r
}

// Specifies whether to include credit operations in the report.
func (r ApiGetCustomerOperationsRequest) Credit(credit bool) ApiGetCustomerOperationsRequest {	r.credit = &credit
	return r
}

// Specifies whether to include debit operations in the report.
func (r ApiGetCustomerOperationsRequest) Debit(debit bool) ApiGetCustomerOperationsRequest {	r.debit = &debit
	return r
}

// The operation type to filter by.
func (r ApiGetCustomerOperationsRequest) Type_(type_ OperationType) ApiGetCustomerOperationsRequest {	r.type_ = &type_
	return r
}

// The operation status to filter by.
func (r ApiGetCustomerOperationsRequest) Status(status OperationStatus) ApiGetCustomerOperationsRequest {	r.status = &status
	return r
}

// The field to order by.
func (r ApiGetCustomerOperationsRequest) OrderBy(orderBy string) ApiGetCustomerOperationsRequest {	r.orderBy = &orderBy
	return r
}

// Order direction: Ascending or Descending.
func (r ApiGetCustomerOperationsRequest) OrderType(orderType OperationOrderType) ApiGetCustomerOperationsRequest {	r.orderType = &orderType
	return r
}

func (r ApiGetCustomerOperationsRequest) Execute() (*ReportWrapper, *http.Response, error) {
	return r.ApiService.GetCustomerOperationsExecute(r)
}

// GetCustomerOperations Get the customer operations
//
// Returns the report of customer operations from the accounting service.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-customer-operations/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetCustomerOperationsRequest
func (a *PortalPaymentAPIService) GetCustomerOperations(ctx context.Context) ApiGetCustomerOperationsRequest {
	return ApiGetCustomerOperationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReportWrapper
func (a *PortalPaymentAPIService) GetCustomerOperationsExecute(r ApiGetCustomerOperationsRequest) (*ReportWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReportWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.GetCustomerOperations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/customer/operations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
			}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
			}
	if r.serviceName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ServiceName", r.serviceName, "form", "")
			}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "StartDate", r.startDate, "form", "")
			}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "EndDate", r.endDate, "form", "")
			}
	if r.participantName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ParticipantName", r.participantName, "form", "")
			}
	if r.credit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Credit", r.credit, "form", "")
			}
	if r.debit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Debit", r.debit, "form", "")
			}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Type", r.type_, "form", "")
			}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Status", r.status, "form", "")
			}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "OrderBy", r.orderBy, "form", "")
			}
	if r.orderType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "OrderType", r.orderType, "form", "")
			}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCustomerOperationsReportRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
}

func (r ApiGetCustomerOperationsReportRequest) Execute() (*DocumentBuilderTaskWrapper, *http.Response, error) {
	return r.ApiService.GetCustomerOperationsReportExecute(r)
}

// GetCustomerOperationsReport Get the status of the customer operations report generation
//
// Returns the status of generating a customer operations report.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-customer-operations-report/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetCustomerOperationsReportRequest
func (a *PortalPaymentAPIService) GetCustomerOperationsReport(ctx context.Context) ApiGetCustomerOperationsReportRequest {
	return ApiGetCustomerOperationsReportRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DocumentBuilderTaskWrapper
func (a *PortalPaymentAPIService) GetCustomerOperationsReportExecute(r ApiGetCustomerOperationsReportRequest) (*DocumentBuilderTaskWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DocumentBuilderTaskWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.GetCustomerOperationsReport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/customer/operationsreport"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPaymentAccountRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
	backUrl *string
}

// The URL where the user will be redirected after payment processing.
func (r ApiGetPaymentAccountRequest) BackUrl(backUrl string) ApiGetPaymentAccountRequest {	r.backUrl = &backUrl
	return r
}

func (r ApiGetPaymentAccountRequest) Execute() (*StringWrapper, *http.Response, error) {
	return r.ApiService.GetPaymentAccountExecute(r)
}

// GetPaymentAccount Get the payment account
//
// Returns the URL to the payment account.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-payment-account/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetPaymentAccountRequest
func (a *PortalPaymentAPIService) GetPaymentAccount(ctx context.Context) ApiGetPaymentAccountRequest {
	return ApiGetPaymentAccountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StringWrapper
func (a *PortalPaymentAPIService) GetPaymentAccountExecute(r ApiGetPaymentAccountRequest) (*StringWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StringWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.GetPaymentAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/account"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.backUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "backUrl", r.backUrl, "form", "")
			}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPaymentCurrenciesRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
}

func (r ApiGetPaymentCurrenciesRequest) Execute() (*CurrenciesArrayWrapper, *http.Response, error) {
	return r.ApiService.GetPaymentCurrenciesExecute(r)
}

// GetPaymentCurrencies Get currencies
//
// Returns the available portal currencies.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-payment-currencies/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetPaymentCurrenciesRequest
func (a *PortalPaymentAPIService) GetPaymentCurrencies(ctx context.Context) ApiGetPaymentCurrenciesRequest {
	return ApiGetPaymentCurrenciesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CurrenciesArrayWrapper
func (a *PortalPaymentAPIService) GetPaymentCurrenciesExecute(r ApiGetPaymentCurrenciesRequest) (*CurrenciesArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CurrenciesArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.GetPaymentCurrencies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/currencies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPaymentQuotasRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
	wallet *bool
}

// Specifies whether to return the wallet quotas only.
func (r ApiGetPaymentQuotasRequest) Wallet(wallet bool) ApiGetPaymentQuotasRequest {	r.wallet = &wallet
	return r
}

func (r ApiGetPaymentQuotasRequest) Execute() (*QuotaArrayWrapper, *http.Response, error) {
	return r.ApiService.GetPaymentQuotasExecute(r)
}

// GetPaymentQuotas Get quotas
//
// Returns the available portal quotas.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-payment-quotas/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetPaymentQuotasRequest
func (a *PortalPaymentAPIService) GetPaymentQuotas(ctx context.Context) ApiGetPaymentQuotasRequest {
	return ApiGetPaymentQuotasRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return QuotaArrayWrapper
func (a *PortalPaymentAPIService) GetPaymentQuotasExecute(r ApiGetPaymentQuotasRequest) (*QuotaArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QuotaArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.GetPaymentQuotas")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/quotas"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.wallet != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "wallet", r.wallet, "form", "")
			}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPaymentUrlRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
	paymentUrlRequestDto *PaymentUrlRequestDto
}

func (r ApiGetPaymentUrlRequest) PaymentUrlRequestDto(paymentUrlRequestDto PaymentUrlRequestDto) ApiGetPaymentUrlRequest {	r.paymentUrlRequestDto = &paymentUrlRequestDto
	return r
}

func (r ApiGetPaymentUrlRequest) Execute() (*StringWrapper, *http.Response, error) {
	return r.ApiService.GetPaymentUrlExecute(r)
}

// GetPaymentUrl Get the payment page URL
//
// Returns the URL to the payment page.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-payment-url/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetPaymentUrlRequest
func (a *PortalPaymentAPIService) GetPaymentUrl(ctx context.Context) ApiGetPaymentUrlRequest {
	return ApiGetPaymentUrlRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StringWrapper
func (a *PortalPaymentAPIService) GetPaymentUrlExecute(r ApiGetPaymentUrlRequest) (*StringWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StringWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.GetPaymentUrl")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/url"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.paymentUrlRequestDto
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPortalPricesRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
}

func (r ApiGetPortalPricesRequest) Execute() (*GetPortalPrices200Response, *http.Response, error) {
	return r.ApiService.GetPortalPricesExecute(r)
}

// GetPortalPrices Get prices
//
// Returns the available portal prices.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-portal-prices/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetPortalPricesRequest
func (a *PortalPaymentAPIService) GetPortalPrices(ctx context.Context) ApiGetPortalPricesRequest {
	return ApiGetPortalPricesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetPortalPrices200Response
func (a *PortalPaymentAPIService) GetPortalPricesExecute(r ApiGetPortalPricesRequest) (*GetPortalPrices200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetPortalPrices200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.GetPortalPrices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/prices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetQuotaPaymentInformationRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
	refresh *bool
}

// Specifies whether to refresh the payment information cache or not.
func (r ApiGetQuotaPaymentInformationRequest) Refresh(refresh bool) ApiGetQuotaPaymentInformationRequest {	r.refresh = &refresh
	return r
}

func (r ApiGetQuotaPaymentInformationRequest) Execute() (*QuotaWrapper, *http.Response, error) {
	return r.ApiService.GetQuotaPaymentInformationExecute(r)
}

// GetQuotaPaymentInformation Get quota payment information
//
// Returns the payment information about the current portal quota.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-quota-payment-information/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetQuotaPaymentInformationRequest
func (a *PortalPaymentAPIService) GetQuotaPaymentInformation(ctx context.Context) ApiGetQuotaPaymentInformationRequest {
	return ApiGetQuotaPaymentInformationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return QuotaWrapper
func (a *PortalPaymentAPIService) GetQuotaPaymentInformationExecute(r ApiGetQuotaPaymentInformationRequest) (*QuotaWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QuotaWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.GetQuotaPaymentInformation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/quota"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.refresh != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "refresh", r.refresh, "form", "")
			}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetRestrictedAiModelsRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
}

func (r ApiGetRestrictedAiModelsRequest) Execute() (*RestrictedModelsResponseWrapper, *http.Response, error) {
	return r.ApiService.GetRestrictedAiModelsExecute(r)
}

// GetRestrictedAiModels Get restricted AI models
//
// Returns the list of AI chat model IDs that are restricted (disabled) for the current tenant.  Restricted models cannot be used for AI chat conversations by any user within the portal.  Only DocSpace administrators can access this endpoint.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-restricted-ai-models/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetRestrictedAiModelsRequest
func (a *PortalPaymentAPIService) GetRestrictedAiModels(ctx context.Context) ApiGetRestrictedAiModelsRequest {
	return ApiGetRestrictedAiModelsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RestrictedModelsResponseWrapper
func (a *PortalPaymentAPIService) GetRestrictedAiModelsExecute(r ApiGetRestrictedAiModelsRequest) (*RestrictedModelsResponseWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RestrictedModelsResponseWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.GetRestrictedAiModels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/ai-model/restrictions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetTenantWalletServiceSettingsRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
}

func (r ApiGetTenantWalletServiceSettingsRequest) Execute() (*TenantWalletServiceSettingsWrapper, *http.Response, error) {
	return r.ApiService.GetTenantWalletServiceSettingsExecute(r)
}

// GetTenantWalletServiceSettings Gets the wallet service settings for the tenant.
//
// Retrieves configuration settings related to the wallet service associated with the current tenant.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-tenant-wallet-service-settings/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetTenantWalletServiceSettingsRequest
func (a *PortalPaymentAPIService) GetTenantWalletServiceSettings(ctx context.Context) ApiGetTenantWalletServiceSettingsRequest {
	return ApiGetTenantWalletServiceSettingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TenantWalletServiceSettingsWrapper
func (a *PortalPaymentAPIService) GetTenantWalletServiceSettingsExecute(r ApiGetTenantWalletServiceSettingsRequest) (*TenantWalletServiceSettingsWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TenantWalletServiceSettingsWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.GetTenantWalletServiceSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/servicessettings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetTenantWalletSettingsRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
}

func (r ApiGetTenantWalletSettingsRequest) Execute() (*TenantWalletSettingsWrapper, *http.Response, error) {
	return r.ApiService.GetTenantWalletSettingsExecute(r)
}

// GetTenantWalletSettings Gets the tenant wallet auto top up settings
//
// Returns the wallet auto top up settings for the current tenant.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-tenant-wallet-settings/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetTenantWalletSettingsRequest
func (a *PortalPaymentAPIService) GetTenantWalletSettings(ctx context.Context) ApiGetTenantWalletSettingsRequest {
	return ApiGetTenantWalletSettingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TenantWalletSettingsWrapper
func (a *PortalPaymentAPIService) GetTenantWalletSettingsExecute(r ApiGetTenantWalletSettingsRequest) (*TenantWalletSettingsWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TenantWalletSettingsWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.GetTenantWalletSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/topupsettings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetWalletServiceRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
	service *TenantWalletService
}

// The wallet service type.
func (r ApiGetWalletServiceRequest) Service(service TenantWalletService) ApiGetWalletServiceRequest {	r.service = &service
	return r
}

func (r ApiGetWalletServiceRequest) Execute() (*WalletServiceWrapper, *http.Response, error) {
	return r.ApiService.GetWalletServiceExecute(r)
}

// GetWalletService Get wallet service
//
// Returns the specified wallet service.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-wallet-service/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetWalletServiceRequest
func (a *PortalPaymentAPIService) GetWalletService(ctx context.Context) ApiGetWalletServiceRequest {
	return ApiGetWalletServiceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WalletServiceWrapper
func (a *PortalPaymentAPIService) GetWalletServiceExecute(r ApiGetWalletServiceRequest) (*WalletServiceWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WalletServiceWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.GetWalletService")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/walletservice"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.service == nil {
		return localVarReturnValue, nil, reportError("service is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "service", r.service, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetWalletServicesRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
}

func (r ApiGetWalletServicesRequest) Execute() (*WalletServiceArrayWrapper, *http.Response, error) {
	return r.ApiService.GetWalletServicesExecute(r)
}

// GetWalletServices Get wallet services
//
// Returns the available wallet services.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/get-wallet-services/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiGetWalletServicesRequest
func (a *PortalPaymentAPIService) GetWalletServices(ctx context.Context) ApiGetWalletServicesRequest {
	return ApiGetWalletServicesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WalletServiceArrayWrapper
func (a *PortalPaymentAPIService) GetWalletServicesExecute(r ApiGetWalletServicesRequest) (*WalletServiceArrayWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WalletServiceArrayWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.GetWalletServices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/walletservices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSendPaymentRequestRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
	salesRequestsDto *SalesRequestsDto
}

func (r ApiSendPaymentRequestRequest) SalesRequestsDto(salesRequestsDto SalesRequestsDto) ApiSendPaymentRequestRequest {	r.salesRequestsDto = &salesRequestsDto
	return r
}

func (r ApiSendPaymentRequestRequest) Execute() (*http.Response, error) {
	return r.ApiService.SendPaymentRequestExecute(r)
}

// SendPaymentRequest Send a payment request
//
// Sends a request for the portal payment.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/send-payment-request/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiSendPaymentRequestRequest
func (a *PortalPaymentAPIService) SendPaymentRequest(ctx context.Context) ApiSendPaymentRequestRequest {
	return ApiSendPaymentRequestRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PortalPaymentAPIService) SendPaymentRequestExecute(r ApiSendPaymentRequestRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.SendPaymentRequest")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/request"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.salesRequestsDto
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiSetRestrictedAiModelsRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
	setRestrictedAiModelsRequestDto *SetRestrictedAiModelsRequestDto
}

func (r ApiSetRestrictedAiModelsRequest) SetRestrictedAiModelsRequestDto(setRestrictedAiModelsRequestDto SetRestrictedAiModelsRequestDto) ApiSetRestrictedAiModelsRequest {	r.setRestrictedAiModelsRequestDto = &setRestrictedAiModelsRequestDto
	return r
}

func (r ApiSetRestrictedAiModelsRequest) Execute() (*RestrictedModelsResponseWrapper, *http.Response, error) {
	return r.ApiService.SetRestrictedAiModelsExecute(r)
}

// SetRestrictedAiModels Set restricted AI models
//
// Overwrites the entire set of restricted AI model IDs for the current tenant.  The request body must contain the complete desired set — to add a restriction, include the new model alongside existing ones;  to remove one, omit it. An empty set lifts all restrictions. Only the portal payer can perform this action.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/set-restricted-ai-models/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiSetRestrictedAiModelsRequest
func (a *PortalPaymentAPIService) SetRestrictedAiModels(ctx context.Context) ApiSetRestrictedAiModelsRequest {
	return ApiSetRestrictedAiModelsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RestrictedModelsResponseWrapper
func (a *PortalPaymentAPIService) SetRestrictedAiModelsExecute(r ApiSetRestrictedAiModelsRequest) (*RestrictedModelsResponseWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RestrictedModelsResponseWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.SetRestrictedAiModels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/ai-model/restrictions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.setRestrictedAiModelsRequestDto
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSetTenantWalletSettingsRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
	tenantWalletSettingsWrapper *TenantWalletSettingsWrapper
}

func (r ApiSetTenantWalletSettingsRequest) TenantWalletSettingsWrapper(tenantWalletSettingsWrapper TenantWalletSettingsWrapper) ApiSetTenantWalletSettingsRequest {	r.tenantWalletSettingsWrapper = &tenantWalletSettingsWrapper
	return r
}

func (r ApiSetTenantWalletSettingsRequest) Execute() (*TenantWalletSettingsWrapper, *http.Response, error) {
	return r.ApiService.SetTenantWalletSettingsExecute(r)
}

// SetTenantWalletSettings Set the wallet auto top up settings
//
// Updates the wallet auto top up settings for the current tenant.  Requires the tariff service to be configured and the user to be authorized as a payer.  Returns null if the tariff service is not configured or customer information/balance cannot be retrieved.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/set-tenant-wallet-settings/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiSetTenantWalletSettingsRequest
func (a *PortalPaymentAPIService) SetTenantWalletSettings(ctx context.Context) ApiSetTenantWalletSettingsRequest {
	return ApiSetTenantWalletSettingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TenantWalletSettingsWrapper
func (a *PortalPaymentAPIService) SetTenantWalletSettingsExecute(r ApiSetTenantWalletSettingsRequest) (*TenantWalletSettingsWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TenantWalletSettingsWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.SetTenantWalletSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/topupsettings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.tenantWalletSettingsWrapper
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiTerminateCustomerOperationsReportRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
}

func (r ApiTerminateCustomerOperationsReportRequest) Execute() (*http.Response, error) {
	return r.ApiService.TerminateCustomerOperationsReportExecute(r)
}

// TerminateCustomerOperationsReport Terminate the customer operations report generation
//
// Terminates generating a customer operations report.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/terminate-customer-operations-report/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiTerminateCustomerOperationsReportRequest
func (a *PortalPaymentAPIService) TerminateCustomerOperationsReport(ctx context.Context) ApiTerminateCustomerOperationsReportRequest {
	return ApiTerminateCustomerOperationsReportRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PortalPaymentAPIService) TerminateCustomerOperationsReportExecute(r ApiTerminateCustomerOperationsReportRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.TerminateCustomerOperationsReport")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/customer/operationsreport"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiTopUpDepositRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
	topUpDepositRequestDto *TopUpDepositRequestDto
}

func (r ApiTopUpDepositRequest) TopUpDepositRequestDto(topUpDepositRequestDto TopUpDepositRequestDto) ApiTopUpDepositRequest {	r.topUpDepositRequestDto = &topUpDepositRequestDto
	return r
}

func (r ApiTopUpDepositRequest) Execute() (*BooleanWrapper, *http.Response, error) {
	return r.ApiService.TopUpDepositExecute(r)
}

// TopUpDeposit Put money on deposit
//
// Returns the result of putting money on deposit.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/top-up-deposit/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiTopUpDepositRequest
func (a *PortalPaymentAPIService) TopUpDeposit(ctx context.Context) ApiTopUpDepositRequest {
	return ApiTopUpDepositRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BooleanWrapper
func (a *PortalPaymentAPIService) TopUpDepositExecute(r ApiTopUpDepositRequest) (*BooleanWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BooleanWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.TopUpDeposit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/deposit"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.topUpDepositRequestDto
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdatePaymentRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
	quantityRequestDto *QuantityRequestDto
}

func (r ApiUpdatePaymentRequest) QuantityRequestDto(quantityRequestDto QuantityRequestDto) ApiUpdatePaymentRequest {	r.quantityRequestDto = &quantityRequestDto
	return r
}

func (r ApiUpdatePaymentRequest) Execute() (*BooleanWrapper, *http.Response, error) {
	return r.ApiService.UpdatePaymentExecute(r)
}

// UpdatePayment Update the payment quantity
//
// Updates the payment quantity with the parameters specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/update-payment/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiUpdatePaymentRequest
func (a *PortalPaymentAPIService) UpdatePayment(ctx context.Context) ApiUpdatePaymentRequest {
	return ApiUpdatePaymentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BooleanWrapper
func (a *PortalPaymentAPIService) UpdatePaymentExecute(r ApiUpdatePaymentRequest) (*BooleanWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BooleanWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.UpdatePayment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/update"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.quantityRequestDto
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateWalletPaymentRequest struct {
	ctx context.Context
	ApiService *PortalPaymentAPIService
	walletQuantityRequestDto *WalletQuantityRequestDto
}

func (r ApiUpdateWalletPaymentRequest) WalletQuantityRequestDto(walletQuantityRequestDto WalletQuantityRequestDto) ApiUpdateWalletPaymentRequest {	r.walletQuantityRequestDto = &walletQuantityRequestDto
	return r
}

func (r ApiUpdateWalletPaymentRequest) Execute() (*BooleanWrapper, *http.Response, error) {
	return r.ApiService.UpdateWalletPaymentExecute(r)
}

// UpdateWalletPayment Update the wallet payment quantity
//
// Updates the wallet payment quantity with the parameters specified in the request.
//
// See also: https://api.onlyoffice.com/docspace/api-backend/usage-api/update-wallet-payment/
//
// @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return ApiUpdateWalletPaymentRequest
func (a *PortalPaymentAPIService) UpdateWalletPayment(ctx context.Context) ApiUpdateWalletPaymentRequest {
	return ApiUpdateWalletPaymentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BooleanWrapper
func (a *PortalPaymentAPIService) UpdateWalletPaymentExecute(r ApiUpdateWalletPaymentRequest) (*BooleanWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BooleanWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PortalPaymentAPIService.UpdateWalletPayment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/2.0/portal/payment/updatewallet"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.walletQuantityRequestDto
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyBearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["ApiKeyBearer"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
