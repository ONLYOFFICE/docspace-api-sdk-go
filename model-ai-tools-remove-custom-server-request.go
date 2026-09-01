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

// checks if the AiToolsRemoveCustomServerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiToolsRemoveCustomServerRequest{}

// AiToolsRemoveCustomServerRequest struct for AiToolsRemoveCustomServerRequest
type AiToolsRemoveCustomServerRequest struct {
	Name string `json:"name"`
	EntityId *string `json:"entityId,omitempty"`
}

type _AiToolsRemoveCustomServerRequest AiToolsRemoveCustomServerRequest

// NewAiToolsRemoveCustomServerRequest instantiates a new AiToolsRemoveCustomServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiToolsRemoveCustomServerRequest(name string) *AiToolsRemoveCustomServerRequest {
	this := AiToolsRemoveCustomServerRequest{}
	this.Name = name
	return &this
}

// NewAiToolsRemoveCustomServerRequestWithDefaults instantiates a new AiToolsRemoveCustomServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiToolsRemoveCustomServerRequestWithDefaults() *AiToolsRemoveCustomServerRequest {
	this := AiToolsRemoveCustomServerRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AiToolsRemoveCustomServerRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiToolsRemoveCustomServerRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiToolsRemoveCustomServerRequest) SetName(v string) {
	o.Name = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiToolsRemoveCustomServerRequest) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiToolsRemoveCustomServerRequest) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiToolsRemoveCustomServerRequest) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiToolsRemoveCustomServerRequest) SetEntityId(v string) {
	o.EntityId = &v
}

func (o AiToolsRemoveCustomServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiToolsRemoveCustomServerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	return toSerialize, nil
}

func (o *AiToolsRemoveCustomServerRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varAiToolsRemoveCustomServerRequest := _AiToolsRemoveCustomServerRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiToolsRemoveCustomServerRequest)

	if err != nil {
		return err
	}

	*o = AiToolsRemoveCustomServerRequest(varAiToolsRemoveCustomServerRequest)

	return err
}

type NullableAiToolsRemoveCustomServerRequest struct {
	value *AiToolsRemoveCustomServerRequest
	isSet bool
}

func (v NullableAiToolsRemoveCustomServerRequest) Get() *AiToolsRemoveCustomServerRequest {
	return v.value
}

func (v *NullableAiToolsRemoveCustomServerRequest) Set(val *AiToolsRemoveCustomServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiToolsRemoveCustomServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiToolsRemoveCustomServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiToolsRemoveCustomServerRequest(val *AiToolsRemoveCustomServerRequest) *NullableAiToolsRemoveCustomServerRequest {
	return &NullableAiToolsRemoveCustomServerRequest{value: val, isSet: true}
}

func (v NullableAiToolsRemoveCustomServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiToolsRemoveCustomServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

