# CreateClientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The display name shown to the user on the consent screen. It has to be between 3 and 256 characters long. | 
**Description** | Pointer to **string** | The free-text description shown next to the name on the consent screen, at most 255 characters. | [optional] 
**Logo** | **string** | The client logo as a data URI carrying base64 image data, shown on the consent screen. Only png, jpeg, jpg and svg+xml are accepted, the whole string may not exceed 2000000 characters and the decoded image may not exceed 256000 bytes. | 
**Scopes** | **[]string** | The permissions the client may ask for, named as they appear in the tenant scope catalogue - for example files:read, rooms:write or openid. A client cannot request a scope that is not listed here. | 
**AllowPkce** | Pointer to **bool** | Whether the client may use PKCE. Turning it on lets the client authenticate with the none method and prove itself with a code verifier instead of sending a secret, which is what a client that cannot keep a secret needs. | [optional] 
**WebsiteUrl** | **string** | The URL of the client home page, offered to the user before they consent. The value has to be an http or https URL. | 
**TermsUrl** | **string** | The URL of the client terms of service, linked from the consent screen. The value has to be an http or https URL. | 
**PolicyUrl** | **string** | The URL of the client privacy policy, linked from the consent screen. The value has to be an http or https URL. | 
**RedirectUris** | **[]string** | The URIs an authorization code may be delivered to. An authorization request naming any other URI is refused, and the set holds between 1 and 12 addresses. | 
**AllowedOrigins** | **[]string** | The web origins allowed to call the portal on behalf of this client, used for the CORS check. The set holds between 1 and 12 addresses. | 
**LogoutRedirectUri** | **string** | The single URI the user may be sent back to once they have logged out. The value has to be an http or https URL. | 
**IsPublic** | Pointer to **bool** | Whether the client is offered to third-party tenants rather than only to the tenant that registers it. | [optional] 

## Methods

### NewCreateClientRequest

`func NewCreateClientRequest(name string, logo string, scopes []string, websiteUrl string, termsUrl string, policyUrl string, redirectUris []string, allowedOrigins []string, logoutRedirectUri string, ) *CreateClientRequest`

NewCreateClientRequest instantiates a new CreateClientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClientRequestWithDefaults

`func NewCreateClientRequestWithDefaults() *CreateClientRequest`

NewCreateClientRequestWithDefaults instantiates a new CreateClientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateClientRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateClientRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateClientRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateClientRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateClientRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateClientRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateClientRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLogo

`func (o *CreateClientRequest) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *CreateClientRequest) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *CreateClientRequest) SetLogo(v string)`

SetLogo sets Logo field to given value.


### GetScopes

`func (o *CreateClientRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CreateClientRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CreateClientRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetAllowPkce

`func (o *CreateClientRequest) GetAllowPkce() bool`

GetAllowPkce returns the AllowPkce field if non-nil, zero value otherwise.

### GetAllowPkceOk

`func (o *CreateClientRequest) GetAllowPkceOk() (*bool, bool)`

GetAllowPkceOk returns a tuple with the AllowPkce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPkce

`func (o *CreateClientRequest) SetAllowPkce(v bool)`

SetAllowPkce sets AllowPkce field to given value.

### HasAllowPkce

`func (o *CreateClientRequest) HasAllowPkce() bool`

HasAllowPkce returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *CreateClientRequest) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *CreateClientRequest) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *CreateClientRequest) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.


### GetTermsUrl

`func (o *CreateClientRequest) GetTermsUrl() string`

GetTermsUrl returns the TermsUrl field if non-nil, zero value otherwise.

### GetTermsUrlOk

`func (o *CreateClientRequest) GetTermsUrlOk() (*string, bool)`

GetTermsUrlOk returns a tuple with the TermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsUrl

`func (o *CreateClientRequest) SetTermsUrl(v string)`

SetTermsUrl sets TermsUrl field to given value.


### GetPolicyUrl

`func (o *CreateClientRequest) GetPolicyUrl() string`

GetPolicyUrl returns the PolicyUrl field if non-nil, zero value otherwise.

### GetPolicyUrlOk

`func (o *CreateClientRequest) GetPolicyUrlOk() (*string, bool)`

GetPolicyUrlOk returns a tuple with the PolicyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUrl

`func (o *CreateClientRequest) SetPolicyUrl(v string)`

SetPolicyUrl sets PolicyUrl field to given value.


### GetRedirectUris

`func (o *CreateClientRequest) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *CreateClientRequest) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *CreateClientRequest) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.


### GetAllowedOrigins

`func (o *CreateClientRequest) GetAllowedOrigins() []string`

GetAllowedOrigins returns the AllowedOrigins field if non-nil, zero value otherwise.

### GetAllowedOriginsOk

`func (o *CreateClientRequest) GetAllowedOriginsOk() (*[]string, bool)`

GetAllowedOriginsOk returns a tuple with the AllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrigins

`func (o *CreateClientRequest) SetAllowedOrigins(v []string)`

SetAllowedOrigins sets AllowedOrigins field to given value.


### GetLogoutRedirectUri

`func (o *CreateClientRequest) GetLogoutRedirectUri() string`

GetLogoutRedirectUri returns the LogoutRedirectUri field if non-nil, zero value otherwise.

### GetLogoutRedirectUriOk

`func (o *CreateClientRequest) GetLogoutRedirectUriOk() (*string, bool)`

GetLogoutRedirectUriOk returns a tuple with the LogoutRedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutRedirectUri

`func (o *CreateClientRequest) SetLogoutRedirectUri(v string)`

SetLogoutRedirectUri sets LogoutRedirectUri field to given value.


### GetIsPublic

`func (o *CreateClientRequest) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *CreateClientRequest) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *CreateClientRequest) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *CreateClientRequest) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


