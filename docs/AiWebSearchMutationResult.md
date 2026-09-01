# AiWebSearchMutationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | True when the configuration was persisted. | 
**Config** | Pointer to [**AiWebSearchConfig**](AiWebSearchConfig.md) | The persisted web-search configuration. Present on success. | [optional] 
**Error** | Pointer to [**AiTErrorData**](AiTErrorData.md) | Why the configuration was rejected. Present on failure. | [optional] 

## Methods

### NewAiWebSearchMutationResult

`func NewAiWebSearchMutationResult(success bool, ) *AiWebSearchMutationResult`

NewAiWebSearchMutationResult instantiates a new AiWebSearchMutationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiWebSearchMutationResultWithDefaults

`func NewAiWebSearchMutationResultWithDefaults() *AiWebSearchMutationResult`

NewAiWebSearchMutationResultWithDefaults instantiates a new AiWebSearchMutationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AiWebSearchMutationResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AiWebSearchMutationResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AiWebSearchMutationResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetConfig

`func (o *AiWebSearchMutationResult) GetConfig() AiWebSearchConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AiWebSearchMutationResult) GetConfigOk() (*AiWebSearchConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AiWebSearchMutationResult) SetConfig(v AiWebSearchConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AiWebSearchMutationResult) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetError

`func (o *AiWebSearchMutationResult) GetError() AiTErrorData`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AiWebSearchMutationResult) GetErrorOk() (*AiTErrorData, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AiWebSearchMutationResult) SetError(v AiTErrorData)`

SetError sets Error field to given value.

### HasError

`func (o *AiWebSearchMutationResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


