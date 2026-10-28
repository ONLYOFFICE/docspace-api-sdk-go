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

// checks if the UpdateRoomsRoomIdsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRoomsRoomIdsRequestDto{}

// UpdateRoomsRoomIdsRequestDto The rooms that are to go back to the default storage limit of the portal.
type UpdateRoomsRoomIdsRequestDto struct {
	// The rooms to reset, named by the identifiers that `GET api/2.0/files/rooms` reports. Only whole numbers are  processed, so identifiers of rooms kept in a connected third-party account are skipped without an error.
	RoomIds []DuplicateRequestDtoAllOfFileIds `json:"roomIds,omitempty"`
}

// NewUpdateRoomsRoomIdsRequestDto instantiates a new UpdateRoomsRoomIdsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRoomsRoomIdsRequestDto() *UpdateRoomsRoomIdsRequestDto {
	this := UpdateRoomsRoomIdsRequestDto{}
	return &this
}

// NewUpdateRoomsRoomIdsRequestDtoWithDefaults instantiates a new UpdateRoomsRoomIdsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoomsRoomIdsRequestDtoWithDefaults() *UpdateRoomsRoomIdsRequestDto {
	this := UpdateRoomsRoomIdsRequestDto{}
	return &this
}

// GetRoomIds returns the RoomIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateRoomsRoomIdsRequestDto) GetRoomIds() []DuplicateRequestDtoAllOfFileIds {
	if o == nil {
		var ret []DuplicateRequestDtoAllOfFileIds
		return ret
	}
	return o.RoomIds
}

// GetRoomIdsOk returns a tuple with the RoomIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateRoomsRoomIdsRequestDto) GetRoomIdsOk() ([]DuplicateRequestDtoAllOfFileIds, bool) {
	if o == nil || IsNil(o.RoomIds) {
		return nil, false
	}
	return o.RoomIds, true
}

// HasRoomIds returns a boolean if a field has been set.
func (o *UpdateRoomsRoomIdsRequestDto) IsRoomIdsSet() bool {
	if o != nil && !IsNil(o.RoomIds) {
		return true
	}

	return false
}

// SetRoomIds gets a reference to the given []DuplicateRequestDtoAllOfFileIds and assigns it to the RoomIds field.
func (o *UpdateRoomsRoomIdsRequestDto) SetRoomIds(v []DuplicateRequestDtoAllOfFileIds) {
	o.RoomIds = v
}

func (o UpdateRoomsRoomIdsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRoomsRoomIdsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RoomIds != nil {
		toSerialize["roomIds"] = o.RoomIds
	}
	return toSerialize, nil
}

type NullableUpdateRoomsRoomIdsRequestDto struct {
	value *UpdateRoomsRoomIdsRequestDto
	isSet bool
}

func (v NullableUpdateRoomsRoomIdsRequestDto) Get() *UpdateRoomsRoomIdsRequestDto {
	return v.value
}

func (v *NullableUpdateRoomsRoomIdsRequestDto) Set(val *UpdateRoomsRoomIdsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRoomsRoomIdsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRoomsRoomIdsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRoomsRoomIdsRequestDto(val *UpdateRoomsRoomIdsRequestDto) *NullableUpdateRoomsRoomIdsRequestDto {
	return &NullableUpdateRoomsRoomIdsRequestDto{value: val, isSet: true}
}

func (v NullableUpdateRoomsRoomIdsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRoomsRoomIdsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

