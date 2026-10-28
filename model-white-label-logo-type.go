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

// WhiteLabelLogoType [1 - Light small, 2 - Login page, 3 - Favicon, 4 - Docs editor, 5 - Docs editor embed, 6 - Left menu, 7 - About page, 8 - Notification, 9 - Spreadsheet editor, 10 - Spreadsheet editor embed, 11 - Presentation editor, 12 - Presentation editor embed, 13 - Pdf editor, 14 - Pdf editor embed, 15 - Diagram editor, 16 - Diagram editor embed]
type WhiteLabelLogoType int32

// List of WhiteLabelLogoType
const (
	WHITELABELLOGOTYPE_LightSmall WhiteLabelLogoType = 1
	WHITELABELLOGOTYPE_LoginPage WhiteLabelLogoType = 2
	WHITELABELLOGOTYPE_Favicon WhiteLabelLogoType = 3
	WHITELABELLOGOTYPE_DocsEditor WhiteLabelLogoType = 4
	WHITELABELLOGOTYPE_DocsEditorEmbed WhiteLabelLogoType = 5
	WHITELABELLOGOTYPE_LeftMenu WhiteLabelLogoType = 6
	WHITELABELLOGOTYPE_AboutPage WhiteLabelLogoType = 7
	WHITELABELLOGOTYPE_Notification WhiteLabelLogoType = 8
	WHITELABELLOGOTYPE_SpreadsheetEditor WhiteLabelLogoType = 9
	WHITELABELLOGOTYPE_SpreadsheetEditorEmbed WhiteLabelLogoType = 10
	WHITELABELLOGOTYPE_PresentationEditor WhiteLabelLogoType = 11
	WHITELABELLOGOTYPE_PresentationEditorEmbed WhiteLabelLogoType = 12
	WHITELABELLOGOTYPE_PdfEditor WhiteLabelLogoType = 13
	WHITELABELLOGOTYPE_PdfEditorEmbed WhiteLabelLogoType = 14
	WHITELABELLOGOTYPE_DiagramEditor WhiteLabelLogoType = 15
	WHITELABELLOGOTYPE_DiagramEditorEmbed WhiteLabelLogoType = 16
)

// All allowed values of WhiteLabelLogoType enum
var AllowedWhiteLabelLogoTypeEnumValues = []WhiteLabelLogoType{
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
	11,
	12,
	13,
	14,
	15,
	16,
}

func (v *WhiteLabelLogoType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WhiteLabelLogoType(value)
	for _, existing := range AllowedWhiteLabelLogoTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WhiteLabelLogoType", value)
}

// NewWhiteLabelLogoTypeFromValue returns a pointer to a valid WhiteLabelLogoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWhiteLabelLogoTypeFromValue(v int32) (*WhiteLabelLogoType, error) {
	ev := WhiteLabelLogoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WhiteLabelLogoType: valid values are %v", v, AllowedWhiteLabelLogoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WhiteLabelLogoType) IsValid() bool {
	for _, existing := range AllowedWhiteLabelLogoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WhiteLabelLogoType value
func (v WhiteLabelLogoType) Ptr() *WhiteLabelLogoType {
	return &v
}

type NullableWhiteLabelLogoType struct {
	value *WhiteLabelLogoType
	isSet bool
}

func (v NullableWhiteLabelLogoType) Get() *WhiteLabelLogoType {
	return v.value
}

func (v *NullableWhiteLabelLogoType) Set(val *WhiteLabelLogoType) {
	v.value = val
	v.isSet = true
}

func (v NullableWhiteLabelLogoType) IsSet() bool {
	return v.isSet
}

func (v *NullableWhiteLabelLogoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhiteLabelLogoType(val *WhiteLabelLogoType) *NullableWhiteLabelLogoType {
	return &NullableWhiteLabelLogoType{value: val, isSet: true}
}

func (v NullableWhiteLabelLogoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhiteLabelLogoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

