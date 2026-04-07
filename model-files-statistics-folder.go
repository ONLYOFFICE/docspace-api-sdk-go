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

// checks if the FilesStatisticsFolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesStatisticsFolder{}

// FilesStatisticsFolder The file statictics folder parameters.
type FilesStatisticsFolder struct {
	// The folder title.
	Title NullableString `json:"title,omitempty"`
	// The used space in the folder.
	UsedSpace *int64 `json:"usedSpace,omitempty"`
}

// NewFilesStatisticsFolder instantiates a new FilesStatisticsFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesStatisticsFolder() *FilesStatisticsFolder {
	this := FilesStatisticsFolder{}
	return &this
}

// NewFilesStatisticsFolderWithDefaults instantiates a new FilesStatisticsFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesStatisticsFolderWithDefaults() *FilesStatisticsFolder {
	this := FilesStatisticsFolder{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FilesStatisticsFolder) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesStatisticsFolder) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *FilesStatisticsFolder) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *FilesStatisticsFolder) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *FilesStatisticsFolder) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *FilesStatisticsFolder) UnsetTitle() {
	o.Title.Unset()
}

// GetUsedSpace returns the UsedSpace field value if set, zero value otherwise.
func (o *FilesStatisticsFolder) GetUsedSpace() int64 {
	if o == nil || IsNil(o.UsedSpace) {
		var ret int64
		return ret
	}
	return *o.UsedSpace
}

// GetUsedSpaceOk returns a tuple with the UsedSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesStatisticsFolder) GetUsedSpaceOk() (*int64, bool) {
	if o == nil || IsNil(o.UsedSpace) {
		return nil, false
	}
	return o.UsedSpace, true
}

// HasUsedSpace returns a boolean if a field has been set.
func (o *FilesStatisticsFolder) IsUsedSpaceSet() bool {
	if o != nil && !IsNil(o.UsedSpace) {
		return true
	}

	return false
}

// SetUsedSpace gets a reference to the given int64 and assigns it to the UsedSpace field.
func (o *FilesStatisticsFolder) SetUsedSpace(v int64) {
	o.UsedSpace = &v
}

func (o FilesStatisticsFolder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesStatisticsFolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if !IsNil(o.UsedSpace) {
		toSerialize["usedSpace"] = o.UsedSpace
	}
	return toSerialize, nil
}

type NullableFilesStatisticsFolder struct {
	value *FilesStatisticsFolder
	isSet bool
}

func (v NullableFilesStatisticsFolder) Get() *FilesStatisticsFolder {
	return v.value
}

func (v *NullableFilesStatisticsFolder) Set(val *FilesStatisticsFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesStatisticsFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesStatisticsFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesStatisticsFolder(val *FilesStatisticsFolder) *NullableFilesStatisticsFolder {
	return &NullableFilesStatisticsFolder{value: val, isSet: true}
}

func (v NullableFilesStatisticsFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesStatisticsFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

