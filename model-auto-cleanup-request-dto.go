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

// checks if the AutoCleanupRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoCleanupRequestDto{}

// AutoCleanupRequestDto The request parameters for updating the trash bin auto-clearing setting.
type AutoCleanupRequestDto struct {
	// Specifies whether to enable the auto-clearing or not.
	Set *bool `json:"set,omitempty"`
	Gap *DateToAutoCleanUp `json:"gap,omitempty"`
}

// NewAutoCleanupRequestDto instantiates a new AutoCleanupRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoCleanupRequestDto() *AutoCleanupRequestDto {
	this := AutoCleanupRequestDto{}
	return &this
}

// NewAutoCleanupRequestDtoWithDefaults instantiates a new AutoCleanupRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoCleanupRequestDtoWithDefaults() *AutoCleanupRequestDto {
	this := AutoCleanupRequestDto{}
	return &this
}

// GetSet returns the Set field value if set, zero value otherwise.
func (o *AutoCleanupRequestDto) GetSet() bool {
	if o == nil || IsNil(o.Set) {
		var ret bool
		return ret
	}
	return *o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoCleanupRequestDto) GetSetOk() (*bool, bool) {
	if o == nil || IsNil(o.Set) {
		return nil, false
	}
	return o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *AutoCleanupRequestDto) IsSetSet() bool {
	if o != nil && !IsNil(o.Set) {
		return true
	}

	return false
}

// SetSet gets a reference to the given bool and assigns it to the Set field.
func (o *AutoCleanupRequestDto) SetSet(v bool) {
	o.Set = &v
}

// GetGap returns the Gap field value if set, zero value otherwise.
func (o *AutoCleanupRequestDto) GetGap() DateToAutoCleanUp {
	if o == nil || IsNil(o.Gap) {
		var ret DateToAutoCleanUp
		return ret
	}
	return *o.Gap
}

// GetGapOk returns a tuple with the Gap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoCleanupRequestDto) GetGapOk() (*DateToAutoCleanUp, bool) {
	if o == nil || IsNil(o.Gap) {
		return nil, false
	}
	return o.Gap, true
}

// HasGap returns a boolean if a field has been set.
func (o *AutoCleanupRequestDto) IsGapSet() bool {
	if o != nil && !IsNil(o.Gap) {
		return true
	}

	return false
}

// SetGap gets a reference to the given DateToAutoCleanUp and assigns it to the Gap field.
func (o *AutoCleanupRequestDto) SetGap(v DateToAutoCleanUp) {
	o.Gap = &v
}

func (o AutoCleanupRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoCleanupRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Set) {
		toSerialize["set"] = o.Set
	}
	if !IsNil(o.Gap) {
		toSerialize["gap"] = o.Gap
	}
	return toSerialize, nil
}

type NullableAutoCleanupRequestDto struct {
	value *AutoCleanupRequestDto
	isSet bool
}

func (v NullableAutoCleanupRequestDto) Get() *AutoCleanupRequestDto {
	return v.value
}

func (v *NullableAutoCleanupRequestDto) Set(val *AutoCleanupRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoCleanupRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoCleanupRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoCleanupRequestDto(val *AutoCleanupRequestDto) *NullableAutoCleanupRequestDto {
	return &NullableAutoCleanupRequestDto{value: val, isSet: true}
}

func (v NullableAutoCleanupRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoCleanupRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

