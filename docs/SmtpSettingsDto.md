# SmtpSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** | The host name or address of the mail server. On a cloud portal that has saved no relay of its own every  field of this object comes back empty, because the installation's own server is not disclosed - only  `isDefaultSettings` is set there. | [optional] 
**Port** | Pointer to **NullableInt32** | The port the mail server is reached on - conventionally 25 or 587 without encryption from the start, 465  with it. It is empty when no port was stored, in which case the portal falls back to its own default. | [optional] 
**SenderAddress** | Pointer to **NullableString** | The address the letters are sent from, which appears in the From header and is what a reply goes to. | [optional] 
**SenderDisplayName** | Pointer to **NullableString** | The name shown beside that address in a recipient's mailbox. | [optional] 
**CredentialsUserName** | Pointer to **NullableString** | The account the portal signs in to the mail server as, meaningful only while `enableAuth` is `true`. | [optional] 
**CredentialsUserPassword** | Pointer to **NullableString** | Always empty here: the stored password is never returned, so a client that sends these settings back has  to supply it again rather than echoing what it read. | [optional] 
**EnableSSL** | Pointer to **bool** | Whether the connection to the mail server is encrypted. | [optional] 
**EnableAuth** | Pointer to **bool** | Whether the portal signs in to the mail server at all. While it is `false` the credentials above are  ignored and the server is expected to accept mail unauthenticated. | [optional] 
**UseNtlm** | Pointer to **bool** | Always `false` here: the flag is accepted when settings are saved but is not stored, so it never comes  back set and says nothing about how the portal authenticates. | [optional] 
**IsDefaultSettings** | Pointer to **bool** | Whether the portal is still on the mail configuration of the installation rather than on a relay of its  own. `DELETE api/2.0/smtpsettings/smtp` puts it back to `true`, and while it is `true` on a cloud portal  the fields above are blank rather than showing the installation's server. | [optional] 

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


