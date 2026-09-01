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

// checks if the AiOpenAIChatCompletionChunk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiOpenAIChatCompletionChunk{}

// AiOpenAIChatCompletionChunk One `chat.completion.chunk` of an OpenAI-compatible streaming response. Only the fields this service can populate are emitted - an OpenAI client tolerates the rest as absent.
type AiOpenAIChatCompletionChunk struct {
	// The completion identifier, stable across every chunk of one response.
	Id string `json:"id"`
	// Always `chat.completion.chunk`.
	Object string `json:"object"`
	// When the completion started, in Unix seconds.
	Created float32 `json:"created"`
	// The model that produced the completion - the resolved profile's model.
	Model string `json:"model"`
	// The choices carried by this chunk. This service emits exactly one.
	Choices []AiOpenAIChunkChoice `json:"choices"`
}

type _AiOpenAIChatCompletionChunk AiOpenAIChatCompletionChunk

// NewAiOpenAIChatCompletionChunk instantiates a new AiOpenAIChatCompletionChunk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiOpenAIChatCompletionChunk(id string, object string, created float32, model string, choices []AiOpenAIChunkChoice) *AiOpenAIChatCompletionChunk {
	this := AiOpenAIChatCompletionChunk{}
	this.Id = id
	this.Object = object
	this.Created = created
	this.Model = model
	this.Choices = choices
	return &this
}

// NewAiOpenAIChatCompletionChunkWithDefaults instantiates a new AiOpenAIChatCompletionChunk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiOpenAIChatCompletionChunkWithDefaults() *AiOpenAIChatCompletionChunk {
	this := AiOpenAIChatCompletionChunk{}
	return &this
}

// GetId returns the Id field value
func (o *AiOpenAIChatCompletionChunk) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiOpenAIChatCompletionChunk) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AiOpenAIChatCompletionChunk) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *AiOpenAIChatCompletionChunk) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *AiOpenAIChatCompletionChunk) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *AiOpenAIChatCompletionChunk) SetObject(v string) {
	o.Object = v
}

// GetCreated returns the Created field value
func (o *AiOpenAIChatCompletionChunk) GetCreated() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *AiOpenAIChatCompletionChunk) GetCreatedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *AiOpenAIChatCompletionChunk) SetCreated(v float32) {
	o.Created = v
}

// GetModel returns the Model field value
func (o *AiOpenAIChatCompletionChunk) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *AiOpenAIChatCompletionChunk) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *AiOpenAIChatCompletionChunk) SetModel(v string) {
	o.Model = v
}

// GetChoices returns the Choices field value
func (o *AiOpenAIChatCompletionChunk) GetChoices() []AiOpenAIChunkChoice {
	if o == nil {
		var ret []AiOpenAIChunkChoice
		return ret
	}

	return o.Choices
}

// GetChoicesOk returns a tuple with the Choices field value
// and a boolean to check if the value has been set.
func (o *AiOpenAIChatCompletionChunk) GetChoicesOk() ([]AiOpenAIChunkChoice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Choices, true
}

// SetChoices sets field value
func (o *AiOpenAIChatCompletionChunk) SetChoices(v []AiOpenAIChunkChoice) {
	o.Choices = v
}

func (o AiOpenAIChatCompletionChunk) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiOpenAIChatCompletionChunk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["created"] = o.Created
	toSerialize["model"] = o.Model
	toSerialize["choices"] = o.Choices
	return toSerialize, nil
}

func (o *AiOpenAIChatCompletionChunk) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"created",
		"model",
		"choices",
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

	varAiOpenAIChatCompletionChunk := _AiOpenAIChatCompletionChunk{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiOpenAIChatCompletionChunk)

	if err != nil {
		return err
	}

	*o = AiOpenAIChatCompletionChunk(varAiOpenAIChatCompletionChunk)

	return err
}

type NullableAiOpenAIChatCompletionChunk struct {
	value *AiOpenAIChatCompletionChunk
	isSet bool
}

func (v NullableAiOpenAIChatCompletionChunk) Get() *AiOpenAIChatCompletionChunk {
	return v.value
}

func (v *NullableAiOpenAIChatCompletionChunk) Set(val *AiOpenAIChatCompletionChunk) {
	v.value = val
	v.isSet = true
}

func (v NullableAiOpenAIChatCompletionChunk) IsSet() bool {
	return v.isSet
}

func (v *NullableAiOpenAIChatCompletionChunk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiOpenAIChatCompletionChunk(val *AiOpenAIChatCompletionChunk) *NullableAiOpenAIChatCompletionChunk {
	return &NullableAiOpenAIChatCompletionChunk{value: val, isSet: true}
}

func (v NullableAiOpenAIChatCompletionChunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiOpenAIChatCompletionChunk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

