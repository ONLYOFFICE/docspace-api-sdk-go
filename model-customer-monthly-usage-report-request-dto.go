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

// checks if the CustomerMonthlyUsageReportRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerMonthlyUsageReportRequestDto{}

// CustomerMonthlyUsageReportRequestDto The request parameters for generating a customer monthly usage report.
type CustomerMonthlyUsageReportRequestDto struct {
	// The report start date.
	StartDate NullableTime `json:"startDate,omitempty"`
	// The report end date.
	EndDate NullableTime `json:"endDate,omitempty"`
}

// NewCustomerMonthlyUsageReportRequestDto instantiates a new CustomerMonthlyUsageReportRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerMonthlyUsageReportRequestDto() *CustomerMonthlyUsageReportRequestDto {
	this := CustomerMonthlyUsageReportRequestDto{}
	return &this
}

// NewCustomerMonthlyUsageReportRequestDtoWithDefaults instantiates a new CustomerMonthlyUsageReportRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerMonthlyUsageReportRequestDtoWithDefaults() *CustomerMonthlyUsageReportRequestDto {
	this := CustomerMonthlyUsageReportRequestDto{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerMonthlyUsageReportRequestDto) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerMonthlyUsageReportRequestDto) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *CustomerMonthlyUsageReportRequestDto) IsStartDateSet() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableTime and assigns it to the StartDate field.
func (o *CustomerMonthlyUsageReportRequestDto) SetStartDate(v time.Time) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *CustomerMonthlyUsageReportRequestDto) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *CustomerMonthlyUsageReportRequestDto) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerMonthlyUsageReportRequestDto) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerMonthlyUsageReportRequestDto) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *CustomerMonthlyUsageReportRequestDto) IsEndDateSet() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *CustomerMonthlyUsageReportRequestDto) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *CustomerMonthlyUsageReportRequestDto) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *CustomerMonthlyUsageReportRequestDto) UnsetEndDate() {
	o.EndDate.Unset()
}

func (o CustomerMonthlyUsageReportRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerMonthlyUsageReportRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.StartDate.IsSet() {
		toSerialize["startDate"] = o.StartDate.Get()
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	return toSerialize, nil
}

type NullableCustomerMonthlyUsageReportRequestDto struct {
	value *CustomerMonthlyUsageReportRequestDto
	isSet bool
}

func (v NullableCustomerMonthlyUsageReportRequestDto) Get() *CustomerMonthlyUsageReportRequestDto {
	return v.value
}

func (v *NullableCustomerMonthlyUsageReportRequestDto) Set(val *CustomerMonthlyUsageReportRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerMonthlyUsageReportRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerMonthlyUsageReportRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerMonthlyUsageReportRequestDto(val *CustomerMonthlyUsageReportRequestDto) *NullableCustomerMonthlyUsageReportRequestDto {
	return &NullableCustomerMonthlyUsageReportRequestDto{value: val, isSet: true}
}

func (v NullableCustomerMonthlyUsageReportRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerMonthlyUsageReportRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

