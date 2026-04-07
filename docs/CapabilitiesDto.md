# CapabilitiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdapEnabled** | **bool** | Specifies if the LDAP settings are enabled or not. | 
**LdapDomain** | Pointer to **NullableString** | The LDAP domain. | [optional] 
**Providers** | **[]string** | The list of providers. | 
**SsoLabel** | **NullableString** | The SP login label. | 
**OauthEnabled** | **bool** | Specifies if OAuth is enabled or not. | 
**SsoUrl** | **NullableString** | The SSO URL. If this parameter is empty, then the SSO settings are disabled. | 
**IdentityServerEnabled** | **bool** | Specifies if an identity server is enabled or not. | 

## Methods

### NewCapabilitiesDto

`func NewCapabilitiesDto(ldapEnabled bool, providers []string, ssoLabel NullableString, oauthEnabled bool, ssoUrl NullableString, identityServerEnabled bool, ) *CapabilitiesDto`

NewCapabilitiesDto instantiates a new CapabilitiesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitiesDtoWithDefaults

`func NewCapabilitiesDtoWithDefaults() *CapabilitiesDto`

NewCapabilitiesDtoWithDefaults instantiates a new CapabilitiesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdapEnabled

`func (o *CapabilitiesDto) GetLdapEnabled() bool`

GetLdapEnabled returns the LdapEnabled field if non-nil, zero value otherwise.

### GetLdapEnabledOk

`func (o *CapabilitiesDto) GetLdapEnabledOk() (*bool, bool)`

GetLdapEnabledOk returns a tuple with the LdapEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapEnabled

`func (o *CapabilitiesDto) SetLdapEnabled(v bool)`

SetLdapEnabled sets LdapEnabled field to given value.


### GetLdapDomain

`func (o *CapabilitiesDto) GetLdapDomain() string`

GetLdapDomain returns the LdapDomain field if non-nil, zero value otherwise.

### GetLdapDomainOk

`func (o *CapabilitiesDto) GetLdapDomainOk() (*string, bool)`

GetLdapDomainOk returns a tuple with the LdapDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapDomain

`func (o *CapabilitiesDto) SetLdapDomain(v string)`

SetLdapDomain sets LdapDomain field to given value.

### HasLdapDomain

`func (o *CapabilitiesDto) HasLdapDomain() bool`

HasLdapDomain returns a boolean if a field has been set.

### SetLdapDomainNil

`func (o *CapabilitiesDto) SetLdapDomainNil(b bool)`

 SetLdapDomainNil sets the value for LdapDomain to be an explicit nil

### UnsetLdapDomain
`func (o *CapabilitiesDto) UnsetLdapDomain()`

UnsetLdapDomain ensures that no value is present for LdapDomain, not even an explicit nil
### GetProviders

`func (o *CapabilitiesDto) GetProviders() []string`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *CapabilitiesDto) GetProvidersOk() (*[]string, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *CapabilitiesDto) SetProviders(v []string)`

SetProviders sets Providers field to given value.


### SetProvidersNil

`func (o *CapabilitiesDto) SetProvidersNil(b bool)`

 SetProvidersNil sets the value for Providers to be an explicit nil

### UnsetProviders
`func (o *CapabilitiesDto) UnsetProviders()`

UnsetProviders ensures that no value is present for Providers, not even an explicit nil
### GetSsoLabel

`func (o *CapabilitiesDto) GetSsoLabel() string`

GetSsoLabel returns the SsoLabel field if non-nil, zero value otherwise.

### GetSsoLabelOk

`func (o *CapabilitiesDto) GetSsoLabelOk() (*string, bool)`

GetSsoLabelOk returns a tuple with the SsoLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoLabel

`func (o *CapabilitiesDto) SetSsoLabel(v string)`

SetSsoLabel sets SsoLabel field to given value.


### SetSsoLabelNil

`func (o *CapabilitiesDto) SetSsoLabelNil(b bool)`

 SetSsoLabelNil sets the value for SsoLabel to be an explicit nil

### UnsetSsoLabel
`func (o *CapabilitiesDto) UnsetSsoLabel()`

UnsetSsoLabel ensures that no value is present for SsoLabel, not even an explicit nil
### GetOauthEnabled

`func (o *CapabilitiesDto) GetOauthEnabled() bool`

GetOauthEnabled returns the OauthEnabled field if non-nil, zero value otherwise.

### GetOauthEnabledOk

`func (o *CapabilitiesDto) GetOauthEnabledOk() (*bool, bool)`

GetOauthEnabledOk returns a tuple with the OauthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthEnabled

`func (o *CapabilitiesDto) SetOauthEnabled(v bool)`

SetOauthEnabled sets OauthEnabled field to given value.


### GetSsoUrl

`func (o *CapabilitiesDto) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *CapabilitiesDto) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *CapabilitiesDto) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.


### SetSsoUrlNil

`func (o *CapabilitiesDto) SetSsoUrlNil(b bool)`

 SetSsoUrlNil sets the value for SsoUrl to be an explicit nil

### UnsetSsoUrl
`func (o *CapabilitiesDto) UnsetSsoUrl()`

UnsetSsoUrl ensures that no value is present for SsoUrl, not even an explicit nil
### GetIdentityServerEnabled

`func (o *CapabilitiesDto) GetIdentityServerEnabled() bool`

GetIdentityServerEnabled returns the IdentityServerEnabled field if non-nil, zero value otherwise.

### GetIdentityServerEnabledOk

`func (o *CapabilitiesDto) GetIdentityServerEnabledOk() (*bool, bool)`

GetIdentityServerEnabledOk returns a tuple with the IdentityServerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityServerEnabled

`func (o *CapabilitiesDto) SetIdentityServerEnabled(v bool)`

SetIdentityServerEnabled sets IdentityServerEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


