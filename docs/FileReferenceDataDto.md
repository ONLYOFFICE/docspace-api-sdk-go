# FileReferenceDataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileKey** | Pointer to **NullableString** | The id of the document inside the portal named below. | [optional] 
**InstanceId** | Pointer to **NullableString** | The portal the document lives in. A reference whose value is not this portal cannot be resolved by the file  key and falls back to the path or the link. | [optional] 
**RoomId** | Pointer to **NullableString** | The room the document lies in. It is filled in only for a document opened in a virtual data room, and stays  empty everywhere else. | [optional] 
**CanEditRoom** | Pointer to **bool** | Whether the caller may manage the room named above; it is only meaningful together with it. | [optional] 

## Methods

### NewFileReferenceDataDto

`func NewFileReferenceDataDto() *FileReferenceDataDto`

NewFileReferenceDataDto instantiates a new FileReferenceDataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileReferenceDataDtoWithDefaults

`func NewFileReferenceDataDtoWithDefaults() *FileReferenceDataDto`

NewFileReferenceDataDtoWithDefaults instantiates a new FileReferenceDataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileKey

`func (o *FileReferenceDataDto) GetFileKey() string`

GetFileKey returns the FileKey field if non-nil, zero value otherwise.

### GetFileKeyOk

`func (o *FileReferenceDataDto) GetFileKeyOk() (*string, bool)`

GetFileKeyOk returns a tuple with the FileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileKey

`func (o *FileReferenceDataDto) SetFileKey(v string)`

SetFileKey sets FileKey field to given value.

### HasFileKey

`func (o *FileReferenceDataDto) HasFileKey() bool`

HasFileKey returns a boolean if a field has been set.

### SetFileKeyNil

`func (o *FileReferenceDataDto) SetFileKeyNil(b bool)`

 SetFileKeyNil sets the value for FileKey to be an explicit nil

### UnsetFileKey
`func (o *FileReferenceDataDto) UnsetFileKey()`

UnsetFileKey ensures that no value is present for FileKey, not even an explicit nil
### GetInstanceId

`func (o *FileReferenceDataDto) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FileReferenceDataDto) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FileReferenceDataDto) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *FileReferenceDataDto) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *FileReferenceDataDto) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *FileReferenceDataDto) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetRoomId

`func (o *FileReferenceDataDto) GetRoomId() string`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *FileReferenceDataDto) GetRoomIdOk() (*string, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *FileReferenceDataDto) SetRoomId(v string)`

SetRoomId sets RoomId field to given value.

### HasRoomId

`func (o *FileReferenceDataDto) HasRoomId() bool`

HasRoomId returns a boolean if a field has been set.

### SetRoomIdNil

`func (o *FileReferenceDataDto) SetRoomIdNil(b bool)`

 SetRoomIdNil sets the value for RoomId to be an explicit nil

### UnsetRoomId
`func (o *FileReferenceDataDto) UnsetRoomId()`

UnsetRoomId ensures that no value is present for RoomId, not even an explicit nil
### GetCanEditRoom

`func (o *FileReferenceDataDto) GetCanEditRoom() bool`

GetCanEditRoom returns the CanEditRoom field if non-nil, zero value otherwise.

### GetCanEditRoomOk

`func (o *FileReferenceDataDto) GetCanEditRoomOk() (*bool, bool)`

GetCanEditRoomOk returns a tuple with the CanEditRoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEditRoom

`func (o *FileReferenceDataDto) SetCanEditRoom(v bool)`

SetCanEditRoom sets CanEditRoom field to given value.

### HasCanEditRoom

`func (o *FileReferenceDataDto) HasCanEditRoom() bool`

HasCanEditRoom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


