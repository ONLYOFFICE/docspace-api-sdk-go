# AiAiApproveToolCallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**AllowAlways** | Pointer to **bool** | Persist auto-approve for this tool's name. | [optional] 
**ThreadId** | **string** | Thread the assistant message belongs to. | 
**MessageId** | **string** | Storage id of the assistant message holding the tool call. | 
**Idx** | **float32** | Index of the tool-call content part inside `message.content`. | 
**Message** | [**AiThreadMessageLike**](AiThreadMessageLike.md) | Snapshot of the assistant message at the time the tool call surfaced. | 
**ActionArgs** | Pointer to [**AiAiActionArgs**](AiAiActionArgs.md) | Per-request engine options: extra tools, reasoning, prompt override. | [optional] 
**EntityId** | Pointer to **string** | Optional entity (room) scope for profile resolution. | [optional] 
**ProfileId** | Pointer to **string** | Session-level profile override for this request only. | [optional] 

## Methods

### NewAiAiApproveToolCallRequest

`func NewAiAiApproveToolCallRequest(result interface{}, threadId string, messageId string, idx float32, message AiThreadMessageLike, ) *AiAiApproveToolCallRequest`

NewAiAiApproveToolCallRequest instantiates a new AiAiApproveToolCallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAiApproveToolCallRequestWithDefaults

`func NewAiAiApproveToolCallRequestWithDefaults() *AiAiApproveToolCallRequest`

NewAiAiApproveToolCallRequestWithDefaults instantiates a new AiAiApproveToolCallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *AiAiApproveToolCallRequest) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AiAiApproveToolCallRequest) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AiAiApproveToolCallRequest) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *AiAiApproveToolCallRequest) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *AiAiApproveToolCallRequest) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetAllowAlways

`func (o *AiAiApproveToolCallRequest) GetAllowAlways() bool`

GetAllowAlways returns the AllowAlways field if non-nil, zero value otherwise.

### GetAllowAlwaysOk

`func (o *AiAiApproveToolCallRequest) GetAllowAlwaysOk() (*bool, bool)`

GetAllowAlwaysOk returns a tuple with the AllowAlways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAlways

`func (o *AiAiApproveToolCallRequest) SetAllowAlways(v bool)`

SetAllowAlways sets AllowAlways field to given value.

### HasAllowAlways

`func (o *AiAiApproveToolCallRequest) HasAllowAlways() bool`

HasAllowAlways returns a boolean if a field has been set.

### GetThreadId

`func (o *AiAiApproveToolCallRequest) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *AiAiApproveToolCallRequest) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *AiAiApproveToolCallRequest) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### GetMessageId

`func (o *AiAiApproveToolCallRequest) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *AiAiApproveToolCallRequest) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *AiAiApproveToolCallRequest) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetIdx

`func (o *AiAiApproveToolCallRequest) GetIdx() float32`

GetIdx returns the Idx field if non-nil, zero value otherwise.

### GetIdxOk

`func (o *AiAiApproveToolCallRequest) GetIdxOk() (*float32, bool)`

GetIdxOk returns a tuple with the Idx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdx

`func (o *AiAiApproveToolCallRequest) SetIdx(v float32)`

SetIdx sets Idx field to given value.


### GetMessage

`func (o *AiAiApproveToolCallRequest) GetMessage() AiThreadMessageLike`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiAiApproveToolCallRequest) GetMessageOk() (*AiThreadMessageLike, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiAiApproveToolCallRequest) SetMessage(v AiThreadMessageLike)`

SetMessage sets Message field to given value.


### GetActionArgs

`func (o *AiAiApproveToolCallRequest) GetActionArgs() AiAiActionArgs`

GetActionArgs returns the ActionArgs field if non-nil, zero value otherwise.

### GetActionArgsOk

`func (o *AiAiApproveToolCallRequest) GetActionArgsOk() (*AiAiActionArgs, bool)`

GetActionArgsOk returns a tuple with the ActionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionArgs

`func (o *AiAiApproveToolCallRequest) SetActionArgs(v AiAiActionArgs)`

SetActionArgs sets ActionArgs field to given value.

### HasActionArgs

`func (o *AiAiApproveToolCallRequest) HasActionArgs() bool`

HasActionArgs returns a boolean if a field has been set.

### GetEntityId

`func (o *AiAiApproveToolCallRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiAiApproveToolCallRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiAiApproveToolCallRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiAiApproveToolCallRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetProfileId

`func (o *AiAiApproveToolCallRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AiAiApproveToolCallRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AiAiApproveToolCallRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *AiAiApproveToolCallRequest) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


