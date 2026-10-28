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

// checks if the AiToolsAddCustomServerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiToolsAddCustomServerRequest{}

// AiToolsAddCustomServerRequest struct for AiToolsAddCustomServerRequest
type AiToolsAddCustomServerRequest struct {
	// Server name (unique within scope).
	Name string `json:"name"`
	// Server transport configuration.
	Config map[string]interface{} `json:"config"`
	EntityId *string `json:"entityId,omitempty"`
}

type _AiToolsAddCustomServerRequest AiToolsAddCustomServerRequest

// NewAiToolsAddCustomServerRequest instantiates a new AiToolsAddCustomServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiToolsAddCustomServerRequest(name string, config map[string]interface{}) *AiToolsAddCustomServerRequest {
	this := AiToolsAddCustomServerRequest{}
	this.Name = name
	this.Config = config
	return &this
}

// NewAiToolsAddCustomServerRequestWithDefaults instantiates a new AiToolsAddCustomServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiToolsAddCustomServerRequestWithDefaults() *AiToolsAddCustomServerRequest {
	this := AiToolsAddCustomServerRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AiToolsAddCustomServerRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiToolsAddCustomServerRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiToolsAddCustomServerRequest) SetName(v string) {
	o.Name = v
}

// GetConfig returns the Config field value
func (o *AiToolsAddCustomServerRequest) GetConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *AiToolsAddCustomServerRequest) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// SetConfig sets field value
func (o *AiToolsAddCustomServerRequest) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiToolsAddCustomServerRequest) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiToolsAddCustomServerRequest) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiToolsAddCustomServerRequest) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiToolsAddCustomServerRequest) SetEntityId(v string) {
	o.EntityId = &v
}

func (o AiToolsAddCustomServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiToolsAddCustomServerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["config"] = o.Config
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	return toSerialize, nil
}

func (o *AiToolsAddCustomServerRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"config",
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

	varAiToolsAddCustomServerRequest := _AiToolsAddCustomServerRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiToolsAddCustomServerRequest)

	if err != nil {
		return err
	}

	*o = AiToolsAddCustomServerRequest(varAiToolsAddCustomServerRequest)

	return err
}

type NullableAiToolsAddCustomServerRequest struct {
	value *AiToolsAddCustomServerRequest
	isSet bool
}

func (v NullableAiToolsAddCustomServerRequest) Get() *AiToolsAddCustomServerRequest {
	return v.value
}

func (v *NullableAiToolsAddCustomServerRequest) Set(val *AiToolsAddCustomServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiToolsAddCustomServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiToolsAddCustomServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiToolsAddCustomServerRequest(val *AiToolsAddCustomServerRequest) *NullableAiToolsAddCustomServerRequest {
	return &NullableAiToolsAddCustomServerRequest{value: val, isSet: true}
}

func (v NullableAiToolsAddCustomServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiToolsAddCustomServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

