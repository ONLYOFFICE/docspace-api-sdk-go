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

// checks if the LoginSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginSettingsDto{}

// LoginSettingsDto The brute-force protection of the sign-in form: how many failures, over how long, cost how long a block.
type LoginSettingsDto struct {
	// How many failed attempts inside one window are tolerated before the offender is blocked. Attempts are  counted per user name and client address together, so one member being blocked leaves the rest of the  portal signing in normally.
	AttemptCount int32 `json:"attemptCount"`
	// How long, in seconds, a blocked user name and address pair stays refused. While the block lasts the  sign-in is refused even once the password is correct.
	BlockTime int32 `json:"blockTime"`
	// The length, in seconds, of the rolling window the failures are counted over. It is not a request timeout: a  wider window makes the same `attemptCount` stricter, because failures further apart still add up.
	CheckPeriod int32 `json:"checkPeriod"`
	// Whether the three numbers above still match the ones the installation ships with. It turns `false` as soon  as any of them is saved differently, and `true` again after  `DELETE api/2.0/settings/security/loginsettings`.
	IsDefault bool `json:"isDefault"`
}

type _LoginSettingsDto LoginSettingsDto

// NewLoginSettingsDto instantiates a new LoginSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginSettingsDto(attemptCount int32, blockTime int32, checkPeriod int32, isDefault bool) *LoginSettingsDto {
	this := LoginSettingsDto{}
	this.AttemptCount = attemptCount
	this.BlockTime = blockTime
	this.CheckPeriod = checkPeriod
	this.IsDefault = isDefault
	return &this
}

// NewLoginSettingsDtoWithDefaults instantiates a new LoginSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginSettingsDtoWithDefaults() *LoginSettingsDto {
	this := LoginSettingsDto{}
	return &this
}

// GetAttemptCount returns the AttemptCount field value
func (o *LoginSettingsDto) GetAttemptCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AttemptCount
}

// GetAttemptCountOk returns a tuple with the AttemptCount field value
// and a boolean to check if the value has been set.
func (o *LoginSettingsDto) GetAttemptCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttemptCount, true
}

// SetAttemptCount sets field value
func (o *LoginSettingsDto) SetAttemptCount(v int32) {
	o.AttemptCount = v
}

// GetBlockTime returns the BlockTime field value
func (o *LoginSettingsDto) GetBlockTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockTime
}

// GetBlockTimeOk returns a tuple with the BlockTime field value
// and a boolean to check if the value has been set.
func (o *LoginSettingsDto) GetBlockTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockTime, true
}

// SetBlockTime sets field value
func (o *LoginSettingsDto) SetBlockTime(v int32) {
	o.BlockTime = v
}

// GetCheckPeriod returns the CheckPeriod field value
func (o *LoginSettingsDto) GetCheckPeriod() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CheckPeriod
}

// GetCheckPeriodOk returns a tuple with the CheckPeriod field value
// and a boolean to check if the value has been set.
func (o *LoginSettingsDto) GetCheckPeriodOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckPeriod, true
}

// SetCheckPeriod sets field value
func (o *LoginSettingsDto) SetCheckPeriod(v int32) {
	o.CheckPeriod = v
}

// GetIsDefault returns the IsDefault field value
func (o *LoginSettingsDto) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *LoginSettingsDto) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *LoginSettingsDto) SetIsDefault(v bool) {
	o.IsDefault = v
}

func (o LoginSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attemptCount"] = o.AttemptCount
	toSerialize["blockTime"] = o.BlockTime
	toSerialize["checkPeriod"] = o.CheckPeriod
	toSerialize["isDefault"] = o.IsDefault
	return toSerialize, nil
}

func (o *LoginSettingsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attemptCount",
		"blockTime",
		"checkPeriod",
		"isDefault",
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

	varLoginSettingsDto := _LoginSettingsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varLoginSettingsDto)

	if err != nil {
		return err
	}

	*o = LoginSettingsDto(varLoginSettingsDto)

	return err
}

type NullableLoginSettingsDto struct {
	value *LoginSettingsDto
	isSet bool
}

func (v NullableLoginSettingsDto) Get() *LoginSettingsDto {
	return v.value
}

func (v *NullableLoginSettingsDto) Set(val *LoginSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginSettingsDto(val *LoginSettingsDto) *NullableLoginSettingsDto {
	return &NullableLoginSettingsDto{value: val, isSet: true}
}

func (v NullableLoginSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

