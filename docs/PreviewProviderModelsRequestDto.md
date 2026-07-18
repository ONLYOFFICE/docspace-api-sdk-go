# PreviewProviderModelsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ProviderType**](ProviderType.md) |  | [optional] 
**Url** | Pointer to **NullableString** | The API endpoint URL. Required for OpenAiCompatible type; optional for other types that have default URLs. | [optional] 
**Key** | **NullableString** | The authentication API key for the AI provider. | 

## Methods

### NewPreviewProviderModelsRequestDto

`func NewPreviewProviderModelsRequestDto(key NullableString, ) *PreviewProviderModelsRequestDto`

NewPreviewProviderModelsRequestDto instantiates a new PreviewProviderModelsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewProviderModelsRequestDtoWithDefaults

`func NewPreviewProviderModelsRequestDtoWithDefaults() *PreviewProviderModelsRequestDto`

NewPreviewProviderModelsRequestDtoWithDefaults instantiates a new PreviewProviderModelsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PreviewProviderModelsRequestDto) GetType() ProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PreviewProviderModelsRequestDto) GetTypeOk() (*ProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PreviewProviderModelsRequestDto) SetType(v ProviderType)`

SetType sets Type field to given value.

### HasType

`func (o *PreviewProviderModelsRequestDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *PreviewProviderModelsRequestDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PreviewProviderModelsRequestDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PreviewProviderModelsRequestDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PreviewProviderModelsRequestDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *PreviewProviderModelsRequestDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *PreviewProviderModelsRequestDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetKey

`func (o *PreviewProviderModelsRequestDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PreviewProviderModelsRequestDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PreviewProviderModelsRequestDto) SetKey(v string)`

SetKey sets Key field to given value.


### SetKeyNil

`func (o *PreviewProviderModelsRequestDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *PreviewProviderModelsRequestDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


