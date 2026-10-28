# AiPriceCurrencyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **NullableString** | The ISO 4217 code of the currency the prices are quoted in. | 
**Symbol** | **NullableString** | The display symbol of the currency. | 

## Methods

### NewAiPriceCurrencyDto

`func NewAiPriceCurrencyDto(code NullableString, symbol NullableString, ) *AiPriceCurrencyDto`

NewAiPriceCurrencyDto instantiates a new AiPriceCurrencyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiPriceCurrencyDtoWithDefaults

`func NewAiPriceCurrencyDtoWithDefaults() *AiPriceCurrencyDto`

NewAiPriceCurrencyDtoWithDefaults instantiates a new AiPriceCurrencyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AiPriceCurrencyDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AiPriceCurrencyDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AiPriceCurrencyDto) SetCode(v string)`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *AiPriceCurrencyDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *AiPriceCurrencyDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetSymbol

`func (o *AiPriceCurrencyDto) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *AiPriceCurrencyDto) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *AiPriceCurrencyDto) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### SetSymbolNil

`func (o *AiPriceCurrencyDto) SetSymbolNil(b bool)`

 SetSymbolNil sets the value for Symbol to be an explicit nil

### UnsetSymbol
`func (o *AiPriceCurrencyDto) UnsetSymbol()`

UnsetSymbol ensures that no value is present for Symbol, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


