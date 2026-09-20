# DeleteRoomRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteAfter** | Pointer to **bool** | Carried by the contract but not acted upon: the deletion behaves the same either way, and the record of the  finished job is kept until it is read once. | [optional] 

## Methods

### NewDeleteRoomRequest

`func NewDeleteRoomRequest() *DeleteRoomRequest`

NewDeleteRoomRequest instantiates a new DeleteRoomRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteRoomRequestWithDefaults

`func NewDeleteRoomRequestWithDefaults() *DeleteRoomRequest`

NewDeleteRoomRequestWithDefaults instantiates a new DeleteRoomRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteAfter

`func (o *DeleteRoomRequest) GetDeleteAfter() bool`

GetDeleteAfter returns the DeleteAfter field if non-nil, zero value otherwise.

### GetDeleteAfterOk

`func (o *DeleteRoomRequest) GetDeleteAfterOk() (*bool, bool)`

GetDeleteAfterOk returns a tuple with the DeleteAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAfter

`func (o *DeleteRoomRequest) SetDeleteAfter(v bool)`

SetDeleteAfter sets DeleteAfter field to given value.

### HasDeleteAfter

`func (o *DeleteRoomRequest) HasDeleteAfter() bool`

HasDeleteAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


