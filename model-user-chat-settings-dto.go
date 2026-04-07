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

// checks if the UserChatSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserChatSettingsDto{}

// UserChatSettingsDto The user chat settings.
type UserChatSettingsDto struct {
	// Indicates whether the AI assistant is allowed to perform web searches when generating responses in this room.
	WebSearchEnabled *bool `json:"webSearchEnabled,omitempty"`
	ReasoningEffort *ChatReasoningEffort `json:"reasoningEffort,omitempty"`
}

// NewUserChatSettingsDto instantiates a new UserChatSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserChatSettingsDto() *UserChatSettingsDto {
	this := UserChatSettingsDto{}
	return &this
}

// NewUserChatSettingsDtoWithDefaults instantiates a new UserChatSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserChatSettingsDtoWithDefaults() *UserChatSettingsDto {
	this := UserChatSettingsDto{}
	return &this
}

// GetWebSearchEnabled returns the WebSearchEnabled field value if set, zero value otherwise.
func (o *UserChatSettingsDto) GetWebSearchEnabled() bool {
	if o == nil || IsNil(o.WebSearchEnabled) {
		var ret bool
		return ret
	}
	return *o.WebSearchEnabled
}

// GetWebSearchEnabledOk returns a tuple with the WebSearchEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserChatSettingsDto) GetWebSearchEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.WebSearchEnabled) {
		return nil, false
	}
	return o.WebSearchEnabled, true
}

// HasWebSearchEnabled returns a boolean if a field has been set.
func (o *UserChatSettingsDto) IsWebSearchEnabledSet() bool {
	if o != nil && !IsNil(o.WebSearchEnabled) {
		return true
	}

	return false
}

// SetWebSearchEnabled gets a reference to the given bool and assigns it to the WebSearchEnabled field.
func (o *UserChatSettingsDto) SetWebSearchEnabled(v bool) {
	o.WebSearchEnabled = &v
}

// GetReasoningEffort returns the ReasoningEffort field value if set, zero value otherwise.
func (o *UserChatSettingsDto) GetReasoningEffort() ChatReasoningEffort {
	if o == nil || IsNil(o.ReasoningEffort) {
		var ret ChatReasoningEffort
		return ret
	}
	return *o.ReasoningEffort
}

// GetReasoningEffortOk returns a tuple with the ReasoningEffort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserChatSettingsDto) GetReasoningEffortOk() (*ChatReasoningEffort, bool) {
	if o == nil || IsNil(o.ReasoningEffort) {
		return nil, false
	}
	return o.ReasoningEffort, true
}

// HasReasoningEffort returns a boolean if a field has been set.
func (o *UserChatSettingsDto) IsReasoningEffortSet() bool {
	if o != nil && !IsNil(o.ReasoningEffort) {
		return true
	}

	return false
}

// SetReasoningEffort gets a reference to the given ChatReasoningEffort and assigns it to the ReasoningEffort field.
func (o *UserChatSettingsDto) SetReasoningEffort(v ChatReasoningEffort) {
	o.ReasoningEffort = &v
}

func (o UserChatSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserChatSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WebSearchEnabled) {
		toSerialize["webSearchEnabled"] = o.WebSearchEnabled
	}
	if !IsNil(o.ReasoningEffort) {
		toSerialize["reasoningEffort"] = o.ReasoningEffort
	}
	return toSerialize, nil
}

type NullableUserChatSettingsDto struct {
	value *UserChatSettingsDto
	isSet bool
}

func (v NullableUserChatSettingsDto) Get() *UserChatSettingsDto {
	return v.value
}

func (v *NullableUserChatSettingsDto) Set(val *UserChatSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUserChatSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUserChatSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserChatSettingsDto(val *UserChatSettingsDto) *NullableUserChatSettingsDto {
	return &NullableUserChatSettingsDto{value: val, isSet: true}
}

func (v NullableUserChatSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserChatSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

