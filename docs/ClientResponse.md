# ClientResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The display name shown to the user on the consent screen, between 3 and 256 characters. | [optional] 
**Description** | Pointer to **string** | The free-text description shown next to the name on the consent screen, at most 255 characters. | [optional] 
**Tenant** | Pointer to **int64** | The identifier of the portal the client belongs to. A client is visible only inside its own tenant, apart from the unauthenticated public info read. | [optional] 
**Scopes** | Pointer to **[]string** | The permissions the client may ask for, named as they appear in the tenant scope catalogue - for example files:read, rooms:write or openid. A client cannot request a scope that is not listed here. | [optional] 
**Enabled** | Pointer to **bool** | Whether the client may currently obtain tokens. A disabled client keeps its registration and the tokens already issued to it, but new authorization requests for it are refused. | [optional] 
**ClientId** | Pointer to **string** | The generated identifier of the client, sent as client_id in every OAuth2 request. It is assigned when the client is registered and never changes afterwards. | [optional] 
**ClientSecret** | Pointer to **string** | The client secret, which the client presents at the token endpoint when it authenticates with client_secret_post. It is omitted from the response rather than sent as null when the client has none. | [optional] 
**WebsiteUrl** | Pointer to **string** | The URL of the client home page, offered to the user before they consent. | [optional] 
**TermsUrl** | Pointer to **string** | The URL of the client terms of service, linked from the consent screen. | [optional] 
**PolicyUrl** | Pointer to **string** | The URL of the client privacy policy, linked from the consent screen. | [optional] 
**Logo** | Pointer to **string** | The client logo as a data URI carrying base64 image data, shown on the consent screen. Only png, jpeg, jpg and svg+xml are accepted, the whole string may not exceed 2000000 characters and the decoded image may not exceed 256000 bytes. | [optional] 
**AuthenticationMethods** | Pointer to **[]string** | How the client authenticates itself at the token endpoint: client_secret_post for a confidential client that sends its secret, none for a public client that proves itself with PKCE instead. | [optional] 
**RedirectUris** | Pointer to **[]string** | The URIs an authorization code may be delivered to. An authorization request naming any other URI is refused, and the set holds between 1 and 12 addresses. | [optional] 
**AllowedOrigins** | Pointer to **[]string** | The web origins allowed to call the portal on behalf of this client, used for the CORS check. The set holds between 1 and 12 addresses. | [optional] 
**LogoutRedirectUris** | Pointer to **[]string** | The URIs the user may be sent back to once they have logged out. | [optional] 
**CreatedOn** | Pointer to **time.Time** | When the client was registered, as an ISO-8601 timestamp with a zone offset. | [optional] 
**CreatedBy** | Pointer to **string** | The identifier of the user who registered the client. A plain user may read and change only the clients where this is their own identifier. | [optional] 
**ModifiedOn** | Pointer to **time.Time** | When the client was last changed, as an ISO-8601 timestamp with a zone offset. | [optional] 
**ModifiedBy** | Pointer to **string** | The identifier of the user who last changed the client. | [optional] 
**IsPublic** | Pointer to **bool** | Whether the client is offered to third-party tenants rather than only to the tenant that registered it. | [optional] 

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


