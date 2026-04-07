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

// checks if the QuantityRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuantityRequestDto{}

// QuantityRequestDto The request parameters for specifying payment quantity.
type QuantityRequestDto struct {
	// The mapping of item identifiers to their respective quantities in the payment.
	Quantity map[string]int32 `json:"quantity,omitempty"`
}

// NewQuantityRequestDto instantiates a new QuantityRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuantityRequestDto() *QuantityRequestDto {
	this := QuantityRequestDto{}
	return &this
}

// NewQuantityRequestDtoWithDefaults instantiates a new QuantityRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuantityRequestDtoWithDefaults() *QuantityRequestDto {
	this := QuantityRequestDto{}
	return &this
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QuantityRequestDto) GetQuantity() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}
	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QuantityRequestDto) GetQuantityOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return &o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *QuantityRequestDto) IsQuantitySet() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given map[string]int32 and assigns it to the Quantity field.
func (o *QuantityRequestDto) SetQuantity(v map[string]int32) {
	o.Quantity = v
}

func (o QuantityRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuantityRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	return toSerialize, nil
}

type NullableQuantityRequestDto struct {
	value *QuantityRequestDto
	isSet bool
}

func (v NullableQuantityRequestDto) Get() *QuantityRequestDto {
	return v.value
}

func (v *NullableQuantityRequestDto) Set(val *QuantityRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableQuantityRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableQuantityRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuantityRequestDto(val *QuantityRequestDto) *NullableQuantityRequestDto {
	return &NullableQuantityRequestDto{value: val, isSet: true}
}

func (v NullableQuantityRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuantityRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

