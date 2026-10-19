# SsoSpCertificateAdvancedDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SigningAlgorithm** | Pointer to **NullableString** | The certificate signing algorithm. | [optional] 
**SignAuthRequests** | Pointer to **bool** | Specifies if SP will sign the SAML authentication requests sent to IdP or not. | [optional] 
**SignLogoutRequests** | Pointer to **bool** | Specifies if SP will sign the SAML logout requests sent to IdP or not. | [optional] 
**SignLogoutResponses** | Pointer to **bool** | Specifies if SP will sign the SAML logout responses sent to IdP or not. | [optional] 
**EncryptAlgorithm** | Pointer to **NullableString** | The certificate encryption algorithm. | [optional] 
**DecryptAlgorithm** | Pointer to **NullableString** | The certificate decryption algorithm. | [optional] 
**EncryptAssertions** | Pointer to **bool** | Specifies if the assertions will be encrypted or not. | [optional] 

## Methods

### NewSsoSpCertificateAdvancedDto

`func NewSsoSpCertificateAdvancedDto() *SsoSpCertificateAdvancedDto`

NewSsoSpCertificateAdvancedDto instantiates a new SsoSpCertificateAdvancedDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoSpCertificateAdvancedDtoWithDefaults

`func NewSsoSpCertificateAdvancedDtoWithDefaults() *SsoSpCertificateAdvancedDto`

NewSsoSpCertificateAdvancedDtoWithDefaults instantiates a new SsoSpCertificateAdvancedDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigningAlgorithm

`func (o *SsoSpCertificateAdvancedDto) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *SsoSpCertificateAdvancedDto) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *SsoSpCertificateAdvancedDto) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.

### HasSigningAlgorithm

`func (o *SsoSpCertificateAdvancedDto) HasSigningAlgorithm() bool`

HasSigningAlgorithm returns a boolean if a field has been set.

### SetSigningAlgorithmNil

`func (o *SsoSpCertificateAdvancedDto) SetSigningAlgorithmNil(b bool)`

 SetSigningAlgorithmNil sets the value for SigningAlgorithm to be an explicit nil

### UnsetSigningAlgorithm
`func (o *SsoSpCertificateAdvancedDto) UnsetSigningAlgorithm()`

UnsetSigningAlgorithm ensures that no value is present for SigningAlgorithm, not even an explicit nil
### GetSignAuthRequests

`func (o *SsoSpCertificateAdvancedDto) GetSignAuthRequests() bool`

GetSignAuthRequests returns the SignAuthRequests field if non-nil, zero value otherwise.

### GetSignAuthRequestsOk

`func (o *SsoSpCertificateAdvancedDto) GetSignAuthRequestsOk() (*bool, bool)`

GetSignAuthRequestsOk returns a tuple with the SignAuthRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAuthRequests

`func (o *SsoSpCertificateAdvancedDto) SetSignAuthRequests(v bool)`

SetSignAuthRequests sets SignAuthRequests field to given value.

### HasSignAuthRequests

`func (o *SsoSpCertificateAdvancedDto) HasSignAuthRequests() bool`

HasSignAuthRequests returns a boolean if a field has been set.

### GetSignLogoutRequests

`func (o *SsoSpCertificateAdvancedDto) GetSignLogoutRequests() bool`

GetSignLogoutRequests returns the SignLogoutRequests field if non-nil, zero value otherwise.

### GetSignLogoutRequestsOk

`func (o *SsoSpCertificateAdvancedDto) GetSignLogoutRequestsOk() (*bool, bool)`

GetSignLogoutRequestsOk returns a tuple with the SignLogoutRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignLogoutRequests

`func (o *SsoSpCertificateAdvancedDto) SetSignLogoutRequests(v bool)`

SetSignLogoutRequests sets SignLogoutRequests field to given value.

### HasSignLogoutRequests

`func (o *SsoSpCertificateAdvancedDto) HasSignLogoutRequests() bool`

HasSignLogoutRequests returns a boolean if a field has been set.

### GetSignLogoutResponses

`func (o *SsoSpCertificateAdvancedDto) GetSignLogoutResponses() bool`

GetSignLogoutResponses returns the SignLogoutResponses field if non-nil, zero value otherwise.

### GetSignLogoutResponsesOk

`func (o *SsoSpCertificateAdvancedDto) GetSignLogoutResponsesOk() (*bool, bool)`

GetSignLogoutResponsesOk returns a tuple with the SignLogoutResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignLogoutResponses

`func (o *SsoSpCertificateAdvancedDto) SetSignLogoutResponses(v bool)`

SetSignLogoutResponses sets SignLogoutResponses field to given value.

### HasSignLogoutResponses

`func (o *SsoSpCertificateAdvancedDto) HasSignLogoutResponses() bool`

HasSignLogoutResponses returns a boolean if a field has been set.

### GetEncryptAlgorithm

`func (o *SsoSpCertificateAdvancedDto) GetEncryptAlgorithm() string`

GetEncryptAlgorithm returns the EncryptAlgorithm field if non-nil, zero value otherwise.

### GetEncryptAlgorithmOk

`func (o *SsoSpCertificateAdvancedDto) GetEncryptAlgorithmOk() (*string, bool)`

GetEncryptAlgorithmOk returns a tuple with the EncryptAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAlgorithm

`func (o *SsoSpCertificateAdvancedDto) SetEncryptAlgorithm(v string)`

SetEncryptAlgorithm sets EncryptAlgorithm field to given value.

### HasEncryptAlgorithm

`func (o *SsoSpCertificateAdvancedDto) HasEncryptAlgorithm() bool`

HasEncryptAlgorithm returns a boolean if a field has been set.

### SetEncryptAlgorithmNil

`func (o *SsoSpCertificateAdvancedDto) SetEncryptAlgorithmNil(b bool)`

 SetEncryptAlgorithmNil sets the value for EncryptAlgorithm to be an explicit nil

### UnsetEncryptAlgorithm
`func (o *SsoSpCertificateAdvancedDto) UnsetEncryptAlgorithm()`

UnsetEncryptAlgorithm ensures that no value is present for EncryptAlgorithm, not even an explicit nil
### GetDecryptAlgorithm

`func (o *SsoSpCertificateAdvancedDto) GetDecryptAlgorithm() string`

GetDecryptAlgorithm returns the DecryptAlgorithm field if non-nil, zero value otherwise.

### GetDecryptAlgorithmOk

`func (o *SsoSpCertificateAdvancedDto) GetDecryptAlgorithmOk() (*string, bool)`

GetDecryptAlgorithmOk returns a tuple with the DecryptAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecryptAlgorithm

`func (o *SsoSpCertificateAdvancedDto) SetDecryptAlgorithm(v string)`

SetDecryptAlgorithm sets DecryptAlgorithm field to given value.

### HasDecryptAlgorithm

`func (o *SsoSpCertificateAdvancedDto) HasDecryptAlgorithm() bool`

HasDecryptAlgorithm returns a boolean if a field has been set.

### SetDecryptAlgorithmNil

`func (o *SsoSpCertificateAdvancedDto) SetDecryptAlgorithmNil(b bool)`

 SetDecryptAlgorithmNil sets the value for DecryptAlgorithm to be an explicit nil

### UnsetDecryptAlgorithm
`func (o *SsoSpCertificateAdvancedDto) UnsetDecryptAlgorithm()`

UnsetDecryptAlgorithm ensures that no value is present for DecryptAlgorithm, not even an explicit nil
### GetEncryptAssertions

`func (o *SsoSpCertificateAdvancedDto) GetEncryptAssertions() bool`

GetEncryptAssertions returns the EncryptAssertions field if non-nil, zero value otherwise.

### GetEncryptAssertionsOk

`func (o *SsoSpCertificateAdvancedDto) GetEncryptAssertionsOk() (*bool, bool)`

GetEncryptAssertionsOk returns a tuple with the EncryptAssertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAssertions

`func (o *SsoSpCertificateAdvancedDto) SetEncryptAssertions(v bool)`

SetEncryptAssertions sets EncryptAssertions field to given value.

### HasEncryptAssertions

`func (o *SsoSpCertificateAdvancedDto) HasEncryptAssertions() bool`

HasEncryptAssertions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


