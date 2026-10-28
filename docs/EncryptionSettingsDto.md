# EncryptionSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **NullableString** | Always an empty string: the encryption password is never returned. | [optional] 
**Status** | Pointer to [**EncryptionStatus**](EncryptionStatus.md) | Whether the storage is encrypted, decrypted, or on its way to either. | [optional] 
**NotifyUsers** | Pointer to **bool** | Whether the users are notified when the operation starts and ends. | [optional] 

## Methods

### NewEncryptionSettingsDto

`func NewEncryptionSettingsDto() *EncryptionSettingsDto`

NewEncryptionSettingsDto instantiates a new EncryptionSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionSettingsDtoWithDefaults

`func NewEncryptionSettingsDtoWithDefaults() *EncryptionSettingsDto`

NewEncryptionSettingsDtoWithDefaults instantiates a new EncryptionSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *EncryptionSettingsDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EncryptionSettingsDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EncryptionSettingsDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EncryptionSettingsDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *EncryptionSettingsDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *EncryptionSettingsDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetStatus

`func (o *EncryptionSettingsDto) GetStatus() EncryptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EncryptionSettingsDto) GetStatusOk() (*EncryptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EncryptionSettingsDto) SetStatus(v EncryptionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EncryptionSettingsDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNotifyUsers

`func (o *EncryptionSettingsDto) GetNotifyUsers() bool`

GetNotifyUsers returns the NotifyUsers field if non-nil, zero value otherwise.

### GetNotifyUsersOk

`func (o *EncryptionSettingsDto) GetNotifyUsersOk() (*bool, bool)`

GetNotifyUsersOk returns a tuple with the NotifyUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUsers

`func (o *EncryptionSettingsDto) SetNotifyUsers(v bool)`

SetNotifyUsers sets NotifyUsers field to given value.

### HasNotifyUsers

`func (o *EncryptionSettingsDto) HasNotifyUsers() bool`

HasNotifyUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


