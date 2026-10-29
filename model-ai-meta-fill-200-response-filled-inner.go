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

// checks if the AiMetaFill200ResponseFilledInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiMetaFill200ResponseFilledInner{}

// AiMetaFill200ResponseFilledInner struct for AiMetaFill200ResponseFilledInner
type AiMetaFill200ResponseFilledInner struct {
	// Absent for a custom field.
	FieldId *float32 `json:"fieldId,omitempty"`
	Name string `json:"name"`
	Value AiMetaFill200ResponseFilledInnerValue `json:"value"`
}

type _AiMetaFill200ResponseFilledInner AiMetaFill200ResponseFilledInner

// NewAiMetaFill200ResponseFilledInner instantiates a new AiMetaFill200ResponseFilledInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiMetaFill200ResponseFilledInner(name string, value AiMetaFill200ResponseFilledInnerValue) *AiMetaFill200ResponseFilledInner {
	this := AiMetaFill200ResponseFilledInner{}
	this.Name = name
	this.Value = value
	return &this
}

// NewAiMetaFill200ResponseFilledInnerWithDefaults instantiates a new AiMetaFill200ResponseFilledInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiMetaFill200ResponseFilledInnerWithDefaults() *AiMetaFill200ResponseFilledInner {
	this := AiMetaFill200ResponseFilledInner{}
	return &this
}

// GetFieldId returns the FieldId field value if set, zero value otherwise.
func (o *AiMetaFill200ResponseFilledInner) GetFieldId() float32 {
	if o == nil || IsNil(o.FieldId) {
		var ret float32
		return ret
	}
	return *o.FieldId
}

// GetFieldIdOk returns a tuple with the FieldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMetaFill200ResponseFilledInner) GetFieldIdOk() (*float32, bool) {
	if o == nil || IsNil(o.FieldId) {
		return nil, false
	}
	return o.FieldId, true
}

// HasFieldId returns a boolean if a field has been set.
func (o *AiMetaFill200ResponseFilledInner) IsFieldIdSet() bool {
	if o != nil && !IsNil(o.FieldId) {
		return true
	}

	return false
}

// SetFieldId gets a reference to the given float32 and assigns it to the FieldId field.
func (o *AiMetaFill200ResponseFilledInner) SetFieldId(v float32) {
	o.FieldId = &v
}

// GetName returns the Name field value
func (o *AiMetaFill200ResponseFilledInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiMetaFill200ResponseFilledInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiMetaFill200ResponseFilledInner) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *AiMetaFill200ResponseFilledInner) GetValue() AiMetaFill200ResponseFilledInnerValue {
	if o == nil {
		var ret AiMetaFill200ResponseFilledInnerValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AiMetaFill200ResponseFilledInner) GetValueOk() (*AiMetaFill200ResponseFilledInnerValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AiMetaFill200ResponseFilledInner) SetValue(v AiMetaFill200ResponseFilledInnerValue) {
	o.Value = v
}

func (o AiMetaFill200ResponseFilledInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiMetaFill200ResponseFilledInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FieldId) {
		toSerialize["fieldId"] = o.FieldId
	}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *AiMetaFill200ResponseFilledInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"value",
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

	varAiMetaFill200ResponseFilledInner := _AiMetaFill200ResponseFilledInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiMetaFill200ResponseFilledInner)

	if err != nil {
		return err
	}

	*o = AiMetaFill200ResponseFilledInner(varAiMetaFill200ResponseFilledInner)

	return err
}

type NullableAiMetaFill200ResponseFilledInner struct {
	value *AiMetaFill200ResponseFilledInner
	isSet bool
}

func (v NullableAiMetaFill200ResponseFilledInner) Get() *AiMetaFill200ResponseFilledInner {
	return v.value
}

func (v *NullableAiMetaFill200ResponseFilledInner) Set(val *AiMetaFill200ResponseFilledInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAiMetaFill200ResponseFilledInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAiMetaFill200ResponseFilledInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiMetaFill200ResponseFilledInner(val *AiMetaFill200ResponseFilledInner) *NullableAiMetaFill200ResponseFilledInner {
	return &NullableAiMetaFill200ResponseFilledInner{value: val, isSet: true}
}

func (v NullableAiMetaFill200ResponseFilledInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiMetaFill200ResponseFilledInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

