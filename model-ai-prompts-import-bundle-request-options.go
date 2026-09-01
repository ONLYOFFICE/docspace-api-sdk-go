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
)

// checks if the AiPromptsImportBundleRequestOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiPromptsImportBundleRequestOptions{}

// AiPromptsImportBundleRequestOptions Import options.
type AiPromptsImportBundleRequestOptions struct {
	Mode *AiImportMode `json:"mode,omitempty"`
}

// NewAiPromptsImportBundleRequestOptions instantiates a new AiPromptsImportBundleRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiPromptsImportBundleRequestOptions() *AiPromptsImportBundleRequestOptions {
	this := AiPromptsImportBundleRequestOptions{}
	return &this
}

// NewAiPromptsImportBundleRequestOptionsWithDefaults instantiates a new AiPromptsImportBundleRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiPromptsImportBundleRequestOptionsWithDefaults() *AiPromptsImportBundleRequestOptions {
	this := AiPromptsImportBundleRequestOptions{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *AiPromptsImportBundleRequestOptions) GetMode() AiImportMode {
	if o == nil || IsNil(o.Mode) {
		var ret AiImportMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiPromptsImportBundleRequestOptions) GetModeOk() (*AiImportMode, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *AiPromptsImportBundleRequestOptions) IsModeSet() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given AiImportMode and assigns it to the Mode field.
func (o *AiPromptsImportBundleRequestOptions) SetMode(v AiImportMode) {
	o.Mode = &v
}

func (o AiPromptsImportBundleRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiPromptsImportBundleRequestOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableAiPromptsImportBundleRequestOptions struct {
	value *AiPromptsImportBundleRequestOptions
	isSet bool
}

func (v NullableAiPromptsImportBundleRequestOptions) Get() *AiPromptsImportBundleRequestOptions {
	return v.value
}

func (v *NullableAiPromptsImportBundleRequestOptions) Set(val *AiPromptsImportBundleRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAiPromptsImportBundleRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAiPromptsImportBundleRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiPromptsImportBundleRequestOptions(val *AiPromptsImportBundleRequestOptions) *NullableAiPromptsImportBundleRequestOptions {
	return &NullableAiPromptsImportBundleRequestOptions{value: val, isSet: true}
}

func (v NullableAiPromptsImportBundleRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiPromptsImportBundleRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

