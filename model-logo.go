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

// checks if the Logo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Logo{}

// Logo The room logo information.
type Logo struct {
	// The original logo.
	Original NullableString `json:"original"`
	// The large logo.
	Large NullableString `json:"large"`
	// The medium logo.
	Medium NullableString `json:"medium"`
	// The small logo.
	Small NullableString `json:"small"`
	// The logo color.
	Color NullableString `json:"color,omitempty"`
	// The logo cover.
	Cover *LogoCover `json:"cover,omitempty"`
}

type _Logo Logo

// NewLogo instantiates a new Logo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogo(original NullableString, large NullableString, medium NullableString, small NullableString) *Logo {
	this := Logo{}
	this.Original = original
	this.Large = large
	this.Medium = medium
	this.Small = small
	return &this
}

// NewLogoWithDefaults instantiates a new Logo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogoWithDefaults() *Logo {
	this := Logo{}
	return &this
}

// GetOriginal returns the Original field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Logo) GetOriginal() string {
	if o == nil || o.Original.Get() == nil {
		var ret string
		return ret
	}

	return *o.Original.Get()
}

// GetOriginalOk returns a tuple with the Original field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Logo) GetOriginalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Original.Get(), o.Original.IsSet()
}

// SetOriginal sets field value
func (o *Logo) SetOriginal(v string) {
	o.Original.Set(&v)
}

// GetLarge returns the Large field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Logo) GetLarge() string {
	if o == nil || o.Large.Get() == nil {
		var ret string
		return ret
	}

	return *o.Large.Get()
}

// GetLargeOk returns a tuple with the Large field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Logo) GetLargeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Large.Get(), o.Large.IsSet()
}

// SetLarge sets field value
func (o *Logo) SetLarge(v string) {
	o.Large.Set(&v)
}

// GetMedium returns the Medium field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Logo) GetMedium() string {
	if o == nil || o.Medium.Get() == nil {
		var ret string
		return ret
	}

	return *o.Medium.Get()
}

// GetMediumOk returns a tuple with the Medium field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Logo) GetMediumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Medium.Get(), o.Medium.IsSet()
}

// SetMedium sets field value
func (o *Logo) SetMedium(v string) {
	o.Medium.Set(&v)
}

// GetSmall returns the Small field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Logo) GetSmall() string {
	if o == nil || o.Small.Get() == nil {
		var ret string
		return ret
	}

	return *o.Small.Get()
}

// GetSmallOk returns a tuple with the Small field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Logo) GetSmallOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Small.Get(), o.Small.IsSet()
}

// SetSmall sets field value
func (o *Logo) SetSmall(v string) {
	o.Small.Set(&v)
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Logo) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Logo) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *Logo) IsColorSet() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *Logo) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *Logo) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *Logo) UnsetColor() {
	o.Color.Unset()
}

// GetCover returns the Cover field value if set, zero value otherwise.
func (o *Logo) GetCover() LogoCover {
	if o == nil || IsNil(o.Cover) {
		var ret LogoCover
		return ret
	}
	return *o.Cover
}

// GetCoverOk returns a tuple with the Cover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Logo) GetCoverOk() (*LogoCover, bool) {
	if o == nil || IsNil(o.Cover) {
		return nil, false
	}
	return o.Cover, true
}

// HasCover returns a boolean if a field has been set.
func (o *Logo) IsCoverSet() bool {
	if o != nil && !IsNil(o.Cover) {
		return true
	}

	return false
}

// SetCover gets a reference to the given LogoCover and assigns it to the Cover field.
func (o *Logo) SetCover(v LogoCover) {
	o.Cover = &v
}

func (o Logo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Logo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["original"] = o.Original.Get()
	toSerialize["large"] = o.Large.Get()
	toSerialize["medium"] = o.Medium.Get()
	toSerialize["small"] = o.Small.Get()
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	if !IsNil(o.Cover) {
		toSerialize["cover"] = o.Cover
	}
	return toSerialize, nil
}

func (o *Logo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"original",
		"large",
		"medium",
		"small",
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

	varLogo := _Logo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varLogo)

	if err != nil {
		return err
	}

	*o = Logo(varLogo)

	return err
}

type NullableLogo struct {
	value *Logo
	isSet bool
}

func (v NullableLogo) Get() *Logo {
	return v.value
}

func (v *NullableLogo) Set(val *Logo) {
	v.value = val
	v.isSet = true
}

func (v NullableLogo) IsSet() bool {
	return v.isSet
}

func (v *NullableLogo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogo(val *Logo) *NullableLogo {
	return &NullableLogo{value: val, isSet: true}
}

func (v NullableLogo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

