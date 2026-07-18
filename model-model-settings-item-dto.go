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

// checks if the ModelSettingsItemDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSettingsItemDto{}

// ModelSettingsItemDto A single model settings entry within a provider create or update request.
type ModelSettingsItemDto struct {
	// The model identifier.
	ModelId string `json:"modelId"`
	// Whether the model is enabled for use in chat.
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// The display name for the model. Only applies to non-recommended models.
	Alias NullableString `json:"alias,omitempty"`
	Capabilities *AiModelCapabilities `json:"capabilities,omitempty"`
}

type _ModelSettingsItemDto ModelSettingsItemDto

// NewModelSettingsItemDto instantiates a new ModelSettingsItemDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSettingsItemDto(modelId string) *ModelSettingsItemDto {
	this := ModelSettingsItemDto{}
	this.ModelId = modelId
	return &this
}

// NewModelSettingsItemDtoWithDefaults instantiates a new ModelSettingsItemDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSettingsItemDtoWithDefaults() *ModelSettingsItemDto {
	this := ModelSettingsItemDto{}
	return &this
}

// GetModelId returns the ModelId field value
func (o *ModelSettingsItemDto) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *ModelSettingsItemDto) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value
func (o *ModelSettingsItemDto) SetModelId(v string) {
	o.ModelId = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *ModelSettingsItemDto) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSettingsItemDto) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *ModelSettingsItemDto) IsIsEnabledSet() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *ModelSettingsItemDto) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelSettingsItemDto) GetAlias() string {
	if o == nil || IsNil(o.Alias.Get()) {
		var ret string
		return ret
	}
	return *o.Alias.Get()
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSettingsItemDto) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alias.Get(), o.Alias.IsSet()
}

// HasAlias returns a boolean if a field has been set.
func (o *ModelSettingsItemDto) IsAliasSet() bool {
	if o != nil && o.Alias.IsSet() {
		return true
	}

	return false
}

// SetAlias gets a reference to the given NullableString and assigns it to the Alias field.
func (o *ModelSettingsItemDto) SetAlias(v string) {
	o.Alias.Set(&v)
}
// SetAliasNil sets the value for Alias to be an explicit nil
func (o *ModelSettingsItemDto) SetAliasNil() {
	o.Alias.Set(nil)
}

// UnsetAlias ensures that no value is present for Alias, not even an explicit nil
func (o *ModelSettingsItemDto) UnsetAlias() {
	o.Alias.Unset()
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *ModelSettingsItemDto) GetCapabilities() AiModelCapabilities {
	if o == nil || IsNil(o.Capabilities) {
		var ret AiModelCapabilities
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSettingsItemDto) GetCapabilitiesOk() (*AiModelCapabilities, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *ModelSettingsItemDto) IsCapabilitiesSet() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given AiModelCapabilities and assigns it to the Capabilities field.
func (o *ModelSettingsItemDto) SetCapabilities(v AiModelCapabilities) {
	o.Capabilities = &v
}

func (o ModelSettingsItemDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSettingsItemDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["modelId"] = o.ModelId
	if !IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.Alias.IsSet() {
		toSerialize["alias"] = o.Alias.Get()
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	return toSerialize, nil
}

func (o *ModelSettingsItemDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"modelId",
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

	varModelSettingsItemDto := _ModelSettingsItemDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varModelSettingsItemDto)

	if err != nil {
		return err
	}

	*o = ModelSettingsItemDto(varModelSettingsItemDto)

	return err
}

type NullableModelSettingsItemDto struct {
	value *ModelSettingsItemDto
	isSet bool
}

func (v NullableModelSettingsItemDto) Get() *ModelSettingsItemDto {
	return v.value
}

func (v *NullableModelSettingsItemDto) Set(val *ModelSettingsItemDto) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSettingsItemDto) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSettingsItemDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSettingsItemDto(val *ModelSettingsItemDto) *NullableModelSettingsItemDto {
	return &NullableModelSettingsItemDto{value: val, isSet: true}
}

func (v NullableModelSettingsItemDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSettingsItemDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

