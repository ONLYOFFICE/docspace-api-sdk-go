# AiSendStreamBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadId** | Pointer to **string** | Target thread; a new one is created (with an auto title) when omitted. | [optional] 
**UserMessage** | [**AiThreadMessageLike**](AiThreadMessageLike.md) | The user turn to send. | 
**ActionArgs** | Pointer to [**AiActionArgs**](AiActionArgs.md) | Per-request engine options: extra tools, reasoning, prompt override. | [optional] 
**EntityId** | Pointer to **string** | Optional entity (room) scope for profile resolution. | [optional] 
**ProfileId** | Pointer to **string** | Session-level profile override for this request only. | [optional] 

## Methods

### NewAiSendStreamBody

`func NewAiSendStreamBody(userMessage AiThreadMessageLike, ) *AiSendStreamBody`

NewAiSendStreamBody instantiates a new AiSendStreamBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSendStreamBodyWithDefaults

`func NewAiSendStreamBodyWithDefaults() *AiSendStreamBody`

NewAiSendStreamBodyWithDefaults instantiates a new AiSendStreamBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadId

`func (o *AiSendStreamBody) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *AiSendStreamBody) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *AiSendStreamBody) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *AiSendStreamBody) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetUserMessage

`func (o *AiSendStreamBody) GetUserMessage() AiThreadMessageLike`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### GetUserMessageOk

`func (o *AiSendStreamBody) GetUserMessageOk() (*AiThreadMessageLike, bool)`

GetUserMessageOk returns a tuple with the UserMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMessage

`func (o *AiSendStreamBody) SetUserMessage(v AiThreadMessageLike)`

SetUserMessage sets UserMessage field to given value.


### GetActionArgs

`func (o *AiSendStreamBody) GetActionArgs() AiActionArgs`

GetActionArgs returns the ActionArgs field if non-nil, zero value otherwise.

### GetActionArgsOk

`func (o *AiSendStreamBody) GetActionArgsOk() (*AiActionArgs, bool)`

GetActionArgsOk returns a tuple with the ActionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionArgs

`func (o *AiSendStreamBody) SetActionArgs(v AiActionArgs)`

SetActionArgs sets ActionArgs field to given value.

### HasActionArgs

`func (o *AiSendStreamBody) HasActionArgs() bool`

HasActionArgs returns a boolean if a field has been set.

### GetEntityId

`func (o *AiSendStreamBody) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiSendStreamBody) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiSendStreamBody) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiSendStreamBody) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetProfileId

`func (o *AiSendStreamBody) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AiSendStreamBody) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AiSendStreamBody) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *AiSendStreamBody) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


