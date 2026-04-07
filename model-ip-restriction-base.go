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

// checks if the IpRestrictionBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpRestrictionBase{}

// IpRestrictionBase struct for IpRestrictionBase
type IpRestrictionBase struct {
	Ip NullableString `json:"ip"`
	ForAdmin *bool `json:"forAdmin,omitempty"`
}

type _IpRestrictionBase IpRestrictionBase

// NewIpRestrictionBase instantiates a new IpRestrictionBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpRestrictionBase(ip NullableString) *IpRestrictionBase {
	this := IpRestrictionBase{}
	this.Ip = ip
	return &this
}

// NewIpRestrictionBaseWithDefaults instantiates a new IpRestrictionBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpRestrictionBaseWithDefaults() *IpRestrictionBase {
	this := IpRestrictionBase{}
	return &this
}

// GetIp returns the Ip field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IpRestrictionBase) GetIp() string {
	if o == nil || o.Ip.Get() == nil {
		var ret string
		return ret
	}

	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpRestrictionBase) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// SetIp sets field value
func (o *IpRestrictionBase) SetIp(v string) {
	o.Ip.Set(&v)
}

// GetForAdmin returns the ForAdmin field value if set, zero value otherwise.
func (o *IpRestrictionBase) GetForAdmin() bool {
	if o == nil || IsNil(o.ForAdmin) {
		var ret bool
		return ret
	}
	return *o.ForAdmin
}

// GetForAdminOk returns a tuple with the ForAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpRestrictionBase) GetForAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.ForAdmin) {
		return nil, false
	}
	return o.ForAdmin, true
}

// HasForAdmin returns a boolean if a field has been set.
func (o *IpRestrictionBase) IsForAdminSet() bool {
	if o != nil && !IsNil(o.ForAdmin) {
		return true
	}

	return false
}

// SetForAdmin gets a reference to the given bool and assigns it to the ForAdmin field.
func (o *IpRestrictionBase) SetForAdmin(v bool) {
	o.ForAdmin = &v
}

func (o IpRestrictionBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpRestrictionBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ip"] = o.Ip.Get()
	if !IsNil(o.ForAdmin) {
		toSerialize["forAdmin"] = o.ForAdmin
	}
	return toSerialize, nil
}

func (o *IpRestrictionBase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ip",
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

	varIpRestrictionBase := _IpRestrictionBase{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varIpRestrictionBase)

	if err != nil {
		return err
	}

	*o = IpRestrictionBase(varIpRestrictionBase)

	return err
}

type NullableIpRestrictionBase struct {
	value *IpRestrictionBase
	isSet bool
}

func (v NullableIpRestrictionBase) Get() *IpRestrictionBase {
	return v.value
}

func (v *NullableIpRestrictionBase) Set(val *IpRestrictionBase) {
	v.value = val
	v.isSet = true
}

func (v NullableIpRestrictionBase) IsSet() bool {
	return v.isSet
}

func (v *NullableIpRestrictionBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpRestrictionBase(val *IpRestrictionBase) *NullableIpRestrictionBase {
	return &NullableIpRestrictionBase{value: val, isSet: true}
}

func (v NullableIpRestrictionBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpRestrictionBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

