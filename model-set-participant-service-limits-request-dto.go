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

// checks if the SetParticipantServiceLimitsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetParticipantServiceLimitsRequestDto{}

// SetParticipantServiceLimitsRequestDto The personal spending limit on the AI services to set for each of several portal users.
type SetParticipantServiceLimitsRequestDto struct {
	// The IDs of the portal users the limit is set for, as returned in `id` by the people operations.
	UserIds []string `json:"userIds"`
	// The most money the AI services may cost each user per period, in their currency, which cannot be chosen. Zero  blocks the AI services completely.
	AmountValue NullableFloat64 `json:"amountValue,omitempty"`
	// The largest quantity the AI services may consume per period for each user, in their common unit. Zero blocks the  AI services completely.
	QuantityValue NullableInt32 `json:"quantityValue,omitempty"`
	// The window the thresholds apply to, in UTC. A day, week or month is calendar-aligned rather than rolling - a  daily limit resets at midnight UTC, a weekly one on Monday - while a lifetime limit counts from its creation and  never resets.
	Period *ServiceLimitPeriod `json:"period,omitempty"`
}

type _SetParticipantServiceLimitsRequestDto SetParticipantServiceLimitsRequestDto

// NewSetParticipantServiceLimitsRequestDto instantiates a new SetParticipantServiceLimitsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetParticipantServiceLimitsRequestDto(userIds []string) *SetParticipantServiceLimitsRequestDto {
	this := SetParticipantServiceLimitsRequestDto{}
	this.UserIds = userIds
	return &this
}

// NewSetParticipantServiceLimitsRequestDtoWithDefaults instantiates a new SetParticipantServiceLimitsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetParticipantServiceLimitsRequestDtoWithDefaults() *SetParticipantServiceLimitsRequestDto {
	this := SetParticipantServiceLimitsRequestDto{}
	return &this
}

// GetUserIds returns the UserIds field value
func (o *SetParticipantServiceLimitsRequestDto) GetUserIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value
// and a boolean to check if the value has been set.
func (o *SetParticipantServiceLimitsRequestDto) GetUserIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserIds, true
}

// SetUserIds sets field value
func (o *SetParticipantServiceLimitsRequestDto) SetUserIds(v []string) {
	o.UserIds = v
}

// GetAmountValue returns the AmountValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetParticipantServiceLimitsRequestDto) GetAmountValue() float64 {
	if o == nil || IsNil(o.AmountValue.Get()) {
		var ret float64
		return ret
	}
	return *o.AmountValue.Get()
}

// GetAmountValueOk returns a tuple with the AmountValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetParticipantServiceLimitsRequestDto) GetAmountValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountValue.Get(), o.AmountValue.IsSet()
}

// HasAmountValue returns a boolean if a field has been set.
func (o *SetParticipantServiceLimitsRequestDto) IsAmountValueSet() bool {
	if o != nil && o.AmountValue.IsSet() {
		return true
	}

	return false
}

// SetAmountValue gets a reference to the given NullableFloat64 and assigns it to the AmountValue field.
func (o *SetParticipantServiceLimitsRequestDto) SetAmountValue(v float64) {
	o.AmountValue.Set(&v)
}
// SetAmountValueNil sets the value for AmountValue to be an explicit nil
func (o *SetParticipantServiceLimitsRequestDto) SetAmountValueNil() {
	o.AmountValue.Set(nil)
}

// UnsetAmountValue ensures that no value is present for AmountValue, not even an explicit nil
func (o *SetParticipantServiceLimitsRequestDto) UnsetAmountValue() {
	o.AmountValue.Unset()
}

// GetQuantityValue returns the QuantityValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetParticipantServiceLimitsRequestDto) GetQuantityValue() int32 {
	if o == nil || IsNil(o.QuantityValue.Get()) {
		var ret int32
		return ret
	}
	return *o.QuantityValue.Get()
}

// GetQuantityValueOk returns a tuple with the QuantityValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetParticipantServiceLimitsRequestDto) GetQuantityValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.QuantityValue.Get(), o.QuantityValue.IsSet()
}

// HasQuantityValue returns a boolean if a field has been set.
func (o *SetParticipantServiceLimitsRequestDto) IsQuantityValueSet() bool {
	if o != nil && o.QuantityValue.IsSet() {
		return true
	}

	return false
}

// SetQuantityValue gets a reference to the given NullableInt32 and assigns it to the QuantityValue field.
func (o *SetParticipantServiceLimitsRequestDto) SetQuantityValue(v int32) {
	o.QuantityValue.Set(&v)
}
// SetQuantityValueNil sets the value for QuantityValue to be an explicit nil
func (o *SetParticipantServiceLimitsRequestDto) SetQuantityValueNil() {
	o.QuantityValue.Set(nil)
}

// UnsetQuantityValue ensures that no value is present for QuantityValue, not even an explicit nil
func (o *SetParticipantServiceLimitsRequestDto) UnsetQuantityValue() {
	o.QuantityValue.Unset()
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *SetParticipantServiceLimitsRequestDto) GetPeriod() ServiceLimitPeriod {
	if o == nil || IsNil(o.Period) {
		var ret ServiceLimitPeriod
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetParticipantServiceLimitsRequestDto) GetPeriodOk() (*ServiceLimitPeriod, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *SetParticipantServiceLimitsRequestDto) IsPeriodSet() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given ServiceLimitPeriod and assigns it to the Period field.
func (o *SetParticipantServiceLimitsRequestDto) SetPeriod(v ServiceLimitPeriod) {
	o.Period = &v
}

func (o SetParticipantServiceLimitsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetParticipantServiceLimitsRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userIds"] = o.UserIds
	if o.AmountValue.IsSet() {
		toSerialize["amountValue"] = o.AmountValue.Get()
	}
	if o.QuantityValue.IsSet() {
		toSerialize["quantityValue"] = o.QuantityValue.Get()
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	return toSerialize, nil
}

func (o *SetParticipantServiceLimitsRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userIds",
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

	varSetParticipantServiceLimitsRequestDto := _SetParticipantServiceLimitsRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSetParticipantServiceLimitsRequestDto)

	if err != nil {
		return err
	}

	*o = SetParticipantServiceLimitsRequestDto(varSetParticipantServiceLimitsRequestDto)

	return err
}

type NullableSetParticipantServiceLimitsRequestDto struct {
	value *SetParticipantServiceLimitsRequestDto
	isSet bool
}

func (v NullableSetParticipantServiceLimitsRequestDto) Get() *SetParticipantServiceLimitsRequestDto {
	return v.value
}

func (v *NullableSetParticipantServiceLimitsRequestDto) Set(val *SetParticipantServiceLimitsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSetParticipantServiceLimitsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSetParticipantServiceLimitsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetParticipantServiceLimitsRequestDto(val *SetParticipantServiceLimitsRequestDto) *NullableSetParticipantServiceLimitsRequestDto {
	return &NullableSetParticipantServiceLimitsRequestDto{value: val, isSet: true}
}

func (v NullableSetParticipantServiceLimitsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetParticipantServiceLimitsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

