# SsoNameIdFormatTypeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Saml11Unspecified** | Pointer to **NullableString** | The SAML 1.1 unspecified name ID format. | [optional] [readonly] 
**Saml11EmailAddress** | Pointer to **NullableString** | The SAML 1.1 email address name ID format. | [optional] [readonly] 
**Saml20Entity** | Pointer to **NullableString** | The SAML 2.0 entity name ID format. | [optional] [readonly] 
**Saml20Transient** | Pointer to **NullableString** | The SAML 2.0 transient name ID format, whose identifier differs from one session to the next. It is what  the built-in configuration uses. | [optional] [readonly] 
**Saml20Persistent** | Pointer to **NullableString** | The SAML 2.0 persistent name ID format, whose identifier stays the same for one person across sessions. | [optional] [readonly] 
**Saml20Encrypted** | Pointer to **NullableString** | The SAML 2.0 encrypted name ID format. | [optional] [readonly] 
**Saml20Unspecified** | Pointer to **NullableString** | The SAML 2.0 unspecified name ID format. | [optional] [readonly] 
**Saml11X509SubjectName** | Pointer to **NullableString** | The SAML 1.1 X.509 subject name name ID format. | [optional] [readonly] 
**Saml11WindowsDomainQualifiedName** | Pointer to **NullableString** | The SAML 1.1 Windows domain qualified name name ID format. | [optional] [readonly] 
**Saml20Kerberos** | Pointer to **NullableString** | The SAML 2.0 Kerberos name ID format. | [optional] [readonly] 

## Methods

### NewSsoNameIdFormatTypeDto

`func NewSsoNameIdFormatTypeDto() *SsoNameIdFormatTypeDto`

NewSsoNameIdFormatTypeDto instantiates a new SsoNameIdFormatTypeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoNameIdFormatTypeDtoWithDefaults

`func NewSsoNameIdFormatTypeDtoWithDefaults() *SsoNameIdFormatTypeDto`

NewSsoNameIdFormatTypeDtoWithDefaults instantiates a new SsoNameIdFormatTypeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSaml11Unspecified

`func (o *SsoNameIdFormatTypeDto) GetSaml11Unspecified() string`

GetSaml11Unspecified returns the Saml11Unspecified field if non-nil, zero value otherwise.

### GetSaml11UnspecifiedOk

`func (o *SsoNameIdFormatTypeDto) GetSaml11UnspecifiedOk() (*string, bool)`

GetSaml11UnspecifiedOk returns a tuple with the Saml11Unspecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml11Unspecified

`func (o *SsoNameIdFormatTypeDto) SetSaml11Unspecified(v string)`

SetSaml11Unspecified sets Saml11Unspecified field to given value.

### HasSaml11Unspecified

`func (o *SsoNameIdFormatTypeDto) HasSaml11Unspecified() bool`

HasSaml11Unspecified returns a boolean if a field has been set.

### SetSaml11UnspecifiedNil

`func (o *SsoNameIdFormatTypeDto) SetSaml11UnspecifiedNil(b bool)`

 SetSaml11UnspecifiedNil sets the value for Saml11Unspecified to be an explicit nil

### UnsetSaml11Unspecified
`func (o *SsoNameIdFormatTypeDto) UnsetSaml11Unspecified()`

UnsetSaml11Unspecified ensures that no value is present for Saml11Unspecified, not even an explicit nil
### GetSaml11EmailAddress

`func (o *SsoNameIdFormatTypeDto) GetSaml11EmailAddress() string`

GetSaml11EmailAddress returns the Saml11EmailAddress field if non-nil, zero value otherwise.

### GetSaml11EmailAddressOk

`func (o *SsoNameIdFormatTypeDto) GetSaml11EmailAddressOk() (*string, bool)`

GetSaml11EmailAddressOk returns a tuple with the Saml11EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml11EmailAddress

`func (o *SsoNameIdFormatTypeDto) SetSaml11EmailAddress(v string)`

SetSaml11EmailAddress sets Saml11EmailAddress field to given value.

### HasSaml11EmailAddress

`func (o *SsoNameIdFormatTypeDto) HasSaml11EmailAddress() bool`

HasSaml11EmailAddress returns a boolean if a field has been set.

### SetSaml11EmailAddressNil

`func (o *SsoNameIdFormatTypeDto) SetSaml11EmailAddressNil(b bool)`

 SetSaml11EmailAddressNil sets the value for Saml11EmailAddress to be an explicit nil

### UnsetSaml11EmailAddress
`func (o *SsoNameIdFormatTypeDto) UnsetSaml11EmailAddress()`

UnsetSaml11EmailAddress ensures that no value is present for Saml11EmailAddress, not even an explicit nil
### GetSaml20Entity

`func (o *SsoNameIdFormatTypeDto) GetSaml20Entity() string`

GetSaml20Entity returns the Saml20Entity field if non-nil, zero value otherwise.

### GetSaml20EntityOk

`func (o *SsoNameIdFormatTypeDto) GetSaml20EntityOk() (*string, bool)`

GetSaml20EntityOk returns a tuple with the Saml20Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml20Entity

`func (o *SsoNameIdFormatTypeDto) SetSaml20Entity(v string)`

SetSaml20Entity sets Saml20Entity field to given value.

### HasSaml20Entity

`func (o *SsoNameIdFormatTypeDto) HasSaml20Entity() bool`

HasSaml20Entity returns a boolean if a field has been set.

### SetSaml20EntityNil

`func (o *SsoNameIdFormatTypeDto) SetSaml20EntityNil(b bool)`

 SetSaml20EntityNil sets the value for Saml20Entity to be an explicit nil

### UnsetSaml20Entity
`func (o *SsoNameIdFormatTypeDto) UnsetSaml20Entity()`

UnsetSaml20Entity ensures that no value is present for Saml20Entity, not even an explicit nil
### GetSaml20Transient

`func (o *SsoNameIdFormatTypeDto) GetSaml20Transient() string`

GetSaml20Transient returns the Saml20Transient field if non-nil, zero value otherwise.

### GetSaml20TransientOk

`func (o *SsoNameIdFormatTypeDto) GetSaml20TransientOk() (*string, bool)`

GetSaml20TransientOk returns a tuple with the Saml20Transient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml20Transient

`func (o *SsoNameIdFormatTypeDto) SetSaml20Transient(v string)`

SetSaml20Transient sets Saml20Transient field to given value.

### HasSaml20Transient

`func (o *SsoNameIdFormatTypeDto) HasSaml20Transient() bool`

HasSaml20Transient returns a boolean if a field has been set.

### SetSaml20TransientNil

`func (o *SsoNameIdFormatTypeDto) SetSaml20TransientNil(b bool)`

 SetSaml20TransientNil sets the value for Saml20Transient to be an explicit nil

### UnsetSaml20Transient
`func (o *SsoNameIdFormatTypeDto) UnsetSaml20Transient()`

UnsetSaml20Transient ensures that no value is present for Saml20Transient, not even an explicit nil
### GetSaml20Persistent

`func (o *SsoNameIdFormatTypeDto) GetSaml20Persistent() string`

GetSaml20Persistent returns the Saml20Persistent field if non-nil, zero value otherwise.

### GetSaml20PersistentOk

`func (o *SsoNameIdFormatTypeDto) GetSaml20PersistentOk() (*string, bool)`

GetSaml20PersistentOk returns a tuple with the Saml20Persistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml20Persistent

`func (o *SsoNameIdFormatTypeDto) SetSaml20Persistent(v string)`

SetSaml20Persistent sets Saml20Persistent field to given value.

### HasSaml20Persistent

`func (o *SsoNameIdFormatTypeDto) HasSaml20Persistent() bool`

HasSaml20Persistent returns a boolean if a field has been set.

### SetSaml20PersistentNil

`func (o *SsoNameIdFormatTypeDto) SetSaml20PersistentNil(b bool)`

 SetSaml20PersistentNil sets the value for Saml20Persistent to be an explicit nil

### UnsetSaml20Persistent
`func (o *SsoNameIdFormatTypeDto) UnsetSaml20Persistent()`

UnsetSaml20Persistent ensures that no value is present for Saml20Persistent, not even an explicit nil
### GetSaml20Encrypted

`func (o *SsoNameIdFormatTypeDto) GetSaml20Encrypted() string`

GetSaml20Encrypted returns the Saml20Encrypted field if non-nil, zero value otherwise.

### GetSaml20EncryptedOk

`func (o *SsoNameIdFormatTypeDto) GetSaml20EncryptedOk() (*string, bool)`

GetSaml20EncryptedOk returns a tuple with the Saml20Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml20Encrypted

`func (o *SsoNameIdFormatTypeDto) SetSaml20Encrypted(v string)`

SetSaml20Encrypted sets Saml20Encrypted field to given value.

### HasSaml20Encrypted

`func (o *SsoNameIdFormatTypeDto) HasSaml20Encrypted() bool`

HasSaml20Encrypted returns a boolean if a field has been set.

### SetSaml20EncryptedNil

`func (o *SsoNameIdFormatTypeDto) SetSaml20EncryptedNil(b bool)`

 SetSaml20EncryptedNil sets the value for Saml20Encrypted to be an explicit nil

### UnsetSaml20Encrypted
`func (o *SsoNameIdFormatTypeDto) UnsetSaml20Encrypted()`

UnsetSaml20Encrypted ensures that no value is present for Saml20Encrypted, not even an explicit nil
### GetSaml20Unspecified

`func (o *SsoNameIdFormatTypeDto) GetSaml20Unspecified() string`

GetSaml20Unspecified returns the Saml20Unspecified field if non-nil, zero value otherwise.

### GetSaml20UnspecifiedOk

`func (o *SsoNameIdFormatTypeDto) GetSaml20UnspecifiedOk() (*string, bool)`

GetSaml20UnspecifiedOk returns a tuple with the Saml20Unspecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml20Unspecified

`func (o *SsoNameIdFormatTypeDto) SetSaml20Unspecified(v string)`

SetSaml20Unspecified sets Saml20Unspecified field to given value.

### HasSaml20Unspecified

`func (o *SsoNameIdFormatTypeDto) HasSaml20Unspecified() bool`

HasSaml20Unspecified returns a boolean if a field has been set.

### SetSaml20UnspecifiedNil

`func (o *SsoNameIdFormatTypeDto) SetSaml20UnspecifiedNil(b bool)`

 SetSaml20UnspecifiedNil sets the value for Saml20Unspecified to be an explicit nil

### UnsetSaml20Unspecified
`func (o *SsoNameIdFormatTypeDto) UnsetSaml20Unspecified()`

UnsetSaml20Unspecified ensures that no value is present for Saml20Unspecified, not even an explicit nil
### GetSaml11X509SubjectName

`func (o *SsoNameIdFormatTypeDto) GetSaml11X509SubjectName() string`

GetSaml11X509SubjectName returns the Saml11X509SubjectName field if non-nil, zero value otherwise.

### GetSaml11X509SubjectNameOk

`func (o *SsoNameIdFormatTypeDto) GetSaml11X509SubjectNameOk() (*string, bool)`

GetSaml11X509SubjectNameOk returns a tuple with the Saml11X509SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml11X509SubjectName

`func (o *SsoNameIdFormatTypeDto) SetSaml11X509SubjectName(v string)`

SetSaml11X509SubjectName sets Saml11X509SubjectName field to given value.

### HasSaml11X509SubjectName

`func (o *SsoNameIdFormatTypeDto) HasSaml11X509SubjectName() bool`

HasSaml11X509SubjectName returns a boolean if a field has been set.

### SetSaml11X509SubjectNameNil

`func (o *SsoNameIdFormatTypeDto) SetSaml11X509SubjectNameNil(b bool)`

 SetSaml11X509SubjectNameNil sets the value for Saml11X509SubjectName to be an explicit nil

### UnsetSaml11X509SubjectName
`func (o *SsoNameIdFormatTypeDto) UnsetSaml11X509SubjectName()`

UnsetSaml11X509SubjectName ensures that no value is present for Saml11X509SubjectName, not even an explicit nil
### GetSaml11WindowsDomainQualifiedName

`func (o *SsoNameIdFormatTypeDto) GetSaml11WindowsDomainQualifiedName() string`

GetSaml11WindowsDomainQualifiedName returns the Saml11WindowsDomainQualifiedName field if non-nil, zero value otherwise.

### GetSaml11WindowsDomainQualifiedNameOk

`func (o *SsoNameIdFormatTypeDto) GetSaml11WindowsDomainQualifiedNameOk() (*string, bool)`

GetSaml11WindowsDomainQualifiedNameOk returns a tuple with the Saml11WindowsDomainQualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml11WindowsDomainQualifiedName

`func (o *SsoNameIdFormatTypeDto) SetSaml11WindowsDomainQualifiedName(v string)`

SetSaml11WindowsDomainQualifiedName sets Saml11WindowsDomainQualifiedName field to given value.

### HasSaml11WindowsDomainQualifiedName

`func (o *SsoNameIdFormatTypeDto) HasSaml11WindowsDomainQualifiedName() bool`

HasSaml11WindowsDomainQualifiedName returns a boolean if a field has been set.

### SetSaml11WindowsDomainQualifiedNameNil

`func (o *SsoNameIdFormatTypeDto) SetSaml11WindowsDomainQualifiedNameNil(b bool)`

 SetSaml11WindowsDomainQualifiedNameNil sets the value for Saml11WindowsDomainQualifiedName to be an explicit nil

### UnsetSaml11WindowsDomainQualifiedName
`func (o *SsoNameIdFormatTypeDto) UnsetSaml11WindowsDomainQualifiedName()`

UnsetSaml11WindowsDomainQualifiedName ensures that no value is present for Saml11WindowsDomainQualifiedName, not even an explicit nil
### GetSaml20Kerberos

`func (o *SsoNameIdFormatTypeDto) GetSaml20Kerberos() string`

GetSaml20Kerberos returns the Saml20Kerberos field if non-nil, zero value otherwise.

### GetSaml20KerberosOk

`func (o *SsoNameIdFormatTypeDto) GetSaml20KerberosOk() (*string, bool)`

GetSaml20KerberosOk returns a tuple with the Saml20Kerberos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml20Kerberos

`func (o *SsoNameIdFormatTypeDto) SetSaml20Kerberos(v string)`

SetSaml20Kerberos sets Saml20Kerberos field to given value.

### HasSaml20Kerberos

`func (o *SsoNameIdFormatTypeDto) HasSaml20Kerberos() bool`

HasSaml20Kerberos returns a boolean if a field has been set.

### SetSaml20KerberosNil

`func (o *SsoNameIdFormatTypeDto) SetSaml20KerberosNil(b bool)`

 SetSaml20KerberosNil sets the value for Saml20Kerberos to be an explicit nil

### UnsetSaml20Kerberos
`func (o *SsoNameIdFormatTypeDto) UnsetSaml20Kerberos()`

UnsetSaml20Kerberos ensures that no value is present for Saml20Kerberos, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


