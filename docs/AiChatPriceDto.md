# AiChatPriceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prompt** | Pointer to **float64** | The cost of one million tokens sent to the model, which includes the conversation history resent with  every turn and not just the newest message. | [optional] 
**Completion** | Pointer to **float64** | The cost of one million tokens the model writes back. It is normally the dearer of the two directions. | [optional] 
**PromptCacheRead** | Pointer to **NullableFloat64** | The cost of one million prompt tokens served from the prompt cache. It is absent when the model does not  support prompt caching. | [optional] 
**PromptCacheWrite** | Pointer to **NullableFloat64** | The cost of one million prompt tokens written to the prompt cache with the default lifetime. It is absent  when the model does not support prompt caching. | [optional] 
**PromptCacheWrite1H** | Pointer to **NullableFloat64** | The cost of one million prompt tokens written to the prompt cache with a one-hour lifetime. It is absent  when the model offers no such option. | [optional] 

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

### GetPromptCacheRead

`func (o *AiChatPriceDto) GetPromptCacheRead() float64`

GetPromptCacheRead returns the PromptCacheRead field if non-nil, zero value otherwise.

### GetPromptCacheReadOk

`func (o *AiChatPriceDto) GetPromptCacheReadOk() (*float64, bool)`

GetPromptCacheReadOk returns a tuple with the PromptCacheRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptCacheRead

`func (o *AiChatPriceDto) SetPromptCacheRead(v float64)`

SetPromptCacheRead sets PromptCacheRead field to given value.

### HasPromptCacheRead

`func (o *AiChatPriceDto) HasPromptCacheRead() bool`

HasPromptCacheRead returns a boolean if a field has been set.

### SetPromptCacheReadNil

`func (o *AiChatPriceDto) SetPromptCacheReadNil(b bool)`

 SetPromptCacheReadNil sets the value for PromptCacheRead to be an explicit nil

### UnsetPromptCacheRead
`func (o *AiChatPriceDto) UnsetPromptCacheRead()`

UnsetPromptCacheRead ensures that no value is present for PromptCacheRead, not even an explicit nil
### GetPromptCacheWrite

`func (o *AiChatPriceDto) GetPromptCacheWrite() float64`

GetPromptCacheWrite returns the PromptCacheWrite field if non-nil, zero value otherwise.

### GetPromptCacheWriteOk

`func (o *AiChatPriceDto) GetPromptCacheWriteOk() (*float64, bool)`

GetPromptCacheWriteOk returns a tuple with the PromptCacheWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptCacheWrite

`func (o *AiChatPriceDto) SetPromptCacheWrite(v float64)`

SetPromptCacheWrite sets PromptCacheWrite field to given value.

### HasPromptCacheWrite

`func (o *AiChatPriceDto) HasPromptCacheWrite() bool`

HasPromptCacheWrite returns a boolean if a field has been set.

### SetPromptCacheWriteNil

`func (o *AiChatPriceDto) SetPromptCacheWriteNil(b bool)`

 SetPromptCacheWriteNil sets the value for PromptCacheWrite to be an explicit nil

### UnsetPromptCacheWrite
`func (o *AiChatPriceDto) UnsetPromptCacheWrite()`

UnsetPromptCacheWrite ensures that no value is present for PromptCacheWrite, not even an explicit nil
### GetPromptCacheWrite1H

`func (o *AiChatPriceDto) GetPromptCacheWrite1H() float64`

GetPromptCacheWrite1H returns the PromptCacheWrite1H field if non-nil, zero value otherwise.

### GetPromptCacheWrite1HOk

`func (o *AiChatPriceDto) GetPromptCacheWrite1HOk() (*float64, bool)`

GetPromptCacheWrite1HOk returns a tuple with the PromptCacheWrite1H field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptCacheWrite1H

`func (o *AiChatPriceDto) SetPromptCacheWrite1H(v float64)`

SetPromptCacheWrite1H sets PromptCacheWrite1H field to given value.

### HasPromptCacheWrite1H

`func (o *AiChatPriceDto) HasPromptCacheWrite1H() bool`

HasPromptCacheWrite1H returns a boolean if a field has been set.

### SetPromptCacheWrite1HNil

`func (o *AiChatPriceDto) SetPromptCacheWrite1HNil(b bool)`

 SetPromptCacheWrite1HNil sets the value for PromptCacheWrite1H to be an explicit nil

### UnsetPromptCacheWrite1H
`func (o *AiChatPriceDto) UnsetPromptCacheWrite1H()`

UnsetPromptCacheWrite1H ensures that no value is present for PromptCacheWrite1H, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


