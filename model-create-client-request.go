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
	// The client name.
	Name *string `json:"name,omitempty"`
	// The description of the client
	Description *string `json:"description,omitempty"`
	// The logo of the client in base64 format
	Logo *string `json:"logo,omitempty" validate:"regexp=^data:image\\/(?:png|jpeg|jpg|svg\\\\+xml);base64,.*.{1,}"`
	// The scopes for the client
	Scopes []string `json:"scopes,omitempty"`
	Public *bool `json:"public,omitempty"`
	// Indicates whether PKCE is allowed for the client
	AllowPkce *bool `json:"allow_pkce,omitempty"`
	// Indicates if the client is public
	IsPublic *bool `json:"is_public,omitempty"`
	// The website URL of the client
	WebsiteUrl *string `json:"website_url,omitempty" validate:"regexp=^(https?:\\/\\/)?([a-zA-Z0-9-]+\\\\.)+[a-zA-Z]{2,}(:\\\\d+)?(\\/[a-zA-Z0-9-._~:\\/?#\\\\[\\\\]@!$&'()*+,;=]*)?$|^https?:\\/\\/(\\\\d{1,3}\\\\.){3}\\\\d{1,3}(:\\\\d+)?(\\/[a-zA-Z0-9-._~:\\/?#\\\\[\\\\]@!$&'()*+,;=]*)?$"`
	// The terms URL of the client
	TermsUrl *string `json:"terms_url,omitempty" validate:"regexp=^(https?:\\/\\/)?([a-zA-Z0-9-]+\\\\.)+[a-zA-Z]{2,}(:\\\\d+)?(\\/[a-zA-Z0-9-._~:\\/?#\\\\[\\\\]@!$&'()*+,;=]*)?$|^https?:\\/\\/(\\\\d{1,3}\\\\.){3}\\\\d{1,3}(:\\\\d+)?(\\/[a-zA-Z0-9-._~:\\/?#\\\\[\\\\]@!$&'()*+,;=]*)?$"`
	// The policy URL of the client
	PolicyUrl *string `json:"policy_url,omitempty" validate:"regexp=^(https?:\\/\\/)?([a-zA-Z0-9-]+\\\\.)+[a-zA-Z]{2,}(:\\\\d+)?(\\/[a-zA-Z0-9-._~:\\/?#\\\\[\\\\]@!$&'()*+,;=]*)?$|^https?:\\/\\/(\\\\d{1,3}\\\\.){3}\\\\d{1,3}(:\\\\d+)?(\\/[a-zA-Z0-9-._~:\\/?#\\\\[\\\\]@!$&'()*+,;=]*)?$"`
	// The redirect URIs for the client
	RedirectUris []string `json:"redirect_uris"`
	// The allowed origins for the client
	AllowedOrigins []string `json:"allowed_origins"`
	// The logout redirect URI for the client
	LogoutRedirectUri *string `json:"logout_redirect_uri,omitempty" validate:"regexp=^(https?:\\/\\/)?([a-zA-Z0-9-]+\\\\.)+[a-zA-Z]{2,}(:\\\\d+)?(\\/[a-zA-Z0-9-._~:\\/?#\\\\[\\\\]@!$&'()*+,;=]*)?$|^https?:\\/\\/(\\\\d{1,3}\\\\.){3}\\\\d{1,3}(:\\\\d+)?(\\/[a-zA-Z0-9-._~:\\/?#\\\\[\\\\]@!$&'()*+,;=]*)?$"`
}

type _CreateClientRequest CreateClientRequest

// NewCreateClientRequest instantiates a new CreateClientRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClientRequest(redirectUris []string, allowedOrigins []string) *CreateClientRequest {
	this := CreateClientRequest{}
	this.RedirectUris = redirectUris
	this.AllowedOrigins = allowedOrigins
	return &this
}

// NewCreateClientRequestWithDefaults instantiates a new CreateClientRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClientRequestWithDefaults() *CreateClientRequest {
	this := CreateClientRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateClientRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateClientRequest) IsNameSet() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateClientRequest) SetName(v string) {
	o.Name = &v
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

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *CreateClientRequest) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *CreateClientRequest) IsLogoSet() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *CreateClientRequest) SetLogo(v string) {
	o.Logo = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *CreateClientRequest) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *CreateClientRequest) IsScopesSet() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *CreateClientRequest) SetScopes(v []string) {
	o.Scopes = v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *CreateClientRequest) GetPublic() bool {
	if o == nil || IsNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *CreateClientRequest) IsPublicSet() bool {
	if o != nil && !IsNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *CreateClientRequest) SetPublic(v bool) {
	o.Public = &v
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

// GetWebsiteUrl returns the WebsiteUrl field value if set, zero value otherwise.
func (o *CreateClientRequest) GetWebsiteUrl() string {
	if o == nil || IsNil(o.WebsiteUrl) {
		var ret string
		return ret
	}
	return *o.WebsiteUrl
}

// GetWebsiteUrlOk returns a tuple with the WebsiteUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetWebsiteUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebsiteUrl) {
		return nil, false
	}
	return o.WebsiteUrl, true
}

// HasWebsiteUrl returns a boolean if a field has been set.
func (o *CreateClientRequest) IsWebsiteUrlSet() bool {
	if o != nil && !IsNil(o.WebsiteUrl) {
		return true
	}

	return false
}

// SetWebsiteUrl gets a reference to the given string and assigns it to the WebsiteUrl field.
func (o *CreateClientRequest) SetWebsiteUrl(v string) {
	o.WebsiteUrl = &v
}

// GetTermsUrl returns the TermsUrl field value if set, zero value otherwise.
func (o *CreateClientRequest) GetTermsUrl() string {
	if o == nil || IsNil(o.TermsUrl) {
		var ret string
		return ret
	}
	return *o.TermsUrl
}

// GetTermsUrlOk returns a tuple with the TermsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetTermsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TermsUrl) {
		return nil, false
	}
	return o.TermsUrl, true
}

// HasTermsUrl returns a boolean if a field has been set.
func (o *CreateClientRequest) IsTermsUrlSet() bool {
	if o != nil && !IsNil(o.TermsUrl) {
		return true
	}

	return false
}

// SetTermsUrl gets a reference to the given string and assigns it to the TermsUrl field.
func (o *CreateClientRequest) SetTermsUrl(v string) {
	o.TermsUrl = &v
}

// GetPolicyUrl returns the PolicyUrl field value if set, zero value otherwise.
func (o *CreateClientRequest) GetPolicyUrl() string {
	if o == nil || IsNil(o.PolicyUrl) {
		var ret string
		return ret
	}
	return *o.PolicyUrl
}

// GetPolicyUrlOk returns a tuple with the PolicyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetPolicyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyUrl) {
		return nil, false
	}
	return o.PolicyUrl, true
}

// HasPolicyUrl returns a boolean if a field has been set.
func (o *CreateClientRequest) IsPolicyUrlSet() bool {
	if o != nil && !IsNil(o.PolicyUrl) {
		return true
	}

	return false
}

// SetPolicyUrl gets a reference to the given string and assigns it to the PolicyUrl field.
func (o *CreateClientRequest) SetPolicyUrl(v string) {
	o.PolicyUrl = &v
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

// GetLogoutRedirectUri returns the LogoutRedirectUri field value if set, zero value otherwise.
func (o *CreateClientRequest) GetLogoutRedirectUri() string {
	if o == nil || IsNil(o.LogoutRedirectUri) {
		var ret string
		return ret
	}
	return *o.LogoutRedirectUri
}

// GetLogoutRedirectUriOk returns a tuple with the LogoutRedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClientRequest) GetLogoutRedirectUriOk() (*string, bool) {
	if o == nil || IsNil(o.LogoutRedirectUri) {
		return nil, false
	}
	return o.LogoutRedirectUri, true
}

// HasLogoutRedirectUri returns a boolean if a field has been set.
func (o *CreateClientRequest) IsLogoutRedirectUriSet() bool {
	if o != nil && !IsNil(o.LogoutRedirectUri) {
		return true
	}

	return false
}

// SetLogoutRedirectUri gets a reference to the given string and assigns it to the LogoutRedirectUri field.
func (o *CreateClientRequest) SetLogoutRedirectUri(v string) {
	o.LogoutRedirectUri = &v
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	if !IsNil(o.AllowPkce) {
		toSerialize["allow_pkce"] = o.AllowPkce
	}
	if !IsNil(o.IsPublic) {
		toSerialize["is_public"] = o.IsPublic
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
	toSerialize["redirect_uris"] = o.RedirectUris
	toSerialize["allowed_origins"] = o.AllowedOrigins
	if !IsNil(o.LogoutRedirectUri) {
		toSerialize["logout_redirect_uri"] = o.LogoutRedirectUri
	}
	return toSerialize, nil
}

func (o *CreateClientRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"redirect_uris",
		"allowed_origins",
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

