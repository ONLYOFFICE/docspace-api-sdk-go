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

// checks if the AiEmbeddingPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiEmbeddingPrice{}

// AiEmbeddingPrice struct for AiEmbeddingPrice
type AiEmbeddingPrice struct {
	Prompt *float64 `json:"prompt,omitempty"`
}

// NewAiEmbeddingPrice instantiates a new AiEmbeddingPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiEmbeddingPrice() *AiEmbeddingPrice {
	this := AiEmbeddingPrice{}
	return &this
}

// NewAiEmbeddingPriceWithDefaults instantiates a new AiEmbeddingPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiEmbeddingPriceWithDefaults() *AiEmbeddingPrice {
	this := AiEmbeddingPrice{}
	return &this
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (o *AiEmbeddingPrice) GetPrompt() float64 {
	if o == nil || IsNil(o.Prompt) {
		var ret float64
		return ret
	}
	return *o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiEmbeddingPrice) GetPromptOk() (*float64, bool) {
	if o == nil || IsNil(o.Prompt) {
		return nil, false
	}
	return o.Prompt, true
}

// HasPrompt returns a boolean if a field has been set.
func (o *AiEmbeddingPrice) IsPromptSet() bool {
	if o != nil && !IsNil(o.Prompt) {
		return true
	}

	return false
}

// SetPrompt gets a reference to the given float64 and assigns it to the Prompt field.
func (o *AiEmbeddingPrice) SetPrompt(v float64) {
	o.Prompt = &v
}

func (o AiEmbeddingPrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiEmbeddingPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Prompt) {
		toSerialize["prompt"] = o.Prompt
	}
	return toSerialize, nil
}

type NullableAiEmbeddingPrice struct {
	value *AiEmbeddingPrice
	isSet bool
}

func (v NullableAiEmbeddingPrice) Get() *AiEmbeddingPrice {
	return v.value
}

func (v *NullableAiEmbeddingPrice) Set(val *AiEmbeddingPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableAiEmbeddingPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableAiEmbeddingPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiEmbeddingPrice(val *AiEmbeddingPrice) *NullableAiEmbeddingPrice {
	return &NullableAiEmbeddingPrice{value: val, isSet: true}
}

func (v NullableAiEmbeddingPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiEmbeddingPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

