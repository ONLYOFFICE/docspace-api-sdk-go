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

// checks if the InfoConfigDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InfoConfigDto{}

// InfoConfigDto The information config parameters.
type InfoConfigDto struct {
	// Specifies if the file is favorite or not.
	Favorite NullableBool `json:"favorite,omitempty"`
	// The folder of the file.
	Folder NullableString `json:"folder,omitempty"`
	// The file owner.
	Owner NullableString `json:"owner,omitempty"`
	// The sharing settings of the file.
	SharingSettings []AceShortWrapper `json:"sharingSettings,omitempty"`
	// The editor type of the file.
	Type *EditorType `json:"type,omitempty"`
	// The uploaded file.
	Uploaded NullableString `json:"uploaded,omitempty"`
}

// NewInfoConfigDto instantiates a new InfoConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfoConfigDto() *InfoConfigDto {
	this := InfoConfigDto{}
	return &this
}

// NewInfoConfigDtoWithDefaults instantiates a new InfoConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfoConfigDtoWithDefaults() *InfoConfigDto {
	this := InfoConfigDto{}
	return &this
}

// GetFavorite returns the Favorite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InfoConfigDto) GetFavorite() bool {
	if o == nil || IsNil(o.Favorite.Get()) {
		var ret bool
		return ret
	}
	return *o.Favorite.Get()
}

// GetFavoriteOk returns a tuple with the Favorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InfoConfigDto) GetFavoriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Favorite.Get(), o.Favorite.IsSet()
}

// HasFavorite returns a boolean if a field has been set.
func (o *InfoConfigDto) IsFavoriteSet() bool {
	if o != nil && o.Favorite.IsSet() {
		return true
	}

	return false
}

// SetFavorite gets a reference to the given NullableBool and assigns it to the Favorite field.
func (o *InfoConfigDto) SetFavorite(v bool) {
	o.Favorite.Set(&v)
}
// SetFavoriteNil sets the value for Favorite to be an explicit nil
func (o *InfoConfigDto) SetFavoriteNil() {
	o.Favorite.Set(nil)
}

// UnsetFavorite ensures that no value is present for Favorite, not even an explicit nil
func (o *InfoConfigDto) UnsetFavorite() {
	o.Favorite.Unset()
}

// GetFolder returns the Folder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InfoConfigDto) GetFolder() string {
	if o == nil || IsNil(o.Folder.Get()) {
		var ret string
		return ret
	}
	return *o.Folder.Get()
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InfoConfigDto) GetFolderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Folder.Get(), o.Folder.IsSet()
}

// HasFolder returns a boolean if a field has been set.
func (o *InfoConfigDto) IsFolderSet() bool {
	if o != nil && o.Folder.IsSet() {
		return true
	}

	return false
}

// SetFolder gets a reference to the given NullableString and assigns it to the Folder field.
func (o *InfoConfigDto) SetFolder(v string) {
	o.Folder.Set(&v)
}
// SetFolderNil sets the value for Folder to be an explicit nil
func (o *InfoConfigDto) SetFolderNil() {
	o.Folder.Set(nil)
}

// UnsetFolder ensures that no value is present for Folder, not even an explicit nil
func (o *InfoConfigDto) UnsetFolder() {
	o.Folder.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InfoConfigDto) GetOwner() string {
	if o == nil || IsNil(o.Owner.Get()) {
		var ret string
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InfoConfigDto) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *InfoConfigDto) IsOwnerSet() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableString and assigns it to the Owner field.
func (o *InfoConfigDto) SetOwner(v string) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *InfoConfigDto) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *InfoConfigDto) UnsetOwner() {
	o.Owner.Unset()
}

// GetSharingSettings returns the SharingSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InfoConfigDto) GetSharingSettings() []AceShortWrapper {
	if o == nil {
		var ret []AceShortWrapper
		return ret
	}
	return o.SharingSettings
}

// GetSharingSettingsOk returns a tuple with the SharingSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InfoConfigDto) GetSharingSettingsOk() ([]AceShortWrapper, bool) {
	if o == nil || IsNil(o.SharingSettings) {
		return nil, false
	}
	return o.SharingSettings, true
}

// HasSharingSettings returns a boolean if a field has been set.
func (o *InfoConfigDto) IsSharingSettingsSet() bool {
	if o != nil && !IsNil(o.SharingSettings) {
		return true
	}

	return false
}

// SetSharingSettings gets a reference to the given []AceShortWrapper and assigns it to the SharingSettings field.
func (o *InfoConfigDto) SetSharingSettings(v []AceShortWrapper) {
	o.SharingSettings = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InfoConfigDto) GetType() EditorType {
	if o == nil || IsNil(o.Type) {
		var ret EditorType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfoConfigDto) GetTypeOk() (*EditorType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InfoConfigDto) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EditorType and assigns it to the Type field.
func (o *InfoConfigDto) SetType(v EditorType) {
	o.Type = &v
}

// GetUploaded returns the Uploaded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InfoConfigDto) GetUploaded() string {
	if o == nil || IsNil(o.Uploaded.Get()) {
		var ret string
		return ret
	}
	return *o.Uploaded.Get()
}

// GetUploadedOk returns a tuple with the Uploaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InfoConfigDto) GetUploadedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uploaded.Get(), o.Uploaded.IsSet()
}

// HasUploaded returns a boolean if a field has been set.
func (o *InfoConfigDto) IsUploadedSet() bool {
	if o != nil && o.Uploaded.IsSet() {
		return true
	}

	return false
}

// SetUploaded gets a reference to the given NullableString and assigns it to the Uploaded field.
func (o *InfoConfigDto) SetUploaded(v string) {
	o.Uploaded.Set(&v)
}
// SetUploadedNil sets the value for Uploaded to be an explicit nil
func (o *InfoConfigDto) SetUploadedNil() {
	o.Uploaded.Set(nil)
}

// UnsetUploaded ensures that no value is present for Uploaded, not even an explicit nil
func (o *InfoConfigDto) UnsetUploaded() {
	o.Uploaded.Unset()
}

func (o InfoConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InfoConfigDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Favorite.IsSet() {
		toSerialize["favorite"] = o.Favorite.Get()
	}
	if o.Folder.IsSet() {
		toSerialize["folder"] = o.Folder.Get()
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if o.SharingSettings != nil {
		toSerialize["sharingSettings"] = o.SharingSettings
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Uploaded.IsSet() {
		toSerialize["uploaded"] = o.Uploaded.Get()
	}
	return toSerialize, nil
}

type NullableInfoConfigDto struct {
	value *InfoConfigDto
	isSet bool
}

func (v NullableInfoConfigDto) Get() *InfoConfigDto {
	return v.value
}

func (v *NullableInfoConfigDto) Set(val *InfoConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullableInfoConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullableInfoConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfoConfigDto(val *InfoConfigDto) *NullableInfoConfigDto {
	return &NullableInfoConfigDto{value: val, isSet: true}
}

func (v NullableInfoConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfoConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

