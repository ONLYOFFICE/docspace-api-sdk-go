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

// checks if the SubAccountDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubAccountDto{}

// SubAccountDto A sub-account of the wallet: its currency and balance.
type SubAccountDto struct {
	// The three-character ISO 4217 currency symbol.
	Currency NullableString `json:"currency,omitempty"`
	// The amount in the specified currency.
	Amount *float64 `json:"amount,omitempty"`
}

// NewSubAccountDto instantiates a new SubAccountDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubAccountDto() *SubAccountDto {
	this := SubAccountDto{}
	return &this
}

// NewSubAccountDtoWithDefaults instantiates a new SubAccountDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubAccountDtoWithDefaults() *SubAccountDto {
	this := SubAccountDto{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubAccountDto) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubAccountDto) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *SubAccountDto) IsCurrencySet() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *SubAccountDto) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *SubAccountDto) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *SubAccountDto) UnsetCurrency() {
	o.Currency.Unset()
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *SubAccountDto) GetAmount() float64 {
	if o == nil || IsNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubAccountDto) GetAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *SubAccountDto) IsAmountSet() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *SubAccountDto) SetAmount(v float64) {
	o.Amount = &v
}

func (o SubAccountDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubAccountDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableSubAccountDto struct {
	value *SubAccountDto
	isSet bool
}

func (v NullableSubAccountDto) Get() *SubAccountDto {
	return v.value
}

func (v *NullableSubAccountDto) Set(val *SubAccountDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSubAccountDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSubAccountDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubAccountDto(val *SubAccountDto) *NullableSubAccountDto {
	return &NullableSubAccountDto{value: val, isSet: true}
}

func (v NullableSubAccountDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubAccountDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

