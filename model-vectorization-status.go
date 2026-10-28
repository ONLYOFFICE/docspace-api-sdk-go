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

// VectorizationStatus [0 - In Progress, 1 - Completed, 2 - Failed]
type VectorizationStatus int32

// List of VectorizationStatus
const (
	VECTORIZATIONSTATUS_InProgress VectorizationStatus = 0
	VECTORIZATIONSTATUS_Completed VectorizationStatus = 1
	VECTORIZATIONSTATUS_Failed VectorizationStatus = 2
)

// All allowed values of VectorizationStatus enum
var AllowedVectorizationStatusEnumValues = []VectorizationStatus{
	0,
	1,
	2,
}

func (v *VectorizationStatus) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VectorizationStatus(value)
	for _, existing := range AllowedVectorizationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VectorizationStatus", value)
}

// NewVectorizationStatusFromValue returns a pointer to a valid VectorizationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVectorizationStatusFromValue(v int32) (*VectorizationStatus, error) {
	ev := VectorizationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VectorizationStatus: valid values are %v", v, AllowedVectorizationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VectorizationStatus) IsValid() bool {
	for _, existing := range AllowedVectorizationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VectorizationStatus value
func (v VectorizationStatus) Ptr() *VectorizationStatus {
	return &v
}

type NullableVectorizationStatus struct {
	value *VectorizationStatus
	isSet bool
}

func (v NullableVectorizationStatus) Get() *VectorizationStatus {
	return v.value
}

func (v *NullableVectorizationStatus) Set(val *VectorizationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVectorizationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVectorizationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVectorizationStatus(val *VectorizationStatus) *NullableVectorizationStatus {
	return &NullableVectorizationStatus{value: val, isSet: true}
}

func (v NullableVectorizationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVectorizationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

