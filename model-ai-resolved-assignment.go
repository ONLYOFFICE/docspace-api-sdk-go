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

// checks if the AiResolvedAssignment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiResolvedAssignment{}

// AiResolvedAssignment Resolved profile for an action — both the storage row and its ID.
type AiResolvedAssignment struct {
	// The identifier of the resolved profile.
	ProfileId string `json:"profileId"`
	// The resolved profile itself.
	Profile AiProfile `json:"profile"`
}

type _AiResolvedAssignment AiResolvedAssignment

// NewAiResolvedAssignment instantiates a new AiResolvedAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiResolvedAssignment(profileId string, profile AiProfile) *AiResolvedAssignment {
	this := AiResolvedAssignment{}
	this.ProfileId = profileId
	this.Profile = profile
	return &this
}

// NewAiResolvedAssignmentWithDefaults instantiates a new AiResolvedAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiResolvedAssignmentWithDefaults() *AiResolvedAssignment {
	this := AiResolvedAssignment{}
	return &this
}

// GetProfileId returns the ProfileId field value
func (o *AiResolvedAssignment) GetProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value
// and a boolean to check if the value has been set.
func (o *AiResolvedAssignment) GetProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfileId, true
}

// SetProfileId sets field value
func (o *AiResolvedAssignment) SetProfileId(v string) {
	o.ProfileId = v
}

// GetProfile returns the Profile field value
func (o *AiResolvedAssignment) GetProfile() AiProfile {
	if o == nil {
		var ret AiProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *AiResolvedAssignment) GetProfileOk() (*AiProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *AiResolvedAssignment) SetProfile(v AiProfile) {
	o.Profile = v
}

func (o AiResolvedAssignment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiResolvedAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["profileId"] = o.ProfileId
	toSerialize["profile"] = o.Profile
	return toSerialize, nil
}

func (o *AiResolvedAssignment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"profileId",
		"profile",
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

	varAiResolvedAssignment := _AiResolvedAssignment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiResolvedAssignment)

	if err != nil {
		return err
	}

	*o = AiResolvedAssignment(varAiResolvedAssignment)

	return err
}

type NullableAiResolvedAssignment struct {
	value *AiResolvedAssignment
	isSet bool
}

func (v NullableAiResolvedAssignment) Get() *AiResolvedAssignment {
	return v.value
}

func (v *NullableAiResolvedAssignment) Set(val *AiResolvedAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableAiResolvedAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableAiResolvedAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiResolvedAssignment(val *AiResolvedAssignment) *NullableAiResolvedAssignment {
	return &NullableAiResolvedAssignment{value: val, isSet: true}
}

func (v NullableAiResolvedAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiResolvedAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

