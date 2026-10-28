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

// checks if the AiPromptFolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiPromptFolder{}

// AiPromptFolder Folder for organizing saved prompts.
type AiPromptFolder struct {
	// Unique folder identifier (UUID).
	Id string `json:"id"`
	// Folder display name.
	Name string `json:"name"`
	// Timestamp (ms since epoch) when the folder was created.
	CreatedAt float32 `json:"createdAt"`
	// Timestamp (ms since epoch) of the last folder modification.
	UpdatedAt float32 `json:"updatedAt"`
}

type _AiPromptFolder AiPromptFolder

// NewAiPromptFolder instantiates a new AiPromptFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiPromptFolder(id string, name string, createdAt float32, updatedAt float32) *AiPromptFolder {
	this := AiPromptFolder{}
	this.Id = id
	this.Name = name
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewAiPromptFolderWithDefaults instantiates a new AiPromptFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiPromptFolderWithDefaults() *AiPromptFolder {
	this := AiPromptFolder{}
	return &this
}

// GetId returns the Id field value
func (o *AiPromptFolder) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiPromptFolder) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiPromptFolder) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AiPromptFolder) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiPromptFolder) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiPromptFolder) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AiPromptFolder) GetCreatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiPromptFolder) GetCreatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AiPromptFolder) SetCreatedAt(v float32) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *AiPromptFolder) GetUpdatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AiPromptFolder) GetUpdatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *AiPromptFolder) SetUpdatedAt(v float32) {
	o.UpdatedAt = v
}

func (o AiPromptFolder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiPromptFolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *AiPromptFolder) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"createdAt",
		"updatedAt",
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

	varAiPromptFolder := _AiPromptFolder{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiPromptFolder)

	if err != nil {
		return err
	}

	*o = AiPromptFolder(varAiPromptFolder)

	return err
}

type NullableAiPromptFolder struct {
	value *AiPromptFolder
	isSet bool
}

func (v NullableAiPromptFolder) Get() *AiPromptFolder {
	return v.value
}

func (v *NullableAiPromptFolder) Set(val *AiPromptFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableAiPromptFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableAiPromptFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiPromptFolder(val *AiPromptFolder) *NullableAiPromptFolder {
	return &NullableAiPromptFolder{value: val, isSet: true}
}

func (v NullableAiPromptFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiPromptFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

