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

// checks if the AiAssignmentsAssignRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiAssignmentsAssignRequest{}

// AiAssignmentsAssignRequest struct for AiAssignmentsAssignRequest
type AiAssignmentsAssignRequest struct {
	// Action the assignment applies to.
	ActionType AiActionType `json:"actionType"`
	// Profile id to bind.
	ProfileId string `json:"profileId"`
}

type _AiAssignmentsAssignRequest AiAssignmentsAssignRequest

// NewAiAssignmentsAssignRequest instantiates a new AiAssignmentsAssignRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAssignmentsAssignRequest(actionType AiActionType, profileId string) *AiAssignmentsAssignRequest {
	this := AiAssignmentsAssignRequest{}
	this.ActionType = actionType
	this.ProfileId = profileId
	return &this
}

// NewAiAssignmentsAssignRequestWithDefaults instantiates a new AiAssignmentsAssignRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAssignmentsAssignRequestWithDefaults() *AiAssignmentsAssignRequest {
	this := AiAssignmentsAssignRequest{}
	return &this
}

// GetActionType returns the ActionType field value
func (o *AiAssignmentsAssignRequest) GetActionType() AiActionType {
	if o == nil {
		var ret AiActionType
		return ret
	}

	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *AiAssignmentsAssignRequest) GetActionTypeOk() (*AiActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value
func (o *AiAssignmentsAssignRequest) SetActionType(v AiActionType) {
	o.ActionType = v
}

// GetProfileId returns the ProfileId field value
func (o *AiAssignmentsAssignRequest) GetProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value
// and a boolean to check if the value has been set.
func (o *AiAssignmentsAssignRequest) GetProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfileId, true
}

// SetProfileId sets field value
func (o *AiAssignmentsAssignRequest) SetProfileId(v string) {
	o.ProfileId = v
}

func (o AiAssignmentsAssignRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiAssignmentsAssignRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["actionType"] = o.ActionType
	toSerialize["profileId"] = o.ProfileId
	return toSerialize, nil
}

func (o *AiAssignmentsAssignRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"actionType",
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

	varAiAssignmentsAssignRequest := _AiAssignmentsAssignRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiAssignmentsAssignRequest)

	if err != nil {
		return err
	}

	*o = AiAssignmentsAssignRequest(varAiAssignmentsAssignRequest)

	return err
}

type NullableAiAssignmentsAssignRequest struct {
	value *AiAssignmentsAssignRequest
	isSet bool
}

func (v NullableAiAssignmentsAssignRequest) Get() *AiAssignmentsAssignRequest {
	return v.value
}

func (v *NullableAiAssignmentsAssignRequest) Set(val *AiAssignmentsAssignRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAssignmentsAssignRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAssignmentsAssignRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAssignmentsAssignRequest(val *AiAssignmentsAssignRequest) *NullableAiAssignmentsAssignRequest {
	return &NullableAiAssignmentsAssignRequest{value: val, isSet: true}
}

func (v NullableAiAssignmentsAssignRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAssignmentsAssignRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

