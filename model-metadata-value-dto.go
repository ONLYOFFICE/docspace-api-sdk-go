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

// checks if the MetadataValueDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataValueDto{}

// MetadataValueDto The value of a metadata field on an entry. Exactly one of the value properties is set, the one matching the field type:  `stringValue` for a string field, `numberValue` for a number field, `dateValue` for a date field,  `optionIds` for a single or multiple choice field.
type MetadataValueDto struct {
	// The string value.
	StringValue NullableString `json:"stringValue,omitempty"`
	// The number value.
	NumberValue NullableInt64 `json:"numberValue,omitempty"`
	// The date value.
	DateValue *ApiDateTime `json:"dateValue,omitempty"`
	// The selected choice option IDs.
	OptionIds []string `json:"optionIds,omitempty"`
}

// NewMetadataValueDto instantiates a new MetadataValueDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataValueDto() *MetadataValueDto {
	this := MetadataValueDto{}
	return &this
}

// NewMetadataValueDtoWithDefaults instantiates a new MetadataValueDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataValueDtoWithDefaults() *MetadataValueDto {
	this := MetadataValueDto{}
	return &this
}

// GetStringValue returns the StringValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataValueDto) GetStringValue() string {
	if o == nil || IsNil(o.StringValue.Get()) {
		var ret string
		return ret
	}
	return *o.StringValue.Get()
}

// GetStringValueOk returns a tuple with the StringValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataValueDto) GetStringValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StringValue.Get(), o.StringValue.IsSet()
}

// HasStringValue returns a boolean if a field has been set.
func (o *MetadataValueDto) IsStringValueSet() bool {
	if o != nil && o.StringValue.IsSet() {
		return true
	}

	return false
}

// SetStringValue gets a reference to the given NullableString and assigns it to the StringValue field.
func (o *MetadataValueDto) SetStringValue(v string) {
	o.StringValue.Set(&v)
}
// SetStringValueNil sets the value for StringValue to be an explicit nil
func (o *MetadataValueDto) SetStringValueNil() {
	o.StringValue.Set(nil)
}

// UnsetStringValue ensures that no value is present for StringValue, not even an explicit nil
func (o *MetadataValueDto) UnsetStringValue() {
	o.StringValue.Unset()
}

// GetNumberValue returns the NumberValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataValueDto) GetNumberValue() int64 {
	if o == nil || IsNil(o.NumberValue.Get()) {
		var ret int64
		return ret
	}
	return *o.NumberValue.Get()
}

// GetNumberValueOk returns a tuple with the NumberValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataValueDto) GetNumberValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumberValue.Get(), o.NumberValue.IsSet()
}

// HasNumberValue returns a boolean if a field has been set.
func (o *MetadataValueDto) IsNumberValueSet() bool {
	if o != nil && o.NumberValue.IsSet() {
		return true
	}

	return false
}

// SetNumberValue gets a reference to the given NullableInt64 and assigns it to the NumberValue field.
func (o *MetadataValueDto) SetNumberValue(v int64) {
	o.NumberValue.Set(&v)
}
// SetNumberValueNil sets the value for NumberValue to be an explicit nil
func (o *MetadataValueDto) SetNumberValueNil() {
	o.NumberValue.Set(nil)
}

// UnsetNumberValue ensures that no value is present for NumberValue, not even an explicit nil
func (o *MetadataValueDto) UnsetNumberValue() {
	o.NumberValue.Unset()
}

// GetDateValue returns the DateValue field value if set, zero value otherwise.
func (o *MetadataValueDto) GetDateValue() ApiDateTime {
	if o == nil || IsNil(o.DateValue) {
		var ret ApiDateTime
		return ret
	}
	return *o.DateValue
}

// GetDateValueOk returns a tuple with the DateValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataValueDto) GetDateValueOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.DateValue) {
		return nil, false
	}
	return o.DateValue, true
}

// HasDateValue returns a boolean if a field has been set.
func (o *MetadataValueDto) IsDateValueSet() bool {
	if o != nil && !IsNil(o.DateValue) {
		return true
	}

	return false
}

// SetDateValue gets a reference to the given ApiDateTime and assigns it to the DateValue field.
func (o *MetadataValueDto) SetDateValue(v ApiDateTime) {
	o.DateValue = &v
}

// GetOptionIds returns the OptionIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataValueDto) GetOptionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OptionIds
}

// GetOptionIdsOk returns a tuple with the OptionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataValueDto) GetOptionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OptionIds) {
		return nil, false
	}
	return o.OptionIds, true
}

// HasOptionIds returns a boolean if a field has been set.
func (o *MetadataValueDto) IsOptionIdsSet() bool {
	if o != nil && !IsNil(o.OptionIds) {
		return true
	}

	return false
}

// SetOptionIds gets a reference to the given []string and assigns it to the OptionIds field.
func (o *MetadataValueDto) SetOptionIds(v []string) {
	o.OptionIds = v
}

func (o MetadataValueDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataValueDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.StringValue.IsSet() {
		toSerialize["stringValue"] = o.StringValue.Get()
	}
	if o.NumberValue.IsSet() {
		toSerialize["numberValue"] = o.NumberValue.Get()
	}
	if !IsNil(o.DateValue) {
		toSerialize["dateValue"] = o.DateValue
	}
	if o.OptionIds != nil {
		toSerialize["optionIds"] = o.OptionIds
	}
	return toSerialize, nil
}

type NullableMetadataValueDto struct {
	value *MetadataValueDto
	isSet bool
}

func (v NullableMetadataValueDto) Get() *MetadataValueDto {
	return v.value
}

func (v *NullableMetadataValueDto) Set(val *MetadataValueDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataValueDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataValueDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataValueDto(val *MetadataValueDto) *NullableMetadataValueDto {
	return &NullableMetadataValueDto{value: val, isSet: true}
}

func (v NullableMetadataValueDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataValueDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

