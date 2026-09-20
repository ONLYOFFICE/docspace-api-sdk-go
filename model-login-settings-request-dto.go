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

// checks if the LoginSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginSettingsRequestDto{}

// LoginSettingsRequestDto The brute-force protection of the sign-in form: how many failures, over how long, cost how long a block.
type LoginSettingsRequestDto struct {
	// How many failed sign-in attempts inside one window are tolerated before the offender is blocked. Attempts are  counted per user name and client address together, so one member being blocked leaves the rest of the portal  signing in normally.
	AttemptCount *int32 `json:"attemptCount,omitempty"`
	// How long, in seconds, a blocked user name and address pair stays refused. While the block lasts the sign-in  is refused even when the password is finally correct.
	BlockTime *int32 `json:"blockTime,omitempty"`
	// The length, in seconds, of the rolling window the failed attempts are counted over. A wider window makes the  same `attemptCount` stricter, because failures further apart still add up.
	CheckPeriod *int32 `json:"checkPeriod,omitempty"`
}

// NewLoginSettingsRequestDto instantiates a new LoginSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginSettingsRequestDto() *LoginSettingsRequestDto {
	this := LoginSettingsRequestDto{}
	return &this
}

// NewLoginSettingsRequestDtoWithDefaults instantiates a new LoginSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginSettingsRequestDtoWithDefaults() *LoginSettingsRequestDto {
	this := LoginSettingsRequestDto{}
	return &this
}

// GetAttemptCount returns the AttemptCount field value if set, zero value otherwise.
func (o *LoginSettingsRequestDto) GetAttemptCount() int32 {
	if o == nil || IsNil(o.AttemptCount) {
		var ret int32
		return ret
	}
	return *o.AttemptCount
}

// GetAttemptCountOk returns a tuple with the AttemptCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginSettingsRequestDto) GetAttemptCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AttemptCount) {
		return nil, false
	}
	return o.AttemptCount, true
}

// HasAttemptCount returns a boolean if a field has been set.
func (o *LoginSettingsRequestDto) IsAttemptCountSet() bool {
	if o != nil && !IsNil(o.AttemptCount) {
		return true
	}

	return false
}

// SetAttemptCount gets a reference to the given int32 and assigns it to the AttemptCount field.
func (o *LoginSettingsRequestDto) SetAttemptCount(v int32) {
	o.AttemptCount = &v
}

// GetBlockTime returns the BlockTime field value if set, zero value otherwise.
func (o *LoginSettingsRequestDto) GetBlockTime() int32 {
	if o == nil || IsNil(o.BlockTime) {
		var ret int32
		return ret
	}
	return *o.BlockTime
}

// GetBlockTimeOk returns a tuple with the BlockTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginSettingsRequestDto) GetBlockTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.BlockTime) {
		return nil, false
	}
	return o.BlockTime, true
}

// HasBlockTime returns a boolean if a field has been set.
func (o *LoginSettingsRequestDto) IsBlockTimeSet() bool {
	if o != nil && !IsNil(o.BlockTime) {
		return true
	}

	return false
}

// SetBlockTime gets a reference to the given int32 and assigns it to the BlockTime field.
func (o *LoginSettingsRequestDto) SetBlockTime(v int32) {
	o.BlockTime = &v
}

// GetCheckPeriod returns the CheckPeriod field value if set, zero value otherwise.
func (o *LoginSettingsRequestDto) GetCheckPeriod() int32 {
	if o == nil || IsNil(o.CheckPeriod) {
		var ret int32
		return ret
	}
	return *o.CheckPeriod
}

// GetCheckPeriodOk returns a tuple with the CheckPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginSettingsRequestDto) GetCheckPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.CheckPeriod) {
		return nil, false
	}
	return o.CheckPeriod, true
}

// HasCheckPeriod returns a boolean if a field has been set.
func (o *LoginSettingsRequestDto) IsCheckPeriodSet() bool {
	if o != nil && !IsNil(o.CheckPeriod) {
		return true
	}

	return false
}

// SetCheckPeriod gets a reference to the given int32 and assigns it to the CheckPeriod field.
func (o *LoginSettingsRequestDto) SetCheckPeriod(v int32) {
	o.CheckPeriod = &v
}

func (o LoginSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttemptCount) {
		toSerialize["attemptCount"] = o.AttemptCount
	}
	if !IsNil(o.BlockTime) {
		toSerialize["blockTime"] = o.BlockTime
	}
	if !IsNil(o.CheckPeriod) {
		toSerialize["checkPeriod"] = o.CheckPeriod
	}
	return toSerialize, nil
}

type NullableLoginSettingsRequestDto struct {
	value *LoginSettingsRequestDto
	isSet bool
}

func (v NullableLoginSettingsRequestDto) Get() *LoginSettingsRequestDto {
	return v.value
}

func (v *NullableLoginSettingsRequestDto) Set(val *LoginSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginSettingsRequestDto(val *LoginSettingsRequestDto) *NullableLoginSettingsRequestDto {
	return &NullableLoginSettingsRequestDto{value: val, isSet: true}
}

func (v NullableLoginSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

