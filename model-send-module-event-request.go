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

// checks if the SendModuleEventRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendModuleEventRequest{}

// SendModuleEventRequest A socket event from an integration module: who receives it in the portal, its name and its payload.
type SendModuleEventRequest struct {
	// A room of the module in this portal: everyone subscribed to `plugin-{module}-{room}` receives the event. Latin  letters, digits and the characters `_ . : -`.
	Room NullableString `json:"room,omitempty"`
	// The e-mail of the portal user who receives the event; with single sign-on it is the same e-mail the module's own  application knows the person by. Only that user's open sessions receive it.
	User NullableString `json:"user,omitempty"`
	// The event name. Browsers receive it as `plugin:{module}:{event}`. Latin letters, digits and the characters  `_ . : -`.
	Event NullableString `json:"event,omitempty"`
	// Passed to the browser as it is, as `{ message }`. Meant as a signal to re-read, such as an unread count; use  JSON text for structured data and keep personal content out of it.
	Message NullableString `json:"message,omitempty"`
}

// NewSendModuleEventRequest instantiates a new SendModuleEventRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendModuleEventRequest() *SendModuleEventRequest {
	this := SendModuleEventRequest{}
	return &this
}

// NewSendModuleEventRequestWithDefaults instantiates a new SendModuleEventRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendModuleEventRequestWithDefaults() *SendModuleEventRequest {
	this := SendModuleEventRequest{}
	return &this
}

// GetRoom returns the Room field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendModuleEventRequest) GetRoom() string {
	if o == nil || IsNil(o.Room.Get()) {
		var ret string
		return ret
	}
	return *o.Room.Get()
}

// GetRoomOk returns a tuple with the Room field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendModuleEventRequest) GetRoomOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Room.Get(), o.Room.IsSet()
}

// HasRoom returns a boolean if a field has been set.
func (o *SendModuleEventRequest) IsRoomSet() bool {
	if o != nil && o.Room.IsSet() {
		return true
	}

	return false
}

// SetRoom gets a reference to the given NullableString and assigns it to the Room field.
func (o *SendModuleEventRequest) SetRoom(v string) {
	o.Room.Set(&v)
}
// SetRoomNil sets the value for Room to be an explicit nil
func (o *SendModuleEventRequest) SetRoomNil() {
	o.Room.Set(nil)
}

// UnsetRoom ensures that no value is present for Room, not even an explicit nil
func (o *SendModuleEventRequest) UnsetRoom() {
	o.Room.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendModuleEventRequest) GetUser() string {
	if o == nil || IsNil(o.User.Get()) {
		var ret string
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendModuleEventRequest) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *SendModuleEventRequest) IsUserSet() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableString and assigns it to the User field.
func (o *SendModuleEventRequest) SetUser(v string) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *SendModuleEventRequest) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *SendModuleEventRequest) UnsetUser() {
	o.User.Unset()
}

// GetEvent returns the Event field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendModuleEventRequest) GetEvent() string {
	if o == nil || IsNil(o.Event.Get()) {
		var ret string
		return ret
	}
	return *o.Event.Get()
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendModuleEventRequest) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Event.Get(), o.Event.IsSet()
}

// HasEvent returns a boolean if a field has been set.
func (o *SendModuleEventRequest) IsEventSet() bool {
	if o != nil && o.Event.IsSet() {
		return true
	}

	return false
}

// SetEvent gets a reference to the given NullableString and assigns it to the Event field.
func (o *SendModuleEventRequest) SetEvent(v string) {
	o.Event.Set(&v)
}
// SetEventNil sets the value for Event to be an explicit nil
func (o *SendModuleEventRequest) SetEventNil() {
	o.Event.Set(nil)
}

// UnsetEvent ensures that no value is present for Event, not even an explicit nil
func (o *SendModuleEventRequest) UnsetEvent() {
	o.Event.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendModuleEventRequest) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendModuleEventRequest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *SendModuleEventRequest) IsMessageSet() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *SendModuleEventRequest) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *SendModuleEventRequest) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *SendModuleEventRequest) UnsetMessage() {
	o.Message.Unset()
}

func (o SendModuleEventRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendModuleEventRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Room.IsSet() {
		toSerialize["room"] = o.Room.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.Event.IsSet() {
		toSerialize["event"] = o.Event.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableSendModuleEventRequest struct {
	value *SendModuleEventRequest
	isSet bool
}

func (v NullableSendModuleEventRequest) Get() *SendModuleEventRequest {
	return v.value
}

func (v *NullableSendModuleEventRequest) Set(val *SendModuleEventRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendModuleEventRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendModuleEventRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendModuleEventRequest(val *SendModuleEventRequest) *NullableSendModuleEventRequest {
	return &NullableSendModuleEventRequest{value: val, isSet: true}
}

func (v NullableSendModuleEventRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendModuleEventRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

