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

// checks if the UpdateRoomsQuotaRequestDtoInteger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRoomsQuotaRequestDtoInteger{}

// UpdateRoomsQuotaRequestDtoInteger The request parameters for updating the room quota.
type UpdateRoomsQuotaRequestDtoInteger struct {
	// The list of room IDs.
	RoomIds []ContinueChatBodyFilesInner `json:"roomIds,omitempty"`
	// The room quota.
	Quota *int64 `json:"quota,omitempty"`
}

// NewUpdateRoomsQuotaRequestDtoInteger instantiates a new UpdateRoomsQuotaRequestDtoInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRoomsQuotaRequestDtoInteger() *UpdateRoomsQuotaRequestDtoInteger {
	this := UpdateRoomsQuotaRequestDtoInteger{}
	return &this
}

// NewUpdateRoomsQuotaRequestDtoIntegerWithDefaults instantiates a new UpdateRoomsQuotaRequestDtoInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoomsQuotaRequestDtoIntegerWithDefaults() *UpdateRoomsQuotaRequestDtoInteger {
	this := UpdateRoomsQuotaRequestDtoInteger{}
	return &this
}

// GetRoomIds returns the RoomIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateRoomsQuotaRequestDtoInteger) GetRoomIds() []ContinueChatBodyFilesInner {
	if o == nil {
		var ret []ContinueChatBodyFilesInner
		return ret
	}
	return o.RoomIds
}

// GetRoomIdsOk returns a tuple with the RoomIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateRoomsQuotaRequestDtoInteger) GetRoomIdsOk() ([]ContinueChatBodyFilesInner, bool) {
	if o == nil || IsNil(o.RoomIds) {
		return nil, false
	}
	return o.RoomIds, true
}

// HasRoomIds returns a boolean if a field has been set.
func (o *UpdateRoomsQuotaRequestDtoInteger) IsRoomIdsSet() bool {
	if o != nil && !IsNil(o.RoomIds) {
		return true
	}

	return false
}

// SetRoomIds gets a reference to the given []ContinueChatBodyFilesInner and assigns it to the RoomIds field.
func (o *UpdateRoomsQuotaRequestDtoInteger) SetRoomIds(v []ContinueChatBodyFilesInner) {
	o.RoomIds = v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *UpdateRoomsQuotaRequestDtoInteger) GetQuota() int64 {
	if o == nil || IsNil(o.Quota) {
		var ret int64
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRoomsQuotaRequestDtoInteger) GetQuotaOk() (*int64, bool) {
	if o == nil || IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *UpdateRoomsQuotaRequestDtoInteger) IsQuotaSet() bool {
	if o != nil && !IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given int64 and assigns it to the Quota field.
func (o *UpdateRoomsQuotaRequestDtoInteger) SetQuota(v int64) {
	o.Quota = &v
}

func (o UpdateRoomsQuotaRequestDtoInteger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRoomsQuotaRequestDtoInteger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RoomIds != nil {
		toSerialize["roomIds"] = o.RoomIds
	}
	if !IsNil(o.Quota) {
		toSerialize["quota"] = o.Quota
	}
	return toSerialize, nil
}

type NullableUpdateRoomsQuotaRequestDtoInteger struct {
	value *UpdateRoomsQuotaRequestDtoInteger
	isSet bool
}

func (v NullableUpdateRoomsQuotaRequestDtoInteger) Get() *UpdateRoomsQuotaRequestDtoInteger {
	return v.value
}

func (v *NullableUpdateRoomsQuotaRequestDtoInteger) Set(val *UpdateRoomsQuotaRequestDtoInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRoomsQuotaRequestDtoInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRoomsQuotaRequestDtoInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRoomsQuotaRequestDtoInteger(val *UpdateRoomsQuotaRequestDtoInteger) *NullableUpdateRoomsQuotaRequestDtoInteger {
	return &NullableUpdateRoomsQuotaRequestDtoInteger{value: val, isSet: true}
}

func (v NullableUpdateRoomsQuotaRequestDtoInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRoomsQuotaRequestDtoInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

