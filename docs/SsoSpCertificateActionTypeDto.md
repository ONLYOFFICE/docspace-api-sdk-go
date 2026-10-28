# SsoSpCertificateActionTypeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signing** | Pointer to **NullableString** | The key pair signs the requests the portal sends and nothing else. | [optional] [readonly] 
**Encrypt** | Pointer to **NullableString** | The key pair encrypts what the portal sends and decrypts what comes back, but signs nothing. | [optional] [readonly] 
**SigningAndEncrypt** | Pointer to **NullableString** | The key pair does both, which is what one pair configured on its own has to be set to. | [optional] [readonly] 

## Methods

### NewSsoSpCertificateActionTypeDto

`func NewSsoSpCertificateActionTypeDto() *SsoSpCertificateActionTypeDto`

NewSsoSpCertificateActionTypeDto instantiates a new SsoSpCertificateActionTypeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoSpCertificateActionTypeDtoWithDefaults

`func NewSsoSpCertificateActionTypeDtoWithDefaults() *SsoSpCertificateActionTypeDto`

NewSsoSpCertificateActionTypeDtoWithDefaults instantiates a new SsoSpCertificateActionTypeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigning

`func (o *SsoSpCertificateActionTypeDto) GetSigning() string`

GetSigning returns the Signing field if non-nil, zero value otherwise.

### GetSigningOk

`func (o *SsoSpCertificateActionTypeDto) GetSigningOk() (*string, bool)`

GetSigningOk returns a tuple with the Signing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigning

`func (o *SsoSpCertificateActionTypeDto) SetSigning(v string)`

SetSigning sets Signing field to given value.

### HasSigning

`func (o *SsoSpCertificateActionTypeDto) HasSigning() bool`

HasSigning returns a boolean if a field has been set.

### SetSigningNil

`func (o *SsoSpCertificateActionTypeDto) SetSigningNil(b bool)`

 SetSigningNil sets the value for Signing to be an explicit nil

### UnsetSigning
`func (o *SsoSpCertificateActionTypeDto) UnsetSigning()`

UnsetSigning ensures that no value is present for Signing, not even an explicit nil
### GetEncrypt

`func (o *SsoSpCertificateActionTypeDto) GetEncrypt() string`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *SsoSpCertificateActionTypeDto) GetEncryptOk() (*string, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *SsoSpCertificateActionTypeDto) SetEncrypt(v string)`

SetEncrypt sets Encrypt field to given value.

### HasEncrypt

`func (o *SsoSpCertificateActionTypeDto) HasEncrypt() bool`

HasEncrypt returns a boolean if a field has been set.

### SetEncryptNil

`func (o *SsoSpCertificateActionTypeDto) SetEncryptNil(b bool)`

 SetEncryptNil sets the value for Encrypt to be an explicit nil

### UnsetEncrypt
`func (o *SsoSpCertificateActionTypeDto) UnsetEncrypt()`

UnsetEncrypt ensures that no value is present for Encrypt, not even an explicit nil
### GetSigningAndEncrypt

`func (o *SsoSpCertificateActionTypeDto) GetSigningAndEncrypt() string`

GetSigningAndEncrypt returns the SigningAndEncrypt field if non-nil, zero value otherwise.

### GetSigningAndEncryptOk

`func (o *SsoSpCertificateActionTypeDto) GetSigningAndEncryptOk() (*string, bool)`

GetSigningAndEncryptOk returns a tuple with the SigningAndEncrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAndEncrypt

`func (o *SsoSpCertificateActionTypeDto) SetSigningAndEncrypt(v string)`

SetSigningAndEncrypt sets SigningAndEncrypt field to given value.

### HasSigningAndEncrypt

`func (o *SsoSpCertificateActionTypeDto) HasSigningAndEncrypt() bool`

HasSigningAndEncrypt returns a boolean if a field has been set.

### SetSigningAndEncryptNil

`func (o *SsoSpCertificateActionTypeDto) SetSigningAndEncryptNil(b bool)`

 SetSigningAndEncryptNil sets the value for SigningAndEncrypt to be an explicit nil

### UnsetSigningAndEncrypt
`func (o *SsoSpCertificateActionTypeDto) UnsetSigningAndEncrypt()`

UnsetSigningAndEncrypt ensures that no value is present for SigningAndEncrypt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


