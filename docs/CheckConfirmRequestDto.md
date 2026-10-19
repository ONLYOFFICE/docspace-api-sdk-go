# CheckConfirmRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **NullableString** | The email validation key. | [optional] 
**EmplType** | Pointer to [**EmployeeType**](EmployeeType.md) | The user type. | [optional] 
**Email** | Pointer to **NullableString** | The email address. | [optional] 
**EncEmail** | Pointer to **NullableString** | The encrypted email address. | [optional] 
**UiD** | Pointer to **NullableString** | The user ID. | [optional] 
**Type** | Pointer to [**ConfirmType**](ConfirmType.md) | The confirmation email type. | [optional] 
**First** | Pointer to **NullableString** | Specifies whether it is the first time account access or not. | [optional] 
**RoomId** | Pointer to **NullableString** | The room ID. | [optional] 

## Methods

### NewCheckConfirmRequestDto

`func NewCheckConfirmRequestDto() *CheckConfirmRequestDto`

NewCheckConfirmRequestDto instantiates a new CheckConfirmRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckConfirmRequestDtoWithDefaults

`func NewCheckConfirmRequestDtoWithDefaults() *CheckConfirmRequestDto`

NewCheckConfirmRequestDtoWithDefaults instantiates a new CheckConfirmRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CheckConfirmRequestDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CheckConfirmRequestDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CheckConfirmRequestDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CheckConfirmRequestDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *CheckConfirmRequestDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *CheckConfirmRequestDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetEmplType

`func (o *CheckConfirmRequestDto) GetEmplType() EmployeeType`

GetEmplType returns the EmplType field if non-nil, zero value otherwise.

### GetEmplTypeOk

`func (o *CheckConfirmRequestDto) GetEmplTypeOk() (*EmployeeType, bool)`

GetEmplTypeOk returns a tuple with the EmplType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmplType

`func (o *CheckConfirmRequestDto) SetEmplType(v EmployeeType)`

SetEmplType sets EmplType field to given value.

### HasEmplType

`func (o *CheckConfirmRequestDto) HasEmplType() bool`

HasEmplType returns a boolean if a field has been set.

### GetEmail

`func (o *CheckConfirmRequestDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CheckConfirmRequestDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CheckConfirmRequestDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CheckConfirmRequestDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CheckConfirmRequestDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CheckConfirmRequestDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEncEmail

`func (o *CheckConfirmRequestDto) GetEncEmail() string`

GetEncEmail returns the EncEmail field if non-nil, zero value otherwise.

### GetEncEmailOk

`func (o *CheckConfirmRequestDto) GetEncEmailOk() (*string, bool)`

GetEncEmailOk returns a tuple with the EncEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncEmail

`func (o *CheckConfirmRequestDto) SetEncEmail(v string)`

SetEncEmail sets EncEmail field to given value.

### HasEncEmail

`func (o *CheckConfirmRequestDto) HasEncEmail() bool`

HasEncEmail returns a boolean if a field has been set.

### SetEncEmailNil

`func (o *CheckConfirmRequestDto) SetEncEmailNil(b bool)`

 SetEncEmailNil sets the value for EncEmail to be an explicit nil

### UnsetEncEmail
`func (o *CheckConfirmRequestDto) UnsetEncEmail()`

UnsetEncEmail ensures that no value is present for EncEmail, not even an explicit nil
### GetUiD

`func (o *CheckConfirmRequestDto) GetUiD() string`

GetUiD returns the UiD field if non-nil, zero value otherwise.

### GetUiDOk

`func (o *CheckConfirmRequestDto) GetUiDOk() (*string, bool)`

GetUiDOk returns a tuple with the UiD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiD

`func (o *CheckConfirmRequestDto) SetUiD(v string)`

SetUiD sets UiD field to given value.

### HasUiD

`func (o *CheckConfirmRequestDto) HasUiD() bool`

HasUiD returns a boolean if a field has been set.

### SetUiDNil

`func (o *CheckConfirmRequestDto) SetUiDNil(b bool)`

 SetUiDNil sets the value for UiD to be an explicit nil

### UnsetUiD
`func (o *CheckConfirmRequestDto) UnsetUiD()`

UnsetUiD ensures that no value is present for UiD, not even an explicit nil
### GetType

`func (o *CheckConfirmRequestDto) GetType() ConfirmType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckConfirmRequestDto) GetTypeOk() (*ConfirmType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckConfirmRequestDto) SetType(v ConfirmType)`

SetType sets Type field to given value.

### HasType

`func (o *CheckConfirmRequestDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFirst

`func (o *CheckConfirmRequestDto) GetFirst() string`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *CheckConfirmRequestDto) GetFirstOk() (*string, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *CheckConfirmRequestDto) SetFirst(v string)`

SetFirst sets First field to given value.

### HasFirst

`func (o *CheckConfirmRequestDto) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### SetFirstNil

`func (o *CheckConfirmRequestDto) SetFirstNil(b bool)`

 SetFirstNil sets the value for First to be an explicit nil

### UnsetFirst
`func (o *CheckConfirmRequestDto) UnsetFirst()`

UnsetFirst ensures that no value is present for First, not even an explicit nil
### GetRoomId

`func (o *CheckConfirmRequestDto) GetRoomId() string`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *CheckConfirmRequestDto) GetRoomIdOk() (*string, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *CheckConfirmRequestDto) SetRoomId(v string)`

SetRoomId sets RoomId field to given value.

### HasRoomId

`func (o *CheckConfirmRequestDto) HasRoomId() bool`

HasRoomId returns a boolean if a field has been set.

### SetRoomIdNil

`func (o *CheckConfirmRequestDto) SetRoomIdNil(b bool)`

 SetRoomIdNil sets the value for RoomId to be an explicit nil

### UnsetRoomId
`func (o *CheckConfirmRequestDto) UnsetRoomId()`

UnsetRoomId ensures that no value is present for RoomId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


