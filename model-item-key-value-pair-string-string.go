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

// checks if the ItemKeyValuePairStringString type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemKeyValuePairStringString{}

// ItemKeyValuePairStringString One entry of a keyed collection, carried as an explicit pair of `key` and `value` fields instead of as a member  of a JSON object, so that the key is not restricted to a string and the entries keep the order they are sent in.
type ItemKeyValuePairStringString struct {
	// The left half of the pair. Where the pair configures something, this is the identifier the value belongs to -  a setting name, a module id, a logo slot; where the pair reports the result of a call, this is the result  itself, such as the flag telling whether the call succeeded. Which of the two it is, and which keys are  accepted, is stated by the operation that sends or returns the pair.
	Key NullableString `json:"key,omitempty"`
	// The right half of the pair: what is assigned to the key next to it, or what is reported for it. Its meaning  and its accepted values follow from the key, so read them from the operation that sends or returns the pair.
	Value NullableString `json:"value,omitempty"`
}

// NewItemKeyValuePairStringString instantiates a new ItemKeyValuePairStringString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemKeyValuePairStringString() *ItemKeyValuePairStringString {
	this := ItemKeyValuePairStringString{}
	return &this
}

// NewItemKeyValuePairStringStringWithDefaults instantiates a new ItemKeyValuePairStringString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemKeyValuePairStringStringWithDefaults() *ItemKeyValuePairStringString {
	this := ItemKeyValuePairStringString{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemKeyValuePairStringString) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemKeyValuePairStringString) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *ItemKeyValuePairStringString) IsKeySet() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *ItemKeyValuePairStringString) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *ItemKeyValuePairStringString) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *ItemKeyValuePairStringString) UnsetKey() {
	o.Key.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemKeyValuePairStringString) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemKeyValuePairStringString) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ItemKeyValuePairStringString) IsValueSet() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *ItemKeyValuePairStringString) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *ItemKeyValuePairStringString) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ItemKeyValuePairStringString) UnsetValue() {
	o.Value.Unset()
}

func (o ItemKeyValuePairStringString) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemKeyValuePairStringString) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableItemKeyValuePairStringString struct {
	value *ItemKeyValuePairStringString
	isSet bool
}

func (v NullableItemKeyValuePairStringString) Get() *ItemKeyValuePairStringString {
	return v.value
}

func (v *NullableItemKeyValuePairStringString) Set(val *ItemKeyValuePairStringString) {
	v.value = val
	v.isSet = true
}

func (v NullableItemKeyValuePairStringString) IsSet() bool {
	return v.isSet
}

func (v *NullableItemKeyValuePairStringString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemKeyValuePairStringString(val *ItemKeyValuePairStringString) *NullableItemKeyValuePairStringString {
	return &NullableItemKeyValuePairStringString{value: val, isSet: true}
}

func (v NullableItemKeyValuePairStringString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemKeyValuePairStringString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

