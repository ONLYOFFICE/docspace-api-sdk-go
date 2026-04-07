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

// checks if the WhiteLabelRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WhiteLabelRequestsDto{}

// WhiteLabelRequestsDto The request parameters for configuring the white label branding settings.
type WhiteLabelRequestsDto struct {
	// The text to display alongside or in place of the logo.
	LogoText NullableString `json:"logoText,omitempty"`
	// The white label tenant IDs with their logos (light or dark).
	Logo []ItemKeyValuePairStringLogoRequestsDto `json:"logo,omitempty"`
}

// NewWhiteLabelRequestsDto instantiates a new WhiteLabelRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhiteLabelRequestsDto() *WhiteLabelRequestsDto {
	this := WhiteLabelRequestsDto{}
	return &this
}

// NewWhiteLabelRequestsDtoWithDefaults instantiates a new WhiteLabelRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhiteLabelRequestsDtoWithDefaults() *WhiteLabelRequestsDto {
	this := WhiteLabelRequestsDto{}
	return &this
}

// GetLogoText returns the LogoText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WhiteLabelRequestsDto) GetLogoText() string {
	if o == nil || IsNil(o.LogoText.Get()) {
		var ret string
		return ret
	}
	return *o.LogoText.Get()
}

// GetLogoTextOk returns a tuple with the LogoText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WhiteLabelRequestsDto) GetLogoTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoText.Get(), o.LogoText.IsSet()
}

// HasLogoText returns a boolean if a field has been set.
func (o *WhiteLabelRequestsDto) IsLogoTextSet() bool {
	if o != nil && o.LogoText.IsSet() {
		return true
	}

	return false
}

// SetLogoText gets a reference to the given NullableString and assigns it to the LogoText field.
func (o *WhiteLabelRequestsDto) SetLogoText(v string) {
	o.LogoText.Set(&v)
}
// SetLogoTextNil sets the value for LogoText to be an explicit nil
func (o *WhiteLabelRequestsDto) SetLogoTextNil() {
	o.LogoText.Set(nil)
}

// UnsetLogoText ensures that no value is present for LogoText, not even an explicit nil
func (o *WhiteLabelRequestsDto) UnsetLogoText() {
	o.LogoText.Unset()
}

// GetLogo returns the Logo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WhiteLabelRequestsDto) GetLogo() []ItemKeyValuePairStringLogoRequestsDto {
	if o == nil {
		var ret []ItemKeyValuePairStringLogoRequestsDto
		return ret
	}
	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WhiteLabelRequestsDto) GetLogoOk() ([]ItemKeyValuePairStringLogoRequestsDto, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *WhiteLabelRequestsDto) IsLogoSet() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given []ItemKeyValuePairStringLogoRequestsDto and assigns it to the Logo field.
func (o *WhiteLabelRequestsDto) SetLogo(v []ItemKeyValuePairStringLogoRequestsDto) {
	o.Logo = v
}

func (o WhiteLabelRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WhiteLabelRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LogoText.IsSet() {
		toSerialize["logoText"] = o.LogoText.Get()
	}
	if o.Logo != nil {
		toSerialize["logo"] = o.Logo
	}
	return toSerialize, nil
}

type NullableWhiteLabelRequestsDto struct {
	value *WhiteLabelRequestsDto
	isSet bool
}

func (v NullableWhiteLabelRequestsDto) Get() *WhiteLabelRequestsDto {
	return v.value
}

func (v *NullableWhiteLabelRequestsDto) Set(val *WhiteLabelRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWhiteLabelRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWhiteLabelRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhiteLabelRequestsDto(val *WhiteLabelRequestsDto) *NullableWhiteLabelRequestsDto {
	return &NullableWhiteLabelRequestsDto{value: val, isSet: true}
}

func (v NullableWhiteLabelRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhiteLabelRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

