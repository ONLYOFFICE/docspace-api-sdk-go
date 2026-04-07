# ToolDecisionRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Decision** | Pointer to [**ToolExecutionDecision**](ToolExecutionDecision.md) |  | [optional] 

## Methods

### NewToolDecisionRequestBody

`func NewToolDecisionRequestBody() *ToolDecisionRequestBody`

NewToolDecisionRequestBody instantiates a new ToolDecisionRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolDecisionRequestBodyWithDefaults

`func NewToolDecisionRequestBodyWithDefaults() *ToolDecisionRequestBody`

NewToolDecisionRequestBodyWithDefaults instantiates a new ToolDecisionRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecision

`func (o *ToolDecisionRequestBody) GetDecision() ToolExecutionDecision`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *ToolDecisionRequestBody) GetDecisionOk() (*ToolExecutionDecision, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *ToolDecisionRequestBody) SetDecision(v ToolExecutionDecision)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *ToolDecisionRequestBody) HasDecision() bool`

HasDecision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


