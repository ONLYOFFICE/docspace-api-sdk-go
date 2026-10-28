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

// checks if the AiFileEntryDtoAllOfAvailableShareRights type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiFileEntryDtoAllOfAvailableShareRights{}

// AiFileEntryDtoAllOfAvailableShareRights Which access levels may be handed out on this entry, listed per kind of recipient, so that a client offers  only levels the entry actually supports - a room for filling forms and a plain folder do not accept the same  ones.
type AiFileEntryDtoAllOfAvailableShareRights struct {
	User []string `json:"User,omitempty"`
	ExternalLink []string `json:"ExternalLink,omitempty"`
	Group []string `json:"Group,omitempty"`
	InvitationLink []string `json:"InvitationLink,omitempty"`
	PrimaryExternalLink []string `json:"PrimaryExternalLink,omitempty"`
}

// NewAiFileEntryDtoAllOfAvailableShareRights instantiates a new AiFileEntryDtoAllOfAvailableShareRights object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiFileEntryDtoAllOfAvailableShareRights() *AiFileEntryDtoAllOfAvailableShareRights {
	this := AiFileEntryDtoAllOfAvailableShareRights{}
	return &this
}

// NewAiFileEntryDtoAllOfAvailableShareRightsWithDefaults instantiates a new AiFileEntryDtoAllOfAvailableShareRights object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiFileEntryDtoAllOfAvailableShareRightsWithDefaults() *AiFileEntryDtoAllOfAvailableShareRights {
	this := AiFileEntryDtoAllOfAvailableShareRights{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfAvailableShareRights) GetUser() []string {
	if o == nil || IsNil(o.User) {
		var ret []string
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfAvailableShareRights) GetUserOk() ([]string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfAvailableShareRights) IsUserSet() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given []string and assigns it to the User field.
func (o *AiFileEntryDtoAllOfAvailableShareRights) SetUser(v []string) {
	o.User = v
}

// GetExternalLink returns the ExternalLink field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfAvailableShareRights) GetExternalLink() []string {
	if o == nil || IsNil(o.ExternalLink) {
		var ret []string
		return ret
	}
	return o.ExternalLink
}

// GetExternalLinkOk returns a tuple with the ExternalLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfAvailableShareRights) GetExternalLinkOk() ([]string, bool) {
	if o == nil || IsNil(o.ExternalLink) {
		return nil, false
	}
	return o.ExternalLink, true
}

// HasExternalLink returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfAvailableShareRights) IsExternalLinkSet() bool {
	if o != nil && !IsNil(o.ExternalLink) {
		return true
	}

	return false
}

// SetExternalLink gets a reference to the given []string and assigns it to the ExternalLink field.
func (o *AiFileEntryDtoAllOfAvailableShareRights) SetExternalLink(v []string) {
	o.ExternalLink = v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfAvailableShareRights) GetGroup() []string {
	if o == nil || IsNil(o.Group) {
		var ret []string
		return ret
	}
	return o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfAvailableShareRights) GetGroupOk() ([]string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfAvailableShareRights) IsGroupSet() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given []string and assigns it to the Group field.
func (o *AiFileEntryDtoAllOfAvailableShareRights) SetGroup(v []string) {
	o.Group = v
}

// GetInvitationLink returns the InvitationLink field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfAvailableShareRights) GetInvitationLink() []string {
	if o == nil || IsNil(o.InvitationLink) {
		var ret []string
		return ret
	}
	return o.InvitationLink
}

// GetInvitationLinkOk returns a tuple with the InvitationLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfAvailableShareRights) GetInvitationLinkOk() ([]string, bool) {
	if o == nil || IsNil(o.InvitationLink) {
		return nil, false
	}
	return o.InvitationLink, true
}

// HasInvitationLink returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfAvailableShareRights) IsInvitationLinkSet() bool {
	if o != nil && !IsNil(o.InvitationLink) {
		return true
	}

	return false
}

// SetInvitationLink gets a reference to the given []string and assigns it to the InvitationLink field.
func (o *AiFileEntryDtoAllOfAvailableShareRights) SetInvitationLink(v []string) {
	o.InvitationLink = v
}

// GetPrimaryExternalLink returns the PrimaryExternalLink field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfAvailableShareRights) GetPrimaryExternalLink() []string {
	if o == nil || IsNil(o.PrimaryExternalLink) {
		var ret []string
		return ret
	}
	return o.PrimaryExternalLink
}

// GetPrimaryExternalLinkOk returns a tuple with the PrimaryExternalLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfAvailableShareRights) GetPrimaryExternalLinkOk() ([]string, bool) {
	if o == nil || IsNil(o.PrimaryExternalLink) {
		return nil, false
	}
	return o.PrimaryExternalLink, true
}

// HasPrimaryExternalLink returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfAvailableShareRights) IsPrimaryExternalLinkSet() bool {
	if o != nil && !IsNil(o.PrimaryExternalLink) {
		return true
	}

	return false
}

// SetPrimaryExternalLink gets a reference to the given []string and assigns it to the PrimaryExternalLink field.
func (o *AiFileEntryDtoAllOfAvailableShareRights) SetPrimaryExternalLink(v []string) {
	o.PrimaryExternalLink = v
}

func (o AiFileEntryDtoAllOfAvailableShareRights) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiFileEntryDtoAllOfAvailableShareRights) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["User"] = o.User
	}
	if !IsNil(o.ExternalLink) {
		toSerialize["ExternalLink"] = o.ExternalLink
	}
	if !IsNil(o.Group) {
		toSerialize["Group"] = o.Group
	}
	if !IsNil(o.InvitationLink) {
		toSerialize["InvitationLink"] = o.InvitationLink
	}
	if !IsNil(o.PrimaryExternalLink) {
		toSerialize["PrimaryExternalLink"] = o.PrimaryExternalLink
	}
	return toSerialize, nil
}

type NullableAiFileEntryDtoAllOfAvailableShareRights struct {
	value *AiFileEntryDtoAllOfAvailableShareRights
	isSet bool
}

func (v NullableAiFileEntryDtoAllOfAvailableShareRights) Get() *AiFileEntryDtoAllOfAvailableShareRights {
	return v.value
}

func (v *NullableAiFileEntryDtoAllOfAvailableShareRights) Set(val *AiFileEntryDtoAllOfAvailableShareRights) {
	v.value = val
	v.isSet = true
}

func (v NullableAiFileEntryDtoAllOfAvailableShareRights) IsSet() bool {
	return v.isSet
}

func (v *NullableAiFileEntryDtoAllOfAvailableShareRights) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiFileEntryDtoAllOfAvailableShareRights(val *AiFileEntryDtoAllOfAvailableShareRights) *NullableAiFileEntryDtoAllOfAvailableShareRights {
	return &NullableAiFileEntryDtoAllOfAvailableShareRights{value: val, isSet: true}
}

func (v NullableAiFileEntryDtoAllOfAvailableShareRights) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiFileEntryDtoAllOfAvailableShareRights) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

