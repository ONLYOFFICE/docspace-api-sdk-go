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

// checks if the AiToolsSetAllowAlwaysRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiToolsSetAllowAlwaysRequest{}

// AiToolsSetAllowAlwaysRequest struct for AiToolsSetAllowAlwaysRequest
type AiToolsSetAllowAlwaysRequest struct {
	ServerType string `json:"serverType"`
	ToolName string `json:"toolName"`
	// Whether the tool is always allowed.
	Value bool `json:"value"`
	EntityId *string `json:"entityId,omitempty"`
}

type _AiToolsSetAllowAlwaysRequest AiToolsSetAllowAlwaysRequest

// NewAiToolsSetAllowAlwaysRequest instantiates a new AiToolsSetAllowAlwaysRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiToolsSetAllowAlwaysRequest(serverType string, toolName string, value bool) *AiToolsSetAllowAlwaysRequest {
	this := AiToolsSetAllowAlwaysRequest{}
	this.ServerType = serverType
	this.ToolName = toolName
	this.Value = value
	return &this
}

// NewAiToolsSetAllowAlwaysRequestWithDefaults instantiates a new AiToolsSetAllowAlwaysRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiToolsSetAllowAlwaysRequestWithDefaults() *AiToolsSetAllowAlwaysRequest {
	this := AiToolsSetAllowAlwaysRequest{}
	return &this
}

// GetServerType returns the ServerType field value
func (o *AiToolsSetAllowAlwaysRequest) GetServerType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerType
}

// GetServerTypeOk returns a tuple with the ServerType field value
// and a boolean to check if the value has been set.
func (o *AiToolsSetAllowAlwaysRequest) GetServerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerType, true
}

// SetServerType sets field value
func (o *AiToolsSetAllowAlwaysRequest) SetServerType(v string) {
	o.ServerType = v
}

// GetToolName returns the ToolName field value
func (o *AiToolsSetAllowAlwaysRequest) GetToolName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToolName
}

// GetToolNameOk returns a tuple with the ToolName field value
// and a boolean to check if the value has been set.
func (o *AiToolsSetAllowAlwaysRequest) GetToolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToolName, true
}

// SetToolName sets field value
func (o *AiToolsSetAllowAlwaysRequest) SetToolName(v string) {
	o.ToolName = v
}

// GetValue returns the Value field value
func (o *AiToolsSetAllowAlwaysRequest) GetValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AiToolsSetAllowAlwaysRequest) GetValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AiToolsSetAllowAlwaysRequest) SetValue(v bool) {
	o.Value = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiToolsSetAllowAlwaysRequest) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiToolsSetAllowAlwaysRequest) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiToolsSetAllowAlwaysRequest) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiToolsSetAllowAlwaysRequest) SetEntityId(v string) {
	o.EntityId = &v
}

func (o AiToolsSetAllowAlwaysRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiToolsSetAllowAlwaysRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serverType"] = o.ServerType
	toSerialize["toolName"] = o.ToolName
	toSerialize["value"] = o.Value
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	return toSerialize, nil
}

func (o *AiToolsSetAllowAlwaysRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"serverType",
		"toolName",
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

	varAiToolsSetAllowAlwaysRequest := _AiToolsSetAllowAlwaysRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiToolsSetAllowAlwaysRequest)

	if err != nil {
		return err
	}

	*o = AiToolsSetAllowAlwaysRequest(varAiToolsSetAllowAlwaysRequest)

	return err
}

type NullableAiToolsSetAllowAlwaysRequest struct {
	value *AiToolsSetAllowAlwaysRequest
	isSet bool
}

func (v NullableAiToolsSetAllowAlwaysRequest) Get() *AiToolsSetAllowAlwaysRequest {
	return v.value
}

func (v *NullableAiToolsSetAllowAlwaysRequest) Set(val *AiToolsSetAllowAlwaysRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiToolsSetAllowAlwaysRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiToolsSetAllowAlwaysRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiToolsSetAllowAlwaysRequest(val *AiToolsSetAllowAlwaysRequest) *NullableAiToolsSetAllowAlwaysRequest {
	return &NullableAiToolsSetAllowAlwaysRequest{value: val, isSet: true}
}

func (v NullableAiToolsSetAllowAlwaysRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiToolsSetAllowAlwaysRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

