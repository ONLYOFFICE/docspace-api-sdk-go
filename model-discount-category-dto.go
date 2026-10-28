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
	"time"
)

// checks if the DiscountCategoryDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscountCategoryDto{}

// DiscountCategoryDto Represents a discount category applied to the price.
type DiscountCategoryDto struct {
	// The discount category unique identifier.
	Id *int32 `json:"id,omitempty"`
	// The discount value.
	ValueDiscount *float64 `json:"valueDiscount,omitempty"`
	// The discount category description.
	Description NullableString `json:"description,omitempty"`
	// The date and time when the discount category was created.
	Created *time.Time `json:"created,omitempty"`
}

// NewDiscountCategoryDto instantiates a new DiscountCategoryDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscountCategoryDto() *DiscountCategoryDto {
	this := DiscountCategoryDto{}
	return &this
}

// NewDiscountCategoryDtoWithDefaults instantiates a new DiscountCategoryDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscountCategoryDtoWithDefaults() *DiscountCategoryDto {
	this := DiscountCategoryDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DiscountCategoryDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountCategoryDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DiscountCategoryDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DiscountCategoryDto) SetId(v int32) {
	o.Id = &v
}

// GetValueDiscount returns the ValueDiscount field value if set, zero value otherwise.
func (o *DiscountCategoryDto) GetValueDiscount() float64 {
	if o == nil || IsNil(o.ValueDiscount) {
		var ret float64
		return ret
	}
	return *o.ValueDiscount
}

// GetValueDiscountOk returns a tuple with the ValueDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountCategoryDto) GetValueDiscountOk() (*float64, bool) {
	if o == nil || IsNil(o.ValueDiscount) {
		return nil, false
	}
	return o.ValueDiscount, true
}

// HasValueDiscount returns a boolean if a field has been set.
func (o *DiscountCategoryDto) IsValueDiscountSet() bool {
	if o != nil && !IsNil(o.ValueDiscount) {
		return true
	}

	return false
}

// SetValueDiscount gets a reference to the given float64 and assigns it to the ValueDiscount field.
func (o *DiscountCategoryDto) SetValueDiscount(v float64) {
	o.ValueDiscount = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscountCategoryDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscountCategoryDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DiscountCategoryDto) IsDescriptionSet() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *DiscountCategoryDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *DiscountCategoryDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *DiscountCategoryDto) UnsetDescription() {
	o.Description.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *DiscountCategoryDto) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscountCategoryDto) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *DiscountCategoryDto) IsCreatedSet() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *DiscountCategoryDto) SetCreated(v time.Time) {
	o.Created = &v
}

func (o DiscountCategoryDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscountCategoryDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ValueDiscount) {
		toSerialize["valueDiscount"] = o.ValueDiscount
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	return toSerialize, nil
}

type NullableDiscountCategoryDto struct {
	value *DiscountCategoryDto
	isSet bool
}

func (v NullableDiscountCategoryDto) Get() *DiscountCategoryDto {
	return v.value
}

func (v *NullableDiscountCategoryDto) Set(val *DiscountCategoryDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscountCategoryDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscountCategoryDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscountCategoryDto(val *DiscountCategoryDto) *NullableDiscountCategoryDto {
	return &NullableDiscountCategoryDto{value: val, isSet: true}
}

func (v NullableDiscountCategoryDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscountCategoryDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

