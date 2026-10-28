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

// checks if the UserInvitation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInvitation{}

// UserInvitation Which pending room invitations are to be sent again.
type UserInvitation struct {
	// The accounts to write to, taken from `GET api/2.0/files/rooms/{id}/share`. Anyone who has already joined, is  not in the room, or is invisible to the caller is skipped without an error, and the field is ignored once  every pending invitation is being resent.
	UsersIds []string `json:"usersIds,omitempty"`
	// Whether every invitation of the room that is still waiting is sent again. With it on the list of accounts is  ignored, and with it off an empty list means that nothing is sent at all.
	ResendAll *bool `json:"resendAll,omitempty"`
}

// NewUserInvitation instantiates a new UserInvitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInvitation() *UserInvitation {
	this := UserInvitation{}
	return &this
}

// NewUserInvitationWithDefaults instantiates a new UserInvitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInvitationWithDefaults() *UserInvitation {
	this := UserInvitation{}
	return &this
}

// GetUsersIds returns the UsersIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInvitation) GetUsersIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UsersIds
}

// GetUsersIdsOk returns a tuple with the UsersIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInvitation) GetUsersIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.UsersIds) {
		return nil, false
	}
	return o.UsersIds, true
}

// HasUsersIds returns a boolean if a field has been set.
func (o *UserInvitation) IsUsersIdsSet() bool {
	if o != nil && !IsNil(o.UsersIds) {
		return true
	}

	return false
}

// SetUsersIds gets a reference to the given []string and assigns it to the UsersIds field.
func (o *UserInvitation) SetUsersIds(v []string) {
	o.UsersIds = v
}

// GetResendAll returns the ResendAll field value if set, zero value otherwise.
func (o *UserInvitation) GetResendAll() bool {
	if o == nil || IsNil(o.ResendAll) {
		var ret bool
		return ret
	}
	return *o.ResendAll
}

// GetResendAllOk returns a tuple with the ResendAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvitation) GetResendAllOk() (*bool, bool) {
	if o == nil || IsNil(o.ResendAll) {
		return nil, false
	}
	return o.ResendAll, true
}

// HasResendAll returns a boolean if a field has been set.
func (o *UserInvitation) IsResendAllSet() bool {
	if o != nil && !IsNil(o.ResendAll) {
		return true
	}

	return false
}

// SetResendAll gets a reference to the given bool and assigns it to the ResendAll field.
func (o *UserInvitation) SetResendAll(v bool) {
	o.ResendAll = &v
}

func (o UserInvitation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInvitation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UsersIds != nil {
		toSerialize["usersIds"] = o.UsersIds
	}
	if !IsNil(o.ResendAll) {
		toSerialize["resendAll"] = o.ResendAll
	}
	return toSerialize, nil
}

type NullableUserInvitation struct {
	value *UserInvitation
	isSet bool
}

func (v NullableUserInvitation) Get() *UserInvitation {
	return v.value
}

func (v *NullableUserInvitation) Set(val *UserInvitation) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInvitation(val *UserInvitation) *NullableUserInvitation {
	return &NullableUserInvitation{value: val, isSet: true}
}

func (v NullableUserInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

