# AiAiSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VectorizationEnabled** | Pointer to **bool** | Indicates whether document vectorization is enabled. | [optional] 
**VectorizationNeedReset** | Pointer to **bool** | Indicates whether the embedding provider API key needs to be reconfigured. | [optional] 
**AiReady** | Pointer to **bool** | Indicates whether the AI subsystem is fully configured and operational. | [optional] 
**EmbeddingModel** | **NullableString** | The name of the embedding model used for document vectorization. | 
**SystemAiEnabled** | Pointer to **bool** | Indicates whether the system-level AI provider is enabled. | [optional] 
**RecommendedModelForForms** | Pointer to **NullableString** | The identifier of the model recommended for form generation. | [optional] 

## Methods

### NewAiAiSettingsDto

`func NewAiAiSettingsDto(embeddingModel NullableString, ) *AiAiSettingsDto`

NewAiAiSettingsDto instantiates a new AiAiSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAiSettingsDtoWithDefaults

`func NewAiAiSettingsDtoWithDefaults() *AiAiSettingsDto`

NewAiAiSettingsDtoWithDefaults instantiates a new AiAiSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVectorizationEnabled

`func (o *AiAiSettingsDto) GetVectorizationEnabled() bool`

GetVectorizationEnabled returns the VectorizationEnabled field if non-nil, zero value otherwise.

### GetVectorizationEnabledOk

`func (o *AiAiSettingsDto) GetVectorizationEnabledOk() (*bool, bool)`

GetVectorizationEnabledOk returns a tuple with the VectorizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorizationEnabled

`func (o *AiAiSettingsDto) SetVectorizationEnabled(v bool)`

SetVectorizationEnabled sets VectorizationEnabled field to given value.

### HasVectorizationEnabled

`func (o *AiAiSettingsDto) HasVectorizationEnabled() bool`

HasVectorizationEnabled returns a boolean if a field has been set.

### GetVectorizationNeedReset

`func (o *AiAiSettingsDto) GetVectorizationNeedReset() bool`

GetVectorizationNeedReset returns the VectorizationNeedReset field if non-nil, zero value otherwise.

### GetVectorizationNeedResetOk

`func (o *AiAiSettingsDto) GetVectorizationNeedResetOk() (*bool, bool)`

GetVectorizationNeedResetOk returns a tuple with the VectorizationNeedReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorizationNeedReset

`func (o *AiAiSettingsDto) SetVectorizationNeedReset(v bool)`

SetVectorizationNeedReset sets VectorizationNeedReset field to given value.

### HasVectorizationNeedReset

`func (o *AiAiSettingsDto) HasVectorizationNeedReset() bool`

HasVectorizationNeedReset returns a boolean if a field has been set.

### GetAiReady

`func (o *AiAiSettingsDto) GetAiReady() bool`

GetAiReady returns the AiReady field if non-nil, zero value otherwise.

### GetAiReadyOk

`func (o *AiAiSettingsDto) GetAiReadyOk() (*bool, bool)`

GetAiReadyOk returns a tuple with the AiReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiReady

`func (o *AiAiSettingsDto) SetAiReady(v bool)`

SetAiReady sets AiReady field to given value.

### HasAiReady

`func (o *AiAiSettingsDto) HasAiReady() bool`

HasAiReady returns a boolean if a field has been set.

### GetEmbeddingModel

`func (o *AiAiSettingsDto) GetEmbeddingModel() string`

GetEmbeddingModel returns the EmbeddingModel field if non-nil, zero value otherwise.

### GetEmbeddingModelOk

`func (o *AiAiSettingsDto) GetEmbeddingModelOk() (*string, bool)`

GetEmbeddingModelOk returns a tuple with the EmbeddingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingModel

`func (o *AiAiSettingsDto) SetEmbeddingModel(v string)`

SetEmbeddingModel sets EmbeddingModel field to given value.


### SetEmbeddingModelNil

`func (o *AiAiSettingsDto) SetEmbeddingModelNil(b bool)`

 SetEmbeddingModelNil sets the value for EmbeddingModel to be an explicit nil

### UnsetEmbeddingModel
`func (o *AiAiSettingsDto) UnsetEmbeddingModel()`

UnsetEmbeddingModel ensures that no value is present for EmbeddingModel, not even an explicit nil
### GetSystemAiEnabled

`func (o *AiAiSettingsDto) GetSystemAiEnabled() bool`

GetSystemAiEnabled returns the SystemAiEnabled field if non-nil, zero value otherwise.

### GetSystemAiEnabledOk

`func (o *AiAiSettingsDto) GetSystemAiEnabledOk() (*bool, bool)`

GetSystemAiEnabledOk returns a tuple with the SystemAiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAiEnabled

`func (o *AiAiSettingsDto) SetSystemAiEnabled(v bool)`

SetSystemAiEnabled sets SystemAiEnabled field to given value.

### HasSystemAiEnabled

`func (o *AiAiSettingsDto) HasSystemAiEnabled() bool`

HasSystemAiEnabled returns a boolean if a field has been set.

### GetRecommendedModelForForms

`func (o *AiAiSettingsDto) GetRecommendedModelForForms() string`

GetRecommendedModelForForms returns the RecommendedModelForForms field if non-nil, zero value otherwise.

### GetRecommendedModelForFormsOk

`func (o *AiAiSettingsDto) GetRecommendedModelForFormsOk() (*string, bool)`

GetRecommendedModelForFormsOk returns a tuple with the RecommendedModelForForms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedModelForForms

`func (o *AiAiSettingsDto) SetRecommendedModelForForms(v string)`

SetRecommendedModelForForms sets RecommendedModelForForms field to given value.

### HasRecommendedModelForForms

`func (o *AiAiSettingsDto) HasRecommendedModelForForms() bool`

HasRecommendedModelForForms returns a boolean if a field has been set.

### SetRecommendedModelForFormsNil

`func (o *AiAiSettingsDto) SetRecommendedModelForFormsNil(b bool)`

 SetRecommendedModelForFormsNil sets the value for RecommendedModelForForms to be an explicit nil

### UnsetRecommendedModelForForms
`func (o *AiAiSettingsDto) UnsetRecommendedModelForForms()`

UnsetRecommendedModelForForms ensures that no value is present for RecommendedModelForForms, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


