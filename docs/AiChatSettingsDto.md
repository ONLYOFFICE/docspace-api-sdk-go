# AiChatSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prompt** | Pointer to **NullableString** | The system prompt for the chat. | [optional] 

## Methods

### NewAiChatSettingsDto

`func NewAiChatSettingsDto() *AiChatSettingsDto`

NewAiChatSettingsDto instantiates a new AiChatSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiChatSettingsDtoWithDefaults

`func NewAiChatSettingsDtoWithDefaults() *AiChatSettingsDto`

NewAiChatSettingsDtoWithDefaults instantiates a new AiChatSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrompt

`func (o *AiChatSettingsDto) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AiChatSettingsDto) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AiChatSettingsDto) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *AiChatSettingsDto) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### SetPromptNil

`func (o *AiChatSettingsDto) SetPromptNil(b bool)`

 SetPromptNil sets the value for Prompt to be an explicit nil

### UnsetPrompt
`func (o *AiChatSettingsDto) UnsetPrompt()`

UnsetPrompt ensures that no value is present for Prompt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


