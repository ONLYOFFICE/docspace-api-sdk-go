# ExternalDbSyncTaskDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The task ID. | 
**Error** | Pointer to **NullableString** | The error message if the synchronization failed. | [optional] 
**Percentage** | **int32** | The progress percentage of the synchronization. | 
**IsCompleted** | **bool** | Specifies whether the synchronization is completed or not. | 
**Status** | [**DistributedTaskStatus**](DistributedTaskStatus.md) | The status of the synchronization task. | 
**Forms** | [**[]ExternalDbSyncFormResultDto**](ExternalDbSyncFormResultDto.md) | The synchronization results for all original forms in the room. | 

## Methods

### NewExternalDbSyncTaskDto

`func NewExternalDbSyncTaskDto(id NullableString, percentage int32, isCompleted bool, status DistributedTaskStatus, forms []ExternalDbSyncFormResultDto, ) *ExternalDbSyncTaskDto`

NewExternalDbSyncTaskDto instantiates a new ExternalDbSyncTaskDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalDbSyncTaskDtoWithDefaults

`func NewExternalDbSyncTaskDtoWithDefaults() *ExternalDbSyncTaskDto`

NewExternalDbSyncTaskDtoWithDefaults instantiates a new ExternalDbSyncTaskDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalDbSyncTaskDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalDbSyncTaskDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalDbSyncTaskDto) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ExternalDbSyncTaskDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExternalDbSyncTaskDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetError

`func (o *ExternalDbSyncTaskDto) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ExternalDbSyncTaskDto) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ExternalDbSyncTaskDto) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ExternalDbSyncTaskDto) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ExternalDbSyncTaskDto) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ExternalDbSyncTaskDto) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetPercentage

`func (o *ExternalDbSyncTaskDto) GetPercentage() int32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *ExternalDbSyncTaskDto) GetPercentageOk() (*int32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *ExternalDbSyncTaskDto) SetPercentage(v int32)`

SetPercentage sets Percentage field to given value.


### GetIsCompleted

`func (o *ExternalDbSyncTaskDto) GetIsCompleted() bool`

GetIsCompleted returns the IsCompleted field if non-nil, zero value otherwise.

### GetIsCompletedOk

`func (o *ExternalDbSyncTaskDto) GetIsCompletedOk() (*bool, bool)`

GetIsCompletedOk returns a tuple with the IsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompleted

`func (o *ExternalDbSyncTaskDto) SetIsCompleted(v bool)`

SetIsCompleted sets IsCompleted field to given value.


### GetStatus

`func (o *ExternalDbSyncTaskDto) GetStatus() DistributedTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalDbSyncTaskDto) GetStatusOk() (*DistributedTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalDbSyncTaskDto) SetStatus(v DistributedTaskStatus)`

SetStatus sets Status field to given value.


### GetForms

`func (o *ExternalDbSyncTaskDto) GetForms() []ExternalDbSyncFormResultDto`

GetForms returns the Forms field if non-nil, zero value otherwise.

### GetFormsOk

`func (o *ExternalDbSyncTaskDto) GetFormsOk() (*[]ExternalDbSyncFormResultDto, bool)`

GetFormsOk returns a tuple with the Forms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForms

`func (o *ExternalDbSyncTaskDto) SetForms(v []ExternalDbSyncFormResultDto)`

SetForms sets Forms field to given value.


### SetFormsNil

`func (o *ExternalDbSyncTaskDto) SetFormsNil(b bool)`

 SetFormsNil sets the value for Forms to be an explicit nil

### UnsetForms
`func (o *ExternalDbSyncTaskDto) UnsetForms()`

UnsetForms ensures that no value is present for Forms, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


