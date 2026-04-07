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
	"bytes"
	"fmt"
)

// checks if the Icon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Icon{}

// Icon struct for Icon
type Icon struct {
	Icon48 NullableString `json:"icon48"`
	Icon32 NullableString `json:"icon32"`
	Icon24 NullableString `json:"icon24"`
	Icon16 NullableString `json:"icon16"`
}

type _Icon Icon

// NewIcon instantiates a new Icon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIcon(icon48 NullableString, icon32 NullableString, icon24 NullableString, icon16 NullableString) *Icon {
	this := Icon{}
	this.Icon48 = icon48
	this.Icon32 = icon32
	this.Icon24 = icon24
	this.Icon16 = icon16
	return &this
}

// NewIconWithDefaults instantiates a new Icon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIconWithDefaults() *Icon {
	this := Icon{}
	return &this
}

// GetIcon48 returns the Icon48 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Icon) GetIcon48() string {
	if o == nil || o.Icon48.Get() == nil {
		var ret string
		return ret
	}

	return *o.Icon48.Get()
}

// GetIcon48Ok returns a tuple with the Icon48 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Icon) GetIcon48Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon48.Get(), o.Icon48.IsSet()
}

// SetIcon48 sets field value
func (o *Icon) SetIcon48(v string) {
	o.Icon48.Set(&v)
}

// GetIcon32 returns the Icon32 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Icon) GetIcon32() string {
	if o == nil || o.Icon32.Get() == nil {
		var ret string
		return ret
	}

	return *o.Icon32.Get()
}

// GetIcon32Ok returns a tuple with the Icon32 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Icon) GetIcon32Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon32.Get(), o.Icon32.IsSet()
}

// SetIcon32 sets field value
func (o *Icon) SetIcon32(v string) {
	o.Icon32.Set(&v)
}

// GetIcon24 returns the Icon24 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Icon) GetIcon24() string {
	if o == nil || o.Icon24.Get() == nil {
		var ret string
		return ret
	}

	return *o.Icon24.Get()
}

// GetIcon24Ok returns a tuple with the Icon24 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Icon) GetIcon24Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon24.Get(), o.Icon24.IsSet()
}

// SetIcon24 sets field value
func (o *Icon) SetIcon24(v string) {
	o.Icon24.Set(&v)
}

// GetIcon16 returns the Icon16 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Icon) GetIcon16() string {
	if o == nil || o.Icon16.Get() == nil {
		var ret string
		return ret
	}

	return *o.Icon16.Get()
}

// GetIcon16Ok returns a tuple with the Icon16 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Icon) GetIcon16Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon16.Get(), o.Icon16.IsSet()
}

// SetIcon16 sets field value
func (o *Icon) SetIcon16(v string) {
	o.Icon16.Set(&v)
}

func (o Icon) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Icon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["icon48"] = o.Icon48.Get()
	toSerialize["icon32"] = o.Icon32.Get()
	toSerialize["icon24"] = o.Icon24.Get()
	toSerialize["icon16"] = o.Icon16.Get()
	return toSerialize, nil
}

func (o *Icon) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"icon48",
		"icon32",
		"icon24",
		"icon16",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varIcon := _Icon{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varIcon)

	if err != nil {
		return err
	}

	*o = Icon(varIcon)

	return err
}

type NullableIcon struct {
	value *Icon
	isSet bool
}

func (v NullableIcon) Get() *Icon {
	return v.value
}

func (v *NullableIcon) Set(val *Icon) {
	v.value = val
	v.isSet = true
}

func (v NullableIcon) IsSet() bool {
	return v.isSet
}

func (v *NullableIcon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIcon(val *Icon) *NullableIcon {
	return &NullableIcon{value: val, isSet: true}
}

func (v NullableIcon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIcon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

