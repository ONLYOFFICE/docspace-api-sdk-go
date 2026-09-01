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

// checks if the IpRestrictionsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpRestrictionsDto{}

// IpRestrictionsDto The parameters for configuring new IP restriction settings.
type IpRestrictionsDto struct {
	// The list of IP restriction addresses.
	IpRestrictions []IpRestrictionBase `json:"ipRestrictions"`
	// Specifies whether to enable IP restrictions or not.
	Enable NullableBool `json:"enable,omitempty"`
}

type _IpRestrictionsDto IpRestrictionsDto

// NewIpRestrictionsDto instantiates a new IpRestrictionsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpRestrictionsDto(ipRestrictions []IpRestrictionBase) *IpRestrictionsDto {
	this := IpRestrictionsDto{}
	this.IpRestrictions = ipRestrictions
	return &this
}

// NewIpRestrictionsDtoWithDefaults instantiates a new IpRestrictionsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpRestrictionsDtoWithDefaults() *IpRestrictionsDto {
	this := IpRestrictionsDto{}
	return &this
}

// GetIpRestrictions returns the IpRestrictions field value
// If the value is explicit nil, the zero value for []IpRestrictionBase will be returned
func (o *IpRestrictionsDto) GetIpRestrictions() []IpRestrictionBase {
	if o == nil {
		var ret []IpRestrictionBase
		return ret
	}

	return o.IpRestrictions
}

// GetIpRestrictionsOk returns a tuple with the IpRestrictions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpRestrictionsDto) GetIpRestrictionsOk() ([]IpRestrictionBase, bool) {
	if o == nil || IsNil(o.IpRestrictions) {
		return nil, false
	}
	return o.IpRestrictions, true
}

// SetIpRestrictions sets field value
func (o *IpRestrictionsDto) SetIpRestrictions(v []IpRestrictionBase) {
	o.IpRestrictions = v
}

// GetEnable returns the Enable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IpRestrictionsDto) GetEnable() bool {
	if o == nil || IsNil(o.Enable.Get()) {
		var ret bool
		return ret
	}
	return *o.Enable.Get()
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpRestrictionsDto) GetEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enable.Get(), o.Enable.IsSet()
}

// HasEnable returns a boolean if a field has been set.
func (o *IpRestrictionsDto) IsEnableSet() bool {
	if o != nil && o.Enable.IsSet() {
		return true
	}

	return false
}

// SetEnable gets a reference to the given NullableBool and assigns it to the Enable field.
func (o *IpRestrictionsDto) SetEnable(v bool) {
	o.Enable.Set(&v)
}
// SetEnableNil sets the value for Enable to be an explicit nil
func (o *IpRestrictionsDto) SetEnableNil() {
	o.Enable.Set(nil)
}

// UnsetEnable ensures that no value is present for Enable, not even an explicit nil
func (o *IpRestrictionsDto) UnsetEnable() {
	o.Enable.Unset()
}

func (o IpRestrictionsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpRestrictionsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IpRestrictions != nil {
		toSerialize["ipRestrictions"] = o.IpRestrictions
	}
	if o.Enable.IsSet() {
		toSerialize["enable"] = o.Enable.Get()
	}
	return toSerialize, nil
}

func (o *IpRestrictionsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ipRestrictions",
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

	varIpRestrictionsDto := _IpRestrictionsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varIpRestrictionsDto)

	if err != nil {
		return err
	}

	*o = IpRestrictionsDto(varIpRestrictionsDto)

	return err
}

type NullableIpRestrictionsDto struct {
	value *IpRestrictionsDto
	isSet bool
}

func (v NullableIpRestrictionsDto) Get() *IpRestrictionsDto {
	return v.value
}

func (v *NullableIpRestrictionsDto) Set(val *IpRestrictionsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableIpRestrictionsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableIpRestrictionsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpRestrictionsDto(val *IpRestrictionsDto) *NullableIpRestrictionsDto {
	return &NullableIpRestrictionsDto{value: val, isSet: true}
}

func (v NullableIpRestrictionsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpRestrictionsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

