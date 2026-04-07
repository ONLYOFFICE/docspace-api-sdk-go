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

// checks if the Cron type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cron{}

// Cron The backup cron parameters.
type Cron struct {
	Period *BackupPeriod `json:"period,omitempty"`
	// The time of the day to start the backup process.
	Hour *int32 `json:"hour,omitempty"`
	// The day of the week to start the backup process.
	Day NullableInt32 `json:"day,omitempty"`
}

// NewCron instantiates a new Cron object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCron() *Cron {
	this := Cron{}
	return &this
}

// NewCronWithDefaults instantiates a new Cron object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCronWithDefaults() *Cron {
	this := Cron{}
	return &this
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *Cron) GetPeriod() BackupPeriod {
	if o == nil || IsNil(o.Period) {
		var ret BackupPeriod
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cron) GetPeriodOk() (*BackupPeriod, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *Cron) IsPeriodSet() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given BackupPeriod and assigns it to the Period field.
func (o *Cron) SetPeriod(v BackupPeriod) {
	o.Period = &v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *Cron) GetHour() int32 {
	if o == nil || IsNil(o.Hour) {
		var ret int32
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cron) GetHourOk() (*int32, bool) {
	if o == nil || IsNil(o.Hour) {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *Cron) IsHourSet() bool {
	if o != nil && !IsNil(o.Hour) {
		return true
	}

	return false
}

// SetHour gets a reference to the given int32 and assigns it to the Hour field.
func (o *Cron) SetHour(v int32) {
	o.Hour = &v
}

// GetDay returns the Day field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cron) GetDay() int32 {
	if o == nil || IsNil(o.Day.Get()) {
		var ret int32
		return ret
	}
	return *o.Day.Get()
}

// GetDayOk returns a tuple with the Day field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cron) GetDayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Day.Get(), o.Day.IsSet()
}

// HasDay returns a boolean if a field has been set.
func (o *Cron) IsDaySet() bool {
	if o != nil && o.Day.IsSet() {
		return true
	}

	return false
}

// SetDay gets a reference to the given NullableInt32 and assigns it to the Day field.
func (o *Cron) SetDay(v int32) {
	o.Day.Set(&v)
}
// SetDayNil sets the value for Day to be an explicit nil
func (o *Cron) SetDayNil() {
	o.Day.Set(nil)
}

// UnsetDay ensures that no value is present for Day, not even an explicit nil
func (o *Cron) UnsetDay() {
	o.Day.Unset()
}

func (o Cron) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cron) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !IsNil(o.Hour) {
		toSerialize["hour"] = o.Hour
	}
	if o.Day.IsSet() {
		toSerialize["day"] = o.Day.Get()
	}
	return toSerialize, nil
}

type NullableCron struct {
	value *Cron
	isSet bool
}

func (v NullableCron) Get() *Cron {
	return v.value
}

func (v *NullableCron) Set(val *Cron) {
	v.value = val
	v.isSet = true
}

func (v NullableCron) IsSet() bool {
	return v.isSet
}

func (v *NullableCron) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCron(val *Cron) *NullableCron {
	return &NullableCron{value: val, isSet: true}
}

func (v NullableCron) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCron) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

