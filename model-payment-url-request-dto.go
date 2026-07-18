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
	"bytes"
	"fmt"
)

// checks if the PaymentUrlRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentUrlRequestDto{}

// PaymentUrlRequestDto The request parameters for the payment URL configuration with quantity information.
type PaymentUrlRequestDto struct {
	// The URL where the user will be redirected after payment cancellation.
	BackUrl string `json:"backUrl"`
	// The URL where the user will be redirected after successful payment.
	SuccessUrl string `json:"successUrl"`
	// The payment quantity.
	Quantity map[string]int32 `json:"quantity,omitempty"`
}

type _PaymentUrlRequestDto PaymentUrlRequestDto

// NewPaymentUrlRequestDto instantiates a new PaymentUrlRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentUrlRequestDto(backUrl string, successUrl string) *PaymentUrlRequestDto {
	this := PaymentUrlRequestDto{}
	this.BackUrl = backUrl
	this.SuccessUrl = successUrl
	return &this
}

// NewPaymentUrlRequestDtoWithDefaults instantiates a new PaymentUrlRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentUrlRequestDtoWithDefaults() *PaymentUrlRequestDto {
	this := PaymentUrlRequestDto{}
	return &this
}

// GetBackUrl returns the BackUrl field value
func (o *PaymentUrlRequestDto) GetBackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackUrl
}

// GetBackUrlOk returns a tuple with the BackUrl field value
// and a boolean to check if the value has been set.
func (o *PaymentUrlRequestDto) GetBackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackUrl, true
}

// SetBackUrl sets field value
func (o *PaymentUrlRequestDto) SetBackUrl(v string) {
	o.BackUrl = v
}

// GetSuccessUrl returns the SuccessUrl field value
func (o *PaymentUrlRequestDto) GetSuccessUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SuccessUrl
}

// GetSuccessUrlOk returns a tuple with the SuccessUrl field value
// and a boolean to check if the value has been set.
func (o *PaymentUrlRequestDto) GetSuccessUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccessUrl, true
}

// SetSuccessUrl sets field value
func (o *PaymentUrlRequestDto) SetSuccessUrl(v string) {
	o.SuccessUrl = v
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
	toSerialize["backUrl"] = o.BackUrl
	toSerialize["successUrl"] = o.SuccessUrl
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

func (o *PaymentUrlRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"backUrl",
		"successUrl",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPaymentUrlRequestDto := _PaymentUrlRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varPaymentUrlRequestDto)

	if err != nil {
		return err
	}

	*o = PaymentUrlRequestDto(varPaymentUrlRequestDto)

	return err
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

