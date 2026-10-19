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

// checks if the ThirdPartySaveAsPdfRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartySaveAsPdfRequest{}

// ThirdPartySaveAsPdfRequest The place and the name the PDF copy of a file is stored under.
type ThirdPartySaveAsPdfRequest struct {
	// The folder the PDF is created in; the caller has to be allowed to create files there.
	FolderId NullableString `json:"folderId"`
	// The name of the PDF, without an extension - `.pdf` is appended. Left empty, the name of the source file is  reused with its extension replaced.
	Title NullableString `json:"title"`
}

type _ThirdPartySaveAsPdfRequest ThirdPartySaveAsPdfRequest

// NewThirdPartySaveAsPdfRequest instantiates a new ThirdPartySaveAsPdfRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartySaveAsPdfRequest(folderId NullableString, title NullableString) *ThirdPartySaveAsPdfRequest {
	this := ThirdPartySaveAsPdfRequest{}
	this.FolderId = folderId
	this.Title = title
	return &this
}

// NewThirdPartySaveAsPdfRequestWithDefaults instantiates a new ThirdPartySaveAsPdfRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartySaveAsPdfRequestWithDefaults() *ThirdPartySaveAsPdfRequest {
	this := ThirdPartySaveAsPdfRequest{}
	return &this
}

// GetFolderId returns the FolderId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ThirdPartySaveAsPdfRequest) GetFolderId() string {
	if o == nil || o.FolderId.Get() == nil {
		var ret string
		return ret
	}

	return *o.FolderId.Get()
}

// GetFolderIdOk returns a tuple with the FolderId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartySaveAsPdfRequest) GetFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderId.Get(), o.FolderId.IsSet()
}

// SetFolderId sets field value
func (o *ThirdPartySaveAsPdfRequest) SetFolderId(v string) {
	o.FolderId.Set(&v)
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ThirdPartySaveAsPdfRequest) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartySaveAsPdfRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *ThirdPartySaveAsPdfRequest) SetTitle(v string) {
	o.Title.Set(&v)
}

func (o ThirdPartySaveAsPdfRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartySaveAsPdfRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["folderId"] = o.FolderId.Get()
	toSerialize["title"] = o.Title.Get()
	return toSerialize, nil
}

func (o *ThirdPartySaveAsPdfRequest) UnmarshalJSON(data []byte) (err error) {
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

	varThirdPartySaveAsPdfRequest := _ThirdPartySaveAsPdfRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varThirdPartySaveAsPdfRequest)

	if err != nil {
		return err
	}

	*o = ThirdPartySaveAsPdfRequest(varThirdPartySaveAsPdfRequest)

	return err
}

type NullableThirdPartySaveAsPdfRequest struct {
	value *ThirdPartySaveAsPdfRequest
	isSet bool
}

func (v NullableThirdPartySaveAsPdfRequest) Get() *ThirdPartySaveAsPdfRequest {
	return v.value
}

func (v *NullableThirdPartySaveAsPdfRequest) Set(val *ThirdPartySaveAsPdfRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartySaveAsPdfRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartySaveAsPdfRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartySaveAsPdfRequest(val *ThirdPartySaveAsPdfRequest) *NullableThirdPartySaveAsPdfRequest {
	return &NullableThirdPartySaveAsPdfRequest{value: val, isSet: true}
}

func (v NullableThirdPartySaveAsPdfRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartySaveAsPdfRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

