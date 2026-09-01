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

// checks if the ConnectionTestResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionTestResult{}

// ConnectionTestResult The outcome of a connection test against an external database.
type ConnectionTestResult struct {
	// Specifies whether the connection to the database succeeded.
	Success *bool `json:"success,omitempty"`
	// The reason the connection failed, or null when it succeeded.
	Error NullableString `json:"error,omitempty"`
}

// NewConnectionTestResult instantiates a new ConnectionTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionTestResult() *ConnectionTestResult {
	this := ConnectionTestResult{}
	return &this
}

// NewConnectionTestResultWithDefaults instantiates a new ConnectionTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionTestResultWithDefaults() *ConnectionTestResult {
	this := ConnectionTestResult{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ConnectionTestResult) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionTestResult) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ConnectionTestResult) IsSuccessSet() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ConnectionTestResult) SetSuccess(v bool) {
	o.Success = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionTestResult) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionTestResult) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *ConnectionTestResult) IsErrorSet() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *ConnectionTestResult) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *ConnectionTestResult) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *ConnectionTestResult) UnsetError() {
	o.Error.Unset()
}

func (o ConnectionTestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionTestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	return toSerialize, nil
}

type NullableConnectionTestResult struct {
	value *ConnectionTestResult
	isSet bool
}

func (v NullableConnectionTestResult) Get() *ConnectionTestResult {
	return v.value
}

func (v *NullableConnectionTestResult) Set(val *ConnectionTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionTestResult(val *ConnectionTestResult) *NullableConnectionTestResult {
	return &NullableConnectionTestResult{value: val, isSet: true}
}

func (v NullableConnectionTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

