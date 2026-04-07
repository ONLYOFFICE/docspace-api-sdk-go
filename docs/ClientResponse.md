# ClientResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The client name. | [optional] 
**Description** | Pointer to **string** | The client description. | [optional] 
**Tenant** | Pointer to **int64** | The tenant ID associated with the client. | [optional] 
**Scopes** | Pointer to **[]string** | The client scopes. | [optional] 
**Enabled** | Pointer to **bool** | Specifies if the client is currently enabled or not. | [optional] 
**ClientId** | Pointer to **string** | The client identifier issued to the client during registration. | [optional] 
**ClientSecret** | Pointer to **string** | The client secret issued to the client during registration. | [optional] 
**WebsiteUrl** | Pointer to **string** | The URL to the client's website. | [optional] 
**TermsUrl** | Pointer to **string** | The URL to the client's terms of service. | [optional] 
**PolicyUrl** | Pointer to **string** | The URL to the client's privacy policy. | [optional] 
**Logo** | Pointer to **string** | The URL to the client's logo. | [optional] 
**AuthenticationMethods** | Pointer to **[]string** | The authentication methods supported by the client. | [optional] 
**RedirectUris** | Pointer to **[]string** | The list of allowed redirect URIs. | [optional] 
**AllowedOrigins** | Pointer to **[]string** | The list of allowed CORS origins. | [optional] 
**LogoutRedirectUris** | Pointer to **[]string** | The list of allowed logout redirect URIs. | [optional] 
**CreatedOn** | Pointer to **time.Time** | The date and time when the client was created. | [optional] 
**CreatedBy** | Pointer to **string** | The user who created the client. | [optional] 
**ModifiedOn** | Pointer to **time.Time** | The date and time when the client was last modified. | [optional] 
**ModifiedBy** | Pointer to **string** | The user who last modified the client. | [optional] 
**IsPublic** | Pointer to **bool** | Indicates whether the client is accessible by third-party tenants. | [optional] 

## Methods

### NewClientResponse

`func NewClientResponse() *ClientResponse`

NewClientResponse instantiates a new ClientResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientResponseWithDefaults

`func NewClientResponseWithDefaults() *ClientResponse`

NewClientResponseWithDefaults instantiates a new ClientResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClientResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ClientResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClientResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClientResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClientResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTenant

`func (o *ClientResponse) GetTenant() int64`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ClientResponse) GetTenantOk() (*int64, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ClientResponse) SetTenant(v int64)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ClientResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetScopes

`func (o *ClientResponse) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ClientResponse) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ClientResponse) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ClientResponse) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetEnabled

`func (o *ClientResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClientResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClientResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ClientResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetClientId

`func (o *ClientResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClientResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClientResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ClientResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *ClientResponse) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ClientResponse) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ClientResponse) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ClientResponse) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *ClientResponse) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *ClientResponse) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *ClientResponse) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *ClientResponse) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetTermsUrl

`func (o *ClientResponse) GetTermsUrl() string`

GetTermsUrl returns the TermsUrl field if non-nil, zero value otherwise.

### GetTermsUrlOk

`func (o *ClientResponse) GetTermsUrlOk() (*string, bool)`

GetTermsUrlOk returns a tuple with the TermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsUrl

`func (o *ClientResponse) SetTermsUrl(v string)`

SetTermsUrl sets TermsUrl field to given value.

### HasTermsUrl

`func (o *ClientResponse) HasTermsUrl() bool`

HasTermsUrl returns a boolean if a field has been set.

### GetPolicyUrl

`func (o *ClientResponse) GetPolicyUrl() string`

GetPolicyUrl returns the PolicyUrl field if non-nil, zero value otherwise.

### GetPolicyUrlOk

`func (o *ClientResponse) GetPolicyUrlOk() (*string, bool)`

GetPolicyUrlOk returns a tuple with the PolicyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUrl

`func (o *ClientResponse) SetPolicyUrl(v string)`

SetPolicyUrl sets PolicyUrl field to given value.

### HasPolicyUrl

`func (o *ClientResponse) HasPolicyUrl() bool`

HasPolicyUrl returns a boolean if a field has been set.

### GetLogo

`func (o *ClientResponse) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ClientResponse) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ClientResponse) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ClientResponse) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetAuthenticationMethods

`func (o *ClientResponse) GetAuthenticationMethods() []string`

GetAuthenticationMethods returns the AuthenticationMethods field if non-nil, zero value otherwise.

### GetAuthenticationMethodsOk

`func (o *ClientResponse) GetAuthenticationMethodsOk() (*[]string, bool)`

GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethods

`func (o *ClientResponse) SetAuthenticationMethods(v []string)`

SetAuthenticationMethods sets AuthenticationMethods field to given value.

### HasAuthenticationMethods

`func (o *ClientResponse) HasAuthenticationMethods() bool`

HasAuthenticationMethods returns a boolean if a field has been set.

### GetRedirectUris

`func (o *ClientResponse) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *ClientResponse) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *ClientResponse) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *ClientResponse) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetAllowedOrigins

`func (o *ClientResponse) GetAllowedOrigins() []string`

GetAllowedOrigins returns the AllowedOrigins field if non-nil, zero value otherwise.

### GetAllowedOriginsOk

`func (o *ClientResponse) GetAllowedOriginsOk() (*[]string, bool)`

GetAllowedOriginsOk returns a tuple with the AllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrigins

`func (o *ClientResponse) SetAllowedOrigins(v []string)`

SetAllowedOrigins sets AllowedOrigins field to given value.

### HasAllowedOrigins

`func (o *ClientResponse) HasAllowedOrigins() bool`

HasAllowedOrigins returns a boolean if a field has been set.

### GetLogoutRedirectUris

`func (o *ClientResponse) GetLogoutRedirectUris() []string`

GetLogoutRedirectUris returns the LogoutRedirectUris field if non-nil, zero value otherwise.

### GetLogoutRedirectUrisOk

`func (o *ClientResponse) GetLogoutRedirectUrisOk() (*[]string, bool)`

GetLogoutRedirectUrisOk returns a tuple with the LogoutRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutRedirectUris

`func (o *ClientResponse) SetLogoutRedirectUris(v []string)`

SetLogoutRedirectUris sets LogoutRedirectUris field to given value.

### HasLogoutRedirectUris

`func (o *ClientResponse) HasLogoutRedirectUris() bool`

HasLogoutRedirectUris returns a boolean if a field has been set.

### GetCreatedOn

`func (o *ClientResponse) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *ClientResponse) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *ClientResponse) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *ClientResponse) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ClientResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ClientResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ClientResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ClientResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetModifiedOn

`func (o *ClientResponse) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *ClientResponse) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *ClientResponse) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *ClientResponse) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetModifiedBy

`func (o *ClientResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *ClientResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *ClientResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *ClientResponse) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetIsPublic

`func (o *ClientResponse) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *ClientResponse) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *ClientResponse) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *ClientResponse) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


