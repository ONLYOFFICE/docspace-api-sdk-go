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

// checks if the TemplatesRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplatesRequestDto{}

// TemplatesRequestDto The request parameters for adding files to the template list.
type TemplatesRequestDto struct {
	// The list of file IDs.
	FileIds []int32 `json:"fileIds,omitempty"`
}

// NewTemplatesRequestDto instantiates a new TemplatesRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplatesRequestDto() *TemplatesRequestDto {
	this := TemplatesRequestDto{}
	return &this
}

// NewTemplatesRequestDtoWithDefaults instantiates a new TemplatesRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplatesRequestDtoWithDefaults() *TemplatesRequestDto {
	this := TemplatesRequestDto{}
	return &this
}

// GetFileIds returns the FileIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplatesRequestDto) GetFileIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.FileIds
}

// GetFileIdsOk returns a tuple with the FileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplatesRequestDto) GetFileIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.FileIds) {
		return nil, false
	}
	return o.FileIds, true
}

// HasFileIds returns a boolean if a field has been set.
func (o *TemplatesRequestDto) IsFileIdsSet() bool {
	if o != nil && !IsNil(o.FileIds) {
		return true
	}

	return false
}

// SetFileIds gets a reference to the given []int32 and assigns it to the FileIds field.
func (o *TemplatesRequestDto) SetFileIds(v []int32) {
	o.FileIds = v
}

func (o TemplatesRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplatesRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FileIds != nil {
		toSerialize["fileIds"] = o.FileIds
	}
	return toSerialize, nil
}

type NullableTemplatesRequestDto struct {
	value *TemplatesRequestDto
	isSet bool
}

func (v NullableTemplatesRequestDto) Get() *TemplatesRequestDto {
	return v.value
}

func (v *NullableTemplatesRequestDto) Set(val *TemplatesRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplatesRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplatesRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplatesRequestDto(val *TemplatesRequestDto) *NullableTemplatesRequestDto {
	return &NullableTemplatesRequestDto{value: val, isSet: true}
}

func (v NullableTemplatesRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplatesRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

