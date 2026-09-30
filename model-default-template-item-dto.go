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
	"time"
	"bytes"
	"fmt"
)

// checks if the DefaultTemplateItemDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultTemplateItemDto{}

// DefaultTemplateItemDto The blank document configured for one extension.
type DefaultTemplateItemDto struct {
	// The copy stored in the portal that serves as the blank for this extension. A null means no custom blank has  been chosen and new documents start from the portal's built-in one; the other fields of the entry are then  empty as well.
	SelectedFile NullableInt32 `json:"selectedFile,omitempty"`
	// The extension the entry describes, in lower case with the leading dot. It is the value to send back when this  blank is replaced or reset.
	FileExtension NullableString `json:"fileExtension"`
	// The name the custom blank was copied under, useful for showing which document was chosen. Empty while the  built-in blank is in use.
	FileTitle NullableString `json:"fileTitle,omitempty"`
	// When the custom blank was last changed, in the time zone of the portal. Null while the built-in blank is in  use.
	LastModified NullableTime `json:"lastModified,omitempty"`
	// The size of the custom blank in bytes. Null while the built-in blank is in use.
	FileSize NullableInt64 `json:"fileSize,omitempty"`
	// The address the custom blank can be downloaded from, already carrying the access key of the calling account.  Empty while the built-in blank is in use.
	ViewUrl NullableString `json:"viewUrl,omitempty"`
}

type _DefaultTemplateItemDto DefaultTemplateItemDto

// NewDefaultTemplateItemDto instantiates a new DefaultTemplateItemDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultTemplateItemDto(fileExtension NullableString) *DefaultTemplateItemDto {
	this := DefaultTemplateItemDto{}
	this.FileExtension = fileExtension
	return &this
}

// NewDefaultTemplateItemDtoWithDefaults instantiates a new DefaultTemplateItemDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultTemplateItemDtoWithDefaults() *DefaultTemplateItemDto {
	this := DefaultTemplateItemDto{}
	return &this
}

// GetSelectedFile returns the SelectedFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultTemplateItemDto) GetSelectedFile() int32 {
	if o == nil || IsNil(o.SelectedFile.Get()) {
		var ret int32
		return ret
	}
	return *o.SelectedFile.Get()
}

// GetSelectedFileOk returns a tuple with the SelectedFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultTemplateItemDto) GetSelectedFileOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelectedFile.Get(), o.SelectedFile.IsSet()
}

// HasSelectedFile returns a boolean if a field has been set.
func (o *DefaultTemplateItemDto) IsSelectedFileSet() bool {
	if o != nil && o.SelectedFile.IsSet() {
		return true
	}

	return false
}

// SetSelectedFile gets a reference to the given NullableInt32 and assigns it to the SelectedFile field.
func (o *DefaultTemplateItemDto) SetSelectedFile(v int32) {
	o.SelectedFile.Set(&v)
}
// SetSelectedFileNil sets the value for SelectedFile to be an explicit nil
func (o *DefaultTemplateItemDto) SetSelectedFileNil() {
	o.SelectedFile.Set(nil)
}

// UnsetSelectedFile ensures that no value is present for SelectedFile, not even an explicit nil
func (o *DefaultTemplateItemDto) UnsetSelectedFile() {
	o.SelectedFile.Unset()
}

// GetFileExtension returns the FileExtension field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DefaultTemplateItemDto) GetFileExtension() string {
	if o == nil || o.FileExtension.Get() == nil {
		var ret string
		return ret
	}

	return *o.FileExtension.Get()
}

// GetFileExtensionOk returns a tuple with the FileExtension field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultTemplateItemDto) GetFileExtensionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileExtension.Get(), o.FileExtension.IsSet()
}

// SetFileExtension sets field value
func (o *DefaultTemplateItemDto) SetFileExtension(v string) {
	o.FileExtension.Set(&v)
}

// GetFileTitle returns the FileTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultTemplateItemDto) GetFileTitle() string {
	if o == nil || IsNil(o.FileTitle.Get()) {
		var ret string
		return ret
	}
	return *o.FileTitle.Get()
}

// GetFileTitleOk returns a tuple with the FileTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultTemplateItemDto) GetFileTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileTitle.Get(), o.FileTitle.IsSet()
}

// HasFileTitle returns a boolean if a field has been set.
func (o *DefaultTemplateItemDto) IsFileTitleSet() bool {
	if o != nil && o.FileTitle.IsSet() {
		return true
	}

	return false
}

// SetFileTitle gets a reference to the given NullableString and assigns it to the FileTitle field.
func (o *DefaultTemplateItemDto) SetFileTitle(v string) {
	o.FileTitle.Set(&v)
}
// SetFileTitleNil sets the value for FileTitle to be an explicit nil
func (o *DefaultTemplateItemDto) SetFileTitleNil() {
	o.FileTitle.Set(nil)
}

// UnsetFileTitle ensures that no value is present for FileTitle, not even an explicit nil
func (o *DefaultTemplateItemDto) UnsetFileTitle() {
	o.FileTitle.Unset()
}

// GetLastModified returns the LastModified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultTemplateItemDto) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastModified.Get()
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultTemplateItemDto) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModified.Get(), o.LastModified.IsSet()
}

// HasLastModified returns a boolean if a field has been set.
func (o *DefaultTemplateItemDto) IsLastModifiedSet() bool {
	if o != nil && o.LastModified.IsSet() {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given NullableTime and assigns it to the LastModified field.
func (o *DefaultTemplateItemDto) SetLastModified(v time.Time) {
	o.LastModified.Set(&v)
}
// SetLastModifiedNil sets the value for LastModified to be an explicit nil
func (o *DefaultTemplateItemDto) SetLastModifiedNil() {
	o.LastModified.Set(nil)
}

// UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
func (o *DefaultTemplateItemDto) UnsetLastModified() {
	o.LastModified.Unset()
}

// GetFileSize returns the FileSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultTemplateItemDto) GetFileSize() int64 {
	if o == nil || IsNil(o.FileSize.Get()) {
		var ret int64
		return ret
	}
	return *o.FileSize.Get()
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultTemplateItemDto) GetFileSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileSize.Get(), o.FileSize.IsSet()
}

// HasFileSize returns a boolean if a field has been set.
func (o *DefaultTemplateItemDto) IsFileSizeSet() bool {
	if o != nil && o.FileSize.IsSet() {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given NullableInt64 and assigns it to the FileSize field.
func (o *DefaultTemplateItemDto) SetFileSize(v int64) {
	o.FileSize.Set(&v)
}
// SetFileSizeNil sets the value for FileSize to be an explicit nil
func (o *DefaultTemplateItemDto) SetFileSizeNil() {
	o.FileSize.Set(nil)
}

// UnsetFileSize ensures that no value is present for FileSize, not even an explicit nil
func (o *DefaultTemplateItemDto) UnsetFileSize() {
	o.FileSize.Unset()
}

// GetViewUrl returns the ViewUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultTemplateItemDto) GetViewUrl() string {
	if o == nil || IsNil(o.ViewUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ViewUrl.Get()
}

// GetViewUrlOk returns a tuple with the ViewUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultTemplateItemDto) GetViewUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewUrl.Get(), o.ViewUrl.IsSet()
}

// HasViewUrl returns a boolean if a field has been set.
func (o *DefaultTemplateItemDto) IsViewUrlSet() bool {
	if o != nil && o.ViewUrl.IsSet() {
		return true
	}

	return false
}

// SetViewUrl gets a reference to the given NullableString and assigns it to the ViewUrl field.
func (o *DefaultTemplateItemDto) SetViewUrl(v string) {
	o.ViewUrl.Set(&v)
}
// SetViewUrlNil sets the value for ViewUrl to be an explicit nil
func (o *DefaultTemplateItemDto) SetViewUrlNil() {
	o.ViewUrl.Set(nil)
}

// UnsetViewUrl ensures that no value is present for ViewUrl, not even an explicit nil
func (o *DefaultTemplateItemDto) UnsetViewUrl() {
	o.ViewUrl.Unset()
}

func (o DefaultTemplateItemDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultTemplateItemDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SelectedFile.IsSet() {
		toSerialize["selectedFile"] = o.SelectedFile.Get()
	}
	toSerialize["fileExtension"] = o.FileExtension.Get()
	if o.FileTitle.IsSet() {
		toSerialize["fileTitle"] = o.FileTitle.Get()
	}
	if o.LastModified.IsSet() {
		toSerialize["lastModified"] = o.LastModified.Get()
	}
	if o.FileSize.IsSet() {
		toSerialize["fileSize"] = o.FileSize.Get()
	}
	if o.ViewUrl.IsSet() {
		toSerialize["viewUrl"] = o.ViewUrl.Get()
	}
	return toSerialize, nil
}

func (o *DefaultTemplateItemDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fileExtension",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDefaultTemplateItemDto := _DefaultTemplateItemDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varDefaultTemplateItemDto)

	if err != nil {
		return err
	}

	*o = DefaultTemplateItemDto(varDefaultTemplateItemDto)

	return err
}

type NullableDefaultTemplateItemDto struct {
	value *DefaultTemplateItemDto
	isSet bool
}

func (v NullableDefaultTemplateItemDto) Get() *DefaultTemplateItemDto {
	return v.value
}

func (v *NullableDefaultTemplateItemDto) Set(val *DefaultTemplateItemDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultTemplateItemDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultTemplateItemDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultTemplateItemDto(val *DefaultTemplateItemDto) *NullableDefaultTemplateItemDto {
	return &NullableDefaultTemplateItemDto{value: val, isSet: true}
}

func (v NullableDefaultTemplateItemDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultTemplateItemDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

