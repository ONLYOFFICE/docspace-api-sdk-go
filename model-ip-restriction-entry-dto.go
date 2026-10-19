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

// checks if the IpRestrictionEntryDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpRestrictionEntryDto{}

// IpRestrictionEntryDto One allowed address.
type IpRestrictionEntryDto struct {
	// The IPv4 or IPv6 address.
	Ip NullableString `json:"ip"`
	// Whether the address admits administrators only.
	ForAdmin *bool `json:"forAdmin,omitempty"`
}

type _IpRestrictionEntryDto IpRestrictionEntryDto

// NewIpRestrictionEntryDto instantiates a new IpRestrictionEntryDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpRestrictionEntryDto(ip NullableString) *IpRestrictionEntryDto {
	this := IpRestrictionEntryDto{}
	this.Ip = ip
	return &this
}

// NewIpRestrictionEntryDtoWithDefaults instantiates a new IpRestrictionEntryDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpRestrictionEntryDtoWithDefaults() *IpRestrictionEntryDto {
	this := IpRestrictionEntryDto{}
	return &this
}

// GetIp returns the Ip field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IpRestrictionEntryDto) GetIp() string {
	if o == nil || o.Ip.Get() == nil {
		var ret string
		return ret
	}

	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IpRestrictionEntryDto) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// SetIp sets field value
func (o *IpRestrictionEntryDto) SetIp(v string) {
	o.Ip.Set(&v)
}

// GetForAdmin returns the ForAdmin field value if set, zero value otherwise.
func (o *IpRestrictionEntryDto) GetForAdmin() bool {
	if o == nil || IsNil(o.ForAdmin) {
		var ret bool
		return ret
	}
	return *o.ForAdmin
}

// GetForAdminOk returns a tuple with the ForAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpRestrictionEntryDto) GetForAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.ForAdmin) {
		return nil, false
	}
	return o.ForAdmin, true
}

// HasForAdmin returns a boolean if a field has been set.
func (o *IpRestrictionEntryDto) IsForAdminSet() bool {
	if o != nil && !IsNil(o.ForAdmin) {
		return true
	}

	return false
}

// SetForAdmin gets a reference to the given bool and assigns it to the ForAdmin field.
func (o *IpRestrictionEntryDto) SetForAdmin(v bool) {
	o.ForAdmin = &v
}

func (o IpRestrictionEntryDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpRestrictionEntryDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ip"] = o.Ip.Get()
	if !IsNil(o.ForAdmin) {
		toSerialize["forAdmin"] = o.ForAdmin
	}
	return toSerialize, nil
}

func (o *IpRestrictionEntryDto) UnmarshalJSON(data []byte) (err error) {
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

	varIpRestrictionEntryDto := _IpRestrictionEntryDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varIpRestrictionEntryDto)

	if err != nil {
		return err
	}

	*o = IpRestrictionEntryDto(varIpRestrictionEntryDto)

	return err
}

type NullableIpRestrictionEntryDto struct {
	value *IpRestrictionEntryDto
	isSet bool
}

func (v NullableIpRestrictionEntryDto) Get() *IpRestrictionEntryDto {
	return v.value
}

func (v *NullableIpRestrictionEntryDto) Set(val *IpRestrictionEntryDto) {
	v.value = val
	v.isSet = true
}

func (v NullableIpRestrictionEntryDto) IsSet() bool {
	return v.isSet
}

func (v *NullableIpRestrictionEntryDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpRestrictionEntryDto(val *IpRestrictionEntryDto) *NullableIpRestrictionEntryDto {
	return &NullableIpRestrictionEntryDto{value: val, isSet: true}
}

func (v NullableIpRestrictionEntryDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpRestrictionEntryDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

