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

// checks if the PreviewProviderModelsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreviewProviderModelsRequestDto{}

// PreviewProviderModelsRequestDto Request parameters for previewing models available from a provider before saving it.
type PreviewProviderModelsRequestDto struct {
	Type *ProviderType `json:"type,omitempty"`
	// The API endpoint URL. Required for OpenAiCompatible type; optional for other types that have default URLs.
	Url NullableString `json:"url,omitempty"`
	// The authentication API key for the AI provider.
	Key NullableString `json:"key"`
}

type _PreviewProviderModelsRequestDto PreviewProviderModelsRequestDto

// NewPreviewProviderModelsRequestDto instantiates a new PreviewProviderModelsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreviewProviderModelsRequestDto(key NullableString) *PreviewProviderModelsRequestDto {
	this := PreviewProviderModelsRequestDto{}
	this.Key = key
	return &this
}

// NewPreviewProviderModelsRequestDtoWithDefaults instantiates a new PreviewProviderModelsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreviewProviderModelsRequestDtoWithDefaults() *PreviewProviderModelsRequestDto {
	this := PreviewProviderModelsRequestDto{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PreviewProviderModelsRequestDto) GetType() ProviderType {
	if o == nil || IsNil(o.Type) {
		var ret ProviderType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreviewProviderModelsRequestDto) GetTypeOk() (*ProviderType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PreviewProviderModelsRequestDto) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ProviderType and assigns it to the Type field.
func (o *PreviewProviderModelsRequestDto) SetType(v ProviderType) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PreviewProviderModelsRequestDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PreviewProviderModelsRequestDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *PreviewProviderModelsRequestDto) IsUrlSet() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *PreviewProviderModelsRequestDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *PreviewProviderModelsRequestDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *PreviewProviderModelsRequestDto) UnsetUrl() {
	o.Url.Unset()
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PreviewProviderModelsRequestDto) GetKey() string {
	if o == nil || o.Key.Get() == nil {
		var ret string
		return ret
	}

	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PreviewProviderModelsRequestDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// SetKey sets field value
func (o *PreviewProviderModelsRequestDto) SetKey(v string) {
	o.Key.Set(&v)
}

func (o PreviewProviderModelsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreviewProviderModelsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	toSerialize["key"] = o.Key.Get()
	return toSerialize, nil
}

func (o *PreviewProviderModelsRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varPreviewProviderModelsRequestDto := _PreviewProviderModelsRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varPreviewProviderModelsRequestDto)

	if err != nil {
		return err
	}

	*o = PreviewProviderModelsRequestDto(varPreviewProviderModelsRequestDto)

	return err
}

type NullablePreviewProviderModelsRequestDto struct {
	value *PreviewProviderModelsRequestDto
	isSet bool
}

func (v NullablePreviewProviderModelsRequestDto) Get() *PreviewProviderModelsRequestDto {
	return v.value
}

func (v *NullablePreviewProviderModelsRequestDto) Set(val *PreviewProviderModelsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePreviewProviderModelsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePreviewProviderModelsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreviewProviderModelsRequestDto(val *PreviewProviderModelsRequestDto) *NullablePreviewProviderModelsRequestDto {
	return &NullablePreviewProviderModelsRequestDto{value: val, isSet: true}
}

func (v NullablePreviewProviderModelsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreviewProviderModelsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

