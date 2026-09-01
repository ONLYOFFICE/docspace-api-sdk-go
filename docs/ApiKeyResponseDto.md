# ApiKeyResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The API key unique identifier. | 
**Name** | **NullableString** | The API key name. | 
**Key** | **NullableString** | The full API key value (only returned when creating a new key). | 
**KeyPostfix** | Pointer to **NullableString** | The API key postfix (used for identification). | [optional] 
**Permissions** | **[]string** | The list of permissions granted to the API key. | 
**LastUsed** | Pointer to **NullableTime** | The date and time when the API key was last used. | [optional] 
**CreateOn** | Pointer to **NullableTime** | The date and time when the API key was created. | [optional] 
**CreateBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) | The identifier of the user who created the API key. | [optional] 
**ExpiresAt** | Pointer to **NullableTime** | The date and time when the API key expires. | [optional] 
**IsActive** | **bool** | Indicates whether the API key is active or not. | 

## Methods

### NewApiKeyResponseDto

`func NewApiKeyResponseDto(id string, name NullableString, key NullableString, permissions []string, isActive bool, ) *ApiKeyResponseDto`

NewApiKeyResponseDto instantiates a new ApiKeyResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyResponseDtoWithDefaults

`func NewApiKeyResponseDtoWithDefaults() *ApiKeyResponseDto`

NewApiKeyResponseDtoWithDefaults instantiates a new ApiKeyResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiKeyResponseDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiKeyResponseDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiKeyResponseDto) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ApiKeyResponseDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKeyResponseDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKeyResponseDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ApiKeyResponseDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApiKeyResponseDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetKey

`func (o *ApiKeyResponseDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApiKeyResponseDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApiKeyResponseDto) SetKey(v string)`

SetKey sets Key field to given value.


### SetKeyNil

`func (o *ApiKeyResponseDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *ApiKeyResponseDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetKeyPostfix

`func (o *ApiKeyResponseDto) GetKeyPostfix() string`

GetKeyPostfix returns the KeyPostfix field if non-nil, zero value otherwise.

### GetKeyPostfixOk

`func (o *ApiKeyResponseDto) GetKeyPostfixOk() (*string, bool)`

GetKeyPostfixOk returns a tuple with the KeyPostfix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPostfix

`func (o *ApiKeyResponseDto) SetKeyPostfix(v string)`

SetKeyPostfix sets KeyPostfix field to given value.

### HasKeyPostfix

`func (o *ApiKeyResponseDto) HasKeyPostfix() bool`

HasKeyPostfix returns a boolean if a field has been set.

### SetKeyPostfixNil

`func (o *ApiKeyResponseDto) SetKeyPostfixNil(b bool)`

 SetKeyPostfixNil sets the value for KeyPostfix to be an explicit nil

### UnsetKeyPostfix
`func (o *ApiKeyResponseDto) UnsetKeyPostfix()`

UnsetKeyPostfix ensures that no value is present for KeyPostfix, not even an explicit nil
### GetPermissions

`func (o *ApiKeyResponseDto) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ApiKeyResponseDto) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ApiKeyResponseDto) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### SetPermissionsNil

`func (o *ApiKeyResponseDto) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *ApiKeyResponseDto) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetLastUsed

`func (o *ApiKeyResponseDto) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *ApiKeyResponseDto) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *ApiKeyResponseDto) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *ApiKeyResponseDto) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### SetLastUsedNil

`func (o *ApiKeyResponseDto) SetLastUsedNil(b bool)`

 SetLastUsedNil sets the value for LastUsed to be an explicit nil

### UnsetLastUsed
`func (o *ApiKeyResponseDto) UnsetLastUsed()`

UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
### GetCreateOn

`func (o *ApiKeyResponseDto) GetCreateOn() time.Time`

GetCreateOn returns the CreateOn field if non-nil, zero value otherwise.

### GetCreateOnOk

`func (o *ApiKeyResponseDto) GetCreateOnOk() (*time.Time, bool)`

GetCreateOnOk returns a tuple with the CreateOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateOn

`func (o *ApiKeyResponseDto) SetCreateOn(v time.Time)`

SetCreateOn sets CreateOn field to given value.

### HasCreateOn

`func (o *ApiKeyResponseDto) HasCreateOn() bool`

HasCreateOn returns a boolean if a field has been set.

### SetCreateOnNil

`func (o *ApiKeyResponseDto) SetCreateOnNil(b bool)`

 SetCreateOnNil sets the value for CreateOn to be an explicit nil

### UnsetCreateOn
`func (o *ApiKeyResponseDto) UnsetCreateOn()`

UnsetCreateOn ensures that no value is present for CreateOn, not even an explicit nil
### GetCreateBy

`func (o *ApiKeyResponseDto) GetCreateBy() EmployeeDto`

GetCreateBy returns the CreateBy field if non-nil, zero value otherwise.

### GetCreateByOk

`func (o *ApiKeyResponseDto) GetCreateByOk() (*EmployeeDto, bool)`

GetCreateByOk returns a tuple with the CreateBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBy

`func (o *ApiKeyResponseDto) SetCreateBy(v EmployeeDto)`

SetCreateBy sets CreateBy field to given value.

### HasCreateBy

`func (o *ApiKeyResponseDto) HasCreateBy() bool`

HasCreateBy returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ApiKeyResponseDto) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ApiKeyResponseDto) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ApiKeyResponseDto) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ApiKeyResponseDto) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *ApiKeyResponseDto) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *ApiKeyResponseDto) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetIsActive

`func (o *ApiKeyResponseDto) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ApiKeyResponseDto) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ApiKeyResponseDto) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


