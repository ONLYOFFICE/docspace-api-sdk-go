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

// checks if the MailDomainSettingsRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MailDomainSettingsRequestsDto{}

// MailDomainSettingsRequestsDto Which email domains the portal treats as already verified, and how their users join.
type MailDomainSettingsRequestsDto struct {
	// How trusted domains are decided: no domain is trusted, every domain is, or only the ones listed in `domains`.  Only the custom mode reads `domains`; under the other two the list is ignored rather than refused.
	Type TenantTrustedDomainsType `json:"type"`
	// The trusted domains, as bare hostnames such as `example.com` without a scheme or an `@`. This is the whole  list that is to hold afterwards and not a list of additions. Each entry is lowercased before it is stored,  and one entry that is not a valid hostname - or an empty list in the custom mode - fails the whole call  without saving anything.
	Domains []string `json:"domains"`
	// What a user joining through a trusted domain becomes: `true` admits them as a guest, `false` as a full  member. It applies to joins made from now on and does not change anybody who has already joined.
	InviteUsersAsVisitors bool `json:"inviteUsersAsVisitors"`
}

type _MailDomainSettingsRequestsDto MailDomainSettingsRequestsDto

// NewMailDomainSettingsRequestsDto instantiates a new MailDomainSettingsRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMailDomainSettingsRequestsDto(type_ TenantTrustedDomainsType, domains []string, inviteUsersAsVisitors bool) *MailDomainSettingsRequestsDto {
	this := MailDomainSettingsRequestsDto{}
	this.Type = type_
	this.Domains = domains
	this.InviteUsersAsVisitors = inviteUsersAsVisitors
	return &this
}

// NewMailDomainSettingsRequestsDtoWithDefaults instantiates a new MailDomainSettingsRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMailDomainSettingsRequestsDtoWithDefaults() *MailDomainSettingsRequestsDto {
	this := MailDomainSettingsRequestsDto{}
	return &this
}

// GetType returns the Type field value
func (o *MailDomainSettingsRequestsDto) GetType() TenantTrustedDomainsType {
	if o == nil {
		var ret TenantTrustedDomainsType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MailDomainSettingsRequestsDto) GetTypeOk() (*TenantTrustedDomainsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MailDomainSettingsRequestsDto) SetType(v TenantTrustedDomainsType) {
	o.Type = v
}

// GetDomains returns the Domains field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *MailDomainSettingsRequestsDto) GetDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MailDomainSettingsRequestsDto) GetDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// SetDomains sets field value
func (o *MailDomainSettingsRequestsDto) SetDomains(v []string) {
	o.Domains = v
}

// GetInviteUsersAsVisitors returns the InviteUsersAsVisitors field value
func (o *MailDomainSettingsRequestsDto) GetInviteUsersAsVisitors() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InviteUsersAsVisitors
}

// GetInviteUsersAsVisitorsOk returns a tuple with the InviteUsersAsVisitors field value
// and a boolean to check if the value has been set.
func (o *MailDomainSettingsRequestsDto) GetInviteUsersAsVisitorsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InviteUsersAsVisitors, true
}

// SetInviteUsersAsVisitors sets field value
func (o *MailDomainSettingsRequestsDto) SetInviteUsersAsVisitors(v bool) {
	o.InviteUsersAsVisitors = v
}

func (o MailDomainSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MailDomainSettingsRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}
	toSerialize["inviteUsersAsVisitors"] = o.InviteUsersAsVisitors
	return toSerialize, nil
}

func (o *MailDomainSettingsRequestsDto) UnmarshalJSON(data []byte) (err error) {
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

	varMailDomainSettingsRequestsDto := _MailDomainSettingsRequestsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varMailDomainSettingsRequestsDto)

	if err != nil {
		return err
	}

	*o = MailDomainSettingsRequestsDto(varMailDomainSettingsRequestsDto)

	return err
}

type NullableMailDomainSettingsRequestsDto struct {
	value *MailDomainSettingsRequestsDto
	isSet bool
}

func (v NullableMailDomainSettingsRequestsDto) Get() *MailDomainSettingsRequestsDto {
	return v.value
}

func (v *NullableMailDomainSettingsRequestsDto) Set(val *MailDomainSettingsRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMailDomainSettingsRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMailDomainSettingsRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMailDomainSettingsRequestsDto(val *MailDomainSettingsRequestsDto) *NullableMailDomainSettingsRequestsDto {
	return &NullableMailDomainSettingsRequestsDto{value: val, isSet: true}
}

func (v NullableMailDomainSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMailDomainSettingsRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

