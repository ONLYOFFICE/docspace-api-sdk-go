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

// WebhookTrigger [0 - *, 1 - user.created, 2 - user.invited, 4 - user.updated, 8 - user.deleted, 16 - group.created, 32 - group.updated, 64 - group.deleted, 128 - file.created, 256 - file.uploaded, 512 - file.updated, 1024 - file.trashed, 2048 - file.deleted, 4096 - file.restored, 8192 - file.copied, 16384 - file.moved, 32768 - folder.created, 65536 - folder.updated, 131072 - folder.trashed, 262144 - folder.deleted, 524288 - folder.restored, 1048576 - folder.copied, 2097152 - folder.moved, 4194304 - room.created, 8388608 - room.updated, 16777216 - room.archived, 33554432 - room.deleted, 67108864 - room.restored, 134217728 - room.copied, 268435456 - form.submit, 536870912 - form.filled.out, 1073741824 - form.stopped]
type WebhookTrigger int32

// List of WebhookTrigger
const (
	WEBHOOKTRIGGER_All WebhookTrigger = 0
	WEBHOOKTRIGGER_UserCreated WebhookTrigger = 1
	WEBHOOKTRIGGER_UserInvited WebhookTrigger = 2
	WEBHOOKTRIGGER_UserUpdated WebhookTrigger = 4
	WEBHOOKTRIGGER_UserDeleted WebhookTrigger = 8
	WEBHOOKTRIGGER_GroupCreated WebhookTrigger = 16
	WEBHOOKTRIGGER_GroupUpdated WebhookTrigger = 32
	WEBHOOKTRIGGER_GroupDeleted WebhookTrigger = 64
	WEBHOOKTRIGGER_FileCreated WebhookTrigger = 128
	WEBHOOKTRIGGER_FileUploaded WebhookTrigger = 256
	WEBHOOKTRIGGER_FileUpdated WebhookTrigger = 512
	WEBHOOKTRIGGER_FileTrashed WebhookTrigger = 1024
	WEBHOOKTRIGGER_FileDeleted WebhookTrigger = 2048
	WEBHOOKTRIGGER_FileRestored WebhookTrigger = 4096
	WEBHOOKTRIGGER_FileCopied WebhookTrigger = 8192
	WEBHOOKTRIGGER_FileMoved WebhookTrigger = 16384
	WEBHOOKTRIGGER_FolderCreated WebhookTrigger = 32768
	WEBHOOKTRIGGER_FolderUpdated WebhookTrigger = 65536
	WEBHOOKTRIGGER_FolderTrashed WebhookTrigger = 131072
	WEBHOOKTRIGGER_FolderDeleted WebhookTrigger = 262144
	WEBHOOKTRIGGER_FolderRestored WebhookTrigger = 524288
	WEBHOOKTRIGGER_FolderCopied WebhookTrigger = 1048576
	WEBHOOKTRIGGER_FolderMoved WebhookTrigger = 2097152
	WEBHOOKTRIGGER_RoomCreated WebhookTrigger = 4194304
	WEBHOOKTRIGGER_RoomUpdated WebhookTrigger = 8388608
	WEBHOOKTRIGGER_RoomArchived WebhookTrigger = 16777216
	WEBHOOKTRIGGER_RoomDeleted WebhookTrigger = 33554432
	WEBHOOKTRIGGER_RoomRestored WebhookTrigger = 67108864
	WEBHOOKTRIGGER_RoomCopied WebhookTrigger = 134217728
	WEBHOOKTRIGGER_FormSubmit WebhookTrigger = 268435456
	WEBHOOKTRIGGER_FormFilledOut WebhookTrigger = 536870912
	WEBHOOKTRIGGER_FormStopped WebhookTrigger = 1073741824
)

// All allowed values of WebhookTrigger enum
var AllowedWebhookTriggerEnumValues = []WebhookTrigger{
	0,
	1,
	2,
	4,
	8,
	16,
	32,
	64,
	128,
	256,
	512,
	1024,
	2048,
	4096,
	8192,
	16384,
	32768,
	65536,
	131072,
	262144,
	524288,
	1048576,
	2097152,
	4194304,
	8388608,
	16777216,
	33554432,
	67108864,
	134217728,
	268435456,
	536870912,
	1073741824,
}

func (v *WebhookTrigger) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebhookTrigger(value)
	for _, existing := range AllowedWebhookTriggerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebhookTrigger", value)
}

// NewWebhookTriggerFromValue returns a pointer to a valid WebhookTrigger
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhookTriggerFromValue(v int32) (*WebhookTrigger, error) {
	ev := WebhookTrigger(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebhookTrigger: valid values are %v", v, AllowedWebhookTriggerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhookTrigger) IsValid() bool {
	for _, existing := range AllowedWebhookTriggerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebhookTrigger value
func (v WebhookTrigger) Ptr() *WebhookTrigger {
	return &v
}

type NullableWebhookTrigger struct {
	value *WebhookTrigger
	isSet bool
}

func (v NullableWebhookTrigger) Get() *WebhookTrigger {
	return v.value
}

func (v *NullableWebhookTrigger) Set(val *WebhookTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookTrigger(val *WebhookTrigger) *NullableWebhookTrigger {
	return &NullableWebhookTrigger{value: val, isSet: true}
}

func (v NullableWebhookTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

