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

// checks if the OrdersRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrdersRequestDto{}

// OrdersRequestDto The request that moves several files and folders to given positions.
type OrdersRequestDto struct {
	// The entries to move, applied one after another in the order they are sent, so each of them shifts the  neighbours the ones before it left behind.
	Items []OrdersItemRequestDto `json:"items"`
}

type _OrdersRequestDto OrdersRequestDto

// NewOrdersRequestDto instantiates a new OrdersRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrdersRequestDto(items []OrdersItemRequestDto) *OrdersRequestDto {
	this := OrdersRequestDto{}
	this.Items = items
	return &this
}

// NewOrdersRequestDtoWithDefaults instantiates a new OrdersRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrdersRequestDtoWithDefaults() *OrdersRequestDto {
	this := OrdersRequestDto{}
	return &this
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []OrdersItemRequestDto will be returned
func (o *OrdersRequestDto) GetItems() []OrdersItemRequestDto {
	if o == nil {
		var ret []OrdersItemRequestDto
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrdersRequestDto) GetItemsOk() ([]OrdersItemRequestDto, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *OrdersRequestDto) SetItems(v []OrdersItemRequestDto) {
	o.Items = v
}

func (o OrdersRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrdersRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

func (o *OrdersRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
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

	varOrdersRequestDto := _OrdersRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varOrdersRequestDto)

	if err != nil {
		return err
	}

	*o = OrdersRequestDto(varOrdersRequestDto)

	return err
}

type NullableOrdersRequestDto struct {
	value *OrdersRequestDto
	isSet bool
}

func (v NullableOrdersRequestDto) Get() *OrdersRequestDto {
	return v.value
}

func (v *NullableOrdersRequestDto) Set(val *OrdersRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOrdersRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOrdersRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrdersRequestDto(val *OrdersRequestDto) *NullableOrdersRequestDto {
	return &NullableOrdersRequestDto{value: val, isSet: true}
}

func (v NullableOrdersRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrdersRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

