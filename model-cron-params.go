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

// checks if the CronParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CronParams{}

// CronParams The time a scheduled backup runs at.
type CronParams struct {
	// How often the backup runs: 0 for every day, 1 for every week and 2 for every month.
	Period *BackupPeriod `json:"period,omitempty"`
	// The hour of the day the backup starts at, from 0 to 23.
	Hour *int32 `json:"hour,omitempty"`
	// The day the backup runs on: the day of the week from 1 to 7, Sunday being 1, for a weekly schedule,  and the day of the month from 1 to 31 for a monthly one. It is 0 for a daily schedule.
	Day *int32 `json:"day,omitempty"`
}

// NewCronParams instantiates a new CronParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCronParams() *CronParams {
	this := CronParams{}
	return &this
}

// NewCronParamsWithDefaults instantiates a new CronParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCronParamsWithDefaults() *CronParams {
	this := CronParams{}
	return &this
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *CronParams) GetPeriod() BackupPeriod {
	if o == nil || IsNil(o.Period) {
		var ret BackupPeriod
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronParams) GetPeriodOk() (*BackupPeriod, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *CronParams) IsPeriodSet() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given BackupPeriod and assigns it to the Period field.
func (o *CronParams) SetPeriod(v BackupPeriod) {
	o.Period = &v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *CronParams) GetHour() int32 {
	if o == nil || IsNil(o.Hour) {
		var ret int32
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronParams) GetHourOk() (*int32, bool) {
	if o == nil || IsNil(o.Hour) {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *CronParams) IsHourSet() bool {
	if o != nil && !IsNil(o.Hour) {
		return true
	}

	return false
}

// SetHour gets a reference to the given int32 and assigns it to the Hour field.
func (o *CronParams) SetHour(v int32) {
	o.Hour = &v
}

// GetDay returns the Day field value if set, zero value otherwise.
func (o *CronParams) GetDay() int32 {
	if o == nil || IsNil(o.Day) {
		var ret int32
		return ret
	}
	return *o.Day
}

// GetDayOk returns a tuple with the Day field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronParams) GetDayOk() (*int32, bool) {
	if o == nil || IsNil(o.Day) {
		return nil, false
	}
	return o.Day, true
}

// HasDay returns a boolean if a field has been set.
func (o *CronParams) IsDaySet() bool {
	if o != nil && !IsNil(o.Day) {
		return true
	}

	return false
}

// SetDay gets a reference to the given int32 and assigns it to the Day field.
func (o *CronParams) SetDay(v int32) {
	o.Day = &v
}

func (o CronParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CronParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !IsNil(o.Hour) {
		toSerialize["hour"] = o.Hour
	}
	if !IsNil(o.Day) {
		toSerialize["day"] = o.Day
	}
	return toSerialize, nil
}

type NullableCronParams struct {
	value *CronParams
	isSet bool
}

func (v NullableCronParams) Get() *CronParams {
	return v.value
}

func (v *NullableCronParams) Set(val *CronParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCronParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCronParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCronParams(val *CronParams) *NullableCronParams {
	return &NullableCronParams{value: val, isSet: true}
}

func (v NullableCronParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCronParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

