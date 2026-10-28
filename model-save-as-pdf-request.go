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

// checks if the SaveAsPdfRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaveAsPdfRequest{}

// SaveAsPdfRequest The place and the name the PDF copy of a file is stored under.
type SaveAsPdfRequest struct {
	// The folder the PDF is created in; the caller has to be allowed to create files there.
	FolderId int32 `json:"folderId"`
	// The name of the PDF, without an extension - `.pdf` is appended. Left empty, the name of the source file is  reused with its extension replaced.
	Title NullableString `json:"title"`
}

type _SaveAsPdfRequest SaveAsPdfRequest

// NewSaveAsPdfRequest instantiates a new SaveAsPdfRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaveAsPdfRequest(folderId int32, title NullableString) *SaveAsPdfRequest {
	this := SaveAsPdfRequest{}
	this.FolderId = folderId
	this.Title = title
	return &this
}

// NewSaveAsPdfRequestWithDefaults instantiates a new SaveAsPdfRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaveAsPdfRequestWithDefaults() *SaveAsPdfRequest {
	this := SaveAsPdfRequest{}
	return &this
}

// GetFolderId returns the FolderId field value
func (o *SaveAsPdfRequest) GetFolderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value
// and a boolean to check if the value has been set.
func (o *SaveAsPdfRequest) GetFolderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FolderId, true
}

// SetFolderId sets field value
func (o *SaveAsPdfRequest) SetFolderId(v int32) {
	o.FolderId = v
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SaveAsPdfRequest) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SaveAsPdfRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *SaveAsPdfRequest) SetTitle(v string) {
	o.Title.Set(&v)
}

func (o SaveAsPdfRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaveAsPdfRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["folderId"] = o.FolderId
	toSerialize["title"] = o.Title.Get()
	return toSerialize, nil
}

func (o *SaveAsPdfRequest) UnmarshalJSON(data []byte) (err error) {
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

	varSaveAsPdfRequest := _SaveAsPdfRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSaveAsPdfRequest)

	if err != nil {
		return err
	}

	*o = SaveAsPdfRequest(varSaveAsPdfRequest)

	return err
}

type NullableSaveAsPdfRequest struct {
	value *SaveAsPdfRequest
	isSet bool
}

func (v NullableSaveAsPdfRequest) Get() *SaveAsPdfRequest {
	return v.value
}

func (v *NullableSaveAsPdfRequest) Set(val *SaveAsPdfRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveAsPdfRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveAsPdfRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveAsPdfRequest(val *SaveAsPdfRequest) *NullableSaveAsPdfRequest {
	return &NullableSaveAsPdfRequest{value: val, isSet: true}
}

func (v NullableSaveAsPdfRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveAsPdfRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

