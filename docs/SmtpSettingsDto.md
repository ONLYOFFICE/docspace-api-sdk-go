# SmtpSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** | The SMTP host. | [optional] 
**Port** | Pointer to **NullableInt32** | The SMTP port. | [optional] 
**SenderAddress** | Pointer to **NullableString** | The sender address. | [optional] 
**SenderDisplayName** | Pointer to **NullableString** | The sender display name. | [optional] 
**CredentialsUserName** | Pointer to **NullableString** | The credentials username. | [optional] 
**CredentialsUserPassword** | Pointer to **NullableString** | The credentials user password. | [optional] 
**EnableSSL** | Pointer to **bool** | Specifies whether the SSL is enabled or not. | [optional] 
**EnableAuth** | Pointer to **bool** | Specifies whether the authentication is enabled or not. | [optional] 
**UseNtlm** | Pointer to **bool** | Specifies whether to use NTLM or not. | [optional] 
**IsDefaultSettings** | Pointer to **bool** | Specifies if the current settings are default or not. | [optional] 

## Methods

### NewSmtpSettingsDto

`func NewSmtpSettingsDto() *SmtpSettingsDto`

NewSmtpSettingsDto instantiates a new SmtpSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpSettingsDtoWithDefaults

`func NewSmtpSettingsDtoWithDefaults() *SmtpSettingsDto`

NewSmtpSettingsDtoWithDefaults instantiates a new SmtpSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SmtpSettingsDto) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SmtpSettingsDto) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SmtpSettingsDto) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SmtpSettingsDto) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *SmtpSettingsDto) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *SmtpSettingsDto) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *SmtpSettingsDto) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SmtpSettingsDto) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SmtpSettingsDto) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SmtpSettingsDto) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *SmtpSettingsDto) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *SmtpSettingsDto) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetSenderAddress

`func (o *SmtpSettingsDto) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *SmtpSettingsDto) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *SmtpSettingsDto) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.

### HasSenderAddress

`func (o *SmtpSettingsDto) HasSenderAddress() bool`

HasSenderAddress returns a boolean if a field has been set.

### SetSenderAddressNil

`func (o *SmtpSettingsDto) SetSenderAddressNil(b bool)`

 SetSenderAddressNil sets the value for SenderAddress to be an explicit nil

### UnsetSenderAddress
`func (o *SmtpSettingsDto) UnsetSenderAddress()`

UnsetSenderAddress ensures that no value is present for SenderAddress, not even an explicit nil
### GetSenderDisplayName

`func (o *SmtpSettingsDto) GetSenderDisplayName() string`

GetSenderDisplayName returns the SenderDisplayName field if non-nil, zero value otherwise.

### GetSenderDisplayNameOk

`func (o *SmtpSettingsDto) GetSenderDisplayNameOk() (*string, bool)`

GetSenderDisplayNameOk returns a tuple with the SenderDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderDisplayName

`func (o *SmtpSettingsDto) SetSenderDisplayName(v string)`

SetSenderDisplayName sets SenderDisplayName field to given value.

### HasSenderDisplayName

`func (o *SmtpSettingsDto) HasSenderDisplayName() bool`

HasSenderDisplayName returns a boolean if a field has been set.

### SetSenderDisplayNameNil

`func (o *SmtpSettingsDto) SetSenderDisplayNameNil(b bool)`

 SetSenderDisplayNameNil sets the value for SenderDisplayName to be an explicit nil

### UnsetSenderDisplayName
`func (o *SmtpSettingsDto) UnsetSenderDisplayName()`

UnsetSenderDisplayName ensures that no value is present for SenderDisplayName, not even an explicit nil
### GetCredentialsUserName

`func (o *SmtpSettingsDto) GetCredentialsUserName() string`

GetCredentialsUserName returns the CredentialsUserName field if non-nil, zero value otherwise.

### GetCredentialsUserNameOk

`func (o *SmtpSettingsDto) GetCredentialsUserNameOk() (*string, bool)`

GetCredentialsUserNameOk returns a tuple with the CredentialsUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsUserName

`func (o *SmtpSettingsDto) SetCredentialsUserName(v string)`

SetCredentialsUserName sets CredentialsUserName field to given value.

### HasCredentialsUserName

`func (o *SmtpSettingsDto) HasCredentialsUserName() bool`

HasCredentialsUserName returns a boolean if a field has been set.

### SetCredentialsUserNameNil

`func (o *SmtpSettingsDto) SetCredentialsUserNameNil(b bool)`

 SetCredentialsUserNameNil sets the value for CredentialsUserName to be an explicit nil

### UnsetCredentialsUserName
`func (o *SmtpSettingsDto) UnsetCredentialsUserName()`

UnsetCredentialsUserName ensures that no value is present for CredentialsUserName, not even an explicit nil
### GetCredentialsUserPassword

`func (o *SmtpSettingsDto) GetCredentialsUserPassword() string`

GetCredentialsUserPassword returns the CredentialsUserPassword field if non-nil, zero value otherwise.

### GetCredentialsUserPasswordOk

`func (o *SmtpSettingsDto) GetCredentialsUserPasswordOk() (*string, bool)`

GetCredentialsUserPasswordOk returns a tuple with the CredentialsUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsUserPassword

`func (o *SmtpSettingsDto) SetCredentialsUserPassword(v string)`

SetCredentialsUserPassword sets CredentialsUserPassword field to given value.

### HasCredentialsUserPassword

`func (o *SmtpSettingsDto) HasCredentialsUserPassword() bool`

HasCredentialsUserPassword returns a boolean if a field has been set.

### SetCredentialsUserPasswordNil

`func (o *SmtpSettingsDto) SetCredentialsUserPasswordNil(b bool)`

 SetCredentialsUserPasswordNil sets the value for CredentialsUserPassword to be an explicit nil

### UnsetCredentialsUserPassword
`func (o *SmtpSettingsDto) UnsetCredentialsUserPassword()`

UnsetCredentialsUserPassword ensures that no value is present for CredentialsUserPassword, not even an explicit nil
### GetEnableSSL

`func (o *SmtpSettingsDto) GetEnableSSL() bool`

GetEnableSSL returns the EnableSSL field if non-nil, zero value otherwise.

### GetEnableSSLOk

`func (o *SmtpSettingsDto) GetEnableSSLOk() (*bool, bool)`

GetEnableSSLOk returns a tuple with the EnableSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSSL

`func (o *SmtpSettingsDto) SetEnableSSL(v bool)`

SetEnableSSL sets EnableSSL field to given value.

### HasEnableSSL

`func (o *SmtpSettingsDto) HasEnableSSL() bool`

HasEnableSSL returns a boolean if a field has been set.

### GetEnableAuth

`func (o *SmtpSettingsDto) GetEnableAuth() bool`

GetEnableAuth returns the EnableAuth field if non-nil, zero value otherwise.

### GetEnableAuthOk

`func (o *SmtpSettingsDto) GetEnableAuthOk() (*bool, bool)`

GetEnableAuthOk returns a tuple with the EnableAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAuth

`func (o *SmtpSettingsDto) SetEnableAuth(v bool)`

SetEnableAuth sets EnableAuth field to given value.

### HasEnableAuth

`func (o *SmtpSettingsDto) HasEnableAuth() bool`

HasEnableAuth returns a boolean if a field has been set.

### GetUseNtlm

`func (o *SmtpSettingsDto) GetUseNtlm() bool`

GetUseNtlm returns the UseNtlm field if non-nil, zero value otherwise.

### GetUseNtlmOk

`func (o *SmtpSettingsDto) GetUseNtlmOk() (*bool, bool)`

GetUseNtlmOk returns a tuple with the UseNtlm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNtlm

`func (o *SmtpSettingsDto) SetUseNtlm(v bool)`

SetUseNtlm sets UseNtlm field to given value.

### HasUseNtlm

`func (o *SmtpSettingsDto) HasUseNtlm() bool`

HasUseNtlm returns a boolean if a field has been set.

### GetIsDefaultSettings

`func (o *SmtpSettingsDto) GetIsDefaultSettings() bool`

GetIsDefaultSettings returns the IsDefaultSettings field if non-nil, zero value otherwise.

### GetIsDefaultSettingsOk

`func (o *SmtpSettingsDto) GetIsDefaultSettingsOk() (*bool, bool)`

GetIsDefaultSettingsOk returns a tuple with the IsDefaultSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultSettings

`func (o *SmtpSettingsDto) SetIsDefaultSettings(v bool)`

SetIsDefaultSettings sets IsDefaultSettings field to given value.

### HasIsDefaultSettings

`func (o *SmtpSettingsDto) HasIsDefaultSettings() bool`

HasIsDefaultSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


