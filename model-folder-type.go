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

// FolderType [0 - Default, 1 - Coomon, 2 - Bunch, 3 - Trash, 5 - User, 6 - Share, 8 - Projects, 10 - Favourites, 11 - Recent, 12 - Templates, 13 - Privacy, 14 - Virtual rooms, 15 - Filling forms room, 16 - Editing room, 19 - Custom room, 20 - Archive, 21 - Thirdparty backup, 22 - Public room, 25 - Ready form folder, 26 - In process form folder, 27 - Form filling folder done, 28 - Form filling folder in progress, 29 - Virtual Data Room, 30 - Room templates folder, 31 - AI Room, 32 - Knowledge, 33 - Chat outputs, 34 - AI Agents, 35 - Default Templates, 36 - Forms]
type FolderType int32

// List of FolderType
const (
	FOLDERTYPE_DEFAULT FolderType = 0
	FOLDERTYPE_COMMON FolderType = 1
	FOLDERTYPE_BUNCH FolderType = 2
	FOLDERTYPE_TRASH FolderType = 3
	FOLDERTYPE_USER FolderType = 5
	FOLDERTYPE_SHARE FolderType = 6
	FOLDERTYPE_Projects FolderType = 8
	FOLDERTYPE_Favorites FolderType = 10
	FOLDERTYPE_Recent FolderType = 11
	FOLDERTYPE_Templates FolderType = 12
	FOLDERTYPE_Privacy FolderType = 13
	FOLDERTYPE_VirtualRooms FolderType = 14
	FOLDERTYPE_FillingFormsRoom FolderType = 15
	FOLDERTYPE_EditingRoom FolderType = 16
	FOLDERTYPE_CustomRoom FolderType = 19
	FOLDERTYPE_Archive FolderType = 20
	FOLDERTYPE_ThirdpartyBackup FolderType = 21
	FOLDERTYPE_PublicRoom FolderType = 22
	FOLDERTYPE_ReadyFormFolder FolderType = 25
	FOLDERTYPE_InProcessFormFolder FolderType = 26
	FOLDERTYPE_FormFillingFolderDone FolderType = 27
	FOLDERTYPE_FormFillingFolderInProgress FolderType = 28
	FOLDERTYPE_VirtualDataRoom FolderType = 29
	FOLDERTYPE_RoomTemplates FolderType = 30
	FOLDERTYPE_AiRoom FolderType = 31
	FOLDERTYPE_Knowledge FolderType = 32
	FOLDERTYPE_ChatOutputs FolderType = 33
	FOLDERTYPE_AiAgents FolderType = 34
	FOLDERTYPE_DefaultTemplates FolderType = 35
	FOLDERTYPE_Forms FolderType = 36
)

// All allowed values of FolderType enum
var AllowedFolderTypeEnumValues = []FolderType{
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

func (v *FolderType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FolderType(value)
	for _, existing := range AllowedFolderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FolderType", value)
}

// NewFolderTypeFromValue returns a pointer to a valid FolderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFolderTypeFromValue(v int32) (*FolderType, error) {
	ev := FolderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FolderType: valid values are %v", v, AllowedFolderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FolderType) IsValid() bool {
	for _, existing := range AllowedFolderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FolderType value
func (v FolderType) Ptr() *FolderType {
	return &v
}

type NullableFolderType struct {
	value *FolderType
	isSet bool
}

func (v NullableFolderType) Get() *FolderType {
	return v.value
}

func (v *NullableFolderType) Set(val *FolderType) {
	v.value = val
	v.isSet = true
}

func (v NullableFolderType) IsSet() bool {
	return v.isSet
}

func (v *NullableFolderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolderType(val *FolderType) *NullableFolderType {
	return &NullableFolderType{value: val, isSet: true}
}

func (v NullableFolderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

