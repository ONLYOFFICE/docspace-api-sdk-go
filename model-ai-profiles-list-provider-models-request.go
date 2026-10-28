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

// checks if the AiProfilesListProviderModelsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiProfilesListProviderModelsRequest{}

// AiProfilesListProviderModelsRequest struct for AiProfilesListProviderModelsRequest
type AiProfilesListProviderModelsRequest struct {
	// Provider whose catalog to list.
	ProviderType AiProviderType `json:"providerType"`
	// Provider API base URL.
	BaseUrl string `json:"baseUrl"`
	// Provider API key. Omit it for a provider that needs none; the request is then made without one.
	ApiKey *string `json:"apiKey,omitempty"`
}

type _AiProfilesListProviderModelsRequest AiProfilesListProviderModelsRequest

// NewAiProfilesListProviderModelsRequest instantiates a new AiProfilesListProviderModelsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiProfilesListProviderModelsRequest(providerType AiProviderType, baseUrl string) *AiProfilesListProviderModelsRequest {
	this := AiProfilesListProviderModelsRequest{}
	this.ProviderType = providerType
	this.BaseUrl = baseUrl
	return &this
}

// NewAiProfilesListProviderModelsRequestWithDefaults instantiates a new AiProfilesListProviderModelsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiProfilesListProviderModelsRequestWithDefaults() *AiProfilesListProviderModelsRequest {
	this := AiProfilesListProviderModelsRequest{}
	return &this
}

// GetProviderType returns the ProviderType field value
func (o *AiProfilesListProviderModelsRequest) GetProviderType() AiProviderType {
	if o == nil {
		var ret AiProviderType
		return ret
	}

	return o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value
// and a boolean to check if the value has been set.
func (o *AiProfilesListProviderModelsRequest) GetProviderTypeOk() (*AiProviderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderType, true
}

// SetProviderType sets field value
func (o *AiProfilesListProviderModelsRequest) SetProviderType(v AiProviderType) {
	o.ProviderType = v
}

// GetBaseUrl returns the BaseUrl field value
func (o *AiProfilesListProviderModelsRequest) GetBaseUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value
// and a boolean to check if the value has been set.
func (o *AiProfilesListProviderModelsRequest) GetBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseUrl, true
}

// SetBaseUrl sets field value
func (o *AiProfilesListProviderModelsRequest) SetBaseUrl(v string) {
	o.BaseUrl = v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *AiProfilesListProviderModelsRequest) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProfilesListProviderModelsRequest) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *AiProfilesListProviderModelsRequest) IsApiKeySet() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *AiProfilesListProviderModelsRequest) SetApiKey(v string) {
	o.ApiKey = &v
}

func (o AiProfilesListProviderModelsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiProfilesListProviderModelsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["providerType"] = o.ProviderType
	toSerialize["baseUrl"] = o.BaseUrl
	if !IsNil(o.ApiKey) {
		toSerialize["apiKey"] = o.ApiKey
	}
	return toSerialize, nil
}

func (o *AiProfilesListProviderModelsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"providerType",
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

	varAiProfilesListProviderModelsRequest := _AiProfilesListProviderModelsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiProfilesListProviderModelsRequest)

	if err != nil {
		return err
	}

	*o = AiProfilesListProviderModelsRequest(varAiProfilesListProviderModelsRequest)

	return err
}

type NullableAiProfilesListProviderModelsRequest struct {
	value *AiProfilesListProviderModelsRequest
	isSet bool
}

func (v NullableAiProfilesListProviderModelsRequest) Get() *AiProfilesListProviderModelsRequest {
	return v.value
}

func (v *NullableAiProfilesListProviderModelsRequest) Set(val *AiProfilesListProviderModelsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiProfilesListProviderModelsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiProfilesListProviderModelsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiProfilesListProviderModelsRequest(val *AiProfilesListProviderModelsRequest) *NullableAiProfilesListProviderModelsRequest {
	return &NullableAiProfilesListProviderModelsRequest{value: val, isSet: true}
}

func (v NullableAiProfilesListProviderModelsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiProfilesListProviderModelsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

