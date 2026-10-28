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

// EncryptionStatus [0 - Decrypted, 1 - Encryption started, 2 - Encrypted, 3 - Decryption started]
type EncryptionStatus int32

// List of EncryptionStatus
const (
	ENCRYPTIONSTATUS_Decrypted EncryptionStatus = 0
	ENCRYPTIONSTATUS_EncryptionStarted EncryptionStatus = 1
	ENCRYPTIONSTATUS_Encrypted EncryptionStatus = 2
	ENCRYPTIONSTATUS_DecryptionStarted EncryptionStatus = 3
)

// All allowed values of EncryptionStatus enum
var AllowedEncryptionStatusEnumValues = []EncryptionStatus{
	0,
	1,
	2,
	3,
}

func (v *EncryptionStatus) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EncryptionStatus(value)
	for _, existing := range AllowedEncryptionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EncryptionStatus", value)
}

// NewEncryptionStatusFromValue returns a pointer to a valid EncryptionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEncryptionStatusFromValue(v int32) (*EncryptionStatus, error) {
	ev := EncryptionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EncryptionStatus: valid values are %v", v, AllowedEncryptionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EncryptionStatus) IsValid() bool {
	for _, existing := range AllowedEncryptionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EncryptionStatus value
func (v EncryptionStatus) Ptr() *EncryptionStatus {
	return &v
}

type NullableEncryptionStatus struct {
	value *EncryptionStatus
	isSet bool
}

func (v NullableEncryptionStatus) Get() *EncryptionStatus {
	return v.value
}

func (v *NullableEncryptionStatus) Set(val *EncryptionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptionStatus(val *EncryptionStatus) *NullableEncryptionStatus {
	return &NullableEncryptionStatus{value: val, isSet: true}
}

func (v NullableEncryptionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

