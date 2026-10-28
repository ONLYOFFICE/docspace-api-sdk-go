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

// checks if the AiPromptBundle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiPromptBundle{}

// AiPromptBundle Versioned, self-contained bundle of every saved prompt and folder. Stable wire format — `version` lets the import path migrate older shapes if the schema ever changes.
type AiPromptBundle struct {
	// The bundle format version, so an import can migrate an older export.
	Version float32 `json:"version"`
	// Every exported prompt folder.
	Folders []AiPromptFolder `json:"folders"`
	// Every exported prompt.
	Prompts []AiPrompt `json:"prompts"`
}

type _AiPromptBundle AiPromptBundle

// NewAiPromptBundle instantiates a new AiPromptBundle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiPromptBundle(version float32, folders []AiPromptFolder, prompts []AiPrompt) *AiPromptBundle {
	this := AiPromptBundle{}
	this.Version = version
	this.Folders = folders
	this.Prompts = prompts
	return &this
}

// NewAiPromptBundleWithDefaults instantiates a new AiPromptBundle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiPromptBundleWithDefaults() *AiPromptBundle {
	this := AiPromptBundle{}
	return &this
}

// GetVersion returns the Version field value
func (o *AiPromptBundle) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *AiPromptBundle) GetVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *AiPromptBundle) SetVersion(v float32) {
	o.Version = v
}

// GetFolders returns the Folders field value
func (o *AiPromptBundle) GetFolders() []AiPromptFolder {
	if o == nil {
		var ret []AiPromptFolder
		return ret
	}

	return o.Folders
}

// GetFoldersOk returns a tuple with the Folders field value
// and a boolean to check if the value has been set.
func (o *AiPromptBundle) GetFoldersOk() ([]AiPromptFolder, bool) {
	if o == nil {
		return nil, false
	}
	return o.Folders, true
}

// SetFolders sets field value
func (o *AiPromptBundle) SetFolders(v []AiPromptFolder) {
	o.Folders = v
}

// GetPrompts returns the Prompts field value
func (o *AiPromptBundle) GetPrompts() []AiPrompt {
	if o == nil {
		var ret []AiPrompt
		return ret
	}

	return o.Prompts
}

// GetPromptsOk returns a tuple with the Prompts field value
// and a boolean to check if the value has been set.
func (o *AiPromptBundle) GetPromptsOk() ([]AiPrompt, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prompts, true
}

// SetPrompts sets field value
func (o *AiPromptBundle) SetPrompts(v []AiPrompt) {
	o.Prompts = v
}

func (o AiPromptBundle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiPromptBundle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["folders"] = o.Folders
	toSerialize["prompts"] = o.Prompts
	return toSerialize, nil
}

func (o *AiPromptBundle) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"version",
		"folders",
		"prompts",
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

	varAiPromptBundle := _AiPromptBundle{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiPromptBundle)

	if err != nil {
		return err
	}

	*o = AiPromptBundle(varAiPromptBundle)

	return err
}

type NullableAiPromptBundle struct {
	value *AiPromptBundle
	isSet bool
}

func (v NullableAiPromptBundle) Get() *AiPromptBundle {
	return v.value
}

func (v *NullableAiPromptBundle) Set(val *AiPromptBundle) {
	v.value = val
	v.isSet = true
}

func (v NullableAiPromptBundle) IsSet() bool {
	return v.isSet
}

func (v *NullableAiPromptBundle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiPromptBundle(val *AiPromptBundle) *NullableAiPromptBundle {
	return &NullableAiPromptBundle{value: val, isSet: true}
}

func (v NullableAiPromptBundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiPromptBundle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

