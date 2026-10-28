# AiOpenOrCreateResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadId** | **string** | The thread that was opened, or the one just created. | 
**Title** | **string** | Empty string for existing threads — the engine doesn't re-fetch. | 
**PriorMessages** | [**[]AiThreadMessageLike**](AiThreadMessageLike.md) | The messages already in the thread - empty for a thread that was just created. | 

## Methods

### NewAiOpenOrCreateResult

`func NewAiOpenOrCreateResult(threadId string, title string, priorMessages []AiThreadMessageLike, ) *AiOpenOrCreateResult`

NewAiOpenOrCreateResult instantiates a new AiOpenOrCreateResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiOpenOrCreateResultWithDefaults

`func NewAiOpenOrCreateResultWithDefaults() *AiOpenOrCreateResult`

NewAiOpenOrCreateResultWithDefaults instantiates a new AiOpenOrCreateResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadId

`func (o *AiOpenOrCreateResult) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *AiOpenOrCreateResult) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *AiOpenOrCreateResult) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### GetTitle

`func (o *AiOpenOrCreateResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AiOpenOrCreateResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AiOpenOrCreateResult) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPriorMessages

`func (o *AiOpenOrCreateResult) GetPriorMessages() []AiThreadMessageLike`

GetPriorMessages returns the PriorMessages field if non-nil, zero value otherwise.

### GetPriorMessagesOk

`func (o *AiOpenOrCreateResult) GetPriorMessagesOk() (*[]AiThreadMessageLike, bool)`

GetPriorMessagesOk returns a tuple with the PriorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorMessages

`func (o *AiOpenOrCreateResult) SetPriorMessages(v []AiThreadMessageLike)`

SetPriorMessages sets PriorMessages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


