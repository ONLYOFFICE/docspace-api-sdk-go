# AiOpenAIChunkChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **float32** | The zero-based position of the choice. This service emits a single choice, so always 0. | 
**Delta** | [**AiOpenAIChoiceDelta**](AiOpenAIChoiceDelta.md) | What this chunk adds to the choice. | 
**FinishReason** | [**NullableAiOpenAIFinishReason**](AiOpenAIFinishReason.md) | Why the completion stopped, or null while it is still streaming. | 

## Methods

### NewAiOpenAIChunkChoice

`func NewAiOpenAIChunkChoice(index float32, delta AiOpenAIChoiceDelta, finishReason NullableAiOpenAIFinishReason, ) *AiOpenAIChunkChoice`

NewAiOpenAIChunkChoice instantiates a new AiOpenAIChunkChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiOpenAIChunkChoiceWithDefaults

`func NewAiOpenAIChunkChoiceWithDefaults() *AiOpenAIChunkChoice`

NewAiOpenAIChunkChoiceWithDefaults instantiates a new AiOpenAIChunkChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *AiOpenAIChunkChoice) GetIndex() float32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *AiOpenAIChunkChoice) GetIndexOk() (*float32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *AiOpenAIChunkChoice) SetIndex(v float32)`

SetIndex sets Index field to given value.


### GetDelta

`func (o *AiOpenAIChunkChoice) GetDelta() AiOpenAIChoiceDelta`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *AiOpenAIChunkChoice) GetDeltaOk() (*AiOpenAIChoiceDelta, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *AiOpenAIChunkChoice) SetDelta(v AiOpenAIChoiceDelta)`

SetDelta sets Delta field to given value.


### GetFinishReason

`func (o *AiOpenAIChunkChoice) GetFinishReason() AiOpenAIFinishReason`

GetFinishReason returns the FinishReason field if non-nil, zero value otherwise.

### GetFinishReasonOk

`func (o *AiOpenAIChunkChoice) GetFinishReasonOk() (*AiOpenAIFinishReason, bool)`

GetFinishReasonOk returns a tuple with the FinishReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishReason

`func (o *AiOpenAIChunkChoice) SetFinishReason(v AiOpenAIFinishReason)`

SetFinishReason sets FinishReason field to given value.


### SetFinishReasonNil

`func (o *AiOpenAIChunkChoice) SetFinishReasonNil(b bool)`

 SetFinishReasonNil sets the value for FinishReason to be an explicit nil

### UnsetFinishReason
`func (o *AiOpenAIChunkChoice) UnsetFinishReason()`

UnsetFinishReason ensures that no value is present for FinishReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


