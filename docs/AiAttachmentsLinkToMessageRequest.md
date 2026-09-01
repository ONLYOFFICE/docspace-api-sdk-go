# AiAttachmentsLinkToMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | Attachment ids to bind. | 
**MessageId** | **string** | Owning message id. | 
**ThreadId** | **string** | Owning thread id. | 

## Methods

### NewAiAttachmentsLinkToMessageRequest

`func NewAiAttachmentsLinkToMessageRequest(ids []string, messageId string, threadId string, ) *AiAttachmentsLinkToMessageRequest`

NewAiAttachmentsLinkToMessageRequest instantiates a new AiAttachmentsLinkToMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAttachmentsLinkToMessageRequestWithDefaults

`func NewAiAttachmentsLinkToMessageRequestWithDefaults() *AiAttachmentsLinkToMessageRequest`

NewAiAttachmentsLinkToMessageRequestWithDefaults instantiates a new AiAttachmentsLinkToMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *AiAttachmentsLinkToMessageRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *AiAttachmentsLinkToMessageRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *AiAttachmentsLinkToMessageRequest) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetMessageId

`func (o *AiAttachmentsLinkToMessageRequest) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *AiAttachmentsLinkToMessageRequest) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *AiAttachmentsLinkToMessageRequest) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetThreadId

`func (o *AiAttachmentsLinkToMessageRequest) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *AiAttachmentsLinkToMessageRequest) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *AiAttachmentsLinkToMessageRequest) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


