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

// WebhookGroupStatus [0 - None, 1 - Not sent, 2 - Status2xx, 4 - Status3xx, 8 - Status4xx, 16 - Status5xx]
type WebhookGroupStatus int32

// List of WebhookGroupStatus
const (
	WEBHOOKGROUPSTATUS_None WebhookGroupStatus = 0
	WEBHOOKGROUPSTATUS_NotSent WebhookGroupStatus = 1
	WEBHOOKGROUPSTATUS_Status2xx WebhookGroupStatus = 2
	WEBHOOKGROUPSTATUS_Status3xx WebhookGroupStatus = 4
	WEBHOOKGROUPSTATUS_Status4xx WebhookGroupStatus = 8
	WEBHOOKGROUPSTATUS_Status5xx WebhookGroupStatus = 16
)

// All allowed values of WebhookGroupStatus enum
var AllowedWebhookGroupStatusEnumValues = []WebhookGroupStatus{
	0,
	1,
	2,
	4,
	8,
	16,
}

func (v *WebhookGroupStatus) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebhookGroupStatus(value)
	for _, existing := range AllowedWebhookGroupStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebhookGroupStatus", value)
}

// NewWebhookGroupStatusFromValue returns a pointer to a valid WebhookGroupStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhookGroupStatusFromValue(v int32) (*WebhookGroupStatus, error) {
	ev := WebhookGroupStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebhookGroupStatus: valid values are %v", v, AllowedWebhookGroupStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhookGroupStatus) IsValid() bool {
	for _, existing := range AllowedWebhookGroupStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebhookGroupStatus value
func (v WebhookGroupStatus) Ptr() *WebhookGroupStatus {
	return &v
}

type NullableWebhookGroupStatus struct {
	value *WebhookGroupStatus
	isSet bool
}

func (v NullableWebhookGroupStatus) Get() *WebhookGroupStatus {
	return v.value
}

func (v *NullableWebhookGroupStatus) Set(val *WebhookGroupStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookGroupStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookGroupStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookGroupStatus(val *WebhookGroupStatus) *NullableWebhookGroupStatus {
	return &NullableWebhookGroupStatus{value: val, isSet: true}
}

func (v NullableWebhookGroupStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookGroupStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

