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

// checks if the AiProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiProvider{}

// AiProvider Minimal provider connection configuration. Used to connect to a provider API.
type AiProvider struct {
	// Provider type identifier.
	Type AiProviderType `json:"type"`
	// User-defined display name for this provider connection.
	Name string `json:"name"`
	// API key or token. Optional for local providers (Ollama, LM Studio).
	Key *string `json:"key,omitempty"`
	// Base URL of the provider API.
	BaseUrl string `json:"baseUrl"`
}

type _AiProvider AiProvider

// NewAiProvider instantiates a new AiProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiProvider(type_ AiProviderType, name string, baseUrl string) *AiProvider {
	this := AiProvider{}
	this.Type = type_
	this.Name = name
	this.BaseUrl = baseUrl
	return &this
}

// NewAiProviderWithDefaults instantiates a new AiProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiProviderWithDefaults() *AiProvider {
	this := AiProvider{}
	return &this
}

// GetType returns the Type field value
func (o *AiProvider) GetType() AiProviderType {
	if o == nil {
		var ret AiProviderType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiProvider) GetTypeOk() (*AiProviderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AiProvider) SetType(v AiProviderType) {
	o.Type = v
}

// GetName returns the Name field value
func (o *AiProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiProvider) SetName(v string) {
	o.Name = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AiProvider) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProvider) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AiProvider) IsKeySet() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AiProvider) SetKey(v string) {
	o.Key = &v
}

// GetBaseUrl returns the BaseUrl field value
func (o *AiProvider) GetBaseUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value
// and a boolean to check if the value has been set.
func (o *AiProvider) GetBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseUrl, true
}

// SetBaseUrl sets field value
func (o *AiProvider) SetBaseUrl(v string) {
	o.BaseUrl = v
}

func (o AiProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	toSerialize["baseUrl"] = o.BaseUrl
	return toSerialize, nil
}

func (o *AiProvider) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"name",
		"baseUrl",
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

	varAiProvider := _AiProvider{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiProvider)

	if err != nil {
		return err
	}

	*o = AiProvider(varAiProvider)

	return err
}

type NullableAiProvider struct {
	value *AiProvider
	isSet bool
}

func (v NullableAiProvider) Get() *AiProvider {
	return v.value
}

func (v *NullableAiProvider) Set(val *AiProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableAiProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableAiProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiProvider(val *AiProvider) *NullableAiProvider {
	return &NullableAiProvider{value: val, isSet: true}
}

func (v NullableAiProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

