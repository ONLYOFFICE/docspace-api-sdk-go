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

// TariffState [0 - Trial, 1 - Paid, 2 - Delay, 3 - Not paid]
type TariffState int32

// List of TariffState
const (
	TARIFFSTATE_Trial TariffState = 0
	TARIFFSTATE_Paid TariffState = 1
	TARIFFSTATE_Delay TariffState = 2
	TARIFFSTATE_NotPaid TariffState = 3
)

// All allowed values of TariffState enum
var AllowedTariffStateEnumValues = []TariffState{
	0,
	1,
	2,
	3,
}

func (v *TariffState) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TariffState(value)
	for _, existing := range AllowedTariffStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TariffState", value)
}

// NewTariffStateFromValue returns a pointer to a valid TariffState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTariffStateFromValue(v int32) (*TariffState, error) {
	ev := TariffState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TariffState: valid values are %v", v, AllowedTariffStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TariffState) IsValid() bool {
	for _, existing := range AllowedTariffStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TariffState value
func (v TariffState) Ptr() *TariffState {
	return &v
}

type NullableTariffState struct {
	value *TariffState
	isSet bool
}

func (v NullableTariffState) Get() *TariffState {
	return v.value
}

func (v *NullableTariffState) Set(val *TariffState) {
	v.value = val
	v.isSet = true
}

func (v NullableTariffState) IsSet() bool {
	return v.isSet
}

func (v *NullableTariffState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTariffState(val *TariffState) *NullableTariffState {
	return &NullableTariffState{value: val, isSet: true}
}

func (v NullableTariffState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTariffState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

