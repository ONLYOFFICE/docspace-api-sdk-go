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

// checks if the AiToolsBulkResultErrorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiToolsBulkResultErrorsInner{}

// AiToolsBulkResultErrorsInner struct for AiToolsBulkResultErrorsInner
type AiToolsBulkResultErrorsInner struct {
	Name string `json:"name"`
	Error AiTErrorData `json:"error"`
}

type _AiToolsBulkResultErrorsInner AiToolsBulkResultErrorsInner

// NewAiToolsBulkResultErrorsInner instantiates a new AiToolsBulkResultErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiToolsBulkResultErrorsInner(name string, error_ AiTErrorData) *AiToolsBulkResultErrorsInner {
	this := AiToolsBulkResultErrorsInner{}
	this.Name = name
	this.Error = error_
	return &this
}

// NewAiToolsBulkResultErrorsInnerWithDefaults instantiates a new AiToolsBulkResultErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiToolsBulkResultErrorsInnerWithDefaults() *AiToolsBulkResultErrorsInner {
	this := AiToolsBulkResultErrorsInner{}
	return &this
}

// GetName returns the Name field value
func (o *AiToolsBulkResultErrorsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiToolsBulkResultErrorsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiToolsBulkResultErrorsInner) SetName(v string) {
	o.Name = v
}

// GetError returns the Error field value
func (o *AiToolsBulkResultErrorsInner) GetError() AiTErrorData {
	if o == nil {
		var ret AiTErrorData
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *AiToolsBulkResultErrorsInner) GetErrorOk() (*AiTErrorData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *AiToolsBulkResultErrorsInner) SetError(v AiTErrorData) {
	o.Error = v
}

func (o AiToolsBulkResultErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiToolsBulkResultErrorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *AiToolsBulkResultErrorsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"error",
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

	varAiToolsBulkResultErrorsInner := _AiToolsBulkResultErrorsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiToolsBulkResultErrorsInner)

	if err != nil {
		return err
	}

	*o = AiToolsBulkResultErrorsInner(varAiToolsBulkResultErrorsInner)

	return err
}

type NullableAiToolsBulkResultErrorsInner struct {
	value *AiToolsBulkResultErrorsInner
	isSet bool
}

func (v NullableAiToolsBulkResultErrorsInner) Get() *AiToolsBulkResultErrorsInner {
	return v.value
}

func (v *NullableAiToolsBulkResultErrorsInner) Set(val *AiToolsBulkResultErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAiToolsBulkResultErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAiToolsBulkResultErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiToolsBulkResultErrorsInner(val *AiToolsBulkResultErrorsInner) *NullableAiToolsBulkResultErrorsInner {
	return &NullableAiToolsBulkResultErrorsInner{value: val, isSet: true}
}

func (v NullableAiToolsBulkResultErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiToolsBulkResultErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

