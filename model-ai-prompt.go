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

// checks if the AiPrompt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiPrompt{}

// AiPrompt Saved prompt template that users can quickly insert into the chat.
type AiPrompt struct {
	// Unique prompt identifier (UUID).
	Id string `json:"id"`
	// Prompt display name shown in the prompt picker.
	Name string `json:"name"`
	// Prompt template text. May contain placeholder tokens.
	Text string `json:"text"`
	// Optional parent folder ID. `undefined` means the prompt is at the root level.
	FolderId *string `json:"folderId,omitempty"`
	// Timestamp (ms since epoch) when the prompt was created.
	CreatedAt float32 `json:"createdAt"`
	// Timestamp (ms since epoch) of the last prompt modification.
	UpdatedAt float32 `json:"updatedAt"`
}

type _AiPrompt AiPrompt

// NewAiPrompt instantiates a new AiPrompt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiPrompt(id string, name string, text string, createdAt float32, updatedAt float32) *AiPrompt {
	this := AiPrompt{}
	this.Id = id
	this.Name = name
	this.Text = text
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewAiPromptWithDefaults instantiates a new AiPrompt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiPromptWithDefaults() *AiPrompt {
	this := AiPrompt{}
	return &this
}

// GetId returns the Id field value
func (o *AiPrompt) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiPrompt) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiPrompt) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AiPrompt) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiPrompt) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiPrompt) SetName(v string) {
	o.Name = v
}

// GetText returns the Text field value
func (o *AiPrompt) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *AiPrompt) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *AiPrompt) SetText(v string) {
	o.Text = v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *AiPrompt) GetFolderId() string {
	if o == nil || IsNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiPrompt) GetFolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.FolderId) {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *AiPrompt) IsFolderIdSet() bool {
	if o != nil && !IsNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *AiPrompt) SetFolderId(v string) {
	o.FolderId = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AiPrompt) GetCreatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiPrompt) GetCreatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AiPrompt) SetCreatedAt(v float32) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *AiPrompt) GetUpdatedAt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AiPrompt) GetUpdatedAtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *AiPrompt) SetUpdatedAt(v float32) {
	o.UpdatedAt = v
}

func (o AiPrompt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiPrompt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["text"] = o.Text
	if !IsNil(o.FolderId) {
		toSerialize["folderId"] = o.FolderId
	}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *AiPrompt) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"text",
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

	varAiPrompt := _AiPrompt{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiPrompt)

	if err != nil {
		return err
	}

	*o = AiPrompt(varAiPrompt)

	return err
}

type NullableAiPrompt struct {
	value *AiPrompt
	isSet bool
}

func (v NullableAiPrompt) Get() *AiPrompt {
	return v.value
}

func (v *NullableAiPrompt) Set(val *AiPrompt) {
	v.value = val
	v.isSet = true
}

func (v NullableAiPrompt) IsSet() bool {
	return v.isSet
}

func (v *NullableAiPrompt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiPrompt(val *AiPrompt) *NullableAiPrompt {
	return &NullableAiPrompt{value: val, isSet: true}
}

func (v NullableAiPrompt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiPrompt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

