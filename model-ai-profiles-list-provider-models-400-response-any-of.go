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

// checks if the AiProfilesListProviderModels400ResponseAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiProfilesListProviderModels400ResponseAnyOf{}

// AiProfilesListProviderModels400ResponseAnyOf struct for AiProfilesListProviderModels400ResponseAnyOf
type AiProfilesListProviderModels400ResponseAnyOf struct {
	// The error message, ready to be shown to the caller.
	Error string `json:"error"`
	// Name of the request field that was missing or rejected.
	Field string `json:"field"`
}

type _AiProfilesListProviderModels400ResponseAnyOf AiProfilesListProviderModels400ResponseAnyOf

// NewAiProfilesListProviderModels400ResponseAnyOf instantiates a new AiProfilesListProviderModels400ResponseAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiProfilesListProviderModels400ResponseAnyOf(error_ string, field string) *AiProfilesListProviderModels400ResponseAnyOf {
	this := AiProfilesListProviderModels400ResponseAnyOf{}
	this.Error = error_
	this.Field = field
	return &this
}

// NewAiProfilesListProviderModels400ResponseAnyOfWithDefaults instantiates a new AiProfilesListProviderModels400ResponseAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiProfilesListProviderModels400ResponseAnyOfWithDefaults() *AiProfilesListProviderModels400ResponseAnyOf {
	this := AiProfilesListProviderModels400ResponseAnyOf{}
	return &this
}

// GetError returns the Error field value
func (o *AiProfilesListProviderModels400ResponseAnyOf) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *AiProfilesListProviderModels400ResponseAnyOf) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *AiProfilesListProviderModels400ResponseAnyOf) SetError(v string) {
	o.Error = v
}

// GetField returns the Field field value
func (o *AiProfilesListProviderModels400ResponseAnyOf) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *AiProfilesListProviderModels400ResponseAnyOf) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *AiProfilesListProviderModels400ResponseAnyOf) SetField(v string) {
	o.Field = v
}

func (o AiProfilesListProviderModels400ResponseAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiProfilesListProviderModels400ResponseAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	toSerialize["field"] = o.Field
	return toSerialize, nil
}

func (o *AiProfilesListProviderModels400ResponseAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error",
		"field",
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

	varAiProfilesListProviderModels400ResponseAnyOf := _AiProfilesListProviderModels400ResponseAnyOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiProfilesListProviderModels400ResponseAnyOf)

	if err != nil {
		return err
	}

	*o = AiProfilesListProviderModels400ResponseAnyOf(varAiProfilesListProviderModels400ResponseAnyOf)

	return err
}

type NullableAiProfilesListProviderModels400ResponseAnyOf struct {
	value *AiProfilesListProviderModels400ResponseAnyOf
	isSet bool
}

func (v NullableAiProfilesListProviderModels400ResponseAnyOf) Get() *AiProfilesListProviderModels400ResponseAnyOf {
	return v.value
}

func (v *NullableAiProfilesListProviderModels400ResponseAnyOf) Set(val *AiProfilesListProviderModels400ResponseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAiProfilesListProviderModels400ResponseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAiProfilesListProviderModels400ResponseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiProfilesListProviderModels400ResponseAnyOf(val *AiProfilesListProviderModels400ResponseAnyOf) *NullableAiProfilesListProviderModels400ResponseAnyOf {
	return &NullableAiProfilesListProviderModels400ResponseAnyOf{value: val, isSet: true}
}

func (v NullableAiProfilesListProviderModels400ResponseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiProfilesListProviderModels400ResponseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

