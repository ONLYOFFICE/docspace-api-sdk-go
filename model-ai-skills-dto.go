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

// checks if the AiSkillsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiSkillsDto{}

// AiSkillsDto Instructions rooms with their skills.
type AiSkillsDto struct {
	// The rooms, sorted by title.
	Rooms []AiSkillsRoomDto `json:"rooms"`
}

type _AiSkillsDto AiSkillsDto

// NewAiSkillsDto instantiates a new AiSkillsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiSkillsDto(rooms []AiSkillsRoomDto) *AiSkillsDto {
	this := AiSkillsDto{}
	this.Rooms = rooms
	return &this
}

// NewAiSkillsDtoWithDefaults instantiates a new AiSkillsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiSkillsDtoWithDefaults() *AiSkillsDto {
	this := AiSkillsDto{}
	return &this
}

// GetRooms returns the Rooms field value
// If the value is explicit nil, the zero value for []AiSkillsRoomDto will be returned
func (o *AiSkillsDto) GetRooms() []AiSkillsRoomDto {
	if o == nil {
		var ret []AiSkillsRoomDto
		return ret
	}

	return o.Rooms
}

// GetRoomsOk returns a tuple with the Rooms field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiSkillsDto) GetRoomsOk() ([]AiSkillsRoomDto, bool) {
	if o == nil || IsNil(o.Rooms) {
		return nil, false
	}
	return o.Rooms, true
}

// SetRooms sets field value
func (o *AiSkillsDto) SetRooms(v []AiSkillsRoomDto) {
	o.Rooms = v
}

func (o AiSkillsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiSkillsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Rooms != nil {
		toSerialize["rooms"] = o.Rooms
	}
	return toSerialize, nil
}

func (o *AiSkillsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rooms",
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

	varAiSkillsDto := _AiSkillsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiSkillsDto)

	if err != nil {
		return err
	}

	*o = AiSkillsDto(varAiSkillsDto)

	return err
}

type NullableAiSkillsDto struct {
	value *AiSkillsDto
	isSet bool
}

func (v NullableAiSkillsDto) Get() *AiSkillsDto {
	return v.value
}

func (v *NullableAiSkillsDto) Set(val *AiSkillsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiSkillsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiSkillsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiSkillsDto(val *AiSkillsDto) *NullableAiSkillsDto {
	return &NullableAiSkillsDto{value: val, isSet: true}
}

func (v NullableAiSkillsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiSkillsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

