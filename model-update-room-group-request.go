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

// checks if the UpdateRoomGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRoomGroupRequest{}

// UpdateRoomGroupRequest The changes to apply to a room group: its name and the rooms to add or remove.
type UpdateRoomGroupRequest struct {
	// The list of room IDs to add to the group.
	RoomsToAdd []DuplicateRequestDtoAllOfFileIds `json:"roomsToAdd,omitempty"`
	// The list of room IDs to remove from the group.
	RoomsToRemove []DuplicateRequestDtoAllOfFileIds `json:"roomsToRemove,omitempty"`
	// The group name.
	GroupName NullableString `json:"groupName,omitempty"`
}

// NewUpdateRoomGroupRequest instantiates a new UpdateRoomGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRoomGroupRequest() *UpdateRoomGroupRequest {
	this := UpdateRoomGroupRequest{}
	return &this
}

// NewUpdateRoomGroupRequestWithDefaults instantiates a new UpdateRoomGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRoomGroupRequestWithDefaults() *UpdateRoomGroupRequest {
	this := UpdateRoomGroupRequest{}
	return &this
}

// GetRoomsToAdd returns the RoomsToAdd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateRoomGroupRequest) GetRoomsToAdd() []DuplicateRequestDtoAllOfFileIds {
	if o == nil {
		var ret []DuplicateRequestDtoAllOfFileIds
		return ret
	}
	return o.RoomsToAdd
}

// GetRoomsToAddOk returns a tuple with the RoomsToAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateRoomGroupRequest) GetRoomsToAddOk() ([]DuplicateRequestDtoAllOfFileIds, bool) {
	if o == nil || IsNil(o.RoomsToAdd) {
		return nil, false
	}
	return o.RoomsToAdd, true
}

// HasRoomsToAdd returns a boolean if a field has been set.
func (o *UpdateRoomGroupRequest) IsRoomsToAddSet() bool {
	if o != nil && !IsNil(o.RoomsToAdd) {
		return true
	}

	return false
}

// SetRoomsToAdd gets a reference to the given []DuplicateRequestDtoAllOfFileIds and assigns it to the RoomsToAdd field.
func (o *UpdateRoomGroupRequest) SetRoomsToAdd(v []DuplicateRequestDtoAllOfFileIds) {
	o.RoomsToAdd = v
}

// GetRoomsToRemove returns the RoomsToRemove field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateRoomGroupRequest) GetRoomsToRemove() []DuplicateRequestDtoAllOfFileIds {
	if o == nil {
		var ret []DuplicateRequestDtoAllOfFileIds
		return ret
	}
	return o.RoomsToRemove
}

// GetRoomsToRemoveOk returns a tuple with the RoomsToRemove field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateRoomGroupRequest) GetRoomsToRemoveOk() ([]DuplicateRequestDtoAllOfFileIds, bool) {
	if o == nil || IsNil(o.RoomsToRemove) {
		return nil, false
	}
	return o.RoomsToRemove, true
}

// HasRoomsToRemove returns a boolean if a field has been set.
func (o *UpdateRoomGroupRequest) IsRoomsToRemoveSet() bool {
	if o != nil && !IsNil(o.RoomsToRemove) {
		return true
	}

	return false
}

// SetRoomsToRemove gets a reference to the given []DuplicateRequestDtoAllOfFileIds and assigns it to the RoomsToRemove field.
func (o *UpdateRoomGroupRequest) SetRoomsToRemove(v []DuplicateRequestDtoAllOfFileIds) {
	o.RoomsToRemove = v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateRoomGroupRequest) GetGroupName() string {
	if o == nil || IsNil(o.GroupName.Get()) {
		var ret string
		return ret
	}
	return *o.GroupName.Get()
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateRoomGroupRequest) GetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupName.Get(), o.GroupName.IsSet()
}

// HasGroupName returns a boolean if a field has been set.
func (o *UpdateRoomGroupRequest) IsGroupNameSet() bool {
	if o != nil && o.GroupName.IsSet() {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given NullableString and assigns it to the GroupName field.
func (o *UpdateRoomGroupRequest) SetGroupName(v string) {
	o.GroupName.Set(&v)
}
// SetGroupNameNil sets the value for GroupName to be an explicit nil
func (o *UpdateRoomGroupRequest) SetGroupNameNil() {
	o.GroupName.Set(nil)
}

// UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
func (o *UpdateRoomGroupRequest) UnsetGroupName() {
	o.GroupName.Unset()
}

func (o UpdateRoomGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRoomGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RoomsToAdd != nil {
		toSerialize["roomsToAdd"] = o.RoomsToAdd
	}
	if o.RoomsToRemove != nil {
		toSerialize["roomsToRemove"] = o.RoomsToRemove
	}
	if o.GroupName.IsSet() {
		toSerialize["groupName"] = o.GroupName.Get()
	}
	return toSerialize, nil
}

type NullableUpdateRoomGroupRequest struct {
	value *UpdateRoomGroupRequest
	isSet bool
}

func (v NullableUpdateRoomGroupRequest) Get() *UpdateRoomGroupRequest {
	return v.value
}

func (v *NullableUpdateRoomGroupRequest) Set(val *UpdateRoomGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRoomGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRoomGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRoomGroupRequest(val *UpdateRoomGroupRequest) *NullableUpdateRoomGroupRequest {
	return &NullableUpdateRoomGroupRequest{value: val, isSet: true}
}

func (v NullableUpdateRoomGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRoomGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

