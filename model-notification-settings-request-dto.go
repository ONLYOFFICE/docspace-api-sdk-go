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

// checks if the NotificationSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationSettingsRequestDto{}

// NotificationSettingsRequestDto Which kind of notification the calling user switches, and which way.
type NotificationSettingsRequestDto struct {
	// The kind of notification being switched. A value outside the defined set is echoed back while nothing is  stored, so confirm the result with `GET api/2.0/settings/notification/{type}` rather than trusting the  answer.
	Type NotificationType `json:"type"`
	// Whether that kind reaches the calling account. It applies to the caller own account alone and to every room  at once; a single room is silenced with `POST api/2.0/settings/notification/rooms` instead.
	IsEnabled *bool `json:"isEnabled,omitempty"`
}

type _NotificationSettingsRequestDto NotificationSettingsRequestDto

// NewNotificationSettingsRequestDto instantiates a new NotificationSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSettingsRequestDto(type_ NotificationType) *NotificationSettingsRequestDto {
	this := NotificationSettingsRequestDto{}
	this.Type = type_
	return &this
}

// NewNotificationSettingsRequestDtoWithDefaults instantiates a new NotificationSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSettingsRequestDtoWithDefaults() *NotificationSettingsRequestDto {
	this := NotificationSettingsRequestDto{}
	return &this
}

// GetType returns the Type field value
func (o *NotificationSettingsRequestDto) GetType() NotificationType {
	if o == nil {
		var ret NotificationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotificationSettingsRequestDto) GetTypeOk() (*NotificationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NotificationSettingsRequestDto) SetType(v NotificationType) {
	o.Type = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *NotificationSettingsRequestDto) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSettingsRequestDto) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *NotificationSettingsRequestDto) IsIsEnabledSet() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *NotificationSettingsRequestDto) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

func (o NotificationSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	return toSerialize, nil
}

func (o *NotificationSettingsRequestDto) UnmarshalJSON(data []byte) (err error) {
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

	varNotificationSettingsRequestDto := _NotificationSettingsRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varNotificationSettingsRequestDto)

	if err != nil {
		return err
	}

	*o = NotificationSettingsRequestDto(varNotificationSettingsRequestDto)

	return err
}

type NullableNotificationSettingsRequestDto struct {
	value *NotificationSettingsRequestDto
	isSet bool
}

func (v NullableNotificationSettingsRequestDto) Get() *NotificationSettingsRequestDto {
	return v.value
}

func (v *NullableNotificationSettingsRequestDto) Set(val *NotificationSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSettingsRequestDto(val *NotificationSettingsRequestDto) *NullableNotificationSettingsRequestDto {
	return &NullableNotificationSettingsRequestDto{value: val, isSet: true}
}

func (v NullableNotificationSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

