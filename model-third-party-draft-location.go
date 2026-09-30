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

// checks if the ThirdPartyDraftLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyDraftLocation{}

// ThirdPartyDraftLocation Where the caller's own filling draft of a form is kept.
type ThirdPartyDraftLocation struct {
	// The folder holding the draft: the sub-folder that the room for filling keeps for drafts of this particular  form.
	FolderId NullableString `json:"folderId,omitempty"`
	// The title of that folder, which the portal takes from the form itself when the form is released for filling.
	FolderTitle NullableString `json:"folderTitle,omitempty"`
	// The draft itself - the copy the caller fills in, not the original form, and the identifier to pass to the file  operations while filling.
	FileId NullableString `json:"fileId,omitempty"`
	// The title of the draft, which the portal builds from the name of the person filling it and the name of the  form. Null when the draft the record points at no longer exists.
	FileTitle NullableString `json:"fileTitle,omitempty"`
}

// NewThirdPartyDraftLocation instantiates a new ThirdPartyDraftLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyDraftLocation() *ThirdPartyDraftLocation {
	this := ThirdPartyDraftLocation{}
	return &this
}

// NewThirdPartyDraftLocationWithDefaults instantiates a new ThirdPartyDraftLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyDraftLocationWithDefaults() *ThirdPartyDraftLocation {
	this := ThirdPartyDraftLocation{}
	return &this
}

// GetFolderId returns the FolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyDraftLocation) GetFolderId() string {
	if o == nil || IsNil(o.FolderId.Get()) {
		var ret string
		return ret
	}
	return *o.FolderId.Get()
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyDraftLocation) GetFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderId.Get(), o.FolderId.IsSet()
}

// HasFolderId returns a boolean if a field has been set.
func (o *ThirdPartyDraftLocation) IsFolderIdSet() bool {
	if o != nil && o.FolderId.IsSet() {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given NullableString and assigns it to the FolderId field.
func (o *ThirdPartyDraftLocation) SetFolderId(v string) {
	o.FolderId.Set(&v)
}
// SetFolderIdNil sets the value for FolderId to be an explicit nil
func (o *ThirdPartyDraftLocation) SetFolderIdNil() {
	o.FolderId.Set(nil)
}

// UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
func (o *ThirdPartyDraftLocation) UnsetFolderId() {
	o.FolderId.Unset()
}

// GetFolderTitle returns the FolderTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyDraftLocation) GetFolderTitle() string {
	if o == nil || IsNil(o.FolderTitle.Get()) {
		var ret string
		return ret
	}
	return *o.FolderTitle.Get()
}

// GetFolderTitleOk returns a tuple with the FolderTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyDraftLocation) GetFolderTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderTitle.Get(), o.FolderTitle.IsSet()
}

// HasFolderTitle returns a boolean if a field has been set.
func (o *ThirdPartyDraftLocation) IsFolderTitleSet() bool {
	if o != nil && o.FolderTitle.IsSet() {
		return true
	}

	return false
}

// SetFolderTitle gets a reference to the given NullableString and assigns it to the FolderTitle field.
func (o *ThirdPartyDraftLocation) SetFolderTitle(v string) {
	o.FolderTitle.Set(&v)
}
// SetFolderTitleNil sets the value for FolderTitle to be an explicit nil
func (o *ThirdPartyDraftLocation) SetFolderTitleNil() {
	o.FolderTitle.Set(nil)
}

// UnsetFolderTitle ensures that no value is present for FolderTitle, not even an explicit nil
func (o *ThirdPartyDraftLocation) UnsetFolderTitle() {
	o.FolderTitle.Unset()
}

// GetFileId returns the FileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyDraftLocation) GetFileId() string {
	if o == nil || IsNil(o.FileId.Get()) {
		var ret string
		return ret
	}
	return *o.FileId.Get()
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyDraftLocation) GetFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileId.Get(), o.FileId.IsSet()
}

// HasFileId returns a boolean if a field has been set.
func (o *ThirdPartyDraftLocation) IsFileIdSet() bool {
	if o != nil && o.FileId.IsSet() {
		return true
	}

	return false
}

// SetFileId gets a reference to the given NullableString and assigns it to the FileId field.
func (o *ThirdPartyDraftLocation) SetFileId(v string) {
	o.FileId.Set(&v)
}
// SetFileIdNil sets the value for FileId to be an explicit nil
func (o *ThirdPartyDraftLocation) SetFileIdNil() {
	o.FileId.Set(nil)
}

// UnsetFileId ensures that no value is present for FileId, not even an explicit nil
func (o *ThirdPartyDraftLocation) UnsetFileId() {
	o.FileId.Unset()
}

// GetFileTitle returns the FileTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyDraftLocation) GetFileTitle() string {
	if o == nil || IsNil(o.FileTitle.Get()) {
		var ret string
		return ret
	}
	return *o.FileTitle.Get()
}

// GetFileTitleOk returns a tuple with the FileTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyDraftLocation) GetFileTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileTitle.Get(), o.FileTitle.IsSet()
}

// HasFileTitle returns a boolean if a field has been set.
func (o *ThirdPartyDraftLocation) IsFileTitleSet() bool {
	if o != nil && o.FileTitle.IsSet() {
		return true
	}

	return false
}

// SetFileTitle gets a reference to the given NullableString and assigns it to the FileTitle field.
func (o *ThirdPartyDraftLocation) SetFileTitle(v string) {
	o.FileTitle.Set(&v)
}
// SetFileTitleNil sets the value for FileTitle to be an explicit nil
func (o *ThirdPartyDraftLocation) SetFileTitleNil() {
	o.FileTitle.Set(nil)
}

// UnsetFileTitle ensures that no value is present for FileTitle, not even an explicit nil
func (o *ThirdPartyDraftLocation) UnsetFileTitle() {
	o.FileTitle.Unset()
}

func (o ThirdPartyDraftLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyDraftLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FolderId.IsSet() {
		toSerialize["folderId"] = o.FolderId.Get()
	}
	if o.FolderTitle.IsSet() {
		toSerialize["folderTitle"] = o.FolderTitle.Get()
	}
	if o.FileId.IsSet() {
		toSerialize["fileId"] = o.FileId.Get()
	}
	if o.FileTitle.IsSet() {
		toSerialize["fileTitle"] = o.FileTitle.Get()
	}
	return toSerialize, nil
}

type NullableThirdPartyDraftLocation struct {
	value *ThirdPartyDraftLocation
	isSet bool
}

func (v NullableThirdPartyDraftLocation) Get() *ThirdPartyDraftLocation {
	return v.value
}

func (v *NullableThirdPartyDraftLocation) Set(val *ThirdPartyDraftLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyDraftLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyDraftLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyDraftLocation(val *ThirdPartyDraftLocation) *NullableThirdPartyDraftLocation {
	return &NullableThirdPartyDraftLocation{value: val, isSet: true}
}

func (v NullableThirdPartyDraftLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyDraftLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

