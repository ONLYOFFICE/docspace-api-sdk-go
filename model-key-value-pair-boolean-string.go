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

// checks if the KeyValuePairBooleanString type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyValuePairBooleanString{}

// KeyValuePairBooleanString struct for KeyValuePairBooleanString
type KeyValuePairBooleanString struct {
	Key *bool `json:"key,omitempty"`
	Value NullableString `json:"value,omitempty"`
}

// NewKeyValuePairBooleanString instantiates a new KeyValuePairBooleanString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyValuePairBooleanString() *KeyValuePairBooleanString {
	this := KeyValuePairBooleanString{}
	return &this
}

// NewKeyValuePairBooleanStringWithDefaults instantiates a new KeyValuePairBooleanString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyValuePairBooleanStringWithDefaults() *KeyValuePairBooleanString {
	this := KeyValuePairBooleanString{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *KeyValuePairBooleanString) GetKey() bool {
	if o == nil || IsNil(o.Key) {
		var ret bool
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyValuePairBooleanString) GetKeyOk() (*bool, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *KeyValuePairBooleanString) IsKeySet() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given bool and assigns it to the Key field.
func (o *KeyValuePairBooleanString) SetKey(v bool) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeyValuePairBooleanString) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeyValuePairBooleanString) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *KeyValuePairBooleanString) IsValueSet() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *KeyValuePairBooleanString) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *KeyValuePairBooleanString) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *KeyValuePairBooleanString) UnsetValue() {
	o.Value.Unset()
}

func (o KeyValuePairBooleanString) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyValuePairBooleanString) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableKeyValuePairBooleanString struct {
	value *KeyValuePairBooleanString
	isSet bool
}

func (v NullableKeyValuePairBooleanString) Get() *KeyValuePairBooleanString {
	return v.value
}

func (v *NullableKeyValuePairBooleanString) Set(val *KeyValuePairBooleanString) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyValuePairBooleanString) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyValuePairBooleanString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyValuePairBooleanString(val *KeyValuePairBooleanString) *NullableKeyValuePairBooleanString {
	return &NullableKeyValuePairBooleanString{value: val, isSet: true}
}

func (v NullableKeyValuePairBooleanString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyValuePairBooleanString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

