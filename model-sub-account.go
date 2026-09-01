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

// checks if the SubAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubAccount{}

// SubAccount Represents a sub-account with a specific currency and balance.
type SubAccount struct {
	// The three-character ISO 4217 currency symbol.
	Currency NullableString `json:"currency,omitempty"`
	// The amount in the specified currency.
	Amount *float64 `json:"amount,omitempty"`
}

// NewSubAccount instantiates a new SubAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubAccount() *SubAccount {
	this := SubAccount{}
	return &this
}

// NewSubAccountWithDefaults instantiates a new SubAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubAccountWithDefaults() *SubAccount {
	this := SubAccount{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubAccount) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubAccount) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *SubAccount) IsCurrencySet() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *SubAccount) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *SubAccount) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *SubAccount) UnsetCurrency() {
	o.Currency.Unset()
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *SubAccount) GetAmount() float64 {
	if o == nil || IsNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccount) GetAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *SubAccount) IsAmountSet() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *SubAccount) SetAmount(v float64) {
	o.Amount = &v
}

func (o SubAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableSubAccount struct {
	value *SubAccount
	isSet bool
}

func (v NullableSubAccount) Get() *SubAccount {
	return v.value
}

func (v *NullableSubAccount) Set(val *SubAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableSubAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableSubAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubAccount(val *SubAccount) *NullableSubAccount {
	return &NullableSubAccount{value: val, isSet: true}
}

func (v NullableSubAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

