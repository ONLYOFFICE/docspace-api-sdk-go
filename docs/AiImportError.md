# AiImportError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | `folder` or `prompt`, plus the offending name or id. | 
**Ref** | **string** | The offending entry - its name or its id. | 
**Error** | [**AiTErrorData**](AiTErrorData.md) | Why the entry was rejected. | 

## Methods

### NewAiImportError

`func NewAiImportError(kind string, ref string, error_ AiTErrorData, ) *AiImportError`

NewAiImportError instantiates a new AiImportError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiImportErrorWithDefaults

`func NewAiImportErrorWithDefaults() *AiImportError`

NewAiImportErrorWithDefaults instantiates a new AiImportError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AiImportError) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AiImportError) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AiImportError) SetKind(v string)`

SetKind sets Kind field to given value.


### GetRef

`func (o *AiImportError) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *AiImportError) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *AiImportError) SetRef(v string)`

SetRef sets Ref field to given value.


### GetError

`func (o *AiImportError) GetError() AiTErrorData`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AiImportError) GetErrorOk() (*AiTErrorData, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AiImportError) SetError(v AiTErrorData)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


