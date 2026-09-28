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

// checks if the RoomGroupDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomGroupDto{}

// RoomGroupDto A personal collection of rooms: the name and icon it was given, the account that owns it, and the rooms it gathers  at the moment it was read.
type RoomGroupDto struct {
	// The identifier of the group, which addresses it in every other group operation and is kept for as long as the  group exists.
	Id *int32 `json:"id,omitempty"`
	// The name its owner gave the group, stored trimmed of surrounding spaces. Names are not unique, so two groups  of the same account can be told apart only by their identifier.
	Name NullableString `json:"name,omitempty"`
	// The built-in cover chosen for the group, carrying the cover identifier and its rendering in each available  size. Null when the group has no icon, either because it was never given one or because the icon was cleared  by setting it to an empty value.
	Icon *MultiSizeLogoCover `json:"icon,omitempty"`
	// The account that created the group and the only one able to read, change or delete it; for any other member of  the portal the group does not exist.
	UserId *string `json:"userId,omitempty"`
	// The section the group belongs to, which categorizes it within the application's structure. This property determines  which area of the interface the group is associated with and affects how its rooms are filtered and displayed.  Common values include Active for standard rooms, Forms for form-based rooms, Archive for archived content, and  Templates for template rooms. The search area ensures that when retrieving a group, only rooms that belong to  the specified section are included in the results, maintaining proper organizational boundaries within the system.
	SearchArea *SearchArea `json:"searchArea,omitempty"`
	// The rooms the group gathers, those stored in the portal first and those on connected third-party accounts  after them. Null when the group was asked for without its members, and an empty array when the group holds no  room the caller can still see. A room moved to the archive is left out until it is taken out of the archive.
	Rooms []FileEntryBaseDto `json:"rooms,omitempty"`
	// How many rooms the group shows: the same rooms `rooms` lists, so archived ones are not counted either. It is  filled even when the rooms themselves were not asked for, which makes it the cheap way to tell an empty group  from a populated one.
	TotalRooms *int32 `json:"totalRooms,omitempty"`
}

// NewRoomGroupDto instantiates a new RoomGroupDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomGroupDto() *RoomGroupDto {
	this := RoomGroupDto{}
	return &this
}

// NewRoomGroupDtoWithDefaults instantiates a new RoomGroupDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomGroupDtoWithDefaults() *RoomGroupDto {
	this := RoomGroupDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoomGroupDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomGroupDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoomGroupDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RoomGroupDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomGroupDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomGroupDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *RoomGroupDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *RoomGroupDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *RoomGroupDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *RoomGroupDto) UnsetName() {
	o.Name.Unset()
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *RoomGroupDto) GetIcon() MultiSizeLogoCover {
	if o == nil || IsNil(o.Icon) {
		var ret MultiSizeLogoCover
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomGroupDto) GetIconOk() (*MultiSizeLogoCover, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *RoomGroupDto) IsIconSet() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given MultiSizeLogoCover and assigns it to the Icon field.
func (o *RoomGroupDto) SetIcon(v MultiSizeLogoCover) {
	o.Icon = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *RoomGroupDto) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomGroupDto) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *RoomGroupDto) IsUserIdSet() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *RoomGroupDto) SetUserId(v string) {
	o.UserId = &v
}

// GetSearchArea returns the SearchArea field value if set, zero value otherwise.
func (o *RoomGroupDto) GetSearchArea() SearchArea {
	if o == nil || IsNil(o.SearchArea) {
		var ret SearchArea
		return ret
	}
	return *o.SearchArea
}

// GetSearchAreaOk returns a tuple with the SearchArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomGroupDto) GetSearchAreaOk() (*SearchArea, bool) {
	if o == nil || IsNil(o.SearchArea) {
		return nil, false
	}
	return o.SearchArea, true
}

// HasSearchArea returns a boolean if a field has been set.
func (o *RoomGroupDto) IsSearchAreaSet() bool {
	if o != nil && !IsNil(o.SearchArea) {
		return true
	}

	return false
}

// SetSearchArea gets a reference to the given SearchArea and assigns it to the SearchArea field.
func (o *RoomGroupDto) SetSearchArea(v SearchArea) {
	o.SearchArea = &v
}

// GetRooms returns the Rooms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomGroupDto) GetRooms() []FileEntryBaseDto {
	if o == nil {
		var ret []FileEntryBaseDto
		return ret
	}
	return o.Rooms
}

// GetRoomsOk returns a tuple with the Rooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomGroupDto) GetRoomsOk() ([]FileEntryBaseDto, bool) {
	if o == nil || IsNil(o.Rooms) {
		return nil, false
	}
	return o.Rooms, true
}

// HasRooms returns a boolean if a field has been set.
func (o *RoomGroupDto) IsRoomsSet() bool {
	if o != nil && !IsNil(o.Rooms) {
		return true
	}

	return false
}

// SetRooms gets a reference to the given []FileEntryBaseDto and assigns it to the Rooms field.
func (o *RoomGroupDto) SetRooms(v []FileEntryBaseDto) {
	o.Rooms = v
}

// GetTotalRooms returns the TotalRooms field value if set, zero value otherwise.
func (o *RoomGroupDto) GetTotalRooms() int32 {
	if o == nil || IsNil(o.TotalRooms) {
		var ret int32
		return ret
	}
	return *o.TotalRooms
}

// GetTotalRoomsOk returns a tuple with the TotalRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomGroupDto) GetTotalRoomsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRooms) {
		return nil, false
	}
	return o.TotalRooms, true
}

// HasTotalRooms returns a boolean if a field has been set.
func (o *RoomGroupDto) IsTotalRoomsSet() bool {
	if o != nil && !IsNil(o.TotalRooms) {
		return true
	}

	return false
}

// SetTotalRooms gets a reference to the given int32 and assigns it to the TotalRooms field.
func (o *RoomGroupDto) SetTotalRooms(v int32) {
	o.TotalRooms = &v
}

func (o RoomGroupDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomGroupDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.SearchArea) {
		toSerialize["searchArea"] = o.SearchArea
	}
	if o.Rooms != nil {
		toSerialize["rooms"] = o.Rooms
	}
	if !IsNil(o.TotalRooms) {
		toSerialize["totalRooms"] = o.TotalRooms
	}
	return toSerialize, nil
}

type NullableRoomGroupDto struct {
	value *RoomGroupDto
	isSet bool
}

func (v NullableRoomGroupDto) Get() *RoomGroupDto {
	return v.value
}

func (v *NullableRoomGroupDto) Set(val *RoomGroupDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomGroupDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomGroupDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomGroupDto(val *RoomGroupDto) *NullableRoomGroupDto {
	return &NullableRoomGroupDto{value: val, isSet: true}
}

func (v NullableRoomGroupDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomGroupDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

