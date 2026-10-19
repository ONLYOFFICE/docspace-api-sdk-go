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
	// Indicates whether document vectorization is enabled.
	VectorizationEnabled *bool `json:"vectorizationEnabled,omitempty"`
	// Indicates whether the embedding provider API key needs to be reconfigured.
	VectorizationNeedReset *bool `json:"vectorizationNeedReset,omitempty"`
	// Indicates whether the AI subsystem is fully configured and operational.
	AiReady *bool `json:"aiReady,omitempty"`
	// The name of the embedding model used for document vectorization.
	EmbeddingModel NullableString `json:"embeddingModel"`
	// Indicates whether the system-level AI provider is enabled.
	SystemAiEnabled *bool `json:"systemAiEnabled,omitempty"`
	// The identifier of the model recommended for form generation.
	RecommendedModelForForms NullableString `json:"recommendedModelForForms,omitempty"`
	// How tool calls made by the model are approved for the current user. The default applies while the user has stored nothing.
	ToolPermissionMode *AiToolPermissionMode `json:"toolPermissionMode,omitempty"`
}

type _AiSettingsDto AiSettingsDto

// NewAiSettingsDto instantiates a new AiSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiSettingsDto(embeddingModel NullableString) *AiSettingsDto {
	this := AiSettingsDto{}
	this.EmbeddingModel = embeddingModel
	return &this
}

// NewAiSettingsDtoWithDefaults instantiates a new AiSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiSettingsDtoWithDefaults() *AiSettingsDto {
	this := AiSettingsDto{}
	return &this
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

// GetRecommendedModelForForms returns the RecommendedModelForForms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiSettingsDto) GetRecommendedModelForForms() string {
	if o == nil || IsNil(o.RecommendedModelForForms.Get()) {
		var ret string
		return ret
	}
	return *o.RecommendedModelForForms.Get()
}

// GetRecommendedModelForFormsOk returns a tuple with the RecommendedModelForForms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiSettingsDto) GetRecommendedModelForFormsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecommendedModelForForms.Get(), o.RecommendedModelForForms.IsSet()
}

// HasRecommendedModelForForms returns a boolean if a field has been set.
func (o *AiSettingsDto) IsRecommendedModelForFormsSet() bool {
	if o != nil && o.RecommendedModelForForms.IsSet() {
		return true
	}

	return false
}

// SetRecommendedModelForForms gets a reference to the given NullableString and assigns it to the RecommendedModelForForms field.
func (o *AiSettingsDto) SetRecommendedModelForForms(v string) {
	o.RecommendedModelForForms.Set(&v)
}
// SetRecommendedModelForFormsNil sets the value for RecommendedModelForForms to be an explicit nil
func (o *AiSettingsDto) SetRecommendedModelForFormsNil() {
	o.RecommendedModelForForms.Set(nil)
}

// UnsetRecommendedModelForForms ensures that no value is present for RecommendedModelForForms, not even an explicit nil
func (o *AiSettingsDto) UnsetRecommendedModelForForms() {
	o.RecommendedModelForForms.Unset()
}

// GetToolPermissionMode returns the ToolPermissionMode field value if set, zero value otherwise.
func (o *AiSettingsDto) GetToolPermissionMode() AiToolPermissionMode {
	if o == nil || IsNil(o.ToolPermissionMode) {
		var ret AiToolPermissionMode
		return ret
	}
	return *o.ToolPermissionMode
}

// GetToolPermissionModeOk returns a tuple with the ToolPermissionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiSettingsDto) GetToolPermissionModeOk() (*AiToolPermissionMode, bool) {
	if o == nil || IsNil(o.ToolPermissionMode) {
		return nil, false
	}
	return o.ToolPermissionMode, true
}

// HasToolPermissionMode returns a boolean if a field has been set.
func (o *AiSettingsDto) IsToolPermissionModeSet() bool {
	if o != nil && !IsNil(o.ToolPermissionMode) {
		return true
	}

	return false
}

// SetToolPermissionMode gets a reference to the given AiToolPermissionMode and assigns it to the ToolPermissionMode field.
func (o *AiSettingsDto) SetToolPermissionMode(v AiToolPermissionMode) {
	o.ToolPermissionMode = &v
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
	if !IsNil(o.VectorizationEnabled) {
		toSerialize["vectorizationEnabled"] = o.VectorizationEnabled
	}
	if !IsNil(o.VectorizationNeedReset) {
		toSerialize["vectorizationNeedReset"] = o.VectorizationNeedReset
	}
	if !IsNil(o.AiReady) {
		toSerialize["aiReady"] = o.AiReady
	}
	toSerialize["embeddingModel"] = o.EmbeddingModel.Get()
	if !IsNil(o.SystemAiEnabled) {
		toSerialize["systemAiEnabled"] = o.SystemAiEnabled
	}
	if o.RecommendedModelForForms.IsSet() {
		toSerialize["recommendedModelForForms"] = o.RecommendedModelForForms.Get()
	}
	if !IsNil(o.ToolPermissionMode) {
		toSerialize["toolPermissionMode"] = o.ToolPermissionMode
	}
	return toSerialize, nil
}

func (o *AiSettingsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"embeddingModel",
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

