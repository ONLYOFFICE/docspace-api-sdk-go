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

// checks if the RecentConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecentConfig{}

// RecentConfig The presence or absence of the documents in the Open Recent... menu option.
type RecentConfig struct {
	// The folder where the document is stored.
	Folder NullableString `json:"folder,omitempty"`
	// The document title that will be displayed in the Open Recent... menu option.
	Title NullableString `json:"title,omitempty"`
	// The absolute URL to the document where it is stored.
	Url NullableString `json:"url,omitempty"`
}

// NewRecentConfig instantiates a new RecentConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecentConfig() *RecentConfig {
	this := RecentConfig{}
	return &this
}

// NewRecentConfigWithDefaults instantiates a new RecentConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecentConfigWithDefaults() *RecentConfig {
	this := RecentConfig{}
	return &this
}

// GetFolder returns the Folder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecentConfig) GetFolder() string {
	if o == nil || IsNil(o.Folder.Get()) {
		var ret string
		return ret
	}
	return *o.Folder.Get()
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecentConfig) GetFolderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Folder.Get(), o.Folder.IsSet()
}

// HasFolder returns a boolean if a field has been set.
func (o *RecentConfig) IsFolderSet() bool {
	if o != nil && o.Folder.IsSet() {
		return true
	}

	return false
}

// SetFolder gets a reference to the given NullableString and assigns it to the Folder field.
func (o *RecentConfig) SetFolder(v string) {
	o.Folder.Set(&v)
}
// SetFolderNil sets the value for Folder to be an explicit nil
func (o *RecentConfig) SetFolderNil() {
	o.Folder.Set(nil)
}

// UnsetFolder ensures that no value is present for Folder, not even an explicit nil
func (o *RecentConfig) UnsetFolder() {
	o.Folder.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecentConfig) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecentConfig) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *RecentConfig) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *RecentConfig) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *RecentConfig) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *RecentConfig) UnsetTitle() {
	o.Title.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecentConfig) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecentConfig) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *RecentConfig) IsUrlSet() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *RecentConfig) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *RecentConfig) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *RecentConfig) UnsetUrl() {
	o.Url.Unset()
}

func (o RecentConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecentConfig) ToMap() (map[string]interface{}, error) {
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

type NullableRecentConfig struct {
	value *RecentConfig
	isSet bool
}

func (v NullableRecentConfig) Get() *RecentConfig {
	return v.value
}

func (v *NullableRecentConfig) Set(val *RecentConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRecentConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRecentConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecentConfig(val *RecentConfig) *NullableRecentConfig {
	return &NullableRecentConfig{value: val, isSet: true}
}

func (v NullableRecentConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecentConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

