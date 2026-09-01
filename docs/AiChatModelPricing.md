# AiChatModelPricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The identifier of the model, as the provider expects it on the wire. | 
**Alias** | Pointer to **NullableString** | The display name of the model. | [optional] 
**OwnedBy** | Pointer to **NullableString** | The owner of the model, as reported by the provider. | [optional] 
**Provider** | Pointer to **NullableString** | The provider that serves the model. | [optional] 
**Link** | Pointer to **NullableString** | The link to the pricing page of the model. | [optional] 
**Price** | [**AiChatPrice**](AiChatPrice.md) | The price of a chat model, per token. | 

## Methods

### NewAiChatModelPricing

`func NewAiChatModelPricing(id NullableString, price AiChatPrice, ) *AiChatModelPricing`

NewAiChatModelPricing instantiates a new AiChatModelPricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiChatModelPricingWithDefaults

`func NewAiChatModelPricingWithDefaults() *AiChatModelPricing`

NewAiChatModelPricingWithDefaults instantiates a new AiChatModelPricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiChatModelPricing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiChatModelPricing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiChatModelPricing) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *AiChatModelPricing) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AiChatModelPricing) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAlias

`func (o *AiChatModelPricing) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *AiChatModelPricing) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *AiChatModelPricing) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *AiChatModelPricing) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *AiChatModelPricing) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *AiChatModelPricing) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetOwnedBy

`func (o *AiChatModelPricing) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *AiChatModelPricing) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *AiChatModelPricing) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *AiChatModelPricing) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### SetOwnedByNil

`func (o *AiChatModelPricing) SetOwnedByNil(b bool)`

 SetOwnedByNil sets the value for OwnedBy to be an explicit nil

### UnsetOwnedBy
`func (o *AiChatModelPricing) UnsetOwnedBy()`

UnsetOwnedBy ensures that no value is present for OwnedBy, not even an explicit nil
### GetProvider

`func (o *AiChatModelPricing) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AiChatModelPricing) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AiChatModelPricing) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AiChatModelPricing) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *AiChatModelPricing) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *AiChatModelPricing) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetLink

`func (o *AiChatModelPricing) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AiChatModelPricing) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AiChatModelPricing) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *AiChatModelPricing) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *AiChatModelPricing) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *AiChatModelPricing) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetPrice

`func (o *AiChatModelPricing) GetPrice() AiChatPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AiChatModelPricing) GetPriceOk() (*AiChatPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AiChatModelPricing) SetPrice(v AiChatPrice)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


