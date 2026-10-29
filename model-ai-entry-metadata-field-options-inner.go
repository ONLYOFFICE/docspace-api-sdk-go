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

// checks if the AiEntryMetadataFieldOptionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiEntryMetadataFieldOptionsInner{}

// AiEntryMetadataFieldOptionsInner struct for AiEntryMetadataFieldOptionsInner
type AiEntryMetadataFieldOptionsInner struct {
	Id string `json:"id"`
	Value string `json:"value"`
}

type _AiEntryMetadataFieldOptionsInner AiEntryMetadataFieldOptionsInner

// NewAiEntryMetadataFieldOptionsInner instantiates a new AiEntryMetadataFieldOptionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiEntryMetadataFieldOptionsInner(id string, value string) *AiEntryMetadataFieldOptionsInner {
	this := AiEntryMetadataFieldOptionsInner{}
	this.Id = id
	this.Value = value
	return &this
}

// NewAiEntryMetadataFieldOptionsInnerWithDefaults instantiates a new AiEntryMetadataFieldOptionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiEntryMetadataFieldOptionsInnerWithDefaults() *AiEntryMetadataFieldOptionsInner {
	this := AiEntryMetadataFieldOptionsInner{}
	return &this
}

// GetId returns the Id field value
func (o *AiEntryMetadataFieldOptionsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiEntryMetadataFieldOptionsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiEntryMetadataFieldOptionsInner) SetId(v string) {
	o.Id = v
}

// GetValue returns the Value field value
func (o *AiEntryMetadataFieldOptionsInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AiEntryMetadataFieldOptionsInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AiEntryMetadataFieldOptionsInner) SetValue(v string) {
	o.Value = v
}

func (o AiEntryMetadataFieldOptionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiEntryMetadataFieldOptionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *AiEntryMetadataFieldOptionsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varAiEntryMetadataFieldOptionsInner := _AiEntryMetadataFieldOptionsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiEntryMetadataFieldOptionsInner)

	if err != nil {
		return err
	}

	*o = AiEntryMetadataFieldOptionsInner(varAiEntryMetadataFieldOptionsInner)

	return err
}

type NullableAiEntryMetadataFieldOptionsInner struct {
	value *AiEntryMetadataFieldOptionsInner
	isSet bool
}

func (v NullableAiEntryMetadataFieldOptionsInner) Get() *AiEntryMetadataFieldOptionsInner {
	return v.value
}

func (v *NullableAiEntryMetadataFieldOptionsInner) Set(val *AiEntryMetadataFieldOptionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAiEntryMetadataFieldOptionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAiEntryMetadataFieldOptionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiEntryMetadataFieldOptionsInner(val *AiEntryMetadataFieldOptionsInner) *NullableAiEntryMetadataFieldOptionsInner {
	return &NullableAiEntryMetadataFieldOptionsInner{value: val, isSet: true}
}

func (v NullableAiEntryMetadataFieldOptionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiEntryMetadataFieldOptionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

