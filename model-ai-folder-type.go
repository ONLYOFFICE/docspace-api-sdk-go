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

// AiFolderType [0 - Default, 1 - Coomon, 2 - Bunch, 3 - Trash, 5 - User, 6 - Share, 8 - Projects, 10 - Favourites, 11 - Recent, 12 - Templates, 13 - Privacy, 14 - Virtual rooms, 15 - Filling forms room, 16 - Editing room, 19 - Custom room, 20 - Archive, 21 - Thirdparty backup, 22 - Public room, 25 - Ready form folder, 26 - In process form folder, 27 - Form filling folder done, 28 - Form filling folder in progress, 29 - Virtual Data Room, 30 - Room templates folder, 31 - AI Room, 32 - Knowledge, 33 - Result storage, 34 - AI Agents, 35 - Default Templates, 36 - Forms]
type AiFolderType int32

// List of AiFolderType
const (
	AIFOLDERTYPE_DEFAULT AiFolderType = 0
	AIFOLDERTYPE_COMMON AiFolderType = 1
	AIFOLDERTYPE_BUNCH AiFolderType = 2
	AIFOLDERTYPE_TRASH AiFolderType = 3
	AIFOLDERTYPE_USER AiFolderType = 5
	AIFOLDERTYPE_SHARE AiFolderType = 6
	AIFOLDERTYPE_Projects AiFolderType = 8
	AIFOLDERTYPE_Favorites AiFolderType = 10
	AIFOLDERTYPE_Recent AiFolderType = 11
	AIFOLDERTYPE_Templates AiFolderType = 12
	AIFOLDERTYPE_Privacy AiFolderType = 13
	AIFOLDERTYPE_VirtualRooms AiFolderType = 14
	AIFOLDERTYPE_FillingFormsRoom AiFolderType = 15
	AIFOLDERTYPE_EditingRoom AiFolderType = 16
	AIFOLDERTYPE_CustomRoom AiFolderType = 19
	AIFOLDERTYPE_Archive AiFolderType = 20
	AIFOLDERTYPE_ThirdpartyBackup AiFolderType = 21
	AIFOLDERTYPE_PublicRoom AiFolderType = 22
	AIFOLDERTYPE_ReadyFormFolder AiFolderType = 25
	AIFOLDERTYPE_InProcessFormFolder AiFolderType = 26
	AIFOLDERTYPE_FormFillingFolderDone AiFolderType = 27
	AIFOLDERTYPE_FormFillingFolderInProgress AiFolderType = 28
	AIFOLDERTYPE_VirtualDataRoom AiFolderType = 29
	AIFOLDERTYPE_RoomTemplates AiFolderType = 30
	AIFOLDERTYPE_AiRoom AiFolderType = 31
	AIFOLDERTYPE_Knowledge AiFolderType = 32
	AIFOLDERTYPE_ResultStorage AiFolderType = 33
	AIFOLDERTYPE_AiAgents AiFolderType = 34
	AIFOLDERTYPE_DefaultTemplates AiFolderType = 35
	AIFOLDERTYPE_Forms AiFolderType = 36
)

// All allowed values of AiFolderType enum
var AllowedAiFolderTypeEnumValues = []AiFolderType{
	0,
	1,
	2,
	3,
	5,
	6,
	8,
	10,
	11,
	12,
	13,
	14,
	15,
	16,
	19,
	20,
	21,
	22,
	25,
	26,
	27,
	28,
	29,
	30,
	31,
	32,
	33,
	34,
	35,
	36,
}

func (v *AiFolderType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AiFolderType(value)
	for _, existing := range AllowedAiFolderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AiFolderType", value)
}

// NewAiFolderTypeFromValue returns a pointer to a valid AiFolderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAiFolderTypeFromValue(v int32) (*AiFolderType, error) {
	ev := AiFolderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AiFolderType: valid values are %v", v, AllowedAiFolderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AiFolderType) IsValid() bool {
	for _, existing := range AllowedAiFolderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AiFolderType value
func (v AiFolderType) Ptr() *AiFolderType {
	return &v
}

type NullableAiFolderType struct {
	value *AiFolderType
	isSet bool
}

func (v NullableAiFolderType) Get() *AiFolderType {
	return v.value
}

func (v *NullableAiFolderType) Set(val *AiFolderType) {
	v.value = val
	v.isSet = true
}

func (v NullableAiFolderType) IsSet() bool {
	return v.isSet
}

func (v *NullableAiFolderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiFolderType(val *AiFolderType) *NullableAiFolderType {
	return &NullableAiFolderType{value: val, isSet: true}
}

func (v NullableAiFolderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiFolderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

