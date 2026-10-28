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
	"bytes"
	"fmt"
)

// checks if the WebPluginRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebPluginRequest{}

// WebPluginRequest The state the portal keeps for an installed web plugin: whether it runs, and its own settings blob.
type WebPluginRequest struct {
	// Whether the plugin runs in this portal. Switching it on adds the domains its manifest declares to the portal  Content Security Policy and switching it off takes them away again; connected clients are told of the new  state without a reload.
	Enabled *bool `json:"enabled,omitempty"`
	// The configuration the plugin reads at run time, as a JSON document serialised into a string. Its shape is  defined by the plugin and not by the portal, which stores it encrypted for this portal alone. It replaces  whatever was stored rather than merging into it, so send `{}` when there is nothing to keep.
	Settings NullableString `json:"settings"`
}

type _WebPluginRequest WebPluginRequest

// NewWebPluginRequest instantiates a new WebPluginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebPluginRequest(settings NullableString) *WebPluginRequest {
	this := WebPluginRequest{}
	this.Settings = settings
	return &this
}

// NewWebPluginRequestWithDefaults instantiates a new WebPluginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebPluginRequestWithDefaults() *WebPluginRequest {
	this := WebPluginRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WebPluginRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebPluginRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WebPluginRequest) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WebPluginRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSettings returns the Settings field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebPluginRequest) GetSettings() string {
	if o == nil || o.Settings.Get() == nil {
		var ret string
		return ret
	}

	return *o.Settings.Get()
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebPluginRequest) GetSettingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Settings.Get(), o.Settings.IsSet()
}

// SetSettings sets field value
func (o *WebPluginRequest) SetSettings(v string) {
	o.Settings.Set(&v)
}

func (o WebPluginRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebPluginRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["settings"] = o.Settings.Get()
	return toSerialize, nil
}

func (o *WebPluginRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"settings",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWebPluginRequest := _WebPluginRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varWebPluginRequest)

	if err != nil {
		return err
	}

	*o = WebPluginRequest(varWebPluginRequest)

	return err
}

type NullableWebPluginRequest struct {
	value *WebPluginRequest
	isSet bool
}

func (v NullableWebPluginRequest) Get() *WebPluginRequest {
	return v.value
}

func (v *NullableWebPluginRequest) Set(val *WebPluginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWebPluginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWebPluginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebPluginRequest(val *WebPluginRequest) *NullableWebPluginRequest {
	return &NullableWebPluginRequest{value: val, isSet: true}
}

func (v NullableWebPluginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebPluginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

