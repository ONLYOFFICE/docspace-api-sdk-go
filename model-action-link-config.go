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

// checks if the ActionLinkConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionLinkConfig{}

// ActionLinkConfig The config parameter which contains the information about the action in the document that will be scrolled to.
type ActionLinkConfig struct {
	Action *ActionConfig `json:"action,omitempty"`
}

// NewActionLinkConfig instantiates a new ActionLinkConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionLinkConfig() *ActionLinkConfig {
	this := ActionLinkConfig{}
	return &this
}

// NewActionLinkConfigWithDefaults instantiates a new ActionLinkConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionLinkConfigWithDefaults() *ActionLinkConfig {
	this := ActionLinkConfig{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ActionLinkConfig) GetAction() ActionConfig {
	if o == nil || IsNil(o.Action) {
		var ret ActionConfig
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionLinkConfig) GetActionOk() (*ActionConfig, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ActionLinkConfig) IsActionSet() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given ActionConfig and assigns it to the Action field.
func (o *ActionLinkConfig) SetAction(v ActionConfig) {
	o.Action = &v
}

func (o ActionLinkConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionLinkConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	return toSerialize, nil
}

type NullableActionLinkConfig struct {
	value *ActionLinkConfig
	isSet bool
}

func (v NullableActionLinkConfig) Get() *ActionLinkConfig {
	return v.value
}

func (v *NullableActionLinkConfig) Set(val *ActionLinkConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableActionLinkConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableActionLinkConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionLinkConfig(val *ActionLinkConfig) *NullableActionLinkConfig {
	return &NullableActionLinkConfig{value: val, isSet: true}
}

func (v NullableActionLinkConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionLinkConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

