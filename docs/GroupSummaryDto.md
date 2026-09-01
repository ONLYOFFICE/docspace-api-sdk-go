# GroupSummaryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The group ID. | 
**Name** | **NullableString** | The group name. | 
**Manager** | Pointer to **NullableString** | The group manager. | [optional] 
**IsSystem** | Pointer to **NullableBool** | Indicates whether the group is a system group. | [optional] 

## Methods

### NewGroupSummaryDto

`func NewGroupSummaryDto(id string, name NullableString, ) *GroupSummaryDto`

NewGroupSummaryDto instantiates a new GroupSummaryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupSummaryDtoWithDefaults

`func NewGroupSummaryDtoWithDefaults() *GroupSummaryDto`

NewGroupSummaryDtoWithDefaults instantiates a new GroupSummaryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupSummaryDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupSummaryDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupSummaryDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GroupSummaryDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupSummaryDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupSummaryDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *GroupSummaryDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GroupSummaryDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetManager

`func (o *GroupSummaryDto) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *GroupSummaryDto) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *GroupSummaryDto) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *GroupSummaryDto) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *GroupSummaryDto) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *GroupSummaryDto) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetIsSystem

`func (o *GroupSummaryDto) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *GroupSummaryDto) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *GroupSummaryDto) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.

### HasIsSystem

`func (o *GroupSummaryDto) HasIsSystem() bool`

HasIsSystem returns a boolean if a field has been set.

### SetIsSystemNil

`func (o *GroupSummaryDto) SetIsSystemNil(b bool)`

 SetIsSystemNil sets the value for IsSystem to be an explicit nil

### UnsetIsSystem
`func (o *GroupSummaryDto) UnsetIsSystem()`

UnsetIsSystem ensures that no value is present for IsSystem, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


