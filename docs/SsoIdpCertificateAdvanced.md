# SsoIdpCertificateAdvanced

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerifyAlgorithm** | Pointer to **NullableString** | The certificate verification algorithm. | [optional] 
**VerifyAuthResponsesSign** | Pointer to **bool** | Specifies if the signatures of the SAML authentication responses sent to SP will be verified or not. | [optional] 
**VerifyLogoutRequestsSign** | Pointer to **bool** | Specifies if the signatures of the SAML logout requests sent to SP will be verified or not. | [optional] 
**VerifyLogoutResponsesSign** | Pointer to **bool** | Specifies if the signatures of the SAML logout responses sent to SP will be verified or not. | [optional] 
**DecryptAlgorithm** | Pointer to **NullableString** | The certificate decryption algorithm. | [optional] 
**DecryptAssertions** | Pointer to **bool** | Specifies if the assertions will be decrypted or not. | [optional] 

## Methods

### NewSsoIdpCertificateAdvanced

`func NewSsoIdpCertificateAdvanced() *SsoIdpCertificateAdvanced`

NewSsoIdpCertificateAdvanced instantiates a new SsoIdpCertificateAdvanced object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoIdpCertificateAdvancedWithDefaults

`func NewSsoIdpCertificateAdvancedWithDefaults() *SsoIdpCertificateAdvanced`

NewSsoIdpCertificateAdvancedWithDefaults instantiates a new SsoIdpCertificateAdvanced object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerifyAlgorithm

`func (o *SsoIdpCertificateAdvanced) GetVerifyAlgorithm() string`

GetVerifyAlgorithm returns the VerifyAlgorithm field if non-nil, zero value otherwise.

### GetVerifyAlgorithmOk

`func (o *SsoIdpCertificateAdvanced) GetVerifyAlgorithmOk() (*string, bool)`

GetVerifyAlgorithmOk returns a tuple with the VerifyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyAlgorithm

`func (o *SsoIdpCertificateAdvanced) SetVerifyAlgorithm(v string)`

SetVerifyAlgorithm sets VerifyAlgorithm field to given value.

### HasVerifyAlgorithm

`func (o *SsoIdpCertificateAdvanced) HasVerifyAlgorithm() bool`

HasVerifyAlgorithm returns a boolean if a field has been set.

### SetVerifyAlgorithmNil

`func (o *SsoIdpCertificateAdvanced) SetVerifyAlgorithmNil(b bool)`

 SetVerifyAlgorithmNil sets the value for VerifyAlgorithm to be an explicit nil

### UnsetVerifyAlgorithm
`func (o *SsoIdpCertificateAdvanced) UnsetVerifyAlgorithm()`

UnsetVerifyAlgorithm ensures that no value is present for VerifyAlgorithm, not even an explicit nil
### GetVerifyAuthResponsesSign

`func (o *SsoIdpCertificateAdvanced) GetVerifyAuthResponsesSign() bool`

GetVerifyAuthResponsesSign returns the VerifyAuthResponsesSign field if non-nil, zero value otherwise.

### GetVerifyAuthResponsesSignOk

`func (o *SsoIdpCertificateAdvanced) GetVerifyAuthResponsesSignOk() (*bool, bool)`

GetVerifyAuthResponsesSignOk returns a tuple with the VerifyAuthResponsesSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyAuthResponsesSign

`func (o *SsoIdpCertificateAdvanced) SetVerifyAuthResponsesSign(v bool)`

SetVerifyAuthResponsesSign sets VerifyAuthResponsesSign field to given value.

### HasVerifyAuthResponsesSign

`func (o *SsoIdpCertificateAdvanced) HasVerifyAuthResponsesSign() bool`

HasVerifyAuthResponsesSign returns a boolean if a field has been set.

### GetVerifyLogoutRequestsSign

`func (o *SsoIdpCertificateAdvanced) GetVerifyLogoutRequestsSign() bool`

GetVerifyLogoutRequestsSign returns the VerifyLogoutRequestsSign field if non-nil, zero value otherwise.

### GetVerifyLogoutRequestsSignOk

`func (o *SsoIdpCertificateAdvanced) GetVerifyLogoutRequestsSignOk() (*bool, bool)`

GetVerifyLogoutRequestsSignOk returns a tuple with the VerifyLogoutRequestsSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyLogoutRequestsSign

`func (o *SsoIdpCertificateAdvanced) SetVerifyLogoutRequestsSign(v bool)`

SetVerifyLogoutRequestsSign sets VerifyLogoutRequestsSign field to given value.

### HasVerifyLogoutRequestsSign

`func (o *SsoIdpCertificateAdvanced) HasVerifyLogoutRequestsSign() bool`

HasVerifyLogoutRequestsSign returns a boolean if a field has been set.

### GetVerifyLogoutResponsesSign

`func (o *SsoIdpCertificateAdvanced) GetVerifyLogoutResponsesSign() bool`

GetVerifyLogoutResponsesSign returns the VerifyLogoutResponsesSign field if non-nil, zero value otherwise.

### GetVerifyLogoutResponsesSignOk

`func (o *SsoIdpCertificateAdvanced) GetVerifyLogoutResponsesSignOk() (*bool, bool)`

GetVerifyLogoutResponsesSignOk returns a tuple with the VerifyLogoutResponsesSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyLogoutResponsesSign

`func (o *SsoIdpCertificateAdvanced) SetVerifyLogoutResponsesSign(v bool)`

SetVerifyLogoutResponsesSign sets VerifyLogoutResponsesSign field to given value.

### HasVerifyLogoutResponsesSign

`func (o *SsoIdpCertificateAdvanced) HasVerifyLogoutResponsesSign() bool`

HasVerifyLogoutResponsesSign returns a boolean if a field has been set.

### GetDecryptAlgorithm

`func (o *SsoIdpCertificateAdvanced) GetDecryptAlgorithm() string`

GetDecryptAlgorithm returns the DecryptAlgorithm field if non-nil, zero value otherwise.

### GetDecryptAlgorithmOk

`func (o *SsoIdpCertificateAdvanced) GetDecryptAlgorithmOk() (*string, bool)`

GetDecryptAlgorithmOk returns a tuple with the DecryptAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecryptAlgorithm

`func (o *SsoIdpCertificateAdvanced) SetDecryptAlgorithm(v string)`

SetDecryptAlgorithm sets DecryptAlgorithm field to given value.

### HasDecryptAlgorithm

`func (o *SsoIdpCertificateAdvanced) HasDecryptAlgorithm() bool`

HasDecryptAlgorithm returns a boolean if a field has been set.

### SetDecryptAlgorithmNil

`func (o *SsoIdpCertificateAdvanced) SetDecryptAlgorithmNil(b bool)`

 SetDecryptAlgorithmNil sets the value for DecryptAlgorithm to be an explicit nil

### UnsetDecryptAlgorithm
`func (o *SsoIdpCertificateAdvanced) UnsetDecryptAlgorithm()`

UnsetDecryptAlgorithm ensures that no value is present for DecryptAlgorithm, not even an explicit nil
### GetDecryptAssertions

`func (o *SsoIdpCertificateAdvanced) GetDecryptAssertions() bool`

GetDecryptAssertions returns the DecryptAssertions field if non-nil, zero value otherwise.

### GetDecryptAssertionsOk

`func (o *SsoIdpCertificateAdvanced) GetDecryptAssertionsOk() (*bool, bool)`

GetDecryptAssertionsOk returns a tuple with the DecryptAssertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecryptAssertions

`func (o *SsoIdpCertificateAdvanced) SetDecryptAssertions(v bool)`

SetDecryptAssertions sets DecryptAssertions field to given value.

### HasDecryptAssertions

`func (o *SsoIdpCertificateAdvanced) HasDecryptAssertions() bool`

HasDecryptAssertions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


