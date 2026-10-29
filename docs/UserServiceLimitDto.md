# UserServiceLimitDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The service limit ID, used to read, change, switch off or delete the limit through the  `api/2.0/portal/payment/servicelimit` operations. For the limit for every user it is the same for all users. | [optional] 
**Scope** | Pointer to [**ServiceLimitScope**](ServiceLimitScope.md) | Whether the limit is the user's own (`Custom`) or the limit for every user (`Global`) the user falls back to. | [optional] 
**AmountValue** | Pointer to **NullableFloat64** | The most money the AI services may cost the user per period, in `currency`. It is the only threshold a limit on  the AI services has: their services count in different units, so a threshold on a quantity is refused for them. | [optional] 
**AmountConsumed** | Pointer to **float64** | The money the AI services cost the user in the current period, in `currency`; 0 when nothing was spent yet,  including a user who has never used the AI services. | [optional] 
**Currency** | Pointer to **NullableString** | The currency `amountValue` and `amountConsumed` are expressed in; null when the limit has no money threshold. | [optional] 
**Period** | Pointer to [**ServiceLimitPeriod**](ServiceLimitPeriod.md) | The window the threshold applies to. It resets at the start of the next day, week or month in UTC (a week on  Monday), not a fixed time after the first charge; a lifetime limit never resets. | [optional] 
**PeriodEnd** | Pointer to **NullableTime** | The moment the current period ends and the spending starts over from zero, in UTC. Null for a lifetime limit,  which never resets. | [optional] 

## Methods

### NewUserServiceLimitDto

`func NewUserServiceLimitDto() *UserServiceLimitDto`

NewUserServiceLimitDto instantiates a new UserServiceLimitDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserServiceLimitDtoWithDefaults

`func NewUserServiceLimitDtoWithDefaults() *UserServiceLimitDto`

NewUserServiceLimitDtoWithDefaults instantiates a new UserServiceLimitDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserServiceLimitDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserServiceLimitDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserServiceLimitDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserServiceLimitDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScope

`func (o *UserServiceLimitDto) GetScope() ServiceLimitScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UserServiceLimitDto) GetScopeOk() (*ServiceLimitScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UserServiceLimitDto) SetScope(v ServiceLimitScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *UserServiceLimitDto) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetAmountValue

`func (o *UserServiceLimitDto) GetAmountValue() float64`

GetAmountValue returns the AmountValue field if non-nil, zero value otherwise.

### GetAmountValueOk

`func (o *UserServiceLimitDto) GetAmountValueOk() (*float64, bool)`

GetAmountValueOk returns a tuple with the AmountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountValue

`func (o *UserServiceLimitDto) SetAmountValue(v float64)`

SetAmountValue sets AmountValue field to given value.

### HasAmountValue

`func (o *UserServiceLimitDto) HasAmountValue() bool`

HasAmountValue returns a boolean if a field has been set.

### SetAmountValueNil

`func (o *UserServiceLimitDto) SetAmountValueNil(b bool)`

 SetAmountValueNil sets the value for AmountValue to be an explicit nil

### UnsetAmountValue
`func (o *UserServiceLimitDto) UnsetAmountValue()`

UnsetAmountValue ensures that no value is present for AmountValue, not even an explicit nil
### GetAmountConsumed

`func (o *UserServiceLimitDto) GetAmountConsumed() float64`

GetAmountConsumed returns the AmountConsumed field if non-nil, zero value otherwise.

### GetAmountConsumedOk

`func (o *UserServiceLimitDto) GetAmountConsumedOk() (*float64, bool)`

GetAmountConsumedOk returns a tuple with the AmountConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountConsumed

`func (o *UserServiceLimitDto) SetAmountConsumed(v float64)`

SetAmountConsumed sets AmountConsumed field to given value.

### HasAmountConsumed

`func (o *UserServiceLimitDto) HasAmountConsumed() bool`

HasAmountConsumed returns a boolean if a field has been set.

### GetCurrency

`func (o *UserServiceLimitDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UserServiceLimitDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UserServiceLimitDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UserServiceLimitDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *UserServiceLimitDto) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *UserServiceLimitDto) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetPeriod

`func (o *UserServiceLimitDto) GetPeriod() ServiceLimitPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *UserServiceLimitDto) GetPeriodOk() (*ServiceLimitPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *UserServiceLimitDto) SetPeriod(v ServiceLimitPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *UserServiceLimitDto) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *UserServiceLimitDto) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *UserServiceLimitDto) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *UserServiceLimitDto) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *UserServiceLimitDto) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### SetPeriodEndNil

`func (o *UserServiceLimitDto) SetPeriodEndNil(b bool)`

 SetPeriodEndNil sets the value for PeriodEnd to be an explicit nil

### UnsetPeriodEnd
`func (o *UserServiceLimitDto) UnsetPeriodEnd()`

UnsetPeriodEnd ensures that no value is present for PeriodEnd, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


