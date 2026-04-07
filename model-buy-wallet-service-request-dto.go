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

// checks if the BuyWalletServiceRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuyWalletServiceRequestDto{}

// BuyWalletServiceRequestDto The request parameters for buying wallet service.
type BuyWalletServiceRequestDto struct {
	// Number of services provided.
	Quantity *int32 `json:"quantity,omitempty"`
	// The service name.
	ServiceName NullableString `json:"serviceName,omitempty"`
}

// NewBuyWalletServiceRequestDto instantiates a new BuyWalletServiceRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyWalletServiceRequestDto() *BuyWalletServiceRequestDto {
	this := BuyWalletServiceRequestDto{}
	return &this
}

// NewBuyWalletServiceRequestDtoWithDefaults instantiates a new BuyWalletServiceRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyWalletServiceRequestDtoWithDefaults() *BuyWalletServiceRequestDto {
	this := BuyWalletServiceRequestDto{}
	return &this
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *BuyWalletServiceRequestDto) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyWalletServiceRequestDto) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *BuyWalletServiceRequestDto) IsQuantitySet() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *BuyWalletServiceRequestDto) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BuyWalletServiceRequestDto) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName.Get()) {
		var ret string
		return ret
	}
	return *o.ServiceName.Get()
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BuyWalletServiceRequestDto) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceName.Get(), o.ServiceName.IsSet()
}

// HasServiceName returns a boolean if a field has been set.
func (o *BuyWalletServiceRequestDto) IsServiceNameSet() bool {
	if o != nil && o.ServiceName.IsSet() {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given NullableString and assigns it to the ServiceName field.
func (o *BuyWalletServiceRequestDto) SetServiceName(v string) {
	o.ServiceName.Set(&v)
}
// SetServiceNameNil sets the value for ServiceName to be an explicit nil
func (o *BuyWalletServiceRequestDto) SetServiceNameNil() {
	o.ServiceName.Set(nil)
}

// UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil
func (o *BuyWalletServiceRequestDto) UnsetServiceName() {
	o.ServiceName.Unset()
}

func (o BuyWalletServiceRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuyWalletServiceRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if o.ServiceName.IsSet() {
		toSerialize["serviceName"] = o.ServiceName.Get()
	}
	return toSerialize, nil
}

type NullableBuyWalletServiceRequestDto struct {
	value *BuyWalletServiceRequestDto
	isSet bool
}

func (v NullableBuyWalletServiceRequestDto) Get() *BuyWalletServiceRequestDto {
	return v.value
}

func (v *NullableBuyWalletServiceRequestDto) Set(val *BuyWalletServiceRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyWalletServiceRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyWalletServiceRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyWalletServiceRequestDto(val *BuyWalletServiceRequestDto) *NullableBuyWalletServiceRequestDto {
	return &NullableBuyWalletServiceRequestDto{value: val, isSet: true}
}

func (v NullableBuyWalletServiceRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyWalletServiceRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

