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

// checks if the CustomFilterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFilterRequest{}

// CustomFilterRequest The Custom Filter state a spreadsheet is to be put into.
type CustomFilterRequest struct {
	// The state to reach: `true` turns the mode on, so that the sorting and filtering each person applies stays  visible to that person alone, and drops the others out of a running editing session; `false` turns it off and  makes filtering shared again.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewCustomFilterRequest instantiates a new CustomFilterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFilterRequest() *CustomFilterRequest {
	this := CustomFilterRequest{}
	return &this
}

// NewCustomFilterRequestWithDefaults instantiates a new CustomFilterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFilterRequestWithDefaults() *CustomFilterRequest {
	this := CustomFilterRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CustomFilterRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFilterRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CustomFilterRequest) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CustomFilterRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o CustomFilterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFilterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableCustomFilterRequest struct {
	value *CustomFilterRequest
	isSet bool
}

func (v NullableCustomFilterRequest) Get() *CustomFilterRequest {
	return v.value
}

func (v *NullableCustomFilterRequest) Set(val *CustomFilterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFilterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFilterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFilterRequest(val *CustomFilterRequest) *NullableCustomFilterRequest {
	return &NullableCustomFilterRequest{value: val, isSet: true}
}

func (v NullableCustomFilterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFilterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

