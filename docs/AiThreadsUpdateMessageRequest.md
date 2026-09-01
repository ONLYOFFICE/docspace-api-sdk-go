# AiThreadsUpdateMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **string** |  | 
**Message** | [**AiThreadMessageLike**](AiThreadMessageLike.md) | Replacement message content. | 

## Methods

### NewAiThreadsUpdateMessageRequest

`func NewAiThreadsUpdateMessageRequest(messageId string, message AiThreadMessageLike, ) *AiThreadsUpdateMessageRequest`

NewAiThreadsUpdateMessageRequest instantiates a new AiThreadsUpdateMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiThreadsUpdateMessageRequestWithDefaults

`func NewAiThreadsUpdateMessageRequestWithDefaults() *AiThreadsUpdateMessageRequest`

NewAiThreadsUpdateMessageRequestWithDefaults instantiates a new AiThreadsUpdateMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *AiThreadsUpdateMessageRequest) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *AiThreadsUpdateMessageRequest) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *AiThreadsUpdateMessageRequest) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetMessage

`func (o *AiThreadsUpdateMessageRequest) GetMessage() AiThreadMessageLike`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiThreadsUpdateMessageRequest) GetMessageOk() (*AiThreadMessageLike, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiThreadsUpdateMessageRequest) SetMessage(v AiThreadMessageLike)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


