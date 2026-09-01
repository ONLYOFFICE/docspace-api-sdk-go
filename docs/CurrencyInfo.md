# CurrencyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **NullableString** | The ISO 4217 code of the currency the prices are quoted in. | 
**Symbol** | **NullableString** | The display symbol of the currency. | 

## Methods

### NewCurrencyInfo

`func NewCurrencyInfo(code NullableString, symbol NullableString, ) *CurrencyInfo`

NewCurrencyInfo instantiates a new CurrencyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyInfoWithDefaults

`func NewCurrencyInfoWithDefaults() *CurrencyInfo`

NewCurrencyInfoWithDefaults instantiates a new CurrencyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CurrencyInfo) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CurrencyInfo) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CurrencyInfo) SetCode(v string)`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *CurrencyInfo) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *CurrencyInfo) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetSymbol

`func (o *CurrencyInfo) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CurrencyInfo) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CurrencyInfo) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### SetSymbolNil

`func (o *CurrencyInfo) SetSymbolNil(b bool)`

 SetSymbolNil sets the value for Symbol to be an explicit nil

### UnsetSymbol
`func (o *CurrencyInfo) UnsetSymbol()`

UnsetSymbol ensures that no value is present for Symbol, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


