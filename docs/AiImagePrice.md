# AiImagePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prompt** | Pointer to **float64** | The price of a single prompt token. | [optional] 
**Completion** | Pointer to **float64** | The cost associated with the completion of a prompt in an AI model. | [optional] 
**Image** | Pointer to **float64** | The price of a single generated image. | [optional] 

## Methods

### NewAiImagePrice

`func NewAiImagePrice() *AiImagePrice`

NewAiImagePrice instantiates a new AiImagePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiImagePriceWithDefaults

`func NewAiImagePriceWithDefaults() *AiImagePrice`

NewAiImagePriceWithDefaults instantiates a new AiImagePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrompt

`func (o *AiImagePrice) GetPrompt() float64`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AiImagePrice) GetPromptOk() (*float64, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AiImagePrice) SetPrompt(v float64)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *AiImagePrice) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetCompletion

`func (o *AiImagePrice) GetCompletion() float64`

GetCompletion returns the Completion field if non-nil, zero value otherwise.

### GetCompletionOk

`func (o *AiImagePrice) GetCompletionOk() (*float64, bool)`

GetCompletionOk returns a tuple with the Completion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletion

`func (o *AiImagePrice) SetCompletion(v float64)`

SetCompletion sets Completion field to given value.

### HasCompletion

`func (o *AiImagePrice) HasCompletion() bool`

HasCompletion returns a boolean if a field has been set.

### GetImage

`func (o *AiImagePrice) GetImage() float64`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AiImagePrice) GetImageOk() (*float64, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AiImagePrice) SetImage(v float64)`

SetImage sets Image field to given value.

### HasImage

`func (o *AiImagePrice) HasImage() bool`

HasImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


