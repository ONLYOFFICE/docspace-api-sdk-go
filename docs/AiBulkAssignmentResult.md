# AiBulkAssignmentResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | True when every entry was persisted. | 
**Errors** | Pointer to [**[]AiBulkAssignmentResultErrorsInner**](AiBulkAssignmentResultErrorsInner.md) | What was rejected, per action. Present on failure - and then no entry was persisted. | [optional] 

## Methods

### NewAiBulkAssignmentResult

`func NewAiBulkAssignmentResult(success bool, ) *AiBulkAssignmentResult`

NewAiBulkAssignmentResult instantiates a new AiBulkAssignmentResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiBulkAssignmentResultWithDefaults

`func NewAiBulkAssignmentResultWithDefaults() *AiBulkAssignmentResult`

NewAiBulkAssignmentResultWithDefaults instantiates a new AiBulkAssignmentResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AiBulkAssignmentResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AiBulkAssignmentResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AiBulkAssignmentResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetErrors

`func (o *AiBulkAssignmentResult) GetErrors() []AiBulkAssignmentResultErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AiBulkAssignmentResult) GetErrorsOk() (*[]AiBulkAssignmentResultErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AiBulkAssignmentResult) SetErrors(v []AiBulkAssignmentResultErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AiBulkAssignmentResult) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


