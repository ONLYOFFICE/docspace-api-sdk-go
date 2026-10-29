# AiMetaFillRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The entry's table. Only `file` for now: a folder or a room answers 400. | 
**Id** | [**AiMetaGenerateRequestId**](AiMetaGenerateRequestId.md) |  | 
**Prompt** | Pointer to **string** | What to pay attention to; a hint for the model. | [optional] 
**Overwrite** | Pointer to **bool** | Also replace the fields that already carry a value. Default false. | [optional] 
**DryRun** | Pointer to **bool** | Resolve the values but write nothing; the caller writes what the user confirms. Default true: a write needs an explicit false. | [optional] 

## Methods

### NewAiMetaFillRequest

`func NewAiMetaFillRequest(kind string, id AiMetaGenerateRequestId, ) *AiMetaFillRequest`

NewAiMetaFillRequest instantiates a new AiMetaFillRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiMetaFillRequestWithDefaults

`func NewAiMetaFillRequestWithDefaults() *AiMetaFillRequest`

NewAiMetaFillRequestWithDefaults instantiates a new AiMetaFillRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AiMetaFillRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AiMetaFillRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AiMetaFillRequest) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *AiMetaFillRequest) GetId() AiMetaGenerateRequestId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiMetaFillRequest) GetIdOk() (*AiMetaGenerateRequestId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiMetaFillRequest) SetId(v AiMetaGenerateRequestId)`

SetId sets Id field to given value.


### GetPrompt

`func (o *AiMetaFillRequest) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AiMetaFillRequest) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AiMetaFillRequest) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *AiMetaFillRequest) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetOverwrite

`func (o *AiMetaFillRequest) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *AiMetaFillRequest) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *AiMetaFillRequest) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *AiMetaFillRequest) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetDryRun

`func (o *AiMetaFillRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *AiMetaFillRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *AiMetaFillRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *AiMetaFillRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


