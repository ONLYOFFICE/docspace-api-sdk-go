# AiAiSendStreamBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadId** | Pointer to **string** | Target thread; a new one is created (with an auto title) when omitted. | [optional] 
**UserMessage** | [**AiThreadMessageLike**](AiThreadMessageLike.md) | The user turn to send. | 
**ActionArgs** | Pointer to [**AiAiActionArgs**](AiAiActionArgs.md) | Per-request engine options: extra tools, reasoning, prompt override. | [optional] 
**EntityId** | Pointer to **string** | Optional entity (room) scope for profile resolution. | [optional] 
**ProfileId** | Pointer to **string** | Session-level profile override for this request only. | [optional] 

## Methods

### NewAiAiSendStreamBody

`func NewAiAiSendStreamBody(userMessage AiThreadMessageLike, ) *AiAiSendStreamBody`

NewAiAiSendStreamBody instantiates a new AiAiSendStreamBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAiSendStreamBodyWithDefaults

`func NewAiAiSendStreamBodyWithDefaults() *AiAiSendStreamBody`

NewAiAiSendStreamBodyWithDefaults instantiates a new AiAiSendStreamBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadId

`func (o *AiAiSendStreamBody) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *AiAiSendStreamBody) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *AiAiSendStreamBody) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *AiAiSendStreamBody) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetUserMessage

`func (o *AiAiSendStreamBody) GetUserMessage() AiThreadMessageLike`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### GetUserMessageOk

`func (o *AiAiSendStreamBody) GetUserMessageOk() (*AiThreadMessageLike, bool)`

GetUserMessageOk returns a tuple with the UserMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMessage

`func (o *AiAiSendStreamBody) SetUserMessage(v AiThreadMessageLike)`

SetUserMessage sets UserMessage field to given value.


### GetActionArgs

`func (o *AiAiSendStreamBody) GetActionArgs() AiAiActionArgs`

GetActionArgs returns the ActionArgs field if non-nil, zero value otherwise.

### GetActionArgsOk

`func (o *AiAiSendStreamBody) GetActionArgsOk() (*AiAiActionArgs, bool)`

GetActionArgsOk returns a tuple with the ActionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionArgs

`func (o *AiAiSendStreamBody) SetActionArgs(v AiAiActionArgs)`

SetActionArgs sets ActionArgs field to given value.

### HasActionArgs

`func (o *AiAiSendStreamBody) HasActionArgs() bool`

HasActionArgs returns a boolean if a field has been set.

### GetEntityId

`func (o *AiAiSendStreamBody) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiAiSendStreamBody) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiAiSendStreamBody) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiAiSendStreamBody) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetProfileId

`func (o *AiAiSendStreamBody) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AiAiSendStreamBody) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AiAiSendStreamBody) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *AiAiSendStreamBody) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


