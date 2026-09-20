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

// checks if the QuotaSettingsRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuotaSettingsRequestsDto{}

// QuotaSettingsRequestsDto The default storage limit given to newly created users, rooms or AI agents, and whether it is enforced.
type QuotaSettingsRequestsDto struct {
	// Whether the limit is enforced at all. While it is false the size is ignored and nothing created afterwards  carries a limit; objects that already have one keep it either way.
	EnableQuota *bool `json:"enableQuota,omitempty"`
	DefaultQuota QuotaSettingsRequestsDtoDefaultQuota `json:"defaultQuota"`
}

type _QuotaSettingsRequestsDto QuotaSettingsRequestsDto

// NewQuotaSettingsRequestsDto instantiates a new QuotaSettingsRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaSettingsRequestsDto(defaultQuota QuotaSettingsRequestsDtoDefaultQuota) *QuotaSettingsRequestsDto {
	this := QuotaSettingsRequestsDto{}
	this.DefaultQuota = defaultQuota
	return &this
}

// NewQuotaSettingsRequestsDtoWithDefaults instantiates a new QuotaSettingsRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaSettingsRequestsDtoWithDefaults() *QuotaSettingsRequestsDto {
	this := QuotaSettingsRequestsDto{}
	return &this
}

// GetEnableQuota returns the EnableQuota field value if set, zero value otherwise.
func (o *QuotaSettingsRequestsDto) GetEnableQuota() bool {
	if o == nil || IsNil(o.EnableQuota) {
		var ret bool
		return ret
	}
	return *o.EnableQuota
}

// GetEnableQuotaOk returns a tuple with the EnableQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaSettingsRequestsDto) GetEnableQuotaOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableQuota) {
		return nil, false
	}
	return o.EnableQuota, true
}

// HasEnableQuota returns a boolean if a field has been set.
func (o *QuotaSettingsRequestsDto) IsEnableQuotaSet() bool {
	if o != nil && !IsNil(o.EnableQuota) {
		return true
	}

	return false
}

// SetEnableQuota gets a reference to the given bool and assigns it to the EnableQuota field.
func (o *QuotaSettingsRequestsDto) SetEnableQuota(v bool) {
	o.EnableQuota = &v
}

// GetDefaultQuota returns the DefaultQuota field value
func (o *QuotaSettingsRequestsDto) GetDefaultQuota() QuotaSettingsRequestsDtoDefaultQuota {
	if o == nil {
		var ret QuotaSettingsRequestsDtoDefaultQuota
		return ret
	}

	return o.DefaultQuota
}

// GetDefaultQuotaOk returns a tuple with the DefaultQuota field value
// and a boolean to check if the value has been set.
func (o *QuotaSettingsRequestsDto) GetDefaultQuotaOk() (*QuotaSettingsRequestsDtoDefaultQuota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultQuota, true
}

// SetDefaultQuota sets field value
func (o *QuotaSettingsRequestsDto) SetDefaultQuota(v QuotaSettingsRequestsDtoDefaultQuota) {
	o.DefaultQuota = v
}

func (o QuotaSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuotaSettingsRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableQuota) {
		toSerialize["enableQuota"] = o.EnableQuota
	}
	toSerialize["defaultQuota"] = o.DefaultQuota
	return toSerialize, nil
}

func (o *QuotaSettingsRequestsDto) UnmarshalJSON(data []byte) (err error) {
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

	varQuotaSettingsRequestsDto := _QuotaSettingsRequestsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varQuotaSettingsRequestsDto)

	if err != nil {
		return err
	}

	*o = QuotaSettingsRequestsDto(varQuotaSettingsRequestsDto)

	return err
}

type NullableQuotaSettingsRequestsDto struct {
	value *QuotaSettingsRequestsDto
	isSet bool
}

func (v NullableQuotaSettingsRequestsDto) Get() *QuotaSettingsRequestsDto {
	return v.value
}

func (v *NullableQuotaSettingsRequestsDto) Set(val *QuotaSettingsRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaSettingsRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaSettingsRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaSettingsRequestsDto(val *QuotaSettingsRequestsDto) *NullableQuotaSettingsRequestsDto {
	return &NullableQuotaSettingsRequestsDto{value: val, isSet: true}
}

func (v NullableQuotaSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaSettingsRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

