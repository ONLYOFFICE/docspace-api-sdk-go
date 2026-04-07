# FormRoleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleName** | **NullableString** | The role name. | 
**RoleColor** | Pointer to **NullableString** | The role color. | [optional] 
**User** | Pointer to [**EmployeeFullDto**](EmployeeFullDto.md) |  | [optional] 
**Sequence** | **int32** | The role sequence. | 
**Submitted** | **bool** | Specifies if the role is submitted. | 
**StopedBy** | Pointer to [**EmployeeFullDto**](EmployeeFullDto.md) |  | [optional] 
**History** | Pointer to [**map[string]time.Time**](time.Time.md) | The role history. | [optional] 
**RoleStatus** | Pointer to [**FormFillingStatus**](FormFillingStatus.md) |  | [optional] 

## Methods

### NewFormRoleDto

`func NewFormRoleDto(roleName NullableString, sequence int32, submitted bool, ) *FormRoleDto`

NewFormRoleDto instantiates a new FormRoleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormRoleDtoWithDefaults

`func NewFormRoleDtoWithDefaults() *FormRoleDto`

NewFormRoleDtoWithDefaults instantiates a new FormRoleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleName

`func (o *FormRoleDto) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *FormRoleDto) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *FormRoleDto) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### SetRoleNameNil

`func (o *FormRoleDto) SetRoleNameNil(b bool)`

 SetRoleNameNil sets the value for RoleName to be an explicit nil

### UnsetRoleName
`func (o *FormRoleDto) UnsetRoleName()`

UnsetRoleName ensures that no value is present for RoleName, not even an explicit nil
### GetRoleColor

`func (o *FormRoleDto) GetRoleColor() string`

GetRoleColor returns the RoleColor field if non-nil, zero value otherwise.

### GetRoleColorOk

`func (o *FormRoleDto) GetRoleColorOk() (*string, bool)`

GetRoleColorOk returns a tuple with the RoleColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleColor

`func (o *FormRoleDto) SetRoleColor(v string)`

SetRoleColor sets RoleColor field to given value.

### HasRoleColor

`func (o *FormRoleDto) HasRoleColor() bool`

HasRoleColor returns a boolean if a field has been set.

### SetRoleColorNil

`func (o *FormRoleDto) SetRoleColorNil(b bool)`

 SetRoleColorNil sets the value for RoleColor to be an explicit nil

### UnsetRoleColor
`func (o *FormRoleDto) UnsetRoleColor()`

UnsetRoleColor ensures that no value is present for RoleColor, not even an explicit nil
### GetUser

`func (o *FormRoleDto) GetUser() EmployeeFullDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FormRoleDto) GetUserOk() (*EmployeeFullDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FormRoleDto) SetUser(v EmployeeFullDto)`

SetUser sets User field to given value.

### HasUser

`func (o *FormRoleDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetSequence

`func (o *FormRoleDto) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *FormRoleDto) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *FormRoleDto) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetSubmitted

`func (o *FormRoleDto) GetSubmitted() bool`

GetSubmitted returns the Submitted field if non-nil, zero value otherwise.

### GetSubmittedOk

`func (o *FormRoleDto) GetSubmittedOk() (*bool, bool)`

GetSubmittedOk returns a tuple with the Submitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitted

`func (o *FormRoleDto) SetSubmitted(v bool)`

SetSubmitted sets Submitted field to given value.


### GetStopedBy

`func (o *FormRoleDto) GetStopedBy() EmployeeFullDto`

GetStopedBy returns the StopedBy field if non-nil, zero value otherwise.

### GetStopedByOk

`func (o *FormRoleDto) GetStopedByOk() (*EmployeeFullDto, bool)`

GetStopedByOk returns a tuple with the StopedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopedBy

`func (o *FormRoleDto) SetStopedBy(v EmployeeFullDto)`

SetStopedBy sets StopedBy field to given value.

### HasStopedBy

`func (o *FormRoleDto) HasStopedBy() bool`

HasStopedBy returns a boolean if a field has been set.

### GetHistory

`func (o *FormRoleDto) GetHistory() map[string]time.Time`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *FormRoleDto) GetHistoryOk() (*map[string]time.Time, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *FormRoleDto) SetHistory(v map[string]time.Time)`

SetHistory sets History field to given value.

### HasHistory

`func (o *FormRoleDto) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### SetHistoryNil

`func (o *FormRoleDto) SetHistoryNil(b bool)`

 SetHistoryNil sets the value for History to be an explicit nil

### UnsetHistory
`func (o *FormRoleDto) UnsetHistory()`

UnsetHistory ensures that no value is present for History, not even an explicit nil
### GetRoleStatus

`func (o *FormRoleDto) GetRoleStatus() FormFillingStatus`

GetRoleStatus returns the RoleStatus field if non-nil, zero value otherwise.

### GetRoleStatusOk

`func (o *FormRoleDto) GetRoleStatusOk() (*FormFillingStatus, bool)`

GetRoleStatusOk returns a tuple with the RoleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleStatus

`func (o *FormRoleDto) SetRoleStatus(v FormFillingStatus)`

SetRoleStatus sets RoleStatus field to given value.

### HasRoleStatus

`func (o *FormRoleDto) HasRoleStatus() bool`

HasRoleStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


