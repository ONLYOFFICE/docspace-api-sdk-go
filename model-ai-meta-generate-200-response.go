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

// checks if the AiMetaGenerate200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiMetaGenerate200Response{}

// AiMetaGenerate200Response struct for AiMetaGenerate200Response
type AiMetaGenerate200Response struct {
	// The proposed custom metadata fields with the values the text supports.
	Fields []AiMetaGenerate200ResponseFieldsInner `json:"fields"`
	// The documents the proposal was read from.
	Sources []AiMetaGenerate200ResponseSourcesInner `json:"sources"`
	// The entries that were left out, with the reason for each.
	Skipped []AiMetaGenerate200ResponseSkippedInner `json:"skipped"`
}

type _AiMetaGenerate200Response AiMetaGenerate200Response

// NewAiMetaGenerate200Response instantiates a new AiMetaGenerate200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiMetaGenerate200Response(fields []AiMetaGenerate200ResponseFieldsInner, sources []AiMetaGenerate200ResponseSourcesInner, skipped []AiMetaGenerate200ResponseSkippedInner) *AiMetaGenerate200Response {
	this := AiMetaGenerate200Response{}
	this.Fields = fields
	this.Sources = sources
	this.Skipped = skipped
	return &this
}

// NewAiMetaGenerate200ResponseWithDefaults instantiates a new AiMetaGenerate200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiMetaGenerate200ResponseWithDefaults() *AiMetaGenerate200Response {
	this := AiMetaGenerate200Response{}
	return &this
}

// GetFields returns the Fields field value
func (o *AiMetaGenerate200Response) GetFields() []AiMetaGenerate200ResponseFieldsInner {
	if o == nil {
		var ret []AiMetaGenerate200ResponseFieldsInner
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *AiMetaGenerate200Response) GetFieldsOk() ([]AiMetaGenerate200ResponseFieldsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *AiMetaGenerate200Response) SetFields(v []AiMetaGenerate200ResponseFieldsInner) {
	o.Fields = v
}

// GetSources returns the Sources field value
func (o *AiMetaGenerate200Response) GetSources() []AiMetaGenerate200ResponseSourcesInner {
	if o == nil {
		var ret []AiMetaGenerate200ResponseSourcesInner
		return ret
	}

	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *AiMetaGenerate200Response) GetSourcesOk() ([]AiMetaGenerate200ResponseSourcesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sources, true
}

// SetSources sets field value
func (o *AiMetaGenerate200Response) SetSources(v []AiMetaGenerate200ResponseSourcesInner) {
	o.Sources = v
}

// GetSkipped returns the Skipped field value
func (o *AiMetaGenerate200Response) GetSkipped() []AiMetaGenerate200ResponseSkippedInner {
	if o == nil {
		var ret []AiMetaGenerate200ResponseSkippedInner
		return ret
	}

	return o.Skipped
}

// GetSkippedOk returns a tuple with the Skipped field value
// and a boolean to check if the value has been set.
func (o *AiMetaGenerate200Response) GetSkippedOk() ([]AiMetaGenerate200ResponseSkippedInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Skipped, true
}

// SetSkipped sets field value
func (o *AiMetaGenerate200Response) SetSkipped(v []AiMetaGenerate200ResponseSkippedInner) {
	o.Skipped = v
}

func (o AiMetaGenerate200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiMetaGenerate200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	toSerialize["sources"] = o.Sources
	toSerialize["skipped"] = o.Skipped
	return toSerialize, nil
}

func (o *AiMetaGenerate200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fields",
		"sources",
		"skipped",
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

	varAiMetaGenerate200Response := _AiMetaGenerate200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiMetaGenerate200Response)

	if err != nil {
		return err
	}

	*o = AiMetaGenerate200Response(varAiMetaGenerate200Response)

	return err
}

type NullableAiMetaGenerate200Response struct {
	value *AiMetaGenerate200Response
	isSet bool
}

func (v NullableAiMetaGenerate200Response) Get() *AiMetaGenerate200Response {
	return v.value
}

func (v *NullableAiMetaGenerate200Response) Set(val *AiMetaGenerate200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAiMetaGenerate200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAiMetaGenerate200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiMetaGenerate200Response(val *AiMetaGenerate200Response) *NullableAiMetaGenerate200Response {
	return &NullableAiMetaGenerate200Response{value: val, isSet: true}
}

func (v NullableAiMetaGenerate200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiMetaGenerate200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

