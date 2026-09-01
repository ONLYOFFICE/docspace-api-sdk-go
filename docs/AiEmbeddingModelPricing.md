# AiEmbeddingModelPricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The identifier of the model, as the provider expects it on the wire. | 
**Alias** | Pointer to **NullableString** | The display name of the model. | [optional] 
**OwnedBy** | Pointer to **NullableString** | The owner of the model, as reported by the provider. | [optional] 
**Provider** | Pointer to **NullableString** | The provider that serves the model. | [optional] 
**Link** | Pointer to **NullableString** | The link to the pricing page of the model. | [optional] 
**Price** | [**AiEmbeddingPrice**](AiEmbeddingPrice.md) | The price of an embedding model, per token. | 

## Methods

### NewAiEmbeddingModelPricing

`func NewAiEmbeddingModelPricing(id NullableString, price AiEmbeddingPrice, ) *AiEmbeddingModelPricing`

NewAiEmbeddingModelPricing instantiates a new AiEmbeddingModelPricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiEmbeddingModelPricingWithDefaults

`func NewAiEmbeddingModelPricingWithDefaults() *AiEmbeddingModelPricing`

NewAiEmbeddingModelPricingWithDefaults instantiates a new AiEmbeddingModelPricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiEmbeddingModelPricing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiEmbeddingModelPricing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiEmbeddingModelPricing) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *AiEmbeddingModelPricing) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AiEmbeddingModelPricing) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAlias

`func (o *AiEmbeddingModelPricing) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *AiEmbeddingModelPricing) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *AiEmbeddingModelPricing) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *AiEmbeddingModelPricing) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *AiEmbeddingModelPricing) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *AiEmbeddingModelPricing) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetOwnedBy

`func (o *AiEmbeddingModelPricing) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *AiEmbeddingModelPricing) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *AiEmbeddingModelPricing) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *AiEmbeddingModelPricing) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### SetOwnedByNil

`func (o *AiEmbeddingModelPricing) SetOwnedByNil(b bool)`

 SetOwnedByNil sets the value for OwnedBy to be an explicit nil

### UnsetOwnedBy
`func (o *AiEmbeddingModelPricing) UnsetOwnedBy()`

UnsetOwnedBy ensures that no value is present for OwnedBy, not even an explicit nil
### GetProvider

`func (o *AiEmbeddingModelPricing) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AiEmbeddingModelPricing) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AiEmbeddingModelPricing) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AiEmbeddingModelPricing) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *AiEmbeddingModelPricing) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *AiEmbeddingModelPricing) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetLink

`func (o *AiEmbeddingModelPricing) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AiEmbeddingModelPricing) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AiEmbeddingModelPricing) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *AiEmbeddingModelPricing) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *AiEmbeddingModelPricing) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *AiEmbeddingModelPricing) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetPrice

`func (o *AiEmbeddingModelPricing) GetPrice() AiEmbeddingPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AiEmbeddingModelPricing) GetPriceOk() (*AiEmbeddingPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AiEmbeddingModelPricing) SetPrice(v AiEmbeddingPrice)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


