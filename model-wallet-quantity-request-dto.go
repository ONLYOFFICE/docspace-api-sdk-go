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

// checks if the WalletQuantityRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletQuantityRequestDto{}

// WalletQuantityRequestDto The wallet service being bought or scheduled, and the way its quantity is applied.
type WalletQuantityRequestDto struct {
	// The wallet service and the number of units of it, as a single pair. The key is the `serviceName` of a service  from `GET api/2.0/portal/payment/walletservices`, and the value is read according to  `productQuantityType`: the units to add, or the total the service is to have in the next period. Minimum  quantities apply per service - disk storage starts at 100 units, the Docs Connect Dev Pack at 10, and the  administrators may not be fewer than the portal already has. Exactly one pair is accepted, and a null or zero  value cancels a change scheduled earlier rather than buying nothing.
	Quantity map[string]*int32 `json:"quantity"`
	// How the number in `quantity` is applied. `Add` buys the units straight away and charges them to the portal  wallet, while `Set` charges nothing now and records the quantity the service is to have from the next period.  Only these two are accepted here; `Sub` and `Renew` are refused with 400.
	ProductQuantityType *ProductQuantityType `json:"productQuantityType,omitempty"`
}

type _WalletQuantityRequestDto WalletQuantityRequestDto

// NewWalletQuantityRequestDto instantiates a new WalletQuantityRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletQuantityRequestDto(quantity map[string]*int32) *WalletQuantityRequestDto {
	this := WalletQuantityRequestDto{}
	this.Quantity = quantity
	return &this
}

// NewWalletQuantityRequestDtoWithDefaults instantiates a new WalletQuantityRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletQuantityRequestDtoWithDefaults() *WalletQuantityRequestDto {
	this := WalletQuantityRequestDto{}
	return &this
}

// GetQuantity returns the Quantity field value
func (o *WalletQuantityRequestDto) GetQuantity() map[string]*int32 {
	if o == nil {
		var ret map[string]*int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *WalletQuantityRequestDto) GetQuantityOk() (map[string]*int32, bool) {
	if o == nil {
		return map[string]*int32{}, false
	}
	return o.Quantity, true
}

// SetQuantity sets field value
func (o *WalletQuantityRequestDto) SetQuantity(v map[string]*int32) {
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
	toSerialize["quantity"] = o.Quantity
	if !IsNil(o.ProductQuantityType) {
		toSerialize["productQuantityType"] = o.ProductQuantityType
	}
	return toSerialize, nil
}

func (o *WalletQuantityRequestDto) UnmarshalJSON(data []byte) (err error) {
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

	varWalletQuantityRequestDto := _WalletQuantityRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varWalletQuantityRequestDto)

	if err != nil {
		return err
	}

	*o = WalletQuantityRequestDto(varWalletQuantityRequestDto)

	return err
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

