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

// checks if the SubscriptionBalanceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionBalanceInfo{}

// SubscriptionBalanceInfo The information about the current subscription and its unused balance.
type SubscriptionBalanceInfo struct {
	// The total cost of the current billing period (the sum across all subscription items).
	TotalCost *float64 `json:"totalCost,omitempty"`
	// The three-character ISO 4217 currency symbol of the subscription.
	Currency NullableString `json:"currency,omitempty"`
	// The start of the current billing period.
	PeriodStart *time.Time `json:"periodStart,omitempty"`
	// The end of the current billing period.
	PeriodEnd *time.Time `json:"periodEnd,omitempty"`
	// The boundary of the used part of the period (the moment of the request).
	PeriodUsedUntil *time.Time `json:"periodUsedUntil,omitempty"`
	// The number of days elapsed since the start of the period (inclusive).
	DaysElapsed *int32 `json:"daysElapsed,omitempty"`
	// The unused balance of the subscription, in the subscription currency.
	RemainingBalance *float64 `json:"remainingBalance,omitempty"`
	// The unused balance of the subscription, converted to the wallet currency.
	RemainingBalanceInWalletCurrency *float64 `json:"remainingBalanceInWalletCurrency,omitempty"`
	// The three-character ISO 4217 currency symbol of the wallet.
	WalletCurrency NullableString `json:"walletCurrency,omitempty"`
}

// NewSubscriptionBalanceInfo instantiates a new SubscriptionBalanceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionBalanceInfo() *SubscriptionBalanceInfo {
	this := SubscriptionBalanceInfo{}
	return &this
}

// NewSubscriptionBalanceInfoWithDefaults instantiates a new SubscriptionBalanceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionBalanceInfoWithDefaults() *SubscriptionBalanceInfo {
	this := SubscriptionBalanceInfo{}
	return &this
}

// GetTotalCost returns the TotalCost field value if set, zero value otherwise.
func (o *SubscriptionBalanceInfo) GetTotalCost() float64 {
	if o == nil || IsNil(o.TotalCost) {
		var ret float64
		return ret
	}
	return *o.TotalCost
}

// GetTotalCostOk returns a tuple with the TotalCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionBalanceInfo) GetTotalCostOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalCost) {
		return nil, false
	}
	return o.TotalCost, true
}

// HasTotalCost returns a boolean if a field has been set.
func (o *SubscriptionBalanceInfo) IsTotalCostSet() bool {
	if o != nil && !IsNil(o.TotalCost) {
		return true
	}

	return false
}

// SetTotalCost gets a reference to the given float64 and assigns it to the TotalCost field.
func (o *SubscriptionBalanceInfo) SetTotalCost(v float64) {
	o.TotalCost = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionBalanceInfo) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionBalanceInfo) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *SubscriptionBalanceInfo) IsCurrencySet() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *SubscriptionBalanceInfo) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *SubscriptionBalanceInfo) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *SubscriptionBalanceInfo) UnsetCurrency() {
	o.Currency.Unset()
}

// GetPeriodStart returns the PeriodStart field value if set, zero value otherwise.
func (o *SubscriptionBalanceInfo) GetPeriodStart() time.Time {
	if o == nil || IsNil(o.PeriodStart) {
		var ret time.Time
		return ret
	}
	return *o.PeriodStart
}

// GetPeriodStartOk returns a tuple with the PeriodStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionBalanceInfo) GetPeriodStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PeriodStart) {
		return nil, false
	}
	return o.PeriodStart, true
}

// HasPeriodStart returns a boolean if a field has been set.
func (o *SubscriptionBalanceInfo) IsPeriodStartSet() bool {
	if o != nil && !IsNil(o.PeriodStart) {
		return true
	}

	return false
}

// SetPeriodStart gets a reference to the given time.Time and assigns it to the PeriodStart field.
func (o *SubscriptionBalanceInfo) SetPeriodStart(v time.Time) {
	o.PeriodStart = &v
}

// GetPeriodEnd returns the PeriodEnd field value if set, zero value otherwise.
func (o *SubscriptionBalanceInfo) GetPeriodEnd() time.Time {
	if o == nil || IsNil(o.PeriodEnd) {
		var ret time.Time
		return ret
	}
	return *o.PeriodEnd
}

// GetPeriodEndOk returns a tuple with the PeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionBalanceInfo) GetPeriodEndOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PeriodEnd) {
		return nil, false
	}
	return o.PeriodEnd, true
}

// HasPeriodEnd returns a boolean if a field has been set.
func (o *SubscriptionBalanceInfo) IsPeriodEndSet() bool {
	if o != nil && !IsNil(o.PeriodEnd) {
		return true
	}

	return false
}

// SetPeriodEnd gets a reference to the given time.Time and assigns it to the PeriodEnd field.
func (o *SubscriptionBalanceInfo) SetPeriodEnd(v time.Time) {
	o.PeriodEnd = &v
}

// GetPeriodUsedUntil returns the PeriodUsedUntil field value if set, zero value otherwise.
func (o *SubscriptionBalanceInfo) GetPeriodUsedUntil() time.Time {
	if o == nil || IsNil(o.PeriodUsedUntil) {
		var ret time.Time
		return ret
	}
	return *o.PeriodUsedUntil
}

// GetPeriodUsedUntilOk returns a tuple with the PeriodUsedUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionBalanceInfo) GetPeriodUsedUntilOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PeriodUsedUntil) {
		return nil, false
	}
	return o.PeriodUsedUntil, true
}

// HasPeriodUsedUntil returns a boolean if a field has been set.
func (o *SubscriptionBalanceInfo) IsPeriodUsedUntilSet() bool {
	if o != nil && !IsNil(o.PeriodUsedUntil) {
		return true
	}

	return false
}

// SetPeriodUsedUntil gets a reference to the given time.Time and assigns it to the PeriodUsedUntil field.
func (o *SubscriptionBalanceInfo) SetPeriodUsedUntil(v time.Time) {
	o.PeriodUsedUntil = &v
}

// GetDaysElapsed returns the DaysElapsed field value if set, zero value otherwise.
func (o *SubscriptionBalanceInfo) GetDaysElapsed() int32 {
	if o == nil || IsNil(o.DaysElapsed) {
		var ret int32
		return ret
	}
	return *o.DaysElapsed
}

// GetDaysElapsedOk returns a tuple with the DaysElapsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionBalanceInfo) GetDaysElapsedOk() (*int32, bool) {
	if o == nil || IsNil(o.DaysElapsed) {
		return nil, false
	}
	return o.DaysElapsed, true
}

// HasDaysElapsed returns a boolean if a field has been set.
func (o *SubscriptionBalanceInfo) IsDaysElapsedSet() bool {
	if o != nil && !IsNil(o.DaysElapsed) {
		return true
	}

	return false
}

// SetDaysElapsed gets a reference to the given int32 and assigns it to the DaysElapsed field.
func (o *SubscriptionBalanceInfo) SetDaysElapsed(v int32) {
	o.DaysElapsed = &v
}

// GetRemainingBalance returns the RemainingBalance field value if set, zero value otherwise.
func (o *SubscriptionBalanceInfo) GetRemainingBalance() float64 {
	if o == nil || IsNil(o.RemainingBalance) {
		var ret float64
		return ret
	}
	return *o.RemainingBalance
}

// GetRemainingBalanceOk returns a tuple with the RemainingBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionBalanceInfo) GetRemainingBalanceOk() (*float64, bool) {
	if o == nil || IsNil(o.RemainingBalance) {
		return nil, false
	}
	return o.RemainingBalance, true
}

// HasRemainingBalance returns a boolean if a field has been set.
func (o *SubscriptionBalanceInfo) IsRemainingBalanceSet() bool {
	if o != nil && !IsNil(o.RemainingBalance) {
		return true
	}

	return false
}

// SetRemainingBalance gets a reference to the given float64 and assigns it to the RemainingBalance field.
func (o *SubscriptionBalanceInfo) SetRemainingBalance(v float64) {
	o.RemainingBalance = &v
}

// GetRemainingBalanceInWalletCurrency returns the RemainingBalanceInWalletCurrency field value if set, zero value otherwise.
func (o *SubscriptionBalanceInfo) GetRemainingBalanceInWalletCurrency() float64 {
	if o == nil || IsNil(o.RemainingBalanceInWalletCurrency) {
		var ret float64
		return ret
	}
	return *o.RemainingBalanceInWalletCurrency
}

// GetRemainingBalanceInWalletCurrencyOk returns a tuple with the RemainingBalanceInWalletCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionBalanceInfo) GetRemainingBalanceInWalletCurrencyOk() (*float64, bool) {
	if o == nil || IsNil(o.RemainingBalanceInWalletCurrency) {
		return nil, false
	}
	return o.RemainingBalanceInWalletCurrency, true
}

// HasRemainingBalanceInWalletCurrency returns a boolean if a field has been set.
func (o *SubscriptionBalanceInfo) IsRemainingBalanceInWalletCurrencySet() bool {
	if o != nil && !IsNil(o.RemainingBalanceInWalletCurrency) {
		return true
	}

	return false
}

// SetRemainingBalanceInWalletCurrency gets a reference to the given float64 and assigns it to the RemainingBalanceInWalletCurrency field.
func (o *SubscriptionBalanceInfo) SetRemainingBalanceInWalletCurrency(v float64) {
	o.RemainingBalanceInWalletCurrency = &v
}

// GetWalletCurrency returns the WalletCurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionBalanceInfo) GetWalletCurrency() string {
	if o == nil || IsNil(o.WalletCurrency.Get()) {
		var ret string
		return ret
	}
	return *o.WalletCurrency.Get()
}

// GetWalletCurrencyOk returns a tuple with the WalletCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionBalanceInfo) GetWalletCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WalletCurrency.Get(), o.WalletCurrency.IsSet()
}

// HasWalletCurrency returns a boolean if a field has been set.
func (o *SubscriptionBalanceInfo) IsWalletCurrencySet() bool {
	if o != nil && o.WalletCurrency.IsSet() {
		return true
	}

	return false
}

// SetWalletCurrency gets a reference to the given NullableString and assigns it to the WalletCurrency field.
func (o *SubscriptionBalanceInfo) SetWalletCurrency(v string) {
	o.WalletCurrency.Set(&v)
}
// SetWalletCurrencyNil sets the value for WalletCurrency to be an explicit nil
func (o *SubscriptionBalanceInfo) SetWalletCurrencyNil() {
	o.WalletCurrency.Set(nil)
}

// UnsetWalletCurrency ensures that no value is present for WalletCurrency, not even an explicit nil
func (o *SubscriptionBalanceInfo) UnsetWalletCurrency() {
	o.WalletCurrency.Unset()
}

func (o SubscriptionBalanceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionBalanceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCost) {
		toSerialize["totalCost"] = o.TotalCost
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if !IsNil(o.PeriodStart) {
		toSerialize["periodStart"] = o.PeriodStart
	}
	if !IsNil(o.PeriodEnd) {
		toSerialize["periodEnd"] = o.PeriodEnd
	}
	if !IsNil(o.PeriodUsedUntil) {
		toSerialize["periodUsedUntil"] = o.PeriodUsedUntil
	}
	if !IsNil(o.DaysElapsed) {
		toSerialize["daysElapsed"] = o.DaysElapsed
	}
	if !IsNil(o.RemainingBalance) {
		toSerialize["remainingBalance"] = o.RemainingBalance
	}
	if !IsNil(o.RemainingBalanceInWalletCurrency) {
		toSerialize["remainingBalanceInWalletCurrency"] = o.RemainingBalanceInWalletCurrency
	}
	if o.WalletCurrency.IsSet() {
		toSerialize["walletCurrency"] = o.WalletCurrency.Get()
	}
	return toSerialize, nil
}

type NullableSubscriptionBalanceInfo struct {
	value *SubscriptionBalanceInfo
	isSet bool
}

func (v NullableSubscriptionBalanceInfo) Get() *SubscriptionBalanceInfo {
	return v.value
}

func (v *NullableSubscriptionBalanceInfo) Set(val *SubscriptionBalanceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionBalanceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionBalanceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionBalanceInfo(val *SubscriptionBalanceInfo) *NullableSubscriptionBalanceInfo {
	return &NullableSubscriptionBalanceInfo{value: val, isSet: true}
}

func (v NullableSubscriptionBalanceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionBalanceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

