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

// checks if the ItemKeyValuePairObjectObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemKeyValuePairObjectObject{}

// ItemKeyValuePairObjectObject struct for ItemKeyValuePairObjectObject
type ItemKeyValuePairObjectObject struct {
	Key interface{} `json:"key,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// NewItemKeyValuePairObjectObject instantiates a new ItemKeyValuePairObjectObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemKeyValuePairObjectObject() *ItemKeyValuePairObjectObject {
	this := ItemKeyValuePairObjectObject{}
	return &this
}

// NewItemKeyValuePairObjectObjectWithDefaults instantiates a new ItemKeyValuePairObjectObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemKeyValuePairObjectObjectWithDefaults() *ItemKeyValuePairObjectObject {
	this := ItemKeyValuePairObjectObject{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemKeyValuePairObjectObject) GetKey() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemKeyValuePairObjectObject) GetKeyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return &o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ItemKeyValuePairObjectObject) IsKeySet() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given interface{} and assigns it to the Key field.
func (o *ItemKeyValuePairObjectObject) SetKey(v interface{}) {
	o.Key = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemKeyValuePairObjectObject) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemKeyValuePairObjectObject) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ItemKeyValuePairObjectObject) IsValueSet() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *ItemKeyValuePairObjectObject) SetValue(v interface{}) {
	o.Value = v
}

func (o ItemKeyValuePairObjectObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemKeyValuePairObjectObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableItemKeyValuePairObjectObject struct {
	value *ItemKeyValuePairObjectObject
	isSet bool
}

func (v NullableItemKeyValuePairObjectObject) Get() *ItemKeyValuePairObjectObject {
	return v.value
}

func (v *NullableItemKeyValuePairObjectObject) Set(val *ItemKeyValuePairObjectObject) {
	v.value = val
	v.isSet = true
}

func (v NullableItemKeyValuePairObjectObject) IsSet() bool {
	return v.isSet
}

func (v *NullableItemKeyValuePairObjectObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemKeyValuePairObjectObject(val *ItemKeyValuePairObjectObject) *NullableItemKeyValuePairObjectObject {
	return &NullableItemKeyValuePairObjectObject{value: val, isSet: true}
}

func (v NullableItemKeyValuePairObjectObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemKeyValuePairObjectObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

