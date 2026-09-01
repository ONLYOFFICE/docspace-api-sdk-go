# AiAssignmentMutationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | True when the assignment was persisted. | 
**Error** | Pointer to [**AiTErrorData**](AiTErrorData.md) | Why the assignment was rejected. Present on failure. | [optional] 

## Methods

### NewAiAssignmentMutationResult

`func NewAiAssignmentMutationResult(success bool, ) *AiAssignmentMutationResult`

NewAiAssignmentMutationResult instantiates a new AiAssignmentMutationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAssignmentMutationResultWithDefaults

`func NewAiAssignmentMutationResultWithDefaults() *AiAssignmentMutationResult`

NewAiAssignmentMutationResultWithDefaults instantiates a new AiAssignmentMutationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AiAssignmentMutationResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AiAssignmentMutationResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AiAssignmentMutationResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *AiAssignmentMutationResult) GetError() AiTErrorData`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AiAssignmentMutationResult) GetErrorOk() (*AiTErrorData, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AiAssignmentMutationResult) SetError(v AiTErrorData)`

SetError sets Error field to given value.

### HasError

`func (o *AiAssignmentMutationResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


