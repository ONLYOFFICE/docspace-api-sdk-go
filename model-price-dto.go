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

// checks if the PriceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceDto{}

// PriceDto What a quota costs, and the currency that amount is in.
type PriceDto struct {
	// The amount for one billing period, per unit for a quota sold by the unit. It is empty for a quota that is  not sold for money - the free, trial and non-profit ones - and for a quota this installation has no price  list entry for.
	Value NullableFloat64 `json:"value,omitempty"`
	// The symbol to print in front of `value`, such as `$`. It is chosen for the currency, not for the portal  language, so it is not a localised format.
	CurrencySymbol NullableString `json:"currencySymbol,omitempty"`
	// The currency as a three-letter ISO 4217 code, which is the value to compare on when `currencySymbol` is  ambiguous between currencies that share a sign.
	IsoCurrencySymbol NullableString `json:"isoCurrencySymbol,omitempty"`
}

// NewPriceDto instantiates a new PriceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceDto() *PriceDto {
	this := PriceDto{}
	return &this
}

// NewPriceDtoWithDefaults instantiates a new PriceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceDtoWithDefaults() *PriceDto {
	this := PriceDto{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceDto) GetValue() float64 {
	if o == nil || IsNil(o.Value.Get()) {
		var ret float64
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceDto) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *PriceDto) IsValueSet() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableFloat64 and assigns it to the Value field.
func (o *PriceDto) SetValue(v float64) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *PriceDto) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *PriceDto) UnsetValue() {
	o.Value.Unset()
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceDto) GetCurrencySymbol() string {
	if o == nil || IsNil(o.CurrencySymbol.Get()) {
		var ret string
		return ret
	}
	return *o.CurrencySymbol.Get()
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceDto) GetCurrencySymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrencySymbol.Get(), o.CurrencySymbol.IsSet()
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *PriceDto) IsCurrencySymbolSet() bool {
	if o != nil && o.CurrencySymbol.IsSet() {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given NullableString and assigns it to the CurrencySymbol field.
func (o *PriceDto) SetCurrencySymbol(v string) {
	o.CurrencySymbol.Set(&v)
}
// SetCurrencySymbolNil sets the value for CurrencySymbol to be an explicit nil
func (o *PriceDto) SetCurrencySymbolNil() {
	o.CurrencySymbol.Set(nil)
}

// UnsetCurrencySymbol ensures that no value is present for CurrencySymbol, not even an explicit nil
func (o *PriceDto) UnsetCurrencySymbol() {
	o.CurrencySymbol.Unset()
}

// GetIsoCurrencySymbol returns the IsoCurrencySymbol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceDto) GetIsoCurrencySymbol() string {
	if o == nil || IsNil(o.IsoCurrencySymbol.Get()) {
		var ret string
		return ret
	}
	return *o.IsoCurrencySymbol.Get()
}

// GetIsoCurrencySymbolOk returns a tuple with the IsoCurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceDto) GetIsoCurrencySymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsoCurrencySymbol.Get(), o.IsoCurrencySymbol.IsSet()
}

// HasIsoCurrencySymbol returns a boolean if a field has been set.
func (o *PriceDto) IsIsoCurrencySymbolSet() bool {
	if o != nil && o.IsoCurrencySymbol.IsSet() {
		return true
	}

	return false
}

// SetIsoCurrencySymbol gets a reference to the given NullableString and assigns it to the IsoCurrencySymbol field.
func (o *PriceDto) SetIsoCurrencySymbol(v string) {
	o.IsoCurrencySymbol.Set(&v)
}
// SetIsoCurrencySymbolNil sets the value for IsoCurrencySymbol to be an explicit nil
func (o *PriceDto) SetIsoCurrencySymbolNil() {
	o.IsoCurrencySymbol.Set(nil)
}

// UnsetIsoCurrencySymbol ensures that no value is present for IsoCurrencySymbol, not even an explicit nil
func (o *PriceDto) UnsetIsoCurrencySymbol() {
	o.IsoCurrencySymbol.Unset()
}

func (o PriceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.CurrencySymbol.IsSet() {
		toSerialize["currencySymbol"] = o.CurrencySymbol.Get()
	}
	if o.IsoCurrencySymbol.IsSet() {
		toSerialize["isoCurrencySymbol"] = o.IsoCurrencySymbol.Get()
	}
	return toSerialize, nil
}

type NullablePriceDto struct {
	value *PriceDto
	isSet bool
}

func (v NullablePriceDto) Get() *PriceDto {
	return v.value
}

func (v *NullablePriceDto) Set(val *PriceDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceDto(val *PriceDto) *NullablePriceDto {
	return &NullablePriceDto{value: val, isSet: true}
}

func (v NullablePriceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

