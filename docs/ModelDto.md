# ModelDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | Pointer to **int32** | The unique identifier of the AI provider that offers this model. | [optional] 
**ProviderTitle** | **NullableString** | The human-readable display name of the AI provider (e.g., OpenAI, Anthropic). | 
**ModelId** | **NullableString** | The model identifier as recognized by the AI provider (e.g., gpt-4o, claude-sonnet-4-20250514). | 
**Alias** | Pointer to **NullableString** | The display name for the model. | [optional] 
**Capabilities** | Pointer to [**AiModelCapabilities**](AiModelCapabilities.md) |  | [optional] 
**Price** | Pointer to [**AiChatPrice**](AiChatPrice.md) |  | [optional] 
**Currency** | Pointer to [**CurrencyInfo**](CurrencyInfo.md) |  | [optional] 

## Methods

### NewModelDto

`func NewModelDto(providerTitle NullableString, modelId NullableString, ) *ModelDto`

NewModelDto instantiates a new ModelDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelDtoWithDefaults

`func NewModelDtoWithDefaults() *ModelDto`

NewModelDtoWithDefaults instantiates a new ModelDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *ModelDto) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ModelDto) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ModelDto) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *ModelDto) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetProviderTitle

`func (o *ModelDto) GetProviderTitle() string`

GetProviderTitle returns the ProviderTitle field if non-nil, zero value otherwise.

### GetProviderTitleOk

`func (o *ModelDto) GetProviderTitleOk() (*string, bool)`

GetProviderTitleOk returns a tuple with the ProviderTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderTitle

`func (o *ModelDto) SetProviderTitle(v string)`

SetProviderTitle sets ProviderTitle field to given value.


### SetProviderTitleNil

`func (o *ModelDto) SetProviderTitleNil(b bool)`

 SetProviderTitleNil sets the value for ProviderTitle to be an explicit nil

### UnsetProviderTitle
`func (o *ModelDto) UnsetProviderTitle()`

UnsetProviderTitle ensures that no value is present for ProviderTitle, not even an explicit nil
### GetModelId

`func (o *ModelDto) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *ModelDto) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *ModelDto) SetModelId(v string)`

SetModelId sets ModelId field to given value.


### SetModelIdNil

`func (o *ModelDto) SetModelIdNil(b bool)`

 SetModelIdNil sets the value for ModelId to be an explicit nil

### UnsetModelId
`func (o *ModelDto) UnsetModelId()`

UnsetModelId ensures that no value is present for ModelId, not even an explicit nil
### GetAlias

`func (o *ModelDto) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ModelDto) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ModelDto) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ModelDto) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *ModelDto) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *ModelDto) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetCapabilities

`func (o *ModelDto) GetCapabilities() AiModelCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ModelDto) GetCapabilitiesOk() (*AiModelCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ModelDto) SetCapabilities(v AiModelCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ModelDto) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetPrice

`func (o *ModelDto) GetPrice() AiChatPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ModelDto) GetPriceOk() (*AiChatPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ModelDto) SetPrice(v AiChatPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ModelDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *ModelDto) GetCurrency() CurrencyInfo`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ModelDto) GetCurrencyOk() (*CurrencyInfo, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ModelDto) SetCurrency(v CurrencyInfo)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ModelDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


