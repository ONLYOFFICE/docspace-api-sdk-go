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

// checks if the NotificationChannelDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationChannelDto{}

// NotificationChannelDto The notification channel information.
type NotificationChannelDto struct {
	// The notification channel name.
	Name NullableString `json:"name"`
	// Specifies whether the notification channel is enabled.
	IsEnabled bool `json:"isEnabled"`
}

type _NotificationChannelDto NotificationChannelDto

// NewNotificationChannelDto instantiates a new NotificationChannelDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationChannelDto(name NullableString, isEnabled bool) *NotificationChannelDto {
	this := NotificationChannelDto{}
	this.Name = name
	this.IsEnabled = isEnabled
	return &this
}

// NewNotificationChannelDtoWithDefaults instantiates a new NotificationChannelDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationChannelDtoWithDefaults() *NotificationChannelDto {
	this := NotificationChannelDto{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NotificationChannelDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationChannelDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *NotificationChannelDto) SetName(v string) {
	o.Name.Set(&v)
}

// GetIsEnabled returns the IsEnabled field value
func (o *NotificationChannelDto) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *NotificationChannelDto) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *NotificationChannelDto) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

func (o NotificationChannelDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationChannelDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name.Get()
	toSerialize["isEnabled"] = o.IsEnabled
	return toSerialize, nil
}

func (o *NotificationChannelDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"isEnabled",
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

	varNotificationChannelDto := _NotificationChannelDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varNotificationChannelDto)

	if err != nil {
		return err
	}

	*o = NotificationChannelDto(varNotificationChannelDto)

	return err
}

type NullableNotificationChannelDto struct {
	value *NotificationChannelDto
	isSet bool
}

func (v NullableNotificationChannelDto) Get() *NotificationChannelDto {
	return v.value
}

func (v *NullableNotificationChannelDto) Set(val *NotificationChannelDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationChannelDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationChannelDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationChannelDto(val *NotificationChannelDto) *NullableNotificationChannelDto {
	return &NullableNotificationChannelDto{value: val, isSet: true}
}

func (v NullableNotificationChannelDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationChannelDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

