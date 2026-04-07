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

// checks if the SetUserChatSettingsRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetUserChatSettingsRequestBody{}

// SetUserChatSettingsRequestBody Parameters for updating user chat settings.
type SetUserChatSettingsRequestBody struct {
	// Indicates whether the AI assistant is allowed to perform web searches when generating responses.
	WebSearchEnabled NullableBool `json:"webSearchEnabled,omitempty"`
	ReasoningEffort *ChatReasoningEffort `json:"reasoningEffort,omitempty"`
}

// NewSetUserChatSettingsRequestBody instantiates a new SetUserChatSettingsRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetUserChatSettingsRequestBody() *SetUserChatSettingsRequestBody {
	this := SetUserChatSettingsRequestBody{}
	return &this
}

// NewSetUserChatSettingsRequestBodyWithDefaults instantiates a new SetUserChatSettingsRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetUserChatSettingsRequestBodyWithDefaults() *SetUserChatSettingsRequestBody {
	this := SetUserChatSettingsRequestBody{}
	return &this
}

// GetWebSearchEnabled returns the WebSearchEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetUserChatSettingsRequestBody) GetWebSearchEnabled() bool {
	if o == nil || IsNil(o.WebSearchEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.WebSearchEnabled.Get()
}

// GetWebSearchEnabledOk returns a tuple with the WebSearchEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetUserChatSettingsRequestBody) GetWebSearchEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebSearchEnabled.Get(), o.WebSearchEnabled.IsSet()
}

// HasWebSearchEnabled returns a boolean if a field has been set.
func (o *SetUserChatSettingsRequestBody) IsWebSearchEnabledSet() bool {
	if o != nil && o.WebSearchEnabled.IsSet() {
		return true
	}

	return false
}

// SetWebSearchEnabled gets a reference to the given NullableBool and assigns it to the WebSearchEnabled field.
func (o *SetUserChatSettingsRequestBody) SetWebSearchEnabled(v bool) {
	o.WebSearchEnabled.Set(&v)
}
// SetWebSearchEnabledNil sets the value for WebSearchEnabled to be an explicit nil
func (o *SetUserChatSettingsRequestBody) SetWebSearchEnabledNil() {
	o.WebSearchEnabled.Set(nil)
}

// UnsetWebSearchEnabled ensures that no value is present for WebSearchEnabled, not even an explicit nil
func (o *SetUserChatSettingsRequestBody) UnsetWebSearchEnabled() {
	o.WebSearchEnabled.Unset()
}

// GetReasoningEffort returns the ReasoningEffort field value if set, zero value otherwise.
func (o *SetUserChatSettingsRequestBody) GetReasoningEffort() ChatReasoningEffort {
	if o == nil || IsNil(o.ReasoningEffort) {
		var ret ChatReasoningEffort
		return ret
	}
	return *o.ReasoningEffort
}

// GetReasoningEffortOk returns a tuple with the ReasoningEffort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetUserChatSettingsRequestBody) GetReasoningEffortOk() (*ChatReasoningEffort, bool) {
	if o == nil || IsNil(o.ReasoningEffort) {
		return nil, false
	}
	return o.ReasoningEffort, true
}

// HasReasoningEffort returns a boolean if a field has been set.
func (o *SetUserChatSettingsRequestBody) IsReasoningEffortSet() bool {
	if o != nil && !IsNil(o.ReasoningEffort) {
		return true
	}

	return false
}

// SetReasoningEffort gets a reference to the given ChatReasoningEffort and assigns it to the ReasoningEffort field.
func (o *SetUserChatSettingsRequestBody) SetReasoningEffort(v ChatReasoningEffort) {
	o.ReasoningEffort = &v
}

func (o SetUserChatSettingsRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetUserChatSettingsRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.WebSearchEnabled.IsSet() {
		toSerialize["webSearchEnabled"] = o.WebSearchEnabled.Get()
	}
	if !IsNil(o.ReasoningEffort) {
		toSerialize["reasoningEffort"] = o.ReasoningEffort
	}
	return toSerialize, nil
}

type NullableSetUserChatSettingsRequestBody struct {
	value *SetUserChatSettingsRequestBody
	isSet bool
}

func (v NullableSetUserChatSettingsRequestBody) Get() *SetUserChatSettingsRequestBody {
	return v.value
}

func (v *NullableSetUserChatSettingsRequestBody) Set(val *SetUserChatSettingsRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSetUserChatSettingsRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSetUserChatSettingsRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetUserChatSettingsRequestBody(val *SetUserChatSettingsRequestBody) *NullableSetUserChatSettingsRequestBody {
	return &NullableSetUserChatSettingsRequestBody{value: val, isSet: true}
}

func (v NullableSetUserChatSettingsRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetUserChatSettingsRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

