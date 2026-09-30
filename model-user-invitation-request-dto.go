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

// checks if the UserInvitationRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInvitationRequestDto{}

// UserInvitationRequestDto The user invitation parameters.
type UserInvitationRequestDto struct {
	// The address of somebody who has no portal account yet. An invitation is sent to it and an account is created  once it is accepted, so this is the field to use instead of an account identifier when the person is new to  the portal.
	Email *string `json:"email,omitempty"`
	// The user type.
	Type *EmployeeType `json:"type,omitempty"`
}

// NewUserInvitationRequestDto instantiates a new UserInvitationRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInvitationRequestDto() *UserInvitationRequestDto {
	this := UserInvitationRequestDto{}
	return &this
}

// NewUserInvitationRequestDtoWithDefaults instantiates a new UserInvitationRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInvitationRequestDtoWithDefaults() *UserInvitationRequestDto {
	this := UserInvitationRequestDto{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserInvitationRequestDto) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvitationRequestDto) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserInvitationRequestDto) IsEmailSet() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserInvitationRequestDto) SetEmail(v string) {
	o.Email = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserInvitationRequestDto) GetType() EmployeeType {
	if o == nil || IsNil(o.Type) {
		var ret EmployeeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvitationRequestDto) GetTypeOk() (*EmployeeType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserInvitationRequestDto) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EmployeeType and assigns it to the Type field.
func (o *UserInvitationRequestDto) SetType(v EmployeeType) {
	o.Type = &v
}

func (o UserInvitationRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInvitationRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableUserInvitationRequestDto struct {
	value *UserInvitationRequestDto
	isSet bool
}

func (v NullableUserInvitationRequestDto) Get() *UserInvitationRequestDto {
	return v.value
}

func (v *NullableUserInvitationRequestDto) Set(val *UserInvitationRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInvitationRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInvitationRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInvitationRequestDto(val *UserInvitationRequestDto) *NullableUserInvitationRequestDto {
	return &NullableUserInvitationRequestDto{value: val, isSet: true}
}

func (v NullableUserInvitationRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInvitationRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

