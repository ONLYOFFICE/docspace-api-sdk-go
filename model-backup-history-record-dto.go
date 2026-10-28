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

// checks if the BackupHistoryRecordDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupHistoryRecordDto{}

// BackupHistoryRecordDto One stored backup of a portal.
type BackupHistoryRecordDto struct {
	// The ID of the backup, which is the same value as the `taskId` the backup was started with. Pass it to  `DELETE api/2.0/backup/deletebackup/{id}` or as the `backupId` of  `POST api/2.0/backup/startrestore`.
	Id string `json:"id"`
	// The name of the stored archive. It is built from the portal alias and the moment the backup started,  or from `workspace` instead of the alias for a backup of the whole server.
	FileName NullableString `json:"fileName"`
	// The storage the archive was written to, reported as a number rather than as a name.
	StorageType BackupStorageType `json:"storageType"`
	// The date and time the backup was stored at, in UTC.
	CreatedOn time.Time `json:"createdOn"`
	// The date and time a background cleaner removes this backup at. Only a backup written to `DataStore`  expires, one day after it was stored; for every other storage type this is `0001-01-01T00:00:00`,  which means the backup is kept until it is deleted by hand or pushed out by the stored-copies limit  of a schedule.
	ExpiresOn time.Time `json:"expiresOn"`
}

type _BackupHistoryRecordDto BackupHistoryRecordDto

// NewBackupHistoryRecordDto instantiates a new BackupHistoryRecordDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupHistoryRecordDto(id string, fileName NullableString, storageType BackupStorageType, createdOn time.Time, expiresOn time.Time) *BackupHistoryRecordDto {
	this := BackupHistoryRecordDto{}
	this.Id = id
	this.FileName = fileName
	this.StorageType = storageType
	this.CreatedOn = createdOn
	this.ExpiresOn = expiresOn
	return &this
}

// NewBackupHistoryRecordDtoWithDefaults instantiates a new BackupHistoryRecordDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupHistoryRecordDtoWithDefaults() *BackupHistoryRecordDto {
	this := BackupHistoryRecordDto{}
	return &this
}

// GetId returns the Id field value
func (o *BackupHistoryRecordDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BackupHistoryRecordDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BackupHistoryRecordDto) SetId(v string) {
	o.Id = v
}

// GetFileName returns the FileName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BackupHistoryRecordDto) GetFileName() string {
	if o == nil || o.FileName.Get() == nil {
		var ret string
		return ret
	}

	return *o.FileName.Get()
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupHistoryRecordDto) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileName.Get(), o.FileName.IsSet()
}

// SetFileName sets field value
func (o *BackupHistoryRecordDto) SetFileName(v string) {
	o.FileName.Set(&v)
}

// GetStorageType returns the StorageType field value
func (o *BackupHistoryRecordDto) GetStorageType() BackupStorageType {
	if o == nil {
		var ret BackupStorageType
		return ret
	}

	return o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value
// and a boolean to check if the value has been set.
func (o *BackupHistoryRecordDto) GetStorageTypeOk() (*BackupStorageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageType, true
}

// SetStorageType sets field value
func (o *BackupHistoryRecordDto) SetStorageType(v BackupStorageType) {
	o.StorageType = v
}

// GetCreatedOn returns the CreatedOn field value
func (o *BackupHistoryRecordDto) GetCreatedOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *BackupHistoryRecordDto) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *BackupHistoryRecordDto) SetCreatedOn(v time.Time) {
	o.CreatedOn = v
}

// GetExpiresOn returns the ExpiresOn field value
func (o *BackupHistoryRecordDto) GetExpiresOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresOn
}

// GetExpiresOnOk returns a tuple with the ExpiresOn field value
// and a boolean to check if the value has been set.
func (o *BackupHistoryRecordDto) GetExpiresOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresOn, true
}

// SetExpiresOn sets field value
func (o *BackupHistoryRecordDto) SetExpiresOn(v time.Time) {
	o.ExpiresOn = v
}

func (o BackupHistoryRecordDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupHistoryRecordDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["fileName"] = o.FileName.Get()
	toSerialize["storageType"] = o.StorageType
	toSerialize["createdOn"] = o.CreatedOn
	toSerialize["expiresOn"] = o.ExpiresOn
	return toSerialize, nil
}

func (o *BackupHistoryRecordDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"fileName",
		"storageType",
		"createdOn",
		"expiresOn",
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

	varBackupHistoryRecordDto := _BackupHistoryRecordDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varBackupHistoryRecordDto)

	if err != nil {
		return err
	}

	*o = BackupHistoryRecordDto(varBackupHistoryRecordDto)

	return err
}

type NullableBackupHistoryRecordDto struct {
	value *BackupHistoryRecordDto
	isSet bool
}

func (v NullableBackupHistoryRecordDto) Get() *BackupHistoryRecordDto {
	return v.value
}

func (v *NullableBackupHistoryRecordDto) Set(val *BackupHistoryRecordDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupHistoryRecordDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupHistoryRecordDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupHistoryRecordDto(val *BackupHistoryRecordDto) *NullableBackupHistoryRecordDto {
	return &NullableBackupHistoryRecordDto{value: val, isSet: true}
}

func (v NullableBackupHistoryRecordDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupHistoryRecordDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

