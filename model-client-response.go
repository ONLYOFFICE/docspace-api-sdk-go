// (c) Copyright Ascensio System SIA 2026
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docspace_api_sdk

import (
	"encoding/json"
	"time"
)

// checks if the ClientResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientResponse{}

// ClientResponse The whole stored record of an OAuth2 client, including the secret and every address the client is allowed to use.
type ClientResponse struct {
	// The display name shown to the user on the consent screen, between 3 and 256 characters.
	Name *string `json:"name,omitempty"`
	// The free-text description shown next to the name on the consent screen, at most 255 characters.
	Description *string `json:"description,omitempty"`
	// The identifier of the portal the client belongs to. A client is visible only inside its own tenant, apart from the unauthenticated public info read.
	Tenant *int64 `json:"tenant,omitempty"`
	// The permissions the client may ask for, named as they appear in the tenant scope catalogue - for example files:read, rooms:write or openid. A client cannot request a scope that is not listed here.
	Scopes []string `json:"scopes,omitempty"`
	// Whether the client may currently obtain tokens. A disabled client keeps its registration and the tokens already issued to it, but new authorization requests for it are refused.
	Enabled *bool `json:"enabled,omitempty"`
	// The generated identifier of the client, sent as client_id in every OAuth2 request. It is assigned when the client is registered and never changes afterwards.
	ClientId *string `json:"client_id,omitempty"`
	// The client secret, which the client presents at the token endpoint when it authenticates with client_secret_post. It is omitted from the response rather than sent as null when the client has none.
	ClientSecret *string `json:"client_secret,omitempty"`
	// The URL of the client home page, offered to the user before they consent.
	WebsiteUrl *string `json:"website_url,omitempty"`
	// The URL of the client terms of service, linked from the consent screen.
	TermsUrl *string `json:"terms_url,omitempty"`
	// The URL of the client privacy policy, linked from the consent screen.
	PolicyUrl *string `json:"policy_url,omitempty"`
	// The client logo as a data URI carrying base64 image data, shown on the consent screen. Only png, jpeg, jpg and svg+xml are accepted, the whole string may not exceed 2000000 characters and the decoded image may not exceed 256000 bytes.
	Logo *string `json:"logo,omitempty"`
	// How the client authenticates itself at the token endpoint: client_secret_post for a confidential client that sends its secret, none for a public client that proves itself with PKCE instead.
	AuthenticationMethods []string `json:"authentication_methods,omitempty"`
	// The URIs an authorization code may be delivered to. An authorization request naming any other URI is refused, and the set holds between 1 and 12 addresses.
	RedirectUris []string `json:"redirect_uris,omitempty"`
	// The web origins allowed to call the portal on behalf of this client, used for the CORS check. The set holds between 1 and 12 addresses.
	AllowedOrigins []string `json:"allowed_origins,omitempty"`
	// The URIs the user may be sent back to once they have logged out.
	LogoutRedirectUris []string `json:"logout_redirect_uris,omitempty"`
	// When the client was registered, as an ISO-8601 timestamp with a zone offset.
	CreatedOn *time.Time `json:"created_on,omitempty"`
	// The identifier of the user who registered the client. A plain user may read and change only the clients where this is their own identifier.
	CreatedBy *string `json:"created_by,omitempty"`
	// When the client was last changed, as an ISO-8601 timestamp with a zone offset.
	ModifiedOn *time.Time `json:"modified_on,omitempty"`
	// The identifier of the user who last changed the client.
	ModifiedBy *string `json:"modified_by,omitempty"`
	// Whether the client is offered to third-party tenants rather than only to the tenant that registered it.
	IsPublic *bool `json:"is_public,omitempty"`
}

// NewClientResponse instantiates a new ClientResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientResponse() *ClientResponse {
	this := ClientResponse{}
	return &this
}

// NewClientResponseWithDefaults instantiates a new ClientResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientResponseWithDefaults() *ClientResponse {
	this := ClientResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClientResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClientResponse) IsNameSet() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClientResponse) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClientResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClientResponse) IsDescriptionSet() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClientResponse) SetDescription(v string) {
	o.Description = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *ClientResponse) GetTenant() int64 {
	if o == nil || IsNil(o.Tenant) {
		var ret int64
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetTenantOk() (*int64, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *ClientResponse) IsTenantSet() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given int64 and assigns it to the Tenant field.
func (o *ClientResponse) SetTenant(v int64) {
	o.Tenant = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ClientResponse) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ClientResponse) IsScopesSet() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ClientResponse) SetScopes(v []string) {
	o.Scopes = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ClientResponse) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ClientResponse) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ClientResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ClientResponse) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ClientResponse) IsClientIdSet() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ClientResponse) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *ClientResponse) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *ClientResponse) IsClientSecretSet() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *ClientResponse) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetWebsiteUrl returns the WebsiteUrl field value if set, zero value otherwise.
func (o *ClientResponse) GetWebsiteUrl() string {
	if o == nil || IsNil(o.WebsiteUrl) {
		var ret string
		return ret
	}
	return *o.WebsiteUrl
}

// GetWebsiteUrlOk returns a tuple with the WebsiteUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetWebsiteUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebsiteUrl) {
		return nil, false
	}
	return o.WebsiteUrl, true
}

// HasWebsiteUrl returns a boolean if a field has been set.
func (o *ClientResponse) IsWebsiteUrlSet() bool {
	if o != nil && !IsNil(o.WebsiteUrl) {
		return true
	}

	return false
}

// SetWebsiteUrl gets a reference to the given string and assigns it to the WebsiteUrl field.
func (o *ClientResponse) SetWebsiteUrl(v string) {
	o.WebsiteUrl = &v
}

// GetTermsUrl returns the TermsUrl field value if set, zero value otherwise.
func (o *ClientResponse) GetTermsUrl() string {
	if o == nil || IsNil(o.TermsUrl) {
		var ret string
		return ret
	}
	return *o.TermsUrl
}

// GetTermsUrlOk returns a tuple with the TermsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetTermsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TermsUrl) {
		return nil, false
	}
	return o.TermsUrl, true
}

// HasTermsUrl returns a boolean if a field has been set.
func (o *ClientResponse) IsTermsUrlSet() bool {
	if o != nil && !IsNil(o.TermsUrl) {
		return true
	}

	return false
}

// SetTermsUrl gets a reference to the given string and assigns it to the TermsUrl field.
func (o *ClientResponse) SetTermsUrl(v string) {
	o.TermsUrl = &v
}

// GetPolicyUrl returns the PolicyUrl field value if set, zero value otherwise.
func (o *ClientResponse) GetPolicyUrl() string {
	if o == nil || IsNil(o.PolicyUrl) {
		var ret string
		return ret
	}
	return *o.PolicyUrl
}

// GetPolicyUrlOk returns a tuple with the PolicyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetPolicyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyUrl) {
		return nil, false
	}
	return o.PolicyUrl, true
}

// HasPolicyUrl returns a boolean if a field has been set.
func (o *ClientResponse) IsPolicyUrlSet() bool {
	if o != nil && !IsNil(o.PolicyUrl) {
		return true
	}

	return false
}

// SetPolicyUrl gets a reference to the given string and assigns it to the PolicyUrl field.
func (o *ClientResponse) SetPolicyUrl(v string) {
	o.PolicyUrl = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *ClientResponse) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *ClientResponse) IsLogoSet() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *ClientResponse) SetLogo(v string) {
	o.Logo = &v
}

// GetAuthenticationMethods returns the AuthenticationMethods field value if set, zero value otherwise.
func (o *ClientResponse) GetAuthenticationMethods() []string {
	if o == nil || IsNil(o.AuthenticationMethods) {
		var ret []string
		return ret
	}
	return o.AuthenticationMethods
}

// GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetAuthenticationMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.AuthenticationMethods) {
		return nil, false
	}
	return o.AuthenticationMethods, true
}

// HasAuthenticationMethods returns a boolean if a field has been set.
func (o *ClientResponse) IsAuthenticationMethodsSet() bool {
	if o != nil && !IsNil(o.AuthenticationMethods) {
		return true
	}

	return false
}

// SetAuthenticationMethods gets a reference to the given []string and assigns it to the AuthenticationMethods field.
func (o *ClientResponse) SetAuthenticationMethods(v []string) {
	o.AuthenticationMethods = v
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *ClientResponse) GetRedirectUris() []string {
	if o == nil || IsNil(o.RedirectUris) {
		var ret []string
		return ret
	}
	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetRedirectUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.RedirectUris) {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *ClientResponse) IsRedirectUrisSet() bool {
	if o != nil && !IsNil(o.RedirectUris) {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *ClientResponse) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

// GetAllowedOrigins returns the AllowedOrigins field value if set, zero value otherwise.
func (o *ClientResponse) GetAllowedOrigins() []string {
	if o == nil || IsNil(o.AllowedOrigins) {
		var ret []string
		return ret
	}
	return o.AllowedOrigins
}

// GetAllowedOriginsOk returns a tuple with the AllowedOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetAllowedOriginsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedOrigins) {
		return nil, false
	}
	return o.AllowedOrigins, true
}

// HasAllowedOrigins returns a boolean if a field has been set.
func (o *ClientResponse) IsAllowedOriginsSet() bool {
	if o != nil && !IsNil(o.AllowedOrigins) {
		return true
	}

	return false
}

// SetAllowedOrigins gets a reference to the given []string and assigns it to the AllowedOrigins field.
func (o *ClientResponse) SetAllowedOrigins(v []string) {
	o.AllowedOrigins = v
}

// GetLogoutRedirectUris returns the LogoutRedirectUris field value if set, zero value otherwise.
func (o *ClientResponse) GetLogoutRedirectUris() []string {
	if o == nil || IsNil(o.LogoutRedirectUris) {
		var ret []string
		return ret
	}
	return o.LogoutRedirectUris
}

// GetLogoutRedirectUrisOk returns a tuple with the LogoutRedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetLogoutRedirectUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.LogoutRedirectUris) {
		return nil, false
	}
	return o.LogoutRedirectUris, true
}

// HasLogoutRedirectUris returns a boolean if a field has been set.
func (o *ClientResponse) IsLogoutRedirectUrisSet() bool {
	if o != nil && !IsNil(o.LogoutRedirectUris) {
		return true
	}

	return false
}

// SetLogoutRedirectUris gets a reference to the given []string and assigns it to the LogoutRedirectUris field.
func (o *ClientResponse) SetLogoutRedirectUris(v []string) {
	o.LogoutRedirectUris = v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *ClientResponse) GetCreatedOn() time.Time {
	if o == nil || IsNil(o.CreatedOn) {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedOn) {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *ClientResponse) IsCreatedOnSet() bool {
	if o != nil && !IsNil(o.CreatedOn) {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *ClientResponse) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ClientResponse) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ClientResponse) IsCreatedBySet() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ClientResponse) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *ClientResponse) GetModifiedOn() time.Time {
	if o == nil || IsNil(o.ModifiedOn) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModifiedOn) {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *ClientResponse) IsModifiedOnSet() bool {
	if o != nil && !IsNil(o.ModifiedOn) {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given time.Time and assigns it to the ModifiedOn field.
func (o *ClientResponse) SetModifiedOn(v time.Time) {
	o.ModifiedOn = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *ClientResponse) GetModifiedBy() string {
	if o == nil || IsNil(o.ModifiedBy) {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetModifiedByOk() (*string, bool) {
	if o == nil || IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *ClientResponse) IsModifiedBySet() bool {
	if o != nil && !IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *ClientResponse) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *ClientResponse) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *ClientResponse) IsIsPublicSet() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *ClientResponse) SetIsPublic(v bool) {
	o.IsPublic = &v
}

func (o ClientResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.ClientSecret) {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if !IsNil(o.WebsiteUrl) {
		toSerialize["website_url"] = o.WebsiteUrl
	}
	if !IsNil(o.TermsUrl) {
		toSerialize["terms_url"] = o.TermsUrl
	}
	if !IsNil(o.PolicyUrl) {
		toSerialize["policy_url"] = o.PolicyUrl
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.AuthenticationMethods) {
		toSerialize["authentication_methods"] = o.AuthenticationMethods
	}
	if !IsNil(o.RedirectUris) {
		toSerialize["redirect_uris"] = o.RedirectUris
	}
	if !IsNil(o.AllowedOrigins) {
		toSerialize["allowed_origins"] = o.AllowedOrigins
	}
	if !IsNil(o.LogoutRedirectUris) {
		toSerialize["logout_redirect_uris"] = o.LogoutRedirectUris
	}
	if !IsNil(o.CreatedOn) {
		toSerialize["created_on"] = o.CreatedOn
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.ModifiedOn) {
		toSerialize["modified_on"] = o.ModifiedOn
	}
	if !IsNil(o.ModifiedBy) {
		toSerialize["modified_by"] = o.ModifiedBy
	}
	if !IsNil(o.IsPublic) {
		toSerialize["is_public"] = o.IsPublic
	}
	return toSerialize, nil
}

type NullableClientResponse struct {
	value *ClientResponse
	isSet bool
}

func (v NullableClientResponse) Get() *ClientResponse {
	return v.value
}

func (v *NullableClientResponse) Set(val *ClientResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClientResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClientResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientResponse(val *ClientResponse) *NullableClientResponse {
	return &NullableClientResponse{value: val, isSet: true}
}

func (v NullableClientResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

