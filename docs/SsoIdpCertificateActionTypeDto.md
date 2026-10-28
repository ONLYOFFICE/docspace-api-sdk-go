# SsoIdpCertificateActionTypeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verification** | Pointer to **NullableString** | The certificate verifies the signatures on what the provider sends, and nothing else - the counterpart of  the service provider's signing action. | [optional] [readonly] 
**Decrypt** | Pointer to **NullableString** | The certificate is used to decrypt what the provider sends, but verifies no signature. | [optional] [readonly] 
**VerificationAndDecrypt** | Pointer to **NullableString** | The certificate does both, which is what a single provider certificate has to be set to. | [optional] [readonly] 

## Methods

### NewSsoIdpCertificateActionTypeDto

`func NewSsoIdpCertificateActionTypeDto() *SsoIdpCertificateActionTypeDto`

NewSsoIdpCertificateActionTypeDto instantiates a new SsoIdpCertificateActionTypeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoIdpCertificateActionTypeDtoWithDefaults

`func NewSsoIdpCertificateActionTypeDtoWithDefaults() *SsoIdpCertificateActionTypeDto`

NewSsoIdpCertificateActionTypeDtoWithDefaults instantiates a new SsoIdpCertificateActionTypeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerification

`func (o *SsoIdpCertificateActionTypeDto) GetVerification() string`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *SsoIdpCertificateActionTypeDto) GetVerificationOk() (*string, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *SsoIdpCertificateActionTypeDto) SetVerification(v string)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *SsoIdpCertificateActionTypeDto) HasVerification() bool`

HasVerification returns a boolean if a field has been set.

### SetVerificationNil

`func (o *SsoIdpCertificateActionTypeDto) SetVerificationNil(b bool)`

 SetVerificationNil sets the value for Verification to be an explicit nil

### UnsetVerification
`func (o *SsoIdpCertificateActionTypeDto) UnsetVerification()`

UnsetVerification ensures that no value is present for Verification, not even an explicit nil
### GetDecrypt

`func (o *SsoIdpCertificateActionTypeDto) GetDecrypt() string`

GetDecrypt returns the Decrypt field if non-nil, zero value otherwise.

### GetDecryptOk

`func (o *SsoIdpCertificateActionTypeDto) GetDecryptOk() (*string, bool)`

GetDecryptOk returns a tuple with the Decrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecrypt

`func (o *SsoIdpCertificateActionTypeDto) SetDecrypt(v string)`

SetDecrypt sets Decrypt field to given value.

### HasDecrypt

`func (o *SsoIdpCertificateActionTypeDto) HasDecrypt() bool`

HasDecrypt returns a boolean if a field has been set.

### SetDecryptNil

`func (o *SsoIdpCertificateActionTypeDto) SetDecryptNil(b bool)`

 SetDecryptNil sets the value for Decrypt to be an explicit nil

### UnsetDecrypt
`func (o *SsoIdpCertificateActionTypeDto) UnsetDecrypt()`

UnsetDecrypt ensures that no value is present for Decrypt, not even an explicit nil
### GetVerificationAndDecrypt

`func (o *SsoIdpCertificateActionTypeDto) GetVerificationAndDecrypt() string`

GetVerificationAndDecrypt returns the VerificationAndDecrypt field if non-nil, zero value otherwise.

### GetVerificationAndDecryptOk

`func (o *SsoIdpCertificateActionTypeDto) GetVerificationAndDecryptOk() (*string, bool)`

GetVerificationAndDecryptOk returns a tuple with the VerificationAndDecrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationAndDecrypt

`func (o *SsoIdpCertificateActionTypeDto) SetVerificationAndDecrypt(v string)`

SetVerificationAndDecrypt sets VerificationAndDecrypt field to given value.

### HasVerificationAndDecrypt

`func (o *SsoIdpCertificateActionTypeDto) HasVerificationAndDecrypt() bool`

HasVerificationAndDecrypt returns a boolean if a field has been set.

### SetVerificationAndDecryptNil

`func (o *SsoIdpCertificateActionTypeDto) SetVerificationAndDecryptNil(b bool)`

 SetVerificationAndDecryptNil sets the value for VerificationAndDecrypt to be an explicit nil

### UnsetVerificationAndDecrypt
`func (o *SsoIdpCertificateActionTypeDto) UnsetVerificationAndDecrypt()`

UnsetVerificationAndDecrypt ensures that no value is present for VerificationAndDecrypt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


