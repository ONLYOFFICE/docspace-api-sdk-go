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

// checks if the ProductAdministratorDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAdministratorDto{}

// ProductAdministratorDto The product administrator parameters.
type ProductAdministratorDto struct {
	// The product ID.
	ProductId string `json:"productId"`
	// The user unique identifier.
	UserId string `json:"userId"`
	// Indicates whether the user has administrator privileges for the product.
	Administrator bool `json:"administrator"`
}

type _ProductAdministratorDto ProductAdministratorDto

// NewProductAdministratorDto instantiates a new ProductAdministratorDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAdministratorDto(productId string, userId string, administrator bool) *ProductAdministratorDto {
	this := ProductAdministratorDto{}
	this.ProductId = productId
	this.UserId = userId
	this.Administrator = administrator
	return &this
}

// NewProductAdministratorDtoWithDefaults instantiates a new ProductAdministratorDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAdministratorDtoWithDefaults() *ProductAdministratorDto {
	this := ProductAdministratorDto{}
	return &this
}

// GetProductId returns the ProductId field value
func (o *ProductAdministratorDto) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *ProductAdministratorDto) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *ProductAdministratorDto) SetProductId(v string) {
	o.ProductId = v
}

// GetUserId returns the UserId field value
func (o *ProductAdministratorDto) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ProductAdministratorDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ProductAdministratorDto) SetUserId(v string) {
	o.UserId = v
}

// GetAdministrator returns the Administrator field value
func (o *ProductAdministratorDto) GetAdministrator() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Administrator
}

// GetAdministratorOk returns a tuple with the Administrator field value
// and a boolean to check if the value has been set.
func (o *ProductAdministratorDto) GetAdministratorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Administrator, true
}

// SetAdministrator sets field value
func (o *ProductAdministratorDto) SetAdministrator(v bool) {
	o.Administrator = v
}

func (o ProductAdministratorDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAdministratorDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productId"] = o.ProductId
	toSerialize["userId"] = o.UserId
	toSerialize["administrator"] = o.Administrator
	return toSerialize, nil
}

func (o *ProductAdministratorDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"productId",
		"userId",
		"administrator",
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

	varProductAdministratorDto := _ProductAdministratorDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varProductAdministratorDto)

	if err != nil {
		return err
	}

	*o = ProductAdministratorDto(varProductAdministratorDto)

	return err
}

type NullableProductAdministratorDto struct {
	value *ProductAdministratorDto
	isSet bool
}

func (v NullableProductAdministratorDto) Get() *ProductAdministratorDto {
	return v.value
}

func (v *NullableProductAdministratorDto) Set(val *ProductAdministratorDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAdministratorDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAdministratorDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAdministratorDto(val *ProductAdministratorDto) *NullableProductAdministratorDto {
	return &NullableProductAdministratorDto{value: val, isSet: true}
}

func (v NullableProductAdministratorDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAdministratorDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

