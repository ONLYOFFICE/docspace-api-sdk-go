# AiSuggestedQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Question** | **string** | The question as it is shown to the user. | 
**Prompt** | **string** | The prompt sent to the chat when the user picks the question. | 

## Methods

### NewAiSuggestedQuestion

`func NewAiSuggestedQuestion(question string, prompt string, ) *AiSuggestedQuestion`

NewAiSuggestedQuestion instantiates a new AiSuggestedQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSuggestedQuestionWithDefaults

`func NewAiSuggestedQuestionWithDefaults() *AiSuggestedQuestion`

NewAiSuggestedQuestionWithDefaults instantiates a new AiSuggestedQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestion

`func (o *AiSuggestedQuestion) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *AiSuggestedQuestion) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *AiSuggestedQuestion) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetPrompt

`func (o *AiSuggestedQuestion) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AiSuggestedQuestion) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AiSuggestedQuestion) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


