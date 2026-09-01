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

// checks if the UpdateFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFile{}

// UpdateFile The parameters for updating a file.
type UpdateFile struct {
	// The file title to update.
	Title NullableString `json:"title,omitempty"`
	// The number of the latest file version.
	LastVersion *int32 `json:"lastVersion,omitempty"`
}

// NewUpdateFile instantiates a new UpdateFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFile() *UpdateFile {
	this := UpdateFile{}
	return &this
}

// NewUpdateFileWithDefaults instantiates a new UpdateFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFileWithDefaults() *UpdateFile {
	this := UpdateFile{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFile) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFile) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *UpdateFile) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *UpdateFile) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *UpdateFile) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *UpdateFile) UnsetTitle() {
	o.Title.Unset()
}

// GetLastVersion returns the LastVersion field value if set, zero value otherwise.
func (o *UpdateFile) GetLastVersion() int32 {
	if o == nil || IsNil(o.LastVersion) {
		var ret int32
		return ret
	}
	return *o.LastVersion
}

// GetLastVersionOk returns a tuple with the LastVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFile) GetLastVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.LastVersion) {
		return nil, false
	}
	return o.LastVersion, true
}

// HasLastVersion returns a boolean if a field has been set.
func (o *UpdateFile) IsLastVersionSet() bool {
	if o != nil && !IsNil(o.LastVersion) {
		return true
	}

	return false
}

// SetLastVersion gets a reference to the given int32 and assigns it to the LastVersion field.
func (o *UpdateFile) SetLastVersion(v int32) {
	o.LastVersion = &v
}

func (o UpdateFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if !IsNil(o.LastVersion) {
		toSerialize["lastVersion"] = o.LastVersion
	}
	return toSerialize, nil
}

type NullableUpdateFile struct {
	value *UpdateFile
	isSet bool
}

func (v NullableUpdateFile) Get() *UpdateFile {
	return v.value
}

func (v *NullableUpdateFile) Set(val *UpdateFile) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFile) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFile(val *UpdateFile) *NullableUpdateFile {
	return &NullableUpdateFile{value: val, isSet: true}
}

func (v NullableUpdateFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

