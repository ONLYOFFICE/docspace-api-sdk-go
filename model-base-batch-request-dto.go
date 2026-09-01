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

// checks if the BaseBatchRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseBatchRequestDto{}

// BaseBatchRequestDto The base batch request parameters.
type BaseBatchRequestDto struct {
	// Specifies whether to return only the current operation
	ReturnSingleOperation *bool `json:"returnSingleOperation,omitempty"`
	// The list of folder IDs of the base batch request.
	FolderIds []BaseBatchRequestDtoAllOfFolderIds `json:"folderIds,omitempty"`
	// The list of file IDs of the base batch request.
	FileIds []BaseBatchRequestDtoAllOfFileIds `json:"fileIds,omitempty"`
}

// NewBaseBatchRequestDto instantiates a new BaseBatchRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseBatchRequestDto() *BaseBatchRequestDto {
	this := BaseBatchRequestDto{}
	return &this
}

// NewBaseBatchRequestDtoWithDefaults instantiates a new BaseBatchRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseBatchRequestDtoWithDefaults() *BaseBatchRequestDto {
	this := BaseBatchRequestDto{}
	return &this
}

// GetReturnSingleOperation returns the ReturnSingleOperation field value if set, zero value otherwise.
func (o *BaseBatchRequestDto) GetReturnSingleOperation() bool {
	if o == nil || IsNil(o.ReturnSingleOperation) {
		var ret bool
		return ret
	}
	return *o.ReturnSingleOperation
}

// GetReturnSingleOperationOk returns a tuple with the ReturnSingleOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseBatchRequestDto) GetReturnSingleOperationOk() (*bool, bool) {
	if o == nil || IsNil(o.ReturnSingleOperation) {
		return nil, false
	}
	return o.ReturnSingleOperation, true
}

// HasReturnSingleOperation returns a boolean if a field has been set.
func (o *BaseBatchRequestDto) IsReturnSingleOperationSet() bool {
	if o != nil && !IsNil(o.ReturnSingleOperation) {
		return true
	}

	return false
}

// SetReturnSingleOperation gets a reference to the given bool and assigns it to the ReturnSingleOperation field.
func (o *BaseBatchRequestDto) SetReturnSingleOperation(v bool) {
	o.ReturnSingleOperation = &v
}

// GetFolderIds returns the FolderIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseBatchRequestDto) GetFolderIds() []BaseBatchRequestDtoAllOfFolderIds {
	if o == nil {
		var ret []BaseBatchRequestDtoAllOfFolderIds
		return ret
	}
	return o.FolderIds
}

// GetFolderIdsOk returns a tuple with the FolderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseBatchRequestDto) GetFolderIdsOk() ([]BaseBatchRequestDtoAllOfFolderIds, bool) {
	if o == nil || IsNil(o.FolderIds) {
		return nil, false
	}
	return o.FolderIds, true
}

// HasFolderIds returns a boolean if a field has been set.
func (o *BaseBatchRequestDto) IsFolderIdsSet() bool {
	if o != nil && !IsNil(o.FolderIds) {
		return true
	}

	return false
}

// SetFolderIds gets a reference to the given []BaseBatchRequestDtoAllOfFolderIds and assigns it to the FolderIds field.
func (o *BaseBatchRequestDto) SetFolderIds(v []BaseBatchRequestDtoAllOfFolderIds) {
	o.FolderIds = v
}

// GetFileIds returns the FileIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseBatchRequestDto) GetFileIds() []BaseBatchRequestDtoAllOfFileIds {
	if o == nil {
		var ret []BaseBatchRequestDtoAllOfFileIds
		return ret
	}
	return o.FileIds
}

// GetFileIdsOk returns a tuple with the FileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseBatchRequestDto) GetFileIdsOk() ([]BaseBatchRequestDtoAllOfFileIds, bool) {
	if o == nil || IsNil(o.FileIds) {
		return nil, false
	}
	return o.FileIds, true
}

// HasFileIds returns a boolean if a field has been set.
func (o *BaseBatchRequestDto) IsFileIdsSet() bool {
	if o != nil && !IsNil(o.FileIds) {
		return true
	}

	return false
}

// SetFileIds gets a reference to the given []BaseBatchRequestDtoAllOfFileIds and assigns it to the FileIds field.
func (o *BaseBatchRequestDto) SetFileIds(v []BaseBatchRequestDtoAllOfFileIds) {
	o.FileIds = v
}

func (o BaseBatchRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseBatchRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReturnSingleOperation) {
		toSerialize["returnSingleOperation"] = o.ReturnSingleOperation
	}
	if o.FolderIds != nil {
		toSerialize["folderIds"] = o.FolderIds
	}
	if o.FileIds != nil {
		toSerialize["fileIds"] = o.FileIds
	}
	return toSerialize, nil
}

type NullableBaseBatchRequestDto struct {
	value *BaseBatchRequestDto
	isSet bool
}

func (v NullableBaseBatchRequestDto) Get() *BaseBatchRequestDto {
	return v.value
}

func (v *NullableBaseBatchRequestDto) Set(val *BaseBatchRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseBatchRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseBatchRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseBatchRequestDto(val *BaseBatchRequestDto) *NullableBaseBatchRequestDto {
	return &NullableBaseBatchRequestDto{value: val, isSet: true}
}

func (v NullableBaseBatchRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseBatchRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

