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

// checks if the ThirdPartySaveAsPdf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartySaveAsPdf{}

// ThirdPartySaveAsPdf The place and the name the PDF copy of a file is stored under.
type ThirdPartySaveAsPdf struct {
	// The folder the PDF is created in; the caller has to be allowed to create files there.
	FolderId NullableString `json:"folderId"`
	// The name of the PDF, without an extension - `.pdf` is appended. Left empty, the name of the source file is  reused with its extension replaced.
	Title NullableString `json:"title"`
}

type _ThirdPartySaveAsPdf ThirdPartySaveAsPdf

// NewThirdPartySaveAsPdf instantiates a new ThirdPartySaveAsPdf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartySaveAsPdf(folderId NullableString, title NullableString) *ThirdPartySaveAsPdf {
	this := ThirdPartySaveAsPdf{}
	this.FolderId = folderId
	this.Title = title
	return &this
}

// NewThirdPartySaveAsPdfWithDefaults instantiates a new ThirdPartySaveAsPdf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartySaveAsPdfWithDefaults() *ThirdPartySaveAsPdf {
	this := ThirdPartySaveAsPdf{}
	return &this
}

// GetFolderId returns the FolderId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ThirdPartySaveAsPdf) GetFolderId() string {
	if o == nil || o.FolderId.Get() == nil {
		var ret string
		return ret
	}

	return *o.FolderId.Get()
}

// GetFolderIdOk returns a tuple with the FolderId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartySaveAsPdf) GetFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderId.Get(), o.FolderId.IsSet()
}

// SetFolderId sets field value
func (o *ThirdPartySaveAsPdf) SetFolderId(v string) {
	o.FolderId.Set(&v)
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ThirdPartySaveAsPdf) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartySaveAsPdf) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *ThirdPartySaveAsPdf) SetTitle(v string) {
	o.Title.Set(&v)
}

func (o ThirdPartySaveAsPdf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartySaveAsPdf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["folderId"] = o.FolderId.Get()
	toSerialize["title"] = o.Title.Get()
	return toSerialize, nil
}

func (o *ThirdPartySaveAsPdf) UnmarshalJSON(data []byte) (err error) {
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

	varThirdPartySaveAsPdf := _ThirdPartySaveAsPdf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varThirdPartySaveAsPdf)

	if err != nil {
		return err
	}

	*o = ThirdPartySaveAsPdf(varThirdPartySaveAsPdf)

	return err
}

type NullableThirdPartySaveAsPdf struct {
	value *ThirdPartySaveAsPdf
	isSet bool
}

func (v NullableThirdPartySaveAsPdf) Get() *ThirdPartySaveAsPdf {
	return v.value
}

func (v *NullableThirdPartySaveAsPdf) Set(val *ThirdPartySaveAsPdf) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartySaveAsPdf) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartySaveAsPdf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartySaveAsPdf(val *ThirdPartySaveAsPdf) *NullableThirdPartySaveAsPdf {
	return &NullableThirdPartySaveAsPdf{value: val, isSet: true}
}

func (v NullableThirdPartySaveAsPdf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartySaveAsPdf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

