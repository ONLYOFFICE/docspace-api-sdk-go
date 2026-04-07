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

// checks if the WalletQuantityRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletQuantityRequestDto{}

// WalletQuantityRequestDto The request parameters for specifying wallet payment quantity.
type WalletQuantityRequestDto struct {
	// The mapping of item identifiers to their respective quantities in the payment.
	Quantity map[string]int32 `json:"quantity,omitempty"`
	ProductQuantityType *ProductQuantityType `json:"productQuantityType,omitempty"`
}

// NewWalletQuantityRequestDto instantiates a new WalletQuantityRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletQuantityRequestDto() *WalletQuantityRequestDto {
	this := WalletQuantityRequestDto{}
	return &this
}

// NewWalletQuantityRequestDtoWithDefaults instantiates a new WalletQuantityRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletQuantityRequestDtoWithDefaults() *WalletQuantityRequestDto {
	this := WalletQuantityRequestDto{}
	return &this
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WalletQuantityRequestDto) GetQuantity() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}
	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WalletQuantityRequestDto) GetQuantityOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return &o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *WalletQuantityRequestDto) IsQuantitySet() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given map[string]int32 and assigns it to the Quantity field.
func (o *WalletQuantityRequestDto) SetQuantity(v map[string]int32) {
	o.Quantity = v
}

// GetProductQuantityType returns the ProductQuantityType field value if set, zero value otherwise.
func (o *WalletQuantityRequestDto) GetProductQuantityType() ProductQuantityType {
	if o == nil || IsNil(o.ProductQuantityType) {
		var ret ProductQuantityType
		return ret
	}
	return *o.ProductQuantityType
}

// GetProductQuantityTypeOk returns a tuple with the ProductQuantityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletQuantityRequestDto) GetProductQuantityTypeOk() (*ProductQuantityType, bool) {
	if o == nil || IsNil(o.ProductQuantityType) {
		return nil, false
	}
	return o.ProductQuantityType, true
}

// HasProductQuantityType returns a boolean if a field has been set.
func (o *WalletQuantityRequestDto) IsProductQuantityTypeSet() bool {
	if o != nil && !IsNil(o.ProductQuantityType) {
		return true
	}

	return false
}

// SetProductQuantityType gets a reference to the given ProductQuantityType and assigns it to the ProductQuantityType field.
func (o *WalletQuantityRequestDto) SetProductQuantityType(v ProductQuantityType) {
	o.ProductQuantityType = &v
}

func (o WalletQuantityRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletQuantityRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.ProductQuantityType) {
		toSerialize["productQuantityType"] = o.ProductQuantityType
	}
	return toSerialize, nil
}

type NullableWalletQuantityRequestDto struct {
	value *WalletQuantityRequestDto
	isSet bool
}

func (v NullableWalletQuantityRequestDto) Get() *WalletQuantityRequestDto {
	return v.value
}

func (v *NullableWalletQuantityRequestDto) Set(val *WalletQuantityRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletQuantityRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletQuantityRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletQuantityRequestDto(val *WalletQuantityRequestDto) *NullableWalletQuantityRequestDto {
	return &NullableWalletQuantityRequestDto{value: val, isSet: true}
}

func (v NullableWalletQuantityRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletQuantityRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

