# SetEmbeddingConfigRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**EmbeddingProviderType**](EmbeddingProviderType.md) |  | [optional] 
**Key** | Pointer to **NullableString** | The API key for the selected embedding provider. Pass null to keep the existing key unchanged. | [optional] 

## Methods

### NewSetEmbeddingConfigRequestBody

`func NewSetEmbeddingConfigRequestBody() *SetEmbeddingConfigRequestBody`

NewSetEmbeddingConfigRequestBody instantiates a new SetEmbeddingConfigRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetEmbeddingConfigRequestBodyWithDefaults

`func NewSetEmbeddingConfigRequestBodyWithDefaults() *SetEmbeddingConfigRequestBody`

NewSetEmbeddingConfigRequestBodyWithDefaults instantiates a new SetEmbeddingConfigRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SetEmbeddingConfigRequestBody) GetType() EmbeddingProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SetEmbeddingConfigRequestBody) GetTypeOk() (*EmbeddingProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SetEmbeddingConfigRequestBody) SetType(v EmbeddingProviderType)`

SetType sets Type field to given value.

### HasType

`func (o *SetEmbeddingConfigRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetKey

`func (o *SetEmbeddingConfigRequestBody) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SetEmbeddingConfigRequestBody) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SetEmbeddingConfigRequestBody) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SetEmbeddingConfigRequestBody) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *SetEmbeddingConfigRequestBody) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *SetEmbeddingConfigRequestBody) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


