# TransactionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The three-character ISO 4217 currency symbol. | [optional] 
**Amount** | Pointer to **float64** | The amount in the specified currency. | [optional] 
**Date** | Pointer to **time.Time** | The date and time when the credit transaction occurred. | [optional] 

## Methods

### NewTransactionInfo

`func NewTransactionInfo() *TransactionInfo`

NewTransactionInfo instantiates a new TransactionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionInfoWithDefaults

`func NewTransactionInfoWithDefaults() *TransactionInfo`

NewTransactionInfoWithDefaults instantiates a new TransactionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *TransactionInfo) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionInfo) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionInfo) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *TransactionInfo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionInfo) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionInfo) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionInfo) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionInfo) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDate

`func (o *TransactionInfo) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TransactionInfo) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TransactionInfo) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *TransactionInfo) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


