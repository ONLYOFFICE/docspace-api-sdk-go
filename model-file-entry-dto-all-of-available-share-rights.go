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

// checks if the FileEntryDtoAllOfAvailableShareRights type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileEntryDtoAllOfAvailableShareRights{}

// FileEntryDtoAllOfAvailableShareRights Which access levels may be handed out on this entry, listed per kind of recipient, so that a client offers  only levels the entry actually supports - a room for filling forms and a plain folder do not accept the same  ones.
type FileEntryDtoAllOfAvailableShareRights struct {
	User []string `json:"User,omitempty"`
	ExternalLink []string `json:"ExternalLink,omitempty"`
	Group []string `json:"Group,omitempty"`
	InvitationLink []string `json:"InvitationLink,omitempty"`
	PrimaryExternalLink []string `json:"PrimaryExternalLink,omitempty"`
	AiProject []string `json:"AiProject,omitempty"`
}

// NewFileEntryDtoAllOfAvailableShareRights instantiates a new FileEntryDtoAllOfAvailableShareRights object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileEntryDtoAllOfAvailableShareRights() *FileEntryDtoAllOfAvailableShareRights {
	this := FileEntryDtoAllOfAvailableShareRights{}
	return &this
}

// NewFileEntryDtoAllOfAvailableShareRightsWithDefaults instantiates a new FileEntryDtoAllOfAvailableShareRights object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileEntryDtoAllOfAvailableShareRightsWithDefaults() *FileEntryDtoAllOfAvailableShareRights {
	this := FileEntryDtoAllOfAvailableShareRights{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *FileEntryDtoAllOfAvailableShareRights) GetUser() []string {
	if o == nil || IsNil(o.User) {
		var ret []string
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntryDtoAllOfAvailableShareRights) GetUserOk() ([]string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *FileEntryDtoAllOfAvailableShareRights) IsUserSet() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given []string and assigns it to the User field.
func (o *FileEntryDtoAllOfAvailableShareRights) SetUser(v []string) {
	o.User = v
}

// GetExternalLink returns the ExternalLink field value if set, zero value otherwise.
func (o *FileEntryDtoAllOfAvailableShareRights) GetExternalLink() []string {
	if o == nil || IsNil(o.ExternalLink) {
		var ret []string
		return ret
	}
	return o.ExternalLink
}

// GetExternalLinkOk returns a tuple with the ExternalLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntryDtoAllOfAvailableShareRights) GetExternalLinkOk() ([]string, bool) {
	if o == nil || IsNil(o.ExternalLink) {
		return nil, false
	}
	return o.ExternalLink, true
}

// HasExternalLink returns a boolean if a field has been set.
func (o *FileEntryDtoAllOfAvailableShareRights) IsExternalLinkSet() bool {
	if o != nil && !IsNil(o.ExternalLink) {
		return true
	}

	return false
}

// SetExternalLink gets a reference to the given []string and assigns it to the ExternalLink field.
func (o *FileEntryDtoAllOfAvailableShareRights) SetExternalLink(v []string) {
	o.ExternalLink = v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *FileEntryDtoAllOfAvailableShareRights) GetGroup() []string {
	if o == nil || IsNil(o.Group) {
		var ret []string
		return ret
	}
	return o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntryDtoAllOfAvailableShareRights) GetGroupOk() ([]string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *FileEntryDtoAllOfAvailableShareRights) IsGroupSet() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given []string and assigns it to the Group field.
func (o *FileEntryDtoAllOfAvailableShareRights) SetGroup(v []string) {
	o.Group = v
}

// GetInvitationLink returns the InvitationLink field value if set, zero value otherwise.
func (o *FileEntryDtoAllOfAvailableShareRights) GetInvitationLink() []string {
	if o == nil || IsNil(o.InvitationLink) {
		var ret []string
		return ret
	}
	return o.InvitationLink
}

// GetInvitationLinkOk returns a tuple with the InvitationLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntryDtoAllOfAvailableShareRights) GetInvitationLinkOk() ([]string, bool) {
	if o == nil || IsNil(o.InvitationLink) {
		return nil, false
	}
	return o.InvitationLink, true
}

// HasInvitationLink returns a boolean if a field has been set.
func (o *FileEntryDtoAllOfAvailableShareRights) IsInvitationLinkSet() bool {
	if o != nil && !IsNil(o.InvitationLink) {
		return true
	}

	return false
}

// SetInvitationLink gets a reference to the given []string and assigns it to the InvitationLink field.
func (o *FileEntryDtoAllOfAvailableShareRights) SetInvitationLink(v []string) {
	o.InvitationLink = v
}

// GetPrimaryExternalLink returns the PrimaryExternalLink field value if set, zero value otherwise.
func (o *FileEntryDtoAllOfAvailableShareRights) GetPrimaryExternalLink() []string {
	if o == nil || IsNil(o.PrimaryExternalLink) {
		var ret []string
		return ret
	}
	return o.PrimaryExternalLink
}

// GetPrimaryExternalLinkOk returns a tuple with the PrimaryExternalLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntryDtoAllOfAvailableShareRights) GetPrimaryExternalLinkOk() ([]string, bool) {
	if o == nil || IsNil(o.PrimaryExternalLink) {
		return nil, false
	}
	return o.PrimaryExternalLink, true
}

// HasPrimaryExternalLink returns a boolean if a field has been set.
func (o *FileEntryDtoAllOfAvailableShareRights) IsPrimaryExternalLinkSet() bool {
	if o != nil && !IsNil(o.PrimaryExternalLink) {
		return true
	}

	return false
}

// SetPrimaryExternalLink gets a reference to the given []string and assigns it to the PrimaryExternalLink field.
func (o *FileEntryDtoAllOfAvailableShareRights) SetPrimaryExternalLink(v []string) {
	o.PrimaryExternalLink = v
}

// GetAiProject returns the AiProject field value if set, zero value otherwise.
func (o *FileEntryDtoAllOfAvailableShareRights) GetAiProject() []string {
	if o == nil || IsNil(o.AiProject) {
		var ret []string
		return ret
	}
	return o.AiProject
}

// GetAiProjectOk returns a tuple with the AiProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileEntryDtoAllOfAvailableShareRights) GetAiProjectOk() ([]string, bool) {
	if o == nil || IsNil(o.AiProject) {
		return nil, false
	}
	return o.AiProject, true
}

// HasAiProject returns a boolean if a field has been set.
func (o *FileEntryDtoAllOfAvailableShareRights) IsAiProjectSet() bool {
	if o != nil && !IsNil(o.AiProject) {
		return true
	}

	return false
}

// SetAiProject gets a reference to the given []string and assigns it to the AiProject field.
func (o *FileEntryDtoAllOfAvailableShareRights) SetAiProject(v []string) {
	o.AiProject = v
}

func (o FileEntryDtoAllOfAvailableShareRights) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileEntryDtoAllOfAvailableShareRights) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AiProject) {
		toSerialize["AiProject"] = o.AiProject
	}
	return toSerialize, nil
}

type NullableFileEntryDtoAllOfAvailableShareRights struct {
	value *FileEntryDtoAllOfAvailableShareRights
	isSet bool
}

func (v NullableFileEntryDtoAllOfAvailableShareRights) Get() *FileEntryDtoAllOfAvailableShareRights {
	return v.value
}

func (v *NullableFileEntryDtoAllOfAvailableShareRights) Set(val *FileEntryDtoAllOfAvailableShareRights) {
	v.value = val
	v.isSet = true
}

func (v NullableFileEntryDtoAllOfAvailableShareRights) IsSet() bool {
	return v.isSet
}

func (v *NullableFileEntryDtoAllOfAvailableShareRights) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileEntryDtoAllOfAvailableShareRights(val *FileEntryDtoAllOfAvailableShareRights) *NullableFileEntryDtoAllOfAvailableShareRights {
	return &NullableFileEntryDtoAllOfAvailableShareRights{value: val, isSet: true}
}

func (v NullableFileEntryDtoAllOfAvailableShareRights) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileEntryDtoAllOfAvailableShareRights) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

