# AiRegenerateStreamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadId** | **string** | Target thread (must already exist). | 
**MessageId** | Pointer to **string** | The user message to re-roll from: everything stored after it is deleted. When omitted, or no longer in the thread, the last user message is used. | [optional] 
**KeepCount** | Pointer to **int32** | How many messages from the start of the thread to keep, used when `messageId` does not resolve. The message at `keepCount - 1` has to be a user message; otherwise the value is ignored. | [optional] 
**ExpectedLength** | Pointer to **int32** | The thread length the caller saw when it computed `keepCount`. When it differs from the stored length, `keepCount` is discarded and the last user message is used instead. | [optional] 
**ActionArgs** | Pointer to [**AiActionArgs**](AiActionArgs.md) | Per-request options of the round: extra tools, reasoning, a prompt override and the connected context. | [optional] 
**EntityId** | Pointer to **string** | Optional entity (room) scope for profile resolution. | [optional] 
**ContextEntityId** | Pointer to **string** | The folder or room the user is in. Tools run against it and the model is told it is the user's location; falls back to `entityId` when omitted. | [optional] 
**ProfileId** | Pointer to **string** | Session-level profile override for this request only. | [optional] 

## Methods

### NewAiRegenerateStreamRequest

`func NewAiRegenerateStreamRequest(threadId string, ) *AiRegenerateStreamRequest`

NewAiRegenerateStreamRequest instantiates a new AiRegenerateStreamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiRegenerateStreamRequestWithDefaults

`func NewAiRegenerateStreamRequestWithDefaults() *AiRegenerateStreamRequest`

NewAiRegenerateStreamRequestWithDefaults instantiates a new AiRegenerateStreamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadId

`func (o *AiRegenerateStreamRequest) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *AiRegenerateStreamRequest) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *AiRegenerateStreamRequest) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### GetMessageId

`func (o *AiRegenerateStreamRequest) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *AiRegenerateStreamRequest) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *AiRegenerateStreamRequest) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *AiRegenerateStreamRequest) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetKeepCount

`func (o *AiRegenerateStreamRequest) GetKeepCount() int32`

GetKeepCount returns the KeepCount field if non-nil, zero value otherwise.

### GetKeepCountOk

`func (o *AiRegenerateStreamRequest) GetKeepCountOk() (*int32, bool)`

GetKeepCountOk returns a tuple with the KeepCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepCount

`func (o *AiRegenerateStreamRequest) SetKeepCount(v int32)`

SetKeepCount sets KeepCount field to given value.

### HasKeepCount

`func (o *AiRegenerateStreamRequest) HasKeepCount() bool`

HasKeepCount returns a boolean if a field has been set.

### GetExpectedLength

`func (o *AiRegenerateStreamRequest) GetExpectedLength() int32`

GetExpectedLength returns the ExpectedLength field if non-nil, zero value otherwise.

### GetExpectedLengthOk

`func (o *AiRegenerateStreamRequest) GetExpectedLengthOk() (*int32, bool)`

GetExpectedLengthOk returns a tuple with the ExpectedLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedLength

`func (o *AiRegenerateStreamRequest) SetExpectedLength(v int32)`

SetExpectedLength sets ExpectedLength field to given value.

### HasExpectedLength

`func (o *AiRegenerateStreamRequest) HasExpectedLength() bool`

HasExpectedLength returns a boolean if a field has been set.

### GetActionArgs

`func (o *AiRegenerateStreamRequest) GetActionArgs() AiActionArgs`

GetActionArgs returns the ActionArgs field if non-nil, zero value otherwise.

### GetActionArgsOk

`func (o *AiRegenerateStreamRequest) GetActionArgsOk() (*AiActionArgs, bool)`

GetActionArgsOk returns a tuple with the ActionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionArgs

`func (o *AiRegenerateStreamRequest) SetActionArgs(v AiActionArgs)`

SetActionArgs sets ActionArgs field to given value.

### HasActionArgs

`func (o *AiRegenerateStreamRequest) HasActionArgs() bool`

HasActionArgs returns a boolean if a field has been set.

### GetEntityId

`func (o *AiRegenerateStreamRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiRegenerateStreamRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiRegenerateStreamRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiRegenerateStreamRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetContextEntityId

`func (o *AiRegenerateStreamRequest) GetContextEntityId() string`

GetContextEntityId returns the ContextEntityId field if non-nil, zero value otherwise.

### GetContextEntityIdOk

`func (o *AiRegenerateStreamRequest) GetContextEntityIdOk() (*string, bool)`

GetContextEntityIdOk returns a tuple with the ContextEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextEntityId

`func (o *AiRegenerateStreamRequest) SetContextEntityId(v string)`

SetContextEntityId sets ContextEntityId field to given value.

### HasContextEntityId

`func (o *AiRegenerateStreamRequest) HasContextEntityId() bool`

HasContextEntityId returns a boolean if a field has been set.

### GetProfileId

`func (o *AiRegenerateStreamRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AiRegenerateStreamRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AiRegenerateStreamRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *AiRegenerateStreamRequest) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


