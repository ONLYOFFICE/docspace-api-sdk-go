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
	"time"
)

// checks if the TimeBoundDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeBoundDto{}

// TimeBoundDto Represents the period the price is effective in.
type TimeBoundDto struct {
	// The date and time when the period starts.
	StartDate *time.Time `json:"startDate,omitempty"`
	// The date and time when the period ends.
	EndDate NullableTime `json:"endDate,omitempty"`
}

// NewTimeBoundDto instantiates a new TimeBoundDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeBoundDto() *TimeBoundDto {
	this := TimeBoundDto{}
	return &this
}

// NewTimeBoundDtoWithDefaults instantiates a new TimeBoundDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeBoundDtoWithDefaults() *TimeBoundDto {
	this := TimeBoundDto{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *TimeBoundDto) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeBoundDto) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *TimeBoundDto) IsStartDateSet() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *TimeBoundDto) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeBoundDto) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeBoundDto) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *TimeBoundDto) IsEndDateSet() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *TimeBoundDto) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *TimeBoundDto) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *TimeBoundDto) UnsetEndDate() {
	o.EndDate.Unset()
}

func (o TimeBoundDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeBoundDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	return toSerialize, nil
}

type NullableTimeBoundDto struct {
	value *TimeBoundDto
	isSet bool
}

func (v NullableTimeBoundDto) Get() *TimeBoundDto {
	return v.value
}

func (v *NullableTimeBoundDto) Set(val *TimeBoundDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeBoundDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeBoundDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeBoundDto(val *TimeBoundDto) *NullableTimeBoundDto {
	return &NullableTimeBoundDto{value: val, isSet: true}
}

func (v NullableTimeBoundDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeBoundDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

