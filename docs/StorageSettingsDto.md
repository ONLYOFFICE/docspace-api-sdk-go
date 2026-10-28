# StorageSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | Pointer to **NullableString** | The storage module, or `null` when the built-in storage is used. | [optional] 
**Props** | Pointer to **map[string]string** | The connection properties stored for the module. | [optional] 
**LastModified** | Pointer to **time.Time** | When the settings were last stored. | [optional] 

## Methods

### NewStorageSettingsDto

`func NewStorageSettingsDto() *StorageSettingsDto`

NewStorageSettingsDto instantiates a new StorageSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSettingsDtoWithDefaults

`func NewStorageSettingsDtoWithDefaults() *StorageSettingsDto`

NewStorageSettingsDtoWithDefaults instantiates a new StorageSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *StorageSettingsDto) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *StorageSettingsDto) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *StorageSettingsDto) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *StorageSettingsDto) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *StorageSettingsDto) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *StorageSettingsDto) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetProps

`func (o *StorageSettingsDto) GetProps() map[string]*string`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *StorageSettingsDto) GetPropsOk() (*map[string]*string, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *StorageSettingsDto) SetProps(v map[string]*string)`

SetProps sets Props field to given value.

### HasProps

`func (o *StorageSettingsDto) HasProps() bool`

HasProps returns a boolean if a field has been set.

### GetLastModified

`func (o *StorageSettingsDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *StorageSettingsDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *StorageSettingsDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *StorageSettingsDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


