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

// checks if the CoEditingConfigDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoEditingConfigDto{}

// CoEditingConfigDto How co-editing is preset when the document opens, and whether the user may switch it afterwards.
type CoEditingConfigDto struct {
	// Whether the user may switch between the two co-editing modes from the editor interface, or is held to the one  the portal preset.
	Change *bool `json:"change,omitempty"`
	// Whether other participants see each change as it is typed. Left off, changes are exchanged only when a  participant saves, and the paragraph being edited is locked for the others meanwhile.
	Fast *bool `json:"fast,omitempty"`
	// The mode the two settings above amount to, as the editors name it.
	Mode *CoEditingConfigMode `json:"mode,omitempty"`
}

// NewCoEditingConfigDto instantiates a new CoEditingConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoEditingConfigDto() *CoEditingConfigDto {
	this := CoEditingConfigDto{}
	return &this
}

// NewCoEditingConfigDtoWithDefaults instantiates a new CoEditingConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoEditingConfigDtoWithDefaults() *CoEditingConfigDto {
	this := CoEditingConfigDto{}
	return &this
}

// GetChange returns the Change field value if set, zero value otherwise.
func (o *CoEditingConfigDto) GetChange() bool {
	if o == nil || IsNil(o.Change) {
		var ret bool
		return ret
	}
	return *o.Change
}

// GetChangeOk returns a tuple with the Change field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoEditingConfigDto) GetChangeOk() (*bool, bool) {
	if o == nil || IsNil(o.Change) {
		return nil, false
	}
	return o.Change, true
}

// HasChange returns a boolean if a field has been set.
func (o *CoEditingConfigDto) IsChangeSet() bool {
	if o != nil && !IsNil(o.Change) {
		return true
	}

	return false
}

// SetChange gets a reference to the given bool and assigns it to the Change field.
func (o *CoEditingConfigDto) SetChange(v bool) {
	o.Change = &v
}

// GetFast returns the Fast field value if set, zero value otherwise.
func (o *CoEditingConfigDto) GetFast() bool {
	if o == nil || IsNil(o.Fast) {
		var ret bool
		return ret
	}
	return *o.Fast
}

// GetFastOk returns a tuple with the Fast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoEditingConfigDto) GetFastOk() (*bool, bool) {
	if o == nil || IsNil(o.Fast) {
		return nil, false
	}
	return o.Fast, true
}

// HasFast returns a boolean if a field has been set.
func (o *CoEditingConfigDto) IsFastSet() bool {
	if o != nil && !IsNil(o.Fast) {
		return true
	}

	return false
}

// SetFast gets a reference to the given bool and assigns it to the Fast field.
func (o *CoEditingConfigDto) SetFast(v bool) {
	o.Fast = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *CoEditingConfigDto) GetMode() CoEditingConfigMode {
	if o == nil || IsNil(o.Mode) {
		var ret CoEditingConfigMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoEditingConfigDto) GetModeOk() (*CoEditingConfigMode, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *CoEditingConfigDto) IsModeSet() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given CoEditingConfigMode and assigns it to the Mode field.
func (o *CoEditingConfigDto) SetMode(v CoEditingConfigMode) {
	o.Mode = &v
}

func (o CoEditingConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoEditingConfigDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Change) {
		toSerialize["change"] = o.Change
	}
	if !IsNil(o.Fast) {
		toSerialize["fast"] = o.Fast
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableCoEditingConfigDto struct {
	value *CoEditingConfigDto
	isSet bool
}

func (v NullableCoEditingConfigDto) Get() *CoEditingConfigDto {
	return v.value
}

func (v *NullableCoEditingConfigDto) Set(val *CoEditingConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCoEditingConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCoEditingConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoEditingConfigDto(val *CoEditingConfigDto) *NullableCoEditingConfigDto {
	return &NullableCoEditingConfigDto{value: val, isSet: true}
}

func (v NullableCoEditingConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoEditingConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

