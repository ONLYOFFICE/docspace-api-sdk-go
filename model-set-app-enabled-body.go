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

// checks if the SetAppEnabledBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetAppEnabledBody{}

// SetAppEnabledBody Whether a portal application is switched on.
type SetAppEnabledBody struct {
	// Whether the application is available in this portal. Switching it off leaves its settings document stored, so  switching it back on restores the configuration it had; connected clients are told of the new state without a  reload.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewSetAppEnabledBody instantiates a new SetAppEnabledBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetAppEnabledBody() *SetAppEnabledBody {
	this := SetAppEnabledBody{}
	return &this
}

// NewSetAppEnabledBodyWithDefaults instantiates a new SetAppEnabledBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetAppEnabledBodyWithDefaults() *SetAppEnabledBody {
	this := SetAppEnabledBody{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SetAppEnabledBody) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetAppEnabledBody) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SetAppEnabledBody) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SetAppEnabledBody) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o SetAppEnabledBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetAppEnabledBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableSetAppEnabledBody struct {
	value *SetAppEnabledBody
	isSet bool
}

func (v NullableSetAppEnabledBody) Get() *SetAppEnabledBody {
	return v.value
}

func (v *NullableSetAppEnabledBody) Set(val *SetAppEnabledBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSetAppEnabledBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSetAppEnabledBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetAppEnabledBody(val *SetAppEnabledBody) *NullableSetAppEnabledBody {
	return &NullableSetAppEnabledBody{value: val, isSet: true}
}

func (v NullableSetAppEnabledBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetAppEnabledBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

