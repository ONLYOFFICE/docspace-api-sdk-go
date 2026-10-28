# AiOpenAIChatCompletionChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The completion identifier, stable across every chunk of one response. | 
**Object** | **string** | Always `chat.completion.chunk`. | 
**Created** | **float32** | When the completion started, in Unix seconds. | 
**Model** | **string** | The model that produced the completion - the resolved profile's model. | 
**Choices** | [**[]AiOpenAIChunkChoice**](AiOpenAIChunkChoice.md) | The choices carried by this chunk. This service emits exactly one. | 

## Methods

### NewAiOpenAIChatCompletionChunk

`func NewAiOpenAIChatCompletionChunk(id string, object string, created float32, model string, choices []AiOpenAIChunkChoice, ) *AiOpenAIChatCompletionChunk`

NewAiOpenAIChatCompletionChunk instantiates a new AiOpenAIChatCompletionChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiOpenAIChatCompletionChunkWithDefaults

`func NewAiOpenAIChatCompletionChunkWithDefaults() *AiOpenAIChatCompletionChunk`

NewAiOpenAIChatCompletionChunkWithDefaults instantiates a new AiOpenAIChatCompletionChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiOpenAIChatCompletionChunk) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiOpenAIChatCompletionChunk) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiOpenAIChatCompletionChunk) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *AiOpenAIChatCompletionChunk) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AiOpenAIChatCompletionChunk) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AiOpenAIChatCompletionChunk) SetObject(v string)`

SetObject sets Object field to given value.


### GetCreated

`func (o *AiOpenAIChatCompletionChunk) GetCreated() float32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AiOpenAIChatCompletionChunk) GetCreatedOk() (*float32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AiOpenAIChatCompletionChunk) SetCreated(v float32)`

SetCreated sets Created field to given value.


### GetModel

`func (o *AiOpenAIChatCompletionChunk) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiOpenAIChatCompletionChunk) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiOpenAIChatCompletionChunk) SetModel(v string)`

SetModel sets Model field to given value.


### GetChoices

`func (o *AiOpenAIChatCompletionChunk) GetChoices() []AiOpenAIChunkChoice`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *AiOpenAIChatCompletionChunk) GetChoicesOk() (*[]AiOpenAIChunkChoice, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *AiOpenAIChatCompletionChunk) SetChoices(v []AiOpenAIChunkChoice)`

SetChoices sets Choices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


