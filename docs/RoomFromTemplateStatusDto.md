# RoomFromTemplateStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoomId** | **int32** | The room ID. | 
**Progress** | **float64** | The progress of creating a room from the template. | 
**Error** | **NullableString** | The error message that is sent when a room is not created successfully from the template. | 
**IsCompleted** | **bool** | Specifies whether the process of creating a room from the template is completed. | 

## Methods

### NewRoomFromTemplateStatusDto

`func NewRoomFromTemplateStatusDto(roomId int32, progress float64, error_ NullableString, isCompleted bool, ) *RoomFromTemplateStatusDto`

NewRoomFromTemplateStatusDto instantiates a new RoomFromTemplateStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomFromTemplateStatusDtoWithDefaults

`func NewRoomFromTemplateStatusDtoWithDefaults() *RoomFromTemplateStatusDto`

NewRoomFromTemplateStatusDtoWithDefaults instantiates a new RoomFromTemplateStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoomId

`func (o *RoomFromTemplateStatusDto) GetRoomId() int32`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *RoomFromTemplateStatusDto) GetRoomIdOk() (*int32, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *RoomFromTemplateStatusDto) SetRoomId(v int32)`

SetRoomId sets RoomId field to given value.


### GetProgress

`func (o *RoomFromTemplateStatusDto) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *RoomFromTemplateStatusDto) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *RoomFromTemplateStatusDto) SetProgress(v float64)`

SetProgress sets Progress field to given value.


### GetError

`func (o *RoomFromTemplateStatusDto) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RoomFromTemplateStatusDto) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RoomFromTemplateStatusDto) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *RoomFromTemplateStatusDto) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *RoomFromTemplateStatusDto) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetIsCompleted

`func (o *RoomFromTemplateStatusDto) GetIsCompleted() bool`

GetIsCompleted returns the IsCompleted field if non-nil, zero value otherwise.

### GetIsCompletedOk

`func (o *RoomFromTemplateStatusDto) GetIsCompletedOk() (*bool, bool)`

GetIsCompletedOk returns a tuple with the IsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompleted

`func (o *RoomFromTemplateStatusDto) SetIsCompleted(v bool)`

SetIsCompleted sets IsCompleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


