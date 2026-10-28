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

// checks if the MetadataFieldOptionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataFieldOptionDto{}

// MetadataFieldOptionDto The metadata field choice option.
type MetadataFieldOptionDto struct {
	// The option ID.
	Id *string `json:"id,omitempty"`
	// The option value.
	Value NullableString `json:"value,omitempty"`
}

// NewMetadataFieldOptionDto instantiates a new MetadataFieldOptionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataFieldOptionDto() *MetadataFieldOptionDto {
	this := MetadataFieldOptionDto{}
	return &this
}

// NewMetadataFieldOptionDtoWithDefaults instantiates a new MetadataFieldOptionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataFieldOptionDtoWithDefaults() *MetadataFieldOptionDto {
	this := MetadataFieldOptionDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetadataFieldOptionDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataFieldOptionDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetadataFieldOptionDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MetadataFieldOptionDto) SetId(v string) {
	o.Id = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataFieldOptionDto) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataFieldOptionDto) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *MetadataFieldOptionDto) IsValueSet() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *MetadataFieldOptionDto) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *MetadataFieldOptionDto) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *MetadataFieldOptionDto) UnsetValue() {
	o.Value.Unset()
}

func (o MetadataFieldOptionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataFieldOptionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableMetadataFieldOptionDto struct {
	value *MetadataFieldOptionDto
	isSet bool
}

func (v NullableMetadataFieldOptionDto) Get() *MetadataFieldOptionDto {
	return v.value
}

func (v *NullableMetadataFieldOptionDto) Set(val *MetadataFieldOptionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataFieldOptionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataFieldOptionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataFieldOptionDto(val *MetadataFieldOptionDto) *NullableMetadataFieldOptionDto {
	return &NullableMetadataFieldOptionDto{value: val, isSet: true}
}

func (v NullableMetadataFieldOptionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataFieldOptionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

