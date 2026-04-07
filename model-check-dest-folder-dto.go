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

// checks if the CheckDestFolderDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckDestFolderDto{}

// CheckDestFolderDto The result of checking whether files can be moved or copied to the specified folder.
type CheckDestFolderDto struct {
	Result *CheckDestFolderResult `json:"result,omitempty"`
	// The list of files in the destination folder.
	Files []FileEntryBaseDto `json:"files,omitempty"`
}

// NewCheckDestFolderDto instantiates a new CheckDestFolderDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckDestFolderDto() *CheckDestFolderDto {
	this := CheckDestFolderDto{}
	return &this
}

// NewCheckDestFolderDtoWithDefaults instantiates a new CheckDestFolderDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckDestFolderDtoWithDefaults() *CheckDestFolderDto {
	this := CheckDestFolderDto{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CheckDestFolderDto) GetResult() CheckDestFolderResult {
	if o == nil || IsNil(o.Result) {
		var ret CheckDestFolderResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckDestFolderDto) GetResultOk() (*CheckDestFolderResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CheckDestFolderDto) IsResultSet() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given CheckDestFolderResult and assigns it to the Result field.
func (o *CheckDestFolderDto) SetResult(v CheckDestFolderResult) {
	o.Result = &v
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckDestFolderDto) GetFiles() []FileEntryBaseDto {
	if o == nil {
		var ret []FileEntryBaseDto
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckDestFolderDto) GetFilesOk() ([]FileEntryBaseDto, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *CheckDestFolderDto) IsFilesSet() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FileEntryBaseDto and assigns it to the Files field.
func (o *CheckDestFolderDto) SetFiles(v []FileEntryBaseDto) {
	o.Files = v
}

func (o CheckDestFolderDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckDestFolderDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	return toSerialize, nil
}

type NullableCheckDestFolderDto struct {
	value *CheckDestFolderDto
	isSet bool
}

func (v NullableCheckDestFolderDto) Get() *CheckDestFolderDto {
	return v.value
}

func (v *NullableCheckDestFolderDto) Set(val *CheckDestFolderDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckDestFolderDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckDestFolderDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckDestFolderDto(val *CheckDestFolderDto) *NullableCheckDestFolderDto {
	return &NullableCheckDestFolderDto{value: val, isSet: true}
}

func (v NullableCheckDestFolderDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckDestFolderDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

