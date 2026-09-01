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

// checks if the CurrencyAmount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrencyAmount{}

// CurrencyAmount An amount of money together with its currency.
type CurrencyAmount struct {
	// The three-character ISO 4217 currency symbol.
	Currency *string `json:"currency,omitempty"`
	// The amount in the specified currency.
	Amount *float64 `json:"amount,omitempty"`
}

// NewCurrencyAmount instantiates a new CurrencyAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrencyAmount() *CurrencyAmount {
	this := CurrencyAmount{}
	return &this
}

// NewCurrencyAmountWithDefaults instantiates a new CurrencyAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrencyAmountWithDefaults() *CurrencyAmount {
	this := CurrencyAmount{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CurrencyAmount) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrencyAmount) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CurrencyAmount) IsCurrencySet() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *CurrencyAmount) SetCurrency(v string) {
	o.Currency = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CurrencyAmount) GetAmount() float64 {
	if o == nil || IsNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrencyAmount) GetAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CurrencyAmount) IsAmountSet() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *CurrencyAmount) SetAmount(v float64) {
	o.Amount = &v
}

func (o CurrencyAmount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrencyAmount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableCurrencyAmount struct {
	value *CurrencyAmount
	isSet bool
}

func (v NullableCurrencyAmount) Get() *CurrencyAmount {
	return v.value
}

func (v *NullableCurrencyAmount) Set(val *CurrencyAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrencyAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrencyAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrencyAmount(val *CurrencyAmount) *NullableCurrencyAmount {
	return &NullableCurrencyAmount{value: val, isSet: true}
}

func (v NullableCurrencyAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrencyAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

