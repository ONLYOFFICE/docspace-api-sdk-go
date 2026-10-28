# TransactionInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **time.Time** | The date and time when the credit transaction occurred. | [optional] 
**Currency** | Pointer to **NullableString** | The three-character ISO 4217 currency symbol. | [optional] 
**Amount** | Pointer to **float64** | The amount in the specified currency. | [optional] 

## Methods

### NewTransactionInfoDto

`func NewTransactionInfoDto() *TransactionInfoDto`

NewTransactionInfoDto instantiates a new TransactionInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionInfoDtoWithDefaults

`func NewTransactionInfoDtoWithDefaults() *TransactionInfoDto`

NewTransactionInfoDtoWithDefaults instantiates a new TransactionInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *TransactionInfoDto) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TransactionInfoDto) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TransactionInfoDto) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *TransactionInfoDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetCurrency

`func (o *TransactionInfoDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionInfoDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionInfoDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *TransactionInfoDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *TransactionInfoDto) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *TransactionInfoDto) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetAmount

`func (o *TransactionInfoDto) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionInfoDto) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionInfoDto) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionInfoDto) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


