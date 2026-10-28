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

// checks if the DeleteVersionBatchRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteVersionBatchRequestDto{}

// DeleteVersionBatchRequestDto The file whose versions are deleted, and the versions to delete.
type DeleteVersionBatchRequestDto struct {
	// Which operations the answer carries: `true` returns the operation this call started and nothing else, `false`  returns every operation of the same kind that the caller has running or unread. When nothing was queued, which  happens for an empty selection, `true` falls back to the full list.
	ReturnSingleOperation *bool `json:"returnSingleOperation,omitempty"`
	// Whether the finished operation is still reported: `false` keeps its final record readable through  `GET api/2.0/files/fileops` until it has been read once, `true` drops the record as soon as the work is done.  It does not postpone the deletion and does not delete anything of its own.
	DeleteAfter *bool `json:"deleteAfter,omitempty"`
	// The file whose history the versions are taken from; only files stored in the portal itself are addressed here.
	FileId int32 `json:"fileId"`
	// The version numbers to remove, as reported by `GET api/2.0/files/file/{fileId}/history`. At least one number  has to be sent: an empty list removes the file itself instead of one of its versions. The number of the  current version is refused outright, while a number that no longer exists is passed over without a complaint.
	Versions []int32 `json:"versions"`
}

type _DeleteVersionBatchRequestDto DeleteVersionBatchRequestDto

// NewDeleteVersionBatchRequestDto instantiates a new DeleteVersionBatchRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteVersionBatchRequestDto(fileId int32, versions []int32) *DeleteVersionBatchRequestDto {
	this := DeleteVersionBatchRequestDto{}
	this.FileId = fileId
	this.Versions = versions
	return &this
}

// NewDeleteVersionBatchRequestDtoWithDefaults instantiates a new DeleteVersionBatchRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteVersionBatchRequestDtoWithDefaults() *DeleteVersionBatchRequestDto {
	this := DeleteVersionBatchRequestDto{}
	return &this
}

// GetReturnSingleOperation returns the ReturnSingleOperation field value if set, zero value otherwise.
func (o *DeleteVersionBatchRequestDto) GetReturnSingleOperation() bool {
	if o == nil || IsNil(o.ReturnSingleOperation) {
		var ret bool
		return ret
	}
	return *o.ReturnSingleOperation
}

// GetReturnSingleOperationOk returns a tuple with the ReturnSingleOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteVersionBatchRequestDto) GetReturnSingleOperationOk() (*bool, bool) {
	if o == nil || IsNil(o.ReturnSingleOperation) {
		return nil, false
	}
	return o.ReturnSingleOperation, true
}

// HasReturnSingleOperation returns a boolean if a field has been set.
func (o *DeleteVersionBatchRequestDto) IsReturnSingleOperationSet() bool {
	if o != nil && !IsNil(o.ReturnSingleOperation) {
		return true
	}

	return false
}

// SetReturnSingleOperation gets a reference to the given bool and assigns it to the ReturnSingleOperation field.
func (o *DeleteVersionBatchRequestDto) SetReturnSingleOperation(v bool) {
	o.ReturnSingleOperation = &v
}

// GetDeleteAfter returns the DeleteAfter field value if set, zero value otherwise.
func (o *DeleteVersionBatchRequestDto) GetDeleteAfter() bool {
	if o == nil || IsNil(o.DeleteAfter) {
		var ret bool
		return ret
	}
	return *o.DeleteAfter
}

// GetDeleteAfterOk returns a tuple with the DeleteAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteVersionBatchRequestDto) GetDeleteAfterOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteAfter) {
		return nil, false
	}
	return o.DeleteAfter, true
}

// HasDeleteAfter returns a boolean if a field has been set.
func (o *DeleteVersionBatchRequestDto) IsDeleteAfterSet() bool {
	if o != nil && !IsNil(o.DeleteAfter) {
		return true
	}

	return false
}

// SetDeleteAfter gets a reference to the given bool and assigns it to the DeleteAfter field.
func (o *DeleteVersionBatchRequestDto) SetDeleteAfter(v bool) {
	o.DeleteAfter = &v
}

// GetFileId returns the FileId field value
func (o *DeleteVersionBatchRequestDto) GetFileId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value
// and a boolean to check if the value has been set.
func (o *DeleteVersionBatchRequestDto) GetFileIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileId, true
}

// SetFileId sets field value
func (o *DeleteVersionBatchRequestDto) SetFileId(v int32) {
	o.FileId = v
}

// GetVersions returns the Versions field value
// If the value is explicit nil, the zero value for []int32 will be returned
func (o *DeleteVersionBatchRequestDto) GetVersions() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteVersionBatchRequestDto) GetVersionsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// SetVersions sets field value
func (o *DeleteVersionBatchRequestDto) SetVersions(v []int32) {
	o.Versions = v
}

func (o DeleteVersionBatchRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteVersionBatchRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReturnSingleOperation) {
		toSerialize["returnSingleOperation"] = o.ReturnSingleOperation
	}
	if !IsNil(o.DeleteAfter) {
		toSerialize["deleteAfter"] = o.DeleteAfter
	}
	toSerialize["fileId"] = o.FileId
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

func (o *DeleteVersionBatchRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fileId",
		"versions",
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

	varDeleteVersionBatchRequestDto := _DeleteVersionBatchRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varDeleteVersionBatchRequestDto)

	if err != nil {
		return err
	}

	*o = DeleteVersionBatchRequestDto(varDeleteVersionBatchRequestDto)

	return err
}

type NullableDeleteVersionBatchRequestDto struct {
	value *DeleteVersionBatchRequestDto
	isSet bool
}

func (v NullableDeleteVersionBatchRequestDto) Get() *DeleteVersionBatchRequestDto {
	return v.value
}

func (v *NullableDeleteVersionBatchRequestDto) Set(val *DeleteVersionBatchRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteVersionBatchRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteVersionBatchRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteVersionBatchRequestDto(val *DeleteVersionBatchRequestDto) *NullableDeleteVersionBatchRequestDto {
	return &NullableDeleteVersionBatchRequestDto{value: val, isSet: true}
}

func (v NullableDeleteVersionBatchRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteVersionBatchRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

