# AiSuggestedQuestionsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | `ready` - the questions are there; `pending` - they are being generated and arrive over the socket, ask again later; `unavailable` - the attachment is not a form that can have any. | 
**Questions** | [**[]AiSuggestedQuestion**](AiSuggestedQuestion.md) | The questions; empty unless `status` is `ready`. | 

## Methods

### NewAiSuggestedQuestionsResult

`func NewAiSuggestedQuestionsResult(status string, questions []AiSuggestedQuestion, ) *AiSuggestedQuestionsResult`

NewAiSuggestedQuestionsResult instantiates a new AiSuggestedQuestionsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSuggestedQuestionsResultWithDefaults

`func NewAiSuggestedQuestionsResultWithDefaults() *AiSuggestedQuestionsResult`

NewAiSuggestedQuestionsResultWithDefaults instantiates a new AiSuggestedQuestionsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AiSuggestedQuestionsResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AiSuggestedQuestionsResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AiSuggestedQuestionsResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetQuestions

`func (o *AiSuggestedQuestionsResult) GetQuestions() []AiSuggestedQuestion`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *AiSuggestedQuestionsResult) GetQuestionsOk() (*[]AiSuggestedQuestion, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *AiSuggestedQuestionsResult) SetQuestions(v []AiSuggestedQuestion)`

SetQuestions sets Questions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


