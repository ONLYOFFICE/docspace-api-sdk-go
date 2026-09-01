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

// checks if the AiTMCPItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiTMCPItem{}

// AiTMCPItem Descriptor for a tool exposed by an MCP server.
type AiTMCPItem struct {
	// Tool name as registered on the MCP server (e.g. `web_search`, `insert_text`).
	Name string `json:"name"`
	// Human-readable description shown to the AI model and in the tools list UI.
	Description string `json:"description"`
	// JSON Schema describing the tool's input parameters.
	InputSchema map[string]interface{} `json:"inputSchema"`
	// Whether this tool is currently enabled. Disabled tools are hidden from the AI model.
	Enabled *bool `json:"enabled,omitempty"`
	// Server type (MCP server name / host tool group id) this tool belongs to — the key the persisted disabled map is stored under. Set by the source that enumerated the tool, so a caller-supplied tool can still be attributed to its group after being flattened into a single list: that is what lets the engine apply the disabled map to `actionArgs.tools` instead of trusting the caller to pre-filter. Wire-serializable, so it survives a remote (server-side) engine.
	ServerType *string `json:"serverType,omitempty"`
	// Whether the consumer must show an approval dialog before this tool runs. The engine reads it when deciding the `autoAllow` flag on a `tool-call-pending` event: `requireApproval === false` auto-allows the call (no dialog), `true` always prompts. `undefined` leaves the decision to the persisted always-allow list alone — so MCP / custom-server tools (which never set it) keep prompting as before, while host tools opt into auto-allow by default. Wire-serializable, so it survives a remote (server-side) engine.
	RequireApproval *bool `json:"requireApproval,omitempty"`
}

type _AiTMCPItem AiTMCPItem

// NewAiTMCPItem instantiates a new AiTMCPItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiTMCPItem(name string, description string, inputSchema map[string]interface{}) *AiTMCPItem {
	this := AiTMCPItem{}
	this.Name = name
	this.Description = description
	this.InputSchema = inputSchema
	return &this
}

// NewAiTMCPItemWithDefaults instantiates a new AiTMCPItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiTMCPItemWithDefaults() *AiTMCPItem {
	this := AiTMCPItem{}
	return &this
}

// GetName returns the Name field value
func (o *AiTMCPItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiTMCPItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiTMCPItem) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *AiTMCPItem) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AiTMCPItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AiTMCPItem) SetDescription(v string) {
	o.Description = v
}

// GetInputSchema returns the InputSchema field value
func (o *AiTMCPItem) GetInputSchema() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.InputSchema
}

// GetInputSchemaOk returns a tuple with the InputSchema field value
// and a boolean to check if the value has been set.
func (o *AiTMCPItem) GetInputSchemaOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.InputSchema, true
}

// SetInputSchema sets field value
func (o *AiTMCPItem) SetInputSchema(v map[string]interface{}) {
	o.InputSchema = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AiTMCPItem) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiTMCPItem) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AiTMCPItem) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AiTMCPItem) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetServerType returns the ServerType field value if set, zero value otherwise.
func (o *AiTMCPItem) GetServerType() string {
	if o == nil || IsNil(o.ServerType) {
		var ret string
		return ret
	}
	return *o.ServerType
}

// GetServerTypeOk returns a tuple with the ServerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiTMCPItem) GetServerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ServerType) {
		return nil, false
	}
	return o.ServerType, true
}

// HasServerType returns a boolean if a field has been set.
func (o *AiTMCPItem) IsServerTypeSet() bool {
	if o != nil && !IsNil(o.ServerType) {
		return true
	}

	return false
}

// SetServerType gets a reference to the given string and assigns it to the ServerType field.
func (o *AiTMCPItem) SetServerType(v string) {
	o.ServerType = &v
}

// GetRequireApproval returns the RequireApproval field value if set, zero value otherwise.
func (o *AiTMCPItem) GetRequireApproval() bool {
	if o == nil || IsNil(o.RequireApproval) {
		var ret bool
		return ret
	}
	return *o.RequireApproval
}

// GetRequireApprovalOk returns a tuple with the RequireApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiTMCPItem) GetRequireApprovalOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireApproval) {
		return nil, false
	}
	return o.RequireApproval, true
}

// HasRequireApproval returns a boolean if a field has been set.
func (o *AiTMCPItem) IsRequireApprovalSet() bool {
	if o != nil && !IsNil(o.RequireApproval) {
		return true
	}

	return false
}

// SetRequireApproval gets a reference to the given bool and assigns it to the RequireApproval field.
func (o *AiTMCPItem) SetRequireApproval(v bool) {
	o.RequireApproval = &v
}

func (o AiTMCPItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiTMCPItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["inputSchema"] = o.InputSchema
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ServerType) {
		toSerialize["serverType"] = o.ServerType
	}
	if !IsNil(o.RequireApproval) {
		toSerialize["requireApproval"] = o.RequireApproval
	}
	return toSerialize, nil
}

func (o *AiTMCPItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
		"inputSchema",
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

	varAiTMCPItem := _AiTMCPItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiTMCPItem)

	if err != nil {
		return err
	}

	*o = AiTMCPItem(varAiTMCPItem)

	return err
}

type NullableAiTMCPItem struct {
	value *AiTMCPItem
	isSet bool
}

func (v NullableAiTMCPItem) Get() *AiTMCPItem {
	return v.value
}

func (v *NullableAiTMCPItem) Set(val *AiTMCPItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAiTMCPItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAiTMCPItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiTMCPItem(val *AiTMCPItem) *NullableAiTMCPItem {
	return &NullableAiTMCPItem{value: val, isSet: true}
}

func (v NullableAiTMCPItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiTMCPItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

