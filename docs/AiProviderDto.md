# AiProviderDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | AI provider identifier. | [optional] 
**Title** | **NullableString** | AI provider display title. | 
**Type** | Pointer to [**ProviderType**](ProviderType.md) |  | [optional] 
**Url** | Pointer to **NullableString** | API endpoint URL for the AI provider. | [optional] 
**CreatedOn** | [**ApiDateTime**](ApiDateTime.md) |  | 
**ModifiedOn** | [**ApiDateTime**](ApiDateTime.md) |  | 
**NeedReset** | Pointer to **bool** | Indicates whether the provider's API key needs to be reset. | [optional] 
**IsDefault** | Pointer to **bool** | Indicates whether this provider is the default provider for the tenant. | [optional] 

## Methods

### NewAiProviderDto

`func NewAiProviderDto(title NullableString, createdOn ApiDateTime, modifiedOn ApiDateTime, ) *AiProviderDto`

NewAiProviderDto instantiates a new AiProviderDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiProviderDtoWithDefaults

`func NewAiProviderDtoWithDefaults() *AiProviderDto`

NewAiProviderDtoWithDefaults instantiates a new AiProviderDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiProviderDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiProviderDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiProviderDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AiProviderDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *AiProviderDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AiProviderDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AiProviderDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *AiProviderDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AiProviderDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetType

`func (o *AiProviderDto) GetType() ProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiProviderDto) GetTypeOk() (*ProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiProviderDto) SetType(v ProviderType)`

SetType sets Type field to given value.

### HasType

`func (o *AiProviderDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *AiProviderDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AiProviderDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AiProviderDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AiProviderDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *AiProviderDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AiProviderDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetCreatedOn

`func (o *AiProviderDto) GetCreatedOn() ApiDateTime`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *AiProviderDto) GetCreatedOnOk() (*ApiDateTime, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *AiProviderDto) SetCreatedOn(v ApiDateTime)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedOn

`func (o *AiProviderDto) GetModifiedOn() ApiDateTime`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *AiProviderDto) GetModifiedOnOk() (*ApiDateTime, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *AiProviderDto) SetModifiedOn(v ApiDateTime)`

SetModifiedOn sets ModifiedOn field to given value.


### GetNeedReset

`func (o *AiProviderDto) GetNeedReset() bool`

GetNeedReset returns the NeedReset field if non-nil, zero value otherwise.

### GetNeedResetOk

`func (o *AiProviderDto) GetNeedResetOk() (*bool, bool)`

GetNeedResetOk returns a tuple with the NeedReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedReset

`func (o *AiProviderDto) SetNeedReset(v bool)`

SetNeedReset sets NeedReset field to given value.

### HasNeedReset

`func (o *AiProviderDto) HasNeedReset() bool`

HasNeedReset returns a boolean if a field has been set.

### GetIsDefault

`func (o *AiProviderDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *AiProviderDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *AiProviderDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *AiProviderDto) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


