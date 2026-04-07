# CdnStorageSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | Pointer to **NullableString** |  | [optional] 
**Props** | Pointer to **map[string]string** |  | [optional] 
**LastModified** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCdnStorageSettings

`func NewCdnStorageSettings() *CdnStorageSettings`

NewCdnStorageSettings instantiates a new CdnStorageSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnStorageSettingsWithDefaults

`func NewCdnStorageSettingsWithDefaults() *CdnStorageSettings`

NewCdnStorageSettingsWithDefaults instantiates a new CdnStorageSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *CdnStorageSettings) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *CdnStorageSettings) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *CdnStorageSettings) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *CdnStorageSettings) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *CdnStorageSettings) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *CdnStorageSettings) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetProps

`func (o *CdnStorageSettings) GetProps() map[string]string`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *CdnStorageSettings) GetPropsOk() (*map[string]string, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *CdnStorageSettings) SetProps(v map[string]string)`

SetProps sets Props field to given value.

### HasProps

`func (o *CdnStorageSettings) HasProps() bool`

HasProps returns a boolean if a field has been set.

### SetPropsNil

`func (o *CdnStorageSettings) SetPropsNil(b bool)`

 SetPropsNil sets the value for Props to be an explicit nil

### UnsetProps
`func (o *CdnStorageSettings) UnsetProps()`

UnsetProps ensures that no value is present for Props, not even an explicit nil
### GetLastModified

`func (o *CdnStorageSettings) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *CdnStorageSettings) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *CdnStorageSettings) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *CdnStorageSettings) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


