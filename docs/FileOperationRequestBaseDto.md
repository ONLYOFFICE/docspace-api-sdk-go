# FileOperationRequestBaseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnSingleOperation** | Pointer to **bool** | Which operations the answer carries: `true` returns the operation this call started and nothing else, `false`  returns every operation of the same kind that the caller has running or unread. When nothing was queued, which  happens for an empty selection, `true` falls back to the full list. | [optional] 

## Methods

### NewFileOperationRequestBaseDto

`func NewFileOperationRequestBaseDto() *FileOperationRequestBaseDto`

NewFileOperationRequestBaseDto instantiates a new FileOperationRequestBaseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileOperationRequestBaseDtoWithDefaults

`func NewFileOperationRequestBaseDtoWithDefaults() *FileOperationRequestBaseDto`

NewFileOperationRequestBaseDtoWithDefaults instantiates a new FileOperationRequestBaseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnSingleOperation

`func (o *FileOperationRequestBaseDto) GetReturnSingleOperation() bool`

GetReturnSingleOperation returns the ReturnSingleOperation field if non-nil, zero value otherwise.

### GetReturnSingleOperationOk

`func (o *FileOperationRequestBaseDto) GetReturnSingleOperationOk() (*bool, bool)`

GetReturnSingleOperationOk returns a tuple with the ReturnSingleOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnSingleOperation

`func (o *FileOperationRequestBaseDto) SetReturnSingleOperation(v bool)`

SetReturnSingleOperation sets ReturnSingleOperation field to given value.

### HasReturnSingleOperation

`func (o *FileOperationRequestBaseDto) HasReturnSingleOperation() bool`

HasReturnSingleOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


