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

// checks if the Culture type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Culture{}

// Culture The culture name parameters.
type Culture struct {
	// The user culture name (en-US, de, fr, es, ...).
	CultureName string `json:"cultureName"`
}

type _Culture Culture

// NewCulture instantiates a new Culture object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCulture(cultureName string) *Culture {
	this := Culture{}
	this.CultureName = cultureName
	return &this
}

// NewCultureWithDefaults instantiates a new Culture object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCultureWithDefaults() *Culture {
	this := Culture{}
	return &this
}

// GetCultureName returns the CultureName field value
func (o *Culture) GetCultureName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CultureName
}

// GetCultureNameOk returns a tuple with the CultureName field value
// and a boolean to check if the value has been set.
func (o *Culture) GetCultureNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CultureName, true
}

// SetCultureName sets field value
func (o *Culture) SetCultureName(v string) {
	o.CultureName = v
}

func (o Culture) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Culture) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cultureName"] = o.CultureName
	return toSerialize, nil
}

func (o *Culture) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cultureName",
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

	varCulture := _Culture{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCulture)

	if err != nil {
		return err
	}

	*o = Culture(varCulture)

	return err
}

type NullableCulture struct {
	value *Culture
	isSet bool
}

func (v NullableCulture) Get() *Culture {
	return v.value
}

func (v *NullableCulture) Set(val *Culture) {
	v.value = val
	v.isSet = true
}

func (v NullableCulture) IsSet() bool {
	return v.isSet
}

func (v *NullableCulture) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCulture(val *Culture) *NullableCulture {
	return &NullableCulture{value: val, isSet: true}
}

func (v NullableCulture) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCulture) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

