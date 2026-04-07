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

// BackupPeriod [0 - Every day, 1 - Every week, 2 - Every month]
type BackupPeriod int32

// List of BackupPeriod
const (
	BACKUPPERIOD_EveryDay BackupPeriod = 0
	BACKUPPERIOD_EveryWeek BackupPeriod = 1
	BACKUPPERIOD_EveryMonth BackupPeriod = 2
)

// All allowed values of BackupPeriod enum
var AllowedBackupPeriodEnumValues = []BackupPeriod{
	0,
	1,
	2,
}

func (v *BackupPeriod) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BackupPeriod(value)
	for _, existing := range AllowedBackupPeriodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BackupPeriod", value)
}

// NewBackupPeriodFromValue returns a pointer to a valid BackupPeriod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBackupPeriodFromValue(v int32) (*BackupPeriod, error) {
	ev := BackupPeriod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BackupPeriod: valid values are %v", v, AllowedBackupPeriodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BackupPeriod) IsValid() bool {
	for _, existing := range AllowedBackupPeriodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackupPeriod value
func (v BackupPeriod) Ptr() *BackupPeriod {
	return &v
}

type NullableBackupPeriod struct {
	value *BackupPeriod
	isSet bool
}

func (v NullableBackupPeriod) Get() *BackupPeriod {
	return v.value
}

func (v *NullableBackupPeriod) Set(val *BackupPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupPeriod(val *BackupPeriod) *NullableBackupPeriod {
	return &NullableBackupPeriod{value: val, isSet: true}
}

func (v NullableBackupPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

