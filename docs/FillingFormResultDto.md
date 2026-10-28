# FillingFormResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormNumber** | **int32** | The number this copy was given among the copies made of the same form, counting up from 1. It is the number  the results of the form are ordered by and the one the title of the copy carries. | 
**CompletedForm** | Pointer to [**FileDto**](FileDto.md) | The filled copy that the session produced, as an ordinary file: it can be read and downloaded with the file  operations of this API. | [optional] 
**OriginalForm** | Pointer to [**FileDto**](FileDto.md) | The form the copy was made from, so that a client can offer filling it once more. | [optional] 
**Manager** | Pointer to [**EmployeeFullDto**](EmployeeFullDto.md) | The account that owns the original form, reported with its email address, so that the person who has just  filled the form knows who receives it and whom to ask about it. | [optional] 
**RoomId** | **int32** | The room the form was filled in. It comes back as 0 when the session was reached through a link shared for  that single form rather than for its room, in which case there is no room the caller could be sent to. | 
**IsRoomMember** | Pointer to **bool** | Tells whether the calling account may open that room: true for a member of the room and for a portal  administrator, in which case a client can offer going to the room; false for the anonymous caller who filled  the form through a link and can only be shown the copy itself. | [optional] 

## Methods

### NewFillingFormResultDto

`func NewFillingFormResultDto(formNumber int32, roomId int32, ) *FillingFormResultDto`

NewFillingFormResultDto instantiates a new FillingFormResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFillingFormResultDtoWithDefaults

`func NewFillingFormResultDtoWithDefaults() *FillingFormResultDto`

NewFillingFormResultDtoWithDefaults instantiates a new FillingFormResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormNumber

`func (o *FillingFormResultDto) GetFormNumber() int32`

GetFormNumber returns the FormNumber field if non-nil, zero value otherwise.

### GetFormNumberOk

`func (o *FillingFormResultDto) GetFormNumberOk() (*int32, bool)`

GetFormNumberOk returns a tuple with the FormNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormNumber

`func (o *FillingFormResultDto) SetFormNumber(v int32)`

SetFormNumber sets FormNumber field to given value.


### GetCompletedForm

`func (o *FillingFormResultDto) GetCompletedForm() FileDto`

GetCompletedForm returns the CompletedForm field if non-nil, zero value otherwise.

### GetCompletedFormOk

`func (o *FillingFormResultDto) GetCompletedFormOk() (*FileDto, bool)`

GetCompletedFormOk returns a tuple with the CompletedForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedForm

`func (o *FillingFormResultDto) SetCompletedForm(v FileDto)`

SetCompletedForm sets CompletedForm field to given value.

### HasCompletedForm

`func (o *FillingFormResultDto) HasCompletedForm() bool`

HasCompletedForm returns a boolean if a field has been set.

### GetOriginalForm

`func (o *FillingFormResultDto) GetOriginalForm() FileDto`

GetOriginalForm returns the OriginalForm field if non-nil, zero value otherwise.

### GetOriginalFormOk

`func (o *FillingFormResultDto) GetOriginalFormOk() (*FileDto, bool)`

GetOriginalFormOk returns a tuple with the OriginalForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalForm

`func (o *FillingFormResultDto) SetOriginalForm(v FileDto)`

SetOriginalForm sets OriginalForm field to given value.

### HasOriginalForm

`func (o *FillingFormResultDto) HasOriginalForm() bool`

HasOriginalForm returns a boolean if a field has been set.

### GetManager

`func (o *FillingFormResultDto) GetManager() EmployeeFullDto`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *FillingFormResultDto) GetManagerOk() (*EmployeeFullDto, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *FillingFormResultDto) SetManager(v EmployeeFullDto)`

SetManager sets Manager field to given value.

### HasManager

`func (o *FillingFormResultDto) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetRoomId

`func (o *FillingFormResultDto) GetRoomId() int32`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *FillingFormResultDto) GetRoomIdOk() (*int32, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *FillingFormResultDto) SetRoomId(v int32)`

SetRoomId sets RoomId field to given value.


### GetIsRoomMember

`func (o *FillingFormResultDto) GetIsRoomMember() bool`

GetIsRoomMember returns the IsRoomMember field if non-nil, zero value otherwise.

### GetIsRoomMemberOk

`func (o *FillingFormResultDto) GetIsRoomMemberOk() (*bool, bool)`

GetIsRoomMemberOk returns a tuple with the IsRoomMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoomMember

`func (o *FillingFormResultDto) SetIsRoomMember(v bool)`

SetIsRoomMember sets IsRoomMember field to given value.

### HasIsRoomMember

`func (o *FillingFormResultDto) HasIsRoomMember() bool`

HasIsRoomMember returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


