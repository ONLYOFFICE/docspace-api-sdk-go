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

// checks if the AutoCleanUpDataDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoCleanUpDataDto{}

// AutoCleanUpDataDto The trash auto-clearing setting of an account.
type AutoCleanUpDataDto struct {
	// Whether the trash of the account is cleared automatically. While it is false nothing is removed by the portal  and the interval below is kept but unused.
	IsAutoCleanUp *bool `json:"isAutoCleanUp,omitempty"`
	// How long an item may stay in the trash before it is removed for good. It is reported even while clearing is  off, and it is what the moment in the `autoDelete` field of a trashed entry is computed from.
	Gap *DateToAutoCleanUp `json:"gap,omitempty"`
}

// NewAutoCleanUpDataDto instantiates a new AutoCleanUpDataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoCleanUpDataDto() *AutoCleanUpDataDto {
	this := AutoCleanUpDataDto{}
	return &this
}

// NewAutoCleanUpDataDtoWithDefaults instantiates a new AutoCleanUpDataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoCleanUpDataDtoWithDefaults() *AutoCleanUpDataDto {
	this := AutoCleanUpDataDto{}
	return &this
}

// GetIsAutoCleanUp returns the IsAutoCleanUp field value if set, zero value otherwise.
func (o *AutoCleanUpDataDto) GetIsAutoCleanUp() bool {
	if o == nil || IsNil(o.IsAutoCleanUp) {
		var ret bool
		return ret
	}
	return *o.IsAutoCleanUp
}

// GetIsAutoCleanUpOk returns a tuple with the IsAutoCleanUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoCleanUpDataDto) GetIsAutoCleanUpOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAutoCleanUp) {
		return nil, false
	}
	return o.IsAutoCleanUp, true
}

// HasIsAutoCleanUp returns a boolean if a field has been set.
func (o *AutoCleanUpDataDto) IsIsAutoCleanUpSet() bool {
	if o != nil && !IsNil(o.IsAutoCleanUp) {
		return true
	}

	return false
}

// SetIsAutoCleanUp gets a reference to the given bool and assigns it to the IsAutoCleanUp field.
func (o *AutoCleanUpDataDto) SetIsAutoCleanUp(v bool) {
	o.IsAutoCleanUp = &v
}

// GetGap returns the Gap field value if set, zero value otherwise.
func (o *AutoCleanUpDataDto) GetGap() DateToAutoCleanUp {
	if o == nil || IsNil(o.Gap) {
		var ret DateToAutoCleanUp
		return ret
	}
	return *o.Gap
}

// GetGapOk returns a tuple with the Gap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoCleanUpDataDto) GetGapOk() (*DateToAutoCleanUp, bool) {
	if o == nil || IsNil(o.Gap) {
		return nil, false
	}
	return o.Gap, true
}

// HasGap returns a boolean if a field has been set.
func (o *AutoCleanUpDataDto) IsGapSet() bool {
	if o != nil && !IsNil(o.Gap) {
		return true
	}

	return false
}

// SetGap gets a reference to the given DateToAutoCleanUp and assigns it to the Gap field.
func (o *AutoCleanUpDataDto) SetGap(v DateToAutoCleanUp) {
	o.Gap = &v
}

func (o AutoCleanUpDataDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoCleanUpDataDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAutoCleanUp) {
		toSerialize["isAutoCleanUp"] = o.IsAutoCleanUp
	}
	if !IsNil(o.Gap) {
		toSerialize["gap"] = o.Gap
	}
	return toSerialize, nil
}

type NullableAutoCleanUpDataDto struct {
	value *AutoCleanUpDataDto
	isSet bool
}

func (v NullableAutoCleanUpDataDto) Get() *AutoCleanUpDataDto {
	return v.value
}

func (v *NullableAutoCleanUpDataDto) Set(val *AutoCleanUpDataDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoCleanUpDataDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoCleanUpDataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoCleanUpDataDto(val *AutoCleanUpDataDto) *NullableAutoCleanUpDataDto {
	return &NullableAutoCleanUpDataDto{value: val, isSet: true}
}

func (v NullableAutoCleanUpDataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoCleanUpDataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

