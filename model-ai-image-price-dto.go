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

// checks if the AiImagePriceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiImagePriceDto{}

// AiImagePriceDto What an image model charges: the tokens of the request and the images that come out of it.
type AiImagePriceDto struct {
	// The cost of one million tokens sent to the image model, which is the prompt describing the picture.
	Prompt *float64 `json:"prompt,omitempty"`
	// The cost of one million tokens the image model writes back alongside the picture.
	Completion *float64 `json:"completion,omitempty"`
	// The cost of one produced image, charged on top of the token amounts above.
	Image *float64 `json:"image,omitempty"`
}

// NewAiImagePriceDto instantiates a new AiImagePriceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiImagePriceDto() *AiImagePriceDto {
	this := AiImagePriceDto{}
	return &this
}

// NewAiImagePriceDtoWithDefaults instantiates a new AiImagePriceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiImagePriceDtoWithDefaults() *AiImagePriceDto {
	this := AiImagePriceDto{}
	return &this
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (o *AiImagePriceDto) GetPrompt() float64 {
	if o == nil || IsNil(o.Prompt) {
		var ret float64
		return ret
	}
	return *o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiImagePriceDto) GetPromptOk() (*float64, bool) {
	if o == nil || IsNil(o.Prompt) {
		return nil, false
	}
	return o.Prompt, true
}

// HasPrompt returns a boolean if a field has been set.
func (o *AiImagePriceDto) IsPromptSet() bool {
	if o != nil && !IsNil(o.Prompt) {
		return true
	}

	return false
}

// SetPrompt gets a reference to the given float64 and assigns it to the Prompt field.
func (o *AiImagePriceDto) SetPrompt(v float64) {
	o.Prompt = &v
}

// GetCompletion returns the Completion field value if set, zero value otherwise.
func (o *AiImagePriceDto) GetCompletion() float64 {
	if o == nil || IsNil(o.Completion) {
		var ret float64
		return ret
	}
	return *o.Completion
}

// GetCompletionOk returns a tuple with the Completion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiImagePriceDto) GetCompletionOk() (*float64, bool) {
	if o == nil || IsNil(o.Completion) {
		return nil, false
	}
	return o.Completion, true
}

// HasCompletion returns a boolean if a field has been set.
func (o *AiImagePriceDto) IsCompletionSet() bool {
	if o != nil && !IsNil(o.Completion) {
		return true
	}

	return false
}

// SetCompletion gets a reference to the given float64 and assigns it to the Completion field.
func (o *AiImagePriceDto) SetCompletion(v float64) {
	o.Completion = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *AiImagePriceDto) GetImage() float64 {
	if o == nil || IsNil(o.Image) {
		var ret float64
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiImagePriceDto) GetImageOk() (*float64, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *AiImagePriceDto) IsImageSet() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given float64 and assigns it to the Image field.
func (o *AiImagePriceDto) SetImage(v float64) {
	o.Image = &v
}

func (o AiImagePriceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiImagePriceDto) ToMap() (map[string]interface{}, error) {
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

type NullableAiImagePriceDto struct {
	value *AiImagePriceDto
	isSet bool
}

func (v NullableAiImagePriceDto) Get() *AiImagePriceDto {
	return v.value
}

func (v *NullableAiImagePriceDto) Set(val *AiImagePriceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiImagePriceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiImagePriceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiImagePriceDto(val *AiImagePriceDto) *NullableAiImagePriceDto {
	return &NullableAiImagePriceDto{value: val, isSet: true}
}

func (v NullableAiImagePriceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiImagePriceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

