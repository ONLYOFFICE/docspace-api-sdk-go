# AiRoomDataLifetimeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletePermanently** | Pointer to **bool** | Specifies whether to permanently delete the room data or not. | [optional] 
**Period** | Pointer to [**AiRoomDataLifetimePeriod**](AiRoomDataLifetimePeriod.md) | Specifies the time period type of the room data lifetime. | [optional] 
**Value** | Pointer to **NullableInt32** | Specifies the time period value of the room data lifetime. | [optional] 
**Enabled** | Pointer to **NullableBool** | Specifies whether the room data lifetime setting is enabled or not. | [optional] 

## Methods

### NewAiRoomDataLifetimeDto

`func NewAiRoomDataLifetimeDto() *AiRoomDataLifetimeDto`

NewAiRoomDataLifetimeDto instantiates a new AiRoomDataLifetimeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiRoomDataLifetimeDtoWithDefaults

`func NewAiRoomDataLifetimeDtoWithDefaults() *AiRoomDataLifetimeDto`

NewAiRoomDataLifetimeDtoWithDefaults instantiates a new AiRoomDataLifetimeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletePermanently

`func (o *AiRoomDataLifetimeDto) GetDeletePermanently() bool`

GetDeletePermanently returns the DeletePermanently field if non-nil, zero value otherwise.

### GetDeletePermanentlyOk

`func (o *AiRoomDataLifetimeDto) GetDeletePermanentlyOk() (*bool, bool)`

GetDeletePermanentlyOk returns a tuple with the DeletePermanently field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletePermanently

`func (o *AiRoomDataLifetimeDto) SetDeletePermanently(v bool)`

SetDeletePermanently sets DeletePermanently field to given value.

### HasDeletePermanently

`func (o *AiRoomDataLifetimeDto) HasDeletePermanently() bool`

HasDeletePermanently returns a boolean if a field has been set.

### GetPeriod

`func (o *AiRoomDataLifetimeDto) GetPeriod() AiRoomDataLifetimePeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *AiRoomDataLifetimeDto) GetPeriodOk() (*AiRoomDataLifetimePeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *AiRoomDataLifetimeDto) SetPeriod(v AiRoomDataLifetimePeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *AiRoomDataLifetimeDto) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetValue

`func (o *AiRoomDataLifetimeDto) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AiRoomDataLifetimeDto) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AiRoomDataLifetimeDto) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *AiRoomDataLifetimeDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *AiRoomDataLifetimeDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *AiRoomDataLifetimeDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetEnabled

`func (o *AiRoomDataLifetimeDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AiRoomDataLifetimeDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AiRoomDataLifetimeDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AiRoomDataLifetimeDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *AiRoomDataLifetimeDto) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *AiRoomDataLifetimeDto) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


