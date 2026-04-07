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
	"bytes"
	"fmt"
)

// checks if the NotificationSettingsRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationSettingsRequestsDto{}

// NotificationSettingsRequestsDto The request parameters for configuring notification settings.
type NotificationSettingsRequestsDto struct {
	Type NotificationType `json:"type"`
	// Specifies if the specified notification type is enabled or not.
	IsEnabled *bool `json:"isEnabled,omitempty"`
}

type _NotificationSettingsRequestsDto NotificationSettingsRequestsDto

// NewNotificationSettingsRequestsDto instantiates a new NotificationSettingsRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSettingsRequestsDto(type_ NotificationType) *NotificationSettingsRequestsDto {
	this := NotificationSettingsRequestsDto{}
	this.Type = type_
	return &this
}

// NewNotificationSettingsRequestsDtoWithDefaults instantiates a new NotificationSettingsRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSettingsRequestsDtoWithDefaults() *NotificationSettingsRequestsDto {
	this := NotificationSettingsRequestsDto{}
	return &this
}

// GetType returns the Type field value
func (o *NotificationSettingsRequestsDto) GetType() NotificationType {
	if o == nil {
		var ret NotificationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotificationSettingsRequestsDto) GetTypeOk() (*NotificationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NotificationSettingsRequestsDto) SetType(v NotificationType) {
	o.Type = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *NotificationSettingsRequestsDto) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSettingsRequestsDto) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *NotificationSettingsRequestsDto) IsIsEnabledSet() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *NotificationSettingsRequestsDto) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

func (o NotificationSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationSettingsRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	return toSerialize, nil
}

func (o *NotificationSettingsRequestsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNotificationSettingsRequestsDto := _NotificationSettingsRequestsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varNotificationSettingsRequestsDto)

	if err != nil {
		return err
	}

	*o = NotificationSettingsRequestsDto(varNotificationSettingsRequestsDto)

	return err
}

type NullableNotificationSettingsRequestsDto struct {
	value *NotificationSettingsRequestsDto
	isSet bool
}

func (v NullableNotificationSettingsRequestsDto) Get() *NotificationSettingsRequestsDto {
	return v.value
}

func (v *NullableNotificationSettingsRequestsDto) Set(val *NotificationSettingsRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSettingsRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSettingsRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSettingsRequestsDto(val *NotificationSettingsRequestsDto) *NullableNotificationSettingsRequestsDto {
	return &NullableNotificationSettingsRequestsDto{value: val, isSet: true}
}

func (v NullableNotificationSettingsRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSettingsRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

