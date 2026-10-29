# UpdateServiceLimitRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the service limit, as returned in `id` by the operation that created it or by the reads of limits. | 
**AmountValue** | Pointer to **NullableFloat64** | The most money the AI services may cost per period, in their currency. Zero blocks the AI services completely. | [optional] 
**QuantityValue** | Pointer to **NullableInt32** | The largest quantity the AI services may consume per period, in their common unit. Zero blocks the AI services  completely. | [optional] 
**Enabled** | Pointer to **NullableBool** | Whether the limit is in force: false switches it off at once, true switches it back on. | [optional] 

## Methods

### NewUpdateServiceLimitRequestDto

`func NewUpdateServiceLimitRequestDto(id int32, ) *UpdateServiceLimitRequestDto`

NewUpdateServiceLimitRequestDto instantiates a new UpdateServiceLimitRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServiceLimitRequestDtoWithDefaults

`func NewUpdateServiceLimitRequestDtoWithDefaults() *UpdateServiceLimitRequestDto`

NewUpdateServiceLimitRequestDtoWithDefaults instantiates a new UpdateServiceLimitRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateServiceLimitRequestDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateServiceLimitRequestDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateServiceLimitRequestDto) SetId(v int32)`

SetId sets Id field to given value.


### GetAmountValue

`func (o *UpdateServiceLimitRequestDto) GetAmountValue() float64`

GetAmountValue returns the AmountValue field if non-nil, zero value otherwise.

### GetAmountValueOk

`func (o *UpdateServiceLimitRequestDto) GetAmountValueOk() (*float64, bool)`

GetAmountValueOk returns a tuple with the AmountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountValue

`func (o *UpdateServiceLimitRequestDto) SetAmountValue(v float64)`

SetAmountValue sets AmountValue field to given value.

### HasAmountValue

`func (o *UpdateServiceLimitRequestDto) HasAmountValue() bool`

HasAmountValue returns a boolean if a field has been set.

### SetAmountValueNil

`func (o *UpdateServiceLimitRequestDto) SetAmountValueNil(b bool)`

 SetAmountValueNil sets the value for AmountValue to be an explicit nil

### UnsetAmountValue
`func (o *UpdateServiceLimitRequestDto) UnsetAmountValue()`

UnsetAmountValue ensures that no value is present for AmountValue, not even an explicit nil
### GetQuantityValue

`func (o *UpdateServiceLimitRequestDto) GetQuantityValue() int32`

GetQuantityValue returns the QuantityValue field if non-nil, zero value otherwise.

### GetQuantityValueOk

`func (o *UpdateServiceLimitRequestDto) GetQuantityValueOk() (*int32, bool)`

GetQuantityValueOk returns a tuple with the QuantityValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityValue

`func (o *UpdateServiceLimitRequestDto) SetQuantityValue(v int32)`

SetQuantityValue sets QuantityValue field to given value.

### HasQuantityValue

`func (o *UpdateServiceLimitRequestDto) HasQuantityValue() bool`

HasQuantityValue returns a boolean if a field has been set.

### SetQuantityValueNil

`func (o *UpdateServiceLimitRequestDto) SetQuantityValueNil(b bool)`

 SetQuantityValueNil sets the value for QuantityValue to be an explicit nil

### UnsetQuantityValue
`func (o *UpdateServiceLimitRequestDto) UnsetQuantityValue()`

UnsetQuantityValue ensures that no value is present for QuantityValue, not even an explicit nil
### GetEnabled

`func (o *UpdateServiceLimitRequestDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateServiceLimitRequestDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateServiceLimitRequestDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateServiceLimitRequestDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *UpdateServiceLimitRequestDto) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *UpdateServiceLimitRequestDto) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


