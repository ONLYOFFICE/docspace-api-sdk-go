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
)

// checks if the NotificationSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationSettingsDto{}

// NotificationSettingsDto Whether one kind of notification is switched on for the calling user.
type NotificationSettingsDto struct {
	// Which kind of notification the flag belongs to, echoed from the request. It is published as a number:  badges, room activity, the daily feed, and the tips.
	Type *NotificationType `json:"type,omitempty"`
	// Whether the caller receives that kind of notification. It describes the caller's own account and nobody  else's; a fresh account has the badges on and the other three off, because those are subscriptions that  only `POST api/2.0/settings/notification` creates.
	IsEnabled *bool `json:"isEnabled,omitempty"`
}

// NewNotificationSettingsDto instantiates a new NotificationSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSettingsDto() *NotificationSettingsDto {
	this := NotificationSettingsDto{}
	return &this
}

// NewNotificationSettingsDtoWithDefaults instantiates a new NotificationSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSettingsDtoWithDefaults() *NotificationSettingsDto {
	this := NotificationSettingsDto{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NotificationSettingsDto) GetType() NotificationType {
	if o == nil || IsNil(o.Type) {
		var ret NotificationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSettingsDto) GetTypeOk() (*NotificationType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NotificationSettingsDto) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given NotificationType and assigns it to the Type field.
func (o *NotificationSettingsDto) SetType(v NotificationType) {
	o.Type = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *NotificationSettingsDto) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSettingsDto) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *NotificationSettingsDto) IsIsEnabledSet() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *NotificationSettingsDto) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

func (o NotificationSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	return toSerialize, nil
}

type NullableNotificationSettingsDto struct {
	value *NotificationSettingsDto
	isSet bool
}

func (v NullableNotificationSettingsDto) Get() *NotificationSettingsDto {
	return v.value
}

func (v *NullableNotificationSettingsDto) Set(val *NotificationSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSettingsDto(val *NotificationSettingsDto) *NullableNotificationSettingsDto {
	return &NullableNotificationSettingsDto{value: val, isSet: true}
}

func (v NullableNotificationSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

