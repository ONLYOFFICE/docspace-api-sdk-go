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
	"bytes"
	"fmt"
)

// checks if the MetadataValueRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataValueRequest{}

// MetadataValueRequest The parameters of a metadata field value.
type MetadataValueRequest struct {
	// The field ID.
	FieldId int32 `json:"fieldId"`
	// The string value.
	StringValue NullableString `json:"stringValue,omitempty"`
	// The number value.
	NumberValue NullableInt64 `json:"numberValue,omitempty"`
	// The date value. A value without a time zone offset is treated as UTC, the same way the metadata filters treat their date bounds.
	DateValue NullableTime `json:"dateValue,omitempty"`
	// The selected choice option IDs.
	OptionIds []string `json:"optionIds,omitempty"`
}

type _MetadataValueRequest MetadataValueRequest

// NewMetadataValueRequest instantiates a new MetadataValueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataValueRequest(fieldId int32) *MetadataValueRequest {
	this := MetadataValueRequest{}
	this.FieldId = fieldId
	return &this
}

// NewMetadataValueRequestWithDefaults instantiates a new MetadataValueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataValueRequestWithDefaults() *MetadataValueRequest {
	this := MetadataValueRequest{}
	return &this
}

// GetFieldId returns the FieldId field value
func (o *MetadataValueRequest) GetFieldId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FieldId
}

// GetFieldIdOk returns a tuple with the FieldId field value
// and a boolean to check if the value has been set.
func (o *MetadataValueRequest) GetFieldIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldId, true
}

// SetFieldId sets field value
func (o *MetadataValueRequest) SetFieldId(v int32) {
	o.FieldId = v
}

// GetStringValue returns the StringValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataValueRequest) GetStringValue() string {
	if o == nil || IsNil(o.StringValue.Get()) {
		var ret string
		return ret
	}
	return *o.StringValue.Get()
}

// GetStringValueOk returns a tuple with the StringValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataValueRequest) GetStringValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StringValue.Get(), o.StringValue.IsSet()
}

// HasStringValue returns a boolean if a field has been set.
func (o *MetadataValueRequest) IsStringValueSet() bool {
	if o != nil && o.StringValue.IsSet() {
		return true
	}

	return false
}

// SetStringValue gets a reference to the given NullableString and assigns it to the StringValue field.
func (o *MetadataValueRequest) SetStringValue(v string) {
	o.StringValue.Set(&v)
}
// SetStringValueNil sets the value for StringValue to be an explicit nil
func (o *MetadataValueRequest) SetStringValueNil() {
	o.StringValue.Set(nil)
}

// UnsetStringValue ensures that no value is present for StringValue, not even an explicit nil
func (o *MetadataValueRequest) UnsetStringValue() {
	o.StringValue.Unset()
}

// GetNumberValue returns the NumberValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataValueRequest) GetNumberValue() int64 {
	if o == nil || IsNil(o.NumberValue.Get()) {
		var ret int64
		return ret
	}
	return *o.NumberValue.Get()
}

// GetNumberValueOk returns a tuple with the NumberValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataValueRequest) GetNumberValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumberValue.Get(), o.NumberValue.IsSet()
}

// HasNumberValue returns a boolean if a field has been set.
func (o *MetadataValueRequest) IsNumberValueSet() bool {
	if o != nil && o.NumberValue.IsSet() {
		return true
	}

	return false
}

// SetNumberValue gets a reference to the given NullableInt64 and assigns it to the NumberValue field.
func (o *MetadataValueRequest) SetNumberValue(v int64) {
	o.NumberValue.Set(&v)
}
// SetNumberValueNil sets the value for NumberValue to be an explicit nil
func (o *MetadataValueRequest) SetNumberValueNil() {
	o.NumberValue.Set(nil)
}

// UnsetNumberValue ensures that no value is present for NumberValue, not even an explicit nil
func (o *MetadataValueRequest) UnsetNumberValue() {
	o.NumberValue.Unset()
}

// GetDateValue returns the DateValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataValueRequest) GetDateValue() time.Time {
	if o == nil || IsNil(o.DateValue.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DateValue.Get()
}

// GetDateValueOk returns a tuple with the DateValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataValueRequest) GetDateValueOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateValue.Get(), o.DateValue.IsSet()
}

// HasDateValue returns a boolean if a field has been set.
func (o *MetadataValueRequest) IsDateValueSet() bool {
	if o != nil && o.DateValue.IsSet() {
		return true
	}

	return false
}

// SetDateValue gets a reference to the given NullableTime and assigns it to the DateValue field.
func (o *MetadataValueRequest) SetDateValue(v time.Time) {
	o.DateValue.Set(&v)
}
// SetDateValueNil sets the value for DateValue to be an explicit nil
func (o *MetadataValueRequest) SetDateValueNil() {
	o.DateValue.Set(nil)
}

// UnsetDateValue ensures that no value is present for DateValue, not even an explicit nil
func (o *MetadataValueRequest) UnsetDateValue() {
	o.DateValue.Unset()
}

// GetOptionIds returns the OptionIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataValueRequest) GetOptionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OptionIds
}

// GetOptionIdsOk returns a tuple with the OptionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataValueRequest) GetOptionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OptionIds) {
		return nil, false
	}
	return o.OptionIds, true
}

// HasOptionIds returns a boolean if a field has been set.
func (o *MetadataValueRequest) IsOptionIdsSet() bool {
	if o != nil && !IsNil(o.OptionIds) {
		return true
	}

	return false
}

// SetOptionIds gets a reference to the given []string and assigns it to the OptionIds field.
func (o *MetadataValueRequest) SetOptionIds(v []string) {
	o.OptionIds = v
}

func (o MetadataValueRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataValueRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fieldId"] = o.FieldId
	if o.StringValue.IsSet() {
		toSerialize["stringValue"] = o.StringValue.Get()
	}
	if o.NumberValue.IsSet() {
		toSerialize["numberValue"] = o.NumberValue.Get()
	}
	if o.DateValue.IsSet() {
		toSerialize["dateValue"] = o.DateValue.Get()
	}
	if o.OptionIds != nil {
		toSerialize["optionIds"] = o.OptionIds
	}
	return toSerialize, nil
}

func (o *MetadataValueRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fieldId",
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

	varMetadataValueRequest := _MetadataValueRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varMetadataValueRequest)

	if err != nil {
		return err
	}

	*o = MetadataValueRequest(varMetadataValueRequest)

	return err
}

type NullableMetadataValueRequest struct {
	value *MetadataValueRequest
	isSet bool
}

func (v NullableMetadataValueRequest) Get() *MetadataValueRequest {
	return v.value
}

func (v *NullableMetadataValueRequest) Set(val *MetadataValueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataValueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataValueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataValueRequest(val *MetadataValueRequest) *NullableMetadataValueRequest {
	return &NullableMetadataValueRequest{value: val, isSet: true}
}

func (v NullableMetadataValueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataValueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

