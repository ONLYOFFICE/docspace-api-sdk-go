# DocsCloudPaymentDto

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

### NewDocsCloudPaymentDto

`func NewDocsCloudPaymentDto() *DocsCloudPaymentDto`

NewDocsCloudPaymentDto instantiates a new DocsCloudPaymentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudPaymentDtoWithDefaults

`func NewDocsCloudPaymentDtoWithDefaults() *DocsCloudPaymentDto`

NewDocsCloudPaymentDtoWithDefaults instantiates a new DocsCloudPaymentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCartId

`func (o *DocsCloudPaymentDto) GetCartId() string`

GetCartId returns the CartId field if non-nil, zero value otherwise.

### GetCartIdOk

`func (o *DocsCloudPaymentDto) GetCartIdOk() (*string, bool)`

GetCartIdOk returns a tuple with the CartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartId

`func (o *DocsCloudPaymentDto) SetCartId(v string)`

SetCartId sets CartId field to given value.

### HasCartId

`func (o *DocsCloudPaymentDto) HasCartId() bool`

HasCartId returns a boolean if a field has been set.

### SetCartIdNil

`func (o *DocsCloudPaymentDto) SetCartIdNil(b bool)`

 SetCartIdNil sets the value for CartId to be an explicit nil

### UnsetCartId
`func (o *DocsCloudPaymentDto) UnsetCartId()`

UnsetCartId ensures that no value is present for CartId, not even an explicit nil
### GetProductId

`func (o *DocsCloudPaymentDto) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *DocsCloudPaymentDto) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *DocsCloudPaymentDto) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *DocsCloudPaymentDto) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetStatus

`func (o *DocsCloudPaymentDto) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocsCloudPaymentDto) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocsCloudPaymentDto) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DocsCloudPaymentDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIntervalUnit

`func (o *DocsCloudPaymentDto) GetIntervalUnit() int32`

GetIntervalUnit returns the IntervalUnit field if non-nil, zero value otherwise.

### GetIntervalUnitOk

`func (o *DocsCloudPaymentDto) GetIntervalUnitOk() (*int32, bool)`

GetIntervalUnitOk returns a tuple with the IntervalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalUnit

`func (o *DocsCloudPaymentDto) SetIntervalUnit(v int32)`

SetIntervalUnit sets IntervalUnit field to given value.

### HasIntervalUnit

`func (o *DocsCloudPaymentDto) HasIntervalUnit() bool`

HasIntervalUnit returns a boolean if a field has been set.

### GetIsYear

`func (o *DocsCloudPaymentDto) GetIsYear() bool`

GetIsYear returns the IsYear field if non-nil, zero value otherwise.

### GetIsYearOk

`func (o *DocsCloudPaymentDto) GetIsYearOk() (*bool, bool)`

GetIsYearOk returns a tuple with the IsYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYear

`func (o *DocsCloudPaymentDto) SetIsYear(v bool)`

SetIsYear sets IsYear field to given value.

### HasIsYear

`func (o *DocsCloudPaymentDto) HasIsYear() bool`

HasIsYear returns a boolean if a field has been set.

### GetIsPrepaid

`func (o *DocsCloudPaymentDto) GetIsPrepaid() bool`

GetIsPrepaid returns the IsPrepaid field if non-nil, zero value otherwise.

### GetIsPrepaidOk

`func (o *DocsCloudPaymentDto) GetIsPrepaidOk() (*bool, bool)`

GetIsPrepaidOk returns a tuple with the IsPrepaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrepaid

`func (o *DocsCloudPaymentDto) SetIsPrepaid(v bool)`

SetIsPrepaid sets IsPrepaid field to given value.

### HasIsPrepaid

`func (o *DocsCloudPaymentDto) HasIsPrepaid() bool`

HasIsPrepaid returns a boolean if a field has been set.

### GetQuantity

`func (o *DocsCloudPaymentDto) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *DocsCloudPaymentDto) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *DocsCloudPaymentDto) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *DocsCloudPaymentDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetCurrency

`func (o *DocsCloudPaymentDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DocsCloudPaymentDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DocsCloudPaymentDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DocsCloudPaymentDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *DocsCloudPaymentDto) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *DocsCloudPaymentDto) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


