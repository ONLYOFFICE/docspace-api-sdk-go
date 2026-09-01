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

// checks if the AiAiSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiAiSettingsDto{}

// AiAiSettingsDto The AI module settings.
type AiAiSettingsDto struct {
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
}

type _AiAiSettingsDto AiAiSettingsDto

// NewAiAiSettingsDto instantiates a new AiAiSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAiSettingsDto(embeddingModel NullableString) *AiAiSettingsDto {
	this := AiAiSettingsDto{}
	this.EmbeddingModel = embeddingModel
	return &this
}

// NewAiAiSettingsDtoWithDefaults instantiates a new AiAiSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAiSettingsDtoWithDefaults() *AiAiSettingsDto {
	this := AiAiSettingsDto{}
	return &this
}

// GetVectorizationEnabled returns the VectorizationEnabled field value if set, zero value otherwise.
func (o *AiAiSettingsDto) GetVectorizationEnabled() bool {
	if o == nil || IsNil(o.VectorizationEnabled) {
		var ret bool
		return ret
	}
	return *o.VectorizationEnabled
}

// GetVectorizationEnabledOk returns a tuple with the VectorizationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAiSettingsDto) GetVectorizationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.VectorizationEnabled) {
		return nil, false
	}
	return o.VectorizationEnabled, true
}

// HasVectorizationEnabled returns a boolean if a field has been set.
func (o *AiAiSettingsDto) IsVectorizationEnabledSet() bool {
	if o != nil && !IsNil(o.VectorizationEnabled) {
		return true
	}

	return false
}

// SetVectorizationEnabled gets a reference to the given bool and assigns it to the VectorizationEnabled field.
func (o *AiAiSettingsDto) SetVectorizationEnabled(v bool) {
	o.VectorizationEnabled = &v
}

// GetVectorizationNeedReset returns the VectorizationNeedReset field value if set, zero value otherwise.
func (o *AiAiSettingsDto) GetVectorizationNeedReset() bool {
	if o == nil || IsNil(o.VectorizationNeedReset) {
		var ret bool
		return ret
	}
	return *o.VectorizationNeedReset
}

// GetVectorizationNeedResetOk returns a tuple with the VectorizationNeedReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAiSettingsDto) GetVectorizationNeedResetOk() (*bool, bool) {
	if o == nil || IsNil(o.VectorizationNeedReset) {
		return nil, false
	}
	return o.VectorizationNeedReset, true
}

// HasVectorizationNeedReset returns a boolean if a field has been set.
func (o *AiAiSettingsDto) IsVectorizationNeedResetSet() bool {
	if o != nil && !IsNil(o.VectorizationNeedReset) {
		return true
	}

	return false
}

// SetVectorizationNeedReset gets a reference to the given bool and assigns it to the VectorizationNeedReset field.
func (o *AiAiSettingsDto) SetVectorizationNeedReset(v bool) {
	o.VectorizationNeedReset = &v
}

// GetAiReady returns the AiReady field value if set, zero value otherwise.
func (o *AiAiSettingsDto) GetAiReady() bool {
	if o == nil || IsNil(o.AiReady) {
		var ret bool
		return ret
	}
	return *o.AiReady
}

// GetAiReadyOk returns a tuple with the AiReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAiSettingsDto) GetAiReadyOk() (*bool, bool) {
	if o == nil || IsNil(o.AiReady) {
		return nil, false
	}
	return o.AiReady, true
}

// HasAiReady returns a boolean if a field has been set.
func (o *AiAiSettingsDto) IsAiReadySet() bool {
	if o != nil && !IsNil(o.AiReady) {
		return true
	}

	return false
}

// SetAiReady gets a reference to the given bool and assigns it to the AiReady field.
func (o *AiAiSettingsDto) SetAiReady(v bool) {
	o.AiReady = &v
}

// GetEmbeddingModel returns the EmbeddingModel field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiAiSettingsDto) GetEmbeddingModel() string {
	if o == nil || o.EmbeddingModel.Get() == nil {
		var ret string
		return ret
	}

	return *o.EmbeddingModel.Get()
}

// GetEmbeddingModelOk returns a tuple with the EmbeddingModel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiAiSettingsDto) GetEmbeddingModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmbeddingModel.Get(), o.EmbeddingModel.IsSet()
}

// SetEmbeddingModel sets field value
func (o *AiAiSettingsDto) SetEmbeddingModel(v string) {
	o.EmbeddingModel.Set(&v)
}

// GetSystemAiEnabled returns the SystemAiEnabled field value if set, zero value otherwise.
func (o *AiAiSettingsDto) GetSystemAiEnabled() bool {
	if o == nil || IsNil(o.SystemAiEnabled) {
		var ret bool
		return ret
	}
	return *o.SystemAiEnabled
}

// GetSystemAiEnabledOk returns a tuple with the SystemAiEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAiSettingsDto) GetSystemAiEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SystemAiEnabled) {
		return nil, false
	}
	return o.SystemAiEnabled, true
}

// HasSystemAiEnabled returns a boolean if a field has been set.
func (o *AiAiSettingsDto) IsSystemAiEnabledSet() bool {
	if o != nil && !IsNil(o.SystemAiEnabled) {
		return true
	}

	return false
}

// SetSystemAiEnabled gets a reference to the given bool and assigns it to the SystemAiEnabled field.
func (o *AiAiSettingsDto) SetSystemAiEnabled(v bool) {
	o.SystemAiEnabled = &v
}

// GetRecommendedModelForForms returns the RecommendedModelForForms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiAiSettingsDto) GetRecommendedModelForForms() string {
	if o == nil || IsNil(o.RecommendedModelForForms.Get()) {
		var ret string
		return ret
	}
	return *o.RecommendedModelForForms.Get()
}

// GetRecommendedModelForFormsOk returns a tuple with the RecommendedModelForForms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiAiSettingsDto) GetRecommendedModelForFormsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecommendedModelForForms.Get(), o.RecommendedModelForForms.IsSet()
}

// HasRecommendedModelForForms returns a boolean if a field has been set.
func (o *AiAiSettingsDto) IsRecommendedModelForFormsSet() bool {
	if o != nil && o.RecommendedModelForForms.IsSet() {
		return true
	}

	return false
}

// SetRecommendedModelForForms gets a reference to the given NullableString and assigns it to the RecommendedModelForForms field.
func (o *AiAiSettingsDto) SetRecommendedModelForForms(v string) {
	o.RecommendedModelForForms.Set(&v)
}
// SetRecommendedModelForFormsNil sets the value for RecommendedModelForForms to be an explicit nil
func (o *AiAiSettingsDto) SetRecommendedModelForFormsNil() {
	o.RecommendedModelForForms.Set(nil)
}

// UnsetRecommendedModelForForms ensures that no value is present for RecommendedModelForForms, not even an explicit nil
func (o *AiAiSettingsDto) UnsetRecommendedModelForForms() {
	o.RecommendedModelForForms.Unset()
}

func (o AiAiSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiAiSettingsDto) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

func (o *AiAiSettingsDto) UnmarshalJSON(data []byte) (err error) {
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

	varAiAiSettingsDto := _AiAiSettingsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiAiSettingsDto)

	if err != nil {
		return err
	}

	*o = AiAiSettingsDto(varAiAiSettingsDto)

	return err
}

type NullableAiAiSettingsDto struct {
	value *AiAiSettingsDto
	isSet bool
}

func (v NullableAiAiSettingsDto) Get() *AiAiSettingsDto {
	return v.value
}

func (v *NullableAiAiSettingsDto) Set(val *AiAiSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAiSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAiSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAiSettingsDto(val *AiAiSettingsDto) *NullableAiAiSettingsDto {
	return &NullableAiAiSettingsDto{value: val, isSet: true}
}

func (v NullableAiAiSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAiSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

