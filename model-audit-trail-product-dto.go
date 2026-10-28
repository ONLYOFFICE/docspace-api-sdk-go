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

// checks if the AuditTrailProductDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditTrailProductDto{}

// AuditTrailProductDto The audit trail actions of one product, grouped by module.
type AuditTrailProductDto struct {
	// The product this branch of the tree belongs to, as the `productType` filter of this operation spells it and  as `GET api/2.0/security/audit/types` lists it under `productTypes`.
	ProductType NullableString `json:"productType,omitempty"`
	// The locations inside the product. It is empty when `moduleType` was passed and this product has no module  of that name, which is why a product can come back with nothing under it.
	Modules []AuditTrailModuleDto `json:"modules,omitempty"`
}

// NewAuditTrailProductDto instantiates a new AuditTrailProductDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditTrailProductDto() *AuditTrailProductDto {
	this := AuditTrailProductDto{}
	return &this
}

// NewAuditTrailProductDtoWithDefaults instantiates a new AuditTrailProductDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditTrailProductDtoWithDefaults() *AuditTrailProductDto {
	this := AuditTrailProductDto{}
	return &this
}

// GetProductType returns the ProductType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditTrailProductDto) GetProductType() string {
	if o == nil || IsNil(o.ProductType.Get()) {
		var ret string
		return ret
	}
	return *o.ProductType.Get()
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditTrailProductDto) GetProductTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductType.Get(), o.ProductType.IsSet()
}

// HasProductType returns a boolean if a field has been set.
func (o *AuditTrailProductDto) IsProductTypeSet() bool {
	if o != nil && o.ProductType.IsSet() {
		return true
	}

	return false
}

// SetProductType gets a reference to the given NullableString and assigns it to the ProductType field.
func (o *AuditTrailProductDto) SetProductType(v string) {
	o.ProductType.Set(&v)
}
// SetProductTypeNil sets the value for ProductType to be an explicit nil
func (o *AuditTrailProductDto) SetProductTypeNil() {
	o.ProductType.Set(nil)
}

// UnsetProductType ensures that no value is present for ProductType, not even an explicit nil
func (o *AuditTrailProductDto) UnsetProductType() {
	o.ProductType.Unset()
}

// GetModules returns the Modules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditTrailProductDto) GetModules() []AuditTrailModuleDto {
	if o == nil {
		var ret []AuditTrailModuleDto
		return ret
	}
	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditTrailProductDto) GetModulesOk() ([]AuditTrailModuleDto, bool) {
	if o == nil || IsNil(o.Modules) {
		return nil, false
	}
	return o.Modules, true
}

// HasModules returns a boolean if a field has been set.
func (o *AuditTrailProductDto) IsModulesSet() bool {
	if o != nil && !IsNil(o.Modules) {
		return true
	}

	return false
}

// SetModules gets a reference to the given []AuditTrailModuleDto and assigns it to the Modules field.
func (o *AuditTrailProductDto) SetModules(v []AuditTrailModuleDto) {
	o.Modules = v
}

func (o AuditTrailProductDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditTrailProductDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ProductType.IsSet() {
		toSerialize["productType"] = o.ProductType.Get()
	}
	if o.Modules != nil {
		toSerialize["modules"] = o.Modules
	}
	return toSerialize, nil
}

type NullableAuditTrailProductDto struct {
	value *AuditTrailProductDto
	isSet bool
}

func (v NullableAuditTrailProductDto) Get() *AuditTrailProductDto {
	return v.value
}

func (v *NullableAuditTrailProductDto) Set(val *AuditTrailProductDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditTrailProductDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditTrailProductDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditTrailProductDto(val *AuditTrailProductDto) *NullableAuditTrailProductDto {
	return &NullableAuditTrailProductDto{value: val, isSet: true}
}

func (v NullableAuditTrailProductDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditTrailProductDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

