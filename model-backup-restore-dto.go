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
	"bytes"
	"fmt"
)

// checks if the BackupRestoreDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupRestoreDto{}

// BackupRestoreDto The request parameters for restoring a portal from a backup.
type BackupRestoreDto struct {
	// The ID of the backup to restore from, as listed by `GET api/2.0/backup/getbackuphistory`. Send  anything that is not a GUID to restore from a file given by `storageParams` instead; an all-zero GUID  selects neither, because it parses as a GUID and then matches no record.
	BackupId NullableString `json:"backupId"`
	// The storage the archive is read from. It defaults to `Documents` and is only used when `backupId` is  not a GUID, because a known backup carries the storage of its own record.
	StorageType *BackupStorageType `json:"storageType,omitempty"`
	// The location of the archive, as an array of key and value pairs. The key read here is `filePath` -  not the `folderId` a backup is started with - and it holds a file ID for `Documents`, a  provider-specific file ID for `ThridpartyDocuments` and a path on the server for `Local`. It is only  used when `backupId` is not a GUID.
	StorageParams []ItemKeyValuePairObjectObject `json:"storageParams,omitempty"`
	// Chooses who is emailed when the restoring starts and when it finishes: every active user of the  portal when true, and its owner alone when false. Mail goes only to accounts that have been  activated, so this decides the audience rather than whether anybody is notified at all.
	Notify *bool `json:"notify,omitempty"`
	// Restores the whole server rather than this one portal. It requires the space access permission.
	Dump *bool `json:"dump,omitempty"`
}

type _BackupRestoreDto BackupRestoreDto

// NewBackupRestoreDto instantiates a new BackupRestoreDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupRestoreDto(backupId NullableString) *BackupRestoreDto {
	this := BackupRestoreDto{}
	this.BackupId = backupId
	return &this
}

// NewBackupRestoreDtoWithDefaults instantiates a new BackupRestoreDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupRestoreDtoWithDefaults() *BackupRestoreDto {
	this := BackupRestoreDto{}
	return &this
}

// GetBackupId returns the BackupId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BackupRestoreDto) GetBackupId() string {
	if o == nil || o.BackupId.Get() == nil {
		var ret string
		return ret
	}

	return *o.BackupId.Get()
}

// GetBackupIdOk returns a tuple with the BackupId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreDto) GetBackupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupId.Get(), o.BackupId.IsSet()
}

// SetBackupId sets field value
func (o *BackupRestoreDto) SetBackupId(v string) {
	o.BackupId.Set(&v)
}

// GetStorageType returns the StorageType field value if set, zero value otherwise.
func (o *BackupRestoreDto) GetStorageType() BackupStorageType {
	if o == nil || IsNil(o.StorageType) {
		var ret BackupStorageType
		return ret
	}
	return *o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreDto) GetStorageTypeOk() (*BackupStorageType, bool) {
	if o == nil || IsNil(o.StorageType) {
		return nil, false
	}
	return o.StorageType, true
}

// HasStorageType returns a boolean if a field has been set.
func (o *BackupRestoreDto) IsStorageTypeSet() bool {
	if o != nil && !IsNil(o.StorageType) {
		return true
	}

	return false
}

// SetStorageType gets a reference to the given BackupStorageType and assigns it to the StorageType field.
func (o *BackupRestoreDto) SetStorageType(v BackupStorageType) {
	o.StorageType = &v
}

// GetStorageParams returns the StorageParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupRestoreDto) GetStorageParams() []ItemKeyValuePairObjectObject {
	if o == nil {
		var ret []ItemKeyValuePairObjectObject
		return ret
	}
	return o.StorageParams
}

// GetStorageParamsOk returns a tuple with the StorageParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupRestoreDto) GetStorageParamsOk() ([]ItemKeyValuePairObjectObject, bool) {
	if o == nil || IsNil(o.StorageParams) {
		return nil, false
	}
	return o.StorageParams, true
}

// HasStorageParams returns a boolean if a field has been set.
func (o *BackupRestoreDto) IsStorageParamsSet() bool {
	if o != nil && !IsNil(o.StorageParams) {
		return true
	}

	return false
}

// SetStorageParams gets a reference to the given []ItemKeyValuePairObjectObject and assigns it to the StorageParams field.
func (o *BackupRestoreDto) SetStorageParams(v []ItemKeyValuePairObjectObject) {
	o.StorageParams = v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *BackupRestoreDto) GetNotify() bool {
	if o == nil || IsNil(o.Notify) {
		var ret bool
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreDto) GetNotifyOk() (*bool, bool) {
	if o == nil || IsNil(o.Notify) {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *BackupRestoreDto) IsNotifySet() bool {
	if o != nil && !IsNil(o.Notify) {
		return true
	}

	return false
}

// SetNotify gets a reference to the given bool and assigns it to the Notify field.
func (o *BackupRestoreDto) SetNotify(v bool) {
	o.Notify = &v
}

// GetDump returns the Dump field value if set, zero value otherwise.
func (o *BackupRestoreDto) GetDump() bool {
	if o == nil || IsNil(o.Dump) {
		var ret bool
		return ret
	}
	return *o.Dump
}

// GetDumpOk returns a tuple with the Dump field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreDto) GetDumpOk() (*bool, bool) {
	if o == nil || IsNil(o.Dump) {
		return nil, false
	}
	return o.Dump, true
}

// HasDump returns a boolean if a field has been set.
func (o *BackupRestoreDto) IsDumpSet() bool {
	if o != nil && !IsNil(o.Dump) {
		return true
	}

	return false
}

// SetDump gets a reference to the given bool and assigns it to the Dump field.
func (o *BackupRestoreDto) SetDump(v bool) {
	o.Dump = &v
}

func (o BackupRestoreDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupRestoreDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["backupId"] = o.BackupId.Get()
	if !IsNil(o.StorageType) {
		toSerialize["storageType"] = o.StorageType
	}
	if o.StorageParams != nil {
		toSerialize["storageParams"] = o.StorageParams
	}
	if !IsNil(o.Notify) {
		toSerialize["notify"] = o.Notify
	}
	if !IsNil(o.Dump) {
		toSerialize["dump"] = o.Dump
	}
	return toSerialize, nil
}

func (o *BackupRestoreDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"backupId",
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

	varBackupRestoreDto := _BackupRestoreDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varBackupRestoreDto)

	if err != nil {
		return err
	}

	*o = BackupRestoreDto(varBackupRestoreDto)

	return err
}

type NullableBackupRestoreDto struct {
	value *BackupRestoreDto
	isSet bool
}

func (v NullableBackupRestoreDto) Get() *BackupRestoreDto {
	return v.value
}

func (v *NullableBackupRestoreDto) Set(val *BackupRestoreDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRestoreDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRestoreDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRestoreDto(val *BackupRestoreDto) *NullableBackupRestoreDto {
	return &NullableBackupRestoreDto{value: val, isSet: true}
}

func (v NullableBackupRestoreDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRestoreDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

