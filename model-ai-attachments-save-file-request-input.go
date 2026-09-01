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

// checks if the AiAttachmentsSaveFileRequestInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiAttachmentsSaveFileRequestInput{}

// AiAttachmentsSaveFileRequestInput A file attachment draft to persist.
type AiAttachmentsSaveFileRequestInput struct {
	// Storage path/key of the file.
	Path string `json:"path"`
	// File contents.
	Content string `json:"content"`
	// File type discriminator.
	Type float32 `json:"type"`
	// Optional display title.
	Title *string `json:"title,omitempty"`
}

type _AiAttachmentsSaveFileRequestInput AiAttachmentsSaveFileRequestInput

// NewAiAttachmentsSaveFileRequestInput instantiates a new AiAttachmentsSaveFileRequestInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAttachmentsSaveFileRequestInput(path string, content string, type_ float32) *AiAttachmentsSaveFileRequestInput {
	this := AiAttachmentsSaveFileRequestInput{}
	this.Path = path
	this.Content = content
	this.Type = type_
	return &this
}

// NewAiAttachmentsSaveFileRequestInputWithDefaults instantiates a new AiAttachmentsSaveFileRequestInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAttachmentsSaveFileRequestInputWithDefaults() *AiAttachmentsSaveFileRequestInput {
	this := AiAttachmentsSaveFileRequestInput{}
	return &this
}

// GetPath returns the Path field value
func (o *AiAttachmentsSaveFileRequestInput) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *AiAttachmentsSaveFileRequestInput) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *AiAttachmentsSaveFileRequestInput) SetPath(v string) {
	o.Path = v
}

// GetContent returns the Content field value
func (o *AiAttachmentsSaveFileRequestInput) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *AiAttachmentsSaveFileRequestInput) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *AiAttachmentsSaveFileRequestInput) SetContent(v string) {
	o.Content = v
}

// GetType returns the Type field value
func (o *AiAttachmentsSaveFileRequestInput) GetType() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiAttachmentsSaveFileRequestInput) GetTypeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AiAttachmentsSaveFileRequestInput) SetType(v float32) {
	o.Type = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AiAttachmentsSaveFileRequestInput) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAttachmentsSaveFileRequestInput) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AiAttachmentsSaveFileRequestInput) IsTitleSet() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AiAttachmentsSaveFileRequestInput) SetTitle(v string) {
	o.Title = &v
}

func (o AiAttachmentsSaveFileRequestInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiAttachmentsSaveFileRequestInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	toSerialize["content"] = o.Content
	toSerialize["type"] = o.Type
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

func (o *AiAttachmentsSaveFileRequestInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"path",
		"content",
		"type",
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

	varAiAttachmentsSaveFileRequestInput := _AiAttachmentsSaveFileRequestInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiAttachmentsSaveFileRequestInput)

	if err != nil {
		return err
	}

	*o = AiAttachmentsSaveFileRequestInput(varAiAttachmentsSaveFileRequestInput)

	return err
}

type NullableAiAttachmentsSaveFileRequestInput struct {
	value *AiAttachmentsSaveFileRequestInput
	isSet bool
}

func (v NullableAiAttachmentsSaveFileRequestInput) Get() *AiAttachmentsSaveFileRequestInput {
	return v.value
}

func (v *NullableAiAttachmentsSaveFileRequestInput) Set(val *AiAttachmentsSaveFileRequestInput) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAttachmentsSaveFileRequestInput) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAttachmentsSaveFileRequestInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAttachmentsSaveFileRequestInput(val *AiAttachmentsSaveFileRequestInput) *NullableAiAttachmentsSaveFileRequestInput {
	return &NullableAiAttachmentsSaveFileRequestInput{value: val, isSet: true}
}

func (v NullableAiAttachmentsSaveFileRequestInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAttachmentsSaveFileRequestInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

