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

// checks if the ScheduleDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduleDto{}

// ScheduleDto The backup schedule parameters.
type ScheduleDto struct {
	// The backup storage type.
	StorageType BackupStorageType `json:"storageType"`
	// The backup storage parameters.
	StorageParams map[string]*string `json:"storageParams"`
	// The backup cron parameters.
	CronParams CronParams `json:"cronParams"`
	// The maximum number of the stored backup copies.
	BackupsStored NullableInt32 `json:"backupsStored,omitempty"`
	// The date and time when the last backup was reated.
	LastBackupTime time.Time `json:"lastBackupTime"`
	// Specifies if a dump will be created or not.
	Dump bool `json:"dump"`
}

type _ScheduleDto ScheduleDto

// NewScheduleDto instantiates a new ScheduleDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleDto(storageType BackupStorageType, storageParams map[string]*string, cronParams CronParams, lastBackupTime time.Time, dump bool) *ScheduleDto {
	this := ScheduleDto{}
	this.StorageType = storageType
	this.StorageParams = storageParams
	this.CronParams = cronParams
	this.LastBackupTime = lastBackupTime
	this.Dump = dump
	return &this
}

// NewScheduleDtoWithDefaults instantiates a new ScheduleDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleDtoWithDefaults() *ScheduleDto {
	this := ScheduleDto{}
	return &this
}

// GetStorageType returns the StorageType field value
func (o *ScheduleDto) GetStorageType() BackupStorageType {
	if o == nil {
		var ret BackupStorageType
		return ret
	}

	return o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value
// and a boolean to check if the value has been set.
func (o *ScheduleDto) GetStorageTypeOk() (*BackupStorageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageType, true
}

// SetStorageType sets field value
func (o *ScheduleDto) SetStorageType(v BackupStorageType) {
	o.StorageType = v
}

// GetStorageParams returns the StorageParams field value
func (o *ScheduleDto) GetStorageParams() map[string]*string {
	if o == nil {
		var ret map[string]*string
		return ret
	}

	return o.StorageParams
}

// GetStorageParamsOk returns a tuple with the StorageParams field value
// and a boolean to check if the value has been set.
func (o *ScheduleDto) GetStorageParamsOk() (map[string]*string, bool) {
	if o == nil {
		return map[string]*string{}, false
	}
	return o.StorageParams, true
}

// SetStorageParams sets field value
func (o *ScheduleDto) SetStorageParams(v map[string]*string) {
	o.StorageParams = v
}

// GetCronParams returns the CronParams field value
func (o *ScheduleDto) GetCronParams() CronParams {
	if o == nil {
		var ret CronParams
		return ret
	}

	return o.CronParams
}

// GetCronParamsOk returns a tuple with the CronParams field value
// and a boolean to check if the value has been set.
func (o *ScheduleDto) GetCronParamsOk() (*CronParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CronParams, true
}

// SetCronParams sets field value
func (o *ScheduleDto) SetCronParams(v CronParams) {
	o.CronParams = v
}

// GetBackupsStored returns the BackupsStored field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScheduleDto) GetBackupsStored() int32 {
	if o == nil || IsNil(o.BackupsStored.Get()) {
		var ret int32
		return ret
	}
	return *o.BackupsStored.Get()
}

// GetBackupsStoredOk returns a tuple with the BackupsStored field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScheduleDto) GetBackupsStoredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupsStored.Get(), o.BackupsStored.IsSet()
}

// HasBackupsStored returns a boolean if a field has been set.
func (o *ScheduleDto) IsBackupsStoredSet() bool {
	if o != nil && o.BackupsStored.IsSet() {
		return true
	}

	return false
}

// SetBackupsStored gets a reference to the given NullableInt32 and assigns it to the BackupsStored field.
func (o *ScheduleDto) SetBackupsStored(v int32) {
	o.BackupsStored.Set(&v)
}
// SetBackupsStoredNil sets the value for BackupsStored to be an explicit nil
func (o *ScheduleDto) SetBackupsStoredNil() {
	o.BackupsStored.Set(nil)
}

// UnsetBackupsStored ensures that no value is present for BackupsStored, not even an explicit nil
func (o *ScheduleDto) UnsetBackupsStored() {
	o.BackupsStored.Unset()
}

// GetLastBackupTime returns the LastBackupTime field value
func (o *ScheduleDto) GetLastBackupTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastBackupTime
}

// GetLastBackupTimeOk returns a tuple with the LastBackupTime field value
// and a boolean to check if the value has been set.
func (o *ScheduleDto) GetLastBackupTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastBackupTime, true
}

// SetLastBackupTime sets field value
func (o *ScheduleDto) SetLastBackupTime(v time.Time) {
	o.LastBackupTime = v
}

// GetDump returns the Dump field value
func (o *ScheduleDto) GetDump() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Dump
}

// GetDumpOk returns a tuple with the Dump field value
// and a boolean to check if the value has been set.
func (o *ScheduleDto) GetDumpOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dump, true
}

// SetDump sets field value
func (o *ScheduleDto) SetDump(v bool) {
	o.Dump = v
}

func (o ScheduleDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduleDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["storageType"] = o.StorageType
	toSerialize["storageParams"] = o.StorageParams
	toSerialize["cronParams"] = o.CronParams
	if o.BackupsStored.IsSet() {
		toSerialize["backupsStored"] = o.BackupsStored.Get()
	}
	toSerialize["lastBackupTime"] = o.LastBackupTime
	toSerialize["dump"] = o.Dump
	return toSerialize, nil
}

func (o *ScheduleDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"storageType",
		"storageParams",
		"cronParams",
		"lastBackupTime",
		"dump",
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

	varScheduleDto := _ScheduleDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varScheduleDto)

	if err != nil {
		return err
	}

	*o = ScheduleDto(varScheduleDto)

	return err
}

type NullableScheduleDto struct {
	value *ScheduleDto
	isSet bool
}

func (v NullableScheduleDto) Get() *ScheduleDto {
	return v.value
}

func (v *NullableScheduleDto) Set(val *ScheduleDto) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleDto) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleDto(val *ScheduleDto) *NullableScheduleDto {
	return &NullableScheduleDto{value: val, isSet: true}
}

func (v NullableScheduleDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

