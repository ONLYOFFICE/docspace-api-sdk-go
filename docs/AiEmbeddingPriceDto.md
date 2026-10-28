# AiEmbeddingPriceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prompt** | Pointer to **float64** | The cost of one million tokens turned into vectors. Embedding produces no completion, so this single  figure is the whole price. | [optional] 

## Methods

### NewAiEmbeddingPriceDto

`func NewAiEmbeddingPriceDto() *AiEmbeddingPriceDto`

NewAiEmbeddingPriceDto instantiates a new AiEmbeddingPriceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiEmbeddingPriceDtoWithDefaults

`func NewAiEmbeddingPriceDtoWithDefaults() *AiEmbeddingPriceDto`

NewAiEmbeddingPriceDtoWithDefaults instantiates a new AiEmbeddingPriceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrompt

`func (o *AiEmbeddingPriceDto) GetPrompt() float64`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AiEmbeddingPriceDto) GetPromptOk() (*float64, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AiEmbeddingPriceDto) SetPrompt(v float64)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *AiEmbeddingPriceDto) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


