# SsoSpCertificateAdvanced

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

### NewSsoSpCertificateAdvanced

`func NewSsoSpCertificateAdvanced() *SsoSpCertificateAdvanced`

NewSsoSpCertificateAdvanced instantiates a new SsoSpCertificateAdvanced object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoSpCertificateAdvancedWithDefaults

`func NewSsoSpCertificateAdvancedWithDefaults() *SsoSpCertificateAdvanced`

NewSsoSpCertificateAdvancedWithDefaults instantiates a new SsoSpCertificateAdvanced object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigningAlgorithm

`func (o *SsoSpCertificateAdvanced) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *SsoSpCertificateAdvanced) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *SsoSpCertificateAdvanced) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.

### HasSigningAlgorithm

`func (o *SsoSpCertificateAdvanced) HasSigningAlgorithm() bool`

HasSigningAlgorithm returns a boolean if a field has been set.

### SetSigningAlgorithmNil

`func (o *SsoSpCertificateAdvanced) SetSigningAlgorithmNil(b bool)`

 SetSigningAlgorithmNil sets the value for SigningAlgorithm to be an explicit nil

### UnsetSigningAlgorithm
`func (o *SsoSpCertificateAdvanced) UnsetSigningAlgorithm()`

UnsetSigningAlgorithm ensures that no value is present for SigningAlgorithm, not even an explicit nil
### GetSignAuthRequests

`func (o *SsoSpCertificateAdvanced) GetSignAuthRequests() bool`

GetSignAuthRequests returns the SignAuthRequests field if non-nil, zero value otherwise.

### GetSignAuthRequestsOk

`func (o *SsoSpCertificateAdvanced) GetSignAuthRequestsOk() (*bool, bool)`

GetSignAuthRequestsOk returns a tuple with the SignAuthRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAuthRequests

`func (o *SsoSpCertificateAdvanced) SetSignAuthRequests(v bool)`

SetSignAuthRequests sets SignAuthRequests field to given value.

### HasSignAuthRequests

`func (o *SsoSpCertificateAdvanced) HasSignAuthRequests() bool`

HasSignAuthRequests returns a boolean if a field has been set.

### GetSignLogoutRequests

`func (o *SsoSpCertificateAdvanced) GetSignLogoutRequests() bool`

GetSignLogoutRequests returns the SignLogoutRequests field if non-nil, zero value otherwise.

### GetSignLogoutRequestsOk

`func (o *SsoSpCertificateAdvanced) GetSignLogoutRequestsOk() (*bool, bool)`

GetSignLogoutRequestsOk returns a tuple with the SignLogoutRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignLogoutRequests

`func (o *SsoSpCertificateAdvanced) SetSignLogoutRequests(v bool)`

SetSignLogoutRequests sets SignLogoutRequests field to given value.

### HasSignLogoutRequests

`func (o *SsoSpCertificateAdvanced) HasSignLogoutRequests() bool`

HasSignLogoutRequests returns a boolean if a field has been set.

### GetSignLogoutResponses

`func (o *SsoSpCertificateAdvanced) GetSignLogoutResponses() bool`

GetSignLogoutResponses returns the SignLogoutResponses field if non-nil, zero value otherwise.

### GetSignLogoutResponsesOk

`func (o *SsoSpCertificateAdvanced) GetSignLogoutResponsesOk() (*bool, bool)`

GetSignLogoutResponsesOk returns a tuple with the SignLogoutResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignLogoutResponses

`func (o *SsoSpCertificateAdvanced) SetSignLogoutResponses(v bool)`

SetSignLogoutResponses sets SignLogoutResponses field to given value.

### HasSignLogoutResponses

`func (o *SsoSpCertificateAdvanced) HasSignLogoutResponses() bool`

HasSignLogoutResponses returns a boolean if a field has been set.

### GetEncryptAlgorithm

`func (o *SsoSpCertificateAdvanced) GetEncryptAlgorithm() string`

GetEncryptAlgorithm returns the EncryptAlgorithm field if non-nil, zero value otherwise.

### GetEncryptAlgorithmOk

`func (o *SsoSpCertificateAdvanced) GetEncryptAlgorithmOk() (*string, bool)`

GetEncryptAlgorithmOk returns a tuple with the EncryptAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAlgorithm

`func (o *SsoSpCertificateAdvanced) SetEncryptAlgorithm(v string)`

SetEncryptAlgorithm sets EncryptAlgorithm field to given value.

### HasEncryptAlgorithm

`func (o *SsoSpCertificateAdvanced) HasEncryptAlgorithm() bool`

HasEncryptAlgorithm returns a boolean if a field has been set.

### SetEncryptAlgorithmNil

`func (o *SsoSpCertificateAdvanced) SetEncryptAlgorithmNil(b bool)`

 SetEncryptAlgorithmNil sets the value for EncryptAlgorithm to be an explicit nil

### UnsetEncryptAlgorithm
`func (o *SsoSpCertificateAdvanced) UnsetEncryptAlgorithm()`

UnsetEncryptAlgorithm ensures that no value is present for EncryptAlgorithm, not even an explicit nil
### GetDecryptAlgorithm

`func (o *SsoSpCertificateAdvanced) GetDecryptAlgorithm() string`

GetDecryptAlgorithm returns the DecryptAlgorithm field if non-nil, zero value otherwise.

### GetDecryptAlgorithmOk

`func (o *SsoSpCertificateAdvanced) GetDecryptAlgorithmOk() (*string, bool)`

GetDecryptAlgorithmOk returns a tuple with the DecryptAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecryptAlgorithm

`func (o *SsoSpCertificateAdvanced) SetDecryptAlgorithm(v string)`

SetDecryptAlgorithm sets DecryptAlgorithm field to given value.

### HasDecryptAlgorithm

`func (o *SsoSpCertificateAdvanced) HasDecryptAlgorithm() bool`

HasDecryptAlgorithm returns a boolean if a field has been set.

### SetDecryptAlgorithmNil

`func (o *SsoSpCertificateAdvanced) SetDecryptAlgorithmNil(b bool)`

 SetDecryptAlgorithmNil sets the value for DecryptAlgorithm to be an explicit nil

### UnsetDecryptAlgorithm
`func (o *SsoSpCertificateAdvanced) UnsetDecryptAlgorithm()`

UnsetDecryptAlgorithm ensures that no value is present for DecryptAlgorithm, not even an explicit nil
### GetEncryptAssertions

`func (o *SsoSpCertificateAdvanced) GetEncryptAssertions() bool`

GetEncryptAssertions returns the EncryptAssertions field if non-nil, zero value otherwise.

### GetEncryptAssertionsOk

`func (o *SsoSpCertificateAdvanced) GetEncryptAssertionsOk() (*bool, bool)`

GetEncryptAssertionsOk returns a tuple with the EncryptAssertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAssertions

`func (o *SsoSpCertificateAdvanced) SetEncryptAssertions(v bool)`

SetEncryptAssertions sets EncryptAssertions field to given value.

### HasEncryptAssertions

`func (o *SsoSpCertificateAdvanced) HasEncryptAssertions() bool`

HasEncryptAssertions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


