# CreditAiBalanceRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float64** | The amount to transfer from the main balance to the AI subaccount. | [optional] 
**Currency** | Pointer to **NullableString** | The three-character ISO 4217 currency symbol. | [optional] 

## Methods

### NewCreditAiBalanceRequestDto

`func NewCreditAiBalanceRequestDto() *CreditAiBalanceRequestDto`

NewCreditAiBalanceRequestDto instantiates a new CreditAiBalanceRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditAiBalanceRequestDtoWithDefaults

`func NewCreditAiBalanceRequestDtoWithDefaults() *CreditAiBalanceRequestDto`

NewCreditAiBalanceRequestDtoWithDefaults instantiates a new CreditAiBalanceRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CreditAiBalanceRequestDto) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreditAiBalanceRequestDto) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreditAiBalanceRequestDto) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CreditAiBalanceRequestDto) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *CreditAiBalanceRequestDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreditAiBalanceRequestDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreditAiBalanceRequestDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreditAiBalanceRequestDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *CreditAiBalanceRequestDto) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CreditAiBalanceRequestDto) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


