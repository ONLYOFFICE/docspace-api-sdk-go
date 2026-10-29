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

// checks if the AiMetaGenerateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiMetaGenerateRequest{}

// AiMetaGenerateRequest struct for AiMetaGenerateRequest
type AiMetaGenerateRequest struct {
	// The entry's table. Only `file` for now: a folder or a room answers 400.
	Kind string `json:"kind"`
	Id AiMetaGenerateRequestId `json:"id"`
	// What kind of metadata fields the caller is interested in; a hint for the model.
	Prompt *string `json:"prompt,omitempty"`
}

type _AiMetaGenerateRequest AiMetaGenerateRequest

// NewAiMetaGenerateRequest instantiates a new AiMetaGenerateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiMetaGenerateRequest(kind string, id AiMetaGenerateRequestId) *AiMetaGenerateRequest {
	this := AiMetaGenerateRequest{}
	this.Kind = kind
	this.Id = id
	return &this
}

// NewAiMetaGenerateRequestWithDefaults instantiates a new AiMetaGenerateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiMetaGenerateRequestWithDefaults() *AiMetaGenerateRequest {
	this := AiMetaGenerateRequest{}
	return &this
}

// GetKind returns the Kind field value
func (o *AiMetaGenerateRequest) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *AiMetaGenerateRequest) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *AiMetaGenerateRequest) SetKind(v string) {
	o.Kind = v
}

// GetId returns the Id field value
func (o *AiMetaGenerateRequest) GetId() AiMetaGenerateRequestId {
	if o == nil {
		var ret AiMetaGenerateRequestId
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiMetaGenerateRequest) GetIdOk() (*AiMetaGenerateRequestId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiMetaGenerateRequest) SetId(v AiMetaGenerateRequestId) {
	o.Id = v
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (o *AiMetaGenerateRequest) GetPrompt() string {
	if o == nil || IsNil(o.Prompt) {
		var ret string
		return ret
	}
	return *o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiMetaGenerateRequest) GetPromptOk() (*string, bool) {
	if o == nil || IsNil(o.Prompt) {
		return nil, false
	}
	return o.Prompt, true
}

// HasPrompt returns a boolean if a field has been set.
func (o *AiMetaGenerateRequest) IsPromptSet() bool {
	if o != nil && !IsNil(o.Prompt) {
		return true
	}

	return false
}

// SetPrompt gets a reference to the given string and assigns it to the Prompt field.
func (o *AiMetaGenerateRequest) SetPrompt(v string) {
	o.Prompt = &v
}

func (o AiMetaGenerateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiMetaGenerateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["id"] = o.Id
	if !IsNil(o.Prompt) {
		toSerialize["prompt"] = o.Prompt
	}
	return toSerialize, nil
}

func (o *AiMetaGenerateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kind",
		"id",
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

	varAiMetaGenerateRequest := _AiMetaGenerateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiMetaGenerateRequest)

	if err != nil {
		return err
	}

	*o = AiMetaGenerateRequest(varAiMetaGenerateRequest)

	return err
}

type NullableAiMetaGenerateRequest struct {
	value *AiMetaGenerateRequest
	isSet bool
}

func (v NullableAiMetaGenerateRequest) Get() *AiMetaGenerateRequest {
	return v.value
}

func (v *NullableAiMetaGenerateRequest) Set(val *AiMetaGenerateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiMetaGenerateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiMetaGenerateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiMetaGenerateRequest(val *AiMetaGenerateRequest) *NullableAiMetaGenerateRequest {
	return &NullableAiMetaGenerateRequest{value: val, isSet: true}
}

func (v NullableAiMetaGenerateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiMetaGenerateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

