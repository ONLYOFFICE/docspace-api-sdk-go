# AiFolderMutationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | True when the folder was persisted. | 
**Folder** | Pointer to [**AiPromptFolder**](AiPromptFolder.md) | The persisted folder. Present on success. | [optional] 
**Error** | Pointer to [**AiTErrorData**](AiTErrorData.md) | Why the folder was rejected. Present on failure. | [optional] 

## Methods

### NewAiFolderMutationResult

`func NewAiFolderMutationResult(success bool, ) *AiFolderMutationResult`

NewAiFolderMutationResult instantiates a new AiFolderMutationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiFolderMutationResultWithDefaults

`func NewAiFolderMutationResultWithDefaults() *AiFolderMutationResult`

NewAiFolderMutationResultWithDefaults instantiates a new AiFolderMutationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AiFolderMutationResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AiFolderMutationResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AiFolderMutationResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetFolder

`func (o *AiFolderMutationResult) GetFolder() AiPromptFolder`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *AiFolderMutationResult) GetFolderOk() (*AiPromptFolder, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *AiFolderMutationResult) SetFolder(v AiPromptFolder)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *AiFolderMutationResult) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetError

`func (o *AiFolderMutationResult) GetError() AiTErrorData`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AiFolderMutationResult) GetErrorOk() (*AiTErrorData, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AiFolderMutationResult) SetError(v AiTErrorData)`

SetError sets Error field to given value.

### HasError

`func (o *AiFolderMutationResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


