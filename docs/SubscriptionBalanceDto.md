# SubscriptionBalanceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCost** | Pointer to **float64** | The total cost of the current billing period (the sum across all subscription items). | [optional] 
**Currency** | Pointer to **NullableString** | The three-character ISO 4217 currency symbol of the subscription. | [optional] 
**PeriodStart** | Pointer to **time.Time** | The start of the current billing period. | [optional] 
**PeriodEnd** | Pointer to **time.Time** | The end of the current billing period. | [optional] 
**PeriodUsedUntil** | Pointer to **time.Time** | The boundary of the used part of the period (the moment of the request). | [optional] 
**DaysElapsed** | Pointer to **int32** | The number of days elapsed since the start of the period (inclusive). | [optional] 
**RemainingBalance** | Pointer to **float64** | The unused balance of the subscription, in the subscription currency. | [optional] 
**RemainingBalanceInWalletCurrency** | Pointer to **float64** | The unused balance of the subscription, converted to the wallet currency. | [optional] 
**WalletCurrency** | Pointer to **NullableString** | The three-character ISO 4217 currency symbol of the wallet. | [optional] 

## Methods

### NewSubscriptionBalanceDto

`func NewSubscriptionBalanceDto() *SubscriptionBalanceDto`

NewSubscriptionBalanceDto instantiates a new SubscriptionBalanceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionBalanceDtoWithDefaults

`func NewSubscriptionBalanceDtoWithDefaults() *SubscriptionBalanceDto`

NewSubscriptionBalanceDtoWithDefaults instantiates a new SubscriptionBalanceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCost

`func (o *SubscriptionBalanceDto) GetTotalCost() float64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *SubscriptionBalanceDto) GetTotalCostOk() (*float64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *SubscriptionBalanceDto) SetTotalCost(v float64)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *SubscriptionBalanceDto) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetCurrency

`func (o *SubscriptionBalanceDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SubscriptionBalanceDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SubscriptionBalanceDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SubscriptionBalanceDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *SubscriptionBalanceDto) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *SubscriptionBalanceDto) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetPeriodStart

`func (o *SubscriptionBalanceDto) GetPeriodStart() time.Time`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *SubscriptionBalanceDto) GetPeriodStartOk() (*time.Time, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *SubscriptionBalanceDto) SetPeriodStart(v time.Time)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *SubscriptionBalanceDto) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *SubscriptionBalanceDto) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *SubscriptionBalanceDto) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *SubscriptionBalanceDto) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *SubscriptionBalanceDto) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodUsedUntil

`func (o *SubscriptionBalanceDto) GetPeriodUsedUntil() time.Time`

GetPeriodUsedUntil returns the PeriodUsedUntil field if non-nil, zero value otherwise.

### GetPeriodUsedUntilOk

`func (o *SubscriptionBalanceDto) GetPeriodUsedUntilOk() (*time.Time, bool)`

GetPeriodUsedUntilOk returns a tuple with the PeriodUsedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodUsedUntil

`func (o *SubscriptionBalanceDto) SetPeriodUsedUntil(v time.Time)`

SetPeriodUsedUntil sets PeriodUsedUntil field to given value.

### HasPeriodUsedUntil

`func (o *SubscriptionBalanceDto) HasPeriodUsedUntil() bool`

HasPeriodUsedUntil returns a boolean if a field has been set.

### GetDaysElapsed

`func (o *SubscriptionBalanceDto) GetDaysElapsed() int32`

GetDaysElapsed returns the DaysElapsed field if non-nil, zero value otherwise.

### GetDaysElapsedOk

`func (o *SubscriptionBalanceDto) GetDaysElapsedOk() (*int32, bool)`

GetDaysElapsedOk returns a tuple with the DaysElapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysElapsed

`func (o *SubscriptionBalanceDto) SetDaysElapsed(v int32)`

SetDaysElapsed sets DaysElapsed field to given value.

### HasDaysElapsed

`func (o *SubscriptionBalanceDto) HasDaysElapsed() bool`

HasDaysElapsed returns a boolean if a field has been set.

### GetRemainingBalance

`func (o *SubscriptionBalanceDto) GetRemainingBalance() float64`

GetRemainingBalance returns the RemainingBalance field if non-nil, zero value otherwise.

### GetRemainingBalanceOk

`func (o *SubscriptionBalanceDto) GetRemainingBalanceOk() (*float64, bool)`

GetRemainingBalanceOk returns a tuple with the RemainingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingBalance

`func (o *SubscriptionBalanceDto) SetRemainingBalance(v float64)`

SetRemainingBalance sets RemainingBalance field to given value.

### HasRemainingBalance

`func (o *SubscriptionBalanceDto) HasRemainingBalance() bool`

HasRemainingBalance returns a boolean if a field has been set.

### GetRemainingBalanceInWalletCurrency

`func (o *SubscriptionBalanceDto) GetRemainingBalanceInWalletCurrency() float64`

GetRemainingBalanceInWalletCurrency returns the RemainingBalanceInWalletCurrency field if non-nil, zero value otherwise.

### GetRemainingBalanceInWalletCurrencyOk

`func (o *SubscriptionBalanceDto) GetRemainingBalanceInWalletCurrencyOk() (*float64, bool)`

GetRemainingBalanceInWalletCurrencyOk returns a tuple with the RemainingBalanceInWalletCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingBalanceInWalletCurrency

`func (o *SubscriptionBalanceDto) SetRemainingBalanceInWalletCurrency(v float64)`

SetRemainingBalanceInWalletCurrency sets RemainingBalanceInWalletCurrency field to given value.

### HasRemainingBalanceInWalletCurrency

`func (o *SubscriptionBalanceDto) HasRemainingBalanceInWalletCurrency() bool`

HasRemainingBalanceInWalletCurrency returns a boolean if a field has been set.

### GetWalletCurrency

`func (o *SubscriptionBalanceDto) GetWalletCurrency() string`

GetWalletCurrency returns the WalletCurrency field if non-nil, zero value otherwise.

### GetWalletCurrencyOk

`func (o *SubscriptionBalanceDto) GetWalletCurrencyOk() (*string, bool)`

GetWalletCurrencyOk returns a tuple with the WalletCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletCurrency

`func (o *SubscriptionBalanceDto) SetWalletCurrency(v string)`

SetWalletCurrency sets WalletCurrency field to given value.

### HasWalletCurrency

`func (o *SubscriptionBalanceDto) HasWalletCurrency() bool`

HasWalletCurrency returns a boolean if a field has been set.

### SetWalletCurrencyNil

`func (o *SubscriptionBalanceDto) SetWalletCurrencyNil(b bool)`

 SetWalletCurrencyNil sets the value for WalletCurrency to be an explicit nil

### UnsetWalletCurrency
`func (o *SubscriptionBalanceDto) UnsetWalletCurrency()`

UnsetWalletCurrency ensures that no value is present for WalletCurrency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


