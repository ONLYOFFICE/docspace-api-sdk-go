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

// checks if the RoomsNotificationsSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomsNotificationsSettingsRequestDto{}

// RoomsNotificationsSettingsRequestDto The request parameters for configuring notification settings for the chat or collaboration rooms.
type RoomsNotificationsSettingsRequestDto struct {
	RoomsId interface{} `json:"roomsId,omitempty"`
	// Specifies whether the notifications will be delivered to the specified room or not.
	Mute *bool `json:"mute,omitempty"`
}

// NewRoomsNotificationsSettingsRequestDto instantiates a new RoomsNotificationsSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomsNotificationsSettingsRequestDto() *RoomsNotificationsSettingsRequestDto {
	this := RoomsNotificationsSettingsRequestDto{}
	return &this
}

// NewRoomsNotificationsSettingsRequestDtoWithDefaults instantiates a new RoomsNotificationsSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomsNotificationsSettingsRequestDtoWithDefaults() *RoomsNotificationsSettingsRequestDto {
	this := RoomsNotificationsSettingsRequestDto{}
	return &this
}

// GetRoomsId returns the RoomsId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomsNotificationsSettingsRequestDto) GetRoomsId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RoomsId
}

// GetRoomsIdOk returns a tuple with the RoomsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomsNotificationsSettingsRequestDto) GetRoomsIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RoomsId) {
		return nil, false
	}
	return &o.RoomsId, true
}

// HasRoomsId returns a boolean if a field has been set.
func (o *RoomsNotificationsSettingsRequestDto) IsRoomsIdSet() bool {
	if o != nil && !IsNil(o.RoomsId) {
		return true
	}

	return false
}

// SetRoomsId gets a reference to the given interface{} and assigns it to the RoomsId field.
func (o *RoomsNotificationsSettingsRequestDto) SetRoomsId(v interface{}) {
	o.RoomsId = v
}

// GetMute returns the Mute field value if set, zero value otherwise.
func (o *RoomsNotificationsSettingsRequestDto) GetMute() bool {
	if o == nil || IsNil(o.Mute) {
		var ret bool
		return ret
	}
	return *o.Mute
}

// GetMuteOk returns a tuple with the Mute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomsNotificationsSettingsRequestDto) GetMuteOk() (*bool, bool) {
	if o == nil || IsNil(o.Mute) {
		return nil, false
	}
	return o.Mute, true
}

// HasMute returns a boolean if a field has been set.
func (o *RoomsNotificationsSettingsRequestDto) IsMuteSet() bool {
	if o != nil && !IsNil(o.Mute) {
		return true
	}

	return false
}

// SetMute gets a reference to the given bool and assigns it to the Mute field.
func (o *RoomsNotificationsSettingsRequestDto) SetMute(v bool) {
	o.Mute = &v
}

func (o RoomsNotificationsSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomsNotificationsSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RoomsId != nil {
		toSerialize["roomsId"] = o.RoomsId
	}
	if !IsNil(o.Mute) {
		toSerialize["mute"] = o.Mute
	}
	return toSerialize, nil
}

type NullableRoomsNotificationsSettingsRequestDto struct {
	value *RoomsNotificationsSettingsRequestDto
	isSet bool
}

func (v NullableRoomsNotificationsSettingsRequestDto) Get() *RoomsNotificationsSettingsRequestDto {
	return v.value
}

func (v *NullableRoomsNotificationsSettingsRequestDto) Set(val *RoomsNotificationsSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomsNotificationsSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomsNotificationsSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomsNotificationsSettingsRequestDto(val *RoomsNotificationsSettingsRequestDto) *NullableRoomsNotificationsSettingsRequestDto {
	return &NullableRoomsNotificationsSettingsRequestDto{value: val, isSet: true}
}

func (v NullableRoomsNotificationsSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomsNotificationsSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

