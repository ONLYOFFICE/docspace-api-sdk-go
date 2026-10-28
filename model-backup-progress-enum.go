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

// BackupProgressEnum [0 - Backup, 1 - Restore, 2 - Transfer]
type BackupProgressEnum int32

// List of BackupProgressEnum
const (
	BACKUPPROGRESSENUM_Backup BackupProgressEnum = 0
	BACKUPPROGRESSENUM_Restore BackupProgressEnum = 1
	BACKUPPROGRESSENUM_Transfer BackupProgressEnum = 2
)

// All allowed values of BackupProgressEnum enum
var AllowedBackupProgressEnumEnumValues = []BackupProgressEnum{
	0,
	1,
	2,
}

func (v *BackupProgressEnum) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BackupProgressEnum(value)
	for _, existing := range AllowedBackupProgressEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BackupProgressEnum", value)
}

// NewBackupProgressEnumFromValue returns a pointer to a valid BackupProgressEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBackupProgressEnumFromValue(v int32) (*BackupProgressEnum, error) {
	ev := BackupProgressEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BackupProgressEnum: valid values are %v", v, AllowedBackupProgressEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BackupProgressEnum) IsValid() bool {
	for _, existing := range AllowedBackupProgressEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackupProgressEnum value
func (v BackupProgressEnum) Ptr() *BackupProgressEnum {
	return &v
}

type NullableBackupProgressEnum struct {
	value *BackupProgressEnum
	isSet bool
}

func (v NullableBackupProgressEnum) Get() *BackupProgressEnum {
	return v.value
}

func (v *NullableBackupProgressEnum) Set(val *BackupProgressEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupProgressEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupProgressEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupProgressEnum(val *BackupProgressEnum) *NullableBackupProgressEnum {
	return &NullableBackupProgressEnum{value: val, isSet: true}
}

func (v NullableBackupProgressEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupProgressEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

