# VectorizationSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**EmbeddingProviderType**](EmbeddingProviderType.md) |  | [optional] 
**NeedReset** | Pointer to **bool** | Indicates whether the embedding provider API key needs to be reconfigured. | [optional] 

## Methods

### NewVectorizationSettingsDto

`func NewVectorizationSettingsDto() *VectorizationSettingsDto`

NewVectorizationSettingsDto instantiates a new VectorizationSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorizationSettingsDtoWithDefaults

`func NewVectorizationSettingsDtoWithDefaults() *VectorizationSettingsDto`

NewVectorizationSettingsDtoWithDefaults instantiates a new VectorizationSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VectorizationSettingsDto) GetType() EmbeddingProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VectorizationSettingsDto) GetTypeOk() (*EmbeddingProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VectorizationSettingsDto) SetType(v EmbeddingProviderType)`

SetType sets Type field to given value.

### HasType

`func (o *VectorizationSettingsDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNeedReset

`func (o *VectorizationSettingsDto) GetNeedReset() bool`

GetNeedReset returns the NeedReset field if non-nil, zero value otherwise.

### GetNeedResetOk

`func (o *VectorizationSettingsDto) GetNeedResetOk() (*bool, bool)`

GetNeedResetOk returns a tuple with the NeedReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedReset

`func (o *VectorizationSettingsDto) SetNeedReset(v bool)`

SetNeedReset sets NeedReset field to given value.

### HasNeedReset

`func (o *VectorizationSettingsDto) HasNeedReset() bool`

HasNeedReset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


