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

// checks if the ExportMessageRequestBodyInteger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportMessageRequestBodyInteger{}

// ExportMessageRequestBodyInteger Parameters for exporting an AI chat message to a document.
type ExportMessageRequestBodyInteger struct {
	// The identifier of the destination folder where the exported document will be saved.
	FolderId int32 `json:"folderId"`
	// The file name (without extension) to use for the exported document.
	Title NullableString `json:"title"`
}

type _ExportMessageRequestBodyInteger ExportMessageRequestBodyInteger

// NewExportMessageRequestBodyInteger instantiates a new ExportMessageRequestBodyInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportMessageRequestBodyInteger(folderId int32, title NullableString) *ExportMessageRequestBodyInteger {
	this := ExportMessageRequestBodyInteger{}
	this.FolderId = folderId
	this.Title = title
	return &this
}

// NewExportMessageRequestBodyIntegerWithDefaults instantiates a new ExportMessageRequestBodyInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportMessageRequestBodyIntegerWithDefaults() *ExportMessageRequestBodyInteger {
	this := ExportMessageRequestBodyInteger{}
	return &this
}

// GetFolderId returns the FolderId field value
func (o *ExportMessageRequestBodyInteger) GetFolderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value
// and a boolean to check if the value has been set.
func (o *ExportMessageRequestBodyInteger) GetFolderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FolderId, true
}

// SetFolderId sets field value
func (o *ExportMessageRequestBodyInteger) SetFolderId(v int32) {
	o.FolderId = v
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ExportMessageRequestBodyInteger) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportMessageRequestBodyInteger) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *ExportMessageRequestBodyInteger) SetTitle(v string) {
	o.Title.Set(&v)
}

func (o ExportMessageRequestBodyInteger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportMessageRequestBodyInteger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["folderId"] = o.FolderId
	toSerialize["title"] = o.Title.Get()
	return toSerialize, nil
}

func (o *ExportMessageRequestBodyInteger) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"folderId",
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

	varExportMessageRequestBodyInteger := _ExportMessageRequestBodyInteger{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varExportMessageRequestBodyInteger)

	if err != nil {
		return err
	}

	*o = ExportMessageRequestBodyInteger(varExportMessageRequestBodyInteger)

	return err
}

type NullableExportMessageRequestBodyInteger struct {
	value *ExportMessageRequestBodyInteger
	isSet bool
}

func (v NullableExportMessageRequestBodyInteger) Get() *ExportMessageRequestBodyInteger {
	return v.value
}

func (v *NullableExportMessageRequestBodyInteger) Set(val *ExportMessageRequestBodyInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableExportMessageRequestBodyInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableExportMessageRequestBodyInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportMessageRequestBodyInteger(val *ExportMessageRequestBodyInteger) *NullableExportMessageRequestBodyInteger {
	return &NullableExportMessageRequestBodyInteger{value: val, isSet: true}
}

func (v NullableExportMessageRequestBodyInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportMessageRequestBodyInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

