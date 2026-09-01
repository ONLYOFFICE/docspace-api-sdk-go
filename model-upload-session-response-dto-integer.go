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

// checks if the UploadSessionResponseDtoInteger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadSessionResponseDtoInteger{}

// UploadSessionResponseDtoInteger The upload session response parameters.
type UploadSessionResponseDtoInteger struct {
	// The upload session ID.
	Id *int32 `json:"id,omitempty"`
	// The folder ID where the file is being uploaded.
	FolderId *int32 `json:"folderId,omitempty"`
	// The file version number.
	Version *int32 `json:"version,omitempty"`
	// The file title.
	Title NullableString `json:"title,omitempty"`
	// The third-party provider key.
	ProviderKey NullableString `json:"providerKey,omitempty"`
	// Specifies whether the file has been uploaded.
	Uploaded *bool `json:"uploaded,omitempty"`
	// The uploaded file information.
	File *FileDtoInteger `json:"file,omitempty"`
}

// NewUploadSessionResponseDtoInteger instantiates a new UploadSessionResponseDtoInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadSessionResponseDtoInteger() *UploadSessionResponseDtoInteger {
	this := UploadSessionResponseDtoInteger{}
	return &this
}

// NewUploadSessionResponseDtoIntegerWithDefaults instantiates a new UploadSessionResponseDtoInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadSessionResponseDtoIntegerWithDefaults() *UploadSessionResponseDtoInteger {
	this := UploadSessionResponseDtoInteger{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UploadSessionResponseDtoInteger) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadSessionResponseDtoInteger) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UploadSessionResponseDtoInteger) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UploadSessionResponseDtoInteger) SetId(v int32) {
	o.Id = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *UploadSessionResponseDtoInteger) GetFolderId() int32 {
	if o == nil || IsNil(o.FolderId) {
		var ret int32
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadSessionResponseDtoInteger) GetFolderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FolderId) {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *UploadSessionResponseDtoInteger) IsFolderIdSet() bool {
	if o != nil && !IsNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given int32 and assigns it to the FolderId field.
func (o *UploadSessionResponseDtoInteger) SetFolderId(v int32) {
	o.FolderId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *UploadSessionResponseDtoInteger) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadSessionResponseDtoInteger) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *UploadSessionResponseDtoInteger) IsVersionSet() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *UploadSessionResponseDtoInteger) SetVersion(v int32) {
	o.Version = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UploadSessionResponseDtoInteger) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UploadSessionResponseDtoInteger) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *UploadSessionResponseDtoInteger) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *UploadSessionResponseDtoInteger) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *UploadSessionResponseDtoInteger) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *UploadSessionResponseDtoInteger) UnsetTitle() {
	o.Title.Unset()
}

// GetProviderKey returns the ProviderKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UploadSessionResponseDtoInteger) GetProviderKey() string {
	if o == nil || IsNil(o.ProviderKey.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderKey.Get()
}

// GetProviderKeyOk returns a tuple with the ProviderKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UploadSessionResponseDtoInteger) GetProviderKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderKey.Get(), o.ProviderKey.IsSet()
}

// HasProviderKey returns a boolean if a field has been set.
func (o *UploadSessionResponseDtoInteger) IsProviderKeySet() bool {
	if o != nil && o.ProviderKey.IsSet() {
		return true
	}

	return false
}

// SetProviderKey gets a reference to the given NullableString and assigns it to the ProviderKey field.
func (o *UploadSessionResponseDtoInteger) SetProviderKey(v string) {
	o.ProviderKey.Set(&v)
}
// SetProviderKeyNil sets the value for ProviderKey to be an explicit nil
func (o *UploadSessionResponseDtoInteger) SetProviderKeyNil() {
	o.ProviderKey.Set(nil)
}

// UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil
func (o *UploadSessionResponseDtoInteger) UnsetProviderKey() {
	o.ProviderKey.Unset()
}

// GetUploaded returns the Uploaded field value if set, zero value otherwise.
func (o *UploadSessionResponseDtoInteger) GetUploaded() bool {
	if o == nil || IsNil(o.Uploaded) {
		var ret bool
		return ret
	}
	return *o.Uploaded
}

// GetUploadedOk returns a tuple with the Uploaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadSessionResponseDtoInteger) GetUploadedOk() (*bool, bool) {
	if o == nil || IsNil(o.Uploaded) {
		return nil, false
	}
	return o.Uploaded, true
}

// HasUploaded returns a boolean if a field has been set.
func (o *UploadSessionResponseDtoInteger) IsUploadedSet() bool {
	if o != nil && !IsNil(o.Uploaded) {
		return true
	}

	return false
}

// SetUploaded gets a reference to the given bool and assigns it to the Uploaded field.
func (o *UploadSessionResponseDtoInteger) SetUploaded(v bool) {
	o.Uploaded = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *UploadSessionResponseDtoInteger) GetFile() FileDtoInteger {
	if o == nil || IsNil(o.File) {
		var ret FileDtoInteger
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadSessionResponseDtoInteger) GetFileOk() (*FileDtoInteger, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *UploadSessionResponseDtoInteger) IsFileSet() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given FileDtoInteger and assigns it to the File field.
func (o *UploadSessionResponseDtoInteger) SetFile(v FileDtoInteger) {
	o.File = &v
}

func (o UploadSessionResponseDtoInteger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadSessionResponseDtoInteger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.FolderId) {
		toSerialize["folderId"] = o.FolderId
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.ProviderKey.IsSet() {
		toSerialize["providerKey"] = o.ProviderKey.Get()
	}
	if !IsNil(o.Uploaded) {
		toSerialize["uploaded"] = o.Uploaded
	}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	return toSerialize, nil
}

type NullableUploadSessionResponseDtoInteger struct {
	value *UploadSessionResponseDtoInteger
	isSet bool
}

func (v NullableUploadSessionResponseDtoInteger) Get() *UploadSessionResponseDtoInteger {
	return v.value
}

func (v *NullableUploadSessionResponseDtoInteger) Set(val *UploadSessionResponseDtoInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadSessionResponseDtoInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadSessionResponseDtoInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadSessionResponseDtoInteger(val *UploadSessionResponseDtoInteger) *NullableUploadSessionResponseDtoInteger {
	return &NullableUploadSessionResponseDtoInteger{value: val, isSet: true}
}

func (v NullableUploadSessionResponseDtoInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadSessionResponseDtoInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

