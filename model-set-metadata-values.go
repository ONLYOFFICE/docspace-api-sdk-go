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

// checks if the SetMetadataValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetMetadataValues{}

// SetMetadataValues The parameters for setting metadata values.
type SetMetadataValues struct {
	// The metadata field values. An empty value clears the field.
	Values []MetadataValueRequest `json:"values"`
}

type _SetMetadataValues SetMetadataValues

// NewSetMetadataValues instantiates a new SetMetadataValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetMetadataValues(values []MetadataValueRequest) *SetMetadataValues {
	this := SetMetadataValues{}
	this.Values = values
	return &this
}

// NewSetMetadataValuesWithDefaults instantiates a new SetMetadataValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetMetadataValuesWithDefaults() *SetMetadataValues {
	this := SetMetadataValues{}
	return &this
}

// GetValues returns the Values field value
func (o *SetMetadataValues) GetValues() []MetadataValueRequest {
	if o == nil {
		var ret []MetadataValueRequest
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *SetMetadataValues) GetValuesOk() ([]MetadataValueRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *SetMetadataValues) SetValues(v []MetadataValueRequest) {
	o.Values = v
}

func (o SetMetadataValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetMetadataValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

func (o *SetMetadataValues) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"values",
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

	varSetMetadataValues := _SetMetadataValues{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSetMetadataValues)

	if err != nil {
		return err
	}

	*o = SetMetadataValues(varSetMetadataValues)

	return err
}

type NullableSetMetadataValues struct {
	value *SetMetadataValues
	isSet bool
}

func (v NullableSetMetadataValues) Get() *SetMetadataValues {
	return v.value
}

func (v *NullableSetMetadataValues) Set(val *SetMetadataValues) {
	v.value = val
	v.isSet = true
}

func (v NullableSetMetadataValues) IsSet() bool {
	return v.isSet
}

func (v *NullableSetMetadataValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetMetadataValues(val *SetMetadataValues) *NullableSetMetadataValues {
	return &NullableSetMetadataValues{value: val, isSet: true}
}

func (v NullableSetMetadataValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetMetadataValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

