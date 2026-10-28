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

// checks if the AiPreferencesSetDeepModeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiPreferencesSetDeepModeRequest{}

// AiPreferencesSetDeepModeRequest struct for AiPreferencesSetDeepModeRequest
type AiPreferencesSetDeepModeRequest struct {
	// New deep-mode value.
	Value bool `json:"value"`
	EntityId *string `json:"entityId,omitempty"`
}

type _AiPreferencesSetDeepModeRequest AiPreferencesSetDeepModeRequest

// NewAiPreferencesSetDeepModeRequest instantiates a new AiPreferencesSetDeepModeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiPreferencesSetDeepModeRequest(value bool) *AiPreferencesSetDeepModeRequest {
	this := AiPreferencesSetDeepModeRequest{}
	this.Value = value
	return &this
}

// NewAiPreferencesSetDeepModeRequestWithDefaults instantiates a new AiPreferencesSetDeepModeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiPreferencesSetDeepModeRequestWithDefaults() *AiPreferencesSetDeepModeRequest {
	this := AiPreferencesSetDeepModeRequest{}
	return &this
}

// GetValue returns the Value field value
func (o *AiPreferencesSetDeepModeRequest) GetValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AiPreferencesSetDeepModeRequest) GetValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AiPreferencesSetDeepModeRequest) SetValue(v bool) {
	o.Value = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiPreferencesSetDeepModeRequest) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiPreferencesSetDeepModeRequest) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiPreferencesSetDeepModeRequest) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiPreferencesSetDeepModeRequest) SetEntityId(v string) {
	o.EntityId = &v
}

func (o AiPreferencesSetDeepModeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiPreferencesSetDeepModeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	return toSerialize, nil
}

func (o *AiPreferencesSetDeepModeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
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

	varAiPreferencesSetDeepModeRequest := _AiPreferencesSetDeepModeRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiPreferencesSetDeepModeRequest)

	if err != nil {
		return err
	}

	*o = AiPreferencesSetDeepModeRequest(varAiPreferencesSetDeepModeRequest)

	return err
}

type NullableAiPreferencesSetDeepModeRequest struct {
	value *AiPreferencesSetDeepModeRequest
	isSet bool
}

func (v NullableAiPreferencesSetDeepModeRequest) Get() *AiPreferencesSetDeepModeRequest {
	return v.value
}

func (v *NullableAiPreferencesSetDeepModeRequest) Set(val *AiPreferencesSetDeepModeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiPreferencesSetDeepModeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiPreferencesSetDeepModeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiPreferencesSetDeepModeRequest(val *AiPreferencesSetDeepModeRequest) *NullableAiPreferencesSetDeepModeRequest {
	return &NullableAiPreferencesSetDeepModeRequest{value: val, isSet: true}
}

func (v NullableAiPreferencesSetDeepModeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiPreferencesSetDeepModeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

