# AiActionArgs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tools** | Pointer to [**[]AiMCPItem**](AiMCPItem.md) | Extra tools offered to the model for this request. | [optional] 
**IsReasoning** | Pointer to **bool** | Legacy extended-thinking switch; stands for `medium`. `reasoningLevel` wins when both are set. | [optional] 
**ReasoningLevel** | Pointer to [**AiReasoningLevel**](AiReasoningLevel.md) | Depth of extended thinking for the round; providers clamp it to what the model accepts. | [optional] 
**Prompt** | Pointer to [**AiActionArgsPrompt**](AiActionArgsPrompt.md) |  | [optional] 

## Methods

### NewAiActionArgs

`func NewAiActionArgs() *AiActionArgs`

NewAiActionArgs instantiates a new AiActionArgs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiActionArgsWithDefaults

`func NewAiActionArgsWithDefaults() *AiActionArgs`

NewAiActionArgsWithDefaults instantiates a new AiActionArgs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTools

`func (o *AiActionArgs) GetTools() []AiMCPItem`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AiActionArgs) GetToolsOk() (*[]AiMCPItem, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AiActionArgs) SetTools(v []AiMCPItem)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AiActionArgs) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetIsReasoning

`func (o *AiActionArgs) GetIsReasoning() bool`

GetIsReasoning returns the IsReasoning field if non-nil, zero value otherwise.

### GetIsReasoningOk

`func (o *AiActionArgs) GetIsReasoningOk() (*bool, bool)`

GetIsReasoningOk returns a tuple with the IsReasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReasoning

`func (o *AiActionArgs) SetIsReasoning(v bool)`

SetIsReasoning sets IsReasoning field to given value.

### HasIsReasoning

`func (o *AiActionArgs) HasIsReasoning() bool`

HasIsReasoning returns a boolean if a field has been set.

### GetReasoningLevel

`func (o *AiActionArgs) GetReasoningLevel() AiReasoningLevel`

GetReasoningLevel returns the ReasoningLevel field if non-nil, zero value otherwise.

### GetReasoningLevelOk

`func (o *AiActionArgs) GetReasoningLevelOk() (*AiReasoningLevel, bool)`

GetReasoningLevelOk returns a tuple with the ReasoningLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoningLevel

`func (o *AiActionArgs) SetReasoningLevel(v AiReasoningLevel)`

SetReasoningLevel sets ReasoningLevel field to given value.

### HasReasoningLevel

`func (o *AiActionArgs) HasReasoningLevel() bool`

HasReasoningLevel returns a boolean if a field has been set.

### GetPrompt

`func (o *AiActionArgs) GetPrompt() AiActionArgsPrompt`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AiActionArgs) GetPromptOk() (*AiActionArgsPrompt, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AiActionArgs) SetPrompt(v AiActionArgsPrompt)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *AiActionArgs) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


