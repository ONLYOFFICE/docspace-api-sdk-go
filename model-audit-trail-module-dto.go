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

// checks if the AuditTrailModuleDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditTrailModuleDto{}

// AuditTrailModuleDto The audit trail actions of one module.
type AuditTrailModuleDto struct {
	// The location inside the product, as the `moduleType` filter of `GET api/2.0/security/audit/events/filter`  spells it.
	ModuleType NullableString `json:"moduleType,omitempty"`
	// Every action this module can record. Each action appears under exactly one module, so this tree is where a  caller learns which module a given action belongs to.
	Actions []AuditTrailActionDto `json:"actions,omitempty"`
}

// NewAuditTrailModuleDto instantiates a new AuditTrailModuleDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditTrailModuleDto() *AuditTrailModuleDto {
	this := AuditTrailModuleDto{}
	return &this
}

// NewAuditTrailModuleDtoWithDefaults instantiates a new AuditTrailModuleDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditTrailModuleDtoWithDefaults() *AuditTrailModuleDto {
	this := AuditTrailModuleDto{}
	return &this
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditTrailModuleDto) GetModuleType() string {
	if o == nil || IsNil(o.ModuleType.Get()) {
		var ret string
		return ret
	}
	return *o.ModuleType.Get()
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditTrailModuleDto) GetModuleTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleType.Get(), o.ModuleType.IsSet()
}

// HasModuleType returns a boolean if a field has been set.
func (o *AuditTrailModuleDto) IsModuleTypeSet() bool {
	if o != nil && o.ModuleType.IsSet() {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given NullableString and assigns it to the ModuleType field.
func (o *AuditTrailModuleDto) SetModuleType(v string) {
	o.ModuleType.Set(&v)
}
// SetModuleTypeNil sets the value for ModuleType to be an explicit nil
func (o *AuditTrailModuleDto) SetModuleTypeNil() {
	o.ModuleType.Set(nil)
}

// UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
func (o *AuditTrailModuleDto) UnsetModuleType() {
	o.ModuleType.Unset()
}

// GetActions returns the Actions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditTrailModuleDto) GetActions() []AuditTrailActionDto {
	if o == nil {
		var ret []AuditTrailActionDto
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditTrailModuleDto) GetActionsOk() ([]AuditTrailActionDto, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *AuditTrailModuleDto) IsActionsSet() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []AuditTrailActionDto and assigns it to the Actions field.
func (o *AuditTrailModuleDto) SetActions(v []AuditTrailActionDto) {
	o.Actions = v
}

func (o AuditTrailModuleDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditTrailModuleDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ModuleType.IsSet() {
		toSerialize["moduleType"] = o.ModuleType.Get()
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	return toSerialize, nil
}

type NullableAuditTrailModuleDto struct {
	value *AuditTrailModuleDto
	isSet bool
}

func (v NullableAuditTrailModuleDto) Get() *AuditTrailModuleDto {
	return v.value
}

func (v *NullableAuditTrailModuleDto) Set(val *AuditTrailModuleDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditTrailModuleDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditTrailModuleDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditTrailModuleDto(val *AuditTrailModuleDto) *NullableAuditTrailModuleDto {
	return &NullableAuditTrailModuleDto{value: val, isSet: true}
}

func (v NullableAuditTrailModuleDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditTrailModuleDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

