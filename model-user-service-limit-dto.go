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

// checks if the UserServiceLimitDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserServiceLimitDto{}

// UserServiceLimitDto The spending limit on the AI services in force for one user: the user's personal limit, or else the limit for every  user, together with the money the user spent against it in the current period.
type UserServiceLimitDto struct {
	// The service limit ID, used to read, change, switch off or delete the limit through the  `api/2.0/portal/payment/servicelimit` operations. For the limit for every user it is the same for all users.
	Id *int32 `json:"id,omitempty"`
	// Whether the limit is the user's own (`Custom`) or the limit for every user (`Global`) the user falls back to.
	Scope *ServiceLimitScope `json:"scope,omitempty"`
	// The most money the AI services may cost the user per period, in `currency`. It is the only threshold a limit on  the AI services has: their services count in different units, so a threshold on a quantity is refused for them.
	AmountValue NullableFloat64 `json:"amountValue,omitempty"`
	// The money the AI services cost the user in the current period, in `currency`; 0 when nothing was spent yet,  including a user who has never used the AI services.
	AmountConsumed *float64 `json:"amountConsumed,omitempty"`
	// The currency `amountValue` and `amountConsumed` are expressed in; null when the limit has no money threshold.
	Currency NullableString `json:"currency,omitempty"`
	// The window the threshold applies to. It resets at the start of the next day, week or month in UTC (a week on  Monday), not a fixed time after the first charge; a lifetime limit never resets.
	Period *ServiceLimitPeriod `json:"period,omitempty"`
	// The moment the current period ends and the spending starts over from zero, in UTC. Null for a lifetime limit,  which never resets.
	PeriodEnd NullableTime `json:"periodEnd,omitempty"`
}

// NewUserServiceLimitDto instantiates a new UserServiceLimitDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserServiceLimitDto() *UserServiceLimitDto {
	this := UserServiceLimitDto{}
	return &this
}

// NewUserServiceLimitDtoWithDefaults instantiates a new UserServiceLimitDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserServiceLimitDtoWithDefaults() *UserServiceLimitDto {
	this := UserServiceLimitDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserServiceLimitDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserServiceLimitDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserServiceLimitDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UserServiceLimitDto) SetId(v int32) {
	o.Id = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *UserServiceLimitDto) GetScope() ServiceLimitScope {
	if o == nil || IsNil(o.Scope) {
		var ret ServiceLimitScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserServiceLimitDto) GetScopeOk() (*ServiceLimitScope, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *UserServiceLimitDto) IsScopeSet() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given ServiceLimitScope and assigns it to the Scope field.
func (o *UserServiceLimitDto) SetScope(v ServiceLimitScope) {
	o.Scope = &v
}

// GetAmountValue returns the AmountValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserServiceLimitDto) GetAmountValue() float64 {
	if o == nil || IsNil(o.AmountValue.Get()) {
		var ret float64
		return ret
	}
	return *o.AmountValue.Get()
}

// GetAmountValueOk returns a tuple with the AmountValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserServiceLimitDto) GetAmountValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountValue.Get(), o.AmountValue.IsSet()
}

// HasAmountValue returns a boolean if a field has been set.
func (o *UserServiceLimitDto) IsAmountValueSet() bool {
	if o != nil && o.AmountValue.IsSet() {
		return true
	}

	return false
}

// SetAmountValue gets a reference to the given NullableFloat64 and assigns it to the AmountValue field.
func (o *UserServiceLimitDto) SetAmountValue(v float64) {
	o.AmountValue.Set(&v)
}
// SetAmountValueNil sets the value for AmountValue to be an explicit nil
func (o *UserServiceLimitDto) SetAmountValueNil() {
	o.AmountValue.Set(nil)
}

// UnsetAmountValue ensures that no value is present for AmountValue, not even an explicit nil
func (o *UserServiceLimitDto) UnsetAmountValue() {
	o.AmountValue.Unset()
}

// GetAmountConsumed returns the AmountConsumed field value if set, zero value otherwise.
func (o *UserServiceLimitDto) GetAmountConsumed() float64 {
	if o == nil || IsNil(o.AmountConsumed) {
		var ret float64
		return ret
	}
	return *o.AmountConsumed
}

// GetAmountConsumedOk returns a tuple with the AmountConsumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserServiceLimitDto) GetAmountConsumedOk() (*float64, bool) {
	if o == nil || IsNil(o.AmountConsumed) {
		return nil, false
	}
	return o.AmountConsumed, true
}

// HasAmountConsumed returns a boolean if a field has been set.
func (o *UserServiceLimitDto) IsAmountConsumedSet() bool {
	if o != nil && !IsNil(o.AmountConsumed) {
		return true
	}

	return false
}

// SetAmountConsumed gets a reference to the given float64 and assigns it to the AmountConsumed field.
func (o *UserServiceLimitDto) SetAmountConsumed(v float64) {
	o.AmountConsumed = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserServiceLimitDto) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserServiceLimitDto) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *UserServiceLimitDto) IsCurrencySet() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *UserServiceLimitDto) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *UserServiceLimitDto) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *UserServiceLimitDto) UnsetCurrency() {
	o.Currency.Unset()
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *UserServiceLimitDto) GetPeriod() ServiceLimitPeriod {
	if o == nil || IsNil(o.Period) {
		var ret ServiceLimitPeriod
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserServiceLimitDto) GetPeriodOk() (*ServiceLimitPeriod, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *UserServiceLimitDto) IsPeriodSet() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given ServiceLimitPeriod and assigns it to the Period field.
func (o *UserServiceLimitDto) SetPeriod(v ServiceLimitPeriod) {
	o.Period = &v
}

// GetPeriodEnd returns the PeriodEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserServiceLimitDto) GetPeriodEnd() time.Time {
	if o == nil || IsNil(o.PeriodEnd.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PeriodEnd.Get()
}

// GetPeriodEndOk returns a tuple with the PeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserServiceLimitDto) GetPeriodEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PeriodEnd.Get(), o.PeriodEnd.IsSet()
}

// HasPeriodEnd returns a boolean if a field has been set.
func (o *UserServiceLimitDto) IsPeriodEndSet() bool {
	if o != nil && o.PeriodEnd.IsSet() {
		return true
	}

	return false
}

// SetPeriodEnd gets a reference to the given NullableTime and assigns it to the PeriodEnd field.
func (o *UserServiceLimitDto) SetPeriodEnd(v time.Time) {
	o.PeriodEnd.Set(&v)
}
// SetPeriodEndNil sets the value for PeriodEnd to be an explicit nil
func (o *UserServiceLimitDto) SetPeriodEndNil() {
	o.PeriodEnd.Set(nil)
}

// UnsetPeriodEnd ensures that no value is present for PeriodEnd, not even an explicit nil
func (o *UserServiceLimitDto) UnsetPeriodEnd() {
	o.PeriodEnd.Unset()
}

func (o UserServiceLimitDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserServiceLimitDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if o.AmountValue.IsSet() {
		toSerialize["amountValue"] = o.AmountValue.Get()
	}
	if !IsNil(o.AmountConsumed) {
		toSerialize["amountConsumed"] = o.AmountConsumed
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if o.PeriodEnd.IsSet() {
		toSerialize["periodEnd"] = o.PeriodEnd.Get()
	}
	return toSerialize, nil
}

type NullableUserServiceLimitDto struct {
	value *UserServiceLimitDto
	isSet bool
}

func (v NullableUserServiceLimitDto) Get() *UserServiceLimitDto {
	return v.value
}

func (v *NullableUserServiceLimitDto) Set(val *UserServiceLimitDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUserServiceLimitDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUserServiceLimitDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserServiceLimitDto(val *UserServiceLimitDto) *NullableUserServiceLimitDto {
	return &NullableUserServiceLimitDto{value: val, isSet: true}
}

func (v NullableUserServiceLimitDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserServiceLimitDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

