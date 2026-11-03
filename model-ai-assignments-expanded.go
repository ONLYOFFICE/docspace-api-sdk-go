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

// checks if the AiAssignmentsExpanded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiAssignmentsExpanded{}

// AiAssignmentsExpanded The bindings with `expand=profile`: the bound profiles themselves, plus how many profiles the portal has, so a chat can start on this one read.
type AiAssignmentsExpanded struct {
	// The bound profiles keyed by action type, as `GET api/2.0/ai/profiles/list` describes them. A binding whose profile no longer exists is left out.
	Assignments map[string]AiProfile `json:"assignments"`
	// How many profiles the portal has.
	ProfilesTotal int32 `json:"profilesTotal"`
	// How many of those profiles can chat; a profile without declared capabilities counts as one that can.
	ChatProfilesTotal int32 `json:"chatProfilesTotal"`
}

type _AiAssignmentsExpanded AiAssignmentsExpanded

// NewAiAssignmentsExpanded instantiates a new AiAssignmentsExpanded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAssignmentsExpanded(assignments map[string]AiProfile, profilesTotal int32, chatProfilesTotal int32) *AiAssignmentsExpanded {
	this := AiAssignmentsExpanded{}
	this.Assignments = assignments
	this.ProfilesTotal = profilesTotal
	this.ChatProfilesTotal = chatProfilesTotal
	return &this
}

// NewAiAssignmentsExpandedWithDefaults instantiates a new AiAssignmentsExpanded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAssignmentsExpandedWithDefaults() *AiAssignmentsExpanded {
	this := AiAssignmentsExpanded{}
	return &this
}

// GetAssignments returns the Assignments field value
func (o *AiAssignmentsExpanded) GetAssignments() map[string]AiProfile {
	if o == nil {
		var ret map[string]AiProfile
		return ret
	}

	return o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value
// and a boolean to check if the value has been set.
func (o *AiAssignmentsExpanded) GetAssignmentsOk() (map[string]AiProfile, bool) {
	if o == nil {
		return map[string]AiProfile{}, false
	}
	return o.Assignments, true
}

// SetAssignments sets field value
func (o *AiAssignmentsExpanded) SetAssignments(v map[string]AiProfile) {
	o.Assignments = v
}

// GetProfilesTotal returns the ProfilesTotal field value
func (o *AiAssignmentsExpanded) GetProfilesTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProfilesTotal
}

// GetProfilesTotalOk returns a tuple with the ProfilesTotal field value
// and a boolean to check if the value has been set.
func (o *AiAssignmentsExpanded) GetProfilesTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfilesTotal, true
}

// SetProfilesTotal sets field value
func (o *AiAssignmentsExpanded) SetProfilesTotal(v int32) {
	o.ProfilesTotal = v
}

// GetChatProfilesTotal returns the ChatProfilesTotal field value
func (o *AiAssignmentsExpanded) GetChatProfilesTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChatProfilesTotal
}

// GetChatProfilesTotalOk returns a tuple with the ChatProfilesTotal field value
// and a boolean to check if the value has been set.
func (o *AiAssignmentsExpanded) GetChatProfilesTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatProfilesTotal, true
}

// SetChatProfilesTotal sets field value
func (o *AiAssignmentsExpanded) SetChatProfilesTotal(v int32) {
	o.ChatProfilesTotal = v
}

func (o AiAssignmentsExpanded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiAssignmentsExpanded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assignments"] = o.Assignments
	toSerialize["profilesTotal"] = o.ProfilesTotal
	toSerialize["chatProfilesTotal"] = o.ChatProfilesTotal
	return toSerialize, nil
}

func (o *AiAssignmentsExpanded) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"assignments",
		"profilesTotal",
		"chatProfilesTotal",
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

	varAiAssignmentsExpanded := _AiAssignmentsExpanded{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiAssignmentsExpanded)

	if err != nil {
		return err
	}

	*o = AiAssignmentsExpanded(varAiAssignmentsExpanded)

	return err
}

type NullableAiAssignmentsExpanded struct {
	value *AiAssignmentsExpanded
	isSet bool
}

func (v NullableAiAssignmentsExpanded) Get() *AiAssignmentsExpanded {
	return v.value
}

func (v *NullableAiAssignmentsExpanded) Set(val *AiAssignmentsExpanded) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAssignmentsExpanded) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAssignmentsExpanded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAssignmentsExpanded(val *AiAssignmentsExpanded) *NullableAiAssignmentsExpanded {
	return &NullableAiAssignmentsExpanded{value: val, isSet: true}
}

func (v NullableAiAssignmentsExpanded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAssignmentsExpanded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

