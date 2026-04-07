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

// checks if the PaymentUrlRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentUrlRequestDto{}

// PaymentUrlRequestDto The request parameters for the payment URL configuration with quantity information.
type PaymentUrlRequestDto struct {
	// The URL where the user will be redirected after payment processing.
	BackUrl NullableString `json:"backUrl,omitempty"`
	// The payment quantity.
	Quantity map[string]int32 `json:"quantity,omitempty"`
}

// NewPaymentUrlRequestDto instantiates a new PaymentUrlRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentUrlRequestDto() *PaymentUrlRequestDto {
	this := PaymentUrlRequestDto{}
	return &this
}

// NewPaymentUrlRequestDtoWithDefaults instantiates a new PaymentUrlRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentUrlRequestDtoWithDefaults() *PaymentUrlRequestDto {
	this := PaymentUrlRequestDto{}
	return &this
}

// GetBackUrl returns the BackUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentUrlRequestDto) GetBackUrl() string {
	if o == nil || IsNil(o.BackUrl.Get()) {
		var ret string
		return ret
	}
	return *o.BackUrl.Get()
}

// GetBackUrlOk returns a tuple with the BackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentUrlRequestDto) GetBackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackUrl.Get(), o.BackUrl.IsSet()
}

// HasBackUrl returns a boolean if a field has been set.
func (o *PaymentUrlRequestDto) IsBackUrlSet() bool {
	if o != nil && o.BackUrl.IsSet() {
		return true
	}

	return false
}

// SetBackUrl gets a reference to the given NullableString and assigns it to the BackUrl field.
func (o *PaymentUrlRequestDto) SetBackUrl(v string) {
	o.BackUrl.Set(&v)
}
// SetBackUrlNil sets the value for BackUrl to be an explicit nil
func (o *PaymentUrlRequestDto) SetBackUrlNil() {
	o.BackUrl.Set(nil)
}

// UnsetBackUrl ensures that no value is present for BackUrl, not even an explicit nil
func (o *PaymentUrlRequestDto) UnsetBackUrl() {
	o.BackUrl.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentUrlRequestDto) GetQuantity() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}
	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentUrlRequestDto) GetQuantityOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return &o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *PaymentUrlRequestDto) IsQuantitySet() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given map[string]int32 and assigns it to the Quantity field.
func (o *PaymentUrlRequestDto) SetQuantity(v map[string]int32) {
	o.Quantity = v
}

func (o PaymentUrlRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentUrlRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BackUrl.IsSet() {
		toSerialize["backUrl"] = o.BackUrl.Get()
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullablePaymentUrlRequestDto struct {
	value *PaymentUrlRequestDto
	isSet bool
}

func (v NullablePaymentUrlRequestDto) Get() *PaymentUrlRequestDto {
	return v.value
}

func (v *NullablePaymentUrlRequestDto) Set(val *PaymentUrlRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentUrlRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentUrlRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentUrlRequestDto(val *PaymentUrlRequestDto) *NullablePaymentUrlRequestDto {
	return &NullablePaymentUrlRequestDto{value: val, isSet: true}
}

func (v NullablePaymentUrlRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentUrlRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

