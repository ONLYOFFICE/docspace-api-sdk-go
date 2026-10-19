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

// checks if the AiMCPItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiMCPItem{}

// AiMCPItem Descriptor for a tool exposed by an MCP server.
type AiMCPItem struct {
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
	// Whether the consumer must show an approval dialog before this tool runs. Feeds the `autoAllow` flag on a `tool-call-pending` event together with the user's tool permission mode: `false` skips the dialog under the auto mode, the default (under ask only the persisted always-allow list does), `true` and `undefined` defer to that list; allow skips it for every tool. Host tools set it and default to `false`; MCP / custom-server tools leave it unset. Wire-serializable, so it survives a remote (server-side) engine.
	RequireApproval *bool `json:"requireApproval,omitempty"`
	// The MCP tool annotations as the server declared them in `tools/list` (kept verbatim on the descriptor; never sent to the model). The approval flow reads two of them under the auto permission mode: `readOnlyHint: true` and `destructiveHint: false` run without the dialog, a destructive or unannotated tool keeps asking — see `resolveAutoAllow`. Wire-serializable.
	Annotations *AiToolAnnotations `json:"annotations,omitempty"`
}

type _AiMCPItem AiMCPItem

// NewAiMCPItem instantiates a new AiMCPItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiMCPItem(name string, description string, inputSchema map[string]interface{}) *AiMCPItem {
	this := AiMCPItem{}
	this.Name = name
	this.Description = description
	this.InputSchema = inputSchema
	return &this
}

// NewAiMCPItemWithDefaults instantiates a new AiMCPItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiMCPItemWithDefaults() *AiMCPItem {
	this := AiMCPItem{}
	return &this
}

// GetName returns the Name field value
func (o *AiMCPItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiMCPItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiMCPItem) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *AiMCPItem) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AiMCPItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AiMCPItem) SetDescription(v string) {
	o.Description = v
}

// GetInputSchema returns the InputSchema field value
func (o *AiMCPItem) GetInputSchema() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.InputSchema
}

// GetInputSchemaOk returns a tuple with the InputSchema field value
// and a boolean to check if the value has been set.
func (o *AiMCPItem) GetInputSchemaOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.InputSchema, true
}

// SetInputSchema sets field value
func (o *AiMCPItem) SetInputSchema(v map[string]interface{}) {
	o.InputSchema = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AiMCPItem) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMCPItem) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AiMCPItem) IsEnabledSet() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AiMCPItem) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetServerType returns the ServerType field value if set, zero value otherwise.
func (o *AiMCPItem) GetServerType() string {
	if o == nil || IsNil(o.ServerType) {
		var ret string
		return ret
	}
	return *o.ServerType
}

// GetServerTypeOk returns a tuple with the ServerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMCPItem) GetServerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ServerType) {
		return nil, false
	}
	return o.ServerType, true
}

// HasServerType returns a boolean if a field has been set.
func (o *AiMCPItem) IsServerTypeSet() bool {
	if o != nil && !IsNil(o.ServerType) {
		return true
	}

	return false
}

// SetServerType gets a reference to the given string and assigns it to the ServerType field.
func (o *AiMCPItem) SetServerType(v string) {
	o.ServerType = &v
}

// GetRequireApproval returns the RequireApproval field value if set, zero value otherwise.
func (o *AiMCPItem) GetRequireApproval() bool {
	if o == nil || IsNil(o.RequireApproval) {
		var ret bool
		return ret
	}
	return *o.RequireApproval
}

// GetRequireApprovalOk returns a tuple with the RequireApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMCPItem) GetRequireApprovalOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireApproval) {
		return nil, false
	}
	return o.RequireApproval, true
}

// HasRequireApproval returns a boolean if a field has been set.
func (o *AiMCPItem) IsRequireApprovalSet() bool {
	if o != nil && !IsNil(o.RequireApproval) {
		return true
	}

	return false
}

// SetRequireApproval gets a reference to the given bool and assigns it to the RequireApproval field.
func (o *AiMCPItem) SetRequireApproval(v bool) {
	o.RequireApproval = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *AiMCPItem) GetAnnotations() AiToolAnnotations {
	if o == nil || IsNil(o.Annotations) {
		var ret AiToolAnnotations
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMCPItem) GetAnnotationsOk() (*AiToolAnnotations, bool) {
	if o == nil || IsNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *AiMCPItem) IsAnnotationsSet() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given AiToolAnnotations and assigns it to the Annotations field.
func (o *AiMCPItem) SetAnnotations(v AiToolAnnotations) {
	o.Annotations = &v
}

func (o AiMCPItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiMCPItem) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	return toSerialize, nil
}

func (o *AiMCPItem) UnmarshalJSON(data []byte) (err error) {
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

	varAiMCPItem := _AiMCPItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiMCPItem)

	if err != nil {
		return err
	}

	*o = AiMCPItem(varAiMCPItem)

	return err
}

type NullableAiMCPItem struct {
	value *AiMCPItem
	isSet bool
}

func (v NullableAiMCPItem) Get() *AiMCPItem {
	return v.value
}

func (v *NullableAiMCPItem) Set(val *AiMCPItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAiMCPItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAiMCPItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiMCPItem(val *AiMCPItem) *NullableAiMCPItem {
	return &NullableAiMCPItem{value: val, isSet: true}
}

func (v NullableAiMCPItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiMCPItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

