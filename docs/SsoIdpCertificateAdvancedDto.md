# SsoIdpCertificateAdvancedDto

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

### NewSsoIdpCertificateAdvancedDto

`func NewSsoIdpCertificateAdvancedDto() *SsoIdpCertificateAdvancedDto`

NewSsoIdpCertificateAdvancedDto instantiates a new SsoIdpCertificateAdvancedDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoIdpCertificateAdvancedDtoWithDefaults

`func NewSsoIdpCertificateAdvancedDtoWithDefaults() *SsoIdpCertificateAdvancedDto`

NewSsoIdpCertificateAdvancedDtoWithDefaults instantiates a new SsoIdpCertificateAdvancedDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerifyAlgorithm

`func (o *SsoIdpCertificateAdvancedDto) GetVerifyAlgorithm() string`

GetVerifyAlgorithm returns the VerifyAlgorithm field if non-nil, zero value otherwise.

### GetVerifyAlgorithmOk

`func (o *SsoIdpCertificateAdvancedDto) GetVerifyAlgorithmOk() (*string, bool)`

GetVerifyAlgorithmOk returns a tuple with the VerifyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyAlgorithm

`func (o *SsoIdpCertificateAdvancedDto) SetVerifyAlgorithm(v string)`

SetVerifyAlgorithm sets VerifyAlgorithm field to given value.

### HasVerifyAlgorithm

`func (o *SsoIdpCertificateAdvancedDto) HasVerifyAlgorithm() bool`

HasVerifyAlgorithm returns a boolean if a field has been set.

### SetVerifyAlgorithmNil

`func (o *SsoIdpCertificateAdvancedDto) SetVerifyAlgorithmNil(b bool)`

 SetVerifyAlgorithmNil sets the value for VerifyAlgorithm to be an explicit nil

### UnsetVerifyAlgorithm
`func (o *SsoIdpCertificateAdvancedDto) UnsetVerifyAlgorithm()`

UnsetVerifyAlgorithm ensures that no value is present for VerifyAlgorithm, not even an explicit nil
### GetVerifyAuthResponsesSign

`func (o *SsoIdpCertificateAdvancedDto) GetVerifyAuthResponsesSign() bool`

GetVerifyAuthResponsesSign returns the VerifyAuthResponsesSign field if non-nil, zero value otherwise.

### GetVerifyAuthResponsesSignOk

`func (o *SsoIdpCertificateAdvancedDto) GetVerifyAuthResponsesSignOk() (*bool, bool)`

GetVerifyAuthResponsesSignOk returns a tuple with the VerifyAuthResponsesSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyAuthResponsesSign

`func (o *SsoIdpCertificateAdvancedDto) SetVerifyAuthResponsesSign(v bool)`

SetVerifyAuthResponsesSign sets VerifyAuthResponsesSign field to given value.

### HasVerifyAuthResponsesSign

`func (o *SsoIdpCertificateAdvancedDto) HasVerifyAuthResponsesSign() bool`

HasVerifyAuthResponsesSign returns a boolean if a field has been set.

### GetVerifyLogoutRequestsSign

`func (o *SsoIdpCertificateAdvancedDto) GetVerifyLogoutRequestsSign() bool`

GetVerifyLogoutRequestsSign returns the VerifyLogoutRequestsSign field if non-nil, zero value otherwise.

### GetVerifyLogoutRequestsSignOk

`func (o *SsoIdpCertificateAdvancedDto) GetVerifyLogoutRequestsSignOk() (*bool, bool)`

GetVerifyLogoutRequestsSignOk returns a tuple with the VerifyLogoutRequestsSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyLogoutRequestsSign

`func (o *SsoIdpCertificateAdvancedDto) SetVerifyLogoutRequestsSign(v bool)`

SetVerifyLogoutRequestsSign sets VerifyLogoutRequestsSign field to given value.

### HasVerifyLogoutRequestsSign

`func (o *SsoIdpCertificateAdvancedDto) HasVerifyLogoutRequestsSign() bool`

HasVerifyLogoutRequestsSign returns a boolean if a field has been set.

### GetVerifyLogoutResponsesSign

`func (o *SsoIdpCertificateAdvancedDto) GetVerifyLogoutResponsesSign() bool`

GetVerifyLogoutResponsesSign returns the VerifyLogoutResponsesSign field if non-nil, zero value otherwise.

### GetVerifyLogoutResponsesSignOk

`func (o *SsoIdpCertificateAdvancedDto) GetVerifyLogoutResponsesSignOk() (*bool, bool)`

GetVerifyLogoutResponsesSignOk returns a tuple with the VerifyLogoutResponsesSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyLogoutResponsesSign

`func (o *SsoIdpCertificateAdvancedDto) SetVerifyLogoutResponsesSign(v bool)`

SetVerifyLogoutResponsesSign sets VerifyLogoutResponsesSign field to given value.

### HasVerifyLogoutResponsesSign

`func (o *SsoIdpCertificateAdvancedDto) HasVerifyLogoutResponsesSign() bool`

HasVerifyLogoutResponsesSign returns a boolean if a field has been set.

### GetDecryptAlgorithm

`func (o *SsoIdpCertificateAdvancedDto) GetDecryptAlgorithm() string`

GetDecryptAlgorithm returns the DecryptAlgorithm field if non-nil, zero value otherwise.

### GetDecryptAlgorithmOk

`func (o *SsoIdpCertificateAdvancedDto) GetDecryptAlgorithmOk() (*string, bool)`

GetDecryptAlgorithmOk returns a tuple with the DecryptAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecryptAlgorithm

`func (o *SsoIdpCertificateAdvancedDto) SetDecryptAlgorithm(v string)`

SetDecryptAlgorithm sets DecryptAlgorithm field to given value.

### HasDecryptAlgorithm

`func (o *SsoIdpCertificateAdvancedDto) HasDecryptAlgorithm() bool`

HasDecryptAlgorithm returns a boolean if a field has been set.

### SetDecryptAlgorithmNil

`func (o *SsoIdpCertificateAdvancedDto) SetDecryptAlgorithmNil(b bool)`

 SetDecryptAlgorithmNil sets the value for DecryptAlgorithm to be an explicit nil

### UnsetDecryptAlgorithm
`func (o *SsoIdpCertificateAdvancedDto) UnsetDecryptAlgorithm()`

UnsetDecryptAlgorithm ensures that no value is present for DecryptAlgorithm, not even an explicit nil
### GetDecryptAssertions

`func (o *SsoIdpCertificateAdvancedDto) GetDecryptAssertions() bool`

GetDecryptAssertions returns the DecryptAssertions field if non-nil, zero value otherwise.

### GetDecryptAssertionsOk

`func (o *SsoIdpCertificateAdvancedDto) GetDecryptAssertionsOk() (*bool, bool)`

GetDecryptAssertionsOk returns a tuple with the DecryptAssertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecryptAssertions

`func (o *SsoIdpCertificateAdvancedDto) SetDecryptAssertions(v bool)`

SetDecryptAssertions sets DecryptAssertions field to given value.

### HasDecryptAssertions

`func (o *SsoIdpCertificateAdvancedDto) HasDecryptAssertions() bool`

HasDecryptAssertions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


