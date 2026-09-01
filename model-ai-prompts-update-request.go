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

// checks if the AiPromptsUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiPromptsUpdateRequest{}

// AiPromptsUpdateRequest struct for AiPromptsUpdateRequest
type AiPromptsUpdateRequest struct {
	// Prompt id to update.
	Id string `json:"id"`
	Updates AiPromptsUpdateRequestUpdates `json:"updates"`
}

type _AiPromptsUpdateRequest AiPromptsUpdateRequest

// NewAiPromptsUpdateRequest instantiates a new AiPromptsUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiPromptsUpdateRequest(id string, updates AiPromptsUpdateRequestUpdates) *AiPromptsUpdateRequest {
	this := AiPromptsUpdateRequest{}
	this.Id = id
	this.Updates = updates
	return &this
}

// NewAiPromptsUpdateRequestWithDefaults instantiates a new AiPromptsUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiPromptsUpdateRequestWithDefaults() *AiPromptsUpdateRequest {
	this := AiPromptsUpdateRequest{}
	return &this
}

// GetId returns the Id field value
func (o *AiPromptsUpdateRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiPromptsUpdateRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiPromptsUpdateRequest) SetId(v string) {
	o.Id = v
}

// GetUpdates returns the Updates field value
func (o *AiPromptsUpdateRequest) GetUpdates() AiPromptsUpdateRequestUpdates {
	if o == nil {
		var ret AiPromptsUpdateRequestUpdates
		return ret
	}

	return o.Updates
}

// GetUpdatesOk returns a tuple with the Updates field value
// and a boolean to check if the value has been set.
func (o *AiPromptsUpdateRequest) GetUpdatesOk() (*AiPromptsUpdateRequestUpdates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updates, true
}

// SetUpdates sets field value
func (o *AiPromptsUpdateRequest) SetUpdates(v AiPromptsUpdateRequestUpdates) {
	o.Updates = v
}

func (o AiPromptsUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiPromptsUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["updates"] = o.Updates
	return toSerialize, nil
}

func (o *AiPromptsUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"updates",
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

	varAiPromptsUpdateRequest := _AiPromptsUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiPromptsUpdateRequest)

	if err != nil {
		return err
	}

	*o = AiPromptsUpdateRequest(varAiPromptsUpdateRequest)

	return err
}

type NullableAiPromptsUpdateRequest struct {
	value *AiPromptsUpdateRequest
	isSet bool
}

func (v NullableAiPromptsUpdateRequest) Get() *AiPromptsUpdateRequest {
	return v.value
}

func (v *NullableAiPromptsUpdateRequest) Set(val *AiPromptsUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiPromptsUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiPromptsUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiPromptsUpdateRequest(val *AiPromptsUpdateRequest) *NullableAiPromptsUpdateRequest {
	return &NullableAiPromptsUpdateRequest{value: val, isSet: true}
}

func (v NullableAiPromptsUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiPromptsUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

