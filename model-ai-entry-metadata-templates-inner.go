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

// checks if the AiEntryMetadataTemplatesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiEntryMetadataTemplatesInner{}

// AiEntryMetadataTemplatesInner struct for AiEntryMetadataTemplatesInner
type AiEntryMetadataTemplatesInner struct {
	Id float32 `json:"id"`
	Name string `json:"name"`
	Visible bool `json:"visible"`
	Fields []AiEntryMetadataField `json:"fields"`
}

type _AiEntryMetadataTemplatesInner AiEntryMetadataTemplatesInner

// NewAiEntryMetadataTemplatesInner instantiates a new AiEntryMetadataTemplatesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiEntryMetadataTemplatesInner(id float32, name string, visible bool, fields []AiEntryMetadataField) *AiEntryMetadataTemplatesInner {
	this := AiEntryMetadataTemplatesInner{}
	this.Id = id
	this.Name = name
	this.Visible = visible
	this.Fields = fields
	return &this
}

// NewAiEntryMetadataTemplatesInnerWithDefaults instantiates a new AiEntryMetadataTemplatesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiEntryMetadataTemplatesInnerWithDefaults() *AiEntryMetadataTemplatesInner {
	this := AiEntryMetadataTemplatesInner{}
	return &this
}

// GetId returns the Id field value
func (o *AiEntryMetadataTemplatesInner) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiEntryMetadataTemplatesInner) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiEntryMetadataTemplatesInner) SetId(v float32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AiEntryMetadataTemplatesInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiEntryMetadataTemplatesInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiEntryMetadataTemplatesInner) SetName(v string) {
	o.Name = v
}

// GetVisible returns the Visible field value
func (o *AiEntryMetadataTemplatesInner) GetVisible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value
// and a boolean to check if the value has been set.
func (o *AiEntryMetadataTemplatesInner) GetVisibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visible, true
}

// SetVisible sets field value
func (o *AiEntryMetadataTemplatesInner) SetVisible(v bool) {
	o.Visible = v
}

// GetFields returns the Fields field value
func (o *AiEntryMetadataTemplatesInner) GetFields() []AiEntryMetadataField {
	if o == nil {
		var ret []AiEntryMetadataField
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *AiEntryMetadataTemplatesInner) GetFieldsOk() ([]AiEntryMetadataField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *AiEntryMetadataTemplatesInner) SetFields(v []AiEntryMetadataField) {
	o.Fields = v
}

func (o AiEntryMetadataTemplatesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiEntryMetadataTemplatesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["visible"] = o.Visible
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *AiEntryMetadataTemplatesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"visible",
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

	varAiEntryMetadataTemplatesInner := _AiEntryMetadataTemplatesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiEntryMetadataTemplatesInner)

	if err != nil {
		return err
	}

	*o = AiEntryMetadataTemplatesInner(varAiEntryMetadataTemplatesInner)

	return err
}

type NullableAiEntryMetadataTemplatesInner struct {
	value *AiEntryMetadataTemplatesInner
	isSet bool
}

func (v NullableAiEntryMetadataTemplatesInner) Get() *AiEntryMetadataTemplatesInner {
	return v.value
}

func (v *NullableAiEntryMetadataTemplatesInner) Set(val *AiEntryMetadataTemplatesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAiEntryMetadataTemplatesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAiEntryMetadataTemplatesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiEntryMetadataTemplatesInner(val *AiEntryMetadataTemplatesInner) *NullableAiEntryMetadataTemplatesInner {
	return &NullableAiEntryMetadataTemplatesInner{value: val, isSet: true}
}

func (v NullableAiEntryMetadataTemplatesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiEntryMetadataTemplatesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

