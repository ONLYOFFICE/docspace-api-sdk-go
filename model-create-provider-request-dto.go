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

// checks if the CreateProviderRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProviderRequestDto{}

// CreateProviderRequestDto Request parameters for creating a new AI provider.
type CreateProviderRequestDto struct {
	Type *ProviderType `json:"type,omitempty"`
	// The display title for the AI provider.
	Title NullableString `json:"title"`
	// The API endpoint URL for the AI provider. Required for OpenAiCompatible type; optional for other types that have default URLs.
	Url NullableString `json:"url,omitempty"`
	// The authentication API key for the AI provider.
	Key NullableString `json:"key"`
	// Optional list of model settings to configure atomically with the provider creation.
	ModelSettings []ModelSettingsItemDto `json:"modelSettings,omitempty"`
}

type _CreateProviderRequestDto CreateProviderRequestDto

// NewCreateProviderRequestDto instantiates a new CreateProviderRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProviderRequestDto(title NullableString, key NullableString) *CreateProviderRequestDto {
	this := CreateProviderRequestDto{}
	this.Title = title
	this.Key = key
	return &this
}

// NewCreateProviderRequestDtoWithDefaults instantiates a new CreateProviderRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProviderRequestDtoWithDefaults() *CreateProviderRequestDto {
	this := CreateProviderRequestDto{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateProviderRequestDto) GetType() ProviderType {
	if o == nil || IsNil(o.Type) {
		var ret ProviderType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProviderRequestDto) GetTypeOk() (*ProviderType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateProviderRequestDto) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ProviderType and assigns it to the Type field.
func (o *CreateProviderRequestDto) SetType(v ProviderType) {
	o.Type = &v
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateProviderRequestDto) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProviderRequestDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *CreateProviderRequestDto) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProviderRequestDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProviderRequestDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *CreateProviderRequestDto) IsUrlSet() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *CreateProviderRequestDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *CreateProviderRequestDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *CreateProviderRequestDto) UnsetUrl() {
	o.Url.Unset()
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateProviderRequestDto) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}

	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProviderRequestDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// SetKey sets field value
func (o *CreateProviderRequestDto) SetKey(v string) {
	o.Key.Set(&v)
}

// GetModelSettings returns the ModelSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProviderRequestDto) GetModelSettings() []ModelSettingsItemDto {
	if o == nil {
		var ret []ModelSettingsItemDto
		return ret
	}
	return o.ModelSettings
}

// GetModelSettingsOk returns a tuple with the ModelSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProviderRequestDto) GetModelSettingsOk() ([]ModelSettingsItemDto, bool) {
	if o == nil || IsNil(o.ModelSettings) {
		return nil, false
	}
	return o.ModelSettings, true
}

// HasModelSettings returns a boolean if a field has been set.
func (o *CreateProviderRequestDto) IsModelSettingsSet() bool {
	if o != nil && !IsNil(o.ModelSettings) {
		return true
	}

	return false
}

// SetModelSettings gets a reference to the given []ModelSettingsItemDto and assigns it to the ModelSettings field.
func (o *CreateProviderRequestDto) SetModelSettings(v []ModelSettingsItemDto) {
	o.ModelSettings = v
}

func (o CreateProviderRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProviderRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["title"] = o.Title.Get()
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	toSerialize["key"] = o.Key.Get()
	if o.ModelSettings != nil {
		toSerialize["modelSettings"] = o.ModelSettings
	}
	return toSerialize, nil
}

func (o *CreateProviderRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"key",
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

	varCreateProviderRequestDto := _CreateProviderRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCreateProviderRequestDto)

	if err != nil {
		return err
	}

	*o = CreateProviderRequestDto(varCreateProviderRequestDto)

	return err
}

type NullableCreateProviderRequestDto struct {
	value *CreateProviderRequestDto
	isSet bool
}

func (v NullableCreateProviderRequestDto) Get() *CreateProviderRequestDto {
	return v.value
}

func (v *NullableCreateProviderRequestDto) Set(val *CreateProviderRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProviderRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProviderRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProviderRequestDto(val *CreateProviderRequestDto) *NullableCreateProviderRequestDto {
	return &NullableCreateProviderRequestDto{value: val, isSet: true}
}

func (v NullableCreateProviderRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProviderRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

