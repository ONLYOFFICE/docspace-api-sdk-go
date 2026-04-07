# PluginsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Specifies if the plugins are enabled or not. | [optional] 
**Upload** | Pointer to **bool** | Specifies if the plugins can be uploaded or not. | [optional] 
**Delete** | Pointer to **bool** | Specifies if the plugins can be deleted or not. | [optional] 

## Methods

### NewPluginsDto

`func NewPluginsDto() *PluginsDto`

NewPluginsDto instantiates a new PluginsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginsDtoWithDefaults

`func NewPluginsDtoWithDefaults() *PluginsDto`

NewPluginsDtoWithDefaults instantiates a new PluginsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *PluginsDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PluginsDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PluginsDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PluginsDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUpload

`func (o *PluginsDto) GetUpload() bool`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *PluginsDto) GetUploadOk() (*bool, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *PluginsDto) SetUpload(v bool)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *PluginsDto) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetDelete

`func (o *PluginsDto) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *PluginsDto) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *PluginsDto) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *PluginsDto) HasDelete() bool`

HasDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


