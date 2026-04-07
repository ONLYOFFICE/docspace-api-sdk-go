# GroupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | The group name. | 
**Parent** | Pointer to **NullableString** | The parent group ID. | [optional] 
**Category** | **string** | The group category ID. | 
**Id** | **string** | The group ID. | 
**IsLDAP** | **bool** | Specifies if the LDAP settings are enabled for the group or not. | 
**IsSystem** | Pointer to **NullableBool** | Indicates whether the group is a system group. | [optional] 
**Manager** | Pointer to [**EmployeeFullDto**](EmployeeFullDto.md) |  | [optional] 
**Members** | Pointer to [**[]EmployeeFullDto**](EmployeeFullDto.md) | The list of group members. | [optional] 
**Shared** | Pointer to **NullableBool** | Specifies whether the group can be shared or not. | [optional] 
**MembersCount** | Pointer to **int32** | The number of group members. | [optional] 

## Methods

### NewGroupDto

`func NewGroupDto(name NullableString, category string, id string, isLDAP bool, ) *GroupDto`

NewGroupDto instantiates a new GroupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupDtoWithDefaults

`func NewGroupDtoWithDefaults() *GroupDto`

NewGroupDtoWithDefaults instantiates a new GroupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GroupDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *GroupDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GroupDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParent

`func (o *GroupDto) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GroupDto) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GroupDto) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *GroupDto) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *GroupDto) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *GroupDto) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetCategory

`func (o *GroupDto) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GroupDto) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GroupDto) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetId

`func (o *GroupDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupDto) SetId(v string)`

SetId sets Id field to given value.


### GetIsLDAP

`func (o *GroupDto) GetIsLDAP() bool`

GetIsLDAP returns the IsLDAP field if non-nil, zero value otherwise.

### GetIsLDAPOk

`func (o *GroupDto) GetIsLDAPOk() (*bool, bool)`

GetIsLDAPOk returns a tuple with the IsLDAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLDAP

`func (o *GroupDto) SetIsLDAP(v bool)`

SetIsLDAP sets IsLDAP field to given value.


### GetIsSystem

`func (o *GroupDto) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *GroupDto) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *GroupDto) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.

### HasIsSystem

`func (o *GroupDto) HasIsSystem() bool`

HasIsSystem returns a boolean if a field has been set.

### SetIsSystemNil

`func (o *GroupDto) SetIsSystemNil(b bool)`

 SetIsSystemNil sets the value for IsSystem to be an explicit nil

### UnsetIsSystem
`func (o *GroupDto) UnsetIsSystem()`

UnsetIsSystem ensures that no value is present for IsSystem, not even an explicit nil
### GetManager

`func (o *GroupDto) GetManager() EmployeeFullDto`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *GroupDto) GetManagerOk() (*EmployeeFullDto, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *GroupDto) SetManager(v EmployeeFullDto)`

SetManager sets Manager field to given value.

### HasManager

`func (o *GroupDto) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetMembers

`func (o *GroupDto) GetMembers() []EmployeeFullDto`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GroupDto) GetMembersOk() (*[]EmployeeFullDto, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GroupDto) SetMembers(v []EmployeeFullDto)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GroupDto) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembersNil

`func (o *GroupDto) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *GroupDto) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil
### GetShared

`func (o *GroupDto) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *GroupDto) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *GroupDto) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *GroupDto) HasShared() bool`

HasShared returns a boolean if a field has been set.

### SetSharedNil

`func (o *GroupDto) SetSharedNil(b bool)`

 SetSharedNil sets the value for Shared to be an explicit nil

### UnsetShared
`func (o *GroupDto) UnsetShared()`

UnsetShared ensures that no value is present for Shared, not even an explicit nil
### GetMembersCount

`func (o *GroupDto) GetMembersCount() int32`

GetMembersCount returns the MembersCount field if non-nil, zero value otherwise.

### GetMembersCountOk

`func (o *GroupDto) GetMembersCountOk() (*int32, bool)`

GetMembersCountOk returns a tuple with the MembersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCount

`func (o *GroupDto) SetMembersCount(v int32)`

SetMembersCount sets MembersCount field to given value.

### HasMembersCount

`func (o *GroupDto) HasMembersCount() bool`

HasMembersCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


