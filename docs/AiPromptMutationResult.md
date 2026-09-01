# AiPromptMutationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | True when the prompt was persisted. | 
**Prompt** | Pointer to [**AiPrompt**](AiPrompt.md) | The persisted prompt. Present on success. | [optional] 
**Error** | Pointer to [**AiTErrorData**](AiTErrorData.md) | Why the prompt was rejected. Present on failure. | [optional] 

## Methods

### NewAiPromptMutationResult

`func NewAiPromptMutationResult(success bool, ) *AiPromptMutationResult`

NewAiPromptMutationResult instantiates a new AiPromptMutationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiPromptMutationResultWithDefaults

`func NewAiPromptMutationResultWithDefaults() *AiPromptMutationResult`

NewAiPromptMutationResultWithDefaults instantiates a new AiPromptMutationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AiPromptMutationResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AiPromptMutationResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AiPromptMutationResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetPrompt

`func (o *AiPromptMutationResult) GetPrompt() AiPrompt`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AiPromptMutationResult) GetPromptOk() (*AiPrompt, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AiPromptMutationResult) SetPrompt(v AiPrompt)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *AiPromptMutationResult) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetError

`func (o *AiPromptMutationResult) GetError() AiTErrorData`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AiPromptMutationResult) GetErrorOk() (*AiTErrorData, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AiPromptMutationResult) SetError(v AiTErrorData)`

SetError sets Error field to given value.

### HasError

`func (o *AiPromptMutationResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


