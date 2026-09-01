# AiAiToolCallData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadId** | **string** | Thread the assistant message belongs to. | 
**MessageId** | **string** | Storage id of the assistant message holding the tool call. | 
**Idx** | **float32** | Index of the tool-call content part inside `message.content`. | 
**Message** | [**AiThreadMessageLike**](AiThreadMessageLike.md) | Snapshot of the assistant message at the time the tool call surfaced. | 
**ActionArgs** | Pointer to [**AiAiActionArgs**](AiAiActionArgs.md) | Per-request engine options: extra tools, reasoning, prompt override. | [optional] 
**EntityId** | Pointer to **string** | Optional entity (room) scope for profile resolution. | [optional] 
**ProfileId** | Pointer to **string** | Session-level profile override for this request only. | [optional] 

## Methods

### NewAiAiToolCallData

`func NewAiAiToolCallData(threadId string, messageId string, idx float32, message AiThreadMessageLike, ) *AiAiToolCallData`

NewAiAiToolCallData instantiates a new AiAiToolCallData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAiToolCallDataWithDefaults

`func NewAiAiToolCallDataWithDefaults() *AiAiToolCallData`

NewAiAiToolCallDataWithDefaults instantiates a new AiAiToolCallData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadId

`func (o *AiAiToolCallData) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *AiAiToolCallData) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *AiAiToolCallData) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### GetMessageId

`func (o *AiAiToolCallData) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *AiAiToolCallData) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *AiAiToolCallData) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetIdx

`func (o *AiAiToolCallData) GetIdx() float32`

GetIdx returns the Idx field if non-nil, zero value otherwise.

### GetIdxOk

`func (o *AiAiToolCallData) GetIdxOk() (*float32, bool)`

GetIdxOk returns a tuple with the Idx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdx

`func (o *AiAiToolCallData) SetIdx(v float32)`

SetIdx sets Idx field to given value.


### GetMessage

`func (o *AiAiToolCallData) GetMessage() AiThreadMessageLike`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiAiToolCallData) GetMessageOk() (*AiThreadMessageLike, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiAiToolCallData) SetMessage(v AiThreadMessageLike)`

SetMessage sets Message field to given value.


### GetActionArgs

`func (o *AiAiToolCallData) GetActionArgs() AiAiActionArgs`

GetActionArgs returns the ActionArgs field if non-nil, zero value otherwise.

### GetActionArgsOk

`func (o *AiAiToolCallData) GetActionArgsOk() (*AiAiActionArgs, bool)`

GetActionArgsOk returns a tuple with the ActionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionArgs

`func (o *AiAiToolCallData) SetActionArgs(v AiAiActionArgs)`

SetActionArgs sets ActionArgs field to given value.

### HasActionArgs

`func (o *AiAiToolCallData) HasActionArgs() bool`

HasActionArgs returns a boolean if a field has been set.

### GetEntityId

`func (o *AiAiToolCallData) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiAiToolCallData) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiAiToolCallData) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiAiToolCallData) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetProfileId

`func (o *AiAiToolCallData) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AiAiToolCallData) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AiAiToolCallData) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *AiAiToolCallData) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


