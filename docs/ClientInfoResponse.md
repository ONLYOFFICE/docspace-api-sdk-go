# ClientInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The client name. | [optional] 
**Description** | Pointer to **string** | The client description. | [optional] 
**Scopes** | Pointer to **[]string** | The client scopes. | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**ClientId** | Pointer to **string** | The client ID. | [optional] 
**WebsiteUrl** | Pointer to **string** | The URL to the client's website | [optional] 
**TermsUrl** | Pointer to **string** | The URL to the client's terms of service. | [optional] 
**PolicyUrl** | Pointer to **string** | The URL to the client's privacy policy. | [optional] 
**Logo** | Pointer to **string** | The client logo in base64 format. | [optional] 
**AuthenticationMethods** | Pointer to **[]string** | The authentication methods supported by the client. | [optional] 
**IsPublic** | Pointer to **bool** | Indicates whether the client is accessible by third-party tenants. | [optional] 
**CreatedOn** | Pointer to **time.Time** | The date and time when the client was created. | [optional] 
**CreatedBy** | Pointer to **string** | The user who created the client. | [optional] 
**ModifiedOn** | Pointer to **time.Time** | The date and time when the client was last modified. | [optional] 
**ModifiedBy** | Pointer to **string** | The user who last modified the client. | [optional] 

## Methods

### NewClientInfoResponse

`func NewClientInfoResponse() *ClientInfoResponse`

NewClientInfoResponse instantiates a new ClientInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientInfoResponseWithDefaults

`func NewClientInfoResponseWithDefaults() *ClientInfoResponse`

NewClientInfoResponseWithDefaults instantiates a new ClientInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClientInfoResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientInfoResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientInfoResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientInfoResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ClientInfoResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClientInfoResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClientInfoResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClientInfoResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScopes

`func (o *ClientInfoResponse) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ClientInfoResponse) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ClientInfoResponse) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ClientInfoResponse) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetPublic

`func (o *ClientInfoResponse) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *ClientInfoResponse) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *ClientInfoResponse) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *ClientInfoResponse) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetClientId

`func (o *ClientInfoResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClientInfoResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClientInfoResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ClientInfoResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *ClientInfoResponse) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *ClientInfoResponse) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *ClientInfoResponse) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *ClientInfoResponse) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetTermsUrl

`func (o *ClientInfoResponse) GetTermsUrl() string`

GetTermsUrl returns the TermsUrl field if non-nil, zero value otherwise.

### GetTermsUrlOk

`func (o *ClientInfoResponse) GetTermsUrlOk() (*string, bool)`

GetTermsUrlOk returns a tuple with the TermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsUrl

`func (o *ClientInfoResponse) SetTermsUrl(v string)`

SetTermsUrl sets TermsUrl field to given value.

### HasTermsUrl

`func (o *ClientInfoResponse) HasTermsUrl() bool`

HasTermsUrl returns a boolean if a field has been set.

### GetPolicyUrl

`func (o *ClientInfoResponse) GetPolicyUrl() string`

GetPolicyUrl returns the PolicyUrl field if non-nil, zero value otherwise.

### GetPolicyUrlOk

`func (o *ClientInfoResponse) GetPolicyUrlOk() (*string, bool)`

GetPolicyUrlOk returns a tuple with the PolicyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUrl

`func (o *ClientInfoResponse) SetPolicyUrl(v string)`

SetPolicyUrl sets PolicyUrl field to given value.

### HasPolicyUrl

`func (o *ClientInfoResponse) HasPolicyUrl() bool`

HasPolicyUrl returns a boolean if a field has been set.

### GetLogo

`func (o *ClientInfoResponse) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ClientInfoResponse) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ClientInfoResponse) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ClientInfoResponse) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetAuthenticationMethods

`func (o *ClientInfoResponse) GetAuthenticationMethods() []string`

GetAuthenticationMethods returns the AuthenticationMethods field if non-nil, zero value otherwise.

### GetAuthenticationMethodsOk

`func (o *ClientInfoResponse) GetAuthenticationMethodsOk() (*[]string, bool)`

GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethods

`func (o *ClientInfoResponse) SetAuthenticationMethods(v []string)`

SetAuthenticationMethods sets AuthenticationMethods field to given value.

### HasAuthenticationMethods

`func (o *ClientInfoResponse) HasAuthenticationMethods() bool`

HasAuthenticationMethods returns a boolean if a field has been set.

### GetIsPublic

`func (o *ClientInfoResponse) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *ClientInfoResponse) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *ClientInfoResponse) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *ClientInfoResponse) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetCreatedOn

`func (o *ClientInfoResponse) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *ClientInfoResponse) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *ClientInfoResponse) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *ClientInfoResponse) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ClientInfoResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ClientInfoResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ClientInfoResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ClientInfoResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetModifiedOn

`func (o *ClientInfoResponse) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *ClientInfoResponse) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *ClientInfoResponse) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *ClientInfoResponse) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetModifiedBy

`func (o *ClientInfoResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *ClientInfoResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *ClientInfoResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *ClientInfoResponse) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


