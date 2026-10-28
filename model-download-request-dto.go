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

// checks if the DownloadRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DownloadRequestDto{}

// DownloadRequestDto The files and folders to pack into one archive, together with the formats they are converted to.
type DownloadRequestDto struct {
	// Which operations the answer carries: `true` returns the operation this call started and nothing else, `false`  returns every operation of the same kind that the caller has running or unread. When nothing was queued, which  happens for an empty selection, `true` falls back to the full list.
	ReturnSingleOperation *bool `json:"returnSingleOperation,omitempty"`
	// The folders to pack, by id; everything inside them that the caller may read goes into the archive. A number  addresses a folder stored in the portal itself, a string addresses a folder on a connected third-party  account, and both kinds may be sent in one list.
	FolderIds []DownloadRequestDtoAllOfFolderIds `json:"folderIds,omitempty"`
	// The files to pack as they are, by id, without conversion. A number addresses a file stored in the portal  itself, a string addresses a file on a connected third-party account, and both kinds may be sent in one list.
	FileIds []DownloadRequestDtoAllOfFileIds `json:"fileIds,omitempty"`
	// The files to convert before they are packed, each named together with the format it is converted to. A file  listed here does not have to be repeated in `fileIds`.
	FileConvertIds []DownloadRequestItemDto `json:"fileConvertIds,omitempty"`
}

// NewDownloadRequestDto instantiates a new DownloadRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDownloadRequestDto() *DownloadRequestDto {
	this := DownloadRequestDto{}
	return &this
}

// NewDownloadRequestDtoWithDefaults instantiates a new DownloadRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDownloadRequestDtoWithDefaults() *DownloadRequestDto {
	this := DownloadRequestDto{}
	return &this
}

// GetReturnSingleOperation returns the ReturnSingleOperation field value if set, zero value otherwise.
func (o *DownloadRequestDto) GetReturnSingleOperation() bool {
	if o == nil || IsNil(o.ReturnSingleOperation) {
		var ret bool
		return ret
	}
	return *o.ReturnSingleOperation
}

// GetReturnSingleOperationOk returns a tuple with the ReturnSingleOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadRequestDto) GetReturnSingleOperationOk() (*bool, bool) {
	if o == nil || IsNil(o.ReturnSingleOperation) {
		return nil, false
	}
	return o.ReturnSingleOperation, true
}

// HasReturnSingleOperation returns a boolean if a field has been set.
func (o *DownloadRequestDto) IsReturnSingleOperationSet() bool {
	if o != nil && !IsNil(o.ReturnSingleOperation) {
		return true
	}

	return false
}

// SetReturnSingleOperation gets a reference to the given bool and assigns it to the ReturnSingleOperation field.
func (o *DownloadRequestDto) SetReturnSingleOperation(v bool) {
	o.ReturnSingleOperation = &v
}

// GetFolderIds returns the FolderIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadRequestDto) GetFolderIds() []DownloadRequestDtoAllOfFolderIds {
	if o == nil {
		var ret []DownloadRequestDtoAllOfFolderIds
		return ret
	}
	return o.FolderIds
}

// GetFolderIdsOk returns a tuple with the FolderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadRequestDto) GetFolderIdsOk() ([]DownloadRequestDtoAllOfFolderIds, bool) {
	if o == nil || IsNil(o.FolderIds) {
		return nil, false
	}
	return o.FolderIds, true
}

// HasFolderIds returns a boolean if a field has been set.
func (o *DownloadRequestDto) IsFolderIdsSet() bool {
	if o != nil && !IsNil(o.FolderIds) {
		return true
	}

	return false
}

// SetFolderIds gets a reference to the given []DownloadRequestDtoAllOfFolderIds and assigns it to the FolderIds field.
func (o *DownloadRequestDto) SetFolderIds(v []DownloadRequestDtoAllOfFolderIds) {
	o.FolderIds = v
}

// GetFileIds returns the FileIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadRequestDto) GetFileIds() []DownloadRequestDtoAllOfFileIds {
	if o == nil {
		var ret []DownloadRequestDtoAllOfFileIds
		return ret
	}
	return o.FileIds
}

// GetFileIdsOk returns a tuple with the FileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadRequestDto) GetFileIdsOk() ([]DownloadRequestDtoAllOfFileIds, bool) {
	if o == nil || IsNil(o.FileIds) {
		return nil, false
	}
	return o.FileIds, true
}

// HasFileIds returns a boolean if a field has been set.
func (o *DownloadRequestDto) IsFileIdsSet() bool {
	if o != nil && !IsNil(o.FileIds) {
		return true
	}

	return false
}

// SetFileIds gets a reference to the given []DownloadRequestDtoAllOfFileIds and assigns it to the FileIds field.
func (o *DownloadRequestDto) SetFileIds(v []DownloadRequestDtoAllOfFileIds) {
	o.FileIds = v
}

// GetFileConvertIds returns the FileConvertIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DownloadRequestDto) GetFileConvertIds() []DownloadRequestItemDto {
	if o == nil {
		var ret []DownloadRequestItemDto
		return ret
	}
	return o.FileConvertIds
}

// GetFileConvertIdsOk returns a tuple with the FileConvertIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DownloadRequestDto) GetFileConvertIdsOk() ([]DownloadRequestItemDto, bool) {
	if o == nil || IsNil(o.FileConvertIds) {
		return nil, false
	}
	return o.FileConvertIds, true
}

// HasFileConvertIds returns a boolean if a field has been set.
func (o *DownloadRequestDto) IsFileConvertIdsSet() bool {
	if o != nil && !IsNil(o.FileConvertIds) {
		return true
	}

	return false
}

// SetFileConvertIds gets a reference to the given []DownloadRequestItemDto and assigns it to the FileConvertIds field.
func (o *DownloadRequestDto) SetFileConvertIds(v []DownloadRequestItemDto) {
	o.FileConvertIds = v
}

func (o DownloadRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DownloadRequestDto) ToMap() (map[string]interface{}, error) {
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
	if o.FileConvertIds != nil {
		toSerialize["fileConvertIds"] = o.FileConvertIds
	}
	return toSerialize, nil
}

type NullableDownloadRequestDto struct {
	value *DownloadRequestDto
	isSet bool
}

func (v NullableDownloadRequestDto) Get() *DownloadRequestDto {
	return v.value
}

func (v *NullableDownloadRequestDto) Set(val *DownloadRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDownloadRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDownloadRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownloadRequestDto(val *DownloadRequestDto) *NullableDownloadRequestDto {
	return &NullableDownloadRequestDto{value: val, isSet: true}
}

func (v NullableDownloadRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownloadRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

