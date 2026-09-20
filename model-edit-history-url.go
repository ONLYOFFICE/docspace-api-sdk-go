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

// checks if the EditHistoryUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditHistoryUrl{}

// EditHistoryUrl The address, document key and format of the revision a comparison is made against.
type EditHistoryUrl struct {
	// The document key of that revision. When the file has no earlier revision the portal generates a fresh key for  the template it falls back to, so the value is not always one an earlier revision ever had.
	Key NullableString `json:"key,omitempty"`
	// The address that revision's content is served from. It is meant for the editing service and carries its own  key, which is valid for a limited time.
	Url NullableString `json:"url,omitempty"`
	// The format of that revision, as an extension without the leading dot.
	FileType NullableString `json:"fileType,omitempty"`
}

// NewEditHistoryUrl instantiates a new EditHistoryUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditHistoryUrl() *EditHistoryUrl {
	this := EditHistoryUrl{}
	return &this
}

// NewEditHistoryUrlWithDefaults instantiates a new EditHistoryUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditHistoryUrlWithDefaults() *EditHistoryUrl {
	this := EditHistoryUrl{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHistoryUrl) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHistoryUrl) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *EditHistoryUrl) IsKeySet() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *EditHistoryUrl) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *EditHistoryUrl) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *EditHistoryUrl) UnsetKey() {
	o.Key.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHistoryUrl) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHistoryUrl) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *EditHistoryUrl) IsUrlSet() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *EditHistoryUrl) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *EditHistoryUrl) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *EditHistoryUrl) UnsetUrl() {
	o.Url.Unset()
}

// GetFileType returns the FileType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHistoryUrl) GetFileType() string {
	if o == nil || IsNil(o.FileType.Get()) {
		var ret string
		return ret
	}
	return *o.FileType.Get()
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHistoryUrl) GetFileTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileType.Get(), o.FileType.IsSet()
}

// HasFileType returns a boolean if a field has been set.
func (o *EditHistoryUrl) IsFileTypeSet() bool {
	if o != nil && o.FileType.IsSet() {
		return true
	}

	return false
}

// SetFileType gets a reference to the given NullableString and assigns it to the FileType field.
func (o *EditHistoryUrl) SetFileType(v string) {
	o.FileType.Set(&v)
}
// SetFileTypeNil sets the value for FileType to be an explicit nil
func (o *EditHistoryUrl) SetFileTypeNil() {
	o.FileType.Set(nil)
}

// UnsetFileType ensures that no value is present for FileType, not even an explicit nil
func (o *EditHistoryUrl) UnsetFileType() {
	o.FileType.Unset()
}

func (o EditHistoryUrl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditHistoryUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.FileType.IsSet() {
		toSerialize["fileType"] = o.FileType.Get()
	}
	return toSerialize, nil
}

type NullableEditHistoryUrl struct {
	value *EditHistoryUrl
	isSet bool
}

func (v NullableEditHistoryUrl) Get() *EditHistoryUrl {
	return v.value
}

func (v *NullableEditHistoryUrl) Set(val *EditHistoryUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableEditHistoryUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableEditHistoryUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditHistoryUrl(val *EditHistoryUrl) *NullableEditHistoryUrl {
	return &NullableEditHistoryUrl{value: val, isSet: true}
}

func (v NullableEditHistoryUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditHistoryUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

