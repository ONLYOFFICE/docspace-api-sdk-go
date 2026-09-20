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

// checks if the DraftLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DraftLocation{}

// DraftLocation Where the caller's own filling draft of a form is kept.
type DraftLocation struct {
	// The folder holding the draft: the sub-folder that the room for filling keeps for drafts of this particular  form.
	FolderId *int32 `json:"folderId,omitempty"`
	// The title of that folder, which the portal takes from the form itself when the form is released for filling.
	FolderTitle NullableString `json:"folderTitle,omitempty"`
	// The draft itself - the copy the caller fills in, not the original form, and the identifier to pass to the file  operations while filling.
	FileId *int32 `json:"fileId,omitempty"`
	// The title of the draft, which the portal builds from the name of the person filling it and the name of the  form. Null when the draft the record points at no longer exists.
	FileTitle NullableString `json:"fileTitle,omitempty"`
}

// NewDraftLocation instantiates a new DraftLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDraftLocation() *DraftLocation {
	this := DraftLocation{}
	return &this
}

// NewDraftLocationWithDefaults instantiates a new DraftLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDraftLocationWithDefaults() *DraftLocation {
	this := DraftLocation{}
	return &this
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *DraftLocation) GetFolderId() int32 {
	if o == nil || IsNil(o.FolderId) {
		var ret int32
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftLocation) GetFolderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FolderId) {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *DraftLocation) IsFolderIdSet() bool {
	if o != nil && !IsNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given int32 and assigns it to the FolderId field.
func (o *DraftLocation) SetFolderId(v int32) {
	o.FolderId = &v
}

// GetFolderTitle returns the FolderTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DraftLocation) GetFolderTitle() string {
	if o == nil || IsNil(o.FolderTitle.Get()) {
		var ret string
		return ret
	}
	return *o.FolderTitle.Get()
}

// GetFolderTitleOk returns a tuple with the FolderTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DraftLocation) GetFolderTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderTitle.Get(), o.FolderTitle.IsSet()
}

// HasFolderTitle returns a boolean if a field has been set.
func (o *DraftLocation) IsFolderTitleSet() bool {
	if o != nil && o.FolderTitle.IsSet() {
		return true
	}

	return false
}

// SetFolderTitle gets a reference to the given NullableString and assigns it to the FolderTitle field.
func (o *DraftLocation) SetFolderTitle(v string) {
	o.FolderTitle.Set(&v)
}
// SetFolderTitleNil sets the value for FolderTitle to be an explicit nil
func (o *DraftLocation) SetFolderTitleNil() {
	o.FolderTitle.Set(nil)
}

// UnsetFolderTitle ensures that no value is present for FolderTitle, not even an explicit nil
func (o *DraftLocation) UnsetFolderTitle() {
	o.FolderTitle.Unset()
}

// GetFileId returns the FileId field value if set, zero value otherwise.
func (o *DraftLocation) GetFileId() int32 {
	if o == nil || IsNil(o.FileId) {
		var ret int32
		return ret
	}
	return *o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DraftLocation) GetFileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FileId) {
		return nil, false
	}
	return o.FileId, true
}

// HasFileId returns a boolean if a field has been set.
func (o *DraftLocation) IsFileIdSet() bool {
	if o != nil && !IsNil(o.FileId) {
		return true
	}

	return false
}

// SetFileId gets a reference to the given int32 and assigns it to the FileId field.
func (o *DraftLocation) SetFileId(v int32) {
	o.FileId = &v
}

// GetFileTitle returns the FileTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DraftLocation) GetFileTitle() string {
	if o == nil || IsNil(o.FileTitle.Get()) {
		var ret string
		return ret
	}
	return *o.FileTitle.Get()
}

// GetFileTitleOk returns a tuple with the FileTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DraftLocation) GetFileTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileTitle.Get(), o.FileTitle.IsSet()
}

// HasFileTitle returns a boolean if a field has been set.
func (o *DraftLocation) IsFileTitleSet() bool {
	if o != nil && o.FileTitle.IsSet() {
		return true
	}

	return false
}

// SetFileTitle gets a reference to the given NullableString and assigns it to the FileTitle field.
func (o *DraftLocation) SetFileTitle(v string) {
	o.FileTitle.Set(&v)
}
// SetFileTitleNil sets the value for FileTitle to be an explicit nil
func (o *DraftLocation) SetFileTitleNil() {
	o.FileTitle.Set(nil)
}

// UnsetFileTitle ensures that no value is present for FileTitle, not even an explicit nil
func (o *DraftLocation) UnsetFileTitle() {
	o.FileTitle.Unset()
}

func (o DraftLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DraftLocation) ToMap() (map[string]interface{}, error) {
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

type NullableDraftLocation struct {
	value *DraftLocation
	isSet bool
}

func (v NullableDraftLocation) Get() *DraftLocation {
	return v.value
}

func (v *NullableDraftLocation) Set(val *DraftLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableDraftLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableDraftLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDraftLocation(val *DraftLocation) *NullableDraftLocation {
	return &NullableDraftLocation{value: val, isSet: true}
}

func (v NullableDraftLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDraftLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

