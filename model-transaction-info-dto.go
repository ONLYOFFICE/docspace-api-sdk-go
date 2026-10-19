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

// checks if the TransactionInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionInfoDto{}

// TransactionInfoDto Represents information about the transaction applied to an account.
type TransactionInfoDto struct {
	// The date and time when the credit transaction occurred.
	Date *time.Time `json:"date,omitempty"`
	// The three-character ISO 4217 currency symbol.
	Currency NullableString `json:"currency,omitempty"`
	// The amount in the specified currency.
	Amount *float64 `json:"amount,omitempty"`
}

// NewTransactionInfoDto instantiates a new TransactionInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionInfoDto() *TransactionInfoDto {
	this := TransactionInfoDto{}
	return &this
}

// NewTransactionInfoDtoWithDefaults instantiates a new TransactionInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionInfoDtoWithDefaults() *TransactionInfoDto {
	this := TransactionInfoDto{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *TransactionInfoDto) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionInfoDto) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *TransactionInfoDto) IsDateSet() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *TransactionInfoDto) SetDate(v time.Time) {
	o.Date = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionInfoDto) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionInfoDto) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *TransactionInfoDto) IsCurrencySet() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *TransactionInfoDto) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *TransactionInfoDto) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *TransactionInfoDto) UnsetCurrency() {
	o.Currency.Unset()
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *TransactionInfoDto) GetAmount() float64 {
	if o == nil || IsNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionInfoDto) GetAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *TransactionInfoDto) IsAmountSet() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *TransactionInfoDto) SetAmount(v float64) {
	o.Amount = &v
}

func (o TransactionInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableTransactionInfoDto struct {
	value *TransactionInfoDto
	isSet bool
}

func (v NullableTransactionInfoDto) Get() *TransactionInfoDto {
	return v.value
}

func (v *NullableTransactionInfoDto) Set(val *TransactionInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionInfoDto(val *TransactionInfoDto) *NullableTransactionInfoDto {
	return &NullableTransactionInfoDto{value: val, isSet: true}
}

func (v NullableTransactionInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

