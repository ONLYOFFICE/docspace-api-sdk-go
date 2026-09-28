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

// checks if the SetAppSettingsBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetAppSettingsBody{}

// SetAppSettingsBody The configuration document a portal application keeps.
type SetAppSettingsBody struct {
	Settings interface{} `json:"settings,omitempty"`
}

// NewSetAppSettingsBody instantiates a new SetAppSettingsBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetAppSettingsBody() *SetAppSettingsBody {
	this := SetAppSettingsBody{}
	return &this
}

// NewSetAppSettingsBodyWithDefaults instantiates a new SetAppSettingsBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetAppSettingsBodyWithDefaults() *SetAppSettingsBody {
	this := SetAppSettingsBody{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetAppSettingsBody) GetSettings() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetAppSettingsBody) GetSettingsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return &o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *SetAppSettingsBody) IsSettingsSet() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given interface{} and assigns it to the Settings field.
func (o *SetAppSettingsBody) SetSettings(v interface{}) {
	o.Settings = v
}

func (o SetAppSettingsBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetAppSettingsBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	return toSerialize, nil
}

type NullableSetAppSettingsBody struct {
	value *SetAppSettingsBody
	isSet bool
}

func (v NullableSetAppSettingsBody) Get() *SetAppSettingsBody {
	return v.value
}

func (v *NullableSetAppSettingsBody) Set(val *SetAppSettingsBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSetAppSettingsBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSetAppSettingsBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetAppSettingsBody(val *SetAppSettingsBody) *NullableSetAppSettingsBody {
	return &NullableSetAppSettingsBody{value: val, isSet: true}
}

func (v NullableSetAppSettingsBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetAppSettingsBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

