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

// checks if the SetMcpToolsRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetMcpToolsRequestBody{}

// SetMcpToolsRequestBody Parameters for updating the disabled tools list of an MCP server in a room.
type SetMcpToolsRequestBody struct {
	// List of tool names to disable. Tools not included in this list will remain enabled. Pass an empty list to enable all tools.
	DisabledTools []string `json:"disabledTools"`
}

type _SetMcpToolsRequestBody SetMcpToolsRequestBody

// NewSetMcpToolsRequestBody instantiates a new SetMcpToolsRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetMcpToolsRequestBody(disabledTools []string) *SetMcpToolsRequestBody {
	this := SetMcpToolsRequestBody{}
	this.DisabledTools = disabledTools
	return &this
}

// NewSetMcpToolsRequestBodyWithDefaults instantiates a new SetMcpToolsRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetMcpToolsRequestBodyWithDefaults() *SetMcpToolsRequestBody {
	this := SetMcpToolsRequestBody{}
	return &this
}

// GetDisabledTools returns the DisabledTools field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *SetMcpToolsRequestBody) GetDisabledTools() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DisabledTools
}

// GetDisabledToolsOk returns a tuple with the DisabledTools field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetMcpToolsRequestBody) GetDisabledToolsOk() ([]string, bool) {
	if o == nil || IsNil(o.DisabledTools) {
		return nil, false
	}
	return o.DisabledTools, true
}

// SetDisabledTools sets field value
func (o *SetMcpToolsRequestBody) SetDisabledTools(v []string) {
	o.DisabledTools = v
}

func (o SetMcpToolsRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetMcpToolsRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DisabledTools != nil {
		toSerialize["disabledTools"] = o.DisabledTools
	}
	return toSerialize, nil
}

func (o *SetMcpToolsRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"disabledTools",
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

	varSetMcpToolsRequestBody := _SetMcpToolsRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSetMcpToolsRequestBody)

	if err != nil {
		return err
	}

	*o = SetMcpToolsRequestBody(varSetMcpToolsRequestBody)

	return err
}

type NullableSetMcpToolsRequestBody struct {
	value *SetMcpToolsRequestBody
	isSet bool
}

func (v NullableSetMcpToolsRequestBody) Get() *SetMcpToolsRequestBody {
	return v.value
}

func (v *NullableSetMcpToolsRequestBody) Set(val *SetMcpToolsRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSetMcpToolsRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSetMcpToolsRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetMcpToolsRequestBody(val *SetMcpToolsRequestBody) *NullableSetMcpToolsRequestBody {
	return &NullableSetMcpToolsRequestBody{value: val, isSet: true}
}

func (v NullableSetMcpToolsRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetMcpToolsRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

