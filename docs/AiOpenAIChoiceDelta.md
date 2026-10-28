# AiOpenAIChoiceDelta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** | Sent on the first chunk only, always `assistant`. | [optional] 
**Content** | Pointer to **NullableString** | The text this chunk appends. Null when the chunk carries no text. | [optional] 
**ToolCalls** | Pointer to [**[]AiOpenAIToolCallDelta**](AiOpenAIToolCallDelta.md) | The tool calls the model requested, emitted in place of text. | [optional] 

## Methods

### NewAiOpenAIChoiceDelta

`func NewAiOpenAIChoiceDelta() *AiOpenAIChoiceDelta`

NewAiOpenAIChoiceDelta instantiates a new AiOpenAIChoiceDelta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiOpenAIChoiceDeltaWithDefaults

`func NewAiOpenAIChoiceDeltaWithDefaults() *AiOpenAIChoiceDelta`

NewAiOpenAIChoiceDeltaWithDefaults instantiates a new AiOpenAIChoiceDelta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *AiOpenAIChoiceDelta) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AiOpenAIChoiceDelta) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AiOpenAIChoiceDelta) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AiOpenAIChoiceDelta) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetContent

`func (o *AiOpenAIChoiceDelta) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AiOpenAIChoiceDelta) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AiOpenAIChoiceDelta) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *AiOpenAIChoiceDelta) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *AiOpenAIChoiceDelta) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *AiOpenAIChoiceDelta) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetToolCalls

`func (o *AiOpenAIChoiceDelta) GetToolCalls() []AiOpenAIToolCallDelta`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *AiOpenAIChoiceDelta) GetToolCallsOk() (*[]AiOpenAIToolCallDelta, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *AiOpenAIChoiceDelta) SetToolCalls(v []AiOpenAIToolCallDelta)`

SetToolCalls sets ToolCalls field to given value.

### HasToolCalls

`func (o *AiOpenAIChoiceDelta) HasToolCalls() bool`

HasToolCalls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


