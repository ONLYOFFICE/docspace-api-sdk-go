# ServiceLimitDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The service limit ID, used to read, change, switch off or delete the limit. In the list of the users the same  ID repeats for every user the limit for every user applies to. | [optional] 
**ServiceGroupName** | Pointer to **NullableString** | The limited service group; the AI services are `ai-services`. | [optional] 
**CustomerAccountNumber** | Pointer to **int32** | The number of the billing account of the portal the limit belongs to. | [optional] 
**Scope** | Pointer to [**ServiceLimitScope**](ServiceLimitScope.md) | Whom the limit applies to. | [optional] 
**UserId** | Pointer to **NullableString** | The ID of the portal user the limit applies to: the user of a personal limit, or, in the list of the users, the  user a limit for every user applies to. Null for a limit of the portal as a whole, and for a limit for every  user read on its own. | [optional] 
**AmountValue** | Pointer to **NullableFloat64** | The most money the AI services may cost per period, in `currency`; null when the limit has no money threshold. | [optional] 
**Currency** | Pointer to **NullableString** | The currency `amountValue` is expressed in; null when the limit has no money threshold. | [optional] 
**QuantityValue** | Pointer to **NullableInt32** | The largest quantity the AI services may consume per period, in `serviceUnit`; null when the limit has no  quantity threshold. | [optional] 
**ServiceUnit** | Pointer to **NullableString** | The unit `quantityValue` is expressed in; null when the limit has no quantity threshold. | [optional] 
**Period** | Pointer to [**ServiceLimitPeriod**](ServiceLimitPeriod.md) | The window the thresholds apply to. A day, week or month resets at the start of the next one in UTC (a week on  Monday), not a fixed time after the first charge; a lifetime limit never resets. | [optional] 
**PeriodEnd** | Pointer to **NullableTime** | The last moment of the current period, in UTC, after which the spending starts over from zero; null for a  lifetime limit, which never resets. Filled for every limit, the limit for every user included, whose spending is  counted per user. | [optional] 
**UsageFrom** | Pointer to **NullableTime** | The moment a lifetime limit counts from, in UTC; null for the other periods. | [optional] 
**Enabled** | Pointer to **bool** | Whether the limit is in force. A switched-off limit no longer restricts the AI services and can be switched  back on. | [optional] 
**Created** | Pointer to **time.Time** | The date and time when the limit was created, in UTC. | [optional] 
**Modified** | Pointer to **time.Time** | The date and time when the limit was last changed, in UTC. | [optional] 
**Usage** | Pointer to [**ServiceLimitUsageDto**](ServiceLimitUsageDto.md) | What was spent against the limit in its current period and what is left. Filled for the limit of the portal as  a whole and in the list of the users; null where the spending is not read, such as for a limit read by its ID  or the limit for every user, whose spending is counted per user. | [optional] 

## Methods

### NewServiceLimitDto

`func NewServiceLimitDto() *ServiceLimitDto`

NewServiceLimitDto instantiates a new ServiceLimitDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceLimitDtoWithDefaults

`func NewServiceLimitDtoWithDefaults() *ServiceLimitDto`

NewServiceLimitDtoWithDefaults instantiates a new ServiceLimitDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceLimitDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceLimitDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceLimitDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceLimitDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetServiceGroupName

`func (o *ServiceLimitDto) GetServiceGroupName() string`

GetServiceGroupName returns the ServiceGroupName field if non-nil, zero value otherwise.

### GetServiceGroupNameOk

`func (o *ServiceLimitDto) GetServiceGroupNameOk() (*string, bool)`

GetServiceGroupNameOk returns a tuple with the ServiceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceGroupName

`func (o *ServiceLimitDto) SetServiceGroupName(v string)`

SetServiceGroupName sets ServiceGroupName field to given value.

### HasServiceGroupName

`func (o *ServiceLimitDto) HasServiceGroupName() bool`

HasServiceGroupName returns a boolean if a field has been set.

### SetServiceGroupNameNil

`func (o *ServiceLimitDto) SetServiceGroupNameNil(b bool)`

 SetServiceGroupNameNil sets the value for ServiceGroupName to be an explicit nil

### UnsetServiceGroupName
`func (o *ServiceLimitDto) UnsetServiceGroupName()`

UnsetServiceGroupName ensures that no value is present for ServiceGroupName, not even an explicit nil
### GetCustomerAccountNumber

`func (o *ServiceLimitDto) GetCustomerAccountNumber() int32`

GetCustomerAccountNumber returns the CustomerAccountNumber field if non-nil, zero value otherwise.

### GetCustomerAccountNumberOk

`func (o *ServiceLimitDto) GetCustomerAccountNumberOk() (*int32, bool)`

GetCustomerAccountNumberOk returns a tuple with the CustomerAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAccountNumber

`func (o *ServiceLimitDto) SetCustomerAccountNumber(v int32)`

SetCustomerAccountNumber sets CustomerAccountNumber field to given value.

### HasCustomerAccountNumber

`func (o *ServiceLimitDto) HasCustomerAccountNumber() bool`

HasCustomerAccountNumber returns a boolean if a field has been set.

### GetScope

`func (o *ServiceLimitDto) GetScope() ServiceLimitScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ServiceLimitDto) GetScopeOk() (*ServiceLimitScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ServiceLimitDto) SetScope(v ServiceLimitScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ServiceLimitDto) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetUserId

`func (o *ServiceLimitDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ServiceLimitDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ServiceLimitDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ServiceLimitDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ServiceLimitDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ServiceLimitDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetAmountValue

`func (o *ServiceLimitDto) GetAmountValue() float64`

GetAmountValue returns the AmountValue field if non-nil, zero value otherwise.

### GetAmountValueOk

`func (o *ServiceLimitDto) GetAmountValueOk() (*float64, bool)`

GetAmountValueOk returns a tuple with the AmountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountValue

`func (o *ServiceLimitDto) SetAmountValue(v float64)`

SetAmountValue sets AmountValue field to given value.

### HasAmountValue

`func (o *ServiceLimitDto) HasAmountValue() bool`

HasAmountValue returns a boolean if a field has been set.

### SetAmountValueNil

`func (o *ServiceLimitDto) SetAmountValueNil(b bool)`

 SetAmountValueNil sets the value for AmountValue to be an explicit nil

### UnsetAmountValue
`func (o *ServiceLimitDto) UnsetAmountValue()`

UnsetAmountValue ensures that no value is present for AmountValue, not even an explicit nil
### GetCurrency

`func (o *ServiceLimitDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ServiceLimitDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ServiceLimitDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ServiceLimitDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *ServiceLimitDto) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ServiceLimitDto) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetQuantityValue

`func (o *ServiceLimitDto) GetQuantityValue() int32`

GetQuantityValue returns the QuantityValue field if non-nil, zero value otherwise.

### GetQuantityValueOk

`func (o *ServiceLimitDto) GetQuantityValueOk() (*int32, bool)`

GetQuantityValueOk returns a tuple with the QuantityValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityValue

`func (o *ServiceLimitDto) SetQuantityValue(v int32)`

SetQuantityValue sets QuantityValue field to given value.

### HasQuantityValue

`func (o *ServiceLimitDto) HasQuantityValue() bool`

HasQuantityValue returns a boolean if a field has been set.

### SetQuantityValueNil

`func (o *ServiceLimitDto) SetQuantityValueNil(b bool)`

 SetQuantityValueNil sets the value for QuantityValue to be an explicit nil

### UnsetQuantityValue
`func (o *ServiceLimitDto) UnsetQuantityValue()`

UnsetQuantityValue ensures that no value is present for QuantityValue, not even an explicit nil
### GetServiceUnit

`func (o *ServiceLimitDto) GetServiceUnit() string`

GetServiceUnit returns the ServiceUnit field if non-nil, zero value otherwise.

### GetServiceUnitOk

`func (o *ServiceLimitDto) GetServiceUnitOk() (*string, bool)`

GetServiceUnitOk returns a tuple with the ServiceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUnit

`func (o *ServiceLimitDto) SetServiceUnit(v string)`

SetServiceUnit sets ServiceUnit field to given value.

### HasServiceUnit

`func (o *ServiceLimitDto) HasServiceUnit() bool`

HasServiceUnit returns a boolean if a field has been set.

### SetServiceUnitNil

`func (o *ServiceLimitDto) SetServiceUnitNil(b bool)`

 SetServiceUnitNil sets the value for ServiceUnit to be an explicit nil

### UnsetServiceUnit
`func (o *ServiceLimitDto) UnsetServiceUnit()`

UnsetServiceUnit ensures that no value is present for ServiceUnit, not even an explicit nil
### GetPeriod

`func (o *ServiceLimitDto) GetPeriod() ServiceLimitPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ServiceLimitDto) GetPeriodOk() (*ServiceLimitPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ServiceLimitDto) SetPeriod(v ServiceLimitPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *ServiceLimitDto) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *ServiceLimitDto) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *ServiceLimitDto) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *ServiceLimitDto) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *ServiceLimitDto) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### SetPeriodEndNil

`func (o *ServiceLimitDto) SetPeriodEndNil(b bool)`

 SetPeriodEndNil sets the value for PeriodEnd to be an explicit nil

### UnsetPeriodEnd
`func (o *ServiceLimitDto) UnsetPeriodEnd()`

UnsetPeriodEnd ensures that no value is present for PeriodEnd, not even an explicit nil
### GetUsageFrom

`func (o *ServiceLimitDto) GetUsageFrom() time.Time`

GetUsageFrom returns the UsageFrom field if non-nil, zero value otherwise.

### GetUsageFromOk

`func (o *ServiceLimitDto) GetUsageFromOk() (*time.Time, bool)`

GetUsageFromOk returns a tuple with the UsageFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageFrom

`func (o *ServiceLimitDto) SetUsageFrom(v time.Time)`

SetUsageFrom sets UsageFrom field to given value.

### HasUsageFrom

`func (o *ServiceLimitDto) HasUsageFrom() bool`

HasUsageFrom returns a boolean if a field has been set.

### SetUsageFromNil

`func (o *ServiceLimitDto) SetUsageFromNil(b bool)`

 SetUsageFromNil sets the value for UsageFrom to be an explicit nil

### UnsetUsageFrom
`func (o *ServiceLimitDto) UnsetUsageFrom()`

UnsetUsageFrom ensures that no value is present for UsageFrom, not even an explicit nil
### GetEnabled

`func (o *ServiceLimitDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServiceLimitDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServiceLimitDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServiceLimitDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreated

`func (o *ServiceLimitDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ServiceLimitDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ServiceLimitDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ServiceLimitDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *ServiceLimitDto) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ServiceLimitDto) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ServiceLimitDto) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ServiceLimitDto) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetUsage

`func (o *ServiceLimitDto) GetUsage() ServiceLimitUsageDto`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ServiceLimitDto) GetUsageOk() (*ServiceLimitUsageDto, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ServiceLimitDto) SetUsage(v ServiceLimitUsageDto)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ServiceLimitDto) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


