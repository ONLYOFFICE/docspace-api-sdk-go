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

// checks if the EditorToolCallStateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditorToolCallStateDto{}

// EditorToolCallStateDto A generation the editor is expected to run as soon as the document opens, left behind by an AI agent that created  the file but not its content.
type EditorToolCallStateDto struct {
	// Which generation to run, which also decides the shape of the parameters below.
	ToolName NullableString `json:"toolName"`
	// The arguments of the generation named above.
	Parameters EditorToolCallParametersDto `json:"parameters"`
}

type _EditorToolCallStateDto EditorToolCallStateDto

// NewEditorToolCallStateDto instantiates a new EditorToolCallStateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditorToolCallStateDto(toolName NullableString, parameters EditorToolCallParametersDto) *EditorToolCallStateDto {
	this := EditorToolCallStateDto{}
	this.ToolName = toolName
	this.Parameters = parameters
	return &this
}

// NewEditorToolCallStateDtoWithDefaults instantiates a new EditorToolCallStateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditorToolCallStateDtoWithDefaults() *EditorToolCallStateDto {
	this := EditorToolCallStateDto{}
	return &this
}

// GetToolName returns the ToolName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EditorToolCallStateDto) GetToolName() string {
	if o == nil || o.ToolName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ToolName.Get()
}

// GetToolNameOk returns a tuple with the ToolName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditorToolCallStateDto) GetToolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToolName.Get(), o.ToolName.IsSet()
}

// SetToolName sets field value
func (o *EditorToolCallStateDto) SetToolName(v string) {
	o.ToolName.Set(&v)
}

// GetParameters returns the Parameters field value
func (o *EditorToolCallStateDto) GetParameters() EditorToolCallParametersDto {
	if o == nil {
		var ret EditorToolCallParametersDto
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *EditorToolCallStateDto) GetParametersOk() (*EditorToolCallParametersDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value
func (o *EditorToolCallStateDto) SetParameters(v EditorToolCallParametersDto) {
	o.Parameters = v
}

func (o EditorToolCallStateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditorToolCallStateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["toolName"] = o.ToolName.Get()
	toSerialize["parameters"] = o.Parameters
	return toSerialize, nil
}

func (o *EditorToolCallStateDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"toolName",
		"parameters",
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

	varEditorToolCallStateDto := _EditorToolCallStateDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varEditorToolCallStateDto)

	if err != nil {
		return err
	}

	*o = EditorToolCallStateDto(varEditorToolCallStateDto)

	return err
}

type NullableEditorToolCallStateDto struct {
	value *EditorToolCallStateDto
	isSet bool
}

func (v NullableEditorToolCallStateDto) Get() *EditorToolCallStateDto {
	return v.value
}

func (v *NullableEditorToolCallStateDto) Set(val *EditorToolCallStateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEditorToolCallStateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEditorToolCallStateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditorToolCallStateDto(val *EditorToolCallStateDto) *NullableEditorToolCallStateDto {
	return &NullableEditorToolCallStateDto{value: val, isSet: true}
}

func (v NullableEditorToolCallStateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditorToolCallStateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

