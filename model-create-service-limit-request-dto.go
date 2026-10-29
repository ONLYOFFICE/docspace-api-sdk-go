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

// checks if the CreateServiceLimitRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateServiceLimitRequestDto{}

// CreateServiceLimitRequestDto The spending limit to create on the AI services: for the portal as a whole, for every user, or for one user.
type CreateServiceLimitRequestDto struct {
	// Whom the limit applies to.
	Scope *ServiceLimitScope `json:"scope,omitempty"`
	// The ID of the portal user a personal limit is set for, as returned in `id` by the people operations.
	UserId NullableString `json:"userId,omitempty"`
	// The most money the AI services may cost per period, in their currency, which cannot be chosen. Zero blocks the  AI services completely.
	AmountValue NullableFloat64 `json:"amountValue,omitempty"`
	// The largest quantity the AI services may consume per period, in their common unit. Zero blocks the AI services  completely.
	QuantityValue NullableInt32 `json:"quantityValue,omitempty"`
	// The window the thresholds apply to, in UTC. A day, week or month is calendar-aligned rather than rolling - a  daily limit resets at midnight UTC, a weekly one on Monday - while a lifetime limit counts from its creation and  never resets.
	Period *ServiceLimitPeriod `json:"period,omitempty"`
}

// NewCreateServiceLimitRequestDto instantiates a new CreateServiceLimitRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServiceLimitRequestDto() *CreateServiceLimitRequestDto {
	this := CreateServiceLimitRequestDto{}
	return &this
}

// NewCreateServiceLimitRequestDtoWithDefaults instantiates a new CreateServiceLimitRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServiceLimitRequestDtoWithDefaults() *CreateServiceLimitRequestDto {
	this := CreateServiceLimitRequestDto{}
	return &this
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *CreateServiceLimitRequestDto) GetScope() ServiceLimitScope {
	if o == nil || IsNil(o.Scope) {
		var ret ServiceLimitScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServiceLimitRequestDto) GetScopeOk() (*ServiceLimitScope, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *CreateServiceLimitRequestDto) IsScopeSet() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given ServiceLimitScope and assigns it to the Scope field.
func (o *CreateServiceLimitRequestDto) SetScope(v ServiceLimitScope) {
	o.Scope = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateServiceLimitRequestDto) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateServiceLimitRequestDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *CreateServiceLimitRequestDto) IsUserIdSet() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *CreateServiceLimitRequestDto) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *CreateServiceLimitRequestDto) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *CreateServiceLimitRequestDto) UnsetUserId() {
	o.UserId.Unset()
}

// GetAmountValue returns the AmountValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateServiceLimitRequestDto) GetAmountValue() float64 {
	if o == nil || IsNil(o.AmountValue.Get()) {
		var ret float64
		return ret
	}
	return *o.AmountValue.Get()
}

// GetAmountValueOk returns a tuple with the AmountValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateServiceLimitRequestDto) GetAmountValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountValue.Get(), o.AmountValue.IsSet()
}

// HasAmountValue returns a boolean if a field has been set.
func (o *CreateServiceLimitRequestDto) IsAmountValueSet() bool {
	if o != nil && o.AmountValue.IsSet() {
		return true
	}

	return false
}

// SetAmountValue gets a reference to the given NullableFloat64 and assigns it to the AmountValue field.
func (o *CreateServiceLimitRequestDto) SetAmountValue(v float64) {
	o.AmountValue.Set(&v)
}
// SetAmountValueNil sets the value for AmountValue to be an explicit nil
func (o *CreateServiceLimitRequestDto) SetAmountValueNil() {
	o.AmountValue.Set(nil)
}

// UnsetAmountValue ensures that no value is present for AmountValue, not even an explicit nil
func (o *CreateServiceLimitRequestDto) UnsetAmountValue() {
	o.AmountValue.Unset()
}

// GetQuantityValue returns the QuantityValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateServiceLimitRequestDto) GetQuantityValue() int32 {
	if o == nil || IsNil(o.QuantityValue.Get()) {
		var ret int32
		return ret
	}
	return *o.QuantityValue.Get()
}

// GetQuantityValueOk returns a tuple with the QuantityValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateServiceLimitRequestDto) GetQuantityValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.QuantityValue.Get(), o.QuantityValue.IsSet()
}

// HasQuantityValue returns a boolean if a field has been set.
func (o *CreateServiceLimitRequestDto) IsQuantityValueSet() bool {
	if o != nil && o.QuantityValue.IsSet() {
		return true
	}

	return false
}

// SetQuantityValue gets a reference to the given NullableInt32 and assigns it to the QuantityValue field.
func (o *CreateServiceLimitRequestDto) SetQuantityValue(v int32) {
	o.QuantityValue.Set(&v)
}
// SetQuantityValueNil sets the value for QuantityValue to be an explicit nil
func (o *CreateServiceLimitRequestDto) SetQuantityValueNil() {
	o.QuantityValue.Set(nil)
}

// UnsetQuantityValue ensures that no value is present for QuantityValue, not even an explicit nil
func (o *CreateServiceLimitRequestDto) UnsetQuantityValue() {
	o.QuantityValue.Unset()
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *CreateServiceLimitRequestDto) GetPeriod() ServiceLimitPeriod {
	if o == nil || IsNil(o.Period) {
		var ret ServiceLimitPeriod
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServiceLimitRequestDto) GetPeriodOk() (*ServiceLimitPeriod, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *CreateServiceLimitRequestDto) IsPeriodSet() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given ServiceLimitPeriod and assigns it to the Period field.
func (o *CreateServiceLimitRequestDto) SetPeriod(v ServiceLimitPeriod) {
	o.Period = &v
}

func (o CreateServiceLimitRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateServiceLimitRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
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

type NullableCreateServiceLimitRequestDto struct {
	value *CreateServiceLimitRequestDto
	isSet bool
}

func (v NullableCreateServiceLimitRequestDto) Get() *CreateServiceLimitRequestDto {
	return v.value
}

func (v *NullableCreateServiceLimitRequestDto) Set(val *CreateServiceLimitRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServiceLimitRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServiceLimitRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServiceLimitRequestDto(val *CreateServiceLimitRequestDto) *NullableCreateServiceLimitRequestDto {
	return &NullableCreateServiceLimitRequestDto{value: val, isSet: true}
}

func (v NullableCreateServiceLimitRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServiceLimitRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

