# AiEntryPricingDtoAiEmbeddingPriceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The model identifier to send to the AI operations. It is the value to branch on, while `alias` is for display  only. | 
**Alias** | **NullableString** | The model name as the vendor writes it, meant to be shown to a person rather than matched on. | 
**Provider** | **NullableString** | Who runs the model. Two entries can share a provider, and one provider's models can be priced quite  differently, so the price always belongs to the entry and never to the provider. | 
**Image** | **NullableString** | The absolute URL of the provider's icon, for rendering next to the entry. | 
**Price** | [**AiEmbeddingPriceDto**](AiEmbeddingPriceDto.md) | What the entry costs, in the currency the answer names. Amounts per token are normalised per million  tokens, so they are not the price of a single call. | 
**Link** | **NullableString** | The provider's own page for the model, for a person to read the model's terms. It is empty when the  provider publishes none. | 

## Methods

### NewAiEntryPricingDtoAiEmbeddingPriceDto

`func NewAiEntryPricingDtoAiEmbeddingPriceDto(id NullableString, alias NullableString, provider NullableString, image NullableString, price AiEmbeddingPriceDto, link NullableString, ) *AiEntryPricingDtoAiEmbeddingPriceDto`

NewAiEntryPricingDtoAiEmbeddingPriceDto instantiates a new AiEntryPricingDtoAiEmbeddingPriceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiEntryPricingDtoAiEmbeddingPriceDtoWithDefaults

`func NewAiEntryPricingDtoAiEmbeddingPriceDtoWithDefaults() *AiEntryPricingDtoAiEmbeddingPriceDto`

NewAiEntryPricingDtoAiEmbeddingPriceDtoWithDefaults instantiates a new AiEntryPricingDtoAiEmbeddingPriceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAlias

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) SetAlias(v string)`

SetAlias sets Alias field to given value.


### SetAliasNil

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetProvider

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) SetProvider(v string)`

SetProvider sets Provider field to given value.


### SetProviderNil

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetImage

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) SetImage(v string)`

SetImage sets Image field to given value.


### SetImageNil

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetPrice

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) GetPrice() AiEmbeddingPriceDto`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) GetPriceOk() (*AiEmbeddingPriceDto, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) SetPrice(v AiEmbeddingPriceDto)`

SetPrice sets Price field to given value.


### GetLink

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) SetLink(v string)`

SetLink sets Link field to given value.


### SetLinkNil

`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *AiEntryPricingDtoAiEmbeddingPriceDto) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


