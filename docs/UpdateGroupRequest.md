# UpdateGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MembersToAdd** | Pointer to **[]string** | The list of user IDs to add to the group. | [optional] 
**MembersToRemove** | Pointer to **[]string** | The list of user IDs to remove from the group. | [optional] 
**GroupManager** | Pointer to **string** | The group manager ID. | [optional] 
**GroupName** | Pointer to **NullableString** | The group name. | [optional] 

## Methods

### NewUpdateGroupRequest

`func NewUpdateGroupRequest() *UpdateGroupRequest`

NewUpdateGroupRequest instantiates a new UpdateGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGroupRequestWithDefaults

`func NewUpdateGroupRequestWithDefaults() *UpdateGroupRequest`

NewUpdateGroupRequestWithDefaults instantiates a new UpdateGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembersToAdd

`func (o *UpdateGroupRequest) GetMembersToAdd() []string`

GetMembersToAdd returns the MembersToAdd field if non-nil, zero value otherwise.

### GetMembersToAddOk

`func (o *UpdateGroupRequest) GetMembersToAddOk() (*[]string, bool)`

GetMembersToAddOk returns a tuple with the MembersToAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersToAdd

`func (o *UpdateGroupRequest) SetMembersToAdd(v []string)`

SetMembersToAdd sets MembersToAdd field to given value.

### HasMembersToAdd

`func (o *UpdateGroupRequest) HasMembersToAdd() bool`

HasMembersToAdd returns a boolean if a field has been set.

### SetMembersToAddNil

`func (o *UpdateGroupRequest) SetMembersToAddNil(b bool)`

 SetMembersToAddNil sets the value for MembersToAdd to be an explicit nil

### UnsetMembersToAdd
`func (o *UpdateGroupRequest) UnsetMembersToAdd()`

UnsetMembersToAdd ensures that no value is present for MembersToAdd, not even an explicit nil
### GetMembersToRemove

`func (o *UpdateGroupRequest) GetMembersToRemove() []string`

GetMembersToRemove returns the MembersToRemove field if non-nil, zero value otherwise.

### GetMembersToRemoveOk

`func (o *UpdateGroupRequest) GetMembersToRemoveOk() (*[]string, bool)`

GetMembersToRemoveOk returns a tuple with the MembersToRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersToRemove

`func (o *UpdateGroupRequest) SetMembersToRemove(v []string)`

SetMembersToRemove sets MembersToRemove field to given value.

### HasMembersToRemove

`func (o *UpdateGroupRequest) HasMembersToRemove() bool`

HasMembersToRemove returns a boolean if a field has been set.

### SetMembersToRemoveNil

`func (o *UpdateGroupRequest) SetMembersToRemoveNil(b bool)`

 SetMembersToRemoveNil sets the value for MembersToRemove to be an explicit nil

### UnsetMembersToRemove
`func (o *UpdateGroupRequest) UnsetMembersToRemove()`

UnsetMembersToRemove ensures that no value is present for MembersToRemove, not even an explicit nil
### GetGroupManager

`func (o *UpdateGroupRequest) GetGroupManager() string`

GetGroupManager returns the GroupManager field if non-nil, zero value otherwise.

### GetGroupManagerOk

`func (o *UpdateGroupRequest) GetGroupManagerOk() (*string, bool)`

GetGroupManagerOk returns a tuple with the GroupManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupManager

`func (o *UpdateGroupRequest) SetGroupManager(v string)`

SetGroupManager sets GroupManager field to given value.

### HasGroupManager

`func (o *UpdateGroupRequest) HasGroupManager() bool`

HasGroupManager returns a boolean if a field has been set.

### GetGroupName

`func (o *UpdateGroupRequest) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *UpdateGroupRequest) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *UpdateGroupRequest) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *UpdateGroupRequest) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *UpdateGroupRequest) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *UpdateGroupRequest) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


