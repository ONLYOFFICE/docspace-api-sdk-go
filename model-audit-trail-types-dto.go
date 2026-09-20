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

// checks if the AuditTrailTypesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditTrailTypesDto{}

// AuditTrailTypesDto The vocabularies the audit and login-history filters accept, one array of names per dimension of an event.
type AuditTrailTypesDto struct {
	// Every action name the build can record, spelled as the `action` filter of  `GET api/2.0/security/audit/events/filter` and `GET api/2.0/security/audit/login/filter` expects it. It is  the whole vocabulary, not the actions this portal has recorded, and only a handful of the names are the  sign-in actions the login filter accepts.
	Actions []string `json:"actions,omitempty"`
	// The kinds of change an action can stand for, spelled as the `actionType` filter of  `GET api/2.0/security/audit/events/filter` expects it.
	ActionTypes []string `json:"actionTypes,omitempty"`
	// The products an action can belong to, spelled as the `productType` filter of  `GET api/2.0/security/audit/mappers` expects it. The audit trail itself cannot be filtered by product.
	ProductTypes []string `json:"productTypes,omitempty"`
	// The locations inside those products, spelled as the `moduleType` filter of  `GET api/2.0/security/audit/events/filter` and `GET api/2.0/security/audit/mappers` expects it.
	ModuleTypes []string `json:"moduleTypes,omitempty"`
	// The kinds of object an action can be applied to, spelled as the `entryType` filter of  `GET api/2.0/security/audit/events/filter` expects it.
	EntryTypes []string `json:"entryTypes,omitempty"`
}

// NewAuditTrailTypesDto instantiates a new AuditTrailTypesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditTrailTypesDto() *AuditTrailTypesDto {
	this := AuditTrailTypesDto{}
	return &this
}

// NewAuditTrailTypesDtoWithDefaults instantiates a new AuditTrailTypesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditTrailTypesDtoWithDefaults() *AuditTrailTypesDto {
	this := AuditTrailTypesDto{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditTrailTypesDto) GetActions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditTrailTypesDto) GetActionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *AuditTrailTypesDto) IsActionsSet() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []string and assigns it to the Actions field.
func (o *AuditTrailTypesDto) SetActions(v []string) {
	o.Actions = v
}

// GetActionTypes returns the ActionTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditTrailTypesDto) GetActionTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ActionTypes
}

// GetActionTypesOk returns a tuple with the ActionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditTrailTypesDto) GetActionTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ActionTypes) {
		return nil, false
	}
	return o.ActionTypes, true
}

// HasActionTypes returns a boolean if a field has been set.
func (o *AuditTrailTypesDto) IsActionTypesSet() bool {
	if o != nil && !IsNil(o.ActionTypes) {
		return true
	}

	return false
}

// SetActionTypes gets a reference to the given []string and assigns it to the ActionTypes field.
func (o *AuditTrailTypesDto) SetActionTypes(v []string) {
	o.ActionTypes = v
}

// GetProductTypes returns the ProductTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditTrailTypesDto) GetProductTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProductTypes
}

// GetProductTypesOk returns a tuple with the ProductTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditTrailTypesDto) GetProductTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductTypes) {
		return nil, false
	}
	return o.ProductTypes, true
}

// HasProductTypes returns a boolean if a field has been set.
func (o *AuditTrailTypesDto) IsProductTypesSet() bool {
	if o != nil && !IsNil(o.ProductTypes) {
		return true
	}

	return false
}

// SetProductTypes gets a reference to the given []string and assigns it to the ProductTypes field.
func (o *AuditTrailTypesDto) SetProductTypes(v []string) {
	o.ProductTypes = v
}

// GetModuleTypes returns the ModuleTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditTrailTypesDto) GetModuleTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ModuleTypes
}

// GetModuleTypesOk returns a tuple with the ModuleTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditTrailTypesDto) GetModuleTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ModuleTypes) {
		return nil, false
	}
	return o.ModuleTypes, true
}

// HasModuleTypes returns a boolean if a field has been set.
func (o *AuditTrailTypesDto) IsModuleTypesSet() bool {
	if o != nil && !IsNil(o.ModuleTypes) {
		return true
	}

	return false
}

// SetModuleTypes gets a reference to the given []string and assigns it to the ModuleTypes field.
func (o *AuditTrailTypesDto) SetModuleTypes(v []string) {
	o.ModuleTypes = v
}

// GetEntryTypes returns the EntryTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditTrailTypesDto) GetEntryTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EntryTypes
}

// GetEntryTypesOk returns a tuple with the EntryTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditTrailTypesDto) GetEntryTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.EntryTypes) {
		return nil, false
	}
	return o.EntryTypes, true
}

// HasEntryTypes returns a boolean if a field has been set.
func (o *AuditTrailTypesDto) IsEntryTypesSet() bool {
	if o != nil && !IsNil(o.EntryTypes) {
		return true
	}

	return false
}

// SetEntryTypes gets a reference to the given []string and assigns it to the EntryTypes field.
func (o *AuditTrailTypesDto) SetEntryTypes(v []string) {
	o.EntryTypes = v
}

func (o AuditTrailTypesDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditTrailTypesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.ActionTypes != nil {
		toSerialize["actionTypes"] = o.ActionTypes
	}
	if o.ProductTypes != nil {
		toSerialize["productTypes"] = o.ProductTypes
	}
	if o.ModuleTypes != nil {
		toSerialize["moduleTypes"] = o.ModuleTypes
	}
	if o.EntryTypes != nil {
		toSerialize["entryTypes"] = o.EntryTypes
	}
	return toSerialize, nil
}

type NullableAuditTrailTypesDto struct {
	value *AuditTrailTypesDto
	isSet bool
}

func (v NullableAuditTrailTypesDto) Get() *AuditTrailTypesDto {
	return v.value
}

func (v *NullableAuditTrailTypesDto) Set(val *AuditTrailTypesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditTrailTypesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditTrailTypesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditTrailTypesDto(val *AuditTrailTypesDto) *NullableAuditTrailTypesDto {
	return &NullableAuditTrailTypesDto{value: val, isSet: true}
}

func (v NullableAuditTrailTypesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditTrailTypesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

