# AiThreadMessageLikeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | `running`, `requires-action`, `complete` or `incomplete`. | 
**Reason** | Pointer to **string** | Why the message is in that state: `tool-calls` or `interrupt` for `requires-action`; `stop` or `unknown` for `complete`; `cancelled`, `tool-calls`, `length`, `content-filter`, `other` or `error` for `incomplete`. An `incomplete` message the chat stored carries `cancelled` (the round was stopped) or `error`. | [optional] 
**Error** | Pointer to [**AiChatErrorPayload**](AiChatErrorPayload.md) | What failed, on an `incomplete` message whose `reason` is `error`. | [optional] 

## Methods

### NewAiThreadMessageLikeStatus

`func NewAiThreadMessageLikeStatus(type_ string, ) *AiThreadMessageLikeStatus`

NewAiThreadMessageLikeStatus instantiates a new AiThreadMessageLikeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiThreadMessageLikeStatusWithDefaults

`func NewAiThreadMessageLikeStatusWithDefaults() *AiThreadMessageLikeStatus`

NewAiThreadMessageLikeStatusWithDefaults instantiates a new AiThreadMessageLikeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AiThreadMessageLikeStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiThreadMessageLikeStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiThreadMessageLikeStatus) SetType(v string)`

SetType sets Type field to given value.


### GetReason

`func (o *AiThreadMessageLikeStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AiThreadMessageLikeStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AiThreadMessageLikeStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AiThreadMessageLikeStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetError

`func (o *AiThreadMessageLikeStatus) GetError() AiChatErrorPayload`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AiThreadMessageLikeStatus) GetErrorOk() (*AiChatErrorPayload, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AiThreadMessageLikeStatus) SetError(v AiChatErrorPayload)`

SetError sets Error field to given value.

### HasError

`func (o *AiThreadMessageLikeStatus) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


