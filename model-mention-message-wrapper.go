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

// checks if the MentionMessageWrapper type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MentionMessageWrapper{}

// MentionMessageWrapper The mention notification to send: what to say, whom to tell and where in the document the mention sits.
type MentionMessageWrapper struct {
	// The place in the document the notification link should open at, as the editor reports it when the mention is  made. Left out, the link opens the file at its beginning.
	ActionLink *ActionLinkConfig `json:"actionLink,omitempty"`
	// The addresses to notify. Only an address that belongs to a portal account receives a mail; an unknown address  is skipped, and the answer then carries the access list of the file so that the client can invite its owner.
	Emails []string `json:"emails,omitempty"`
	// The note shown next to the link in the mail. Only its first 200 characters are sent, and a value longer than  the field allows is refused.
	Message NullableString `json:"message,omitempty"`
}

// NewMentionMessageWrapper instantiates a new MentionMessageWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMentionMessageWrapper() *MentionMessageWrapper {
	this := MentionMessageWrapper{}
	return &this
}

// NewMentionMessageWrapperWithDefaults instantiates a new MentionMessageWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMentionMessageWrapperWithDefaults() *MentionMessageWrapper {
	this := MentionMessageWrapper{}
	return &this
}

// GetActionLink returns the ActionLink field value if set, zero value otherwise.
func (o *MentionMessageWrapper) GetActionLink() ActionLinkConfig {
	if o == nil || IsNil(o.ActionLink) {
		var ret ActionLinkConfig
		return ret
	}
	return *o.ActionLink
}

// GetActionLinkOk returns a tuple with the ActionLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MentionMessageWrapper) GetActionLinkOk() (*ActionLinkConfig, bool) {
	if o == nil || IsNil(o.ActionLink) {
		return nil, false
	}
	return o.ActionLink, true
}

// HasActionLink returns a boolean if a field has been set.
func (o *MentionMessageWrapper) IsActionLinkSet() bool {
	if o != nil && !IsNil(o.ActionLink) {
		return true
	}

	return false
}

// SetActionLink gets a reference to the given ActionLinkConfig and assigns it to the ActionLink field.
func (o *MentionMessageWrapper) SetActionLink(v ActionLinkConfig) {
	o.ActionLink = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MentionMessageWrapper) GetEmails() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MentionMessageWrapper) GetEmailsOk() ([]string, bool) {
	if o == nil || IsNil(o.Emails) {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *MentionMessageWrapper) IsEmailsSet() bool {
	if o != nil && !IsNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *MentionMessageWrapper) SetEmails(v []string) {
	o.Emails = v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MentionMessageWrapper) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MentionMessageWrapper) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *MentionMessageWrapper) IsMessageSet() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *MentionMessageWrapper) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *MentionMessageWrapper) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *MentionMessageWrapper) UnsetMessage() {
	o.Message.Unset()
}

func (o MentionMessageWrapper) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MentionMessageWrapper) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionLink) {
		toSerialize["actionLink"] = o.ActionLink
	}
	if o.Emails != nil {
		toSerialize["emails"] = o.Emails
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableMentionMessageWrapper struct {
	value *MentionMessageWrapper
	isSet bool
}

func (v NullableMentionMessageWrapper) Get() *MentionMessageWrapper {
	return v.value
}

func (v *NullableMentionMessageWrapper) Set(val *MentionMessageWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullableMentionMessageWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullableMentionMessageWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMentionMessageWrapper(val *MentionMessageWrapper) *NullableMentionMessageWrapper {
	return &NullableMentionMessageWrapper{value: val, isSet: true}
}

func (v NullableMentionMessageWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMentionMessageWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

