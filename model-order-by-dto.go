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

// checks if the OrderByDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderByDto{}

// OrderByDto The sorting parameters.
type OrderByDto struct {
	// Specifies if the order is ascending.
	IsAsc *bool `json:"is_asc,omitempty"`
	// The parameters by which the files will be sorted.
	Property *SortedByType `json:"property,omitempty"`
}

// NewOrderByDto instantiates a new OrderByDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderByDto() *OrderByDto {
	this := OrderByDto{}
	return &this
}

// NewOrderByDtoWithDefaults instantiates a new OrderByDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderByDtoWithDefaults() *OrderByDto {
	this := OrderByDto{}
	return &this
}

// GetIsAsc returns the IsAsc field value if set, zero value otherwise.
func (o *OrderByDto) GetIsAsc() bool {
	if o == nil || IsNil(o.IsAsc) {
		var ret bool
		return ret
	}
	return *o.IsAsc
}

// GetIsAscOk returns a tuple with the IsAsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderByDto) GetIsAscOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAsc) {
		return nil, false
	}
	return o.IsAsc, true
}

// HasIsAsc returns a boolean if a field has been set.
func (o *OrderByDto) IsIsAscSet() bool {
	if o != nil && !IsNil(o.IsAsc) {
		return true
	}

	return false
}

// SetIsAsc gets a reference to the given bool and assigns it to the IsAsc field.
func (o *OrderByDto) SetIsAsc(v bool) {
	o.IsAsc = &v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *OrderByDto) GetProperty() SortedByType {
	if o == nil || IsNil(o.Property) {
		var ret SortedByType
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderByDto) GetPropertyOk() (*SortedByType, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *OrderByDto) IsPropertySet() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given SortedByType and assigns it to the Property field.
func (o *OrderByDto) SetProperty(v SortedByType) {
	o.Property = &v
}

func (o OrderByDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderByDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAsc) {
		toSerialize["is_asc"] = o.IsAsc
	}
	if !IsNil(o.Property) {
		toSerialize["property"] = o.Property
	}
	return toSerialize, nil
}

type NullableOrderByDto struct {
	value *OrderByDto
	isSet bool
}

func (v NullableOrderByDto) Get() *OrderByDto {
	return v.value
}

func (v *NullableOrderByDto) Set(val *OrderByDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderByDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderByDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderByDto(val *OrderByDto) *NullableOrderByDto {
	return &NullableOrderByDto{value: val, isSet: true}
}

func (v NullableOrderByDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderByDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

