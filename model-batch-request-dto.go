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

// checks if the BatchRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchRequestDto{}

// BatchRequestDto The files and folders to move or copy, the folder they go to, and the way name clashes are settled.
type BatchRequestDto struct {
	// Which operations the answer carries: `true` returns the operation this call started and nothing else, `false`  returns every operation of the same kind that the caller has running or unread. When nothing was queued, which  happens for an empty selection, `true` falls back to the full list.
	ReturnSingleOperation *bool `json:"returnSingleOperation,omitempty"`
	// The folders to move or copy, by id. A number addresses a folder stored in the portal itself, a string  addresses a folder on a connected third-party account, and both kinds may be sent in one list.
	FolderIds []BatchRequestDtoAllOfFolderIds `json:"folderIds,omitempty"`
	// The files to move or copy, by id. A number addresses a file stored in the portal itself, a string addresses a  file on a connected third-party account, and both kinds may be sent in one list.
	FileIds []BatchRequestDtoAllOfFileIds `json:"fileIds,omitempty"`
	DestFolderId *BatchRequestDtoAllOfDestFolderId `json:"destFolderId,omitempty"`
	// What happens to an item whose name is already taken in the destination folder: `skip` leaves it where it is,  `overwrite` replaces the entry at the destination, and `duplicate` places it beside that entry under a name  with a numeric suffix. `GET api/2.0/files/fileops/move` reports which items would clash.
	ConflictResolveType *FileConflictResolveType `json:"conflictResolveType,omitempty"`
	// Whether the finished operation is still reported: `false` keeps its final record readable through  `GET api/2.0/files/fileops` until it has been read once, `true` drops the record as soon as the work is done.  It deletes nothing: a move takes the sources away in any case, and a copy always leaves them.
	DeleteAfter *bool `json:"deleteAfter,omitempty"`
	// What is taken from a listed folder: `false` moves or copies the folder itself, `true` takes only what it  contains, so its files and subfolders land in the destination and the folder is not recreated there.
	Content *bool `json:"content,omitempty"`
	// Marks every copied PDF form as a draft prepared for filling, which is how such a copy reports its filling  status in a virtual data room. Files that are not forms are left unaffected.
	ToFillOut *bool `json:"toFillOut,omitempty"`
}

// NewBatchRequestDto instantiates a new BatchRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchRequestDto() *BatchRequestDto {
	this := BatchRequestDto{}
	return &this
}

// NewBatchRequestDtoWithDefaults instantiates a new BatchRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchRequestDtoWithDefaults() *BatchRequestDto {
	this := BatchRequestDto{}
	return &this
}

// GetReturnSingleOperation returns the ReturnSingleOperation field value if set, zero value otherwise.
func (o *BatchRequestDto) GetReturnSingleOperation() bool {
	if o == nil || IsNil(o.ReturnSingleOperation) {
		var ret bool
		return ret
	}
	return *o.ReturnSingleOperation
}

// GetReturnSingleOperationOk returns a tuple with the ReturnSingleOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRequestDto) GetReturnSingleOperationOk() (*bool, bool) {
	if o == nil || IsNil(o.ReturnSingleOperation) {
		return nil, false
	}
	return o.ReturnSingleOperation, true
}

// HasReturnSingleOperation returns a boolean if a field has been set.
func (o *BatchRequestDto) IsReturnSingleOperationSet() bool {
	if o != nil && !IsNil(o.ReturnSingleOperation) {
		return true
	}

	return false
}

// SetReturnSingleOperation gets a reference to the given bool and assigns it to the ReturnSingleOperation field.
func (o *BatchRequestDto) SetReturnSingleOperation(v bool) {
	o.ReturnSingleOperation = &v
}

// GetFolderIds returns the FolderIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BatchRequestDto) GetFolderIds() []BatchRequestDtoAllOfFolderIds {
	if o == nil {
		var ret []BatchRequestDtoAllOfFolderIds
		return ret
	}
	return o.FolderIds
}

// GetFolderIdsOk returns a tuple with the FolderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BatchRequestDto) GetFolderIdsOk() ([]BatchRequestDtoAllOfFolderIds, bool) {
	if o == nil || IsNil(o.FolderIds) {
		return nil, false
	}
	return o.FolderIds, true
}

// HasFolderIds returns a boolean if a field has been set.
func (o *BatchRequestDto) IsFolderIdsSet() bool {
	if o != nil && !IsNil(o.FolderIds) {
		return true
	}

	return false
}

// SetFolderIds gets a reference to the given []BatchRequestDtoAllOfFolderIds and assigns it to the FolderIds field.
func (o *BatchRequestDto) SetFolderIds(v []BatchRequestDtoAllOfFolderIds) {
	o.FolderIds = v
}

// GetFileIds returns the FileIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BatchRequestDto) GetFileIds() []BatchRequestDtoAllOfFileIds {
	if o == nil {
		var ret []BatchRequestDtoAllOfFileIds
		return ret
	}
	return o.FileIds
}

// GetFileIdsOk returns a tuple with the FileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BatchRequestDto) GetFileIdsOk() ([]BatchRequestDtoAllOfFileIds, bool) {
	if o == nil || IsNil(o.FileIds) {
		return nil, false
	}
	return o.FileIds, true
}

// HasFileIds returns a boolean if a field has been set.
func (o *BatchRequestDto) IsFileIdsSet() bool {
	if o != nil && !IsNil(o.FileIds) {
		return true
	}

	return false
}

// SetFileIds gets a reference to the given []BatchRequestDtoAllOfFileIds and assigns it to the FileIds field.
func (o *BatchRequestDto) SetFileIds(v []BatchRequestDtoAllOfFileIds) {
	o.FileIds = v
}

// GetDestFolderId returns the DestFolderId field value if set, zero value otherwise.
func (o *BatchRequestDto) GetDestFolderId() BatchRequestDtoAllOfDestFolderId {
	if o == nil || IsNil(o.DestFolderId) {
		var ret BatchRequestDtoAllOfDestFolderId
		return ret
	}
	return *o.DestFolderId
}

// GetDestFolderIdOk returns a tuple with the DestFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRequestDto) GetDestFolderIdOk() (*BatchRequestDtoAllOfDestFolderId, bool) {
	if o == nil || IsNil(o.DestFolderId) {
		return nil, false
	}
	return o.DestFolderId, true
}

// HasDestFolderId returns a boolean if a field has been set.
func (o *BatchRequestDto) IsDestFolderIdSet() bool {
	if o != nil && !IsNil(o.DestFolderId) {
		return true
	}

	return false
}

// SetDestFolderId gets a reference to the given BatchRequestDtoAllOfDestFolderId and assigns it to the DestFolderId field.
func (o *BatchRequestDto) SetDestFolderId(v BatchRequestDtoAllOfDestFolderId) {
	o.DestFolderId = &v
}

// GetConflictResolveType returns the ConflictResolveType field value if set, zero value otherwise.
func (o *BatchRequestDto) GetConflictResolveType() FileConflictResolveType {
	if o == nil || IsNil(o.ConflictResolveType) {
		var ret FileConflictResolveType
		return ret
	}
	return *o.ConflictResolveType
}

// GetConflictResolveTypeOk returns a tuple with the ConflictResolveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRequestDto) GetConflictResolveTypeOk() (*FileConflictResolveType, bool) {
	if o == nil || IsNil(o.ConflictResolveType) {
		return nil, false
	}
	return o.ConflictResolveType, true
}

// HasConflictResolveType returns a boolean if a field has been set.
func (o *BatchRequestDto) IsConflictResolveTypeSet() bool {
	if o != nil && !IsNil(o.ConflictResolveType) {
		return true
	}

	return false
}

// SetConflictResolveType gets a reference to the given FileConflictResolveType and assigns it to the ConflictResolveType field.
func (o *BatchRequestDto) SetConflictResolveType(v FileConflictResolveType) {
	o.ConflictResolveType = &v
}

// GetDeleteAfter returns the DeleteAfter field value if set, zero value otherwise.
func (o *BatchRequestDto) GetDeleteAfter() bool {
	if o == nil || IsNil(o.DeleteAfter) {
		var ret bool
		return ret
	}
	return *o.DeleteAfter
}

// GetDeleteAfterOk returns a tuple with the DeleteAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRequestDto) GetDeleteAfterOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteAfter) {
		return nil, false
	}
	return o.DeleteAfter, true
}

// HasDeleteAfter returns a boolean if a field has been set.
func (o *BatchRequestDto) IsDeleteAfterSet() bool {
	if o != nil && !IsNil(o.DeleteAfter) {
		return true
	}

	return false
}

// SetDeleteAfter gets a reference to the given bool and assigns it to the DeleteAfter field.
func (o *BatchRequestDto) SetDeleteAfter(v bool) {
	o.DeleteAfter = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *BatchRequestDto) GetContent() bool {
	if o == nil || IsNil(o.Content) {
		var ret bool
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRequestDto) GetContentOk() (*bool, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *BatchRequestDto) IsContentSet() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given bool and assigns it to the Content field.
func (o *BatchRequestDto) SetContent(v bool) {
	o.Content = &v
}

// GetToFillOut returns the ToFillOut field value if set, zero value otherwise.
func (o *BatchRequestDto) GetToFillOut() bool {
	if o == nil || IsNil(o.ToFillOut) {
		var ret bool
		return ret
	}
	return *o.ToFillOut
}

// GetToFillOutOk returns a tuple with the ToFillOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRequestDto) GetToFillOutOk() (*bool, bool) {
	if o == nil || IsNil(o.ToFillOut) {
		return nil, false
	}
	return o.ToFillOut, true
}

// HasToFillOut returns a boolean if a field has been set.
func (o *BatchRequestDto) IsToFillOutSet() bool {
	if o != nil && !IsNil(o.ToFillOut) {
		return true
	}

	return false
}

// SetToFillOut gets a reference to the given bool and assigns it to the ToFillOut field.
func (o *BatchRequestDto) SetToFillOut(v bool) {
	o.ToFillOut = &v
}

func (o BatchRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchRequestDto) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DestFolderId) {
		toSerialize["destFolderId"] = o.DestFolderId
	}
	if !IsNil(o.ConflictResolveType) {
		toSerialize["conflictResolveType"] = o.ConflictResolveType
	}
	if !IsNil(o.DeleteAfter) {
		toSerialize["deleteAfter"] = o.DeleteAfter
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.ToFillOut) {
		toSerialize["toFillOut"] = o.ToFillOut
	}
	return toSerialize, nil
}

type NullableBatchRequestDto struct {
	value *BatchRequestDto
	isSet bool
}

func (v NullableBatchRequestDto) Get() *BatchRequestDto {
	return v.value
}

func (v *NullableBatchRequestDto) Set(val *BatchRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchRequestDto(val *BatchRequestDto) *NullableBatchRequestDto {
	return &NullableBatchRequestDto{value: val, isSet: true}
}

func (v NullableBatchRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

