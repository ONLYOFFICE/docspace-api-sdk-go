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

// checks if the McpToolDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &McpToolDto{}

// McpToolDto Represents a single tool exposed by an MCP server, along with its enabled or disabled state within a room.
type McpToolDto struct {
	// Name of the tool as reported by the MCP server.
	Name NullableString `json:"name"`
	// Indicates whether this tool is enabled (true) or disabled (false) for use in AI chat sessions within the room.
	Enabled *bool `json:"enabled,omitempty"`
}

type _McpToolDto McpToolDto

// NewMcpToolDto instantiates a new McpToolDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMcpToolDto(name NullableString) *McpToolDto {
	this := McpToolDto{}
	this.Name = name
	return &this
}

// NewMcpToolDtoWithDefaults instantiates a new McpToolDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMcpToolDtoWithDefaults() *McpToolDto {
	this := McpToolDto{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *McpToolDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *McpToolDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *McpToolDto) SetName(v string) {
	o.Name.Set(&v)
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *McpToolDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McpToolDto) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *McpToolDto) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *McpToolDto) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o McpToolDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o McpToolDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name.Get()
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

func (o *McpToolDto) UnmarshalJSON(data []byte) (err error) {
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

	varMcpToolDto := _McpToolDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varMcpToolDto)

	if err != nil {
		return err
	}

	*o = McpToolDto(varMcpToolDto)

	return err
}

type NullableMcpToolDto struct {
	value *McpToolDto
	isSet bool
}

func (v NullableMcpToolDto) Get() *McpToolDto {
	return v.value
}

func (v *NullableMcpToolDto) Set(val *McpToolDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMcpToolDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMcpToolDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMcpToolDto(val *McpToolDto) *NullableMcpToolDto {
	return &NullableMcpToolDto{value: val, isSet: true}
}

func (v NullableMcpToolDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMcpToolDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

