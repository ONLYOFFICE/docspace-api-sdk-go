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

// checks if the MentionMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MentionMessageRequest{}

// MentionMessageRequest The mention notification to send: what to say, whom to tell and where in the document the mention sits.
type MentionMessageRequest struct {
	// The place in the document the notification link should open at, as the editor reports it when the mention is  made. Left out, the link opens the file at its beginning.
	ActionLink *ActionLinkRequest `json:"actionLink,omitempty"`
	// The addresses to notify. Only an address that belongs to a portal account receives a mail; an unknown address  is skipped, and the answer then carries the access list of the file so that the client can invite its owner.
	Emails []string `json:"emails,omitempty"`
	// The note shown next to the link in the mail. Only its first 200 characters are sent, and a value longer than  the field allows is refused.
	Message NullableString `json:"message,omitempty"`
}

// NewMentionMessageRequest instantiates a new MentionMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMentionMessageRequest() *MentionMessageRequest {
	this := MentionMessageRequest{}
	return &this
}

// NewMentionMessageRequestWithDefaults instantiates a new MentionMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMentionMessageRequestWithDefaults() *MentionMessageRequest {
	this := MentionMessageRequest{}
	return &this
}

// GetActionLink returns the ActionLink field value if set, zero value otherwise.
func (o *MentionMessageRequest) GetActionLink() ActionLinkRequest {
	if o == nil || IsNil(o.ActionLink) {
		var ret ActionLinkRequest
		return ret
	}
	return *o.ActionLink
}

// GetActionLinkOk returns a tuple with the ActionLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MentionMessageRequest) GetActionLinkOk() (*ActionLinkRequest, bool) {
	if o == nil || IsNil(o.ActionLink) {
		return nil, false
	}
	return o.ActionLink, true
}

// HasActionLink returns a boolean if a field has been set.
func (o *MentionMessageRequest) IsActionLinkSet() bool {
	if o != nil && !IsNil(o.ActionLink) {
		return true
	}

	return false
}

// SetActionLink gets a reference to the given ActionLinkRequest and assigns it to the ActionLink field.
func (o *MentionMessageRequest) SetActionLink(v ActionLinkRequest) {
	o.ActionLink = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MentionMessageRequest) GetEmails() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MentionMessageRequest) GetEmailsOk() ([]string, bool) {
	if o == nil || IsNil(o.Emails) {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *MentionMessageRequest) IsEmailsSet() bool {
	if o != nil && !IsNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *MentionMessageRequest) SetEmails(v []string) {
	o.Emails = v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MentionMessageRequest) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MentionMessageRequest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *MentionMessageRequest) IsMessageSet() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *MentionMessageRequest) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *MentionMessageRequest) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *MentionMessageRequest) UnsetMessage() {
	o.Message.Unset()
}

func (o MentionMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MentionMessageRequest) ToMap() (map[string]interface{}, error) {
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

type NullableMentionMessageRequest struct {
	value *MentionMessageRequest
	isSet bool
}

func (v NullableMentionMessageRequest) Get() *MentionMessageRequest {
	return v.value
}

func (v *NullableMentionMessageRequest) Set(val *MentionMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMentionMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMentionMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMentionMessageRequest(val *MentionMessageRequest) *NullableMentionMessageRequest {
	return &NullableMentionMessageRequest{value: val, isSet: true}
}

func (v NullableMentionMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMentionMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

