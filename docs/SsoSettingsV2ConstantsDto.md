# SsoSettingsV2ConstantsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsoNameIdFormatType** | Pointer to [**SsoNameIdFormatTypeDto**](SsoNameIdFormatTypeDto.md) | The values the `nameIdFormat` of the identity provider settings accepts. The built-in configuration uses  the SAML 2.0 transient format. | [optional] 
**SsoBindingType** | Pointer to [**SsoBindingTypeDto**](SsoBindingTypeDto.md) | The values the `ssoBinding` and `sloBinding` of the identity provider settings accept - how the portal  sends its sign-in and sign-out requests. The built-in configuration uses HTTP POST for both. | [optional] 
**SsoSigningAlgorithmType** | Pointer to [**SsoSigningAlgorithmTypeDto**](SsoSigningAlgorithmTypeDto.md) | The values the `signingAlgorithm` of the service provider certificate and the `verifyAlgorithm` of the  identity provider certificate accept. The built-in configuration uses RSA-SHA1 for both. | [optional] 
**SsoEncryptAlgorithmType** | Pointer to [**SsoEncryptAlgorithmTypeDto**](SsoEncryptAlgorithmTypeDto.md) | The values the `encryptAlgorithm` and `decryptAlgorithm` of the certificate settings accept. The built-in  configuration uses AES-128 everywhere. | [optional] 
**SsoSpCertificateActionType** | Pointer to [**SsoSpCertificateActionTypeDto**](SsoSpCertificateActionTypeDto.md) | The values the `action` of a service provider certificate accepts, which is what the portal's own key  pair may be used for. | [optional] 
**SsoIdpCertificateActionType** | Pointer to [**SsoIdpCertificateActionTypeDto**](SsoIdpCertificateActionTypeDto.md) | The values the `action` of an identity provider certificate accepts, which is what the provider's  certificate may be used for - the mirror image of the service provider actions. | [optional] 

## Methods

### NewSsoSettingsV2ConstantsDto

`func NewSsoSettingsV2ConstantsDto() *SsoSettingsV2ConstantsDto`

NewSsoSettingsV2ConstantsDto instantiates a new SsoSettingsV2ConstantsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoSettingsV2ConstantsDtoWithDefaults

`func NewSsoSettingsV2ConstantsDtoWithDefaults() *SsoSettingsV2ConstantsDto`

NewSsoSettingsV2ConstantsDtoWithDefaults instantiates a new SsoSettingsV2ConstantsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsoNameIdFormatType

`func (o *SsoSettingsV2ConstantsDto) GetSsoNameIdFormatType() SsoNameIdFormatTypeDto`

GetSsoNameIdFormatType returns the SsoNameIdFormatType field if non-nil, zero value otherwise.

### GetSsoNameIdFormatTypeOk

`func (o *SsoSettingsV2ConstantsDto) GetSsoNameIdFormatTypeOk() (*SsoNameIdFormatTypeDto, bool)`

GetSsoNameIdFormatTypeOk returns a tuple with the SsoNameIdFormatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoNameIdFormatType

`func (o *SsoSettingsV2ConstantsDto) SetSsoNameIdFormatType(v SsoNameIdFormatTypeDto)`

SetSsoNameIdFormatType sets SsoNameIdFormatType field to given value.

### HasSsoNameIdFormatType

`func (o *SsoSettingsV2ConstantsDto) HasSsoNameIdFormatType() bool`

HasSsoNameIdFormatType returns a boolean if a field has been set.

### GetSsoBindingType

`func (o *SsoSettingsV2ConstantsDto) GetSsoBindingType() SsoBindingTypeDto`

GetSsoBindingType returns the SsoBindingType field if non-nil, zero value otherwise.

### GetSsoBindingTypeOk

`func (o *SsoSettingsV2ConstantsDto) GetSsoBindingTypeOk() (*SsoBindingTypeDto, bool)`

GetSsoBindingTypeOk returns a tuple with the SsoBindingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoBindingType

`func (o *SsoSettingsV2ConstantsDto) SetSsoBindingType(v SsoBindingTypeDto)`

SetSsoBindingType sets SsoBindingType field to given value.

### HasSsoBindingType

`func (o *SsoSettingsV2ConstantsDto) HasSsoBindingType() bool`

HasSsoBindingType returns a boolean if a field has been set.

### GetSsoSigningAlgorithmType

`func (o *SsoSettingsV2ConstantsDto) GetSsoSigningAlgorithmType() SsoSigningAlgorithmTypeDto`

GetSsoSigningAlgorithmType returns the SsoSigningAlgorithmType field if non-nil, zero value otherwise.

### GetSsoSigningAlgorithmTypeOk

`func (o *SsoSettingsV2ConstantsDto) GetSsoSigningAlgorithmTypeOk() (*SsoSigningAlgorithmTypeDto, bool)`

GetSsoSigningAlgorithmTypeOk returns a tuple with the SsoSigningAlgorithmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoSigningAlgorithmType

`func (o *SsoSettingsV2ConstantsDto) SetSsoSigningAlgorithmType(v SsoSigningAlgorithmTypeDto)`

SetSsoSigningAlgorithmType sets SsoSigningAlgorithmType field to given value.

### HasSsoSigningAlgorithmType

`func (o *SsoSettingsV2ConstantsDto) HasSsoSigningAlgorithmType() bool`

HasSsoSigningAlgorithmType returns a boolean if a field has been set.

### GetSsoEncryptAlgorithmType

`func (o *SsoSettingsV2ConstantsDto) GetSsoEncryptAlgorithmType() SsoEncryptAlgorithmTypeDto`

GetSsoEncryptAlgorithmType returns the SsoEncryptAlgorithmType field if non-nil, zero value otherwise.

### GetSsoEncryptAlgorithmTypeOk

`func (o *SsoSettingsV2ConstantsDto) GetSsoEncryptAlgorithmTypeOk() (*SsoEncryptAlgorithmTypeDto, bool)`

GetSsoEncryptAlgorithmTypeOk returns a tuple with the SsoEncryptAlgorithmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoEncryptAlgorithmType

`func (o *SsoSettingsV2ConstantsDto) SetSsoEncryptAlgorithmType(v SsoEncryptAlgorithmTypeDto)`

SetSsoEncryptAlgorithmType sets SsoEncryptAlgorithmType field to given value.

### HasSsoEncryptAlgorithmType

`func (o *SsoSettingsV2ConstantsDto) HasSsoEncryptAlgorithmType() bool`

HasSsoEncryptAlgorithmType returns a boolean if a field has been set.

### GetSsoSpCertificateActionType

`func (o *SsoSettingsV2ConstantsDto) GetSsoSpCertificateActionType() SsoSpCertificateActionTypeDto`

GetSsoSpCertificateActionType returns the SsoSpCertificateActionType field if non-nil, zero value otherwise.

### GetSsoSpCertificateActionTypeOk

`func (o *SsoSettingsV2ConstantsDto) GetSsoSpCertificateActionTypeOk() (*SsoSpCertificateActionTypeDto, bool)`

GetSsoSpCertificateActionTypeOk returns a tuple with the SsoSpCertificateActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoSpCertificateActionType

`func (o *SsoSettingsV2ConstantsDto) SetSsoSpCertificateActionType(v SsoSpCertificateActionTypeDto)`

SetSsoSpCertificateActionType sets SsoSpCertificateActionType field to given value.

### HasSsoSpCertificateActionType

`func (o *SsoSettingsV2ConstantsDto) HasSsoSpCertificateActionType() bool`

HasSsoSpCertificateActionType returns a boolean if a field has been set.

### GetSsoIdpCertificateActionType

`func (o *SsoSettingsV2ConstantsDto) GetSsoIdpCertificateActionType() SsoIdpCertificateActionTypeDto`

GetSsoIdpCertificateActionType returns the SsoIdpCertificateActionType field if non-nil, zero value otherwise.

### GetSsoIdpCertificateActionTypeOk

`func (o *SsoSettingsV2ConstantsDto) GetSsoIdpCertificateActionTypeOk() (*SsoIdpCertificateActionTypeDto, bool)`

GetSsoIdpCertificateActionTypeOk returns a tuple with the SsoIdpCertificateActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoIdpCertificateActionType

`func (o *SsoSettingsV2ConstantsDto) SetSsoIdpCertificateActionType(v SsoIdpCertificateActionTypeDto)`

SetSsoIdpCertificateActionType sets SsoIdpCertificateActionType field to given value.

### HasSsoIdpCertificateActionType

`func (o *SsoSettingsV2ConstantsDto) HasSsoIdpCertificateActionType() bool`

HasSsoIdpCertificateActionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


