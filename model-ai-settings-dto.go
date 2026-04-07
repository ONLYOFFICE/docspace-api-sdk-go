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

// checks if the AiSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiSettingsDto{}

// AiSettingsDto The AI module settings.
type AiSettingsDto struct {
	// Indicates whether web search is enabled for AI chat sessions.
	WebSearchEnabled *bool `json:"webSearchEnabled,omitempty"`
	// Indicates whether the web search API key needs to be reconfigured.
	WebSearchNeedReset *bool `json:"webSearchNeedReset,omitempty"`
	// Indicates whether document vectorization is enabled.
	VectorizationEnabled *bool `json:"vectorizationEnabled,omitempty"`
	// Indicates whether the embedding provider API key needs to be reconfigured.
	VectorizationNeedReset *bool `json:"vectorizationNeedReset,omitempty"`
	// Indicates whether the AI subsystem is fully configured and operational.
	AiReady *bool `json:"aiReady,omitempty"`
	// Indicates whether the AI provider API key needs to be reconfigured.
	AiReadyNeedReset *bool `json:"aiReadyNeedReset,omitempty"`
	// The unique identifier of the portal-level MCP server, if configured.
	PortalMcpServerId NullableString `json:"portalMcpServerId,omitempty"`
	// The name of the embedding model used for document vectorization.
	EmbeddingModel NullableString `json:"embeddingModel"`
	// Mapping of model identifiers to human-readable aliases.
	ModelAliases map[string]string `json:"modelAliases"`
	// The tool name used by the AI assistant for knowledge base search.
	KnowledgeSearchToolName NullableString `json:"knowledgeSearchToolName"`
	// The tool name used by the AI assistant for web search.
	WebSearchToolName NullableString `json:"webSearchToolName"`
	// The tool name used by the AI assistant for web page crawling.
	WebCrawlingToolName NullableString `json:"webCrawlingToolName"`
	// The tool name used by the AI to launch docx creation in the editor.
	GenerateDocxToolName NullableString `json:"generateDocxToolName"`
	// The tool name used by the AI assistant to launch form creation in the editor.
	GenerateFormToolName NullableString `json:"generateFormToolName"`
	// The tool name used by the AI assistant to launch presentation creation in the editor.
	GeneratePresentationToolName NullableString `json:"generatePresentationToolName"`
	// Indicates whether the system-level AI provider is enabled.
	SystemAiEnabled *bool `json:"systemAiEnabled,omitempty"`
}

type _AiSettingsDto AiSettingsDto

// NewAiSettingsDto instantiates a new AiSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiSettingsDto(embeddingModel NullableString, modelAliases map[string]string, knowledgeSearchToolName NullableString, webSearchToolName NullableString, webCrawlingToolName NullableString, generateDocxToolName NullableString, generateFormToolName NullableString, generatePresentationToolName NullableString) *AiSettingsDto {
	this := AiSettingsDto{}
	this.EmbeddingModel = embeddingModel
	this.ModelAliases = modelAliases
	this.KnowledgeSearchToolName = knowledgeSearchToolName
	this.WebSearchToolName = webSearchToolName
	this.WebCrawlingToolName = webCrawlingToolName
	this.GenerateDocxToolName = generateDocxToolName
	this.GenerateFormToolName = generateFormToolName
	this.GeneratePresentationToolName = generatePresentationToolName
	return &this
}

// NewAiSettingsDtoWithDefaults instantiates a new AiSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiSettingsDtoWithDefaults() *AiSettingsDto {
	this := AiSettingsDto{}
	return &this
}

// GetWebSearchEnabled returns the WebSearchEnabled field value if set, zero value otherwise.
func (o *AiSettingsDto) GetWebSearchEnabled() bool {
	if o == nil || IsNil(o.WebSearchEnabled) {
		var ret bool
		return ret
	}
	return *o.WebSearchEnabled
}

// GetWebSearchEnabledOk returns a tuple with the WebSearchEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiSettingsDto) GetWebSearchEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.WebSearchEnabled) {
		return nil, false
	}
	return o.WebSearchEnabled, true
}

// HasWebSearchEnabled returns a boolean if a field has been set.
func (o *AiSettingsDto) IsWebSearchEnabledSet() bool {
	if o != nil && !IsNil(o.WebSearchEnabled) {
		return true
	}

	return false
}

// SetWebSearchEnabled gets a reference to the given bool and assigns it to the WebSearchEnabled field.
func (o *AiSettingsDto) SetWebSearchEnabled(v bool) {
	o.WebSearchEnabled = &v
}

// GetWebSearchNeedReset returns the WebSearchNeedReset field value if set, zero value otherwise.
func (o *AiSettingsDto) GetWebSearchNeedReset() bool {
	if o == nil || IsNil(o.WebSearchNeedReset) {
		var ret bool
		return ret
	}
	return *o.WebSearchNeedReset
}

// GetWebSearchNeedResetOk returns a tuple with the WebSearchNeedReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiSettingsDto) GetWebSearchNeedResetOk() (*bool, bool) {
	if o == nil || IsNil(o.WebSearchNeedReset) {
		return nil, false
	}
	return o.WebSearchNeedReset, true
}

// HasWebSearchNeedReset returns a boolean if a field has been set.
func (o *AiSettingsDto) IsWebSearchNeedResetSet() bool {
	if o != nil && !IsNil(o.WebSearchNeedReset) {
		return true
	}

	return false
}

// SetWebSearchNeedReset gets a reference to the given bool and assigns it to the WebSearchNeedReset field.
func (o *AiSettingsDto) SetWebSearchNeedReset(v bool) {
	o.WebSearchNeedReset = &v
}

// GetVectorizationEnabled returns the VectorizationEnabled field value if set, zero value otherwise.
func (o *AiSettingsDto) GetVectorizationEnabled() bool {
	if o == nil || IsNil(o.VectorizationEnabled) {
		var ret bool
		return ret
	}
	return *o.VectorizationEnabled
}

// GetVectorizationEnabledOk returns a tuple with the VectorizationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiSettingsDto) GetVectorizationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.VectorizationEnabled) {
		return nil, false
	}
	return o.VectorizationEnabled, true
}

// HasVectorizationEnabled returns a boolean if a field has been set.
func (o *AiSettingsDto) IsVectorizationEnabledSet() bool {
	if o != nil && !IsNil(o.VectorizationEnabled) {
		return true
	}

	return false
}

// SetVectorizationEnabled gets a reference to the given bool and assigns it to the VectorizationEnabled field.
func (o *AiSettingsDto) SetVectorizationEnabled(v bool) {
	o.VectorizationEnabled = &v
}

// GetVectorizationNeedReset returns the VectorizationNeedReset field value if set, zero value otherwise.
func (o *AiSettingsDto) GetVectorizationNeedReset() bool {
	if o == nil || IsNil(o.VectorizationNeedReset) {
		var ret bool
		return ret
	}
	return *o.VectorizationNeedReset
}

// GetVectorizationNeedResetOk returns a tuple with the VectorizationNeedReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiSettingsDto) GetVectorizationNeedResetOk() (*bool, bool) {
	if o == nil || IsNil(o.VectorizationNeedReset) {
		return nil, false
	}
	return o.VectorizationNeedReset, true
}

// HasVectorizationNeedReset returns a boolean if a field has been set.
func (o *AiSettingsDto) IsVectorizationNeedResetSet() bool {
	if o != nil && !IsNil(o.VectorizationNeedReset) {
		return true
	}

	return false
}

// SetVectorizationNeedReset gets a reference to the given bool and assigns it to the VectorizationNeedReset field.
func (o *AiSettingsDto) SetVectorizationNeedReset(v bool) {
	o.VectorizationNeedReset = &v
}

// GetAiReady returns the AiReady field value if set, zero value otherwise.
func (o *AiSettingsDto) GetAiReady() bool {
	if o == nil || IsNil(o.AiReady) {
		var ret bool
		return ret
	}
	return *o.AiReady
}

// GetAiReadyOk returns a tuple with the AiReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiSettingsDto) GetAiReadyOk() (*bool, bool) {
	if o == nil || IsNil(o.AiReady) {
		return nil, false
	}
	return o.AiReady, true
}

// HasAiReady returns a boolean if a field has been set.
func (o *AiSettingsDto) IsAiReadySet() bool {
	if o != nil && !IsNil(o.AiReady) {
		return true
	}

	return false
}

// SetAiReady gets a reference to the given bool and assigns it to the AiReady field.
func (o *AiSettingsDto) SetAiReady(v bool) {
	o.AiReady = &v
}

// GetAiReadyNeedReset returns the AiReadyNeedReset field value if set, zero value otherwise.
func (o *AiSettingsDto) GetAiReadyNeedReset() bool {
	if o == nil || IsNil(o.AiReadyNeedReset) {
		var ret bool
		return ret
	}
	return *o.AiReadyNeedReset
}

// GetAiReadyNeedResetOk returns a tuple with the AiReadyNeedReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiSettingsDto) GetAiReadyNeedResetOk() (*bool, bool) {
	if o == nil || IsNil(o.AiReadyNeedReset) {
		return nil, false
	}
	return o.AiReadyNeedReset, true
}

// HasAiReadyNeedReset returns a boolean if a field has been set.
func (o *AiSettingsDto) IsAiReadyNeedResetSet() bool {
	if o != nil && !IsNil(o.AiReadyNeedReset) {
		return true
	}

	return false
}

// SetAiReadyNeedReset gets a reference to the given bool and assigns it to the AiReadyNeedReset field.
func (o *AiSettingsDto) SetAiReadyNeedReset(v bool) {
	o.AiReadyNeedReset = &v
}

// GetPortalMcpServerId returns the PortalMcpServerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiSettingsDto) GetPortalMcpServerId() string {
	if o == nil || IsNil(o.PortalMcpServerId.Get()) {
		var ret string
		return ret
	}
	return *o.PortalMcpServerId.Get()
}

// GetPortalMcpServerIdOk returns a tuple with the PortalMcpServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiSettingsDto) GetPortalMcpServerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PortalMcpServerId.Get(), o.PortalMcpServerId.IsSet()
}

// HasPortalMcpServerId returns a boolean if a field has been set.
func (o *AiSettingsDto) IsPortalMcpServerIdSet() bool {
	if o != nil && o.PortalMcpServerId.IsSet() {
		return true
	}

	return false
}

// SetPortalMcpServerId gets a reference to the given NullableString and assigns it to the PortalMcpServerId field.
func (o *AiSettingsDto) SetPortalMcpServerId(v string) {
	o.PortalMcpServerId.Set(&v)
}
// SetPortalMcpServerIdNil sets the value for PortalMcpServerId to be an explicit nil
func (o *AiSettingsDto) SetPortalMcpServerIdNil() {
	o.PortalMcpServerId.Set(nil)
}

// UnsetPortalMcpServerId ensures that no value is present for PortalMcpServerId, not even an explicit nil
func (o *AiSettingsDto) UnsetPortalMcpServerId() {
	o.PortalMcpServerId.Unset()
}

// GetEmbeddingModel returns the EmbeddingModel field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiSettingsDto) GetEmbeddingModel() string {
	if o == nil || o.EmbeddingModel.Get() == nil {
		var ret string
		return ret
	}

	return *o.EmbeddingModel.Get()
}

// GetEmbeddingModelOk returns a tuple with the EmbeddingModel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiSettingsDto) GetEmbeddingModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmbeddingModel.Get(), o.EmbeddingModel.IsSet()
}

// SetEmbeddingModel sets field value
func (o *AiSettingsDto) SetEmbeddingModel(v string) {
	o.EmbeddingModel.Set(&v)
}

// GetModelAliases returns the ModelAliases field value
// If the value is explicit nil, the zero value for map[string]string will be returned
func (o *AiSettingsDto) GetModelAliases() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.ModelAliases
}

// GetModelAliasesOk returns a tuple with the ModelAliases field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiSettingsDto) GetModelAliasesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ModelAliases) {
		return nil, false
	}
	return &o.ModelAliases, true
}

// SetModelAliases sets field value
func (o *AiSettingsDto) SetModelAliases(v map[string]string) {
	o.ModelAliases = v
}

// GetKnowledgeSearchToolName returns the KnowledgeSearchToolName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiSettingsDto) GetKnowledgeSearchToolName() string {
	if o == nil || o.KnowledgeSearchToolName.Get() == nil {
		var ret string
		return ret
	}

	return *o.KnowledgeSearchToolName.Get()
}

// GetKnowledgeSearchToolNameOk returns a tuple with the KnowledgeSearchToolName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiSettingsDto) GetKnowledgeSearchToolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KnowledgeSearchToolName.Get(), o.KnowledgeSearchToolName.IsSet()
}

// SetKnowledgeSearchToolName sets field value
func (o *AiSettingsDto) SetKnowledgeSearchToolName(v string) {
	o.KnowledgeSearchToolName.Set(&v)
}

// GetWebSearchToolName returns the WebSearchToolName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiSettingsDto) GetWebSearchToolName() string {
	if o == nil || o.WebSearchToolName.Get() == nil {
		var ret string
		return ret
	}

	return *o.WebSearchToolName.Get()
}

// GetWebSearchToolNameOk returns a tuple with the WebSearchToolName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiSettingsDto) GetWebSearchToolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebSearchToolName.Get(), o.WebSearchToolName.IsSet()
}

// SetWebSearchToolName sets field value
func (o *AiSettingsDto) SetWebSearchToolName(v string) {
	o.WebSearchToolName.Set(&v)
}

// GetWebCrawlingToolName returns the WebCrawlingToolName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiSettingsDto) GetWebCrawlingToolName() string {
	if o == nil || o.WebCrawlingToolName.Get() == nil {
		var ret string
		return ret
	}

	return *o.WebCrawlingToolName.Get()
}

// GetWebCrawlingToolNameOk returns a tuple with the WebCrawlingToolName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiSettingsDto) GetWebCrawlingToolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebCrawlingToolName.Get(), o.WebCrawlingToolName.IsSet()
}

// SetWebCrawlingToolName sets field value
func (o *AiSettingsDto) SetWebCrawlingToolName(v string) {
	o.WebCrawlingToolName.Set(&v)
}

// GetGenerateDocxToolName returns the GenerateDocxToolName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiSettingsDto) GetGenerateDocxToolName() string {
	if o == nil || o.GenerateDocxToolName.Get() == nil {
		var ret string
		return ret
	}

	return *o.GenerateDocxToolName.Get()
}

// GetGenerateDocxToolNameOk returns a tuple with the GenerateDocxToolName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiSettingsDto) GetGenerateDocxToolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GenerateDocxToolName.Get(), o.GenerateDocxToolName.IsSet()
}

// SetGenerateDocxToolName sets field value
func (o *AiSettingsDto) SetGenerateDocxToolName(v string) {
	o.GenerateDocxToolName.Set(&v)
}

// GetGenerateFormToolName returns the GenerateFormToolName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiSettingsDto) GetGenerateFormToolName() string {
	if o == nil || o.GenerateFormToolName.Get() == nil {
		var ret string
		return ret
	}

	return *o.GenerateFormToolName.Get()
}

// GetGenerateFormToolNameOk returns a tuple with the GenerateFormToolName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiSettingsDto) GetGenerateFormToolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GenerateFormToolName.Get(), o.GenerateFormToolName.IsSet()
}

// SetGenerateFormToolName sets field value
func (o *AiSettingsDto) SetGenerateFormToolName(v string) {
	o.GenerateFormToolName.Set(&v)
}

// GetGeneratePresentationToolName returns the GeneratePresentationToolName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiSettingsDto) GetGeneratePresentationToolName() string {
	if o == nil || o.GeneratePresentationToolName.Get() == nil {
		var ret string
		return ret
	}

	return *o.GeneratePresentationToolName.Get()
}

// GetGeneratePresentationToolNameOk returns a tuple with the GeneratePresentationToolName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiSettingsDto) GetGeneratePresentationToolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GeneratePresentationToolName.Get(), o.GeneratePresentationToolName.IsSet()
}

// SetGeneratePresentationToolName sets field value
func (o *AiSettingsDto) SetGeneratePresentationToolName(v string) {
	o.GeneratePresentationToolName.Set(&v)
}

// GetSystemAiEnabled returns the SystemAiEnabled field value if set, zero value otherwise.
func (o *AiSettingsDto) GetSystemAiEnabled() bool {
	if o == nil || IsNil(o.SystemAiEnabled) {
		var ret bool
		return ret
	}
	return *o.SystemAiEnabled
}

// GetSystemAiEnabledOk returns a tuple with the SystemAiEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiSettingsDto) GetSystemAiEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SystemAiEnabled) {
		return nil, false
	}
	return o.SystemAiEnabled, true
}

// HasSystemAiEnabled returns a boolean if a field has been set.
func (o *AiSettingsDto) IsSystemAiEnabledSet() bool {
	if o != nil && !IsNil(o.SystemAiEnabled) {
		return true
	}

	return false
}

// SetSystemAiEnabled gets a reference to the given bool and assigns it to the SystemAiEnabled field.
func (o *AiSettingsDto) SetSystemAiEnabled(v bool) {
	o.SystemAiEnabled = &v
}

func (o AiSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WebSearchEnabled) {
		toSerialize["webSearchEnabled"] = o.WebSearchEnabled
	}
	if !IsNil(o.WebSearchNeedReset) {
		toSerialize["webSearchNeedReset"] = o.WebSearchNeedReset
	}
	if !IsNil(o.VectorizationEnabled) {
		toSerialize["vectorizationEnabled"] = o.VectorizationEnabled
	}
	if !IsNil(o.VectorizationNeedReset) {
		toSerialize["vectorizationNeedReset"] = o.VectorizationNeedReset
	}
	if !IsNil(o.AiReady) {
		toSerialize["aiReady"] = o.AiReady
	}
	if !IsNil(o.AiReadyNeedReset) {
		toSerialize["aiReadyNeedReset"] = o.AiReadyNeedReset
	}
	if o.PortalMcpServerId.IsSet() {
		toSerialize["portalMcpServerId"] = o.PortalMcpServerId.Get()
	}
	toSerialize["embeddingModel"] = o.EmbeddingModel.Get()
	if o.ModelAliases != nil {
		toSerialize["modelAliases"] = o.ModelAliases
	}
	toSerialize["knowledgeSearchToolName"] = o.KnowledgeSearchToolName.Get()
	toSerialize["webSearchToolName"] = o.WebSearchToolName.Get()
	toSerialize["webCrawlingToolName"] = o.WebCrawlingToolName.Get()
	toSerialize["generateDocxToolName"] = o.GenerateDocxToolName.Get()
	toSerialize["generateFormToolName"] = o.GenerateFormToolName.Get()
	toSerialize["generatePresentationToolName"] = o.GeneratePresentationToolName.Get()
	if !IsNil(o.SystemAiEnabled) {
		toSerialize["systemAiEnabled"] = o.SystemAiEnabled
	}
	return toSerialize, nil
}

func (o *AiSettingsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"embeddingModel",
		"modelAliases",
		"knowledgeSearchToolName",
		"webSearchToolName",
		"webCrawlingToolName",
		"generateDocxToolName",
		"generateFormToolName",
		"generatePresentationToolName",
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

	varAiSettingsDto := _AiSettingsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiSettingsDto)

	if err != nil {
		return err
	}

	*o = AiSettingsDto(varAiSettingsDto)

	return err
}

type NullableAiSettingsDto struct {
	value *AiSettingsDto
	isSet bool
}

func (v NullableAiSettingsDto) Get() *AiSettingsDto {
	return v.value
}

func (v *NullableAiSettingsDto) Set(val *AiSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiSettingsDto(val *AiSettingsDto) *NullableAiSettingsDto {
	return &NullableAiSettingsDto{value: val, isSet: true}
}

func (v NullableAiSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

