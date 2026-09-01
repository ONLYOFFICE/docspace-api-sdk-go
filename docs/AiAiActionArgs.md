# AiAiActionArgs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tools** | Pointer to [**[]AiTMCPItem**](AiTMCPItem.md) | Extra tools offered to the model for this request. | [optional] 
**IsReasoning** | Pointer to **bool** | Enable extended thinking / reasoning for this request. | [optional] 
**Prompt** | Pointer to [**AiAiActionArgsPrompt**](AiAiActionArgsPrompt.md) |  | [optional] 

## Methods

### NewAiAiActionArgs

`func NewAiAiActionArgs() *AiAiActionArgs`

NewAiAiActionArgs instantiates a new AiAiActionArgs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAiActionArgsWithDefaults

`func NewAiAiActionArgsWithDefaults() *AiAiActionArgs`

NewAiAiActionArgsWithDefaults instantiates a new AiAiActionArgs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTools

`func (o *AiAiActionArgs) GetTools() []AiTMCPItem`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AiAiActionArgs) GetToolsOk() (*[]AiTMCPItem, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AiAiActionArgs) SetTools(v []AiTMCPItem)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AiAiActionArgs) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetIsReasoning

`func (o *AiAiActionArgs) GetIsReasoning() bool`

GetIsReasoning returns the IsReasoning field if non-nil, zero value otherwise.

### GetIsReasoningOk

`func (o *AiAiActionArgs) GetIsReasoningOk() (*bool, bool)`

GetIsReasoningOk returns a tuple with the IsReasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReasoning

`func (o *AiAiActionArgs) SetIsReasoning(v bool)`

SetIsReasoning sets IsReasoning field to given value.

### HasIsReasoning

`func (o *AiAiActionArgs) HasIsReasoning() bool`

HasIsReasoning returns a boolean if a field has been set.

### GetPrompt

`func (o *AiAiActionArgs) GetPrompt() AiAiActionArgsPrompt`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AiAiActionArgs) GetPromptOk() (*AiAiActionArgsPrompt, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AiAiActionArgs) SetPrompt(v AiAiActionArgsPrompt)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *AiAiActionArgs) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


