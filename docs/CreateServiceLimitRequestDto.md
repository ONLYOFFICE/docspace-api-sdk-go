# CreateServiceLimitRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to [**ServiceLimitScope**](ServiceLimitScope.md) | Whom the limit applies to. | [optional] 
**UserId** | Pointer to **NullableString** | The ID of the portal user a personal limit is set for, as returned in `id` by the people operations. | [optional] 
**AmountValue** | Pointer to **NullableFloat64** | The most money the AI services may cost per period, in their currency, which cannot be chosen. Zero blocks the  AI services completely. | [optional] 
**QuantityValue** | Pointer to **NullableInt32** | The largest quantity the AI services may consume per period, in their common unit. Zero blocks the AI services  completely. | [optional] 
**Period** | Pointer to [**ServiceLimitPeriod**](ServiceLimitPeriod.md) | The window the thresholds apply to, in UTC. A day, week or month is calendar-aligned rather than rolling - a  daily limit resets at midnight UTC, a weekly one on Monday - while a lifetime limit counts from its creation and  never resets. | [optional] 

## Methods

### NewCreateServiceLimitRequestDto

`func NewCreateServiceLimitRequestDto() *CreateServiceLimitRequestDto`

NewCreateServiceLimitRequestDto instantiates a new CreateServiceLimitRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceLimitRequestDtoWithDefaults

`func NewCreateServiceLimitRequestDtoWithDefaults() *CreateServiceLimitRequestDto`

NewCreateServiceLimitRequestDtoWithDefaults instantiates a new CreateServiceLimitRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *CreateServiceLimitRequestDto) GetScope() ServiceLimitScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateServiceLimitRequestDto) GetScopeOk() (*ServiceLimitScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateServiceLimitRequestDto) SetScope(v ServiceLimitScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CreateServiceLimitRequestDto) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetUserId

`func (o *CreateServiceLimitRequestDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateServiceLimitRequestDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateServiceLimitRequestDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreateServiceLimitRequestDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *CreateServiceLimitRequestDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *CreateServiceLimitRequestDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetAmountValue

`func (o *CreateServiceLimitRequestDto) GetAmountValue() float64`

GetAmountValue returns the AmountValue field if non-nil, zero value otherwise.

### GetAmountValueOk

`func (o *CreateServiceLimitRequestDto) GetAmountValueOk() (*float64, bool)`

GetAmountValueOk returns a tuple with the AmountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountValue

`func (o *CreateServiceLimitRequestDto) SetAmountValue(v float64)`

SetAmountValue sets AmountValue field to given value.

### HasAmountValue

`func (o *CreateServiceLimitRequestDto) HasAmountValue() bool`

HasAmountValue returns a boolean if a field has been set.

### SetAmountValueNil

`func (o *CreateServiceLimitRequestDto) SetAmountValueNil(b bool)`

 SetAmountValueNil sets the value for AmountValue to be an explicit nil

### UnsetAmountValue
`func (o *CreateServiceLimitRequestDto) UnsetAmountValue()`

UnsetAmountValue ensures that no value is present for AmountValue, not even an explicit nil
### GetQuantityValue

`func (o *CreateServiceLimitRequestDto) GetQuantityValue() int32`

GetQuantityValue returns the QuantityValue field if non-nil, zero value otherwise.

### GetQuantityValueOk

`func (o *CreateServiceLimitRequestDto) GetQuantityValueOk() (*int32, bool)`

GetQuantityValueOk returns a tuple with the QuantityValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityValue

`func (o *CreateServiceLimitRequestDto) SetQuantityValue(v int32)`

SetQuantityValue sets QuantityValue field to given value.

### HasQuantityValue

`func (o *CreateServiceLimitRequestDto) HasQuantityValue() bool`

HasQuantityValue returns a boolean if a field has been set.

### SetQuantityValueNil

`func (o *CreateServiceLimitRequestDto) SetQuantityValueNil(b bool)`

 SetQuantityValueNil sets the value for QuantityValue to be an explicit nil

### UnsetQuantityValue
`func (o *CreateServiceLimitRequestDto) UnsetQuantityValue()`

UnsetQuantityValue ensures that no value is present for QuantityValue, not even an explicit nil
### GetPeriod

`func (o *CreateServiceLimitRequestDto) GetPeriod() ServiceLimitPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CreateServiceLimitRequestDto) GetPeriodOk() (*ServiceLimitPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CreateServiceLimitRequestDto) SetPeriod(v ServiceLimitPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *CreateServiceLimitRequestDto) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


