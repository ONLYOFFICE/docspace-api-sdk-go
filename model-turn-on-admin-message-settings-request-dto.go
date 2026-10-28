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

// checks if the TurnOnAdminMessageSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TurnOnAdminMessageSettingsRequestDto{}

// TurnOnAdminMessageSettingsRequestDto Whether the sign-in page offers the form for writing to the portal administrators.
type TurnOnAdminMessageSettingsRequestDto struct {
	// Whether the form is offered. Switching it off hides the form for everybody and makes the operation that  submits it refuse new messages; letters already sent are untouched.
	TurnOn *bool `json:"turnOn,omitempty"`
}

// NewTurnOnAdminMessageSettingsRequestDto instantiates a new TurnOnAdminMessageSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTurnOnAdminMessageSettingsRequestDto() *TurnOnAdminMessageSettingsRequestDto {
	this := TurnOnAdminMessageSettingsRequestDto{}
	return &this
}

// NewTurnOnAdminMessageSettingsRequestDtoWithDefaults instantiates a new TurnOnAdminMessageSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTurnOnAdminMessageSettingsRequestDtoWithDefaults() *TurnOnAdminMessageSettingsRequestDto {
	this := TurnOnAdminMessageSettingsRequestDto{}
	return &this
}

// GetTurnOn returns the TurnOn field value if set, zero value otherwise.
func (o *TurnOnAdminMessageSettingsRequestDto) GetTurnOn() bool {
	if o == nil || IsNil(o.TurnOn) {
		var ret bool
		return ret
	}
	return *o.TurnOn
}

// GetTurnOnOk returns a tuple with the TurnOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TurnOnAdminMessageSettingsRequestDto) GetTurnOnOk() (*bool, bool) {
	if o == nil || IsNil(o.TurnOn) {
		return nil, false
	}
	return o.TurnOn, true
}

// HasTurnOn returns a boolean if a field has been set.
func (o *TurnOnAdminMessageSettingsRequestDto) IsTurnOnSet() bool {
	if o != nil && !IsNil(o.TurnOn) {
		return true
	}

	return false
}

// SetTurnOn gets a reference to the given bool and assigns it to the TurnOn field.
func (o *TurnOnAdminMessageSettingsRequestDto) SetTurnOn(v bool) {
	o.TurnOn = &v
}

func (o TurnOnAdminMessageSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TurnOnAdminMessageSettingsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TurnOn) {
		toSerialize["turnOn"] = o.TurnOn
	}
	return toSerialize, nil
}

type NullableTurnOnAdminMessageSettingsRequestDto struct {
	value *TurnOnAdminMessageSettingsRequestDto
	isSet bool
}

func (v NullableTurnOnAdminMessageSettingsRequestDto) Get() *TurnOnAdminMessageSettingsRequestDto {
	return v.value
}

func (v *NullableTurnOnAdminMessageSettingsRequestDto) Set(val *TurnOnAdminMessageSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTurnOnAdminMessageSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTurnOnAdminMessageSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTurnOnAdminMessageSettingsRequestDto(val *TurnOnAdminMessageSettingsRequestDto) *NullableTurnOnAdminMessageSettingsRequestDto {
	return &NullableTurnOnAdminMessageSettingsRequestDto{value: val, isSet: true}
}

func (v NullableTurnOnAdminMessageSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTurnOnAdminMessageSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

