# DeleteFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteAfter** | Pointer to **bool** | When to delete: `true` waits until the editing session on the file has ended, `false` deletes at once, pulling  the file away from whoever is working on it. | [optional] 
**Immediately** | Pointer to **bool** | Where the file goes: `false` moves it to Trash, from where it can be restored, `true` deletes it for good.  Inside a room, where there is no Trash, deletion is always final. | [optional] 

## Methods

### NewDeleteFileRequest

`func NewDeleteFileRequest() *DeleteFileRequest`

NewDeleteFileRequest instantiates a new DeleteFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteFileRequestWithDefaults

`func NewDeleteFileRequestWithDefaults() *DeleteFileRequest`

NewDeleteFileRequestWithDefaults instantiates a new DeleteFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteAfter

`func (o *DeleteFileRequest) GetDeleteAfter() bool`

GetDeleteAfter returns the DeleteAfter field if non-nil, zero value otherwise.

### GetDeleteAfterOk

`func (o *DeleteFileRequest) GetDeleteAfterOk() (*bool, bool)`

GetDeleteAfterOk returns a tuple with the DeleteAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAfter

`func (o *DeleteFileRequest) SetDeleteAfter(v bool)`

SetDeleteAfter sets DeleteAfter field to given value.

### HasDeleteAfter

`func (o *DeleteFileRequest) HasDeleteAfter() bool`

HasDeleteAfter returns a boolean if a field has been set.

### GetImmediately

`func (o *DeleteFileRequest) GetImmediately() bool`

GetImmediately returns the Immediately field if non-nil, zero value otherwise.

### GetImmediatelyOk

`func (o *DeleteFileRequest) GetImmediatelyOk() (*bool, bool)`

GetImmediatelyOk returns a tuple with the Immediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediately

`func (o *DeleteFileRequest) SetImmediately(v bool)`

SetImmediately sets Immediately field to given value.

### HasImmediately

`func (o *DeleteFileRequest) HasImmediately() bool`

HasImmediately returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


