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

// checks if the StartBackupRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartBackupRequestDto{}

// StartBackupRequestDto The request parameters for starting a backup.
type StartBackupRequestDto struct {
	// The storage the archive is written to. It defaults to `Documents`, and it decides which keys  `storageParams` has to carry.
	StorageType *BackupStorageType `json:"storageType,omitempty"`
	// The settings of the chosen storage, as an array of key and value pairs. `Documents` needs an integer  `folderId`, `ThridpartyDocuments` a provider-specific non-integer `folderId`, `Local` a `filePath`,  `ThirdPartyConsumer` a `module` plus the settings of that consumer, and `DataStore` none. The  `subdir` key is added by the operation itself and must not be sent.
	StorageParams []ItemKeyValuePairObjectObject `json:"storageParams,omitempty"`
	// Backs up the whole server rather than this one portal. It requires the space access permission and  works on a standalone installation only.
	Dump *bool `json:"dump,omitempty"`
}

// NewStartBackupRequestDto instantiates a new StartBackupRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartBackupRequestDto() *StartBackupRequestDto {
	this := StartBackupRequestDto{}
	return &this
}

// NewStartBackupRequestDtoWithDefaults instantiates a new StartBackupRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartBackupRequestDtoWithDefaults() *StartBackupRequestDto {
	this := StartBackupRequestDto{}
	return &this
}

// GetStorageType returns the StorageType field value if set, zero value otherwise.
func (o *StartBackupRequestDto) GetStorageType() BackupStorageType {
	if o == nil || IsNil(o.StorageType) {
		var ret BackupStorageType
		return ret
	}
	return *o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartBackupRequestDto) GetStorageTypeOk() (*BackupStorageType, bool) {
	if o == nil || IsNil(o.StorageType) {
		return nil, false
	}
	return o.StorageType, true
}

// HasStorageType returns a boolean if a field has been set.
func (o *StartBackupRequestDto) IsStorageTypeSet() bool {
	if o != nil && !IsNil(o.StorageType) {
		return true
	}

	return false
}

// SetStorageType gets a reference to the given BackupStorageType and assigns it to the StorageType field.
func (o *StartBackupRequestDto) SetStorageType(v BackupStorageType) {
	o.StorageType = &v
}

// GetStorageParams returns the StorageParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StartBackupRequestDto) GetStorageParams() []ItemKeyValuePairObjectObject {
	if o == nil {
		var ret []ItemKeyValuePairObjectObject
		return ret
	}
	return o.StorageParams
}

// GetStorageParamsOk returns a tuple with the StorageParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StartBackupRequestDto) GetStorageParamsOk() ([]ItemKeyValuePairObjectObject, bool) {
	if o == nil || IsNil(o.StorageParams) {
		return nil, false
	}
	return o.StorageParams, true
}

// HasStorageParams returns a boolean if a field has been set.
func (o *StartBackupRequestDto) IsStorageParamsSet() bool {
	if o != nil && !IsNil(o.StorageParams) {
		return true
	}

	return false
}

// SetStorageParams gets a reference to the given []ItemKeyValuePairObjectObject and assigns it to the StorageParams field.
func (o *StartBackupRequestDto) SetStorageParams(v []ItemKeyValuePairObjectObject) {
	o.StorageParams = v
}

// GetDump returns the Dump field value if set, zero value otherwise.
func (o *StartBackupRequestDto) GetDump() bool {
	if o == nil || IsNil(o.Dump) {
		var ret bool
		return ret
	}
	return *o.Dump
}

// GetDumpOk returns a tuple with the Dump field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartBackupRequestDto) GetDumpOk() (*bool, bool) {
	if o == nil || IsNil(o.Dump) {
		return nil, false
	}
	return o.Dump, true
}

// HasDump returns a boolean if a field has been set.
func (o *StartBackupRequestDto) IsDumpSet() bool {
	if o != nil && !IsNil(o.Dump) {
		return true
	}

	return false
}

// SetDump gets a reference to the given bool and assigns it to the Dump field.
func (o *StartBackupRequestDto) SetDump(v bool) {
	o.Dump = &v
}

func (o StartBackupRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartBackupRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StorageType) {
		toSerialize["storageType"] = o.StorageType
	}
	if o.StorageParams != nil {
		toSerialize["storageParams"] = o.StorageParams
	}
	if !IsNil(o.Dump) {
		toSerialize["dump"] = o.Dump
	}
	return toSerialize, nil
}

type NullableStartBackupRequestDto struct {
	value *StartBackupRequestDto
	isSet bool
}

func (v NullableStartBackupRequestDto) Get() *StartBackupRequestDto {
	return v.value
}

func (v *NullableStartBackupRequestDto) Set(val *StartBackupRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStartBackupRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStartBackupRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartBackupRequestDto(val *StartBackupRequestDto) *NullableStartBackupRequestDto {
	return &NullableStartBackupRequestDto{value: val, isSet: true}
}

func (v NullableStartBackupRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartBackupRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

