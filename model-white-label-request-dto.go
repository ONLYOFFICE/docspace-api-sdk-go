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

// checks if the WhiteLabelRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WhiteLabelRequestDto{}

// WhiteLabelRequestDto The branding a portal is given: the wordmark, the logo images, or both.
type WhiteLabelRequestDto struct {
	// The wordmark printed next to or instead of a logo image, on the login page, in the editors and in  notification letters. An empty or blank value, and the built-in `ONLYOFFICE` itself, clear the setting rather  than store it. The text is not rendered into the logo images, which carry their own wordmark.
	LogoText NullableString `json:"logoText,omitempty"`
	// The logo images to store, each entry naming a logo slot in its `key` - the numeric `type` published by  `GET api/2.0/settings/whitelabel/logos` - and carrying the two theme images in its value. A slot left out of  the list keeps the image it has, so this is a partial update rather than a replacement of the whole branding.  Saving the login-page slot also rebuilds the notification logo from it.
	Logo []ItemKeyValuePairStringLogoRequestDto `json:"logo,omitempty"`
}

// NewWhiteLabelRequestDto instantiates a new WhiteLabelRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhiteLabelRequestDto() *WhiteLabelRequestDto {
	this := WhiteLabelRequestDto{}
	return &this
}

// NewWhiteLabelRequestDtoWithDefaults instantiates a new WhiteLabelRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhiteLabelRequestDtoWithDefaults() *WhiteLabelRequestDto {
	this := WhiteLabelRequestDto{}
	return &this
}

// GetLogoText returns the LogoText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WhiteLabelRequestDto) GetLogoText() string {
	if o == nil || IsNil(o.LogoText.Get()) {
		var ret string
		return ret
	}
	return *o.LogoText.Get()
}

// GetLogoTextOk returns a tuple with the LogoText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WhiteLabelRequestDto) GetLogoTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoText.Get(), o.LogoText.IsSet()
}

// HasLogoText returns a boolean if a field has been set.
func (o *WhiteLabelRequestDto) IsLogoTextSet() bool {
	if o != nil && o.LogoText.IsSet() {
		return true
	}

	return false
}

// SetLogoText gets a reference to the given NullableString and assigns it to the LogoText field.
func (o *WhiteLabelRequestDto) SetLogoText(v string) {
	o.LogoText.Set(&v)
}
// SetLogoTextNil sets the value for LogoText to be an explicit nil
func (o *WhiteLabelRequestDto) SetLogoTextNil() {
	o.LogoText.Set(nil)
}

// UnsetLogoText ensures that no value is present for LogoText, not even an explicit nil
func (o *WhiteLabelRequestDto) UnsetLogoText() {
	o.LogoText.Unset()
}

// GetLogo returns the Logo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WhiteLabelRequestDto) GetLogo() []ItemKeyValuePairStringLogoRequestDto {
	if o == nil {
		var ret []ItemKeyValuePairStringLogoRequestDto
		return ret
	}
	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WhiteLabelRequestDto) GetLogoOk() ([]ItemKeyValuePairStringLogoRequestDto, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *WhiteLabelRequestDto) IsLogoSet() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given []ItemKeyValuePairStringLogoRequestDto and assigns it to the Logo field.
func (o *WhiteLabelRequestDto) SetLogo(v []ItemKeyValuePairStringLogoRequestDto) {
	o.Logo = v
}

func (o WhiteLabelRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WhiteLabelRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LogoText.IsSet() {
		toSerialize["logoText"] = o.LogoText.Get()
	}
	if o.Logo != nil {
		toSerialize["logo"] = o.Logo
	}
	return toSerialize, nil
}

type NullableWhiteLabelRequestDto struct {
	value *WhiteLabelRequestDto
	isSet bool
}

func (v NullableWhiteLabelRequestDto) Get() *WhiteLabelRequestDto {
	return v.value
}

func (v *NullableWhiteLabelRequestDto) Set(val *WhiteLabelRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWhiteLabelRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWhiteLabelRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhiteLabelRequestDto(val *WhiteLabelRequestDto) *NullableWhiteLabelRequestDto {
	return &NullableWhiteLabelRequestDto{value: val, isSet: true}
}

func (v NullableWhiteLabelRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhiteLabelRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

