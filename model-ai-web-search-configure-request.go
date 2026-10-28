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

// checks if the AiWebSearchConfigureRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiWebSearchConfigureRequest{}

// AiWebSearchConfigureRequest struct for AiWebSearchConfigureRequest
type AiWebSearchConfigureRequest struct {
	Config AiWebSearchConfig `json:"config"`
	EntityId *string `json:"entityId,omitempty"`
}

type _AiWebSearchConfigureRequest AiWebSearchConfigureRequest

// NewAiWebSearchConfigureRequest instantiates a new AiWebSearchConfigureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiWebSearchConfigureRequest(config AiWebSearchConfig) *AiWebSearchConfigureRequest {
	this := AiWebSearchConfigureRequest{}
	this.Config = config
	return &this
}

// NewAiWebSearchConfigureRequestWithDefaults instantiates a new AiWebSearchConfigureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiWebSearchConfigureRequestWithDefaults() *AiWebSearchConfigureRequest {
	this := AiWebSearchConfigureRequest{}
	return &this
}

// GetConfig returns the Config field value
func (o *AiWebSearchConfigureRequest) GetConfig() AiWebSearchConfig {
	if o == nil {
		var ret AiWebSearchConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *AiWebSearchConfigureRequest) GetConfigOk() (*AiWebSearchConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *AiWebSearchConfigureRequest) SetConfig(v AiWebSearchConfig) {
	o.Config = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiWebSearchConfigureRequest) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiWebSearchConfigureRequest) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiWebSearchConfigureRequest) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiWebSearchConfigureRequest) SetEntityId(v string) {
	o.EntityId = &v
}

func (o AiWebSearchConfigureRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiWebSearchConfigureRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["config"] = o.Config
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	return toSerialize, nil
}

func (o *AiWebSearchConfigureRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varAiWebSearchConfigureRequest := _AiWebSearchConfigureRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiWebSearchConfigureRequest)

	if err != nil {
		return err
	}

	*o = AiWebSearchConfigureRequest(varAiWebSearchConfigureRequest)

	return err
}

type NullableAiWebSearchConfigureRequest struct {
	value *AiWebSearchConfigureRequest
	isSet bool
}

func (v NullableAiWebSearchConfigureRequest) Get() *AiWebSearchConfigureRequest {
	return v.value
}

func (v *NullableAiWebSearchConfigureRequest) Set(val *AiWebSearchConfigureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiWebSearchConfigureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiWebSearchConfigureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiWebSearchConfigureRequest(val *AiWebSearchConfigureRequest) *NullableAiWebSearchConfigureRequest {
	return &NullableAiWebSearchConfigureRequest{value: val, isSet: true}
}

func (v NullableAiWebSearchConfigureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiWebSearchConfigureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

