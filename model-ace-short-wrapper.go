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

// checks if the AceShortWrapper type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AceShortWrapper{}

// AceShortWrapper One line of a document sharing list in display form: who the document is shared with and the label of their access  level, rather than an access record with identifiers. Entries that deny access and invitation links are left out,  so the list names only the subjects and links that can currently open the document.
type AceShortWrapper struct {
	// Who or what the line stands for, as a display string: the display name of a member, the name of a group, or  the title given to a shared link when `isLink` is true. It is empty when the subject has no name to show - a  shared link that was never given a title, for instance.
	User NullableString `json:"user,omitempty"`
	// The access level of that subject as a localized label, not a code: inside a room it usually names the role the  subject holds there (Viewer, Editor, Room Manager), while outside a room it names the access itself  (Read Only, Full Access). The wording comes from the portal resources and is translated for the current  language, so show it to a person rather than compare it in code.
	Permissions NullableString `json:"permissions,omitempty"`
	// Whether the line stands for a shared link instead of a member or a group. Clients use it to draw a link badge  where they would otherwise draw an avatar.
	IsLink *bool `json:"isLink,omitempty"`
}

// NewAceShortWrapper instantiates a new AceShortWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAceShortWrapper() *AceShortWrapper {
	this := AceShortWrapper{}
	return &this
}

// NewAceShortWrapperWithDefaults instantiates a new AceShortWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAceShortWrapperWithDefaults() *AceShortWrapper {
	this := AceShortWrapper{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AceShortWrapper) GetUser() string {
	if o == nil || IsNil(o.User.Get()) {
		var ret string
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AceShortWrapper) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *AceShortWrapper) IsUserSet() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableString and assigns it to the User field.
func (o *AceShortWrapper) SetUser(v string) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *AceShortWrapper) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *AceShortWrapper) UnsetUser() {
	o.User.Unset()
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AceShortWrapper) GetPermissions() string {
	if o == nil || IsNil(o.Permissions.Get()) {
		var ret string
		return ret
	}
	return *o.Permissions.Get()
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AceShortWrapper) GetPermissionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions.Get(), o.Permissions.IsSet()
}

// HasPermissions returns a boolean if a field has been set.
func (o *AceShortWrapper) IsPermissionsSet() bool {
	if o != nil && o.Permissions.IsSet() {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given NullableString and assigns it to the Permissions field.
func (o *AceShortWrapper) SetPermissions(v string) {
	o.Permissions.Set(&v)
}
// SetPermissionsNil sets the value for Permissions to be an explicit nil
func (o *AceShortWrapper) SetPermissionsNil() {
	o.Permissions.Set(nil)
}

// UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
func (o *AceShortWrapper) UnsetPermissions() {
	o.Permissions.Unset()
}

// GetIsLink returns the IsLink field value if set, zero value otherwise.
func (o *AceShortWrapper) GetIsLink() bool {
	if o == nil || IsNil(o.IsLink) {
		var ret bool
		return ret
	}
	return *o.IsLink
}

// GetIsLinkOk returns a tuple with the IsLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AceShortWrapper) GetIsLinkOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLink) {
		return nil, false
	}
	return o.IsLink, true
}

// HasIsLink returns a boolean if a field has been set.
func (o *AceShortWrapper) IsIsLinkSet() bool {
	if o != nil && !IsNil(o.IsLink) {
		return true
	}

	return false
}

// SetIsLink gets a reference to the given bool and assigns it to the IsLink field.
func (o *AceShortWrapper) SetIsLink(v bool) {
	o.IsLink = &v
}

func (o AceShortWrapper) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AceShortWrapper) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.Permissions.IsSet() {
		toSerialize["permissions"] = o.Permissions.Get()
	}
	if !IsNil(o.IsLink) {
		toSerialize["isLink"] = o.IsLink
	}
	return toSerialize, nil
}

type NullableAceShortWrapper struct {
	value *AceShortWrapper
	isSet bool
}

func (v NullableAceShortWrapper) Get() *AceShortWrapper {
	return v.value
}

func (v *NullableAceShortWrapper) Set(val *AceShortWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullableAceShortWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullableAceShortWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAceShortWrapper(val *AceShortWrapper) *NullableAceShortWrapper {
	return &NullableAceShortWrapper{value: val, isSet: true}
}

func (v NullableAceShortWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAceShortWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

