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

// NotificationType [0 - Badges, 1 - Rooms activity, 2 - Daily feed, 3 - Usefull tips]
type NotificationType int32

// List of NotificationType
const (
	NOTIFICATIONTYPE_Badges NotificationType = 0
	NOTIFICATIONTYPE_RoomsActivity NotificationType = 1
	NOTIFICATIONTYPE_DailyFeed NotificationType = 2
	NOTIFICATIONTYPE_UsefullTips NotificationType = 3
)

// All allowed values of NotificationType enum
var AllowedNotificationTypeEnumValues = []NotificationType{
	0,
	1,
	2,
	3,
}

func (v *NotificationType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotificationType(value)
	for _, existing := range AllowedNotificationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotificationType", value)
}

// NewNotificationTypeFromValue returns a pointer to a valid NotificationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotificationTypeFromValue(v int32) (*NotificationType, error) {
	ev := NotificationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotificationType: valid values are %v", v, AllowedNotificationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotificationType) IsValid() bool {
	for _, existing := range AllowedNotificationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationType value
func (v NotificationType) Ptr() *NotificationType {
	return &v
}

type NullableNotificationType struct {
	value *NotificationType
	isSet bool
}

func (v NullableNotificationType) Get() *NotificationType {
	return v.value
}

func (v *NullableNotificationType) Set(val *NotificationType) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationType) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationType(val *NotificationType) *NullableNotificationType {
	return &NullableNotificationType{value: val, isSet: true}
}

func (v NullableNotificationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

