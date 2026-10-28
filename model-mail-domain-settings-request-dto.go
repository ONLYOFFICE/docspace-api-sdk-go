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

// checks if the MailDomainSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MailDomainSettingsRequestDto{}

// MailDomainSettingsRequestDto Which email domains the portal treats as already verified, and how their users join.
type MailDomainSettingsRequestDto struct {
	// How trusted domains are decided: no domain is trusted, every domain is, or only the ones listed in `domains`.  Only the custom mode reads `domains`; under the other two the list is ignored rather than refused.
	Type TenantTrustedDomainsType `json:"type"`
	// The trusted domains, as bare hostnames such as `example.com` without a scheme or an `@`. This is the whole  list that is to hold afterwards and not a list of additions. Each entry is lowercased before it is stored,  and one entry that is not a valid hostname - or an empty list in the custom mode - fails the whole call  without saving anything.
	Domains []string `json:"domains"`
	// What a user joining through a trusted domain becomes: `true` admits them as a guest, `false` as a full  member. It applies to joins made from now on and does not change anybody who has already joined.
	InviteUsersAsVisitors bool `json:"inviteUsersAsVisitors"`
}

type _MailDomainSettingsRequestDto MailDomainSettingsRequestDto

// NewMailDomainSettingsRequestDto instantiates a new MailDomainSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMailDomainSettingsRequestDto(type_ TenantTrustedDomainsType, domains []string, inviteUsersAsVisitors bool) *MailDomainSettingsRequestDto {
	this := MailDomainSettingsRequestDto{}
	this.Type = type_
	this.Domains = domains
	this.InviteUsersAsVisitors = inviteUsersAsVisitors
	return &this
}

// NewMailDomainSettingsRequestDtoWithDefaults instantiates a new MailDomainSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMailDomainSettingsRequestDtoWithDefaults() *MailDomainSettingsRequestDto {
	this := MailDomainSettingsRequestDto{}
	return &this
}

// GetType returns the Type field value
func (o *MailDomainSettingsRequestDto) GetType() TenantTrustedDomainsType {
	if o == nil {
		var ret TenantTrustedDomainsType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MailDomainSettingsRequestDto) GetTypeOk() (*TenantTrustedDomainsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MailDomainSettingsRequestDto) SetType(v TenantTrustedDomainsType) {
	o.Type = v
}

// GetDomains returns the Domains field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *MailDomainSettingsRequestDto) GetDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MailDomainSettingsRequestDto) GetDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// SetDomains sets field value
func (o *MailDomainSettingsRequestDto) SetDomains(v []string) {
	o.Domains = v
}

// GetInviteUsersAsVisitors returns the InviteUsersAsVisitors field value
func (o *MailDomainSettingsRequestDto) GetInviteUsersAsVisitors() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InviteUsersAsVisitors
}

// GetInviteUsersAsVisitorsOk returns a tuple with the InviteUsersAsVisitors field value
// and a boolean to check if the value has been set.
func (o *MailDomainSettingsRequestDto) GetInviteUsersAsVisitorsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InviteUsersAsVisitors, true
}

// SetInviteUsersAsVisitors sets field value
func (o *MailDomainSettingsRequestDto) SetInviteUsersAsVisitors(v bool) {
	o.InviteUsersAsVisitors = v
}

func (o MailDomainSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MailDomainSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}
	toSerialize["inviteUsersAsVisitors"] = o.InviteUsersAsVisitors
	return toSerialize, nil
}

func (o *MailDomainSettingsRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"domains",
		"inviteUsersAsVisitors",
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

	varMailDomainSettingsRequestDto := _MailDomainSettingsRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varMailDomainSettingsRequestDto)

	if err != nil {
		return err
	}

	*o = MailDomainSettingsRequestDto(varMailDomainSettingsRequestDto)

	return err
}

type NullableMailDomainSettingsRequestDto struct {
	value *MailDomainSettingsRequestDto
	isSet bool
}

func (v NullableMailDomainSettingsRequestDto) Get() *MailDomainSettingsRequestDto {
	return v.value
}

func (v *NullableMailDomainSettingsRequestDto) Set(val *MailDomainSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMailDomainSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMailDomainSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMailDomainSettingsRequestDto(val *MailDomainSettingsRequestDto) *NullableMailDomainSettingsRequestDto {
	return &NullableMailDomainSettingsRequestDto{value: val, isSet: true}
}

func (v NullableMailDomainSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMailDomainSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

