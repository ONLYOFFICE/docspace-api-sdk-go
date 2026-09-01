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

// checks if the DeleteBatchRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteBatchRequestDto{}

// DeleteBatchRequestDto The request parameters for deleting files.
type DeleteBatchRequestDto struct {
	// Specifies whether to return only the current operation
	ReturnSingleOperation *bool `json:"returnSingleOperation,omitempty"`
	// The list of folder IDs to be deleted.
	FolderIds []DeleteBatchRequestDtoAllOfFolderIds `json:"folderIds,omitempty"`
	// The list of file IDs to be deleted.
	FileIds []DeleteBatchRequestDtoAllOfFileIds `json:"fileIds,omitempty"`
	// Specifies whether to delete a file after the editing session is finished or not
	DeleteAfter *bool `json:"deleteAfter,omitempty"`
	// Specifies whether to move a file to the \\Trash\\ folder or delete it immediately.
	Immediately *bool `json:"immediately,omitempty"`
}

// NewDeleteBatchRequestDto instantiates a new DeleteBatchRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteBatchRequestDto() *DeleteBatchRequestDto {
	this := DeleteBatchRequestDto{}
	return &this
}

// NewDeleteBatchRequestDtoWithDefaults instantiates a new DeleteBatchRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteBatchRequestDtoWithDefaults() *DeleteBatchRequestDto {
	this := DeleteBatchRequestDto{}
	return &this
}

// GetReturnSingleOperation returns the ReturnSingleOperation field value if set, zero value otherwise.
func (o *DeleteBatchRequestDto) GetReturnSingleOperation() bool {
	if o == nil || IsNil(o.ReturnSingleOperation) {
		var ret bool
		return ret
	}
	return *o.ReturnSingleOperation
}

// GetReturnSingleOperationOk returns a tuple with the ReturnSingleOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteBatchRequestDto) GetReturnSingleOperationOk() (*bool, bool) {
	if o == nil || IsNil(o.ReturnSingleOperation) {
		return nil, false
	}
	return o.ReturnSingleOperation, true
}

// HasReturnSingleOperation returns a boolean if a field has been set.
func (o *DeleteBatchRequestDto) IsReturnSingleOperationSet() bool {
	if o != nil && !IsNil(o.ReturnSingleOperation) {
		return true
	}

	return false
}

// SetReturnSingleOperation gets a reference to the given bool and assigns it to the ReturnSingleOperation field.
func (o *DeleteBatchRequestDto) SetReturnSingleOperation(v bool) {
	o.ReturnSingleOperation = &v
}

// GetFolderIds returns the FolderIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteBatchRequestDto) GetFolderIds() []DeleteBatchRequestDtoAllOfFolderIds {
	if o == nil {
		var ret []DeleteBatchRequestDtoAllOfFolderIds
		return ret
	}
	return o.FolderIds
}

// GetFolderIdsOk returns a tuple with the FolderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteBatchRequestDto) GetFolderIdsOk() ([]DeleteBatchRequestDtoAllOfFolderIds, bool) {
	if o == nil || IsNil(o.FolderIds) {
		return nil, false
	}
	return o.FolderIds, true
}

// HasFolderIds returns a boolean if a field has been set.
func (o *DeleteBatchRequestDto) IsFolderIdsSet() bool {
	if o != nil && !IsNil(o.FolderIds) {
		return true
	}

	return false
}

// SetFolderIds gets a reference to the given []DeleteBatchRequestDtoAllOfFolderIds and assigns it to the FolderIds field.
func (o *DeleteBatchRequestDto) SetFolderIds(v []DeleteBatchRequestDtoAllOfFolderIds) {
	o.FolderIds = v
}

// GetFileIds returns the FileIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteBatchRequestDto) GetFileIds() []DeleteBatchRequestDtoAllOfFileIds {
	if o == nil {
		var ret []DeleteBatchRequestDtoAllOfFileIds
		return ret
	}
	return o.FileIds
}

// GetFileIdsOk returns a tuple with the FileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteBatchRequestDto) GetFileIdsOk() ([]DeleteBatchRequestDtoAllOfFileIds, bool) {
	if o == nil || IsNil(o.FileIds) {
		return nil, false
	}
	return o.FileIds, true
}

// HasFileIds returns a boolean if a field has been set.
func (o *DeleteBatchRequestDto) IsFileIdsSet() bool {
	if o != nil && !IsNil(o.FileIds) {
		return true
	}

	return false
}

// SetFileIds gets a reference to the given []DeleteBatchRequestDtoAllOfFileIds and assigns it to the FileIds field.
func (o *DeleteBatchRequestDto) SetFileIds(v []DeleteBatchRequestDtoAllOfFileIds) {
	o.FileIds = v
}

// GetDeleteAfter returns the DeleteAfter field value if set, zero value otherwise.
func (o *DeleteBatchRequestDto) GetDeleteAfter() bool {
	if o == nil || IsNil(o.DeleteAfter) {
		var ret bool
		return ret
	}
	return *o.DeleteAfter
}

// GetDeleteAfterOk returns a tuple with the DeleteAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteBatchRequestDto) GetDeleteAfterOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteAfter) {
		return nil, false
	}
	return o.DeleteAfter, true
}

// HasDeleteAfter returns a boolean if a field has been set.
func (o *DeleteBatchRequestDto) IsDeleteAfterSet() bool {
	if o != nil && !IsNil(o.DeleteAfter) {
		return true
	}

	return false
}

// SetDeleteAfter gets a reference to the given bool and assigns it to the DeleteAfter field.
func (o *DeleteBatchRequestDto) SetDeleteAfter(v bool) {
	o.DeleteAfter = &v
}

// GetImmediately returns the Immediately field value if set, zero value otherwise.
func (o *DeleteBatchRequestDto) GetImmediately() bool {
	if o == nil || IsNil(o.Immediately) {
		var ret bool
		return ret
	}
	return *o.Immediately
}

// GetImmediatelyOk returns a tuple with the Immediately field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteBatchRequestDto) GetImmediatelyOk() (*bool, bool) {
	if o == nil || IsNil(o.Immediately) {
		return nil, false
	}
	return o.Immediately, true
}

// HasImmediately returns a boolean if a field has been set.
func (o *DeleteBatchRequestDto) IsImmediatelySet() bool {
	if o != nil && !IsNil(o.Immediately) {
		return true
	}

	return false
}

// SetImmediately gets a reference to the given bool and assigns it to the Immediately field.
func (o *DeleteBatchRequestDto) SetImmediately(v bool) {
	o.Immediately = &v
}

func (o DeleteBatchRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteBatchRequestDto) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DeleteAfter) {
		toSerialize["deleteAfter"] = o.DeleteAfter
	}
	if !IsNil(o.Immediately) {
		toSerialize["immediately"] = o.Immediately
	}
	return toSerialize, nil
}

type NullableDeleteBatchRequestDto struct {
	value *DeleteBatchRequestDto
	isSet bool
}

func (v NullableDeleteBatchRequestDto) Get() *DeleteBatchRequestDto {
	return v.value
}

func (v *NullableDeleteBatchRequestDto) Set(val *DeleteBatchRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteBatchRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteBatchRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteBatchRequestDto(val *DeleteBatchRequestDto) *NullableDeleteBatchRequestDto {
	return &NullableDeleteBatchRequestDto{value: val, isSet: true}
}

func (v NullableDeleteBatchRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteBatchRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

