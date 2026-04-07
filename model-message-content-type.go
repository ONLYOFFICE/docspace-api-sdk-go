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

// MessageContentType [0 - Text, 1 - Tool, 2 - Attachment]
type MessageContentType int32

// List of MessageContentType
const (
	MESSAGECONTENTTYPE_Text MessageContentType = 0
	MESSAGECONTENTTYPE_Tool MessageContentType = 1
	MESSAGECONTENTTYPE_Attachment MessageContentType = 2
	MESSAGECONTENTTYPE_Data MessageContentType = 3
)

// All allowed values of MessageContentType enum
var AllowedMessageContentTypeEnumValues = []MessageContentType{
	0,
	1,
	2,
	3,
}

func (v *MessageContentType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MessageContentType(value)
	for _, existing := range AllowedMessageContentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MessageContentType", value)
}

// NewMessageContentTypeFromValue returns a pointer to a valid MessageContentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMessageContentTypeFromValue(v int32) (*MessageContentType, error) {
	ev := MessageContentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MessageContentType: valid values are %v", v, AllowedMessageContentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MessageContentType) IsValid() bool {
	for _, existing := range AllowedMessageContentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MessageContentType value
func (v MessageContentType) Ptr() *MessageContentType {
	return &v
}

type NullableMessageContentType struct {
	value *MessageContentType
	isSet bool
}

func (v NullableMessageContentType) Get() *MessageContentType {
	return v.value
}

func (v *NullableMessageContentType) Set(val *MessageContentType) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageContentType) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageContentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageContentType(val *MessageContentType) *NullableMessageContentType {
	return &NullableMessageContentType{value: val, isSet: true}
}

func (v NullableMessageContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageContentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

