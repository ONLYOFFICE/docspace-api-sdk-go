# ArchiveRoomRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteAfter** | Pointer to **bool** | Specifies whether to archive a room after the editing session is finished or not. | [optional] 

## Methods

### NewArchiveRoomRequest

`func NewArchiveRoomRequest() *ArchiveRoomRequest`

NewArchiveRoomRequest instantiates a new ArchiveRoomRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveRoomRequestWithDefaults

`func NewArchiveRoomRequestWithDefaults() *ArchiveRoomRequest`

NewArchiveRoomRequestWithDefaults instantiates a new ArchiveRoomRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteAfter

`func (o *ArchiveRoomRequest) GetDeleteAfter() bool`

GetDeleteAfter returns the DeleteAfter field if non-nil, zero value otherwise.

### GetDeleteAfterOk

`func (o *ArchiveRoomRequest) GetDeleteAfterOk() (*bool, bool)`

GetDeleteAfterOk returns a tuple with the DeleteAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAfter

`func (o *ArchiveRoomRequest) SetDeleteAfter(v bool)`

SetDeleteAfter sets DeleteAfter field to given value.

### HasDeleteAfter

`func (o *ArchiveRoomRequest) HasDeleteAfter() bool`

HasDeleteAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


