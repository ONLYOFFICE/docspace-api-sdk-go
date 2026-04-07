# GroupRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to **[]string** | The list of group member IDs. | [optional] 
**GroupManager** | **string** | The group manager ID. | 
**GroupName** | Pointer to **NullableString** | The group name. | [optional] 

## Methods

### NewGroupRequestDto

`func NewGroupRequestDto(groupManager string, ) *GroupRequestDto`

NewGroupRequestDto instantiates a new GroupRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRequestDtoWithDefaults

`func NewGroupRequestDtoWithDefaults() *GroupRequestDto`

NewGroupRequestDtoWithDefaults instantiates a new GroupRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *GroupRequestDto) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GroupRequestDto) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GroupRequestDto) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GroupRequestDto) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembersNil

`func (o *GroupRequestDto) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *GroupRequestDto) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil
### GetGroupManager

`func (o *GroupRequestDto) GetGroupManager() string`

GetGroupManager returns the GroupManager field if non-nil, zero value otherwise.

### GetGroupManagerOk

`func (o *GroupRequestDto) GetGroupManagerOk() (*string, bool)`

GetGroupManagerOk returns a tuple with the GroupManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupManager

`func (o *GroupRequestDto) SetGroupManager(v string)`

SetGroupManager sets GroupManager field to given value.


### GetGroupName

`func (o *GroupRequestDto) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GroupRequestDto) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GroupRequestDto) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *GroupRequestDto) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *GroupRequestDto) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *GroupRequestDto) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


