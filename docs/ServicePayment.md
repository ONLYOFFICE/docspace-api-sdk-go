# ServicePayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationId** | Pointer to **int32** | The payment operation ID. | [optional] 
**Amount** | Pointer to **float64** | The balance of the sub-account in the specified currency. | [optional] 
**Currency** | Pointer to **NullableString** | The three-character ISO 4217 currency symbol. | [optional] 
**Quantity** | Pointer to **int32** | Total quantity of operations. | [optional] 
**SubscriptionId** | Pointer to **NullableInt32** | The subscription ID | [optional] 
**StartDate** | Pointer to **NullableTime** | The subscription start date. | [optional] 
**EndDate** | Pointer to **NullableTime** | The subscription end date. | [optional] 

## Methods

### NewServicePayment

`func NewServicePayment() *ServicePayment`

NewServicePayment instantiates a new ServicePayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePaymentWithDefaults

`func NewServicePaymentWithDefaults() *ServicePayment`

NewServicePaymentWithDefaults instantiates a new ServicePayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationId

`func (o *ServicePayment) GetOperationId() int32`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *ServicePayment) GetOperationIdOk() (*int32, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *ServicePayment) SetOperationId(v int32)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *ServicePayment) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### GetAmount

`func (o *ServicePayment) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ServicePayment) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ServicePayment) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ServicePayment) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *ServicePayment) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ServicePayment) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ServicePayment) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ServicePayment) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *ServicePayment) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ServicePayment) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetQuantity

`func (o *ServicePayment) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ServicePayment) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ServicePayment) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ServicePayment) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *ServicePayment) GetSubscriptionId() int32`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ServicePayment) GetSubscriptionIdOk() (*int32, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ServicePayment) SetSubscriptionId(v int32)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *ServicePayment) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *ServicePayment) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *ServicePayment) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetStartDate

`func (o *ServicePayment) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ServicePayment) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ServicePayment) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ServicePayment) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *ServicePayment) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *ServicePayment) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *ServicePayment) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ServicePayment) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ServicePayment) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ServicePayment) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *ServicePayment) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *ServicePayment) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


