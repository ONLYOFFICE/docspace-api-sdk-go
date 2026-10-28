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

// checks if the UpdateRoomsQuotaRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRoomsQuotaRequestDto{}

// UpdateRoomsQuotaRequestDto The rooms whose storage limit is to be changed, and the limit to give them.
type UpdateRoomsQuotaRequestDto struct {
	// The rooms to change, named by the identifiers that `GET api/2.0/files/rooms` reports. Only whole numbers are  processed, so identifiers of rooms kept in a connected third-party account are skipped without an error.
	RoomIds []DuplicateRequestDtoAllOfFileIds `json:"roomIds,omitempty"`
	// The storage each of the listed rooms may take, in bytes. It has to stay inside the portal own limit, and the  per-room quota feature has to be on, otherwise nothing is changed.
	Quota *int64 `json:"quota,omitempty"`
}

// NewUpdateRoomsQuotaRequestDto instantiates a new UpdateRoomsQuotaRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRoomsQuotaRequestDto() *UpdateRoomsQuotaRequestDto {
	this := UpdateRoomsQuotaRequestDto{}
	return &this
}

// NewUpdateRoomsQuotaRequestDtoWithDefaults instantiates a new UpdateRoomsQuotaRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoomsQuotaRequestDtoWithDefaults() *UpdateRoomsQuotaRequestDto {
	this := UpdateRoomsQuotaRequestDto{}
	return &this
}

// GetRoomIds returns the RoomIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateRoomsQuotaRequestDto) GetRoomIds() []DuplicateRequestDtoAllOfFileIds {
	if o == nil {
		var ret []DuplicateRequestDtoAllOfFileIds
		return ret
	}
	return o.RoomIds
}

// GetRoomIdsOk returns a tuple with the RoomIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateRoomsQuotaRequestDto) GetRoomIdsOk() ([]DuplicateRequestDtoAllOfFileIds, bool) {
	if o == nil || IsNil(o.RoomIds) {
		return nil, false
	}
	return o.RoomIds, true
}

// HasRoomIds returns a boolean if a field has been set.
func (o *UpdateRoomsQuotaRequestDto) IsRoomIdsSet() bool {
	if o != nil && !IsNil(o.RoomIds) {
		return true
	}

	return false
}

// SetRoomIds gets a reference to the given []DuplicateRequestDtoAllOfFileIds and assigns it to the RoomIds field.
func (o *UpdateRoomsQuotaRequestDto) SetRoomIds(v []DuplicateRequestDtoAllOfFileIds) {
	o.RoomIds = v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *UpdateRoomsQuotaRequestDto) GetQuota() int64 {
	if o == nil || IsNil(o.Quota) {
		var ret int64
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoomsQuotaRequestDto) GetQuotaOk() (*int64, bool) {
	if o == nil || IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *UpdateRoomsQuotaRequestDto) IsQuotaSet() bool {
	if o != nil && !IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given int64 and assigns it to the Quota field.
func (o *UpdateRoomsQuotaRequestDto) SetQuota(v int64) {
	o.Quota = &v
}

func (o UpdateRoomsQuotaRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRoomsQuotaRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RoomIds != nil {
		toSerialize["roomIds"] = o.RoomIds
	}
	if !IsNil(o.Quota) {
		toSerialize["quota"] = o.Quota
	}
	return toSerialize, nil
}

type NullableUpdateRoomsQuotaRequestDto struct {
	value *UpdateRoomsQuotaRequestDto
	isSet bool
}

func (v NullableUpdateRoomsQuotaRequestDto) Get() *UpdateRoomsQuotaRequestDto {
	return v.value
}

func (v *NullableUpdateRoomsQuotaRequestDto) Set(val *UpdateRoomsQuotaRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRoomsQuotaRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRoomsQuotaRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRoomsQuotaRequestDto(val *UpdateRoomsQuotaRequestDto) *NullableUpdateRoomsQuotaRequestDto {
	return &NullableUpdateRoomsQuotaRequestDto{value: val, isSet: true}
}

func (v NullableUpdateRoomsQuotaRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRoomsQuotaRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

