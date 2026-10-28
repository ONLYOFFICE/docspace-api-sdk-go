# FormRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoomId** | Pointer to **int32** | The ID of the room the form is in. It is stored with the role as sent, so pass the room the form lives in. | [optional] 
**RoleName** | Pointer to **NullableString** | The name of a role the form defines, such as the one the form author gave a group of fields. | [optional] 
**RoleColor** | Pointer to **NullableString** | The color the editor marks the fields of this role with, as a hex code. | [optional] 
**UserId** | Pointer to **string** | The account that fills this role. It is notified once filling starts, unless it is the caller. | [optional] 
**Sequence** | Pointer to **int32** | Accepted for compatibility and ignored: the position of the role in the list sets the filling order. | [optional] 
**Submitted** | Pointer to **bool** | Whether this role counts as already submitted. It is stored as sent; send false when filling starts. | [optional] 
**OpenedAt** | Pointer to **time.Time** | Accepted for compatibility and ignored: the portal records when the role is opened. | [optional] 
**SubmissionDate** | Pointer to **time.Time** | Accepted for compatibility and ignored: the portal records when the role is submitted. | [optional] 

## Methods

### NewFormRoleRequest

`func NewFormRoleRequest() *FormRoleRequest`

NewFormRoleRequest instantiates a new FormRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormRoleRequestWithDefaults

`func NewFormRoleRequestWithDefaults() *FormRoleRequest`

NewFormRoleRequestWithDefaults instantiates a new FormRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoomId

`func (o *FormRoleRequest) GetRoomId() int32`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *FormRoleRequest) GetRoomIdOk() (*int32, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *FormRoleRequest) SetRoomId(v int32)`

SetRoomId sets RoomId field to given value.

### HasRoomId

`func (o *FormRoleRequest) HasRoomId() bool`

HasRoomId returns a boolean if a field has been set.

### GetRoleName

`func (o *FormRoleRequest) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *FormRoleRequest) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *FormRoleRequest) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *FormRoleRequest) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### SetRoleNameNil

`func (o *FormRoleRequest) SetRoleNameNil(b bool)`

 SetRoleNameNil sets the value for RoleName to be an explicit nil

### UnsetRoleName
`func (o *FormRoleRequest) UnsetRoleName()`

UnsetRoleName ensures that no value is present for RoleName, not even an explicit nil
### GetRoleColor

`func (o *FormRoleRequest) GetRoleColor() string`

GetRoleColor returns the RoleColor field if non-nil, zero value otherwise.

### GetRoleColorOk

`func (o *FormRoleRequest) GetRoleColorOk() (*string, bool)`

GetRoleColorOk returns a tuple with the RoleColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleColor

`func (o *FormRoleRequest) SetRoleColor(v string)`

SetRoleColor sets RoleColor field to given value.

### HasRoleColor

`func (o *FormRoleRequest) HasRoleColor() bool`

HasRoleColor returns a boolean if a field has been set.

### SetRoleColorNil

`func (o *FormRoleRequest) SetRoleColorNil(b bool)`

 SetRoleColorNil sets the value for RoleColor to be an explicit nil

### UnsetRoleColor
`func (o *FormRoleRequest) UnsetRoleColor()`

UnsetRoleColor ensures that no value is present for RoleColor, not even an explicit nil
### GetUserId

`func (o *FormRoleRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FormRoleRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FormRoleRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FormRoleRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSequence

`func (o *FormRoleRequest) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *FormRoleRequest) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *FormRoleRequest) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *FormRoleRequest) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSubmitted

`func (o *FormRoleRequest) GetSubmitted() bool`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *FormRoleRequest) GetSubmittedOk() (*bool, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *FormRoleRequest) SetSubmitted(v bool)`

SetSubmitted sets Submitted field to given value.

### HasSubmitted

`func (o *FormRoleRequest) HasSubmitted() bool`

HasSubmitted returns a boolean if a field has been set.

### GetOpenedAt

`func (o *FormRoleRequest) GetOpenedAt() time.Time`

GetOpenedAt returns the OpenedAt field if non-nil, zero value otherwise.

### GetOpenedAtOk

`func (o *FormRoleRequest) GetOpenedAtOk() (*time.Time, bool)`

GetOpenedAtOk returns a tuple with the OpenedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedAt

`func (o *FormRoleRequest) SetOpenedAt(v time.Time)`

SetOpenedAt sets OpenedAt field to given value.

### HasOpenedAt

`func (o *FormRoleRequest) HasOpenedAt() bool`

HasOpenedAt returns a boolean if a field has been set.

### GetSubmissionDate

`func (o *FormRoleRequest) GetSubmissionDate() time.Time`

GetSubmissionDate returns the SubmissionDate field if non-nil, zero value otherwise.

### GetSubmissionDateOk

`func (o *FormRoleRequest) GetSubmissionDateOk() (*time.Time, bool)`

GetSubmissionDateOk returns a tuple with the SubmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDate

`func (o *FormRoleRequest) SetSubmissionDate(v time.Time)`

SetSubmissionDate sets SubmissionDate field to given value.

### HasSubmissionDate

`func (o *FormRoleRequest) HasSubmissionDate() bool`

HasSubmissionDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


