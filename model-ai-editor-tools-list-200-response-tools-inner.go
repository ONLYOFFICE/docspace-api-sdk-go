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

// checks if the AiEditorToolsList200ResponseToolsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiEditorToolsList200ResponseToolsInner{}

// AiEditorToolsList200ResponseToolsInner struct for AiEditorToolsList200ResponseToolsInner
type AiEditorToolsList200ResponseToolsInner struct {
	// Tool name, as it is passed back to the call endpoint.
	Name string `json:"name"`
	// What the tool does, empty when the server declares nothing.
	Description string `json:"description"`
	// JSON Schema of the tool arguments.
	InputSchema map[string]*interface{} `json:"inputSchema"`
	// Whether the editor has to ask the user before running the tool. Read-only operations arrive with this off.
	RequireApproval bool `json:"requireApproval"`
}

type _AiEditorToolsList200ResponseToolsInner AiEditorToolsList200ResponseToolsInner

// NewAiEditorToolsList200ResponseToolsInner instantiates a new AiEditorToolsList200ResponseToolsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiEditorToolsList200ResponseToolsInner(name string, description string, inputSchema map[string]*interface{}, requireApproval bool) *AiEditorToolsList200ResponseToolsInner {
	this := AiEditorToolsList200ResponseToolsInner{}
	this.Name = name
	this.Description = description
	this.InputSchema = inputSchema
	this.RequireApproval = requireApproval
	return &this
}

// NewAiEditorToolsList200ResponseToolsInnerWithDefaults instantiates a new AiEditorToolsList200ResponseToolsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiEditorToolsList200ResponseToolsInnerWithDefaults() *AiEditorToolsList200ResponseToolsInner {
	this := AiEditorToolsList200ResponseToolsInner{}
	return &this
}

// GetName returns the Name field value
func (o *AiEditorToolsList200ResponseToolsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiEditorToolsList200ResponseToolsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiEditorToolsList200ResponseToolsInner) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *AiEditorToolsList200ResponseToolsInner) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AiEditorToolsList200ResponseToolsInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AiEditorToolsList200ResponseToolsInner) SetDescription(v string) {
	o.Description = v
}

// GetInputSchema returns the InputSchema field value
func (o *AiEditorToolsList200ResponseToolsInner) GetInputSchema() map[string]*interface{} {
	if o == nil {
		var ret map[string]*interface{}
		return ret
	}

	return o.InputSchema
}

// GetInputSchemaOk returns a tuple with the InputSchema field value
// and a boolean to check if the value has been set.
func (o *AiEditorToolsList200ResponseToolsInner) GetInputSchemaOk() (map[string]*interface{}, bool) {
	if o == nil {
		return map[string]*interface{}{}, false
	}
	return o.InputSchema, true
}

// SetInputSchema sets field value
func (o *AiEditorToolsList200ResponseToolsInner) SetInputSchema(v map[string]*interface{}) {
	o.InputSchema = v
}

// GetRequireApproval returns the RequireApproval field value
func (o *AiEditorToolsList200ResponseToolsInner) GetRequireApproval() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequireApproval
}

// GetRequireApprovalOk returns a tuple with the RequireApproval field value
// and a boolean to check if the value has been set.
func (o *AiEditorToolsList200ResponseToolsInner) GetRequireApprovalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireApproval, true
}

// SetRequireApproval sets field value
func (o *AiEditorToolsList200ResponseToolsInner) SetRequireApproval(v bool) {
	o.RequireApproval = v
}

func (o AiEditorToolsList200ResponseToolsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiEditorToolsList200ResponseToolsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["inputSchema"] = o.InputSchema
	toSerialize["requireApproval"] = o.RequireApproval
	return toSerialize, nil
}

func (o *AiEditorToolsList200ResponseToolsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
		"inputSchema",
		"requireApproval",
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

	varAiEditorToolsList200ResponseToolsInner := _AiEditorToolsList200ResponseToolsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiEditorToolsList200ResponseToolsInner)

	if err != nil {
		return err
	}

	*o = AiEditorToolsList200ResponseToolsInner(varAiEditorToolsList200ResponseToolsInner)

	return err
}

type NullableAiEditorToolsList200ResponseToolsInner struct {
	value *AiEditorToolsList200ResponseToolsInner
	isSet bool
}

func (v NullableAiEditorToolsList200ResponseToolsInner) Get() *AiEditorToolsList200ResponseToolsInner {
	return v.value
}

func (v *NullableAiEditorToolsList200ResponseToolsInner) Set(val *AiEditorToolsList200ResponseToolsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAiEditorToolsList200ResponseToolsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAiEditorToolsList200ResponseToolsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiEditorToolsList200ResponseToolsInner(val *AiEditorToolsList200ResponseToolsInner) *NullableAiEditorToolsList200ResponseToolsInner {
	return &NullableAiEditorToolsList200ResponseToolsInner{value: val, isSet: true}
}

func (v NullableAiEditorToolsList200ResponseToolsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiEditorToolsList200ResponseToolsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

