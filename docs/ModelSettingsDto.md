# ModelSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The model identifier. | 
**Alias** | Pointer to **NullableString** | The display name for the model. | [optional] 
**IsEnabled** | Pointer to **bool** | Whether the model is enabled for use in chat. | [optional] 
**IsRecommended** | Pointer to **bool** | Whether the model is recommended (defined in configuration). | [optional] 
**Capabilities** | [**AiModelCapabilities**](AiModelCapabilities.md) |  | 

## Methods

### NewModelSettingsDto

`func NewModelSettingsDto(id NullableString, capabilities AiModelCapabilities, ) *ModelSettingsDto`

NewModelSettingsDto instantiates a new ModelSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSettingsDtoWithDefaults

`func NewModelSettingsDtoWithDefaults() *ModelSettingsDto`

NewModelSettingsDtoWithDefaults instantiates a new ModelSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelSettingsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelSettingsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelSettingsDto) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ModelSettingsDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ModelSettingsDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAlias

`func (o *ModelSettingsDto) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ModelSettingsDto) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ModelSettingsDto) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ModelSettingsDto) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *ModelSettingsDto) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *ModelSettingsDto) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetIsEnabled

`func (o *ModelSettingsDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ModelSettingsDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ModelSettingsDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ModelSettingsDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsRecommended

`func (o *ModelSettingsDto) GetIsRecommended() bool`

GetIsRecommended returns the IsRecommended field if non-nil, zero value otherwise.

### GetIsRecommendedOk

`func (o *ModelSettingsDto) GetIsRecommendedOk() (*bool, bool)`

GetIsRecommendedOk returns a tuple with the IsRecommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecommended

`func (o *ModelSettingsDto) SetIsRecommended(v bool)`

SetIsRecommended sets IsRecommended field to given value.

### HasIsRecommended

`func (o *ModelSettingsDto) HasIsRecommended() bool`

HasIsRecommended returns a boolean if a field has been set.

### GetCapabilities

`func (o *ModelSettingsDto) GetCapabilities() AiModelCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ModelSettingsDto) GetCapabilitiesOk() (*AiModelCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ModelSettingsDto) SetCapabilities(v AiModelCapabilities)`

SetCapabilities sets Capabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


