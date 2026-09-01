# RoomGroupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The group ID. | [optional] 
**Name** | Pointer to **NullableString** | Group name | [optional] 
**Icon** | Pointer to [**MultiSizeLogoCover**](MultiSizeLogoCover.md) | Group icon | [optional] 
**UserId** | Pointer to **string** | The user ID. | [optional] 
**Rooms** | Pointer to [**[]FileEntryBaseDto**](FileEntryBaseDto.md) | The list of rooms in the group. | [optional] 
**TotalRooms** | Pointer to **int32** | Total number of rooms in the group. | [optional] 

## Methods

### NewRoomGroupDto

`func NewRoomGroupDto() *RoomGroupDto`

NewRoomGroupDto instantiates a new RoomGroupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomGroupDtoWithDefaults

`func NewRoomGroupDtoWithDefaults() *RoomGroupDto`

NewRoomGroupDtoWithDefaults instantiates a new RoomGroupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoomGroupDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoomGroupDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoomGroupDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RoomGroupDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RoomGroupDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoomGroupDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoomGroupDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoomGroupDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RoomGroupDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RoomGroupDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIcon

`func (o *RoomGroupDto) GetIcon() MultiSizeLogoCover`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *RoomGroupDto) GetIconOk() (*MultiSizeLogoCover, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *RoomGroupDto) SetIcon(v MultiSizeLogoCover)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *RoomGroupDto) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetUserId

`func (o *RoomGroupDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RoomGroupDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RoomGroupDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *RoomGroupDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetRooms

`func (o *RoomGroupDto) GetRooms() []FileEntryBaseDto`

GetRooms returns the Rooms field if non-nil, zero value otherwise.

### GetRoomsOk

`func (o *RoomGroupDto) GetRoomsOk() (*[]FileEntryBaseDto, bool)`

GetRoomsOk returns a tuple with the Rooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooms

`func (o *RoomGroupDto) SetRooms(v []FileEntryBaseDto)`

SetRooms sets Rooms field to given value.

### HasRooms

`func (o *RoomGroupDto) HasRooms() bool`

HasRooms returns a boolean if a field has been set.

### SetRoomsNil

`func (o *RoomGroupDto) SetRoomsNil(b bool)`

 SetRoomsNil sets the value for Rooms to be an explicit nil

### UnsetRooms
`func (o *RoomGroupDto) UnsetRooms()`

UnsetRooms ensures that no value is present for Rooms, not even an explicit nil
### GetTotalRooms

`func (o *RoomGroupDto) GetTotalRooms() int32`

GetTotalRooms returns the TotalRooms field if non-nil, zero value otherwise.

### GetTotalRoomsOk

`func (o *RoomGroupDto) GetTotalRoomsOk() (*int32, bool)`

GetTotalRoomsOk returns a tuple with the TotalRooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRooms

`func (o *RoomGroupDto) SetTotalRooms(v int32)`

SetTotalRooms sets TotalRooms field to given value.

### HasTotalRooms

`func (o *RoomGroupDto) HasTotalRooms() bool`

HasTotalRooms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


