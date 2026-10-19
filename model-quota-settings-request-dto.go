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

// checks if the QuotaSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuotaSettingsRequestDto{}

// QuotaSettingsRequestDto The default storage limit given to newly created users, rooms or AI agents, and whether it is enforced.
type QuotaSettingsRequestDto struct {
	// Whether the limit is enforced at all. While it is false the size is ignored and nothing created afterwards  carries a limit; objects that already have one keep it either way.
	EnableQuota *bool `json:"enableQuota,omitempty"`
	DefaultQuota QuotaSettingsRequestDtoDefaultQuota `json:"defaultQuota"`
}

type _QuotaSettingsRequestDto QuotaSettingsRequestDto

// NewQuotaSettingsRequestDto instantiates a new QuotaSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaSettingsRequestDto(defaultQuota QuotaSettingsRequestDtoDefaultQuota) *QuotaSettingsRequestDto {
	this := QuotaSettingsRequestDto{}
	this.DefaultQuota = defaultQuota
	return &this
}

// NewQuotaSettingsRequestDtoWithDefaults instantiates a new QuotaSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaSettingsRequestDtoWithDefaults() *QuotaSettingsRequestDto {
	this := QuotaSettingsRequestDto{}
	return &this
}

// GetEnableQuota returns the EnableQuota field value if set, zero value otherwise.
func (o *QuotaSettingsRequestDto) GetEnableQuota() bool {
	if o == nil || IsNil(o.EnableQuota) {
		var ret bool
		return ret
	}
	return *o.EnableQuota
}

// GetEnableQuotaOk returns a tuple with the EnableQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaSettingsRequestDto) GetEnableQuotaOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableQuota) {
		return nil, false
	}
	return o.EnableQuota, true
}

// HasEnableQuota returns a boolean if a field has been set.
func (o *QuotaSettingsRequestDto) IsEnableQuotaSet() bool {
	if o != nil && !IsNil(o.EnableQuota) {
		return true
	}

	return false
}

// SetEnableQuota gets a reference to the given bool and assigns it to the EnableQuota field.
func (o *QuotaSettingsRequestDto) SetEnableQuota(v bool) {
	o.EnableQuota = &v
}

// GetDefaultQuota returns the DefaultQuota field value
func (o *QuotaSettingsRequestDto) GetDefaultQuota() QuotaSettingsRequestDtoDefaultQuota {
	if o == nil {
		var ret QuotaSettingsRequestDtoDefaultQuota
		return ret
	}

	return o.DefaultQuota
}

// GetDefaultQuotaOk returns a tuple with the DefaultQuota field value
// and a boolean to check if the value has been set.
func (o *QuotaSettingsRequestDto) GetDefaultQuotaOk() (*QuotaSettingsRequestDtoDefaultQuota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultQuota, true
}

// SetDefaultQuota sets field value
func (o *QuotaSettingsRequestDto) SetDefaultQuota(v QuotaSettingsRequestDtoDefaultQuota) {
	o.DefaultQuota = v
}

func (o QuotaSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuotaSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableQuota) {
		toSerialize["enableQuota"] = o.EnableQuota
	}
	toSerialize["defaultQuota"] = o.DefaultQuota
	return toSerialize, nil
}

func (o *QuotaSettingsRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"defaultQuota",
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

	varQuotaSettingsRequestDto := _QuotaSettingsRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varQuotaSettingsRequestDto)

	if err != nil {
		return err
	}

	*o = QuotaSettingsRequestDto(varQuotaSettingsRequestDto)

	return err
}

type NullableQuotaSettingsRequestDto struct {
	value *QuotaSettingsRequestDto
	isSet bool
}

func (v NullableQuotaSettingsRequestDto) Get() *QuotaSettingsRequestDto {
	return v.value
}

func (v *NullableQuotaSettingsRequestDto) Set(val *QuotaSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaSettingsRequestDto(val *QuotaSettingsRequestDto) *NullableQuotaSettingsRequestDto {
	return &NullableQuotaSettingsRequestDto{value: val, isSet: true}
}

func (v NullableQuotaSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

