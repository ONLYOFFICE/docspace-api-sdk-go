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

// checks if the ServiceLimitUsageDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceLimitUsageDto{}

// ServiceLimitUsageDto What has been spent against a spending limit in its current period, and what is left.
type ServiceLimitUsageDto struct {
	// The start of the current period, in UTC.
	PeriodStart *time.Time `json:"periodStart,omitempty"`
	// The end of the current period, in UTC: the moment the spending starts again from zero. Null for a lifetime  limit, which never resets.
	PeriodEnd NullableTime `json:"periodEnd,omitempty"`
	// The money charged in the current period, in the currency of the limit.
	AmountConsumed *float64 `json:"amountConsumed,omitempty"`
	// The money still available in the current period: the threshold less what was charged and what is reserved for  operations in progress, never below zero. Null when the limit has no money threshold.
	AmountAvailable NullableFloat64 `json:"amountAvailable,omitempty"`
	// The quantity consumed in the current period, in the unit of the limit.
	QuantityConsumed *int32 `json:"quantityConsumed,omitempty"`
	// The quantity still available in the current period, never below zero; null when the limit has no quantity  threshold.
	QuantityAvailable NullableInt32 `json:"quantityAvailable,omitempty"`
}

// NewServiceLimitUsageDto instantiates a new ServiceLimitUsageDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceLimitUsageDto() *ServiceLimitUsageDto {
	this := ServiceLimitUsageDto{}
	return &this
}

// NewServiceLimitUsageDtoWithDefaults instantiates a new ServiceLimitUsageDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceLimitUsageDtoWithDefaults() *ServiceLimitUsageDto {
	this := ServiceLimitUsageDto{}
	return &this
}

// GetPeriodStart returns the PeriodStart field value if set, zero value otherwise.
func (o *ServiceLimitUsageDto) GetPeriodStart() time.Time {
	if o == nil || IsNil(o.PeriodStart) {
		var ret time.Time
		return ret
	}
	return *o.PeriodStart
}

// GetPeriodStartOk returns a tuple with the PeriodStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLimitUsageDto) GetPeriodStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PeriodStart) {
		return nil, false
	}
	return o.PeriodStart, true
}

// HasPeriodStart returns a boolean if a field has been set.
func (o *ServiceLimitUsageDto) IsPeriodStartSet() bool {
	if o != nil && !IsNil(o.PeriodStart) {
		return true
	}

	return false
}

// SetPeriodStart gets a reference to the given time.Time and assigns it to the PeriodStart field.
func (o *ServiceLimitUsageDto) SetPeriodStart(v time.Time) {
	o.PeriodStart = &v
}

// GetPeriodEnd returns the PeriodEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceLimitUsageDto) GetPeriodEnd() time.Time {
	if o == nil || IsNil(o.PeriodEnd.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PeriodEnd.Get()
}

// GetPeriodEndOk returns a tuple with the PeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceLimitUsageDto) GetPeriodEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PeriodEnd.Get(), o.PeriodEnd.IsSet()
}

// HasPeriodEnd returns a boolean if a field has been set.
func (o *ServiceLimitUsageDto) IsPeriodEndSet() bool {
	if o != nil && o.PeriodEnd.IsSet() {
		return true
	}

	return false
}

// SetPeriodEnd gets a reference to the given NullableTime and assigns it to the PeriodEnd field.
func (o *ServiceLimitUsageDto) SetPeriodEnd(v time.Time) {
	o.PeriodEnd.Set(&v)
}
// SetPeriodEndNil sets the value for PeriodEnd to be an explicit nil
func (o *ServiceLimitUsageDto) SetPeriodEndNil() {
	o.PeriodEnd.Set(nil)
}

// UnsetPeriodEnd ensures that no value is present for PeriodEnd, not even an explicit nil
func (o *ServiceLimitUsageDto) UnsetPeriodEnd() {
	o.PeriodEnd.Unset()
}

// GetAmountConsumed returns the AmountConsumed field value if set, zero value otherwise.
func (o *ServiceLimitUsageDto) GetAmountConsumed() float64 {
	if o == nil || IsNil(o.AmountConsumed) {
		var ret float64
		return ret
	}
	return *o.AmountConsumed
}

// GetAmountConsumedOk returns a tuple with the AmountConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLimitUsageDto) GetAmountConsumedOk() (*float64, bool) {
	if o == nil || IsNil(o.AmountConsumed) {
		return nil, false
	}
	return o.AmountConsumed, true
}

// HasAmountConsumed returns a boolean if a field has been set.
func (o *ServiceLimitUsageDto) IsAmountConsumedSet() bool {
	if o != nil && !IsNil(o.AmountConsumed) {
		return true
	}

	return false
}

// SetAmountConsumed gets a reference to the given float64 and assigns it to the AmountConsumed field.
func (o *ServiceLimitUsageDto) SetAmountConsumed(v float64) {
	o.AmountConsumed = &v
}

// GetAmountAvailable returns the AmountAvailable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceLimitUsageDto) GetAmountAvailable() float64 {
	if o == nil || IsNil(o.AmountAvailable.Get()) {
		var ret float64
		return ret
	}
	return *o.AmountAvailable.Get()
}

// GetAmountAvailableOk returns a tuple with the AmountAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceLimitUsageDto) GetAmountAvailableOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountAvailable.Get(), o.AmountAvailable.IsSet()
}

// HasAmountAvailable returns a boolean if a field has been set.
func (o *ServiceLimitUsageDto) IsAmountAvailableSet() bool {
	if o != nil && o.AmountAvailable.IsSet() {
		return true
	}

	return false
}

// SetAmountAvailable gets a reference to the given NullableFloat64 and assigns it to the AmountAvailable field.
func (o *ServiceLimitUsageDto) SetAmountAvailable(v float64) {
	o.AmountAvailable.Set(&v)
}
// SetAmountAvailableNil sets the value for AmountAvailable to be an explicit nil
func (o *ServiceLimitUsageDto) SetAmountAvailableNil() {
	o.AmountAvailable.Set(nil)
}

// UnsetAmountAvailable ensures that no value is present for AmountAvailable, not even an explicit nil
func (o *ServiceLimitUsageDto) UnsetAmountAvailable() {
	o.AmountAvailable.Unset()
}

// GetQuantityConsumed returns the QuantityConsumed field value if set, zero value otherwise.
func (o *ServiceLimitUsageDto) GetQuantityConsumed() int32 {
	if o == nil || IsNil(o.QuantityConsumed) {
		var ret int32
		return ret
	}
	return *o.QuantityConsumed
}

// GetQuantityConsumedOk returns a tuple with the QuantityConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLimitUsageDto) GetQuantityConsumedOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityConsumed) {
		return nil, false
	}
	return o.QuantityConsumed, true
}

// HasQuantityConsumed returns a boolean if a field has been set.
func (o *ServiceLimitUsageDto) IsQuantityConsumedSet() bool {
	if o != nil && !IsNil(o.QuantityConsumed) {
		return true
	}

	return false
}

// SetQuantityConsumed gets a reference to the given int32 and assigns it to the QuantityConsumed field.
func (o *ServiceLimitUsageDto) SetQuantityConsumed(v int32) {
	o.QuantityConsumed = &v
}

// GetQuantityAvailable returns the QuantityAvailable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceLimitUsageDto) GetQuantityAvailable() int32 {
	if o == nil || IsNil(o.QuantityAvailable.Get()) {
		var ret int32
		return ret
	}
	return *o.QuantityAvailable.Get()
}

// GetQuantityAvailableOk returns a tuple with the QuantityAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceLimitUsageDto) GetQuantityAvailableOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.QuantityAvailable.Get(), o.QuantityAvailable.IsSet()
}

// HasQuantityAvailable returns a boolean if a field has been set.
func (o *ServiceLimitUsageDto) IsQuantityAvailableSet() bool {
	if o != nil && o.QuantityAvailable.IsSet() {
		return true
	}

	return false
}

// SetQuantityAvailable gets a reference to the given NullableInt32 and assigns it to the QuantityAvailable field.
func (o *ServiceLimitUsageDto) SetQuantityAvailable(v int32) {
	o.QuantityAvailable.Set(&v)
}
// SetQuantityAvailableNil sets the value for QuantityAvailable to be an explicit nil
func (o *ServiceLimitUsageDto) SetQuantityAvailableNil() {
	o.QuantityAvailable.Set(nil)
}

// UnsetQuantityAvailable ensures that no value is present for QuantityAvailable, not even an explicit nil
func (o *ServiceLimitUsageDto) UnsetQuantityAvailable() {
	o.QuantityAvailable.Unset()
}

func (o ServiceLimitUsageDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceLimitUsageDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PeriodStart) {
		toSerialize["periodStart"] = o.PeriodStart
	}
	if o.PeriodEnd.IsSet() {
		toSerialize["periodEnd"] = o.PeriodEnd.Get()
	}
	if !IsNil(o.AmountConsumed) {
		toSerialize["amountConsumed"] = o.AmountConsumed
	}
	if o.AmountAvailable.IsSet() {
		toSerialize["amountAvailable"] = o.AmountAvailable.Get()
	}
	if !IsNil(o.QuantityConsumed) {
		toSerialize["quantityConsumed"] = o.QuantityConsumed
	}
	if o.QuantityAvailable.IsSet() {
		toSerialize["quantityAvailable"] = o.QuantityAvailable.Get()
	}
	return toSerialize, nil
}

type NullableServiceLimitUsageDto struct {
	value *ServiceLimitUsageDto
	isSet bool
}

func (v NullableServiceLimitUsageDto) Get() *ServiceLimitUsageDto {
	return v.value
}

func (v *NullableServiceLimitUsageDto) Set(val *ServiceLimitUsageDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceLimitUsageDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceLimitUsageDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceLimitUsageDto(val *ServiceLimitUsageDto) *NullableServiceLimitUsageDto {
	return &NullableServiceLimitUsageDto{value: val, isSet: true}
}

func (v NullableServiceLimitUsageDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceLimitUsageDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

