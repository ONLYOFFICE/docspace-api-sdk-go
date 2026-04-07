# EncryptionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **NullableString** | The encryption password. | [optional] 
**Status** | Pointer to [**EncryprtionStatus**](EncryprtionStatus.md) |  | [optional] 
**NotifyUsers** | Pointer to **bool** | Specifies if the users will be notified about the encryption operation or not. | [optional] 

## Methods

### NewEncryptionSettings

`func NewEncryptionSettings() *EncryptionSettings`

NewEncryptionSettings instantiates a new EncryptionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionSettingsWithDefaults

`func NewEncryptionSettingsWithDefaults() *EncryptionSettings`

NewEncryptionSettingsWithDefaults instantiates a new EncryptionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *EncryptionSettings) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EncryptionSettings) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EncryptionSettings) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EncryptionSettings) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *EncryptionSettings) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *EncryptionSettings) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetStatus

`func (o *EncryptionSettings) GetStatus() EncryprtionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EncryptionSettings) GetStatusOk() (*EncryprtionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EncryptionSettings) SetStatus(v EncryprtionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EncryptionSettings) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNotifyUsers

`func (o *EncryptionSettings) GetNotifyUsers() bool`

GetNotifyUsers returns the NotifyUsers field if non-nil, zero value otherwise.

### GetNotifyUsersOk

`func (o *EncryptionSettings) GetNotifyUsersOk() (*bool, bool)`

GetNotifyUsersOk returns a tuple with the NotifyUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUsers

`func (o *EncryptionSettings) SetNotifyUsers(v bool)`

SetNotifyUsers sets NotifyUsers field to given value.

### HasNotifyUsers

`func (o *EncryptionSettings) HasNotifyUsers() bool`

HasNotifyUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


