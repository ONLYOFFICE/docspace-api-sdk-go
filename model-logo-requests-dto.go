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

// checks if the LogoRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogoRequestsDto{}

// LogoRequestsDto The two theme variants of one branding logo.
type LogoRequestsDto struct {
	// The image used on a light background, either as a `data:image/png;base64,...` payload - `png`, `jpg` and  `svg` are accepted - or as the name of a file already put in the temporary store.
	Light NullableString `json:"light,omitempty"`
	// The image used on a dark background, in the same two forms as `light`. It is only stored for the slots that  have a dark variant and is ignored for the favicon and the editor logos.
	Dark NullableString `json:"dark,omitempty"`
}

// NewLogoRequestsDto instantiates a new LogoRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogoRequestsDto() *LogoRequestsDto {
	this := LogoRequestsDto{}
	return &this
}

// NewLogoRequestsDtoWithDefaults instantiates a new LogoRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogoRequestsDtoWithDefaults() *LogoRequestsDto {
	this := LogoRequestsDto{}
	return &this
}

// GetLight returns the Light field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogoRequestsDto) GetLight() string {
	if o == nil || IsNil(o.Light.Get()) {
		var ret string
		return ret
	}
	return *o.Light.Get()
}

// GetLightOk returns a tuple with the Light field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogoRequestsDto) GetLightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Light.Get(), o.Light.IsSet()
}

// HasLight returns a boolean if a field has been set.
func (o *LogoRequestsDto) IsLightSet() bool {
	if o != nil && o.Light.IsSet() {
		return true
	}

	return false
}

// SetLight gets a reference to the given NullableString and assigns it to the Light field.
func (o *LogoRequestsDto) SetLight(v string) {
	o.Light.Set(&v)
}
// SetLightNil sets the value for Light to be an explicit nil
func (o *LogoRequestsDto) SetLightNil() {
	o.Light.Set(nil)
}

// UnsetLight ensures that no value is present for Light, not even an explicit nil
func (o *LogoRequestsDto) UnsetLight() {
	o.Light.Unset()
}

// GetDark returns the Dark field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogoRequestsDto) GetDark() string {
	if o == nil || IsNil(o.Dark.Get()) {
		var ret string
		return ret
	}
	return *o.Dark.Get()
}

// GetDarkOk returns a tuple with the Dark field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogoRequestsDto) GetDarkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dark.Get(), o.Dark.IsSet()
}

// HasDark returns a boolean if a field has been set.
func (o *LogoRequestsDto) IsDarkSet() bool {
	if o != nil && o.Dark.IsSet() {
		return true
	}

	return false
}

// SetDark gets a reference to the given NullableString and assigns it to the Dark field.
func (o *LogoRequestsDto) SetDark(v string) {
	o.Dark.Set(&v)
}
// SetDarkNil sets the value for Dark to be an explicit nil
func (o *LogoRequestsDto) SetDarkNil() {
	o.Dark.Set(nil)
}

// UnsetDark ensures that no value is present for Dark, not even an explicit nil
func (o *LogoRequestsDto) UnsetDark() {
	o.Dark.Unset()
}

func (o LogoRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogoRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Light.IsSet() {
		toSerialize["light"] = o.Light.Get()
	}
	if o.Dark.IsSet() {
		toSerialize["dark"] = o.Dark.Get()
	}
	return toSerialize, nil
}

type NullableLogoRequestsDto struct {
	value *LogoRequestsDto
	isSet bool
}

func (v NullableLogoRequestsDto) Get() *LogoRequestsDto {
	return v.value
}

func (v *NullableLogoRequestsDto) Set(val *LogoRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableLogoRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableLogoRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogoRequestsDto(val *LogoRequestsDto) *NullableLogoRequestsDto {
	return &NullableLogoRequestsDto{value: val, isSet: true}
}

func (v NullableLogoRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogoRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

