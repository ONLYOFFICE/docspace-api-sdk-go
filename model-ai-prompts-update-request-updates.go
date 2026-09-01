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

// checks if the AiPromptsUpdateRequestUpdates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiPromptsUpdateRequestUpdates{}

// AiPromptsUpdateRequestUpdates Fields to change.
type AiPromptsUpdateRequestUpdates struct {
	Name *string `json:"name,omitempty"`
	Text *string `json:"text,omitempty"`
	FolderId NullableString `json:"folderId,omitempty"`
}

// NewAiPromptsUpdateRequestUpdates instantiates a new AiPromptsUpdateRequestUpdates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiPromptsUpdateRequestUpdates() *AiPromptsUpdateRequestUpdates {
	this := AiPromptsUpdateRequestUpdates{}
	return &this
}

// NewAiPromptsUpdateRequestUpdatesWithDefaults instantiates a new AiPromptsUpdateRequestUpdates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiPromptsUpdateRequestUpdatesWithDefaults() *AiPromptsUpdateRequestUpdates {
	this := AiPromptsUpdateRequestUpdates{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AiPromptsUpdateRequestUpdates) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiPromptsUpdateRequestUpdates) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AiPromptsUpdateRequestUpdates) IsNameSet() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AiPromptsUpdateRequestUpdates) SetName(v string) {
	o.Name = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *AiPromptsUpdateRequestUpdates) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiPromptsUpdateRequestUpdates) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *AiPromptsUpdateRequestUpdates) IsTextSet() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *AiPromptsUpdateRequestUpdates) SetText(v string) {
	o.Text = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiPromptsUpdateRequestUpdates) GetFolderId() string {
	if o == nil || IsNil(o.FolderId.Get()) {
		var ret string
		return ret
	}
	return *o.FolderId.Get()
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiPromptsUpdateRequestUpdates) GetFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderId.Get(), o.FolderId.IsSet()
}

// HasFolderId returns a boolean if a field has been set.
func (o *AiPromptsUpdateRequestUpdates) IsFolderIdSet() bool {
	if o != nil && o.FolderId.IsSet() {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given NullableString and assigns it to the FolderId field.
func (o *AiPromptsUpdateRequestUpdates) SetFolderId(v string) {
	o.FolderId.Set(&v)
}
// SetFolderIdNil sets the value for FolderId to be an explicit nil
func (o *AiPromptsUpdateRequestUpdates) SetFolderIdNil() {
	o.FolderId.Set(nil)
}

// UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
func (o *AiPromptsUpdateRequestUpdates) UnsetFolderId() {
	o.FolderId.Unset()
}

func (o AiPromptsUpdateRequestUpdates) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiPromptsUpdateRequestUpdates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if o.FolderId.IsSet() {
		toSerialize["folderId"] = o.FolderId.Get()
	}
	return toSerialize, nil
}

type NullableAiPromptsUpdateRequestUpdates struct {
	value *AiPromptsUpdateRequestUpdates
	isSet bool
}

func (v NullableAiPromptsUpdateRequestUpdates) Get() *AiPromptsUpdateRequestUpdates {
	return v.value
}

func (v *NullableAiPromptsUpdateRequestUpdates) Set(val *AiPromptsUpdateRequestUpdates) {
	v.value = val
	v.isSet = true
}

func (v NullableAiPromptsUpdateRequestUpdates) IsSet() bool {
	return v.isSet
}

func (v *NullableAiPromptsUpdateRequestUpdates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiPromptsUpdateRequestUpdates(val *AiPromptsUpdateRequestUpdates) *NullableAiPromptsUpdateRequestUpdates {
	return &NullableAiPromptsUpdateRequestUpdates{value: val, isSet: true}
}

func (v NullableAiPromptsUpdateRequestUpdates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiPromptsUpdateRequestUpdates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

