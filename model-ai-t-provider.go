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

// checks if the AiTProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiTProvider{}

// AiTProvider Minimal provider connection configuration. Used to connect to a provider API.
type AiTProvider struct {
	// Provider type identifier.
	Type AiProviderType `json:"type"`
	// User-defined display name for this provider connection.
	Name string `json:"name"`
	// API key or token. Optional for local providers (Ollama, LM Studio).
	Key *string `json:"key,omitempty"`
	// Base URL of the provider API.
	BaseUrl string `json:"baseUrl"`
}

type _AiTProvider AiTProvider

// NewAiTProvider instantiates a new AiTProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiTProvider(type_ AiProviderType, name string, baseUrl string) *AiTProvider {
	this := AiTProvider{}
	this.Type = type_
	this.Name = name
	this.BaseUrl = baseUrl
	return &this
}

// NewAiTProviderWithDefaults instantiates a new AiTProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiTProviderWithDefaults() *AiTProvider {
	this := AiTProvider{}
	return &this
}

// GetType returns the Type field value
func (o *AiTProvider) GetType() AiProviderType {
	if o == nil {
		var ret AiProviderType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiTProvider) GetTypeOk() (*AiProviderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AiTProvider) SetType(v AiProviderType) {
	o.Type = v
}

// GetName returns the Name field value
func (o *AiTProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiTProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiTProvider) SetName(v string) {
	o.Name = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AiTProvider) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiTProvider) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AiTProvider) IsKeySet() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AiTProvider) SetKey(v string) {
	o.Key = &v
}

// GetBaseUrl returns the BaseUrl field value
func (o *AiTProvider) GetBaseUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value
// and a boolean to check if the value has been set.
func (o *AiTProvider) GetBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseUrl, true
}

// SetBaseUrl sets field value
func (o *AiTProvider) SetBaseUrl(v string) {
	o.BaseUrl = v
}

func (o AiTProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiTProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	toSerialize["baseUrl"] = o.BaseUrl
	return toSerialize, nil
}

func (o *AiTProvider) UnmarshalJSON(data []byte) (err error) {
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

	varAiTProvider := _AiTProvider{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiTProvider)

	if err != nil {
		return err
	}

	*o = AiTProvider(varAiTProvider)

	return err
}

type NullableAiTProvider struct {
	value *AiTProvider
	isSet bool
}

func (v NullableAiTProvider) Get() *AiTProvider {
	return v.value
}

func (v *NullableAiTProvider) Set(val *AiTProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableAiTProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableAiTProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiTProvider(val *AiTProvider) *NullableAiTProvider {
	return &NullableAiTProvider{value: val, isSet: true}
}

func (v NullableAiTProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiTProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

