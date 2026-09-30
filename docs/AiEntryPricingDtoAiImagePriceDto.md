# AiEntryPricingDtoAiImagePriceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The model identifier to send to the AI operations. It is the value to branch on, while `alias` is for display  only. | 
**Alias** | **NullableString** | The model name as the vendor writes it, meant to be shown to a person rather than matched on. | 
**Provider** | **NullableString** | Who runs the model. Two entries can share a provider, and one provider's models can be priced quite  differently, so the price always belongs to the entry and never to the provider. | 
**Image** | **NullableString** | The absolute URL of the provider's icon, for rendering next to the entry. | 
**Price** | [**AiImagePriceDto**](AiImagePriceDto.md) | What the entry costs, in the currency the answer names. Amounts per token are normalised per million  tokens, so they are not the price of a single call. | 
**Link** | **NullableString** | The provider's own page for the model, for a person to read the model's terms. It is empty when the  provider publishes none. | 

## Methods

### NewAiEntryPricingDtoAiImagePriceDto

`func NewAiEntryPricingDtoAiImagePriceDto(id NullableString, alias NullableString, provider NullableString, image NullableString, price AiImagePriceDto, link NullableString, ) *AiEntryPricingDtoAiImagePriceDto`

NewAiEntryPricingDtoAiImagePriceDto instantiates a new AiEntryPricingDtoAiImagePriceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiEntryPricingDtoAiImagePriceDtoWithDefaults

`func NewAiEntryPricingDtoAiImagePriceDtoWithDefaults() *AiEntryPricingDtoAiImagePriceDto`

NewAiEntryPricingDtoAiImagePriceDtoWithDefaults instantiates a new AiEntryPricingDtoAiImagePriceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiEntryPricingDtoAiImagePriceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiEntryPricingDtoAiImagePriceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiEntryPricingDtoAiImagePriceDto) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *AiEntryPricingDtoAiImagePriceDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AiEntryPricingDtoAiImagePriceDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAlias

`func (o *AiEntryPricingDtoAiImagePriceDto) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *AiEntryPricingDtoAiImagePriceDto) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *AiEntryPricingDtoAiImagePriceDto) SetAlias(v string)`

SetAlias sets Alias field to given value.


### SetAliasNil

`func (o *AiEntryPricingDtoAiImagePriceDto) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *AiEntryPricingDtoAiImagePriceDto) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetProvider

`func (o *AiEntryPricingDtoAiImagePriceDto) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AiEntryPricingDtoAiImagePriceDto) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AiEntryPricingDtoAiImagePriceDto) SetProvider(v string)`

SetProvider sets Provider field to given value.


### SetProviderNil

`func (o *AiEntryPricingDtoAiImagePriceDto) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *AiEntryPricingDtoAiImagePriceDto) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetImage

`func (o *AiEntryPricingDtoAiImagePriceDto) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AiEntryPricingDtoAiImagePriceDto) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AiEntryPricingDtoAiImagePriceDto) SetImage(v string)`

SetImage sets Image field to given value.


### SetImageNil

`func (o *AiEntryPricingDtoAiImagePriceDto) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *AiEntryPricingDtoAiImagePriceDto) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetPrice

`func (o *AiEntryPricingDtoAiImagePriceDto) GetPrice() AiImagePriceDto`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AiEntryPricingDtoAiImagePriceDto) GetPriceOk() (*AiImagePriceDto, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AiEntryPricingDtoAiImagePriceDto) SetPrice(v AiImagePriceDto)`

SetPrice sets Price field to given value.


### GetLink

`func (o *AiEntryPricingDtoAiImagePriceDto) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AiEntryPricingDtoAiImagePriceDto) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AiEntryPricingDtoAiImagePriceDto) SetLink(v string)`

SetLink sets Link field to given value.


### SetLinkNil

`func (o *AiEntryPricingDtoAiImagePriceDto) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *AiEntryPricingDtoAiImagePriceDto) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


