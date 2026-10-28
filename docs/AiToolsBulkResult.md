# AiToolsBulkResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | True when every custom MCP server was persisted. | 
**Errors** | Pointer to [**[]AiToolsBulkResultErrorsInner**](AiToolsBulkResultErrorsInner.md) | What was rejected, per server. Present on failure - and then no server was persisted. | [optional] 

## Methods

### NewAiToolsBulkResult

`func NewAiToolsBulkResult(success bool, ) *AiToolsBulkResult`

NewAiToolsBulkResult instantiates a new AiToolsBulkResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiToolsBulkResultWithDefaults

`func NewAiToolsBulkResultWithDefaults() *AiToolsBulkResult`

NewAiToolsBulkResultWithDefaults instantiates a new AiToolsBulkResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AiToolsBulkResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AiToolsBulkResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AiToolsBulkResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetErrors

`func (o *AiToolsBulkResult) GetErrors() []AiToolsBulkResultErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AiToolsBulkResult) GetErrorsOk() (*[]AiToolsBulkResultErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AiToolsBulkResult) SetErrors(v []AiToolsBulkResultErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AiToolsBulkResult) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


