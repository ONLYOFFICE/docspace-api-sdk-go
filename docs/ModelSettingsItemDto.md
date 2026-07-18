# ModelSettingsItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelId** | **string** | The model identifier. | 
**IsEnabled** | Pointer to **bool** | Whether the model is enabled for use in chat. | [optional] 
**Alias** | Pointer to **NullableString** | The display name for the model. Only applies to non-recommended models. | [optional] 
**Capabilities** | Pointer to [**AiModelCapabilities**](AiModelCapabilities.md) |  | [optional] 

## Methods

### NewModelSettingsItemDto

`func NewModelSettingsItemDto(modelId string, ) *ModelSettingsItemDto`

NewModelSettingsItemDto instantiates a new ModelSettingsItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSettingsItemDtoWithDefaults

`func NewModelSettingsItemDtoWithDefaults() *ModelSettingsItemDto`

NewModelSettingsItemDtoWithDefaults instantiates a new ModelSettingsItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelId

`func (o *ModelSettingsItemDto) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *ModelSettingsItemDto) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *ModelSettingsItemDto) SetModelId(v string)`

SetModelId sets ModelId field to given value.


### GetIsEnabled

`func (o *ModelSettingsItemDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ModelSettingsItemDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ModelSettingsItemDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ModelSettingsItemDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetAlias

`func (o *ModelSettingsItemDto) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ModelSettingsItemDto) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ModelSettingsItemDto) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ModelSettingsItemDto) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *ModelSettingsItemDto) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *ModelSettingsItemDto) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetCapabilities

`func (o *ModelSettingsItemDto) GetCapabilities() AiModelCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ModelSettingsItemDto) GetCapabilitiesOk() (*AiModelCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ModelSettingsItemDto) SetCapabilities(v AiModelCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ModelSettingsItemDto) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


