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

// checks if the AiFileEntryDtoAllOfShareSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiFileEntryDtoAllOfShareSettings{}

// AiFileEntryDtoAllOfShareSettings How many links of each kind currently exist for the entry, counted separately for the primary link and the  additional ones. Kinds with no links are left out, and the whole field is null when the caller may not change  the access or no link exists at all.
type AiFileEntryDtoAllOfShareSettings struct {
	User *int32 `json:"User,omitempty"`
	ExternalLink *int32 `json:"ExternalLink,omitempty"`
	Group *int32 `json:"Group,omitempty"`
	InvitationLink *int32 `json:"InvitationLink,omitempty"`
	PrimaryExternalLink *int32 `json:"PrimaryExternalLink,omitempty"`
}

// NewAiFileEntryDtoAllOfShareSettings instantiates a new AiFileEntryDtoAllOfShareSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiFileEntryDtoAllOfShareSettings() *AiFileEntryDtoAllOfShareSettings {
	this := AiFileEntryDtoAllOfShareSettings{}
	return &this
}

// NewAiFileEntryDtoAllOfShareSettingsWithDefaults instantiates a new AiFileEntryDtoAllOfShareSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiFileEntryDtoAllOfShareSettingsWithDefaults() *AiFileEntryDtoAllOfShareSettings {
	this := AiFileEntryDtoAllOfShareSettings{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfShareSettings) GetUser() int32 {
	if o == nil || IsNil(o.User) {
		var ret int32
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfShareSettings) GetUserOk() (*int32, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfShareSettings) IsUserSet() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given int32 and assigns it to the User field.
func (o *AiFileEntryDtoAllOfShareSettings) SetUser(v int32) {
	o.User = &v
}

// GetExternalLink returns the ExternalLink field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfShareSettings) GetExternalLink() int32 {
	if o == nil || IsNil(o.ExternalLink) {
		var ret int32
		return ret
	}
	return *o.ExternalLink
}

// GetExternalLinkOk returns a tuple with the ExternalLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfShareSettings) GetExternalLinkOk() (*int32, bool) {
	if o == nil || IsNil(o.ExternalLink) {
		return nil, false
	}
	return o.ExternalLink, true
}

// HasExternalLink returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfShareSettings) IsExternalLinkSet() bool {
	if o != nil && !IsNil(o.ExternalLink) {
		return true
	}

	return false
}

// SetExternalLink gets a reference to the given int32 and assigns it to the ExternalLink field.
func (o *AiFileEntryDtoAllOfShareSettings) SetExternalLink(v int32) {
	o.ExternalLink = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfShareSettings) GetGroup() int32 {
	if o == nil || IsNil(o.Group) {
		var ret int32
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfShareSettings) GetGroupOk() (*int32, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfShareSettings) IsGroupSet() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given int32 and assigns it to the Group field.
func (o *AiFileEntryDtoAllOfShareSettings) SetGroup(v int32) {
	o.Group = &v
}

// GetInvitationLink returns the InvitationLink field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfShareSettings) GetInvitationLink() int32 {
	if o == nil || IsNil(o.InvitationLink) {
		var ret int32
		return ret
	}
	return *o.InvitationLink
}

// GetInvitationLinkOk returns a tuple with the InvitationLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfShareSettings) GetInvitationLinkOk() (*int32, bool) {
	if o == nil || IsNil(o.InvitationLink) {
		return nil, false
	}
	return o.InvitationLink, true
}

// HasInvitationLink returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfShareSettings) IsInvitationLinkSet() bool {
	if o != nil && !IsNil(o.InvitationLink) {
		return true
	}

	return false
}

// SetInvitationLink gets a reference to the given int32 and assigns it to the InvitationLink field.
func (o *AiFileEntryDtoAllOfShareSettings) SetInvitationLink(v int32) {
	o.InvitationLink = &v
}

// GetPrimaryExternalLink returns the PrimaryExternalLink field value if set, zero value otherwise.
func (o *AiFileEntryDtoAllOfShareSettings) GetPrimaryExternalLink() int32 {
	if o == nil || IsNil(o.PrimaryExternalLink) {
		var ret int32
		return ret
	}
	return *o.PrimaryExternalLink
}

// GetPrimaryExternalLinkOk returns a tuple with the PrimaryExternalLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileEntryDtoAllOfShareSettings) GetPrimaryExternalLinkOk() (*int32, bool) {
	if o == nil || IsNil(o.PrimaryExternalLink) {
		return nil, false
	}
	return o.PrimaryExternalLink, true
}

// HasPrimaryExternalLink returns a boolean if a field has been set.
func (o *AiFileEntryDtoAllOfShareSettings) IsPrimaryExternalLinkSet() bool {
	if o != nil && !IsNil(o.PrimaryExternalLink) {
		return true
	}

	return false
}

// SetPrimaryExternalLink gets a reference to the given int32 and assigns it to the PrimaryExternalLink field.
func (o *AiFileEntryDtoAllOfShareSettings) SetPrimaryExternalLink(v int32) {
	o.PrimaryExternalLink = &v
}

func (o AiFileEntryDtoAllOfShareSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiFileEntryDtoAllOfShareSettings) ToMap() (map[string]interface{}, error) {
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

type NullableAiFileEntryDtoAllOfShareSettings struct {
	value *AiFileEntryDtoAllOfShareSettings
	isSet bool
}

func (v NullableAiFileEntryDtoAllOfShareSettings) Get() *AiFileEntryDtoAllOfShareSettings {
	return v.value
}

func (v *NullableAiFileEntryDtoAllOfShareSettings) Set(val *AiFileEntryDtoAllOfShareSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAiFileEntryDtoAllOfShareSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAiFileEntryDtoAllOfShareSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiFileEntryDtoAllOfShareSettings(val *AiFileEntryDtoAllOfShareSettings) *NullableAiFileEntryDtoAllOfShareSettings {
	return &NullableAiFileEntryDtoAllOfShareSettings{value: val, isSet: true}
}

func (v NullableAiFileEntryDtoAllOfShareSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiFileEntryDtoAllOfShareSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

