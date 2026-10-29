# AiApproveToolCallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **interface{}** |  | 
**AllowAlways** | Pointer to **bool** | Persist auto-approve for this tool's name. | [optional] 
**ThreadId** | **string** | Thread the assistant message belongs to. | 
**MessageId** | **string** | Storage id of the assistant message holding the tool call. | 
**Idx** | **float32** | Index of the tool-call content part inside `message.content`. | 
**Message** | Pointer to [**AiThreadMessageLike**](AiThreadMessageLike.md) | Snapshot of the assistant message at the time the tool call surfaced. Optional and ignored: the engine resumes from the stored message, and the copy inflated the body past the parser limit (Bug 84608). | [optional] 
**ActionArgs** | Pointer to [**AiActionArgs**](AiActionArgs.md) | Per-request engine options: extra tools, reasoning, prompt override. | [optional] 
**EntityId** | Pointer to **string** | Optional entity (room) scope for profile resolution. | [optional] 
**ProfileId** | Pointer to **string** | Session-level profile override for this request only. | [optional] 

## Methods

### NewAiApproveToolCallRequest

`func NewAiApproveToolCallRequest(result interface{}, threadId string, messageId string, idx float32, ) *AiApproveToolCallRequest`

NewAiApproveToolCallRequest instantiates a new AiApproveToolCallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiApproveToolCallRequestWithDefaults

`func NewAiApproveToolCallRequestWithDefaults() *AiApproveToolCallRequest`

NewAiApproveToolCallRequestWithDefaults instantiates a new AiApproveToolCallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *AiApproveToolCallRequest) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AiApproveToolCallRequest) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AiApproveToolCallRequest) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *AiApproveToolCallRequest) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *AiApproveToolCallRequest) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetAllowAlways

`func (o *AiApproveToolCallRequest) GetAllowAlways() bool`

GetAllowAlways returns the AllowAlways field if non-nil, zero value otherwise.

### GetAllowAlwaysOk

`func (o *AiApproveToolCallRequest) GetAllowAlwaysOk() (*bool, bool)`

GetAllowAlwaysOk returns a tuple with the AllowAlways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAlways

`func (o *AiApproveToolCallRequest) SetAllowAlways(v bool)`

SetAllowAlways sets AllowAlways field to given value.

### HasAllowAlways

`func (o *AiApproveToolCallRequest) HasAllowAlways() bool`

HasAllowAlways returns a boolean if a field has been set.

### GetThreadId

`func (o *AiApproveToolCallRequest) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *AiApproveToolCallRequest) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *AiApproveToolCallRequest) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### GetMessageId

`func (o *AiApproveToolCallRequest) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *AiApproveToolCallRequest) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *AiApproveToolCallRequest) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetIdx

`func (o *AiApproveToolCallRequest) GetIdx() float32`

GetIdx returns the Idx field if non-nil, zero value otherwise.

### GetIdxOk

`func (o *AiApproveToolCallRequest) GetIdxOk() (*float32, bool)`

GetIdxOk returns a tuple with the Idx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdx

`func (o *AiApproveToolCallRequest) SetIdx(v float32)`

SetIdx sets Idx field to given value.


### GetMessage

`func (o *AiApproveToolCallRequest) GetMessage() AiThreadMessageLike`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiApproveToolCallRequest) GetMessageOk() (*AiThreadMessageLike, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiApproveToolCallRequest) SetMessage(v AiThreadMessageLike)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AiApproveToolCallRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetActionArgs

`func (o *AiApproveToolCallRequest) GetActionArgs() AiActionArgs`

GetActionArgs returns the ActionArgs field if non-nil, zero value otherwise.

### GetActionArgsOk

`func (o *AiApproveToolCallRequest) GetActionArgsOk() (*AiActionArgs, bool)`

GetActionArgsOk returns a tuple with the ActionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionArgs

`func (o *AiApproveToolCallRequest) SetActionArgs(v AiActionArgs)`

SetActionArgs sets ActionArgs field to given value.

### HasActionArgs

`func (o *AiApproveToolCallRequest) HasActionArgs() bool`

HasActionArgs returns a boolean if a field has been set.

### GetEntityId

`func (o *AiApproveToolCallRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiApproveToolCallRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiApproveToolCallRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiApproveToolCallRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetProfileId

`func (o *AiApproveToolCallRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AiApproveToolCallRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AiApproveToolCallRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *AiApproveToolCallRequest) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


