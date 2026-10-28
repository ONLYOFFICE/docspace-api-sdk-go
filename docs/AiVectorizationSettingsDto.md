# AiVectorizationSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**AiEmbeddingProviderType**](AiEmbeddingProviderType.md) | The type of embedding provider configured for document vectorization. | [optional] 
**NeedReset** | Pointer to **bool** | Indicates whether the embedding provider API key needs to be reconfigured. | [optional] 

## Methods

### NewAiVectorizationSettingsDto

`func NewAiVectorizationSettingsDto() *AiVectorizationSettingsDto`

NewAiVectorizationSettingsDto instantiates a new AiVectorizationSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiVectorizationSettingsDtoWithDefaults

`func NewAiVectorizationSettingsDtoWithDefaults() *AiVectorizationSettingsDto`

NewAiVectorizationSettingsDtoWithDefaults instantiates a new AiVectorizationSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AiVectorizationSettingsDto) GetType() AiEmbeddingProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiVectorizationSettingsDto) GetTypeOk() (*AiEmbeddingProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiVectorizationSettingsDto) SetType(v AiEmbeddingProviderType)`

SetType sets Type field to given value.

### HasType

`func (o *AiVectorizationSettingsDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNeedReset

`func (o *AiVectorizationSettingsDto) GetNeedReset() bool`

GetNeedReset returns the NeedReset field if non-nil, zero value otherwise.

### GetNeedResetOk

`func (o *AiVectorizationSettingsDto) GetNeedResetOk() (*bool, bool)`

GetNeedResetOk returns a tuple with the NeedReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedReset

`func (o *AiVectorizationSettingsDto) SetNeedReset(v bool)`

SetNeedReset sets NeedReset field to given value.

### HasNeedReset

`func (o *AiVectorizationSettingsDto) HasNeedReset() bool`

HasNeedReset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


