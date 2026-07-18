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

// checks if the ExportMessageRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportMessageRequestBody{}

// ExportMessageRequestBody Parameters for exporting an AI chat message to a document.
type ExportMessageRequestBody struct {
	FolderId ExportChatRequestBodyFolderId `json:"folderId"`
	// The file name (without extension) to use for the exported document.
	Title NullableString `json:"title"`
}

type _ExportMessageRequestBody ExportMessageRequestBody

// NewExportMessageRequestBody instantiates a new ExportMessageRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportMessageRequestBody(folderId ExportChatRequestBodyFolderId, title NullableString) *ExportMessageRequestBody {
	this := ExportMessageRequestBody{}
	this.FolderId = folderId
	this.Title = title
	return &this
}

// NewExportMessageRequestBodyWithDefaults instantiates a new ExportMessageRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportMessageRequestBodyWithDefaults() *ExportMessageRequestBody {
	this := ExportMessageRequestBody{}
	return &this
}

// GetFolderId returns the FolderId field value
func (o *ExportMessageRequestBody) GetFolderId() ExportChatRequestBodyFolderId {
	if o == nil {
		var ret ExportChatRequestBodyFolderId
		return ret
	}

	return o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value
// and a boolean to check if the value has been set.
func (o *ExportMessageRequestBody) GetFolderIdOk() (*ExportChatRequestBodyFolderId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FolderId, true
}

// SetFolderId sets field value
func (o *ExportMessageRequestBody) SetFolderId(v ExportChatRequestBodyFolderId) {
	o.FolderId = v
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ExportMessageRequestBody) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportMessageRequestBody) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *ExportMessageRequestBody) SetTitle(v string) {
	o.Title.Set(&v)
}

func (o ExportMessageRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportMessageRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["folderId"] = o.FolderId
	toSerialize["title"] = o.Title.Get()
	return toSerialize, nil
}

func (o *ExportMessageRequestBody) UnmarshalJSON(data []byte) (err error) {
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

	varExportMessageRequestBody := _ExportMessageRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varExportMessageRequestBody)

	if err != nil {
		return err
	}

	*o = ExportMessageRequestBody(varExportMessageRequestBody)

	return err
}

type NullableExportMessageRequestBody struct {
	value *ExportMessageRequestBody
	isSet bool
}

func (v NullableExportMessageRequestBody) Get() *ExportMessageRequestBody {
	return v.value
}

func (v *NullableExportMessageRequestBody) Set(val *ExportMessageRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableExportMessageRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableExportMessageRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportMessageRequestBody(val *ExportMessageRequestBody) *NullableExportMessageRequestBody {
	return &NullableExportMessageRequestBody{value: val, isSet: true}
}

func (v NullableExportMessageRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportMessageRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

