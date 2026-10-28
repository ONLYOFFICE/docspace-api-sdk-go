# ChatSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | Pointer to **int32** | The provider ID. | [optional] 
**ModelId** | Pointer to **NullableString** | The model ID. | [optional] 
**Prompt** | Pointer to **NullableString** | The prompt. | [optional] 
**Internal** | Pointer to **bool** | Specifies whether the provider is internal or not. | [optional] [readonly] 

## Methods

### NewChatSettings

`func NewChatSettings() *ChatSettings`

NewChatSettings instantiates a new ChatSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatSettingsWithDefaults

`func NewChatSettingsWithDefaults() *ChatSettings`

NewChatSettingsWithDefaults instantiates a new ChatSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *ChatSettings) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ChatSettings) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ChatSettings) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *ChatSettings) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetModelId

`func (o *ChatSettings) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *ChatSettings) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *ChatSettings) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *ChatSettings) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### SetModelIdNil

`func (o *ChatSettings) SetModelIdNil(b bool)`

 SetModelIdNil sets the value for ModelId to be an explicit nil

### UnsetModelId
`func (o *ChatSettings) UnsetModelId()`

UnsetModelId ensures that no value is present for ModelId, not even an explicit nil
### GetPrompt

`func (o *ChatSettings) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *ChatSettings) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *ChatSettings) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *ChatSettings) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### SetPromptNil

`func (o *ChatSettings) SetPromptNil(b bool)`

 SetPromptNil sets the value for Prompt to be an explicit nil

### UnsetPrompt
`func (o *ChatSettings) UnsetPrompt()`

UnsetPrompt ensures that no value is present for Prompt, not even an explicit nil
### GetInternal

`func (o *ChatSettings) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *ChatSettings) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *ChatSettings) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *ChatSettings) HasInternal() bool`

HasInternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


