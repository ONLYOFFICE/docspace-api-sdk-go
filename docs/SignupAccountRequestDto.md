# SignupAccountRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployeeType** | Pointer to [**EmployeeType**](EmployeeType.md) |  | [optional] 
**Key** | **NullableString** | The user link key. | 
**Culture** | Pointer to **NullableString** | The user culture code. | [optional] 
**SerializedProfile** | **NullableString** | The third-party profile in the serialized format | 

## Methods

### NewSignupAccountRequestDto

`func NewSignupAccountRequestDto(key NullableString, serializedProfile NullableString, ) *SignupAccountRequestDto`

NewSignupAccountRequestDto instantiates a new SignupAccountRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignupAccountRequestDtoWithDefaults

`func NewSignupAccountRequestDtoWithDefaults() *SignupAccountRequestDto`

NewSignupAccountRequestDtoWithDefaults instantiates a new SignupAccountRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployeeType

`func (o *SignupAccountRequestDto) GetEmployeeType() EmployeeType`

GetEmployeeType returns the EmployeeType field if non-nil, zero value otherwise.

### GetEmployeeTypeOk

`func (o *SignupAccountRequestDto) GetEmployeeTypeOk() (*EmployeeType, bool)`

GetEmployeeTypeOk returns a tuple with the EmployeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeType

`func (o *SignupAccountRequestDto) SetEmployeeType(v EmployeeType)`

SetEmployeeType sets EmployeeType field to given value.

### HasEmployeeType

`func (o *SignupAccountRequestDto) HasEmployeeType() bool`

HasEmployeeType returns a boolean if a field has been set.

### GetKey

`func (o *SignupAccountRequestDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SignupAccountRequestDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SignupAccountRequestDto) SetKey(v string)`

SetKey sets Key field to given value.


### SetKeyNil

`func (o *SignupAccountRequestDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *SignupAccountRequestDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetCulture

`func (o *SignupAccountRequestDto) GetCulture() string`

GetCulture returns the Culture field if non-nil, zero value otherwise.

### GetCultureOk

`func (o *SignupAccountRequestDto) GetCultureOk() (*string, bool)`

GetCultureOk returns a tuple with the Culture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCulture

`func (o *SignupAccountRequestDto) SetCulture(v string)`

SetCulture sets Culture field to given value.

### HasCulture

`func (o *SignupAccountRequestDto) HasCulture() bool`

HasCulture returns a boolean if a field has been set.

### SetCultureNil

`func (o *SignupAccountRequestDto) SetCultureNil(b bool)`

 SetCultureNil sets the value for Culture to be an explicit nil

### UnsetCulture
`func (o *SignupAccountRequestDto) UnsetCulture()`

UnsetCulture ensures that no value is present for Culture, not even an explicit nil
### GetSerializedProfile

`func (o *SignupAccountRequestDto) GetSerializedProfile() string`

GetSerializedProfile returns the SerializedProfile field if non-nil, zero value otherwise.

### GetSerializedProfileOk

`func (o *SignupAccountRequestDto) GetSerializedProfileOk() (*string, bool)`

GetSerializedProfileOk returns a tuple with the SerializedProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedProfile

`func (o *SignupAccountRequestDto) SetSerializedProfile(v string)`

SetSerializedProfile sets SerializedProfile field to given value.


### SetSerializedProfileNil

`func (o *SignupAccountRequestDto) SetSerializedProfileNil(b bool)`

 SetSerializedProfileNil sets the value for SerializedProfile to be an explicit nil

### UnsetSerializedProfile
`func (o *SignupAccountRequestDto) UnsetSerializedProfile()`

UnsetSerializedProfile ensures that no value is present for SerializedProfile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


