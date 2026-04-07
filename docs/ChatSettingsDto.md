# ChatSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | Pointer to **int32** | The AI provider ID. | [optional] 
**ModelId** | Pointer to **NullableString** | The AI model ID used for chat completions. | [optional] 
**ModelAlias** | Pointer to **NullableString** | The AI model display alias. | [optional] 
**Prompt** | Pointer to **NullableString** | The system prompt for the chat. | [optional] 
**Multimodal** | Pointer to [**ChatMultimodalSettingsDto**](ChatMultimodalSettingsDto.md) |  | [optional] 
**Thinking** | Pointer to **bool** | Indicates whether the model supports extended thinking mode. | [optional] 
**Internal** | Pointer to **bool** | Indicates whether this is an internal AI gateway provider. | [optional] [readonly] 

## Methods

### NewChatSettingsDto

`func NewChatSettingsDto() *ChatSettingsDto`

NewChatSettingsDto instantiates a new ChatSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatSettingsDtoWithDefaults

`func NewChatSettingsDtoWithDefaults() *ChatSettingsDto`

NewChatSettingsDtoWithDefaults instantiates a new ChatSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *ChatSettingsDto) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ChatSettingsDto) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ChatSettingsDto) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *ChatSettingsDto) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetModelId

`func (o *ChatSettingsDto) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *ChatSettingsDto) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *ChatSettingsDto) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *ChatSettingsDto) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### SetModelIdNil

`func (o *ChatSettingsDto) SetModelIdNil(b bool)`

 SetModelIdNil sets the value for ModelId to be an explicit nil

### UnsetModelId
`func (o *ChatSettingsDto) UnsetModelId()`

UnsetModelId ensures that no value is present for ModelId, not even an explicit nil
### GetModelAlias

`func (o *ChatSettingsDto) GetModelAlias() string`

GetModelAlias returns the ModelAlias field if non-nil, zero value otherwise.

### GetModelAliasOk

`func (o *ChatSettingsDto) GetModelAliasOk() (*string, bool)`

GetModelAliasOk returns a tuple with the ModelAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelAlias

`func (o *ChatSettingsDto) SetModelAlias(v string)`

SetModelAlias sets ModelAlias field to given value.

### HasModelAlias

`func (o *ChatSettingsDto) HasModelAlias() bool`

HasModelAlias returns a boolean if a field has been set.

### SetModelAliasNil

`func (o *ChatSettingsDto) SetModelAliasNil(b bool)`

 SetModelAliasNil sets the value for ModelAlias to be an explicit nil

### UnsetModelAlias
`func (o *ChatSettingsDto) UnsetModelAlias()`

UnsetModelAlias ensures that no value is present for ModelAlias, not even an explicit nil
### GetPrompt

`func (o *ChatSettingsDto) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *ChatSettingsDto) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *ChatSettingsDto) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *ChatSettingsDto) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### SetPromptNil

`func (o *ChatSettingsDto) SetPromptNil(b bool)`

 SetPromptNil sets the value for Prompt to be an explicit nil

### UnsetPrompt
`func (o *ChatSettingsDto) UnsetPrompt()`

UnsetPrompt ensures that no value is present for Prompt, not even an explicit nil
### GetMultimodal

`func (o *ChatSettingsDto) GetMultimodal() ChatMultimodalSettingsDto`

GetMultimodal returns the Multimodal field if non-nil, zero value otherwise.

### GetMultimodalOk

`func (o *ChatSettingsDto) GetMultimodalOk() (*ChatMultimodalSettingsDto, bool)`

GetMultimodalOk returns a tuple with the Multimodal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultimodal

`func (o *ChatSettingsDto) SetMultimodal(v ChatMultimodalSettingsDto)`

SetMultimodal sets Multimodal field to given value.

### HasMultimodal

`func (o *ChatSettingsDto) HasMultimodal() bool`

HasMultimodal returns a boolean if a field has been set.

### GetThinking

`func (o *ChatSettingsDto) GetThinking() bool`

GetThinking returns the Thinking field if non-nil, zero value otherwise.

### GetThinkingOk

`func (o *ChatSettingsDto) GetThinkingOk() (*bool, bool)`

GetThinkingOk returns a tuple with the Thinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinking

`func (o *ChatSettingsDto) SetThinking(v bool)`

SetThinking sets Thinking field to given value.

### HasThinking

`func (o *ChatSettingsDto) HasThinking() bool`

HasThinking returns a boolean if a field has been set.

### GetInternal

`func (o *ChatSettingsDto) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *ChatSettingsDto) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *ChatSettingsDto) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *ChatSettingsDto) HasInternal() bool`

HasInternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


