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

// checks if the AiCreatePromptInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiCreatePromptInput{}

// AiCreatePromptInput Input for creating a prompt — the engine generates `id`/`createdAt`/`updatedAt`.
type AiCreatePromptInput struct {
	// The prompt name.
	Name string `json:"name"`
	// The prompt body.
	Text string `json:"text"`
	// The folder to file the prompt under. Omit or send null to leave it outside any folder.
	FolderId NullableString `json:"folderId,omitempty"`
}

type _AiCreatePromptInput AiCreatePromptInput

// NewAiCreatePromptInput instantiates a new AiCreatePromptInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiCreatePromptInput(name string, text string) *AiCreatePromptInput {
	this := AiCreatePromptInput{}
	this.Name = name
	this.Text = text
	return &this
}

// NewAiCreatePromptInputWithDefaults instantiates a new AiCreatePromptInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiCreatePromptInputWithDefaults() *AiCreatePromptInput {
	this := AiCreatePromptInput{}
	return &this
}

// GetName returns the Name field value
func (o *AiCreatePromptInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiCreatePromptInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiCreatePromptInput) SetName(v string) {
	o.Name = v
}

// GetText returns the Text field value
func (o *AiCreatePromptInput) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *AiCreatePromptInput) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *AiCreatePromptInput) SetText(v string) {
	o.Text = v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiCreatePromptInput) GetFolderId() string {
	if o == nil || IsNil(o.FolderId.Get()) {
		var ret string
		return ret
	}
	return *o.FolderId.Get()
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiCreatePromptInput) GetFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderId.Get(), o.FolderId.IsSet()
}

// HasFolderId returns a boolean if a field has been set.
func (o *AiCreatePromptInput) IsFolderIdSet() bool {
	if o != nil && o.FolderId.IsSet() {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given NullableString and assigns it to the FolderId field.
func (o *AiCreatePromptInput) SetFolderId(v string) {
	o.FolderId.Set(&v)
}
// SetFolderIdNil sets the value for FolderId to be an explicit nil
func (o *AiCreatePromptInput) SetFolderIdNil() {
	o.FolderId.Set(nil)
}

// UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
func (o *AiCreatePromptInput) UnsetFolderId() {
	o.FolderId.Unset()
}

func (o AiCreatePromptInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiCreatePromptInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["text"] = o.Text
	if o.FolderId.IsSet() {
		toSerialize["folderId"] = o.FolderId.Get()
	}
	return toSerialize, nil
}

func (o *AiCreatePromptInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"text",
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

	varAiCreatePromptInput := _AiCreatePromptInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiCreatePromptInput)

	if err != nil {
		return err
	}

	*o = AiCreatePromptInput(varAiCreatePromptInput)

	return err
}

type NullableAiCreatePromptInput struct {
	value *AiCreatePromptInput
	isSet bool
}

func (v NullableAiCreatePromptInput) Get() *AiCreatePromptInput {
	return v.value
}

func (v *NullableAiCreatePromptInput) Set(val *AiCreatePromptInput) {
	v.value = val
	v.isSet = true
}

func (v NullableAiCreatePromptInput) IsSet() bool {
	return v.isSet
}

func (v *NullableAiCreatePromptInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiCreatePromptInput(val *AiCreatePromptInput) *NullableAiCreatePromptInput {
	return &NullableAiCreatePromptInput{value: val, isSet: true}
}

func (v NullableAiCreatePromptInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiCreatePromptInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

