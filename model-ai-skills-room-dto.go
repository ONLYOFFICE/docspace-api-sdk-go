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

// checks if the AiSkillsRoomDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiSkillsRoomDto{}

// AiSkillsRoomDto An Instructions room with its skills.
type AiSkillsRoomDto struct {
	// ID of the room.
	Id int32 `json:"id"`
	// Room title.
	Title NullableString `json:"title"`
	// Room logo.
	Logo LogoDto `json:"logo"`
	// Skills of the room, sorted by file title.
	Skills []AiSkillDto `json:"skills"`
}

type _AiSkillsRoomDto AiSkillsRoomDto

// NewAiSkillsRoomDto instantiates a new AiSkillsRoomDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiSkillsRoomDto(id int32, title NullableString, logo LogoDto, skills []AiSkillDto) *AiSkillsRoomDto {
	this := AiSkillsRoomDto{}
	this.Id = id
	this.Title = title
	this.Logo = logo
	this.Skills = skills
	return &this
}

// NewAiSkillsRoomDtoWithDefaults instantiates a new AiSkillsRoomDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiSkillsRoomDtoWithDefaults() *AiSkillsRoomDto {
	this := AiSkillsRoomDto{}
	return &this
}

// GetId returns the Id field value
func (o *AiSkillsRoomDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiSkillsRoomDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiSkillsRoomDto) SetId(v int32) {
	o.Id = v
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiSkillsRoomDto) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiSkillsRoomDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *AiSkillsRoomDto) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetLogo returns the Logo field value
func (o *AiSkillsRoomDto) GetLogo() LogoDto {
	if o == nil {
		var ret LogoDto
		return ret
	}

	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value
// and a boolean to check if the value has been set.
func (o *AiSkillsRoomDto) GetLogoOk() (*LogoDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logo, true
}

// SetLogo sets field value
func (o *AiSkillsRoomDto) SetLogo(v LogoDto) {
	o.Logo = v
}

// GetSkills returns the Skills field value
// If the value is explicit nil, the zero value for []AiSkillDto will be returned
func (o *AiSkillsRoomDto) GetSkills() []AiSkillDto {
	if o == nil {
		var ret []AiSkillDto
		return ret
	}

	return o.Skills
}

// GetSkillsOk returns a tuple with the Skills field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiSkillsRoomDto) GetSkillsOk() ([]AiSkillDto, bool) {
	if o == nil || IsNil(o.Skills) {
		return nil, false
	}
	return o.Skills, true
}

// SetSkills sets field value
func (o *AiSkillsRoomDto) SetSkills(v []AiSkillDto) {
	o.Skills = v
}

func (o AiSkillsRoomDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiSkillsRoomDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title.Get()
	toSerialize["logo"] = o.Logo
	if o.Skills != nil {
		toSerialize["skills"] = o.Skills
	}
	return toSerialize, nil
}

func (o *AiSkillsRoomDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"title",
		"logo",
		"skills",
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

	varAiSkillsRoomDto := _AiSkillsRoomDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiSkillsRoomDto)

	if err != nil {
		return err
	}

	*o = AiSkillsRoomDto(varAiSkillsRoomDto)

	return err
}

type NullableAiSkillsRoomDto struct {
	value *AiSkillsRoomDto
	isSet bool
}

func (v NullableAiSkillsRoomDto) Get() *AiSkillsRoomDto {
	return v.value
}

func (v *NullableAiSkillsRoomDto) Set(val *AiSkillsRoomDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiSkillsRoomDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiSkillsRoomDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiSkillsRoomDto(val *AiSkillsRoomDto) *NullableAiSkillsRoomDto {
	return &NullableAiSkillsRoomDto{value: val, isSet: true}
}

func (v NullableAiSkillsRoomDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiSkillsRoomDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

