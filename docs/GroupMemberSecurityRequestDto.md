# GroupMemberSecurityRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**EmployeeFullDto**](EmployeeFullDto.md) | The member the line is about, as the portal reports the account: the display name, the avatar and the portal  role to show next to the access level. | 
**GroupAccess** | [**FileShare**](FileShare.md) | The level granted to the group as a whole on this file or folder. It belongs to the group record rather than  to the member, so the same value repeats on every line of the answer; a group whose record was set back to  none is answered with an empty list instead. | 
**UserAccess** | Pointer to [**FileShare**](FileShare.md) | The level granted to this member alone on the same file or folder, or `null` when the member has no record of  their own and the group level is what applies. The member who created the file or folder is always reported  here as a room manager, whatever their own record says. | [optional] 
**Overridden** | **bool** | Whether `userAccess` is the level that decides what the member may do. When it is false the member inherits  `groupAccess`, and the creator of the file or folder is always reported as overridden because of the room  manager level forced onto them. | 
**CanEditAccess** | **bool** | Whether the caller may still change the level of this member. It comes back false on the line of the member  who created the file or folder, on the line of the caller themselves, and on every line at once when the  caller may read the file or folder but not manage access to it. | 
**Owner** | **bool** | Whether this member created the file or folder - the owner of the entry, not the owner of the group. Their  level is reported as a room manager one and cannot be taken away through this group. | 

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


