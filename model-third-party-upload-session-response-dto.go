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

// checks if the ThirdPartyUploadSessionResponseDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyUploadSessionResponseDto{}

// ThirdPartyUploadSessionResponseDto How far a chunked upload has got, and the file it produced once the last byte has arrived.
type ThirdPartyUploadSessionResponseDto struct {
	// The file the parts are being written into. An upload that took over a file of the same title carries it from  the start, while an upload that creates a new file has nothing to name yet and reports 0 until the answer that  sets `uploaded` to true.
	Id NullableString `json:"id,omitempty"`
	// The folder receiving the file. It is the folder the upload was reserved against, or the sub-folder created for  it when the reservation declared a relative path.
	FolderId NullableString `json:"folderId,omitempty"`
	// The revision the content is being written as: 1 for a file that did not exist, the next number when the upload  took over a file of the same title, and the unchanged current number for an upload opened over an existing  file, which replaces its content in place.
	Version *int32 `json:"version,omitempty"`
	// The title the file is stored under, after characters a title cannot hold were replaced and, where a second  copy was asked for, a numeric suffix was added - so it can differ from the name that was sent.
	Title NullableString `json:"title,omitempty"`
	// The third-party service holding the destination, such as `GoogleDrive` or `OneDrive`, and null for a folder  stored on the portal itself.
	ProviderKey NullableString `json:"providerKey,omitempty"`
	// False while bytes are still missing, when the answer only reports progress; true in the answer that reports  the stored file, which is also the answer that arrives with 201.
	Uploaded *bool `json:"uploaded,omitempty"`
	// The file as it stands. It is filled in both answers, but while `uploaded` is false it describes a file that  has not been written yet, so its identifier, size and links are only worth reading once that flag turns true.
	File *ThirdPartyFileDto `json:"file,omitempty"`
}

// NewThirdPartyUploadSessionResponseDto instantiates a new ThirdPartyUploadSessionResponseDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyUploadSessionResponseDto() *ThirdPartyUploadSessionResponseDto {
	this := ThirdPartyUploadSessionResponseDto{}
	return &this
}

// NewThirdPartyUploadSessionResponseDtoWithDefaults instantiates a new ThirdPartyUploadSessionResponseDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyUploadSessionResponseDtoWithDefaults() *ThirdPartyUploadSessionResponseDto {
	this := ThirdPartyUploadSessionResponseDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyUploadSessionResponseDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyUploadSessionResponseDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ThirdPartyUploadSessionResponseDto) IsIdSet() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ThirdPartyUploadSessionResponseDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ThirdPartyUploadSessionResponseDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ThirdPartyUploadSessionResponseDto) UnsetId() {
	o.Id.Unset()
}

// GetFolderId returns the FolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyUploadSessionResponseDto) GetFolderId() string {
	if o == nil || IsNil(o.FolderId.Get()) {
		var ret string
		return ret
	}
	return *o.FolderId.Get()
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyUploadSessionResponseDto) GetFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FolderId.Get(), o.FolderId.IsSet()
}

// HasFolderId returns a boolean if a field has been set.
func (o *ThirdPartyUploadSessionResponseDto) IsFolderIdSet() bool {
	if o != nil && o.FolderId.IsSet() {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given NullableString and assigns it to the FolderId field.
func (o *ThirdPartyUploadSessionResponseDto) SetFolderId(v string) {
	o.FolderId.Set(&v)
}
// SetFolderIdNil sets the value for FolderId to be an explicit nil
func (o *ThirdPartyUploadSessionResponseDto) SetFolderIdNil() {
	o.FolderId.Set(nil)
}

// UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
func (o *ThirdPartyUploadSessionResponseDto) UnsetFolderId() {
	o.FolderId.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ThirdPartyUploadSessionResponseDto) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyUploadSessionResponseDto) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ThirdPartyUploadSessionResponseDto) IsVersionSet() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *ThirdPartyUploadSessionResponseDto) SetVersion(v int32) {
	o.Version = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyUploadSessionResponseDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyUploadSessionResponseDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ThirdPartyUploadSessionResponseDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ThirdPartyUploadSessionResponseDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ThirdPartyUploadSessionResponseDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ThirdPartyUploadSessionResponseDto) UnsetTitle() {
	o.Title.Unset()
}

// GetProviderKey returns the ProviderKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyUploadSessionResponseDto) GetProviderKey() string {
	if o == nil || IsNil(o.ProviderKey.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderKey.Get()
}

// GetProviderKeyOk returns a tuple with the ProviderKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyUploadSessionResponseDto) GetProviderKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderKey.Get(), o.ProviderKey.IsSet()
}

// HasProviderKey returns a boolean if a field has been set.
func (o *ThirdPartyUploadSessionResponseDto) IsProviderKeySet() bool {
	if o != nil && o.ProviderKey.IsSet() {
		return true
	}

	return false
}

// SetProviderKey gets a reference to the given NullableString and assigns it to the ProviderKey field.
func (o *ThirdPartyUploadSessionResponseDto) SetProviderKey(v string) {
	o.ProviderKey.Set(&v)
}
// SetProviderKeyNil sets the value for ProviderKey to be an explicit nil
func (o *ThirdPartyUploadSessionResponseDto) SetProviderKeyNil() {
	o.ProviderKey.Set(nil)
}

// UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil
func (o *ThirdPartyUploadSessionResponseDto) UnsetProviderKey() {
	o.ProviderKey.Unset()
}

// GetUploaded returns the Uploaded field value if set, zero value otherwise.
func (o *ThirdPartyUploadSessionResponseDto) GetUploaded() bool {
	if o == nil || IsNil(o.Uploaded) {
		var ret bool
		return ret
	}
	return *o.Uploaded
}

// GetUploadedOk returns a tuple with the Uploaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyUploadSessionResponseDto) GetUploadedOk() (*bool, bool) {
	if o == nil || IsNil(o.Uploaded) {
		return nil, false
	}
	return o.Uploaded, true
}

// HasUploaded returns a boolean if a field has been set.
func (o *ThirdPartyUploadSessionResponseDto) IsUploadedSet() bool {
	if o != nil && !IsNil(o.Uploaded) {
		return true
	}

	return false
}

// SetUploaded gets a reference to the given bool and assigns it to the Uploaded field.
func (o *ThirdPartyUploadSessionResponseDto) SetUploaded(v bool) {
	o.Uploaded = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *ThirdPartyUploadSessionResponseDto) GetFile() ThirdPartyFileDto {
	if o == nil || IsNil(o.File) {
		var ret ThirdPartyFileDto
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyUploadSessionResponseDto) GetFileOk() (*ThirdPartyFileDto, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *ThirdPartyUploadSessionResponseDto) IsFileSet() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given ThirdPartyFileDto and assigns it to the File field.
func (o *ThirdPartyUploadSessionResponseDto) SetFile(v ThirdPartyFileDto) {
	o.File = &v
}

func (o ThirdPartyUploadSessionResponseDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyUploadSessionResponseDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.FolderId.IsSet() {
		toSerialize["folderId"] = o.FolderId.Get()
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

type NullableThirdPartyUploadSessionResponseDto struct {
	value *ThirdPartyUploadSessionResponseDto
	isSet bool
}

func (v NullableThirdPartyUploadSessionResponseDto) Get() *ThirdPartyUploadSessionResponseDto {
	return v.value
}

func (v *NullableThirdPartyUploadSessionResponseDto) Set(val *ThirdPartyUploadSessionResponseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyUploadSessionResponseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyUploadSessionResponseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyUploadSessionResponseDto(val *ThirdPartyUploadSessionResponseDto) *NullableThirdPartyUploadSessionResponseDto {
	return &NullableThirdPartyUploadSessionResponseDto{value: val, isSet: true}
}

func (v NullableThirdPartyUploadSessionResponseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyUploadSessionResponseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

