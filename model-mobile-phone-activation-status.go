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

// MobilePhoneActivationStatus [0 - Not activated, 1 - Activated]
type MobilePhoneActivationStatus int32

// List of MobilePhoneActivationStatus
const (
	MOBILEPHONEACTIVATIONSTATUS_NotActivated MobilePhoneActivationStatus = 0
	MOBILEPHONEACTIVATIONSTATUS_Activated MobilePhoneActivationStatus = 1
)

// All allowed values of MobilePhoneActivationStatus enum
var AllowedMobilePhoneActivationStatusEnumValues = []MobilePhoneActivationStatus{
	0,
	1,
}

func (v *MobilePhoneActivationStatus) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MobilePhoneActivationStatus(value)
	for _, existing := range AllowedMobilePhoneActivationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MobilePhoneActivationStatus", value)
}

// NewMobilePhoneActivationStatusFromValue returns a pointer to a valid MobilePhoneActivationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMobilePhoneActivationStatusFromValue(v int32) (*MobilePhoneActivationStatus, error) {
	ev := MobilePhoneActivationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MobilePhoneActivationStatus: valid values are %v", v, AllowedMobilePhoneActivationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MobilePhoneActivationStatus) IsValid() bool {
	for _, existing := range AllowedMobilePhoneActivationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MobilePhoneActivationStatus value
func (v MobilePhoneActivationStatus) Ptr() *MobilePhoneActivationStatus {
	return &v
}

type NullableMobilePhoneActivationStatus struct {
	value *MobilePhoneActivationStatus
	isSet bool
}

func (v NullableMobilePhoneActivationStatus) Get() *MobilePhoneActivationStatus {
	return v.value
}

func (v *NullableMobilePhoneActivationStatus) Set(val *MobilePhoneActivationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMobilePhoneActivationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMobilePhoneActivationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobilePhoneActivationStatus(val *MobilePhoneActivationStatus) *NullableMobilePhoneActivationStatus {
	return &NullableMobilePhoneActivationStatus{value: val, isSet: true}
}

func (v NullableMobilePhoneActivationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobilePhoneActivationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

