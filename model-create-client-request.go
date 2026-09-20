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
	"bytes"
	"fmt"
)

// checks if the CreateClientRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateClientRequest{}

// CreateClientRequest Client creation request containing client details
type CreateClientRequest struct {
	// The display name shown to the user on the consent screen. It has to be between 3 and 256 characters long.
	Name string `json:"name"`
	// The free-text description shown next to the name on the consent screen, at most 255 characters.
	Description *string `json:"description,omitempty"`
	// The client logo as a data URI carrying base64 image data, shown on the consent screen. Only png, jpeg, jpg and svg+xml are accepted, the whole string may not exceed 2000000 characters and the decoded image may not exceed 256000 bytes.
	Logo string `json:"logo" validate:"regexp=^data:image/(?:png|jpeg|jpg|svg\\+xml);base64\\,.*.{1\\,}"`
	// The permissions the client may ask for, named as they appear in the tenant scope catalogue - for example files:read, rooms:write or openid. A client cannot request a scope that is not listed here.
	Scopes []string `json:"scopes"`
	// Whether the client may use PKCE. Turning it on lets the client authenticate with the none method and prove itself with a code verifier instead of sending a secret, which is what a client that cannot keep a secret needs.
	AllowPkce *bool `json:"allow_pkce,omitempty"`
	// The URL of the client home page, offered to the user before they consent. The value has to be an http or https URL.
	WebsiteUrl string `json:"website_url" validate:"regexp=^(https?://)?([a-zA-Z0-9-]+\\.)+[a-zA-Z]{2\\,}(:\\d+)?(/[a-zA-Z0-9-._~:/?#\\[\\]@!$&'()*+\\,;=]*)?$|^https?://(\\d{1\\,3}\\.){3}\\d{1\\,3}(:\\d+)?(/[a-zA-Z0-9-._~:/?#\\[\\]@!$&'()*+\\,;=]*)?$"`
	// The URL of the client terms of service, linked from the consent screen. The value has to be an http or https URL.
	TermsUrl string `json:"terms_url" validate:"regexp=^(https?://)?([a-zA-Z0-9-]+\\.)+[a-zA-Z]{2\\,}(:\\d+)?(/[a-zA-Z0-9-._~:/?#\\[\\]@!$&'()*+\\,;=]*)?$|^https?://(\\d{1\\,3}\\.){3}\\d{1\\,3}(:\\d+)?(/[a-zA-Z0-9-._~:/?#\\[\\]@!$&'()*+\\,;=]*)?$"`
	// The URL of the client privacy policy, linked from the consent screen. The value has to be an http or https URL.
	PolicyUrl string `json:"policy_url" validate:"regexp=^(https?://)?([a-zA-Z0-9-]+\\.)+[a-zA-Z]{2\\,}(:\\d+)?(/[a-zA-Z0-9-._~:/?#\\[\\]@!$&'()*+\\,;=]*)?$|^https?://(\\d{1\\,3}\\.){3}\\d{1\\,3}(:\\d+)?(/[a-zA-Z0-9-._~:/?#\\[\\]@!$&'()*+\\,;=]*)?$"`
	// The URIs an authorization code may be delivered to. An authorization request naming any other URI is refused, and the set holds between 1 and 12 addresses.
	RedirectUris []string `json:"redirect_uris"`
	// The web origins allowed to call the portal on behalf of this client, used for the CORS check. The set holds between 1 and 12 addresses.
	AllowedOrigins []string `json:"allowed_origins"`
	// The single URI the user may be sent back to once they have logged out. The value has to be an http or https URL.
	LogoutRedirectUri string `json:"logout_redirect_uri" validate:"regexp=^(https?://)?([a-zA-Z0-9-]+\\.)+[a-zA-Z]{2\\,}(:\\d+)?(/[a-zA-Z0-9-._~:/?#\\[\\]@!$&'()*+\\,;=]*)?$|^https?://(\\d{1\\,3}\\.){3}\\d{1\\,3}(:\\d+)?(/[a-zA-Z0-9-._~:/?#\\[\\]@!$&'()*+\\,;=]*)?$"`
	// Whether the client is offered to third-party tenants rather than only to the tenant that registers it.
	IsPublic *bool `json:"is_public,omitempty"`
}

type _CreateClientRequest CreateClientRequest

// NewCreateClientRequest instantiates a new CreateClientRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClientRequest(name string, logo string, scopes []string, websiteUrl string, termsUrl string, policyUrl string, redirectUris []string, allowedOrigins []string, logoutRedirectUri string) *CreateClientRequest {
	this := CreateClientRequest{}
	this.Name = name
	this.Logo = logo
	this.Scopes = scopes
	this.WebsiteUrl = websiteUrl
	this.TermsUrl = termsUrl
	this.PolicyUrl = policyUrl
	this.RedirectUris = redirectUris
	this.AllowedOrigins = allowedOrigins
	this.LogoutRedirectUri = logoutRedirectUri
	return &this
}

// NewCreateClientRequestWithDefaults instantiates a new CreateClientRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClientRequestWithDefaults() *CreateClientRequest {
	this := CreateClientRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateClientRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateClientRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateClientRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateClientRequest) IsDescriptionSet() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateClientRequest) SetDescription(v string) {
	o.Description = &v
}

// GetLogo returns the Logo field value
func (o *CreateClientRequest) GetLogo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logo, true
}

// SetLogo sets field value
func (o *CreateClientRequest) SetLogo(v string) {
	o.Logo = v
}

// GetScopes returns the Scopes field value
func (o *CreateClientRequest) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *CreateClientRequest) SetScopes(v []string) {
	o.Scopes = v
}

// GetAllowPkce returns the AllowPkce field value if set, zero value otherwise.
func (o *CreateClientRequest) GetAllowPkce() bool {
	if o == nil || IsNil(o.AllowPkce) {
		var ret bool
		return ret
	}
	return *o.AllowPkce
}

// GetAllowPkceOk returns a tuple with the AllowPkce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetAllowPkceOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPkce) {
		return nil, false
	}
	return o.AllowPkce, true
}

// HasAllowPkce returns a boolean if a field has been set.
func (o *CreateClientRequest) IsAllowPkceSet() bool {
	if o != nil && !IsNil(o.AllowPkce) {
		return true
	}

	return false
}

// SetAllowPkce gets a reference to the given bool and assigns it to the AllowPkce field.
func (o *CreateClientRequest) SetAllowPkce(v bool) {
	o.AllowPkce = &v
}

// GetWebsiteUrl returns the WebsiteUrl field value
func (o *CreateClientRequest) GetWebsiteUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebsiteUrl
}

// GetWebsiteUrlOk returns a tuple with the WebsiteUrl field value
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetWebsiteUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebsiteUrl, true
}

// SetWebsiteUrl sets field value
func (o *CreateClientRequest) SetWebsiteUrl(v string) {
	o.WebsiteUrl = v
}

// GetTermsUrl returns the TermsUrl field value
func (o *CreateClientRequest) GetTermsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TermsUrl
}

// GetTermsUrlOk returns a tuple with the TermsUrl field value
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetTermsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TermsUrl, true
}

// SetTermsUrl sets field value
func (o *CreateClientRequest) SetTermsUrl(v string) {
	o.TermsUrl = v
}

// GetPolicyUrl returns the PolicyUrl field value
func (o *CreateClientRequest) GetPolicyUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyUrl
}

// GetPolicyUrlOk returns a tuple with the PolicyUrl field value
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetPolicyUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyUrl, true
}

// SetPolicyUrl sets field value
func (o *CreateClientRequest) SetPolicyUrl(v string) {
	o.PolicyUrl = v
}

// GetRedirectUris returns the RedirectUris field value
func (o *CreateClientRequest) GetRedirectUris() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetRedirectUrisOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// SetRedirectUris sets field value
func (o *CreateClientRequest) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

// GetAllowedOrigins returns the AllowedOrigins field value
func (o *CreateClientRequest) GetAllowedOrigins() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllowedOrigins
}

// GetAllowedOriginsOk returns a tuple with the AllowedOrigins field value
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetAllowedOriginsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedOrigins, true
}

// SetAllowedOrigins sets field value
func (o *CreateClientRequest) SetAllowedOrigins(v []string) {
	o.AllowedOrigins = v
}

// GetLogoutRedirectUri returns the LogoutRedirectUri field value
func (o *CreateClientRequest) GetLogoutRedirectUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogoutRedirectUri
}

// GetLogoutRedirectUriOk returns a tuple with the LogoutRedirectUri field value
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetLogoutRedirectUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogoutRedirectUri, true
}

// SetLogoutRedirectUri sets field value
func (o *CreateClientRequest) SetLogoutRedirectUri(v string) {
	o.LogoutRedirectUri = v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *CreateClientRequest) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *CreateClientRequest) IsIsPublicSet() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *CreateClientRequest) SetIsPublic(v bool) {
	o.IsPublic = &v
}

func (o CreateClientRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateClientRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["logo"] = o.Logo
	toSerialize["scopes"] = o.Scopes
	if !IsNil(o.AllowPkce) {
		toSerialize["allow_pkce"] = o.AllowPkce
	}
	toSerialize["website_url"] = o.WebsiteUrl
	toSerialize["terms_url"] = o.TermsUrl
	toSerialize["policy_url"] = o.PolicyUrl
	toSerialize["redirect_uris"] = o.RedirectUris
	toSerialize["allowed_origins"] = o.AllowedOrigins
	toSerialize["logout_redirect_uri"] = o.LogoutRedirectUri
	if !IsNil(o.IsPublic) {
		toSerialize["is_public"] = o.IsPublic
	}
	return toSerialize, nil
}

func (o *CreateClientRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"logo",
		"scopes",
		"website_url",
		"terms_url",
		"policy_url",
		"redirect_uris",
		"allowed_origins",
		"logout_redirect_uri",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateClientRequest := _CreateClientRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCreateClientRequest)

	if err != nil {
		return err
	}

	*o = CreateClientRequest(varCreateClientRequest)

	return err
}

type NullableCreateClientRequest struct {
	value *CreateClientRequest
	isSet bool
}

func (v NullableCreateClientRequest) Get() *CreateClientRequest {
	return v.value
}

func (v *NullableCreateClientRequest) Set(val *CreateClientRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClientRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClientRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClientRequest(val *CreateClientRequest) *NullableCreateClientRequest {
	return &NullableCreateClientRequest{value: val, isSet: true}
}

func (v NullableCreateClientRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClientRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

