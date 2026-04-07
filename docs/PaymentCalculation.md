# PaymentCalculation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationId** | Pointer to **int64** | The operation unique identifier. | [optional] 
**Amount** | Pointer to **float64** | The calculated payment amount. | [optional] 
**Currency** | Pointer to **NullableString** | The three-character ISO 4217 currency symbol used for the payment calculation. | [optional] 
**Quantity** | Pointer to **int32** | The quantity associated with the payment calculation. | [optional] 

## Methods

### NewPaymentCalculation

`func NewPaymentCalculation() *PaymentCalculation`

NewPaymentCalculation instantiates a new PaymentCalculation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentCalculationWithDefaults

`func NewPaymentCalculationWithDefaults() *PaymentCalculation`

NewPaymentCalculationWithDefaults instantiates a new PaymentCalculation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationId

`func (o *PaymentCalculation) GetOperationId() int64`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *PaymentCalculation) GetOperationIdOk() (*int64, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *PaymentCalculation) SetOperationId(v int64)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *PaymentCalculation) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentCalculation) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentCalculation) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentCalculation) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentCalculation) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentCalculation) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentCalculation) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentCalculation) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentCalculation) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *PaymentCalculation) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *PaymentCalculation) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetQuantity

`func (o *PaymentCalculation) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PaymentCalculation) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PaymentCalculation) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PaymentCalculation) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


