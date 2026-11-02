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

// checks if the AiProjectSourcesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiProjectSourcesDto{}

// AiProjectSourcesDto The sources of an AI project: the Instructions and Knowledge rooms linked to it.
type AiProjectSourcesDto struct {
	// Linked Instructions rooms, sorted by title.
	InstructionsRooms []AiProjectSourceDto `json:"instructionsRooms"`
	// Linked Knowledge rooms, sorted by title.
	KnowledgeRooms []AiProjectSourceDto `json:"knowledgeRooms"`
}

type _AiProjectSourcesDto AiProjectSourcesDto

// NewAiProjectSourcesDto instantiates a new AiProjectSourcesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiProjectSourcesDto(instructionsRooms []AiProjectSourceDto, knowledgeRooms []AiProjectSourceDto) *AiProjectSourcesDto {
	this := AiProjectSourcesDto{}
	this.InstructionsRooms = instructionsRooms
	this.KnowledgeRooms = knowledgeRooms
	return &this
}

// NewAiProjectSourcesDtoWithDefaults instantiates a new AiProjectSourcesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiProjectSourcesDtoWithDefaults() *AiProjectSourcesDto {
	this := AiProjectSourcesDto{}
	return &this
}

// GetInstructionsRooms returns the InstructionsRooms field value
// If the value is explicit nil, the zero value for []AiProjectSourceDto will be returned
func (o *AiProjectSourcesDto) GetInstructionsRooms() []AiProjectSourceDto {
	if o == nil {
		var ret []AiProjectSourceDto
		return ret
	}

	return o.InstructionsRooms
}

// GetInstructionsRoomsOk returns a tuple with the InstructionsRooms field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiProjectSourcesDto) GetInstructionsRoomsOk() ([]AiProjectSourceDto, bool) {
	if o == nil || IsNil(o.InstructionsRooms) {
		return nil, false
	}
	return o.InstructionsRooms, true
}

// SetInstructionsRooms sets field value
func (o *AiProjectSourcesDto) SetInstructionsRooms(v []AiProjectSourceDto) {
	o.InstructionsRooms = v
}

// GetKnowledgeRooms returns the KnowledgeRooms field value
// If the value is explicit nil, the zero value for []AiProjectSourceDto will be returned
func (o *AiProjectSourcesDto) GetKnowledgeRooms() []AiProjectSourceDto {
	if o == nil {
		var ret []AiProjectSourceDto
		return ret
	}

	return o.KnowledgeRooms
}

// GetKnowledgeRoomsOk returns a tuple with the KnowledgeRooms field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiProjectSourcesDto) GetKnowledgeRoomsOk() ([]AiProjectSourceDto, bool) {
	if o == nil || IsNil(o.KnowledgeRooms) {
		return nil, false
	}
	return o.KnowledgeRooms, true
}

// SetKnowledgeRooms sets field value
func (o *AiProjectSourcesDto) SetKnowledgeRooms(v []AiProjectSourceDto) {
	o.KnowledgeRooms = v
}

func (o AiProjectSourcesDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiProjectSourcesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.InstructionsRooms != nil {
		toSerialize["instructionsRooms"] = o.InstructionsRooms
	}
	if o.KnowledgeRooms != nil {
		toSerialize["knowledgeRooms"] = o.KnowledgeRooms
	}
	return toSerialize, nil
}

func (o *AiProjectSourcesDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"instructionsRooms",
		"knowledgeRooms",
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

	varAiProjectSourcesDto := _AiProjectSourcesDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiProjectSourcesDto)

	if err != nil {
		return err
	}

	*o = AiProjectSourcesDto(varAiProjectSourcesDto)

	return err
}

type NullableAiProjectSourcesDto struct {
	value *AiProjectSourcesDto
	isSet bool
}

func (v NullableAiProjectSourcesDto) Get() *AiProjectSourcesDto {
	return v.value
}

func (v *NullableAiProjectSourcesDto) Set(val *AiProjectSourcesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiProjectSourcesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiProjectSourcesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiProjectSourcesDto(val *AiProjectSourcesDto) *NullableAiProjectSourcesDto {
	return &NullableAiProjectSourcesDto{value: val, isSet: true}
}

func (v NullableAiProjectSourcesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiProjectSourcesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

