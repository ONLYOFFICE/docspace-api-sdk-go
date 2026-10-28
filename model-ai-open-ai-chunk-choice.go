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

// checks if the AiOpenAIChunkChoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiOpenAIChunkChoice{}

// AiOpenAIChunkChoice One choice of a streaming completion, carrying the part this chunk adds.
type AiOpenAIChunkChoice struct {
	// The zero-based position of the choice. This service emits a single choice, so always 0.
	Index float32 `json:"index"`
	// What this chunk adds to the choice.
	Delta AiOpenAIChoiceDelta `json:"delta"`
	// Why the completion stopped, or null while it is still streaming.
	FinishReason NullableAiOpenAIFinishReason `json:"finish_reason"`
}

type _AiOpenAIChunkChoice AiOpenAIChunkChoice

// NewAiOpenAIChunkChoice instantiates a new AiOpenAIChunkChoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiOpenAIChunkChoice(index float32, delta AiOpenAIChoiceDelta, finishReason NullableAiOpenAIFinishReason) *AiOpenAIChunkChoice {
	this := AiOpenAIChunkChoice{}
	this.Index = index
	this.Delta = delta
	this.FinishReason = finishReason
	return &this
}

// NewAiOpenAIChunkChoiceWithDefaults instantiates a new AiOpenAIChunkChoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiOpenAIChunkChoiceWithDefaults() *AiOpenAIChunkChoice {
	this := AiOpenAIChunkChoice{}
	return &this
}

// GetIndex returns the Index field value
func (o *AiOpenAIChunkChoice) GetIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *AiOpenAIChunkChoice) GetIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *AiOpenAIChunkChoice) SetIndex(v float32) {
	o.Index = v
}

// GetDelta returns the Delta field value
func (o *AiOpenAIChunkChoice) GetDelta() AiOpenAIChoiceDelta {
	if o == nil {
		var ret AiOpenAIChoiceDelta
		return ret
	}

	return o.Delta
}

// GetDeltaOk returns a tuple with the Delta field value
// and a boolean to check if the value has been set.
func (o *AiOpenAIChunkChoice) GetDeltaOk() (*AiOpenAIChoiceDelta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delta, true
}

// SetDelta sets field value
func (o *AiOpenAIChunkChoice) SetDelta(v AiOpenAIChoiceDelta) {
	o.Delta = v
}

// GetFinishReason returns the FinishReason field value
// If the value is explicit nil, the zero value for AiOpenAIFinishReason will be returned
func (o *AiOpenAIChunkChoice) GetFinishReason() AiOpenAIFinishReason {
	if o == nil || o.FinishReason.Get() == nil {
		var ret AiOpenAIFinishReason
		return ret
	}

	return *o.FinishReason.Get()
}

// GetFinishReasonOk returns a tuple with the FinishReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiOpenAIChunkChoice) GetFinishReasonOk() (*AiOpenAIFinishReason, bool) {
	if o == nil {
		return nil, false
	}
	return o.FinishReason.Get(), o.FinishReason.IsSet()
}

// SetFinishReason sets field value
func (o *AiOpenAIChunkChoice) SetFinishReason(v AiOpenAIFinishReason) {
	o.FinishReason.Set(&v)
}

func (o AiOpenAIChunkChoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiOpenAIChunkChoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["delta"] = o.Delta
	toSerialize["finish_reason"] = o.FinishReason.Get()
	return toSerialize, nil
}

func (o *AiOpenAIChunkChoice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"index",
		"delta",
		"finish_reason",
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

	varAiOpenAIChunkChoice := _AiOpenAIChunkChoice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiOpenAIChunkChoice)

	if err != nil {
		return err
	}

	*o = AiOpenAIChunkChoice(varAiOpenAIChunkChoice)

	return err
}

type NullableAiOpenAIChunkChoice struct {
	value *AiOpenAIChunkChoice
	isSet bool
}

func (v NullableAiOpenAIChunkChoice) Get() *AiOpenAIChunkChoice {
	return v.value
}

func (v *NullableAiOpenAIChunkChoice) Set(val *AiOpenAIChunkChoice) {
	v.value = val
	v.isSet = true
}

func (v NullableAiOpenAIChunkChoice) IsSet() bool {
	return v.isSet
}

func (v *NullableAiOpenAIChunkChoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiOpenAIChunkChoice(val *AiOpenAIChunkChoice) *NullableAiOpenAIChunkChoice {
	return &NullableAiOpenAIChunkChoice{value: val, isSet: true}
}

func (v NullableAiOpenAIChunkChoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiOpenAIChunkChoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

