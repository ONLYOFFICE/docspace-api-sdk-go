# DefaultProviderDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | Pointer to **int32** | AI provider identifier. | [optional] 
**DefaultModel** | **NullableString** | Default model identifier used with this provider. | 
**ProviderTitle** | Pointer to **NullableString** | AI provider title. | [optional] 
**ProviderType** | Pointer to [**ProviderType**](ProviderType.md) |  | [optional] 
**DefaultModelAlias** | Pointer to **NullableString** | Display alias of the default model. | [optional] 

## Methods

### NewDefaultProviderDto

`func NewDefaultProviderDto(defaultModel NullableString, ) *DefaultProviderDto`

NewDefaultProviderDto instantiates a new DefaultProviderDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultProviderDtoWithDefaults

`func NewDefaultProviderDtoWithDefaults() *DefaultProviderDto`

NewDefaultProviderDtoWithDefaults instantiates a new DefaultProviderDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *DefaultProviderDto) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *DefaultProviderDto) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *DefaultProviderDto) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *DefaultProviderDto) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetDefaultModel

`func (o *DefaultProviderDto) GetDefaultModel() string`

GetDefaultModel returns the DefaultModel field if non-nil, zero value otherwise.

### GetDefaultModelOk

`func (o *DefaultProviderDto) GetDefaultModelOk() (*string, bool)`

GetDefaultModelOk returns a tuple with the DefaultModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultModel

`func (o *DefaultProviderDto) SetDefaultModel(v string)`

SetDefaultModel sets DefaultModel field to given value.


### SetDefaultModelNil

`func (o *DefaultProviderDto) SetDefaultModelNil(b bool)`

 SetDefaultModelNil sets the value for DefaultModel to be an explicit nil

### UnsetDefaultModel
`func (o *DefaultProviderDto) UnsetDefaultModel()`

UnsetDefaultModel ensures that no value is present for DefaultModel, not even an explicit nil
### GetProviderTitle

`func (o *DefaultProviderDto) GetProviderTitle() string`

GetProviderTitle returns the ProviderTitle field if non-nil, zero value otherwise.

### GetProviderTitleOk

`func (o *DefaultProviderDto) GetProviderTitleOk() (*string, bool)`

GetProviderTitleOk returns a tuple with the ProviderTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderTitle

`func (o *DefaultProviderDto) SetProviderTitle(v string)`

SetProviderTitle sets ProviderTitle field to given value.

### HasProviderTitle

`func (o *DefaultProviderDto) HasProviderTitle() bool`

HasProviderTitle returns a boolean if a field has been set.

### SetProviderTitleNil

`func (o *DefaultProviderDto) SetProviderTitleNil(b bool)`

 SetProviderTitleNil sets the value for ProviderTitle to be an explicit nil

### UnsetProviderTitle
`func (o *DefaultProviderDto) UnsetProviderTitle()`

UnsetProviderTitle ensures that no value is present for ProviderTitle, not even an explicit nil
### GetProviderType

`func (o *DefaultProviderDto) GetProviderType() ProviderType`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *DefaultProviderDto) GetProviderTypeOk() (*ProviderType, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *DefaultProviderDto) SetProviderType(v ProviderType)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *DefaultProviderDto) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetDefaultModelAlias

`func (o *DefaultProviderDto) GetDefaultModelAlias() string`

GetDefaultModelAlias returns the DefaultModelAlias field if non-nil, zero value otherwise.

### GetDefaultModelAliasOk

`func (o *DefaultProviderDto) GetDefaultModelAliasOk() (*string, bool)`

GetDefaultModelAliasOk returns a tuple with the DefaultModelAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultModelAlias

`func (o *DefaultProviderDto) SetDefaultModelAlias(v string)`

SetDefaultModelAlias sets DefaultModelAlias field to given value.

### HasDefaultModelAlias

`func (o *DefaultProviderDto) HasDefaultModelAlias() bool`

HasDefaultModelAlias returns a boolean if a field has been set.

### SetDefaultModelAliasNil

`func (o *DefaultProviderDto) SetDefaultModelAliasNil(b bool)`

 SetDefaultModelAliasNil sets the value for DefaultModelAlias to be an explicit nil

### UnsetDefaultModelAlias
`func (o *DefaultProviderDto) UnsetDefaultModelAlias()`

UnsetDefaultModelAlias ensures that no value is present for DefaultModelAlias, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


