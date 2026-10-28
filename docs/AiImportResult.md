# AiImportResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | True when the whole bundle was imported. | 
**Imported** | Pointer to [**AiImportResultImported**](AiImportResultImported.md) |  | [optional] 
**Errors** | Pointer to [**[]AiImportError**](AiImportError.md) | What was rejected, per entry. Present on failure - and then nothing was imported. | [optional] 

## Methods

### NewAiImportResult

`func NewAiImportResult(success bool, ) *AiImportResult`

NewAiImportResult instantiates a new AiImportResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiImportResultWithDefaults

`func NewAiImportResultWithDefaults() *AiImportResult`

NewAiImportResultWithDefaults instantiates a new AiImportResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AiImportResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AiImportResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AiImportResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetImported

`func (o *AiImportResult) GetImported() AiImportResultImported`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *AiImportResult) GetImportedOk() (*AiImportResultImported, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *AiImportResult) SetImported(v AiImportResultImported)`

SetImported sets Imported field to given value.

### HasImported

`func (o *AiImportResult) HasImported() bool`

HasImported returns a boolean if a field has been set.

### GetErrors

`func (o *AiImportResult) GetErrors() []AiImportError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AiImportResult) GetErrorsOk() (*[]AiImportError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AiImportResult) SetErrors(v []AiImportError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AiImportResult) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


