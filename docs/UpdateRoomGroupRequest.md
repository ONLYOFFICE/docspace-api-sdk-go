# UpdateRoomGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoomsToAdd** | Pointer to [**[]DuplicateRequestDtoAllOfFileIds**](DuplicateRequestDtoAllOfFileIds.md) | The rooms to attach to the group, each given as a number for a room stored in the portal or as a string for a  room on a connected third-party account. Every identifier has to name a room the caller can read; repeats and  rooms the group already holds are collapsed rather than refused. | [optional] 
**RoomsToRemove** | Pointer to [**[]DuplicateRequestDtoAllOfFileIds**](DuplicateRequestDtoAllOfFileIds.md) | The rooms to detach from the group, in the same two forms. Detaching leaves the room and its content  untouched, and a room the group already holds can be detached even when the caller has lost access to it in  the meantime. | [optional] 
**GroupName** | Pointer to **NullableString** | The new name of the group, trimmed of surrounding spaces before it is stored. Leaving the member out keeps the  current name, and a name that is blank once trimmed is refused. | [optional] 

## Methods

### NewUpdateRoomGroupRequest

`func NewUpdateRoomGroupRequest() *UpdateRoomGroupRequest`

NewUpdateRoomGroupRequest instantiates a new UpdateRoomGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoomGroupRequestWithDefaults

`func NewUpdateRoomGroupRequestWithDefaults() *UpdateRoomGroupRequest`

NewUpdateRoomGroupRequestWithDefaults instantiates a new UpdateRoomGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoomsToAdd

`func (o *UpdateRoomGroupRequest) GetRoomsToAdd() []DuplicateRequestDtoAllOfFileIds`

GetRoomsToAdd returns the RoomsToAdd field if non-nil, zero value otherwise.

### GetRoomsToAddOk

`func (o *UpdateRoomGroupRequest) GetRoomsToAddOk() (*[]DuplicateRequestDtoAllOfFileIds, bool)`

GetRoomsToAddOk returns a tuple with the RoomsToAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomsToAdd

`func (o *UpdateRoomGroupRequest) SetRoomsToAdd(v []DuplicateRequestDtoAllOfFileIds)`

SetRoomsToAdd sets RoomsToAdd field to given value.

### HasRoomsToAdd

`func (o *UpdateRoomGroupRequest) HasRoomsToAdd() bool`

HasRoomsToAdd returns a boolean if a field has been set.

### SetRoomsToAddNil

`func (o *UpdateRoomGroupRequest) SetRoomsToAddNil(b bool)`

 SetRoomsToAddNil sets the value for RoomsToAdd to be an explicit nil

### UnsetRoomsToAdd
`func (o *UpdateRoomGroupRequest) UnsetRoomsToAdd()`

UnsetRoomsToAdd ensures that no value is present for RoomsToAdd, not even an explicit nil
### GetRoomsToRemove

`func (o *UpdateRoomGroupRequest) GetRoomsToRemove() []DuplicateRequestDtoAllOfFileIds`

GetRoomsToRemove returns the RoomsToRemove field if non-nil, zero value otherwise.

### GetRoomsToRemoveOk

`func (o *UpdateRoomGroupRequest) GetRoomsToRemoveOk() (*[]DuplicateRequestDtoAllOfFileIds, bool)`

GetRoomsToRemoveOk returns a tuple with the RoomsToRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomsToRemove

`func (o *UpdateRoomGroupRequest) SetRoomsToRemove(v []DuplicateRequestDtoAllOfFileIds)`

SetRoomsToRemove sets RoomsToRemove field to given value.

### HasRoomsToRemove

`func (o *UpdateRoomGroupRequest) HasRoomsToRemove() bool`

HasRoomsToRemove returns a boolean if a field has been set.

### SetRoomsToRemoveNil

`func (o *UpdateRoomGroupRequest) SetRoomsToRemoveNil(b bool)`

 SetRoomsToRemoveNil sets the value for RoomsToRemove to be an explicit nil

### UnsetRoomsToRemove
`func (o *UpdateRoomGroupRequest) UnsetRoomsToRemove()`

UnsetRoomsToRemove ensures that no value is present for RoomsToRemove, not even an explicit nil
### GetGroupName

`func (o *UpdateRoomGroupRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *UpdateRoomGroupRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *UpdateRoomGroupRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *UpdateRoomGroupRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *UpdateRoomGroupRequest) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *UpdateRoomGroupRequest) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


