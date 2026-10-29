# SetParticipantServiceLimitsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIds** | **[]string** | The IDs of the portal users the limit is set for, as returned in `id` by the people operations. | 
**AmountValue** | Pointer to **NullableFloat64** | The most money the AI services may cost each user per period, in their currency, which cannot be chosen. Zero  blocks the AI services completely. | [optional] 
**QuantityValue** | Pointer to **NullableInt32** | The largest quantity the AI services may consume per period for each user, in their common unit. Zero blocks the  AI services completely. | [optional] 
**Period** | Pointer to [**ServiceLimitPeriod**](ServiceLimitPeriod.md) | The window the thresholds apply to, in UTC. A day, week or month is calendar-aligned rather than rolling - a  daily limit resets at midnight UTC, a weekly one on Monday - while a lifetime limit counts from its creation and  never resets. | [optional] 

## Methods

### NewSetParticipantServiceLimitsRequestDto

`func NewSetParticipantServiceLimitsRequestDto(userIds []string, ) *SetParticipantServiceLimitsRequestDto`

NewSetParticipantServiceLimitsRequestDto instantiates a new SetParticipantServiceLimitsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetParticipantServiceLimitsRequestDtoWithDefaults

`func NewSetParticipantServiceLimitsRequestDtoWithDefaults() *SetParticipantServiceLimitsRequestDto`

NewSetParticipantServiceLimitsRequestDtoWithDefaults instantiates a new SetParticipantServiceLimitsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIds

`func (o *SetParticipantServiceLimitsRequestDto) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *SetParticipantServiceLimitsRequestDto) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *SetParticipantServiceLimitsRequestDto) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.


### GetAmountValue

`func (o *SetParticipantServiceLimitsRequestDto) GetAmountValue() float64`

GetAmountValue returns the AmountValue field if non-nil, zero value otherwise.

### GetAmountValueOk

`func (o *SetParticipantServiceLimitsRequestDto) GetAmountValueOk() (*float64, bool)`

GetAmountValueOk returns a tuple with the AmountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountValue

`func (o *SetParticipantServiceLimitsRequestDto) SetAmountValue(v float64)`

SetAmountValue sets AmountValue field to given value.

### HasAmountValue

`func (o *SetParticipantServiceLimitsRequestDto) HasAmountValue() bool`

HasAmountValue returns a boolean if a field has been set.

### SetAmountValueNil

`func (o *SetParticipantServiceLimitsRequestDto) SetAmountValueNil(b bool)`

 SetAmountValueNil sets the value for AmountValue to be an explicit nil

### UnsetAmountValue
`func (o *SetParticipantServiceLimitsRequestDto) UnsetAmountValue()`

UnsetAmountValue ensures that no value is present for AmountValue, not even an explicit nil
### GetQuantityValue

`func (o *SetParticipantServiceLimitsRequestDto) GetQuantityValue() int32`

GetQuantityValue returns the QuantityValue field if non-nil, zero value otherwise.

### GetQuantityValueOk

`func (o *SetParticipantServiceLimitsRequestDto) GetQuantityValueOk() (*int32, bool)`

GetQuantityValueOk returns a tuple with the QuantityValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityValue

`func (o *SetParticipantServiceLimitsRequestDto) SetQuantityValue(v int32)`

SetQuantityValue sets QuantityValue field to given value.

### HasQuantityValue

`func (o *SetParticipantServiceLimitsRequestDto) HasQuantityValue() bool`

HasQuantityValue returns a boolean if a field has been set.

### SetQuantityValueNil

`func (o *SetParticipantServiceLimitsRequestDto) SetQuantityValueNil(b bool)`

 SetQuantityValueNil sets the value for QuantityValue to be an explicit nil

### UnsetQuantityValue
`func (o *SetParticipantServiceLimitsRequestDto) UnsetQuantityValue()`

UnsetQuantityValue ensures that no value is present for QuantityValue, not even an explicit nil
### GetPeriod

`func (o *SetParticipantServiceLimitsRequestDto) GetPeriod() ServiceLimitPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *SetParticipantServiceLimitsRequestDto) GetPeriodOk() (*ServiceLimitPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *SetParticipantServiceLimitsRequestDto) SetPeriod(v ServiceLimitPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *SetParticipantServiceLimitsRequestDto) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


