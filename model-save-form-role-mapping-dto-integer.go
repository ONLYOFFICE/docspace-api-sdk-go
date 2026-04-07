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

// checks if the SaveFormRoleMappingDtoInteger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaveFormRoleMappingDtoInteger{}

// SaveFormRoleMappingDtoInteger The parameters for saving form role mapping.
type SaveFormRoleMappingDtoInteger struct {
	// The form ID.
	FormId int32 `json:"formId"`
	// The collection of roles.
	Roles []FormRole `json:"roles"`
}

type _SaveFormRoleMappingDtoInteger SaveFormRoleMappingDtoInteger

// NewSaveFormRoleMappingDtoInteger instantiates a new SaveFormRoleMappingDtoInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaveFormRoleMappingDtoInteger(formId int32, roles []FormRole) *SaveFormRoleMappingDtoInteger {
	this := SaveFormRoleMappingDtoInteger{}
	this.FormId = formId
	this.Roles = roles
	return &this
}

// NewSaveFormRoleMappingDtoIntegerWithDefaults instantiates a new SaveFormRoleMappingDtoInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaveFormRoleMappingDtoIntegerWithDefaults() *SaveFormRoleMappingDtoInteger {
	this := SaveFormRoleMappingDtoInteger{}
	return &this
}

// GetFormId returns the FormId field value
func (o *SaveFormRoleMappingDtoInteger) GetFormId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FormId
}

// GetFormIdOk returns a tuple with the FormId field value
// and a boolean to check if the value has been set.
func (o *SaveFormRoleMappingDtoInteger) GetFormIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormId, true
}

// SetFormId sets field value
func (o *SaveFormRoleMappingDtoInteger) SetFormId(v int32) {
	o.FormId = v
}

// GetRoles returns the Roles field value
// If the value is explicit nil, the zero value for []FormRole will be returned
func (o *SaveFormRoleMappingDtoInteger) GetRoles() []FormRole {
	if o == nil {
		var ret []FormRole
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SaveFormRoleMappingDtoInteger) GetRolesOk() ([]FormRole, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *SaveFormRoleMappingDtoInteger) SetRoles(v []FormRole) {
	o.Roles = v
}

func (o SaveFormRoleMappingDtoInteger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaveFormRoleMappingDtoInteger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["formId"] = o.FormId
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

func (o *SaveFormRoleMappingDtoInteger) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"formId",
		"roles",
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

	varSaveFormRoleMappingDtoInteger := _SaveFormRoleMappingDtoInteger{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSaveFormRoleMappingDtoInteger)

	if err != nil {
		return err
	}

	*o = SaveFormRoleMappingDtoInteger(varSaveFormRoleMappingDtoInteger)

	return err
}

type NullableSaveFormRoleMappingDtoInteger struct {
	value *SaveFormRoleMappingDtoInteger
	isSet bool
}

func (v NullableSaveFormRoleMappingDtoInteger) Get() *SaveFormRoleMappingDtoInteger {
	return v.value
}

func (v *NullableSaveFormRoleMappingDtoInteger) Set(val *SaveFormRoleMappingDtoInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveFormRoleMappingDtoInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveFormRoleMappingDtoInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveFormRoleMappingDtoInteger(val *SaveFormRoleMappingDtoInteger) *NullableSaveFormRoleMappingDtoInteger {
	return &NullableSaveFormRoleMappingDtoInteger{value: val, isSet: true}
}

func (v NullableSaveFormRoleMappingDtoInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveFormRoleMappingDtoInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

