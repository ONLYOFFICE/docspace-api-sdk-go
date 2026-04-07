# FillingFormResultDtoInteger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormNumber** | **int32** | The filling form number. | 
**CompletedForm** | Pointer to [**FileDtoInteger**](FileDtoInteger.md) |  | [optional] 
**OriginalForm** | Pointer to [**FileDtoInteger**](FileDtoInteger.md) |  | [optional] 
**Manager** | Pointer to [**EmployeeFullDto**](EmployeeFullDto.md) |  | [optional] 
**RoomId** | **int32** | The room ID where filling the form. | 
**IsRoomMember** | Pointer to **bool** | Specifies if the manager who fills the form is a room member or not. | [optional] 

## Methods

### NewFillingFormResultDtoInteger

`func NewFillingFormResultDtoInteger(formNumber int32, roomId int32, ) *FillingFormResultDtoInteger`

NewFillingFormResultDtoInteger instantiates a new FillingFormResultDtoInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFillingFormResultDtoIntegerWithDefaults

`func NewFillingFormResultDtoIntegerWithDefaults() *FillingFormResultDtoInteger`

NewFillingFormResultDtoIntegerWithDefaults instantiates a new FillingFormResultDtoInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormNumber

`func (o *FillingFormResultDtoInteger) GetFormNumber() int32`

GetFormNumber returns the FormNumber field if non-nil, zero value otherwise.

### GetFormNumberOk

`func (o *FillingFormResultDtoInteger) GetFormNumberOk() (*int32, bool)`

GetFormNumberOk returns a tuple with the FormNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormNumber

`func (o *FillingFormResultDtoInteger) SetFormNumber(v int32)`

SetFormNumber sets FormNumber field to given value.


### GetCompletedForm

`func (o *FillingFormResultDtoInteger) GetCompletedForm() FileDtoInteger`

GetCompletedForm returns the CompletedForm field if non-nil, zero value otherwise.

### GetCompletedFormOk

`func (o *FillingFormResultDtoInteger) GetCompletedFormOk() (*FileDtoInteger, bool)`

GetCompletedFormOk returns a tuple with the CompletedForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedForm

`func (o *FillingFormResultDtoInteger) SetCompletedForm(v FileDtoInteger)`

SetCompletedForm sets CompletedForm field to given value.

### HasCompletedForm

`func (o *FillingFormResultDtoInteger) HasCompletedForm() bool`

HasCompletedForm returns a boolean if a field has been set.

### GetOriginalForm

`func (o *FillingFormResultDtoInteger) GetOriginalForm() FileDtoInteger`

GetOriginalForm returns the OriginalForm field if non-nil, zero value otherwise.

### GetOriginalFormOk

`func (o *FillingFormResultDtoInteger) GetOriginalFormOk() (*FileDtoInteger, bool)`

GetOriginalFormOk returns a tuple with the OriginalForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalForm

`func (o *FillingFormResultDtoInteger) SetOriginalForm(v FileDtoInteger)`

SetOriginalForm sets OriginalForm field to given value.

### HasOriginalForm

`func (o *FillingFormResultDtoInteger) HasOriginalForm() bool`

HasOriginalForm returns a boolean if a field has been set.

### GetManager

`func (o *FillingFormResultDtoInteger) GetManager() EmployeeFullDto`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *FillingFormResultDtoInteger) GetManagerOk() (*EmployeeFullDto, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *FillingFormResultDtoInteger) SetManager(v EmployeeFullDto)`

SetManager sets Manager field to given value.

### HasManager

`func (o *FillingFormResultDtoInteger) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetRoomId

`func (o *FillingFormResultDtoInteger) GetRoomId() int32`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *FillingFormResultDtoInteger) GetRoomIdOk() (*int32, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *FillingFormResultDtoInteger) SetRoomId(v int32)`

SetRoomId sets RoomId field to given value.


### GetIsRoomMember

`func (o *FillingFormResultDtoInteger) GetIsRoomMember() bool`

GetIsRoomMember returns the IsRoomMember field if non-nil, zero value otherwise.

### GetIsRoomMemberOk

`func (o *FillingFormResultDtoInteger) GetIsRoomMemberOk() (*bool, bool)`

GetIsRoomMemberOk returns a tuple with the IsRoomMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoomMember

`func (o *FillingFormResultDtoInteger) SetIsRoomMember(v bool)`

SetIsRoomMember sets IsRoomMember field to given value.

### HasIsRoomMember

`func (o *FillingFormResultDtoInteger) HasIsRoomMember() bool`

HasIsRoomMember returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


