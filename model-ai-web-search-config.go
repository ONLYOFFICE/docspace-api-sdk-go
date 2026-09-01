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

// checks if the AiWebSearchConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiWebSearchConfig{}

// AiWebSearchConfig Web-search provider configuration. Credentials and provider selection for the built-in web-search tool group.
type AiWebSearchConfig struct {
	// Provider identifier (e.g. `exa`).
	Provider string `json:"provider"`
	// API key for the provider. Optional for self-hosted or keyless setups.
	Key *string `json:"key,omitempty"`
	// Optional override for the provider's base URL.
	BaseUrl *string `json:"baseUrl,omitempty"`
	// Whether this provider is cloud-hosted (vs. self-hosted).
	IsCloudProvider *bool `json:"isCloudProvider,omitempty"`
	// Extra HTTP headers sent with each request to the ONLYOFFICE / cloud backend (e.g. `X-Tenant`). Merged after the derived `Authorization` header, so a custom header of the same name wins.
	Headers map[string]string `json:"headers,omitempty"`
}

type _AiWebSearchConfig AiWebSearchConfig

// NewAiWebSearchConfig instantiates a new AiWebSearchConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiWebSearchConfig(provider string) *AiWebSearchConfig {
	this := AiWebSearchConfig{}
	this.Provider = provider
	return &this
}

// NewAiWebSearchConfigWithDefaults instantiates a new AiWebSearchConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiWebSearchConfigWithDefaults() *AiWebSearchConfig {
	this := AiWebSearchConfig{}
	return &this
}

// GetProvider returns the Provider field value
func (o *AiWebSearchConfig) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *AiWebSearchConfig) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *AiWebSearchConfig) SetProvider(v string) {
	o.Provider = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AiWebSearchConfig) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiWebSearchConfig) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AiWebSearchConfig) IsKeySet() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AiWebSearchConfig) SetKey(v string) {
	o.Key = &v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *AiWebSearchConfig) GetBaseUrl() string {
	if o == nil || IsNil(o.BaseUrl) {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiWebSearchConfig) GetBaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BaseUrl) {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *AiWebSearchConfig) IsBaseUrlSet() bool {
	if o != nil && !IsNil(o.BaseUrl) {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *AiWebSearchConfig) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetIsCloudProvider returns the IsCloudProvider field value if set, zero value otherwise.
func (o *AiWebSearchConfig) GetIsCloudProvider() bool {
	if o == nil || IsNil(o.IsCloudProvider) {
		var ret bool
		return ret
	}
	return *o.IsCloudProvider
}

// GetIsCloudProviderOk returns a tuple with the IsCloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiWebSearchConfig) GetIsCloudProviderOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCloudProvider) {
		return nil, false
	}
	return o.IsCloudProvider, true
}

// HasIsCloudProvider returns a boolean if a field has been set.
func (o *AiWebSearchConfig) IsIsCloudProviderSet() bool {
	if o != nil && !IsNil(o.IsCloudProvider) {
		return true
	}

	return false
}

// SetIsCloudProvider gets a reference to the given bool and assigns it to the IsCloudProvider field.
func (o *AiWebSearchConfig) SetIsCloudProvider(v bool) {
	o.IsCloudProvider = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *AiWebSearchConfig) GetHeaders() map[string]string {
	if o == nil || IsNil(o.Headers) {
		var ret map[string]string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiWebSearchConfig) GetHeadersOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return map[string]string{}, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *AiWebSearchConfig) IsHeadersSet() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *AiWebSearchConfig) SetHeaders(v map[string]string) {
	o.Headers = v
}

func (o AiWebSearchConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiWebSearchConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["provider"] = o.Provider
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.BaseUrl) {
		toSerialize["baseUrl"] = o.BaseUrl
	}
	if !IsNil(o.IsCloudProvider) {
		toSerialize["isCloudProvider"] = o.IsCloudProvider
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	return toSerialize, nil
}

func (o *AiWebSearchConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varAiWebSearchConfig := _AiWebSearchConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiWebSearchConfig)

	if err != nil {
		return err
	}

	*o = AiWebSearchConfig(varAiWebSearchConfig)

	return err
}

type NullableAiWebSearchConfig struct {
	value *AiWebSearchConfig
	isSet bool
}

func (v NullableAiWebSearchConfig) Get() *AiWebSearchConfig {
	return v.value
}

func (v *NullableAiWebSearchConfig) Set(val *AiWebSearchConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAiWebSearchConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAiWebSearchConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiWebSearchConfig(val *AiWebSearchConfig) *NullableAiWebSearchConfig {
	return &NullableAiWebSearchConfig{value: val, isSet: true}
}

func (v NullableAiWebSearchConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiWebSearchConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

