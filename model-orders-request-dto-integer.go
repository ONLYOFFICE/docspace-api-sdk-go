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

// checks if the OrdersRequestDtoInteger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrdersRequestDtoInteger{}

// OrdersRequestDtoInteger The collection of items to be ordered.
type OrdersRequestDtoInteger struct {
	// The list of items with their ordering information.
	Items []OrdersItemRequestDtoInteger `json:"items"`
}

type _OrdersRequestDtoInteger OrdersRequestDtoInteger

// NewOrdersRequestDtoInteger instantiates a new OrdersRequestDtoInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrdersRequestDtoInteger(items []OrdersItemRequestDtoInteger) *OrdersRequestDtoInteger {
	this := OrdersRequestDtoInteger{}
	this.Items = items
	return &this
}

// NewOrdersRequestDtoIntegerWithDefaults instantiates a new OrdersRequestDtoInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrdersRequestDtoIntegerWithDefaults() *OrdersRequestDtoInteger {
	this := OrdersRequestDtoInteger{}
	return &this
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []OrdersItemRequestDtoInteger will be returned
func (o *OrdersRequestDtoInteger) GetItems() []OrdersItemRequestDtoInteger {
	if o == nil {
		var ret []OrdersItemRequestDtoInteger
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrdersRequestDtoInteger) GetItemsOk() ([]OrdersItemRequestDtoInteger, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *OrdersRequestDtoInteger) SetItems(v []OrdersItemRequestDtoInteger) {
	o.Items = v
}

func (o OrdersRequestDtoInteger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrdersRequestDtoInteger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

func (o *OrdersRequestDtoInteger) UnmarshalJSON(data []byte) (err error) {
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

	varOrdersRequestDtoInteger := _OrdersRequestDtoInteger{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varOrdersRequestDtoInteger)

	if err != nil {
		return err
	}

	*o = OrdersRequestDtoInteger(varOrdersRequestDtoInteger)

	return err
}

type NullableOrdersRequestDtoInteger struct {
	value *OrdersRequestDtoInteger
	isSet bool
}

func (v NullableOrdersRequestDtoInteger) Get() *OrdersRequestDtoInteger {
	return v.value
}

func (v *NullableOrdersRequestDtoInteger) Set(val *OrdersRequestDtoInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableOrdersRequestDtoInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableOrdersRequestDtoInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrdersRequestDtoInteger(val *OrdersRequestDtoInteger) *NullableOrdersRequestDtoInteger {
	return &NullableOrdersRequestDtoInteger{value: val, isSet: true}
}

func (v NullableOrdersRequestDtoInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrdersRequestDtoInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

