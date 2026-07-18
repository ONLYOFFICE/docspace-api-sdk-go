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

// checks if the ModelSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSettingsDto{}

// ModelSettingsDto AI model settings information.
type ModelSettingsDto struct {
	// The model identifier.
	Id NullableString `json:"id"`
	// The display name for the model.
	Alias NullableString `json:"alias,omitempty"`
	// Whether the model is enabled for use in chat.
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Whether the model is recommended (defined in configuration).
	IsRecommended *bool `json:"isRecommended,omitempty"`
	Capabilities AiModelCapabilities `json:"capabilities"`
}

type _ModelSettingsDto ModelSettingsDto

// NewModelSettingsDto instantiates a new ModelSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSettingsDto(id NullableString, capabilities AiModelCapabilities) *ModelSettingsDto {
	this := ModelSettingsDto{}
	this.Id = id
	this.Capabilities = capabilities
	return &this
}

// NewModelSettingsDtoWithDefaults instantiates a new ModelSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSettingsDtoWithDefaults() *ModelSettingsDto {
	this := ModelSettingsDto{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ModelSettingsDto) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSettingsDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *ModelSettingsDto) SetId(v string) {
	o.Id.Set(&v)
}

// GetAlias returns the Alias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelSettingsDto) GetAlias() string {
	if o == nil || IsNil(o.Alias.Get()) {
		var ret string
		return ret
	}
	return *o.Alias.Get()
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSettingsDto) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alias.Get(), o.Alias.IsSet()
}

// HasAlias returns a boolean if a field has been set.
func (o *ModelSettingsDto) IsAliasSet() bool {
	if o != nil && o.Alias.IsSet() {
		return true
	}

	return false
}

// SetAlias gets a reference to the given NullableString and assigns it to the Alias field.
func (o *ModelSettingsDto) SetAlias(v string) {
	o.Alias.Set(&v)
}
// SetAliasNil sets the value for Alias to be an explicit nil
func (o *ModelSettingsDto) SetAliasNil() {
	o.Alias.Set(nil)
}

// UnsetAlias ensures that no value is present for Alias, not even an explicit nil
func (o *ModelSettingsDto) UnsetAlias() {
	o.Alias.Unset()
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *ModelSettingsDto) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSettingsDto) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *ModelSettingsDto) IsIsEnabledSet() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *ModelSettingsDto) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetIsRecommended returns the IsRecommended field value if set, zero value otherwise.
func (o *ModelSettingsDto) GetIsRecommended() bool {
	if o == nil || IsNil(o.IsRecommended) {
		var ret bool
		return ret
	}
	return *o.IsRecommended
}

// GetIsRecommendedOk returns a tuple with the IsRecommended field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSettingsDto) GetIsRecommendedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRecommended) {
		return nil, false
	}
	return o.IsRecommended, true
}

// HasIsRecommended returns a boolean if a field has been set.
func (o *ModelSettingsDto) IsIsRecommendedSet() bool {
	if o != nil && !IsNil(o.IsRecommended) {
		return true
	}

	return false
}

// SetIsRecommended gets a reference to the given bool and assigns it to the IsRecommended field.
func (o *ModelSettingsDto) SetIsRecommended(v bool) {
	o.IsRecommended = &v
}

// GetCapabilities returns the Capabilities field value
func (o *ModelSettingsDto) GetCapabilities() AiModelCapabilities {
	if o == nil {
		var ret AiModelCapabilities
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *ModelSettingsDto) GetCapabilitiesOk() (*AiModelCapabilities, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capabilities, true
}

// SetCapabilities sets field value
func (o *ModelSettingsDto) SetCapabilities(v AiModelCapabilities) {
	o.Capabilities = v
}

func (o ModelSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	if o.Alias.IsSet() {
		toSerialize["alias"] = o.Alias.Get()
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !IsNil(o.IsRecommended) {
		toSerialize["isRecommended"] = o.IsRecommended
	}
	toSerialize["capabilities"] = o.Capabilities
	return toSerialize, nil
}

func (o *ModelSettingsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"capabilities",
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

	varModelSettingsDto := _ModelSettingsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varModelSettingsDto)

	if err != nil {
		return err
	}

	*o = ModelSettingsDto(varModelSettingsDto)

	return err
}

type NullableModelSettingsDto struct {
	value *ModelSettingsDto
	isSet bool
}

func (v NullableModelSettingsDto) Get() *ModelSettingsDto {
	return v.value
}

func (v *NullableModelSettingsDto) Set(val *ModelSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSettingsDto(val *ModelSettingsDto) *NullableModelSettingsDto {
	return &NullableModelSettingsDto{value: val, isSet: true}
}

func (v NullableModelSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

