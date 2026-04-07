# SsoSettingsV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastModified** | Pointer to **time.Time** | The timestamp indicating when the settings were last modified. | [optional] 
**EnableSso** | Pointer to **NullableBool** | Specifies if the SSO settings are enabled or not. | [optional] 
**IdpSettings** | Pointer to [**SsoIdpSettings**](SsoIdpSettings.md) |  | [optional] 
**IdpCertificates** | Pointer to [**[]SsoCertificate**](SsoCertificate.md) | The list of the IdP certificates. | [optional] 
**IdpCertificateAdvanced** | Pointer to [**SsoIdpCertificateAdvanced**](SsoIdpCertificateAdvanced.md) |  | [optional] 
**SpLoginLabel** | Pointer to **NullableString** | The SP login label. | [optional] 
**SpCertificates** | Pointer to [**[]SsoCertificate**](SsoCertificate.md) | The list of the SP certificates. | [optional] 
**SpCertificateAdvanced** | Pointer to [**SsoSpCertificateAdvanced**](SsoSpCertificateAdvanced.md) |  | [optional] 
**FieldMapping** | Pointer to [**SsoFieldMapping**](SsoFieldMapping.md) |  | [optional] 
**HideAuthPage** | Pointer to **bool** | Specifies if the authentication page will be hidden or not. | [optional] 
**UsersType** | Pointer to **int32** | The user type. | [optional] 
**DisableEmailVerification** | Pointer to **bool** | Specifies if the email verification is disabled or not. | [optional] 

## Methods

### NewSsoSettingsV2

`func NewSsoSettingsV2() *SsoSettingsV2`

NewSsoSettingsV2 instantiates a new SsoSettingsV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoSettingsV2WithDefaults

`func NewSsoSettingsV2WithDefaults() *SsoSettingsV2`

NewSsoSettingsV2WithDefaults instantiates a new SsoSettingsV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastModified

`func (o *SsoSettingsV2) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *SsoSettingsV2) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *SsoSettingsV2) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *SsoSettingsV2) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetEnableSso

`func (o *SsoSettingsV2) GetEnableSso() bool`

GetEnableSso returns the EnableSso field if non-nil, zero value otherwise.

### GetEnableSsoOk

`func (o *SsoSettingsV2) GetEnableSsoOk() (*bool, bool)`

GetEnableSsoOk returns a tuple with the EnableSso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSso

`func (o *SsoSettingsV2) SetEnableSso(v bool)`

SetEnableSso sets EnableSso field to given value.

### HasEnableSso

`func (o *SsoSettingsV2) HasEnableSso() bool`

HasEnableSso returns a boolean if a field has been set.

### SetEnableSsoNil

`func (o *SsoSettingsV2) SetEnableSsoNil(b bool)`

 SetEnableSsoNil sets the value for EnableSso to be an explicit nil

### UnsetEnableSso
`func (o *SsoSettingsV2) UnsetEnableSso()`

UnsetEnableSso ensures that no value is present for EnableSso, not even an explicit nil
### GetIdpSettings

`func (o *SsoSettingsV2) GetIdpSettings() SsoIdpSettings`

GetIdpSettings returns the IdpSettings field if non-nil, zero value otherwise.

### GetIdpSettingsOk

`func (o *SsoSettingsV2) GetIdpSettingsOk() (*SsoIdpSettings, bool)`

GetIdpSettingsOk returns a tuple with the IdpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSettings

`func (o *SsoSettingsV2) SetIdpSettings(v SsoIdpSettings)`

SetIdpSettings sets IdpSettings field to given value.

### HasIdpSettings

`func (o *SsoSettingsV2) HasIdpSettings() bool`

HasIdpSettings returns a boolean if a field has been set.

### GetIdpCertificates

`func (o *SsoSettingsV2) GetIdpCertificates() []SsoCertificate`

GetIdpCertificates returns the IdpCertificates field if non-nil, zero value otherwise.

### GetIdpCertificatesOk

`func (o *SsoSettingsV2) GetIdpCertificatesOk() (*[]SsoCertificate, bool)`

GetIdpCertificatesOk returns a tuple with the IdpCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpCertificates

`func (o *SsoSettingsV2) SetIdpCertificates(v []SsoCertificate)`

SetIdpCertificates sets IdpCertificates field to given value.

### HasIdpCertificates

`func (o *SsoSettingsV2) HasIdpCertificates() bool`

HasIdpCertificates returns a boolean if a field has been set.

### SetIdpCertificatesNil

`func (o *SsoSettingsV2) SetIdpCertificatesNil(b bool)`

 SetIdpCertificatesNil sets the value for IdpCertificates to be an explicit nil

### UnsetIdpCertificates
`func (o *SsoSettingsV2) UnsetIdpCertificates()`

UnsetIdpCertificates ensures that no value is present for IdpCertificates, not even an explicit nil
### GetIdpCertificateAdvanced

`func (o *SsoSettingsV2) GetIdpCertificateAdvanced() SsoIdpCertificateAdvanced`

GetIdpCertificateAdvanced returns the IdpCertificateAdvanced field if non-nil, zero value otherwise.

### GetIdpCertificateAdvancedOk

`func (o *SsoSettingsV2) GetIdpCertificateAdvancedOk() (*SsoIdpCertificateAdvanced, bool)`

GetIdpCertificateAdvancedOk returns a tuple with the IdpCertificateAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpCertificateAdvanced

`func (o *SsoSettingsV2) SetIdpCertificateAdvanced(v SsoIdpCertificateAdvanced)`

SetIdpCertificateAdvanced sets IdpCertificateAdvanced field to given value.

### HasIdpCertificateAdvanced

`func (o *SsoSettingsV2) HasIdpCertificateAdvanced() bool`

HasIdpCertificateAdvanced returns a boolean if a field has been set.

### GetSpLoginLabel

`func (o *SsoSettingsV2) GetSpLoginLabel() string`

GetSpLoginLabel returns the SpLoginLabel field if non-nil, zero value otherwise.

### GetSpLoginLabelOk

`func (o *SsoSettingsV2) GetSpLoginLabelOk() (*string, bool)`

GetSpLoginLabelOk returns a tuple with the SpLoginLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpLoginLabel

`func (o *SsoSettingsV2) SetSpLoginLabel(v string)`

SetSpLoginLabel sets SpLoginLabel field to given value.

### HasSpLoginLabel

`func (o *SsoSettingsV2) HasSpLoginLabel() bool`

HasSpLoginLabel returns a boolean if a field has been set.

### SetSpLoginLabelNil

`func (o *SsoSettingsV2) SetSpLoginLabelNil(b bool)`

 SetSpLoginLabelNil sets the value for SpLoginLabel to be an explicit nil

### UnsetSpLoginLabel
`func (o *SsoSettingsV2) UnsetSpLoginLabel()`

UnsetSpLoginLabel ensures that no value is present for SpLoginLabel, not even an explicit nil
### GetSpCertificates

`func (o *SsoSettingsV2) GetSpCertificates() []SsoCertificate`

GetSpCertificates returns the SpCertificates field if non-nil, zero value otherwise.

### GetSpCertificatesOk

`func (o *SsoSettingsV2) GetSpCertificatesOk() (*[]SsoCertificate, bool)`

GetSpCertificatesOk returns a tuple with the SpCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpCertificates

`func (o *SsoSettingsV2) SetSpCertificates(v []SsoCertificate)`

SetSpCertificates sets SpCertificates field to given value.

### HasSpCertificates

`func (o *SsoSettingsV2) HasSpCertificates() bool`

HasSpCertificates returns a boolean if a field has been set.

### SetSpCertificatesNil

`func (o *SsoSettingsV2) SetSpCertificatesNil(b bool)`

 SetSpCertificatesNil sets the value for SpCertificates to be an explicit nil

### UnsetSpCertificates
`func (o *SsoSettingsV2) UnsetSpCertificates()`

UnsetSpCertificates ensures that no value is present for SpCertificates, not even an explicit nil
### GetSpCertificateAdvanced

`func (o *SsoSettingsV2) GetSpCertificateAdvanced() SsoSpCertificateAdvanced`

GetSpCertificateAdvanced returns the SpCertificateAdvanced field if non-nil, zero value otherwise.

### GetSpCertificateAdvancedOk

`func (o *SsoSettingsV2) GetSpCertificateAdvancedOk() (*SsoSpCertificateAdvanced, bool)`

GetSpCertificateAdvancedOk returns a tuple with the SpCertificateAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpCertificateAdvanced

`func (o *SsoSettingsV2) SetSpCertificateAdvanced(v SsoSpCertificateAdvanced)`

SetSpCertificateAdvanced sets SpCertificateAdvanced field to given value.

### HasSpCertificateAdvanced

`func (o *SsoSettingsV2) HasSpCertificateAdvanced() bool`

HasSpCertificateAdvanced returns a boolean if a field has been set.

### GetFieldMapping

`func (o *SsoSettingsV2) GetFieldMapping() SsoFieldMapping`

GetFieldMapping returns the FieldMapping field if non-nil, zero value otherwise.

### GetFieldMappingOk

`func (o *SsoSettingsV2) GetFieldMappingOk() (*SsoFieldMapping, bool)`

GetFieldMappingOk returns a tuple with the FieldMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMapping

`func (o *SsoSettingsV2) SetFieldMapping(v SsoFieldMapping)`

SetFieldMapping sets FieldMapping field to given value.

### HasFieldMapping

`func (o *SsoSettingsV2) HasFieldMapping() bool`

HasFieldMapping returns a boolean if a field has been set.

### GetHideAuthPage

`func (o *SsoSettingsV2) GetHideAuthPage() bool`

GetHideAuthPage returns the HideAuthPage field if non-nil, zero value otherwise.

### GetHideAuthPageOk

`func (o *SsoSettingsV2) GetHideAuthPageOk() (*bool, bool)`

GetHideAuthPageOk returns a tuple with the HideAuthPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAuthPage

`func (o *SsoSettingsV2) SetHideAuthPage(v bool)`

SetHideAuthPage sets HideAuthPage field to given value.

### HasHideAuthPage

`func (o *SsoSettingsV2) HasHideAuthPage() bool`

HasHideAuthPage returns a boolean if a field has been set.

### GetUsersType

`func (o *SsoSettingsV2) GetUsersType() int32`

GetUsersType returns the UsersType field if non-nil, zero value otherwise.

### GetUsersTypeOk

`func (o *SsoSettingsV2) GetUsersTypeOk() (*int32, bool)`

GetUsersTypeOk returns a tuple with the UsersType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersType

`func (o *SsoSettingsV2) SetUsersType(v int32)`

SetUsersType sets UsersType field to given value.

### HasUsersType

`func (o *SsoSettingsV2) HasUsersType() bool`

HasUsersType returns a boolean if a field has been set.

### GetDisableEmailVerification

`func (o *SsoSettingsV2) GetDisableEmailVerification() bool`

GetDisableEmailVerification returns the DisableEmailVerification field if non-nil, zero value otherwise.

### GetDisableEmailVerificationOk

`func (o *SsoSettingsV2) GetDisableEmailVerificationOk() (*bool, bool)`

GetDisableEmailVerificationOk returns a tuple with the DisableEmailVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEmailVerification

`func (o *SsoSettingsV2) SetDisableEmailVerification(v bool)`

SetDisableEmailVerification sets DisableEmailVerification field to given value.

### HasDisableEmailVerification

`func (o *SsoSettingsV2) HasDisableEmailVerification() bool`

HasDisableEmailVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


