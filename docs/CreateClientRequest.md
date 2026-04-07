# CreateClientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The client name. | [optional] 
**Description** | Pointer to **string** | The description of the client | [optional] 
**Logo** | Pointer to **string** | The logo of the client in base64 format | [optional] 
**Scopes** | Pointer to **[]string** | The scopes for the client | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**AllowPkce** | Pointer to **bool** | Indicates whether PKCE is allowed for the client | [optional] 
**IsPublic** | Pointer to **bool** | Indicates if the client is public | [optional] 
**WebsiteUrl** | Pointer to **string** | The website URL of the client | [optional] 
**TermsUrl** | Pointer to **string** | The terms URL of the client | [optional] 
**PolicyUrl** | Pointer to **string** | The policy URL of the client | [optional] 
**RedirectUris** | **[]string** | The redirect URIs for the client | 
**AllowedOrigins** | **[]string** | The allowed origins for the client | 
**LogoutRedirectUri** | Pointer to **string** | The logout redirect URI for the client | [optional] 

## Methods

### NewCreateClientRequest

`func NewCreateClientRequest(redirectUris []string, allowedOrigins []string, ) *CreateClientRequest`

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

### HasName

`func (o *CreateClientRequest) HasName() bool`

HasName returns a boolean if a field has been set.

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

### HasLogo

`func (o *CreateClientRequest) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

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

### HasScopes

`func (o *CreateClientRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetPublic

`func (o *CreateClientRequest) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *CreateClientRequest) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *CreateClientRequest) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *CreateClientRequest) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

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

### HasWebsiteUrl

`func (o *CreateClientRequest) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

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

### HasTermsUrl

`func (o *CreateClientRequest) HasTermsUrl() bool`

HasTermsUrl returns a boolean if a field has been set.

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

### HasPolicyUrl

`func (o *CreateClientRequest) HasPolicyUrl() bool`

HasPolicyUrl returns a boolean if a field has been set.

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

### HasLogoutRedirectUri

`func (o *CreateClientRequest) HasLogoutRedirectUri() bool`

HasLogoutRedirectUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


