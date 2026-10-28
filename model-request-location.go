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

// RequestLocation [1 - Room, 2 - Documents, 3 - Link]
type RequestLocation int32

// List of RequestLocation
const (
	REQUESTLOCATION_Room RequestLocation = 1
	REQUESTLOCATION_Documents RequestLocation = 2
	REQUESTLOCATION_Link RequestLocation = 3
)

// All allowed values of RequestLocation enum
var AllowedRequestLocationEnumValues = []RequestLocation{
	1,
	2,
	3,
}

func (v *RequestLocation) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestLocation(value)
	for _, existing := range AllowedRequestLocationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestLocation", value)
}

// NewRequestLocationFromValue returns a pointer to a valid RequestLocation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestLocationFromValue(v int32) (*RequestLocation, error) {
	ev := RequestLocation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestLocation: valid values are %v", v, AllowedRequestLocationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestLocation) IsValid() bool {
	for _, existing := range AllowedRequestLocationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RequestLocation value
func (v RequestLocation) Ptr() *RequestLocation {
	return &v
}

type NullableRequestLocation struct {
	value *RequestLocation
	isSet bool
}

func (v NullableRequestLocation) Get() *RequestLocation {
	return v.value
}

func (v *NullableRequestLocation) Set(val *RequestLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestLocation(val *RequestLocation) *NullableRequestLocation {
	return &NullableRequestLocation{value: val, isSet: true}
}

func (v NullableRequestLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

