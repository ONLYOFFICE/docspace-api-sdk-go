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

// checks if the PaymentCalculation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentCalculation{}

// PaymentCalculation The parameters of the calculated payment amount.
type PaymentCalculation struct {
	// The operation unique identifier.
	OperationId *int64 `json:"operationId,omitempty"`
	// The calculated payment amount.
	Amount *float64 `json:"amount,omitempty"`
	// The three-character ISO 4217 currency symbol used for the payment calculation.
	Currency NullableString `json:"currency,omitempty"`
	// The quantity associated with the payment calculation.
	Quantity *int32 `json:"quantity,omitempty"`
}

// NewPaymentCalculation instantiates a new PaymentCalculation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentCalculation() *PaymentCalculation {
	this := PaymentCalculation{}
	return &this
}

// NewPaymentCalculationWithDefaults instantiates a new PaymentCalculation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentCalculationWithDefaults() *PaymentCalculation {
	this := PaymentCalculation{}
	return &this
}

// GetOperationId returns the OperationId field value if set, zero value otherwise.
func (o *PaymentCalculation) GetOperationId() int64 {
	if o == nil || IsNil(o.OperationId) {
		var ret int64
		return ret
	}
	return *o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCalculation) GetOperationIdOk() (*int64, bool) {
	if o == nil || IsNil(o.OperationId) {
		return nil, false
	}
	return o.OperationId, true
}

// HasOperationId returns a boolean if a field has been set.
func (o *PaymentCalculation) IsOperationIdSet() bool {
	if o != nil && !IsNil(o.OperationId) {
		return true
	}

	return false
}

// SetOperationId gets a reference to the given int64 and assigns it to the OperationId field.
func (o *PaymentCalculation) SetOperationId(v int64) {
	o.OperationId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PaymentCalculation) GetAmount() float64 {
	if o == nil || IsNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCalculation) GetAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PaymentCalculation) IsAmountSet() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *PaymentCalculation) SetAmount(v float64) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentCalculation) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentCalculation) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *PaymentCalculation) IsCurrencySet() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *PaymentCalculation) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *PaymentCalculation) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *PaymentCalculation) UnsetCurrency() {
	o.Currency.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *PaymentCalculation) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCalculation) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *PaymentCalculation) IsQuantitySet() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *PaymentCalculation) SetQuantity(v int32) {
	o.Quantity = &v
}

func (o PaymentCalculation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentCalculation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OperationId) {
		toSerialize["operationId"] = o.OperationId
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullablePaymentCalculation struct {
	value *PaymentCalculation
	isSet bool
}

func (v NullablePaymentCalculation) Get() *PaymentCalculation {
	return v.value
}

func (v *NullablePaymentCalculation) Set(val *PaymentCalculation) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentCalculation) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentCalculation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentCalculation(val *PaymentCalculation) *NullablePaymentCalculation {
	return &NullablePaymentCalculation{value: val, isSet: true}
}

func (v NullablePaymentCalculation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentCalculation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

