# AiWebSearchPricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The identifier of the web search provider. | [optional] 
**Provider** | Pointer to **NullableString** | The provider that serves the web search requests. | [optional] 
**Price** | Pointer to **float64** | The price of a single web search request. | [optional] 
**Link** | Pointer to **NullableString** | The link to the pricing page of the provider. | [optional] 

## Methods

### NewAiWebSearchPricing

`func NewAiWebSearchPricing() *AiWebSearchPricing`

NewAiWebSearchPricing instantiates a new AiWebSearchPricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiWebSearchPricingWithDefaults

`func NewAiWebSearchPricingWithDefaults() *AiWebSearchPricing`

NewAiWebSearchPricingWithDefaults instantiates a new AiWebSearchPricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiWebSearchPricing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiWebSearchPricing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiWebSearchPricing) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiWebSearchPricing) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AiWebSearchPricing) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AiWebSearchPricing) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProvider

`func (o *AiWebSearchPricing) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AiWebSearchPricing) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AiWebSearchPricing) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AiWebSearchPricing) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *AiWebSearchPricing) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *AiWebSearchPricing) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetPrice

`func (o *AiWebSearchPricing) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AiWebSearchPricing) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AiWebSearchPricing) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *AiWebSearchPricing) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetLink

`func (o *AiWebSearchPricing) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AiWebSearchPricing) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AiWebSearchPricing) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *AiWebSearchPricing) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *AiWebSearchPricing) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *AiWebSearchPricing) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


