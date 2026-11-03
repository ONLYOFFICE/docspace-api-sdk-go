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

// checks if the AiRoomSkill type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiRoomSkill{}

// AiRoomSkill A skill of an Instructions room: the frontmatter of one Markdown skill file. `name` and `description` are what the model is shown; the body is read through `GET api/2.0/ai/context/get-room-skill`.
type AiRoomSkill struct {
	// The skill's ID, as `GET api/2.0/ai/context/get-room-skill` takes it.
	Id string `json:"id"`
	// The skill's name, from its frontmatter.
	Name string `json:"name"`
	// What the skill is for, from its frontmatter; empty when it declares none.
	Description string `json:"description"`
}

type _AiRoomSkill AiRoomSkill

// NewAiRoomSkill instantiates a new AiRoomSkill object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiRoomSkill(id string, name string, description string) *AiRoomSkill {
	this := AiRoomSkill{}
	this.Id = id
	this.Name = name
	this.Description = description
	return &this
}

// NewAiRoomSkillWithDefaults instantiates a new AiRoomSkill object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiRoomSkillWithDefaults() *AiRoomSkill {
	this := AiRoomSkill{}
	return &this
}

// GetId returns the Id field value
func (o *AiRoomSkill) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiRoomSkill) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiRoomSkill) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AiRoomSkill) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiRoomSkill) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AiRoomSkill) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *AiRoomSkill) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AiRoomSkill) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AiRoomSkill) SetDescription(v string) {
	o.Description = v
}

func (o AiRoomSkill) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiRoomSkill) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

func (o *AiRoomSkill) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"description",
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

	varAiRoomSkill := _AiRoomSkill{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiRoomSkill)

	if err != nil {
		return err
	}

	*o = AiRoomSkill(varAiRoomSkill)

	return err
}

type NullableAiRoomSkill struct {
	value *AiRoomSkill
	isSet bool
}

func (v NullableAiRoomSkill) Get() *AiRoomSkill {
	return v.value
}

func (v *NullableAiRoomSkill) Set(val *AiRoomSkill) {
	v.value = val
	v.isSet = true
}

func (v NullableAiRoomSkill) IsSet() bool {
	return v.isSet
}

func (v *NullableAiRoomSkill) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiRoomSkill(val *AiRoomSkill) *NullableAiRoomSkill {
	return &NullableAiRoomSkill{value: val, isSet: true}
}

func (v NullableAiRoomSkill) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiRoomSkill) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

