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

// checks if the EmailInvitationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailInvitationDto{}

// EmailInvitationDto The email invitation parameters.
type EmailInvitationDto struct {
	// The address of somebody who has no portal account yet. An invitation is sent to it and an account is created  once it is accepted, so this is the field to use instead of an account identifier when the person is new to  the portal.
	Email NullableString `json:"email,omitempty"`
}

// NewEmailInvitationDto instantiates a new EmailInvitationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailInvitationDto() *EmailInvitationDto {
	this := EmailInvitationDto{}
	return &this
}

// NewEmailInvitationDtoWithDefaults instantiates a new EmailInvitationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailInvitationDtoWithDefaults() *EmailInvitationDto {
	this := EmailInvitationDto{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailInvitationDto) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailInvitationDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *EmailInvitationDto) IsEmailSet() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *EmailInvitationDto) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *EmailInvitationDto) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *EmailInvitationDto) UnsetEmail() {
	o.Email.Unset()
}

func (o EmailInvitationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailInvitationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	return toSerialize, nil
}

type NullableEmailInvitationDto struct {
	value *EmailInvitationDto
	isSet bool
}

func (v NullableEmailInvitationDto) Get() *EmailInvitationDto {
	return v.value
}

func (v *NullableEmailInvitationDto) Set(val *EmailInvitationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailInvitationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailInvitationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailInvitationDto(val *EmailInvitationDto) *NullableEmailInvitationDto {
	return &NullableEmailInvitationDto{value: val, isSet: true}
}

func (v NullableEmailInvitationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailInvitationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

