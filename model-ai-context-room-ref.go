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

// checks if the AiContextRoomRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiContextRoomRef{}

// AiContextRoomRef One Instructions room a round runs with as context: the account and room ID from `GET api/2.0/ai/context/get-context-folders`, plus the skills of the room the round may use.
type AiContextRoomRef struct {
	// The account the room belongs to; `docspace` for a room of this portal.
	Cloud string `json:"cloud"`
	// The room ID.
	Id string `json:"id"`
	// The skills of the room the round may use, as `GET api/2.0/ai/context/get-room-skills` lists them.
	Skills []AiRoomSkill `json:"skills,omitempty"`
}

type _AiContextRoomRef AiContextRoomRef

// NewAiContextRoomRef instantiates a new AiContextRoomRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiContextRoomRef(cloud string, id string) *AiContextRoomRef {
	this := AiContextRoomRef{}
	this.Cloud = cloud
	this.Id = id
	return &this
}

// NewAiContextRoomRefWithDefaults instantiates a new AiContextRoomRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiContextRoomRefWithDefaults() *AiContextRoomRef {
	this := AiContextRoomRef{}
	return &this
}

// GetCloud returns the Cloud field value
func (o *AiContextRoomRef) GetCloud() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value
// and a boolean to check if the value has been set.
func (o *AiContextRoomRef) GetCloudOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cloud, true
}

// SetCloud sets field value
func (o *AiContextRoomRef) SetCloud(v string) {
	o.Cloud = v
}

// GetId returns the Id field value
func (o *AiContextRoomRef) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiContextRoomRef) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiContextRoomRef) SetId(v string) {
	o.Id = v
}

// GetSkills returns the Skills field value if set, zero value otherwise.
func (o *AiContextRoomRef) GetSkills() []AiRoomSkill {
	if o == nil || IsNil(o.Skills) {
		var ret []AiRoomSkill
		return ret
	}
	return o.Skills
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiContextRoomRef) GetSkillsOk() ([]AiRoomSkill, bool) {
	if o == nil || IsNil(o.Skills) {
		return nil, false
	}
	return o.Skills, true
}

// HasSkills returns a boolean if a field has been set.
func (o *AiContextRoomRef) IsSkillsSet() bool {
	if o != nil && !IsNil(o.Skills) {
		return true
	}

	return false
}

// SetSkills gets a reference to the given []AiRoomSkill and assigns it to the Skills field.
func (o *AiContextRoomRef) SetSkills(v []AiRoomSkill) {
	o.Skills = v
}

func (o AiContextRoomRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiContextRoomRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloud"] = o.Cloud
	toSerialize["id"] = o.Id
	if !IsNil(o.Skills) {
		toSerialize["skills"] = o.Skills
	}
	return toSerialize, nil
}

func (o *AiContextRoomRef) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloud",
		"id",
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

	varAiContextRoomRef := _AiContextRoomRef{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiContextRoomRef)

	if err != nil {
		return err
	}

	*o = AiContextRoomRef(varAiContextRoomRef)

	return err
}

type NullableAiContextRoomRef struct {
	value *AiContextRoomRef
	isSet bool
}

func (v NullableAiContextRoomRef) Get() *AiContextRoomRef {
	return v.value
}

func (v *NullableAiContextRoomRef) Set(val *AiContextRoomRef) {
	v.value = val
	v.isSet = true
}

func (v NullableAiContextRoomRef) IsSet() bool {
	return v.isSet
}

func (v *NullableAiContextRoomRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiContextRoomRef(val *AiContextRoomRef) *NullableAiContextRoomRef {
	return &NullableAiContextRoomRef{value: val, isSet: true}
}

func (v NullableAiContextRoomRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiContextRoomRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

