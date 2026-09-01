# CustomerServiceUsageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **NullableString** | The name of the service. | [optional] 
**Title** | Pointer to **NullableString** | The title of the service. | [optional] 
**ServiceUnit** | Pointer to **NullableString** | The unit of measurement for the service. | [optional] 
**Currency** | Pointer to **NullableString** | The three-character ISO 4217 currency symbol of the amounts. | [optional] 
**TotalQuantity** | Pointer to **int32** | The total number of units consumed. | [optional] 
**TotalAmount** | Pointer to **float64** | The total amount charged for the service. | [optional] 
**OperationCount** | Pointer to **int32** | The number of individual purchase operations. | [optional] 
**Price** | Pointer to **float64** | The price of the service. | [optional] 
**Subscription** | Pointer to **bool** | Indicates whether the service is subscription-based. | [optional] 

## Methods

### NewCustomerServiceUsageDto

`func NewCustomerServiceUsageDto() *CustomerServiceUsageDto`

NewCustomerServiceUsageDto instantiates a new CustomerServiceUsageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerServiceUsageDtoWithDefaults

`func NewCustomerServiceUsageDtoWithDefaults() *CustomerServiceUsageDto`

NewCustomerServiceUsageDtoWithDefaults instantiates a new CustomerServiceUsageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *CustomerServiceUsageDto) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CustomerServiceUsageDto) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CustomerServiceUsageDto) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CustomerServiceUsageDto) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *CustomerServiceUsageDto) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *CustomerServiceUsageDto) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetTitle

`func (o *CustomerServiceUsageDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomerServiceUsageDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CustomerServiceUsageDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CustomerServiceUsageDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CustomerServiceUsageDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CustomerServiceUsageDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetServiceUnit

`func (o *CustomerServiceUsageDto) GetServiceUnit() string`

GetServiceUnit returns the ServiceUnit field if non-nil, zero value otherwise.

### GetServiceUnitOk

`func (o *CustomerServiceUsageDto) GetServiceUnitOk() (*string, bool)`

GetServiceUnitOk returns a tuple with the ServiceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUnit

`func (o *CustomerServiceUsageDto) SetServiceUnit(v string)`

SetServiceUnit sets ServiceUnit field to given value.

### HasServiceUnit

`func (o *CustomerServiceUsageDto) HasServiceUnit() bool`

HasServiceUnit returns a boolean if a field has been set.

### SetServiceUnitNil

`func (o *CustomerServiceUsageDto) SetServiceUnitNil(b bool)`

 SetServiceUnitNil sets the value for ServiceUnit to be an explicit nil

### UnsetServiceUnit
`func (o *CustomerServiceUsageDto) UnsetServiceUnit()`

UnsetServiceUnit ensures that no value is present for ServiceUnit, not even an explicit nil
### GetCurrency

`func (o *CustomerServiceUsageDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CustomerServiceUsageDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CustomerServiceUsageDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CustomerServiceUsageDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *CustomerServiceUsageDto) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CustomerServiceUsageDto) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetTotalQuantity

`func (o *CustomerServiceUsageDto) GetTotalQuantity() int32`

GetTotalQuantity returns the TotalQuantity field if non-nil, zero value otherwise.

### GetTotalQuantityOk

`func (o *CustomerServiceUsageDto) GetTotalQuantityOk() (*int32, bool)`

GetTotalQuantityOk returns a tuple with the TotalQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuantity

`func (o *CustomerServiceUsageDto) SetTotalQuantity(v int32)`

SetTotalQuantity sets TotalQuantity field to given value.

### HasTotalQuantity

`func (o *CustomerServiceUsageDto) HasTotalQuantity() bool`

HasTotalQuantity returns a boolean if a field has been set.

### GetTotalAmount

`func (o *CustomerServiceUsageDto) GetTotalAmount() float64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *CustomerServiceUsageDto) GetTotalAmountOk() (*float64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *CustomerServiceUsageDto) SetTotalAmount(v float64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *CustomerServiceUsageDto) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetOperationCount

`func (o *CustomerServiceUsageDto) GetOperationCount() int32`

GetOperationCount returns the OperationCount field if non-nil, zero value otherwise.

### GetOperationCountOk

`func (o *CustomerServiceUsageDto) GetOperationCountOk() (*int32, bool)`

GetOperationCountOk returns a tuple with the OperationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationCount

`func (o *CustomerServiceUsageDto) SetOperationCount(v int32)`

SetOperationCount sets OperationCount field to given value.

### HasOperationCount

`func (o *CustomerServiceUsageDto) HasOperationCount() bool`

HasOperationCount returns a boolean if a field has been set.

### GetPrice

`func (o *CustomerServiceUsageDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CustomerServiceUsageDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CustomerServiceUsageDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CustomerServiceUsageDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSubscription

`func (o *CustomerServiceUsageDto) GetSubscription() bool`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *CustomerServiceUsageDto) GetSubscriptionOk() (*bool, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *CustomerServiceUsageDto) SetSubscription(v bool)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *CustomerServiceUsageDto) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


