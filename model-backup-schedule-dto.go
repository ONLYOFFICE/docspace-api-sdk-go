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

// checks if the BackupScheduleDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupScheduleDto{}

// BackupScheduleDto The backup schedule parameters.
type BackupScheduleDto struct {
	// The backup storage type.
	StorageType *BackupStorageType `json:"storageType,omitempty"`
	// The backup storage parameters.
	StorageParams []ItemKeyValuePairObjectObject `json:"storageParams,omitempty"`
	// The maximum number of the stored backup copies.
	BackupsStored NullableInt32 `json:"backupsStored,omitempty"`
	// The backup cron parameters.
	CronParams *Cron `json:"cronParams,omitempty"`
	// Specifies if a dump will be created or not.
	Dump *bool `json:"dump,omitempty"`
}

// NewBackupScheduleDto instantiates a new BackupScheduleDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupScheduleDto() *BackupScheduleDto {
	this := BackupScheduleDto{}
	return &this
}

// NewBackupScheduleDtoWithDefaults instantiates a new BackupScheduleDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupScheduleDtoWithDefaults() *BackupScheduleDto {
	this := BackupScheduleDto{}
	return &this
}

// GetStorageType returns the StorageType field value if set, zero value otherwise.
func (o *BackupScheduleDto) GetStorageType() BackupStorageType {
	if o == nil || IsNil(o.StorageType) {
		var ret BackupStorageType
		return ret
	}
	return *o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupScheduleDto) GetStorageTypeOk() (*BackupStorageType, bool) {
	if o == nil || IsNil(o.StorageType) {
		return nil, false
	}
	return o.StorageType, true
}

// HasStorageType returns a boolean if a field has been set.
func (o *BackupScheduleDto) IsStorageTypeSet() bool {
	if o != nil && !IsNil(o.StorageType) {
		return true
	}

	return false
}

// SetStorageType gets a reference to the given BackupStorageType and assigns it to the StorageType field.
func (o *BackupScheduleDto) SetStorageType(v BackupStorageType) {
	o.StorageType = &v
}

// GetStorageParams returns the StorageParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupScheduleDto) GetStorageParams() []ItemKeyValuePairObjectObject {
	if o == nil {
		var ret []ItemKeyValuePairObjectObject
		return ret
	}
	return o.StorageParams
}

// GetStorageParamsOk returns a tuple with the StorageParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupScheduleDto) GetStorageParamsOk() ([]ItemKeyValuePairObjectObject, bool) {
	if o == nil || IsNil(o.StorageParams) {
		return nil, false
	}
	return o.StorageParams, true
}

// HasStorageParams returns a boolean if a field has been set.
func (o *BackupScheduleDto) IsStorageParamsSet() bool {
	if o != nil && !IsNil(o.StorageParams) {
		return true
	}

	return false
}

// SetStorageParams gets a reference to the given []ItemKeyValuePairObjectObject and assigns it to the StorageParams field.
func (o *BackupScheduleDto) SetStorageParams(v []ItemKeyValuePairObjectObject) {
	o.StorageParams = v
}

// GetBackupsStored returns the BackupsStored field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupScheduleDto) GetBackupsStored() int32 {
	if o == nil || IsNil(o.BackupsStored.Get()) {
		var ret int32
		return ret
	}
	return *o.BackupsStored.Get()
}

// GetBackupsStoredOk returns a tuple with the BackupsStored field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupScheduleDto) GetBackupsStoredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupsStored.Get(), o.BackupsStored.IsSet()
}

// HasBackupsStored returns a boolean if a field has been set.
func (o *BackupScheduleDto) IsBackupsStoredSet() bool {
	if o != nil && o.BackupsStored.IsSet() {
		return true
	}

	return false
}

// SetBackupsStored gets a reference to the given NullableInt32 and assigns it to the BackupsStored field.
func (o *BackupScheduleDto) SetBackupsStored(v int32) {
	o.BackupsStored.Set(&v)
}
// SetBackupsStoredNil sets the value for BackupsStored to be an explicit nil
func (o *BackupScheduleDto) SetBackupsStoredNil() {
	o.BackupsStored.Set(nil)
}

// UnsetBackupsStored ensures that no value is present for BackupsStored, not even an explicit nil
func (o *BackupScheduleDto) UnsetBackupsStored() {
	o.BackupsStored.Unset()
}

// GetCronParams returns the CronParams field value if set, zero value otherwise.
func (o *BackupScheduleDto) GetCronParams() Cron {
	if o == nil || IsNil(o.CronParams) {
		var ret Cron
		return ret
	}
	return *o.CronParams
}

// GetCronParamsOk returns a tuple with the CronParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupScheduleDto) GetCronParamsOk() (*Cron, bool) {
	if o == nil || IsNil(o.CronParams) {
		return nil, false
	}
	return o.CronParams, true
}

// HasCronParams returns a boolean if a field has been set.
func (o *BackupScheduleDto) IsCronParamsSet() bool {
	if o != nil && !IsNil(o.CronParams) {
		return true
	}

	return false
}

// SetCronParams gets a reference to the given Cron and assigns it to the CronParams field.
func (o *BackupScheduleDto) SetCronParams(v Cron) {
	o.CronParams = &v
}

// GetDump returns the Dump field value if set, zero value otherwise.
func (o *BackupScheduleDto) GetDump() bool {
	if o == nil || IsNil(o.Dump) {
		var ret bool
		return ret
	}
	return *o.Dump
}

// GetDumpOk returns a tuple with the Dump field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupScheduleDto) GetDumpOk() (*bool, bool) {
	if o == nil || IsNil(o.Dump) {
		return nil, false
	}
	return o.Dump, true
}

// HasDump returns a boolean if a field has been set.
func (o *BackupScheduleDto) IsDumpSet() bool {
	if o != nil && !IsNil(o.Dump) {
		return true
	}

	return false
}

// SetDump gets a reference to the given bool and assigns it to the Dump field.
func (o *BackupScheduleDto) SetDump(v bool) {
	o.Dump = &v
}

func (o BackupScheduleDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupScheduleDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StorageType) {
		toSerialize["storageType"] = o.StorageType
	}
	if o.StorageParams != nil {
		toSerialize["storageParams"] = o.StorageParams
	}
	if o.BackupsStored.IsSet() {
		toSerialize["backupsStored"] = o.BackupsStored.Get()
	}
	if !IsNil(o.CronParams) {
		toSerialize["cronParams"] = o.CronParams
	}
	if !IsNil(o.Dump) {
		toSerialize["dump"] = o.Dump
	}
	return toSerialize, nil
}

type NullableBackupScheduleDto struct {
	value *BackupScheduleDto
	isSet bool
}

func (v NullableBackupScheduleDto) Get() *BackupScheduleDto {
	return v.value
}

func (v *NullableBackupScheduleDto) Set(val *BackupScheduleDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupScheduleDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupScheduleDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupScheduleDto(val *BackupScheduleDto) *NullableBackupScheduleDto {
	return &NullableBackupScheduleDto{value: val, isSet: true}
}

func (v NullableBackupScheduleDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupScheduleDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

