# FormRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoomId** | Pointer to **int32** | The room ID. | [optional] 
**RoleName** | Pointer to **NullableString** | The role name. | [optional] 
**RoleColor** | Pointer to **NullableString** | The role color. | [optional] 
**UserId** | Pointer to **string** | The user ID. | [optional] 
**Sequence** | Pointer to **int32** | The role sequence. | [optional] 
**Submitted** | Pointer to **bool** | Specifies if the role was submitted or not. | [optional] 
**OpenedAt** | Pointer to **time.Time** | The date and time when the role was opened. | [optional] 
**SubmissionDate** | Pointer to **time.Time** | The date and time when the role was submitted. | [optional] 

## Methods

### NewFormRole

`func NewFormRole() *FormRole`

NewFormRole instantiates a new FormRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormRoleWithDefaults

`func NewFormRoleWithDefaults() *FormRole`

NewFormRoleWithDefaults instantiates a new FormRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoomId

`func (o *FormRole) GetRoomId() int32`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *FormRole) GetRoomIdOk() (*int32, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *FormRole) SetRoomId(v int32)`

SetRoomId sets RoomId field to given value.

### HasRoomId

`func (o *FormRole) HasRoomId() bool`

HasRoomId returns a boolean if a field has been set.

### GetRoleName

`func (o *FormRole) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *FormRole) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *FormRole) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *FormRole) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### SetRoleNameNil

`func (o *FormRole) SetRoleNameNil(b bool)`

 SetRoleNameNil sets the value for RoleName to be an explicit nil

### UnsetRoleName
`func (o *FormRole) UnsetRoleName()`

UnsetRoleName ensures that no value is present for RoleName, not even an explicit nil
### GetRoleColor

`func (o *FormRole) GetRoleColor() string`

GetRoleColor returns the RoleColor field if non-nil, zero value otherwise.

### GetRoleColorOk

`func (o *FormRole) GetRoleColorOk() (*string, bool)`

GetRoleColorOk returns a tuple with the RoleColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleColor

`func (o *FormRole) SetRoleColor(v string)`

SetRoleColor sets RoleColor field to given value.

### HasRoleColor

`func (o *FormRole) HasRoleColor() bool`

HasRoleColor returns a boolean if a field has been set.

### SetRoleColorNil

`func (o *FormRole) SetRoleColorNil(b bool)`

 SetRoleColorNil sets the value for RoleColor to be an explicit nil

### UnsetRoleColor
`func (o *FormRole) UnsetRoleColor()`

UnsetRoleColor ensures that no value is present for RoleColor, not even an explicit nil
### GetUserId

`func (o *FormRole) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FormRole) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FormRole) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FormRole) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSequence

`func (o *FormRole) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *FormRole) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *FormRole) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *FormRole) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSubmitted

`func (o *FormRole) GetSubmitted() bool`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *FormRole) GetSubmittedOk() (*bool, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *FormRole) SetSubmitted(v bool)`

SetSubmitted sets Submitted field to given value.

### HasSubmitted

`func (o *FormRole) HasSubmitted() bool`

HasSubmitted returns a boolean if a field has been set.

### GetOpenedAt

`func (o *FormRole) GetOpenedAt() time.Time`

GetOpenedAt returns the OpenedAt field if non-nil, zero value otherwise.

### GetOpenedAtOk

`func (o *FormRole) GetOpenedAtOk() (*time.Time, bool)`

GetOpenedAtOk returns a tuple with the OpenedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedAt

`func (o *FormRole) SetOpenedAt(v time.Time)`

SetOpenedAt sets OpenedAt field to given value.

### HasOpenedAt

`func (o *FormRole) HasOpenedAt() bool`

HasOpenedAt returns a boolean if a field has been set.

### GetSubmissionDate

`func (o *FormRole) GetSubmissionDate() time.Time`

GetSubmissionDate returns the SubmissionDate field if non-nil, zero value otherwise.

### GetSubmissionDateOk

`func (o *FormRole) GetSubmissionDateOk() (*time.Time, bool)`

GetSubmissionDateOk returns a tuple with the SubmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDate

`func (o *FormRole) SetSubmissionDate(v time.Time)`

SetSubmissionDate sets SubmissionDate field to given value.

### HasSubmissionDate

`func (o *FormRole) HasSubmissionDate() bool`

HasSubmissionDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


