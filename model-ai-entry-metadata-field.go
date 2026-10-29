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

// checks if the AiEntryMetadataField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiEntryMetadataField{}

// AiEntryMetadataField struct for AiEntryMetadataField
type AiEntryMetadataField struct {
	Id float32 `json:"id"`
	Name string `json:"name"`
	// `MetadataFieldType`: 0 String, 1 Date, 2 Number, 3 SingleChoice, 4 MultiChoice.
	Type float32 `json:"type"`
	Options []AiEntryMetadataFieldOptionsInner `json:"options"`
	Value *AiMetadataValue `json:"value,omitempty"`
}

type _AiEntryMetadataField AiEntryMetadataField

// NewAiEntryMetadataField instantiates a new AiEntryMetadataField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiEntryMetadataField(id float32, name string, type_ float32, options []AiEntryMetadataFieldOptionsInner) *AiEntryMetadataField {
	this := AiEntryMetadataField{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.Options = options
	return &this
}

// NewAiEntryMetadataFieldWithDefaults instantiates a new AiEntryMetadataField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiEntryMetadataFieldWithDefaults() *AiEntryMetadataField {
	this := AiEntryMetadataField{}
	return &this
}

// GetId returns the Id field value
func (o *AiEntryMetadataField) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiEntryMetadataField) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiEntryMetadataField) SetId(v float32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AiEntryMetadataField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiEntryMetadataField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiEntryMetadataField) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *AiEntryMetadataField) GetType() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiEntryMetadataField) GetTypeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AiEntryMetadataField) SetType(v float32) {
	o.Type = v
}

// GetOptions returns the Options field value
func (o *AiEntryMetadataField) GetOptions() []AiEntryMetadataFieldOptionsInner {
	if o == nil {
		var ret []AiEntryMetadataFieldOptionsInner
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *AiEntryMetadataField) GetOptionsOk() ([]AiEntryMetadataFieldOptionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *AiEntryMetadataField) SetOptions(v []AiEntryMetadataFieldOptionsInner) {
	o.Options = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AiEntryMetadataField) GetValue() AiMetadataValue {
	if o == nil || IsNil(o.Value) {
		var ret AiMetadataValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiEntryMetadataField) GetValueOk() (*AiMetadataValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AiEntryMetadataField) IsValueSet() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given AiMetadataValue and assigns it to the Value field.
func (o *AiEntryMetadataField) SetValue(v AiMetadataValue) {
	o.Value = &v
}

func (o AiEntryMetadataField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiEntryMetadataField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["options"] = o.Options
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

func (o *AiEntryMetadataField) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"type",
		"options",
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

	varAiEntryMetadataField := _AiEntryMetadataField{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiEntryMetadataField)

	if err != nil {
		return err
	}

	*o = AiEntryMetadataField(varAiEntryMetadataField)

	return err
}

type NullableAiEntryMetadataField struct {
	value *AiEntryMetadataField
	isSet bool
}

func (v NullableAiEntryMetadataField) Get() *AiEntryMetadataField {
	return v.value
}

func (v *NullableAiEntryMetadataField) Set(val *AiEntryMetadataField) {
	v.value = val
	v.isSet = true
}

func (v NullableAiEntryMetadataField) IsSet() bool {
	return v.isSet
}

func (v *NullableAiEntryMetadataField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiEntryMetadataField(val *AiEntryMetadataField) *NullableAiEntryMetadataField {
	return &NullableAiEntryMetadataField{value: val, isSet: true}
}

func (v NullableAiEntryMetadataField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiEntryMetadataField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

