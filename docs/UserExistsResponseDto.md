# UserExistsResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exists** | **bool** | Specifies whether the user exists or not. | 
**Status** | Pointer to [**EmployeeStatus**](EmployeeStatus.md) | The user status, if the user exists. | [optional] 

## Methods

### NewUserExistsResponseDto

`func NewUserExistsResponseDto(exists bool, ) *UserExistsResponseDto`

NewUserExistsResponseDto instantiates a new UserExistsResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserExistsResponseDtoWithDefaults

`func NewUserExistsResponseDtoWithDefaults() *UserExistsResponseDto`

NewUserExistsResponseDtoWithDefaults instantiates a new UserExistsResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExists

`func (o *UserExistsResponseDto) GetExists() bool`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *UserExistsResponseDto) GetExistsOk() (*bool, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *UserExistsResponseDto) SetExists(v bool)`

SetExists sets Exists field to given value.


### GetStatus

`func (o *UserExistsResponseDto) GetStatus() EmployeeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserExistsResponseDto) GetStatusOk() (*EmployeeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserExistsResponseDto) SetStatus(v EmployeeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserExistsResponseDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


