# AiAiSendCustomRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsStream** | **bool** | Stream the reply (ndjson) when true, else return a single message. | 
**SystemPrompt** | **string** | Caller-supplied system prompt for this one-turn call. | 
**UserMessage** | [**AiThreadMessageLike**](AiThreadMessageLike.md) |  | 
**ActionArgs** | Pointer to [**AiAiActionArgs**](AiAiActionArgs.md) | Per-request engine options: extra tools, reasoning, prompt override. | [optional] 

## Methods

### NewAiAiSendCustomRequest

`func NewAiAiSendCustomRequest(isStream bool, systemPrompt string, userMessage AiThreadMessageLike, ) *AiAiSendCustomRequest`

NewAiAiSendCustomRequest instantiates a new AiAiSendCustomRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAiSendCustomRequestWithDefaults

`func NewAiAiSendCustomRequestWithDefaults() *AiAiSendCustomRequest`

NewAiAiSendCustomRequestWithDefaults instantiates a new AiAiSendCustomRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsStream

`func (o *AiAiSendCustomRequest) GetIsStream() bool`

GetIsStream returns the IsStream field if non-nil, zero value otherwise.

### GetIsStreamOk

`func (o *AiAiSendCustomRequest) GetIsStreamOk() (*bool, bool)`

GetIsStreamOk returns a tuple with the IsStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStream

`func (o *AiAiSendCustomRequest) SetIsStream(v bool)`

SetIsStream sets IsStream field to given value.


### GetSystemPrompt

`func (o *AiAiSendCustomRequest) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *AiAiSendCustomRequest) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *AiAiSendCustomRequest) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.


### GetUserMessage

`func (o *AiAiSendCustomRequest) GetUserMessage() AiThreadMessageLike`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### GetUserMessageOk

`func (o *AiAiSendCustomRequest) GetUserMessageOk() (*AiThreadMessageLike, bool)`

GetUserMessageOk returns a tuple with the UserMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMessage

`func (o *AiAiSendCustomRequest) SetUserMessage(v AiThreadMessageLike)`

SetUserMessage sets UserMessage field to given value.


### GetActionArgs

`func (o *AiAiSendCustomRequest) GetActionArgs() AiAiActionArgs`

GetActionArgs returns the ActionArgs field if non-nil, zero value otherwise.

### GetActionArgsOk

`func (o *AiAiSendCustomRequest) GetActionArgsOk() (*AiAiActionArgs, bool)`

GetActionArgsOk returns a tuple with the ActionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionArgs

`func (o *AiAiSendCustomRequest) SetActionArgs(v AiAiActionArgs)`

SetActionArgs sets ActionArgs field to given value.

### HasActionArgs

`func (o *AiAiSendCustomRequest) HasActionArgs() bool`

HasActionArgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


