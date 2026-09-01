# SecurityDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebItemId** | Pointer to **NullableString** | The module ID. | [optional] 
**Users** | Pointer to [**[]EmployeeDto**](EmployeeDto.md) | The list of users with the access to the module. | [optional] 
**Groups** | Pointer to [**[]GroupSummaryDto**](GroupSummaryDto.md) | The list of groups with the access to the module. | [optional] 
**Enabled** | Pointer to **bool** | Specifies if the security settings are enabled or not. | [optional] 
**IsSubItem** | Pointer to **bool** | Specifies if the module is a subitem or not. | [optional] 

## Methods

### NewSecurityDto

`func NewSecurityDto() *SecurityDto`

NewSecurityDto instantiates a new SecurityDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityDtoWithDefaults

`func NewSecurityDtoWithDefaults() *SecurityDto`

NewSecurityDtoWithDefaults instantiates a new SecurityDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebItemId

`func (o *SecurityDto) GetWebItemId() string`

GetWebItemId returns the WebItemId field if non-nil, zero value otherwise.

### GetWebItemIdOk

`func (o *SecurityDto) GetWebItemIdOk() (*string, bool)`

GetWebItemIdOk returns a tuple with the WebItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebItemId

`func (o *SecurityDto) SetWebItemId(v string)`

SetWebItemId sets WebItemId field to given value.

### HasWebItemId

`func (o *SecurityDto) HasWebItemId() bool`

HasWebItemId returns a boolean if a field has been set.

### SetWebItemIdNil

`func (o *SecurityDto) SetWebItemIdNil(b bool)`

 SetWebItemIdNil sets the value for WebItemId to be an explicit nil

### UnsetWebItemId
`func (o *SecurityDto) UnsetWebItemId()`

UnsetWebItemId ensures that no value is present for WebItemId, not even an explicit nil
### GetUsers

`func (o *SecurityDto) GetUsers() []EmployeeDto`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *SecurityDto) GetUsersOk() (*[]EmployeeDto, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *SecurityDto) SetUsers(v []EmployeeDto)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *SecurityDto) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *SecurityDto) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *SecurityDto) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetGroups

`func (o *SecurityDto) GetGroups() []GroupSummaryDto`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *SecurityDto) GetGroupsOk() (*[]GroupSummaryDto, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *SecurityDto) SetGroups(v []GroupSummaryDto)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *SecurityDto) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *SecurityDto) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *SecurityDto) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetEnabled

`func (o *SecurityDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SecurityDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SecurityDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SecurityDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIsSubItem

`func (o *SecurityDto) GetIsSubItem() bool`

GetIsSubItem returns the IsSubItem field if non-nil, zero value otherwise.

### GetIsSubItemOk

`func (o *SecurityDto) GetIsSubItemOk() (*bool, bool)`

GetIsSubItemOk returns a tuple with the IsSubItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubItem

`func (o *SecurityDto) SetIsSubItem(v bool)`

SetIsSubItem sets IsSubItem field to given value.

### HasIsSubItem

`func (o *SecurityDto) HasIsSubItem() bool`

HasIsSubItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


