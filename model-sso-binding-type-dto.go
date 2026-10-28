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
)

// checks if the SsoBindingTypeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsoBindingTypeDto{}

// SsoBindingTypeDto The SAML bindings the SSO settings accept.
type SsoBindingTypeDto struct {
	// The SAML 2.0 HTTP POST binding, which carries the request in a self-submitting form. It is what the  built-in configuration uses and the one to pick when requests are signed, since it has no length limit.
	Saml20HttpPost NullableString `json:"saml20HttpPost,omitempty"`
	// The SAML 2.0 HTTP redirect binding, which carries the request in the query string and is therefore bound  by the length a URL may have.
	Saml20HttpRedirect NullableString `json:"saml20HttpRedirect,omitempty"`
}

// NewSsoBindingTypeDto instantiates a new SsoBindingTypeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsoBindingTypeDto() *SsoBindingTypeDto {
	this := SsoBindingTypeDto{}
	return &this
}

// NewSsoBindingTypeDtoWithDefaults instantiates a new SsoBindingTypeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoBindingTypeDtoWithDefaults() *SsoBindingTypeDto {
	this := SsoBindingTypeDto{}
	return &this
}

// GetSaml20HttpPost returns the Saml20HttpPost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoBindingTypeDto) GetSaml20HttpPost() string {
	if o == nil || IsNil(o.Saml20HttpPost.Get()) {
		var ret string
		return ret
	}
	return *o.Saml20HttpPost.Get()
}

// GetSaml20HttpPostOk returns a tuple with the Saml20HttpPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoBindingTypeDto) GetSaml20HttpPostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Saml20HttpPost.Get(), o.Saml20HttpPost.IsSet()
}

// HasSaml20HttpPost returns a boolean if a field has been set.
func (o *SsoBindingTypeDto) IsSaml20HttpPostSet() bool {
	if o != nil && o.Saml20HttpPost.IsSet() {
		return true
	}

	return false
}

// SetSaml20HttpPost gets a reference to the given NullableString and assigns it to the Saml20HttpPost field.
func (o *SsoBindingTypeDto) SetSaml20HttpPost(v string) {
	o.Saml20HttpPost.Set(&v)
}
// SetSaml20HttpPostNil sets the value for Saml20HttpPost to be an explicit nil
func (o *SsoBindingTypeDto) SetSaml20HttpPostNil() {
	o.Saml20HttpPost.Set(nil)
}

// UnsetSaml20HttpPost ensures that no value is present for Saml20HttpPost, not even an explicit nil
func (o *SsoBindingTypeDto) UnsetSaml20HttpPost() {
	o.Saml20HttpPost.Unset()
}

// GetSaml20HttpRedirect returns the Saml20HttpRedirect field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SsoBindingTypeDto) GetSaml20HttpRedirect() string {
	if o == nil || IsNil(o.Saml20HttpRedirect.Get()) {
		var ret string
		return ret
	}
	return *o.Saml20HttpRedirect.Get()
}

// GetSaml20HttpRedirectOk returns a tuple with the Saml20HttpRedirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SsoBindingTypeDto) GetSaml20HttpRedirectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Saml20HttpRedirect.Get(), o.Saml20HttpRedirect.IsSet()
}

// HasSaml20HttpRedirect returns a boolean if a field has been set.
func (o *SsoBindingTypeDto) IsSaml20HttpRedirectSet() bool {
	if o != nil && o.Saml20HttpRedirect.IsSet() {
		return true
	}

	return false
}

// SetSaml20HttpRedirect gets a reference to the given NullableString and assigns it to the Saml20HttpRedirect field.
func (o *SsoBindingTypeDto) SetSaml20HttpRedirect(v string) {
	o.Saml20HttpRedirect.Set(&v)
}
// SetSaml20HttpRedirectNil sets the value for Saml20HttpRedirect to be an explicit nil
func (o *SsoBindingTypeDto) SetSaml20HttpRedirectNil() {
	o.Saml20HttpRedirect.Set(nil)
}

// UnsetSaml20HttpRedirect ensures that no value is present for Saml20HttpRedirect, not even an explicit nil
func (o *SsoBindingTypeDto) UnsetSaml20HttpRedirect() {
	o.Saml20HttpRedirect.Unset()
}

func (o SsoBindingTypeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsoBindingTypeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Saml20HttpPost.IsSet() {
		toSerialize["saml20HttpPost"] = o.Saml20HttpPost.Get()
	}
	if o.Saml20HttpRedirect.IsSet() {
		toSerialize["saml20HttpRedirect"] = o.Saml20HttpRedirect.Get()
	}
	return toSerialize, nil
}

type NullableSsoBindingTypeDto struct {
	value *SsoBindingTypeDto
	isSet bool
}

func (v NullableSsoBindingTypeDto) Get() *SsoBindingTypeDto {
	return v.value
}

func (v *NullableSsoBindingTypeDto) Set(val *SsoBindingTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSsoBindingTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSsoBindingTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsoBindingTypeDto(val *SsoBindingTypeDto) *NullableSsoBindingTypeDto {
	return &NullableSsoBindingTypeDto{value: val, isSet: true}
}

func (v NullableSsoBindingTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsoBindingTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

