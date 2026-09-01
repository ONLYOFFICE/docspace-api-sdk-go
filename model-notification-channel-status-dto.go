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

// checks if the NotificationChannelStatusDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationChannelStatusDto{}

// NotificationChannelStatusDto The notification channel settings.
type NotificationChannelStatusDto struct {
	// The list of notification channels.
	Channels []NotificationChannelDto `json:"channels,omitempty"`
}

// NewNotificationChannelStatusDto instantiates a new NotificationChannelStatusDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationChannelStatusDto() *NotificationChannelStatusDto {
	this := NotificationChannelStatusDto{}
	return &this
}

// NewNotificationChannelStatusDtoWithDefaults instantiates a new NotificationChannelStatusDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationChannelStatusDtoWithDefaults() *NotificationChannelStatusDto {
	this := NotificationChannelStatusDto{}
	return &this
}

// GetChannels returns the Channels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationChannelStatusDto) GetChannels() []NotificationChannelDto {
	if o == nil {
		var ret []NotificationChannelDto
		return ret
	}
	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationChannelStatusDto) GetChannelsOk() ([]NotificationChannelDto, bool) {
	if o == nil || IsNil(o.Channels) {
		return nil, false
	}
	return o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *NotificationChannelStatusDto) IsChannelsSet() bool {
	if o != nil && !IsNil(o.Channels) {
		return true
	}

	return false
}

// SetChannels gets a reference to the given []NotificationChannelDto and assigns it to the Channels field.
func (o *NotificationChannelStatusDto) SetChannels(v []NotificationChannelDto) {
	o.Channels = v
}

func (o NotificationChannelStatusDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationChannelStatusDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Channels != nil {
		toSerialize["channels"] = o.Channels
	}
	return toSerialize, nil
}

type NullableNotificationChannelStatusDto struct {
	value *NotificationChannelStatusDto
	isSet bool
}

func (v NullableNotificationChannelStatusDto) Get() *NotificationChannelStatusDto {
	return v.value
}

func (v *NullableNotificationChannelStatusDto) Set(val *NotificationChannelStatusDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationChannelStatusDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationChannelStatusDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationChannelStatusDto(val *NotificationChannelStatusDto) *NullableNotificationChannelStatusDto {
	return &NullableNotificationChannelStatusDto{value: val, isSet: true}
}

func (v NullableNotificationChannelStatusDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationChannelStatusDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

