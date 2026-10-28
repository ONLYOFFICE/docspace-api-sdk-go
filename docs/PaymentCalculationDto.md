# PaymentCalculationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationId** | Pointer to **int64** | The operation unique identifier. | [optional] 
**Amount** | Pointer to **float64** | The calculated payment amount. | [optional] 
**Currency** | Pointer to **NullableString** | The three-character ISO 4217 currency symbol used for the payment calculation. | [optional] 
**Quantity** | Pointer to **int32** | The quantity associated with the payment calculation. | [optional] 

## Methods

### NewPaymentCalculationDto

`func NewPaymentCalculationDto() *PaymentCalculationDto`

NewPaymentCalculationDto instantiates a new PaymentCalculationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentCalculationDtoWithDefaults

`func NewPaymentCalculationDtoWithDefaults() *PaymentCalculationDto`

NewPaymentCalculationDtoWithDefaults instantiates a new PaymentCalculationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationId

`func (o *PaymentCalculationDto) GetOperationId() int64`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *PaymentCalculationDto) GetOperationIdOk() (*int64, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *PaymentCalculationDto) SetOperationId(v int64)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *PaymentCalculationDto) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentCalculationDto) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentCalculationDto) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentCalculationDto) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentCalculationDto) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentCalculationDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentCalculationDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentCalculationDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentCalculationDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *PaymentCalculationDto) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *PaymentCalculationDto) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetQuantity

`func (o *PaymentCalculationDto) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PaymentCalculationDto) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PaymentCalculationDto) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PaymentCalculationDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


