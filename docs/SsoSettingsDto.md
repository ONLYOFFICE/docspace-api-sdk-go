# SsoSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastModified** | Pointer to **time.Time** | The timestamp indicating when the settings were last modified. | [optional] 
**EnableSso** | Pointer to **NullableBool** | Specifies if the SSO settings are enabled or not. | [optional] 
**IdpSettings** | Pointer to [**SsoIdpSettingsDto**](SsoIdpSettingsDto.md) | The SSO IdP settings. | [optional] 
**IdpCertificates** | Pointer to [**[]SsoCertificateDto**](SsoCertificateDto.md) | The list of the IdP certificates. | [optional] 
**IdpCertificateAdvanced** | Pointer to [**SsoIdpCertificateAdvancedDto**](SsoIdpCertificateAdvancedDto.md) | The IdP advanced certificate. | [optional] 
**SpLoginLabel** | Pointer to **NullableString** | The SP login label. | [optional] 
**SpCertificates** | Pointer to [**[]SsoCertificateDto**](SsoCertificateDto.md) | The list of the SP certificates. | [optional] 
**SpCertificateAdvanced** | Pointer to [**SsoSpCertificateAdvancedDto**](SsoSpCertificateAdvancedDto.md) | The SP advanced certificate. | [optional] 
**FieldMapping** | Pointer to [**SsoFieldMappingDto**](SsoFieldMappingDto.md) | The SSO field mapping. | [optional] 
**HideAuthPage** | Pointer to **bool** | Specifies if the authentication page will be hidden or not. | [optional] 
**UsersType** | Pointer to **int32** | The user type. | [optional] 
**DisableEmailVerification** | Pointer to **bool** | Specifies if the email verification is disabled or not. | [optional] 

## Methods

### NewSsoSettingsDto

`func NewSsoSettingsDto() *SsoSettingsDto`

NewSsoSettingsDto instantiates a new SsoSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoSettingsDtoWithDefaults

`func NewSsoSettingsDtoWithDefaults() *SsoSettingsDto`

NewSsoSettingsDtoWithDefaults instantiates a new SsoSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastModified

`func (o *SsoSettingsDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *SsoSettingsDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *SsoSettingsDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *SsoSettingsDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetEnableSso

`func (o *SsoSettingsDto) GetEnableSso() bool`

GetEnableSso returns the EnableSso field if non-nil, zero value otherwise.

### GetEnableSsoOk

`func (o *SsoSettingsDto) GetEnableSsoOk() (*bool, bool)`

GetEnableSsoOk returns a tuple with the EnableSso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSso

`func (o *SsoSettingsDto) SetEnableSso(v bool)`

SetEnableSso sets EnableSso field to given value.

### HasEnableSso

`func (o *SsoSettingsDto) HasEnableSso() bool`

HasEnableSso returns a boolean if a field has been set.

### SetEnableSsoNil

`func (o *SsoSettingsDto) SetEnableSsoNil(b bool)`

 SetEnableSsoNil sets the value for EnableSso to be an explicit nil

### UnsetEnableSso
`func (o *SsoSettingsDto) UnsetEnableSso()`

UnsetEnableSso ensures that no value is present for EnableSso, not even an explicit nil
### GetIdpSettings

`func (o *SsoSettingsDto) GetIdpSettings() SsoIdpSettingsDto`

GetIdpSettings returns the IdpSettings field if non-nil, zero value otherwise.

### GetIdpSettingsOk

`func (o *SsoSettingsDto) GetIdpSettingsOk() (*SsoIdpSettingsDto, bool)`

GetIdpSettingsOk returns a tuple with the IdpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSettings

`func (o *SsoSettingsDto) SetIdpSettings(v SsoIdpSettingsDto)`

SetIdpSettings sets IdpSettings field to given value.

### HasIdpSettings

`func (o *SsoSettingsDto) HasIdpSettings() bool`

HasIdpSettings returns a boolean if a field has been set.

### GetIdpCertificates

`func (o *SsoSettingsDto) GetIdpCertificates() []SsoCertificateDto`

GetIdpCertificates returns the IdpCertificates field if non-nil, zero value otherwise.

### GetIdpCertificatesOk

`func (o *SsoSettingsDto) GetIdpCertificatesOk() (*[]SsoCertificateDto, bool)`

GetIdpCertificatesOk returns a tuple with the IdpCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpCertificates

`func (o *SsoSettingsDto) SetIdpCertificates(v []SsoCertificateDto)`

SetIdpCertificates sets IdpCertificates field to given value.

### HasIdpCertificates

`func (o *SsoSettingsDto) HasIdpCertificates() bool`

HasIdpCertificates returns a boolean if a field has been set.

### SetIdpCertificatesNil

`func (o *SsoSettingsDto) SetIdpCertificatesNil(b bool)`

 SetIdpCertificatesNil sets the value for IdpCertificates to be an explicit nil

### UnsetIdpCertificates
`func (o *SsoSettingsDto) UnsetIdpCertificates()`

UnsetIdpCertificates ensures that no value is present for IdpCertificates, not even an explicit nil
### GetIdpCertificateAdvanced

`func (o *SsoSettingsDto) GetIdpCertificateAdvanced() SsoIdpCertificateAdvancedDto`

GetIdpCertificateAdvanced returns the IdpCertificateAdvanced field if non-nil, zero value otherwise.

### GetIdpCertificateAdvancedOk

`func (o *SsoSettingsDto) GetIdpCertificateAdvancedOk() (*SsoIdpCertificateAdvancedDto, bool)`

GetIdpCertificateAdvancedOk returns a tuple with the IdpCertificateAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpCertificateAdvanced

`func (o *SsoSettingsDto) SetIdpCertificateAdvanced(v SsoIdpCertificateAdvancedDto)`

SetIdpCertificateAdvanced sets IdpCertificateAdvanced field to given value.

### HasIdpCertificateAdvanced

`func (o *SsoSettingsDto) HasIdpCertificateAdvanced() bool`

HasIdpCertificateAdvanced returns a boolean if a field has been set.

### GetSpLoginLabel

`func (o *SsoSettingsDto) GetSpLoginLabel() string`

GetSpLoginLabel returns the SpLoginLabel field if non-nil, zero value otherwise.

### GetSpLoginLabelOk

`func (o *SsoSettingsDto) GetSpLoginLabelOk() (*string, bool)`

GetSpLoginLabelOk returns a tuple with the SpLoginLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpLoginLabel

`func (o *SsoSettingsDto) SetSpLoginLabel(v string)`

SetSpLoginLabel sets SpLoginLabel field to given value.

### HasSpLoginLabel

`func (o *SsoSettingsDto) HasSpLoginLabel() bool`

HasSpLoginLabel returns a boolean if a field has been set.

### SetSpLoginLabelNil

`func (o *SsoSettingsDto) SetSpLoginLabelNil(b bool)`

 SetSpLoginLabelNil sets the value for SpLoginLabel to be an explicit nil

### UnsetSpLoginLabel
`func (o *SsoSettingsDto) UnsetSpLoginLabel()`

UnsetSpLoginLabel ensures that no value is present for SpLoginLabel, not even an explicit nil
### GetSpCertificates

`func (o *SsoSettingsDto) GetSpCertificates() []SsoCertificateDto`

GetSpCertificates returns the SpCertificates field if non-nil, zero value otherwise.

### GetSpCertificatesOk

`func (o *SsoSettingsDto) GetSpCertificatesOk() (*[]SsoCertificateDto, bool)`

GetSpCertificatesOk returns a tuple with the SpCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpCertificates

`func (o *SsoSettingsDto) SetSpCertificates(v []SsoCertificateDto)`

SetSpCertificates sets SpCertificates field to given value.

### HasSpCertificates

`func (o *SsoSettingsDto) HasSpCertificates() bool`

HasSpCertificates returns a boolean if a field has been set.

### SetSpCertificatesNil

`func (o *SsoSettingsDto) SetSpCertificatesNil(b bool)`

 SetSpCertificatesNil sets the value for SpCertificates to be an explicit nil

### UnsetSpCertificates
`func (o *SsoSettingsDto) UnsetSpCertificates()`

UnsetSpCertificates ensures that no value is present for SpCertificates, not even an explicit nil
### GetSpCertificateAdvanced

`func (o *SsoSettingsDto) GetSpCertificateAdvanced() SsoSpCertificateAdvancedDto`

GetSpCertificateAdvanced returns the SpCertificateAdvanced field if non-nil, zero value otherwise.

### GetSpCertificateAdvancedOk

`func (o *SsoSettingsDto) GetSpCertificateAdvancedOk() (*SsoSpCertificateAdvancedDto, bool)`

GetSpCertificateAdvancedOk returns a tuple with the SpCertificateAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpCertificateAdvanced

`func (o *SsoSettingsDto) SetSpCertificateAdvanced(v SsoSpCertificateAdvancedDto)`

SetSpCertificateAdvanced sets SpCertificateAdvanced field to given value.

### HasSpCertificateAdvanced

`func (o *SsoSettingsDto) HasSpCertificateAdvanced() bool`

HasSpCertificateAdvanced returns a boolean if a field has been set.

### GetFieldMapping

`func (o *SsoSettingsDto) GetFieldMapping() SsoFieldMappingDto`

GetFieldMapping returns the FieldMapping field if non-nil, zero value otherwise.

### GetFieldMappingOk

`func (o *SsoSettingsDto) GetFieldMappingOk() (*SsoFieldMappingDto, bool)`

GetFieldMappingOk returns a tuple with the FieldMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMapping

`func (o *SsoSettingsDto) SetFieldMapping(v SsoFieldMappingDto)`

SetFieldMapping sets FieldMapping field to given value.

### HasFieldMapping

`func (o *SsoSettingsDto) HasFieldMapping() bool`

HasFieldMapping returns a boolean if a field has been set.

### GetHideAuthPage

`func (o *SsoSettingsDto) GetHideAuthPage() bool`

GetHideAuthPage returns the HideAuthPage field if non-nil, zero value otherwise.

### GetHideAuthPageOk

`func (o *SsoSettingsDto) GetHideAuthPageOk() (*bool, bool)`

GetHideAuthPageOk returns a tuple with the HideAuthPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAuthPage

`func (o *SsoSettingsDto) SetHideAuthPage(v bool)`

SetHideAuthPage sets HideAuthPage field to given value.

### HasHideAuthPage

`func (o *SsoSettingsDto) HasHideAuthPage() bool`

HasHideAuthPage returns a boolean if a field has been set.

### GetUsersType

`func (o *SsoSettingsDto) GetUsersType() int32`

GetUsersType returns the UsersType field if non-nil, zero value otherwise.

### GetUsersTypeOk

`func (o *SsoSettingsDto) GetUsersTypeOk() (*int32, bool)`

GetUsersTypeOk returns a tuple with the UsersType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersType

`func (o *SsoSettingsDto) SetUsersType(v int32)`

SetUsersType sets UsersType field to given value.

### HasUsersType

`func (o *SsoSettingsDto) HasUsersType() bool`

HasUsersType returns a boolean if a field has been set.

### GetDisableEmailVerification

`func (o *SsoSettingsDto) GetDisableEmailVerification() bool`

GetDisableEmailVerification returns the DisableEmailVerification field if non-nil, zero value otherwise.

### GetDisableEmailVerificationOk

`func (o *SsoSettingsDto) GetDisableEmailVerificationOk() (*bool, bool)`

GetDisableEmailVerificationOk returns a tuple with the DisableEmailVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEmailVerification

`func (o *SsoSettingsDto) SetDisableEmailVerification(v bool)`

SetDisableEmailVerification sets DisableEmailVerification field to given value.

### HasDisableEmailVerification

`func (o *SsoSettingsDto) HasDisableEmailVerification() bool`

HasDisableEmailVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


