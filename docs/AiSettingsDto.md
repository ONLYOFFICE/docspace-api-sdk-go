# AiSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VectorizationEnabled** | Pointer to **bool** | Indicates whether document vectorization is enabled. | [optional] 
**VectorizationNeedReset** | Pointer to **bool** | Indicates whether the embedding provider API key needs to be reconfigured. | [optional] 
**AiReady** | Pointer to **bool** | Indicates whether the AI subsystem is fully configured and operational. | [optional] 
**EmbeddingModel** | **NullableString** | The name of the embedding model used for document vectorization. | 
**SystemAiEnabled** | Pointer to **bool** | Indicates whether the system-level AI provider is enabled. | [optional] 
**RecommendedModelForForms** | Pointer to **NullableString** | The identifier of the model recommended for form generation. | [optional] 
**ToolPermissionMode** | Pointer to [**AiToolPermissionMode**](AiToolPermissionMode.md) | How tool calls made by the model are approved for the current user. The default applies while the user has stored nothing. | [optional] 

## Methods

### NewAiSettingsDto

`func NewAiSettingsDto(embeddingModel NullableString, ) *AiSettingsDto`

NewAiSettingsDto instantiates a new AiSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSettingsDtoWithDefaults

`func NewAiSettingsDtoWithDefaults() *AiSettingsDto`

NewAiSettingsDtoWithDefaults instantiates a new AiSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVectorizationEnabled

`func (o *AiSettingsDto) GetVectorizationEnabled() bool`

GetVectorizationEnabled returns the VectorizationEnabled field if non-nil, zero value otherwise.

### GetVectorizationEnabledOk

`func (o *AiSettingsDto) GetVectorizationEnabledOk() (*bool, bool)`

GetVectorizationEnabledOk returns a tuple with the VectorizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorizationEnabled

`func (o *AiSettingsDto) SetVectorizationEnabled(v bool)`

SetVectorizationEnabled sets VectorizationEnabled field to given value.

### HasVectorizationEnabled

`func (o *AiSettingsDto) HasVectorizationEnabled() bool`

HasVectorizationEnabled returns a boolean if a field has been set.

### GetVectorizationNeedReset

`func (o *AiSettingsDto) GetVectorizationNeedReset() bool`

GetVectorizationNeedReset returns the VectorizationNeedReset field if non-nil, zero value otherwise.

### GetVectorizationNeedResetOk

`func (o *AiSettingsDto) GetVectorizationNeedResetOk() (*bool, bool)`

GetVectorizationNeedResetOk returns a tuple with the VectorizationNeedReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorizationNeedReset

`func (o *AiSettingsDto) SetVectorizationNeedReset(v bool)`

SetVectorizationNeedReset sets VectorizationNeedReset field to given value.

### HasVectorizationNeedReset

`func (o *AiSettingsDto) HasVectorizationNeedReset() bool`

HasVectorizationNeedReset returns a boolean if a field has been set.

### GetAiReady

`func (o *AiSettingsDto) GetAiReady() bool`

GetAiReady returns the AiReady field if non-nil, zero value otherwise.

### GetAiReadyOk

`func (o *AiSettingsDto) GetAiReadyOk() (*bool, bool)`

GetAiReadyOk returns a tuple with the AiReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiReady

`func (o *AiSettingsDto) SetAiReady(v bool)`

SetAiReady sets AiReady field to given value.

### HasAiReady

`func (o *AiSettingsDto) HasAiReady() bool`

HasAiReady returns a boolean if a field has been set.

### GetEmbeddingModel

`func (o *AiSettingsDto) GetEmbeddingModel() string`

GetEmbeddingModel returns the EmbeddingModel field if non-nil, zero value otherwise.

### GetEmbeddingModelOk

`func (o *AiSettingsDto) GetEmbeddingModelOk() (*string, bool)`

GetEmbeddingModelOk returns a tuple with the EmbeddingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingModel

`func (o *AiSettingsDto) SetEmbeddingModel(v string)`

SetEmbeddingModel sets EmbeddingModel field to given value.


### SetEmbeddingModelNil

`func (o *AiSettingsDto) SetEmbeddingModelNil(b bool)`

 SetEmbeddingModelNil sets the value for EmbeddingModel to be an explicit nil

### UnsetEmbeddingModel
`func (o *AiSettingsDto) UnsetEmbeddingModel()`

UnsetEmbeddingModel ensures that no value is present for EmbeddingModel, not even an explicit nil
### GetSystemAiEnabled

`func (o *AiSettingsDto) GetSystemAiEnabled() bool`

GetSystemAiEnabled returns the SystemAiEnabled field if non-nil, zero value otherwise.

### GetSystemAiEnabledOk

`func (o *AiSettingsDto) GetSystemAiEnabledOk() (*bool, bool)`

GetSystemAiEnabledOk returns a tuple with the SystemAiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAiEnabled

`func (o *AiSettingsDto) SetSystemAiEnabled(v bool)`

SetSystemAiEnabled sets SystemAiEnabled field to given value.

### HasSystemAiEnabled

`func (o *AiSettingsDto) HasSystemAiEnabled() bool`

HasSystemAiEnabled returns a boolean if a field has been set.

### GetRecommendedModelForForms

`func (o *AiSettingsDto) GetRecommendedModelForForms() string`

GetRecommendedModelForForms returns the RecommendedModelForForms field if non-nil, zero value otherwise.

### GetRecommendedModelForFormsOk

`func (o *AiSettingsDto) GetRecommendedModelForFormsOk() (*string, bool)`

GetRecommendedModelForFormsOk returns a tuple with the RecommendedModelForForms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedModelForForms

`func (o *AiSettingsDto) SetRecommendedModelForForms(v string)`

SetRecommendedModelForForms sets RecommendedModelForForms field to given value.

### HasRecommendedModelForForms

`func (o *AiSettingsDto) HasRecommendedModelForForms() bool`

HasRecommendedModelForForms returns a boolean if a field has been set.

### SetRecommendedModelForFormsNil

`func (o *AiSettingsDto) SetRecommendedModelForFormsNil(b bool)`

 SetRecommendedModelForFormsNil sets the value for RecommendedModelForForms to be an explicit nil

### UnsetRecommendedModelForForms
`func (o *AiSettingsDto) UnsetRecommendedModelForForms()`

UnsetRecommendedModelForForms ensures that no value is present for RecommendedModelForForms, not even an explicit nil
### GetToolPermissionMode

`func (o *AiSettingsDto) GetToolPermissionMode() AiToolPermissionMode`

GetToolPermissionMode returns the ToolPermissionMode field if non-nil, zero value otherwise.

### GetToolPermissionModeOk

`func (o *AiSettingsDto) GetToolPermissionModeOk() (*AiToolPermissionMode, bool)`

GetToolPermissionModeOk returns a tuple with the ToolPermissionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolPermissionMode

`func (o *AiSettingsDto) SetToolPermissionMode(v AiToolPermissionMode)`

SetToolPermissionMode sets ToolPermissionMode field to given value.

### HasToolPermissionMode

`func (o *AiSettingsDto) HasToolPermissionMode() bool`

HasToolPermissionMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


