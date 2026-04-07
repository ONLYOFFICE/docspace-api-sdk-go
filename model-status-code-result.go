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

// checks if the StatusCodeResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusCodeResult{}

// StatusCodeResult struct for StatusCodeResult
type StatusCodeResult struct {
	StatusCode *int32 `json:"statusCode,omitempty"`
}

// NewStatusCodeResult instantiates a new StatusCodeResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusCodeResult() *StatusCodeResult {
	this := StatusCodeResult{}
	return &this
}

// NewStatusCodeResultWithDefaults instantiates a new StatusCodeResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusCodeResultWithDefaults() *StatusCodeResult {
	this := StatusCodeResult{}
	return &this
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *StatusCodeResult) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusCodeResult) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *StatusCodeResult) IsStatusCodeSet() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *StatusCodeResult) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o StatusCodeResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusCodeResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusCode) {
		toSerialize["statusCode"] = o.StatusCode
	}
	return toSerialize, nil
}

type NullableStatusCodeResult struct {
	value *StatusCodeResult
	isSet bool
}

func (v NullableStatusCodeResult) Get() *StatusCodeResult {
	return v.value
}

func (v *NullableStatusCodeResult) Set(val *StatusCodeResult) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusCodeResult) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusCodeResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusCodeResult(val *StatusCodeResult) *NullableStatusCodeResult {
	return &NullableStatusCodeResult{value: val, isSet: true}
}

func (v NullableStatusCodeResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusCodeResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

