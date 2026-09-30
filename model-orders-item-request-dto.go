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

// checks if the OrdersItemRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrdersItemRequestDto{}

// OrdersItemRequestDto One entry to move to a given position inside its folder.
type OrdersItemRequestDto struct {
	// The file or folder to move.
	EntryId int32 `json:"entryId"`
	// Which of the two the identifier names, because a file and a folder may carry the same number.
	EntryType FileEntryType `json:"entryType"`
	// The position the entry is to take, counting from 1. The entry that held it, and everything after it, is  shifted to make room. A dotted path such as 1.2.3 is accepted as well, of which only the last segment is  read.
	Order int32 `json:"order"`
}

type _OrdersItemRequestDto OrdersItemRequestDto

// NewOrdersItemRequestDto instantiates a new OrdersItemRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrdersItemRequestDto(entryId int32, entryType FileEntryType, order int32) *OrdersItemRequestDto {
	this := OrdersItemRequestDto{}
	this.EntryId = entryId
	this.EntryType = entryType
	this.Order = order
	return &this
}

// NewOrdersItemRequestDtoWithDefaults instantiates a new OrdersItemRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrdersItemRequestDtoWithDefaults() *OrdersItemRequestDto {
	this := OrdersItemRequestDto{}
	return &this
}

// GetEntryId returns the EntryId field value
func (o *OrdersItemRequestDto) GetEntryId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *OrdersItemRequestDto) GetEntryIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *OrdersItemRequestDto) SetEntryId(v int32) {
	o.EntryId = v
}

// GetEntryType returns the EntryType field value
func (o *OrdersItemRequestDto) GetEntryType() FileEntryType {
	if o == nil {
		var ret FileEntryType
		return ret
	}

	return o.EntryType
}

// GetEntryTypeOk returns a tuple with the EntryType field value
// and a boolean to check if the value has been set.
func (o *OrdersItemRequestDto) GetEntryTypeOk() (*FileEntryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryType, true
}

// SetEntryType sets field value
func (o *OrdersItemRequestDto) SetEntryType(v FileEntryType) {
	o.EntryType = v
}

// GetOrder returns the Order field value
func (o *OrdersItemRequestDto) GetOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *OrdersItemRequestDto) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *OrdersItemRequestDto) SetOrder(v int32) {
	o.Order = v
}

func (o OrdersItemRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrdersItemRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entryId"] = o.EntryId
	toSerialize["entryType"] = o.EntryType
	toSerialize["order"] = o.Order
	return toSerialize, nil
}

func (o *OrdersItemRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entryId",
		"entryType",
		"order",
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

	varOrdersItemRequestDto := _OrdersItemRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varOrdersItemRequestDto)

	if err != nil {
		return err
	}

	*o = OrdersItemRequestDto(varOrdersItemRequestDto)

	return err
}

type NullableOrdersItemRequestDto struct {
	value *OrdersItemRequestDto
	isSet bool
}

func (v NullableOrdersItemRequestDto) Get() *OrdersItemRequestDto {
	return v.value
}

func (v *NullableOrdersItemRequestDto) Set(val *OrdersItemRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOrdersItemRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOrdersItemRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrdersItemRequestDto(val *OrdersItemRequestDto) *NullableOrdersItemRequestDto {
	return &NullableOrdersItemRequestDto{value: val, isSet: true}
}

func (v NullableOrdersItemRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrdersItemRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

