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
	"bytes"
	"fmt"
)

// checks if the InviteUsersRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InviteUsersRequestDto{}

// InviteUsersRequestDto The request parameters for inviting users.
type InviteUsersRequestDto struct {
	// The list of user invitations.
	Invitations []UserInvitationRequestDto `json:"invitations"`
	// The culture code of invitations.
	Culture NullableString `json:"culture,omitempty"`
}

type _InviteUsersRequestDto InviteUsersRequestDto

// NewInviteUsersRequestDto instantiates a new InviteUsersRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteUsersRequestDto(invitations []UserInvitationRequestDto) *InviteUsersRequestDto {
	this := InviteUsersRequestDto{}
	this.Invitations = invitations
	return &this
}

// NewInviteUsersRequestDtoWithDefaults instantiates a new InviteUsersRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteUsersRequestDtoWithDefaults() *InviteUsersRequestDto {
	this := InviteUsersRequestDto{}
	return &this
}

// GetInvitations returns the Invitations field value
func (o *InviteUsersRequestDto) GetInvitations() []UserInvitationRequestDto {
	if o == nil {
		var ret []UserInvitationRequestDto
		return ret
	}

	return o.Invitations
}

// GetInvitationsOk returns a tuple with the Invitations field value
// and a boolean to check if the value has been set.
func (o *InviteUsersRequestDto) GetInvitationsOk() ([]UserInvitationRequestDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Invitations, true
}

// SetInvitations sets field value
func (o *InviteUsersRequestDto) SetInvitations(v []UserInvitationRequestDto) {
	o.Invitations = v
}

// GetCulture returns the Culture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteUsersRequestDto) GetCulture() string {
	if o == nil || IsNil(o.Culture.Get()) {
		var ret string
		return ret
	}
	return *o.Culture.Get()
}

// GetCultureOk returns a tuple with the Culture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteUsersRequestDto) GetCultureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Culture.Get(), o.Culture.IsSet()
}

// HasCulture returns a boolean if a field has been set.
func (o *InviteUsersRequestDto) IsCultureSet() bool {
	if o != nil && o.Culture.IsSet() {
		return true
	}

	return false
}

// SetCulture gets a reference to the given NullableString and assigns it to the Culture field.
func (o *InviteUsersRequestDto) SetCulture(v string) {
	o.Culture.Set(&v)
}
// SetCultureNil sets the value for Culture to be an explicit nil
func (o *InviteUsersRequestDto) SetCultureNil() {
	o.Culture.Set(nil)
}

// UnsetCulture ensures that no value is present for Culture, not even an explicit nil
func (o *InviteUsersRequestDto) UnsetCulture() {
	o.Culture.Unset()
}

func (o InviteUsersRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InviteUsersRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invitations"] = o.Invitations
	if o.Culture.IsSet() {
		toSerialize["culture"] = o.Culture.Get()
	}
	return toSerialize, nil
}

func (o *InviteUsersRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invitations",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varInviteUsersRequestDto := _InviteUsersRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varInviteUsersRequestDto)

	if err != nil {
		return err
	}

	*o = InviteUsersRequestDto(varInviteUsersRequestDto)

	return err
}

type NullableInviteUsersRequestDto struct {
	value *InviteUsersRequestDto
	isSet bool
}

func (v NullableInviteUsersRequestDto) Get() *InviteUsersRequestDto {
	return v.value
}

func (v *NullableInviteUsersRequestDto) Set(val *InviteUsersRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteUsersRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteUsersRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteUsersRequestDto(val *InviteUsersRequestDto) *NullableInviteUsersRequestDto {
	return &NullableInviteUsersRequestDto{value: val, isSet: true}
}

func (v NullableInviteUsersRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteUsersRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

