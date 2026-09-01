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

// checks if the AiImagePrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiImagePrice{}

// AiImagePrice The price of an image model: per prompt token and per generated image.
type AiImagePrice struct {
	// The price of a single prompt token.
	Prompt *float64 `json:"prompt,omitempty"`
	// The cost associated with the completion of a prompt in an AI model.
	Completion *float64 `json:"completion,omitempty"`
	// The price of a single generated image.
	Image *float64 `json:"image,omitempty"`
}

// NewAiImagePrice instantiates a new AiImagePrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiImagePrice() *AiImagePrice {
	this := AiImagePrice{}
	return &this
}

// NewAiImagePriceWithDefaults instantiates a new AiImagePrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiImagePriceWithDefaults() *AiImagePrice {
	this := AiImagePrice{}
	return &this
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (o *AiImagePrice) GetPrompt() float64 {
	if o == nil || IsNil(o.Prompt) {
		var ret float64
		return ret
	}
	return *o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiImagePrice) GetPromptOk() (*float64, bool) {
	if o == nil || IsNil(o.Prompt) {
		return nil, false
	}
	return o.Prompt, true
}

// HasPrompt returns a boolean if a field has been set.
func (o *AiImagePrice) IsPromptSet() bool {
	if o != nil && !IsNil(o.Prompt) {
		return true
	}

	return false
}

// SetPrompt gets a reference to the given float64 and assigns it to the Prompt field.
func (o *AiImagePrice) SetPrompt(v float64) {
	o.Prompt = &v
}

// GetCompletion returns the Completion field value if set, zero value otherwise.
func (o *AiImagePrice) GetCompletion() float64 {
	if o == nil || IsNil(o.Completion) {
		var ret float64
		return ret
	}
	return *o.Completion
}

// GetCompletionOk returns a tuple with the Completion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiImagePrice) GetCompletionOk() (*float64, bool) {
	if o == nil || IsNil(o.Completion) {
		return nil, false
	}
	return o.Completion, true
}

// HasCompletion returns a boolean if a field has been set.
func (o *AiImagePrice) IsCompletionSet() bool {
	if o != nil && !IsNil(o.Completion) {
		return true
	}

	return false
}

// SetCompletion gets a reference to the given float64 and assigns it to the Completion field.
func (o *AiImagePrice) SetCompletion(v float64) {
	o.Completion = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *AiImagePrice) GetImage() float64 {
	if o == nil || IsNil(o.Image) {
		var ret float64
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiImagePrice) GetImageOk() (*float64, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *AiImagePrice) IsImageSet() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given float64 and assigns it to the Image field.
func (o *AiImagePrice) SetImage(v float64) {
	o.Image = &v
}

func (o AiImagePrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiImagePrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Prompt) {
		toSerialize["prompt"] = o.Prompt
	}
	if !IsNil(o.Completion) {
		toSerialize["completion"] = o.Completion
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	return toSerialize, nil
}

type NullableAiImagePrice struct {
	value *AiImagePrice
	isSet bool
}

func (v NullableAiImagePrice) Get() *AiImagePrice {
	return v.value
}

func (v *NullableAiImagePrice) Set(val *AiImagePrice) {
	v.value = val
	v.isSet = true
}

func (v NullableAiImagePrice) IsSet() bool {
	return v.isSet
}

func (v *NullableAiImagePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiImagePrice(val *AiImagePrice) *NullableAiImagePrice {
	return &NullableAiImagePrice{value: val, isSet: true}
}

func (v NullableAiImagePrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiImagePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

