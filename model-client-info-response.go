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

// checks if the ClientInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientInfoResponse{}

// ClientInfoResponse The consent-facing subset of a client: everything needed to render a consent screen, and nothing that would let a caller act as the client.
type ClientInfoResponse struct {
	// The display name shown to the user on the consent screen, between 3 and 256 characters.
	Name *string `json:"name,omitempty"`
	// The free-text description shown next to the name on the consent screen, at most 255 characters.
	Description *string `json:"description,omitempty"`
	// The permissions the client may ask for, named as they appear in the tenant scope catalogue - for example files:read, rooms:write or openid. A client cannot request a scope that is not listed here.
	Scopes []string `json:"scopes,omitempty"`
	// The generated identifier of the client, sent as client_id in every OAuth2 request. It is assigned when the client is registered and never changes afterwards.
	ClientId *string `json:"client_id,omitempty"`
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

// NewClientInfoResponse instantiates a new ClientInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientInfoResponse() *ClientInfoResponse {
	this := ClientInfoResponse{}
	return &this
}

// NewClientInfoResponseWithDefaults instantiates a new ClientInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientInfoResponseWithDefaults() *ClientInfoResponse {
	this := ClientInfoResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClientInfoResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientInfoResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClientInfoResponse) IsNameSet() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClientInfoResponse) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClientInfoResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientInfoResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClientInfoResponse) IsDescriptionSet() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClientInfoResponse) SetDescription(v string) {
	o.Description = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ClientInfoResponse) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientInfoResponse) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ClientInfoResponse) IsScopesSet() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ClientInfoResponse) SetScopes(v []string) {
	o.Scopes = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ClientInfoResponse) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientInfoResponse) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ClientInfoResponse) IsClientIdSet() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ClientInfoResponse) SetClientId(v string) {
	o.ClientId = &v
}

// GetWebsiteUrl returns the WebsiteUrl field value if set, zero value otherwise.
func (o *ClientInfoResponse) GetWebsiteUrl() string {
	if o == nil || IsNil(o.WebsiteUrl) {
		var ret string
		return ret
	}
	return *o.WebsiteUrl
}

// GetWebsiteUrlOk returns a tuple with the WebsiteUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientInfoResponse) GetWebsiteUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebsiteUrl) {
		return nil, false
	}
	return o.WebsiteUrl, true
}

// HasWebsiteUrl returns a boolean if a field has been set.
func (o *ClientInfoResponse) IsWebsiteUrlSet() bool {
	if o != nil && !IsNil(o.WebsiteUrl) {
		return true
	}

	return false
}

// SetWebsiteUrl gets a reference to the given string and assigns it to the WebsiteUrl field.
func (o *ClientInfoResponse) SetWebsiteUrl(v string) {
	o.WebsiteUrl = &v
}

// GetTermsUrl returns the TermsUrl field value if set, zero value otherwise.
func (o *ClientInfoResponse) GetTermsUrl() string {
	if o == nil || IsNil(o.TermsUrl) {
		var ret string
		return ret
	}
	return *o.TermsUrl
}

// GetTermsUrlOk returns a tuple with the TermsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientInfoResponse) GetTermsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TermsUrl) {
		return nil, false
	}
	return o.TermsUrl, true
}

// HasTermsUrl returns a boolean if a field has been set.
func (o *ClientInfoResponse) IsTermsUrlSet() bool {
	if o != nil && !IsNil(o.TermsUrl) {
		return true
	}

	return false
}

// SetTermsUrl gets a reference to the given string and assigns it to the TermsUrl field.
func (o *ClientInfoResponse) SetTermsUrl(v string) {
	o.TermsUrl = &v
}

// GetPolicyUrl returns the PolicyUrl field value if set, zero value otherwise.
func (o *ClientInfoResponse) GetPolicyUrl() string {
	if o == nil || IsNil(o.PolicyUrl) {
		var ret string
		return ret
	}
	return *o.PolicyUrl
}

// GetPolicyUrlOk returns a tuple with the PolicyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientInfoResponse) GetPolicyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyUrl) {
		return nil, false
	}
	return o.PolicyUrl, true
}

// HasPolicyUrl returns a boolean if a field has been set.
func (o *ClientInfoResponse) IsPolicyUrlSet() bool {
	if o != nil && !IsNil(o.PolicyUrl) {
		return true
	}

	return false
}

// SetPolicyUrl gets a reference to the given string and assigns it to the PolicyUrl field.
func (o *ClientInfoResponse) SetPolicyUrl(v string) {
	o.PolicyUrl = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *ClientInfoResponse) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientInfoResponse) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *ClientInfoResponse) IsLogoSet() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *ClientInfoResponse) SetLogo(v string) {
	o.Logo = &v
}

// GetAuthenticationMethods returns the AuthenticationMethods field value if set, zero value otherwise.
func (o *ClientInfoResponse) GetAuthenticationMethods() []string {
	if o == nil || IsNil(o.AuthenticationMethods) {
		var ret []string
		return ret
	}
	return o.AuthenticationMethods
}

// GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientInfoResponse) GetAuthenticationMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.AuthenticationMethods) {
		return nil, false
	}
	return o.AuthenticationMethods, true
}

// HasAuthenticationMethods returns a boolean if a field has been set.
func (o *ClientInfoResponse) IsAuthenticationMethodsSet() bool {
	if o != nil && !IsNil(o.AuthenticationMethods) {
		return true
	}

	return false
}

// SetAuthenticationMethods gets a reference to the given []string and assigns it to the AuthenticationMethods field.
func (o *ClientInfoResponse) SetAuthenticationMethods(v []string) {
	o.AuthenticationMethods = v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *ClientInfoResponse) GetCreatedOn() time.Time {
	if o == nil || IsNil(o.CreatedOn) {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientInfoResponse) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedOn) {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *ClientInfoResponse) IsCreatedOnSet() bool {
	if o != nil && !IsNil(o.CreatedOn) {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *ClientInfoResponse) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ClientInfoResponse) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientInfoResponse) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ClientInfoResponse) IsCreatedBySet() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ClientInfoResponse) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise.
func (o *ClientInfoResponse) GetModifiedOn() time.Time {
	if o == nil || IsNil(o.ModifiedOn) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientInfoResponse) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModifiedOn) {
		return nil, false
	}
	return o.ModifiedOn, true
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *ClientInfoResponse) IsModifiedOnSet() bool {
	if o != nil && !IsNil(o.ModifiedOn) {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given time.Time and assigns it to the ModifiedOn field.
func (o *ClientInfoResponse) SetModifiedOn(v time.Time) {
	o.ModifiedOn = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *ClientInfoResponse) GetModifiedBy() string {
	if o == nil || IsNil(o.ModifiedBy) {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientInfoResponse) GetModifiedByOk() (*string, bool) {
	if o == nil || IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *ClientInfoResponse) IsModifiedBySet() bool {
	if o != nil && !IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *ClientInfoResponse) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *ClientInfoResponse) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientInfoResponse) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *ClientInfoResponse) IsIsPublicSet() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *ClientInfoResponse) SetIsPublic(v bool) {
	o.IsPublic = &v
}

func (o ClientInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
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

type NullableClientInfoResponse struct {
	value *ClientInfoResponse
	isSet bool
}

func (v NullableClientInfoResponse) Get() *ClientInfoResponse {
	return v.value
}

func (v *NullableClientInfoResponse) Set(val *ClientInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClientInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClientInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientInfoResponse(val *ClientInfoResponse) *NullableClientInfoResponse {
	return &NullableClientInfoResponse{value: val, isSet: true}
}

func (v NullableClientInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

