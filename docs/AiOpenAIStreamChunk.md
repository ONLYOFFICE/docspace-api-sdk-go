# AiOpenAIStreamChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The completion identifier, stable across every chunk of one response. | 
**Object** | **string** | Always `chat.completion.chunk`. | 
**Created** | **float32** | When the completion started, in Unix seconds. | 
**Model** | **string** | The model that produced the completion - the resolved profile's model. | 
**Choices** | [**[]AiOpenAIChunkChoice**](AiOpenAIChunkChoice.md) | The choices carried by this chunk. This service emits exactly one. | 
**Error** | [**AiOpenAIStreamErrorError**](AiOpenAIStreamErrorError.md) |  | 

## Methods

### NewAiOpenAIStreamChunk

`func NewAiOpenAIStreamChunk(id string, object string, created float32, model string, choices []AiOpenAIChunkChoice, error_ AiOpenAIStreamErrorError, ) *AiOpenAIStreamChunk`

NewAiOpenAIStreamChunk instantiates a new AiOpenAIStreamChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiOpenAIStreamChunkWithDefaults

`func NewAiOpenAIStreamChunkWithDefaults() *AiOpenAIStreamChunk`

NewAiOpenAIStreamChunkWithDefaults instantiates a new AiOpenAIStreamChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiOpenAIStreamChunk) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiOpenAIStreamChunk) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiOpenAIStreamChunk) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *AiOpenAIStreamChunk) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AiOpenAIStreamChunk) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AiOpenAIStreamChunk) SetObject(v string)`

SetObject sets Object field to given value.


### GetCreated

`func (o *AiOpenAIStreamChunk) GetCreated() float32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AiOpenAIStreamChunk) GetCreatedOk() (*float32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AiOpenAIStreamChunk) SetCreated(v float32)`

SetCreated sets Created field to given value.


### GetModel

`func (o *AiOpenAIStreamChunk) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiOpenAIStreamChunk) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiOpenAIStreamChunk) SetModel(v string)`

SetModel sets Model field to given value.


### GetChoices

`func (o *AiOpenAIStreamChunk) GetChoices() []AiOpenAIChunkChoice`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *AiOpenAIStreamChunk) GetChoicesOk() (*[]AiOpenAIChunkChoice, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *AiOpenAIStreamChunk) SetChoices(v []AiOpenAIChunkChoice)`

SetChoices sets Choices field to given value.


### GetError

`func (o *AiOpenAIStreamChunk) GetError() AiOpenAIStreamErrorError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AiOpenAIStreamChunk) GetErrorOk() (*AiOpenAIStreamErrorError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AiOpenAIStreamChunk) SetError(v AiOpenAIStreamErrorError)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


