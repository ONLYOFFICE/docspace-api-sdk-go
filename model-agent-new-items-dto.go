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

// checks if the AgentNewItemsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentNewItemsDto{}

// AgentNewItemsDto The agent new item's information.
type AgentNewItemsDto struct {
	Agent FileEntryBaseDto `json:"agent"`
	// The list of file entry items.
	Items []FileEntryBaseDto `json:"items"`
}

type _AgentNewItemsDto AgentNewItemsDto

// NewAgentNewItemsDto instantiates a new AgentNewItemsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentNewItemsDto(agent FileEntryBaseDto, items []FileEntryBaseDto) *AgentNewItemsDto {
	this := AgentNewItemsDto{}
	this.Agent = agent
	this.Items = items
	return &this
}

// NewAgentNewItemsDtoWithDefaults instantiates a new AgentNewItemsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentNewItemsDtoWithDefaults() *AgentNewItemsDto {
	this := AgentNewItemsDto{}
	return &this
}

// GetAgent returns the Agent field value
func (o *AgentNewItemsDto) GetAgent() FileEntryBaseDto {
	if o == nil {
		var ret FileEntryBaseDto
		return ret
	}

	return o.Agent
}

// GetAgentOk returns a tuple with the Agent field value
// and a boolean to check if the value has been set.
func (o *AgentNewItemsDto) GetAgentOk() (*FileEntryBaseDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Agent, true
}

// SetAgent sets field value
func (o *AgentNewItemsDto) SetAgent(v FileEntryBaseDto) {
	o.Agent = v
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []FileEntryBaseDto will be returned
func (o *AgentNewItemsDto) GetItems() []FileEntryBaseDto {
	if o == nil {
		var ret []FileEntryBaseDto
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AgentNewItemsDto) GetItemsOk() ([]FileEntryBaseDto, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *AgentNewItemsDto) SetItems(v []FileEntryBaseDto) {
	o.Items = v
}

func (o AgentNewItemsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentNewItemsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agent"] = o.Agent
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

func (o *AgentNewItemsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agent",
		"items",
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

	varAgentNewItemsDto := _AgentNewItemsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAgentNewItemsDto)

	if err != nil {
		return err
	}

	*o = AgentNewItemsDto(varAgentNewItemsDto)

	return err
}

type NullableAgentNewItemsDto struct {
	value *AgentNewItemsDto
	isSet bool
}

func (v NullableAgentNewItemsDto) Get() *AgentNewItemsDto {
	return v.value
}

func (v *NullableAgentNewItemsDto) Set(val *AgentNewItemsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentNewItemsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentNewItemsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentNewItemsDto(val *AgentNewItemsDto) *NullableAgentNewItemsDto {
	return &NullableAgentNewItemsDto{value: val, isSet: true}
}

func (v NullableAgentNewItemsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentNewItemsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

