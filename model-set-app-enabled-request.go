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

// checks if the SetAppEnabledRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetAppEnabledRequest{}

// SetAppEnabledRequest Whether a portal application is switched on.
type SetAppEnabledRequest struct {
	// Whether the application is available in this portal. Switching it off leaves its settings document stored, so  switching it back on restores the configuration it had; connected clients are told of the new state without a  reload. The field is mandatory: a body without it is rejected with 400 and leaves the application as it was.
	Enabled bool `json:"enabled"`
}

type _SetAppEnabledRequest SetAppEnabledRequest

// NewSetAppEnabledRequest instantiates a new SetAppEnabledRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetAppEnabledRequest(enabled bool) *SetAppEnabledRequest {
	this := SetAppEnabledRequest{}
	this.Enabled = enabled
	return &this
}

// NewSetAppEnabledRequestWithDefaults instantiates a new SetAppEnabledRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetAppEnabledRequestWithDefaults() *SetAppEnabledRequest {
	this := SetAppEnabledRequest{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *SetAppEnabledRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SetAppEnabledRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SetAppEnabledRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o SetAppEnabledRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetAppEnabledRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

func (o *SetAppEnabledRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
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

	varSetAppEnabledRequest := _SetAppEnabledRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSetAppEnabledRequest)

	if err != nil {
		return err
	}

	*o = SetAppEnabledRequest(varSetAppEnabledRequest)

	return err
}

type NullableSetAppEnabledRequest struct {
	value *SetAppEnabledRequest
	isSet bool
}

func (v NullableSetAppEnabledRequest) Get() *SetAppEnabledRequest {
	return v.value
}

func (v *NullableSetAppEnabledRequest) Set(val *SetAppEnabledRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetAppEnabledRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetAppEnabledRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetAppEnabledRequest(val *SetAppEnabledRequest) *NullableSetAppEnabledRequest {
	return &NullableSetAppEnabledRequest{value: val, isSet: true}
}

func (v NullableSetAppEnabledRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetAppEnabledRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

