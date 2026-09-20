# AiEntryPricingDtoDecimal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The model identifier to send to the AI operations. It is the value to branch on, while `alias` is for display  only. | 
**Alias** | **NullableString** | The model name as the vendor writes it, meant to be shown to a person rather than matched on. | 
**Provider** | **NullableString** | Who runs the model. Two entries can share a provider, and one provider's models can be priced quite  differently, so the price always belongs to the entry and never to the provider. | 
**Image** | **NullableString** | The absolute URL of the provider's icon, for rendering next to the entry. | 
**Price** | **float64** | What the entry costs, in the currency the answer names. Amounts per token are normalised per million  tokens, so they are not the price of a single call. | 
**Link** | **NullableString** | The provider's own page for the model, for a person to read the model's terms. It is empty when the  provider publishes none. | 

## Methods

### NewAiEntryPricingDtoDecimal

`func NewAiEntryPricingDtoDecimal(id NullableString, alias NullableString, provider NullableString, image NullableString, price float64, link NullableString, ) *AiEntryPricingDtoDecimal`

NewAiEntryPricingDtoDecimal instantiates a new AiEntryPricingDtoDecimal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiEntryPricingDtoDecimalWithDefaults

`func NewAiEntryPricingDtoDecimalWithDefaults() *AiEntryPricingDtoDecimal`

NewAiEntryPricingDtoDecimalWithDefaults instantiates a new AiEntryPricingDtoDecimal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiEntryPricingDtoDecimal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiEntryPricingDtoDecimal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiEntryPricingDtoDecimal) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *AiEntryPricingDtoDecimal) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AiEntryPricingDtoDecimal) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAlias

`func (o *AiEntryPricingDtoDecimal) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *AiEntryPricingDtoDecimal) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *AiEntryPricingDtoDecimal) SetAlias(v string)`

SetAlias sets Alias field to given value.


### SetAliasNil

`func (o *AiEntryPricingDtoDecimal) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *AiEntryPricingDtoDecimal) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetProvider

`func (o *AiEntryPricingDtoDecimal) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AiEntryPricingDtoDecimal) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AiEntryPricingDtoDecimal) SetProvider(v string)`

SetProvider sets Provider field to given value.


### SetProviderNil

`func (o *AiEntryPricingDtoDecimal) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *AiEntryPricingDtoDecimal) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetImage

`func (o *AiEntryPricingDtoDecimal) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AiEntryPricingDtoDecimal) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AiEntryPricingDtoDecimal) SetImage(v string)`

SetImage sets Image field to given value.


### SetImageNil

`func (o *AiEntryPricingDtoDecimal) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *AiEntryPricingDtoDecimal) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetPrice

`func (o *AiEntryPricingDtoDecimal) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AiEntryPricingDtoDecimal) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AiEntryPricingDtoDecimal) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetLink

`func (o *AiEntryPricingDtoDecimal) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AiEntryPricingDtoDecimal) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AiEntryPricingDtoDecimal) SetLink(v string)`

SetLink sets Link field to given value.


### SetLinkNil

`func (o *AiEntryPricingDtoDecimal) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *AiEntryPricingDtoDecimal) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


