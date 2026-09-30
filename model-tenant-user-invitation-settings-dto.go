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

// checks if the TenantUserInvitationSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantUserInvitationSettingsDto{}

// TenantUserInvitationSettingsDto Whether the portal currently lets anyone be invited into it, member and guest kept apart.
type TenantUserInvitationSettingsDto struct {
	// Whether new members may be invited through the Contacts section. Switching it off stops new invitations  from being created; links already handed out keep working and members already invited stay.
	AllowInvitingMembers bool `json:"allowInvitingMembers"`
	// Whether every member, and not only an administrator, may invite an outside guest into a room. It is  independent of `allowInvitingMembers`, and switching it off has the same forward-only effect.
	AllowInvitingGuests bool `json:"allowInvitingGuests"`
}

type _TenantUserInvitationSettingsDto TenantUserInvitationSettingsDto

// NewTenantUserInvitationSettingsDto instantiates a new TenantUserInvitationSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantUserInvitationSettingsDto(allowInvitingMembers bool, allowInvitingGuests bool) *TenantUserInvitationSettingsDto {
	this := TenantUserInvitationSettingsDto{}
	this.AllowInvitingMembers = allowInvitingMembers
	this.AllowInvitingGuests = allowInvitingGuests
	return &this
}

// NewTenantUserInvitationSettingsDtoWithDefaults instantiates a new TenantUserInvitationSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantUserInvitationSettingsDtoWithDefaults() *TenantUserInvitationSettingsDto {
	this := TenantUserInvitationSettingsDto{}
	return &this
}

// GetAllowInvitingMembers returns the AllowInvitingMembers field value
func (o *TenantUserInvitationSettingsDto) GetAllowInvitingMembers() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowInvitingMembers
}

// GetAllowInvitingMembersOk returns a tuple with the AllowInvitingMembers field value
// and a boolean to check if the value has been set.
func (o *TenantUserInvitationSettingsDto) GetAllowInvitingMembersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowInvitingMembers, true
}

// SetAllowInvitingMembers sets field value
func (o *TenantUserInvitationSettingsDto) SetAllowInvitingMembers(v bool) {
	o.AllowInvitingMembers = v
}

// GetAllowInvitingGuests returns the AllowInvitingGuests field value
func (o *TenantUserInvitationSettingsDto) GetAllowInvitingGuests() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowInvitingGuests
}

// GetAllowInvitingGuestsOk returns a tuple with the AllowInvitingGuests field value
// and a boolean to check if the value has been set.
func (o *TenantUserInvitationSettingsDto) GetAllowInvitingGuestsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowInvitingGuests, true
}

// SetAllowInvitingGuests sets field value
func (o *TenantUserInvitationSettingsDto) SetAllowInvitingGuests(v bool) {
	o.AllowInvitingGuests = v
}

func (o TenantUserInvitationSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantUserInvitationSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allowInvitingMembers"] = o.AllowInvitingMembers
	toSerialize["allowInvitingGuests"] = o.AllowInvitingGuests
	return toSerialize, nil
}

func (o *TenantUserInvitationSettingsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"allowInvitingMembers",
		"allowInvitingGuests",
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

	varTenantUserInvitationSettingsDto := _TenantUserInvitationSettingsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varTenantUserInvitationSettingsDto)

	if err != nil {
		return err
	}

	*o = TenantUserInvitationSettingsDto(varTenantUserInvitationSettingsDto)

	return err
}

type NullableTenantUserInvitationSettingsDto struct {
	value *TenantUserInvitationSettingsDto
	isSet bool
}

func (v NullableTenantUserInvitationSettingsDto) Get() *TenantUserInvitationSettingsDto {
	return v.value
}

func (v *NullableTenantUserInvitationSettingsDto) Set(val *TenantUserInvitationSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantUserInvitationSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantUserInvitationSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantUserInvitationSettingsDto(val *TenantUserInvitationSettingsDto) *NullableTenantUserInvitationSettingsDto {
	return &NullableTenantUserInvitationSettingsDto{value: val, isSet: true}
}

func (v NullableTenantUserInvitationSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantUserInvitationSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

