# RoomDataLifetimeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletePermanently** | Pointer to **bool** | Decides what happens to a file that has grown too old: it is erased outright, or it is moved to the trash of  the account that created the room, from where it can still be brought back. | [optional] 
**Period** | Pointer to [**RoomDataLifetimePeriod**](RoomDataLifetimePeriod.md) | The unit the age is counted in. Months and years are counted as calendar ones, so the same number of them  covers a different number of days depending on when the clean-up runs. | [optional] 
**Value** | Pointer to **NullableInt32** | How many periods a file may stay in the room, counted from the moment it was last changed rather than from the  moment the rule was set. Files that are already older than this are removed by the next clean-up. | [optional] 
**Enabled** | Pointer to **NullableBool** | Switches the rule on and off. Switching it off erases the rule instead of keeping it aside, so afterwards the  room reports no rule at all and the other three values have to be sent again to bring it back. | [optional] 

## Methods

### NewRoomDataLifetimeDto

`func NewRoomDataLifetimeDto() *RoomDataLifetimeDto`

NewRoomDataLifetimeDto instantiates a new RoomDataLifetimeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomDataLifetimeDtoWithDefaults

`func NewRoomDataLifetimeDtoWithDefaults() *RoomDataLifetimeDto`

NewRoomDataLifetimeDtoWithDefaults instantiates a new RoomDataLifetimeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletePermanently

`func (o *RoomDataLifetimeDto) GetDeletePermanently() bool`

GetDeletePermanently returns the DeletePermanently field if non-nil, zero value otherwise.

### GetDeletePermanentlyOk

`func (o *RoomDataLifetimeDto) GetDeletePermanentlyOk() (*bool, bool)`

GetDeletePermanentlyOk returns a tuple with the DeletePermanently field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletePermanently

`func (o *RoomDataLifetimeDto) SetDeletePermanently(v bool)`

SetDeletePermanently sets DeletePermanently field to given value.

### HasDeletePermanently

`func (o *RoomDataLifetimeDto) HasDeletePermanently() bool`

HasDeletePermanently returns a boolean if a field has been set.

### GetPeriod

`func (o *RoomDataLifetimeDto) GetPeriod() RoomDataLifetimePeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *RoomDataLifetimeDto) GetPeriodOk() (*RoomDataLifetimePeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *RoomDataLifetimeDto) SetPeriod(v RoomDataLifetimePeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *RoomDataLifetimeDto) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetValue

`func (o *RoomDataLifetimeDto) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RoomDataLifetimeDto) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RoomDataLifetimeDto) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *RoomDataLifetimeDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *RoomDataLifetimeDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *RoomDataLifetimeDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetEnabled

`func (o *RoomDataLifetimeDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RoomDataLifetimeDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RoomDataLifetimeDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RoomDataLifetimeDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *RoomDataLifetimeDto) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *RoomDataLifetimeDto) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


