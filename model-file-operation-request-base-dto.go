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

// checks if the FileOperationRequestBaseDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileOperationRequestBaseDto{}

// FileOperationRequestBaseDto The base operation request parameters.
type FileOperationRequestBaseDto struct {
	// Specifies whether to return only the current operation
	ReturnSingleOperation *bool `json:"returnSingleOperation,omitempty"`
}

// NewFileOperationRequestBaseDto instantiates a new FileOperationRequestBaseDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileOperationRequestBaseDto() *FileOperationRequestBaseDto {
	this := FileOperationRequestBaseDto{}
	return &this
}

// NewFileOperationRequestBaseDtoWithDefaults instantiates a new FileOperationRequestBaseDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileOperationRequestBaseDtoWithDefaults() *FileOperationRequestBaseDto {
	this := FileOperationRequestBaseDto{}
	return &this
}

// GetReturnSingleOperation returns the ReturnSingleOperation field value if set, zero value otherwise.
func (o *FileOperationRequestBaseDto) GetReturnSingleOperation() bool {
	if o == nil || IsNil(o.ReturnSingleOperation) {
		var ret bool
		return ret
	}
	return *o.ReturnSingleOperation
}

// GetReturnSingleOperationOk returns a tuple with the ReturnSingleOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileOperationRequestBaseDto) GetReturnSingleOperationOk() (*bool, bool) {
	if o == nil || IsNil(o.ReturnSingleOperation) {
		return nil, false
	}
	return o.ReturnSingleOperation, true
}

// HasReturnSingleOperation returns a boolean if a field has been set.
func (o *FileOperationRequestBaseDto) IsReturnSingleOperationSet() bool {
	if o != nil && !IsNil(o.ReturnSingleOperation) {
		return true
	}

	return false
}

// SetReturnSingleOperation gets a reference to the given bool and assigns it to the ReturnSingleOperation field.
func (o *FileOperationRequestBaseDto) SetReturnSingleOperation(v bool) {
	o.ReturnSingleOperation = &v
}

func (o FileOperationRequestBaseDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileOperationRequestBaseDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReturnSingleOperation) {
		toSerialize["returnSingleOperation"] = o.ReturnSingleOperation
	}
	return toSerialize, nil
}

type NullableFileOperationRequestBaseDto struct {
	value *FileOperationRequestBaseDto
	isSet bool
}

func (v NullableFileOperationRequestBaseDto) Get() *FileOperationRequestBaseDto {
	return v.value
}

func (v *NullableFileOperationRequestBaseDto) Set(val *FileOperationRequestBaseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFileOperationRequestBaseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFileOperationRequestBaseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileOperationRequestBaseDto(val *FileOperationRequestBaseDto) *NullableFileOperationRequestBaseDto {
	return &NullableFileOperationRequestBaseDto{value: val, isSet: true}
}

func (v NullableFileOperationRequestBaseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileOperationRequestBaseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

