# AiThreadsAppendUserMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadId** | **string** |  | 
**Message** | [**AiThreadMessageLike**](AiThreadMessageLike.md) | Message to persist (id/createdAt are storage-assigned). | 
**ProfileId** | Pointer to **string** |  | [optional] 

## Methods

### NewAiThreadsAppendUserMessageRequest

`func NewAiThreadsAppendUserMessageRequest(threadId string, message AiThreadMessageLike, ) *AiThreadsAppendUserMessageRequest`

NewAiThreadsAppendUserMessageRequest instantiates a new AiThreadsAppendUserMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiThreadsAppendUserMessageRequestWithDefaults

`func NewAiThreadsAppendUserMessageRequestWithDefaults() *AiThreadsAppendUserMessageRequest`

NewAiThreadsAppendUserMessageRequestWithDefaults instantiates a new AiThreadsAppendUserMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadId

`func (o *AiThreadsAppendUserMessageRequest) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *AiThreadsAppendUserMessageRequest) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *AiThreadsAppendUserMessageRequest) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### GetMessage

`func (o *AiThreadsAppendUserMessageRequest) GetMessage() AiThreadMessageLike`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiThreadsAppendUserMessageRequest) GetMessageOk() (*AiThreadMessageLike, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiThreadsAppendUserMessageRequest) SetMessage(v AiThreadMessageLike)`

SetMessage sets Message field to given value.


### GetProfileId

`func (o *AiThreadsAppendUserMessageRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AiThreadsAppendUserMessageRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AiThreadsAppendUserMessageRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *AiThreadsAppendUserMessageRequest) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


