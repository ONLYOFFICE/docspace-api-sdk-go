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

// checks if the CurrencyCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrencyCode{}

// CurrencyCode The currency an amount is expressed in.
type CurrencyCode struct {
	// The three-character ISO 4217 currency symbol.
	Currency NullableString `json:"currency,omitempty"`
}

// NewCurrencyCode instantiates a new CurrencyCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrencyCode() *CurrencyCode {
	this := CurrencyCode{}
	return &this
}

// NewCurrencyCodeWithDefaults instantiates a new CurrencyCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrencyCodeWithDefaults() *CurrencyCode {
	this := CurrencyCode{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrencyCode) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrencyCode) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *CurrencyCode) IsCurrencySet() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *CurrencyCode) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *CurrencyCode) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *CurrencyCode) UnsetCurrency() {
	o.Currency.Unset()
}

func (o CurrencyCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrencyCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	return toSerialize, nil
}

type NullableCurrencyCode struct {
	value *CurrencyCode
	isSet bool
}

func (v NullableCurrencyCode) Get() *CurrencyCode {
	return v.value
}

func (v *NullableCurrencyCode) Set(val *CurrencyCode) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrencyCode) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrencyCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrencyCode(val *CurrencyCode) *NullableCurrencyCode {
	return &NullableCurrencyCode{value: val, isSet: true}
}

func (v NullableCurrencyCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrencyCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

