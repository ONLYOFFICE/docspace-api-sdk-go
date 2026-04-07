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

// checks if the FileReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileReference{}

// FileReference The file reference parameters.
type FileReference struct {
	ReferenceData *FileReferenceData `json:"referenceData,omitempty"`
	// The error message text.
	Error NullableString `json:"error,omitempty"`
	// The file name or relative path for the formula editor.
	Path NullableString `json:"path,omitempty"`
	// The URL address to download the current file.
	Url NullableString `json:"url,omitempty"`
	// An extension of the document specified with the url parameter.
	FileType NullableString `json:"fileType,omitempty"`
	// The unique document identifier used by the service to take the data from the co-editing session.
	Key NullableString `json:"key,omitempty"`
	// The file URL.
	Link NullableString `json:"link,omitempty"`
	// The encrypted signature added to the parameter in the form of a token.
	Token NullableString `json:"token,omitempty"`
}

// NewFileReference instantiates a new FileReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileReference() *FileReference {
	this := FileReference{}
	return &this
}

// NewFileReferenceWithDefaults instantiates a new FileReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileReferenceWithDefaults() *FileReference {
	this := FileReference{}
	return &this
}

// GetReferenceData returns the ReferenceData field value if set, zero value otherwise.
func (o *FileReference) GetReferenceData() FileReferenceData {
	if o == nil || IsNil(o.ReferenceData) {
		var ret FileReferenceData
		return ret
	}
	return *o.ReferenceData
}

// GetReferenceDataOk returns a tuple with the ReferenceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileReference) GetReferenceDataOk() (*FileReferenceData, bool) {
	if o == nil || IsNil(o.ReferenceData) {
		return nil, false
	}
	return o.ReferenceData, true
}

// HasReferenceData returns a boolean if a field has been set.
func (o *FileReference) IsReferenceDataSet() bool {
	if o != nil && !IsNil(o.ReferenceData) {
		return true
	}

	return false
}

// SetReferenceData gets a reference to the given FileReferenceData and assigns it to the ReferenceData field.
func (o *FileReference) SetReferenceData(v FileReferenceData) {
	o.ReferenceData = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileReference) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileReference) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *FileReference) IsErrorSet() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *FileReference) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *FileReference) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *FileReference) UnsetError() {
	o.Error.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileReference) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileReference) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *FileReference) IsPathSet() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *FileReference) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *FileReference) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *FileReference) UnsetPath() {
	o.Path.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileReference) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileReference) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *FileReference) IsUrlSet() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *FileReference) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *FileReference) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *FileReference) UnsetUrl() {
	o.Url.Unset()
}

// GetFileType returns the FileType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileReference) GetFileType() string {
	if o == nil || IsNil(o.FileType.Get()) {
		var ret string
		return ret
	}
	return *o.FileType.Get()
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileReference) GetFileTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileType.Get(), o.FileType.IsSet()
}

// HasFileType returns a boolean if a field has been set.
func (o *FileReference) IsFileTypeSet() bool {
	if o != nil && o.FileType.IsSet() {
		return true
	}

	return false
}

// SetFileType gets a reference to the given NullableString and assigns it to the FileType field.
func (o *FileReference) SetFileType(v string) {
	o.FileType.Set(&v)
}
// SetFileTypeNil sets the value for FileType to be an explicit nil
func (o *FileReference) SetFileTypeNil() {
	o.FileType.Set(nil)
}

// UnsetFileType ensures that no value is present for FileType, not even an explicit nil
func (o *FileReference) UnsetFileType() {
	o.FileType.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileReference) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileReference) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *FileReference) IsKeySet() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *FileReference) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *FileReference) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *FileReference) UnsetKey() {
	o.Key.Unset()
}

// GetLink returns the Link field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileReference) GetLink() string {
	if o == nil || IsNil(o.Link.Get()) {
		var ret string
		return ret
	}
	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileReference) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// HasLink returns a boolean if a field has been set.
func (o *FileReference) IsLinkSet() bool {
	if o != nil && o.Link.IsSet() {
		return true
	}

	return false
}

// SetLink gets a reference to the given NullableString and assigns it to the Link field.
func (o *FileReference) SetLink(v string) {
	o.Link.Set(&v)
}
// SetLinkNil sets the value for Link to be an explicit nil
func (o *FileReference) SetLinkNil() {
	o.Link.Set(nil)
}

// UnsetLink ensures that no value is present for Link, not even an explicit nil
func (o *FileReference) UnsetLink() {
	o.Link.Unset()
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileReference) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileReference) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *FileReference) IsTokenSet() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *FileReference) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *FileReference) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *FileReference) UnsetToken() {
	o.Token.Unset()
}

func (o FileReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReferenceData) {
		toSerialize["referenceData"] = o.ReferenceData
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.FileType.IsSet() {
		toSerialize["fileType"] = o.FileType.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Link.IsSet() {
		toSerialize["link"] = o.Link.Get()
	}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	return toSerialize, nil
}

type NullableFileReference struct {
	value *FileReference
	isSet bool
}

func (v NullableFileReference) Get() *FileReference {
	return v.value
}

func (v *NullableFileReference) Set(val *FileReference) {
	v.value = val
	v.isSet = true
}

func (v NullableFileReference) IsSet() bool {
	return v.isSet
}

func (v *NullableFileReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileReference(val *FileReference) *NullableFileReference {
	return &NullableFileReference{value: val, isSet: true}
}

func (v NullableFileReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

