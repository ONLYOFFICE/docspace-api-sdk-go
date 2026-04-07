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

// checks if the SetWebSearchSettingsRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetWebSearchSettingsRequestBody{}

// SetWebSearchSettingsRequestBody Parameters for configuring web search settings.
type SetWebSearchSettingsRequestBody struct {
	// Indicates whether web search is enabled for AI chat sessions.
	Enabled *bool `json:"enabled,omitempty"`
	Type *EngineType `json:"type,omitempty"`
	// The API key for the selected web search engine. Pass null to keep the existing key unchanged.
	Key NullableString `json:"key,omitempty"`
}

// NewSetWebSearchSettingsRequestBody instantiates a new SetWebSearchSettingsRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetWebSearchSettingsRequestBody() *SetWebSearchSettingsRequestBody {
	this := SetWebSearchSettingsRequestBody{}
	return &this
}

// NewSetWebSearchSettingsRequestBodyWithDefaults instantiates a new SetWebSearchSettingsRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetWebSearchSettingsRequestBodyWithDefaults() *SetWebSearchSettingsRequestBody {
	this := SetWebSearchSettingsRequestBody{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SetWebSearchSettingsRequestBody) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetWebSearchSettingsRequestBody) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SetWebSearchSettingsRequestBody) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SetWebSearchSettingsRequestBody) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SetWebSearchSettingsRequestBody) GetType() EngineType {
	if o == nil || IsNil(o.Type) {
		var ret EngineType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetWebSearchSettingsRequestBody) GetTypeOk() (*EngineType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SetWebSearchSettingsRequestBody) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EngineType and assigns it to the Type field.
func (o *SetWebSearchSettingsRequestBody) SetType(v EngineType) {
	o.Type = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetWebSearchSettingsRequestBody) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetWebSearchSettingsRequestBody) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *SetWebSearchSettingsRequestBody) IsKeySet() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *SetWebSearchSettingsRequestBody) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *SetWebSearchSettingsRequestBody) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *SetWebSearchSettingsRequestBody) UnsetKey() {
	o.Key.Unset()
}

func (o SetWebSearchSettingsRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetWebSearchSettingsRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	return toSerialize, nil
}

type NullableSetWebSearchSettingsRequestBody struct {
	value *SetWebSearchSettingsRequestBody
	isSet bool
}

func (v NullableSetWebSearchSettingsRequestBody) Get() *SetWebSearchSettingsRequestBody {
	return v.value
}

func (v *NullableSetWebSearchSettingsRequestBody) Set(val *SetWebSearchSettingsRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSetWebSearchSettingsRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSetWebSearchSettingsRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetWebSearchSettingsRequestBody(val *SetWebSearchSettingsRequestBody) *NullableSetWebSearchSettingsRequestBody {
	return &NullableSetWebSearchSettingsRequestBody{value: val, isSet: true}
}

func (v NullableSetWebSearchSettingsRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetWebSearchSettingsRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

