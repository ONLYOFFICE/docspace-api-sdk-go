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

// checks if the RunBuilderScriptOutputRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunBuilderScriptOutputRequest{}

// RunBuilderScriptOutputRequest Where one produced file goes.
type RunBuilderScriptOutputRequest struct {
	// The id of a portal file to store the result in as a new version, as reported by a folder listing such as  `GET api/2.0/files/{folderId}`.
	FileId NullableInt32 `json:"fileId,omitempty"`
	// The id of a folder to save the result in as a new file, as reported by a folder listing such as  `GET api/2.0/files/{folderId}`.
	FolderId NullableInt32 `json:"folderId,omitempty"`
	// The title of the new file, when it has to differ from the name the script saved the result under.
	Title NullableString `json:"title,omitempty"`
}

// NewRunBuilderScriptOutputRequest instantiates a new RunBuilderScriptOutputRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunBuilderScriptOutputRequest() *RunBuilderScriptOutputRequest {
	this := RunBuilderScriptOutputRequest{}
	return &this
}

// NewRunBuilderScriptOutputRequestWithDefaults instantiates a new RunBuilderScriptOutputRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunBuilderScriptOutputRequestWithDefaults() *RunBuilderScriptOutputRequest {
	this := RunBuilderScriptOutputRequest{}
	return &this
}

// GetFileId returns the FileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunBuilderScriptOutputRequest) GetFileId() int32 {
	if o == nil || IsNil(o.FileId.Get()) {
		var ret int32
		return ret
	}
	return *o.FileId.Get()
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunBuilderScriptOutputRequest) GetFileIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileId.Get(), o.FileId.IsSet()
}

// HasFileId returns a boolean if a field has been set.
func (o *RunBuilderScriptOutputRequest) IsFileIdSet() bool {
	if o != nil && o.FileId.IsSet() {
		return true
	}

	return false
}

// SetFileId gets a reference to the given NullableInt32 and assigns it to the FileId field.
func (o *RunBuilderScriptOutputRequest) SetFileId(v int32) {
	o.FileId.Set(&v)
}
// SetFileIdNil sets the value for FileId to be an explicit nil
func (o *RunBuilderScriptOutputRequest) SetFileIdNil() {
	o.FileId.Set(nil)
}

// UnsetFileId ensures that no value is present for FileId, not even an explicit nil
func (o *RunBuilderScriptOutputRequest) UnsetFileId() {
	o.FileId.Unset()
}

// GetFolderId returns the FolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunBuilderScriptOutputRequest) GetFolderId() int32 {
	if o == nil || IsNil(o.FolderId.Get()) {
		var ret int32
		return ret
	}
	return *o.FolderId.Get()
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunBuilderScriptOutputRequest) GetFolderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderId.Get(), o.FolderId.IsSet()
}

// HasFolderId returns a boolean if a field has been set.
func (o *RunBuilderScriptOutputRequest) IsFolderIdSet() bool {
	if o != nil && o.FolderId.IsSet() {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given NullableInt32 and assigns it to the FolderId field.
func (o *RunBuilderScriptOutputRequest) SetFolderId(v int32) {
	o.FolderId.Set(&v)
}
// SetFolderIdNil sets the value for FolderId to be an explicit nil
func (o *RunBuilderScriptOutputRequest) SetFolderIdNil() {
	o.FolderId.Set(nil)
}

// UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
func (o *RunBuilderScriptOutputRequest) UnsetFolderId() {
	o.FolderId.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunBuilderScriptOutputRequest) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunBuilderScriptOutputRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *RunBuilderScriptOutputRequest) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *RunBuilderScriptOutputRequest) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *RunBuilderScriptOutputRequest) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *RunBuilderScriptOutputRequest) UnsetTitle() {
	o.Title.Unset()
}

func (o RunBuilderScriptOutputRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunBuilderScriptOutputRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FileId.IsSet() {
		toSerialize["fileId"] = o.FileId.Get()
	}
	if o.FolderId.IsSet() {
		toSerialize["folderId"] = o.FolderId.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	return toSerialize, nil
}

type NullableRunBuilderScriptOutputRequest struct {
	value *RunBuilderScriptOutputRequest
	isSet bool
}

func (v NullableRunBuilderScriptOutputRequest) Get() *RunBuilderScriptOutputRequest {
	return v.value
}

func (v *NullableRunBuilderScriptOutputRequest) Set(val *RunBuilderScriptOutputRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRunBuilderScriptOutputRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRunBuilderScriptOutputRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunBuilderScriptOutputRequest(val *RunBuilderScriptOutputRequest) *NullableRunBuilderScriptOutputRequest {
	return &NullableRunBuilderScriptOutputRequest{value: val, isSet: true}
}

func (v NullableRunBuilderScriptOutputRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunBuilderScriptOutputRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

