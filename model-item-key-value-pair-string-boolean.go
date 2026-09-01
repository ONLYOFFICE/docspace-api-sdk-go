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

// checks if the ItemKeyValuePairStringBoolean type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemKeyValuePairStringBoolean{}

// ItemKeyValuePairStringBoolean A key-value pair of a list item.
type ItemKeyValuePairStringBoolean struct {
	// The key that identifies the item within the list.
	Key NullableString `json:"key,omitempty"`
	// The value associated with the key.
	Value *bool `json:"value,omitempty"`
}

// NewItemKeyValuePairStringBoolean instantiates a new ItemKeyValuePairStringBoolean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemKeyValuePairStringBoolean() *ItemKeyValuePairStringBoolean {
	this := ItemKeyValuePairStringBoolean{}
	return &this
}

// NewItemKeyValuePairStringBooleanWithDefaults instantiates a new ItemKeyValuePairStringBoolean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemKeyValuePairStringBooleanWithDefaults() *ItemKeyValuePairStringBoolean {
	this := ItemKeyValuePairStringBoolean{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemKeyValuePairStringBoolean) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemKeyValuePairStringBoolean) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *ItemKeyValuePairStringBoolean) IsKeySet() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *ItemKeyValuePairStringBoolean) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *ItemKeyValuePairStringBoolean) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *ItemKeyValuePairStringBoolean) UnsetKey() {
	o.Key.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ItemKeyValuePairStringBoolean) GetValue() bool {
	if o == nil || IsNil(o.Value) {
		var ret bool
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemKeyValuePairStringBoolean) GetValueOk() (*bool, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ItemKeyValuePairStringBoolean) IsValueSet() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given bool and assigns it to the Value field.
func (o *ItemKeyValuePairStringBoolean) SetValue(v bool) {
	o.Value = &v
}

func (o ItemKeyValuePairStringBoolean) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemKeyValuePairStringBoolean) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableItemKeyValuePairStringBoolean struct {
	value *ItemKeyValuePairStringBoolean
	isSet bool
}

func (v NullableItemKeyValuePairStringBoolean) Get() *ItemKeyValuePairStringBoolean {
	return v.value
}

func (v *NullableItemKeyValuePairStringBoolean) Set(val *ItemKeyValuePairStringBoolean) {
	v.value = val
	v.isSet = true
}

func (v NullableItemKeyValuePairStringBoolean) IsSet() bool {
	return v.isSet
}

func (v *NullableItemKeyValuePairStringBoolean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemKeyValuePairStringBoolean(val *ItemKeyValuePairStringBoolean) *NullableItemKeyValuePairStringBoolean {
	return &NullableItemKeyValuePairStringBoolean{value: val, isSet: true}
}

func (v NullableItemKeyValuePairStringBoolean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemKeyValuePairStringBoolean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

