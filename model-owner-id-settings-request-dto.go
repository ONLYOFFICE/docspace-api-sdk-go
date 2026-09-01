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

// checks if the OwnerIdSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OwnerIdSettingsRequestDto{}

// OwnerIdSettingsRequestDto The request parameters for managing the owner-specific settings.
type OwnerIdSettingsRequestDto struct {
	// The ID of the owner whose settings are being managed.
	OwnerId string `json:"ownerId"`
}

type _OwnerIdSettingsRequestDto OwnerIdSettingsRequestDto

// NewOwnerIdSettingsRequestDto instantiates a new OwnerIdSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwnerIdSettingsRequestDto(ownerId string) *OwnerIdSettingsRequestDto {
	this := OwnerIdSettingsRequestDto{}
	this.OwnerId = ownerId
	return &this
}

// NewOwnerIdSettingsRequestDtoWithDefaults instantiates a new OwnerIdSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnerIdSettingsRequestDtoWithDefaults() *OwnerIdSettingsRequestDto {
	this := OwnerIdSettingsRequestDto{}
	return &this
}

// GetOwnerId returns the OwnerId field value
func (o *OwnerIdSettingsRequestDto) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *OwnerIdSettingsRequestDto) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *OwnerIdSettingsRequestDto) SetOwnerId(v string) {
	o.OwnerId = v
}

func (o OwnerIdSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OwnerIdSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ownerId"] = o.OwnerId
	return toSerialize, nil
}

func (o *OwnerIdSettingsRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ownerId",
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

	varOwnerIdSettingsRequestDto := _OwnerIdSettingsRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varOwnerIdSettingsRequestDto)

	if err != nil {
		return err
	}

	*o = OwnerIdSettingsRequestDto(varOwnerIdSettingsRequestDto)

	return err
}

type NullableOwnerIdSettingsRequestDto struct {
	value *OwnerIdSettingsRequestDto
	isSet bool
}

func (v NullableOwnerIdSettingsRequestDto) Get() *OwnerIdSettingsRequestDto {
	return v.value
}

func (v *NullableOwnerIdSettingsRequestDto) Set(val *OwnerIdSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnerIdSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnerIdSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnerIdSettingsRequestDto(val *OwnerIdSettingsRequestDto) *NullableOwnerIdSettingsRequestDto {
	return &NullableOwnerIdSettingsRequestDto{value: val, isSet: true}
}

func (v NullableOwnerIdSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnerIdSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

