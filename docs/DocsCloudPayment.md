# DocsCloudPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CartId** | Pointer to **NullableString** | The cart ID. | [optional] 
**ProductId** | Pointer to **int32** | The product ID. | [optional] 
**Status** | Pointer to **int32** | The payment status. | [optional] 
**IntervalUnit** | Pointer to **int32** | The interval unit. | [optional] 
**IsYear** | Pointer to **bool** | Whether the payment interval is yearly. | [optional] 
**IsPrepaid** | Pointer to **bool** | Whether the payment is prepaid. | [optional] 
**Quantity** | Pointer to **int32** | The quantity. | [optional] 
**Currency** | Pointer to **NullableString** | The three-character ISO 4217 currency symbol of the payment. | [optional] 

## Methods

### NewDocsCloudPayment

`func NewDocsCloudPayment() *DocsCloudPayment`

NewDocsCloudPayment instantiates a new DocsCloudPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudPaymentWithDefaults

`func NewDocsCloudPaymentWithDefaults() *DocsCloudPayment`

NewDocsCloudPaymentWithDefaults instantiates a new DocsCloudPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCartId

`func (o *DocsCloudPayment) GetCartId() string`

GetCartId returns the CartId field if non-nil, zero value otherwise.

### GetCartIdOk

`func (o *DocsCloudPayment) GetCartIdOk() (*string, bool)`

GetCartIdOk returns a tuple with the CartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartId

`func (o *DocsCloudPayment) SetCartId(v string)`

SetCartId sets CartId field to given value.

### HasCartId

`func (o *DocsCloudPayment) HasCartId() bool`

HasCartId returns a boolean if a field has been set.

### SetCartIdNil

`func (o *DocsCloudPayment) SetCartIdNil(b bool)`

 SetCartIdNil sets the value for CartId to be an explicit nil

### UnsetCartId
`func (o *DocsCloudPayment) UnsetCartId()`

UnsetCartId ensures that no value is present for CartId, not even an explicit nil
### GetProductId

`func (o *DocsCloudPayment) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *DocsCloudPayment) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *DocsCloudPayment) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *DocsCloudPayment) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetStatus

`func (o *DocsCloudPayment) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocsCloudPayment) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocsCloudPayment) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DocsCloudPayment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIntervalUnit

`func (o *DocsCloudPayment) GetIntervalUnit() int32`

GetIntervalUnit returns the IntervalUnit field if non-nil, zero value otherwise.

### GetIntervalUnitOk

`func (o *DocsCloudPayment) GetIntervalUnitOk() (*int32, bool)`

GetIntervalUnitOk returns a tuple with the IntervalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalUnit

`func (o *DocsCloudPayment) SetIntervalUnit(v int32)`

SetIntervalUnit sets IntervalUnit field to given value.

### HasIntervalUnit

`func (o *DocsCloudPayment) HasIntervalUnit() bool`

HasIntervalUnit returns a boolean if a field has been set.

### GetIsYear

`func (o *DocsCloudPayment) GetIsYear() bool`

GetIsYear returns the IsYear field if non-nil, zero value otherwise.

### GetIsYearOk

`func (o *DocsCloudPayment) GetIsYearOk() (*bool, bool)`

GetIsYearOk returns a tuple with the IsYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYear

`func (o *DocsCloudPayment) SetIsYear(v bool)`

SetIsYear sets IsYear field to given value.

### HasIsYear

`func (o *DocsCloudPayment) HasIsYear() bool`

HasIsYear returns a boolean if a field has been set.

### GetIsPrepaid

`func (o *DocsCloudPayment) GetIsPrepaid() bool`

GetIsPrepaid returns the IsPrepaid field if non-nil, zero value otherwise.

### GetIsPrepaidOk

`func (o *DocsCloudPayment) GetIsPrepaidOk() (*bool, bool)`

GetIsPrepaidOk returns a tuple with the IsPrepaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrepaid

`func (o *DocsCloudPayment) SetIsPrepaid(v bool)`

SetIsPrepaid sets IsPrepaid field to given value.

### HasIsPrepaid

`func (o *DocsCloudPayment) HasIsPrepaid() bool`

HasIsPrepaid returns a boolean if a field has been set.

### GetQuantity

`func (o *DocsCloudPayment) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *DocsCloudPayment) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *DocsCloudPayment) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *DocsCloudPayment) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetCurrency

`func (o *DocsCloudPayment) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DocsCloudPayment) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DocsCloudPayment) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DocsCloudPayment) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *DocsCloudPayment) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *DocsCloudPayment) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


