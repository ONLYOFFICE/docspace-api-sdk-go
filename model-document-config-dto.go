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

// checks if the DocumentConfigDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentConfigDto{}

// DocumentConfigDto The document config parameters.
type DocumentConfigDto struct {
	// The file type of the document.
	FileType NullableString `json:"fileType,omitempty"`
	Info *InfoConfigDto `json:"info,omitempty"`
	// Specifies if the documnet is linked for current user.
	IsLinkedForMe *bool `json:"isLinkedForMe,omitempty"`
	// The document key.
	Key NullableString `json:"key,omitempty"`
	Permissions *PermissionsConfig `json:"permissions,omitempty"`
	// The shared link parameter of the document.
	SharedLinkParam NullableString `json:"sharedLinkParam,omitempty"`
	// The shared link key of the document.
	SharedLinkKey NullableString `json:"sharedLinkKey,omitempty"`
	ReferenceData *FileReferenceData `json:"referenceData,omitempty"`
	// The document title.
	Title NullableString `json:"title,omitempty"`
	// The document url.
	Url NullableString `json:"url,omitempty"`
	// Indicates whether this is a form.
	IsForm *bool `json:"isForm,omitempty"`
	Options *Options `json:"options,omitempty"`
}

// NewDocumentConfigDto instantiates a new DocumentConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentConfigDto() *DocumentConfigDto {
	this := DocumentConfigDto{}
	return &this
}

// NewDocumentConfigDtoWithDefaults instantiates a new DocumentConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentConfigDtoWithDefaults() *DocumentConfigDto {
	this := DocumentConfigDto{}
	return &this
}

// GetFileType returns the FileType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentConfigDto) GetFileType() string {
	if o == nil || IsNil(o.FileType.Get()) {
		var ret string
		return ret
	}
	return *o.FileType.Get()
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentConfigDto) GetFileTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileType.Get(), o.FileType.IsSet()
}

// HasFileType returns a boolean if a field has been set.
func (o *DocumentConfigDto) IsFileTypeSet() bool {
	if o != nil && o.FileType.IsSet() {
		return true
	}

	return false
}

// SetFileType gets a reference to the given NullableString and assigns it to the FileType field.
func (o *DocumentConfigDto) SetFileType(v string) {
	o.FileType.Set(&v)
}
// SetFileTypeNil sets the value for FileType to be an explicit nil
func (o *DocumentConfigDto) SetFileTypeNil() {
	o.FileType.Set(nil)
}

// UnsetFileType ensures that no value is present for FileType, not even an explicit nil
func (o *DocumentConfigDto) UnsetFileType() {
	o.FileType.Unset()
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *DocumentConfigDto) GetInfo() InfoConfigDto {
	if o == nil || IsNil(o.Info) {
		var ret InfoConfigDto
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentConfigDto) GetInfoOk() (*InfoConfigDto, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *DocumentConfigDto) IsInfoSet() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given InfoConfigDto and assigns it to the Info field.
func (o *DocumentConfigDto) SetInfo(v InfoConfigDto) {
	o.Info = &v
}

// GetIsLinkedForMe returns the IsLinkedForMe field value if set, zero value otherwise.
func (o *DocumentConfigDto) GetIsLinkedForMe() bool {
	if o == nil || IsNil(o.IsLinkedForMe) {
		var ret bool
		return ret
	}
	return *o.IsLinkedForMe
}

// GetIsLinkedForMeOk returns a tuple with the IsLinkedForMe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentConfigDto) GetIsLinkedForMeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLinkedForMe) {
		return nil, false
	}
	return o.IsLinkedForMe, true
}

// HasIsLinkedForMe returns a boolean if a field has been set.
func (o *DocumentConfigDto) IsIsLinkedForMeSet() bool {
	if o != nil && !IsNil(o.IsLinkedForMe) {
		return true
	}

	return false
}

// SetIsLinkedForMe gets a reference to the given bool and assigns it to the IsLinkedForMe field.
func (o *DocumentConfigDto) SetIsLinkedForMe(v bool) {
	o.IsLinkedForMe = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentConfigDto) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentConfigDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *DocumentConfigDto) IsKeySet() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *DocumentConfigDto) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *DocumentConfigDto) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *DocumentConfigDto) UnsetKey() {
	o.Key.Unset()
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *DocumentConfigDto) GetPermissions() PermissionsConfig {
	if o == nil || IsNil(o.Permissions) {
		var ret PermissionsConfig
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentConfigDto) GetPermissionsOk() (*PermissionsConfig, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *DocumentConfigDto) IsPermissionsSet() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given PermissionsConfig and assigns it to the Permissions field.
func (o *DocumentConfigDto) SetPermissions(v PermissionsConfig) {
	o.Permissions = &v
}

// GetSharedLinkParam returns the SharedLinkParam field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentConfigDto) GetSharedLinkParam() string {
	if o == nil || IsNil(o.SharedLinkParam.Get()) {
		var ret string
		return ret
	}
	return *o.SharedLinkParam.Get()
}

// GetSharedLinkParamOk returns a tuple with the SharedLinkParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentConfigDto) GetSharedLinkParamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedLinkParam.Get(), o.SharedLinkParam.IsSet()
}

// HasSharedLinkParam returns a boolean if a field has been set.
func (o *DocumentConfigDto) IsSharedLinkParamSet() bool {
	if o != nil && o.SharedLinkParam.IsSet() {
		return true
	}

	return false
}

// SetSharedLinkParam gets a reference to the given NullableString and assigns it to the SharedLinkParam field.
func (o *DocumentConfigDto) SetSharedLinkParam(v string) {
	o.SharedLinkParam.Set(&v)
}
// SetSharedLinkParamNil sets the value for SharedLinkParam to be an explicit nil
func (o *DocumentConfigDto) SetSharedLinkParamNil() {
	o.SharedLinkParam.Set(nil)
}

// UnsetSharedLinkParam ensures that no value is present for SharedLinkParam, not even an explicit nil
func (o *DocumentConfigDto) UnsetSharedLinkParam() {
	o.SharedLinkParam.Unset()
}

// GetSharedLinkKey returns the SharedLinkKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentConfigDto) GetSharedLinkKey() string {
	if o == nil || IsNil(o.SharedLinkKey.Get()) {
		var ret string
		return ret
	}
	return *o.SharedLinkKey.Get()
}

// GetSharedLinkKeyOk returns a tuple with the SharedLinkKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentConfigDto) GetSharedLinkKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedLinkKey.Get(), o.SharedLinkKey.IsSet()
}

// HasSharedLinkKey returns a boolean if a field has been set.
func (o *DocumentConfigDto) IsSharedLinkKeySet() bool {
	if o != nil && o.SharedLinkKey.IsSet() {
		return true
	}

	return false
}

// SetSharedLinkKey gets a reference to the given NullableString and assigns it to the SharedLinkKey field.
func (o *DocumentConfigDto) SetSharedLinkKey(v string) {
	o.SharedLinkKey.Set(&v)
}
// SetSharedLinkKeyNil sets the value for SharedLinkKey to be an explicit nil
func (o *DocumentConfigDto) SetSharedLinkKeyNil() {
	o.SharedLinkKey.Set(nil)
}

// UnsetSharedLinkKey ensures that no value is present for SharedLinkKey, not even an explicit nil
func (o *DocumentConfigDto) UnsetSharedLinkKey() {
	o.SharedLinkKey.Unset()
}

// GetReferenceData returns the ReferenceData field value if set, zero value otherwise.
func (o *DocumentConfigDto) GetReferenceData() FileReferenceData {
	if o == nil || IsNil(o.ReferenceData) {
		var ret FileReferenceData
		return ret
	}
	return *o.ReferenceData
}

// GetReferenceDataOk returns a tuple with the ReferenceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentConfigDto) GetReferenceDataOk() (*FileReferenceData, bool) {
	if o == nil || IsNil(o.ReferenceData) {
		return nil, false
	}
	return o.ReferenceData, true
}

// HasReferenceData returns a boolean if a field has been set.
func (o *DocumentConfigDto) IsReferenceDataSet() bool {
	if o != nil && !IsNil(o.ReferenceData) {
		return true
	}

	return false
}

// SetReferenceData gets a reference to the given FileReferenceData and assigns it to the ReferenceData field.
func (o *DocumentConfigDto) SetReferenceData(v FileReferenceData) {
	o.ReferenceData = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentConfigDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentConfigDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *DocumentConfigDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *DocumentConfigDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *DocumentConfigDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *DocumentConfigDto) UnsetTitle() {
	o.Title.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentConfigDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentConfigDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *DocumentConfigDto) IsUrlSet() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *DocumentConfigDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *DocumentConfigDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *DocumentConfigDto) UnsetUrl() {
	o.Url.Unset()
}

// GetIsForm returns the IsForm field value if set, zero value otherwise.
func (o *DocumentConfigDto) GetIsForm() bool {
	if o == nil || IsNil(o.IsForm) {
		var ret bool
		return ret
	}
	return *o.IsForm
}

// GetIsFormOk returns a tuple with the IsForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentConfigDto) GetIsFormOk() (*bool, bool) {
	if o == nil || IsNil(o.IsForm) {
		return nil, false
	}
	return o.IsForm, true
}

// HasIsForm returns a boolean if a field has been set.
func (o *DocumentConfigDto) IsIsFormSet() bool {
	if o != nil && !IsNil(o.IsForm) {
		return true
	}

	return false
}

// SetIsForm gets a reference to the given bool and assigns it to the IsForm field.
func (o *DocumentConfigDto) SetIsForm(v bool) {
	o.IsForm = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *DocumentConfigDto) GetOptions() Options {
	if o == nil || IsNil(o.Options) {
		var ret Options
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentConfigDto) GetOptionsOk() (*Options, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *DocumentConfigDto) IsOptionsSet() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given Options and assigns it to the Options field.
func (o *DocumentConfigDto) SetOptions(v Options) {
	o.Options = &v
}

func (o DocumentConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentConfigDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FileType.IsSet() {
		toSerialize["fileType"] = o.FileType.Get()
	}
	if !IsNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	if !IsNil(o.IsLinkedForMe) {
		toSerialize["isLinkedForMe"] = o.IsLinkedForMe
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if o.SharedLinkParam.IsSet() {
		toSerialize["sharedLinkParam"] = o.SharedLinkParam.Get()
	}
	if o.SharedLinkKey.IsSet() {
		toSerialize["sharedLinkKey"] = o.SharedLinkKey.Get()
	}
	if !IsNil(o.ReferenceData) {
		toSerialize["referenceData"] = o.ReferenceData
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if !IsNil(o.IsForm) {
		toSerialize["isForm"] = o.IsForm
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableDocumentConfigDto struct {
	value *DocumentConfigDto
	isSet bool
}

func (v NullableDocumentConfigDto) Get() *DocumentConfigDto {
	return v.value
}

func (v *NullableDocumentConfigDto) Set(val *DocumentConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentConfigDto(val *DocumentConfigDto) *NullableDocumentConfigDto {
	return &NullableDocumentConfigDto{value: val, isSet: true}
}

func (v NullableDocumentConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

