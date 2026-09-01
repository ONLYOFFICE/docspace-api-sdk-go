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

// checks if the AiToolsSetDisabledRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiToolsSetDisabledRequest{}

// AiToolsSetDisabledRequest struct for AiToolsSetDisabledRequest
type AiToolsSetDisabledRequest struct {
	ServerType string `json:"serverType"`
	// Tool names to disable.
	ToolNames []string `json:"toolNames"`
	EntityId *string `json:"entityId,omitempty"`
}

type _AiToolsSetDisabledRequest AiToolsSetDisabledRequest

// NewAiToolsSetDisabledRequest instantiates a new AiToolsSetDisabledRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiToolsSetDisabledRequest(serverType string, toolNames []string) *AiToolsSetDisabledRequest {
	this := AiToolsSetDisabledRequest{}
	this.ServerType = serverType
	this.ToolNames = toolNames
	return &this
}

// NewAiToolsSetDisabledRequestWithDefaults instantiates a new AiToolsSetDisabledRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiToolsSetDisabledRequestWithDefaults() *AiToolsSetDisabledRequest {
	this := AiToolsSetDisabledRequest{}
	return &this
}

// GetServerType returns the ServerType field value
func (o *AiToolsSetDisabledRequest) GetServerType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerType
}

// GetServerTypeOk returns a tuple with the ServerType field value
// and a boolean to check if the value has been set.
func (o *AiToolsSetDisabledRequest) GetServerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerType, true
}

// SetServerType sets field value
func (o *AiToolsSetDisabledRequest) SetServerType(v string) {
	o.ServerType = v
}

// GetToolNames returns the ToolNames field value
func (o *AiToolsSetDisabledRequest) GetToolNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ToolNames
}

// GetToolNamesOk returns a tuple with the ToolNames field value
// and a boolean to check if the value has been set.
func (o *AiToolsSetDisabledRequest) GetToolNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToolNames, true
}

// SetToolNames sets field value
func (o *AiToolsSetDisabledRequest) SetToolNames(v []string) {
	o.ToolNames = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiToolsSetDisabledRequest) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiToolsSetDisabledRequest) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiToolsSetDisabledRequest) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiToolsSetDisabledRequest) SetEntityId(v string) {
	o.EntityId = &v
}

func (o AiToolsSetDisabledRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiToolsSetDisabledRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serverType"] = o.ServerType
	toSerialize["toolNames"] = o.ToolNames
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	return toSerialize, nil
}

func (o *AiToolsSetDisabledRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"serverType",
		"toolNames",
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

	varAiToolsSetDisabledRequest := _AiToolsSetDisabledRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiToolsSetDisabledRequest)

	if err != nil {
		return err
	}

	*o = AiToolsSetDisabledRequest(varAiToolsSetDisabledRequest)

	return err
}

type NullableAiToolsSetDisabledRequest struct {
	value *AiToolsSetDisabledRequest
	isSet bool
}

func (v NullableAiToolsSetDisabledRequest) Get() *AiToolsSetDisabledRequest {
	return v.value
}

func (v *NullableAiToolsSetDisabledRequest) Set(val *AiToolsSetDisabledRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiToolsSetDisabledRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiToolsSetDisabledRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiToolsSetDisabledRequest(val *AiToolsSetDisabledRequest) *NullableAiToolsSetDisabledRequest {
	return &NullableAiToolsSetDisabledRequest{value: val, isSet: true}
}

func (v NullableAiToolsSetDisabledRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiToolsSetDisabledRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

