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
	"fmt"
)

// BackupStorageType [0 - Documents, 1 - Thridparty documents, 2 - Custom cloud, 3 - Local, 4 - Data store, 5 - Thirdparty consumer]
type BackupStorageType int32

// List of BackupStorageType
const (
	BACKUPSTORAGETYPE_Documents BackupStorageType = 0
	BACKUPSTORAGETYPE_ThridpartyDocuments BackupStorageType = 1
	BACKUPSTORAGETYPE_CustomCloud BackupStorageType = 2
	BACKUPSTORAGETYPE_Local BackupStorageType = 3
	BACKUPSTORAGETYPE_DataStore BackupStorageType = 4
	BACKUPSTORAGETYPE_ThirdPartyConsumer BackupStorageType = 5
)

// All allowed values of BackupStorageType enum
var AllowedBackupStorageTypeEnumValues = []BackupStorageType{
	0,
	1,
	2,
	3,
	4,
	5,
}

func (v *BackupStorageType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BackupStorageType(value)
	for _, existing := range AllowedBackupStorageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BackupStorageType", value)
}

// NewBackupStorageTypeFromValue returns a pointer to a valid BackupStorageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBackupStorageTypeFromValue(v int32) (*BackupStorageType, error) {
	ev := BackupStorageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BackupStorageType: valid values are %v", v, AllowedBackupStorageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BackupStorageType) IsValid() bool {
	for _, existing := range AllowedBackupStorageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackupStorageType value
func (v BackupStorageType) Ptr() *BackupStorageType {
	return &v
}

type NullableBackupStorageType struct {
	value *BackupStorageType
	isSet bool
}

func (v NullableBackupStorageType) Get() *BackupStorageType {
	return v.value
}

func (v *NullableBackupStorageType) Set(val *BackupStorageType) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupStorageType) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupStorageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupStorageType(val *BackupStorageType) *NullableBackupStorageType {
	return &NullableBackupStorageType{value: val, isSet: true}
}

func (v NullableBackupStorageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupStorageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

