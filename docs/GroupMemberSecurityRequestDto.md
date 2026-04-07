# GroupMemberSecurityRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**EmployeeFullDto**](EmployeeFullDto.md) |  | 
**GroupAccess** | [**FileShare**](FileShare.md) |  | 
**UserAccess** | Pointer to [**FileShare**](FileShare.md) |  | [optional] 
**Overridden** | **bool** | Specifies if the group access rights are overridden or not. | 
**CanEditAccess** | **bool** | Specifies if the group member can edit the group access rights or not. | 
**Owner** | **bool** | Specifies if the group member is a group owner or not. | 

## Methods

### NewGroupMemberSecurityRequestDto

`func NewGroupMemberSecurityRequestDto(user EmployeeFullDto, groupAccess FileShare, overridden bool, canEditAccess bool, owner bool, ) *GroupMemberSecurityRequestDto`

NewGroupMemberSecurityRequestDto instantiates a new GroupMemberSecurityRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMemberSecurityRequestDtoWithDefaults

`func NewGroupMemberSecurityRequestDtoWithDefaults() *GroupMemberSecurityRequestDto`

NewGroupMemberSecurityRequestDtoWithDefaults instantiates a new GroupMemberSecurityRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *GroupMemberSecurityRequestDto) GetUser() EmployeeFullDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GroupMemberSecurityRequestDto) GetUserOk() (*EmployeeFullDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GroupMemberSecurityRequestDto) SetUser(v EmployeeFullDto)`

SetUser sets User field to given value.


### GetGroupAccess

`func (o *GroupMemberSecurityRequestDto) GetGroupAccess() FileShare`

GetGroupAccess returns the GroupAccess field if non-nil, zero value otherwise.

### GetGroupAccessOk

`func (o *GroupMemberSecurityRequestDto) GetGroupAccessOk() (*FileShare, bool)`

GetGroupAccessOk returns a tuple with the GroupAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAccess

`func (o *GroupMemberSecurityRequestDto) SetGroupAccess(v FileShare)`

SetGroupAccess sets GroupAccess field to given value.


### GetUserAccess

`func (o *GroupMemberSecurityRequestDto) GetUserAccess() FileShare`

GetUserAccess returns the UserAccess field if non-nil, zero value otherwise.

### GetUserAccessOk

`func (o *GroupMemberSecurityRequestDto) GetUserAccessOk() (*FileShare, bool)`

GetUserAccessOk returns a tuple with the UserAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccess

`func (o *GroupMemberSecurityRequestDto) SetUserAccess(v FileShare)`

SetUserAccess sets UserAccess field to given value.

### HasUserAccess

`func (o *GroupMemberSecurityRequestDto) HasUserAccess() bool`

HasUserAccess returns a boolean if a field has been set.

### GetOverridden

`func (o *GroupMemberSecurityRequestDto) GetOverridden() bool`

GetOverridden returns the Overridden field if non-nil, zero value otherwise.

### GetOverriddenOk

`func (o *GroupMemberSecurityRequestDto) GetOverriddenOk() (*bool, bool)`

GetOverriddenOk returns a tuple with the Overridden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridden

`func (o *GroupMemberSecurityRequestDto) SetOverridden(v bool)`

SetOverridden sets Overridden field to given value.


### GetCanEditAccess

`func (o *GroupMemberSecurityRequestDto) GetCanEditAccess() bool`

GetCanEditAccess returns the CanEditAccess field if non-nil, zero value otherwise.

### GetCanEditAccessOk

`func (o *GroupMemberSecurityRequestDto) GetCanEditAccessOk() (*bool, bool)`

GetCanEditAccessOk returns a tuple with the CanEditAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditAccess

`func (o *GroupMemberSecurityRequestDto) SetCanEditAccess(v bool)`

SetCanEditAccess sets CanEditAccess field to given value.


### GetOwner

`func (o *GroupMemberSecurityRequestDto) GetOwner() bool`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GroupMemberSecurityRequestDto) GetOwnerOk() (*bool, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GroupMemberSecurityRequestDto) SetOwner(v bool)`

SetOwner sets Owner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


