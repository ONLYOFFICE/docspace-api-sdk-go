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

// checks if the CustomFilterParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFilterParameters{}

// CustomFilterParameters The parameters for setting the Custom Filter editing mode.
type CustomFilterParameters struct {
	// Specifies whether the Custom Filter editing mode is enabled or not.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewCustomFilterParameters instantiates a new CustomFilterParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFilterParameters() *CustomFilterParameters {
	this := CustomFilterParameters{}
	return &this
}

// NewCustomFilterParametersWithDefaults instantiates a new CustomFilterParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFilterParametersWithDefaults() *CustomFilterParameters {
	this := CustomFilterParameters{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CustomFilterParameters) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFilterParameters) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CustomFilterParameters) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CustomFilterParameters) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o CustomFilterParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFilterParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableCustomFilterParameters struct {
	value *CustomFilterParameters
	isSet bool
}

func (v NullableCustomFilterParameters) Get() *CustomFilterParameters {
	return v.value
}

func (v *NullableCustomFilterParameters) Set(val *CustomFilterParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFilterParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFilterParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFilterParameters(val *CustomFilterParameters) *NullableCustomFilterParameters {
	return &NullableCustomFilterParameters{value: val, isSet: true}
}

func (v NullableCustomFilterParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFilterParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

