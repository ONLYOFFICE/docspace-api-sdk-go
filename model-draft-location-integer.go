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

// checks if the DraftLocationInteger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DraftLocationInteger{}

// DraftLocationInteger The file draft parameters.
type DraftLocationInteger struct {
	// The InProcess folder ID of the draft.
	FolderId *int32 `json:"folderId,omitempty"`
	// The InProcess folder title of the draft.
	FolderTitle NullableString `json:"folderTitle,omitempty"`
	// The draft ID.
	FileId *int32 `json:"fileId,omitempty"`
	// The draft title.
	FileTitle NullableString `json:"fileTitle,omitempty"`
}

// NewDraftLocationInteger instantiates a new DraftLocationInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDraftLocationInteger() *DraftLocationInteger {
	this := DraftLocationInteger{}
	return &this
}

// NewDraftLocationIntegerWithDefaults instantiates a new DraftLocationInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDraftLocationIntegerWithDefaults() *DraftLocationInteger {
	this := DraftLocationInteger{}
	return &this
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *DraftLocationInteger) GetFolderId() int32 {
	if o == nil || IsNil(o.FolderId) {
		var ret int32
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftLocationInteger) GetFolderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FolderId) {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *DraftLocationInteger) IsFolderIdSet() bool {
	if o != nil && !IsNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given int32 and assigns it to the FolderId field.
func (o *DraftLocationInteger) SetFolderId(v int32) {
	o.FolderId = &v
}

// GetFolderTitle returns the FolderTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DraftLocationInteger) GetFolderTitle() string {
	if o == nil || IsNil(o.FolderTitle.Get()) {
		var ret string
		return ret
	}
	return *o.FolderTitle.Get()
}

// GetFolderTitleOk returns a tuple with the FolderTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DraftLocationInteger) GetFolderTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderTitle.Get(), o.FolderTitle.IsSet()
}

// HasFolderTitle returns a boolean if a field has been set.
func (o *DraftLocationInteger) IsFolderTitleSet() bool {
	if o != nil && o.FolderTitle.IsSet() {
		return true
	}

	return false
}

// SetFolderTitle gets a reference to the given NullableString and assigns it to the FolderTitle field.
func (o *DraftLocationInteger) SetFolderTitle(v string) {
	o.FolderTitle.Set(&v)
}
// SetFolderTitleNil sets the value for FolderTitle to be an explicit nil
func (o *DraftLocationInteger) SetFolderTitleNil() {
	o.FolderTitle.Set(nil)
}

// UnsetFolderTitle ensures that no value is present for FolderTitle, not even an explicit nil
func (o *DraftLocationInteger) UnsetFolderTitle() {
	o.FolderTitle.Unset()
}

// GetFileId returns the FileId field value if set, zero value otherwise.
func (o *DraftLocationInteger) GetFileId() int32 {
	if o == nil || IsNil(o.FileId) {
		var ret int32
		return ret
	}
	return *o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftLocationInteger) GetFileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FileId) {
		return nil, false
	}
	return o.FileId, true
}

// HasFileId returns a boolean if a field has been set.
func (o *DraftLocationInteger) IsFileIdSet() bool {
	if o != nil && !IsNil(o.FileId) {
		return true
	}

	return false
}

// SetFileId gets a reference to the given int32 and assigns it to the FileId field.
func (o *DraftLocationInteger) SetFileId(v int32) {
	o.FileId = &v
}

// GetFileTitle returns the FileTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DraftLocationInteger) GetFileTitle() string {
	if o == nil || IsNil(o.FileTitle.Get()) {
		var ret string
		return ret
	}
	return *o.FileTitle.Get()
}

// GetFileTitleOk returns a tuple with the FileTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DraftLocationInteger) GetFileTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileTitle.Get(), o.FileTitle.IsSet()
}

// HasFileTitle returns a boolean if a field has been set.
func (o *DraftLocationInteger) IsFileTitleSet() bool {
	if o != nil && o.FileTitle.IsSet() {
		return true
	}

	return false
}

// SetFileTitle gets a reference to the given NullableString and assigns it to the FileTitle field.
func (o *DraftLocationInteger) SetFileTitle(v string) {
	o.FileTitle.Set(&v)
}
// SetFileTitleNil sets the value for FileTitle to be an explicit nil
func (o *DraftLocationInteger) SetFileTitleNil() {
	o.FileTitle.Set(nil)
}

// UnsetFileTitle ensures that no value is present for FileTitle, not even an explicit nil
func (o *DraftLocationInteger) UnsetFileTitle() {
	o.FileTitle.Unset()
}

func (o DraftLocationInteger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DraftLocationInteger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FolderId) {
		toSerialize["folderId"] = o.FolderId
	}
	if o.FolderTitle.IsSet() {
		toSerialize["folderTitle"] = o.FolderTitle.Get()
	}
	if !IsNil(o.FileId) {
		toSerialize["fileId"] = o.FileId
	}
	if o.FileTitle.IsSet() {
		toSerialize["fileTitle"] = o.FileTitle.Get()
	}
	return toSerialize, nil
}

type NullableDraftLocationInteger struct {
	value *DraftLocationInteger
	isSet bool
}

func (v NullableDraftLocationInteger) Get() *DraftLocationInteger {
	return v.value
}

func (v *NullableDraftLocationInteger) Set(val *DraftLocationInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableDraftLocationInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableDraftLocationInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDraftLocationInteger(val *DraftLocationInteger) *NullableDraftLocationInteger {
	return &NullableDraftLocationInteger{value: val, isSet: true}
}

func (v NullableDraftLocationInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDraftLocationInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

