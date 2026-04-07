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

// checks if the UpdateRoomsRoomIdsRequestDtoInteger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRoomsRoomIdsRequestDtoInteger{}

// UpdateRoomsRoomIdsRequestDtoInteger The request parameters for updating the rooms.
type UpdateRoomsRoomIdsRequestDtoInteger struct {
	// The list of room IDs.
	RoomIds []ContinueChatBodyFilesInner `json:"roomIds,omitempty"`
}

// NewUpdateRoomsRoomIdsRequestDtoInteger instantiates a new UpdateRoomsRoomIdsRequestDtoInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRoomsRoomIdsRequestDtoInteger() *UpdateRoomsRoomIdsRequestDtoInteger {
	this := UpdateRoomsRoomIdsRequestDtoInteger{}
	return &this
}

// NewUpdateRoomsRoomIdsRequestDtoIntegerWithDefaults instantiates a new UpdateRoomsRoomIdsRequestDtoInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoomsRoomIdsRequestDtoIntegerWithDefaults() *UpdateRoomsRoomIdsRequestDtoInteger {
	this := UpdateRoomsRoomIdsRequestDtoInteger{}
	return &this
}

// GetRoomIds returns the RoomIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateRoomsRoomIdsRequestDtoInteger) GetRoomIds() []ContinueChatBodyFilesInner {
	if o == nil {
		var ret []ContinueChatBodyFilesInner
		return ret
	}
	return o.RoomIds
}

// GetRoomIdsOk returns a tuple with the RoomIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateRoomsRoomIdsRequestDtoInteger) GetRoomIdsOk() ([]ContinueChatBodyFilesInner, bool) {
	if o == nil || IsNil(o.RoomIds) {
		return nil, false
	}
	return o.RoomIds, true
}

// HasRoomIds returns a boolean if a field has been set.
func (o *UpdateRoomsRoomIdsRequestDtoInteger) IsRoomIdsSet() bool {
	if o != nil && !IsNil(o.RoomIds) {
		return true
	}

	return false
}

// SetRoomIds gets a reference to the given []ContinueChatBodyFilesInner and assigns it to the RoomIds field.
func (o *UpdateRoomsRoomIdsRequestDtoInteger) SetRoomIds(v []ContinueChatBodyFilesInner) {
	o.RoomIds = v
}

func (o UpdateRoomsRoomIdsRequestDtoInteger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRoomsRoomIdsRequestDtoInteger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RoomIds != nil {
		toSerialize["roomIds"] = o.RoomIds
	}
	return toSerialize, nil
}

type NullableUpdateRoomsRoomIdsRequestDtoInteger struct {
	value *UpdateRoomsRoomIdsRequestDtoInteger
	isSet bool
}

func (v NullableUpdateRoomsRoomIdsRequestDtoInteger) Get() *UpdateRoomsRoomIdsRequestDtoInteger {
	return v.value
}

func (v *NullableUpdateRoomsRoomIdsRequestDtoInteger) Set(val *UpdateRoomsRoomIdsRequestDtoInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRoomsRoomIdsRequestDtoInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRoomsRoomIdsRequestDtoInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRoomsRoomIdsRequestDtoInteger(val *UpdateRoomsRoomIdsRequestDtoInteger) *NullableUpdateRoomsRoomIdsRequestDtoInteger {
	return &NullableUpdateRoomsRoomIdsRequestDtoInteger{value: val, isSet: true}
}

func (v NullableUpdateRoomsRoomIdsRequestDtoInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRoomsRoomIdsRequestDtoInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

