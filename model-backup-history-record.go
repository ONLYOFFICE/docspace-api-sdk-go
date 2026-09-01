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

// checks if the BackupHistoryRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupHistoryRecord{}

// BackupHistoryRecord The backup history parameters.
type BackupHistoryRecord struct {
	// The backup ID.
	Id string `json:"id"`
	// The backup file name.
	FileName NullableString `json:"fileName"`
	// The backup storage type.
	StorageType BackupStorageType `json:"storageType"`
	// The backup creation date.
	CreatedOn time.Time `json:"createdOn"`
	// The backup expiration date.
	ExpiresOn time.Time `json:"expiresOn"`
}

type _BackupHistoryRecord BackupHistoryRecord

// NewBackupHistoryRecord instantiates a new BackupHistoryRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupHistoryRecord(id string, fileName NullableString, storageType BackupStorageType, createdOn time.Time, expiresOn time.Time) *BackupHistoryRecord {
	this := BackupHistoryRecord{}
	this.Id = id
	this.FileName = fileName
	this.StorageType = storageType
	this.CreatedOn = createdOn
	this.ExpiresOn = expiresOn
	return &this
}

// NewBackupHistoryRecordWithDefaults instantiates a new BackupHistoryRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupHistoryRecordWithDefaults() *BackupHistoryRecord {
	this := BackupHistoryRecord{}
	return &this
}

// GetId returns the Id field value
func (o *BackupHistoryRecord) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BackupHistoryRecord) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BackupHistoryRecord) SetId(v string) {
	o.Id = v
}

// GetFileName returns the FileName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BackupHistoryRecord) GetFileName() string {
	if o == nil || o.FileName.Get() == nil {
		var ret string
		return ret
	}

	return *o.FileName.Get()
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupHistoryRecord) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileName.Get(), o.FileName.IsSet()
}

// SetFileName sets field value
func (o *BackupHistoryRecord) SetFileName(v string) {
	o.FileName.Set(&v)
}

// GetStorageType returns the StorageType field value
func (o *BackupHistoryRecord) GetStorageType() BackupStorageType {
	if o == nil {
		var ret BackupStorageType
		return ret
	}

	return o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value
// and a boolean to check if the value has been set.
func (o *BackupHistoryRecord) GetStorageTypeOk() (*BackupStorageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageType, true
}

// SetStorageType sets field value
func (o *BackupHistoryRecord) SetStorageType(v BackupStorageType) {
	o.StorageType = v
}

// GetCreatedOn returns the CreatedOn field value
func (o *BackupHistoryRecord) GetCreatedOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *BackupHistoryRecord) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *BackupHistoryRecord) SetCreatedOn(v time.Time) {
	o.CreatedOn = v
}

// GetExpiresOn returns the ExpiresOn field value
func (o *BackupHistoryRecord) GetExpiresOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresOn
}

// GetExpiresOnOk returns a tuple with the ExpiresOn field value
// and a boolean to check if the value has been set.
func (o *BackupHistoryRecord) GetExpiresOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresOn, true
}

// SetExpiresOn sets field value
func (o *BackupHistoryRecord) SetExpiresOn(v time.Time) {
	o.ExpiresOn = v
}

func (o BackupHistoryRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackupHistoryRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["fileName"] = o.FileName.Get()
	toSerialize["storageType"] = o.StorageType
	toSerialize["createdOn"] = o.CreatedOn
	toSerialize["expiresOn"] = o.ExpiresOn
	return toSerialize, nil
}

func (o *BackupHistoryRecord) UnmarshalJSON(data []byte) (err error) {
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

	varBackupHistoryRecord := _BackupHistoryRecord{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varBackupHistoryRecord)

	if err != nil {
		return err
	}

	*o = BackupHistoryRecord(varBackupHistoryRecord)

	return err
}

type NullableBackupHistoryRecord struct {
	value *BackupHistoryRecord
	isSet bool
}

func (v NullableBackupHistoryRecord) Get() *BackupHistoryRecord {
	return v.value
}

func (v *NullableBackupHistoryRecord) Set(val *BackupHistoryRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupHistoryRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupHistoryRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupHistoryRecord(val *BackupHistoryRecord) *NullableBackupHistoryRecord {
	return &NullableBackupHistoryRecord{value: val, isSet: true}
}

func (v NullableBackupHistoryRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupHistoryRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

