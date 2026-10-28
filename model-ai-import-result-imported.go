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

// checks if the AiImportResultImported type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiImportResultImported{}

// AiImportResultImported How many folders and prompts were created. Present on success.
type AiImportResultImported struct {
	Folders float32 `json:"folders"`
	Prompts float32 `json:"prompts"`
}

type _AiImportResultImported AiImportResultImported

// NewAiImportResultImported instantiates a new AiImportResultImported object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiImportResultImported(folders float32, prompts float32) *AiImportResultImported {
	this := AiImportResultImported{}
	this.Folders = folders
	this.Prompts = prompts
	return &this
}

// NewAiImportResultImportedWithDefaults instantiates a new AiImportResultImported object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiImportResultImportedWithDefaults() *AiImportResultImported {
	this := AiImportResultImported{}
	return &this
}

// GetFolders returns the Folders field value
func (o *AiImportResultImported) GetFolders() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Folders
}

// GetFoldersOk returns a tuple with the Folders field value
// and a boolean to check if the value has been set.
func (o *AiImportResultImported) GetFoldersOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Folders, true
}

// SetFolders sets field value
func (o *AiImportResultImported) SetFolders(v float32) {
	o.Folders = v
}

// GetPrompts returns the Prompts field value
func (o *AiImportResultImported) GetPrompts() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Prompts
}

// GetPromptsOk returns a tuple with the Prompts field value
// and a boolean to check if the value has been set.
func (o *AiImportResultImported) GetPromptsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompts, true
}

// SetPrompts sets field value
func (o *AiImportResultImported) SetPrompts(v float32) {
	o.Prompts = v
}

func (o AiImportResultImported) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiImportResultImported) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["folders"] = o.Folders
	toSerialize["prompts"] = o.Prompts
	return toSerialize, nil
}

func (o *AiImportResultImported) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varAiImportResultImported := _AiImportResultImported{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiImportResultImported)

	if err != nil {
		return err
	}

	*o = AiImportResultImported(varAiImportResultImported)

	return err
}

type NullableAiImportResultImported struct {
	value *AiImportResultImported
	isSet bool
}

func (v NullableAiImportResultImported) Get() *AiImportResultImported {
	return v.value
}

func (v *NullableAiImportResultImported) Set(val *AiImportResultImported) {
	v.value = val
	v.isSet = true
}

func (v NullableAiImportResultImported) IsSet() bool {
	return v.isSet
}

func (v *NullableAiImportResultImported) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiImportResultImported(val *AiImportResultImported) *NullableAiImportResultImported {
	return &NullableAiImportResultImported{value: val, isSet: true}
}

func (v NullableAiImportResultImported) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiImportResultImported) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

