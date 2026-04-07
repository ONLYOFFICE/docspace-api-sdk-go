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

// checks if the TopUpDepositRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TopUpDepositRequestDto{}

// TopUpDepositRequestDto The request parameters for putting money on deposit.
type TopUpDepositRequestDto struct {
	// The amount of money for the operation.
	Amount *int32 `json:"amount,omitempty"`
	// The three-character ISO 4217 currency symbol.
	Currency NullableString `json:"currency,omitempty"`
}

// NewTopUpDepositRequestDto instantiates a new TopUpDepositRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopUpDepositRequestDto() *TopUpDepositRequestDto {
	this := TopUpDepositRequestDto{}
	return &this
}

// NewTopUpDepositRequestDtoWithDefaults instantiates a new TopUpDepositRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopUpDepositRequestDtoWithDefaults() *TopUpDepositRequestDto {
	this := TopUpDepositRequestDto{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *TopUpDepositRequestDto) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopUpDepositRequestDto) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *TopUpDepositRequestDto) IsAmountSet() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *TopUpDepositRequestDto) SetAmount(v int32) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopUpDepositRequestDto) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopUpDepositRequestDto) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *TopUpDepositRequestDto) IsCurrencySet() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *TopUpDepositRequestDto) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *TopUpDepositRequestDto) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *TopUpDepositRequestDto) UnsetCurrency() {
	o.Currency.Unset()
}

func (o TopUpDepositRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TopUpDepositRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	return toSerialize, nil
}

type NullableTopUpDepositRequestDto struct {
	value *TopUpDepositRequestDto
	isSet bool
}

func (v NullableTopUpDepositRequestDto) Get() *TopUpDepositRequestDto {
	return v.value
}

func (v *NullableTopUpDepositRequestDto) Set(val *TopUpDepositRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTopUpDepositRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTopUpDepositRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopUpDepositRequestDto(val *TopUpDepositRequestDto) *NullableTopUpDepositRequestDto {
	return &NullableTopUpDepositRequestDto{value: val, isSet: true}
}

func (v NullableTopUpDepositRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopUpDepositRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

