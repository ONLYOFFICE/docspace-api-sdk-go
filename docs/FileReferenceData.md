# FileReferenceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileKey** | Pointer to **NullableString** | The unique document identifier used by the service to get a link to the file. | [optional] 
**InstanceId** | Pointer to **NullableString** | The unique system identifier. | [optional] 
**RoomId** | Pointer to **NullableString** | Room ID | [optional] 
**CanEditRoom** | Pointer to **bool** | Specifies if the room can be edited out or not. | [optional] 

## Methods

### NewFileReferenceData

`func NewFileReferenceData() *FileReferenceData`

NewFileReferenceData instantiates a new FileReferenceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileReferenceDataWithDefaults

`func NewFileReferenceDataWithDefaults() *FileReferenceData`

NewFileReferenceDataWithDefaults instantiates a new FileReferenceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileKey

`func (o *FileReferenceData) GetFileKey() string`

GetFileKey returns the FileKey field if non-nil, zero value otherwise.

### GetFileKeyOk

`func (o *FileReferenceData) GetFileKeyOk() (*string, bool)`

GetFileKeyOk returns a tuple with the FileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileKey

`func (o *FileReferenceData) SetFileKey(v string)`

SetFileKey sets FileKey field to given value.

### HasFileKey

`func (o *FileReferenceData) HasFileKey() bool`

HasFileKey returns a boolean if a field has been set.

### SetFileKeyNil

`func (o *FileReferenceData) SetFileKeyNil(b bool)`

 SetFileKeyNil sets the value for FileKey to be an explicit nil

### UnsetFileKey
`func (o *FileReferenceData) UnsetFileKey()`

UnsetFileKey ensures that no value is present for FileKey, not even an explicit nil
### GetInstanceId

`func (o *FileReferenceData) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FileReferenceData) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FileReferenceData) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *FileReferenceData) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *FileReferenceData) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *FileReferenceData) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetRoomId

`func (o *FileReferenceData) GetRoomId() string`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *FileReferenceData) GetRoomIdOk() (*string, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *FileReferenceData) SetRoomId(v string)`

SetRoomId sets RoomId field to given value.

### HasRoomId

`func (o *FileReferenceData) HasRoomId() bool`

HasRoomId returns a boolean if a field has been set.

### SetRoomIdNil

`func (o *FileReferenceData) SetRoomIdNil(b bool)`

 SetRoomIdNil sets the value for RoomId to be an explicit nil

### UnsetRoomId
`func (o *FileReferenceData) UnsetRoomId()`

UnsetRoomId ensures that no value is present for RoomId, not even an explicit nil
### GetCanEditRoom

`func (o *FileReferenceData) GetCanEditRoom() bool`

GetCanEditRoom returns the CanEditRoom field if non-nil, zero value otherwise.

### GetCanEditRoomOk

`func (o *FileReferenceData) GetCanEditRoomOk() (*bool, bool)`

GetCanEditRoomOk returns a tuple with the CanEditRoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditRoom

`func (o *FileReferenceData) SetCanEditRoom(v bool)`

SetCanEditRoom sets CanEditRoom field to given value.

### HasCanEditRoom

`func (o *FileReferenceData) HasCanEditRoom() bool`

HasCanEditRoom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


