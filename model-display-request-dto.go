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

// checks if the DisplayRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisplayRequestDto{}

// DisplayRequestDto The body of a file settings switch that turns something on or makes it visible.
type DisplayRequestDto struct {
	// The state to store for the setting the operation addresses: true enables it or shows what it governs, false  disables or hides it. What exactly is affected, and whether the value belongs to the calling account or to the  whole portal, are stated by the operation that binds this body. The portal may store a different value than  the one sent when another setting overrides it, so read the answer rather than assuming.
	Set *bool `json:"set,omitempty"`
}

// NewDisplayRequestDto instantiates a new DisplayRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisplayRequestDto() *DisplayRequestDto {
	this := DisplayRequestDto{}
	return &this
}

// NewDisplayRequestDtoWithDefaults instantiates a new DisplayRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisplayRequestDtoWithDefaults() *DisplayRequestDto {
	this := DisplayRequestDto{}
	return &this
}

// GetSet returns the Set field value if set, zero value otherwise.
func (o *DisplayRequestDto) GetSet() bool {
	if o == nil || IsNil(o.Set) {
		var ret bool
		return ret
	}
	return *o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisplayRequestDto) GetSetOk() (*bool, bool) {
	if o == nil || IsNil(o.Set) {
		return nil, false
	}
	return o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *DisplayRequestDto) IsSetSet() bool {
	if o != nil && !IsNil(o.Set) {
		return true
	}

	return false
}

// SetSet gets a reference to the given bool and assigns it to the Set field.
func (o *DisplayRequestDto) SetSet(v bool) {
	o.Set = &v
}

func (o DisplayRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisplayRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Set) {
		toSerialize["set"] = o.Set
	}
	return toSerialize, nil
}

type NullableDisplayRequestDto struct {
	value *DisplayRequestDto
	isSet bool
}

func (v NullableDisplayRequestDto) Get() *DisplayRequestDto {
	return v.value
}

func (v *NullableDisplayRequestDto) Set(val *DisplayRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDisplayRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDisplayRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisplayRequestDto(val *DisplayRequestDto) *NullableDisplayRequestDto {
	return &NullableDisplayRequestDto{value: val, isSet: true}
}

func (v NullableDisplayRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisplayRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

