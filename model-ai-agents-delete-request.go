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

// checks if the AiAgentsDeleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiAgentsDeleteRequest{}

// AiAgentsDeleteRequest struct for AiAgentsDeleteRequest
type AiAgentsDeleteRequest struct {
	// Delete the room after the editing session finishes.
	DeleteAfter *bool `json:"deleteAfter,omitempty"`
}

// NewAiAgentsDeleteRequest instantiates a new AiAgentsDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAgentsDeleteRequest() *AiAgentsDeleteRequest {
	this := AiAgentsDeleteRequest{}
	return &this
}

// NewAiAgentsDeleteRequestWithDefaults instantiates a new AiAgentsDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAgentsDeleteRequestWithDefaults() *AiAgentsDeleteRequest {
	this := AiAgentsDeleteRequest{}
	return &this
}

// GetDeleteAfter returns the DeleteAfter field value if set, zero value otherwise.
func (o *AiAgentsDeleteRequest) GetDeleteAfter() bool {
	if o == nil || IsNil(o.DeleteAfter) {
		var ret bool
		return ret
	}
	return *o.DeleteAfter
}

// GetDeleteAfterOk returns a tuple with the DeleteAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsDeleteRequest) GetDeleteAfterOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteAfter) {
		return nil, false
	}
	return o.DeleteAfter, true
}

// HasDeleteAfter returns a boolean if a field has been set.
func (o *AiAgentsDeleteRequest) IsDeleteAfterSet() bool {
	if o != nil && !IsNil(o.DeleteAfter) {
		return true
	}

	return false
}

// SetDeleteAfter gets a reference to the given bool and assigns it to the DeleteAfter field.
func (o *AiAgentsDeleteRequest) SetDeleteAfter(v bool) {
	o.DeleteAfter = &v
}

func (o AiAgentsDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiAgentsDeleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeleteAfter) {
		toSerialize["deleteAfter"] = o.DeleteAfter
	}
	return toSerialize, nil
}

type NullableAiAgentsDeleteRequest struct {
	value *AiAgentsDeleteRequest
	isSet bool
}

func (v NullableAiAgentsDeleteRequest) Get() *AiAgentsDeleteRequest {
	return v.value
}

func (v *NullableAiAgentsDeleteRequest) Set(val *AiAgentsDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAgentsDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAgentsDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAgentsDeleteRequest(val *AiAgentsDeleteRequest) *NullableAiAgentsDeleteRequest {
	return &NullableAiAgentsDeleteRequest{value: val, isSet: true}
}

func (v NullableAiAgentsDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAgentsDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

