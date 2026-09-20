# ChunkedUploadSessionResponseWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Always true in a body that reaches the caller, because a call that does not succeed answers with an error  status and no body at all. It cannot be used to tell a refusal from a success. | [optional] 
**Data** | Pointer to [**ChunkedUploadSessionResponse**](ChunkedUploadSessionResponse.md) | The reserved upload itself, in the same shape the newer session operations answer with directly. | [optional] 

## Methods

### NewChunkedUploadSessionResponseWrapper

`func NewChunkedUploadSessionResponseWrapper() *ChunkedUploadSessionResponseWrapper`

NewChunkedUploadSessionResponseWrapper instantiates a new ChunkedUploadSessionResponseWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChunkedUploadSessionResponseWrapperWithDefaults

`func NewChunkedUploadSessionResponseWrapperWithDefaults() *ChunkedUploadSessionResponseWrapper`

NewChunkedUploadSessionResponseWrapperWithDefaults instantiates a new ChunkedUploadSessionResponseWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ChunkedUploadSessionResponseWrapper) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ChunkedUploadSessionResponseWrapper) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ChunkedUploadSessionResponseWrapper) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ChunkedUploadSessionResponseWrapper) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *ChunkedUploadSessionResponseWrapper) GetData() ChunkedUploadSessionResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ChunkedUploadSessionResponseWrapper) GetDataOk() (*ChunkedUploadSessionResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ChunkedUploadSessionResponseWrapper) SetData(v ChunkedUploadSessionResponse)`

SetData sets Data field to given value.

### HasData

`func (o *ChunkedUploadSessionResponseWrapper) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


