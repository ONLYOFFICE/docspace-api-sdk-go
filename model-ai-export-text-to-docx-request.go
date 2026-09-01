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

// checks if the AiExportTextToDocxRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiExportTextToDocxRequest{}

// AiExportTextToDocxRequest struct for AiExportTextToDocxRequest
type AiExportTextToDocxRequest struct {
	// Document title (also the file name).
	Title string `json:"title"`
	// Markdown content to convert.
	Content string `json:"content"`
	FolderId AiExportTextToDocxRequestFolderId `json:"folderId"`
}

type _AiExportTextToDocxRequest AiExportTextToDocxRequest

// NewAiExportTextToDocxRequest instantiates a new AiExportTextToDocxRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiExportTextToDocxRequest(title string, content string, folderId AiExportTextToDocxRequestFolderId) *AiExportTextToDocxRequest {
	this := AiExportTextToDocxRequest{}
	this.Title = title
	this.Content = content
	this.FolderId = folderId
	return &this
}

// NewAiExportTextToDocxRequestWithDefaults instantiates a new AiExportTextToDocxRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiExportTextToDocxRequestWithDefaults() *AiExportTextToDocxRequest {
	this := AiExportTextToDocxRequest{}
	return &this
}

// GetTitle returns the Title field value
func (o *AiExportTextToDocxRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AiExportTextToDocxRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *AiExportTextToDocxRequest) SetTitle(v string) {
	o.Title = v
}

// GetContent returns the Content field value
func (o *AiExportTextToDocxRequest) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *AiExportTextToDocxRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *AiExportTextToDocxRequest) SetContent(v string) {
	o.Content = v
}

// GetFolderId returns the FolderId field value
func (o *AiExportTextToDocxRequest) GetFolderId() AiExportTextToDocxRequestFolderId {
	if o == nil {
		var ret AiExportTextToDocxRequestFolderId
		return ret
	}

	return o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value
// and a boolean to check if the value has been set.
func (o *AiExportTextToDocxRequest) GetFolderIdOk() (*AiExportTextToDocxRequestFolderId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FolderId, true
}

// SetFolderId sets field value
func (o *AiExportTextToDocxRequest) SetFolderId(v AiExportTextToDocxRequestFolderId) {
	o.FolderId = v
}

func (o AiExportTextToDocxRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiExportTextToDocxRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["content"] = o.Content
	toSerialize["folderId"] = o.FolderId
	return toSerialize, nil
}

func (o *AiExportTextToDocxRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"content",
		"folderId",
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

	varAiExportTextToDocxRequest := _AiExportTextToDocxRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiExportTextToDocxRequest)

	if err != nil {
		return err
	}

	*o = AiExportTextToDocxRequest(varAiExportTextToDocxRequest)

	return err
}

type NullableAiExportTextToDocxRequest struct {
	value *AiExportTextToDocxRequest
	isSet bool
}

func (v NullableAiExportTextToDocxRequest) Get() *AiExportTextToDocxRequest {
	return v.value
}

func (v *NullableAiExportTextToDocxRequest) Set(val *AiExportTextToDocxRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiExportTextToDocxRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiExportTextToDocxRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiExportTextToDocxRequest(val *AiExportTextToDocxRequest) *NullableAiExportTextToDocxRequest {
	return &NullableAiExportTextToDocxRequest{value: val, isSet: true}
}

func (v NullableAiExportTextToDocxRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiExportTextToDocxRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

