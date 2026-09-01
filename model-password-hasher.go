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

// checks if the PasswordHasher type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordHasher{}

// PasswordHasher The password hash parameters.
type PasswordHasher struct {
	// The password hash size.
	Size *int32 `json:"size,omitempty"`
	// The number of iterations to generate the ppassword hash.
	Iterations *int32 `json:"iterations,omitempty"`
	// The salt to generate the ppassword hash.
	Salt NullableString `json:"salt,omitempty"`
}

// NewPasswordHasher instantiates a new PasswordHasher object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordHasher() *PasswordHasher {
	this := PasswordHasher{}
	return &this
}

// NewPasswordHasherWithDefaults instantiates a new PasswordHasher object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordHasherWithDefaults() *PasswordHasher {
	this := PasswordHasher{}
	return &this
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *PasswordHasher) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordHasher) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *PasswordHasher) IsSizeSet() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *PasswordHasher) SetSize(v int32) {
	o.Size = &v
}

// GetIterations returns the Iterations field value if set, zero value otherwise.
func (o *PasswordHasher) GetIterations() int32 {
	if o == nil || IsNil(o.Iterations) {
		var ret int32
		return ret
	}
	return *o.Iterations
}

// GetIterationsOk returns a tuple with the Iterations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordHasher) GetIterationsOk() (*int32, bool) {
	if o == nil || IsNil(o.Iterations) {
		return nil, false
	}
	return o.Iterations, true
}

// HasIterations returns a boolean if a field has been set.
func (o *PasswordHasher) IsIterationsSet() bool {
	if o != nil && !IsNil(o.Iterations) {
		return true
	}

	return false
}

// SetIterations gets a reference to the given int32 and assigns it to the Iterations field.
func (o *PasswordHasher) SetIterations(v int32) {
	o.Iterations = &v
}

// GetSalt returns the Salt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PasswordHasher) GetSalt() string {
	if o == nil || IsNil(o.Salt.Get()) {
		var ret string
		return ret
	}
	return *o.Salt.Get()
}

// GetSaltOk returns a tuple with the Salt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordHasher) GetSaltOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Salt.Get(), o.Salt.IsSet()
}

// HasSalt returns a boolean if a field has been set.
func (o *PasswordHasher) IsSaltSet() bool {
	if o != nil && o.Salt.IsSet() {
		return true
	}

	return false
}

// SetSalt gets a reference to the given NullableString and assigns it to the Salt field.
func (o *PasswordHasher) SetSalt(v string) {
	o.Salt.Set(&v)
}
// SetSaltNil sets the value for Salt to be an explicit nil
func (o *PasswordHasher) SetSaltNil() {
	o.Salt.Set(nil)
}

// UnsetSalt ensures that no value is present for Salt, not even an explicit nil
func (o *PasswordHasher) UnsetSalt() {
	o.Salt.Unset()
}

func (o PasswordHasher) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordHasher) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Iterations) {
		toSerialize["iterations"] = o.Iterations
	}
	if o.Salt.IsSet() {
		toSerialize["salt"] = o.Salt.Get()
	}
	return toSerialize, nil
}

type NullablePasswordHasher struct {
	value *PasswordHasher
	isSet bool
}

func (v NullablePasswordHasher) Get() *PasswordHasher {
	return v.value
}

func (v *NullablePasswordHasher) Set(val *PasswordHasher) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordHasher) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordHasher) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordHasher(val *PasswordHasher) *NullablePasswordHasher {
	return &NullablePasswordHasher{value: val, isSet: true}
}

func (v NullablePasswordHasher) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordHasher) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

