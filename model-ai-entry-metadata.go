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

// checks if the AiEntryMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiEntryMetadata{}

// AiEntryMetadata struct for AiEntryMetadata
type AiEntryMetadata struct {
	Templates []AiEntryMetadataTemplatesInner `json:"templates"`
	CustomFields []AiEntryMetadataCustomFieldsInner `json:"customFields"`
}

type _AiEntryMetadata AiEntryMetadata

// NewAiEntryMetadata instantiates a new AiEntryMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiEntryMetadata(templates []AiEntryMetadataTemplatesInner, customFields []AiEntryMetadataCustomFieldsInner) *AiEntryMetadata {
	this := AiEntryMetadata{}
	this.Templates = templates
	this.CustomFields = customFields
	return &this
}

// NewAiEntryMetadataWithDefaults instantiates a new AiEntryMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiEntryMetadataWithDefaults() *AiEntryMetadata {
	this := AiEntryMetadata{}
	return &this
}

// GetTemplates returns the Templates field value
func (o *AiEntryMetadata) GetTemplates() []AiEntryMetadataTemplatesInner {
	if o == nil {
		var ret []AiEntryMetadataTemplatesInner
		return ret
	}

	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value
// and a boolean to check if the value has been set.
func (o *AiEntryMetadata) GetTemplatesOk() ([]AiEntryMetadataTemplatesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Templates, true
}

// SetTemplates sets field value
func (o *AiEntryMetadata) SetTemplates(v []AiEntryMetadataTemplatesInner) {
	o.Templates = v
}

// GetCustomFields returns the CustomFields field value
func (o *AiEntryMetadata) GetCustomFields() []AiEntryMetadataCustomFieldsInner {
	if o == nil {
		var ret []AiEntryMetadataCustomFieldsInner
		return ret
	}

	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value
// and a boolean to check if the value has been set.
func (o *AiEntryMetadata) GetCustomFieldsOk() ([]AiEntryMetadataCustomFieldsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// SetCustomFields sets field value
func (o *AiEntryMetadata) SetCustomFields(v []AiEntryMetadataCustomFieldsInner) {
	o.CustomFields = v
}

func (o AiEntryMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiEntryMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["templates"] = o.Templates
	toSerialize["customFields"] = o.CustomFields
	return toSerialize, nil
}

func (o *AiEntryMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"templates",
		"customFields",
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

	varAiEntryMetadata := _AiEntryMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiEntryMetadata)

	if err != nil {
		return err
	}

	*o = AiEntryMetadata(varAiEntryMetadata)

	return err
}

type NullableAiEntryMetadata struct {
	value *AiEntryMetadata
	isSet bool
}

func (v NullableAiEntryMetadata) Get() *AiEntryMetadata {
	return v.value
}

func (v *NullableAiEntryMetadata) Set(val *AiEntryMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableAiEntryMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableAiEntryMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiEntryMetadata(val *AiEntryMetadata) *NullableAiEntryMetadata {
	return &NullableAiEntryMetadata{value: val, isSet: true}
}

func (v NullableAiEntryMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiEntryMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

