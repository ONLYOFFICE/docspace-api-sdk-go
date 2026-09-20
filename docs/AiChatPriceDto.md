# AiChatPriceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prompt** | Pointer to **float64** | The cost of one million tokens sent to the model, which includes the conversation history resent with  every turn and not just the newest message. | [optional] 
**Completion** | Pointer to **float64** | The cost of one million tokens the model writes back. It is normally the dearer of the two directions. | [optional] 

## Methods

### NewAiChatPriceDto

`func NewAiChatPriceDto() *AiChatPriceDto`

NewAiChatPriceDto instantiates a new AiChatPriceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiChatPriceDtoWithDefaults

`func NewAiChatPriceDtoWithDefaults() *AiChatPriceDto`

NewAiChatPriceDtoWithDefaults instantiates a new AiChatPriceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrompt

`func (o *AiChatPriceDto) GetPrompt() float64`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AiChatPriceDto) GetPromptOk() (*float64, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AiChatPriceDto) SetPrompt(v float64)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *AiChatPriceDto) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetCompletion

`func (o *AiChatPriceDto) GetCompletion() float64`

GetCompletion returns the Completion field if non-nil, zero value otherwise.

### GetCompletionOk

`func (o *AiChatPriceDto) GetCompletionOk() (*float64, bool)`

GetCompletionOk returns a tuple with the Completion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletion

`func (o *AiChatPriceDto) SetCompletion(v float64)`

SetCompletion sets Completion field to given value.

### HasCompletion

`func (o *AiChatPriceDto) HasCompletion() bool`

HasCompletion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


