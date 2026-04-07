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

// EncryprtionStatus [0 - Decrypted, 1 - Encryption started, 2 - Encrypted, 3 - Decryption started]
type EncryprtionStatus int32

// List of EncryprtionStatus
const (
	ENCRYPRTIONSTATUS_Decrypted EncryprtionStatus = 0
	ENCRYPRTIONSTATUS_EncryptionStarted EncryprtionStatus = 1
	ENCRYPRTIONSTATUS_Encrypted EncryprtionStatus = 2
	ENCRYPRTIONSTATUS_DecryptionStarted EncryprtionStatus = 3
)

// All allowed values of EncryprtionStatus enum
var AllowedEncryprtionStatusEnumValues = []EncryprtionStatus{
	0,
	1,
	2,
	3,
}

func (v *EncryprtionStatus) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EncryprtionStatus(value)
	for _, existing := range AllowedEncryprtionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EncryprtionStatus", value)
}

// NewEncryprtionStatusFromValue returns a pointer to a valid EncryprtionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEncryprtionStatusFromValue(v int32) (*EncryprtionStatus, error) {
	ev := EncryprtionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EncryprtionStatus: valid values are %v", v, AllowedEncryprtionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EncryprtionStatus) IsValid() bool {
	for _, existing := range AllowedEncryprtionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EncryprtionStatus value
func (v EncryprtionStatus) Ptr() *EncryprtionStatus {
	return &v
}

type NullableEncryprtionStatus struct {
	value *EncryprtionStatus
	isSet bool
}

func (v NullableEncryprtionStatus) Get() *EncryprtionStatus {
	return v.value
}

func (v *NullableEncryprtionStatus) Set(val *EncryprtionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryprtionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryprtionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryprtionStatus(val *EncryprtionStatus) *NullableEncryprtionStatus {
	return &NullableEncryprtionStatus{value: val, isSet: true}
}

func (v NullableEncryprtionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryprtionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

