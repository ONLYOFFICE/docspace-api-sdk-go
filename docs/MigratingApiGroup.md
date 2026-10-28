# MigratingApiGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShouldImport** | Pointer to **bool** | Specifies whether the API entity should be imported. | [optional] 
**GroupName** | Pointer to **NullableString** | The group name. | [optional] 
**ModuleName** | Pointer to **NullableString** | The group module name. | [optional] 
**UserUidList** | Pointer to **[]string** | The list of group user UIDs. | [optional] 

## Methods

### NewMigratingApiGroup

`func NewMigratingApiGroup() *MigratingApiGroup`

NewMigratingApiGroup instantiates a new MigratingApiGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigratingApiGroupWithDefaults

`func NewMigratingApiGroupWithDefaults() *MigratingApiGroup`

NewMigratingApiGroupWithDefaults instantiates a new MigratingApiGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShouldImport

`func (o *MigratingApiGroup) GetShouldImport() bool`

GetShouldImport returns the ShouldImport field if non-nil, zero value otherwise.

### GetShouldImportOk

`func (o *MigratingApiGroup) GetShouldImportOk() (*bool, bool)`

GetShouldImportOk returns a tuple with the ShouldImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldImport

`func (o *MigratingApiGroup) SetShouldImport(v bool)`

SetShouldImport sets ShouldImport field to given value.

### HasShouldImport

`func (o *MigratingApiGroup) HasShouldImport() bool`

HasShouldImport returns a boolean if a field has been set.

### GetGroupName

`func (o *MigratingApiGroup) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *MigratingApiGroup) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *MigratingApiGroup) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *MigratingApiGroup) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *MigratingApiGroup) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *MigratingApiGroup) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetModuleName

`func (o *MigratingApiGroup) GetModuleName() string`

GetModuleName returns the ModuleName field if non-nil, zero value otherwise.

### GetModuleNameOk

`func (o *MigratingApiGroup) GetModuleNameOk() (*string, bool)`

GetModuleNameOk returns a tuple with the ModuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleName

`func (o *MigratingApiGroup) SetModuleName(v string)`

SetModuleName sets ModuleName field to given value.

### HasModuleName

`func (o *MigratingApiGroup) HasModuleName() bool`

HasModuleName returns a boolean if a field has been set.

### SetModuleNameNil

`func (o *MigratingApiGroup) SetModuleNameNil(b bool)`

 SetModuleNameNil sets the value for ModuleName to be an explicit nil

### UnsetModuleName
`func (o *MigratingApiGroup) UnsetModuleName()`

UnsetModuleName ensures that no value is present for ModuleName, not even an explicit nil
### GetUserUidList

`func (o *MigratingApiGroup) GetUserUidList() []string`

GetUserUidList returns the UserUidList field if non-nil, zero value otherwise.

### GetUserUidListOk

`func (o *MigratingApiGroup) GetUserUidListOk() (*[]string, bool)`

GetUserUidListOk returns a tuple with the UserUidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUidList

`func (o *MigratingApiGroup) SetUserUidList(v []string)`

SetUserUidList sets UserUidList field to given value.

### HasUserUidList

`func (o *MigratingApiGroup) HasUserUidList() bool`

HasUserUidList returns a boolean if a field has been set.

### SetUserUidListNil

`func (o *MigratingApiGroup) SetUserUidListNil(b bool)`

 SetUserUidListNil sets the value for UserUidList to be an explicit nil

### UnsetUserUidList
`func (o *MigratingApiGroup) UnsetUserUidList()`

UnsetUserUidList ensures that no value is present for UserUidList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


