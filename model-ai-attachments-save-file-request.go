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

// checks if the AiAttachmentsSaveFileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiAttachmentsSaveFileRequest{}

// AiAttachmentsSaveFileRequest struct for AiAttachmentsSaveFileRequest
type AiAttachmentsSaveFileRequest struct {
	Input AiAttachmentsSaveFileRequestInput `json:"input"`
	// Optional entity (room) scope.
	EntityId *string `json:"entityId,omitempty"`
}

type _AiAttachmentsSaveFileRequest AiAttachmentsSaveFileRequest

// NewAiAttachmentsSaveFileRequest instantiates a new AiAttachmentsSaveFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAttachmentsSaveFileRequest(input AiAttachmentsSaveFileRequestInput) *AiAttachmentsSaveFileRequest {
	this := AiAttachmentsSaveFileRequest{}
	this.Input = input
	return &this
}

// NewAiAttachmentsSaveFileRequestWithDefaults instantiates a new AiAttachmentsSaveFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAttachmentsSaveFileRequestWithDefaults() *AiAttachmentsSaveFileRequest {
	this := AiAttachmentsSaveFileRequest{}
	return &this
}

// GetInput returns the Input field value
func (o *AiAttachmentsSaveFileRequest) GetInput() AiAttachmentsSaveFileRequestInput {
	if o == nil {
		var ret AiAttachmentsSaveFileRequestInput
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *AiAttachmentsSaveFileRequest) GetInputOk() (*AiAttachmentsSaveFileRequestInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *AiAttachmentsSaveFileRequest) SetInput(v AiAttachmentsSaveFileRequestInput) {
	o.Input = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AiAttachmentsSaveFileRequest) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAttachmentsSaveFileRequest) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AiAttachmentsSaveFileRequest) IsEntityIdSet() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AiAttachmentsSaveFileRequest) SetEntityId(v string) {
	o.EntityId = &v
}

func (o AiAttachmentsSaveFileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiAttachmentsSaveFileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["input"] = o.Input
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	return toSerialize, nil
}

func (o *AiAttachmentsSaveFileRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"input",
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

	varAiAttachmentsSaveFileRequest := _AiAttachmentsSaveFileRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiAttachmentsSaveFileRequest)

	if err != nil {
		return err
	}

	*o = AiAttachmentsSaveFileRequest(varAiAttachmentsSaveFileRequest)

	return err
}

type NullableAiAttachmentsSaveFileRequest struct {
	value *AiAttachmentsSaveFileRequest
	isSet bool
}

func (v NullableAiAttachmentsSaveFileRequest) Get() *AiAttachmentsSaveFileRequest {
	return v.value
}

func (v *NullableAiAttachmentsSaveFileRequest) Set(val *AiAttachmentsSaveFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAttachmentsSaveFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAttachmentsSaveFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAttachmentsSaveFileRequest(val *AiAttachmentsSaveFileRequest) *NullableAiAttachmentsSaveFileRequest {
	return &NullableAiAttachmentsSaveFileRequest{value: val, isSet: true}
}

func (v NullableAiAttachmentsSaveFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAttachmentsSaveFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

