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

// checks if the GeneratePresentationToolCallParametersDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeneratePresentationToolCallParametersDto{}

// GeneratePresentationToolCallParametersDto The generate presentation tool call parameters.
type GeneratePresentationToolCallParametersDto struct {
	// What the generated presentation is about.
	Topic NullableString `json:"topic,omitempty"`
	// How many slides to generate, as the request spelled it.
	SlideCount NullableString `json:"slideCount,omitempty"`
	// The visual style the slides should be generated in.
	Style NullableString `json:"style,omitempty"`
}

// NewGeneratePresentationToolCallParametersDto instantiates a new GeneratePresentationToolCallParametersDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneratePresentationToolCallParametersDto() *GeneratePresentationToolCallParametersDto {
	this := GeneratePresentationToolCallParametersDto{}
	return &this
}

// NewGeneratePresentationToolCallParametersDtoWithDefaults instantiates a new GeneratePresentationToolCallParametersDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneratePresentationToolCallParametersDtoWithDefaults() *GeneratePresentationToolCallParametersDto {
	this := GeneratePresentationToolCallParametersDto{}
	return &this
}

// GetTopic returns the Topic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GeneratePresentationToolCallParametersDto) GetTopic() string {
	if o == nil || IsNil(o.Topic.Get()) {
		var ret string
		return ret
	}
	return *o.Topic.Get()
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GeneratePresentationToolCallParametersDto) GetTopicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Topic.Get(), o.Topic.IsSet()
}

// HasTopic returns a boolean if a field has been set.
func (o *GeneratePresentationToolCallParametersDto) IsTopicSet() bool {
	if o != nil && o.Topic.IsSet() {
		return true
	}

	return false
}

// SetTopic gets a reference to the given NullableString and assigns it to the Topic field.
func (o *GeneratePresentationToolCallParametersDto) SetTopic(v string) {
	o.Topic.Set(&v)
}
// SetTopicNil sets the value for Topic to be an explicit nil
func (o *GeneratePresentationToolCallParametersDto) SetTopicNil() {
	o.Topic.Set(nil)
}

// UnsetTopic ensures that no value is present for Topic, not even an explicit nil
func (o *GeneratePresentationToolCallParametersDto) UnsetTopic() {
	o.Topic.Unset()
}

// GetSlideCount returns the SlideCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GeneratePresentationToolCallParametersDto) GetSlideCount() string {
	if o == nil || IsNil(o.SlideCount.Get()) {
		var ret string
		return ret
	}
	return *o.SlideCount.Get()
}

// GetSlideCountOk returns a tuple with the SlideCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GeneratePresentationToolCallParametersDto) GetSlideCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SlideCount.Get(), o.SlideCount.IsSet()
}

// HasSlideCount returns a boolean if a field has been set.
func (o *GeneratePresentationToolCallParametersDto) IsSlideCountSet() bool {
	if o != nil && o.SlideCount.IsSet() {
		return true
	}

	return false
}

// SetSlideCount gets a reference to the given NullableString and assigns it to the SlideCount field.
func (o *GeneratePresentationToolCallParametersDto) SetSlideCount(v string) {
	o.SlideCount.Set(&v)
}
// SetSlideCountNil sets the value for SlideCount to be an explicit nil
func (o *GeneratePresentationToolCallParametersDto) SetSlideCountNil() {
	o.SlideCount.Set(nil)
}

// UnsetSlideCount ensures that no value is present for SlideCount, not even an explicit nil
func (o *GeneratePresentationToolCallParametersDto) UnsetSlideCount() {
	o.SlideCount.Unset()
}

// GetStyle returns the Style field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GeneratePresentationToolCallParametersDto) GetStyle() string {
	if o == nil || IsNil(o.Style.Get()) {
		var ret string
		return ret
	}
	return *o.Style.Get()
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GeneratePresentationToolCallParametersDto) GetStyleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Style.Get(), o.Style.IsSet()
}

// HasStyle returns a boolean if a field has been set.
func (o *GeneratePresentationToolCallParametersDto) IsStyleSet() bool {
	if o != nil && o.Style.IsSet() {
		return true
	}

	return false
}

// SetStyle gets a reference to the given NullableString and assigns it to the Style field.
func (o *GeneratePresentationToolCallParametersDto) SetStyle(v string) {
	o.Style.Set(&v)
}
// SetStyleNil sets the value for Style to be an explicit nil
func (o *GeneratePresentationToolCallParametersDto) SetStyleNil() {
	o.Style.Set(nil)
}

// UnsetStyle ensures that no value is present for Style, not even an explicit nil
func (o *GeneratePresentationToolCallParametersDto) UnsetStyle() {
	o.Style.Unset()
}

func (o GeneratePresentationToolCallParametersDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeneratePresentationToolCallParametersDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Topic.IsSet() {
		toSerialize["topic"] = o.Topic.Get()
	}
	if o.SlideCount.IsSet() {
		toSerialize["slideCount"] = o.SlideCount.Get()
	}
	if o.Style.IsSet() {
		toSerialize["style"] = o.Style.Get()
	}
	return toSerialize, nil
}

type NullableGeneratePresentationToolCallParametersDto struct {
	value *GeneratePresentationToolCallParametersDto
	isSet bool
}

func (v NullableGeneratePresentationToolCallParametersDto) Get() *GeneratePresentationToolCallParametersDto {
	return v.value
}

func (v *NullableGeneratePresentationToolCallParametersDto) Set(val *GeneratePresentationToolCallParametersDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneratePresentationToolCallParametersDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneratePresentationToolCallParametersDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneratePresentationToolCallParametersDto(val *GeneratePresentationToolCallParametersDto) *NullableGeneratePresentationToolCallParametersDto {
	return &NullableGeneratePresentationToolCallParametersDto{value: val, isSet: true}
}

func (v NullableGeneratePresentationToolCallParametersDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneratePresentationToolCallParametersDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

