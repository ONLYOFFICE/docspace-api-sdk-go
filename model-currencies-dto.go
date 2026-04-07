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

// checks if the CurrenciesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrenciesDto{}

// CurrenciesDto The currencies parameters.
type CurrenciesDto struct {
	// The ISO country code.
	IsoCountryCode NullableString `json:"isoCountryCode,omitempty"`
	// The ISO currency symbol.
	IsoCurrencySymbol NullableString `json:"isoCurrencySymbol,omitempty"`
	// The currency native name.
	CurrencyNativeName NullableString `json:"currencyNativeName,omitempty"`
}

// NewCurrenciesDto instantiates a new CurrenciesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrenciesDto() *CurrenciesDto {
	this := CurrenciesDto{}
	return &this
}

// NewCurrenciesDtoWithDefaults instantiates a new CurrenciesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrenciesDtoWithDefaults() *CurrenciesDto {
	this := CurrenciesDto{}
	return &this
}

// GetIsoCountryCode returns the IsoCountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrenciesDto) GetIsoCountryCode() string {
	if o == nil || IsNil(o.IsoCountryCode.Get()) {
		var ret string
		return ret
	}
	return *o.IsoCountryCode.Get()
}

// GetIsoCountryCodeOk returns a tuple with the IsoCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrenciesDto) GetIsoCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsoCountryCode.Get(), o.IsoCountryCode.IsSet()
}

// HasIsoCountryCode returns a boolean if a field has been set.
func (o *CurrenciesDto) IsIsoCountryCodeSet() bool {
	if o != nil && o.IsoCountryCode.IsSet() {
		return true
	}

	return false
}

// SetIsoCountryCode gets a reference to the given NullableString and assigns it to the IsoCountryCode field.
func (o *CurrenciesDto) SetIsoCountryCode(v string) {
	o.IsoCountryCode.Set(&v)
}
// SetIsoCountryCodeNil sets the value for IsoCountryCode to be an explicit nil
func (o *CurrenciesDto) SetIsoCountryCodeNil() {
	o.IsoCountryCode.Set(nil)
}

// UnsetIsoCountryCode ensures that no value is present for IsoCountryCode, not even an explicit nil
func (o *CurrenciesDto) UnsetIsoCountryCode() {
	o.IsoCountryCode.Unset()
}

// GetIsoCurrencySymbol returns the IsoCurrencySymbol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrenciesDto) GetIsoCurrencySymbol() string {
	if o == nil || IsNil(o.IsoCurrencySymbol.Get()) {
		var ret string
		return ret
	}
	return *o.IsoCurrencySymbol.Get()
}

// GetIsoCurrencySymbolOk returns a tuple with the IsoCurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrenciesDto) GetIsoCurrencySymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsoCurrencySymbol.Get(), o.IsoCurrencySymbol.IsSet()
}

// HasIsoCurrencySymbol returns a boolean if a field has been set.
func (o *CurrenciesDto) IsIsoCurrencySymbolSet() bool {
	if o != nil && o.IsoCurrencySymbol.IsSet() {
		return true
	}

	return false
}

// SetIsoCurrencySymbol gets a reference to the given NullableString and assigns it to the IsoCurrencySymbol field.
func (o *CurrenciesDto) SetIsoCurrencySymbol(v string) {
	o.IsoCurrencySymbol.Set(&v)
}
// SetIsoCurrencySymbolNil sets the value for IsoCurrencySymbol to be an explicit nil
func (o *CurrenciesDto) SetIsoCurrencySymbolNil() {
	o.IsoCurrencySymbol.Set(nil)
}

// UnsetIsoCurrencySymbol ensures that no value is present for IsoCurrencySymbol, not even an explicit nil
func (o *CurrenciesDto) UnsetIsoCurrencySymbol() {
	o.IsoCurrencySymbol.Unset()
}

// GetCurrencyNativeName returns the CurrencyNativeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrenciesDto) GetCurrencyNativeName() string {
	if o == nil || IsNil(o.CurrencyNativeName.Get()) {
		var ret string
		return ret
	}
	return *o.CurrencyNativeName.Get()
}

// GetCurrencyNativeNameOk returns a tuple with the CurrencyNativeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrenciesDto) GetCurrencyNativeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrencyNativeName.Get(), o.CurrencyNativeName.IsSet()
}

// HasCurrencyNativeName returns a boolean if a field has been set.
func (o *CurrenciesDto) IsCurrencyNativeNameSet() bool {
	if o != nil && o.CurrencyNativeName.IsSet() {
		return true
	}

	return false
}

// SetCurrencyNativeName gets a reference to the given NullableString and assigns it to the CurrencyNativeName field.
func (o *CurrenciesDto) SetCurrencyNativeName(v string) {
	o.CurrencyNativeName.Set(&v)
}
// SetCurrencyNativeNameNil sets the value for CurrencyNativeName to be an explicit nil
func (o *CurrenciesDto) SetCurrencyNativeNameNil() {
	o.CurrencyNativeName.Set(nil)
}

// UnsetCurrencyNativeName ensures that no value is present for CurrencyNativeName, not even an explicit nil
func (o *CurrenciesDto) UnsetCurrencyNativeName() {
	o.CurrencyNativeName.Unset()
}

func (o CurrenciesDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrenciesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IsoCountryCode.IsSet() {
		toSerialize["isoCountryCode"] = o.IsoCountryCode.Get()
	}
	if o.IsoCurrencySymbol.IsSet() {
		toSerialize["isoCurrencySymbol"] = o.IsoCurrencySymbol.Get()
	}
	if o.CurrencyNativeName.IsSet() {
		toSerialize["currencyNativeName"] = o.CurrencyNativeName.Get()
	}
	return toSerialize, nil
}

type NullableCurrenciesDto struct {
	value *CurrenciesDto
	isSet bool
}

func (v NullableCurrenciesDto) Get() *CurrenciesDto {
	return v.value
}

func (v *NullableCurrenciesDto) Set(val *CurrenciesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrenciesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrenciesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrenciesDto(val *CurrenciesDto) *NullableCurrenciesDto {
	return &NullableCurrenciesDto{value: val, isSet: true}
}

func (v NullableCurrenciesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrenciesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

