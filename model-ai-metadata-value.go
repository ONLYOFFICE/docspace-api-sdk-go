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

// checks if the AiMetadataValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiMetadataValue{}

// AiMetadataValue A metadata value: exactly one member, by the field's type.
type AiMetadataValue struct {
	StringValue *string `json:"stringValue,omitempty"`
	NumberValue *float32 `json:"numberValue,omitempty"`
	// ISO 8601 date-time, UTC midnight of the date.
	DateValue *string `json:"dateValue,omitempty"`
	OptionIds []string `json:"optionIds,omitempty"`
}

// NewAiMetadataValue instantiates a new AiMetadataValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiMetadataValue() *AiMetadataValue {
	this := AiMetadataValue{}
	return &this
}

// NewAiMetadataValueWithDefaults instantiates a new AiMetadataValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiMetadataValueWithDefaults() *AiMetadataValue {
	this := AiMetadataValue{}
	return &this
}

// GetStringValue returns the StringValue field value if set, zero value otherwise.
func (o *AiMetadataValue) GetStringValue() string {
	if o == nil || IsNil(o.StringValue) {
		var ret string
		return ret
	}
	return *o.StringValue
}

// GetStringValueOk returns a tuple with the StringValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMetadataValue) GetStringValueOk() (*string, bool) {
	if o == nil || IsNil(o.StringValue) {
		return nil, false
	}
	return o.StringValue, true
}

// HasStringValue returns a boolean if a field has been set.
func (o *AiMetadataValue) IsStringValueSet() bool {
	if o != nil && !IsNil(o.StringValue) {
		return true
	}

	return false
}

// SetStringValue gets a reference to the given string and assigns it to the StringValue field.
func (o *AiMetadataValue) SetStringValue(v string) {
	o.StringValue = &v
}

// GetNumberValue returns the NumberValue field value if set, zero value otherwise.
func (o *AiMetadataValue) GetNumberValue() float32 {
	if o == nil || IsNil(o.NumberValue) {
		var ret float32
		return ret
	}
	return *o.NumberValue
}

// GetNumberValueOk returns a tuple with the NumberValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMetadataValue) GetNumberValueOk() (*float32, bool) {
	if o == nil || IsNil(o.NumberValue) {
		return nil, false
	}
	return o.NumberValue, true
}

// HasNumberValue returns a boolean if a field has been set.
func (o *AiMetadataValue) IsNumberValueSet() bool {
	if o != nil && !IsNil(o.NumberValue) {
		return true
	}

	return false
}

// SetNumberValue gets a reference to the given float32 and assigns it to the NumberValue field.
func (o *AiMetadataValue) SetNumberValue(v float32) {
	o.NumberValue = &v
}

// GetDateValue returns the DateValue field value if set, zero value otherwise.
func (o *AiMetadataValue) GetDateValue() string {
	if o == nil || IsNil(o.DateValue) {
		var ret string
		return ret
	}
	return *o.DateValue
}

// GetDateValueOk returns a tuple with the DateValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMetadataValue) GetDateValueOk() (*string, bool) {
	if o == nil || IsNil(o.DateValue) {
		return nil, false
	}
	return o.DateValue, true
}

// HasDateValue returns a boolean if a field has been set.
func (o *AiMetadataValue) IsDateValueSet() bool {
	if o != nil && !IsNil(o.DateValue) {
		return true
	}

	return false
}

// SetDateValue gets a reference to the given string and assigns it to the DateValue field.
func (o *AiMetadataValue) SetDateValue(v string) {
	o.DateValue = &v
}

// GetOptionIds returns the OptionIds field value if set, zero value otherwise.
func (o *AiMetadataValue) GetOptionIds() []string {
	if o == nil || IsNil(o.OptionIds) {
		var ret []string
		return ret
	}
	return o.OptionIds
}

// GetOptionIdsOk returns a tuple with the OptionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMetadataValue) GetOptionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OptionIds) {
		return nil, false
	}
	return o.OptionIds, true
}

// HasOptionIds returns a boolean if a field has been set.
func (o *AiMetadataValue) IsOptionIdsSet() bool {
	if o != nil && !IsNil(o.OptionIds) {
		return true
	}

	return false
}

// SetOptionIds gets a reference to the given []string and assigns it to the OptionIds field.
func (o *AiMetadataValue) SetOptionIds(v []string) {
	o.OptionIds = v
}

func (o AiMetadataValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiMetadataValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StringValue) {
		toSerialize["stringValue"] = o.StringValue
	}
	if !IsNil(o.NumberValue) {
		toSerialize["numberValue"] = o.NumberValue
	}
	if !IsNil(o.DateValue) {
		toSerialize["dateValue"] = o.DateValue
	}
	if !IsNil(o.OptionIds) {
		toSerialize["optionIds"] = o.OptionIds
	}
	return toSerialize, nil
}

type NullableAiMetadataValue struct {
	value *AiMetadataValue
	isSet bool
}

func (v NullableAiMetadataValue) Get() *AiMetadataValue {
	return v.value
}

func (v *NullableAiMetadataValue) Set(val *AiMetadataValue) {
	v.value = val
	v.isSet = true
}

func (v NullableAiMetadataValue) IsSet() bool {
	return v.isSet
}

func (v *NullableAiMetadataValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiMetadataValue(val *AiMetadataValue) *NullableAiMetadataValue {
	return &NullableAiMetadataValue{value: val, isSet: true}
}

func (v NullableAiMetadataValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiMetadataValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

