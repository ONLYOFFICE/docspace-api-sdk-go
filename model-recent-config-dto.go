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

// checks if the RecentConfigDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecentConfigDto{}

// RecentConfigDto One entry of the recent-documents list the editor offers.
type RecentConfigDto struct {
	// The folder shown next to the entry, as a readable name rather than an id.
	Folder NullableString `json:"folder,omitempty"`
	// The name shown for the entry.
	Title NullableString `json:"title,omitempty"`
	// Where the entry opens.
	Url NullableString `json:"url,omitempty"`
}

// NewRecentConfigDto instantiates a new RecentConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecentConfigDto() *RecentConfigDto {
	this := RecentConfigDto{}
	return &this
}

// NewRecentConfigDtoWithDefaults instantiates a new RecentConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecentConfigDtoWithDefaults() *RecentConfigDto {
	this := RecentConfigDto{}
	return &this
}

// GetFolder returns the Folder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecentConfigDto) GetFolder() string {
	if o == nil || IsNil(o.Folder.Get()) {
		var ret string
		return ret
	}
	return *o.Folder.Get()
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecentConfigDto) GetFolderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Folder.Get(), o.Folder.IsSet()
}

// HasFolder returns a boolean if a field has been set.
func (o *RecentConfigDto) IsFolderSet() bool {
	if o != nil && o.Folder.IsSet() {
		return true
	}

	return false
}

// SetFolder gets a reference to the given NullableString and assigns it to the Folder field.
func (o *RecentConfigDto) SetFolder(v string) {
	o.Folder.Set(&v)
}
// SetFolderNil sets the value for Folder to be an explicit nil
func (o *RecentConfigDto) SetFolderNil() {
	o.Folder.Set(nil)
}

// UnsetFolder ensures that no value is present for Folder, not even an explicit nil
func (o *RecentConfigDto) UnsetFolder() {
	o.Folder.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecentConfigDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecentConfigDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *RecentConfigDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *RecentConfigDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *RecentConfigDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *RecentConfigDto) UnsetTitle() {
	o.Title.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecentConfigDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecentConfigDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *RecentConfigDto) IsUrlSet() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *RecentConfigDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *RecentConfigDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *RecentConfigDto) UnsetUrl() {
	o.Url.Unset()
}

func (o RecentConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecentConfigDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Folder.IsSet() {
		toSerialize["folder"] = o.Folder.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return toSerialize, nil
}

type NullableRecentConfigDto struct {
	value *RecentConfigDto
	isSet bool
}

func (v NullableRecentConfigDto) Get() *RecentConfigDto {
	return v.value
}

func (v *NullableRecentConfigDto) Set(val *RecentConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRecentConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRecentConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecentConfigDto(val *RecentConfigDto) *NullableRecentConfigDto {
	return &NullableRecentConfigDto{value: val, isSet: true}
}

func (v NullableRecentConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecentConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

