# CurrencyAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The three-character ISO 4217 currency symbol. | [optional] 
**Amount** | Pointer to **float64** | The amount in the specified currency. | [optional] 

## Methods

### NewCurrencyAmount

`func NewCurrencyAmount() *CurrencyAmount`

NewCurrencyAmount instantiates a new CurrencyAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyAmountWithDefaults

`func NewCurrencyAmountWithDefaults() *CurrencyAmount`

NewCurrencyAmountWithDefaults instantiates a new CurrencyAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *CurrencyAmount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CurrencyAmount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CurrencyAmount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CurrencyAmount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAmount

`func (o *CurrencyAmount) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CurrencyAmount) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CurrencyAmount) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CurrencyAmount) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


