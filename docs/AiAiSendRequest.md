# AiAiSendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | [**AiActionType**](AiActionType.md) | Which AI action to run — selects the assignment slot and action. | 
**UserMessage** | [**AiThreadMessageLike**](AiThreadMessageLike.md) | The user turn to send. | 
**ActionArgs** | Pointer to [**AiAiActionArgs**](AiAiActionArgs.md) | Per-request engine options: extra tools, reasoning, prompt override. | [optional] 
**EntityId** | Pointer to **string** | Optional entity (room) scope for profile resolution. | [optional] 

## Methods

### NewAiAiSendRequest

`func NewAiAiSendRequest(actionType AiActionType, userMessage AiThreadMessageLike, ) *AiAiSendRequest`

NewAiAiSendRequest instantiates a new AiAiSendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAiSendRequestWithDefaults

`func NewAiAiSendRequestWithDefaults() *AiAiSendRequest`

NewAiAiSendRequestWithDefaults instantiates a new AiAiSendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *AiAiSendRequest) GetActionType() AiActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *AiAiSendRequest) GetActionTypeOk() (*AiActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *AiAiSendRequest) SetActionType(v AiActionType)`

SetActionType sets ActionType field to given value.


### GetUserMessage

`func (o *AiAiSendRequest) GetUserMessage() AiThreadMessageLike`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### GetUserMessageOk

`func (o *AiAiSendRequest) GetUserMessageOk() (*AiThreadMessageLike, bool)`

GetUserMessageOk returns a tuple with the UserMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMessage

`func (o *AiAiSendRequest) SetUserMessage(v AiThreadMessageLike)`

SetUserMessage sets UserMessage field to given value.


### GetActionArgs

`func (o *AiAiSendRequest) GetActionArgs() AiAiActionArgs`

GetActionArgs returns the ActionArgs field if non-nil, zero value otherwise.

### GetActionArgsOk

`func (o *AiAiSendRequest) GetActionArgsOk() (*AiAiActionArgs, bool)`

GetActionArgsOk returns a tuple with the ActionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionArgs

`func (o *AiAiSendRequest) SetActionArgs(v AiAiActionArgs)`

SetActionArgs sets ActionArgs field to given value.

### HasActionArgs

`func (o *AiAiSendRequest) HasActionArgs() bool`

HasActionArgs returns a boolean if a field has been set.

### GetEntityId

`func (o *AiAiSendRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiAiSendRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiAiSendRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiAiSendRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


