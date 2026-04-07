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
	"os"
)

// checks if the UploadRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadRequestDto{}

// UploadRequestDto The request parameters for uploading a file.
type UploadRequestDto struct {
	// The file to be uploaded.
	File *os.File `json:"file,omitempty"`
	ContentType *ContentType `json:"contentType,omitempty"`
	ContentDisposition *ContentDisposition `json:"contentDisposition,omitempty"`
	// The list of files when specified as multipart/form-data.
	Files []*os.File `json:"files,omitempty"`
	// Specifies whether to create the new file if it already exists or not.
	CreateNewIfExist *bool `json:"createNewIfExist,omitempty"`
	// Specifies whether to upload documents in the original formats as well or not.
	StoreOriginalFileFlag NullableBool `json:"storeOriginalFileFlag,omitempty"`
	// Specifies whether to keep the file converting status or not.
	KeepConvertStatus *bool `json:"keepConvertStatus,omitempty"`
	// The request input stream.
	Stream *os.File `json:"stream,omitempty"`
}

// NewUploadRequestDto instantiates a new UploadRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadRequestDto() *UploadRequestDto {
	this := UploadRequestDto{}
	return &this
}

// NewUploadRequestDtoWithDefaults instantiates a new UploadRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadRequestDtoWithDefaults() *UploadRequestDto {
	this := UploadRequestDto{}
	return &this
}

// GetFile returns the File field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UploadRequestDto) GetFile() *os.File {
	if o == nil || IsNil(o.File) {
		var ret *os.File
		return ret
	}
	return o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UploadRequestDto) GetFileOk() (*os.File, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *UploadRequestDto) IsFileSet() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given *os.File and assigns it to the File field.
func (o *UploadRequestDto) SetFile(v *os.File) {
	o.File = v
}
// SetFileNil sets the value for File to nil
func (o *UploadRequestDto) SetFileNil() {
	o.File = nil
}

// UnsetFile clears the value for File
func (o *UploadRequestDto) UnsetFile() {
	o.File = nil
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *UploadRequestDto) GetContentType() ContentType {
	if o == nil || IsNil(o.ContentType) {
		var ret ContentType
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadRequestDto) GetContentTypeOk() (*ContentType, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *UploadRequestDto) IsContentTypeSet() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given ContentType and assigns it to the ContentType field.
func (o *UploadRequestDto) SetContentType(v ContentType) {
	o.ContentType = &v
}

// GetContentDisposition returns the ContentDisposition field value if set, zero value otherwise.
func (o *UploadRequestDto) GetContentDisposition() ContentDisposition {
	if o == nil || IsNil(o.ContentDisposition) {
		var ret ContentDisposition
		return ret
	}
	return *o.ContentDisposition
}

// GetContentDispositionOk returns a tuple with the ContentDisposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadRequestDto) GetContentDispositionOk() (*ContentDisposition, bool) {
	if o == nil || IsNil(o.ContentDisposition) {
		return nil, false
	}
	return o.ContentDisposition, true
}

// HasContentDisposition returns a boolean if a field has been set.
func (o *UploadRequestDto) IsContentDispositionSet() bool {
	if o != nil && !IsNil(o.ContentDisposition) {
		return true
	}

	return false
}

// SetContentDisposition gets a reference to the given ContentDisposition and assigns it to the ContentDisposition field.
func (o *UploadRequestDto) SetContentDisposition(v ContentDisposition) {
	o.ContentDisposition = &v
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UploadRequestDto) GetFiles() []*os.File {
	if o == nil {
		var ret []*os.File
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UploadRequestDto) GetFilesOk() ([]*os.File, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *UploadRequestDto) IsFilesSet() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []*os.File and assigns it to the Files field.
func (o *UploadRequestDto) SetFiles(v []*os.File) {
	o.Files = v
}

// GetCreateNewIfExist returns the CreateNewIfExist field value if set, zero value otherwise.
func (o *UploadRequestDto) GetCreateNewIfExist() bool {
	if o == nil || IsNil(o.CreateNewIfExist) {
		var ret bool
		return ret
	}
	return *o.CreateNewIfExist
}

// GetCreateNewIfExistOk returns a tuple with the CreateNewIfExist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadRequestDto) GetCreateNewIfExistOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateNewIfExist) {
		return nil, false
	}
	return o.CreateNewIfExist, true
}

// HasCreateNewIfExist returns a boolean if a field has been set.
func (o *UploadRequestDto) IsCreateNewIfExistSet() bool {
	if o != nil && !IsNil(o.CreateNewIfExist) {
		return true
	}

	return false
}

// SetCreateNewIfExist gets a reference to the given bool and assigns it to the CreateNewIfExist field.
func (o *UploadRequestDto) SetCreateNewIfExist(v bool) {
	o.CreateNewIfExist = &v
}

// GetStoreOriginalFileFlag returns the StoreOriginalFileFlag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UploadRequestDto) GetStoreOriginalFileFlag() bool {
	if o == nil || IsNil(o.StoreOriginalFileFlag.Get()) {
		var ret bool
		return ret
	}
	return *o.StoreOriginalFileFlag.Get()
}

// GetStoreOriginalFileFlagOk returns a tuple with the StoreOriginalFileFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UploadRequestDto) GetStoreOriginalFileFlagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoreOriginalFileFlag.Get(), o.StoreOriginalFileFlag.IsSet()
}

// HasStoreOriginalFileFlag returns a boolean if a field has been set.
func (o *UploadRequestDto) IsStoreOriginalFileFlagSet() bool {
	if o != nil && o.StoreOriginalFileFlag.IsSet() {
		return true
	}

	return false
}

// SetStoreOriginalFileFlag gets a reference to the given NullableBool and assigns it to the StoreOriginalFileFlag field.
func (o *UploadRequestDto) SetStoreOriginalFileFlag(v bool) {
	o.StoreOriginalFileFlag.Set(&v)
}
// SetStoreOriginalFileFlagNil sets the value for StoreOriginalFileFlag to be an explicit nil
func (o *UploadRequestDto) SetStoreOriginalFileFlagNil() {
	o.StoreOriginalFileFlag.Set(nil)
}

// UnsetStoreOriginalFileFlag ensures that no value is present for StoreOriginalFileFlag, not even an explicit nil
func (o *UploadRequestDto) UnsetStoreOriginalFileFlag() {
	o.StoreOriginalFileFlag.Unset()
}

// GetKeepConvertStatus returns the KeepConvertStatus field value if set, zero value otherwise.
func (o *UploadRequestDto) GetKeepConvertStatus() bool {
	if o == nil || IsNil(o.KeepConvertStatus) {
		var ret bool
		return ret
	}
	return *o.KeepConvertStatus
}

// GetKeepConvertStatusOk returns a tuple with the KeepConvertStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadRequestDto) GetKeepConvertStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.KeepConvertStatus) {
		return nil, false
	}
	return o.KeepConvertStatus, true
}

// HasKeepConvertStatus returns a boolean if a field has been set.
func (o *UploadRequestDto) IsKeepConvertStatusSet() bool {
	if o != nil && !IsNil(o.KeepConvertStatus) {
		return true
	}

	return false
}

// SetKeepConvertStatus gets a reference to the given bool and assigns it to the KeepConvertStatus field.
func (o *UploadRequestDto) SetKeepConvertStatus(v bool) {
	o.KeepConvertStatus = &v
}

// GetStream returns the Stream field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UploadRequestDto) GetStream() *os.File {
	if o == nil || IsNil(o.Stream) {
		var ret *os.File
		return ret
	}
	return o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UploadRequestDto) GetStreamOk() (*os.File, bool) {
	if o == nil || IsNil(o.Stream) {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *UploadRequestDto) IsStreamSet() bool {
	if o != nil && !IsNil(o.Stream) {
		return true
	}

	return false
}

// SetStream gets a reference to the given *os.File and assigns it to the Stream field.
func (o *UploadRequestDto) SetStream(v *os.File) {
	o.Stream = v
}
// SetStreamNil sets the value for Stream to nil
func (o *UploadRequestDto) SetStreamNil() {
	o.Stream = nil
}

// UnsetStream clears the value for Stream
func (o *UploadRequestDto) UnsetStream() {
	o.Stream = nil
}

func (o UploadRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	if !IsNil(o.ContentType) {
		toSerialize["contentType"] = o.ContentType
	}
	if !IsNil(o.ContentDisposition) {
		toSerialize["contentDisposition"] = o.ContentDisposition
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if !IsNil(o.CreateNewIfExist) {
		toSerialize["createNewIfExist"] = o.CreateNewIfExist
	}
	if o.StoreOriginalFileFlag.IsSet() {
		toSerialize["storeOriginalFileFlag"] = o.StoreOriginalFileFlag.Get()
	}
	if !IsNil(o.KeepConvertStatus) {
		toSerialize["keepConvertStatus"] = o.KeepConvertStatus
	}
	if !IsNil(o.Stream) {
		toSerialize["stream"] = o.Stream
	}
	return toSerialize, nil
}

type NullableUploadRequestDto struct {
	value *UploadRequestDto
	isSet bool
}

func (v NullableUploadRequestDto) Get() *UploadRequestDto {
	return v.value
}

func (v *NullableUploadRequestDto) Set(val *UploadRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadRequestDto(val *UploadRequestDto) *NullableUploadRequestDto {
	return &NullableUploadRequestDto{value: val, isSet: true}
}

func (v NullableUploadRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

