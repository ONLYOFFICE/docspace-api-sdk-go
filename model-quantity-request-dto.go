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

// checks if the QuantityRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuantityRequestDto{}

// QuantityRequestDto The new size of the portal subscription.
type QuantityRequestDto struct {
	// The plan and the number of units it is to cover, as a single pair. While the portal is on a priced plan the  key has to be the `name` of that same plan, which `GET api/2.0/portal/payment/quota` reports, because the  subscription is resized rather than swapped; the value is the total the subscription is to have afterwards,  not the difference. Exactly one pair is accepted, and a value that is already in effect is refused with 400.
	Quantity map[string]int32 `json:"quantity"`
}

type _QuantityRequestDto QuantityRequestDto

// NewQuantityRequestDto instantiates a new QuantityRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuantityRequestDto(quantity map[string]int32) *QuantityRequestDto {
	this := QuantityRequestDto{}
	this.Quantity = quantity
	return &this
}

// NewQuantityRequestDtoWithDefaults instantiates a new QuantityRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuantityRequestDtoWithDefaults() *QuantityRequestDto {
	this := QuantityRequestDto{}
	return &this
}

// GetQuantity returns the Quantity field value
func (o *QuantityRequestDto) GetQuantity() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *QuantityRequestDto) GetQuantityOk() (map[string]int32, bool) {
	if o == nil {
		return map[string]int32{}, false
	}
	return o.Quantity, true
}

// SetQuantity sets field value
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
	toSerialize["quantity"] = o.Quantity
	return toSerialize, nil
}

func (o *QuantityRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varQuantityRequestDto := _QuantityRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varQuantityRequestDto)

	if err != nil {
		return err
	}

	*o = QuantityRequestDto(varQuantityRequestDto)

	return err
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

