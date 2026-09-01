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

// checks if the TenantUserInvitationSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantUserInvitationSettingsRequestDto{}

// TenantUserInvitationSettingsRequestDto The request parameters for updating the user invitation settings.
type TenantUserInvitationSettingsRequestDto struct {
	// Specifies whether to allow inviting new DocSpace members through the Contacts section.
	AllowInvitingMembers *bool `json:"allowInvitingMembers,omitempty"`
	// Specifies whether to allow all DocSpace members to invite external guests to the rooms.
	AllowInvitingGuests *bool `json:"allowInvitingGuests,omitempty"`
}

// NewTenantUserInvitationSettingsRequestDto instantiates a new TenantUserInvitationSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantUserInvitationSettingsRequestDto() *TenantUserInvitationSettingsRequestDto {
	this := TenantUserInvitationSettingsRequestDto{}
	return &this
}

// NewTenantUserInvitationSettingsRequestDtoWithDefaults instantiates a new TenantUserInvitationSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantUserInvitationSettingsRequestDtoWithDefaults() *TenantUserInvitationSettingsRequestDto {
	this := TenantUserInvitationSettingsRequestDto{}
	return &this
}

// GetAllowInvitingMembers returns the AllowInvitingMembers field value if set, zero value otherwise.
func (o *TenantUserInvitationSettingsRequestDto) GetAllowInvitingMembers() bool {
	if o == nil || IsNil(o.AllowInvitingMembers) {
		var ret bool
		return ret
	}
	return *o.AllowInvitingMembers
}

// GetAllowInvitingMembersOk returns a tuple with the AllowInvitingMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantUserInvitationSettingsRequestDto) GetAllowInvitingMembersOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowInvitingMembers) {
		return nil, false
	}
	return o.AllowInvitingMembers, true
}

// HasAllowInvitingMembers returns a boolean if a field has been set.
func (o *TenantUserInvitationSettingsRequestDto) IsAllowInvitingMembersSet() bool {
	if o != nil && !IsNil(o.AllowInvitingMembers) {
		return true
	}

	return false
}

// SetAllowInvitingMembers gets a reference to the given bool and assigns it to the AllowInvitingMembers field.
func (o *TenantUserInvitationSettingsRequestDto) SetAllowInvitingMembers(v bool) {
	o.AllowInvitingMembers = &v
}

// GetAllowInvitingGuests returns the AllowInvitingGuests field value if set, zero value otherwise.
func (o *TenantUserInvitationSettingsRequestDto) GetAllowInvitingGuests() bool {
	if o == nil || IsNil(o.AllowInvitingGuests) {
		var ret bool
		return ret
	}
	return *o.AllowInvitingGuests
}

// GetAllowInvitingGuestsOk returns a tuple with the AllowInvitingGuests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantUserInvitationSettingsRequestDto) GetAllowInvitingGuestsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowInvitingGuests) {
		return nil, false
	}
	return o.AllowInvitingGuests, true
}

// HasAllowInvitingGuests returns a boolean if a field has been set.
func (o *TenantUserInvitationSettingsRequestDto) IsAllowInvitingGuestsSet() bool {
	if o != nil && !IsNil(o.AllowInvitingGuests) {
		return true
	}

	return false
}

// SetAllowInvitingGuests gets a reference to the given bool and assigns it to the AllowInvitingGuests field.
func (o *TenantUserInvitationSettingsRequestDto) SetAllowInvitingGuests(v bool) {
	o.AllowInvitingGuests = &v
}

func (o TenantUserInvitationSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantUserInvitationSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowInvitingMembers) {
		toSerialize["allowInvitingMembers"] = o.AllowInvitingMembers
	}
	if !IsNil(o.AllowInvitingGuests) {
		toSerialize["allowInvitingGuests"] = o.AllowInvitingGuests
	}
	return toSerialize, nil
}

type NullableTenantUserInvitationSettingsRequestDto struct {
	value *TenantUserInvitationSettingsRequestDto
	isSet bool
}

func (v NullableTenantUserInvitationSettingsRequestDto) Get() *TenantUserInvitationSettingsRequestDto {
	return v.value
}

func (v *NullableTenantUserInvitationSettingsRequestDto) Set(val *TenantUserInvitationSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantUserInvitationSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantUserInvitationSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantUserInvitationSettingsRequestDto(val *TenantUserInvitationSettingsRequestDto) *NullableTenantUserInvitationSettingsRequestDto {
	return &NullableTenantUserInvitationSettingsRequestDto{value: val, isSet: true}
}

func (v NullableTenantUserInvitationSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantUserInvitationSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

