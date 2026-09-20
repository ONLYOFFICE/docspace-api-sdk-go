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

// checks if the AiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiModel{}

// AiModel AI model metadata. Describes a single model available from a provider.
type AiModel struct {
	// Model identifier as used by the provider API (e.g. `gpt-4o`, `claude-sonnet-4-20250514`).
	Id string `json:"id"`
	// Human-readable model name for display in the UI.
	Name string `json:"name"`
	// Provider that offers this model.
	Provider AiProviderType `json:"provider"`
	// Whether this model supports extended thinking / chain-of-thought reasoning.
	Reasoning *bool `json:"reasoning,omitempty"`
	// What the model can do with extended thinking, when the provider's catalogue says so (OpenRouter and the ONLYOFFICE route report a per-model `reasoning` object). Copied onto the profile at save time; absent, the widget falls back to the provider's id-based table.
	ReasoningSupport *AiReasoningSupport `json:"reasoningSupport,omitempty"`
	// Bitmask of model capabilities (Chat, Image, Vision, Tools, etc.). Used to filter models per `ActionType`.
	Capabilities *float32 `json:"capabilities,omitempty"`
}

type _AiModel AiModel

// NewAiModel instantiates a new AiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiModel(id string, name string, provider AiProviderType) *AiModel {
	this := AiModel{}
	this.Id = id
	this.Name = name
	this.Provider = provider
	return &this
}

// NewAiModelWithDefaults instantiates a new AiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiModelWithDefaults() *AiModel {
	this := AiModel{}
	return &this
}

// GetId returns the Id field value
func (o *AiModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AiModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiModel) SetName(v string) {
	o.Name = v
}

// GetProvider returns the Provider field value
func (o *AiModel) GetProvider() AiProviderType {
	if o == nil {
		var ret AiProviderType
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *AiModel) GetProviderOk() (*AiProviderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *AiModel) SetProvider(v AiProviderType) {
	o.Provider = v
}

// GetReasoning returns the Reasoning field value if set, zero value otherwise.
func (o *AiModel) GetReasoning() bool {
	if o == nil || IsNil(o.Reasoning) {
		var ret bool
		return ret
	}
	return *o.Reasoning
}

// GetReasoningOk returns a tuple with the Reasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModel) GetReasoningOk() (*bool, bool) {
	if o == nil || IsNil(o.Reasoning) {
		return nil, false
	}
	return o.Reasoning, true
}

// HasReasoning returns a boolean if a field has been set.
func (o *AiModel) IsReasoningSet() bool {
	if o != nil && !IsNil(o.Reasoning) {
		return true
	}

	return false
}

// SetReasoning gets a reference to the given bool and assigns it to the Reasoning field.
func (o *AiModel) SetReasoning(v bool) {
	o.Reasoning = &v
}

// GetReasoningSupport returns the ReasoningSupport field value if set, zero value otherwise.
func (o *AiModel) GetReasoningSupport() AiReasoningSupport {
	if o == nil || IsNil(o.ReasoningSupport) {
		var ret AiReasoningSupport
		return ret
	}
	return *o.ReasoningSupport
}

// GetReasoningSupportOk returns a tuple with the ReasoningSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModel) GetReasoningSupportOk() (*AiReasoningSupport, bool) {
	if o == nil || IsNil(o.ReasoningSupport) {
		return nil, false
	}
	return o.ReasoningSupport, true
}

// HasReasoningSupport returns a boolean if a field has been set.
func (o *AiModel) IsReasoningSupportSet() bool {
	if o != nil && !IsNil(o.ReasoningSupport) {
		return true
	}

	return false
}

// SetReasoningSupport gets a reference to the given AiReasoningSupport and assigns it to the ReasoningSupport field.
func (o *AiModel) SetReasoningSupport(v AiReasoningSupport) {
	o.ReasoningSupport = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *AiModel) GetCapabilities() float32 {
	if o == nil || IsNil(o.Capabilities) {
		var ret float32
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiModel) GetCapabilitiesOk() (*float32, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *AiModel) IsCapabilitiesSet() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given float32 and assigns it to the Capabilities field.
func (o *AiModel) SetCapabilities(v float32) {
	o.Capabilities = &v
}

func (o AiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["provider"] = o.Provider
	if !IsNil(o.Reasoning) {
		toSerialize["reasoning"] = o.Reasoning
	}
	if !IsNil(o.ReasoningSupport) {
		toSerialize["reasoningSupport"] = o.ReasoningSupport
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	return toSerialize, nil
}

func (o *AiModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"provider",
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

	varAiModel := _AiModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiModel)

	if err != nil {
		return err
	}

	*o = AiModel(varAiModel)

	return err
}

type NullableAiModel struct {
	value *AiModel
	isSet bool
}

func (v NullableAiModel) Get() *AiModel {
	return v.value
}

func (v *NullableAiModel) Set(val *AiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiModel(val *AiModel) *NullableAiModel {
	return &NullableAiModel{value: val, isSet: true}
}

func (v NullableAiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

