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

// checks if the ImportableApiEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportableApiEntity{}

// ImportableApiEntity struct for ImportableApiEntity
type ImportableApiEntity struct {
	ShouldImport *bool `json:"shouldImport,omitempty"`
}

// NewImportableApiEntity instantiates a new ImportableApiEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportableApiEntity() *ImportableApiEntity {
	this := ImportableApiEntity{}
	return &this
}

// NewImportableApiEntityWithDefaults instantiates a new ImportableApiEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportableApiEntityWithDefaults() *ImportableApiEntity {
	this := ImportableApiEntity{}
	return &this
}

// GetShouldImport returns the ShouldImport field value if set, zero value otherwise.
func (o *ImportableApiEntity) GetShouldImport() bool {
	if o == nil || IsNil(o.ShouldImport) {
		var ret bool
		return ret
	}
	return *o.ShouldImport
}

// GetShouldImportOk returns a tuple with the ShouldImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportableApiEntity) GetShouldImportOk() (*bool, bool) {
	if o == nil || IsNil(o.ShouldImport) {
		return nil, false
	}
	return o.ShouldImport, true
}

// HasShouldImport returns a boolean if a field has been set.
func (o *ImportableApiEntity) IsShouldImportSet() bool {
	if o != nil && !IsNil(o.ShouldImport) {
		return true
	}

	return false
}

// SetShouldImport gets a reference to the given bool and assigns it to the ShouldImport field.
func (o *ImportableApiEntity) SetShouldImport(v bool) {
	o.ShouldImport = &v
}

func (o ImportableApiEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportableApiEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShouldImport) {
		toSerialize["shouldImport"] = o.ShouldImport
	}
	return toSerialize, nil
}

type NullableImportableApiEntity struct {
	value *ImportableApiEntity
	isSet bool
}

func (v NullableImportableApiEntity) Get() *ImportableApiEntity {
	return v.value
}

func (v *NullableImportableApiEntity) Set(val *ImportableApiEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableImportableApiEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableImportableApiEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportableApiEntity(val *ImportableApiEntity) *NullableImportableApiEntity {
	return &NullableImportableApiEntity{value: val, isSet: true}
}

func (v NullableImportableApiEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportableApiEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

