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

// checks if the AiAssignmentsCascadeProfileDeleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiAssignmentsCascadeProfileDeleteRequest{}

// AiAssignmentsCascadeProfileDeleteRequest struct for AiAssignmentsCascadeProfileDeleteRequest
type AiAssignmentsCascadeProfileDeleteRequest struct {
	// The profile whose assignments are removed. May be sent as the `profileId` query parameter instead of in the body.
	ProfileId string `json:"profileId"`
}

type _AiAssignmentsCascadeProfileDeleteRequest AiAssignmentsCascadeProfileDeleteRequest

// NewAiAssignmentsCascadeProfileDeleteRequest instantiates a new AiAssignmentsCascadeProfileDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAssignmentsCascadeProfileDeleteRequest(profileId string) *AiAssignmentsCascadeProfileDeleteRequest {
	this := AiAssignmentsCascadeProfileDeleteRequest{}
	this.ProfileId = profileId
	return &this
}

// NewAiAssignmentsCascadeProfileDeleteRequestWithDefaults instantiates a new AiAssignmentsCascadeProfileDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAssignmentsCascadeProfileDeleteRequestWithDefaults() *AiAssignmentsCascadeProfileDeleteRequest {
	this := AiAssignmentsCascadeProfileDeleteRequest{}
	return &this
}

// GetProfileId returns the ProfileId field value
func (o *AiAssignmentsCascadeProfileDeleteRequest) GetProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value
// and a boolean to check if the value has been set.
func (o *AiAssignmentsCascadeProfileDeleteRequest) GetProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfileId, true
}

// SetProfileId sets field value
func (o *AiAssignmentsCascadeProfileDeleteRequest) SetProfileId(v string) {
	o.ProfileId = v
}

func (o AiAssignmentsCascadeProfileDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiAssignmentsCascadeProfileDeleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["profileId"] = o.ProfileId
	return toSerialize, nil
}

func (o *AiAssignmentsCascadeProfileDeleteRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"profileId",
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

	varAiAssignmentsCascadeProfileDeleteRequest := _AiAssignmentsCascadeProfileDeleteRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiAssignmentsCascadeProfileDeleteRequest)

	if err != nil {
		return err
	}

	*o = AiAssignmentsCascadeProfileDeleteRequest(varAiAssignmentsCascadeProfileDeleteRequest)

	return err
}

type NullableAiAssignmentsCascadeProfileDeleteRequest struct {
	value *AiAssignmentsCascadeProfileDeleteRequest
	isSet bool
}

func (v NullableAiAssignmentsCascadeProfileDeleteRequest) Get() *AiAssignmentsCascadeProfileDeleteRequest {
	return v.value
}

func (v *NullableAiAssignmentsCascadeProfileDeleteRequest) Set(val *AiAssignmentsCascadeProfileDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAssignmentsCascadeProfileDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAssignmentsCascadeProfileDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAssignmentsCascadeProfileDeleteRequest(val *AiAssignmentsCascadeProfileDeleteRequest) *NullableAiAssignmentsCascadeProfileDeleteRequest {
	return &NullableAiAssignmentsCascadeProfileDeleteRequest{value: val, isSet: true}
}

func (v NullableAiAssignmentsCascadeProfileDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAssignmentsCascadeProfileDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

