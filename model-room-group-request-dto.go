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

// checks if the RoomGroupRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomGroupRequestDto{}

// RoomGroupRequestDto The name, the icon and the rooms of a room group to create.
type RoomGroupRequestDto struct {
	// The name to show the group under. Surrounding spaces are trimmed before it is stored, a name that is blank  once trimmed is refused, and the name does not have to differ from the names of the caller's other groups.
	Name string `json:"name"`
	// The icon of the group, given as the identifier of one of the built-in covers listed by  `GET api/2.0/files/rooms/covers`. An uploaded image cannot be used, and any value that is not one of those  identifiers is refused.
	Icon string `json:"icon"`
	// The rooms to gather in the group, each given as a number for a room stored in the portal or as a string for a  room on a connected third-party account. Every identifier has to name a room the caller can read; repeats are  collapsed, and an element of any other shape - a decimal number, a number sent as a string, null - is refused.
	Rooms []DuplicateRequestDtoAllOfFileIds `json:"rooms"`
}

type _RoomGroupRequestDto RoomGroupRequestDto

// NewRoomGroupRequestDto instantiates a new RoomGroupRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomGroupRequestDto(name string, icon string, rooms []DuplicateRequestDtoAllOfFileIds) *RoomGroupRequestDto {
	this := RoomGroupRequestDto{}
	this.Name = name
	this.Icon = icon
	this.Rooms = rooms
	return &this
}

// NewRoomGroupRequestDtoWithDefaults instantiates a new RoomGroupRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomGroupRequestDtoWithDefaults() *RoomGroupRequestDto {
	this := RoomGroupRequestDto{}
	return &this
}

// GetName returns the Name field value
func (o *RoomGroupRequestDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RoomGroupRequestDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RoomGroupRequestDto) SetName(v string) {
	o.Name = v
}

// GetIcon returns the Icon field value
func (o *RoomGroupRequestDto) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *RoomGroupRequestDto) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *RoomGroupRequestDto) SetIcon(v string) {
	o.Icon = v
}

// GetRooms returns the Rooms field value
func (o *RoomGroupRequestDto) GetRooms() []DuplicateRequestDtoAllOfFileIds {
	if o == nil {
		var ret []DuplicateRequestDtoAllOfFileIds
		return ret
	}

	return o.Rooms
}

// GetRoomsOk returns a tuple with the Rooms field value
// and a boolean to check if the value has been set.
func (o *RoomGroupRequestDto) GetRoomsOk() ([]DuplicateRequestDtoAllOfFileIds, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rooms, true
}

// SetRooms sets field value
func (o *RoomGroupRequestDto) SetRooms(v []DuplicateRequestDtoAllOfFileIds) {
	o.Rooms = v
}

func (o RoomGroupRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomGroupRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["icon"] = o.Icon
	toSerialize["rooms"] = o.Rooms
	return toSerialize, nil
}

func (o *RoomGroupRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"icon",
		"rooms",
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

	varRoomGroupRequestDto := _RoomGroupRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varRoomGroupRequestDto)

	if err != nil {
		return err
	}

	*o = RoomGroupRequestDto(varRoomGroupRequestDto)

	return err
}

type NullableRoomGroupRequestDto struct {
	value *RoomGroupRequestDto
	isSet bool
}

func (v NullableRoomGroupRequestDto) Get() *RoomGroupRequestDto {
	return v.value
}

func (v *NullableRoomGroupRequestDto) Set(val *RoomGroupRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomGroupRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomGroupRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomGroupRequestDto(val *RoomGroupRequestDto) *NullableRoomGroupRequestDto {
	return &NullableRoomGroupRequestDto{value: val, isSet: true}
}

func (v NullableRoomGroupRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomGroupRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

