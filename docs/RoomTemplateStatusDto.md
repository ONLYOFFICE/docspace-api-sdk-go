# RoomTemplateStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | **int32** | The template the job is building. It is meaningful once the job has created the template folder, and the  template can be opened with the room operations only after `isCompleted` turns true. | 
**Progress** | **float64** | How far the job has got. The value climbs while the contents of the room are being copied and reaches its  maximum at the very end, so it is an indication of life rather than a reliable estimate of the time left. | 
**Error** | Pointer to **NullableString** | Why the job stopped. It is empty while the job runs and after a successful one; when it is filled the  half-built template has already been removed, so nothing has to be cleaned up by the caller. | [optional] 
**IsCompleted** | **bool** | Whether the job has ended. It is set both after a successful build and after a failure, so `error` is what  tells the two apart, and the record keeps answering with the same values until another job is started. | 

## Methods

### NewRoomTemplateStatusDto

`func NewRoomTemplateStatusDto(templateId int32, progress float64, isCompleted bool, ) *RoomTemplateStatusDto`

NewRoomTemplateStatusDto instantiates a new RoomTemplateStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomTemplateStatusDtoWithDefaults

`func NewRoomTemplateStatusDtoWithDefaults() *RoomTemplateStatusDto`

NewRoomTemplateStatusDtoWithDefaults instantiates a new RoomTemplateStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *RoomTemplateStatusDto) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *RoomTemplateStatusDto) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *RoomTemplateStatusDto) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.


### GetProgress

`func (o *RoomTemplateStatusDto) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *RoomTemplateStatusDto) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *RoomTemplateStatusDto) SetProgress(v float64)`

SetProgress sets Progress field to given value.


### GetError

`func (o *RoomTemplateStatusDto) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RoomTemplateStatusDto) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RoomTemplateStatusDto) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *RoomTemplateStatusDto) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *RoomTemplateStatusDto) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *RoomTemplateStatusDto) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetIsCompleted

`func (o *RoomTemplateStatusDto) GetIsCompleted() bool`

GetIsCompleted returns the IsCompleted field if non-nil, zero value otherwise.

### GetIsCompletedOk

`func (o *RoomTemplateStatusDto) GetIsCompletedOk() (*bool, bool)`

GetIsCompletedOk returns a tuple with the IsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompleted

`func (o *RoomTemplateStatusDto) SetIsCompleted(v bool)`

SetIsCompleted sets IsCompleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


