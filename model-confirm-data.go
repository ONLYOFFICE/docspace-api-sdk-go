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

// checks if the ConfirmData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfirmData{}

// ConfirmData The additional confirmation data required for authentication.
type ConfirmData struct {
	// The email address to confirm the user's identity.
	Email NullableString `json:"email,omitempty"`
	// Specifies whether this is the first access to the user's account.
	First NullableBool `json:"first,omitempty"`
	// The unique confirmation key for validating user identity.
	Key NullableString `json:"key,omitempty"`
}

// NewConfirmData instantiates a new ConfirmData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmData() *ConfirmData {
	this := ConfirmData{}
	return &this
}

// NewConfirmDataWithDefaults instantiates a new ConfirmData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmDataWithDefaults() *ConfirmData {
	this := ConfirmData{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfirmData) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfirmData) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *ConfirmData) IsEmailSet() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *ConfirmData) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *ConfirmData) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *ConfirmData) UnsetEmail() {
	o.Email.Unset()
}

// GetFirst returns the First field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfirmData) GetFirst() bool {
	if o == nil || IsNil(o.First.Get()) {
		var ret bool
		return ret
	}
	return *o.First.Get()
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfirmData) GetFirstOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.First.Get(), o.First.IsSet()
}

// HasFirst returns a boolean if a field has been set.
func (o *ConfirmData) IsFirstSet() bool {
	if o != nil && o.First.IsSet() {
		return true
	}

	return false
}

// SetFirst gets a reference to the given NullableBool and assigns it to the First field.
func (o *ConfirmData) SetFirst(v bool) {
	o.First.Set(&v)
}
// SetFirstNil sets the value for First to be an explicit nil
func (o *ConfirmData) SetFirstNil() {
	o.First.Set(nil)
}

// UnsetFirst ensures that no value is present for First, not even an explicit nil
func (o *ConfirmData) UnsetFirst() {
	o.First.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfirmData) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfirmData) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *ConfirmData) IsKeySet() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *ConfirmData) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *ConfirmData) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *ConfirmData) UnsetKey() {
	o.Key.Unset()
}

func (o ConfirmData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfirmData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.First.IsSet() {
		toSerialize["first"] = o.First.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	return toSerialize, nil
}

type NullableConfirmData struct {
	value *ConfirmData
	isSet bool
}

func (v NullableConfirmData) Get() *ConfirmData {
	return v.value
}

func (v *NullableConfirmData) Set(val *ConfirmData) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmData) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmData(val *ConfirmData) *NullableConfirmData {
	return &NullableConfirmData{value: val, isSet: true}
}

func (v NullableConfirmData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfirmData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

