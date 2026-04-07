# StorageSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | Pointer to **NullableString** |  | [optional] 
**Props** | Pointer to **map[string]string** |  | [optional] 
**LastModified** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewStorageSettings

`func NewStorageSettings() *StorageSettings`

NewStorageSettings instantiates a new StorageSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSettingsWithDefaults

`func NewStorageSettingsWithDefaults() *StorageSettings`

NewStorageSettingsWithDefaults instantiates a new StorageSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *StorageSettings) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *StorageSettings) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *StorageSettings) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *StorageSettings) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *StorageSettings) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *StorageSettings) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetProps

`func (o *StorageSettings) GetProps() map[string]string`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *StorageSettings) GetPropsOk() (*map[string]string, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *StorageSettings) SetProps(v map[string]string)`

SetProps sets Props field to given value.

### HasProps

`func (o *StorageSettings) HasProps() bool`

HasProps returns a boolean if a field has been set.

### SetPropsNil

`func (o *StorageSettings) SetPropsNil(b bool)`

 SetPropsNil sets the value for Props to be an explicit nil

### UnsetProps
`func (o *StorageSettings) UnsetProps()`

UnsetProps ensures that no value is present for Props, not even an explicit nil
### GetLastModified

`func (o *StorageSettings) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *StorageSettings) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *StorageSettings) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *StorageSettings) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


