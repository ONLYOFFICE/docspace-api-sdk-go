# PriceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **NullableFloat64** | The price value. | [optional] 
**CurrencySymbol** | Pointer to **NullableString** | The currency symbol. | [optional] 
**IsoCurrencySymbol** | Pointer to **NullableString** | The three-character ISO 4217 currency symbol. | [optional] 

## Methods

### NewPriceDto

`func NewPriceDto() *PriceDto`

NewPriceDto instantiates a new PriceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceDtoWithDefaults

`func NewPriceDtoWithDefaults() *PriceDto`

NewPriceDtoWithDefaults instantiates a new PriceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *PriceDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PriceDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PriceDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *PriceDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *PriceDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PriceDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetCurrencySymbol

`func (o *PriceDto) GetCurrencySymbol() string`

GetCurrencySymbol returns the CurrencySymbol field if non-nil, zero value otherwise.

### GetCurrencySymbolOk

`func (o *PriceDto) GetCurrencySymbolOk() (*string, bool)`

GetCurrencySymbolOk returns a tuple with the CurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencySymbol

`func (o *PriceDto) SetCurrencySymbol(v string)`

SetCurrencySymbol sets CurrencySymbol field to given value.

### HasCurrencySymbol

`func (o *PriceDto) HasCurrencySymbol() bool`

HasCurrencySymbol returns a boolean if a field has been set.

### SetCurrencySymbolNil

`func (o *PriceDto) SetCurrencySymbolNil(b bool)`

 SetCurrencySymbolNil sets the value for CurrencySymbol to be an explicit nil

### UnsetCurrencySymbol
`func (o *PriceDto) UnsetCurrencySymbol()`

UnsetCurrencySymbol ensures that no value is present for CurrencySymbol, not even an explicit nil
### GetIsoCurrencySymbol

`func (o *PriceDto) GetIsoCurrencySymbol() string`

GetIsoCurrencySymbol returns the IsoCurrencySymbol field if non-nil, zero value otherwise.

### GetIsoCurrencySymbolOk

`func (o *PriceDto) GetIsoCurrencySymbolOk() (*string, bool)`

GetIsoCurrencySymbolOk returns a tuple with the IsoCurrencySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencySymbol

`func (o *PriceDto) SetIsoCurrencySymbol(v string)`

SetIsoCurrencySymbol sets IsoCurrencySymbol field to given value.

### HasIsoCurrencySymbol

`func (o *PriceDto) HasIsoCurrencySymbol() bool`

HasIsoCurrencySymbol returns a boolean if a field has been set.

### SetIsoCurrencySymbolNil

`func (o *PriceDto) SetIsoCurrencySymbolNil(b bool)`

 SetIsoCurrencySymbolNil sets the value for IsoCurrencySymbol to be an explicit nil

### UnsetIsoCurrencySymbol
`func (o *PriceDto) UnsetIsoCurrencySymbol()`

UnsetIsoCurrencySymbol ensures that no value is present for IsoCurrencySymbol, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


