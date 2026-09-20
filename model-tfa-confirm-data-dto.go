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

// checks if the TfaConfirmDataDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TfaConfirmDataDto{}

// TfaConfirmDataDto The confirmation link the caller has to follow to pass the two-factor step, and the cookie it depends on.
type TfaConfirmDataDto struct {
	// The link to open. Its `type` shows which step it is: phone activation or phone authorization for the SMS  method, and authenticator activation or re-verification for the application method. The whole body is empty  when the portal requires no second factor of the caller.
	Url NullableString `json:"url,omitempty"`
	// The name of the confirmation cookie the link is validated against. It is filled in only for the  authenticator-application method; the SMS method returns `url` alone.
	CookieName NullableString `json:"cookieName,omitempty"`
	// The value of that cookie. The call already set it on the response, so it is repeated here only for a client  that does not keep cookies of its own; it is filled in under the same condition as `cookieName`, and a  later call to this operation replaces it.
	CookieValue NullableString `json:"cookieValue,omitempty"`
}

// NewTfaConfirmDataDto instantiates a new TfaConfirmDataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTfaConfirmDataDto() *TfaConfirmDataDto {
	this := TfaConfirmDataDto{}
	return &this
}

// NewTfaConfirmDataDtoWithDefaults instantiates a new TfaConfirmDataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTfaConfirmDataDtoWithDefaults() *TfaConfirmDataDto {
	this := TfaConfirmDataDto{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TfaConfirmDataDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TfaConfirmDataDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *TfaConfirmDataDto) IsUrlSet() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *TfaConfirmDataDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *TfaConfirmDataDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *TfaConfirmDataDto) UnsetUrl() {
	o.Url.Unset()
}

// GetCookieName returns the CookieName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TfaConfirmDataDto) GetCookieName() string {
	if o == nil || IsNil(o.CookieName.Get()) {
		var ret string
		return ret
	}
	return *o.CookieName.Get()
}

// GetCookieNameOk returns a tuple with the CookieName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TfaConfirmDataDto) GetCookieNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CookieName.Get(), o.CookieName.IsSet()
}

// HasCookieName returns a boolean if a field has been set.
func (o *TfaConfirmDataDto) IsCookieNameSet() bool {
	if o != nil && o.CookieName.IsSet() {
		return true
	}

	return false
}

// SetCookieName gets a reference to the given NullableString and assigns it to the CookieName field.
func (o *TfaConfirmDataDto) SetCookieName(v string) {
	o.CookieName.Set(&v)
}
// SetCookieNameNil sets the value for CookieName to be an explicit nil
func (o *TfaConfirmDataDto) SetCookieNameNil() {
	o.CookieName.Set(nil)
}

// UnsetCookieName ensures that no value is present for CookieName, not even an explicit nil
func (o *TfaConfirmDataDto) UnsetCookieName() {
	o.CookieName.Unset()
}

// GetCookieValue returns the CookieValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TfaConfirmDataDto) GetCookieValue() string {
	if o == nil || IsNil(o.CookieValue.Get()) {
		var ret string
		return ret
	}
	return *o.CookieValue.Get()
}

// GetCookieValueOk returns a tuple with the CookieValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TfaConfirmDataDto) GetCookieValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CookieValue.Get(), o.CookieValue.IsSet()
}

// HasCookieValue returns a boolean if a field has been set.
func (o *TfaConfirmDataDto) IsCookieValueSet() bool {
	if o != nil && o.CookieValue.IsSet() {
		return true
	}

	return false
}

// SetCookieValue gets a reference to the given NullableString and assigns it to the CookieValue field.
func (o *TfaConfirmDataDto) SetCookieValue(v string) {
	o.CookieValue.Set(&v)
}
// SetCookieValueNil sets the value for CookieValue to be an explicit nil
func (o *TfaConfirmDataDto) SetCookieValueNil() {
	o.CookieValue.Set(nil)
}

// UnsetCookieValue ensures that no value is present for CookieValue, not even an explicit nil
func (o *TfaConfirmDataDto) UnsetCookieValue() {
	o.CookieValue.Unset()
}

func (o TfaConfirmDataDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TfaConfirmDataDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.CookieName.IsSet() {
		toSerialize["cookieName"] = o.CookieName.Get()
	}
	if o.CookieValue.IsSet() {
		toSerialize["cookieValue"] = o.CookieValue.Get()
	}
	return toSerialize, nil
}

type NullableTfaConfirmDataDto struct {
	value *TfaConfirmDataDto
	isSet bool
}

func (v NullableTfaConfirmDataDto) Get() *TfaConfirmDataDto {
	return v.value
}

func (v *NullableTfaConfirmDataDto) Set(val *TfaConfirmDataDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTfaConfirmDataDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTfaConfirmDataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTfaConfirmDataDto(val *TfaConfirmDataDto) *NullableTfaConfirmDataDto {
	return &NullableTfaConfirmDataDto{value: val, isSet: true}
}

func (v NullableTfaConfirmDataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTfaConfirmDataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

