# StartUpdateUserTypeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**EmployeeType**](EmployeeType.md) | The new user type. | [optional] 
**UserId** | Pointer to **string** | The user ID. | [optional] 
**ReassignUserId** | Pointer to **NullableString** | The user ID to reassign. | [optional] 

## Methods

### NewStartUpdateUserTypeDto

`func NewStartUpdateUserTypeDto() *StartUpdateUserTypeDto`

NewStartUpdateUserTypeDto instantiates a new StartUpdateUserTypeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartUpdateUserTypeDtoWithDefaults

`func NewStartUpdateUserTypeDtoWithDefaults() *StartUpdateUserTypeDto`

NewStartUpdateUserTypeDtoWithDefaults instantiates a new StartUpdateUserTypeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StartUpdateUserTypeDto) GetType() EmployeeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StartUpdateUserTypeDto) GetTypeOk() (*EmployeeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StartUpdateUserTypeDto) SetType(v EmployeeType)`

SetType sets Type field to given value.

### HasType

`func (o *StartUpdateUserTypeDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *StartUpdateUserTypeDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *StartUpdateUserTypeDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *StartUpdateUserTypeDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *StartUpdateUserTypeDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetReassignUserId

`func (o *StartUpdateUserTypeDto) GetReassignUserId() string`

GetReassignUserId returns the ReassignUserId field if non-nil, zero value otherwise.

### GetReassignUserIdOk

`func (o *StartUpdateUserTypeDto) GetReassignUserIdOk() (*string, bool)`

GetReassignUserIdOk returns a tuple with the ReassignUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReassignUserId

`func (o *StartUpdateUserTypeDto) SetReassignUserId(v string)`

SetReassignUserId sets ReassignUserId field to given value.

### HasReassignUserId

`func (o *StartUpdateUserTypeDto) HasReassignUserId() bool`

HasReassignUserId returns a boolean if a field has been set.

### SetReassignUserIdNil

`func (o *StartUpdateUserTypeDto) SetReassignUserIdNil(b bool)`

 SetReassignUserIdNil sets the value for ReassignUserId to be an explicit nil

### UnsetReassignUserId
`func (o *StartUpdateUserTypeDto) UnsetReassignUserId()`

UnsetReassignUserId ensures that no value is present for ReassignUserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


