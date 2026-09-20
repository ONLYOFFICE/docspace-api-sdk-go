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

// checks if the SaveFormRoleMappingDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaveFormRoleMappingDto{}

// SaveFormRoleMappingDto The people who are to fill in the roles of a PDF form.
type SaveFormRoleMappingDto struct {
	// The PDF form the roles belong to. This is the value the operation reads, rather than the identifier in its  route, and the two are to be sent the same.
	FormId int32 `json:"formId"`
	// The roles with the account taking each of them and the sequence number that decides the turn: the same number  means the roles may be filled in parallel, different ones make a queue. The whole set is replaced on every  call, and an empty set resets the filling.
	Roles []FormRole `json:"roles"`
}

type _SaveFormRoleMappingDto SaveFormRoleMappingDto

// NewSaveFormRoleMappingDto instantiates a new SaveFormRoleMappingDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaveFormRoleMappingDto(formId int32, roles []FormRole) *SaveFormRoleMappingDto {
	this := SaveFormRoleMappingDto{}
	this.FormId = formId
	this.Roles = roles
	return &this
}

// NewSaveFormRoleMappingDtoWithDefaults instantiates a new SaveFormRoleMappingDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaveFormRoleMappingDtoWithDefaults() *SaveFormRoleMappingDto {
	this := SaveFormRoleMappingDto{}
	return &this
}

// GetFormId returns the FormId field value
func (o *SaveFormRoleMappingDto) GetFormId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FormId
}

// GetFormIdOk returns a tuple with the FormId field value
// and a boolean to check if the value has been set.
func (o *SaveFormRoleMappingDto) GetFormIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormId, true
}

// SetFormId sets field value
func (o *SaveFormRoleMappingDto) SetFormId(v int32) {
	o.FormId = v
}

// GetRoles returns the Roles field value
// If the value is explicit nil, the zero value for []FormRole will be returned
func (o *SaveFormRoleMappingDto) GetRoles() []FormRole {
	if o == nil {
		var ret []FormRole
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SaveFormRoleMappingDto) GetRolesOk() ([]FormRole, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *SaveFormRoleMappingDto) SetRoles(v []FormRole) {
	o.Roles = v
}

func (o SaveFormRoleMappingDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaveFormRoleMappingDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["formId"] = o.FormId
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

func (o *SaveFormRoleMappingDto) UnmarshalJSON(data []byte) (err error) {
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

	varSaveFormRoleMappingDto := _SaveFormRoleMappingDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSaveFormRoleMappingDto)

	if err != nil {
		return err
	}

	*o = SaveFormRoleMappingDto(varSaveFormRoleMappingDto)

	return err
}

type NullableSaveFormRoleMappingDto struct {
	value *SaveFormRoleMappingDto
	isSet bool
}

func (v NullableSaveFormRoleMappingDto) Get() *SaveFormRoleMappingDto {
	return v.value
}

func (v *NullableSaveFormRoleMappingDto) Set(val *SaveFormRoleMappingDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveFormRoleMappingDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveFormRoleMappingDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveFormRoleMappingDto(val *SaveFormRoleMappingDto) *NullableSaveFormRoleMappingDto {
	return &NullableSaveFormRoleMappingDto{value: val, isSet: true}
}

func (v NullableSaveFormRoleMappingDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveFormRoleMappingDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

