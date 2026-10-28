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

// checks if the CronParamsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CronParamsDto{}

// CronParamsDto The time a scheduled backup runs at.
type CronParamsDto struct {
	// How often the backup runs: 0 for every day, 1 for every week and 2 for every month.
	Period *BackupPeriod `json:"period,omitempty"`
	// The hour of the day the backup starts at, from 0 to 23.
	Hour *int32 `json:"hour,omitempty"`
	// The day the backup runs on: the day of the week from 1 to 7, Sunday being 1, for a weekly schedule,  and the day of the month from 1 to 31 for a monthly one. It is 0 for a daily schedule.
	Day *int32 `json:"day,omitempty"`
}

// NewCronParamsDto instantiates a new CronParamsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCronParamsDto() *CronParamsDto {
	this := CronParamsDto{}
	return &this
}

// NewCronParamsDtoWithDefaults instantiates a new CronParamsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCronParamsDtoWithDefaults() *CronParamsDto {
	this := CronParamsDto{}
	return &this
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *CronParamsDto) GetPeriod() BackupPeriod {
	if o == nil || IsNil(o.Period) {
		var ret BackupPeriod
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronParamsDto) GetPeriodOk() (*BackupPeriod, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *CronParamsDto) IsPeriodSet() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given BackupPeriod and assigns it to the Period field.
func (o *CronParamsDto) SetPeriod(v BackupPeriod) {
	o.Period = &v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *CronParamsDto) GetHour() int32 {
	if o == nil || IsNil(o.Hour) {
		var ret int32
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronParamsDto) GetHourOk() (*int32, bool) {
	if o == nil || IsNil(o.Hour) {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *CronParamsDto) IsHourSet() bool {
	if o != nil && !IsNil(o.Hour) {
		return true
	}

	return false
}

// SetHour gets a reference to the given int32 and assigns it to the Hour field.
func (o *CronParamsDto) SetHour(v int32) {
	o.Hour = &v
}

// GetDay returns the Day field value if set, zero value otherwise.
func (o *CronParamsDto) GetDay() int32 {
	if o == nil || IsNil(o.Day) {
		var ret int32
		return ret
	}
	return *o.Day
}

// GetDayOk returns a tuple with the Day field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronParamsDto) GetDayOk() (*int32, bool) {
	if o == nil || IsNil(o.Day) {
		return nil, false
	}
	return o.Day, true
}

// HasDay returns a boolean if a field has been set.
func (o *CronParamsDto) IsDaySet() bool {
	if o != nil && !IsNil(o.Day) {
		return true
	}

	return false
}

// SetDay gets a reference to the given int32 and assigns it to the Day field.
func (o *CronParamsDto) SetDay(v int32) {
	o.Day = &v
}

func (o CronParamsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CronParamsDto) ToMap() (map[string]interface{}, error) {
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

type NullableCronParamsDto struct {
	value *CronParamsDto
	isSet bool
}

func (v NullableCronParamsDto) Get() *CronParamsDto {
	return v.value
}

func (v *NullableCronParamsDto) Set(val *CronParamsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCronParamsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCronParamsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCronParamsDto(val *CronParamsDto) *NullableCronParamsDto {
	return &NullableCronParamsDto{value: val, isSet: true}
}

func (v NullableCronParamsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCronParamsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

