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

// checks if the UpdateAiProjectModelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAiProjectModelRequest{}

// UpdateAiProjectModelRequest The AI model to set for the AI project.
type UpdateAiProjectModelRequest struct {
	// ID of the AI model profile that answers in the project's chat; take it from the `id` of a profile that `GET api/2.0/ai/profiles/list` returns.
	ProfileId string `json:"profileId"`
}

type _UpdateAiProjectModelRequest UpdateAiProjectModelRequest

// NewUpdateAiProjectModelRequest instantiates a new UpdateAiProjectModelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAiProjectModelRequest(profileId string) *UpdateAiProjectModelRequest {
	this := UpdateAiProjectModelRequest{}
	this.ProfileId = profileId
	return &this
}

// NewUpdateAiProjectModelRequestWithDefaults instantiates a new UpdateAiProjectModelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAiProjectModelRequestWithDefaults() *UpdateAiProjectModelRequest {
	this := UpdateAiProjectModelRequest{}
	return &this
}

// GetProfileId returns the ProfileId field value
func (o *UpdateAiProjectModelRequest) GetProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value
// and a boolean to check if the value has been set.
func (o *UpdateAiProjectModelRequest) GetProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfileId, true
}

// SetProfileId sets field value
func (o *UpdateAiProjectModelRequest) SetProfileId(v string) {
	o.ProfileId = v
}

func (o UpdateAiProjectModelRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAiProjectModelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["profileId"] = o.ProfileId
	return toSerialize, nil
}

func (o *UpdateAiProjectModelRequest) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateAiProjectModelRequest := _UpdateAiProjectModelRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varUpdateAiProjectModelRequest)

	if err != nil {
		return err
	}

	*o = UpdateAiProjectModelRequest(varUpdateAiProjectModelRequest)

	return err
}

type NullableUpdateAiProjectModelRequest struct {
	value *UpdateAiProjectModelRequest
	isSet bool
}

func (v NullableUpdateAiProjectModelRequest) Get() *UpdateAiProjectModelRequest {
	return v.value
}

func (v *NullableUpdateAiProjectModelRequest) Set(val *UpdateAiProjectModelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAiProjectModelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAiProjectModelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAiProjectModelRequest(val *UpdateAiProjectModelRequest) *NullableUpdateAiProjectModelRequest {
	return &NullableUpdateAiProjectModelRequest{value: val, isSet: true}
}

func (v NullableUpdateAiProjectModelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAiProjectModelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

