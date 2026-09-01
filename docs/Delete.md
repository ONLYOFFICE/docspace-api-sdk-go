# Delete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteAfter** | Pointer to **bool** | Specifies whether to delete a file after the editing session is finished or not. | [optional] 
**Immediately** | Pointer to **bool** | Specifies whether to move a file to the \\Trash\\ folder or delete it immediately. | [optional] 

## Methods

### NewDelete

`func NewDelete() *Delete`

NewDelete instantiates a new Delete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteWithDefaults

`func NewDeleteWithDefaults() *Delete`

NewDeleteWithDefaults instantiates a new Delete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteAfter

`func (o *Delete) GetDeleteAfter() bool`

GetDeleteAfter returns the DeleteAfter field if non-nil, zero value otherwise.

### GetDeleteAfterOk

`func (o *Delete) GetDeleteAfterOk() (*bool, bool)`

GetDeleteAfterOk returns a tuple with the DeleteAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAfter

`func (o *Delete) SetDeleteAfter(v bool)`

SetDeleteAfter sets DeleteAfter field to given value.

### HasDeleteAfter

`func (o *Delete) HasDeleteAfter() bool`

HasDeleteAfter returns a boolean if a field has been set.

### GetImmediately

`func (o *Delete) GetImmediately() bool`

GetImmediately returns the Immediately field if non-nil, zero value otherwise.

### GetImmediatelyOk

`func (o *Delete) GetImmediatelyOk() (*bool, bool)`

GetImmediatelyOk returns a tuple with the Immediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediately

`func (o *Delete) SetImmediately(v bool)`

SetImmediately sets Immediately field to given value.

### HasImmediately

`func (o *Delete) HasImmediately() bool`

HasImmediately returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


