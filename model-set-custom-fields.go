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

// checks if the SetCustomFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetCustomFields{}

// SetCustomFields The parameters for setting custom fields.
type SetCustomFields struct {
	// The custom fields to set on the entry. A listed field gets the value, a null or empty value removes the field  from the entry, the fields not listed are left alone. A name the portal has not seen yet creates the field.
	Fields []CustomFieldRequest `json:"fields"`
}

type _SetCustomFields SetCustomFields

// NewSetCustomFields instantiates a new SetCustomFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetCustomFields(fields []CustomFieldRequest) *SetCustomFields {
	this := SetCustomFields{}
	this.Fields = fields
	return &this
}

// NewSetCustomFieldsWithDefaults instantiates a new SetCustomFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetCustomFieldsWithDefaults() *SetCustomFields {
	this := SetCustomFields{}
	return &this
}

// GetFields returns the Fields field value
func (o *SetCustomFields) GetFields() []CustomFieldRequest {
	if o == nil {
		var ret []CustomFieldRequest
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *SetCustomFields) GetFieldsOk() ([]CustomFieldRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *SetCustomFields) SetFields(v []CustomFieldRequest) {
	o.Fields = v
}

func (o SetCustomFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetCustomFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *SetCustomFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fields",
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

	varSetCustomFields := _SetCustomFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSetCustomFields)

	if err != nil {
		return err
	}

	*o = SetCustomFields(varSetCustomFields)

	return err
}

type NullableSetCustomFields struct {
	value *SetCustomFields
	isSet bool
}

func (v NullableSetCustomFields) Get() *SetCustomFields {
	return v.value
}

func (v *NullableSetCustomFields) Set(val *SetCustomFields) {
	v.value = val
	v.isSet = true
}

func (v NullableSetCustomFields) IsSet() bool {
	return v.isSet
}

func (v *NullableSetCustomFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetCustomFields(val *SetCustomFields) *NullableSetCustomFields {
	return &NullableSetCustomFields{value: val, isSet: true}
}

func (v NullableSetCustomFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetCustomFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

