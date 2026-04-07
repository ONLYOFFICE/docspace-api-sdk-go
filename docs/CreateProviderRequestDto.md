# CreateProviderRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ProviderType**](ProviderType.md) |  | [optional] 
**Title** | **NullableString** | The display title for the AI provider. | 
**Url** | Pointer to **NullableString** | The API endpoint URL for the AI provider. Required for OpenAiCompatible type; optional for other types that have default URLs. | [optional] 
**Key** | **NullableString** | The authentication API key for the AI provider. | 

## Methods

### NewCreateProviderRequestDto

`func NewCreateProviderRequestDto(title NullableString, key NullableString, ) *CreateProviderRequestDto`

NewCreateProviderRequestDto instantiates a new CreateProviderRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProviderRequestDtoWithDefaults

`func NewCreateProviderRequestDtoWithDefaults() *CreateProviderRequestDto`

NewCreateProviderRequestDtoWithDefaults instantiates a new CreateProviderRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateProviderRequestDto) GetType() ProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateProviderRequestDto) GetTypeOk() (*ProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateProviderRequestDto) SetType(v ProviderType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateProviderRequestDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *CreateProviderRequestDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateProviderRequestDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateProviderRequestDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *CreateProviderRequestDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CreateProviderRequestDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUrl

`func (o *CreateProviderRequestDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateProviderRequestDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateProviderRequestDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateProviderRequestDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *CreateProviderRequestDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CreateProviderRequestDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetKey

`func (o *CreateProviderRequestDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateProviderRequestDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateProviderRequestDto) SetKey(v string)`

SetKey sets Key field to given value.


### SetKeyNil

`func (o *CreateProviderRequestDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *CreateProviderRequestDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


