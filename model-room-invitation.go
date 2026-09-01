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

// checks if the RoomInvitation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomInvitation{}

// RoomInvitation The room invitation parameters.
type RoomInvitation struct {
	// The email address.
	Email *string `json:"email,omitempty"`
	// The ID of the user to share a room with.
	Id *string `json:"id,omitempty"`
	// The room sharing rights.
	Access *FileShare `json:"access,omitempty"`
}

// NewRoomInvitation instantiates a new RoomInvitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomInvitation() *RoomInvitation {
	this := RoomInvitation{}
	return &this
}

// NewRoomInvitationWithDefaults instantiates a new RoomInvitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomInvitationWithDefaults() *RoomInvitation {
	this := RoomInvitation{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *RoomInvitation) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomInvitation) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *RoomInvitation) IsEmailSet() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *RoomInvitation) SetEmail(v string) {
	o.Email = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoomInvitation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomInvitation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoomInvitation) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoomInvitation) SetId(v string) {
	o.Id = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *RoomInvitation) GetAccess() FileShare {
	if o == nil || IsNil(o.Access) {
		var ret FileShare
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomInvitation) GetAccessOk() (*FileShare, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *RoomInvitation) IsAccessSet() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given FileShare and assigns it to the Access field.
func (o *RoomInvitation) SetAccess(v FileShare) {
	o.Access = &v
}

func (o RoomInvitation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomInvitation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	return toSerialize, nil
}

type NullableRoomInvitation struct {
	value *RoomInvitation
	isSet bool
}

func (v NullableRoomInvitation) Get() *RoomInvitation {
	return v.value
}

func (v *NullableRoomInvitation) Set(val *RoomInvitation) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomInvitation(val *RoomInvitation) *NullableRoomInvitation {
	return &NullableRoomInvitation{value: val, isSet: true}
}

func (v NullableRoomInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

