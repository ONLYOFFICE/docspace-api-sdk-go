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

// checks if the AiPromptsImportBundleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiPromptsImportBundleRequest{}

// AiPromptsImportBundleRequest struct for AiPromptsImportBundleRequest
type AiPromptsImportBundleRequest struct {
	// Bundle to restore.
	Bundle AiPromptBundle `json:"bundle"`
	Options *AiPromptsImportBundleRequestOptions `json:"options,omitempty"`
}

type _AiPromptsImportBundleRequest AiPromptsImportBundleRequest

// NewAiPromptsImportBundleRequest instantiates a new AiPromptsImportBundleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiPromptsImportBundleRequest(bundle AiPromptBundle) *AiPromptsImportBundleRequest {
	this := AiPromptsImportBundleRequest{}
	this.Bundle = bundle
	return &this
}

// NewAiPromptsImportBundleRequestWithDefaults instantiates a new AiPromptsImportBundleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiPromptsImportBundleRequestWithDefaults() *AiPromptsImportBundleRequest {
	this := AiPromptsImportBundleRequest{}
	return &this
}

// GetBundle returns the Bundle field value
func (o *AiPromptsImportBundleRequest) GetBundle() AiPromptBundle {
	if o == nil {
		var ret AiPromptBundle
		return ret
	}

	return o.Bundle
}

// GetBundleOk returns a tuple with the Bundle field value
// and a boolean to check if the value has been set.
func (o *AiPromptsImportBundleRequest) GetBundleOk() (*AiPromptBundle, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bundle, true
}

// SetBundle sets field value
func (o *AiPromptsImportBundleRequest) SetBundle(v AiPromptBundle) {
	o.Bundle = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *AiPromptsImportBundleRequest) GetOptions() AiPromptsImportBundleRequestOptions {
	if o == nil || IsNil(o.Options) {
		var ret AiPromptsImportBundleRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiPromptsImportBundleRequest) GetOptionsOk() (*AiPromptsImportBundleRequestOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *AiPromptsImportBundleRequest) IsOptionsSet() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given AiPromptsImportBundleRequestOptions and assigns it to the Options field.
func (o *AiPromptsImportBundleRequest) SetOptions(v AiPromptsImportBundleRequestOptions) {
	o.Options = &v
}

func (o AiPromptsImportBundleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiPromptsImportBundleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bundle"] = o.Bundle
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

func (o *AiPromptsImportBundleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bundle",
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

	varAiPromptsImportBundleRequest := _AiPromptsImportBundleRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiPromptsImportBundleRequest)

	if err != nil {
		return err
	}

	*o = AiPromptsImportBundleRequest(varAiPromptsImportBundleRequest)

	return err
}

type NullableAiPromptsImportBundleRequest struct {
	value *AiPromptsImportBundleRequest
	isSet bool
}

func (v NullableAiPromptsImportBundleRequest) Get() *AiPromptsImportBundleRequest {
	return v.value
}

func (v *NullableAiPromptsImportBundleRequest) Set(val *AiPromptsImportBundleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiPromptsImportBundleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiPromptsImportBundleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiPromptsImportBundleRequest(val *AiPromptsImportBundleRequest) *NullableAiPromptsImportBundleRequest {
	return &NullableAiPromptsImportBundleRequest{value: val, isSet: true}
}

func (v NullableAiPromptsImportBundleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiPromptsImportBundleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

