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

// checks if the RoomsNotificationSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomsNotificationSettingsDto{}

// RoomsNotificationSettingsDto The rooms notification settings.
type RoomsNotificationSettingsDto struct {
	// The list of rooms with the disabled notifications.
	DisabledRooms []map[string]interface{} `json:"disabledRooms,omitempty"`
}

// NewRoomsNotificationSettingsDto instantiates a new RoomsNotificationSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomsNotificationSettingsDto() *RoomsNotificationSettingsDto {
	this := RoomsNotificationSettingsDto{}
	return &this
}

// NewRoomsNotificationSettingsDtoWithDefaults instantiates a new RoomsNotificationSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomsNotificationSettingsDtoWithDefaults() *RoomsNotificationSettingsDto {
	this := RoomsNotificationSettingsDto{}
	return &this
}

// GetDisabledRooms returns the DisabledRooms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomsNotificationSettingsDto) GetDisabledRooms() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.DisabledRooms
}

// GetDisabledRoomsOk returns a tuple with the DisabledRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomsNotificationSettingsDto) GetDisabledRoomsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.DisabledRooms) {
		return nil, false
	}
	return o.DisabledRooms, true
}

// HasDisabledRooms returns a boolean if a field has been set.
func (o *RoomsNotificationSettingsDto) IsDisabledRoomsSet() bool {
	if o != nil && !IsNil(o.DisabledRooms) {
		return true
	}

	return false
}

// SetDisabledRooms gets a reference to the given []map[string]interface{} and assigns it to the DisabledRooms field.
func (o *RoomsNotificationSettingsDto) SetDisabledRooms(v []map[string]interface{}) {
	o.DisabledRooms = v
}

func (o RoomsNotificationSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomsNotificationSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DisabledRooms != nil {
		toSerialize["disabledRooms"] = o.DisabledRooms
	}
	return toSerialize, nil
}

type NullableRoomsNotificationSettingsDto struct {
	value *RoomsNotificationSettingsDto
	isSet bool
}

func (v NullableRoomsNotificationSettingsDto) Get() *RoomsNotificationSettingsDto {
	return v.value
}

func (v *NullableRoomsNotificationSettingsDto) Set(val *RoomsNotificationSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomsNotificationSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomsNotificationSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomsNotificationSettingsDto(val *RoomsNotificationSettingsDto) *NullableRoomsNotificationSettingsDto {
	return &NullableRoomsNotificationSettingsDto{value: val, isSet: true}
}

func (v NullableRoomsNotificationSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomsNotificationSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

