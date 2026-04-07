# TaskProgressResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The task progress ID. | 
**Error** | Pointer to **NullableString** | The task progress error message. | [optional] 
**Percentage** | **int32** | The percentage of the task progress. | 
**IsCompleted** | **bool** | Specifies if the task peogress is completed or not. | 
**Status** | [**DistributedTaskStatus**](DistributedTaskStatus.md) |  | 

## Methods

### NewTaskProgressResponseDto

`func NewTaskProgressResponseDto(id NullableString, percentage int32, isCompleted bool, status DistributedTaskStatus, ) *TaskProgressResponseDto`

NewTaskProgressResponseDto instantiates a new TaskProgressResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskProgressResponseDtoWithDefaults

`func NewTaskProgressResponseDtoWithDefaults() *TaskProgressResponseDto`

NewTaskProgressResponseDtoWithDefaults instantiates a new TaskProgressResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskProgressResponseDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskProgressResponseDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskProgressResponseDto) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *TaskProgressResponseDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TaskProgressResponseDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetError

`func (o *TaskProgressResponseDto) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TaskProgressResponseDto) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TaskProgressResponseDto) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *TaskProgressResponseDto) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *TaskProgressResponseDto) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *TaskProgressResponseDto) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetPercentage

`func (o *TaskProgressResponseDto) GetPercentage() int32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *TaskProgressResponseDto) GetPercentageOk() (*int32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *TaskProgressResponseDto) SetPercentage(v int32)`

SetPercentage sets Percentage field to given value.


### GetIsCompleted

`func (o *TaskProgressResponseDto) GetIsCompleted() bool`

GetIsCompleted returns the IsCompleted field if non-nil, zero value otherwise.

### GetIsCompletedOk

`func (o *TaskProgressResponseDto) GetIsCompletedOk() (*bool, bool)`

GetIsCompletedOk returns a tuple with the IsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompleted

`func (o *TaskProgressResponseDto) SetIsCompleted(v bool)`

SetIsCompleted sets IsCompleted field to given value.


### GetStatus

`func (o *TaskProgressResponseDto) GetStatus() DistributedTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskProgressResponseDto) GetStatusOk() (*DistributedTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskProgressResponseDto) SetStatus(v DistributedTaskStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


