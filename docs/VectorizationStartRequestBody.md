# VectorizationStartRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | **[]int32** | The set of file identifiers to submit for vectorization. | 

## Methods

### NewVectorizationStartRequestBody

`func NewVectorizationStartRequestBody(files []int32, ) *VectorizationStartRequestBody`

NewVectorizationStartRequestBody instantiates a new VectorizationStartRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorizationStartRequestBodyWithDefaults

`func NewVectorizationStartRequestBodyWithDefaults() *VectorizationStartRequestBody`

NewVectorizationStartRequestBodyWithDefaults instantiates a new VectorizationStartRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *VectorizationStartRequestBody) GetFiles() []int32`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *VectorizationStartRequestBody) GetFilesOk() (*[]int32, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *VectorizationStartRequestBody) SetFiles(v []int32)`

SetFiles sets Files field to given value.


### SetFilesNil

`func (o *VectorizationStartRequestBody) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *VectorizationStartRequestBody) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


