# RoomSecurityDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to [**[]FileShareDto**](FileShareDto.md) | The access entries of the subjects named in the request, read back after the change was applied. A subject the  caller may not see is missing from it, so comparing this list with the request is the way to learn who was  skipped; it is null when nothing was applied at all. | [optional] 
**Warning** | Pointer to **NullableString** | The reason the first subject that could not be handled was skipped, in the language of the request, while the  rest of the list was still applied. Null when every named subject went through. The text is meant to be shown  to a person, not matched against. | [optional] 
**Error** | Pointer to [**RoomSecurityError**](RoomSecurityError.md) | Reports the one case in which nothing at all was changed: a member being removed still holds a role in a form  of the room, and the request did not ask to remove them anyway. Repeat the call with `force` to remove them  together with the role. | [optional] 

## Methods

### NewRoomSecurityDto

`func NewRoomSecurityDto() *RoomSecurityDto`

NewRoomSecurityDto instantiates a new RoomSecurityDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomSecurityDtoWithDefaults

`func NewRoomSecurityDtoWithDefaults() *RoomSecurityDto`

NewRoomSecurityDtoWithDefaults instantiates a new RoomSecurityDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *RoomSecurityDto) GetMembers() []FileShareDto`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *RoomSecurityDto) GetMembersOk() (*[]FileShareDto, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *RoomSecurityDto) SetMembers(v []FileShareDto)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *RoomSecurityDto) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembersNil

`func (o *RoomSecurityDto) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *RoomSecurityDto) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil
### GetWarning

`func (o *RoomSecurityDto) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *RoomSecurityDto) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *RoomSecurityDto) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *RoomSecurityDto) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### SetWarningNil

`func (o *RoomSecurityDto) SetWarningNil(b bool)`

 SetWarningNil sets the value for Warning to be an explicit nil

### UnsetWarning
`func (o *RoomSecurityDto) UnsetWarning()`

UnsetWarning ensures that no value is present for Warning, not even an explicit nil
### GetError

`func (o *RoomSecurityDto) GetError() RoomSecurityError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RoomSecurityDto) GetErrorOk() (*RoomSecurityError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RoomSecurityDto) SetError(v RoomSecurityError)`

SetError sets Error field to given value.

### HasError

`func (o *RoomSecurityDto) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


