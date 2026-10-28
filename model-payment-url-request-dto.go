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

// PaymentUrlRequestDto The plan being bought and the two pages the hosted checkout returns the buyer to.
type PaymentUrlRequestDto struct {
	// The absolute address the hosted checkout page sends the buyer back to when the purchase is abandoned. It has  to be a well-formed URL and is carried into the checkout page as it is given, so it must be reachable by the  buyer rather than by the portal.
	BackUrl string `json:"backUrl"`
	// The absolute address the hosted checkout page sends the buyer to once the payment provider accepts the  purchase. Reaching it says the provider took the money, not that the portal has already been switched to the  new plan, so a client that lands here reads the plan back rather than assuming it.
	SuccessUrl string `json:"successUrl"`
	// The plan being bought, as a single pair of the plan name and the number of units of it. The key is the `name`  of a monthly, non-wallet quota from `GET api/2.0/portal/payment/quotas`, and the value is how many  administrators the plan is to cover, which has to be greater than zero. Exactly one pair is accepted; yearly  and wallet products are refused with 400, and wallet services are bought through  `PUT api/2.0/portal/payment/updatewallet` instead.
	Quantity map[string]int32 `json:"quantity"`
}

type _PaymentUrlRequestDto PaymentUrlRequestDto

// NewPaymentUrlRequestDto instantiates a new PaymentUrlRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentUrlRequestDto(backUrl string, successUrl string, quantity map[string]int32) *PaymentUrlRequestDto {
	this := PaymentUrlRequestDto{}
	this.BackUrl = backUrl
	this.SuccessUrl = successUrl
	this.Quantity = quantity
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

// GetQuantity returns the Quantity field value
func (o *PaymentUrlRequestDto) GetQuantity() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *PaymentUrlRequestDto) GetQuantityOk() (map[string]int32, bool) {
	if o == nil {
		return map[string]int32{}, false
	}
	return o.Quantity, true
}

// SetQuantity sets field value
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
	toSerialize["quantity"] = o.Quantity
	return toSerialize, nil
}

func (o *PaymentUrlRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"backUrl",
		"successUrl",
		"quantity",
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

