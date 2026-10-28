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

// checks if the RoomNewItemsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomNewItemsDto{}

// RoomNewItemsDto The unseen entries of one room inside a day group.
type RoomNewItemsDto struct {
	// The room the entries were found in, in its short form: only the identifier, the title, the room type and the  logo are filled in.
	Room *FileEntryBaseDto `json:"room,omitempty"`
	// The files of that room the caller has not opened yet, the most recently changed first. Reading them here does  not clear the badges; opening the room itself does.
	Items []FileEntryBaseDto `json:"items,omitempty"`
}

// NewRoomNewItemsDto instantiates a new RoomNewItemsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomNewItemsDto() *RoomNewItemsDto {
	this := RoomNewItemsDto{}
	return &this
}

// NewRoomNewItemsDtoWithDefaults instantiates a new RoomNewItemsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomNewItemsDtoWithDefaults() *RoomNewItemsDto {
	this := RoomNewItemsDto{}
	return &this
}

// GetRoom returns the Room field value if set, zero value otherwise.
func (o *RoomNewItemsDto) GetRoom() FileEntryBaseDto {
	if o == nil || IsNil(o.Room) {
		var ret FileEntryBaseDto
		return ret
	}
	return *o.Room
}

// GetRoomOk returns a tuple with the Room field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomNewItemsDto) GetRoomOk() (*FileEntryBaseDto, bool) {
	if o == nil || IsNil(o.Room) {
		return nil, false
	}
	return o.Room, true
}

// HasRoom returns a boolean if a field has been set.
func (o *RoomNewItemsDto) IsRoomSet() bool {
	if o != nil && !IsNil(o.Room) {
		return true
	}

	return false
}

// SetRoom gets a reference to the given FileEntryBaseDto and assigns it to the Room field.
func (o *RoomNewItemsDto) SetRoom(v FileEntryBaseDto) {
	o.Room = &v
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomNewItemsDto) GetItems() []FileEntryBaseDto {
	if o == nil {
		var ret []FileEntryBaseDto
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomNewItemsDto) GetItemsOk() ([]FileEntryBaseDto, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *RoomNewItemsDto) IsItemsSet() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []FileEntryBaseDto and assigns it to the Items field.
func (o *RoomNewItemsDto) SetItems(v []FileEntryBaseDto) {
	o.Items = v
}

func (o RoomNewItemsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomNewItemsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Room) {
		toSerialize["room"] = o.Room
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableRoomNewItemsDto struct {
	value *RoomNewItemsDto
	isSet bool
}

func (v NullableRoomNewItemsDto) Get() *RoomNewItemsDto {
	return v.value
}

func (v *NullableRoomNewItemsDto) Set(val *RoomNewItemsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomNewItemsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomNewItemsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomNewItemsDto(val *RoomNewItemsDto) *NullableRoomNewItemsDto {
	return &NullableRoomNewItemsDto{value: val, isSet: true}
}

func (v NullableRoomNewItemsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomNewItemsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

