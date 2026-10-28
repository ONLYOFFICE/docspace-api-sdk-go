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
)

// checks if the AiAgentsGet200ResponseAllOfResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiAgentsGet200ResponseAllOfResponse{}

// AiAgentsGet200ResponseAllOfResponse struct for AiAgentsGet200ResponseAllOfResponse
type AiAgentsGet200ResponseAllOfResponse struct {
	// The AI profile bound to this agent, added by this service on top of what the internal service returns. Absent when the agent has no profile assigned.
	ProfileId *string `json:"profileId,omitempty"`
}

// NewAiAgentsGet200ResponseAllOfResponse instantiates a new AiAgentsGet200ResponseAllOfResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAgentsGet200ResponseAllOfResponse() *AiAgentsGet200ResponseAllOfResponse {
	this := AiAgentsGet200ResponseAllOfResponse{}
	return &this
}

// NewAiAgentsGet200ResponseAllOfResponseWithDefaults instantiates a new AiAgentsGet200ResponseAllOfResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAgentsGet200ResponseAllOfResponseWithDefaults() *AiAgentsGet200ResponseAllOfResponse {
	this := AiAgentsGet200ResponseAllOfResponse{}
	return &this
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *AiAgentsGet200ResponseAllOfResponse) GetProfileId() string {
	if o == nil || IsNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsGet200ResponseAllOfResponse) GetProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *AiAgentsGet200ResponseAllOfResponse) IsProfileIdSet() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *AiAgentsGet200ResponseAllOfResponse) SetProfileId(v string) {
	o.ProfileId = &v
}

func (o AiAgentsGet200ResponseAllOfResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiAgentsGet200ResponseAllOfResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	return toSerialize, nil
}

type NullableAiAgentsGet200ResponseAllOfResponse struct {
	value *AiAgentsGet200ResponseAllOfResponse
	isSet bool
}

func (v NullableAiAgentsGet200ResponseAllOfResponse) Get() *AiAgentsGet200ResponseAllOfResponse {
	return v.value
}

func (v *NullableAiAgentsGet200ResponseAllOfResponse) Set(val *AiAgentsGet200ResponseAllOfResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAgentsGet200ResponseAllOfResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAgentsGet200ResponseAllOfResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAgentsGet200ResponseAllOfResponse(val *AiAgentsGet200ResponseAllOfResponse) *NullableAiAgentsGet200ResponseAllOfResponse {
	return &NullableAiAgentsGet200ResponseAllOfResponse{value: val, isSet: true}
}

func (v NullableAiAgentsGet200ResponseAllOfResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAgentsGet200ResponseAllOfResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

