# RoomsNotificationsSettingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoomsId** | Pointer to **interface{}** |  | [optional] 
**Mute** | Pointer to **bool** | Specifies whether the notifications will be delivered to the specified room or not. | [optional] 

## Methods

### NewRoomsNotificationsSettingsRequestDto

`func NewRoomsNotificationsSettingsRequestDto() *RoomsNotificationsSettingsRequestDto`

NewRoomsNotificationsSettingsRequestDto instantiates a new RoomsNotificationsSettingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomsNotificationsSettingsRequestDtoWithDefaults

`func NewRoomsNotificationsSettingsRequestDtoWithDefaults() *RoomsNotificationsSettingsRequestDto`

NewRoomsNotificationsSettingsRequestDtoWithDefaults instantiates a new RoomsNotificationsSettingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoomsId

`func (o *RoomsNotificationsSettingsRequestDto) GetRoomsId() interface{}`

GetRoomsId returns the RoomsId field if non-nil, zero value otherwise.

### GetRoomsIdOk

`func (o *RoomsNotificationsSettingsRequestDto) GetRoomsIdOk() (*interface{}, bool)`

GetRoomsIdOk returns a tuple with the RoomsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomsId

`func (o *RoomsNotificationsSettingsRequestDto) SetRoomsId(v interface{})`

SetRoomsId sets RoomsId field to given value.

### HasRoomsId

`func (o *RoomsNotificationsSettingsRequestDto) HasRoomsId() bool`

HasRoomsId returns a boolean if a field has been set.

### SetRoomsIdNil

`func (o *RoomsNotificationsSettingsRequestDto) SetRoomsIdNil(b bool)`

 SetRoomsIdNil sets the value for RoomsId to be an explicit nil

### UnsetRoomsId
`func (o *RoomsNotificationsSettingsRequestDto) UnsetRoomsId()`

UnsetRoomsId ensures that no value is present for RoomsId, not even an explicit nil
### GetMute

`func (o *RoomsNotificationsSettingsRequestDto) GetMute() bool`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *RoomsNotificationsSettingsRequestDto) GetMuteOk() (*bool, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *RoomsNotificationsSettingsRequestDto) SetMute(v bool)`

SetMute sets Mute field to given value.

### HasMute

`func (o *RoomsNotificationsSettingsRequestDto) HasMute() bool`

HasMute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


