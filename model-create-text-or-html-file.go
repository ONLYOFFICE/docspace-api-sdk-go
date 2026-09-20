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

// checks if the CreateTextOrHtmlFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTextOrHtmlFile{}

// CreateTextOrHtmlFile The parameters of a text or HTML file created from content sent in the request.
type CreateTextOrHtmlFile struct {
	// The title of the file. The extension the operation stands for is appended unless the title already ends with  it, so Notes becomes Notes.txt or Notes.html.
	Title NullableString `json:"title"`
	// The content of the file, as plain text or as HTML markup. A request carrying none is rejected as an invalid  request, and for a text file content that looks like markup makes the portal store it as HTML instead.
	Content NullableString `json:"content,omitempty"`
	// What to do when the folder already holds a file of this title, the other way round than the name reads: `true`  updates that file and adds a version to its history, `false` creates another file and makes its title unique,  as in Notes (1).txt.
	CreateNewIfExist *bool `json:"createNewIfExist,omitempty"`
}

type _CreateTextOrHtmlFile CreateTextOrHtmlFile

// NewCreateTextOrHtmlFile instantiates a new CreateTextOrHtmlFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTextOrHtmlFile(title NullableString) *CreateTextOrHtmlFile {
	this := CreateTextOrHtmlFile{}
	this.Title = title
	return &this
}

// NewCreateTextOrHtmlFileWithDefaults instantiates a new CreateTextOrHtmlFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTextOrHtmlFileWithDefaults() *CreateTextOrHtmlFile {
	this := CreateTextOrHtmlFile{}
	return &this
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateTextOrHtmlFile) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTextOrHtmlFile) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *CreateTextOrHtmlFile) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTextOrHtmlFile) GetContent() string {
	if o == nil || IsNil(o.Content.Get()) {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTextOrHtmlFile) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *CreateTextOrHtmlFile) IsContentSet() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *CreateTextOrHtmlFile) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *CreateTextOrHtmlFile) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *CreateTextOrHtmlFile) UnsetContent() {
	o.Content.Unset()
}

// GetCreateNewIfExist returns the CreateNewIfExist field value if set, zero value otherwise.
func (o *CreateTextOrHtmlFile) GetCreateNewIfExist() bool {
	if o == nil || IsNil(o.CreateNewIfExist) {
		var ret bool
		return ret
	}
	return *o.CreateNewIfExist
}

// GetCreateNewIfExistOk returns a tuple with the CreateNewIfExist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTextOrHtmlFile) GetCreateNewIfExistOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateNewIfExist) {
		return nil, false
	}
	return o.CreateNewIfExist, true
}

// HasCreateNewIfExist returns a boolean if a field has been set.
func (o *CreateTextOrHtmlFile) IsCreateNewIfExistSet() bool {
	if o != nil && !IsNil(o.CreateNewIfExist) {
		return true
	}

	return false
}

// SetCreateNewIfExist gets a reference to the given bool and assigns it to the CreateNewIfExist field.
func (o *CreateTextOrHtmlFile) SetCreateNewIfExist(v bool) {
	o.CreateNewIfExist = &v
}

func (o CreateTextOrHtmlFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTextOrHtmlFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title.Get()
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if !IsNil(o.CreateNewIfExist) {
		toSerialize["createNewIfExist"] = o.CreateNewIfExist
	}
	return toSerialize, nil
}

func (o *CreateTextOrHtmlFile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
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

	varCreateTextOrHtmlFile := _CreateTextOrHtmlFile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCreateTextOrHtmlFile)

	if err != nil {
		return err
	}

	*o = CreateTextOrHtmlFile(varCreateTextOrHtmlFile)

	return err
}

type NullableCreateTextOrHtmlFile struct {
	value *CreateTextOrHtmlFile
	isSet bool
}

func (v NullableCreateTextOrHtmlFile) Get() *CreateTextOrHtmlFile {
	return v.value
}

func (v *NullableCreateTextOrHtmlFile) Set(val *CreateTextOrHtmlFile) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTextOrHtmlFile) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTextOrHtmlFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTextOrHtmlFile(val *CreateTextOrHtmlFile) *NullableCreateTextOrHtmlFile {
	return &NullableCreateTextOrHtmlFile{value: val, isSet: true}
}

func (v NullableCreateTextOrHtmlFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTextOrHtmlFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

