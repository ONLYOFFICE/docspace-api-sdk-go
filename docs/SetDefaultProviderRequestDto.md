# SetDefaultProviderRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | Pointer to **int32** | AI provider identifier. | [optional] 
**DefaultModel** | **NullableString** | Default model identifier to use with this provider. | 

## Methods

### NewSetDefaultProviderRequestDto

`func NewSetDefaultProviderRequestDto(defaultModel NullableString, ) *SetDefaultProviderRequestDto`

NewSetDefaultProviderRequestDto instantiates a new SetDefaultProviderRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetDefaultProviderRequestDtoWithDefaults

`func NewSetDefaultProviderRequestDtoWithDefaults() *SetDefaultProviderRequestDto`

NewSetDefaultProviderRequestDtoWithDefaults instantiates a new SetDefaultProviderRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *SetDefaultProviderRequestDto) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *SetDefaultProviderRequestDto) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *SetDefaultProviderRequestDto) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *SetDefaultProviderRequestDto) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetDefaultModel

`func (o *SetDefaultProviderRequestDto) GetDefaultModel() string`

GetDefaultModel returns the DefaultModel field if non-nil, zero value otherwise.

### GetDefaultModelOk

`func (o *SetDefaultProviderRequestDto) GetDefaultModelOk() (*string, bool)`

GetDefaultModelOk returns a tuple with the DefaultModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultModel

`func (o *SetDefaultProviderRequestDto) SetDefaultModel(v string)`

SetDefaultModel sets DefaultModel field to given value.


### SetDefaultModelNil

`func (o *SetDefaultProviderRequestDto) SetDefaultModelNil(b bool)`

 SetDefaultModelNil sets the value for DefaultModel to be an explicit nil

### UnsetDefaultModel
`func (o *SetDefaultProviderRequestDto) UnsetDefaultModel()`

UnsetDefaultModel ensures that no value is present for DefaultModel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


