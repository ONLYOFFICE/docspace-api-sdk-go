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

// checks if the CustomFieldRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFieldRequest{}

// CustomFieldRequest The parameters of a custom text field.
type CustomFieldRequest struct {
	// The field name. The names are matched without regard to case.
	Name NullableString `json:"name"`
	// The field value. Null or empty removes the field from the entry.
	Value NullableString `json:"value,omitempty"`
}

type _CustomFieldRequest CustomFieldRequest

// NewCustomFieldRequest instantiates a new CustomFieldRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFieldRequest(name NullableString) *CustomFieldRequest {
	this := CustomFieldRequest{}
	this.Name = name
	return &this
}

// NewCustomFieldRequestWithDefaults instantiates a new CustomFieldRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldRequestWithDefaults() *CustomFieldRequest {
	this := CustomFieldRequest{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustomFieldRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFieldRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *CustomFieldRequest) SetName(v string) {
	o.Name.Set(&v)
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFieldRequest) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFieldRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *CustomFieldRequest) IsValueSet() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *CustomFieldRequest) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *CustomFieldRequest) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *CustomFieldRequest) UnsetValue() {
	o.Value.Unset()
}

func (o CustomFieldRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFieldRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name.Get()
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

func (o *CustomFieldRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCustomFieldRequest := _CustomFieldRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCustomFieldRequest)

	if err != nil {
		return err
	}

	*o = CustomFieldRequest(varCustomFieldRequest)

	return err
}

type NullableCustomFieldRequest struct {
	value *CustomFieldRequest
	isSet bool
}

func (v NullableCustomFieldRequest) Get() *CustomFieldRequest {
	return v.value
}

func (v *NullableCustomFieldRequest) Set(val *CustomFieldRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldRequest(val *CustomFieldRequest) *NullableCustomFieldRequest {
	return &NullableCustomFieldRequest{value: val, isSet: true}
}

func (v NullableCustomFieldRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

