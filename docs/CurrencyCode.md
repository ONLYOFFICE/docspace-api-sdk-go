# CurrencyCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **NullableString** | The three-character ISO 4217 currency symbol. | [optional] 

## Methods

### NewCurrencyCode

`func NewCurrencyCode() *CurrencyCode`

NewCurrencyCode instantiates a new CurrencyCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyCodeWithDefaults

`func NewCurrencyCodeWithDefaults() *CurrencyCode`

NewCurrencyCodeWithDefaults instantiates a new CurrencyCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *CurrencyCode) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CurrencyCode) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CurrencyCode) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CurrencyCode) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *CurrencyCode) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CurrencyCode) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


