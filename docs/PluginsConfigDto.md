# PluginsConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginsData** | Pointer to **[]string** | The array of absolute URLs to the plugin configuration files. | [optional] [readonly] 

## Methods

### NewPluginsConfigDto

`func NewPluginsConfigDto() *PluginsConfigDto`

NewPluginsConfigDto instantiates a new PluginsConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginsConfigDtoWithDefaults

`func NewPluginsConfigDtoWithDefaults() *PluginsConfigDto`

NewPluginsConfigDtoWithDefaults instantiates a new PluginsConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginsData

`func (o *PluginsConfigDto) GetPluginsData() []string`

GetPluginsData returns the PluginsData field if non-nil, zero value otherwise.

### GetPluginsDataOk

`func (o *PluginsConfigDto) GetPluginsDataOk() (*[]string, bool)`

GetPluginsDataOk returns a tuple with the PluginsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginsData

`func (o *PluginsConfigDto) SetPluginsData(v []string)`

SetPluginsData sets PluginsData field to given value.

### HasPluginsData

`func (o *PluginsConfigDto) HasPluginsData() bool`

HasPluginsData returns a boolean if a field has been set.

### SetPluginsDataNil

`func (o *PluginsConfigDto) SetPluginsDataNil(b bool)`

 SetPluginsDataNil sets the value for PluginsData to be an explicit nil

### UnsetPluginsData
`func (o *PluginsConfigDto) UnsetPluginsData()`

UnsetPluginsData ensures that no value is present for PluginsData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


