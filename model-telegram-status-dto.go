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

// checks if the TelegramStatusDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelegramStatusDto{}

// TelegramStatusDto Whether the calling user's account is linked to the portal's Telegram bot.
type TelegramStatusDto struct {
	// Where the caller's own account stands: not linked, linked, or a registration link issued and the portal  still waiting for it to be opened in Telegram. The waiting state ends on its own when the link expires,  so it is worth polling rather than treating as final.
	Status RegStatus `json:"status"`
	// The Telegram handle the account is linked to, without the leading `@`. It is filled in only while the  account is linked and comes back empty in the other two states.
	Username NullableString `json:"username,omitempty"`
}

type _TelegramStatusDto TelegramStatusDto

// NewTelegramStatusDto instantiates a new TelegramStatusDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelegramStatusDto(status RegStatus) *TelegramStatusDto {
	this := TelegramStatusDto{}
	this.Status = status
	return &this
}

// NewTelegramStatusDtoWithDefaults instantiates a new TelegramStatusDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelegramStatusDtoWithDefaults() *TelegramStatusDto {
	this := TelegramStatusDto{}
	return &this
}

// GetStatus returns the Status field value
func (o *TelegramStatusDto) GetStatus() RegStatus {
	if o == nil {
		var ret RegStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TelegramStatusDto) GetStatusOk() (*RegStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TelegramStatusDto) SetStatus(v RegStatus) {
	o.Status = v
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TelegramStatusDto) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TelegramStatusDto) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *TelegramStatusDto) IsUsernameSet() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *TelegramStatusDto) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *TelegramStatusDto) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *TelegramStatusDto) UnsetUsername() {
	o.Username.Unset()
}

func (o TelegramStatusDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelegramStatusDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	return toSerialize, nil
}

func (o *TelegramStatusDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varTelegramStatusDto := _TelegramStatusDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varTelegramStatusDto)

	if err != nil {
		return err
	}

	*o = TelegramStatusDto(varTelegramStatusDto)

	return err
}

type NullableTelegramStatusDto struct {
	value *TelegramStatusDto
	isSet bool
}

func (v NullableTelegramStatusDto) Get() *TelegramStatusDto {
	return v.value
}

func (v *NullableTelegramStatusDto) Set(val *TelegramStatusDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTelegramStatusDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTelegramStatusDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelegramStatusDto(val *TelegramStatusDto) *NullableTelegramStatusDto {
	return &NullableTelegramStatusDto{value: val, isSet: true}
}

func (v NullableTelegramStatusDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelegramStatusDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

