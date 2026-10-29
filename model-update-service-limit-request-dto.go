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

// checks if the UpdateServiceLimitRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateServiceLimitRequestDto{}

// UpdateServiceLimitRequestDto The new thresholds and state of an existing service limit.
type UpdateServiceLimitRequestDto struct {
	// The ID of the service limit, as returned in `id` by the operation that created it or by the reads of limits.
	Id int32 `json:"id"`
	// The most money the AI services may cost per period, in their currency. Zero blocks the AI services completely.
	AmountValue NullableFloat64 `json:"amountValue,omitempty"`
	// The largest quantity the AI services may consume per period, in their common unit. Zero blocks the AI services  completely.
	QuantityValue NullableInt32 `json:"quantityValue,omitempty"`
	// Whether the limit is in force: false switches it off at once, true switches it back on.
	Enabled NullableBool `json:"enabled,omitempty"`
}

type _UpdateServiceLimitRequestDto UpdateServiceLimitRequestDto

// NewUpdateServiceLimitRequestDto instantiates a new UpdateServiceLimitRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateServiceLimitRequestDto(id int32) *UpdateServiceLimitRequestDto {
	this := UpdateServiceLimitRequestDto{}
	this.Id = id
	return &this
}

// NewUpdateServiceLimitRequestDtoWithDefaults instantiates a new UpdateServiceLimitRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServiceLimitRequestDtoWithDefaults() *UpdateServiceLimitRequestDto {
	this := UpdateServiceLimitRequestDto{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateServiceLimitRequestDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateServiceLimitRequestDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateServiceLimitRequestDto) SetId(v int32) {
	o.Id = v
}

// GetAmountValue returns the AmountValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateServiceLimitRequestDto) GetAmountValue() float64 {
	if o == nil || IsNil(o.AmountValue.Get()) {
		var ret float64
		return ret
	}
	return *o.AmountValue.Get()
}

// GetAmountValueOk returns a tuple with the AmountValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateServiceLimitRequestDto) GetAmountValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountValue.Get(), o.AmountValue.IsSet()
}

// HasAmountValue returns a boolean if a field has been set.
func (o *UpdateServiceLimitRequestDto) IsAmountValueSet() bool {
	if o != nil && o.AmountValue.IsSet() {
		return true
	}

	return false
}

// SetAmountValue gets a reference to the given NullableFloat64 and assigns it to the AmountValue field.
func (o *UpdateServiceLimitRequestDto) SetAmountValue(v float64) {
	o.AmountValue.Set(&v)
}
// SetAmountValueNil sets the value for AmountValue to be an explicit nil
func (o *UpdateServiceLimitRequestDto) SetAmountValueNil() {
	o.AmountValue.Set(nil)
}

// UnsetAmountValue ensures that no value is present for AmountValue, not even an explicit nil
func (o *UpdateServiceLimitRequestDto) UnsetAmountValue() {
	o.AmountValue.Unset()
}

// GetQuantityValue returns the QuantityValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateServiceLimitRequestDto) GetQuantityValue() int32 {
	if o == nil || IsNil(o.QuantityValue.Get()) {
		var ret int32
		return ret
	}
	return *o.QuantityValue.Get()
}

// GetQuantityValueOk returns a tuple with the QuantityValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateServiceLimitRequestDto) GetQuantityValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.QuantityValue.Get(), o.QuantityValue.IsSet()
}

// HasQuantityValue returns a boolean if a field has been set.
func (o *UpdateServiceLimitRequestDto) IsQuantityValueSet() bool {
	if o != nil && o.QuantityValue.IsSet() {
		return true
	}

	return false
}

// SetQuantityValue gets a reference to the given NullableInt32 and assigns it to the QuantityValue field.
func (o *UpdateServiceLimitRequestDto) SetQuantityValue(v int32) {
	o.QuantityValue.Set(&v)
}
// SetQuantityValueNil sets the value for QuantityValue to be an explicit nil
func (o *UpdateServiceLimitRequestDto) SetQuantityValueNil() {
	o.QuantityValue.Set(nil)
}

// UnsetQuantityValue ensures that no value is present for QuantityValue, not even an explicit nil
func (o *UpdateServiceLimitRequestDto) UnsetQuantityValue() {
	o.QuantityValue.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateServiceLimitRequestDto) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateServiceLimitRequestDto) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateServiceLimitRequestDto) IsEnabledSet() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *UpdateServiceLimitRequestDto) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *UpdateServiceLimitRequestDto) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *UpdateServiceLimitRequestDto) UnsetEnabled() {
	o.Enabled.Unset()
}

func (o UpdateServiceLimitRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateServiceLimitRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.AmountValue.IsSet() {
		toSerialize["amountValue"] = o.AmountValue.Get()
	}
	if o.QuantityValue.IsSet() {
		toSerialize["quantityValue"] = o.QuantityValue.Get()
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	return toSerialize, nil
}

func (o *UpdateServiceLimitRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varUpdateServiceLimitRequestDto := _UpdateServiceLimitRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varUpdateServiceLimitRequestDto)

	if err != nil {
		return err
	}

	*o = UpdateServiceLimitRequestDto(varUpdateServiceLimitRequestDto)

	return err
}

type NullableUpdateServiceLimitRequestDto struct {
	value *UpdateServiceLimitRequestDto
	isSet bool
}

func (v NullableUpdateServiceLimitRequestDto) Get() *UpdateServiceLimitRequestDto {
	return v.value
}

func (v *NullableUpdateServiceLimitRequestDto) Set(val *UpdateServiceLimitRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateServiceLimitRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateServiceLimitRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateServiceLimitRequestDto(val *UpdateServiceLimitRequestDto) *NullableUpdateServiceLimitRequestDto {
	return &NullableUpdateServiceLimitRequestDto{value: val, isSet: true}
}

func (v NullableUpdateServiceLimitRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateServiceLimitRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

