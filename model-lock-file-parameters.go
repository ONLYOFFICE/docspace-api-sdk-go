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

// checks if the LockFileParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LockFileParameters{}

// LockFileParameters The lock state a file is to be put into.
type LockFileParameters struct {
	// The state to reach: `true` locks the file, which blocks editing, renaming and deleting for everybody but the  account that locked it and the room admins, and drops the others out of a running editing session; `false`  releases the lock.
	LockFile *bool `json:"lockFile,omitempty"`
}

// NewLockFileParameters instantiates a new LockFileParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLockFileParameters() *LockFileParameters {
	this := LockFileParameters{}
	return &this
}

// NewLockFileParametersWithDefaults instantiates a new LockFileParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLockFileParametersWithDefaults() *LockFileParameters {
	this := LockFileParameters{}
	return &this
}

// GetLockFile returns the LockFile field value if set, zero value otherwise.
func (o *LockFileParameters) GetLockFile() bool {
	if o == nil || IsNil(o.LockFile) {
		var ret bool
		return ret
	}
	return *o.LockFile
}

// GetLockFileOk returns a tuple with the LockFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockFileParameters) GetLockFileOk() (*bool, bool) {
	if o == nil || IsNil(o.LockFile) {
		return nil, false
	}
	return o.LockFile, true
}

// HasLockFile returns a boolean if a field has been set.
func (o *LockFileParameters) IsLockFileSet() bool {
	if o != nil && !IsNil(o.LockFile) {
		return true
	}

	return false
}

// SetLockFile gets a reference to the given bool and assigns it to the LockFile field.
func (o *LockFileParameters) SetLockFile(v bool) {
	o.LockFile = &v
}

func (o LockFileParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LockFileParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LockFile) {
		toSerialize["lockFile"] = o.LockFile
	}
	return toSerialize, nil
}

type NullableLockFileParameters struct {
	value *LockFileParameters
	isSet bool
}

func (v NullableLockFileParameters) Get() *LockFileParameters {
	return v.value
}

func (v *NullableLockFileParameters) Set(val *LockFileParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableLockFileParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableLockFileParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLockFileParameters(val *LockFileParameters) *NullableLockFileParameters {
	return &NullableLockFileParameters{value: val, isSet: true}
}

func (v NullableLockFileParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLockFileParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

