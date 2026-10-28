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

// FilterType [0 - None, 1 - Files  only, 2 - Folders only, 3 - Documents only, 4 - Presentations only, 5 - Spreadsheets only, 7 - Images only, 8 - By user, 9 - By department, 10 - Archive only, 11 - By extension, 12 - Media only, 13 - Filling forms rooms, 14 - Editing rooms, 17 - Custom rooms, 20 - Public rooms, 22 - Pdf, 23 - Pdf form, 24 - Virtual data rooms, 25 - Diagrams only, 26 - Ai rooms]
type FilterType int32

// List of FilterType
const (
	FILTERTYPE_None FilterType = 0
	FILTERTYPE_FilesOnly FilterType = 1
	FILTERTYPE_FoldersOnly FilterType = 2
	FILTERTYPE_DocumentsOnly FilterType = 3
	FILTERTYPE_PresentationsOnly FilterType = 4
	FILTERTYPE_SpreadsheetsOnly FilterType = 5
	FILTERTYPE_ImagesOnly FilterType = 7
	FILTERTYPE_ByUser FilterType = 8
	FILTERTYPE_ByDepartment FilterType = 9
	FILTERTYPE_ArchiveOnly FilterType = 10
	FILTERTYPE_ByExtension FilterType = 11
	FILTERTYPE_MediaOnly FilterType = 12
	FILTERTYPE_FillingFormsRooms FilterType = 13
	FILTERTYPE_EditingRooms FilterType = 14
	FILTERTYPE_CustomRooms FilterType = 17
	FILTERTYPE_PublicRooms FilterType = 20
	FILTERTYPE_Pdf FilterType = 22
	FILTERTYPE_PdfForm FilterType = 23
	FILTERTYPE_VirtualDataRooms FilterType = 24
	FILTERTYPE_DiagramsOnly FilterType = 25
	FILTERTYPE_AiRooms FilterType = 26
)

// All allowed values of FilterType enum
var AllowedFilterTypeEnumValues = []FilterType{
	0,
	1,
	2,
	3,
	4,
	5,
	7,
	8,
	9,
	10,
	11,
	12,
	13,
	14,
	17,
	20,
	22,
	23,
	24,
	25,
	26,
}

func (v *FilterType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FilterType(value)
	for _, existing := range AllowedFilterTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FilterType", value)
}

// NewFilterTypeFromValue returns a pointer to a valid FilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFilterTypeFromValue(v int32) (*FilterType, error) {
	ev := FilterType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FilterType: valid values are %v", v, AllowedFilterTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FilterType) IsValid() bool {
	for _, existing := range AllowedFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FilterType value
func (v FilterType) Ptr() *FilterType {
	return &v
}

type NullableFilterType struct {
	value *FilterType
	isSet bool
}

func (v NullableFilterType) Get() *FilterType {
	return v.value
}

func (v *NullableFilterType) Set(val *FilterType) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterType) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterType(val *FilterType) *NullableFilterType {
	return &NullableFilterType{value: val, isSet: true}
}

func (v NullableFilterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

