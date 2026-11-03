# AiThreadsAppendUserMessage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | [**AiThreadMessageLike**](AiThreadMessageLike.md) | The stored message itself - with its storage-assigned `id` and `createdAt` - despite the field name, which is kept for the clients that already read it. The message ID is `messageId.id`. | 

## Methods

### NewAiThreadsAppendUserMessage200Response

`func NewAiThreadsAppendUserMessage200Response(messageId AiThreadMessageLike, ) *AiThreadsAppendUserMessage200Response`

NewAiThreadsAppendUserMessage200Response instantiates a new AiThreadsAppendUserMessage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiThreadsAppendUserMessage200ResponseWithDefaults

`func NewAiThreadsAppendUserMessage200ResponseWithDefaults() *AiThreadsAppendUserMessage200Response`

NewAiThreadsAppendUserMessage200ResponseWithDefaults instantiates a new AiThreadsAppendUserMessage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *AiThreadsAppendUserMessage200Response) GetMessageId() AiThreadMessageLike`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *AiThreadsAppendUserMessage200Response) GetMessageIdOk() (*AiThreadMessageLike, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *AiThreadsAppendUserMessage200Response) SetMessageId(v AiThreadMessageLike)`

SetMessageId sets MessageId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


