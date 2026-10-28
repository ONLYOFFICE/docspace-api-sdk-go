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

// checks if the RoomInvitationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomInvitationRequest{}

// RoomInvitationRequest One batch of membership changes for a room.
type RoomInvitationRequest struct {
	// Who is added, changed or removed, one entry per subject. The same subject named twice keeps the level of the  last entry, and an empty list is accepted and changes nothing.
	Invitations []RoomInvitation `json:"invitations,omitempty"`
	// Whether the subjects that gained access are told about it by email. With it off the change is silent, which is  the usual choice when membership is synchronised from another system.
	Notify *bool `json:"notify,omitempty"`
	// The line added to the invitation email. It is used only while the notification is on, and it reaches nobody  whose access was removed.
	Message NullableString `json:"message,omitempty"`
	// The language of the invitation email, as a portal culture name such as en-US. Leaving it out sends each  message in the language of its recipient.
	Culture NullableString `json:"culture,omitempty"`
	// Whether a member who still holds a role in an unfinished form is removed anyway. With it off such a removal is  refused and reported through the error of the answer, so the form can be reassigned first.
	Force *bool `json:"force,omitempty"`
}

// NewRoomInvitationRequest instantiates a new RoomInvitationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomInvitationRequest() *RoomInvitationRequest {
	this := RoomInvitationRequest{}
	return &this
}

// NewRoomInvitationRequestWithDefaults instantiates a new RoomInvitationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomInvitationRequestWithDefaults() *RoomInvitationRequest {
	this := RoomInvitationRequest{}
	return &this
}

// GetInvitations returns the Invitations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomInvitationRequest) GetInvitations() []RoomInvitation {
	if o == nil {
		var ret []RoomInvitation
		return ret
	}
	return o.Invitations
}

// GetInvitationsOk returns a tuple with the Invitations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomInvitationRequest) GetInvitationsOk() ([]RoomInvitation, bool) {
	if o == nil || IsNil(o.Invitations) {
		return nil, false
	}
	return o.Invitations, true
}

// HasInvitations returns a boolean if a field has been set.
func (o *RoomInvitationRequest) IsInvitationsSet() bool {
	if o != nil && !IsNil(o.Invitations) {
		return true
	}

	return false
}

// SetInvitations gets a reference to the given []RoomInvitation and assigns it to the Invitations field.
func (o *RoomInvitationRequest) SetInvitations(v []RoomInvitation) {
	o.Invitations = v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *RoomInvitationRequest) GetNotify() bool {
	if o == nil || IsNil(o.Notify) {
		var ret bool
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomInvitationRequest) GetNotifyOk() (*bool, bool) {
	if o == nil || IsNil(o.Notify) {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *RoomInvitationRequest) IsNotifySet() bool {
	if o != nil && !IsNil(o.Notify) {
		return true
	}

	return false
}

// SetNotify gets a reference to the given bool and assigns it to the Notify field.
func (o *RoomInvitationRequest) SetNotify(v bool) {
	o.Notify = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomInvitationRequest) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomInvitationRequest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *RoomInvitationRequest) IsMessageSet() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *RoomInvitationRequest) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *RoomInvitationRequest) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *RoomInvitationRequest) UnsetMessage() {
	o.Message.Unset()
}

// GetCulture returns the Culture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomInvitationRequest) GetCulture() string {
	if o == nil || IsNil(o.Culture.Get()) {
		var ret string
		return ret
	}
	return *o.Culture.Get()
}

// GetCultureOk returns a tuple with the Culture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomInvitationRequest) GetCultureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Culture.Get(), o.Culture.IsSet()
}

// HasCulture returns a boolean if a field has been set.
func (o *RoomInvitationRequest) IsCultureSet() bool {
	if o != nil && o.Culture.IsSet() {
		return true
	}

	return false
}

// SetCulture gets a reference to the given NullableString and assigns it to the Culture field.
func (o *RoomInvitationRequest) SetCulture(v string) {
	o.Culture.Set(&v)
}
// SetCultureNil sets the value for Culture to be an explicit nil
func (o *RoomInvitationRequest) SetCultureNil() {
	o.Culture.Set(nil)
}

// UnsetCulture ensures that no value is present for Culture, not even an explicit nil
func (o *RoomInvitationRequest) UnsetCulture() {
	o.Culture.Unset()
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *RoomInvitationRequest) GetForce() bool {
	if o == nil || IsNil(o.Force) {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomInvitationRequest) GetForceOk() (*bool, bool) {
	if o == nil || IsNil(o.Force) {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *RoomInvitationRequest) IsForceSet() bool {
	if o != nil && !IsNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *RoomInvitationRequest) SetForce(v bool) {
	o.Force = &v
}

func (o RoomInvitationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomInvitationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Invitations != nil {
		toSerialize["invitations"] = o.Invitations
	}
	if !IsNil(o.Notify) {
		toSerialize["notify"] = o.Notify
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Culture.IsSet() {
		toSerialize["culture"] = o.Culture.Get()
	}
	if !IsNil(o.Force) {
		toSerialize["force"] = o.Force
	}
	return toSerialize, nil
}

type NullableRoomInvitationRequest struct {
	value *RoomInvitationRequest
	isSet bool
}

func (v NullableRoomInvitationRequest) Get() *RoomInvitationRequest {
	return v.value
}

func (v *NullableRoomInvitationRequest) Set(val *RoomInvitationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomInvitationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomInvitationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomInvitationRequest(val *RoomInvitationRequest) *NullableRoomInvitationRequest {
	return &NullableRoomInvitationRequest{value: val, isSet: true}
}

func (v NullableRoomInvitationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomInvitationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

