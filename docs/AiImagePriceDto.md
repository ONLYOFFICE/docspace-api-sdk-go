# AiImagePriceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prompt** | Pointer to **float64** | The cost of one million tokens sent to the image model, which is the prompt describing the picture. | [optional] 
**Completion** | Pointer to **float64** | The cost of one million tokens the image model writes back alongside the picture. | [optional] 
**Image** | Pointer to **float64** | The cost of one produced image, charged on top of the token amounts above. | [optional] 

## Methods

### NewAiImagePriceDto

`func NewAiImagePriceDto() *AiImagePriceDto`

NewAiImagePriceDto instantiates a new AiImagePriceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiImagePriceDtoWithDefaults

`func NewAiImagePriceDtoWithDefaults() *AiImagePriceDto`

NewAiImagePriceDtoWithDefaults instantiates a new AiImagePriceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrompt

`func (o *AiImagePriceDto) GetPrompt() float64`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AiImagePriceDto) GetPromptOk() (*float64, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AiImagePriceDto) SetPrompt(v float64)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *AiImagePriceDto) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetCompletion

`func (o *AiImagePriceDto) GetCompletion() float64`

GetCompletion returns the Completion field if non-nil, zero value otherwise.

### GetCompletionOk

`func (o *AiImagePriceDto) GetCompletionOk() (*float64, bool)`

GetCompletionOk returns a tuple with the Completion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletion

`func (o *AiImagePriceDto) SetCompletion(v float64)`

SetCompletion sets Completion field to given value.

### HasCompletion

`func (o *AiImagePriceDto) HasCompletion() bool`

HasCompletion returns a boolean if a field has been set.

### GetImage

`func (o *AiImagePriceDto) GetImage() float64`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AiImagePriceDto) GetImageOk() (*float64, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AiImagePriceDto) SetImage(v float64)`

SetImage sets Image field to given value.

### HasImage

`func (o *AiImagePriceDto) HasImage() bool`

HasImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


