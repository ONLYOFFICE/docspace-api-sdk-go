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

// checks if the MigratingApiFiles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MigratingApiFiles{}

// MigratingApiFiles struct for MigratingApiFiles
type MigratingApiFiles struct {
	FoldersCount *int32 `json:"foldersCount,omitempty"`
	FilesCount *int32 `json:"filesCount,omitempty"`
	BytesTotal *int64 `json:"bytesTotal,omitempty"`
}

// NewMigratingApiFiles instantiates a new MigratingApiFiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigratingApiFiles() *MigratingApiFiles {
	this := MigratingApiFiles{}
	return &this
}

// NewMigratingApiFilesWithDefaults instantiates a new MigratingApiFiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigratingApiFilesWithDefaults() *MigratingApiFiles {
	this := MigratingApiFiles{}
	return &this
}

// GetFoldersCount returns the FoldersCount field value if set, zero value otherwise.
func (o *MigratingApiFiles) GetFoldersCount() int32 {
	if o == nil || IsNil(o.FoldersCount) {
		var ret int32
		return ret
	}
	return *o.FoldersCount
}

// GetFoldersCountOk returns a tuple with the FoldersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigratingApiFiles) GetFoldersCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FoldersCount) {
		return nil, false
	}
	return o.FoldersCount, true
}

// HasFoldersCount returns a boolean if a field has been set.
func (o *MigratingApiFiles) IsFoldersCountSet() bool {
	if o != nil && !IsNil(o.FoldersCount) {
		return true
	}

	return false
}

// SetFoldersCount gets a reference to the given int32 and assigns it to the FoldersCount field.
func (o *MigratingApiFiles) SetFoldersCount(v int32) {
	o.FoldersCount = &v
}

// GetFilesCount returns the FilesCount field value if set, zero value otherwise.
func (o *MigratingApiFiles) GetFilesCount() int32 {
	if o == nil || IsNil(o.FilesCount) {
		var ret int32
		return ret
	}
	return *o.FilesCount
}

// GetFilesCountOk returns a tuple with the FilesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigratingApiFiles) GetFilesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FilesCount) {
		return nil, false
	}
	return o.FilesCount, true
}

// HasFilesCount returns a boolean if a field has been set.
func (o *MigratingApiFiles) IsFilesCountSet() bool {
	if o != nil && !IsNil(o.FilesCount) {
		return true
	}

	return false
}

// SetFilesCount gets a reference to the given int32 and assigns it to the FilesCount field.
func (o *MigratingApiFiles) SetFilesCount(v int32) {
	o.FilesCount = &v
}

// GetBytesTotal returns the BytesTotal field value if set, zero value otherwise.
func (o *MigratingApiFiles) GetBytesTotal() int64 {
	if o == nil || IsNil(o.BytesTotal) {
		var ret int64
		return ret
	}
	return *o.BytesTotal
}

// GetBytesTotalOk returns a tuple with the BytesTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigratingApiFiles) GetBytesTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.BytesTotal) {
		return nil, false
	}
	return o.BytesTotal, true
}

// HasBytesTotal returns a boolean if a field has been set.
func (o *MigratingApiFiles) IsBytesTotalSet() bool {
	if o != nil && !IsNil(o.BytesTotal) {
		return true
	}

	return false
}

// SetBytesTotal gets a reference to the given int64 and assigns it to the BytesTotal field.
func (o *MigratingApiFiles) SetBytesTotal(v int64) {
	o.BytesTotal = &v
}

func (o MigratingApiFiles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MigratingApiFiles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FoldersCount) {
		toSerialize["foldersCount"] = o.FoldersCount
	}
	if !IsNil(o.FilesCount) {
		toSerialize["filesCount"] = o.FilesCount
	}
	if !IsNil(o.BytesTotal) {
		toSerialize["bytesTotal"] = o.BytesTotal
	}
	return toSerialize, nil
}

type NullableMigratingApiFiles struct {
	value *MigratingApiFiles
	isSet bool
}

func (v NullableMigratingApiFiles) Get() *MigratingApiFiles {
	return v.value
}

func (v *NullableMigratingApiFiles) Set(val *MigratingApiFiles) {
	v.value = val
	v.isSet = true
}

func (v NullableMigratingApiFiles) IsSet() bool {
	return v.isSet
}

func (v *NullableMigratingApiFiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigratingApiFiles(val *MigratingApiFiles) *NullableMigratingApiFiles {
	return &NullableMigratingApiFiles{value: val, isSet: true}
}

func (v NullableMigratingApiFiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigratingApiFiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

