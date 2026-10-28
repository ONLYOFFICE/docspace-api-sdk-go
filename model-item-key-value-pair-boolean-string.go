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

// checks if the ItemKeyValuePairBooleanString type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemKeyValuePairBooleanString{}

// ItemKeyValuePairBooleanString One entry of a keyed collection, carried as an explicit pair of `key` and `value` fields instead of as a member  of a JSON object, so that the key is not restricted to a string and the entries keep the order they are sent in.
type ItemKeyValuePairBooleanString struct {
	// The left half of the pair. Where the pair configures something, this is the identifier the value belongs to -  a setting name, a module id, a logo slot; where the pair reports the result of a call, this is the result  itself, such as the flag telling whether the call succeeded. Which of the two it is, and which keys are  accepted, is stated by the operation that sends or returns the pair.
	Key *bool `json:"key,omitempty"`
	// The right half of the pair: what is assigned to the key next to it, or what is reported for it. Its meaning  and its accepted values follow from the key, so read them from the operation that sends or returns the pair.
	Value NullableString `json:"value,omitempty"`
}

// NewItemKeyValuePairBooleanString instantiates a new ItemKeyValuePairBooleanString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemKeyValuePairBooleanString() *ItemKeyValuePairBooleanString {
	this := ItemKeyValuePairBooleanString{}
	return &this
}

// NewItemKeyValuePairBooleanStringWithDefaults instantiates a new ItemKeyValuePairBooleanString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemKeyValuePairBooleanStringWithDefaults() *ItemKeyValuePairBooleanString {
	this := ItemKeyValuePairBooleanString{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ItemKeyValuePairBooleanString) GetKey() bool {
	if o == nil || IsNil(o.Key) {
		var ret bool
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemKeyValuePairBooleanString) GetKeyOk() (*bool, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ItemKeyValuePairBooleanString) IsKeySet() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given bool and assigns it to the Key field.
func (o *ItemKeyValuePairBooleanString) SetKey(v bool) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemKeyValuePairBooleanString) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemKeyValuePairBooleanString) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ItemKeyValuePairBooleanString) IsValueSet() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *ItemKeyValuePairBooleanString) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *ItemKeyValuePairBooleanString) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ItemKeyValuePairBooleanString) UnsetValue() {
	o.Value.Unset()
}

func (o ItemKeyValuePairBooleanString) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemKeyValuePairBooleanString) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableItemKeyValuePairBooleanString struct {
	value *ItemKeyValuePairBooleanString
	isSet bool
}

func (v NullableItemKeyValuePairBooleanString) Get() *ItemKeyValuePairBooleanString {
	return v.value
}

func (v *NullableItemKeyValuePairBooleanString) Set(val *ItemKeyValuePairBooleanString) {
	v.value = val
	v.isSet = true
}

func (v NullableItemKeyValuePairBooleanString) IsSet() bool {
	return v.isSet
}

func (v *NullableItemKeyValuePairBooleanString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemKeyValuePairBooleanString(val *ItemKeyValuePairBooleanString) *NullableItemKeyValuePairBooleanString {
	return &NullableItemKeyValuePairBooleanString{value: val, isSet: true}
}

func (v NullableItemKeyValuePairBooleanString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemKeyValuePairBooleanString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

